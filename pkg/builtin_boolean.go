package vstruct

import (
	"reflect"
)

func loadBoolean() bool {
	RegisterRule("accepted", Combine(reflect.Bool), func(ctx *Context) string {
		if !ctx.FieldValue.Bool() {
			return translate("accepted", attribute(ctx.AliasName))
		}
		return ""
	})
	return true
}



//noinspection GoUnusedGlobalVariable
var local_boolean = loadBoolean()