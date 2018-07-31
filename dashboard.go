package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
) // the mgo packages are 3rd party and will have to be installed with the command: 'go get gopkg.in/mgo.v2'

type graphData struct { //this struct shows the type of data we are using
	Test        string
	testStatus  bool
	id          string
	timeStarted time.Time
	timeEnded   string // this should eventually become time.Time. However, this is only sample data and no tests were actully running
}

//this is the html code that is being used and the one that will have to be edited to change the page
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
  
  <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
    <script type="text/javascript">
      google.charts.load('current', {'packages':['bar']});
      google.charts.setOnLoadCallback(drawChart);
  	function drawChart() {
        var data = google.visualization.arrayToDataTable([
          ['Year', "value"],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t1', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000],
          ['t2', 1000],
          ['t3', 1000],
          ['t4', 1000]
        ]);

        var options = {
          chart: {
            title: 'Tests',
            subtitle: 'Green: pass ; red: failed',
          },
          colors: ['#e0440e', '#e6693e', '#ec8f6e', '#f3b49f', '#f6c7b6'],
          bars: 'vertical' // Required for Material Bar Charts.
        };

        var chart = new google.charts.Bar(document.getElementById('barchart'));

        chart.draw(data, google.charts.Bar.convertOptions(options));
      }
	
  </script>
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
            <p>6(not actual number)</p>
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
          <div id="barchart" style="width: 9000px; height: 500px;"></div>

  <div class="row">
        <div class="col-sm-4">
          <div class="well">
            <a href = 'https://google.com'>Test: {{ .Test}}</a>
            <p>Status: {{index .Status 0}}</p> 
            <p>Last Attmept Succeded? {{index .Last 0}}</p> 
          </div>
        </div>
        <div class="col-sm-4">
          <div class="well">
            <a href = 'google.com'>Test: {{index .Test 1}}</a>
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
	session, err := mgo.Dial("localhost:27017") //27017 is the default port for mangodb but this can change depending on comands that can be accessed from the mango command line
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("local").C("data") //this specifies where the code is operating in this case local accouunt in a collection called data
	/*err = c.Insert(&graphData{"name of test a", false, "last", time.Now(), time.Now().Format("Jan 3: 7:55 PM")})
	if err != nil {
		log.Fatal(err)
	}*/                   //this bit of code can be used for adding data to the database
	result := graphData{} //this specifies interface
	err = c.Find(bson.M{"test": "name of test a"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Test)

	t := template.Must(template.New("tmpl").Parse(tmpl)) //this parses the tml code from above

	t.Execute(w, result)

}

func main() {
	http.HandleFunc("/view/", viewHandler) // this gets the entered url

	log.Fatal(http.ListenAndServe(":8080", nil)) //says which port to active the server on
}
