# VWO-Golang-SDK Example

This repository provides a basic demo of how server-side works with VWO Golang SDK.


## Requirements

* Works with Go 1.11 +


## Documentation

Refer [VWO Official Server-side Documentation](https://developers.vwo.com/reference#server-side-introduction)


## Setup

1. Install dependencies

```go
go get .

```

2. Update your app with your settings present in `config/dev.yaml`

```yaml
accountID: "REPLACE_THIS_WITH_CORRECT_VALUE"
SDKKey: "REPLACE_THIS_WITH_CORRECT_VALUE"
abCampaignKey: "REPLACE_THIS_WITH_CORRECT_VALUE"
abCampaignGoalIdentifeir: "REPLACE_THIS_WITH_CORRECT_VALUE"
```

3. Run application

```go
go run main.go dev
```


## Basic Usage

**Importing and Instantiation**

```go
import (
	vwo "github.com/decabits/vwo-golang-sdk"
    "github.com/decabits/vwo-golang-sdk/schema"
    "github.com/decabits/vwo-golang-sdk/api"
)

// Get SettingsFile
settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")

// Get UserStorage 
storage := &UserStorageData{}

// Initialize VwoInstance
vwoInstance = vwo.VWOInstance{}

//Create VwoInstance
vwoInstance.Launch("isDevelopmentMode", settingsFile, storage)

// Activate API
variationName = vwoInstance.Activate(VWO, campaignKey, userID)

// GetVariation
options = schema.Options{}
variationName = vwoInstance.GetVariationName(VWO, campaignKey, userID, options)

// Track API
options = schema.Options{}
isSuccessful = vwoInstance.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)

// Push API
isSuccessful = vwoInstance.Push(tagKey, tagValue, userID)
```


**User Storage**

```go

import "github.com/decabits/vwo-golang-sdk/schema"

// UserStorage interface
type UserStorage schema.UserStorage
/*
// UserStorage struct
type UserStorage interface {
	Get(userID, campaignKey string) UserData
	Set(string, string, string)
	Exist() bool
}
*/

// UserStorageData struct
type UserStorageData struct{}

func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	
    //Example code showing how to get userData from DB
    userData, ok := userDatas[campaignKey]
	if ok {
		for _, userdata := range userData {
			if userdata.UserID == userID {
				return userdata
			}
		}
	}
	/*
    // UserData  struct
    type UserData struct {
        UserID        string
        CampaignKey   string
        VariationName string
    }
    */
	return schema.UserData{}
}

func (us *UserStorageData) Set(userID, campaignKey, variationName string) {

    //Example code showing how to store userData in DB
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
}

// Exist function
func (us *UserStorageData) Exist() bool {
	// Set the return value true in case there is a user storage else false
	return true
}
```

## Third-party Resources and Credits

Refer [third-party-attributions.txt](third-party-attribution.txt)

## License

[Apache License, Version 2.0](LICENSE)

