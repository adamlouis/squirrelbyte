package jsonlogicsqlite3

// import (
// 	"encoding/json"
// 	"testing"

// 	"github.com/adamlouis/squirrelbyte/server/internal/pkg/document/jsonlogic"
// 	"github.com/stretchr/testify/require"
// )

// func newTestConverter2() jsonlogic.Converter2 {
// 	return NewConverter2(
// 		map[string]struct{}{
// 			"id":         {},
// 			"header":     {},
// 			"body":       {},
// 			"created_at": {},
// 			"updated_at": {},
// 		},
// 		[]string{
// 			"body",
// 			"header",
// 		},
// 	)
// }

// func TestSelectToSQL2(t *testing.T) {
// 	tcs := []struct {
// 		name string
// 		in   string
// 		out  []string
// 		err  bool
// 	}{
// 		{
// 			name: "empty",
// 			in:   `[{"+":[1,2]}]`,
// 			out:  []string{`('1')+('2')`},
// 		},
// 		{
// 			name: "empty",
// 			in:   `[{"+":["hi",2]}]`,
// 			out:  []string{`('hi')+('2')`},
// 		},
// 		{
// 			name: "empty",
// 			in:   `[{"+":["hi","hey"]}]`,
// 			out:  []string{`('hi')+('hey')`},
// 		},
// 		{
// 			name: "empty",
// 			in:   `[{"avg":["hi","hey"]}]`,
// 			out:  []string{`AVG('hi','hey')`},
// 		},
// 		// 	{
// 		// 		name: "1",
// 		// 		in:   "[1]",
// 		// 		out:  []string{"1"},
// 		// 	},
// 		// 	{
// 		// 		name: "1",
// 		// 		in:   "[1.1]",
// 		// 		out:  []string{"1.1"},
// 		// 	},
// 		// 	{
// 		// 		name: "1",
// 		// 		in:   "[true]",
// 		// 		out:  []string{"true"},
// 		// 	},
// 		// 	{
// 		// 		name: "1",
// 		// 		in:   `["abc"]`,
// 		// 		out:  []string{"abc"},
// 		// 	},
// 		// 	{
// 		// 		name: "select column",
// 		// 		in:   `[{"var":"id"}]`,
// 		// 		out:  []string{"id"},
// 		// 	},
// 		// 	{
// 		// 		name: "select column",
// 		// 		in:   `[{"var":["id"]}]`,
// 		// 		out:  []string{"id"},
// 		// 	},
// 		// 	{
// 		// 		name: "select property",
// 		// 		in:   `[{"var":"body.name"}]`,
// 		// 		out:  []string{"json_extract(body, '$.name')"},
// 		// 	},
// 		// 	{
// 		// 		name: "select 2 properties",
// 		// 		in:   `[{"var":"body.name.first"}]`,
// 		// 		out:  []string{"json_extract(body, '$.name.first')"},
// 		// 	},
// 		// 	{
// 		// 		name: "select property",
// 		// 		in:   `[{"var":["body.name"]}]`,
// 		// 		out:  []string{"json_extract(body, '$.name')"},
// 		// 	},
// 		// 	{
// 		// 		name: "select 2 properties",
// 		// 		in:   `[{"var":["body.name.first"]}]`,
// 		// 		out:  []string{"json_extract(body, '$.name.first')"},
// 		// 	},
// 		// 	{
// 		// 		name: "sum of col",
// 		// 		in:   `[{"sum":{"var":"id"}}]`,
// 		// 		out:  []string{"SUM(id)"},
// 		// 	},
// 		// 	{
// 		// 		name: "count of col",
// 		// 		in:   `[{"count":{"var":"id"}}]`,
// 		// 		out:  []string{"COUNT(id)"},
// 		// 	},
// 		// 	{
// 		// 		name: "min of col",
// 		// 		in:   `[{"min":{"var":"id"}}]`,
// 		// 		out:  []string{"MIN(id)"},
// 		// 	},
// 		// 	{
// 		// 		name: "max of col",
// 		// 		in:   `[{"max":{"var":"id"}}]`,
// 		// 		out:  []string{"MAX(id)"},
// 		// 	},
// 		// 	{
// 		// 		name: "avg of col",
// 		// 		in:   `[{"avg":{"var":"id"}}]`,
// 		// 		out:  []string{"AVG(id)"},
// 		// 	},
// 		// 	{
// 		// 		name: "sum of col",
// 		// 		in:   `[{"sum":[{"var":"id"}]}]`,
// 		// 		out:  []string{"SUM(id)"},
// 		// 	},
// 		// 	{
// 		// 		name: "sum of col",
// 		// 		in:   `[{"distinct":{"sum":{"var":"id"}}}]`,
// 		// 		out:  []string{"DISTINCT(SUM(id))"},
// 		// 	},
// 		// 	{
// 		// 		name: "sum of col",
// 		// 		in:   `[{"distinct":{"sum":[{"var":"id"}]}}]`,
// 		// 		out:  []string{"DISTINCT(SUM(id))"},
// 		// 	},
// 		// 	{
// 		// 		name: "sum of col with property",
// 		// 		in:   `[{"sum":{"var":"body.name"}}]`,
// 		// 		out:  []string{"SUM(json_extract(body, '$.name'))"},
// 		// 	},
// 		// 	{
// 		// 		name: "sum of col with property",
// 		// 		in:   `[{"sum":[{"var":"body.name"}]}]`,
// 		// 		out:  []string{"SUM(json_extract(body, '$.name'))"},
// 		// 	},
// 	}

