package models

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

// UserStorage interface
type UserStorage interface{
	Get(userID, campaignKey string) schema.UserData
	Set(userID, campaignKey, variationName string)
}

// UserStorageData struct
type UserStorageData struct{}

// data is an example of how data is stored
var data = `{
    "abc": [{
            "UserID": "user1",
            "CampaignKey": "abc",
            "VariationName": "Control"
        },
        {
            "UserID": "user2",
            "CampaignKey": "abc",
            "VariationName": "Variation-1"
        }
    ]
}`

// Get function is used to get the data from user storage
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	var userDatas map[string][]schema.UserData
	// Conect your database here to fetch the current data
	// Uncomment the below part to user JSON as data base
	// if err := json.Unmarshal([]byte(data), &userDatas); err != nil {
	// 	fmt.Print("Could not unmarshall")
	// }
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
func (us *UserStorageData) Set(userID, campaignKey, variationName string) {
	var userDatas map[string][]schema.UserData
	// Conect your database here to insert the value
	// Uncomment the below part to user JSON as data base
	// if err := json.Unmarshal([]byte(data), &userDatas); err != nil {
	// 	fmt.Print("Could not unmarshall")
	// }
	userdata := schema.UserData{
		UserID:        userID,
		CampaignKey:   campaignKey,
		VariationName: variationName,
	}
	flag := false
	userData, ok := userDatas[userdata.CampaignKey]
	if ok {
		for _, user := range userData {
			if user.UserID == userdata.UserID {
				flag = true
			}
		}
		if !flag {
			userDatas[userdata.CampaignKey] = append(userDatas[userdata.CampaignKey], userdata)
		}
	} else {
		userDatas[userdata.CampaignKey] = []schema.UserData{
			userdata,
		}
	}
	// This is a part of the above JSON data handling
	// data, _ := json.Marshal(userDatas)
	// fmt.Println(string(data))
}
