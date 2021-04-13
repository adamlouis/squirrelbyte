package jsonlogicsqlite3

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document/jsonlogic"
)

var (
	sqlInfixByJSONLogicOperator = map[string]string{
		string(jsonlogic.OperatorAdd):                "+",
		string(jsonlogic.OperatorMultiply):           "*",
		string(jsonlogic.OperatorSubtract):           "-",
		string(jsonlogic.OperatorDivide):             "/",
		string(jsonlogic.OperatorEqual):              "==",
		string(jsonlogic.OperatorNotEqual):           "!=",
		string(jsonlogic.OperatorGreaterThan):        ">",
		string(jsonlogic.OperatorGreaterThanOrEqual): ">=",
		string(jsonlogic.OperatorLessThan):           "<",
		string(jsonlogic.OperatorLessThanOrEqual):    "<=",
	}

	sqlFunctionByJSONLogicOperator = map[string]string{
		string(jsonlogic.OperatorSum):      "SUM",
		string(jsonlogic.OperatorMax):      "MAX",
		string(jsonlogic.OperatorMin):      "MIN",
		string(jsonlogic.OperatorCount):    "COUNT",
		string(jsonlogic.OperatorAvg):      "AVG",
		string(jsonlogic.OperatorDistinct): "DISTINCT",
		string(jsonlogic.OperatorLike):     "LIKE",
		string(jsonlogic.OperatorGlob):     "GLOB",
		string(jsonlogic.OperatorRandom):   "RANDOM",
		string(jsonlogic.OperatorAbs):      "ABS",
		string(jsonlogic.OperatorRound):    "ROUND",
		string(jsonlogic.OperatorIif):      "IIF",
		string(jsonlogic.OperatorTypeOf):   "TYPEOF",
		string(jsonlogic.OperatorLower):    "LOWER",
		string(jsonlogic.OperatorUpper):    "UPPER",
		string(jsonlogic.OperatorSubstr):   "SUBSTR",
	}

	sqlConjunctionByJSONLogicOperator = map[string]string{
		string(jsonlogic.OperatorAnd): "AND",
		string(jsonlogic.OperatorOr):  "OR",
	}

	sqlOrderByDirectionByJSONLogicOperator = map[string]string{
		string(jsonlogic.OperatorOrderByAsc):  "ASC",
		string(jsonlogic.OperatorOrderByDesc): "DESC",
	}

	sqlAsRegexpLiteral = `^[a-zA-Z0-9_]+$`
	sqlAsNameRegexp    = regexp.MustCompile(sqlAsRegexpLiteral)
)

func NewSQLizer() jsonlogic.SQLizer {
	return &sqlizer{}
}

type sqlizer struct {
}

// build sql from json logic
// warning: this function uses raw string templates rather than prepared statements to build sql queries
//		    there could be bugs & sql injection opportunities. i've done my best. use carefully.
func (s *sqlizer) ToSQL(j interface{}) (string, error) {
	return eval(j)
}

