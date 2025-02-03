package output

import "github.com/fatih/color"

func PrintError(errValue any) {
	switch v := errValue.(type) {
	case string:
		color.Red(v)
	case error:
		color.Red(v.Error())
	case int:
		color.Red("Код ошибки: %d", v)
	default:
		color.Red("Неизвестный тип ошибки")
	}
}
