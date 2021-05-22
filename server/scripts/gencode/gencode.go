// an opinionated code generator for generating boilerplate from yml
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
	Types   yaml.MapSlice `yaml:"types"`
	Objects yaml.MapSlice `yaml:"objects"`
	Routes  yaml.MapSlice `yaml:"routes"`
}

type Route struct {
	Path         string // from parent
	Method       string // from parent
	Name         string `yaml:"name"`
	QueryParams  string `yaml:"query_params"`
	RequestBody  string `yaml:"request_body"`
	ResponseBody string `yaml:"response_body"`
}

type Object struct {
	Name       string
	Properties []*Property
}

type Property struct {
	Name string
	Type string
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

		generateHandlerStubs(&c)
	case "client":
		code += generateHeader(&c, []string{
			`"context"`,
			`"bytes"`,
			`"encoding/json"`,
			`"fmt"`,
			`"io"`,
			`"io/ioutil"`,
			`"net/http"`,
			`"net/url"`,
			`"strconv"`,
			fmt.Sprintf("\"%s\"", flagModelPackage),
		})
		code += generateClientInterface(&c)
		code += generateClientImpl(&c)
	default:
		log.Fatalf("unsupported component %s", flagComponent)
	}

	if err = writeFormattedCode(code, flagOut); err != nil {
		panic(err)
	}
}

func (c *Config) forEachType(fn func(name, value string)) error {
	for _, item := range c.Types {
		k, ok := item.Key.(string)
		if !ok {
			return fmt.Errorf("invalid type name %s", item.Key)
		}
		v, ok := item.Value.(string)
		if !ok {
			return fmt.Errorf("invalid type value %s", item.Value)
		}
		fn(k, v)
	}
	return nil
}

func (c *Config) forEachObject(fn func(o Object)) error {
	for _, item := range c.Objects {
		k, ok := item.Key.(string)
		if !ok {
			return fmt.Errorf("invalid object name %s", item.Key)
		}

		b, err := yaml.Marshal(item.Value)
		if err != nil {
			return err
		}
		v := yaml.MapSlice{}
		yaml.Unmarshal(b, &v)

		ps := []*Property{}
		for _, vitem := range v {
			ps = append(ps, &Property{
				Name: vitem.Key.(string),
				Type: vitem.Value.(string),
			})
		}

		fn(Object{
			Name:       k,
			Properties: ps,
		})
	}
	return nil
}

func (c *Config) getObject(name string) *Object {
	var result *Object
	c.forEachObject(func(o Object) {
		if o.Name == name {
			result = &o
		}
	})
	return result
}

func (c *Config) forEachRoute(fn func(r Route)) error {
	for _, item := range c.Routes {
		path, ok := item.Key.(string)
		if !ok {
			return fmt.Errorf("invalid path name %s", item.Key)
		}

		b, err := yaml.Marshal(item.Value)
		if err != nil {
			return err
		}
		v := yaml.MapSlice{}
		yaml.Unmarshal(b, &v)

		for _, vitem := range v {
			method := vitem.Key.(string)
			b2, err := yaml.Marshal(vitem.Value)
			if err != nil {
				return err
			}
			route := Route{
				Path:   path,
				Method: method,
			}
			yaml.Unmarshal(b2, &route)
			fn(route)
		}
	}
	return nil
}

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

var pathParamRegexp = regexp.MustCompile("{([^}]+)}")

var headerTemplate = `
// GENERATED
// DO NOT EDIT
// GENERATOR: scripts/gencode/gencode.go
// ARGUMENTS: %s
package %s

%s
`

var headerImportTemplate = `
import (
%s
)
`

func generateHeader(c *Config, imports []string) string {
	importCode := ""
	if len(imports) > 0 {
		importCode = fmt.Sprintf(headerImportTemplate, strings.Join(imports, "\n"))
	}

	return fmt.Sprintf(
		headerTemplate,
		strings.Join(os.Args[1:], " "),
		flagPackageName,
		importCode,
	)
}

func generateModels(c *Config) string {
	code := ""

	c.forEachType(func(name, t string) {
		code += fmt.Sprintf("type %s %s\n", name, t)
	})

	c.forEachObject(func(o Object) {
		code += fmt.Sprintf("type %s struct {\n", o.Name)
		for _, p := range o.Properties {
			code += fmt.Sprintf("    %s %s `json:\"%s\"`\n", p.Name, p.Type, toSnakeCase(p.Name))
		}
		code += "}\n"
	})

	c.forEachRoute(func(r Route) {
		pathParams := pathParamRegexp.FindAllString(r.Path, -1)
		if len(pathParams) == 0 {
			return
		}
		code += fmt.Sprintf("type %sPathParams struct {\n", r.Name)
		for _, p := range pathParams {
			cleaned := cleanPathParam(p)
			varName := strings.ToUpper(cleaned[0:1]) + cleaned[1:]
			code += fmt.Sprintf("    %s string\n", varName)
		}
		code += "}\n"
	})

	return code
}

