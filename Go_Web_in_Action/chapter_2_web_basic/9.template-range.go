package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	//创建一个模版
	rangeTemplate := `
		{{if .Kind}}
		{{range $i, $v := .MapContent}}
		{{$i}} => {{$v}} , {{$.OutsideContent}}
		{{end}}
		{{else}}
		{{range .MapContent}}
		{{.}} , {{$.OutsideContent}}
		{{end}}    
		{{end}}`

	str1 := []string{"first time range", "with index and value"}
	str2 := []string{"second time range", "without index and value"}

	type Content struct {
		MapContent     []string
		OutsideContent string
		Kind           bool
	}
	var contents = []Content{
		{str1, "first time outside content", true},
		{str2, "second time outside content", false},
	}

	// create and parse template
	t := template.Must(template.New("range").Parse(rangeTemplate))

	// execute and render template
	for _, c := range contents {
		err := t.Execute(os.Stdout, c)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
