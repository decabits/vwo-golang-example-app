package models

import (
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
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

// var UsersData = map[string][]UserData{
// 	"campaignKey" : [
// 		{
// 			UserID: "user-identifier",
//       CampaignKey: "campaign unique key",
//       VariationName: "Variation-2",
// 		},
// 		{
// 			UserID: "user-identifier",
//       CampaignKey: "campaign unique key",
//       VariationName: "Variation-2",
// 		}
// 	]
// }

// Get function
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	// for _, userData := range config.UserDatas {
	// 	if userData.UserID == userID && userData.CampaignKey == campaignKey {
	// 		return userData
	// 	}
	// }
	return schema.UserData{}
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
