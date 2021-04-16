package jsonlogic

// SQLizer converts the provided JSON logic query to a sql expression
type SQLizer interface {
	ToSQL(j interface{}) (string, error)
}

// AllToSQL converts all provided JSON logic queries to sql expressions
func AllToSQL(s SQLizer, j []interface{}) ([]string, error) {
	r := make([]string, len(j))
	for i := range j {
		v, err := s.ToSQL(j[i])
		if err != nil {
			return nil, err
		}
		r[i] = v
	}
	return r, nil
}

// Operator is a jsonlogic query operator
type Operator string

//nolint:golint
const (
	// binary infix
	OperatorAdd                Operator = "+"
	OperatorMultiply           Operator = "*"
	OperatorSubtract           Operator = "-"
	OperatorDivide             Operator = "/"
	OperatorEqual              Operator = "=="
	OperatorNotEqual           Operator = "!="
	OperatorGreaterThan        Operator = ">"
	OperatorGreaterThanOrEqual Operator = ">="
	OperatorLessThan           Operator = "<"
	OperatorLessThanOrEqual    Operator = "<="

	// fn
	// https://sqlite.org/lang_corefunc.html
	OperatorSum      Operator = "sum"
	OperatorMax      Operator = "max"
	OperatorMin      Operator = "min"
	OperatorCount    Operator = "count"
	OperatorAvg      Operator = "avg"
	OperatorDistinct Operator = "distinct"
	OperatorLike     Operator = "like"
	OperatorGlob     Operator = "glob"
	OperatorRandom   Operator = "random"
	OperatorAbs      Operator = "abs"
	OperatorRound    Operator = "round"
	OperatorIif      Operator = "iif"
	OperatorTypeOf   Operator = "typeof"
	OperatorLower    Operator = "lower"
	OperatorUpper    Operator = "upper"
	OperatorSubstr   Operator = "substr"

	// conj
	OperatorAnd Operator = "and"
	OperatorOr  Operator = "or"

	// other
	OperatorVar          Operator = "var"
	OperatorOrderByAsc   Operator = "asc"
	OperatorOrderByDesc  Operator = "desc"
	OperatorExists       Operator = "exists"
	OperatorDoesNotExist Operator = "!exists"
	OperatorNot          Operator = "not"
	OperatorAs           Operator = "as"
)
