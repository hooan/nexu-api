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

