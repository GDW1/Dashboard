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

var tmplString = `
<!DOCTYPE html>
<html lang="en">
<head>
  <title>Dashboard</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <style>
    /* Set height of the grid so .sidenav can be 100% (adjust as needed) */
    .row.content {height: 550px}
    
    /* Set gray background color and 100% height */
    .sidenav {
      background-color: #f1f1f1;
      height: 500%;
    }
        
    /* On small screens, set height to 'auto' for the grid */
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
      <img id='logo'src="https://cybersecurity-excellence-awards.com/wp-content/uploads/2017/06/598614.png" alt="Trulli" width="350" height="100" >
      <ul class="nav nav-pills nav-stacked">
        <li class="active"><a href="#section1">Dashboard</a></li>
       <!--<li><a href="#section2">Age</a></li>
        <li><a href="#section3">Gender</a></li>
        <li><a href="#section3">Geo</a></li>-->
      </ul><br>
    </div>
    <br>
    
    <div class="col-sm-9">
      <div class="well">
        <h1>Dashboard: All current tests</h4>
      </div>
      <div class="row">
        <div class="col-sm-3">
          <div class="well">
            <h4>Current Tests Going On</h4>
            <p>{{.100}}</p> 
          </div>
        </div>
        <div class="col-sm-3">
          <div class="well">
            <h4><!--Place text here --></h4>
            <p><!--Place text here --></p> 
          </div>
        </div>
        <div class="col-sm-3">
          <div class="well">
            <h4><!--Place text here --></h4>
            <p><!--Place text here --></p> 
          </div>
        </div>
        <div class="col-sm-3">
          <div class="well">
            <h4><!--Place text here --></h4>
            <p><!--Place text here --></p> 
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-4">
          <div class="well">
            <p>Test</p> 
            <p>Status</p> 
            <p>Last Attmept Succeded?</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test</p> 
            <p>Status</p> 
            <p>Last Attmept Succeded?</p>
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test</p> 
            <p>Status</p> 
            <p>Last Attmept Succeded?</p>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-4">
          <div class="well">
            <p>Test</p> 
            <p>Status</p> 
            <p>Last Attmept Succeded?</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test</p> 
            <p>Status</p> 
            <p>Last Attmept Succeded?</p>
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <p>Test</p> 
            <p>Status</p> 
            <p>Last Attmept Succeded?</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

</body>
</html>
`

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
	//b, _ := template.ParseFiles("dashboard.html")
	//b.Execute(w, f)
	tmpl, err := template.New("test").Parse(tmplString)
	tmpl.ExecuteTemplate(os.Stdout, "dashboard", f)

}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
