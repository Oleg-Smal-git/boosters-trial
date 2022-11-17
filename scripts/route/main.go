package main

import (
	"bufio"
	"os"
	"strings"
	"text/template"

	"github.com/Oleg-Smal-git/boosters-trial/app/config"

	helpers "github.com/samber/lo"
)

var (
	// allowedMethods represents a set of allowed methods for endpoint.
	allowedMethods = map[string]struct{}{
		"GET": {}, "PUT": {}, "POST": {}, "DELETE": {}, "PATCH": {},
		"HEAD": {}, "CONNECT": {}, "OPTIONS": {}, "TRACE": {},
	}
)

// interpolationData wraps stuff we put in the template.
type interpolationData struct {
	Handlers    map[string]map[string]string
	Controllers map[string]string
}

// This will just panic if something goes wrong,
// which is fine cuz I didn't sleep last night, and I don't care at this point.
// But also because it's a pre-build stage ^^.
func main() {
	in, err := os.Open(config.BasePath() + "/app/config/routes")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)
	data := interpolationData{
		Handlers:    make(map[string]map[string]string),
		Controllers: make(map[string]string),
	}
	for scanner.Scan() {
		row := strings.Split(strings.ReplaceAll(scanner.Text(), "\t", " "), " ")
		row = helpers.Filter(row, func(s string, i int) bool { return s != "" })
		if len(row) == 0 {
			continue
		}
		if _, ok := allowedMethods[row[0]]; !ok {
			panic(row[0] + "method is not valid")
		}
		if _, ok := data.Handlers[row[0]]; !ok {
			data.Handlers[row[0]] = make(map[string]string)
		}
		upC := strings.Split(row[2], ".")[0]
		lowC := strings.ToLower(upC[:1]) + upC[1:]
		data.Controllers[upC] = lowC
		data.Handlers[row[0]][row[1]] = lowC + "." + strings.Split(row[2], ".")[1]
	}
	rawTemplate, err := os.ReadFile(config.BasePath() + "/scripts/route/_template.go.tmp")
	if err != nil {
		panic(err)
	}
	tmp, err := template.New("routes").Parse(string(rawTemplate))
	if err != nil {
		panic(err)
	}
	_ = os.Remove(config.BasePath() + "/app/controllers/handlers--autogenerated.go")
	out, err := os.Create(config.BasePath() + "/app/controllers/handlers--autogenerated.go")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	err = tmp.Execute(out, data)
	if err != nil {
		panic(err)
	}
}