// 	cv := newTestConverter2()

// 	for i := range tcs {
// 		tc := tcs[i]
// 		t.Run(tc.name, func(t *testing.T) {
// 			ifc := []interface{}{}
// 			err := json.Unmarshal([]byte(tc.in), &ifc)
// 			require.Nil(t, err)

// 			res, err := cv.SelectToSQL(ifc)
// 			if tc.err {
// 				require.NotNil(t, err)
// 			} else {
// 				require.Nil(t, err)
// 				require.Equal(t, tc.out, res)
// 			}
// 		})
// 	}
// }

// func TestGroupByToSQL2(t *testing.T) {
// 	// tcs := []struct {
// 	// 	name string
// 	// 	in   string
// 	// 	out  []string
// 	// 	err  bool
// 	// }{
// 	// 	{
// 	// 		name: "empty",
// 	// 		in:   "[]",
// 	// 		out:  []string{},
// 	// 	},
// 	// 	{
// 	// 		name: "empty",
// 	// 		in:   "[1]",
// 	// 		out:  []string{"1"},
// 	// 	},
// 	// 	{
// 	// 		name: "empty",
// 	// 		in:   `["hi"]`,
// 	// 		out:  []string{"hi"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"var":"id"}]`,
// 	// 		out:  []string{"id"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"var":"body"}]`,
// 	// 		out:  []string{"body"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"var":["body"]}]`,
// 	// 		out:  []string{"body"},
// 	// 	},
// 	// 	{
// 	// 		name: "select property - empty",
// 	// 		in:   `[{"var":"body."}]`,
// 	// 		out:  []string{"json_extract(body, '$.')"},
// 	// 	},
// 	// 	{
// 	// 		name: "select property",
// 	// 		in:   `[{"var":"body.name"}]`,
// 	// 		out:  []string{"json_extract(body, '$.name')"},
// 	// 	},
// 	// 	{
// 	// 		name: "select 2 properties",
// 	// 		in:   `[{"var":"body.name.first"}]`,
// 	// 		out:  []string{"json_extract(body, '$.name.first')"},
// 	// 	},
// 	// }

// 	// cv := newTestConverter2()

// 	// for i := range tcs {
// 	// 	tc := tcs[i]
// 	// 	t.Run(tc.name, func(t *testing.T) {
// 	// 		ifc := []interface{}{}
// 	// 		err := json.Unmarshal([]byte(tc.in), &ifc)
// 	// 		require.Nil(t, err)

