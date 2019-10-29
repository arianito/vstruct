package vstruct

import (
	"fmt"
	"strings"
)

var lang = "en"

var messages = map[string]map[string]interface{}{
	"fa": {
		"accepted":        ":attribute باید پذیرفته شده باشد.",
		"active_url":      "آدرس :attribute معتبر نیست.",
		"after":           ":attribute باید تاریخی بعد از :date باشد.",
		"after_or_equal":  ":attribute باید تاریخی بعد از :date، یا مطابق با آن باشد.",
		"alpha":           ":attribute باید فقط حروف الفبا باشد.",
		"alpha_dash":      ":attribute باید فقط حروف الفبا، اعداد، خط تیره و زیرخط باشد.",
		"alpha_num":       ":attribute باید فقط حروف الفبا و اعداد باشد.",
		"array":           ":attribute باید آرایه باشد.",
		"before":          ":attribute باید تاریخی قبل از :date باشد.",
		"before_or_equal": ":attribute باید تاریخی قبل از :date، یا مطابق با آن باشد.",
		"between": map[string]interface{}{
			"numeric": ":attribute باید بین :min و :max باشد.",
			"file":    ":attribute باید بین :min و :max کیلوبایت باشد.",
			"string":  ":attribute باید بین :min و :max کاراکتر باشد.",
			"array":   ":attribute باید بین :min و :max آیتم باشد.",
		},
		"boolean":        "فیلد :attribute فقط می‌تواند true و یا false باشد.",
		"confirmed":      ":attribute با فیلد تکرار مطابقت ندارد.",
		"date":           ":attribute یک تاریخ معتبر نیست.",
		"date_equals":    ":attribute باید یک تاریخ برابر با تاریخ :date باشد.",
		"date_format":    ":attribute با الگوی :format مطابقت ندارد.",
		"different":      ":attribute و :other باید از یکدیگر متفاوت باشند.",
		"digits":         ":attribute باید :digits رقم باشد.",
		"digits_between": ":attribute باید بین :min و :max رقم باشد.",
		"dimensions":     "ابعاد تصویر :attribute قابل قبول نیست.",
		"distinct":       "فیلد :attribute مقدار تکراری دارد.",
		"email":          ":attribute باید یک ایمیل معتبر باشد.",
		"ends_with":      "فیلد :attribute باید با یکی از مقادیر زیر خاتمه یابد: :values",
		"exists":         ":attribute انتخاب شده، معتبر نیست.",
		"file":           ":attribute باید یک فایل معتبر باشد.",
		"filled":         "فیلد :attribute باید مقدار داشته باشد.",
		"gt": map[string]interface{}{
			"numeric": ":attribute باید بزرگتر از :value باشد.",
			"file":    ":attribute باید بزرگتر از :value کیلوبایت باشد.",
			"string":  ":attribute باید بیشتر از :value کاراکتر داشته باشد.",
			"array":   ":attribute باید بیشتر از :value آیتم داشته باشد.",
		},
		"gte": map[string]interface{}{
			"numeric": ":attribute باید بزرگتر یا مساوی :value باشد.",
			"file":    ":attribute باید بزرگتر یا مساوی :value کیلوبایت باشد.",
			"string":  ":attribute باید بیشتر یا مساوی :value کاراکتر داشته باشد.",
			"array":   ":attribute باید بیشتر یا مساوی :value آیتم داشته باشد.",
		},
		"image":    ":attribute باید یک تصویر معتبر باشد.",
		"in":       ":attribute انتخاب شده، معتبر نیست.",
		"in_array": "فیلد :attribute در لیست :other وجود ندارد.",
		"integer":  ":attribute باید عدد صحیح باشد.",
		"ip":       ":attribute باید آدرس IP معتبر باشد.",
		"ipv4":     ":attribute باید یک آدرس معتبر از نوع IPv4 باشد.",
		"ipv6":     ":attribute باید یک آدرس معتبر از نوع IPv6 باشد.",
		"json":     "فیلد :attribute باید یک رشته از نوع JSON باشد.",
		"lt": map[string]interface{}{
			"numeric": ":attribute باید کوچکتر از :value باشد.",
			"file":    ":attribute باید کوچکتر از :value کیلوبایت باشد.",
			"string":  ":attribute باید کمتر از :value کاراکتر داشته باشد.",
			"array":   ":attribute باید کمتر از :value آیتم داشته باشد.",
		},
		"lte": map[string]interface{}{
			"numeric": ":attribute باید کوچکتر یا مساوی :value باشد.",
			"file":    ":attribute باید کوچکتر یا مساوی :value کیلوبایت باشد.",
			"string":  ":attribute باید کمتر یا مساوی :value کاراکتر داشته باشد.",
			"array":   ":attribute باید کمتر یا مساوی :value آیتم داشته باشد.",
		},
		"max": map[string]interface{}{
			"numeric": ":attribute نباید بزرگتر از :max باشد.",
			"file":    ":attribute نباید بزرگتر از :max کیلوبایت باشد.",
			"string":  ":attribute نباید بیشتر از :max کاراکتر داشته باشد.",
			"array":   ":attribute نباید بیشتر از :max آیتم داشته باشد.",
		},
		"mimes":     "فرمت‌های معتبر فایل عبارتند از: :values.",
		"mimetypes": "فرمت‌های معتبر فایل عبارتند از: :values.",
		"min": map[string]interface{}{
			"numeric": ":attribute نباید کوچکتر از :min باشد.",
			"file":    ":attribute نباید کوچکتر از :min کیلوبایت باشد.",
			"string":  ":attribute نباید کمتر از :min کاراکتر داشته باشد.",
			"array":   ":attribute نباید کمتر از :min آیتم داشته باشد.",
		},
		"not_in":               ":attribute انتخاب شده، معتبر نیست.",
		"not_regex":            "فرمت :attribute معتبر نیست.",
		"numeric":              ":attribute باید عدد یا رشته‌ای از اعداد باشد.",
		"present":              "فیلد :attribute باید در پارامترهای ارسالی وجود داشته باشد.",
		"regex":                "فرمت :attribute معتبر نیست.",
		"required":             "فیلد :attribute الزامی است.",
		"required_if":          "هنگامی که :other برابر با :value است، فیلد :attribute الزامی است.",
		"required_unless":      "فیلد :attribute الزامی است، مگر آنکه :other در :values موجود باشد.",
		"required_with":        "در صورت وجود فیلد :values، فیلد :attribute نیز الزامی است.",
		"required_with_all":    "در صورت وجود فیلدهای :values، فیلد :attribute نیز الزامی است.",
		"required_without":     "در صورت عدم وجود فیلد :values، فیلد :attribute الزامی است.",
		"required_without_all": "در صورت عدم وجود هر یک از فیلدهای :values، فیلد :attribute الزامی است.",
		"same":                 ":attribute و :other باید همانند هم باشند.",
		"size": map[string]interface{}{
			"numeric": ":attribute باید برابر با :size باشد.",
			"file":    ":attribute باید برابر با :size کیلوبایت باشد.",
			"string":  ":attribute باید برابر با :size کاراکتر باشد.",
			"array":   ":attribute باید شامل :size آیتم باشد.",
		},
		"starts_with": ":attribute باید با یکی از این ها شروع شود: :values",
		"string":      "فیلد :attribute باید رشته باشد.",
		"timezone":    "فیلد :attribute باید یک منطقه زمانی معتبر باشد.",
		"unique":      ":attribute قبلا انتخاب شده است.",
		"uploaded":    "بارگذاری فایل :attribute موفقیت آمیز نبود.",
		"url":         ":attribute معتبر نمی‌باشد.",
		"uuid":        ":attribute باید یک UUID معتبر باشد.",
	},
	"en": {
		"accepted":        "The :attribute must be accepted.",
		"active_url":      "The :attribute is not a valid URL.",
		"after":           "The :attribute must be a date after :date.",
		"after_or_equal":  "The :attribute must be a date after or equal to :date.",
		"alpha":           "The :attribute may only contain letters.",
		"alpha_dash":      "The :attribute may only contain letters, numbers, dashes and underscores.",
		"alpha_num":       "The :attribute may only contain letters and numbers.",
		"array":           "The :attribute must be an array.",
		"before":          "The :attribute must be a date before :date.",
		"before_or_equal": "The :attribute must be a date before or equal to :date.",
		"between": map[string]interface{}{
			"numeric": "The :attribute must be between :min and :max.",
			"file":    "The :attribute must be between :min and :max kilobytes.",
			"string":  "The :attribute must be between :min and :max characters.",
			"array":   "The :attribute must have between :min and :max items.",
		},
		"boolean":        "The :attribute field must be true or false.",
		"confirmed":      "The :attribute confirmation does not match.",
		"date":           "The :attribute is not a valid date.",
		"date_equals":    "The :attribute must be a date equal to :date.",
		"date_format":    "The :attribute does not match the format :format.",
		"different":      "The :attribute and :other must be different.",
		"digits":         "The :attribute must be :digits digits.",
		"digits_between": "The :attribute must be between :min and :max digits.",
		"dimensions":     "The :attribute has invalid image dimensions.",
		"distinct":       "The :attribute field has a duplicate value.",
		"email":          "The :attribute must be a valid email address.",
		"ends_with":      "The :attribute must end with one of the following: :values",
		"exists":         "The selected :attribute is invalid.",
		"file":           "The :attribute must be a file.",
		"filled":         "The :attribute field must have a value.",
		"gt": map[string]interface{}{
			"numeric": "The :attribute must be greater than :value.",
			"file":    "The :attribute must be greater than :value kilobytes.",
			"string":  "The :attribute must be greater than :value characters.",
			"array":   "The :attribute must have more than :value items.",
		},
		"gte": map[string]interface{}{
			"numeric": "The :attribute must be greater than or equal :value.",
			"file":    "The :attribute must be greater than or equal :value kilobytes.",
			"string":  "The :attribute must be greater than or equal :value characters.",
			"array":   "The :attribute must have :value items or more.",
		},
		"image":    "The :attribute must be an image.",
		"in":       "The selected :attribute is invalid.",
		"in_array": "The :attribute field does not exist in :other.",
		"integer":  "The :attribute must be an integer.",
		"ip":       "The :attribute must be a valid IP address.",
		"ipv4":     "The :attribute must be a valid IPv4 address.",
		"ipv6":     "The :attribute must be a valid IPv6 address.",
		"json":     "The :attribute must be a valid JSON string.",
		"lt": map[string]interface{}{
			"numeric": "The :attribute must be less than :value.",
			"file":    "The :attribute must be less than :value kilobytes.",
			"string":  "The :attribute must be less than :value characters.",
			"array":   "The :attribute must have less than :value items.",
		},
		"lte": map[string]interface{}{
			"numeric": "The :attribute must be less than or equal :value.",
			"file":    "The :attribute must be less than or equal :value kilobytes.",
			"string":  "The :attribute must be less than or equal :value characters.",
			"array":   "The :attribute must not have more than :value items.",
		},
		"max": map[string]interface{}{
			"numeric": "The :attribute may not be greater than :max.",
			"file":    "The :attribute may not be greater than :max kilobytes.",
			"string":  "The :attribute may not be greater than :max characters.",
			"array":   "The :attribute may not have more than :max items.",
		},
		"mimes":     "The :attribute must be a file of type: :values.",
		"mimetypes": "The :attribute must be a file of type: :values.",
		"min": map[string]interface{}{
			"numeric": "The :attribute must be at least :min.",
			"file":    "The :attribute must be at least :min kilobytes.",
			"string":  "The :attribute must be at least :min characters.",
			"array":   "The :attribute must have at least :min items.",
		},
		"not_in":               "The selected :attribute is invalid.",
		"not_regex":            "The :attribute format is invalid.",
		"numeric":              "The :attribute must be a number.",
		"present":              "The :attribute field must be present.",
		"regex":                "The :attribute format is invalid.",
		"required":             "The :attribute field is required.",
		"required_if":          "The :attribute field is required when :other is :value.",
		"required_unless":      "The :attribute field is required unless :other is in :values.",
		"required_with":        "The :attribute field is required when :values is present.",
		"required_with_all":    "The :attribute field is required when :values are present.",
		"required_without":     "The :attribute field is required when :values is not present.",
		"required_without_all": "The :attribute field is required when none of :values are present.",
		"same":                 "The :attribute and :other must match.",
		"size": map[string]interface{}{
			"numeric": "The :attribute must be :size.",
			"file":    "The :attribute must be :size kilobytes.",
			"string":  "The :attribute must be :size characters.",
			"array":   "The :attribute must contain :size items.",
		},
		"starts_with": "The :attribute must start with one of the following: :values",
		"string":      "The :attribute must be a string.",
		"timezone":    "The :attribute must be a valid zone.",
		"unique":      "The :attribute has already been taken.",
		"uploaded":    "The :attribute failed to upload.",
		"url":         "The :attribute format is invalid.",
		"uuid":        "The :attribute must be a valid UUID.",
	},
}

func translate(key string, args ...interface{}) string {
	obj := messages[lang]
	msg := eval(key, obj)
	if msg == "" {
		return ""
	}
	j := 0
	li := 0
	ln := len(msg)
	st := false
	ot := ""
	aln := len(args)
	for i := 0; i < ln; i++ {
		if msg[i] == ':' {
			st = true
			ot += msg[li:i]
			li = i
		} else if (msg[i] == ' ' || msg[i] == ',' || msg[i] == '.') && st {
			if j < aln {
				ot += fmt.Sprintf("%v", args[j])
				li += len(msg[li:i])
				j++
			}
			st = false
		}
	}
	ot += msg[li:ln]
	return ot
}

func eval(path string, from map[string]interface{}) string {
	firstIndex := strings.Index(path, ".")
	var first string
	var rest string
	if firstIndex > -1 {
		first = path[:firstIndex]
		rest = path[firstIndex+1:]
	} else {
		first = path
		rest = ""
	}

	ifc := from[first]
	if value, ok := ifc.(string); ok {
		return value
	}
	if value, ok := ifc.(map[string]interface{}); ok && rest != "" {
		return eval(rest, value)
	}
	return ""
}
