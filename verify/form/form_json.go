package form

import (
	"github.com/astaxie/beego/validation"
	"fmt"
	"encoding/json"
)

type Err struct {
	Name    string
	Message string
}

type Field struct {
	Value  string
	Errors []Err
}

func NewInstant(Errors []*validation.Error, f map[string]string) []byte {
	var fields = make(map[string]Field)
	var F Field
	var ok bool
	for _, err := range Errors {
		if F, ok = fields[err.Key]; !ok {
			//not exists, add
			F = Field{Value:f[err.Key]}
		}
		F.Errors = append(F.Errors, Err{err.Key, err.Message})
		fields[err.Key] = F
	}
	b, err := json.Marshal(fields)
	if err != nil {
		fmt.Println("json err:", err) //todo err return
	}
	return b
}
