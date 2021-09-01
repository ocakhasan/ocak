package weather

import "text/template"

const templ = `Weather For {{.Name}},
{{- range $index, $word := .Weather}}
Description: {{$word.Description}}
{{- end}}
Temperature: {{.Main.Temp | toCelsius | printf "%.2f"}}°C
Max Temperature: {{.Main.TempMax | toCelsius | printf "%.2f" }}°C
Min Temperature: {{.Main.TempMin | toCelsius | printf "%.2f" }}°C
`

var Report = template.Must(template.New("weather").
	Funcs(template.FuncMap{"toCelsius": toCelsius}).
	Parse(templ))

func toCelsius(temp float32) float32 {
	return temp - 273
}
