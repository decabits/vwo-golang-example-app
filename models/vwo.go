package models

import (
	"fmt"

	"github.com/decabits/vwo-golang-example-app/config"
	vwo "github.com/decabits/vwo-golang-sdk"
	"github.com/decabits/vwo-golang-sdk/pkg/api"
)

// GetVWOInstance function initializes the settings file and launches the
// VWO instance using configuration values and returns the insttance
func GetVWOInstance() *api.VWOInstance {
	config := config.GetConfig()

	//User Storage
	//storage := &UserStorageData{}

	//Custom Logger (Google Logger)
	/*
	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()
	*/

	//Instance with userStorage and customLogger with developmentMode as true
	//instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode(), api.WithUserStorage(storage), api.WithCustomLogger(logs))

	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode())
	if err != nil {
		fmt.Println("error intialising sdk")
	}
	return instance
}
