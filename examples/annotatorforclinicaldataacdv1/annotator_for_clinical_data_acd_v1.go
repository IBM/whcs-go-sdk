package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
	"io"
	"reflect"
)

func main() {
	// Instantiate the AnnotatorForClinicalDataV1 instance
	authenticator := &core.IamAuthenticator{
		ApiKey: "my-iam-apikey",
	}
	ACD, err := annotatorforclinicaldataacdv1.
	  NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
		  URL:           "my-service-url",
		  Version:       "2020-08-21",
		  Authenticator: authenticator,
	  })

	// Check successful instantiation
	if err != nil {
		panic(err)
	}

	// Call the Get Profiles method
	getProfilesOptions := ACD.NewGetProfilesOptions()
	result, detailedResponse, err := ACD.GetProfiles(getProfilesOptions)

  // Check successful call
	if err != nil {
		panic(err)
	}

  // Print the detailed response
	fmt.Println("Get Profiles Response: ", detailedResponse)

	// Print specific feature from Get Profiles result
	if result != nil {
		fmt.Println("List of Profiles by ID")
		for profileId, _ := range result.Profiles {
			fmt.Println("Profile ID: ", profileId)

			// Call the Get Profile By ID method
			getProfilesByIdOptions := ACD.NewGetProfilesByIdOptions(profileId)
			result, detailedResponse, err := ACD.GetProfilesByID(getProfilesByIdOptions)

			// Check successful call
			if err != nil {
				panic(err)
			}

			// Print the detailed response
			fmt.Println("Get Profiles by ID Response: ", detailedResponse)

			// Print specific feature from Get Profile By ID result
			if result != nil {
				fmt.Println("Profile Name: ", result.Name)
			}
		}
	}

	// Call the Get Health Check Status method
	getHealthCheckStatusOptions := ACD.NewGetHealthCheckStatusOptions()
	result, detailedResponse, err := ACD.GetHealthCheckStatus(getHealthCheckStatusOptions)

	// Check successful call
	if err != nil {
		panic(err)
	}

	// Print the detailed response
	fmt.Println("Get Health Check Status Response: ", detailedResponse)

	// Print specific feature from Get Health Check Status result
	if result != nil {
		fmt.Println("Service State: ", result.ServiceState)
	}

}
