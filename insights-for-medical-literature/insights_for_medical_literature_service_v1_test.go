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

package insightsformedicalliteratureservicev1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/whcs-go-sdk/insightsformedicalliteratureservicev1"
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

var _ = Describe(`InsightsForMedicalLiteratureServiceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "https://insightsformedicalliteratureservicev1/api",
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
	Describe(`GetDocuments(getDocumentsOptions *GetDocumentsOptions) - Operation response error`, func() {
		version := "testString"
		getDocumentsPath := "/v1/corpora/testString/documents"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocuments with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentsOptions model
				getDocumentsOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentsOptions)
				getDocumentsOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetDocuments(getDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocuments(getDocumentsOptions *GetDocumentsOptions)`, func() {
		version := "testString"
		getDocumentsPath := "/v1/corpora/testString/documents"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"corpusName": "CorpusName", "ontologies": ["Ontologies"], "descriptiveName": "DescriptiveName", "bvt": false, "elasticsearchIndex": "ElasticsearchIndex"}`)
				}))
			})
			It(`Invoke GetDocuments successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDocuments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentsOptions model
				getDocumentsOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentsOptions)
				getDocumentsOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDocuments(getDocumentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocuments with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentsOptions model
				getDocumentsOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentsOptions)
				getDocumentsOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDocuments(getDocumentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentsOptions model with no property values
				getDocumentsOptionsModelNew := new(insightsformedicalliteratureservicev1.GetDocumentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetDocuments(getDocumentsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddCorpusDocument(addCorpusDocumentOptions *AddCorpusDocumentOptions)`, func() {
		version := "testString"
		addCorpusDocumentPath := "/v1/corpora/testString/documents"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addCorpusDocumentPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddCorpusDocument successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.AddCorpusDocument(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AddCorpusDocumentOptions model
				addCorpusDocumentOptionsModel := new(insightsformedicalliteratureservicev1.AddCorpusDocumentOptions)
				addCorpusDocumentOptionsModel.Corpus = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.Document = make(map[string]interface{})
				addCorpusDocumentOptionsModel.AcdURL = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.ApiKey = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.FlowID = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.AccessToken = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.OtherAnnotators = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				addCorpusDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.AddCorpusDocument(addCorpusDocumentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddCorpusDocument with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AddCorpusDocumentOptions model
				addCorpusDocumentOptionsModel := new(insightsformedicalliteratureservicev1.AddCorpusDocumentOptions)
				addCorpusDocumentOptionsModel.Corpus = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.Document = make(map[string]interface{})
				addCorpusDocumentOptionsModel.AcdURL = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.ApiKey = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.FlowID = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.AccessToken = core.StringPtr("testString")
				addCorpusDocumentOptionsModel.OtherAnnotators = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				addCorpusDocumentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.AddCorpusDocument(addCorpusDocumentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddCorpusDocumentOptions model with no property values
				addCorpusDocumentOptionsModelNew := new(insightsformedicalliteratureservicev1.AddCorpusDocumentOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.AddCorpusDocument(addCorpusDocumentOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocumentInfo(getDocumentInfoOptions *GetDocumentInfoOptions) - Operation response error`, func() {
		version := "testString"
		getDocumentInfoPath := "/v1/corpora/testString/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocumentInfo with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentInfoOptions model
				getDocumentInfoOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentInfoOptions)
				getDocumentInfoOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentInfoOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentInfoOptionsModel.Verbose = core.BoolPtr(true)
				getDocumentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetDocumentInfo(getDocumentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentInfo(getDocumentInfoOptions *GetDocumentInfoOptions)`, func() {
		version := "testString"
		getDocumentInfoPath := "/v1/corpora/testString/documents/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{}`)
				}))
			})
			It(`Invoke GetDocumentInfo successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDocumentInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentInfoOptions model
				getDocumentInfoOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentInfoOptions)
				getDocumentInfoOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentInfoOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentInfoOptionsModel.Verbose = core.BoolPtr(true)
				getDocumentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDocumentInfo(getDocumentInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocumentInfo with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentInfoOptions model
				getDocumentInfoOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentInfoOptions)
				getDocumentInfoOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentInfoOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentInfoOptionsModel.Verbose = core.BoolPtr(true)
				getDocumentInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDocumentInfo(getDocumentInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentInfoOptions model with no property values
				getDocumentInfoOptionsModelNew := new(insightsformedicalliteratureservicev1.GetDocumentInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetDocumentInfo(getDocumentInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentAnnotations(getDocumentAnnotationsOptions *GetDocumentAnnotationsOptions)`, func() {
		version := "testString"
		getDocumentAnnotationsPath := "/v1/corpora/testString/documents/testString/annotations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentAnnotationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["document_section"]).To(Equal([]string{"testString"}))

					// TODO: Add check for include_text query parameter

					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDocumentAnnotations successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetDocumentAnnotations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetDocumentAnnotationsOptions model
				getDocumentAnnotationsOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentAnnotationsOptions)
				getDocumentAnnotationsOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentAnnotationsOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentAnnotationsOptionsModel.DocumentSection = core.StringPtr("testString")
				getDocumentAnnotationsOptionsModel.Cuis = []string{"testString"}
				getDocumentAnnotationsOptionsModel.IncludeText = core.BoolPtr(true)
				getDocumentAnnotationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetDocumentAnnotations(getDocumentAnnotationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke GetDocumentAnnotations with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentAnnotationsOptions model
				getDocumentAnnotationsOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentAnnotationsOptions)
				getDocumentAnnotationsOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentAnnotationsOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentAnnotationsOptionsModel.DocumentSection = core.StringPtr("testString")
				getDocumentAnnotationsOptionsModel.Cuis = []string{"testString"}
				getDocumentAnnotationsOptionsModel.IncludeText = core.BoolPtr(true)
				getDocumentAnnotationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.GetDocumentAnnotations(getDocumentAnnotationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetDocumentAnnotationsOptions model with no property values
				getDocumentAnnotationsOptionsModelNew := new(insightsformedicalliteratureservicev1.GetDocumentAnnotationsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.GetDocumentAnnotations(getDocumentAnnotationsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocumentCategories(getDocumentCategoriesOptions *GetDocumentCategoriesOptions) - Operation response error`, func() {
		version := "testString"
		getDocumentCategoriesPath := "/v1/corpora/testString/documents/testString/categories"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentCategoriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["category"]).To(Equal([]string{"disorders"}))

					// TODO: Add check for only_negated_concepts query parameter

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocumentCategories with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentCategoriesOptions model
				getDocumentCategoriesOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentCategoriesOptions)
				getDocumentCategoriesOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.HighlightTagBegin = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.HighlightTagEnd = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.Types = []string{"testString"}
				getDocumentCategoriesOptionsModel.Category = core.StringPtr("disorders")
				getDocumentCategoriesOptionsModel.OnlyNegatedConcepts = core.BoolPtr(true)
				getDocumentCategoriesOptionsModel.Fields = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getDocumentCategoriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetDocumentCategories(getDocumentCategoriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentCategories(getDocumentCategoriesOptions *GetDocumentCategoriesOptions)`, func() {
		version := "testString"
		getDocumentCategoriesPath := "/v1/corpora/testString/documents/testString/categories"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentCategoriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["category"]).To(Equal([]string{"disorders"}))

					// TODO: Add check for only_negated_concepts query parameter

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"modelLicense": "ModelLicense", "highlightedTitle": {}, "highlightedAbstract": {}, "highlightedBody": {}, "highlightedSections": {"mapKey": {}}, "passages": {"mapKey": {"mapKey": {"id": "ID", "negated": false, "sentences": [{"documentSection": "DocumentSection", "text": {}, "begin": 5, "end": 3, "timestamp": 9}]}}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}`)
				}))
			})
			It(`Invoke GetDocumentCategories successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDocumentCategories(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDocumentCategoriesOptions model
				getDocumentCategoriesOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentCategoriesOptions)
				getDocumentCategoriesOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.HighlightTagBegin = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.HighlightTagEnd = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.Types = []string{"testString"}
				getDocumentCategoriesOptionsModel.Category = core.StringPtr("disorders")
				getDocumentCategoriesOptionsModel.OnlyNegatedConcepts = core.BoolPtr(true)
				getDocumentCategoriesOptionsModel.Fields = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getDocumentCategoriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDocumentCategories(getDocumentCategoriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocumentCategories with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetDocumentCategoriesOptions model
				getDocumentCategoriesOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentCategoriesOptions)
				getDocumentCategoriesOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.HighlightTagBegin = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.HighlightTagEnd = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.Types = []string{"testString"}
				getDocumentCategoriesOptionsModel.Category = core.StringPtr("disorders")
				getDocumentCategoriesOptionsModel.OnlyNegatedConcepts = core.BoolPtr(true)
				getDocumentCategoriesOptionsModel.Fields = core.StringPtr("testString")
				getDocumentCategoriesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getDocumentCategoriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDocumentCategories(getDocumentCategoriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentCategoriesOptions model with no property values
				getDocumentCategoriesOptionsModelNew := new(insightsformedicalliteratureservicev1.GetDocumentCategoriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetDocumentCategories(getDocumentCategoriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptions *GetDocumentMultipleCategoriesOptions) - Operation response error`, func() {
		version := "testString"
		getDocumentMultipleCategoriesPath := "/v1/corpora/testString/documents/testString/categories"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentMultipleCategoriesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDocumentMultipleCategories with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AnnotationModel model
				annotationModelModel := new(insightsformedicalliteratureservicev1.AnnotationModel)
				annotationModelModel.UniqueID = core.Int64Ptr(int64(38))
				annotationModelModel.StickyIds = []int64{int64(38)}
				annotationModelModel.Ontology = core.StringPtr("testString")
				annotationModelModel.Section = core.StringPtr("testString")
				annotationModelModel.PreferredName = core.StringPtr("testString")
				annotationModelModel.Cui = core.StringPtr("testString")
				annotationModelModel.AttributeID = core.StringPtr("testString")
				annotationModelModel.Qualifiers = []string{"testString"}
				annotationModelModel.Type = core.StringPtr("testString")
				annotationModelModel.Negated = core.BoolPtr(true)
				annotationModelModel.Hypothetical = core.BoolPtr(true)
				annotationModelModel.Unit = core.StringPtr("testString")
				annotationModelModel.MinValue = core.StringPtr("testString")
				annotationModelModel.MaxValue = core.StringPtr("testString")
				annotationModelModel.Operator = core.StringPtr("testString")
				annotationModelModel.NluSourceType = core.StringPtr("testString")
				annotationModelModel.NluRelation = core.StringPtr("testString")
				annotationModelModel.NluTargetType = core.StringPtr("testString")
				annotationModelModel.NluEntityIndex = core.StringPtr("testString")
				annotationModelModel.NluMentionIndex = core.StringPtr("testString")
				annotationModelModel.NluRelationID = core.StringPtr("testString")
				annotationModelModel.NluSide = core.StringPtr("testString")
				annotationModelModel.Begin = core.Int64Ptr(int64(38))
				annotationModelModel.End = core.Int64Ptr(int64(38))
				annotationModelModel.Score = core.Float32Ptr(36.0)
				annotationModelModel.Timestamp = core.Int64Ptr(int64(26))
				annotationModelModel.Features = make(map[string]string)
				annotationModelModel.Hits = core.Int64Ptr(int64(38))

				// Construct an instance of the StringBuilder model
				stringBuilderModel := new(insightsformedicalliteratureservicev1.StringBuilder)

				// Construct an instance of the GetDocumentMultipleCategoriesOptions model
				getDocumentMultipleCategoriesOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentMultipleCategoriesOptions)
				getDocumentMultipleCategoriesOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.ModelLicense = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.HighlightedTitle = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedAbstract = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedBody = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedSections = make(map[string]insightsformedicalliteratureservicev1.StringBuilder)
				getDocumentMultipleCategoriesOptionsModel.Passages = make(map[string]map[string]EntryModel)
				getDocumentMultipleCategoriesOptionsModel.Annotations = make(map[string]insightsformedicalliteratureservicev1.AnnotationModel)
				getDocumentMultipleCategoriesOptionsModel.HighlightTagBegin = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.HighlightTagEnd = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.Fields = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getDocumentMultipleCategoriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptions *GetDocumentMultipleCategoriesOptions)`, func() {
		version := "testString"
		getDocumentMultipleCategoriesPath := "/v1/corpora/testString/documents/testString/categories"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getDocumentMultipleCategoriesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["highlight_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"modelLicense": "ModelLicense", "highlightedTitle": {}, "highlightedAbstract": {}, "highlightedBody": {}, "highlightedSections": {"mapKey": {}}, "passages": {"mapKey": {"mapKey": {"id": "ID", "negated": false, "sentences": [{"documentSection": "DocumentSection", "text": {}, "begin": 5, "end": 3, "timestamp": 9}]}}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}`)
				}))
			})
			It(`Invoke GetDocumentMultipleCategories successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetDocumentMultipleCategories(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AnnotationModel model
				annotationModelModel := new(insightsformedicalliteratureservicev1.AnnotationModel)
				annotationModelModel.UniqueID = core.Int64Ptr(int64(38))
				annotationModelModel.StickyIds = []int64{int64(38)}
				annotationModelModel.Ontology = core.StringPtr("testString")
				annotationModelModel.Section = core.StringPtr("testString")
				annotationModelModel.PreferredName = core.StringPtr("testString")
				annotationModelModel.Cui = core.StringPtr("testString")
				annotationModelModel.AttributeID = core.StringPtr("testString")
				annotationModelModel.Qualifiers = []string{"testString"}
				annotationModelModel.Type = core.StringPtr("testString")
				annotationModelModel.Negated = core.BoolPtr(true)
				annotationModelModel.Hypothetical = core.BoolPtr(true)
				annotationModelModel.Unit = core.StringPtr("testString")
				annotationModelModel.MinValue = core.StringPtr("testString")
				annotationModelModel.MaxValue = core.StringPtr("testString")
				annotationModelModel.Operator = core.StringPtr("testString")
				annotationModelModel.NluSourceType = core.StringPtr("testString")
				annotationModelModel.NluRelation = core.StringPtr("testString")
				annotationModelModel.NluTargetType = core.StringPtr("testString")
				annotationModelModel.NluEntityIndex = core.StringPtr("testString")
				annotationModelModel.NluMentionIndex = core.StringPtr("testString")
				annotationModelModel.NluRelationID = core.StringPtr("testString")
				annotationModelModel.NluSide = core.StringPtr("testString")
				annotationModelModel.Begin = core.Int64Ptr(int64(38))
				annotationModelModel.End = core.Int64Ptr(int64(38))
				annotationModelModel.Score = core.Float32Ptr(36.0)
				annotationModelModel.Timestamp = core.Int64Ptr(int64(26))
				annotationModelModel.Features = make(map[string]string)
				annotationModelModel.Hits = core.Int64Ptr(int64(38))

				// Construct an instance of the StringBuilder model
				stringBuilderModel := new(insightsformedicalliteratureservicev1.StringBuilder)

				// Construct an instance of the GetDocumentMultipleCategoriesOptions model
				getDocumentMultipleCategoriesOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentMultipleCategoriesOptions)
				getDocumentMultipleCategoriesOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.ModelLicense = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.HighlightedTitle = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedAbstract = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedBody = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedSections = make(map[string]insightsformedicalliteratureservicev1.StringBuilder)
				getDocumentMultipleCategoriesOptionsModel.Passages = make(map[string]map[string]EntryModel)
				getDocumentMultipleCategoriesOptionsModel.Annotations = make(map[string]insightsformedicalliteratureservicev1.AnnotationModel)
				getDocumentMultipleCategoriesOptionsModel.HighlightTagBegin = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.HighlightTagEnd = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.Fields = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getDocumentMultipleCategoriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetDocumentMultipleCategories with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the AnnotationModel model
				annotationModelModel := new(insightsformedicalliteratureservicev1.AnnotationModel)
				annotationModelModel.UniqueID = core.Int64Ptr(int64(38))
				annotationModelModel.StickyIds = []int64{int64(38)}
				annotationModelModel.Ontology = core.StringPtr("testString")
				annotationModelModel.Section = core.StringPtr("testString")
				annotationModelModel.PreferredName = core.StringPtr("testString")
				annotationModelModel.Cui = core.StringPtr("testString")
				annotationModelModel.AttributeID = core.StringPtr("testString")
				annotationModelModel.Qualifiers = []string{"testString"}
				annotationModelModel.Type = core.StringPtr("testString")
				annotationModelModel.Negated = core.BoolPtr(true)
				annotationModelModel.Hypothetical = core.BoolPtr(true)
				annotationModelModel.Unit = core.StringPtr("testString")
				annotationModelModel.MinValue = core.StringPtr("testString")
				annotationModelModel.MaxValue = core.StringPtr("testString")
				annotationModelModel.Operator = core.StringPtr("testString")
				annotationModelModel.NluSourceType = core.StringPtr("testString")
				annotationModelModel.NluRelation = core.StringPtr("testString")
				annotationModelModel.NluTargetType = core.StringPtr("testString")
				annotationModelModel.NluEntityIndex = core.StringPtr("testString")
				annotationModelModel.NluMentionIndex = core.StringPtr("testString")
				annotationModelModel.NluRelationID = core.StringPtr("testString")
				annotationModelModel.NluSide = core.StringPtr("testString")
				annotationModelModel.Begin = core.Int64Ptr(int64(38))
				annotationModelModel.End = core.Int64Ptr(int64(38))
				annotationModelModel.Score = core.Float32Ptr(36.0)
				annotationModelModel.Timestamp = core.Int64Ptr(int64(26))
				annotationModelModel.Features = make(map[string]string)
				annotationModelModel.Hits = core.Int64Ptr(int64(38))

				// Construct an instance of the StringBuilder model
				stringBuilderModel := new(insightsformedicalliteratureservicev1.StringBuilder)

				// Construct an instance of the GetDocumentMultipleCategoriesOptions model
				getDocumentMultipleCategoriesOptionsModel := new(insightsformedicalliteratureservicev1.GetDocumentMultipleCategoriesOptions)
				getDocumentMultipleCategoriesOptionsModel.Corpus = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.DocumentID = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.ModelLicense = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.HighlightedTitle = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedAbstract = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedBody = stringBuilderModel
				getDocumentMultipleCategoriesOptionsModel.HighlightedSections = make(map[string]insightsformedicalliteratureservicev1.StringBuilder)
				getDocumentMultipleCategoriesOptionsModel.Passages = make(map[string]map[string]EntryModel)
				getDocumentMultipleCategoriesOptionsModel.Annotations = make(map[string]insightsformedicalliteratureservicev1.AnnotationModel)
				getDocumentMultipleCategoriesOptionsModel.HighlightTagBegin = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.HighlightTagEnd = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.Fields = core.StringPtr("testString")
				getDocumentMultipleCategoriesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getDocumentMultipleCategoriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDocumentMultipleCategoriesOptions model with no property values
				getDocumentMultipleCategoriesOptionsModelNew := new(insightsformedicalliteratureservicev1.GetDocumentMultipleCategoriesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSearchMatches(getSearchMatchesOptions *GetSearchMatchesOptions) - Operation response error`, func() {
		version := "testString"
		getSearchMatchesPath := "/v1/corpora/testString/documents/testString/search_matches"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSearchMatchesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for min_score query parameter

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["search_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["search_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["related_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["related_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSearchMatches with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSearchMatchesOptions model
				getSearchMatchesOptionsModel := new(insightsformedicalliteratureservicev1.GetSearchMatchesOptions)
				getSearchMatchesOptionsModel.Corpus = core.StringPtr("testString")
				getSearchMatchesOptionsModel.DocumentID = core.StringPtr("testString")
				getSearchMatchesOptionsModel.MinScore = core.Float32Ptr(36.0)
				getSearchMatchesOptionsModel.Cuis = []string{"testString"}
				getSearchMatchesOptionsModel.Text = []string{"testString"}
				getSearchMatchesOptionsModel.Types = []string{"testString"}
				getSearchMatchesOptionsModel.Attributes = []string{"testString"}
				getSearchMatchesOptionsModel.Values = []string{"testString"}
				getSearchMatchesOptionsModel.NluRelations = []string{"testString"}
				getSearchMatchesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getSearchMatchesOptionsModel.SearchTagBegin = core.StringPtr("testString")
				getSearchMatchesOptionsModel.SearchTagEnd = core.StringPtr("testString")
				getSearchMatchesOptionsModel.RelatedTagBegin = core.StringPtr("testString")
				getSearchMatchesOptionsModel.RelatedTagEnd = core.StringPtr("testString")
				getSearchMatchesOptionsModel.Fields = core.StringPtr("testString")
				getSearchMatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetSearchMatches(getSearchMatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSearchMatches(getSearchMatchesOptions *GetSearchMatchesOptions)`, func() {
		version := "testString"
		getSearchMatchesPath := "/v1/corpora/testString/documents/testString/search_matches"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSearchMatchesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for min_score query parameter

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["search_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["search_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["related_tag_begin"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["related_tag_end"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"externalId": "ExternalID", "documentId": "DocumentID", "parentDocumentId": "ParentDocumentID", "publicationName": "PublicationName", "publicationDate": "PublicationDate", "publicationURL": "PublicationURL", "authors": ["Authors"], "title": "Title", "medlineLicense": "MedlineLicense", "hrefPubMed": "HrefPubMed", "hrefPmc": "HrefPmc", "hrefDoi": "HrefDoi", "pdfUrl": "PdfURL", "referenceUrl": "ReferenceURL", "highlightedTitle": {}, "highlightedAbstract": {}, "highlightedBody": {}, "highlightedSections": {"mapKey": {}}, "passages": {"mapKey": {"mapKey": {"negated": false, "score": 5, "sentences": [{"documentSection": "DocumentSection", "text": {}, "begin": 5, "end": 3, "timestamp": 9}], "id": "ID"}}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}`)
				}))
			})
			It(`Invoke GetSearchMatches successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSearchMatches(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSearchMatchesOptions model
				getSearchMatchesOptionsModel := new(insightsformedicalliteratureservicev1.GetSearchMatchesOptions)
				getSearchMatchesOptionsModel.Corpus = core.StringPtr("testString")
				getSearchMatchesOptionsModel.DocumentID = core.StringPtr("testString")
				getSearchMatchesOptionsModel.MinScore = core.Float32Ptr(36.0)
				getSearchMatchesOptionsModel.Cuis = []string{"testString"}
				getSearchMatchesOptionsModel.Text = []string{"testString"}
				getSearchMatchesOptionsModel.Types = []string{"testString"}
				getSearchMatchesOptionsModel.Attributes = []string{"testString"}
				getSearchMatchesOptionsModel.Values = []string{"testString"}
				getSearchMatchesOptionsModel.NluRelations = []string{"testString"}
				getSearchMatchesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getSearchMatchesOptionsModel.SearchTagBegin = core.StringPtr("testString")
				getSearchMatchesOptionsModel.SearchTagEnd = core.StringPtr("testString")
				getSearchMatchesOptionsModel.RelatedTagBegin = core.StringPtr("testString")
				getSearchMatchesOptionsModel.RelatedTagEnd = core.StringPtr("testString")
				getSearchMatchesOptionsModel.Fields = core.StringPtr("testString")
				getSearchMatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSearchMatches(getSearchMatchesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSearchMatches with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSearchMatchesOptions model
				getSearchMatchesOptionsModel := new(insightsformedicalliteratureservicev1.GetSearchMatchesOptions)
				getSearchMatchesOptionsModel.Corpus = core.StringPtr("testString")
				getSearchMatchesOptionsModel.DocumentID = core.StringPtr("testString")
				getSearchMatchesOptionsModel.MinScore = core.Float32Ptr(36.0)
				getSearchMatchesOptionsModel.Cuis = []string{"testString"}
				getSearchMatchesOptionsModel.Text = []string{"testString"}
				getSearchMatchesOptionsModel.Types = []string{"testString"}
				getSearchMatchesOptionsModel.Attributes = []string{"testString"}
				getSearchMatchesOptionsModel.Values = []string{"testString"}
				getSearchMatchesOptionsModel.NluRelations = []string{"testString"}
				getSearchMatchesOptionsModel.Limit = core.Int64Ptr(int64(38))
				getSearchMatchesOptionsModel.SearchTagBegin = core.StringPtr("testString")
				getSearchMatchesOptionsModel.SearchTagEnd = core.StringPtr("testString")
				getSearchMatchesOptionsModel.RelatedTagBegin = core.StringPtr("testString")
				getSearchMatchesOptionsModel.RelatedTagEnd = core.StringPtr("testString")
				getSearchMatchesOptionsModel.Fields = core.StringPtr("testString")
				getSearchMatchesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetSearchMatches(getSearchMatchesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSearchMatchesOptions model with no property values
				getSearchMatchesOptionsModelNew := new(insightsformedicalliteratureservicev1.GetSearchMatchesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetSearchMatches(getSearchMatchesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "https://insightsformedicalliteratureservicev1/api",
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
	Describe(`GetCorporaConfig(getCorporaConfigOptions *GetCorporaConfigOptions) - Operation response error`, func() {
		version := "testString"
		getCorporaConfigPath := "/v1/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCorporaConfigPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCorporaConfig with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCorporaConfigOptions model
				getCorporaConfigOptionsModel := new(insightsformedicalliteratureservicev1.GetCorporaConfigOptions)
				getCorporaConfigOptionsModel.Verbose = core.BoolPtr(true)
				getCorporaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCorporaConfig(getCorporaConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCorporaConfig(getCorporaConfigOptions *GetCorporaConfigOptions)`, func() {
		version := "testString"
		getCorporaConfigPath := "/v1/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCorporaConfigPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"corpora": [{"corpusName": "CorpusName", "ontologies": ["Ontologies"], "descriptiveName": "DescriptiveName", "bvt": false, "elasticsearchIndex": "ElasticsearchIndex"}]}`)
				}))
			})
			It(`Invoke GetCorporaConfig successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCorporaConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCorporaConfigOptions model
				getCorporaConfigOptionsModel := new(insightsformedicalliteratureservicev1.GetCorporaConfigOptions)
				getCorporaConfigOptionsModel.Verbose = core.BoolPtr(true)
				getCorporaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCorporaConfig(getCorporaConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCorporaConfig with error: Operation request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCorporaConfigOptions model
				getCorporaConfigOptionsModel := new(insightsformedicalliteratureservicev1.GetCorporaConfigOptions)
				getCorporaConfigOptionsModel.Verbose = core.BoolPtr(true)
				getCorporaConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCorporaConfig(getCorporaConfigOptionsModel)
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
	Describe(`SetCorpusSchema(setCorpusSchemaOptions *SetCorpusSchemaOptions) - Operation response error`, func() {
		version := "testString"
		setCorpusSchemaPath := "/v1/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(setCorpusSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCorpusSchema with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SetCorpusSchemaOptions model
				setCorpusSchemaOptionsModel := new(insightsformedicalliteratureservicev1.SetCorpusSchemaOptions)
				setCorpusSchemaOptionsModel.EnrichmentTargets = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				setCorpusSchemaOptionsModel.MetadataFields = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				setCorpusSchemaOptionsModel.CorpusName = core.StringPtr("testString")
				setCorpusSchemaOptionsModel.References = make(map[string]interface{})
				setCorpusSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.SetCorpusSchema(setCorpusSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCorpusSchema(setCorpusSchemaOptions *SetCorpusSchemaOptions)`, func() {
		version := "testString"
		setCorpusSchemaPath := "/v1/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(setCorpusSchemaPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"corpora": [{"corpusName": "CorpusName", "ontologies": ["Ontologies"], "descriptiveName": "DescriptiveName", "bvt": false, "elasticsearchIndex": "ElasticsearchIndex"}]}`)
				}))
			})
			It(`Invoke SetCorpusSchema successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.SetCorpusSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetCorpusSchemaOptions model
				setCorpusSchemaOptionsModel := new(insightsformedicalliteratureservicev1.SetCorpusSchemaOptions)
				setCorpusSchemaOptionsModel.EnrichmentTargets = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				setCorpusSchemaOptionsModel.MetadataFields = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				setCorpusSchemaOptionsModel.CorpusName = core.StringPtr("testString")
				setCorpusSchemaOptionsModel.References = make(map[string]interface{})
				setCorpusSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.SetCorpusSchema(setCorpusSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke SetCorpusSchema with error: Operation request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SetCorpusSchemaOptions model
				setCorpusSchemaOptionsModel := new(insightsformedicalliteratureservicev1.SetCorpusSchemaOptions)
				setCorpusSchemaOptionsModel.EnrichmentTargets = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				setCorpusSchemaOptionsModel.MetadataFields = []interface{}{map[string]interface{}{"anyKey": "anyValue"}}
				setCorpusSchemaOptionsModel.CorpusName = core.StringPtr("testString")
				setCorpusSchemaOptionsModel.References = make(map[string]interface{})
				setCorpusSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.SetCorpusSchema(setCorpusSchemaOptionsModel)
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
	Describe(`DeleteCorpusSchema(deleteCorpusSchemaOptions *DeleteCorpusSchemaOptions) - Operation response error`, func() {
		version := "testString"
		deleteCorpusSchemaPath := "/v1/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteCorpusSchemaPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["instance"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteCorpusSchema with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteCorpusSchemaOptions model
				deleteCorpusSchemaOptionsModel := new(insightsformedicalliteratureservicev1.DeleteCorpusSchemaOptions)
				deleteCorpusSchemaOptionsModel.Instance = core.StringPtr("testString")
				deleteCorpusSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.DeleteCorpusSchema(deleteCorpusSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCorpusSchema(deleteCorpusSchemaOptions *DeleteCorpusSchemaOptions)`, func() {
		version := "testString"
		deleteCorpusSchemaPath := "/v1/corpora"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteCorpusSchemaPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["instance"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"corpora": [{"corpusName": "CorpusName", "ontologies": ["Ontologies"], "descriptiveName": "DescriptiveName", "bvt": false, "elasticsearchIndex": "ElasticsearchIndex"}]}`)
				}))
			})
			It(`Invoke DeleteCorpusSchema successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteCorpusSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteCorpusSchemaOptions model
				deleteCorpusSchemaOptionsModel := new(insightsformedicalliteratureservicev1.DeleteCorpusSchemaOptions)
				deleteCorpusSchemaOptionsModel.Instance = core.StringPtr("testString")
				deleteCorpusSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteCorpusSchema(deleteCorpusSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteCorpusSchema with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteCorpusSchemaOptions model
				deleteCorpusSchemaOptionsModel := new(insightsformedicalliteratureservicev1.DeleteCorpusSchemaOptions)
				deleteCorpusSchemaOptionsModel.Instance = core.StringPtr("testString")
				deleteCorpusSchemaOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteCorpusSchema(deleteCorpusSchemaOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteCorpusSchemaOptions model with no property values
				deleteCorpusSchemaOptionsModelNew := new(insightsformedicalliteratureservicev1.DeleteCorpusSchemaOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteCorpusSchema(deleteCorpusSchemaOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SetCorpusConfig(setCorpusConfigOptions *SetCorpusConfigOptions) - Operation response error`, func() {
		version := "testString"
		setCorpusConfigPath := "/v1/corpora/configure"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(setCorpusConfigPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SetCorpusConfig with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SetCorpusConfigOptions model
				setCorpusConfigOptionsModel := new(insightsformedicalliteratureservicev1.SetCorpusConfigOptions)
				setCorpusConfigOptionsModel.UserName = core.StringPtr("testString")
				setCorpusConfigOptionsModel.Password = core.StringPtr("testString")
				setCorpusConfigOptionsModel.CorpusURI = core.StringPtr("testString")
				setCorpusConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.SetCorpusConfig(setCorpusConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`SetCorpusConfig(setCorpusConfigOptions *SetCorpusConfigOptions)`, func() {
		version := "testString"
		setCorpusConfigPath := "/v1/corpora/configure"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(setCorpusConfigPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"corpora": [{"corpusName": "CorpusName", "ontologies": ["Ontologies"], "descriptiveName": "DescriptiveName", "bvt": false, "elasticsearchIndex": "ElasticsearchIndex"}]}`)
				}))
			})
			It(`Invoke SetCorpusConfig successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.SetCorpusConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SetCorpusConfigOptions model
				setCorpusConfigOptionsModel := new(insightsformedicalliteratureservicev1.SetCorpusConfigOptions)
				setCorpusConfigOptionsModel.UserName = core.StringPtr("testString")
				setCorpusConfigOptionsModel.Password = core.StringPtr("testString")
				setCorpusConfigOptionsModel.CorpusURI = core.StringPtr("testString")
				setCorpusConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.SetCorpusConfig(setCorpusConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke SetCorpusConfig with error: Operation request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SetCorpusConfigOptions model
				setCorpusConfigOptionsModel := new(insightsformedicalliteratureservicev1.SetCorpusConfigOptions)
				setCorpusConfigOptionsModel.UserName = core.StringPtr("testString")
				setCorpusConfigOptionsModel.Password = core.StringPtr("testString")
				setCorpusConfigOptionsModel.CorpusURI = core.StringPtr("testString")
				setCorpusConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.SetCorpusConfig(setCorpusConfigOptionsModel)
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

	Describe(`MonitorCorpus(monitorCorpusOptions *MonitorCorpusOptions)`, func() {
		version := "testString"
		monitorCorpusPath := "/v1/corpora/monitor"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(monitorCorpusPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["apikey"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke MonitorCorpus successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.MonitorCorpus(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the MonitorCorpusOptions model
				monitorCorpusOptionsModel := new(insightsformedicalliteratureservicev1.MonitorCorpusOptions)
				monitorCorpusOptionsModel.Apikey = core.StringPtr("testString")
				monitorCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.MonitorCorpus(monitorCorpusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke MonitorCorpus with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the MonitorCorpusOptions model
				monitorCorpusOptionsModel := new(insightsformedicalliteratureservicev1.MonitorCorpusOptions)
				monitorCorpusOptionsModel.Apikey = core.StringPtr("testString")
				monitorCorpusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.MonitorCorpus(monitorCorpusOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the MonitorCorpusOptions model with no property values
				monitorCorpusOptionsModelNew := new(insightsformedicalliteratureservicev1.MonitorCorpusOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.MonitorCorpus(monitorCorpusOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EnableCorpusSearchTracking(enableCorpusSearchTrackingOptions *EnableCorpusSearchTrackingOptions)`, func() {
		version := "testString"
		enableCorpusSearchTrackingPath := "/v1/corpora/tracking"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(enableCorpusSearchTrackingPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for enable_tracking query parameter

					res.WriteHeader(200)
				}))
			})
			It(`Invoke EnableCorpusSearchTracking successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.EnableCorpusSearchTracking(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the EnableCorpusSearchTrackingOptions model
				enableCorpusSearchTrackingOptionsModel := new(insightsformedicalliteratureservicev1.EnableCorpusSearchTrackingOptions)
				enableCorpusSearchTrackingOptionsModel.EnableTracking = core.BoolPtr(true)
				enableCorpusSearchTrackingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.EnableCorpusSearchTracking(enableCorpusSearchTrackingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke EnableCorpusSearchTracking with error: Operation request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the EnableCorpusSearchTrackingOptions model
				enableCorpusSearchTrackingOptionsModel := new(insightsformedicalliteratureservicev1.EnableCorpusSearchTrackingOptions)
				enableCorpusSearchTrackingOptionsModel.EnableTracking = core.BoolPtr(true)
				enableCorpusSearchTrackingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.EnableCorpusSearchTracking(enableCorpusSearchTrackingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCorpusConfig(getCorpusConfigOptions *GetCorpusConfigOptions) - Operation response error`, func() {
		version := "testString"
		getCorpusConfigPath := "/v1/corpora/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCorpusConfigPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCorpusConfig with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCorpusConfigOptions model
				getCorpusConfigOptionsModel := new(insightsformedicalliteratureservicev1.GetCorpusConfigOptions)
				getCorpusConfigOptionsModel.Corpus = core.StringPtr("testString")
				getCorpusConfigOptionsModel.Verbose = core.BoolPtr(true)
				getCorpusConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCorpusConfig(getCorpusConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCorpusConfig(getCorpusConfigOptions *GetCorpusConfigOptions)`, func() {
		version := "testString"
		getCorpusConfigPath := "/v1/corpora/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCorpusConfigPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"corpora": [{"corpusName": "CorpusName", "ontologies": ["Ontologies"], "descriptiveName": "DescriptiveName", "bvt": false, "elasticsearchIndex": "ElasticsearchIndex"}]}`)
				}))
			})
			It(`Invoke GetCorpusConfig successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCorpusConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCorpusConfigOptions model
				getCorpusConfigOptionsModel := new(insightsformedicalliteratureservicev1.GetCorpusConfigOptions)
				getCorpusConfigOptionsModel.Corpus = core.StringPtr("testString")
				getCorpusConfigOptionsModel.Verbose = core.BoolPtr(true)
				getCorpusConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCorpusConfig(getCorpusConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCorpusConfig with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCorpusConfigOptions model
				getCorpusConfigOptionsModel := new(insightsformedicalliteratureservicev1.GetCorpusConfigOptions)
				getCorpusConfigOptionsModel.Corpus = core.StringPtr("testString")
				getCorpusConfigOptionsModel.Verbose = core.BoolPtr(true)
				getCorpusConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCorpusConfig(getCorpusConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCorpusConfigOptions model with no property values
				getCorpusConfigOptionsModelNew := new(insightsformedicalliteratureservicev1.GetCorpusConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetCorpusConfig(getCorpusConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "https://insightsformedicalliteratureservicev1/api",
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHealthCheckStatusOptions model
				getHealthCheckStatusOptionsModel := new(insightsformedicalliteratureservicev1.GetHealthCheckStatusOptions)
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
					fmt.Fprintf(res, `{"serviceState": "OK", "stateDetails": "StateDetails"}`)
				}))
			})
			It(`Invoke GetHealthCheckStatus successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				getHealthCheckStatusOptionsModel := new(insightsformedicalliteratureservicev1.GetHealthCheckStatusOptions)
				getHealthCheckStatusOptionsModel.Accept = core.StringPtr("application/json")
				getHealthCheckStatusOptionsModel.Format = core.StringPtr("json")
				getHealthCheckStatusOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetHealthCheckStatus(getHealthCheckStatusOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetHealthCheckStatus with error: Operation request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHealthCheckStatusOptions model
				getHealthCheckStatusOptionsModel := new(insightsformedicalliteratureservicev1.GetHealthCheckStatusOptions)
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
	Describe(`Service constructor tests`, func() {
		version := "testString"
		It(`Instantiate service client`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "https://insightsformedicalliteratureservicev1/api",
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
	Describe(`Search(searchOptions *SearchOptions) - Operation response error`, func() {
		version := "testString"
		searchPath := "/v1/corpora/testString/search"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Search with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(insightsformedicalliteratureservicev1.SearchOptions)
				searchOptionsModel.Corpus = core.StringPtr("testString")
				searchOptionsModel.Body = core.StringPtr("testString")
				searchOptionsModel.Verbose = core.BoolPtr(true)
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.Search(searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Search(searchOptions *SearchOptions)`, func() {
		version := "testString"
		searchPath := "/v1/corpora/testString/search"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(searchPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "attributeValues": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "relations": [{"relationId": "RelationID", "relation": "Relation", "source": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}, "target": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}}], "mesh": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "text": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}]}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "attributeValues": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "relations": [{"relationId": "RelationID", "relation": "Relation", "source": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}, "target": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}}], "mesh": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "text": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}]}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "attributeValues": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "relations": [{"relationId": "RelationID", "relation": "Relation", "source": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}, "target": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}}], "mesh": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "text": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}]}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "attributeValues": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "relations": [{"relationId": "RelationID", "relation": "Relation", "source": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}, "target": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}}], "mesh": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "text": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}]}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "attributeValues": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "relations": [{"relationId": "RelationID", "relation": "Relation", "source": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}, "target": {"section": "Section", "begin": 5, "end": 3, "coveredText": "CoveredText", "source": "Source", "type": "Type"}}], "mesh": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "text": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}]}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "attributeValues": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "relations": [{"relationId": "RelationID", "relation": "Relation"}], "mesh": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}], "text": [{"uniqueId": 8, "stickyIds": [9], "section": "Section", "type": "Type", "begin": 5, "end": 3, "coveredText": "CoveredText", "cui": "Cui", "preferredName": "PreferredName", "source": "Source", "negated": false, "hypothetical": true, "timestamp": 9, "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "features": {"mapKey": "Inner"}, "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "hits": 4}]}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "text": {}, "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text", "data": {"concepts": [], "attributeValues": [], "relations": [], "mesh": [], "text": []}}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [{"documentSection": "DocumentSection", "timestamp": 9, "preferredName": "PreferredName"}], "searchTermToPassages": {"mapKey": [{"documentSection": "DocumentSection", "timestamp": 9, "preferredName": "PreferredName"}]}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "parents": {"count": 5, "hits": 4}, "children": {"count": 5, "hits": 4}, "siblings": {"count": 5, "hits": 4}, "related": {"count": 5, "hits": 4}, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": [{"messageType": "expanded_request", "url": "URL", "request": {"anyKey": "anyValue"}, "headers": ["Headers"], "status": 6, "response": {"anyKey": "anyValue"}}]}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": [{"boolOperand": "BoolOperand"}]}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": [{"text": "Text"}]}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "links": {"hrefSearchMatches": "HrefSearchMatches", "hrefCategories": "HrefCategories"}, "passages": {"allPassages": [], "searchTermToPassages": {"mapKey": []}}, "annotations": {"mapKey": {"uniqueId": 8, "stickyIds": [9], "ontology": "Ontology", "section": "Section", "preferredName": "PreferredName", "cui": "Cui", "attributeId": "AttributeID", "qualifiers": ["Qualifiers"], "type": "Type", "negated": false, "hypothetical": true, "unit": "Unit", "minValue": "MinValue", "maxValue": "MaxValue", "operator": "Operator", "nluSourceType": "NluSourceType", "nluRelation": "NluRelation", "nluTargetType": "NluTargetType", "nluEntityIndex": "NluEntityIndex", "nluMentionIndex": "NluMentionIndex", "nluRelationId": "NluRelationID", "nluSide": "NluSide", "begin": 5, "end": 3, "score": 5, "timestamp": 9, "features": {"mapKey": "Inner"}, "hits": 4}}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "types": ["Types"], "attributes": [{"attributeId": "AttributeID", "displayName": "DisplayName", "count": 5}], "values": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "ranges": {"mapKey": {"operator": "Operator", "min": "Min", "max": "Max", "count": 5}}, "typeahead": [{"ontology": "Ontology", "cui": "Cui", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "count": 5, "hitCount": 8, "score": 5, "documentIds": ["DocumentIds"], "dataType": "DataType", "unit": "Unit", "operator": "Operator", "minValue": "MinValue", "maxValue": "MaxValue", "vocab": "Vocab", "properties": ["Properties"]}], "aggregations": {"mapKey": [{"name": "Name", "documentCount": 13}]}, "dateHistograms": {"mapKey": [{"date": "Date", "hits": 4}]}, "qualifiers": [{"id": "ID", "name": "Name"}], "backend": {"messages": []}, "expandedQuery": {"anyKey": "anyValue"}, "parsedBoolExpression": {"boolExpression": "BoolExpression", "boolOperands": []}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": {"unstructured": []}}, "metadata": {"corpus": "Corpus", "corpusDescription": "CorpusDescription", "fields": {"mapKey": ["Inner"]}}, "documents": [{"documentId": "DocumentID", "title": "Title", "metadata": {"mapKey": ["Inner"]}, "sectionName": "SectionName", "sectionId": "SectionID", "corpus": "Corpus", "annotations": {"mapKey": null}}], "subQueries": [{"href": "Href", "pageNumber": 10, "get_limit": 8, "totalDocumentCount": 18, "concepts": [], "types": ["Types"], "attributes": [], "values": [], "ranges": {"mapKey": null}, "typeahead": [], "aggregations": {"mapKey": []}, "dateHistograms": {"mapKey": []}, "qualifiers": [], "expandedQuery": {"anyKey": "anyValue"}, "conceptsExist": {"mapKey": 5}, "cursorId": "CursorID", "vocabs": ["Vocabs"], "annotations": {"mapKey": null}, "documents": [], "subQueries": []}]}]}]}]}]}]}]}]}]}`)
				}))
			})
			It(`Invoke Search successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.Search(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(insightsformedicalliteratureservicev1.SearchOptions)
				searchOptionsModel.Corpus = core.StringPtr("testString")
				searchOptionsModel.Body = core.StringPtr("testString")
				searchOptionsModel.Verbose = core.BoolPtr(true)
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.Search(searchOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke Search with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the SearchOptions model
				searchOptionsModel := new(insightsformedicalliteratureservicev1.SearchOptions)
				searchOptionsModel.Corpus = core.StringPtr("testString")
				searchOptionsModel.Body = core.StringPtr("testString")
				searchOptionsModel.Verbose = core.BoolPtr(true)
				searchOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.Search(searchOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SearchOptions model with no property values
				searchOptionsModelNew := new(insightsformedicalliteratureservicev1.SearchOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.Search(searchOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetFields(getFieldsOptions *GetFieldsOptions) - Operation response error`, func() {
		version := "testString"
		getFieldsPath := "/v1/corpora/testString/search/metadata"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFieldsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetFields with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFieldsOptions model
				getFieldsOptionsModel := new(insightsformedicalliteratureservicev1.GetFieldsOptions)
				getFieldsOptionsModel.Corpus = core.StringPtr("testString")
				getFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetFields(getFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetFields(getFieldsOptions *GetFieldsOptions)`, func() {
		version := "testString"
		getFieldsPath := "/v1/corpora/testString/search/metadata"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getFieldsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"fields": {"mapKey": {"supports": ["Supports"]}}, "sectionFieldNames": ["SectionFieldNames"], "attrSectionFieldNames": ["AttrSectionFieldNames"], "qualifierSectionFieldNames": ["QualifierSectionFieldNames"], "meshSectionFieldNames": ["MeshSectionFieldNames"], "fieldIndexMap": {"mapKey": "Inner"}}`)
				}))
			})
			It(`Invoke GetFields successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetFields(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetFieldsOptions model
				getFieldsOptionsModel := new(insightsformedicalliteratureservicev1.GetFieldsOptions)
				getFieldsOptionsModel.Corpus = core.StringPtr("testString")
				getFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetFields(getFieldsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetFields with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetFieldsOptions model
				getFieldsOptionsModel := new(insightsformedicalliteratureservicev1.GetFieldsOptions)
				getFieldsOptionsModel.Corpus = core.StringPtr("testString")
				getFieldsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetFields(getFieldsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetFieldsOptions model with no property values
				getFieldsOptionsModelNew := new(insightsformedicalliteratureservicev1.GetFieldsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetFields(getFieldsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Typeahead(typeaheadOptions *TypeaheadOptions) - Operation response error`, func() {
		version := "testString"
		typeaheadPath := "/v1/corpora/testString/search/typeahead"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(typeaheadPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["category"]).To(Equal([]string{"disorders"}))

					// TODO: Add check for verbose query parameter

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["max_hit_count"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// TODO: Add check for no_duplicates query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Typeahead with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TypeaheadOptions model
				typeaheadOptionsModel := new(insightsformedicalliteratureservicev1.TypeaheadOptions)
				typeaheadOptionsModel.Corpus = core.StringPtr("testString")
				typeaheadOptionsModel.Query = core.StringPtr("testString")
				typeaheadOptionsModel.Ontologies = []string{"concepts"}
				typeaheadOptionsModel.Types = []string{"testString"}
				typeaheadOptionsModel.Category = core.StringPtr("disorders")
				typeaheadOptionsModel.Verbose = core.BoolPtr(true)
				typeaheadOptionsModel.Limit = core.Int64Ptr(int64(38))
				typeaheadOptionsModel.MaxHitCount = core.Int64Ptr(int64(38))
				typeaheadOptionsModel.NoDuplicates = core.BoolPtr(true)
				typeaheadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.Typeahead(typeaheadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`Typeahead(typeaheadOptions *TypeaheadOptions)`, func() {
		version := "testString"
		typeaheadPath := "/v1/corpora/testString/search/typeahead"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(typeaheadPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["query"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["category"]).To(Equal([]string{"disorders"}))

					// TODO: Add check for verbose query parameter

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["max_hit_count"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					// TODO: Add check for no_duplicates query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"concepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"]}]}`)
				}))
			})
			It(`Invoke Typeahead successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.Typeahead(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TypeaheadOptions model
				typeaheadOptionsModel := new(insightsformedicalliteratureservicev1.TypeaheadOptions)
				typeaheadOptionsModel.Corpus = core.StringPtr("testString")
				typeaheadOptionsModel.Query = core.StringPtr("testString")
				typeaheadOptionsModel.Ontologies = []string{"concepts"}
				typeaheadOptionsModel.Types = []string{"testString"}
				typeaheadOptionsModel.Category = core.StringPtr("disorders")
				typeaheadOptionsModel.Verbose = core.BoolPtr(true)
				typeaheadOptionsModel.Limit = core.Int64Ptr(int64(38))
				typeaheadOptionsModel.MaxHitCount = core.Int64Ptr(int64(38))
				typeaheadOptionsModel.NoDuplicates = core.BoolPtr(true)
				typeaheadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.Typeahead(typeaheadOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke Typeahead with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the TypeaheadOptions model
				typeaheadOptionsModel := new(insightsformedicalliteratureservicev1.TypeaheadOptions)
				typeaheadOptionsModel.Corpus = core.StringPtr("testString")
				typeaheadOptionsModel.Query = core.StringPtr("testString")
				typeaheadOptionsModel.Ontologies = []string{"concepts"}
				typeaheadOptionsModel.Types = []string{"testString"}
				typeaheadOptionsModel.Category = core.StringPtr("disorders")
				typeaheadOptionsModel.Verbose = core.BoolPtr(true)
				typeaheadOptionsModel.Limit = core.Int64Ptr(int64(38))
				typeaheadOptionsModel.MaxHitCount = core.Int64Ptr(int64(38))
				typeaheadOptionsModel.NoDuplicates = core.BoolPtr(true)
				typeaheadOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.Typeahead(typeaheadOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TypeaheadOptions model with no property values
				typeaheadOptionsModelNew := new(insightsformedicalliteratureservicev1.TypeaheadOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.Typeahead(typeaheadOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "{BAD_URL_STRING",
				Version: core.StringPtr(version),
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:     "https://insightsformedicalliteratureservicev1/api",
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
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		version := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					Version: core.StringPtr(version),
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_URL":       "https://insightsformedicalliteratureservicev1/api",
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
				"INSIGHTS_FOR_MEDICAL_LITERATURE_SERVICE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
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
	Describe(`GetConcepts(getConceptsOptions *GetConceptsOptions) - Operation response error`, func() {
		version := "testString"
		getConceptsPath := "/v1/corpora/testString/concepts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getConceptsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					Expect(req.URL.Query()["_sort"]).To(Equal([]string{"hitCount"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConcepts with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetConceptsOptions model
				getConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetConceptsOptions)
				getConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getConceptsOptionsModel.Cuis = []string{"testString"}
				getConceptsOptionsModel.PreferredNames = []string{"testString"}
				getConceptsOptionsModel.SurfaceForms = []string{"testString"}
				getConceptsOptionsModel.Attributes = []string{"testString"}
				getConceptsOptionsModel.Verbose = core.BoolPtr(true)
				getConceptsOptionsModel.Sort = core.StringPtr("hitCount")
				getConceptsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetConcepts(getConceptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetConcepts(getConceptsOptions *GetConceptsOptions)`, func() {
		version := "testString"
		getConceptsPath := "/v1/corpora/testString/concepts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getConceptsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					// TODO: Add check for verbose query parameter

					Expect(req.URL.Query()["_sort"]).To(Equal([]string{"hitCount"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"concepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"]}]}`)
				}))
			})
			It(`Invoke GetConcepts successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetConcepts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConceptsOptions model
				getConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetConceptsOptions)
				getConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getConceptsOptionsModel.Cuis = []string{"testString"}
				getConceptsOptionsModel.PreferredNames = []string{"testString"}
				getConceptsOptionsModel.SurfaceForms = []string{"testString"}
				getConceptsOptionsModel.Attributes = []string{"testString"}
				getConceptsOptionsModel.Verbose = core.BoolPtr(true)
				getConceptsOptionsModel.Sort = core.StringPtr("hitCount")
				getConceptsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetConcepts(getConceptsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetConcepts with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetConceptsOptions model
				getConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetConceptsOptions)
				getConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getConceptsOptionsModel.Cuis = []string{"testString"}
				getConceptsOptionsModel.PreferredNames = []string{"testString"}
				getConceptsOptionsModel.SurfaceForms = []string{"testString"}
				getConceptsOptionsModel.Attributes = []string{"testString"}
				getConceptsOptionsModel.Verbose = core.BoolPtr(true)
				getConceptsOptionsModel.Sort = core.StringPtr("hitCount")
				getConceptsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetConcepts(getConceptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConceptsOptions model with no property values
				getConceptsOptionsModelNew := new(insightsformedicalliteratureservicev1.GetConceptsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetConcepts(getConceptsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddArtifact(addArtifactOptions *AddArtifactOptions)`, func() {
		version := "testString"
		addArtifactPath := "/v1/corpora/testString/concepts/definitions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(addArtifactPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddArtifact successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.AddArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the PossbileValues model
				possbileValuesModel := new(insightsformedicalliteratureservicev1.PossbileValues)
				possbileValuesModel.DisplayValue = core.StringPtr("testString")
				possbileValuesModel.Value = core.StringPtr("testString")

				// Construct an instance of the AttributeEntry model
				attributeEntryModel := new(insightsformedicalliteratureservicev1.AttributeEntry)
				attributeEntryModel.AttrName = core.StringPtr("testString")
				attributeEntryModel.DataType = core.StringPtr("testString")
				attributeEntryModel.DefaultValue = core.StringPtr("testString")
				attributeEntryModel.Description = core.StringPtr("testString")
				attributeEntryModel.DisplayName = core.StringPtr("testString")
				attributeEntryModel.DocID = core.StringPtr("testString")
				attributeEntryModel.FieldValues = []string{"testString"}
				attributeEntryModel.MaximumValue = core.StringPtr("testString")
				attributeEntryModel.MinimumValue = core.StringPtr("testString")
				attributeEntryModel.MultiValue = core.BoolPtr(true)
				attributeEntryModel.Units = core.StringPtr("testString")
				attributeEntryModel.ValueType = core.StringPtr("testString")
				attributeEntryModel.PossibleValues = []insightsformedicalliteratureservicev1.PossbileValues{*possbileValuesModel}

				// Construct an instance of the DictonaryEntry model
				dictonaryEntryModel := new(insightsformedicalliteratureservicev1.DictonaryEntry)
				dictonaryEntryModel.Children = []string{"testString"}
				dictonaryEntryModel.Cui = core.StringPtr("testString")
				dictonaryEntryModel.Definition = []string{"testString"}
				dictonaryEntryModel.Parents = []string{"testString"}
				dictonaryEntryModel.PreferredName = core.StringPtr("testString")
				dictonaryEntryModel.Semtypes = []string{"testString"}
				dictonaryEntryModel.Siblings = []string{"testString"}
				dictonaryEntryModel.SurfaceForms = []string{"testString"}
				dictonaryEntryModel.Variants = []string{"testString"}
				dictonaryEntryModel.Vocab = core.StringPtr("testString")
				dictonaryEntryModel.Related = []string{"testString"}
				dictonaryEntryModel.Source = core.StringPtr("testString")
				dictonaryEntryModel.SourceVersion = core.StringPtr("testString")

				// Construct an instance of the AddArtifactOptions model
				addArtifactOptionsModel := new(insightsformedicalliteratureservicev1.AddArtifactOptions)
				addArtifactOptionsModel.Corpus = core.StringPtr("testString")
				addArtifactOptionsModel.DictionaryEntry = dictonaryEntryModel
				addArtifactOptionsModel.AttributeEntry = attributeEntryModel
				addArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.AddArtifact(addArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddArtifact with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the PossbileValues model
				possbileValuesModel := new(insightsformedicalliteratureservicev1.PossbileValues)
				possbileValuesModel.DisplayValue = core.StringPtr("testString")
				possbileValuesModel.Value = core.StringPtr("testString")

				// Construct an instance of the AttributeEntry model
				attributeEntryModel := new(insightsformedicalliteratureservicev1.AttributeEntry)
				attributeEntryModel.AttrName = core.StringPtr("testString")
				attributeEntryModel.DataType = core.StringPtr("testString")
				attributeEntryModel.DefaultValue = core.StringPtr("testString")
				attributeEntryModel.Description = core.StringPtr("testString")
				attributeEntryModel.DisplayName = core.StringPtr("testString")
				attributeEntryModel.DocID = core.StringPtr("testString")
				attributeEntryModel.FieldValues = []string{"testString"}
				attributeEntryModel.MaximumValue = core.StringPtr("testString")
				attributeEntryModel.MinimumValue = core.StringPtr("testString")
				attributeEntryModel.MultiValue = core.BoolPtr(true)
				attributeEntryModel.Units = core.StringPtr("testString")
				attributeEntryModel.ValueType = core.StringPtr("testString")
				attributeEntryModel.PossibleValues = []insightsformedicalliteratureservicev1.PossbileValues{*possbileValuesModel}

				// Construct an instance of the DictonaryEntry model
				dictonaryEntryModel := new(insightsformedicalliteratureservicev1.DictonaryEntry)
				dictonaryEntryModel.Children = []string{"testString"}
				dictonaryEntryModel.Cui = core.StringPtr("testString")
				dictonaryEntryModel.Definition = []string{"testString"}
				dictonaryEntryModel.Parents = []string{"testString"}
				dictonaryEntryModel.PreferredName = core.StringPtr("testString")
				dictonaryEntryModel.Semtypes = []string{"testString"}
				dictonaryEntryModel.Siblings = []string{"testString"}
				dictonaryEntryModel.SurfaceForms = []string{"testString"}
				dictonaryEntryModel.Variants = []string{"testString"}
				dictonaryEntryModel.Vocab = core.StringPtr("testString")
				dictonaryEntryModel.Related = []string{"testString"}
				dictonaryEntryModel.Source = core.StringPtr("testString")
				dictonaryEntryModel.SourceVersion = core.StringPtr("testString")

				// Construct an instance of the AddArtifactOptions model
				addArtifactOptionsModel := new(insightsformedicalliteratureservicev1.AddArtifactOptions)
				addArtifactOptionsModel.Corpus = core.StringPtr("testString")
				addArtifactOptionsModel.DictionaryEntry = dictonaryEntryModel
				addArtifactOptionsModel.AttributeEntry = attributeEntryModel
				addArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.AddArtifact(addArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddArtifactOptions model with no property values
				addArtifactOptionsModelNew := new(insightsformedicalliteratureservicev1.AddArtifactOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.AddArtifact(addArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCuiInfo(getCuiInfoOptions *GetCuiInfoOptions) - Operation response error`, func() {
		version := "testString"
		getCuiInfoPath := "/v1/corpora/testString/concepts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCuiInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					// TODO: Add check for tree_layout query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCuiInfo with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCuiInfoOptions model
				getCuiInfoOptionsModel := new(insightsformedicalliteratureservicev1.GetCuiInfoOptions)
				getCuiInfoOptionsModel.Corpus = core.StringPtr("testString")
				getCuiInfoOptionsModel.NameOrID = core.StringPtr("testString")
				getCuiInfoOptionsModel.Ontology = core.StringPtr("testString")
				getCuiInfoOptionsModel.Fields = core.StringPtr("testString")
				getCuiInfoOptionsModel.TreeLayout = core.BoolPtr(true)
				getCuiInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCuiInfo(getCuiInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCuiInfo(getCuiInfoOptions *GetCuiInfoOptions)`, func() {
		version := "testString"
		getCuiInfoPath := "/v1/corpora/testString/concepts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCuiInfoPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_fields"]).To(Equal([]string{"testString"}))

					// TODO: Add check for tree_layout query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "semanticTypes": ["SemanticTypes"], "surfaceForms": ["SurfaceForms"], "definition": "Definition", "hasParents": true, "hasChildren": false, "hasSiblings": false}`)
				}))
			})
			It(`Invoke GetCuiInfo successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCuiInfo(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCuiInfoOptions model
				getCuiInfoOptionsModel := new(insightsformedicalliteratureservicev1.GetCuiInfoOptions)
				getCuiInfoOptionsModel.Corpus = core.StringPtr("testString")
				getCuiInfoOptionsModel.NameOrID = core.StringPtr("testString")
				getCuiInfoOptionsModel.Ontology = core.StringPtr("testString")
				getCuiInfoOptionsModel.Fields = core.StringPtr("testString")
				getCuiInfoOptionsModel.TreeLayout = core.BoolPtr(true)
				getCuiInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCuiInfo(getCuiInfoOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCuiInfo with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCuiInfoOptions model
				getCuiInfoOptionsModel := new(insightsformedicalliteratureservicev1.GetCuiInfoOptions)
				getCuiInfoOptionsModel.Corpus = core.StringPtr("testString")
				getCuiInfoOptionsModel.NameOrID = core.StringPtr("testString")
				getCuiInfoOptionsModel.Ontology = core.StringPtr("testString")
				getCuiInfoOptionsModel.Fields = core.StringPtr("testString")
				getCuiInfoOptionsModel.TreeLayout = core.BoolPtr(true)
				getCuiInfoOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCuiInfo(getCuiInfoOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCuiInfoOptions model with no property values
				getCuiInfoOptionsModelNew := new(insightsformedicalliteratureservicev1.GetCuiInfoOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetCuiInfo(getCuiInfoOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetHitCount(getHitCountOptions *GetHitCountOptions) - Operation response error`, func() {
		version := "testString"
		getHitCountPath := "/v1/corpora/testString/concepts/testString/hit_count"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getHitCountPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetHitCount with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHitCountOptions model
				getHitCountOptionsModel := new(insightsformedicalliteratureservicev1.GetHitCountOptions)
				getHitCountOptionsModel.Corpus = core.StringPtr("testString")
				getHitCountOptionsModel.NameOrID = core.StringPtr("testString")
				getHitCountOptionsModel.Ontology = core.StringPtr("testString")
				getHitCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetHitCount(getHitCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetHitCount(getHitCountOptions *GetHitCountOptions)`, func() {
		version := "testString"
		getHitCountPath := "/v1/corpora/testString/concepts/testString/hit_count"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getHitCountPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"hitCount": 8}`)
				}))
			})
			It(`Invoke GetHitCount successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetHitCount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetHitCountOptions model
				getHitCountOptionsModel := new(insightsformedicalliteratureservicev1.GetHitCountOptions)
				getHitCountOptionsModel.Corpus = core.StringPtr("testString")
				getHitCountOptionsModel.NameOrID = core.StringPtr("testString")
				getHitCountOptionsModel.Ontology = core.StringPtr("testString")
				getHitCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetHitCount(getHitCountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetHitCount with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetHitCountOptions model
				getHitCountOptionsModel := new(insightsformedicalliteratureservicev1.GetHitCountOptions)
				getHitCountOptionsModel.Corpus = core.StringPtr("testString")
				getHitCountOptionsModel.NameOrID = core.StringPtr("testString")
				getHitCountOptionsModel.Ontology = core.StringPtr("testString")
				getHitCountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetHitCount(getHitCountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetHitCountOptions model with no property values
				getHitCountOptionsModelNew := new(insightsformedicalliteratureservicev1.GetHitCountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetHitCount(getHitCountOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRelatedConcepts(getRelatedConceptsOptions *GetRelatedConceptsOptions) - Operation response error`, func() {
		version := "testString"
		getRelatedConceptsPath := "/v1/corpora/testString/concepts/testString/related_concepts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRelatedConceptsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["relationship"]).To(Equal([]string{"children"}))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					// TODO: Add check for recursive query parameter

					// TODO: Add check for tree_layout query parameter

					Expect(req.URL.Query()["max_depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRelatedConcepts with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRelatedConceptsOptions model
				getRelatedConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetRelatedConceptsOptions)
				getRelatedConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.NameOrID = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.Relationship = core.StringPtr("children")
				getRelatedConceptsOptionsModel.Ontology = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.RelationshipAttributes = []string{"testString"}
				getRelatedConceptsOptionsModel.Sources = []string{"testString"}
				getRelatedConceptsOptionsModel.Recursive = core.BoolPtr(true)
				getRelatedConceptsOptionsModel.TreeLayout = core.BoolPtr(true)
				getRelatedConceptsOptionsModel.MaxDepth = core.Int64Ptr(int64(38))
				getRelatedConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetRelatedConcepts(getRelatedConceptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetRelatedConcepts(getRelatedConceptsOptions *GetRelatedConceptsOptions)`, func() {
		version := "testString"
		getRelatedConceptsPath := "/v1/corpora/testString/concepts/testString/related_concepts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRelatedConceptsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["relationship"]).To(Equal([]string{"children"}))

					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					// TODO: Add check for recursive query parameter

					// TODO: Add check for tree_layout query parameter

					Expect(req.URL.Query()["max_depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"concepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"], "nextConcepts": []}]}]}]}]}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetRelatedConcepts successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetRelatedConcepts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRelatedConceptsOptions model
				getRelatedConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetRelatedConceptsOptions)
				getRelatedConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.NameOrID = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.Relationship = core.StringPtr("children")
				getRelatedConceptsOptionsModel.Ontology = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.RelationshipAttributes = []string{"testString"}
				getRelatedConceptsOptionsModel.Sources = []string{"testString"}
				getRelatedConceptsOptionsModel.Recursive = core.BoolPtr(true)
				getRelatedConceptsOptionsModel.TreeLayout = core.BoolPtr(true)
				getRelatedConceptsOptionsModel.MaxDepth = core.Int64Ptr(int64(38))
				getRelatedConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetRelatedConcepts(getRelatedConceptsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetRelatedConcepts with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRelatedConceptsOptions model
				getRelatedConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetRelatedConceptsOptions)
				getRelatedConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.NameOrID = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.Relationship = core.StringPtr("children")
				getRelatedConceptsOptionsModel.Ontology = core.StringPtr("testString")
				getRelatedConceptsOptionsModel.RelationshipAttributes = []string{"testString"}
				getRelatedConceptsOptionsModel.Sources = []string{"testString"}
				getRelatedConceptsOptionsModel.Recursive = core.BoolPtr(true)
				getRelatedConceptsOptionsModel.TreeLayout = core.BoolPtr(true)
				getRelatedConceptsOptionsModel.MaxDepth = core.Int64Ptr(int64(38))
				getRelatedConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetRelatedConcepts(getRelatedConceptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRelatedConceptsOptions model with no property values
				getRelatedConceptsOptionsModelNew := new(insightsformedicalliteratureservicev1.GetRelatedConceptsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetRelatedConcepts(getRelatedConceptsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSimilarConcepts(getSimilarConceptsOptions *GetSimilarConceptsOptions) - Operation response error`, func() {
		version := "testString"
		getSimilarConceptsPath := "/v1/corpora/testString/concepts/testString/similar_concepts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSimilarConceptsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSimilarConcepts with error: Operation response processing error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSimilarConceptsOptions model
				getSimilarConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetSimilarConceptsOptions)
				getSimilarConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.NameOrID = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.ReturnOntologies = []string{"testString"}
				getSimilarConceptsOptionsModel.Ontology = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getSimilarConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetSimilarConcepts(getSimilarConceptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetSimilarConcepts(getSimilarConceptsOptions *GetSimilarConceptsOptions)`, func() {
		version := "testString"
		getSimilarConceptsPath := "/v1/corpora/testString/concepts/testString/similar_concepts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getSimilarConceptsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["version"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ontology"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"concepts": [{"cui": "Cui", "ontology": "Ontology", "preferredName": "PreferredName", "alternativeName": "AlternativeName", "semanticType": "SemanticType", "rank": 4, "hitCount": 8, "score": 5, "surfaceForms": ["SurfaceForms"]}]}`)
				}))
			})
			It(`Invoke GetSimilarConcepts successfully`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSimilarConcepts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSimilarConceptsOptions model
				getSimilarConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetSimilarConceptsOptions)
				getSimilarConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.NameOrID = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.ReturnOntologies = []string{"testString"}
				getSimilarConceptsOptionsModel.Ontology = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getSimilarConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSimilarConcepts(getSimilarConceptsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetSimilarConcepts with error: Operation validation and request error`, func() {
				testService, testServiceErr := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
					Version:       core.StringPtr(version),
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetSimilarConceptsOptions model
				getSimilarConceptsOptionsModel := new(insightsformedicalliteratureservicev1.GetSimilarConceptsOptions)
				getSimilarConceptsOptionsModel.Corpus = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.NameOrID = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.ReturnOntologies = []string{"testString"}
				getSimilarConceptsOptionsModel.Ontology = core.StringPtr("testString")
				getSimilarConceptsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getSimilarConceptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetSimilarConcepts(getSimilarConceptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSimilarConceptsOptions model with no property values
				getSimilarConceptsOptionsModelNew := new(insightsformedicalliteratureservicev1.GetSimilarConceptsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetSimilarConcepts(getSimilarConceptsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
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
			testService, _ := insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:           "http://insightsformedicalliteratureservicev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
				Version:       core.StringPtr(version),
			})
			It(`Invoke NewAddArtifactOptions successfully`, func() {
				// Construct an instance of the PossbileValues model
				possbileValuesModel := new(insightsformedicalliteratureservicev1.PossbileValues)
				Expect(possbileValuesModel).ToNot(BeNil())
				possbileValuesModel.DisplayValue = core.StringPtr("testString")
				possbileValuesModel.Value = core.StringPtr("testString")
				Expect(possbileValuesModel.DisplayValue).To(Equal(core.StringPtr("testString")))
				Expect(possbileValuesModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AttributeEntry model
				attributeEntryModel := new(insightsformedicalliteratureservicev1.AttributeEntry)
				Expect(attributeEntryModel).ToNot(BeNil())
				attributeEntryModel.AttrName = core.StringPtr("testString")
				attributeEntryModel.DataType = core.StringPtr("testString")
				attributeEntryModel.DefaultValue = core.StringPtr("testString")
				attributeEntryModel.Description = core.StringPtr("testString")
				attributeEntryModel.DisplayName = core.StringPtr("testString")
				attributeEntryModel.DocID = core.StringPtr("testString")
				attributeEntryModel.FieldValues = []string{"testString"}
				attributeEntryModel.MaximumValue = core.StringPtr("testString")
				attributeEntryModel.MinimumValue = core.StringPtr("testString")
				attributeEntryModel.MultiValue = core.BoolPtr(true)
				attributeEntryModel.Units = core.StringPtr("testString")
				attributeEntryModel.ValueType = core.StringPtr("testString")
				attributeEntryModel.PossibleValues = []insightsformedicalliteratureservicev1.PossbileValues{*possbileValuesModel}
				Expect(attributeEntryModel.AttrName).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.DataType).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.DefaultValue).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.DocID).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.FieldValues).To(Equal([]string{"testString"}))
				Expect(attributeEntryModel.MaximumValue).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.MinimumValue).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.MultiValue).To(Equal(core.BoolPtr(true)))
				Expect(attributeEntryModel.Units).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.ValueType).To(Equal(core.StringPtr("testString")))
				Expect(attributeEntryModel.PossibleValues).To(Equal([]insightsformedicalliteratureservicev1.PossbileValues{*possbileValuesModel}))

				// Construct an instance of the DictonaryEntry model
				dictonaryEntryModel := new(insightsformedicalliteratureservicev1.DictonaryEntry)
				Expect(dictonaryEntryModel).ToNot(BeNil())
				dictonaryEntryModel.Children = []string{"testString"}
				dictonaryEntryModel.Cui = core.StringPtr("testString")
				dictonaryEntryModel.Definition = []string{"testString"}
				dictonaryEntryModel.Parents = []string{"testString"}
				dictonaryEntryModel.PreferredName = core.StringPtr("testString")
				dictonaryEntryModel.Semtypes = []string{"testString"}
				dictonaryEntryModel.Siblings = []string{"testString"}
				dictonaryEntryModel.SurfaceForms = []string{"testString"}
				dictonaryEntryModel.Variants = []string{"testString"}
				dictonaryEntryModel.Vocab = core.StringPtr("testString")
				dictonaryEntryModel.Related = []string{"testString"}
				dictonaryEntryModel.Source = core.StringPtr("testString")
				dictonaryEntryModel.SourceVersion = core.StringPtr("testString")
				Expect(dictonaryEntryModel.Children).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.Cui).To(Equal(core.StringPtr("testString")))
				Expect(dictonaryEntryModel.Definition).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.Parents).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.PreferredName).To(Equal(core.StringPtr("testString")))
				Expect(dictonaryEntryModel.Semtypes).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.Siblings).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.SurfaceForms).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.Variants).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.Vocab).To(Equal(core.StringPtr("testString")))
				Expect(dictonaryEntryModel.Related).To(Equal([]string{"testString"}))
				Expect(dictonaryEntryModel.Source).To(Equal(core.StringPtr("testString")))
				Expect(dictonaryEntryModel.SourceVersion).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddArtifactOptions model
				corpus := "testString"
				addArtifactOptionsModel := testService.NewAddArtifactOptions(corpus)
				addArtifactOptionsModel.SetCorpus("testString")
				addArtifactOptionsModel.SetDictionaryEntry(dictonaryEntryModel)
				addArtifactOptionsModel.SetAttributeEntry(attributeEntryModel)
				addArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addArtifactOptionsModel).ToNot(BeNil())
				Expect(addArtifactOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(addArtifactOptionsModel.DictionaryEntry).To(Equal(dictonaryEntryModel))
				Expect(addArtifactOptionsModel.AttributeEntry).To(Equal(attributeEntryModel))
				Expect(addArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddCorpusDocumentOptions successfully`, func() {
				// Construct an instance of the AddCorpusDocumentOptions model
				corpus := "testString"
				addCorpusDocumentOptionsModel := testService.NewAddCorpusDocumentOptions(corpus)
				addCorpusDocumentOptionsModel.SetCorpus("testString")
				addCorpusDocumentOptionsModel.SetDocument(make(map[string]interface{}))
				addCorpusDocumentOptionsModel.SetAcdURL("testString")
				addCorpusDocumentOptionsModel.SetApiKey("testString")
				addCorpusDocumentOptionsModel.SetFlowID("testString")
				addCorpusDocumentOptionsModel.SetAccessToken("testString")
				addCorpusDocumentOptionsModel.SetOtherAnnotators([]interface{}{map[string]interface{}{"anyKey": "anyValue"}})
				addCorpusDocumentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addCorpusDocumentOptionsModel).ToNot(BeNil())
				Expect(addCorpusDocumentOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusDocumentOptionsModel.Document).To(Equal(make(map[string]interface{})))
				Expect(addCorpusDocumentOptionsModel.AcdURL).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusDocumentOptionsModel.ApiKey).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusDocumentOptionsModel.FlowID).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusDocumentOptionsModel.AccessToken).To(Equal(core.StringPtr("testString")))
				Expect(addCorpusDocumentOptionsModel.OtherAnnotators).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(addCorpusDocumentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCorpusSchemaOptions successfully`, func() {
				// Construct an instance of the DeleteCorpusSchemaOptions model
				instance := "testString"
				deleteCorpusSchemaOptionsModel := testService.NewDeleteCorpusSchemaOptions(instance)
				deleteCorpusSchemaOptionsModel.SetInstance("testString")
				deleteCorpusSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCorpusSchemaOptionsModel).ToNot(BeNil())
				Expect(deleteCorpusSchemaOptionsModel.Instance).To(Equal(core.StringPtr("testString")))
				Expect(deleteCorpusSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnableCorpusSearchTrackingOptions successfully`, func() {
				// Construct an instance of the EnableCorpusSearchTrackingOptions model
				enableCorpusSearchTrackingOptionsModel := testService.NewEnableCorpusSearchTrackingOptions()
				enableCorpusSearchTrackingOptionsModel.SetEnableTracking(true)
				enableCorpusSearchTrackingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(enableCorpusSearchTrackingOptionsModel).ToNot(BeNil())
				Expect(enableCorpusSearchTrackingOptionsModel.EnableTracking).To(Equal(core.BoolPtr(true)))
				Expect(enableCorpusSearchTrackingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConceptsOptions successfully`, func() {
				// Construct an instance of the GetConceptsOptions model
				corpus := "testString"
				getConceptsOptionsModel := testService.NewGetConceptsOptions(corpus)
				getConceptsOptionsModel.SetCorpus("testString")
				getConceptsOptionsModel.SetCuis([]string{"testString"})
				getConceptsOptionsModel.SetPreferredNames([]string{"testString"})
				getConceptsOptionsModel.SetSurfaceForms([]string{"testString"})
				getConceptsOptionsModel.SetAttributes([]string{"testString"})
				getConceptsOptionsModel.SetVerbose(true)
				getConceptsOptionsModel.SetSort("hitCount")
				getConceptsOptionsModel.SetLimit(int64(38))
				getConceptsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConceptsOptionsModel).ToNot(BeNil())
				Expect(getConceptsOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getConceptsOptionsModel.Cuis).To(Equal([]string{"testString"}))
				Expect(getConceptsOptionsModel.PreferredNames).To(Equal([]string{"testString"}))
				Expect(getConceptsOptionsModel.SurfaceForms).To(Equal([]string{"testString"}))
				Expect(getConceptsOptionsModel.Attributes).To(Equal([]string{"testString"}))
				Expect(getConceptsOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(getConceptsOptionsModel.Sort).To(Equal(core.StringPtr("hitCount")))
				Expect(getConceptsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getConceptsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCorporaConfigOptions successfully`, func() {
				// Construct an instance of the GetCorporaConfigOptions model
				getCorporaConfigOptionsModel := testService.NewGetCorporaConfigOptions()
				getCorporaConfigOptionsModel.SetVerbose(true)
				getCorporaConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCorporaConfigOptionsModel).ToNot(BeNil())
				Expect(getCorporaConfigOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(getCorporaConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCorpusConfigOptions successfully`, func() {
				// Construct an instance of the GetCorpusConfigOptions model
				corpus := "testString"
				getCorpusConfigOptionsModel := testService.NewGetCorpusConfigOptions(corpus)
				getCorpusConfigOptionsModel.SetCorpus("testString")
				getCorpusConfigOptionsModel.SetVerbose(true)
				getCorpusConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCorpusConfigOptionsModel).ToNot(BeNil())
				Expect(getCorpusConfigOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getCorpusConfigOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(getCorpusConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCuiInfoOptions successfully`, func() {
				// Construct an instance of the GetCuiInfoOptions model
				corpus := "testString"
				nameOrID := "testString"
				getCuiInfoOptionsModel := testService.NewGetCuiInfoOptions(corpus, nameOrID)
				getCuiInfoOptionsModel.SetCorpus("testString")
				getCuiInfoOptionsModel.SetNameOrID("testString")
				getCuiInfoOptionsModel.SetOntology("testString")
				getCuiInfoOptionsModel.SetFields("testString")
				getCuiInfoOptionsModel.SetTreeLayout(true)
				getCuiInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCuiInfoOptionsModel).ToNot(BeNil())
				Expect(getCuiInfoOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getCuiInfoOptionsModel.NameOrID).To(Equal(core.StringPtr("testString")))
				Expect(getCuiInfoOptionsModel.Ontology).To(Equal(core.StringPtr("testString")))
				Expect(getCuiInfoOptionsModel.Fields).To(Equal(core.StringPtr("testString")))
				Expect(getCuiInfoOptionsModel.TreeLayout).To(Equal(core.BoolPtr(true)))
				Expect(getCuiInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentAnnotationsOptions successfully`, func() {
				// Construct an instance of the GetDocumentAnnotationsOptions model
				corpus := "testString"
				documentID := "testString"
				documentSection := "testString"
				getDocumentAnnotationsOptionsModel := testService.NewGetDocumentAnnotationsOptions(corpus, documentID, documentSection)
				getDocumentAnnotationsOptionsModel.SetCorpus("testString")
				getDocumentAnnotationsOptionsModel.SetDocumentID("testString")
				getDocumentAnnotationsOptionsModel.SetDocumentSection("testString")
				getDocumentAnnotationsOptionsModel.SetCuis([]string{"testString"})
				getDocumentAnnotationsOptionsModel.SetIncludeText(true)
				getDocumentAnnotationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentAnnotationsOptionsModel).ToNot(BeNil())
				Expect(getDocumentAnnotationsOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentAnnotationsOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentAnnotationsOptionsModel.DocumentSection).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentAnnotationsOptionsModel.Cuis).To(Equal([]string{"testString"}))
				Expect(getDocumentAnnotationsOptionsModel.IncludeText).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentAnnotationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentCategoriesOptions successfully`, func() {
				// Construct an instance of the GetDocumentCategoriesOptions model
				corpus := "testString"
				documentID := "testString"
				getDocumentCategoriesOptionsModel := testService.NewGetDocumentCategoriesOptions(corpus, documentID)
				getDocumentCategoriesOptionsModel.SetCorpus("testString")
				getDocumentCategoriesOptionsModel.SetDocumentID("testString")
				getDocumentCategoriesOptionsModel.SetHighlightTagBegin("testString")
				getDocumentCategoriesOptionsModel.SetHighlightTagEnd("testString")
				getDocumentCategoriesOptionsModel.SetTypes([]string{"testString"})
				getDocumentCategoriesOptionsModel.SetCategory("disorders")
				getDocumentCategoriesOptionsModel.SetOnlyNegatedConcepts(true)
				getDocumentCategoriesOptionsModel.SetFields("testString")
				getDocumentCategoriesOptionsModel.SetLimit(int64(38))
				getDocumentCategoriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentCategoriesOptionsModel).ToNot(BeNil())
				Expect(getDocumentCategoriesOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentCategoriesOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentCategoriesOptionsModel.HighlightTagBegin).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentCategoriesOptionsModel.HighlightTagEnd).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentCategoriesOptionsModel.Types).To(Equal([]string{"testString"}))
				Expect(getDocumentCategoriesOptionsModel.Category).To(Equal(core.StringPtr("disorders")))
				Expect(getDocumentCategoriesOptionsModel.OnlyNegatedConcepts).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentCategoriesOptionsModel.Fields).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentCategoriesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getDocumentCategoriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentInfoOptions successfully`, func() {
				// Construct an instance of the GetDocumentInfoOptions model
				corpus := "testString"
				documentID := "testString"
				getDocumentInfoOptionsModel := testService.NewGetDocumentInfoOptions(corpus, documentID)
				getDocumentInfoOptionsModel.SetCorpus("testString")
				getDocumentInfoOptionsModel.SetDocumentID("testString")
				getDocumentInfoOptionsModel.SetVerbose(true)
				getDocumentInfoOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentInfoOptionsModel).ToNot(BeNil())
				Expect(getDocumentInfoOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentInfoOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentInfoOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(getDocumentInfoOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentMultipleCategoriesOptions successfully`, func() {
				// Construct an instance of the AnnotationModel model
				annotationModelModel := new(insightsformedicalliteratureservicev1.AnnotationModel)
				Expect(annotationModelModel).ToNot(BeNil())
				annotationModelModel.UniqueID = core.Int64Ptr(int64(38))
				annotationModelModel.StickyIds = []int64{int64(38)}
				annotationModelModel.Ontology = core.StringPtr("testString")
				annotationModelModel.Section = core.StringPtr("testString")
				annotationModelModel.PreferredName = core.StringPtr("testString")
				annotationModelModel.Cui = core.StringPtr("testString")
				annotationModelModel.AttributeID = core.StringPtr("testString")
				annotationModelModel.Qualifiers = []string{"testString"}
				annotationModelModel.Type = core.StringPtr("testString")
				annotationModelModel.Negated = core.BoolPtr(true)
				annotationModelModel.Hypothetical = core.BoolPtr(true)
				annotationModelModel.Unit = core.StringPtr("testString")
				annotationModelModel.MinValue = core.StringPtr("testString")
				annotationModelModel.MaxValue = core.StringPtr("testString")
				annotationModelModel.Operator = core.StringPtr("testString")
				annotationModelModel.NluSourceType = core.StringPtr("testString")
				annotationModelModel.NluRelation = core.StringPtr("testString")
				annotationModelModel.NluTargetType = core.StringPtr("testString")
				annotationModelModel.NluEntityIndex = core.StringPtr("testString")
				annotationModelModel.NluMentionIndex = core.StringPtr("testString")
				annotationModelModel.NluRelationID = core.StringPtr("testString")
				annotationModelModel.NluSide = core.StringPtr("testString")
				annotationModelModel.Begin = core.Int64Ptr(int64(38))
				annotationModelModel.End = core.Int64Ptr(int64(38))
				annotationModelModel.Score = core.Float32Ptr(36.0)
				annotationModelModel.Timestamp = core.Int64Ptr(int64(26))
				annotationModelModel.Features = make(map[string]string)
				annotationModelModel.Hits = core.Int64Ptr(int64(38))
				Expect(annotationModelModel.UniqueID).To(Equal(core.Int64Ptr(int64(38))))
				Expect(annotationModelModel.StickyIds).To(Equal([]int64{int64(38)}))
				Expect(annotationModelModel.Ontology).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.Section).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.PreferredName).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.Cui).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.AttributeID).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.Qualifiers).To(Equal([]string{"testString"}))
				Expect(annotationModelModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.Negated).To(Equal(core.BoolPtr(true)))
				Expect(annotationModelModel.Hypothetical).To(Equal(core.BoolPtr(true)))
				Expect(annotationModelModel.Unit).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.MinValue).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.MaxValue).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluSourceType).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluRelation).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluTargetType).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluEntityIndex).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluMentionIndex).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluRelationID).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.NluSide).To(Equal(core.StringPtr("testString")))
				Expect(annotationModelModel.Begin).To(Equal(core.Int64Ptr(int64(38))))
				Expect(annotationModelModel.End).To(Equal(core.Int64Ptr(int64(38))))
				Expect(annotationModelModel.Score).To(Equal(core.Float32Ptr(36.0)))
				Expect(annotationModelModel.Timestamp).To(Equal(core.Int64Ptr(int64(26))))
				Expect(annotationModelModel.Features).To(Equal(make(map[string]string)))
				Expect(annotationModelModel.Hits).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the StringBuilder model
				stringBuilderModel := new(insightsformedicalliteratureservicev1.StringBuilder)
				Expect(stringBuilderModel).ToNot(BeNil())

				// Construct an instance of the GetDocumentMultipleCategoriesOptions model
				corpus := "testString"
				documentID := "testString"
				getDocumentMultipleCategoriesOptionsModel := testService.NewGetDocumentMultipleCategoriesOptions(corpus, documentID)
				getDocumentMultipleCategoriesOptionsModel.SetCorpus("testString")
				getDocumentMultipleCategoriesOptionsModel.SetDocumentID("testString")
				getDocumentMultipleCategoriesOptionsModel.SetModelLicense("testString")
				getDocumentMultipleCategoriesOptionsModel.SetHighlightedTitle(stringBuilderModel)
				getDocumentMultipleCategoriesOptionsModel.SetHighlightedAbstract(stringBuilderModel)
				getDocumentMultipleCategoriesOptionsModel.SetHighlightedBody(stringBuilderModel)
				getDocumentMultipleCategoriesOptionsModel.SetHighlightedSections(make(map[string]insightsformedicalliteratureservicev1.StringBuilder))
				getDocumentMultipleCategoriesOptionsModel.SetPassages(make(map[string]map[string]EntryModel))
				getDocumentMultipleCategoriesOptionsModel.SetAnnotations(make(map[string]insightsformedicalliteratureservicev1.AnnotationModel))
				getDocumentMultipleCategoriesOptionsModel.SetHighlightTagBegin("testString")
				getDocumentMultipleCategoriesOptionsModel.SetHighlightTagEnd("testString")
				getDocumentMultipleCategoriesOptionsModel.SetFields("testString")
				getDocumentMultipleCategoriesOptionsModel.SetLimit(int64(38))
				getDocumentMultipleCategoriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentMultipleCategoriesOptionsModel).ToNot(BeNil())
				Expect(getDocumentMultipleCategoriesOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentMultipleCategoriesOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentMultipleCategoriesOptionsModel.ModelLicense).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentMultipleCategoriesOptionsModel.HighlightedTitle).To(Equal(stringBuilderModel))
				Expect(getDocumentMultipleCategoriesOptionsModel.HighlightedAbstract).To(Equal(stringBuilderModel))
				Expect(getDocumentMultipleCategoriesOptionsModel.HighlightedBody).To(Equal(stringBuilderModel))
				Expect(getDocumentMultipleCategoriesOptionsModel.HighlightedSections).To(Equal(make(map[string]insightsformedicalliteratureservicev1.StringBuilder)))
				Expect(getDocumentMultipleCategoriesOptionsModel.Passages).To(Equal(make(map[string]map[string]EntryModel)))
				Expect(getDocumentMultipleCategoriesOptionsModel.Annotations).To(Equal(make(map[string]insightsformedicalliteratureservicev1.AnnotationModel)))
				Expect(getDocumentMultipleCategoriesOptionsModel.HighlightTagBegin).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentMultipleCategoriesOptionsModel.HighlightTagEnd).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentMultipleCategoriesOptionsModel.Fields).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentMultipleCategoriesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getDocumentMultipleCategoriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDocumentsOptions successfully`, func() {
				// Construct an instance of the GetDocumentsOptions model
				corpus := "testString"
				getDocumentsOptionsModel := testService.NewGetDocumentsOptions(corpus)
				getDocumentsOptionsModel.SetCorpus("testString")
				getDocumentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDocumentsOptionsModel).ToNot(BeNil())
				Expect(getDocumentsOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getDocumentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetFieldsOptions successfully`, func() {
				// Construct an instance of the GetFieldsOptions model
				corpus := "testString"
				getFieldsOptionsModel := testService.NewGetFieldsOptions(corpus)
				getFieldsOptionsModel.SetCorpus("testString")
				getFieldsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getFieldsOptionsModel).ToNot(BeNil())
				Expect(getFieldsOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getFieldsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetHitCountOptions successfully`, func() {
				// Construct an instance of the GetHitCountOptions model
				corpus := "testString"
				nameOrID := "testString"
				getHitCountOptionsModel := testService.NewGetHitCountOptions(corpus, nameOrID)
				getHitCountOptionsModel.SetCorpus("testString")
				getHitCountOptionsModel.SetNameOrID("testString")
				getHitCountOptionsModel.SetOntology("testString")
				getHitCountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getHitCountOptionsModel).ToNot(BeNil())
				Expect(getHitCountOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getHitCountOptionsModel.NameOrID).To(Equal(core.StringPtr("testString")))
				Expect(getHitCountOptionsModel.Ontology).To(Equal(core.StringPtr("testString")))
				Expect(getHitCountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRelatedConceptsOptions successfully`, func() {
				// Construct an instance of the GetRelatedConceptsOptions model
				corpus := "testString"
				nameOrID := "testString"
				relationship := "children"
				getRelatedConceptsOptionsModel := testService.NewGetRelatedConceptsOptions(corpus, nameOrID, relationship)
				getRelatedConceptsOptionsModel.SetCorpus("testString")
				getRelatedConceptsOptionsModel.SetNameOrID("testString")
				getRelatedConceptsOptionsModel.SetRelationship("children")
				getRelatedConceptsOptionsModel.SetOntology("testString")
				getRelatedConceptsOptionsModel.SetRelationshipAttributes([]string{"testString"})
				getRelatedConceptsOptionsModel.SetSources([]string{"testString"})
				getRelatedConceptsOptionsModel.SetRecursive(true)
				getRelatedConceptsOptionsModel.SetTreeLayout(true)
				getRelatedConceptsOptionsModel.SetMaxDepth(int64(38))
				getRelatedConceptsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRelatedConceptsOptionsModel).ToNot(BeNil())
				Expect(getRelatedConceptsOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getRelatedConceptsOptionsModel.NameOrID).To(Equal(core.StringPtr("testString")))
				Expect(getRelatedConceptsOptionsModel.Relationship).To(Equal(core.StringPtr("children")))
				Expect(getRelatedConceptsOptionsModel.Ontology).To(Equal(core.StringPtr("testString")))
				Expect(getRelatedConceptsOptionsModel.RelationshipAttributes).To(Equal([]string{"testString"}))
				Expect(getRelatedConceptsOptionsModel.Sources).To(Equal([]string{"testString"}))
				Expect(getRelatedConceptsOptionsModel.Recursive).To(Equal(core.BoolPtr(true)))
				Expect(getRelatedConceptsOptionsModel.TreeLayout).To(Equal(core.BoolPtr(true)))
				Expect(getRelatedConceptsOptionsModel.MaxDepth).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getRelatedConceptsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSearchMatchesOptions successfully`, func() {
				// Construct an instance of the GetSearchMatchesOptions model
				corpus := "testString"
				documentID := "testString"
				minScore := 36.0
				getSearchMatchesOptionsModel := testService.NewGetSearchMatchesOptions(corpus, documentID, minScore)
				getSearchMatchesOptionsModel.SetCorpus("testString")
				getSearchMatchesOptionsModel.SetDocumentID("testString")
				getSearchMatchesOptionsModel.SetMinScore(36.0)
				getSearchMatchesOptionsModel.SetCuis([]string{"testString"})
				getSearchMatchesOptionsModel.SetText([]string{"testString"})
				getSearchMatchesOptionsModel.SetTypes([]string{"testString"})
				getSearchMatchesOptionsModel.SetAttributes([]string{"testString"})
				getSearchMatchesOptionsModel.SetValues([]string{"testString"})
				getSearchMatchesOptionsModel.SetNluRelations([]string{"testString"})
				getSearchMatchesOptionsModel.SetLimit(int64(38))
				getSearchMatchesOptionsModel.SetSearchTagBegin("testString")
				getSearchMatchesOptionsModel.SetSearchTagEnd("testString")
				getSearchMatchesOptionsModel.SetRelatedTagBegin("testString")
				getSearchMatchesOptionsModel.SetRelatedTagEnd("testString")
				getSearchMatchesOptionsModel.SetFields("testString")
				getSearchMatchesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSearchMatchesOptionsModel).ToNot(BeNil())
				Expect(getSearchMatchesOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.DocumentID).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.MinScore).To(Equal(core.Float32Ptr(36.0)))
				Expect(getSearchMatchesOptionsModel.Cuis).To(Equal([]string{"testString"}))
				Expect(getSearchMatchesOptionsModel.Text).To(Equal([]string{"testString"}))
				Expect(getSearchMatchesOptionsModel.Types).To(Equal([]string{"testString"}))
				Expect(getSearchMatchesOptionsModel.Attributes).To(Equal([]string{"testString"}))
				Expect(getSearchMatchesOptionsModel.Values).To(Equal([]string{"testString"}))
				Expect(getSearchMatchesOptionsModel.NluRelations).To(Equal([]string{"testString"}))
				Expect(getSearchMatchesOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getSearchMatchesOptionsModel.SearchTagBegin).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.SearchTagEnd).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.RelatedTagBegin).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.RelatedTagEnd).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.Fields).To(Equal(core.StringPtr("testString")))
				Expect(getSearchMatchesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSimilarConceptsOptions successfully`, func() {
				// Construct an instance of the GetSimilarConceptsOptions model
				corpus := "testString"
				nameOrID := "testString"
				returnOntologies := []string{"testString"}
				getSimilarConceptsOptionsModel := testService.NewGetSimilarConceptsOptions(corpus, nameOrID, returnOntologies)
				getSimilarConceptsOptionsModel.SetCorpus("testString")
				getSimilarConceptsOptionsModel.SetNameOrID("testString")
				getSimilarConceptsOptionsModel.SetReturnOntologies([]string{"testString"})
				getSimilarConceptsOptionsModel.SetOntology("testString")
				getSimilarConceptsOptionsModel.SetLimit(int64(38))
				getSimilarConceptsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSimilarConceptsOptionsModel).ToNot(BeNil())
				Expect(getSimilarConceptsOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(getSimilarConceptsOptionsModel.NameOrID).To(Equal(core.StringPtr("testString")))
				Expect(getSimilarConceptsOptionsModel.ReturnOntologies).To(Equal([]string{"testString"}))
				Expect(getSimilarConceptsOptionsModel.Ontology).To(Equal(core.StringPtr("testString")))
				Expect(getSimilarConceptsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getSimilarConceptsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMonitorCorpusOptions successfully`, func() {
				// Construct an instance of the MonitorCorpusOptions model
				apikey := "testString"
				monitorCorpusOptionsModel := testService.NewMonitorCorpusOptions(apikey)
				monitorCorpusOptionsModel.SetApikey("testString")
				monitorCorpusOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(monitorCorpusOptionsModel).ToNot(BeNil())
				Expect(monitorCorpusOptionsModel.Apikey).To(Equal(core.StringPtr("testString")))
				Expect(monitorCorpusOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSearchOptions successfully`, func() {
				// Construct an instance of the SearchOptions model
				corpus := "testString"
				body := "testString"
				searchOptionsModel := testService.NewSearchOptions(corpus, body)
				searchOptionsModel.SetCorpus("testString")
				searchOptionsModel.SetBody("testString")
				searchOptionsModel.SetVerbose(true)
				searchOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(searchOptionsModel).ToNot(BeNil())
				Expect(searchOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.Body).To(Equal(core.StringPtr("testString")))
				Expect(searchOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(searchOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCorpusConfigOptions successfully`, func() {
				// Construct an instance of the SetCorpusConfigOptions model
				setCorpusConfigOptionsModel := testService.NewSetCorpusConfigOptions()
				setCorpusConfigOptionsModel.SetUserName("testString")
				setCorpusConfigOptionsModel.SetPassword("testString")
				setCorpusConfigOptionsModel.SetCorpusURI("testString")
				setCorpusConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCorpusConfigOptionsModel).ToNot(BeNil())
				Expect(setCorpusConfigOptionsModel.UserName).To(Equal(core.StringPtr("testString")))
				Expect(setCorpusConfigOptionsModel.Password).To(Equal(core.StringPtr("testString")))
				Expect(setCorpusConfigOptionsModel.CorpusURI).To(Equal(core.StringPtr("testString")))
				Expect(setCorpusConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSetCorpusSchemaOptions successfully`, func() {
				// Construct an instance of the SetCorpusSchemaOptions model
				setCorpusSchemaOptionsModel := testService.NewSetCorpusSchemaOptions()
				setCorpusSchemaOptionsModel.SetEnrichmentTargets([]interface{}{map[string]interface{}{"anyKey": "anyValue"}})
				setCorpusSchemaOptionsModel.SetMetadataFields([]interface{}{map[string]interface{}{"anyKey": "anyValue"}})
				setCorpusSchemaOptionsModel.SetCorpusName("testString")
				setCorpusSchemaOptionsModel.SetReferences(make(map[string]interface{}))
				setCorpusSchemaOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(setCorpusSchemaOptionsModel).ToNot(BeNil())
				Expect(setCorpusSchemaOptionsModel.EnrichmentTargets).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(setCorpusSchemaOptionsModel.MetadataFields).To(Equal([]interface{}{map[string]interface{}{"anyKey": "anyValue"}}))
				Expect(setCorpusSchemaOptionsModel.CorpusName).To(Equal(core.StringPtr("testString")))
				Expect(setCorpusSchemaOptionsModel.References).To(Equal(make(map[string]interface{})))
				Expect(setCorpusSchemaOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTypeaheadOptions successfully`, func() {
				// Construct an instance of the TypeaheadOptions model
				corpus := "testString"
				query := "testString"
				typeaheadOptionsModel := testService.NewTypeaheadOptions(corpus, query)
				typeaheadOptionsModel.SetCorpus("testString")
				typeaheadOptionsModel.SetQuery("testString")
				typeaheadOptionsModel.SetOntologies([]string{"concepts"})
				typeaheadOptionsModel.SetTypes([]string{"testString"})
				typeaheadOptionsModel.SetCategory("disorders")
				typeaheadOptionsModel.SetVerbose(true)
				typeaheadOptionsModel.SetLimit(int64(38))
				typeaheadOptionsModel.SetMaxHitCount(int64(38))
				typeaheadOptionsModel.SetNoDuplicates(true)
				typeaheadOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(typeaheadOptionsModel).ToNot(BeNil())
				Expect(typeaheadOptionsModel.Corpus).To(Equal(core.StringPtr("testString")))
				Expect(typeaheadOptionsModel.Query).To(Equal(core.StringPtr("testString")))
				Expect(typeaheadOptionsModel.Ontologies).To(Equal([]string{"concepts"}))
				Expect(typeaheadOptionsModel.Types).To(Equal([]string{"testString"}))
				Expect(typeaheadOptionsModel.Category).To(Equal(core.StringPtr("disorders")))
				Expect(typeaheadOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(typeaheadOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(typeaheadOptionsModel.MaxHitCount).To(Equal(core.Int64Ptr(int64(38))))
				Expect(typeaheadOptionsModel.NoDuplicates).To(Equal(core.BoolPtr(true)))
				Expect(typeaheadOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
