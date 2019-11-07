package vstruct

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ValidationError struct {
	Error    error             `json:"error"`
	Messages map[string]string `json:"messages"`
}

type Validator interface {
	Bind(obj interface{}) Validator
	BindFunc(binder func(obj interface{}) error) Validator
	BindJSON(json string) Validator
	Validate() Validator
	GetError() error
	GetMessages() map[string]string
}

func NewValidator(obj interface{}) Validator {
	vm := &validator{
		obj: obj,
		err: new(ValidationError),
	}
	return vm
}

type validator struct {
	obj interface{}
	err *ValidationError
}

func (v *validator) GetError() error {
	return v.err.Error
}
func (v *validator) GetMessages() map[string]string {
	return v.err.Messages
}


func (v *validator) Bind(obj interface{}) Validator {
	v.obj = obj
	return v
}

func (v *validator) BindJSON(js string) Validator {
	err := json.Unmarshal([]byte(js), v.obj)
	v.err.Error = err
	return v
}
func (v *validator) BindFunc(binder func(obj interface{}) error) Validator {
	err := binder(v.obj)
	v.err.Error = err
	return v
}

func (v *validator) Validate() Validator {
	if v.err.Error != nil {
		return v
	}
	v.err.Messages = make(map[string]string)
	tf := reflect.TypeOf(v.obj).Elem()
	vf := reflect.ValueOf(v.obj).Elem()
	if tf.Kind() != reflect.Struct {
		tf = tf.Elem()
		vf = vf.Elem()
	}

	ln := tf.NumField()
	var failed string
	for i := 0; i < ln; i++ {
		field := tf.Field(i)
		fieldName := getNameFromField(field)
		rule := field.Tag.Get("v")
		if rule != "" {
			var message string
			lexRule(rule, func(rule string, args ...string) bool {
				fn := FindRule(rule, field.Type.Kind())
				if fn != nil {
					message = fn.fn(&Context{
						Index:         i,
						Instance:      v.obj,
						InstanceType:  tf,
						InstanceValue: vf,
						Field:         field,
						FieldName:     field.Name,
						AliasName:     fieldName,
						FieldValue:    vf.Field(i),
						Args:          args,
					})
					return message == ""
				} else {
					failed = rule
					return false
				}
			})
			if failed != "" {
				break
			}
			if message != "" {
				v.err.Messages[fieldName] = message
			}
		}
	}
	if failed != "" {
		v.err.Error = fmt.Errorf("validation rule not found %s", failed)
		return v
	}
	if len(v.err.Messages) > 0 {
		v.err.Error = fmt.Errorf("validation failed")
		return v
	}
	return v
}
