# VWO-Golang-SDK Example

This repository provides a basic demo of how server-side works with VWO Golang SDK.

## Requirements

- Works with Go 1.1x

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
import "github.com/decabits/vwo-golang-sdk/pkg/api"

// Get SettingsFile
settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")

// Default instance of vwoInstance
instance, err := vwo.Launch(settingsFile)
if err != nil {
	//handle err
}

// Instance with custom options
instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode())
if err != nil {
	//handle err
}

// Activate API
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
options["variationTargetingVariables"] = map[string]interface{}{}
options["revenueGoal"] = 12
variationName = vwoInstance.Activate(campaignKey, userID, options)

//Without Custom Variables
variationName = instance.Activate(campaignKey, userID, nil)

// GetVariation
variationName = instance.GetVariationName(campaignKey, userID, nil)

// Track API
options := make(map[string]interface{})
options["revenueGoal"] = 12
isSuccessful = instance.Track(campaignKey, userID, goalIdentifier, options)

// FeatureEnabled API
isSuccessful = instance.IsFeatureEnabled(campaignKey, userID, nil)


// GetFeatureVariableValue API
variableValue = instance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)

// Push API
isSuccessful = instance.Push(tagKey, tagValue, userID)
```

**User Storage**

```go
import vwo "github.com/decabits/vwo-golang-sdk/"
import "github.com/decabits/vwo-golang-sdk/pkg/api"
import "github.com/decabits/vwo-golang-sdk/pkg/schema"

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

	instance, err := vwo.Launch(settingsFile, api.WithUserStorage(storage))
	if err != nil {
		//handle err
	}
}

```

**Custom Logger**

```go
import vwo "github.com/decabits/vwo-golang-sdk"
import "github.com/decabits/vwo-golang-sdk/pkg/api"

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

	instance, err := vwo.Launch(settingsFile, api.WithCustomLogger(logger))
	if err != nil {
		//handle err
	}
}
```

## Third-party Resources and Credits

Refer [third-party-attributions.txt](third-party-attributions.txt)

## License

[Apache License, Version 2.0](LICENSE)
