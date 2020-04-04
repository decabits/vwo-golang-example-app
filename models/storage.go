package models

import (
	"encoding/json"
	"errors"

	"github.com/decabits/vwo-golang-sdk/schema"
	//"github.com/decabits/vwo-golang-example-app/config"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) (schema.UserData, error)
	Set(interface{}) (bool, error)
	Exist() bool
}

// UserStorageData struct
type UserStorageData struct{}

// UsersInfo ...
var UsersInfo = `{
	"campaignKey": [
		{
			"UserID": "user-identifier",
			"CampaignKey": "campaignKey",
			"VariationName": "Variation-2",
		},
		{
			"UserID": "user-identifier",
			"CampaignKey": "campaignKey",
			"VariationName": "Variation-2",
		}
	]
}`

// Get function
/*
{
	struct
}
*/
func (us *UserStorageData) Get(userID, campaignKey string) (schema.UserData, error) {
	var userDatas []schema.UserData
	err := json.Unmarshal([]byte(UsersInfo), &userDatas)
	if err != nil {
		return schema.UserData{}, errors.New("Error: " + err.Error())
	}
	if len(userDatas) == 0 {
		return schema.UserData{}, errors.New("No User Data Found")
	} else if len(userDatas) == 1 {
		return userDatas[0], nil
	} else {
		for _, userdata := range userDatas {
			if userdata.CampaignKey == campaignKey && userdata.UserID == userID {
				return userdata, nil
			}
		}
	}
	return schema.UserData{}, errors.New("No User Data Found")
}

// Set function
func (us *UserStorageData) Set(userData interface{}) (bool, error) {
	userdata, err := json.Marshal(userData)
	if err != nil {
		return false, errors.New("Error: " + err.Error())
	}
	_ = userdata
	return true, nil
}

// Exist function
func (us *UserStorageData) Exist() bool {
	return true
}
