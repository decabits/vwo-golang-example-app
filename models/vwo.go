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

// Init function initializes the settings file and launches the VWO instance using configuration values
func (v *VWO) Init() {
	config := config.GetConfig()
	// storage := &UserStorageData{}
	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	v.vwoInstance = vwo.VWOInstance{}
	err := v.vwoInstance.Launch(config.GetBool("isDevelopmentMode"), settingsFile, nil, nil)
	if err != nil {
		fmt.Println("error intialising sdk")
	}
}

// GetVWOInstance function returns VWO instance
func (v *VWO) GetVWOInstance() vwo.VWOInstance {
	return v.vwoInstance
}
