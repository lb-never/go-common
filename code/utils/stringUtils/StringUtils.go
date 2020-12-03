package stringUtils

import "strconv"

func GetStringValue(val interface{}) (value string, flag bool) {
	if val == nil {
		return "", false
	}
	switch val.(type) {
	case string:
		return val.(string), true
	case int:
		return strconv.Itoa(val.(int)), true
	case int64:
		return strconv.FormatInt(val.(int64), 10), true
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'E', -1, 32), true
	case float64:
		return strconv.FormatFloat(val.(float64), 'E', -1, 64), true
	}
	return "", false
}
