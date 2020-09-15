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
	"bytes"
	"fmt"
	"github.com/IBM/whcs-go-sdk/annotatorforclinicaldataacdv1"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`AnnotatorForClinicalDataAcdV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "https://annotatorforclinicaldataacdv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetProfiles(getProfilesOptions *GetProfilesOptions) - Operation response error`, func() {
		version := "testString"
		getProfilesPath := "/v1/profiles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfiles with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProfilesOptions model
				getProfilesOptionsModel := new(annotatorforclinicaldataacdv1.GetProfilesOptions)
				getProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetProfiles(getProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetProfiles(getProfilesOptions *GetProfilesOptions)`, func() {
		version := "testString"
		getProfilesPath := "/v1/profiles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"data": {}}`)
				}))
			})
			It(`Invoke GetProfiles successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfilesOptions model
				getProfilesOptionsModel := new(annotatorforclinicaldataacdv1.GetProfilesOptions)
				getProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetProfiles(getProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetProfiles with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProfilesOptions model
				getProfilesOptionsModel := new(annotatorforclinicaldataacdv1.GetProfilesOptions)
				getProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetProfiles(getProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateProfile(createProfileOptions *CreateProfileOptions)`, func() {
		version := "testString"
		createProfilePath := "/v1/profiles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createProfilePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateProfile successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.CreateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Annotator model
				annotatorModel := new(annotatorforclinicaldataacdv1.Annotator)
				annotatorModel.Name = core.StringPtr("testString")
				annotatorModel.Parameters = make(map[string][]string)
				annotatorModel.Configurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(annotatorforclinicaldataacdv1.CreateProfileOptions)
				createProfileOptionsModel.ID = core.StringPtr("testString")
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.PublishedDate = core.StringPtr("testString")
				createProfileOptionsModel.Publish = core.BoolPtr(true)
				createProfileOptionsModel.Version = core.StringPtr("testString")
				createProfileOptionsModel.CartridgeID = core.StringPtr("testString")
				createProfileOptionsModel.Annotators = []annotatorforclinicaldataacdv1.Annotator{*annotatorModel}
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(201))
			})
			It(`Invoke CreateProfile with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Annotator model
				annotatorModel := new(annotatorforclinicaldataacdv1.Annotator)
				annotatorModel.Name = core.StringPtr("testString")
				annotatorModel.Parameters = make(map[string][]string)
				annotatorModel.Configurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := new(annotatorforclinicaldataacdv1.CreateProfileOptions)
				createProfileOptionsModel.ID = core.StringPtr("testString")
				createProfileOptionsModel.Name = core.StringPtr("testString")
				createProfileOptionsModel.Description = core.StringPtr("testString")
				createProfileOptionsModel.PublishedDate = core.StringPtr("testString")
				createProfileOptionsModel.Publish = core.BoolPtr(true)
				createProfileOptionsModel.Version = core.StringPtr("testString")
				createProfileOptionsModel.CartridgeID = core.StringPtr("testString")
				createProfileOptionsModel.Annotators = []annotatorforclinicaldataacdv1.Annotator{*annotatorModel}
				createProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.CreateProfile(createProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProfile(getProfilesByIdOptions *GetProfilesByIdOptions) - Operation response error`, func() {
		version := "testString"
		getProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProfile with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfilesByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetProfilesByIdOptions)
				getProfilesByIdOptionsModel.ID = core.StringPtr("testString")
				getProfilesByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetProfilesByID(getProfilesByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetProfile(getProfilesByIdOptions *GetProfilesByIdOptions)`, func() {
		version := "testString"
		getProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "name": "Name", "description": "Description", "publishedDate": "PublishedDate", "publish": false, "version": "Version", "cartridgeId": "CartridgeID", "annotators": [{"name": "Name", "parameters": {"mapKey": ["Inner"]}, "configurations": [{"id": "ID", "type": "Type", "uid": 3, "mergeid": 7}]}]}`)
				}))
			})
			It(`Invoke GetProfile successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetProfilesByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfilesByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetProfilesByIdOptions)
				getProfilesByIdOptionsModel.ID = core.StringPtr("testString")
				getProfilesByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetProfilesByID(getProfilesByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.ID).ToNot(BeNil())
				Expect(result.Name).ToNot(BeNil())
				Expect(result.Description).ToNot(BeNil())
				Expect(result.Annotators).ToNot(BeNil())

			})
			It(`Invoke GetProfile with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetProfileOptions model
				getProfilesByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetProfilesByIdOptions)
				getProfilesByIdOptionsModel.ID = core.StringPtr("testString")
				getProfilesByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetProfilesByID(getProfilesByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProfileOptions model with no property values
				getProfilesByIdOptionsModelNew := new(annotatorforclinicaldataacdv1.GetProfilesByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetProfilesByID(getProfilesByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateProfile(updateProfileOptions *UpdateProfileOptions)`, func() {
		version := "testString"
		updateProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateProfilePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateProfile successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Annotator model
				annotatorModel := new(annotatorforclinicaldataacdv1.Annotator)
				annotatorModel.Name = core.StringPtr("testString")
				annotatorModel.Parameters = make(map[string][]string)
				annotatorModel.Configurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(annotatorforclinicaldataacdv1.UpdateProfileOptions)
				updateProfileOptionsModel.ID = core.StringPtr("testString")
				updateProfileOptionsModel.NewID = core.StringPtr("testString")
				updateProfileOptionsModel.NewName = core.StringPtr("testString")
				updateProfileOptionsModel.NewDescription = core.StringPtr("testString")
				updateProfileOptionsModel.NewPublishedDate = core.StringPtr("testString")
				updateProfileOptionsModel.NewPublish = core.BoolPtr(true)
				updateProfileOptionsModel.NewVersion = core.StringPtr("testString")
				updateProfileOptionsModel.NewCartridgeID = core.StringPtr("testString")
				updateProfileOptionsModel.NewAnnotators = []annotatorforclinicaldataacdv1.Annotator{*annotatorModel}
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateProfile with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Annotator model
				annotatorModel := new(annotatorforclinicaldataacdv1.Annotator)
				annotatorModel.Name = core.StringPtr("testString")
				annotatorModel.Parameters = make(map[string][]string)
				annotatorModel.Configurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}

				// Construct an instance of the UpdateProfileOptions model
				updateProfileOptionsModel := new(annotatorforclinicaldataacdv1.UpdateProfileOptions)
				updateProfileOptionsModel.ID = core.StringPtr("testString")
				updateProfileOptionsModel.NewID = core.StringPtr("testString")
				updateProfileOptionsModel.NewName = core.StringPtr("testString")
				updateProfileOptionsModel.NewDescription = core.StringPtr("testString")
				updateProfileOptionsModel.NewPublishedDate = core.StringPtr("testString")
				updateProfileOptionsModel.NewPublish = core.BoolPtr(true)
				updateProfileOptionsModel.NewVersion = core.StringPtr("testString")
				updateProfileOptionsModel.NewCartridgeID = core.StringPtr("testString")
				updateProfileOptionsModel.NewAnnotators = []annotatorforclinicaldataacdv1.Annotator{*annotatorModel}
				updateProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateProfile(updateProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateProfileOptions model with no property values
				updateProfileOptionsModelNew := new(annotatorforclinicaldataacdv1.UpdateProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateProfile(updateProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteProfile(deleteProfileOptions *DeleteProfileOptions)`, func() {
		version := "testString"
		deleteProfilePath := "/v1/profiles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteProfilePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteProfile successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteProfileOptions model
				deleteProfileOptionsModel := new(annotatorforclinicaldataacdv1.DeleteProfileOptions)
				deleteProfileOptionsModel.ID = core.StringPtr("testString")
				deleteProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteProfile(deleteProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteProfile with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteProfileOptions model
				deleteProfileOptionsModel := new(annotatorforclinicaldataacdv1.DeleteProfileOptions)
				deleteProfileOptionsModel.ID = core.StringPtr("testString")
				deleteProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteProfile(deleteProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())

				// Construct a second instance of the DeleteProfileOptions model with no property values
				deleteProfileOptionsModelNew := new(annotatorforclinicaldataacdv1.DeleteProfileOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteProfile(deleteProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "https://annotatorforclinicaldataacdv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetFlows(getFlowsOptions *GetFlowsOptions) - Operation response error`, func() {
		version := "testString"
		getFlowsPath := "/v1/flows"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFlowsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFlows with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFlowsOptions model
				getFlowsOptionsModel := new(annotatorforclinicaldataacdv1.GetFlowsOptions)
				getFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetFlows(getFlowsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFlows(getFlowsOptions *GetFlowsOptions)`, func() {
		version := "testString"
		getFlowsPath := "/v1/flows"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFlowsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"data": {}}`)
				}))
			})
			It(`Invoke GetFlows successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetFlows(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFlowsOptions model
				getFlowsOptionsModel := new(annotatorforclinicaldataacdv1.GetFlowsOptions)
				getFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetFlows(getFlowsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetFlows with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFlowsOptions model
				getFlowsOptionsModel := new(annotatorforclinicaldataacdv1.GetFlowsOptions)
				getFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetFlows(getFlowsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateFlows(createFlowsOptions *CreateFlowsOptions)`, func() {
		version := "testString"
		createFlowsPath := "/v1/flows"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createFlowsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateFlows successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.CreateFlows(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the CreateFlowsOptions model
				createFlowsOptionsModel := new(annotatorforclinicaldataacdv1.CreateFlowsOptions)
				createFlowsOptionsModel.ID = core.StringPtr("testString")
				createFlowsOptionsModel.Name = core.StringPtr("testString")
				createFlowsOptionsModel.Description = core.StringPtr("testString")
				createFlowsOptionsModel.PublishedDate = core.StringPtr("testString")
				createFlowsOptionsModel.Publish = core.BoolPtr(true)
				createFlowsOptionsModel.Version = core.StringPtr("testString")
				createFlowsOptionsModel.CartridgeID = core.StringPtr("testString")
				createFlowsOptionsModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				createFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.CreateFlows(createFlowsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CreateFlows with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the CreateFlowsOptions model
				createFlowsOptionsModel := new(annotatorforclinicaldataacdv1.CreateFlowsOptions)
				createFlowsOptionsModel.ID = core.StringPtr("testString")
				createFlowsOptionsModel.Name = core.StringPtr("testString")
				createFlowsOptionsModel.Description = core.StringPtr("testString")
				createFlowsOptionsModel.PublishedDate = core.StringPtr("testString")
				createFlowsOptionsModel.Publish = core.BoolPtr(true)
				createFlowsOptionsModel.Version = core.StringPtr("testString")
				createFlowsOptionsModel.CartridgeID = core.StringPtr("testString")
				createFlowsOptionsModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				createFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.CreateFlows(createFlowsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFlowsByID(getFlowsByIdOptions *GetFlowsByIdOptions) - Operation response error`, func() {
		version := "testString"
		getFlowsByIDPath := "/v1/flows/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFlowsByIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFlowsByID with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFlowsByIdOptions model
				getFlowsByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetFlowsByIdOptions)
				getFlowsByIdOptionsModel.ID = core.StringPtr("testString")
				getFlowsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetFlowsByID(getFlowsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFlowsByID(getFlowsByIdOptions *GetFlowsByIdOptions)`, func() {
		version := "testString"
		getFlowsByIDPath := "/v1/flows/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFlowsByIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "name": "Name", "description": "Description", "publishedDate": "PublishedDate", "publish": false, "version": "Version", "cartridgeId": "CartridgeID", "annotatorFlows": [{"profile": "Profile", "flow": {"elements": [{}], "async": false}, "id": "ID", "type": "Type", "data": {"mapKey": [{"id": "ID", "type": "Type", "uid": 3, "mergeid": 7}]}, "metadata": {"mapKey": {"anyKey": "anyValue"}}, "globalConfigurations": [{"id": "ID", "type": "Type", "uid": 3, "mergeid": 7}], "uid": 3}]}`)
				}))
			})
			It(`Invoke GetFlowsByID successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetFlowsByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFlowsByIdOptions model
				getFlowsByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetFlowsByIdOptions)
				getFlowsByIdOptionsModel.ID = core.StringPtr("testString")
				getFlowsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetFlowsByID(getFlowsByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.ID).ToNot(BeNil())
				Expect(result.Name).ToNot(BeNil())
				Expect(result.Description).ToNot(BeNil())
				Expect(result.AnnotatorFlows).ToNot(BeNil())
			})
			It(`Invoke GetFlowsByID with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFlowsByIdOptions model
				getFlowsByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetFlowsByIdOptions)
				getFlowsByIdOptionsModel.ID = core.StringPtr("testString")
				getFlowsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetFlowsByID(getFlowsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct a second instance of the GetFlowsByIdOptions model with no property values
				getFlowsByIdOptionsModelNew := new(annotatorforclinicaldataacdv1.GetFlowsByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetFlowsByID(getFlowsByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateFlows(updateFlowsOptions *UpdateFlowsOptions)`, func() {
		version := "testString"
		updateFlowsPath := "/v1/flows/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateFlowsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateFlows successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateFlows(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the UpdateFlowsOptions model
				updateFlowsOptionsModel := new(annotatorforclinicaldataacdv1.UpdateFlowsOptions)
				updateFlowsOptionsModel.ID = core.StringPtr("testString")
				updateFlowsOptionsModel.NewID = core.StringPtr("testString")
				updateFlowsOptionsModel.NewName = core.StringPtr("testString")
				updateFlowsOptionsModel.NewDescription = core.StringPtr("testString")
				updateFlowsOptionsModel.NewPublishedDate = core.StringPtr("testString")
				updateFlowsOptionsModel.NewPublish = core.BoolPtr(true)
				updateFlowsOptionsModel.NewVersion = core.StringPtr("testString")
				updateFlowsOptionsModel.NewCartridgeID = core.StringPtr("testString")
				updateFlowsOptionsModel.NewAnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				updateFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateFlows(updateFlowsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateFlows with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the UpdateFlowsOptions model
				updateFlowsOptionsModel := new(annotatorforclinicaldataacdv1.UpdateFlowsOptions)
				updateFlowsOptionsModel.ID = core.StringPtr("testString")
				updateFlowsOptionsModel.NewID = core.StringPtr("testString")
				updateFlowsOptionsModel.NewName = core.StringPtr("testString")
				updateFlowsOptionsModel.NewDescription = core.StringPtr("testString")
				updateFlowsOptionsModel.NewPublishedDate = core.StringPtr("testString")
				updateFlowsOptionsModel.NewPublish = core.BoolPtr(true)
				updateFlowsOptionsModel.NewVersion = core.StringPtr("testString")
				updateFlowsOptionsModel.NewCartridgeID = core.StringPtr("testString")
				updateFlowsOptionsModel.NewAnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				updateFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateFlows(updateFlowsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())

				// Construct a second instance of the UpdateFlowsOptions model with no property values
				updateFlowsOptionsModelNew := new(annotatorforclinicaldataacdv1.UpdateFlowsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateFlows(updateFlowsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteFlows(deleteFlowsOptions *DeleteFlowsOptions)`, func() {
		version := "testString"
		deleteFlowsPath := "/v1/flows/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteFlowsPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteFlows successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteFlows(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteFlowsOptions model
				deleteFlowsOptionsModel := new(annotatorforclinicaldataacdv1.DeleteFlowsOptions)
				deleteFlowsOptionsModel.ID = core.StringPtr("testString")
				deleteFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteFlows(deleteFlowsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteFlows with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteFlowsOptions model
				deleteFlowsOptionsModel := new(annotatorforclinicaldataacdv1.DeleteFlowsOptions)
				deleteFlowsOptionsModel.ID = core.StringPtr("testString")
				deleteFlowsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteFlows(deleteFlowsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())

				// Construct a second instance of the DeleteFlowsOptions model with no property values
				deleteFlowsOptionsModelNew := new(annotatorforclinicaldataacdv1.DeleteFlowsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteFlows(deleteFlowsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "https://annotatorforclinicaldataacdv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`RunPipeline(runPipelineOptions *RunPipelineOptions) - Operation response error`, func() {
		version := "testString"
		runPipelinePath := "/v1/analyze"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(runPipelinePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunPipeline with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RunPipelineOptions model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				runPipelineOptionsModel := new(annotatorforclinicaldataacdv1.RunPipelineOptions)
				runPipelineOptionsModel.Unstructured = []annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}
				runPipelineOptionsModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				runPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.RunPipeline(runPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunPipeline(runPipelineOptions *RunPipelineOptions)`, func() {
		version := "testString"
		runPipelinePath := "/v1/analyze"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(runPipelinePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"unstructured":[{"data": {"relations":[{"source":"umls","nodes":[{"entity":{"uid":2}}],"type":"may_treat"}],"nluEntities":[{"relevance":0.828337,"source":"test","type":"test","begin":19,"end":26,"coveredText":"test"}]}}]}`)
				}))
			})
			It(`Invoke RunPipeline successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RunPipeline(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the UnstructuredContainer model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				unstructuredContainerModel.Text = core.StringPtr("testString")
				unstructuredContainerModel.ID = core.StringPtr("testString")
				unstructuredContainerModel.Type = core.StringPtr("testString")
				unstructuredContainerModel.Data = new(annotatorforclinicaldataacdv1.ContainerAnnotation)
				unstructuredContainerModel.Metadata = make(map[string]interface{})
				unstructuredContainerModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the RunPipelineOptions model
				runPipelineOptionsModel := new(annotatorforclinicaldataacdv1.RunPipelineOptions)
				runPipelineOptionsModel.Unstructured = []annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}
				runPipelineOptionsModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				runPipelineOptionsModel.DebugTextRestore = core.BoolPtr(true)
				runPipelineOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				runPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RunPipeline(runPipelineOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke RunPipeline with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the UnstructuredContainer model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				unstructuredContainerModel.Text = core.StringPtr("testString")
				unstructuredContainerModel.ID = core.StringPtr("testString")
				unstructuredContainerModel.Type = core.StringPtr("testString")
				unstructuredContainerModel.Data = new(annotatorforclinicaldataacdv1.ContainerAnnotation)
				unstructuredContainerModel.Metadata = make(map[string]interface{})
				unstructuredContainerModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the RunPipelineOptions model
				runPipelineOptionsModel := new(annotatorforclinicaldataacdv1.RunPipelineOptions)
				runPipelineOptionsModel.Unstructured = []annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}
				runPipelineOptionsModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				runPipelineOptionsModel.DebugTextRestore = core.BoolPtr(true)
				runPipelineOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				runPipelineOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.RunPipeline(runPipelineOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct a second instance of the RunPipeline model with no property values
				runPipelineOptionsModelNew := new(annotatorforclinicaldataacdv1.RunPipelineOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.RunPipeline(runPipelineOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions) - Operation response error`, func() {
		version := "testString"
		runPipelineWithFlowPath := "/v1/analyze/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(runPipelineWithFlowPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunPipelineWithFlow with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RunPipelineWithFlowOptions model
				runPipelineWithFlowOptionsModel := new(annotatorforclinicaldataacdv1.RunPipelineWithFlowOptions)
				runPipelineWithFlowOptionsModel.FlowID = core.StringPtr("testString")
				runPipelineWithFlowOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				runPipelineWithFlowOptionsModel.Body = core.StringPtr("testString")
				runPipelineWithFlowOptionsModel.ContentType = core.StringPtr("application/json")
				runPipelineWithFlowOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.RunPipelineWithFlow(runPipelineWithFlowOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions)`, func() {
		version := "testString"
		runPipelineWithFlowPath := "/v1/analyze/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(runPipelineWithFlowPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke RunPipelineWithFlow successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RunPipelineWithFlow(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the UnstructuredContainer model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				unstructuredContainerModel.Text = core.StringPtr("testString")
				unstructuredContainerModel.ID = core.StringPtr("testString")
				unstructuredContainerModel.Type = core.StringPtr("testString")
				unstructuredContainerModel.Data = new(annotatorforclinicaldataacdv1.ContainerAnnotation)
				unstructuredContainerModel.Metadata = make(map[string]interface{})
				unstructuredContainerModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the AnalyticFlowBeanInput model
				analyticFlowBeanInputModel := new(annotatorforclinicaldataacdv1.AnalyticFlowBeanInput)
				analyticFlowBeanInputModel.Unstructured = []annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}
				analyticFlowBeanInputModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}

				// Construct an instance of the RunPipelineWithFlowOptions model
				runPipelineWithFlowOptionsModel := new(annotatorforclinicaldataacdv1.RunPipelineWithFlowOptions)
				runPipelineWithFlowOptionsModel.FlowID = core.StringPtr("testString")
				runPipelineWithFlowOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				runPipelineWithFlowOptionsModel.AnalyticFlowBeanInput = analyticFlowBeanInputModel
				runPipelineWithFlowOptionsModel.Body = core.StringPtr("testString")
				runPipelineWithFlowOptionsModel.ContentType = core.StringPtr("application/json")
				runPipelineWithFlowOptionsModel.DebugTextRestore = core.BoolPtr(true)
				runPipelineWithFlowOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RunPipelineWithFlow(runPipelineWithFlowOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(result).To(BeNil())
			})
			It(`Invoke RunPipelineWithFlow with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the UnstructuredContainer model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				unstructuredContainerModel.Text = core.StringPtr("testString")
				unstructuredContainerModel.ID = core.StringPtr("testString")
				unstructuredContainerModel.Type = core.StringPtr("testString")
				unstructuredContainerModel.Data = new(annotatorforclinicaldataacdv1.ContainerAnnotation)
				unstructuredContainerModel.Metadata = make(map[string]interface{})
				unstructuredContainerModel.Uid = core.Int64Ptr(int64(26))

				// Construct an instance of the AnalyticFlowBeanInput model
				analyticFlowBeanInputModel := new(annotatorforclinicaldataacdv1.AnalyticFlowBeanInput)
				analyticFlowBeanInputModel.Unstructured = []annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}
				analyticFlowBeanInputModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}

				// Construct an instance of the RunPipelineWithFlowOptions model
				runPipelineWithFlowOptionsModel := new(annotatorforclinicaldataacdv1.RunPipelineWithFlowOptions)
				runPipelineWithFlowOptionsModel.FlowID = core.StringPtr("testString")
				runPipelineWithFlowOptionsModel.ReturnAnalyzedText = core.BoolPtr(true)
				runPipelineWithFlowOptionsModel.AnalyticFlowBeanInput = analyticFlowBeanInputModel
				runPipelineWithFlowOptionsModel.Body = core.StringPtr("testString")
				runPipelineWithFlowOptionsModel.ContentType = core.StringPtr("application/json")
				runPipelineWithFlowOptionsModel.DebugTextRestore = core.BoolPtr(true)
				runPipelineWithFlowOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.RunPipelineWithFlow(runPipelineWithFlowOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct a second instance of the RunPipelineWithFlowOptions model with no property values
				runPipelineWithFlowOptionsModelNew := new(annotatorforclinicaldataacdv1.RunPipelineWithFlowOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.RunPipelineWithFlow(runPipelineWithFlowOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAnnotators(getAnnotatorsOptions *getAnnotatorsOptions) - Operation response error`, func() {
		version := "testString"
		getAnnotatorsPath := "/v1/annotators"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAnnotatorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAnnotators with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAnnotatorsOptions model
				getAnnotatorsOptionsModel := new(annotatorforclinicaldataacdv1.GetAnnotatorsOptions)
				getAnnotatorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetAnnotators(getAnnotatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAnnotators(getAnnotatorsOptions *GetAnnotatorsOptions)`, func() {
		version := "testString"
		getAnnotatorsPath := "/v1/annotators"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAnnotatorsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"data": {}}`)
				}))
			})
			It(`Invoke GetAnnotators successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAnnotators(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())


				// Construct an instance of the GetAnnotatorsOptions model
				getAnnotatorsOptionsModel := new(annotatorforclinicaldataacdv1.GetAnnotatorsOptions)
				getAnnotatorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAnnotators(getAnnotatorsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetAnnotators with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAnnotatorsOptions model
				getAnnotatorsOptionsModel := new(annotatorforclinicaldataacdv1.GetAnnotatorsOptions)
				getAnnotatorsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetAnnotators(getAnnotatorsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAnnotatorsByID(getAnnotatorsByIdOptions *getAnnotatorsByIdOptions) - Operation response error`, func() {
		version := "testString"
		getAnnotatorsByIDPath := "/v1/annotators/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAnnotatorsByIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAnnotatorsByID with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAnnotatorsByIdOptions model
				getAnnotatorsByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetAnnotatorsByIdOptions)
				getAnnotatorsByIdOptionsModel.ID = core.StringPtr("testString")
				getAnnotatorsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetAnnotatorsByID(getAnnotatorsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAnnotatorsByID(getAnnotatorsByIdOptions *GetAnnotatorsByIdOptions)`, func() {
		version := "testString"
		getAnnotatorsByIDPath := "/v1/annotators/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAnnotatorsByIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"description": "Description"}`)
				}))
			})
			It(`Invoke GetAnnotatorsByID successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAnnotatorsByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAnnotatorsByIdOptions model
				getAnnotatorsByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetAnnotatorsByIdOptions)
				getAnnotatorsByIdOptionsModel.ID = core.StringPtr("testString")
				getAnnotatorsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAnnotatorsByID(getAnnotatorsByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.Description).ToNot(BeNil())
			})
			It(`Invoke GetAnnotatorsByID with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAnnotatorsByIdOptions model
				getAnnotatorsByIdOptionsModel := new(annotatorforclinicaldataacdv1.GetAnnotatorsByIdOptions)
				getAnnotatorsByIdOptionsModel.ID = core.StringPtr("testString")
				getAnnotatorsByIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetAnnotatorsByID(getAnnotatorsByIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct a second instance of the GetAnnotatorsByIdOptions model with no property values
				getAnnotatorsByIdOptionsModelNew := new(annotatorforclinicaldataacdv1.GetAnnotatorsByIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetAnnotatorsByID(getAnnotatorsByIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteUserSpecificArtifacts(deleteUserSpecificArtifactsOptions *DeleteUserSpecificArtifactsOptions)`, func() {
		version := "testString"
		deleteUserSpecificArtifactsPath := "/v1/user_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteUserSpecificArtifactsPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteUserSpecificArtifacts successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteUserSpecificArtifacts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteUserSpecificArtifactsOptions model
				deleteUserSpecificArtifactsOptionsModel := new(annotatorforclinicaldataacdv1.DeleteUserSpecificArtifactsOptions)
				deleteUserSpecificArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteUserSpecificArtifacts(deleteUserSpecificArtifactsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(response.StatusCode).To(Equal(204))
			})
			It(`Invoke DeleteUserSpecificArtifacts with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteUserSpecificArtifactsOptions model
				deleteUserSpecificArtifactsOptionsModel := new(annotatorforclinicaldataacdv1.DeleteUserSpecificArtifactsOptions)
				deleteUserSpecificArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteUserSpecificArtifacts(deleteUserSpecificArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "https://annotatorforclinicaldataacdv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CartridgesGet(cartridgesGetOptions *CartridgesGetOptions) - Operation response error`, func() {
		version := "testString"
		cartridgesGetPath := "/v1/cartridges"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesGetPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CartridgesGet with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesGetOptions model
				cartridgesGetOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesGetOptions)
				cartridgesGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CartridgesGet(cartridgesGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CartridgesGet(cartridgesGetOptions *CartridgesGetOptions)`, func() {
		version := "testString"
		cartridgesGetPath := "/v1/cartridges"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesGetPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"data": ["Data"]}`)
				}))
			})
			It(`Invoke CartridgesGet successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CartridgesGet(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CartridgesGetOptions model
				cartridgesGetOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesGetOptions)
				cartridgesGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CartridgesGet(cartridgesGetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.Cartridges).ToNot(BeNil())
			})
			It(`Invoke CartridgesGet with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesGetOptions model
				cartridgesGetOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesGetOptions)
				cartridgesGetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CartridgesGet(cartridgesGetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CartridgesPostMultipart(cartridgesPostMultipartOptions *CartridgesPostMultipartOptions) - Operation response error`, func() {
		version := "testString"
		cartridgesPostMultipartPath := "/v1/cartridges"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesPostMultipartPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CartridgesPostMultipart with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesPostMultipartOptions model
				cartridgesPostMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPostMultipartOptions)
				cartridgesPostMultipartOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				cartridgesPostMultipartOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				cartridgesPostMultipartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CartridgesPostMultipart(cartridgesPostMultipartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CartridgesPostMultipart(cartridgesPostMultipartOptions *CartridgesPostMultipartOptions)`, func() {
		version := "testString"
		cartridgesPostMultipartPath := "/v1/cartridges"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesPostMultipartPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"code": 4, "artifactResponse": [{"code": 4, "message": "Message", "level": "ERROR", "description": "Description", "moreInfo": "MoreInfo", "correlationId": "CorrelationID", "artifact": "Artifact", "href": "Href"}]}`)
				}))
			})
			It(`Invoke CartridgesPostMultipart successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CartridgesPostMultipart(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CartridgesPostMultipartOptions model
				cartridgesPostMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPostMultipartOptions)
				cartridgesPostMultipartOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				cartridgesPostMultipartOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				cartridgesPostMultipartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CartridgesPostMultipart(cartridgesPostMultipartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.Code).ToNot(BeNil())

			})
			It(`Invoke CartridgesPostMultipart with error: Param validation error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesPostMultipartOptions model
				cartridgesPostMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPostMultipartOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := testService.CartridgesPostMultipart(cartridgesPostMultipartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke CartridgesPostMultipart with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesPostMultipartOptions model
				cartridgesPostMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPostMultipartOptions)
				cartridgesPostMultipartOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				cartridgesPostMultipartOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				cartridgesPostMultipartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CartridgesPostMultipart(cartridgesPostMultipartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CartridgesPutMultipart(cartridgesPutMultipartOptions *CartridgesPutMultipartOptions) - Operation response error`, func() {
		version := "testString"
		cartridgesPutMultipartPath := "/v1/cartridges"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesPutMultipartPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CartridgesPutMultipart with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesPutMultipartOptions model
				cartridgesPutMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPutMultipartOptions)
				cartridgesPutMultipartOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				cartridgesPutMultipartOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				cartridgesPutMultipartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CartridgesPutMultipart(cartridgesPutMultipartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CartridgesPutMultipart(cartridgesPutMultipartOptions *CartridgesPutMultipartOptions)`, func() {
		version := "testString"
		cartridgesPutMultipartPath := "/v1/cartridges"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesPutMultipartPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"code": 4, "artifactResponse": [{"code": 4, "message": "Message", "level": "ERROR", "description": "Description", "moreInfo": "MoreInfo", "correlationId": "CorrelationID", "artifact": "Artifact", "href": "Href"}]}`)
				}))
			})
			It(`Invoke CartridgesPutMultipart successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CartridgesPutMultipart(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CartridgesPutMultipartOptions model
				cartridgesPutMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPutMultipartOptions)
				cartridgesPutMultipartOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				cartridgesPutMultipartOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				cartridgesPutMultipartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CartridgesPutMultipart(cartridgesPutMultipartOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CartridgesPutMultipart with error: Param validation error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesPutMultipartOptions model
				cartridgesPutMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPutMultipartOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := testService.CartridgesPutMultipart(cartridgesPutMultipartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke CartridgesPutMultipart with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesPutMultipartOptions model
				cartridgesPutMultipartOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesPutMultipartOptions)
				cartridgesPutMultipartOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				cartridgesPutMultipartOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				cartridgesPutMultipartOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CartridgesPutMultipart(cartridgesPutMultipartOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CartridgesGetID(cartridgesGetIdOptions *CartridgesGetIdOptions) - Operation response error`, func() {
		version := "testString"
		cartridgesGetIDPath := "/v1/cartridges/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesGetIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CartridgesGetID with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesGetIdOptions model
				cartridgesGetIdOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesGetIdOptions)
				cartridgesGetIdOptionsModel.ID = core.StringPtr("testString")
				cartridgesGetIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CartridgesGetID(cartridgesGetIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CartridgesGetID(cartridgesGetIdOptions *CartridgesGetIdOptions)`, func() {
		version := "testString"
		cartridgesGetIDPath := "/v1/cartridges/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(cartridgesGetIDPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "name": "Name", "status": "Status", "statusCode": 10, "statusLocation": "StatusLocation", "startTime": "StartTime", "endTime": "EndTime", "duration": "Duration", "correlationId": "CorrelationID", "artifactResponseCode": 20, "artifactResponse": [{"code": 4, "message": "Message", "level": "ERROR", "description": "Description", "moreInfo": "MoreInfo", "correlationId": "CorrelationID", "artifact": "Artifact", "href": "Href"}]}`)
				}))
			})
			It(`Invoke CartridgesGetID successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CartridgesGetID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CartridgesGetIdOptions model
				cartridgesGetIdOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesGetIdOptions)
				cartridgesGetIdOptionsModel.ID = core.StringPtr("testString")
				cartridgesGetIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CartridgesGetID(cartridgesGetIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.ID).ToNot(BeNil())
			})
			It(`Invoke CartridgesGetID with error: Operation validation and request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CartridgesGetIdOptions model
				cartridgesGetIdOptionsModel := new(annotatorforclinicaldataacdv1.CartridgesGetIdOptions)
				cartridgesGetIdOptionsModel.ID = core.StringPtr("testString")
				cartridgesGetIdOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CartridgesGetID(cartridgesGetIdOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct a second instance of the CartridgesGetIdOptions model with no property values
				cartridgesGetIdOptionsModelNew := new(annotatorforclinicaldataacdv1.CartridgesGetIdOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CartridgesGetID(cartridgesGetIdOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeployCartridge(deployCartridgeOptions *DeployCartridgeOptions) - Operation response error`, func() {
		version := "testString"
		deployCartridgePath := "/v1/deploy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deployCartridgePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeployCartridge with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeployCartridgeOptions model
				deployCartridgeOptionsModel := new(annotatorforclinicaldataacdv1.DeployCartridgeOptions)
				deployCartridgeOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				deployCartridgeOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				deployCartridgeOptionsModel.Update = core.BoolPtr(true)
				deployCartridgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeployCartridge(deployCartridgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeployCartridge(deployCartridgeOptions *DeployCartridgeOptions)`, func() {
		version := "testString"
		deployCartridgePath := "/v1/deploy"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deployCartridgePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"code": 4, "artifactResponse": [{"code": 4, "message": "Message", "level": "ERROR", "description": "Description", "moreInfo": "MoreInfo", "correlationId": "CorrelationID", "artifact": "Artifact", "href": "Href"}]}`)
				}))
			})
			It(`Invoke DeployCartridge successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeployCartridge(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeployCartridgeOptions model
				deployCartridgeOptionsModel := new(annotatorforclinicaldataacdv1.DeployCartridgeOptions)
				deployCartridgeOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				deployCartridgeOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				deployCartridgeOptionsModel.Update = core.BoolPtr(true)
				deployCartridgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeployCartridge(deployCartridgeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeployCartridge with error: Param validation error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeployCartridgeOptions model
				deployCartridgeOptionsModel := new(annotatorforclinicaldataacdv1.DeployCartridgeOptions)
				// Invoke operation with invalid options model (negative test)
				result, response, operationErr := testService.DeployCartridge(deployCartridgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			It(`Invoke DeployCartridge with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeployCartridgeOptions model
				deployCartridgeOptionsModel := new(annotatorforclinicaldataacdv1.DeployCartridgeOptions)
				deployCartridgeOptionsModel.ArchiveFile = CreateMockReader("This is a mock file.")
				deployCartridgeOptionsModel.ArchiveFileContentType = core.StringPtr("testString")
				deployCartridgeOptionsModel.Update = core.BoolPtr(true)
				deployCartridgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeployCartridge(deployCartridgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "https://annotatorforclinicaldataacdv1/api",
				Version: core.StringPtr(version),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:     "https://testService/api",
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					Version: core.StringPtr(version),
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_URL":       "https://annotatorforclinicaldataacdv1/api",
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ANNOTATOR_FOR_CLINICAL_DATA_ACD_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions) - Operation response error`, func() {
		version := "testString"
		getHealthCheckStatusPath := "/v1/status/health_check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getHealthCheckStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"json"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHealthCheckStatus with error: Operation response processing error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHealthCheckStatusOptions model
				getHealthCheckStatusOptionsModel := new(annotatorforclinicaldataacdv1.GetHealthCheckStatusOptions)
				getHealthCheckStatusOptionsModel.Accept = core.StringPtr("application/json")
				getHealthCheckStatusOptionsModel.Format = core.StringPtr("json")
				getHealthCheckStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetHealthCheckStatus(getHealthCheckStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions)`, func() {
		version := "testString"
		getHealthCheckStatusPath := "/v1/status/health_check"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getHealthCheckStatusPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"json"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"version": "Version", "upTime": "UpTime", "serviceState": "OK", "stateDetails": "StateDetails", "hostName": "HostName", "requestCount": 12, "maxMemoryMb": 11, "commitedMemoryMb": 16, "inUseMemoryMb": 13, "availableProcessors": 19, "concurrentRequests": 18, "maxConcurrentRequests": 21, "totalRejectedRequests": 21, "totalBlockedRequests": 20}`)
				}))
			})
			It(`Invoke GetHealthCheckStatus successfully`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetHealthCheckStatus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHealthCheckStatusOptions model
				getHealthCheckStatusOptionsModel := new(annotatorforclinicaldataacdv1.GetHealthCheckStatusOptions)
				getHealthCheckStatusOptionsModel.Accept = core.StringPtr("application/json")
				getHealthCheckStatusOptionsModel.Format = core.StringPtr("json")
				getHealthCheckStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetHealthCheckStatus(getHealthCheckStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(result.ServiceState).ToNot(BeNil())
			})
			It(`Invoke GetHealthCheckStatus with error: Operation request error`, func() {
				testService, testServiceErr := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHealthCheckStatusOptions model
				getHealthCheckStatusOptionsModel := new(annotatorforclinicaldataacdv1.GetHealthCheckStatusOptions)
				getHealthCheckStatusOptionsModel.Accept = core.StringPtr("application/json")
				getHealthCheckStatusOptionsModel.Format = core.StringPtr("json")
				getHealthCheckStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetHealthCheckStatus(getHealthCheckStatusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			version := "testString"
			testService, _ := annotatorforclinicaldataacdv1.NewAnnotatorForClinicalDataAcdV1(&annotatorforclinicaldataacdv1.AnnotatorForClinicalDataAcdV1Options{
				URL:           "http://annotatorforclinicaldataacdv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAnnotator successfully`, func() {
				name := "testString"
				model, err := testService.NewAnnotator(name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAnnotatorFlow successfully`, func() {
				var flow *annotatorforclinicaldataacdv1.Flow = nil
				_, err := testService.NewAnnotatorFlow(flow)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCartridgesGetIdOptions successfully`, func() {
				// Construct an instance of the CartridgesGetIdOptions model
				id := "testString"
				cartridgesGetIdOptionsModel := testService.NewCartridgesGetIdOptions(id)
				cartridgesGetIdOptionsModel.SetID("testString")
				cartridgesGetIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cartridgesGetIdOptionsModel).ToNot(BeNil())
				Expect(cartridgesGetIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(cartridgesGetIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCartridgesGetOptions successfully`, func() {
				// Construct an instance of the CartridgesGetOptions model
				cartridgesGetOptionsModel := testService.NewCartridgesGetOptions()
				cartridgesGetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cartridgesGetOptionsModel).ToNot(BeNil())
				Expect(cartridgesGetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCartridgesPostMultipartOptions successfully`, func() {
				// Construct an instance of the CartridgesPostMultipartOptions model
				cartridgesPostMultipartOptionsModel := testService.NewCartridgesPostMultipartOptions()
				cartridgesPostMultipartOptionsModel.SetArchiveFile(CreateMockReader("This is a mock file."))
				cartridgesPostMultipartOptionsModel.SetArchiveFileContentType("testString")
				cartridgesPostMultipartOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cartridgesPostMultipartOptionsModel).ToNot(BeNil())
				Expect(cartridgesPostMultipartOptionsModel.ArchiveFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(cartridgesPostMultipartOptionsModel.ArchiveFileContentType).To(Equal(core.StringPtr("testString")))
				Expect(cartridgesPostMultipartOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCartridgesPutMultipartOptions successfully`, func() {
				// Construct an instance of the CartridgesPutMultipartOptions model
				cartridgesPutMultipartOptionsModel := testService.NewCartridgesPutMultipartOptions()
				cartridgesPutMultipartOptionsModel.SetArchiveFile(CreateMockReader("This is a mock file."))
				cartridgesPutMultipartOptionsModel.SetArchiveFileContentType("testString")
				cartridgesPutMultipartOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(cartridgesPutMultipartOptionsModel).ToNot(BeNil())
				Expect(cartridgesPutMultipartOptionsModel.ArchiveFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(cartridgesPutMultipartOptionsModel.ArchiveFileContentType).To(Equal(core.StringPtr("testString")))
				Expect(cartridgesPutMultipartOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateFlowsOptions successfully`, func() {
				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)
				Expect(flowEntryModel).ToNot(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				Expect(configurationEntityModel).ToNot(BeNil())
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))
				Expect(configurationEntityModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Uid).To(Equal(core.Int64Ptr(int64(26))))
				Expect(configurationEntityModel.Mergeid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				Expect(flowModel).ToNot(BeNil())
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)
				Expect(flowModel.Elements).To(Equal([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}))
				Expect(flowModel.Async).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				Expect(annotatorFlowModel).ToNot(BeNil())
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))
				Expect(annotatorFlowModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Flow).To(Equal(flowModel))
				Expect(annotatorFlowModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Data).To(Equal(make(map[string][]annotatorforclinicaldataacdv1.Entity)))
				Expect(annotatorFlowModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(annotatorFlowModel.GlobalConfigurations).To(Equal([]annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}))
				Expect(annotatorFlowModel.Uid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the CreateFlowsOptions model
				createFlowsOptionsModel := testService.NewCreateFlowsOptions()
				createFlowsOptionsModel.SetID("testString")
				createFlowsOptionsModel.SetName("testString")
				createFlowsOptionsModel.SetDescription("testString")
				createFlowsOptionsModel.SetPublishedDate("testString")
				createFlowsOptionsModel.SetPublish(true)
				createFlowsOptionsModel.SetVersion("testString")
				createFlowsOptionsModel.SetCartridgeID("testString")
				createFlowsOptionsModel.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel})
				createFlowsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createFlowsOptionsModel).ToNot(BeNil())
				Expect(createFlowsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createFlowsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createFlowsOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createFlowsOptionsModel.PublishedDate).To(Equal(core.StringPtr("testString")))
				Expect(createFlowsOptionsModel.Publish).To(Equal(core.BoolPtr(true)))
				Expect(createFlowsOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(createFlowsOptionsModel.CartridgeID).To(Equal(core.StringPtr("testString")))
				Expect(createFlowsOptionsModel.AnnotatorFlows).To(Equal([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}))
				Expect(createFlowsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProfileOptions successfully`, func() {
				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				Expect(configurationEntityModel).ToNot(BeNil())
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))
				Expect(configurationEntityModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Uid).To(Equal(core.Int64Ptr(int64(26))))
				Expect(configurationEntityModel.Mergeid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Annotator model
				annotatorModel := new(annotatorforclinicaldataacdv1.Annotator)
				Expect(annotatorModel).ToNot(BeNil())
				annotatorModel.Name = core.StringPtr("testString")
				annotatorModel.Parameters = make(map[string][]string)
				annotatorModel.Configurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				Expect(annotatorModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(annotatorModel.Parameters).To(Equal(make(map[string][]string)))
				Expect(annotatorModel.Configurations).To(Equal([]annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}))

				// Construct an instance of the CreateProfileOptions model
				createProfileOptionsModel := testService.NewCreateProfileOptions()
				createProfileOptionsModel.SetID("testString")
				createProfileOptionsModel.SetName("testString")
				createProfileOptionsModel.SetDescription("testString")
				createProfileOptionsModel.SetPublishedDate("testString")
				createProfileOptionsModel.SetPublish(true)
				createProfileOptionsModel.SetVersion("testString")
				createProfileOptionsModel.SetCartridgeID("testString")
				createProfileOptionsModel.SetAnnotators([]annotatorforclinicaldataacdv1.Annotator{*annotatorModel})
				createProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProfileOptionsModel).ToNot(BeNil())
				Expect(createProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.PublishedDate).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Publish).To(Equal(core.BoolPtr(true)))
				Expect(createProfileOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.CartridgeID).To(Equal(core.StringPtr("testString")))
				Expect(createProfileOptionsModel.Annotators).To(Equal([]annotatorforclinicaldataacdv1.Annotator{*annotatorModel}))
				Expect(createProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteFlowsOptions successfully`, func() {
				// Construct an instance of the DeleteFlowsOptions model
				id := "testString"
				deleteFlowsOptionsModel := testService.NewDeleteFlowsOptions(id)
				deleteFlowsOptionsModel.SetID("testString")
				deleteFlowsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteFlowsOptionsModel).ToNot(BeNil())
				Expect(deleteFlowsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteFlowsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProfileOptions successfully`, func() {
				// Construct an instance of the DeleteProfileOptions model
				id := "testString"
				deleteProfileOptionsModel := testService.NewDeleteProfileOptions(id)
				deleteProfileOptionsModel.SetID("testString")
				deleteProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProfileOptionsModel).ToNot(BeNil())
				Expect(deleteProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteUserSpecificArtifactsOptions successfully`, func() {
				// Construct an instance of the DeleteUserSpecificArtifactsOptions model
				deleteUserSpecificArtifactsOptionsModel := testService.NewDeleteUserSpecificArtifactsOptions()
				deleteUserSpecificArtifactsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteUserSpecificArtifactsOptionsModel).ToNot(BeNil())
				Expect(deleteUserSpecificArtifactsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeployCartridgeOptions successfully`, func() {
				// Construct an instance of the DeployCartridgeOptions model
				deployCartridgeOptionsModel := testService.NewDeployCartridgeOptions()
				deployCartridgeOptionsModel.SetArchiveFile(CreateMockReader("This is a mock file."))
				deployCartridgeOptionsModel.SetArchiveFileContentType("testString")
				deployCartridgeOptionsModel.SetUpdate(true)
				deployCartridgeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deployCartridgeOptionsModel).ToNot(BeNil())
				Expect(deployCartridgeOptionsModel.ArchiveFile).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(deployCartridgeOptionsModel.ArchiveFileContentType).To(Equal(core.StringPtr("testString")))
				Expect(deployCartridgeOptionsModel.Update).To(Equal(core.BoolPtr(true)))
				Expect(deployCartridgeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAnnotatorsByIdOptions successfully`, func() {
				// Construct an instance of the GetAnnotatorsByIdOptions model
				id := "testString"
				getAnnotatorsByIdOptionsModel := testService.NewGetAnnotatorsByIdOptions(id)
				getAnnotatorsByIdOptionsModel.SetID("testString")
				getAnnotatorsByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAnnotatorsByIdOptionsModel).ToNot(BeNil())
				Expect(getAnnotatorsByIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAnnotatorsByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAnnotatorsOptions successfully`, func() {
				// Construct an instance of the GetAnnotatorsOptions model
				getAnnotatorsOptionsModel := testService.NewGetAnnotatorsOptions()
				getAnnotatorsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAnnotatorsOptionsModel).ToNot(BeNil())
				Expect(getAnnotatorsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFlowsByIdOptions successfully`, func() {
				// Construct an instance of the GetFlowsByIdOptions model
				id := "testString"
				getFlowsByIdOptionsModel := testService.NewGetFlowsByIdOptions(id)
				getFlowsByIdOptionsModel.SetID("testString")
				getFlowsByIdOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFlowsByIdOptionsModel).ToNot(BeNil())
				Expect(getFlowsByIdOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getFlowsByIdOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFlowsOptions successfully`, func() {
				// Construct an instance of the GetFlowsOptions model
				getFlowsOptionsModel := testService.NewGetFlowsOptions()
				getFlowsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFlowsOptionsModel).ToNot(BeNil())
				Expect(getFlowsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetHealthCheckStatusOptions successfully`, func() {
				// Construct an instance of the GetHealthCheckStatusOptions model
				getHealthCheckStatusOptionsModel := testService.NewGetHealthCheckStatusOptions()
				getHealthCheckStatusOptionsModel.SetAccept("application/json")
				getHealthCheckStatusOptionsModel.SetFormat("json")
				getHealthCheckStatusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHealthCheckStatusOptionsModel).ToNot(BeNil())
				Expect(getHealthCheckStatusOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(getHealthCheckStatusOptionsModel.Format).To(Equal(core.StringPtr("json")))
				Expect(getHealthCheckStatusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfileOptions successfully`, func() {
				// Construct an instance of the GetProfileOptions model
				id := "testString"
				getProfileOptionsModel := testService.NewGetProfilesByIdOptions(id)
				getProfileOptionsModel.SetID("testString")
				getProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfileOptionsModel).ToNot(BeNil())
				Expect(getProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProfilesOptions successfully`, func() {
				// Construct an instance of the GetProfilesOptions model
				getProfilesOptionsModel := testService.NewGetProfilesOptions()
				getProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProfilesOptionsModel).ToNot(BeNil())
				Expect(getProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunPipelineOptions successfully`, func() {
				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)
				Expect(flowEntryModel).ToNot(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				Expect(configurationEntityModel).ToNot(BeNil())
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))
				Expect(configurationEntityModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Uid).To(Equal(core.Int64Ptr(int64(26))))
				Expect(configurationEntityModel.Mergeid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				Expect(flowModel).ToNot(BeNil())
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)
				Expect(flowModel.Elements).To(Equal([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}))
				Expect(flowModel.Async).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				Expect(annotatorFlowModel).ToNot(BeNil())
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))
				Expect(annotatorFlowModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Flow).To(Equal(flowModel))
				Expect(annotatorFlowModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Data).To(Equal(make(map[string][]annotatorforclinicaldataacdv1.Entity)))
				Expect(annotatorFlowModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(annotatorFlowModel.GlobalConfigurations).To(Equal([]annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}))
				Expect(annotatorFlowModel.Uid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of ContainerAnnotation model
				containerAnnotationModel := new(annotatorforclinicaldataacdv1.ContainerAnnotation)
				Expect(containerAnnotationModel).ToNot(BeNil())

				// Construct an instance of the UnstructuredContainer model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				Expect(unstructuredContainerModel).ToNot(BeNil())
				unstructuredContainerModel.Text = core.StringPtr("testString")
				unstructuredContainerModel.ID = core.StringPtr("testString")
				unstructuredContainerModel.Type = core.StringPtr("testString")
				unstructuredContainerModel.Data = containerAnnotationModel
				unstructuredContainerModel.Metadata = make(map[string]interface{})
				unstructuredContainerModel.Uid = core.Int64Ptr(int64(26))
				Expect(unstructuredContainerModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(unstructuredContainerModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unstructuredContainerModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(unstructuredContainerModel.Data).To(Equal(containerAnnotationModel))
				Expect(unstructuredContainerModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(unstructuredContainerModel.Uid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the RunPipelineOptions model
				runPipelineOptionsModel := testService.NewRunPipelineOptions()
				runPipelineOptionsModel.SetUnstructured([]annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel})
				runPipelineOptionsModel.SetAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel})
				runPipelineOptionsModel.SetDebugTextRestore(true)
				runPipelineOptionsModel.SetReturnAnalyzedText(true)
				runPipelineOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runPipelineOptionsModel).ToNot(BeNil())
				Expect(runPipelineOptionsModel.Unstructured).To(Equal([]annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}))
				Expect(runPipelineOptionsModel.AnnotatorFlows).To(Equal([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}))
				Expect(runPipelineOptionsModel.DebugTextRestore).To(Equal(core.BoolPtr(true)))
				Expect(runPipelineOptionsModel.ReturnAnalyzedText).To(Equal(core.BoolPtr(true)))
				Expect(runPipelineOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunPipelineWithFlowOptions successfully`, func() {
				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)
				Expect(flowEntryModel).ToNot(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				Expect(configurationEntityModel).ToNot(BeNil())
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))
				Expect(configurationEntityModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Uid).To(Equal(core.Int64Ptr(int64(26))))
				Expect(configurationEntityModel.Mergeid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				Expect(flowModel).ToNot(BeNil())
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)
				Expect(flowModel.Elements).To(Equal([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}))
				Expect(flowModel.Async).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				Expect(annotatorFlowModel).ToNot(BeNil())
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))
				Expect(annotatorFlowModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Flow).To(Equal(flowModel))
				Expect(annotatorFlowModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Data).To(Equal(make(map[string][]annotatorforclinicaldataacdv1.Entity)))
				Expect(annotatorFlowModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(annotatorFlowModel.GlobalConfigurations).To(Equal([]annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}))
				Expect(annotatorFlowModel.Uid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of ContainerAnnotation model
				containerAnnotationModel := new(annotatorforclinicaldataacdv1.ContainerAnnotation)
				Expect(containerAnnotationModel).ToNot(BeNil())

				// Construct an instance of the UnstructuredContainer model
				unstructuredContainerModel := new(annotatorforclinicaldataacdv1.UnstructuredContainer)
				Expect(unstructuredContainerModel).ToNot(BeNil())
				unstructuredContainerModel.Text = core.StringPtr("testString")
				unstructuredContainerModel.ID = core.StringPtr("testString")
				unstructuredContainerModel.Type = core.StringPtr("testString")
				unstructuredContainerModel.Data = containerAnnotationModel
				unstructuredContainerModel.Metadata = make(map[string]interface{})
				unstructuredContainerModel.Uid = core.Int64Ptr(int64(26))
				Expect(unstructuredContainerModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(unstructuredContainerModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unstructuredContainerModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(unstructuredContainerModel.Data).To(Equal(containerAnnotationModel))
				Expect(unstructuredContainerModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(unstructuredContainerModel.Uid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the AnalyticFlowBeanInput model
				analyticFlowBeanInputModel := new(annotatorforclinicaldataacdv1.AnalyticFlowBeanInput)
				Expect(analyticFlowBeanInputModel).ToNot(BeNil())
				analyticFlowBeanInputModel.Unstructured = []annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}
				analyticFlowBeanInputModel.AnnotatorFlows = []annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}
				Expect(analyticFlowBeanInputModel.Unstructured).To(Equal([]annotatorforclinicaldataacdv1.UnstructuredContainer{*unstructuredContainerModel}))
				Expect(analyticFlowBeanInputModel.AnnotatorFlows).To(Equal([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}))

				// Construct an instance of the RunPipelineWithFlowOptions model
				flowID := "testString"
				returnAnalyzedText := true
				runPipelineWithFlowOptionsModel := testService.NewRunPipelineWithFlowOptions(flowID, returnAnalyzedText)
				runPipelineWithFlowOptionsModel.SetFlowID("testString")
				runPipelineWithFlowOptionsModel.SetReturnAnalyzedText(true)
				runPipelineWithFlowOptionsModel.SetAnalyticFlowBeanInput(analyticFlowBeanInputModel)
				runPipelineWithFlowOptionsModel.SetBody("testString")
				runPipelineWithFlowOptionsModel.SetContentType("application/json")
				runPipelineWithFlowOptionsModel.SetDebugTextRestore(true)
				runPipelineWithFlowOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runPipelineWithFlowOptionsModel).ToNot(BeNil())
				Expect(runPipelineWithFlowOptionsModel.FlowID).To(Equal(core.StringPtr("testString")))
				Expect(runPipelineWithFlowOptionsModel.ReturnAnalyzedText).To(Equal(core.BoolPtr(true)))
				Expect(runPipelineWithFlowOptionsModel.AnalyticFlowBeanInput).To(Equal(analyticFlowBeanInputModel))
				Expect(runPipelineWithFlowOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(runPipelineWithFlowOptionsModel.ContentType).To(Equal(core.StringPtr("application/json")))
				Expect(runPipelineWithFlowOptionsModel.DebugTextRestore).To(Equal(core.BoolPtr(true)))
				Expect(runPipelineWithFlowOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateFlowsOptions successfully`, func() {
				// Construct an instance of the FlowEntry model
				flowEntryModel := new(annotatorforclinicaldataacdv1.FlowEntry)
				Expect(flowEntryModel).ToNot(BeNil())

				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				Expect(configurationEntityModel).ToNot(BeNil())
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))
				Expect(configurationEntityModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Uid).To(Equal(core.Int64Ptr(int64(26))))
				Expect(configurationEntityModel.Mergeid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Flow model
				flowModel := new(annotatorforclinicaldataacdv1.Flow)
				Expect(flowModel).ToNot(BeNil())
				flowModel.Elements = []annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}
				flowModel.Async = core.BoolPtr(true)
				Expect(flowModel.Elements).To(Equal([]annotatorforclinicaldataacdv1.FlowEntry{*flowEntryModel}))
				Expect(flowModel.Async).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the AnnotatorFlow model
				annotatorFlowModel := new(annotatorforclinicaldataacdv1.AnnotatorFlow)
				Expect(annotatorFlowModel).ToNot(BeNil())
				annotatorFlowModel.Profile = core.StringPtr("testString")
				annotatorFlowModel.Flow = flowModel
				annotatorFlowModel.ID = core.StringPtr("testString")
				annotatorFlowModel.Type = core.StringPtr("testString")
				annotatorFlowModel.Data = make(map[string][]annotatorforclinicaldataacdv1.Entity)
				annotatorFlowModel.Metadata = make(map[string]interface{})
				annotatorFlowModel.GlobalConfigurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				annotatorFlowModel.Uid = core.Int64Ptr(int64(26))
				Expect(annotatorFlowModel.Profile).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Flow).To(Equal(flowModel))
				Expect(annotatorFlowModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(annotatorFlowModel.Data).To(Equal(make(map[string][]annotatorforclinicaldataacdv1.Entity)))
				Expect(annotatorFlowModel.Metadata).To(Equal(make(map[string]interface{})))
				Expect(annotatorFlowModel.GlobalConfigurations).To(Equal([]annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}))
				Expect(annotatorFlowModel.Uid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the UpdateFlowsOptions model
				id := "testString"
				updateFlowsOptionsModel := testService.NewUpdateFlowsOptions(id)
				updateFlowsOptionsModel.SetID("testString")
				updateFlowsOptionsModel.SetNewID("testString")
				updateFlowsOptionsModel.SetNewName("testString")
				updateFlowsOptionsModel.SetNewDescription("testString")
				updateFlowsOptionsModel.SetNewPublishedDate("testString")
				updateFlowsOptionsModel.SetNewPublish(true)
				updateFlowsOptionsModel.SetNewVersion("testString")
				updateFlowsOptionsModel.SetNewCartridgeID("testString")
				updateFlowsOptionsModel.SetNewAnnotatorFlows([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel})
				updateFlowsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateFlowsOptionsModel).ToNot(BeNil())
				Expect(updateFlowsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewID).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewName).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewPublishedDate).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewPublish).To(Equal(core.BoolPtr(true)))
				Expect(updateFlowsOptionsModel.NewVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewCartridgeID).To(Equal(core.StringPtr("testString")))
				Expect(updateFlowsOptionsModel.NewAnnotatorFlows).To(Equal([]annotatorforclinicaldataacdv1.AnnotatorFlow{*annotatorFlowModel}))
				Expect(updateFlowsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProfileOptions successfully`, func() {
				// Construct an instance of the ConfigurationEntity model
				configurationEntityModel := new(annotatorforclinicaldataacdv1.ConfigurationEntity)
				Expect(configurationEntityModel).ToNot(BeNil())
				configurationEntityModel.ID = core.StringPtr("testString")
				configurationEntityModel.Type = core.StringPtr("testString")
				configurationEntityModel.Uid = core.Int64Ptr(int64(26))
				configurationEntityModel.Mergeid = core.Int64Ptr(int64(26))
				Expect(configurationEntityModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(configurationEntityModel.Uid).To(Equal(core.Int64Ptr(int64(26))))
				Expect(configurationEntityModel.Mergeid).To(Equal(core.Int64Ptr(int64(26))))

				// Construct an instance of the Annotator model
				annotatorModel := new(annotatorforclinicaldataacdv1.Annotator)
				Expect(annotatorModel).ToNot(BeNil())
				annotatorModel.Name = core.StringPtr("testString")
				annotatorModel.Parameters = make(map[string][]string)
				annotatorModel.Configurations = []annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}
				Expect(annotatorModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(annotatorModel.Parameters).To(Equal(make(map[string][]string)))
				Expect(annotatorModel.Configurations).To(Equal([]annotatorforclinicaldataacdv1.ConfigurationEntity{*configurationEntityModel}))

				// Construct an instance of the UpdateProfileOptions model
				id := "testString"
				updateProfileOptionsModel := testService.NewUpdateProfileOptions(id)
				updateProfileOptionsModel.SetID("testString")
				updateProfileOptionsModel.SetNewID("testString")
				updateProfileOptionsModel.SetNewName("testString")
				updateProfileOptionsModel.SetNewDescription("testString")
				updateProfileOptionsModel.SetNewPublishedDate("testString")
				updateProfileOptionsModel.SetNewPublish(true)
				updateProfileOptionsModel.SetNewVersion("testString")
				updateProfileOptionsModel.SetNewCartridgeID("testString")
				updateProfileOptionsModel.SetNewAnnotators([]annotatorforclinicaldataacdv1.Annotator{*annotatorModel})
				updateProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProfileOptionsModel).ToNot(BeNil())
				Expect(updateProfileOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewName).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewPublishedDate).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewPublish).To(Equal(core.BoolPtr(true)))
				Expect(updateProfileOptionsModel.NewVersion).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewCartridgeID).To(Equal(core.StringPtr("testString")))
				Expect(updateProfileOptionsModel.NewAnnotators).To(Equal([]annotatorforclinicaldataacdv1.Annotator{*annotatorModel}))
				Expect(updateProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