func (route *Route) getDefaultReturnValue() string {
	if route.ResponseBody == "" {
		return "http.StatusInternalServerError, fmt.Errorf(\"unimplemented\")"
	}
	return "nil, http.StatusInternalServerError, fmt.Errorf(\"unimplemented\")"
}

func (route *Route) getClientSig() string {
	sig := fmt.Sprintf("    %s(ctx context.Context", route.Name)
	pathParams := pathParamRegexp.FindAllString(route.Path, -1)
	if len(pathParams) > 0 {
		t := fmt.Sprintf("*%sPathParams", route.Name)
		sig += fmt.Sprintf(", pathParams %s", resolveModel(t))
	}
	if route.QueryParams != "" {
		sig += fmt.Sprintf(", queryParams %s", resolveModel(route.QueryParams))
	}
	if route.RequestBody != "" {
		sig += fmt.Sprintf(", body %s", resolveModel(route.RequestBody))
	}
	sig += ")"
	if route.ResponseBody == "" {
		sig += " (int, error)"
	} else {
		sig += fmt.Sprintf(" (%s, int, error)", resolveModel(route.ResponseBody))
	}
	return sig
}

func (route *Route) getServerSig() string {
	sig := fmt.Sprintf("    %s(ctx context.Context", route.Name)
	pathParams := pathParamRegexp.FindAllString(route.Path, -1)
	if len(pathParams) > 0 {
		t := fmt.Sprintf("*%sPathParams", route.Name)
		sig += fmt.Sprintf(", pathParams %s", resolveModel(t))
	}
	if route.QueryParams != "" {
		sig += fmt.Sprintf(", queryParams %s", resolveModel(route.QueryParams))
	}
	if route.RequestBody != "" {
		sig += fmt.Sprintf(", body %s", resolveModel(route.RequestBody))
	}
	sig += ")"
	if route.ResponseBody == "" {
		sig += " error"
	} else {
		sig += fmt.Sprintf(" (%s, error)", resolveModel(route.ResponseBody))
	}
	return sig
}

func generateServiceInterfaces(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}
	code := "type HTTPHandler interface {\n"
	c.forEachRoute(func(route Route) {
		code += fmt.Sprintf("    %s(w http.ResponseWriter, req *http.Request)\n", route.Name)
	})
	code += "}\n"

	code += "type APIHandler interface {\n"
	c.forEachRoute(func(route Route) {
		code += route.getServerSig() + "\n"
	})
	code += "}\n"

	return code
}

func generateClientInterface(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}

	code := "type Client interface {\n"
	c.forEachRoute(func(route Route) {
		code += route.getClientSig() + "\n"
	})
	code += "}\n"

	return code
}

var clientImplInitTemplate = `

func NewHTTPClient(baseURL string) Client {
	return &client{
		baseURL: baseURL,
	}
}

type client struct {
	baseURL string
}
`

func (r *Route) getClientDefaultReturn(b, c, e string) string {
	if r.ResponseBody == "" {
		return "return " + c + ", " + e
	}
	return "return " + b + ", " + c + ", " + e
}

