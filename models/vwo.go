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

	// User Storage
	// storage := &UserStorageData{}

	// Custom Logger (Google Logger)
	// logs := &LogS{}

	// Instance with userStorage and customLogger with developmentMode as true
	// instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode(), api.WithStorage(storage), api.WithLogger(logs))

	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode())
	if err != nil {
		fmt.Println("error intialising sdk")
	}
	vwo.SetLogLevel(3)
	return instance
}
