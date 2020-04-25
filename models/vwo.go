package models

import (
	"log"

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
	storage := &UserStorageData{}
	settingsFile, err := vwo.GetSettingsFile(config.GetString("accountID"), config.GetString("SDKKey"))
	if err != nil {
		log.Fatal("Unable to fetch settingsFile: ", err)
	}
	v.vwoInstance = vwo.VWOInstance{}
	v.vwoInstance.Launch(config.GetBool("isDevelopmentMode"), settingsFile, storage)
}

// GetVWOInstance ...
func (v *VWO) GetVWOInstance() vwo.VWOInstance {
	return v.vwoInstance
}
