package models

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	//"github.com/decabits/vwo-golang-example-app/config"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userData schema.UserData) bool
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
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	var userDatas map[string][]schema.UserData
	if len(userDatas) == 0 {
		return schema.UserData{}
	}
	userData, ok := userDatas[campaignKey]
	if ok {
		for _, userdata := range userData {
			if userdata.UserID == userID {
				return userdata
			}
		}
	}
	return schema.UserData{}
}

// Set function
func (us *UserStorageData) Set(userData schema.UserData) bool {
	
	return true
}

// Exist function
func (us *UserStorageData) Exist() bool {
	return true
}
