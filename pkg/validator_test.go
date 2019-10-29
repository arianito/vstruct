package vstruct

import (
	"fmt"
	"testing"
)

type Hello struct {
	Name     string `json:"name" v:"alpha_num"`
	Data     string `json:"data" v:"json"`
}

func TestGetStructAttributes(t *testing.T) {
	js := `{
	"name": "aryan alikhani 23",
	"data": "{\"hello\":123}"
}`
	LoadBuiltin()

	_ = SetLocale("fa")
	RegisterAttributes(map[string]string{
		"name": "نام",
	})

	obj := new(Hello)
	if a := NewValidator(obj).BindJSON(js).Validate(); a.GetError() != nil {
		fmt.Println(a.GetError())
		fmt.Println(a.GetMessages())
		return
	}
	fmt.Println(obj)
}
