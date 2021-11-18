package main

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
)

func main() {
	// Instantiate the AnnotatorForClinicalDataV1 instance
	authenticator := &core.IamAuthenticator{
		ApiKey: "my-iam-apikey",
	}
	ACD, err := annotatorforclinicaldataacdv1.
	  NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
		  URL:           "my-service-url",
		  Version:       core.StringPtr("2020-09-08"),
		  Authenticator: authenticator,
	  })

	// Check successful instantiation
	if err != nil {
		panic(err)
	}

	// Call the Get Profiles method
	getProfilesOptions := ACD.NewGetProfilesOptions()
	getProfilesResult, getProfilesDetailedResponse, getProfilesErr := ACD.GetProfiles(getProfilesOptions)

        // Check successful call
	if getProfilesErr != nil {
		panic(getProfilesErr)
	}

        // Print the detailed response
	fmt.Println("Get Profiles Response: ", getProfilesDetailedResponse)

	// Print specific feature from Get Profiles result
	if getProfilesResult != nil {
		fmt.Println("List of Profiles by ID")
		for profileId, _ := range getProfilesResult {
			fmt.Println("Profile ID: ", profileId)

			// Call the Get Profile By ID method
			getProfilesByIdOptions := ACD.NewGetProfilesByIdOptions(profileId)
			getProfilesByIdResult, getProfilesByIdDetailedResponse, getProfilesByIdErr := ACD.GetProfilesByID(getProfilesByIdOptions)

			// Check successful call
			if getProfilesByIdErr != nil {
				panic(getProfilesByIdErr)
			}

			// Print the detailed response
			fmt.Println("Get Profiles by ID Response: ", getProfilesByIdDetailedResponse)

			// Print specific feature from Get Profile By ID result
			if getProfilesByIdResult != nil {
				fmt.Println("Profile Name: ", getProfilesByIdResult.Name)
			}
		}
	}

	// Call the Get Health Check Status method
	getHealthCheckStatusOptions := ACD.NewGetHealthCheckStatusOptions()
	getHealthCheckStatusResult, getHealthCheckStatusDetailedResponse, getHealthCheckStatusErr := ACD.GetHealthCheckStatus(getHealthCheckStatusOptions)

	// Check successful call
	if getHealthCheckStatusErr != nil {
		panic(getHealthCheckStatusErr)
	}

	// Print the detailed response
	fmt.Println("Get Health Check Status Response: ", getHealthCheckStatusDetailedResponse)

	// Print specific feature from Get Health Check Status result
	if getHealthCheckStatusResult != nil {
		fmt.Println("Service State: ", getHealthCheckStatusResult.ServiceState)
	}

}
