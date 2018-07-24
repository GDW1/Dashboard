package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type udata struct {
	UserData []data `json:"data"`
}

type data struct {
	//	NumTest int
	Test   string `json:"test"`
	Status string `json:"status"`
	Last   string `json:"last"`
}

type compiledData struct {
	NumTest int
	Test    []string
	Status  []string
	Last    []string
}

const tmpl = `
	{{.Test}}
	{{.Status}}
	{{.Last}}
`

func viewHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var udata udata

	json.Unmarshal(byteValue, &udata)

	var f compiledData
	for i := 0; i < len(udata.UserData); i++ {
		for e := 1; e <= 3; e++ {
			if e == 1 {
				f.Test = append(f.Test, udata.UserData[i].Test)
			} else if e == 2 {
				f.Status = append(f.Status, udata.UserData[i].Status)
			} else if e == 3 {
				f.Last = append(f.Status, udata.UserData[i].Last)
			}
		}
	}
	f.NumTest = len(udata.UserData)
	//b, _ := template.ParseFiles("Index.html")
	t := template.Must(template.New("tmpl").Parse(tmpl))

	t.Execute(w, f)

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
