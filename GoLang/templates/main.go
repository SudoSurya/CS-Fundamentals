package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

const temp = `   {{ .Name}}
    {{ .Age}}
    {{range .Fruits}}
    {{.FirstName}}
    {{.LastName}}
    {{end}}
`

type Fruits struct {
	FirstName string
	LastName  string
}

func main() {
	var report = template.Must(template.New("report").Parse(temp))
	var data = struct {
		Name   string
		Age    int
		Fruits []Fruits
	}{
		Name:   "John",
		Age:    30,
		Fruits: []Fruits{{FirstName: "Apple", LastName: "Red"}, {"Banana", "Yellow"}},
	}

	/* if _, err := os.Stat(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName), 0751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	} */
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
	fileInfo, _ := os.Stat(dir)
	log.Println(fileInfo.Name())
	err := os.MkdirAll("test", 0751)
	if err != nil {
		log.Fatal(err)
	}
	/* err := report.Execute(os.Stdout, map[string]interface{}{
	    "Name": "John",
	    "Age":  30,
	}) */
	if err != nil {
		log.Fatal(err)
	}
}


func findMaxConsecutiveOnes(nums []int) int {
   var count int = 0
   var mx int = math.MinInt64
   for _,V := range nums{
       if V==1{
           count ++
       }else {
           mx = int(math.Max(float64(mx),float64(count)))
           count = 0
       }
   }
   return mx
}
