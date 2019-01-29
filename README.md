Server code for buying/selling currency app

setup for project:
1. install go: `https://golang.org/dl/`
2. install nodejsC: `https://nodejs.org/download/`
3. go to environment variables in control panel, under "System Variables", click "New"
	Variable name: GOPATH
	Variable value: C:\golang
4. place project at `C:\golang\src\exchange.com`, so the server.go will have path `C:\golang\src\exchange.com\exchange\server.go`
5. setup project database by runnning exchange\exchange_20150729.sql
6. open command prompt, go to "C:\golang\src\exchange.com\exchange"
7. type "go get" to download all libraries which the project is using
8. type "go run server.go" to run the project


to send email in forgot password
1. `Controller/UserController.go` -> Forgot_Password_Action
2. Set a <email> and <password>



guides:
http://stackoverflow.com/questions/572768/styling-an-input-type-file-button