func generateClientImpl(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}

	code := clientImplInitTemplate
	c.forEachRoute(func(route Route) {
		code += fmt.Sprintf("func (c *client) %s {\n", route.getClientSig())
		code += "client := &http.Client{}\n"

		pathParams := pathParamRegexp.FindAllString(route.Path, -1)
		pathVars := []string{}
		for _, p := range pathParams {
			v := cleanPathParam(p)
			v = "pathParams." + strings.ToUpper(v[0:1]) + v[1:]
			pathVars = append(pathVars, v)
		}
		path := pathParamRegexp.ReplaceAllString(route.Path, "%v")

		parseArgs := `fmt.Sprintf("%s` + path + `", c.baseURL`
		if len(pathVars) > 0 {
			parseArgs += ", " + strings.Join(pathVars, ", ")
		}
		parseArgs += `)`

		code += `u, err := url.Parse(` + parseArgs + `)
		if err != nil {
			` + route.getClientDefaultReturn("nil", "-1", "err") + `
		}
		`

		if route.QueryParams != "" {
			queryObj := c.getObject(strings.ReplaceAll(route.QueryParams, "*", ""))
			if queryObj != nil {
				for _, p := range queryObj.Properties {
					v := "queryParams." + p.Name
					if p.Type == "int" {
						v = "strconv.Itoa(queryParams." + p.Name + ")"
					}
					code += fmt.Sprintf("u.Query().Add(\"%s\", %s)\n", toSnakeCase(p.Name), v)
				}
			} else {
				log.Fatalf("no obj %s", route.QueryParams)
			}
		}
		code += "var requestBody io.Reader\n"
		if route.RequestBody != "" {
			code += `if jsonBytes, err := json.Marshal(body); err != nil {
				` + route.getClientDefaultReturn("nil", "-1", "err") + `
			} else {
				requestBody = bytes.NewBuffer(jsonBytes)
			}
			`
		}

		code += "req, err := http.NewRequest(" + golangMethodByMethod[route.Method] + ", u.String(), requestBody)\n"

		code += `req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			` + route.getClientDefaultReturn("nil", "-1", "err") + `
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			respBytes, _ := ioutil.ReadAll(resp.Body)
			` + route.getClientDefaultReturn("nil", "resp.StatusCode", `fmt.Errorf("[%d] %s", resp.StatusCode, string(respBytes))`) + `
		}
		`

		if route.ResponseBody != "" {
			code += fmt.Sprintf(`respBody := %s{}
			err = json.NewDecoder(resp.Body).Decode(&respBody)
			if err != nil {
				return nil, resp.StatusCode, err
			}
			return &respBody, resp.StatusCode, nil
			`, resolveModel(strings.ReplaceAll(route.ResponseBody, "*", "")))
		} else {
			code += "return resp.StatusCode, nil\n"
		}

		code += "}\n"
	})

	return code
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
	code := ""
	code += "func RegisterRouter(apiHandler APIHandler, r *mux.Router, c ErrorCoder) {\n"
	code += "    h := apiHandlerToHTTPHandler(apiHandler, c)\n"
	c.forEachRoute(func(route Route) {
		code += fmt.Sprintf("    r.Handle(\"%s\", http.HandlerFunc(h.%s)).Methods(%s)\n", route.Path, route.Name, golangMethodByMethod[route.Method])
	})
	code += "}\n"
	return code
}

var adapterTemplateInit = `
func apiHandlerToHTTPHandler(apiHandler APIHandler, errorCoder ErrorCoder) HTTPHandler {
	return &httpHandler{
		apiHandler: apiHandler,
		errorCoder: errorCoder,
	}
}

type httpHandler struct {
	apiHandler APIHandler
	errorCoder ErrorCoder
}

type ErrorCoder func(e error) int

// sendError sends an error response
func (h *httpHandler) sendError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(h.errorCoder(err))
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(&errorResponse{
		Message: err.Error(),
	})
}

func sendErrorWithCode(w http.ResponseWriter, code int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(&errorResponse{
		Message: err.Error(),
	})
}

// sendOK sends an success response
func sendOK(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	code := http.StatusOK
	if body == nil {
		code = http.StatusNoContent
	}
	w.WriteHeader(code)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.Encode(body)
}

type errorResponse struct {
	Message string %s
}
`