// 	// 		res, err := cv.GroupByToSQL(ifc)
// 	// 		if tc.err {
// 	// 			require.NotNil(t, err)
// 	// 		} else {
// 	// 			require.Nil(t, err)
// 	// 			require.Equal(t, tc.out, res)
// 	// 		}
// 	// 	})
// 	// }
// }

// func TestOrderByToSQL2(t *testing.T) {
// 	// tcs := []struct {
// 	// 	name string
// 	// 	in   string
// 	// 	out  []string
// 	// 	err  bool
// 	// }{
// 	// 	{
// 	// 		name: "empty",
// 	// 		in:   "[]",
// 	// 		out:  []string{},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"var":"id"}]`,
// 	// 		out:  []string{"id"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"var":"body"}]`,
// 	// 		out:  []string{"body"},
// 	// 	},
// 	// 	{
// 	// 		name: "property - empty",
// 	// 		in:   `[{"var":"body."}]`,
// 	// 		out:  []string{"json_extract(body, '$.')"},
// 	// 	},
// 	// 	{
// 	// 		name: "property",
// 	// 		in:   `[{"var":"body.name"}]`,
// 	// 		out:  []string{"json_extract(body, '$.name')"},
// 	// 	},
// 	// 	{
// 	// 		name: "2 properties",
// 	// 		in:   `[{"var":"body.name.first"}]`,
// 	// 		out:  []string{"json_extract(body, '$.name.first')"},
// 	// 	},
// 	// 	{
// 	// 		name: "sum of col",
// 	// 		in:   `[{"sum":{"var":"id"}}]`,
// 	// 		out:  []string{"SUM(id)"},
// 	// 	},
// 	// 	{
// 	// 		name: "sum of col with property",
// 	// 		in:   `[{"sum":{"var":"body.name"}}]`,
// 	// 		out:  []string{"SUM(json_extract(body, '$.name'))"},
// 	// 	},
// 	// 	{
// 	// 		name: "distinct sum of col",
// 	// 		in:   `[{"distinct":{"sum":{"var":"id"}}}]`,
// 	// 		out:  []string{"DISTINCT(SUM(id))"},
// 	// 	},
// 	// 	{
// 	// 		name: "sum of col distinct",
// 	// 		in:   `[{"sum":{"distinct":{"var":"id"}}}]`,
// 	// 		out:  []string{"SUM(DISTINCT(id))"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"desc":{"var":"id"}}]`,
// 	// 		out:  []string{"id DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"asc":{"var":"id"}}]`,
// 	// 		out:  []string{"id ASC"},
// 	// 	},
// 	// 	{
// 	// 		name: "column",
// 	// 		in:   `[{"desc":{"var":"body"}}]`,
// 	// 		out:  []string{"body DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "property - empty",
// 	// 		in:   `[{"desc":{"var":"body."}}]`,
// 	// 		out:  []string{"json_extract(body, '$.') DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "property",
// 	// 		in:   `[{"desc":{"var":"body.name"}}]`,
// 	// 		out:  []string{"json_extract(body, '$.name') DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "2 properties",
// 	// 		in:   `[{"desc":{"var":"body.name.first"}}]`,
// 	// 		out:  []string{"json_extract(body, '$.name.first') DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "sum of col",
// 	// 		in:   `[{"desc":{"sum":{"var":"id"}}}]`,
// 	// 		out:  []string{"SUM(id) DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "sum of col with property",
// 	// 		in:   `[{"desc":{"sum":{"var":"body.name"}}}]`,
// 	// 		out:  []string{"SUM(json_extract(body, '$.name')) DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "distinct sum of col",
// 	// 		in:   `[{"desc":{"distinct":{"sum":{"var":"id"}}}}]`,
// 	// 		out:  []string{"DISTINCT(SUM(id)) DESC"},
// 	// 	},
// 	// 	{
// 	// 		name: "sum of col distinct",
// 	// 		in:   `[{"desc":{"sum":{"distinct":{"var":"id"}}}}]`,
// 	// 		out:  []string{"SUM(DISTINCT(id)) DESC"},
// 	// 	},
// 	// }

