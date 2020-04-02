package models

import (
	"github.com/decabits/vwo-golang-example-app/config"
	vwo "github.com/decabits/vwo-golang-sdk"
	"github.com/decabits/vwo-golang-sdk/schema"
)

// VWO struct
type VWO struct {
	vwoInstance schema.VwoInstance
}

//Init ...
func (v *VWO) Init() {
	config := config.GetConfig()
	storage := &UserStorageData{}
	VWO := vwo.Default(config.GetString("accountID"), config.GetString("SDKKey"), storage)
	// VWO := vwo.New("./settings.json", storage)
	v.vwoInstance = VWO
}

//GetVWOInstance ...
func (v *VWO) GetVWOInstance() schema.VwoInstance {
	return v.vwoInstance
}
