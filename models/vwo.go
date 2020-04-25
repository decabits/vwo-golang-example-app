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

// Init ...
func (v *VWO) Init() {
	config := config.GetConfig()
	// storage := &UserStorageData{}
	settingsFile := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	v.vwoInstance = vwo.VWOInstance{}
	err := v.vwoInstance.Launch(config.GetBool("isDevelopmentMode"), settingsFile, nil)
	if err != nil {
		fmt.Println("error intialising sdk")
	}
}

// GetVWOInstance ...
func (v *VWO) GetVWOInstance() vwo.VWOInstance {
	return v.vwoInstance
}
