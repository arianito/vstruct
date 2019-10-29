# VSTRUCT
VSTRUCT is a library to validate your structs and objects in golang. <br>

#### INSTALLATION
```bash
go get -u github.com/xeuus/vstruct
```
### USAGE
Simple usage:
```go
package main

import (
	"fmt"
	"github.com/xeuus/vstruct/pkg"
)
type SomeRequest struct {
	 // vstruct comes with builtin validators (full documentation listed below)
	 // you can also register new ones
	Email string `json:"email" v:"required email string min(6)"`
	Password string `json:"password" v:"required string min(6)"`
	Age int `json:"age" v:"required min(18)"`
}

func main()  {
	// loads builtin validators and translations
	vstruct.LoadBuiltin()
	
	// let's create our objects
	obj := &SomeRequest{
		Email: "some@email.come",
		Password: "123456",
	}
	
	// validate your object using vstruct
	if a := vstruct.NewValidator(obj).Validate(); a.GetError() != nil {
		fmt.Println("error:", a.GetError())
		fmt.Println("messages:", a.GetMessages())
		return
	}
}
```
Usage with WEB frameworks (like gin):
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xeuus/vstruct/pkg"
	"net/http"
)
type SomeRequest struct {
	Email string `json:"email" v:"required email string min(6)"`
	Password string `json:"password" v:"required string min(6)"`
	Age int `json:"age" v:"required min(18)"`
}

func main()  {
	// loads builtin validators and translations
	vstruct.LoadBuiltin()
	
	r := gin.Default()
	
	r.Handle("POST", "/some", func(ctx *gin.Context) {
		obj := new(SomeRequest)
		if v := vstruct.NewValidator(obj).Bind(ctx.Bind).Validate(); v.GetError() != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":    v.GetError().Error(),
				"messages": v.GetMessages(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"succeed": true,
			"object":  obj,
		})
	})
}
```
### Builtin Validators
here is the list of builtin validators implemented in codebase.
all builtin validators supports both english and persian locale,
```go
vstruct.SetLocale("fa")
// or
vstruct.SetLocale("en")
```

+ **required** \
The field under validation must be present in the input data and not empty.
+ **accepted** \
The field under validation must be yes, on, 1, or true. This is useful for validating "Terms of Service" acceptance.
+ **filled** \
The field under validation must not be empty when it is present.
+ **alpha** \
The field under validation must be entirely alphabetic characters.
+ **alpha_num** \
The field under validation must be entirely alpha-numeric characters.
+ **alpha_dash** \
The field under validation may have alpha-numeric characters, as well as dashes and underscores.
+ **email** \
The field under validation must be formatted as an e-mail address.
+ **min(value)** \
The field under validation must be greater than or equal to a minimum value. Strings, numerics.
+ **max(value)** \
The field under validation must be less than or equal to a maximum value. Strings, numerics.
+ **size(value)** \
The field under validation must be equal to value. Strings.
+ **between(a, b)** \
The field under validation must have a size between the given min and max. Strings, numerics.
+ **json** \
The field under validation must be a valid JSON string.
+ **confirmed(anotherField)** \
The field under validation must have a matching field.
+ **regex(query)** \
The field under validation must match the given regular expression.
+ **not_regex(query)** \
The field under validation must not match the given regular expression.
+ **string** \
The field under validation must be a string.
+ **ends_with(phrases...)** \
The field under validation must ends with one of the given values.
+ **starts_with(phrases...)** \
The field under validation must start with one of the given values.
+ **in(keys...)** \
The field under validation must be included in the given list of values.
+ **not_in(keys...)** \
The field under validation must not be included in the given list of values.
### Custom Validation Functions
Register custom validation rules
```go
vstruct.RegisterRule("validation_name", vstruct.Combine(reflect.Bool, reflect.Int, ...), func(ctx *Context) string {
	if !ctx.FieldValue.Bool() {
		return vstruct.Translate("dic.name", Attribute(ctx.AliasName))
	}
	// empty string means no validation error found.
	return ""
})
```
extend current local messages
```go
vstruct.Append(map[string]interface{}{
	"hello": "hello, :name"
})
```
register attribute names in translation
```go
vstruct.RegisterAttributes(map[string]interface{}{
	"name": "Full Name",
	"username": "Username"
})
```
use attribute translations
```go
vstruct.Attribute("name")
```
register custom locale translation
```go
vstruct.RegisterLanguage("fr", map[string]interface{}{
	"string": "Le champ :attribute doit être une chaîne de caractères.",
})
```
use translations with arguments
```go
vstruct.Translate("string", "name")
```
