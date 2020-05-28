/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package annotatorforclinicaldataacdv1_test

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "acd-service.env"

var (
	ACD          *annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1
	err          error
	configLoaded bool = false
	version      string
	serviceUrl   string
	iamUrl       string
	apikey       string
	disableSsl   bool   = false
	analyzeText  string = `The patient has cancer and is currently taking 400 ml sisplatin chemotherapy.\nHISTORY:  Patient is allergic to latex.  Patient cannot walk and needs help bathing and getting around.  The lab values were: white blood cell count 4.6, hemoglobin 12.2.  Echocardiogram demonstrated ejection fraction of approx 60%.  Patient cannot dress or feed without help as the patient can not see.  Patient may die soon but has not died yet.  Patient smoked for 20 years.  Patient can not clean up after defacating in toilet.  Jone Doe was seen at Baylor Hospitall in Austin, TX.  Johndoe@testaddress.com - (555) 555-5555'`
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration not found, skipping....")
	}
}

var _ = Describe(`AnnotatorForClinicalDataAcdV1`, func() {
	It("Successfully load the configuration", func() {
		err = godotenv.Load(configFile)
		if err == nil {
			serviceUrl = os.Getenv("URL")
			apikey = os.Getenv("APIKEY")
			version = os.Getenv("VERSION")
			configLoaded = true
		}

		if !configLoaded {
			Skip("External configuration not found, skipping....")
		}
	})

	It("Successfully create the client instance", func() {
		shouldSkipTest()

		if apikey == "" {
			ACD, err = annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:           serviceUrl,
				Version:       core.StringPtr(version),
				Authenticator: &core.NoAuthAuthenticator{},
			})
			ACD.Service.DisableSSLVerification()
		} else {
			ACD, err = annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     serviceUrl,
				Version: core.StringPtr(version),
				Authenticator: &core.IamAuthenticator{
					ApiKey:                 apikey,
					URL:                    iamUrl,
					DisableSSLVerification: disableSsl,
				},
			})
		}
		Expect(err).To(BeNil())
		//		Expect(err).ToNot(BeNil())
		Expect(ACD.Service.Options.URL).To(Not(Equal("")))
	})

	Describe(`GetProfiles(getProfilesOptions *GetProfilesOptions) `, func() {
		It(`Successfully retrieve all profiles`, func() {
			getProfilesOptions := ACD.NewGetProfilesOptions()
			result, detailedResponse, err := ACD.GetProfiles(getProfilesOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetProfilesById(getProfilesByIdOptions *GetProfilesByIdOptions) `, func() {
		It(`Successfully retrieve mathcing profile`, func() {
			getProfilesByIdOptions := ACD.NewGetProfilesByIdOptions("clinical_insights_cartridge_v1.0_profile")
			result, detailedResponse, err := ACD.GetProfilesByID(getProfilesByIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ID).ToNot(BeNil())
			Expect(result.Name).ToNot(BeNil())
			Expect(result.Annotators).ToNot(BeNil())
			annotatorCounter := 0
			for annotatorCounter < len(result.Annotators) {
				annotator := result.Annotators[annotatorCounter]
				Expect(annotator.Name).ToNot(BeNil())
				annotatorCounter++
			}
		})
	})

	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
		It(`Successfully create a profile`, func() {
			createProfilesOptions := ACD.NewCreateProfileOptions()
			createProfilesOptions.SetID("testProfile")
			createProfilesOptions.SetName("testProfile")
			createProfilesOptions.SetDescription("testProfile")
			createProfilesOptions.SetPublishedDate("2020-01-01")
			createProfilesOptions.SetPublish(false)
			createProfilesOptions.SetVersion("1.0")
			createProfilesOptions.SetCartridgeID("testCartridge")
			annotator, err := ACD.NewAnnotator("concept_detection")
			createProfilesOptions.SetAnnotators([]annotatorforclinicaldataacdv1.Annotator{*annotator})
			detailedResponse, err := ACD.CreateProfile(createProfilesOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
		})
	})

	Describe(`UpdateeProfile(updateProfileOptions *UpdateProfileOptions)`, func() {
		It(`Successfully update a profile`, func() {
			updateProfilesOptions := ACD.NewUpdateProfileOptions("testProfile")
			updateProfilesOptions.SetNewID("testProfile")
			updateProfilesOptions.SetNewDescription("testProfile updated")
			updateProfilesOptions.SetNewPublishedDate("2020-01-01")
			updateProfilesOptions.SetNewPublish(false)
			updateProfilesOptions.SetNewVersion("1.1")
			annotator, err := ACD.NewAnnotator("concept_value")
			updateProfilesOptions.SetNewAnnotators([]annotatorforclinicaldataacdv1.Annotator{*annotator})
			detailedResponse, err := ACD.UpdateProfile(updateProfilesOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {
		It(`Successfully delete a profile`, func() {
			deleteProfileOptions := ACD.NewDeleteProfileOptions("testProfile")
			detailedResponse, err := ACD.DeleteProfile(deleteProfileOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetFlows(getFlowsOptions *GetFlowsOptions) `, func() {
		It(`Successfully retrieve all flows`, func() {
			getFlowsOptions := ACD.NewGetFlowsOptions()
			result, detailedResponse, err := ACD.GetFlows(getFlowsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetFlowsById(getFlowsByIdOptions *GetFlowsByIdOptions) `, func() {
		It(`Successfully retrieve mathcing flow`, func() {
			getFlowsByIdOptions := ACD.NewGetFlowsByIdOptions("clinical_insights_cartridge_v5.0_default_flow")
			result, detailedResponse, err := ACD.GetFlowsByID(getFlowsByIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ID).ToNot(BeNil())
			Expect(result.Name).ToNot(BeNil())
			Expect(result.AnnotatorFlows).ToNot(BeNil())
			annotatorCounter := 0
			for annotatorCounter < len(result.AnnotatorFlows) {
				annotator := result.AnnotatorFlows[annotatorCounter]
				Expect(annotator.Profile).ToNot(BeNil())
				annotatorCounter++
			}
		})
	})
	Describe(`CreateFlows(createFlowsOptions *CreateFlowsOptions)`, func() {
		It(`Successfully create a flow`, func() {
			createFlowsOptions := ACD.NewCreateFlowsOptions()
			createFlowsOptions.SetID("testFlow")
			createFlowsOptions.SetName("testFlow")
			createFlowsOptions.SetDescription("testFlow")
			createFlowsOptions.SetPublishedDate("2020-01-01")
			createFlowsOptions.SetPublish(false)
			createFlowsOptions.SetVersion("1.0")
			createFlowsOptions.SetCartridgeID("testCartridge")
			annotator, err := ACD.NewAnnotator("concept_detection")
			flowEntry, err := ACD.NewFlowEntry(annotator)
			flowEntry.Annotator = annotator
			async := false
			flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntry}, core.BoolPtr(async))
			annotatorFlow, err := ACD.NewAnnotatorFlow(flow)
			createFlowsOptions.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
			detailedResponse, err := ACD.CreateFlows(createFlowsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
		})
	})

	Describe(`UpdateFlows(updateFlowsOptions *UpdateFlowsOptions)`, func() {
		It(`Successfully update a flow`, func() {
			updateFlowsOptions := ACD.NewUpdateFlowsOptions("testFlow")
			updateFlowsOptions.SetNewID("testFlow")
			updateFlowsOptions.SetNewDescription("testFlow updated")
			updateFlowsOptions.SetNewPublishedDate("2020-01-01")
			updateFlowsOptions.SetNewPublish(false)
			updateFlowsOptions.SetNewVersion("1.0")
			annotator, err := ACD.NewAnnotator("concept_detection")
			flowEntry, err := ACD.NewFlowEntry(annotator)
			flowEntry.Annotator = annotator
			async := false
			flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntry}, core.BoolPtr(async))
			annotatorFlow, err := ACD.NewAnnotatorFlow(flow)
			updateFlowsOptions.SetNewAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
			detailedResponse, err := ACD.UpdateFlows(updateFlowsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`DeleteFlows(deleteFlowsOptions *DeleteFlowsOptions)`, func() {
		It(`Successfully delete a flow`, func() {
			deleteFlowOptions := ACD.NewDeleteFlowsOptions("testFlow")
			detailedResponse, err := ACD.DeleteFlows(deleteFlowOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetAnnotators(getAnnotatorsOptions *GetAnnotatorsOptions) `, func() {
		It(`Successfully retrieve all flows`, func() {
			getAnnotatorsOptions := ACD.NewGetAnnotatorsOptions()
			detailedResponse, err := ACD.GetAnnotators(getAnnotatorsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetAnnotatorsById(getAnnotatorssByIdOptions *GetAnnotaotrssByIdOptions) `, func() {
		It(`Successfully retrieve mathcing flow`, func() {
			getAnnotatorsByIdOptions := ACD.NewGetAnnotatorsByIdOptions("concept_detection")
			detailedResponse, err := ACD.GetAnnotatorsByID(getAnnotatorsByIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions)`, func() {
		It(`Successfully retrieve service status`, func() {
			getHealthCheckStatusOptions := ACD.NewGetHealthCheckStatusOptions()
			result, detailedResponse, err := ACD.GetHealthCheckStatus(getHealthCheckStatusOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result.ServiceState).ToNot(BeNil())
		})
	})

	Describe(`RunPipeline(runPipelineOptions *RunPipelineOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineOptions()
			annotator, err := ACD.NewAnnotator("concept_detection")
			flowEntry, err := ACD.NewFlowEntry(annotator)
			flowEntry.Annotator = annotator
			async := false
			flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntry}, core.BoolPtr(async))
			annotatorFlow, err := ACD.NewAnnotatorFlow(flow)
			pipelineOptions.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
			container := ACD.NewUnstructuredContainer()
			container.SetText(analyzeText)
			pipelineOptions.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*container})
			result, detailedResponse, err := ACD.RunPipeline(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Data).ToNot(BeNil())
		})
	})

	Describe(`RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineWithFlowOptions("clinical_insights_v1.0_standard_flow", false)
			pipelineOptions.SetBody(analyzeText)
			pipelineOptions.SetContentType("text/plain")
			result, detailedResponse, err := ACD.RunPipelineWithFlow(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Data).ToNot(BeNil())
		})
	})

	Describe(`DeleteUserSpecificArtifacts(deleteUserSpecificArtifactsOptions *DeleteUserSpecificArtifactsOptions)`, func() {
		It(`Successfully delete user specific artifacts`, func() {
			artifactsOptions := ACD.NewDeleteUserSpecificArtifactsOptions()
			detailedResponse, err := ACD.DeleteUserSpecificArtifacts(artifactsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
		})
	})

	Describe(`CartridgesGet(cartridgesGetOptions *CartridgesGetOptions)`, func() {
		It(`Successfully retrieve cartridges`, func() {
			cartridgesGetOptions := ACD.NewCartridgesGetOptions()
			result, detailedResponse, err := ACD.CartridgesGet(cartridgesGetOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`CartridgesGetID(cartridgesGetIdOptions *CartridgesGetIdOptions)`, func() {
		It(`Successfully retrieve the cartridge`, func() {
			cartridgesGetIdOptions := ACD.NewCartridgesGetIdOptions("clinical_insights_cartridge_v1.0")
			result, detailedResponse, err := ACD.CartridgesGetID(cartridgesGetIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ID).ToNot(BeNil())
			Expect(result.Status).ToNot(BeNil())
			Expect(result.StatusLocation).ToNot(BeNil())
			Expect(result.StartTime).ToNot(BeNil())
			Expect(result.EndTime).ToNot(BeNil())
			Expect(result.Duration).ToNot(BeNil())
			Expect(result.CorrelationID).ToNot(BeNil())
			//			Expect(result.ArtifactResponseCode).To(Equal(200)) nil for some reason
			responseCounter := 0
			responses := len(result.ArtifactResponse)
			for responseCounter < responses {
				response := result.ArtifactResponse[responseCounter]
				Expect(response.Code).ToNot(BeNil())
				Expect(response.Message).ToNot(BeNil())
				Expect(response.Level).ToNot(BeNil())
				responseCounter++
			}
			//			ArtifactResponse []ServiceError `json:"artifactResponse,omitempty"`
		})
	})

	Describe(`CartridgesPostMultipart(cartridgesPostMultipartOptions *CartridgesPostMultipartOptions) `, func() {
		It(`Successfully start a cartridge archive deployment`, func() {
			cartridgesPostMultipartOptions := ACD.NewCartridgesPostMultipartOptions()
			cartridgesPostMultipartOptions.SetArchiveFileContentType("application/octet-stream")
			cartridgesPostMultipartOptions.SetArchiveFile(ioutil.NopCloser(strings.NewReader(os.Getenv("ARCHIVE"))))
			result, detailedResponse, err := ACD.CartridgesPostMultipart(cartridgesPostMultipartOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`CartridgesPutMultipart(cartridgesPutMultipartOptions *CartridgesPutMultipartOptions) `, func() {
		It(`Successfully start a cartridge archive deployment`, func() {
			cartridgesPutMultipartOptions := ACD.NewCartridgesPutMultipartOptions()
			cartridgesPutMultipartOptions.SetArchiveFileContentType("application/octet-stream")
			cartridgesPutMultipartOptions.SetArchiveFile(ioutil.NopCloser(strings.NewReader(os.Getenv("ARCHIVE"))))
			result, detailedResponse, err := ACD.CartridgesPutMultipart(cartridgesPutMultipartOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`DeployCartridge(deployCartridgeOptions *DeployCartridgeOptions) `, func() {
		It(`Successfully start a cartridge archive deployment`, func() {
			deployCartridgeOptions := ACD.NewDeployCartridgeOptions()
			deployCartridgeOptions.SetArchiveFileContentType("application/octet-stream")
			deployCartridgeOptions.SetArchiveFile(ioutil.NopCloser(strings.NewReader(os.Getenv("ARCHIVE"))))
			deployCartridgeOptions.SetUpdate(true)
			result, detailedResponse, err := ACD.DeployCartridge(deployCartridgeOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})
})
