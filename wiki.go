package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

const tmpl = `<html lang="en">
	<head>
  <title>Dashboard</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <style>
     
    .row.content {height: 550px}
    
     
    .sidenav {
      background-color: #f1f1f1;
      height: 500%;
    }
        
     
    @media screen and (max-width: 767px) {
      .row.content {height: auto;} 
    }
    img{
        padding: 10px;
        padding-left: 21px;
    }
  </style>
</head>
<body>
<nav class="navbar navbar-inverse visible-xs">
  <div class="container-fluid">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>                        
      </button>
      <a class="navbar-brand" href="#">Logo</a>
    </div>
    <div class="collapse navbar-collapse" id="myNavbar">
      <ul class="nav navbar-nav">
        <li class="active"><a href="#">Dashboard</a></li>
        <li><a href="#">Age</a></li>
        <li><a href="#">Gender</a></li>
        <li><a href="#">Geo</a></li>
      </ul>
    </div>
  </div>
</nav>
<div class="container-fluid">
  <div class="row content">
    <div class="col-sm-3 sidenav hidden-xs">
      <img id="logo" src="https://cybersecurity-excellence-awards.com/wp-content/uploads/2017/06/598614.png" alt="Trulli" width="350" height="100">
      <ul class="nav nav-pills nav-stacked">
        <li class="active"><a href="#section1">Dashboard</a></li>
       
      </ul><br>
    </div>
    <br>
    
    <div class="col-sm-9">
      <div class="well">
        <h1>Dashboard: All current tests</h1>
      </div>
      <div class="row">
        <div class="col-sm-3">
          <div class="well">
            <h4>Current Tests Going On</h4>
            <p>6</p>
          </div>
        </div>
        <div class="col-sm-3">
          <div class="well">
            <h4></h4>
            <p></p> 
          </div>
        </div>
        <div class="col-sm-3">
          <div class="well">
            <h4></h4>
            <p></p> 
          </div>
        </div>
        <div class="col-sm-3">
          <div class="well">
            <h4></h4>
            <p></p> 
          </div>
        </div>
      </div>
      
<div class="row">
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{ .Test}}</p>
            <p>Status: {{index .Status 0}}</p> 
            <p>Last Attmept Succeded? {{index .Last 0}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 1}}</p>
            <p>Status: {{index .Status 1}}</p> 
            <p>Last Attmept Succeded? {{index .Last 1}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 2}}</p>
            <p>Status: {{index .Status 2}}</p> 
            <p>Last Attmept Succeded? {{index .Last 2}}</p> 
          </div>
        </div>
      </div><div class="row">
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 3}}</p>
            <p>Status: {{index .Status 3}}</p> 
            <p>Last Attmept Succeded? {{index .Last 3}}</p>  
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 4}}</p>
            <p>Status: {{index .Status 4}}</p> 
            <p>Last Attmept Succeded? {{index .Last 4}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
           <p>Test: {{index .Test 5}}</p>
            <p>Status: {{index .Status 5}}</p> 
            <p>Last Attmept Succeded? {{index .Last 5}}</p> 
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 6}}</p>
            <p>Status: {{index .Status 6}}</p> 
            <p>Last Attmept Succeded? {{index .Last 6}}</p>  
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
           <p>Test: {{index .Test 7}}</p>
            <p>Status: {{index .Status 7}}</p> 
            <p>Last Attmept Succeded? {{index .Last 7}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 8}}</p>
            <p>Status: {{index .Status 8}}</p> 
            <p>Last Attmept Succeded? {{index .Last 8}}</p> 
          </div>
        </div>
      </div>
<div class="row">
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 9}}</p>
            <p>Status: {{index .Status 9}}</p> 
            <p>Last Attmept Succeded? {{index .Last 9}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 10}}</p>
            <p>Status: {{index .Status 10}}</p> 
            <p>Last Attmept Succeded? {{index .Last 10}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test: {{index .Test 11}}</p>
            <p>Status: {{index .Status 11}}</p> 
            <p>Last Attmept Succeded? {{index .Last 11}}</p> 
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</body></html>`

func viewHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("data")
	//err = c.Insert(&data{"name of test a", "stat", "last"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	result := data{}
	err = c.Find(bson.M{"test": "name of test a"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Test)

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

	t.Execute(w, result)

}

func main() {
	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
