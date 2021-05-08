package main

import (
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Types   map[string]string            `yaml:"types"`
	Objects map[string]map[string]string `yaml:"objects"`
	Routes  map[string]map[string]Route  `yaml:"routes"`
}

type Route struct {
	Name         string `yaml:"name"`
	QueryParams  string `yaml:"query_params"`
	RequestBody  string `yaml:"request_body"`
	ResponseBody string `yaml:"response_body"`
}

var (
	flagFilename     = ""
	flagPackageName  = ""
	flagOut          = ""
	flagOutDir       = ""
	flagComponent    = ""
	flagModelPackage = ""
)

func main() {

	flag.StringVar(&flagFilename, "config", "", "the name of the yml config file with the api declaration")
	flag.StringVar(&flagPackageName, "package", "", "the name of the package containing the generated code")
	flag.StringVar(&flagOut, "out", "", "the name of the file to write output to")
	flag.StringVar(&flagOutDir, "out-dir", "", "the directory to write the output to")
	flag.StringVar(&flagComponent, "component", "", "the name of the component to generate - model, server")
	flag.StringVar(&flagModelPackage, "model-package", "", "the name of the model package")
	flag.Parse()

	c := Config{}

	filename, err := filepath.Abs(flagFilename)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := yaml.NewDecoder(file)
	dec.SetStrict(true)

	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}

	code := ""
	switch flagComponent {
	case "model":
		code += generateHeader(&c, []string{})
		code += generateModels(&c)
	case "server":
		code += generateHeader(&c, []string{
			`"context"`,
			`"encoding/json"`,
			`"fmt"`,
			`"net/http"`,
			`"strconv"`,
			`"github.com/gorilla/mux"`,
			fmt.Sprintf("\"%s\"", flagModelPackage),
		})
		code += generateServiceInterfaces(&c)
		code += generateRouter(&c)
		code += generateInterfaceAdapter(&c)

		if err := generateStubs(&c); err != nil {
			panic(err)
		}
	default:
		log.Fatalf("unsupported component %s", flagComponent)
	}

	if err = writeFormattedCode(code, flagOut); err != nil {
		panic(err)
	}
}

var pathParamRegexp = regexp.MustCompile("{([^}]+)}")

// TODO: use golang templates
var handlerTemplate = `
package %s

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
`

// TODO: use golang templates
var handlerFnTemplate = `
package %s

import (
	"context"
	"fmt"
	"net/http"

	"%s"
)

func (h *hdl) %s(%s) %s {
	return %s
}
`

func writeFormattedCode(code, path string) error {
	fmt.Printf("writing %s\n", path)
	fmt, err := format.Source([]byte(code))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(fmt), 0644)
}

