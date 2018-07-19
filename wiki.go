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
	Test   string `json:"test"`
	Status string `json:"status"`
	Last   bool   `json:"last"`
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	//defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var udata udata

	json.Unmarshal(byteValue, &udata)
	title := r.URL.Path[len("/view/"):]
	var p *Page
	if title != "dashboard" {
		p, err = loadPage(title)
	} else {
		p = &Page{Title: udata.UserData[0].Test, Body: []byte(udata.UserData[0].Status)}
	}
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	if title == "dashboard" {
		renderTemplate(w, "dashboard", p)
	} else {
		renderTemplate(w, "view", p)
	}
}
func editHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}
func renderDash(w http.ResponseWriter, tmpl string, u *Page) {

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, u)
}
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