func generateInterfaceAdapter(c *Config) string {
	if len(c.Routes) == 0 {
		return ""
	}

	code := fmt.Sprintf(
		adapterTemplateInit,
		"`json:\"message\"`",
	)

	c.forEachRoute(func(route Route) {
		code += fmt.Sprintf("func (h *httpHandler) %s(w http.ResponseWriter, req *http.Request) {\n", route.Name)

		pathParams := pathParamRegexp.FindAllString(route.Path, -1)
		if len(pathParams) > 0 {
			code += "vars := mux.Vars(req)\n"
			for _, p := range pathParams {
				varName := cleanPathParam(p)
				code += fmt.Sprintf("%s, ok := vars[\"%s\"]\n", varName, varName)
				code += fmt.Sprintf(`if !ok {
							sendErrorWithCode(w, http.StatusBadRequest, fmt.Errorf("invalid %s path parameter"))
							return
						}`, varName)
				code += "\n"
			}
			code += fmt.Sprintf("pathParams := %sPathParams{\n", resolveModel(route.Name))
			for _, p := range pathParams {
				varName := cleanPathParam(p)
				pubvarName := strings.ToUpper(varName[0:1]) + varName[1:]
				code += fmt.Sprintf("%s: %s,\n", pubvarName, varName)
			}
			code += "}\n"
		}

		if route.QueryParams != "" {
			if !strings.HasPrefix(route.QueryParams, "*") {
				log.Fatalf("generator expects a pointer type for response body but received %s for %s", route.QueryParams, route.Name)
			}
			obj := c.getObject(route.QueryParams[1:])
			if obj == nil {
				log.Fatalf("object %s not found", route.QueryParams[1:])
			}
			for _, p := range obj.Properties {
				if p.Type != "string" && p.Type != "int" {
					log.Fatalf("query parameters must be of type string or int but received %s for %s.%s", p.Type, route.QueryParams, p.Name)
				}
				lowerVarName := strings.ToLower(p.Name[:1]) + p.Name[1:] + "QueryParam"
				if p.Type == "string" {
					code += fmt.Sprintf("    %s := req.URL.Query().Get(\"%s\")\n", lowerVarName, toSnakeCase(p.Name))
				} else {

					code += fmt.Sprintf("%s := 0\n", lowerVarName)
					code += fmt.Sprintf("if req.URL.Query().Get(\"%s\") != \"\" {\n", toSnakeCase(p.Name))
					code += fmt.Sprintf("q, err := strconv.Atoi(req.URL.Query().Get(\"%s\"))\n", toSnakeCase(p.Name))
					code += fmt.Sprintf(`if err != nil {
								sendErrorWithCode(w, http.StatusBadRequest, err)
								return
							}`)
					code += "\n"
					code += fmt.Sprintf("%s = q", lowerVarName)
					code += "}\n"

				}
			}
			code += fmt.Sprintf("    queryParams := %s{\n", resolveModel(route.QueryParams[1:]))
			for _, p := range obj.Properties {
				lowerVarName := strings.ToLower(p.Name[:1]) + p.Name[1:] + "QueryParam"
				code += fmt.Sprintf("    %s: %s,\n", p.Name, lowerVarName)
			}
			code += fmt.Sprintf("    }\n")

		}

		if route.RequestBody != "" {
			if !strings.HasPrefix(route.RequestBody, "*") {
				log.Fatalf("generator expects a pointer type for response body but received %s for %s", route.RequestBody, route.Name)
			}
			code += fmt.Sprintf("    var requestBody %s\n", resolveModel(route.RequestBody[1:]))
			code += fmt.Sprintf(`    if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
	        sendErrorWithCode(w, http.StatusBadRequest, err)
	        return
	    }`)
			code += "\n"
		}

		call := ""
		if route.ResponseBody != "" {
			call += "r, "
		}

		call += "err"
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
		code += call + "\n"
		code += `    if err != nil {
	        h.sendError(w, err)
	    	return
	    }`

		code += "\n"
		if route.ResponseBody != "" {
			code += "sendOK(w, r)\n"
		} else {
			code += "sendOK(w, struct{}{})\n"
		}

		code += "}\n"
	})

	return code
}

var handlerTemplate = `
package %s

func NewAPIHandler() APIHandler {
	return &hdl{}
}

type hdl struct{}
`

var handlerFnTemplate = `
package %s

import (
	"context"
	"fmt"
	"net/http"

	"%s"
)

func (h *hdl) %s {
	return %s
}
`

func generateHandlerStubs(c *Config) {
	handlerCode := fmt.Sprintf(handlerTemplate, flagPackageName)
	handlerPath := filepath.Join(flagOutDir, "handler.go")
	if err := writeFormattedCodeIfDNE(handlerCode, handlerPath); err != nil {
		panic(err)
	}

	c.forEachRoute(func(r Route) {
		handlerFnPath := filepath.Join(flagOutDir, "handler_"+toSnakeCase(r.Name)+".go")
		handlerFnCode := fmt.Sprintf(
			handlerFnTemplate,
			flagPackageName,
			flagModelPackage,
			r.getServerSig(),
			r.getDefaultReturnValue(),
		)
		if err := writeFormattedCodeIfDNE(handlerFnCode, handlerFnPath); err != nil {
			panic(err)
		}
	})
}

func lastComponent(s string) string {
	parts := strings.Split(s, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func resolveModel(s string) string {
	if strings.HasPrefix(s, "*") {
		return strings.Replace(s, "*", "*"+lastComponent(flagModelPackage)+".", 1)
	}
	return lastComponent(flagModelPackage) + "." + s
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
