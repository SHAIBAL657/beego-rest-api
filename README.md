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

### Create DATABASE postgresql (PGADMIN)

    createdb:=CREATE TABLE IF NOT EXISTS public."USER"
    (
    firstname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    lastname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    phone character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    dob character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "USER_pkey" PRIMARY KEY (email)
    )
    
    db.Exec(createdb)
