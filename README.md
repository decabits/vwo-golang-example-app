# VWO-Golang-SDK Example

This repository provides a basic demo of how server-side works with VWO Golang SDK.

## Requirements

- Works with Go 1.11+

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
isDevelopmentMode: bool
```

3. Run application

```go
go run main.go dev
```

## Basic Usage

**Importing and Instantiation**

```go
import vwo "github.com/decabits/vwo-golang-sdk"

// Get SettingsFile
settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")

// Declaration of VwoInstance
vwoInstance = vwo.VWOInstance{}

// Create VwoInstance and handle error if any
err := vwoInstance.Launch("isDevelopmentMode", settingsFile, nil, nil)

// Activate API
variationName = vwoInstance.Activate(campaignKey, userID, nil)

// GetVariation
variationName = vwoInstance.GetVariationName(campaignKey, userID, nil)

// Track API
options := make(map[string]interface{})
options["revenueGoal"] = 12
isSuccessful = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)

// FeatureEnabled API
isSuccessful = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)


// GetFeatureVariableValue API
variableValue = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)

// Push API
isSuccessful = vwoInstance.Push(tagKey, tagValue, userID)
```

**User Storage**

```go
import "github.com/decabits/vwo-golang-sdk/schema"

// declare UserStorage interface with the following Get & Set function signature
type UserStorage interface{
    Get(userID, campaignKey string) UserData
    Set(string, string, string)
}

// declare a UserStorageData struct to implement UserStorage interface
type UserStorageData struct{}

// Get method to fetch user variation from storage
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

// Set method to save user variation to storage
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

func main() {
	settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")
	// create UserStorageData object
	storage := &UserStorageData{}
	v.vwoInstance = vwo.VWOInstance{}
	err := v.vwoInstance.Launch(config.GetBool("isDevelopmentMode"), settingsFile, storage)
	if err != nil {
		fmt.Println("error intialising sdk")
	}
}

```

**Custom Logger**

```go
import vwo "github.com/decabits/vwo-golang-sdk"

// declare Log interface with the following CustomLog function signature
type Log interface {
	CustomLog(level, errorMessage string)
}

// declare a LogS struct to implement Log interface
type LogS struct{}

// Get function to handle logs
func (c *LogS) CustomLog(level, errorMessage string) {}

func main() {
	settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")
	// create LogS object
	logger := &LogS{}
	v.vwoInstance = vwo.VWOInstance{}
	err := v.vwoInstance.LaunchWithLogger(config.GetBool("isDevelopmentMode"), settingsFile, nil, logger)
	if err != nil {
		fmt.Println("error intialising sdk")
	}
}
```

## License

[Apache License, Version 2.0](LICENSE)
