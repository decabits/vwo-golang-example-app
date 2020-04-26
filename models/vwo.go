package models

import (
	"fmt"

	"github.com/decabits/vwo-golang-example-app/config"
	vwo "github.com/decabits/vwo-golang-sdk"
)

// VWO struct
type VWO struct {
	vwoInstance vwo.VWOInstance
}

// GetVWOInstance function initializes the settings file and launches the
// VWO instance using configuration values and returns the insttance
func GetVWOInstance() vwo.VWOInstance {
	config := config.GetConfig()
	// storage := &UserStorageData{}
	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	instance := vwo.VWOInstance{}
	err := instance.Launch(config.GetBool("isDevelopmentMode"), settingsFile, nil, nil)
	if err != nil {
		fmt.Println("error intialising sdk")
	}
	return instance
}
