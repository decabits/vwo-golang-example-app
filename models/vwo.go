package models

import (
	vwo "github.com/Piyushhbhutoria/vwo-go-sdk"
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
	// "github.com/decabits/vwo-golang-example-app/config"
)

// VWO struct
type VWO struct {
	vwoInstance schema.VwoInstance
}

//Init ...
func (v *VWO) Init() {
	// config := config.GetConfig()
	storage := UserStorageData{}
	// vwo := vwo.Default(config.GetString("accountID"), config.GetString("SDKKey"), storage)
	VWO := vwo.New("./settings.json", storage)
	v.vwoInstance = VWO
}

//GetVWOInstance ...
func (v *VWO) GetVWOInstance() schema.VwoInstance {
	return v.vwoInstance
}
