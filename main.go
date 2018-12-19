package main

import (
	"fmt"
	"os"

	ttnsdk "github.com/TheThingsNetwork/go-app-sdk"
	ttnlog "github.com/TheThingsNetwork/go-utils/log"
	"github.com/TheThingsNetwork/go-utils/log/apex"
)

const (
	sdkClientName = "sdk-test"
)

func main() {
	log := apex.Stdout() // We use a cli logger at Stdout
	log.MustParseLevel("debug")
	ttnlog.Set(log) // Set the logger as default for TTN

	appID := os.Getenv("TTN_APP_ID")
	appAccessKey := os.Getenv("TTN_APP_ACCESS_KEY")

	config := ttnsdk.NewCommunityConfig(sdkClientName)
	config.ClientVersion = "2.0.5"

	client := config.NewClient(appID, appAccessKey)
	defer client.Close()

	devices, err := client.ManageDevices()
	if err != nil {
		log.WithError(err).Fatal("my-amazing-app: could not get device manager")
	}
	deviceList, err := devices.List(10, 0)
	if err != nil {
		log.WithError(err).Fatal("my-amazing-app: could not get devices")
	}
	log.Info("my-amazing-app: found devices")
	for _, device := range deviceList {
		fmt.Printf("- %s", device.DevID)
	}

}
