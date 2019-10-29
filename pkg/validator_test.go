package vstruct

import (
	"fmt"
	"testing"
)

type Hello struct {
	Name     string `json:"name" v:"alpha min(3)"`
}

func TestGetStructAttributes(t *testing.T) {
	js := `{
	"name": "aryan"
}`

	obj := new(Hello)
	if a := NewValidator(obj).BindJSON(js).Validate(); a.GetError() != nil {
		fmt.Println(a.GetError())
		fmt.Println(a.GetMessages())
		return
	}
	fmt.Println(obj)
}