// 	// cv := newTestConverter2()

// 	// for i := range tcs {
// 	// 	tc := tcs[i]
// 	// 	t.Run(tc.name, func(t *testing.T) {
// 	// 		ifc := []interface{}{}
// 	// 		err := json.Unmarshal([]byte(tc.in), &ifc)
// 	// 		require.Nil(t, err)

// 	// 		res, err := cv.OrderByToSQL(ifc)
// 	// 		if tc.err {
// 	// 			require.NotNil(t, err)
// 	// 		} else {
// 	// 			require.Nil(t, err)
// 	// 			require.Equal(t, tc.out, res)
// 	// 		}
// 	// 	})
// 	// }
// }

// func TestWhereToSQL2(t *testing.T) {
// 	// tcs := []struct {
// 	// 	name string
// 	// 	in   string
// 	// 	out  interface{}
// 	// 	err  bool
// 	// }{
// 	// 	// {
// 	// 	// 	name: "1",
// 	// 	// 	in:   "1",
// 	// 	// 	out:  sq.Expr("?", "1"),
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "==",
// 	// 	// 	in:   `{"==":[{"var": "id"},"hello"]}`,
// 	// 	// 	out:  sq.Eq{"id": "hello"},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "!=",
// 	// 	// 	// in:   []interface{}{"!=", "id", "world"},
// 	// 	// 	in:  `{"!=":["id", "world"]`,
// 	// 	// 	out: sq.NotEq{"id": "world"},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: ">",
// 	// 	// 	// in:   []interface{}{">", "id", 11},
// 	// 	// 	in:  `{">":["id", 11]`,
// 	// 	// 	out: sq.Gt{"id": 11},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: ">=",
// 	// 	// 	// in:   []interface{}{">=", "id", 11},
// 	// 	// 	in:  `{">=":["id", 11]`,
// 	// 	// 	out: sq.GtOrEq{"id": 11},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "<",
// 	// 	// 	// in:   []interface{}{"<", "id", 11},
// 	// 	// 	in:  `{"<":["id", 11]`,
// 	// 	// 	out: sq.Lt{"id": 11},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "<=",
// 	// 	// 	// in:   []interface{}{"<=", "id", 11},
// 	// 	// 	in:  `{"<=":["id", 11]`,
// 	// 	// 	out: sq.LtOrEq{"id": 11},
// 	// 	// },

