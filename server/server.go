package server

import "github.com/decabits/vwo-golang-example-app/config"

//Init function starts the server on the port specified 
func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("port"))
}
