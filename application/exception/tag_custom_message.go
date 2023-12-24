package exception

var TagCustomMessage = map[string]func(...interface{})string{
	"required": RequiredMessageVal,
}


// RequiredMessageVal is a custom message for required tag
func RequiredMessageVal(data ...interface{}) string {
	field := data[0].(string)
	return "field '" + field + "' tidak boleh kosong"
}
