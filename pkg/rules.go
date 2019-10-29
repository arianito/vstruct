package vstruct

import "reflect"

type ValidatorFunc func(fieldName string, value reflect.Value, args...string) string

type ruleObj struct {
	name string
	kind uint32
	fn ValidatorFunc
}

var rules []*ruleObj


func FindRule(name string, kind reflect.Kind) *ruleObj {
	shifted := uint32(kind) << 1
	for _, rule := range rules {
		if rule.name == name && (rule.kind& shifted) == shifted {
			return rule
		}
	}
	return nil
}

func RegisterRule(name string, kind uint32, fn ValidatorFunc)  {
	rules = append(rules, &ruleObj{
		name: name,
		kind: kind,
		fn:   fn,
	})
}

func Combine(kinds ...reflect.Kind) uint32 {
	var kind uint32 = 0
	for _, k := range kinds {
		kind |= uint32(k) << 1
	}
	return kind
}