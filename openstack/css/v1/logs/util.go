package logs

import (
	"fmt"
	"log"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func WaitForBaseLogStatus(client *golangsdk.ServiceClient, id string, timeout int) error {
	time.Sleep(5 * time.Second)
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		got, err := GetBaseLogConfiguration(client, id)
		if err != nil {
			if _, ok := err.(golangsdk.BaseError); ok {
				return true, err
			}
			log.Printf("Error while waiting for base logging status: %s", err)
			return false, nil
		}
		m, ok := got.(*BaseLogConfiguration)
		if !ok {
			log.Fatal("An interface was expected!")
		}

		return false, fmt.Errorf("base log switch status: %v", m.LogSwitch)
	})
}