func eval(j interface{}) (string, error) {
	if j == nil {
		return "", nil
	}

	if l, err := evalLiteral(j); err == nil {
		return l, nil
	}

	m, err := toMap(j)
	if err != nil {
		return "", err
	}

	if len(m) == 0 {
		return "", nil
	}

	operator, operands, err := getOps(m)
	if err != nil {
		return "", err
	}

	// handle var
	if operator == string(jsonlogic.OperatorVar) {
		if err := requireNOperands(operator, operands, 1); err != nil {
			return "", err
		}

		s, ok := operands[0].(string)
		if !ok {
			return "", fmt.Errorf("%s operator expects a string operand type. received %v", operator, operands)
		}

		return toSQLite3JSONPath(s), nil
	}

	// handle infix
	ifx, ok := sqlInfixByJSONLogicOperator[operator]
	if ok {
		if err := requireNOperands(operator, operands, 2); err != nil {
			return "", err
		}

		l, err := eval(operands[0])
		if err != nil {
			return "", err
		}

		r, err := eval(operands[1])
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("(%s)%s(%s)", l, ifx, r), nil
	}

	// handle fn
	fn, ok := sqlFunctionByJSONLogicOperator[operator]
	if ok {
		children := make([]string, len(operands))
		for i := range operands {
			c, err := eval(operands[i])
			if err != nil {
				return "", err
			}
			children[i] = c
		}

		return fmt.Sprintf("%s(%s)", fn, strings.Join(children, ",")), nil
	}

	// handle conj
	conj, ok := sqlConjunctionByJSONLogicOperator[operator]
	if ok {
		children := make([]string, len(operands))
		for i := range operands {
			c, err := eval(operands[i])
			if err != nil {
				return "", err
			}
			children[i] = fmt.Sprintf("(%s)", c)
		}
		return strings.Join(children, conj), nil
	}

	// handle exists
	if operator == string(jsonlogic.OperatorExists) || operator == string(jsonlogic.OperatorDoesNotExist) {
		if err := requireNOperands(operator, operands, 1); err != nil {
			return "", err
		}

		v, err := eval(operands[0])
		if err != nil {
			return "", err
		}

		neg := ""
		if operator == string(jsonlogic.OperatorDoesNotExist) {
			neg = " NOT"
		}

		return fmt.Sprintf("(%s) IS%s NULL", v, neg), nil
	}

	// handle order by
	if dir, ok := sqlOrderByDirectionByJSONLogicOperator[operator]; ok {
		if err := requireNOperands(operator, operands, 1); err != nil {
			return "", err
		}

		v, err := eval(operands[0])
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s %s", v, dir), nil
	}

	// handle not
	if operator == string(jsonlogic.OperatorNot) {
		if err := requireNOperands(operator, operands, 1); err != nil {
			return "", err
		}

		v, err := eval(operands[0])
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("NOT(%s)", v), nil
	}

	// handle as
	if operator == string(jsonlogic.OperatorAs) {
		if err := requireNOperands(operator, operands, 2); err != nil {
			return "", err
		}

		v, err := eval(operands[0])
		if err != nil {
			return "", err
		}

		as, ok := operands[1].(string)
		if !ok {
			return "", fmt.Errorf("%s operator expects a string operand type at index 1. received %v", operator, operands)
		}

		if !sqlAsNameRegexp.MatchString(as) {
			return "", fmt.Errorf("name provided at index 1 for operator `as` must match the form %s", sqlAsNameRegexp)
		}

		return fmt.Sprintf("(%s) AS %s", v, as), nil
	}

	return "", fmt.Errorf("unsupported operator `%s`", operator)
}

// get and sanitize the json literal
func evalLiteral(i interface{}) (string, error) {
	switch t := i.(type) {
	case string:
		return fmt.Sprintf("'%s'", strings.ReplaceAll(t, `'`, `''`)), nil
	case float64:
		b, err := json.Marshal(i)
		if err != nil {
			return "", err
		}
		return string(b), nil
	case bool:
		b, err := json.Marshal(i)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return "", fmt.Errorf("%v is not a json literal type", i)
}

func toMap(i interface{}) (map[string]interface{}, error) {
	m, ok := i.(map[string]interface{})
	if !ok || m == nil {
		return nil, fmt.Errorf("%v is not a string map type", i)
	}
	return m, nil
}

func requireNOperands(operator string, operands []interface{}, n int) error {
	if len(operands) != n {
		pl := ""
		if n > 1 {
			pl = "s"
		}
		return fmt.Errorf("operator `%s` expects exactly %d operand%s but received %d: %v", operator, n, pl, len(operands), operands)
	}
	return nil
}

// convert the dot notation path to a sqlite3 `json_extract` expression
func toSQLite3JSONPath(path string) string {
	idx := strings.Index(path, ".")

	if idx == -1 {
		return path
	}

	prefix := strings.ReplaceAll(path[:idx], `'`, `''`)
	if idx == len(path)-1 {
		return prefix
	}
	suffix := strings.ReplaceAll(path[idx+1:], `'`, `''`)
	return fmt.Sprintf(`json_extract(%s, '$.%s')`, prefix, suffix)
}

// get the jsonlogic operator & operands from the map
func getOps(m map[string]interface{}) (string, []interface{}, error) {
	if len(m) != 1 {
		return "", nil, fmt.Errorf("expected a object type with a single key, but received %v", m)
	}

	onlyKey := ""
	for k := range m {
		onlyKey = k
	}

	if l, ok := m[onlyKey].([]interface{}); ok {
		return onlyKey, l, nil
	}

	// support jsonlogic syntactic sugar for unary operators to skip the array around values
	return onlyKey, []interface{}{m[onlyKey]}, nil
}
