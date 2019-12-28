package vstruct

func LoadBuiltin() {

	RegisterLanguage("en", map[string]interface{}{
		"accepted":    "The :attribute must be accepted.",
		"alpha":       "The :attribute may only contain letters.",
		"alpha_dash":  "The :attribute may only contain letters, numbers, dashes and underscores.",
		"alpha_num":   "The :attribute may only contain letters and numbers.",
		"username":   "The :attribute may only contain letters, underscores, dots and numbers.",
		"boolean":     "The :attribute field must be true or false.",
		"confirmed":   "The :attribute confirmation does not match.",
		"email":       "The :attribute must be a valid email address.",
		"ends_with":   "The :attribute must end with one of the following, :values",
		"filled":      "The :attribute field must have a value.",
		"in":          "The selected :attribute is invalid.",
		"json":        "The :attribute must be a valid JSON string.",
		"not_in":      "The selected :attribute is invalid.",
		"not_regex":   "The :attribute format is invalid.",
		"regex":       "The :attribute format is invalid.",
		"required":    "The :attribute field is required.",
		"starts_with": "The :attribute must start with one of the following, :values",
		"string":      "The :attribute must be a string.",
		"phone":       "The :attribute must be a valid phone number.",
		"mobile":      "The :attribute must be a valid mobile number.",
		"persian":     "The :attribute may only contain persian letters.",
		"ssn":         "The :attribute must be a valid social security number.",
		"max": map[string]interface{}{
			"numeric": "The :attribute may not be greater than :max.",
			"string":  "The :attribute may not be greater than :max characters.",
		},
		"min": map[string]interface{}{
			"numeric": "The :attribute must be at least :min.",
			"string":  "The :attribute must be at least :min characters.",
		},
		"between": map[string]interface{}{
			"numeric": "The :attribute must be between :min and :max.",
			"string":  "The :attribute must be between :min and :max characters.",
		},
		"size": map[string]interface{}{
			"string":  "The :attribute must be :value characters.",
		},
	})

	RegisterLanguage("fa", map[string]interface{}{
		"accepted":    ":attribute باید پذیرفته شده باشد.",
		"alpha":       ":attribute باید فقط حروف الفبا باشد.",
		"alpha_dash":  ":attribute باید فقط حروف الفبا، اعداد، خط تیره و زیرخط باشد.",
		"alpha_num":   ":attribute باید فقط حروف الفبا و اعداد باشد.",
		"username":  ":attribute باید فقط حروف الفبا، اعداد، نقطه و زیرخط باشد.",
		"boolean":     "فیلد :attribute فقط می‌تواند true و یا false باشد.",
		"confirmed":   ":attribute با فیلد تکرار مطابقت ندارد.",
		"email":       ":attribute باید یک ایمیل معتبر باشد.",
		"ends_with":   "فیلد :attribute باید با یکی از مقادیر زیر خاتمه یابد، :values",
		"filled":      "فیلد :attribute باید مقدار داشته باشد.",
		"in":          ":attribute انتخاب شده، معتبر نیست.",
		"json":        "فیلد :attribute باید یک رشته از نوع JSON باشد.",
		"not_in":      ":attribute انتخاب شده، معتبر نیست.",
		"not_regex":   "فرمت :attribute معتبر نیست.",
		"regex":       "فرمت :attribute معتبر نیست.",
		"required":    "فیلد :attribute الزامی است.",
		"starts_with": ":attribute باید با یکی از این ها شروع شود، :values",
		"string":      "فیلد :attribute باید رشته باشد.",
		"phone":       ":attribute باید یک شماره تلفن معتبر باشد.",
		"mobile":      ":attribute باید یک شماره همراه معتبر باشد.",
		"persian":     ":attribute باید فقط حروف الفبای فارسی باشد.",
		"ssn":         ":attribute باید یک کد ملی معتبر باشد.",
		"max": map[string]interface{}{
			"numeric": ":attribute نباید بزرگتر از :max باشد.",
			"string":  ":attribute نباید بیشتر از :max کاراکتر داشته باشد.",
		},
		"min": map[string]interface{}{
			"numeric": ":attribute نباید کوچکتر از :min باشد.",
			"string":  ":attribute نباید کمتر از :min کاراکتر داشته باشد.",
		},
		"between": map[string]interface{}{
			"numeric": ":attribute باید بین :min و :max باشد.",
			"string":  ":attribute باید بین :min و :max کاراکتر باشد.",
		},
		"size": map[string]interface{}{
			"string":  ":attribute باید :value کاراکتر باشد.",
		},
	})
	loadBoolean()
	loadInt()
	loadString()
	loadAny()
}
