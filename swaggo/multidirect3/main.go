package main

import "swaggo/multidirect3/pkg/ctrl"

// @title swaggo multidirect
// @version 1.0
// @description This is a sample server swaggo multidirect server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host petstore.swagger.io
// @BasePath /v2
func main(){
	r := ctrl.HttpRouter()
	r.Run(":1234")
}