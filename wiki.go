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
	NumTest int
	Test    string `json:"test"`
	Status  bool   `json:"status"`
	Last    bool   `json:"last"`
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var udata udata

	json.Unmarshal(byteValue, &udata)

	//p := &data{NumTest: len(udata.UserData), Test: udata.UserData[0].Test, Status: udata.UserData[0].Status, Last: udata.UserData[0].Last}
	f := make(map[int]interface{})
	var t int = 0
	for i := 0; i < len(udata.UserData); i++ {
		for e := 1; e < 3; e++ {
			if e == 1 {
				f[t] = udata.UserData[i].Test
				t++
			} else if e == 2 {
				f[t] = udata.UserData[i].Status
				t++
			} else if e == 3 {
				f[t] = udata.UserData[i].Last
				t++
			}
		}
	}
	f[100] = len(udata.UserData)
	b, _ := template.ParseFiles("dashboard.html")
	b.Execute(w, udata)

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
