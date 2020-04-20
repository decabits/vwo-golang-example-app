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
    "github.com/decabits/vwo-golang-sdk/service"
)

// Initialize client
// storage should be of type schema.UserStorage
VWO := vwo.Default("accountID", "SDKKey", storage)

// Activate API
variationName = api.Activate(VWO, campaignKey, userID)

// GetVariation
options = {}
variationName = api.GetVariationName(VWO, campaignKey, userID, options)

// Track API
options = {}
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)

// Push API
isSuccessful = api.Push(tagKey, tagValue, userID)
```

**API usage**

** User Storage **

```go

import "github.com/decabits/vwo-golang-sdk/schema"

// UserStorage interface
type UserStorage schema.UserStorage

// UserStorageData struct
type UserStorageData struct{}

func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
    
    //Example code showing how to get userData  from DB
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
```


## License

[Apache License, Version 2.0](LICENSE)

