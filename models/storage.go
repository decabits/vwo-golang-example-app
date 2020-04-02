package models

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	//"github.com/decabits/vwo-golang-example-app/config"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userStorageData interface{})
	Exist() bool
}

// UserStorageData struct
type UserStorageData struct{}

var UsersDatas = `{
	"campaignKey": [
		{
			"UserID": "user-identifier",
			"CampaignKey": "campaign unique key",
			"VariationName": "Variation-2",
		},
		{
			"UserID": "user-identifier",
			"CampaignKey": "campaign unique key",
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
	// var userDatas schema.UserDatas
	// if err := json.Unmarshal([]byte(UsersDatas), &userDatas); err != nil {
	// 	return schema.UserData{}, errors.New("Error: " + err.Error())
	// }
	// for _, userData := range userDatas {
	// 	if userData.UserID == userID && userData.CampaignKey == campaignKey {
	// 		return userData, nil
	// 	}
	// }
	return schema.UserData{}, nil
}

// Set function
func (us *UserStorageData) Set(userStorageData interface{}) bool {
	// if (userIDs.indexOf(userStorageData.userId) === -1) {
	//   userData[userStorageData.campaignKey] = userData[userStorageData.campaignKey] || [];
	//   userData[userStorageData.campaignKey].push(userStorageData);

	//   userIds.push(userStorageData.userId);
	// }
	return true
}

// Exist function
func (us *UserStorageData) Exist() bool {
	return true
}
