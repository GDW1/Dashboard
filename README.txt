Each of the files in the folder do the following:
	Dashboard: This file can be run to access the website at 'localhost:8080/view/dashboard'
	Dashboard.go: this contains a majority of the code running for the program. this should compile into the dashboard file using the terminal command 'go build dashboard.go'
	index.html: this file contains the raw html code and can be opened to see what the file would look like. Editing this will not affect the result in dashboard as there is a separate copy in a variable called tmpl on the top level of dashboard.go
	readme.txt: this contains data on each of the files in the folder and instructions on running the files
	mongodb-osx-ssl-x86_64-4.0.0: this is a copy of mongodb 4.0.0 for Mac OSX and MacOS. Instructions for this will be labeled further in this file
The data folder:
	The data folder contains all the files required for mongodb in the project.If you are having trouble with the folder, delete the content of the db folder inside it and run the command './mongod --dbpath (path to the db folder) and files will appear in it that are used for the database
How to activate the webpage:
	1: turn on mongodb using the command above(you don't have to delete your files)
	2: in the parent folder run 'go build dashboard.go'
	3: run dashboard or type in ./dashboard
	NOTE: the dashboard file that is currently packaged is the most recent build on July 31st 2018 and will get replaced by running the build command
Using mongoDB:
	NOTE: this uses a very simple version of mongodb with a library that enables it with go called mgo
	1: to get started, extract the files from the compressed mongodb folder (macOS should not require any external tools)
	2: take a look at the readme in that folder and continue to the bin folder
	3: run the command './mongod --dbpath (path to the db folder)'
	OPTIONAL 4:To use the javascript enabled shell type in ./mongo
To be done:
	The following steps would be to integrate the backend to load the data as well as begin to debug some of the errors on the front end such as the graph being too long.
	to proceed, I would reconmend writing a script that had to process the information int a database in a clean format. as well as makes sure there are no identical entries
