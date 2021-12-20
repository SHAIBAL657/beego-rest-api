# beego-rest-api


### Prerequisites

    ~ Install go and beego
    ~ Set GOPATH
    ~ Download https://github.com/SHAIBAL657/bee-form-validation and https://github.com/SHAIBAL657/beego-rest-api/sara project.
    ~ Both project should be run on individual port

### How to test
    ~ go mod tidy
    ~ bee run
    ~ go get <library-name> (if not available)
    ~ This run on port 8080 (http://localhost:8080/swagger)
    ~ Enter valid data to create a user. This give response back accordingly.

### Features 
    ~ Valid email,phone,date should be entered.
    ~ Validate data before sending to database.
    ~ Password stored as hase value.