// 	// 	// 		// {
// 	// 	// 		// 	name: "exists",
// 	// 	// 		// 	// in:   []interface{}{"exists", "id"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.NotEq{"id": nil},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "!exists",
// 	// 	// 		// 	// in:   []interface{}{"!exists", "id"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Eq{"id": nil},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "starts-with",
// 	// 	// 		// 	// in:   []interface{}{"starts-with", "id", "abc"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Like{"id": "abc%"},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "!starts-with",
// 	// 	// 		// 	// in:   []interface{}{"!starts-with", "id", "abc"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.NotLike{"id": "abc%"},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "contains",
// 	// 	// 		// 	// in:   []interface{}{"contains", "id", "abc"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Like{"id": "%abc%"},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "!contains",
// 	// 	// 		// 	// in:   []interface{}{"!contains", "id", "abc"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.NotLike{"id": "%abc%"},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "&&",
// 	// 	// 		// 	// in:   []interface{}{"&&", true, false},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.And{sq.Expr("?", true), sq.Expr("?", false)},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "||",
// 	// 	// 		// 	// in:   []interface{}{"||", true, false},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Or{sq.Expr("?", true), sq.Expr("?", false)},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "!",
// 	// 	// 		// 	// in:   []interface{}{"!", true},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Expr("NOT ?", true),
// 	// 	// 		// },
// 	// 	{
// 	// 		name: "&& - rec",
// 	// 		in:   `{"&&":[{"==":[{"var":"id"}, "hello"]}, {"==":[{"var":"id"}, "world"]}]}`,
// 	// 		out:  sq.And{sq.Eq{"id": "hello"}, sq.Eq{"id": "world"}},
// 	// 	},
// 	// 	{
// 	// 		name: "|| - rec",
// 	// 		in:   `{"||":[{"==":[{"var":"id"}, "hello"]}, {"==":[{"var":"id"}, "world"]}]}`,
// 	// 		out:  sq.Or{sq.Eq{"id": "hello"}, sq.Eq{"id": "world"}},
// 	// 	},
// 	// 	{
// 	// 		name: "! - rec",
// 	// 		in:   `{"!":[{"==":[{"var":"id"}, "hello"]}]}`,
// 	// 		out:  sq.Expr("NOT ?", sq.Eq{"id": "hello"}),
// 	// 	},
// 	// 	// {
// 	// 	// 	name: "&& - rec",
// 	// 	// 	in:   `{"&&":[{"==":[{"var":"id"}, "hello"]}, {"==":[{"var":"id"}, "world"]}]}`,
// 	// 	// 	out:  sq.And{sq.Eq{"id": "hello"}, sq.Eq{"id": "world"}},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "|| - rec",
// 	// 	// 	// in:   []interface{}{"||", []interface{}{"exists", "id"}, []interface{}{"<=", "id", 11}},
// 	// 	// 	in:  `{"||":[{"":[]}, {"":[]}]}`,
// 	// 	// 	out: sq.Or{sq.NotEq{"id": nil}, sq.LtOrEq{"id": 11}},
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "! - rec",
// 	// 	// 	// in:   []interface{}{"!", []interface{}{"==", "id", "hello"}},
// 	// 	// 	in:  ``,
// 	// 	// 	out: sq.Expr("NOT ?", sq.Eq{"id": "hello"}),
// 	// 	// },
// 	// 	// {
// 	// 	// 	name: "== - json path",
// 	// 	// 	in:   `{"==":["body.field.here.please", "hello"]}`,
// 	// 	// 	out:  sq.Eq{`json_extract(body, '$.field.here.please')`: "hello"},
// 	// 	// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "== - json path",
// 	// 	// 		// 	// in:   []interface{}{"==", "body.field[0][1].can[2].havearrays", "hello"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Eq{`json_extract(body, '$.field[0][1].can[2].havearrays')`: "hello"},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "no json path single quotes",
// 	// 	// 		// 	// in:   []interface{}{"==", `body.field'`, "hello"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	out: sq.Eq{`json_extract(body, '$.field''')`: "hello"},
// 	// 	// 		// },
// 	// 	// 		// {
// 	// 	// 		// 	name: "== - not string key",
// 	// 	// 		// 	// in:   []interface{}{"==", 7, "hello"},
// 	// 	// 		// 	in:  "",
// 	// 	// 		// 	err: true,
// 	// 	// 		// },
// 	// }

// 	// cv := newTestConverter2()

// 	// for i := range tcs {
// 	// 	tc := tcs[i]
// 	// 	t.Run(tc.name, func(t *testing.T) {
// 	// 		var err error
// 	// 		var res string

// 	// 		ifc := map[string]interface{}{}
// 	// 		err = json.Unmarshal([]byte(tc.in), &ifc)
// 	// 		if err == nil {
// 	// 			res, err = cv.WhereToSQL(ifc)
// 	// 		} else {
// 	// 			res, err = cv.WhereToSQL(tc.in)
// 	// 		}

// 	// 		if tc.err {
// 	// 			require.NotNil(t, err)
// 	// 		} else {
// 	// 			require.Nil(t, err)
// 	// 			require.Equal(t, tc.out, res)
// 	// 		}
// 	// 	})
// 	// }
// }
