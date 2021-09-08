// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
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
	"os"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "ibm-credentials.env"

var (
	ACD          *annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1
	err          error
	configLoaded bool = false
	version      string
	serviceUrl   string
	iamUrl       string
	apikey       string
	disableSsl   bool   = false
	config       map[string]string
	analyzeText  string = "The patient has cancer and patient is currently taking 400 ml sisplatin chemotherapy.  Aspirin from once daily to twice daily.\nHISTORY:  Patient is allergic to latex.  Patient cannot walk and needs help bathing and getting around.  The lab values were: white blood cell count 4.6, hemoglobin 12.2.  Echocardiogram demonstrated ejection fraction of approx 60%.  Patient cannot dress or feed without help as the patient can not see.  Patient may die soon but has not died yet.  Patient smoked for 20 years.  Patient can not clean up after defacating in toilet.  Jone Doe was seen at Baylor Hospitall in Austin, TX.  Johndoe@testaddress.com - (555) 555-5555.  The patient started on metformin because his blood sugar was too high.  She had gallbladder removal September 19 2020"


)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration not found, skipping....")
	}
}

var _ = Describe(`AnnotatorForClinicalDataAcdV1`, func() {
	It("Successfully load the configuration", func() {
		_, err = os.Stat(configFile)
		if err != nil {
			Skip("External configuration not found, skipping....")
		}

		os.Setenv("IBM_CREDENTIALS_FILE", configFile)
		config, err = core.GetServiceProperties("ACD_SERVICE")
		if err != nil {
			Skip("External configuration not found, skipping....")
		}

		serviceUrl = config["URL"]
		apikey = config["APIKEY"]
		version = config["VERSION"]
		configLoaded = true

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
			for profileId, _ := range result {
				Expect(profileId).ToNot(BeNil())
			}
		})
	})

	Describe(`GetProfilesById(getProfilesByIdOptions *GetProfilesByIdOptions) `, func() {
		It(`Successfully retrieve matching profile`, func() {
			getProfilesByIdOptions := ACD.NewGetProfilesByIdOptions("wh_acd.ibm_clinical_insights_v1.0_profile")
			result, detailedResponse, err := ACD.GetProfilesByID(getProfilesByIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ID).ToNot(BeNil())
			Expect(result.Name).ToNot(BeNil())
			Expect(result.Annotators).ToNot(BeNil())
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

	Describe(`UpdateProfile(updateProfileOptions *UpdateProfileOptions)`, func() {
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
			for flowId, _ := range result {
				Expect(flowId).ToNot(BeNil())
			}
		})
	})

	Describe(`GetFlowsById(getFlowsByIdOptions *GetFlowsByIdOptions) `, func() {
		It(`Successfully retrieve matching flow`, func() {
			getFlowsByIdOptions := ACD.NewGetFlowsByIdOptions("wh_acd.ibm_clinical_insights_v1.0_standard_flow")
			result, detailedResponse, err := ACD.GetFlowsByID(getFlowsByIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ID).ToNot(BeNil())
			Expect(result.Name).ToNot(BeNil())
			Expect(result.AnnotatorFlows).ToNot(BeNil())
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
			result, detailedResponse, err := ACD.GetAnnotators(getAnnotatorsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			for annotatorId, _ := range result {
				Expect(annotatorId).ToNot(BeNil())
			}
		})
	})

	Describe(`GetAnnotatorsById(getAnnotatorsByIdOptions *GetAnnotatorsByIdOptions) `, func() {
		It(`Successfully retrieve matching flow`, func() {
			getAnnotatorsByIdOptions := ACD.NewGetAnnotatorsByIdOptions("concept_detection")
			result, detailedResponse, err := ACD.GetAnnotatorsByID(getAnnotatorsByIdOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Description).ToNot(BeNil())
		})
	})

	Describe(`GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions)`, func() {
		It(`Successfully retrieve service status`, func() {
			getHealthCheckStatusOptions := ACD.NewGetHealthCheckStatusOptions()
			result, detailedResponse, err := ACD.GetHealthCheckStatus(getHealthCheckStatusOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ServiceState).ToNot(BeNil())
		})
	})

	Describe(`RunPipeline(runPipelineOptions *RunPipelineOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineOptions()
			cdParams := make(map[string][]string)
			cdParamValue := []string{"true"}
			cdParams["apply_spell_check"] = cdParamValue
			cdAnnotator, err := ACD.NewAnnotator("concept_detection")
			cdAnnotator.Parameters = cdParams
			cdFlowEntry, err := ACD.NewFlowEntry(cdAnnotator)
			cdFlowEntry.Annotator = cdAnnotator
			cvAnnotator, err := ACD.NewAnnotator("concept_value")
			cvFlowEntry, err := ACD.NewFlowEntry(cvAnnotator)
			cvFlowEntry.Annotator = cvAnnotator
			cancerAnnotator, err := ACD.NewAnnotator("cancer")
			cancerFlowEntry, err := ACD.NewFlowEntry(cancerAnnotator)
			cancerFlowEntry.Annotator = cancerAnnotator
			efAnnotator, err := ACD.NewAnnotator("ejection_fraction")
			efFlowEntry, err := ACD.NewFlowEntry(efAnnotator)
			efFlowEntry.Annotator = efAnnotator
			smokingAnnotator, err := ACD.NewAnnotator("smoking")
			smokingFlowEntry, err := ACD.NewFlowEntry(smokingAnnotator)
			smokingFlowEntry.Annotator = smokingAnnotator
			labValueAnnotator, err := ACD.NewAnnotator("lab_value")
			labValueFlowEntry, err := ACD.NewFlowEntry(labValueAnnotator)
			labValueFlowEntry.Annotator = labValueAnnotator
			spellCheckParams := make(map[string][]string)
			spellCheckProfile := []string{"default"}
			spellCheckParams["spell_check_profile"] = spellCheckProfile
			spellAnnotator, err := ACD.NewAnnotator("spell_checker")
			spellAnnotator.Parameters = spellCheckParams
			spellFlowEntry, err := ACD.NewFlowEntry(spellAnnotator)
			spellFlowEntry.Annotator = spellAnnotator
			async := false
			flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*cdFlowEntry, *cvFlowEntry, *cancerFlowEntry, *efFlowEntry, *smokingFlowEntry, *labValueFlowEntry, *spellFlowEntry}, core.BoolPtr(async))
			annotatorFlow, err := ACD.NewAnnotatorFlow(flow)
			pipelineOptions.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
			container := ACD.NewUnstructuredContainer()
			container.SetText(analyzeText)
			pipelineOptions.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*container})
			pipelineOptions.SetDebugTextRestore(false)
			result, detailedResponse, err := ACD.RunPipeline(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			Expect(result.AnnotatorFlows).ToNot(BeNil())
			for _, element := range result.Unstructured {
				Expect(element.Data.Concepts).ToNot(BeNil())
				for _, conceptEntry := range element.Data.Concepts {
					Expect(conceptEntry.Begin).ToNot(BeNil())
					Expect(conceptEntry.End).ToNot(BeNil())
					Expect(conceptEntry.CoveredText).ToNot(BeNil())
					Expect(conceptEntry.Type).ToNot(BeNil())
					Expect(conceptEntry.Cui).ToNot(BeNil())
					Expect(conceptEntry.PreferredName).ToNot(BeNil())
					Expect(conceptEntry.Source).ToNot(BeNil())
				}
				Expect(element.Data.IcaCancerDiagnosisInd).ToNot(BeNil())
				for _, icaCancerEntry := range element.Data.IcaCancerDiagnosisInd {
					Expect(icaCancerEntry.Begin).ToNot(BeNil())
					Expect(icaCancerEntry.End).ToNot(BeNil())
					Expect(icaCancerEntry.CoveredText).ToNot(BeNil())
					Expect(icaCancerEntry.Type).ToNot(BeNil())
					Expect(icaCancerEntry.Modality).ToNot(BeNil())
				}
				Expect(element.Data.EjectionFractionInd).ToNot(BeNil())
				for _, ejectionFractionEntry := range element.Data.EjectionFractionInd {
					Expect(ejectionFractionEntry.Begin).ToNot(BeNil())
					Expect(ejectionFractionEntry.End).ToNot(BeNil())
					Expect(ejectionFractionEntry.CoveredText).ToNot(BeNil())
					Expect(ejectionFractionEntry.Type).ToNot(BeNil())
					Expect(ejectionFractionEntry.IsRange).ToNot(BeNil())
				}
				Expect(element.Data.SmokingInd).ToNot(BeNil())
				for _, smokingEntry := range element.Data.SmokingInd {
					Expect(smokingEntry.Begin).ToNot(BeNil())
					Expect(smokingEntry.End).ToNot(BeNil())
					Expect(smokingEntry.CoveredText).ToNot(BeNil())
					Expect(smokingEntry.Type).ToNot(BeNil())
				}
				Expect(element.Data.LabValueInd).ToNot(BeNil())
				for _, labValueEntry := range element.Data.LabValueInd {
					Expect(labValueEntry.Begin).ToNot(BeNil())
					Expect(labValueEntry.End).ToNot(BeNil())
					Expect(labValueEntry.CoveredText).ToNot(BeNil())
					Expect(labValueEntry.Type).ToNot(BeNil())
				}
			}
		})
	})

	Describe(`RunPipeline(runPipelineOptions *RunPipelineOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineOptions()
			cdAnnotator, err := ACD.NewAnnotator("concept_detection")
			cdFlowEntry, err := ACD.NewFlowEntry(cdAnnotator)
			cdFlowEntry.Annotator = cdAnnotator

			allergyAnnotator, err := ACD.NewAnnotator("allergy")
			allergyFlowEntry, err := ACD.NewFlowEntry(allergyAnnotator)
			allergyFlowEntry.Annotator = allergyAnnotator
			async := false
			flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*cdFlowEntry, *allergyFlowEntry}, core.BoolPtr(async))
			annotatorFlow, err := ACD.NewAnnotatorFlow(flow)

			pipelineOptions.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
			container := ACD.NewUnstructuredContainer()

			container.SetText("Patient is allergic to aspirin, dust, and mold.")
			pipelineOptions.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*container})
			pipelineOptions.SetDebugTextRestore(false)
			result, detailedResponse, err := ACD.RunPipeline(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			Expect(result.AnnotatorFlows).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			for _, element := range result.Unstructured {
				Expect(element.Data).ToNot(BeNil())
				containerAnno := element.Data
				Expect(containerAnno.AllergyInd).ToNot(BeNil())
				for _, allergy := range containerAnno.AllergyInd {
					Expect(allergy.Begin).ToNot(BeNil())
					Expect(allergy.End).ToNot(BeNil())
					Expect(allergy.CoveredText).ToNot(BeNil())
					Expect(allergy.Type).ToNot(BeNil())
                }
				Expect(containerAnno.AllergyMedicationInd).ToNot(BeNil())
				for _, allergyMed := range containerAnno.AllergyMedicationInd {
					Expect(allergyMed.Begin).ToNot(BeNil())
					Expect(allergyMed.End).ToNot(BeNil())
					Expect(allergyMed.CoveredText).ToNot(BeNil())
					Expect(allergyMed.Type).ToNot(BeNil())
        }
      }
		})
	})

	Describe(`RunPipeline(runPipelineOptions *RunPipelineOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineOptions()
			cdAnnotator, err := ACD.NewAnnotator("concept_detection")
			cdFlowEntry, err := ACD.NewFlowEntry(cdAnnotator)
			cdFlowEntry.Annotator = cdAnnotator
			seeingAssistAnnotator, err := ACD.NewAnnotator("seeing_assistance")
			seeingAssistFlowEntry, err := ACD.NewFlowEntry(seeingAssistAnnotator)
			seeingAssistFlowEntry.Annotator = seeingAssistAnnotator
			async := false
			flow, err := ACD.NewFlow([]annotatorforclinicaldataacdv1.FlowEntry{*cdFlowEntry, *seeingAssistFlowEntry}, core.BoolPtr(async))
			annotatorFlow, err := ACD.NewAnnotatorFlow(flow)
			pipelineOptions.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlow})
			container := ACD.NewUnstructuredContainer()
			container.SetText("She has macular degeneration which puts her at high risk for seeing difficulties.")
			pipelineOptions.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*container})
			pipelineOptions.SetDebugTextRestore(false)
			result, detailedResponse, err := ACD.RunPipeline(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			Expect(result.AnnotatorFlows).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			for _, element := range result.Unstructured {
				Expect(element.Data).ToNot(BeNil())
				containerAnno := element.Data
				Expect(containerAnno.SeeingAssistanceInd).ToNot(BeNil())
				for _, seeingAssist := range containerAnno.SeeingAssistanceInd {
					Expect(seeingAssist.Begin).ToNot(BeNil())
					Expect(seeingAssist.End).ToNot(BeNil())
					Expect(seeingAssist.CoveredText).ToNot(BeNil())
					Expect(seeingAssist.Type).ToNot(BeNil())
					Expect(seeingAssist.Modality).ToNot(BeNil())
        }
      }
		})
	})

	Describe(`RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineWithFlowOptions("wh_acd.ibm_clinical_insights_v1.0_standard_flow", false)
			pipelineOptions.SetBody(analyzeText)
			pipelineOptions.SetContentType("text/plain")
			pipelineOptions.SetDebugTextRestore(false)
			result, detailedResponse, err := ACD.RunPipelineWithFlow(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			for _, element := range result.Unstructured {
				Expect(element.Data).ToNot(BeNil())
				containerAnno := element.Data
				Expect(containerAnno.AttributeValues).ToNot(BeNil())
				for _, attributeValue := range containerAnno.AttributeValues {
					Expect(attributeValue.Begin).ToNot(BeNil())
					Expect(attributeValue.End).ToNot(BeNil())
					Expect(attributeValue.CoveredText).ToNot(BeNil())
					Expect(attributeValue.PreferredName).ToNot(BeNil())
					Expect(attributeValue.Negated).ToNot(BeNil())
					Expect(attributeValue.Values).ToNot(BeNil())
					for _, attributeValueEntry := range attributeValue.Values {
						Expect(attributeValueEntry.Value).ToNot(BeNil())
          }
					if (attributeValue.Temporal != nil) {
						for _, element := range attributeValue.Temporal {
							Expect(element.Begin).ToNot(BeNil())
							Expect(element.End).ToNot(BeNil())
							Expect(element.CoveredText).ToNot(BeNil())
						}
					}
				}
				Expect(containerAnno.MedicationInd).ToNot(BeNil())
				for _, medIndEntry := range containerAnno.MedicationInd {
					Expect(medIndEntry.Begin).ToNot(BeNil())
					Expect(medIndEntry.End).ToNot(BeNil())
					Expect(medIndEntry.CoveredText).ToNot(BeNil())
					Expect(medIndEntry.Type).ToNot(BeNil())
					Expect(medIndEntry.Negated).ToNot(BeNil())
					if (medIndEntry.InsightModelData != nil) {
						Expect(medIndEntry.InsightModelData.Medication).ToNot(BeNil())
						medication := medIndEntry.InsightModelData.Medication
						Expect(medication.Usage).ToNot(BeNil())
						Expect(medication.StartedEvent).ToNot(BeNil())
						Expect(medication.StoppedEvent).ToNot(BeNil())
						Expect(medication.AdverseEvent).ToNot(BeNil())
          }
				}
      }
		})
	})
	Describe(`RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineWithFlowOptions("wh_acd.ibm_clinical_insights_v1.0_standard_flow", false)
			pipelineOptions.SetBody("He is taking hydrochlorothiazide for high blood pressure. Systolic blood pressure is 150 mmHg and r diastolic blood pressure is 90 mmHg.\n")
			pipelineOptions.SetContentType("text/plain")
			pipelineOptions.SetDebugTextRestore(false)
			result, detailedResponse, err := ACD.RunPipelineWithFlow(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			for _, element := range result.Unstructured {
				Expect(element.Data).ToNot(BeNil())
				containerAnno := element.Data
				Expect(containerAnno.AttributeValues).ToNot(BeNil())
				for _, attributeValue := range containerAnno.AttributeValues {
					Expect(attributeValue.Begin).ToNot(BeNil())
					Expect(attributeValue.End).ToNot(BeNil())
					Expect(attributeValue.CoveredText).ToNot(BeNil())
                                        if (attributeValue.DisambiguationData != nil) {
					        Expect(attributeValue.DisambiguationData.Validity).ToNot(BeNil())
                                        }
					Expect(attributeValue.PreferredName).ToNot(BeNil())
					Expect(attributeValue.Negated).ToNot(BeNil())
                                        if (attributeValue.DerivedFrom != nil) {
                                                for _, derivedFrom := range attributeValue.DerivedFrom {
                                                        Expect(derivedFrom.UID).ToNot(BeNil())
                                                }
                                        }
                                        if (attributeValue.EvidenceSpans != nil) {
                                                for _, evidenceSpan := range attributeValue.EvidenceSpans {
                                                        Expect(evidenceSpan.UID).ToNot(BeNil())
                                                }
                                        }
					Expect(attributeValue.Values).ToNot(BeNil())
					for _, attributeValueEntry := range attributeValue.Values {
						Expect(attributeValueEntry.Value).ToNot(BeNil())
                                        }
          }
				Expect(containerAnno.MedicationInd).ToNot(BeNil())
				for _, medIndEntry := range containerAnno.MedicationInd {
					Expect(medIndEntry.Begin).ToNot(BeNil())
					Expect(medIndEntry.End).ToNot(BeNil())
					Expect(medIndEntry.CoveredText).ToNot(BeNil())
					Expect(medIndEntry.Type).ToNot(BeNil())
					Expect(medIndEntry.Negated).ToNot(BeNil())
					if (medIndEntry.InsightModelData != nil) {
						Expect(medIndEntry.InsightModelData.Medication).ToNot(BeNil())
						medication := medIndEntry.InsightModelData.Medication
						Expect(medication.Usage).ToNot(BeNil())
						Expect(medication.StartedEvent).ToNot(BeNil())
						Expect(medication.StoppedEvent).ToNot(BeNil())
						Expect(medication.DoseChangedEvent).ToNot(BeNil())
						Expect(medication.AdverseEvent).ToNot(BeNil())
          }
				}
				Expect(containerAnno.SpellCorrectedText).ToNot(BeNil())
				for _, spellCorrectedEntry := range containerAnno.SpellCorrectedText {
					Expect(spellCorrectedEntry.CorrectedText).ToNot(BeNil())
					Expect(spellCorrectedEntry.DebugText).ToNot(BeNil())
				}
                                if (containerAnno.Lines != nil) {
                                        for _, linesEntry := range containerAnno.Lines {
                                                Expect(linesEntry.Begin).ToNot(BeNil())
                                                Expect(linesEntry.End).ToNot(BeNil())
                                        }
          }
                                if (containerAnno.Sentences != nil) {
                                        for _, sentencesEntry := range containerAnno.Sentences {
                                                Expect(sentencesEntry.Begin).ToNot(BeNil())
                                                Expect(sentencesEntry.End).ToNot(BeNil())
                                        }
          }
                                if (containerAnno.Paragraphs != nil) {
                                        for _, paragraphsEntry := range containerAnno.Paragraphs {
                                                Expect(paragraphsEntry.Begin).ToNot(BeNil())
                                                Expect(paragraphsEntry.End).ToNot(BeNil())
                                        }
          }
        }
		})
	})

	Describe(`RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions)`, func() {
		It(`Successfully run analyze pipeline`, func() {
			pipelineOptions := ACD.NewRunPipelineWithFlowOptions("wh_acd.ibm_clinical_insights_v1.0_standard_flow", false)
			container := ACD.NewUnstructuredContainer()
			container.SetText("The CT scan showed a tumor in the left lung.")
			concept := new(annotatorforclinicaldataacdv1.Concept)
			concept.Begin = core.Int64Ptr(int64(4))
			concept.End = core.Int64Ptr(int64(11))
			concept.Cui = core.StringPtr("C0040405")
			concept.PreferredName = core.StringPtr("X-Ray")
			concept.CoveredText = core.StringPtr("CT Scan")
			concept.Type = core.StringPtr("umls.DiagnosticProcedure")
			concept.Source = core.StringPtr("umls")
			concept.SourceVersion = core.StringPtr("2020AA")
			containerAnno := new(annotatorforclinicaldataacdv1.ContainerAnnotation)
			containerAnno.Concepts = []annotatorforclinicaldataacdv1.Concept{*concept}
			container.SetData(containerAnno)
			analyticFlowBeanInput := ACD.NewAnalyticFlowBeanInput()
			analyticFlowBeanInput.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*container})
			pipelineOptions.SetAnalyticFlowBeanInput(analyticFlowBeanInput)
			pipelineOptions.SetContentType("application/json")
			pipelineOptions.SetDebugTextRestore(false)
			result, detailedResponse, err := ACD.RunPipelineWithFlow(pipelineOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Unstructured).ToNot(BeNil())
			for _, element := range result.Unstructured {
				Expect(element.Data).ToNot(BeNil())
				containerAnno := element.Data
				Expect(containerAnno.AttributeValues).ToNot(BeNil())
				for _, attributeValue := range containerAnno.AttributeValues {
					Expect(attributeValue.Begin).ToNot(BeNil())
					Expect(attributeValue.End).ToNot(BeNil())
					Expect(attributeValue.CoveredText).ToNot(BeNil())
                                        if (attributeValue.DisambiguationData != nil) {
                                                Expect(attributeValue.DisambiguationData.Validity).ToNot(BeNil())
                                        }
					Expect(attributeValue.PreferredName).ToNot(BeNil())
					Expect(attributeValue.Negated).ToNot(BeNil())
					Expect(attributeValue.Values).ToNot(BeNil())
					for _, attributeValueEntry := range attributeValue.Values {
						Expect(attributeValueEntry.Value).ToNot(BeNil())
          }
					if (attributeValue.InsightModelData != nil) {
						if (attributeValue.InsightModelData.Normality != nil) {
							norm := attributeValue.InsightModelData.Normality
							if (norm.Evidence != nil) {
								for _, evidence := range norm.Evidence {
									Expect(evidence.Begin).ToNot(BeNil())
									Expect(evidence.End).ToNot(BeNil())
									Expect(evidence.CoveredText).ToNot(BeNil())
								}
							}
						}
						if (attributeValue.InsightModelData.Procedure != nil) {
							proc := attributeValue.InsightModelData.Procedure
							if (proc.Modifiers != nil) {
								if (proc.Modifiers.AssociatedDiagnosis != nil) {
									for _, assocDiag := range proc.Modifiers.AssociatedDiagnosis {
										Expect(assocDiag.Begin).ToNot(BeNil())
										Expect(assocDiag.End).ToNot(BeNil())
										Expect(assocDiag.CoveredText).ToNot(BeNil())
									}
								}
							}
						}
          }
				}
				Expect(containerAnno.MedicationInd).ToNot(BeNil())
				for _, medIndEntry := range containerAnno.MedicationInd {
					Expect(medIndEntry.Begin).ToNot(BeNil())
					Expect(medIndEntry.End).ToNot(BeNil())
					Expect(medIndEntry.CoveredText).ToNot(BeNil())
					Expect(medIndEntry.Type).ToNot(BeNil())
					Expect(medIndEntry.Negated).ToNot(BeNil())
					if (medIndEntry.InsightModelData != nil) {
						Expect(medIndEntry.InsightModelData.Medication).ToNot(BeNil())
						medication := medIndEntry.InsightModelData.Medication
						Expect(medication.Usage).ToNot(BeNil())
						Expect(medication.StartedEvent).ToNot(BeNil())
						Expect(medication.StoppedEvent).ToNot(BeNil())
						Expect(medication.DoseChangedEvent).ToNot(BeNil())
						Expect(medication.AdverseEvent).ToNot(BeNil())
          }
				}

				Expect(containerAnno.ProcedureInd).ToNot(BeNil())
				for _, pIndEntry := range containerAnno.ProcedureInd {
					Expect(pIndEntry.Begin).ToNot(BeNil())
					Expect(pIndEntry.End).ToNot(BeNil())
					Expect(pIndEntry.CoveredText).ToNot(BeNil())
					Expect(pIndEntry.Type).ToNot(BeNil())
					Expect(pIndEntry.Negated).ToNot(BeNil())
					if (pIndEntry.InsightModelData != nil) {
						Expect(pIndEntry.InsightModelData.Normality).ToNot(BeNil())
						norm := pIndEntry.InsightModelData.Normality
						Expect(norm.NormalityUsage).ToNot(BeNil())
						Expect(norm.Evidence).ToNot(BeNil())
          }
				}
				Expect(containerAnno.SpellCorrectedText).ToNot(BeNil())
				for _, spellCorrectedEntry := range containerAnno.SpellCorrectedText {
					Expect(spellCorrectedEntry.CorrectedText).ToNot(BeNil())
					Expect(spellCorrectedEntry.DebugText).ToNot(BeNil())
				}
      }
		})
	})

	Describe(`DeleteUserSpecificArtifacts(deleteUserSpecificArtifactsOptions *DeleteUserSpecificArtifactsOptions)`, func() {
		It(`Successfully delete user specific artifacts`, func() {
			artifactsOptions := ACD.NewDeleteUserSpecificArtifactsOptions()
            sdkHeaders := make(map[string]string)
            sdkHeaders["X-IBM-Client-ID"] = "sdk_test"
            artifactsOptions.SetHeaders(sdkHeaders)
			detailedResponse, err := ACD.DeleteUserSpecificArtifacts(artifactsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
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
			cartridgeCounter := 0
			for cartridgeCounter < len(result.Cartridges) {
				cartridge := result.Cartridges[cartridgeCounter]
				Expect(cartridge.ID).ToNot(BeNil())
				Expect(cartridge.StatusLocation).ToNot(BeNil())
				Expect(cartridge.StartTime).ToNot(BeNil())
				Expect(cartridge.EndTime).ToNot(BeNil())
				cartridgeCounter++
			}
		})
	})

	Describe(`CartridgesGetID(cartridgesGetIdOptions *CartridgesGetIdOptions)`, func() {
		It(`Successfully retrieve the cartridge`, func() {
			cartridgesGetIdOptions := ACD.NewCartridgesGetIdOptions("wh_acd.ibm_clinical_insights_v1.0")
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
			responseCounter := 0
			responses := len(result.ArtifactResponse)
			for responseCounter < responses {
				response := result.ArtifactResponse[responseCounter]
				Expect(response.Code).ToNot(BeNil())
				Expect(response.Message).ToNot(BeNil())
				Expect(response.Level).ToNot(BeNil())
				responseCounter++
			}
		})
	})

	Describe(`CartridgesPostMultipart(cartridgesPostMultipartOptions *CartridgesPostMultipartOptions) `, func() {
		It(`Successfully start a cartridge archive deployment`, func() {
			cartridgesPostMultipartOptions := ACD.NewCartridgesPostMultipartOptions()
			cartridgesPostMultipartOptions.SetArchiveFileContentType("application/octet-stream")
			archiveFile, err := os.Open(config["ARCHIVE"])
			cartridgesPostMultipartOptions.SetArchiveFile(archiveFile)
			result, detailedResponse, err := ACD.CartridgesPostMultipart(cartridgesPostMultipartOptions)
			if (err != nil) {
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(409))
				Expect(result).To(BeNil())
			} else {
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(202))
				Expect(result).ToNot(BeNil())
      }
		})
	})

	Describe(`CartridgesPutMultipart(cartridgesPutMultipartOptions *CartridgesPutMultipartOptions) `, func() {
		It(`Successfully start a cartridge archive deployment`, func() {
			cartridgesPutMultipartOptions := ACD.NewCartridgesPutMultipartOptions()
			cartridgesPutMultipartOptions.SetArchiveFileContentType("application/octet-stream")
			archiveFile, err := os.Open(config["ARCHIVE"])
			cartridgesPutMultipartOptions.SetArchiveFile(archiveFile)
			result, detailedResponse, err := ACD.CartridgesPutMultipart(cartridgesPutMultipartOptions)
			if (err != nil) {
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(409))
				Expect(result).To(BeNil())
			} else {
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(202))
				Expect(result).ToNot(BeNil())
			}
		})
	})

	Describe(`DeployCartridge(deployCartridgeOptions *DeployCartridgeOptions) `, func() {
		It(`Successfully start a cartridge archive deployment`, func() {
			deployCartridgeOptions := ACD.NewDeployCartridgeOptions()
			deployCartridgeOptions.SetArchiveFileContentType("application/octet-stream")
			archiveFile, err := os.Open(config["ARCHIVE"])
			deployCartridgeOptions.SetArchiveFile(archiveFile)
			deployCartridgeOptions.SetUpdate(true)
			result, detailedResponse, err := ACD.DeployCartridge(deployCartridgeOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).Should(BeNumerically(">=", 200))
			Expect(detailedResponse.StatusCode).Should(BeNumerically("<", 299))
			Expect(result).ToNot(BeNil())
			Expect(result.Code).ToNot(BeNil())
		})
	})
})
