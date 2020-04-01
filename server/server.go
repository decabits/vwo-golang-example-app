package server

import "github.com/decabits/vwo-golang-example-app/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("port"))
}
