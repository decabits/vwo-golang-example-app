# VWO-Golang-SDK Example

This repository provides a basic demo of how server-side works with VWO Golang SDK.


## Requirements

* Works with Go 1.11 +


## Documentation

Refer [VWO Official Server-side Documentation](https://developers.vwo.com/reference#server-side-introduction)


## Setup

1. Install dependencies

```go
go get "github.com/decabits/vwo-golang-sdk"
go get "github.com/decabits/vwo-golang-example-app"

```

2. Update your app with your settings present in `config/vwo.json`

```json
"accountID": "REPLACE_THIS_WITH_CORRECT_VALUE"
"SDKKey": "REPLACE_THIS_WITH_CORRECT_VALUE"
"abCampaignKey": "REPLACE_THIS_WITH_CORRECT_VALUE"
"abCampaignGoalIdentifeir": "REPLACE_THIS_WITH_CORRECT_VALUE"
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
VWO := vwo.Default(config.GetString("accountID"), config.GetString("SDKKey"), storage)

// Get Settings
settingsFile = service.SettingsFileManager.GetSettingsFile()

// Get Instance
// path of the required settings file is passed as an arguement 
vwoInstance = api.GetInstance(path)

// Activate API
variationName = api.Activate(vwoInstance, campaignKey, userID)

// GetVariation
options = {}
variationName = api.GetVariationName(vwoInstance, campaignKey, userID, options)

// Track API
options = {}
isSuccessful = api.TrackWithOptions(vwoInstance, campaignKey, userID, goalIdentifier, options)

// FeatureEnabled API
options = {}
isSuccessful = api.IsFeatureEnabled(vwoInstance, campaignKey, userID, options)

// GetFeatureVariableValue API
options = {}
variableValue = api.GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)

// Push API
isSuccessful = api.Push(tagKey, tagValue, userID)
```

**API usage**

**User Define Logger**


## License


