package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

func createFile(file, tmpl string, fd fieldData) error {
	structTmpl := template.Must(template.ParseFiles(fmt.Sprintf("./_template/%s", tmpl)))
	f, err := os.Create(fmt.Sprintf("./day%s/%s", fd.Day, file))
	if err != nil {
		return (err)
	}
	if err := structTmpl.Execute(f, fd); err != nil {
		return (err)
	}
	f.Close()
	return nil
}

type fieldData struct {
	Day string
}

func main() {
	day := time.Now().Format("02")
	dayStr := fmt.Sprintf("day%s", day)

	fd := fieldData{Day: day}

	if err := os.Mkdir(dayStr, 0777); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	if err := os.Mkdir(fmt.Sprintf("%s/cmd", dayStr), 0777); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	os.Create(fmt.Sprintf("./%s/input.txt", dayStr))
	os.Create(fmt.Sprintf("./%s/input_test.txt", dayStr))

	cases := []struct {
		File     string
		Template string
	}{{
		File:     fmt.Sprintf("day%s.go", day),
		Template: "day.tmpl",
	}, {
		File:     "part1.go",
		Template: "part1.tmpl",
	}, {
		File:     "part2.go",
		Template: "part2.tmpl",
	}, {
		File:     fmt.Sprintf("day%s_test.go", day),
		Template: "day_test.tmpl",
	}, {
		File:     "cmd/main.go",
		Template: "cmd/main.tmpl",
	}}
	for _, c := range cases {
		if err := createFile(c.File, c.Template, fd); err != nil {
			log.Fatal(err)
		}
	}
}
