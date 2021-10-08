### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- `go run main.go`
- Import Postman collection from docs to check example

---
### Migrations
- Run command bellow to migration: 
   ...

### Structure Project
 - 	framework
		/helpers: Define helper functions
		/libs: Define function use lib like: redis, rabbit, cloud function..
 -	sdk
		- sdk to call services internal
 -	internal
		/routes: Define routes and point to services
		/services: Define logic , which routes call and response to client
		/repositories: Define query connect db
		/models: Define object fields
		/migrations: Define migration files

	
### Flow start project:
	- call migration -> run application 
#### Docs API: 
 - Open swagger document api: ...
#### Document source
 - 

#### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
