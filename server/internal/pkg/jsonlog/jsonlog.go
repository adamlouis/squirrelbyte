package jsonlog

import (
	"encoding/json"
	"fmt"
)

// LogI attempts to convert the provided interface to json & prints it if it is valid json
func LogI(i interface{}) {
	if b, err := json.Marshal(i); err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(err)
	}
}

// Log wraps the provided values in a json object & prints it if it is valid json
func Log(kv ...interface{}) {
	m := map[string]interface{}{}
	i := 0
	for {
		if i+1 >= len(kv) {
			break
		}
		m[fmt.Sprintf("%v", kv[i])] = kv[i+1]
		i += 2
	}
	LogI(m)
}