func writeFormattedCodeIfDNE(code, path string) error {
	ex, err := exists(path)
	if err != nil {
		return err
	}

	if !ex {
		return writeFormattedCode(code, path)
	}

	fmt.Printf("exists %s\n", path)
	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func generateStubs(c *Config) error {
	handlerCode := fmt.Sprintf(handlerTemplate, flagPackageName)
	handlerPath := filepath.Join(flagOutDir, "handler.go")
	if err := writeFormattedCodeIfDNE(handlerCode, handlerPath); err != nil {
		return err
	}
	for path, methods := range c.Routes {
		for _, route := range methods {
			handlerFnPath := filepath.Join(flagOutDir, toSnakeCase(route.Name)+".go")
			handlerFnCode := fmt.Sprintf(
				handlerFnTemplate,
				flagPackageName,
				flagModelPackage,
				route.Name,
				getArgs(path, route),
				getReturnType(route),
				getDefaultReturnValue(route),
			)
			if err := writeFormattedCodeIfDNE(handlerFnCode, handlerFnPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func generateHeader(c *Config, imports []string) string {
	r := "// GENERATED\n"
	r += "// DO NOT EDIT\n"
	r += "// GENERATOR: scripts/gencode/gencode.go\n"
	r += fmt.Sprintf("// ARGUMENTS: '%s'\n", strings.Join(os.Args[1:], " "))
	r += fmt.Sprintf("\npackage %s\n\n", flagPackageName)

	if len(c.Routes) == 0 {
		return r
	}

	if len(imports) > 0 {
		r += "import (\n"
		r += strings.Join(imports, "\n")
		r += "\n)\n"
	}

	return r
}

func generateModels(c *Config) string {
	r := ""
	// write any explict type declaration
	for name, def := range c.Types {
		r += fmt.Sprintf("type %s %s\n", name, def)
	}
	// write objects
	for name, properties := range c.Objects {
		r += fmt.Sprintf("type %s struct {\n", name)
		for pname, typ := range properties {
			r += fmt.Sprintf("    %s %s `json:\"%s\"`\n", pname, typ, toSnakeCase(pname))
		}
		r += "}\n"
	}
	// write path param objects
	for path, methods := range c.Routes {
		pathParams := pathParamRegexp.FindAllString(path, -1)
		if len(pathParams) == 0 {
			continue
		}
		for _, route := range methods {
			r += fmt.Sprintf("type %sPathParams struct {\n", route.Name)
			for _, p := range pathParams {
				varName := strings.ReplaceAll(p, "}", "")
				varName = strings.ReplaceAll(varName, "{", "")
				varName = strings.ToUpper(varName[0:1]) + varName[1:]
				r += fmt.Sprintf("    %s string\n", varName)
			}
			r += "}\n"
		}
	}
	return r
}

func lastComponent(s string) string {
	parts := strings.Split(s, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func resolve(s string) string {
	if strings.HasPrefix(s, "*") {
		return strings.Replace(s, "*", "*"+lastComponent(flagModelPackage)+".", 1)
	}
	return lastComponent(flagModelPackage) + "." + s
}

func getArgs(path string, route Route) string {
	sig := "ctx context.Context"
	pathParams := pathParamRegexp.FindAllString(path, -1)
	if len(pathParams) > 0 {
		t := fmt.Sprintf("*%sPathParams", route.Name)
		sig += fmt.Sprintf(", pathParams %s", resolve(t))
	}
	if route.QueryParams != "" {
		sig += fmt.Sprintf(", queryParams %s", resolve(route.QueryParams))
	}
	if route.RequestBody != "" {
		sig += fmt.Sprintf(", body %s", resolve(route.RequestBody))
	}
	return sig
}

func getReturnType(route Route) string {
	if route.ResponseBody == "" {
		return "(int, error)"
	}
	return fmt.Sprintf("(%s, int, error)", resolve(route.ResponseBody))
}

func getDefaultReturnValue(route Route) string {
	if route.ResponseBody == "" {
		return "http.StatusInternalServerError, fmt.Errorf(\"unimplemented\")"
	}
	return "nil, http.StatusInternalServerError, fmt.Errorf(\"unimplemented\")"
}

func generateServiceInterfaces(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}
	r := "type HTTPHandler interface {\n"
	for _, methods := range c.Routes {
		for _, route := range methods {
			r += fmt.Sprintf("    %s(w http.ResponseWriter, req *http.Request)\n", route.Name)
		}
	}
	r += "}\n"
	r += "type APIHandler interface {\n"

	for path, methods := range c.Routes {
		for _, route := range methods {
			sig := fmt.Sprintf("    %s(ctx context.Context", route.Name)
			pathParams := pathParamRegexp.FindAllString(path, -1)
			if len(pathParams) > 0 {
				t := fmt.Sprintf("*%sPathParams", route.Name)
				sig += fmt.Sprintf(", pathParams %s", resolve(t))
			}
			if route.QueryParams != "" {
				sig += fmt.Sprintf(", queryParams %s", resolve(route.QueryParams))
			}
			if route.RequestBody != "" {
				sig += fmt.Sprintf(", body %s", resolve(route.RequestBody))
			}
			sig += ")"
			if route.ResponseBody == "" {
				sig += " (int, error)"
			} else {
				sig += fmt.Sprintf(" (%s, int, error)", resolve(route.ResponseBody))
			}
			sig += "\n"
			r += sig
		}
	}
	r += "}\n"
	return r
}

func generateInterfaceAdapter(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}
	r := `func apiHandlerToHTTPHandler(apiHandler APIHandler) HTTPHandler {
			return &httpHandler{
				apiHandler: apiHandler,
			}
		}

		type httpHandler struct {
			apiHandler APIHandler
		}



		// sendError sends an error response
func sendError(w http.ResponseWriter, code int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(&errorResponse{
		Message: err.Error(),
	})
}

// sendOK sends an success response
func sendOK(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(body)
}
`

	r += "type errorResponse struct {\n"
	r += "    Message string `json:\"message\"`\n"
	r += "}\n"

	for path, methods := range c.Routes {
		for _, route := range methods {
			sig := fmt.Sprintf("%s(w http.ResponseWriter, req *http.Request)", route.Name)

			r += "func (h *httpHandler) " + sig + "{\n"

			pathParams := pathParamRegexp.FindAllString(path, -1)
			if len(pathParams) > 0 {
				r += "vars := mux.Vars(req)\n"
				for _, p := range pathParams {
					varName := cleanPathParam(p)
					r += fmt.Sprintf("%s, ok := vars[\"%s\"]\n", varName, varName)
					r += fmt.Sprintf(`if !ok {
						sendError(w, http.StatusInternalServerError, fmt.Errorf("invalid %s path parameter"))
						return
					}`, varName)
					r += "\n"
				}
				r += fmt.Sprintf("pathParams := %sPathParams{\n", resolve(route.Name))
				for _, p := range pathParams {
					varName := cleanPathParam(p)
					pubvarName := strings.ToUpper(varName[0:1]) + varName[1:]
					r += fmt.Sprintf("%s: %s,\n", pubvarName, varName)
				}
				r += "}\n"
			}

			if route.QueryParams != "" {
				if !strings.HasPrefix(route.QueryParams, "*") {
					log.Fatalf("generator expects a pointer type for response body but received %s for %s", route.QueryParams, route.Name)
				}
				for propertyName, typ := range c.Objects[route.QueryParams[1:]] {
					if typ != "string" && typ != "int" {
						log.Fatalf("query parameters must be of type string or int but received %s for %s.%s", typ, route.QueryParams, propertyName)
					}
					lowerVarName := strings.ToLower(propertyName[:1]) + propertyName[1:] + "QueryParam"
					if typ == "string" {
						r += fmt.Sprintf("    %s := req.URL.Query().Get(\"%s\")\n", lowerVarName, toSnakeCase(propertyName))
					} else {
						r += fmt.Sprintf("    %s, err := strconv.Atoi(req.URL.Query().Get(\"%s\"))\n", lowerVarName, toSnakeCase(propertyName))
						r += fmt.Sprintf(`if err != nil {
							sendError(w, http.StatusBadRequest, err)
							return
						}`)
						r += "\n"

					}
				}
				r += fmt.Sprintf("    queryParams := %s{\n", resolve(route.QueryParams[1:]))
				for propertyName := range c.Objects[route.QueryParams[1:]] {
					lowerVarName := strings.ToLower(propertyName[:1]) + propertyName[1:] + "QueryParam"
					r += fmt.Sprintf("    %s: %s,\n", propertyName, lowerVarName)
				}
				r += fmt.Sprintf("    }\n")

			}

			if route.RequestBody != "" {
				if !strings.HasPrefix(route.RequestBody, "*") {
					log.Fatalf("generator expects a pointer type for response body but received %s for %s", route.RequestBody, route.Name)
				}
				r += fmt.Sprintf("    var requestBody %s\n", resolve(route.RequestBody[1:]))
				r += fmt.Sprintf(`    if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
        sendError(w, http.StatusBadRequest, err)
        return
    }`)
				r += "\n"
			}

			call := ""
			if route.ResponseBody != "" {
				call += "r, "
			}

			call += "code, err"
			call += " := "
			call += fmt.Sprintf("h.apiHandler.%s(req.Context()", route.Name)

			if len(pathParams) > 0 {
				call += ", &pathParams"
			}
			if route.QueryParams != "" {
				call += ", &queryParams"
			}
			if route.RequestBody != "" {
				call += ", &requestBody"
			}

			call += ")"
			r += call + "\n"
			r += `    if err != nil {
        sendError(w, code, err)
    	return
    }`

			r += "\n"
			if route.ResponseBody != "" {
				r += "sendOK(w, code, r)\n"
			} else {
				r += "sendOK(w, code, struct{}{})\n"
			}

			r += "}\n"
		}
	}

	return r
}

var golangMethodByMethod = map[string]string{
	"GET":    "http.MethodGet",
	"POST":   "http.MethodPost",
	"PUT":    "http.MethodPut",
	"DELETE": "http.MethodDelete",
}

func generateRouter(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}
	r := ""
	r += "func RegisterRouter(apiHandler APIHandler, r *mux.Router) {\n"
	r += "    h := apiHandlerToHTTPHandler(apiHandler)\n"
	for path, methods := range c.Routes {
		for method, route := range methods {
			r += fmt.Sprintf("    r.Handle(\"%s\", http.HandlerFunc(h.%s)).Methods(%s)\n", path, route.Name, golangMethodByMethod[method])
		}
	}
	r += "}\n"
	return r
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func cleanPathParam(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "}", ""), "{", "")
}
