package weather

import "text/template"

const templ = `Weather For {{.Name}} :
{{- range $index, $word := .Weather}}
Description: {{$word.Description}}
{{- end}}
Temperature: {{.Main.Temp | toCelsius}} C
Max Temperature: {{.Main.TempMax | toCelsius}} C
Min Temperature: {{.Main.TempMin | toCelsius}} C
`

var Report = template.Must(template.New("weather").
	Funcs(template.FuncMap{"toCelsius": toCelsius}).
	Parse(templ))

func toCelsius(temp float32) float32 {
	return temp - 273
}
