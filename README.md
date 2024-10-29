# Nexu Backend Coding Exercise



## Overview
###API Endpoints

```
                              GET    /brands
                              GET    /brands/:id/models
                              POST   /brands
                              POST   /brands/:id/models
                              PUT    /models/:id
                              GET    /models
```
### Data
The project uses migrations to create the database structure, for this it is required to provide an environment variable. 
The database structure is created to preserve consistency and reduce redundancy. 

With the information provided in the .json file, two tables were created, one for Brands and one for Models, which are located in the migrations folder and the scripts for their creation. 

When running the project these will be created automatically. 
Then you can run the following command to populate the tables with the information from the .json file
### Generate data 
```script
go run .\services\data-file.go
```

##Instructions

###Build
```script
go build
```

###Run 
```bash
go run .\main.go
```

###Test 
To run the tests, use:
```bash
go test ./.. -v
```
This will run all the files with *_test.go

I decided to do the project in golang because it is one of the languages that has more projection due to its handling of routines, plus one of the objectives of this exercise was to have fun.

Some aspects of this project are that in the repository layer Dependency injection was used, in order to perform the unit tests of the project, of which only some were added.
In addition, we included the docker-compose file and its respective Dockerfile, as well as the makefile, which will help us if in the future we want to use containers to run this project.   

