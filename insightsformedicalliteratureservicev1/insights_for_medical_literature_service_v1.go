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

// Package insightsformedicalliteratureservicev1 : Operations and models for the InsightsForMedicalLiteratureServiceV1 service
package insightsformedicalliteratureservicev1

import (
	"encoding/json"
	"fmt"
	common "github.com/IBM/whcs-go-sdk/common"
	"github.com/IBM/go-sdk-core/v4/core"
	"reflect"
	"strings"
)

// InsightsForMedicalLiteratureServiceV1 : The Insights for Medical Literature service provides APIs that enable you to
// derive insights from a corpus of medical documents.
//
// Version: 1.0.0 2020-04-20T09:05:55Z
type InsightsForMedicalLiteratureServiceV1 struct {
	Service *core.BaseService

	// The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://insights-for-medical-literature-service.cloud.ibm.com/services/medical_insights/api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "insights_for_medical_literature_service"

// InsightsForMedicalLiteratureServiceV1Options : Service options
type InsightsForMedicalLiteratureServiceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.
	Version *string `validate:"required"`
}

// NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig : constructs an instance of InsightsForMedicalLiteratureServiceV1 with passed in options and external configuration.
func NewInsightsForMedicalLiteratureServiceV1UsingExternalConfig(options *InsightsForMedicalLiteratureServiceV1Options) (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	insightsForMedicalLiteratureService, err = NewInsightsForMedicalLiteratureServiceV1(options)
	if err != nil {
		return
	}

	err = insightsForMedicalLiteratureService.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = insightsForMedicalLiteratureService.Service.SetServiceURL(options.URL)
	}
	return
}

// NewInsightsForMedicalLiteratureServiceV1 : constructs an instance of InsightsForMedicalLiteratureServiceV1 with passed in options.
func NewInsightsForMedicalLiteratureServiceV1(options *InsightsForMedicalLiteratureServiceV1Options) (service *InsightsForMedicalLiteratureServiceV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &InsightsForMedicalLiteratureServiceV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) SetServiceURL(url string) error {
	return insightsForMedicalLiteratureService.Service.SetServiceURL(url)
}

// GetDocuments : Retrieves information about the documents in this corpus
// The response returns the following information: <ul><li>number of documents in the corpus</li><li>corpus
// provider</li></ul>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetDocuments(getDocumentsOptions *GetDocumentsOptions) (result *CorpusInfoModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentsOptions, "getDocumentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentsOptions, "getDocumentsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents"}
	pathParameters := []string{*getDocumentsOptions.Corpus}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetDocuments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorpusInfoModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddCorpusDocument : Define enrichment document
// The response returns whether the document was properly added to the corpus.  <P>This API should be used for adding a
// document to a custom corpus.<P>Example POST body:<pre>{
//   "acdUrl" :
//   "acdApiKeyl" :
//   "flowId" :
//   "document" : {
//    "doc_id" :
//    "field[n]" : "value"
//   }
//   "otherAnnotators" : [   "{    "annotatorUrl    "annotatorApiKey    "containerName   "}  ]
// }
// </pre>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) AddCorpusDocument(addCorpusDocumentOptions *AddCorpusDocumentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addCorpusDocumentOptions, "addCorpusDocumentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addCorpusDocumentOptions, "addCorpusDocumentOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents"}
	pathParameters := []string{*addCorpusDocumentOptions.Corpus}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addCorpusDocumentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "AddCorpusDocument")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))

	body := make(map[string]interface{})
	if addCorpusDocumentOptions.Document != nil {
		body["document"] = addCorpusDocumentOptions.Document
	}
	if addCorpusDocumentOptions.AcdURL != nil {
		body["acdUrl"] = addCorpusDocumentOptions.AcdURL
	}
	if addCorpusDocumentOptions.ApiKey != nil {
		body["apiKey"] = addCorpusDocumentOptions.ApiKey
	}
	if addCorpusDocumentOptions.FlowID != nil {
		body["flowId"] = addCorpusDocumentOptions.FlowID
	}
	if addCorpusDocumentOptions.AccessToken != nil {
		body["accessToken"] = addCorpusDocumentOptions.AccessToken
	}
	if addCorpusDocumentOptions.OtherAnnotators != nil {
		body["otherAnnotators"] = addCorpusDocumentOptions.OtherAnnotators
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = insightsForMedicalLiteratureService.Service.Request(request, nil)

	return
}

// GetDocumentInfo : Retrieves the external ID, title, abstract and text for a document
// The response may return the following fields:<ul><li>external ID (e.g., PubMed
// ID)</li><li>title</li><li>abstract</li><li>body</li><li>pdfUrl</li><li>referenceUrl</li><li>other
// metadata</li></ul>Note, some documents may not have an abstract, or only the abstract may be available without the
// body text.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetDocumentInfo(getDocumentInfoOptions *GetDocumentInfoOptions) (result *CorpusInfoModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentInfoOptions, "getDocumentInfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentInfoOptions, "getDocumentInfoOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents"}
	pathParameters := []string{*getDocumentInfoOptions.Corpus, *getDocumentInfoOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetDocumentInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getDocumentInfoOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*getDocumentInfoOptions.Verbose))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	print(rawResponse)
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorpusInfoModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDocumentAnnotations : Retrieves annotations for a document
// The response returns a list of all the annotations contained in the document.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetDocumentAnnotations(getDocumentAnnotationsOptions *GetDocumentAnnotationsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentAnnotationsOptions, "getDocumentAnnotationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentAnnotationsOptions, "getDocumentAnnotationsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents", "annotations"}
	pathParameters := []string{*getDocumentAnnotationsOptions.Corpus, *getDocumentAnnotationsOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentAnnotationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetDocumentAnnotations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	builder.AddQuery("document_section", fmt.Sprint(*getDocumentAnnotationsOptions.DocumentSection))
	if getDocumentAnnotationsOptions.Cuis != nil {
		builder.AddQuery("cuis", strings.Join(getDocumentAnnotationsOptions.Cuis, ","))
	}
	if getDocumentAnnotationsOptions.IncludeText != nil {
		builder.AddQuery("include_text", fmt.Sprint(*getDocumentAnnotationsOptions.IncludeText))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = insightsForMedicalLiteratureService.Service.Request(request, nil)

	return
}

// GetDocumentCategories : Categorizes concepts in a document
// The response returns a categorized list of text passages in a document.  The sentences are grouped by concept with
// the matching concepts highlighted.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetDocumentCategories(getDocumentCategoriesOptions *GetDocumentCategoriesOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentCategoriesOptions, "getDocumentCategoriesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentCategoriesOptions, "getDocumentCategoriesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents", "categories"}
	pathParameters := []string{*getDocumentCategoriesOptions.Corpus, *getDocumentCategoriesOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentCategoriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetDocumentCategories")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getDocumentCategoriesOptions.HighlightTagBegin != nil {
		builder.AddQuery("highlight_tag_begin", fmt.Sprint(*getDocumentCategoriesOptions.HighlightTagBegin))
	}
	if getDocumentCategoriesOptions.HighlightTagEnd != nil {
		builder.AddQuery("highlight_tag_end", fmt.Sprint(*getDocumentCategoriesOptions.HighlightTagEnd))
	}
	if getDocumentCategoriesOptions.Types != nil {
		builder.AddQuery("types", strings.Join(getDocumentCategoriesOptions.Types, ","))
	}
	if getDocumentCategoriesOptions.Category != nil {
		builder.AddQuery("category", fmt.Sprint(*getDocumentCategoriesOptions.Category))
	}
	if getDocumentCategoriesOptions.OnlyNegatedConcepts != nil {
		builder.AddQuery("only_negated_concepts", fmt.Sprint(*getDocumentCategoriesOptions.OnlyNegatedConcepts))
	}
	if getDocumentCategoriesOptions.Fields != nil {
		builder.AddQuery("_fields", fmt.Sprint(*getDocumentCategoriesOptions.Fields))
	}
	if getDocumentCategoriesOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getDocumentCategoriesOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCategoriesModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDocumentMultipleCategories : Categorizes concepts in a document
// The response returns multiple categorized lists of text passages in a document.  The sentences are grouped by concept
// with the matching concepts highlighted.<P>This API should be used to batch multiple categories in a single request to
// improve performance.<P>Example POST body:<pre>{
//  categories: [
//   {
//    name : 'disorders',
//    category : 'disorders'
//   },
//   {
//    name : 'drugs',
//    category : 'drugs'
//   },
//   {
//    name : 'genes',
//    category : 'genes'
//   },
//   {
//    name : 'negated',
//    category : 'negated'
//   },
//   {
//    name : 'finding','
//    types : ['Finding']
//   },
//  ]
// }
// </pre>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptions *GetDocumentMultipleCategoriesOptions) (result *CategoriesModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDocumentMultipleCategoriesOptions, "getDocumentMultipleCategoriesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDocumentMultipleCategoriesOptions, "getDocumentMultipleCategoriesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents", "categories"}
	pathParameters := []string{*getDocumentMultipleCategoriesOptions.Corpus, *getDocumentMultipleCategoriesOptions.DocumentID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDocumentMultipleCategoriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetDocumentMultipleCategories")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getDocumentMultipleCategoriesOptions.HighlightTagBegin != nil {
		builder.AddQuery("highlight_tag_begin", fmt.Sprint(*getDocumentMultipleCategoriesOptions.HighlightTagBegin))
	}
	if getDocumentMultipleCategoriesOptions.HighlightTagEnd != nil {
		builder.AddQuery("highlight_tag_end", fmt.Sprint(*getDocumentMultipleCategoriesOptions.HighlightTagEnd))
	}
	if getDocumentMultipleCategoriesOptions.Fields != nil {
		builder.AddQuery("_fields", fmt.Sprint(*getDocumentMultipleCategoriesOptions.Fields))
	}
	if getDocumentMultipleCategoriesOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getDocumentMultipleCategoriesOptions.Limit))
	}

	body := make(map[string]interface{})
	if getDocumentMultipleCategoriesOptions.Categories != nil {
		body["categories"] = getDocumentMultipleCategoriesOptions.Categories
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCategoriesModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSearchMatches : Finds concepts in a document matching a set of search concepts
// Returns matching concepts and text passages. The sentences containing each concept are returned with the concept
// highlighted. <p>Extended annotations provide additional details for  each discrete search match detected in the
// document.  An iml-annotation-id attribute is added to each highlight tag which allows an application to easily show
// the annotation details when hovering over a text span.  The iml-annotation-id may also be used to color code the text
// spans.  The ibm_annotation-id is used as a key for the returned annotations. <p>For example, a search match on the
// concept "Breast Carcinoma" will have a class name "iml-breast-carcinoma" inserted in the highlight tag, and the
// returned annotations['umls-breast_carcinoma-hypothetical'] JSON field will contain the detailed annotation data:
// <pre>{
//  "cui": "C0678222"
//  "hypothetical": true
//  "preferredName": "Breast Carcinoma"
//  "semanticType": "umls.NeoplasticProcess"
//  "source": "umls"
//  "type": "umls.NeoplasticProcess"
// }
// </pre>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetSearchMatches(getSearchMatchesOptions *GetSearchMatchesOptions) (result *SearchMatchesModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSearchMatchesOptions, "getSearchMatchesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSearchMatchesOptions, "getSearchMatchesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "documents", "search_matches"}
	pathParameters := []string{*getSearchMatchesOptions.Corpus, *getSearchMatchesOptions.DocumentID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSearchMatchesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetSearchMatches")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	builder.AddQuery("min_score", fmt.Sprint(*getSearchMatchesOptions.MinScore))
	if getSearchMatchesOptions.Cuis != nil {
		builder.AddQuery("cuis", strings.Join(getSearchMatchesOptions.Cuis, ","))
	}
	if getSearchMatchesOptions.Text != nil {
		builder.AddQuery("text", strings.Join(getSearchMatchesOptions.Text, ","))
	}
	if getSearchMatchesOptions.Types != nil {
		builder.AddQuery("types", strings.Join(getSearchMatchesOptions.Types, ","))
	}
	if getSearchMatchesOptions.Attributes != nil {
		builder.AddQuery("attributes", strings.Join(getSearchMatchesOptions.Attributes, ","))
	}
	if getSearchMatchesOptions.Values != nil {
		builder.AddQuery("values", strings.Join(getSearchMatchesOptions.Values, ","))
	}
	if getSearchMatchesOptions.NluRelations != nil {
		builder.AddQuery("nlu_relations", strings.Join(getSearchMatchesOptions.NluRelations, ","))
	}
	if getSearchMatchesOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getSearchMatchesOptions.Limit))
	}
	if getSearchMatchesOptions.SearchTagBegin != nil {
		builder.AddQuery("search_tag_begin", fmt.Sprint(*getSearchMatchesOptions.SearchTagBegin))
	}
	if getSearchMatchesOptions.SearchTagEnd != nil {
		builder.AddQuery("search_tag_end", fmt.Sprint(*getSearchMatchesOptions.SearchTagEnd))
	}
	if getSearchMatchesOptions.RelatedTagBegin != nil {
		builder.AddQuery("related_tag_begin", fmt.Sprint(*getSearchMatchesOptions.RelatedTagBegin))
	}
	if getSearchMatchesOptions.RelatedTagEnd != nil {
		builder.AddQuery("related_tag_end", fmt.Sprint(*getSearchMatchesOptions.RelatedTagEnd))
	}
	if getSearchMatchesOptions.Fields != nil {
		builder.AddQuery("_fields", fmt.Sprint(*getSearchMatchesOptions.Fields))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSearchMatchesModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCorporaConfig : Retrieves the available corpus names and configuration
// The response returns an array of available corpus names and optionally includes detailed configuation parameters.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetCorporaConfig(getCorporaConfigOptions *GetCorporaConfigOptions) (result *CorporaConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCorporaConfigOptions, "getCorporaConfigOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCorporaConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetCorporaConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getCorporaConfigOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*getCorporaConfigOptions.Verbose))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorporaConfig)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetCorpusSchema : Define service repository data model
// The response returns whether the instance schema was properly created.  <P>This API should be used for defining a
// custom corpus schema.<P>Example POST body:<pre>{
//    corpusName : 'string'
//   "enrichmentTargets" : [
//    {
//     "contentField": 'string',
//     "enrichmentField : 'string'
//    }
//   ],
//   "metadataFields" : [
//    {
//     "fieldName": 'string',
//     "fieldType : 'string'
//    }
//   ],
//   "referenceIndices" : {
//    "dictionaryIndex" : "my_umls",
//    "attributeIndex" : "my_attributes",
//    "meshIndex" : "my_mesh",
//   }
// }
// </pre>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) SetCorpusSchema(setCorpusSchemaOptions *SetCorpusSchemaOptions) (result *CorporaConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCorpusSchemaOptions, "setCorpusSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCorpusSchemaOptions, "setCorpusSchemaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCorpusSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "SetCorpusSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))

	body := make(map[string]interface{})
	if setCorpusSchemaOptions.EnrichmentTargets != nil {
		body["enrichmentTargets"] = setCorpusSchemaOptions.EnrichmentTargets
	}
	if setCorpusSchemaOptions.MetadataFields != nil {
		body["metadataFields"] = setCorpusSchemaOptions.MetadataFields
	}
	if setCorpusSchemaOptions.CorpusName != nil {
		body["corpusName"] = setCorpusSchemaOptions.CorpusName
	}
	if setCorpusSchemaOptions.References != nil {
		body["references"] = setCorpusSchemaOptions.References
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorporaConfig)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteCorpusSchema : Delete a corpus
// The response returns whether the instance schema was properly deleted.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) DeleteCorpusSchema(deleteCorpusSchemaOptions *DeleteCorpusSchemaOptions) (result *CorporaConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCorpusSchemaOptions, "deleteCorpusSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCorpusSchemaOptions, "deleteCorpusSchemaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCorpusSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "DeleteCorpusSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	builder.AddQuery("instance", fmt.Sprint(*deleteCorpusSchemaOptions.Instance))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorporaConfig)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SetCorpusConfig : Define service repository
// The response returns whether the service successfully connected to the specified repository.  <P>This API should be
// used for providing a custom enriched corpus.<P>Example POST body:<pre>{
//    userName : 'string',
//    password : 'string'
//    repositoryUri : 'uri'
// }
// </pre>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) SetCorpusConfig(setCorpusConfigOptions *SetCorpusConfigOptions) (result *CorporaConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCorpusConfigOptions, "setCorpusConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCorpusConfigOptions, "setCorpusConfigOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora/configure"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCorpusConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "SetCorpusConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))

	body := make(map[string]interface{})
	if setCorpusConfigOptions.UserName != nil {
		body["userName"] = setCorpusConfigOptions.UserName
	}
	if setCorpusConfigOptions.Password != nil {
		body["password"] = setCorpusConfigOptions.Password
	}
	if setCorpusConfigOptions.CorpusURI != nil {
		body["corpusURI"] = setCorpusConfigOptions.CorpusURI
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorporaConfig)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// MonitorCorpus : Enable monitoring for a custom instance
// This API is meant to be used for IBM Cloud automated monitoring of custom plan instances.  A service api-key with
// read only role can be submitted to enable monitoring.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) MonitorCorpus(monitorCorpusOptions *MonitorCorpusOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(monitorCorpusOptions, "monitorCorpusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(monitorCorpusOptions, "monitorCorpusOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora/monitor"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range monitorCorpusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "MonitorCorpus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	builder.AddQuery("apikey", fmt.Sprint(*monitorCorpusOptions.Apikey))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = insightsForMedicalLiteratureService.Service.Request(request, nil)

	return
}

// EnableCorpusSearchTracking : Toggle Search Activity Tracking
// The response returns whether the tracking was enabled or disabled.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) EnableCorpusSearchTracking(enableCorpusSearchTrackingOptions *EnableCorpusSearchTrackingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(enableCorpusSearchTrackingOptions, "enableCorpusSearchTrackingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora/tracking"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range enableCorpusSearchTrackingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "EnableCorpusSearchTracking")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if enableCorpusSearchTrackingOptions.EnableTracking != nil {
		builder.AddQuery("enable_tracking", fmt.Sprint(*enableCorpusSearchTrackingOptions.EnableTracking))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = insightsForMedicalLiteratureService.Service.Request(request, nil)

	return
}

// GetCorpusConfig : Retrieves the corpus configuration
// The response returns the corpus configuration.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetCorpusConfig(getCorpusConfigOptions *GetCorpusConfigOptions) (result *CorporaConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCorpusConfigOptions, "getCorpusConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCorpusConfigOptions, "getCorpusConfigOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora"}
	pathParameters := []string{*getCorpusConfigOptions.Corpus}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCorpusConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetCorpusConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getCorpusConfigOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*getCorpusConfigOptions.Verbose))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCorporaConfig)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetHealthCheckStatus : Determine if service is running correctly
// This resource differs from /status in that it will will always return a 500 error if the service state is not OK.
// This makes it simpler for service front ends (such as Datapower) to detect a failed service.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions) (result *ServiceStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHealthCheckStatusOptions, "getHealthCheckStatusOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/status/health_check"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHealthCheckStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetHealthCheckStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getHealthCheckStatusOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getHealthCheckStatusOptions.Accept))
	}

	if getHealthCheckStatusOptions.Format != nil {
		builder.AddQuery("format", fmt.Sprint(*getHealthCheckStatusOptions.Format))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceStatus)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// Search : Search for concepts, documents, and authors
// Features include:<ul><li>Concept search</li><li>Keyword search</li><li>Attributes search</li><li>Attributes
// typeahead</li><li>Regular expressions</li><li>Find passages</li><li>Selecting authors</li><li>Selecting
// providers</li><li>Date ranges: publish date</li><li>Pagination</li><li>Aggregation: authors, concepts, and
// documents</li><li>Document date histogram</li></ul>.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) Search(searchOptions *SearchOptions) (result *SearchModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(searchOptions, "searchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(searchOptions, "searchOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "search"}
	pathParameters := []string{*searchOptions.Corpus}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range searchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "Search")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if searchOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*searchOptions.Verbose))
	}

	_, err = builder.SetBodyContentJSON(searchOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSearchModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetFields : Retrieves a list of metadata fields defined in the corpus
// The response returns a list of metadata field names that can be used by the POST search API.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetFields(getFieldsOptions *GetFieldsOptions) (result *MetadataModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getFieldsOptions, "getFieldsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getFieldsOptions, "getFieldsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "search/metadata"}
	pathParameters := []string{*getFieldsOptions.Corpus}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFieldsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetFields")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMetadataModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// Typeahead : Find concepts matching the specified query string
// Searches concepts mentioned in the corpus looking for matches on the query string field. The comparison is not case
// sensitive. The main use of this method is to build query boxes that offer auto-complete, to allow users to select
// valid concepts.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) Typeahead(typeaheadOptions *TypeaheadOptions) (result *ConceptListModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(typeaheadOptions, "typeaheadOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(typeaheadOptions, "typeaheadOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "search/typeahead"}
	pathParameters := []string{*typeaheadOptions.Corpus}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range typeaheadOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "Typeahead")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	builder.AddQuery("query", fmt.Sprint(*typeaheadOptions.Query))
	if typeaheadOptions.Ontologies != nil {
		builder.AddQuery("ontologies", strings.Join(typeaheadOptions.Ontologies, ","))
	}
	if typeaheadOptions.Types != nil {
		builder.AddQuery("types", strings.Join(typeaheadOptions.Types, ","))
	}
	if typeaheadOptions.Category != nil {
		builder.AddQuery("category", fmt.Sprint(*typeaheadOptions.Category))
	}
	if typeaheadOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*typeaheadOptions.Verbose))
	}
	if typeaheadOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*typeaheadOptions.Limit))
	}
	if typeaheadOptions.MaxHitCount != nil {
		builder.AddQuery("max_hit_count", fmt.Sprint(*typeaheadOptions.MaxHitCount))
	}
	if typeaheadOptions.NoDuplicates != nil {
		builder.AddQuery("no_duplicates", fmt.Sprint(*typeaheadOptions.NoDuplicates))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConceptListModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetConcepts : Retrieves information for concepts mentioned in this corpus
// The response returns concepts mentioned in this corpus.  The returned concepts may be selected by CUI, preferred
// name, suface forms and attribute name.  All selected concepts are returned.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetConcepts(getConceptsOptions *GetConceptsOptions) (result *ConceptListModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConceptsOptions, "getConceptsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConceptsOptions, "getConceptsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "concepts"}
	pathParameters := []string{*getConceptsOptions.Corpus}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConceptsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetConcepts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getConceptsOptions.Cuis != nil {
		builder.AddQuery("cuis", strings.Join(getConceptsOptions.Cuis, ","))
	}
	if getConceptsOptions.PreferredNames != nil {
		builder.AddQuery("preferred_names", strings.Join(getConceptsOptions.PreferredNames, ","))
	}
	if getConceptsOptions.SurfaceForms != nil {
		builder.AddQuery("surface_forms", strings.Join(getConceptsOptions.SurfaceForms, ","))
	}
	if getConceptsOptions.Attributes != nil {
		builder.AddQuery("attributes", strings.Join(getConceptsOptions.Attributes, ","))
	}
	if getConceptsOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*getConceptsOptions.Verbose))
	}
	if getConceptsOptions.Sort != nil {
		builder.AddQuery("_sort", fmt.Sprint(*getConceptsOptions.Sort))
	}
	if getConceptsOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getConceptsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConceptListModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddArtifact : Add cartridge artifact
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) AddArtifact(addArtifactOptions *AddArtifactOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addArtifactOptions, "addArtifactOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addArtifactOptions, "addArtifactOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "concepts/definitions"}
	pathParameters := []string{*addArtifactOptions.Corpus}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addArtifactOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "AddArtifact")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))

	body := make(map[string]interface{})
	if addArtifactOptions.DictionaryEntry != nil {
		body["dictionaryEntry"] = addArtifactOptions.DictionaryEntry
	}
	if addArtifactOptions.AttributeEntry != nil {
		body["attributeEntry"] = addArtifactOptions.AttributeEntry
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = insightsForMedicalLiteratureService.Service.Request(request, nil)

	return
}

// GetCuiInfo : Retrieve information for a concept
// The followning fields may be retrieved: <ul><li>Preferred name</li><li>Semantic types</li><li>Surface forms -
// Ontology Dictionary names for this concept</li><li>Definition - Concept definition (if available)</li><li>Related
// Concepts info</li></ul><P>The default is to return all fields.  Individual fields may be selected using the '_fields'
// query parameter.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetCuiInfo(getCuiInfoOptions *GetCuiInfoOptions) (result *ConceptInfoModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCuiInfoOptions, "getCuiInfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCuiInfoOptions, "getCuiInfoOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "concepts"}
	pathParameters := []string{*getCuiInfoOptions.Corpus, *getCuiInfoOptions.NameOrID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCuiInfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetCuiInfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getCuiInfoOptions.Ontology != nil {
		builder.AddQuery("ontology", fmt.Sprint(*getCuiInfoOptions.Ontology))
	}
	if getCuiInfoOptions.Fields != nil {
		builder.AddQuery("_fields", fmt.Sprint(*getCuiInfoOptions.Fields))
	}
	if getCuiInfoOptions.TreeLayout != nil {
		builder.AddQuery("tree_layout", fmt.Sprint(*getCuiInfoOptions.TreeLayout))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConceptInfoModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetHitCount : Retrieves a count of the number of times a concept is mentioned in the corpus
// The response returns the number of times a concept is mentioned (hit count) in the corpus.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetHitCount(getHitCountOptions *GetHitCountOptions) (result *HitCount, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getHitCountOptions, "getHitCountOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getHitCountOptions, "getHitCountOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "concepts", "hit_count"}
	pathParameters := []string{*getHitCountOptions.Corpus, *getHitCountOptions.NameOrID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHitCountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetHitCount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getHitCountOptions.Ontology != nil {
		builder.AddQuery("ontology", fmt.Sprint(*getHitCountOptions.Ontology))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHitCount)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRelatedConcepts : Retrieve concepts related to a concept
// Returns a list of related concepts mentioned in the specified corpus. The following relationships are suppored:
// <ul><li><b>children</b> child concepts</li><li><b>parents</b> parent concepts</li><li><b>siblings</b> sibling
// concepts</li><li><b>synonyms</b> synonym concepts</li><li><b>qualified by</b> qualified by
// concepts</li><li><b>broader</b> broader concepts</li><li><b>narrower</b> narrower concepts</li><li><b>other</b> other
// than synonyms, narrower or broader</li><li><b>related</b> related and posibly synonymous concepts</li></ul><p>If the
// corpus path parameter can be set to 'umls' to look up relationship in the entire UMLS dictionary.  Otherwise, an
// actual corpus name may be specified to limit the output to only those concepts mentioned in a specific corpus.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetRelatedConcepts(getRelatedConceptsOptions *GetRelatedConceptsOptions) (result *RelatedConceptsModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRelatedConceptsOptions, "getRelatedConceptsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRelatedConceptsOptions, "getRelatedConceptsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "concepts", "related_concepts"}
	pathParameters := []string{*getRelatedConceptsOptions.Corpus, *getRelatedConceptsOptions.NameOrID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRelatedConceptsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetRelatedConcepts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("relationship", fmt.Sprint(*getRelatedConceptsOptions.Relationship))
	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	if getRelatedConceptsOptions.Ontology != nil {
		builder.AddQuery("ontology", fmt.Sprint(*getRelatedConceptsOptions.Ontology))
	}
	if getRelatedConceptsOptions.RelationshipAttributes != nil {
		builder.AddQuery("relationship_attributes", strings.Join(getRelatedConceptsOptions.RelationshipAttributes, ","))
	}
	if getRelatedConceptsOptions.Sources != nil {
		builder.AddQuery("sources", strings.Join(getRelatedConceptsOptions.Sources, ","))
	}
	if getRelatedConceptsOptions.Recursive != nil {
		builder.AddQuery("recursive", fmt.Sprint(*getRelatedConceptsOptions.Recursive))
	}
	if getRelatedConceptsOptions.TreeLayout != nil {
		builder.AddQuery("tree_layout", fmt.Sprint(*getRelatedConceptsOptions.TreeLayout))
	}
	if getRelatedConceptsOptions.MaxDepth != nil {
		builder.AddQuery("max_depth", fmt.Sprint(*getRelatedConceptsOptions.MaxDepth))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRelatedConceptsModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSimilarConcepts : Find similar concepts
// The response returns a list of similar concepts.   All ontologies defined in the corpora are searched.  Similarity is
// determined by checking for overlapping surface forms.  The results are sorted in descending order by hit count.
func (insightsForMedicalLiteratureService *InsightsForMedicalLiteratureServiceV1) GetSimilarConcepts(getSimilarConceptsOptions *GetSimilarConceptsOptions) (result *ConceptListModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSimilarConceptsOptions, "getSimilarConceptsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSimilarConceptsOptions, "getSimilarConceptsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/corpora", "concepts", "similar_concepts"}
	pathParameters := []string{*getSimilarConceptsOptions.Corpus, *getSimilarConceptsOptions.NameOrID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(insightsForMedicalLiteratureService.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSimilarConceptsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("insights_for_medical_literature_service", "V1", "GetSimilarConcepts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*insightsForMedicalLiteratureService.Version))
	builder.AddQuery("return_ontologies", strings.Join(getSimilarConceptsOptions.ReturnOntologies, ","))
	if getSimilarConceptsOptions.Ontology != nil {
		builder.AddQuery("ontology", fmt.Sprint(*getSimilarConceptsOptions.Ontology))
	}
	if getSimilarConceptsOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getSimilarConceptsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = insightsForMedicalLiteratureService.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConceptListModel)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AddArtifactOptions : The AddArtifact options.
type AddArtifactOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	DictionaryEntry *DictonaryEntry `json:"dictionaryEntry,omitempty"`

	AttributeEntry *AttributeEntry `json:"attributeEntry,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddArtifactOptions : Instantiate AddArtifactOptions
func (*InsightsForMedicalLiteratureServiceV1) NewAddArtifactOptions(corpus string) *AddArtifactOptions {
	return &AddArtifactOptions{
		Corpus: core.StringPtr(corpus),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *AddArtifactOptions) SetCorpus(corpus string) *AddArtifactOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDictionaryEntry : Allow user to set DictionaryEntry
func (options *AddArtifactOptions) SetDictionaryEntry(dictionaryEntry *DictonaryEntry) *AddArtifactOptions {
	options.DictionaryEntry = dictionaryEntry
	return options
}

// SetAttributeEntry : Allow user to set AttributeEntry
func (options *AddArtifactOptions) SetAttributeEntry(attributeEntry *AttributeEntry) *AddArtifactOptions {
	options.AttributeEntry = attributeEntry
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddArtifactOptions) SetHeaders(param map[string]string) *AddArtifactOptions {
	options.Headers = param
	return options
}

// AddCorpusDocumentOptions : The AddCorpusDocument options.
type AddCorpusDocumentOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// JSON based document for enrichment.
	Document map[string]interface{} `json:"document,omitempty"`

	// Annotator for clincial data url.
	AcdURL *string `json:"acdUrl,omitempty"`

	// Security key.
	ApiKey *string `json:"apiKey,omitempty"`

	// Enrichment flow identifier.
	FlowID *string `json:"flowId,omitempty"`

	// Cloud access token.
	AccessToken *string `json:"accessToken,omitempty"`

	// URLs and API keys for custom annotators.
	OtherAnnotators []interface{} `json:"otherAnnotators,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddCorpusDocumentOptions : Instantiate AddCorpusDocumentOptions
func (*InsightsForMedicalLiteratureServiceV1) NewAddCorpusDocumentOptions(corpus string) *AddCorpusDocumentOptions {
	return &AddCorpusDocumentOptions{
		Corpus: core.StringPtr(corpus),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *AddCorpusDocumentOptions) SetCorpus(corpus string) *AddCorpusDocumentOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDocument : Allow user to set Document
func (options *AddCorpusDocumentOptions) SetDocument(document map[string]interface{}) *AddCorpusDocumentOptions {
	options.Document = document
	return options
}

// SetAcdURL : Allow user to set AcdURL
func (options *AddCorpusDocumentOptions) SetAcdURL(acdURL string) *AddCorpusDocumentOptions {
	options.AcdURL = core.StringPtr(acdURL)
	return options
}

// SetApiKey : Allow user to set ApiKey
func (options *AddCorpusDocumentOptions) SetApiKey(apiKey string) *AddCorpusDocumentOptions {
	options.ApiKey = core.StringPtr(apiKey)
	return options
}

// SetFlowID : Allow user to set FlowID
func (options *AddCorpusDocumentOptions) SetFlowID(flowID string) *AddCorpusDocumentOptions {
	options.FlowID = core.StringPtr(flowID)
	return options
}

// SetAccessToken : Allow user to set AccessToken
func (options *AddCorpusDocumentOptions) SetAccessToken(accessToken string) *AddCorpusDocumentOptions {
	options.AccessToken = core.StringPtr(accessToken)
	return options
}

// SetOtherAnnotators : Allow user to set OtherAnnotators
func (options *AddCorpusDocumentOptions) SetOtherAnnotators(otherAnnotators []interface{}) *AddCorpusDocumentOptions {
	options.OtherAnnotators = otherAnnotators
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddCorpusDocumentOptions) SetHeaders(param map[string]string) *AddCorpusDocumentOptions {
	options.Headers = param
	return options
}

// AttributeEntry : AttributeEntry struct
type AttributeEntry struct {
	AttrName *string `json:"attr_name,omitempty"`

	DataType *string `json:"data_type,omitempty"`

	DefaultValue *string `json:"default_value,omitempty"`

	Description *string `json:"description,omitempty"`

	DisplayName *string `json:"display_name,omitempty"`

	DocID *string `json:"doc_id,omitempty"`

	FieldValues []string `json:"field_values,omitempty"`

	MaximumValue *string `json:"maximum_value,omitempty"`

	MinimumValue *string `json:"minimum_value,omitempty"`

	MultiValue *bool `json:"multi_value,omitempty"`

	Units *string `json:"units,omitempty"`

	ValueType *string `json:"valueType,omitempty"`

	PossibleValues []PossbileValues `json:"possible_values,omitempty"`
}


// UnmarshalAttributeEntry unmarshals an instance of AttributeEntry from the specified map of raw messages.
func UnmarshalAttributeEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttributeEntry)
	err = core.UnmarshalPrimitive(m, "attr_name", &obj.AttrName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_type", &obj.DataType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_value", &obj.DefaultValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "doc_id", &obj.DocID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "field_values", &obj.FieldValues)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maximum_value", &obj.MaximumValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minimum_value", &obj.MinimumValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "multi_value", &obj.MultiValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "units", &obj.Units)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "valueType", &obj.ValueType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "possible_values", &obj.PossibleValues, UnmarshalPossbileValues)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BoolOperand : BoolOperand struct
type BoolOperand struct {
	BoolOperand *string `json:"boolOperand,omitempty"`
}


// UnmarshalBoolOperand unmarshals an instance of BoolOperand from the specified map of raw messages.
func UnmarshalBoolOperand(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BoolOperand)
	err = core.UnmarshalPrimitive(m, "boolOperand", &obj.BoolOperand)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Category : Category struct
type Category struct {
	// Category return name
	Name *string `json:"name,omitempty"`

	// Category Label
	Category *string `json:"category,omitempty"`

	// Semantic types of category
	Types []string `json:"types,omitempty"`
}

func (*InsightsForMedicalLiteratureServiceV1) NewCategory(name string) *Category {
	return &Category{
		Name: core.StringPtr(name),
	}
}

func (options *Category) SetCategory(category string) *Category {
	options.Category = core.StringPtr(category)
	return options
}

func (options *Category) SetTypes(types []string) *Category {
	options.Types = types
	return options
}


// UnmarshalCategory unmarshals an instance of Category from the specified map of raw messages.
func UnmarshalCategory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Category)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesEntry : Categories struct
type CategoriesEntry struct {
	// List of categories
	Categories []Category `json:"categories,omitempty"`
}

func (*InsightsForMedicalLiteratureServiceV1) NewCategoriesEntry(categories []Category) *CategoriesEntry {
	return &CategoriesEntry{
		Categories: categories,
	}
}


// UnmarshalCategoriesEntry unmarshals an instance of CategoriesEntry from the specified map of raw messages.
func UnmarshalCategoriesEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesEntry)
	err = core.UnmarshalModel(m, "categories", &obj.Categories, UnmarshalCategory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteCorpusSchemaOptions : The DeleteCorpusSchema options.
type DeleteCorpusSchemaOptions struct {
	// corpus schema.
	Instance *string `json:"instance" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCorpusSchemaOptions : Instantiate DeleteCorpusSchemaOptions
func (*InsightsForMedicalLiteratureServiceV1) NewDeleteCorpusSchemaOptions(instance string) *DeleteCorpusSchemaOptions {
	return &DeleteCorpusSchemaOptions{
		Instance: core.StringPtr(instance),
	}
}

// SetInstance : Allow user to set Instance
func (options *DeleteCorpusSchemaOptions) SetInstance(instance string) *DeleteCorpusSchemaOptions {
	options.Instance = core.StringPtr(instance)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCorpusSchemaOptions) SetHeaders(param map[string]string) *DeleteCorpusSchemaOptions {
	options.Headers = param
	return options
}

// DictonaryEntry : DictonaryEntry struct
type DictonaryEntry struct {
	Children []string `json:"children,omitempty"`

	Cui *string `json:"cui,omitempty"`

	Definition []string `json:"definition,omitempty"`

	Parents []string `json:"parents,omitempty"`

	PreferredName *string `json:"preferredName,omitempty"`

	Semtypes []string `json:"semtypes,omitempty"`

	Siblings []string `json:"siblings,omitempty"`

	SurfaceForms []string `json:"surfaceForms,omitempty"`

	Variants []string `json:"variants,omitempty"`

	Vocab *string `json:"vocab,omitempty"`

	Related []string `json:"related,omitempty"`

	Source *string `json:"source,omitempty"`

	SourceVersion *string `json:"source_version,omitempty"`
}


// UnmarshalDictonaryEntry unmarshals an instance of DictonaryEntry from the specified map of raw messages.
func UnmarshalDictonaryEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DictonaryEntry)
	err = core.UnmarshalPrimitive(m, "children", &obj.Children)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "definition", &obj.Definition)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parents", &obj.Parents)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semtypes", &obj.Semtypes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "siblings", &obj.Siblings)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "surfaceForms", &obj.SurfaceForms)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "variants", &obj.Variants)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vocab", &obj.Vocab)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "related", &obj.Related)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_version", &obj.SourceVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnableCorpusSearchTrackingOptions : The EnableCorpusSearchTracking options.
type EnableCorpusSearchTrackingOptions struct {
	// Enable corpus read event tracking.  Default is false.
	EnableTracking *bool `json:"enable_tracking,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEnableCorpusSearchTrackingOptions : Instantiate EnableCorpusSearchTrackingOptions
func (*InsightsForMedicalLiteratureServiceV1) NewEnableCorpusSearchTrackingOptions() *EnableCorpusSearchTrackingOptions {
	return &EnableCorpusSearchTrackingOptions{}
}

// SetEnableTracking : Allow user to set EnableTracking
func (options *EnableCorpusSearchTrackingOptions) SetEnableTracking(enableTracking bool) *EnableCorpusSearchTrackingOptions {
	options.EnableTracking = core.BoolPtr(enableTracking)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EnableCorpusSearchTrackingOptions) SetHeaders(param map[string]string) *EnableCorpusSearchTrackingOptions {
	options.Headers = param
	return options
}

// GetConceptsOptions : The GetConcepts options.
type GetConceptsOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Select concepts with the specified CUIs. Each cui is assumed to be from UMLS unless an ontology is explicitly
	// specified using the syntax [ontology:]cui, e.g., 'concepts:C0018787'.
	Cuis []string `json:"cuis,omitempty"`

	// Select concepts with the specified preferred names. Each preferred name is assumed to be from UMLS unless an
	// ontology is explicitly specified using the syntax [ontology:::]preferred_name, e.g., 'concepts:::HEART'.
	PreferredNames []string `json:"preferred_names,omitempty"`

	// Select all concepts having these surface forms. The match is case insensitive. Each surface form is matched against
	// UMLS unless an ontology is explicitly specified using the syntax [ontology:::]surface_form, e.g., 'concepts:::heart
	// attack'.
	SurfaceForms []string `json:"surface_forms,omitempty"`

	// Select all concepts having these attributes. The match is case insensitive.
	Attributes []string `json:"attributes,omitempty"`

	// Verbose output.  Default is false.
	Verbose *bool `json:"verbose,omitempty"`

	// Sort by hitCount (in document count).  Set to ascending order (_sort=+hitCount) or descending order
	// (_sort=-hitCount).
	Sort *string `json:"_sort,omitempty"`

	// Number of possible concepts to return. Default is 250.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConceptsOptions : Instantiate GetConceptsOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetConceptsOptions(corpus string) *GetConceptsOptions {
	return &GetConceptsOptions{
		Corpus: core.StringPtr(corpus),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetConceptsOptions) SetCorpus(corpus string) *GetConceptsOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetCuis : Allow user to set Cuis
func (options *GetConceptsOptions) SetCuis(cuis []string) *GetConceptsOptions {
	options.Cuis = cuis
	return options
}

// SetPreferredNames : Allow user to set PreferredNames
func (options *GetConceptsOptions) SetPreferredNames(preferredNames []string) *GetConceptsOptions {
	options.PreferredNames = preferredNames
	return options
}

// SetSurfaceForms : Allow user to set SurfaceForms
func (options *GetConceptsOptions) SetSurfaceForms(surfaceForms []string) *GetConceptsOptions {
	options.SurfaceForms = surfaceForms
	return options
}

// SetAttributes : Allow user to set Attributes
func (options *GetConceptsOptions) SetAttributes(attributes []string) *GetConceptsOptions {
	options.Attributes = attributes
	return options
}

// SetVerbose : Allow user to set Verbose
func (options *GetConceptsOptions) SetVerbose(verbose bool) *GetConceptsOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetSort : Allow user to set Sort
func (options *GetConceptsOptions) SetSort(sort string) *GetConceptsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetConceptsOptions) SetLimit(limit int64) *GetConceptsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetConceptsOptions) SetHeaders(param map[string]string) *GetConceptsOptions {
	options.Headers = param
	return options
}

// GetCorporaConfigOptions : The GetCorporaConfig options.
type GetCorporaConfigOptions struct {
	// Verbose output.  Default verbose = false.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCorporaConfigOptions : Instantiate GetCorporaConfigOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetCorporaConfigOptions() *GetCorporaConfigOptions {
	return &GetCorporaConfigOptions{}
}

// SetVerbose : Allow user to set Verbose
func (options *GetCorporaConfigOptions) SetVerbose(verbose bool) *GetCorporaConfigOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCorporaConfigOptions) SetHeaders(param map[string]string) *GetCorporaConfigOptions {
	options.Headers = param
	return options
}

// GetCorpusConfigOptions : The GetCorpusConfig options.
type GetCorpusConfigOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Verbose output.  Default verbose = false.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCorpusConfigOptions : Instantiate GetCorpusConfigOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetCorpusConfigOptions(corpus string) *GetCorpusConfigOptions {
	return &GetCorpusConfigOptions{
		Corpus: core.StringPtr(corpus),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetCorpusConfigOptions) SetCorpus(corpus string) *GetCorpusConfigOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetVerbose : Allow user to set Verbose
func (options *GetCorpusConfigOptions) SetVerbose(verbose bool) *GetCorpusConfigOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCorpusConfigOptions) SetHeaders(param map[string]string) *GetCorpusConfigOptions {
	options.Headers = param
	return options
}

// GetCuiInfoOptions : The GetCuiInfo options.
type GetCuiInfoOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Preferred name or concept ID.
	NameOrID *string `json:"name_or_id" validate:"required"`

	// The ontology that defines the cui.
	Ontology *string `json:"ontology,omitempty"`

	// Comma separated list of fields to return: preferredName, semanticTypes, surfaceForms, typeahead, variants,
	// definition.  Defaults to all fields.
	Fields *string `json:"_fields,omitempty"`

	// Generate JSON output that is compatible with a d3 tree layout.  Default is false.
	TreeLayout *bool `json:"tree_layout,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCuiInfoOptions : Instantiate GetCuiInfoOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetCuiInfoOptions(corpus string, nameOrID string) *GetCuiInfoOptions {
	return &GetCuiInfoOptions{
		Corpus: core.StringPtr(corpus),
		NameOrID: core.StringPtr(nameOrID),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetCuiInfoOptions) SetCorpus(corpus string) *GetCuiInfoOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetNameOrID : Allow user to set NameOrID
func (options *GetCuiInfoOptions) SetNameOrID(nameOrID string) *GetCuiInfoOptions {
	options.NameOrID = core.StringPtr(nameOrID)
	return options
}

// SetOntology : Allow user to set Ontology
func (options *GetCuiInfoOptions) SetOntology(ontology string) *GetCuiInfoOptions {
	options.Ontology = core.StringPtr(ontology)
	return options
}

// SetFields : Allow user to set Fields
func (options *GetCuiInfoOptions) SetFields(fields string) *GetCuiInfoOptions {
	options.Fields = core.StringPtr(fields)
	return options
}

// SetTreeLayout : Allow user to set TreeLayout
func (options *GetCuiInfoOptions) SetTreeLayout(treeLayout bool) *GetCuiInfoOptions {
	options.TreeLayout = core.BoolPtr(treeLayout)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCuiInfoOptions) SetHeaders(param map[string]string) *GetCuiInfoOptions {
	options.Headers = param
	return options
}

// GetDocumentAnnotationsOptions : The GetDocumentAnnotations options.
type GetDocumentAnnotationsOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Document ID.
	DocumentID *string `json:"document_id" validate:"required"`

	// Document section to annotate. (e.g., title, abstract, body...
	DocumentSection *string `json:"document_section" validate:"required"`

	// Concepts to show.  Defaults to all concepts.
	Cuis []string `json:"cuis,omitempty"`

	// Include document text.
	IncludeText *bool `json:"include_text,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDocumentAnnotationsOptions : Instantiate GetDocumentAnnotationsOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetDocumentAnnotationsOptions(corpus string, documentID string, documentSection string) *GetDocumentAnnotationsOptions {
	return &GetDocumentAnnotationsOptions{
		Corpus: core.StringPtr(corpus),
		DocumentID: core.StringPtr(documentID),
		DocumentSection: core.StringPtr(documentSection),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetDocumentAnnotationsOptions) SetCorpus(corpus string) *GetDocumentAnnotationsOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentAnnotationsOptions) SetDocumentID(documentID string) *GetDocumentAnnotationsOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetDocumentSection : Allow user to set DocumentSection
func (options *GetDocumentAnnotationsOptions) SetDocumentSection(documentSection string) *GetDocumentAnnotationsOptions {
	options.DocumentSection = core.StringPtr(documentSection)
	return options
}

// SetCuis : Allow user to set Cuis
func (options *GetDocumentAnnotationsOptions) SetCuis(cuis []string) *GetDocumentAnnotationsOptions {
	options.Cuis = cuis
	return options
}

// SetIncludeText : Allow user to set IncludeText
func (options *GetDocumentAnnotationsOptions) SetIncludeText(includeText bool) *GetDocumentAnnotationsOptions {
	options.IncludeText = core.BoolPtr(includeText)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentAnnotationsOptions) SetHeaders(param map[string]string) *GetDocumentAnnotationsOptions {
	options.Headers = param
	return options
}

// GetDocumentCategoriesOptions : The GetDocumentCategories options.
type GetDocumentCategoriesOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Document ID.
	DocumentID *string `json:"document_id" validate:"required"`

	// HTML tag used to highlight concepts found in the text.  Default is '&ltb&gt'.
	HighlightTagBegin *string `json:"highlight_tag_begin,omitempty"`

	// HTML tag used to highlight concepts found in the text.  Default is '&lt/b&gt'.
	HighlightTagEnd *string `json:"highlight_tag_end,omitempty"`

	// Select concepts belonging to these semantic types to return. Semantic types for the corpus can be found using the
	// /v1/corpora/{corpus}/types method.Defaults to 'all'.
	Types []string `json:"types,omitempty"`

	// Select concepts belonging to disorders, drugs or genes.
	Category *string `json:"category,omitempty"`

	// Only return negated concepts?.
	OnlyNegatedConcepts *bool `json:"only_negated_concepts,omitempty"`

	// Comma separated list of fields to return:  passages, annotations, highlightedTitle, highlightedAbstract,
	// highlightedBody, highlightedSections.
	Fields *string `json:"_fields,omitempty"`

	// Limit the number of passages per search concept (1 to 250).  Default is 50.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetDocumentCategoriesOptions.Category property.
// Select concepts belonging to disorders, drugs or genes.
const (
	GetDocumentCategoriesOptions_Category_Disorders = "disorders"
	GetDocumentCategoriesOptions_Category_Drugs = "drugs"
	GetDocumentCategoriesOptions_Category_Genes = "genes"
)

// NewGetDocumentCategoriesOptions : Instantiate GetDocumentCategoriesOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetDocumentCategoriesOptions(corpus string, documentID string) *GetDocumentCategoriesOptions {
	return &GetDocumentCategoriesOptions{
		Corpus: core.StringPtr(corpus),
		DocumentID: core.StringPtr(documentID),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetDocumentCategoriesOptions) SetCorpus(corpus string) *GetDocumentCategoriesOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentCategoriesOptions) SetDocumentID(documentID string) *GetDocumentCategoriesOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetHighlightTagBegin : Allow user to set HighlightTagBegin
func (options *GetDocumentCategoriesOptions) SetHighlightTagBegin(highlightTagBegin string) *GetDocumentCategoriesOptions {
	options.HighlightTagBegin = core.StringPtr(highlightTagBegin)
	return options
}

// SetHighlightTagEnd : Allow user to set HighlightTagEnd
func (options *GetDocumentCategoriesOptions) SetHighlightTagEnd(highlightTagEnd string) *GetDocumentCategoriesOptions {
	options.HighlightTagEnd = core.StringPtr(highlightTagEnd)
	return options
}

// SetTypes : Allow user to set Types
func (options *GetDocumentCategoriesOptions) SetTypes(types []string) *GetDocumentCategoriesOptions {
	options.Types = types
	return options
}

// SetCategory : Allow user to set Category
func (options *GetDocumentCategoriesOptions) SetCategory(category string) *GetDocumentCategoriesOptions {
	options.Category = core.StringPtr(category)
	return options
}

// SetOnlyNegatedConcepts : Allow user to set OnlyNegatedConcepts
func (options *GetDocumentCategoriesOptions) SetOnlyNegatedConcepts(onlyNegatedConcepts bool) *GetDocumentCategoriesOptions {
	options.OnlyNegatedConcepts = core.BoolPtr(onlyNegatedConcepts)
	return options
}

// SetFields : Allow user to set Fields
func (options *GetDocumentCategoriesOptions) SetFields(fields string) *GetDocumentCategoriesOptions {
	options.Fields = core.StringPtr(fields)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetDocumentCategoriesOptions) SetLimit(limit int64) *GetDocumentCategoriesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentCategoriesOptions) SetHeaders(param map[string]string) *GetDocumentCategoriesOptions {
	options.Headers = param
	return options
}

// GetDocumentInfoOptions : The GetDocumentInfo options.
type GetDocumentInfoOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Document ID.
	DocumentID *string `json:"document_id" validate:"required"`

	// Verbose output. If true, text for all document sections is returned.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDocumentInfoOptions : Instantiate GetDocumentInfoOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetDocumentInfoOptions(corpus string, documentID string) *GetDocumentInfoOptions {
	return &GetDocumentInfoOptions{
		Corpus: core.StringPtr(corpus),
		DocumentID: core.StringPtr(documentID),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetDocumentInfoOptions) SetCorpus(corpus string) *GetDocumentInfoOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentInfoOptions) SetDocumentID(documentID string) *GetDocumentInfoOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetVerbose : Allow user to set Verbose
func (options *GetDocumentInfoOptions) SetVerbose(verbose bool) *GetDocumentInfoOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentInfoOptions) SetHeaders(param map[string]string) *GetDocumentInfoOptions {
	options.Headers = param
	return options
}

// GetDocumentInfoResponse : GetDocumentInfoResponse struct
type GetDocumentInfoResponse struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}


// SetProperty allows the user to set an arbitrary property on an instance of GetDocumentInfoResponse
func (o *GetDocumentInfoResponse) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of GetDocumentInfoResponse
func (o *GetDocumentInfoResponse) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of GetDocumentInfoResponse
func (o *GetDocumentInfoResponse) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of GetDocumentInfoResponse
func (o *GetDocumentInfoResponse) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalGetDocumentInfoResponse unmarshals an instance of GetDocumentInfoResponse from the specified map of raw messages.
func UnmarshalGetDocumentInfoResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetDocumentInfoResponse)
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDocumentMultipleCategoriesOptions : The GetDocumentMultipleCategories options.
type GetDocumentMultipleCategoriesOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Document ID.
	DocumentID *string `json:"document_id" validate:"required"`

	// Categories.
	Categories []Category `json:"categories" validate:"required"`

	// HTML tag used to highlight concepts found in the text.  Default is '&ltb&gt'.
	HighlightTagBegin *string `json:"highlight_tag_begin,omitempty"`

	// HTML tag used to highlight concepts found in the text.  Default is '&lt/b&gt'.
	HighlightTagEnd *string `json:"highlight_tag_end,omitempty"`

	// Comma separated list of fields to return:  passages, annotations, highlightedTitle, highlightedAbstract,
	// highlightedBody, highlightedSections.
	Fields *string `json:"_fields,omitempty"`

	// Limit the number of passages per search concept (1 to 250).  Default is 50.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDocumentMultipleCategoriesOptions : Instantiate GetDocumentMultipleCategoriesOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetDocumentMultipleCategoriesOptions(corpus string, documentID string, categories []Category) *GetDocumentMultipleCategoriesOptions {
	return &GetDocumentMultipleCategoriesOptions{
		Corpus: core.StringPtr(corpus),
		DocumentID: core.StringPtr(documentID),
		Categories: categories,
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetDocumentMultipleCategoriesOptions) SetCorpus(corpus string) *GetDocumentMultipleCategoriesOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetDocumentMultipleCategoriesOptions) SetDocumentID(documentID string) *GetDocumentMultipleCategoriesOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetBody : Allow user to set Body
func (options *GetDocumentMultipleCategoriesOptions) SetCategories(categories []Category) *GetDocumentMultipleCategoriesOptions {
	options.Categories = categories
	return options
}

// SetHighlightTagBegin : Allow user to set HighlightTagBegin
func (options *GetDocumentMultipleCategoriesOptions) SetHighlightTagBegin(highlightTagBegin string) *GetDocumentMultipleCategoriesOptions {
	options.HighlightTagBegin = core.StringPtr(highlightTagBegin)
	return options
}

// SetHighlightTagEnd : Allow user to set HighlightTagEnd
func (options *GetDocumentMultipleCategoriesOptions) SetHighlightTagEnd(highlightTagEnd string) *GetDocumentMultipleCategoriesOptions {
	options.HighlightTagEnd = core.StringPtr(highlightTagEnd)
	return options
}

// SetFields : Allow user to set Fields
func (options *GetDocumentMultipleCategoriesOptions) SetFields(fields string) *GetDocumentMultipleCategoriesOptions {
	options.Fields = core.StringPtr(fields)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetDocumentMultipleCategoriesOptions) SetLimit(limit int64) *GetDocumentMultipleCategoriesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentMultipleCategoriesOptions) SetHeaders(param map[string]string) *GetDocumentMultipleCategoriesOptions {
	options.Headers = param
	return options
}

// GetDocumentsOptions : The GetDocuments options.
type GetDocumentsOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDocumentsOptions : Instantiate GetDocumentsOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetDocumentsOptions(corpus string) *GetDocumentsOptions {
	return &GetDocumentsOptions{
		Corpus: core.StringPtr(corpus),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetDocumentsOptions) SetCorpus(corpus string) *GetDocumentsOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDocumentsOptions) SetHeaders(param map[string]string) *GetDocumentsOptions {
	options.Headers = param
	return options
}

// GetFieldsOptions : The GetFields options.
type GetFieldsOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFieldsOptions : Instantiate GetFieldsOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetFieldsOptions(corpus string) *GetFieldsOptions {
	return &GetFieldsOptions{
		Corpus: core.StringPtr(corpus),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetFieldsOptions) SetCorpus(corpus string) *GetFieldsOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFieldsOptions) SetHeaders(param map[string]string) *GetFieldsOptions {
	options.Headers = param
	return options
}

// GetHealthCheckStatusOptions : The GetHealthCheckStatus options.
type GetHealthCheckStatusOptions struct {
	// The type of the response: application/json or application/xml.
	Accept *string `json:"Accept,omitempty"`

	// Override response format.
	Format *string `json:"format,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetHealthCheckStatusOptions.Format property.
// Override response format.
const (
	GetHealthCheckStatusOptions_Format_JSON = "json"
	GetHealthCheckStatusOptions_Format_Xml = "xml"
)

// NewGetHealthCheckStatusOptions : Instantiate GetHealthCheckStatusOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetHealthCheckStatusOptions() *GetHealthCheckStatusOptions {
	return &GetHealthCheckStatusOptions{}
}

// SetAccept : Allow user to set Accept
func (options *GetHealthCheckStatusOptions) SetAccept(accept string) *GetHealthCheckStatusOptions {
	options.Accept = core.StringPtr(accept)
	return options
}

// SetFormat : Allow user to set Format
func (options *GetHealthCheckStatusOptions) SetFormat(format string) *GetHealthCheckStatusOptions {
	options.Format = core.StringPtr(format)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetHealthCheckStatusOptions) SetHeaders(param map[string]string) *GetHealthCheckStatusOptions {
	options.Headers = param
	return options
}

// GetHitCountOptions : The GetHitCount options.
type GetHitCountOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Preferred name or concept ID.
	NameOrID *string `json:"name_or_id" validate:"required"`

	// The ontology that defines the cui.
	Ontology *string `json:"ontology,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetHitCountOptions : Instantiate GetHitCountOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetHitCountOptions(corpus string, nameOrID string) *GetHitCountOptions {
	return &GetHitCountOptions{
		Corpus: core.StringPtr(corpus),
		NameOrID: core.StringPtr(nameOrID),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetHitCountOptions) SetCorpus(corpus string) *GetHitCountOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetNameOrID : Allow user to set NameOrID
func (options *GetHitCountOptions) SetNameOrID(nameOrID string) *GetHitCountOptions {
	options.NameOrID = core.StringPtr(nameOrID)
	return options
}

// SetOntology : Allow user to set Ontology
func (options *GetHitCountOptions) SetOntology(ontology string) *GetHitCountOptions {
	options.Ontology = core.StringPtr(ontology)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetHitCountOptions) SetHeaders(param map[string]string) *GetHitCountOptions {
	options.Headers = param
	return options
}

// GetRelatedConceptsOptions : The GetRelatedConcepts options.
type GetRelatedConceptsOptions struct {
	// Corpus name or null to show all ontology relationships.
	Corpus *string `json:"corpus" validate:"required"`

	// Preferred name or concept ID.
	NameOrID *string `json:"name_or_id" validate:"required"`

	// Select the relationship to retrieve.
	Relationship *string `json:"relationship" validate:"required"`

	// The ontology that defines the cui.
	Ontology *string `json:"ontology,omitempty"`

	// Select UMLS relationship attributes.  If null, all relationship attributes are returned.
	RelationshipAttributes []string `json:"relationship_attributes,omitempty"`

	// Select source vocabularies.  If null, concepts for all source vocabularies are returned.
	Sources []string `json:"sources,omitempty"`

	// Recursively return parents, children, broader and narrower relations.  Default is false.
	Recursive *bool `json:"recursive,omitempty"`

	// Generate JSON output that is compatible with a d3 tree layout.  Default is true.
	TreeLayout *bool `json:"tree_layout,omitempty"`

	// Maximum depth.  Default is 3.
	MaxDepth *int64 `json:"max_depth,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetRelatedConceptsOptions.Relationship property.
// Select the relationship to retrieve.
const (
	GetRelatedConceptsOptions_Relationship_Alike = "alike"
	GetRelatedConceptsOptions_Relationship_Allowedqualifier = "allowedQualifier"
	GetRelatedConceptsOptions_Relationship_Aq = "aq"
	GetRelatedConceptsOptions_Relationship_Broader = "broader"
	GetRelatedConceptsOptions_Relationship_Chd = "chd"
	GetRelatedConceptsOptions_Relationship_Children = "children"
	GetRelatedConceptsOptions_Relationship_Narrower = "narrower"
	GetRelatedConceptsOptions_Relationship_Notrelated = "notRelated"
	GetRelatedConceptsOptions_Relationship_Other = "other"
	GetRelatedConceptsOptions_Relationship_Par = "par"
	GetRelatedConceptsOptions_Relationship_Parents = "parents"
	GetRelatedConceptsOptions_Relationship_Qb = "qb"
	GetRelatedConceptsOptions_Relationship_Qualifiedby = "qualifiedBy"
	GetRelatedConceptsOptions_Relationship_Rb = "rb"
	GetRelatedConceptsOptions_Relationship_Related = "related"
	GetRelatedConceptsOptions_Relationship_Relatedunspecified = "relatedUnspecified"
	GetRelatedConceptsOptions_Relationship_Rl = "rl"
	GetRelatedConceptsOptions_Relationship_Rn = "rn"
	GetRelatedConceptsOptions_Relationship_Ro = "ro"
	GetRelatedConceptsOptions_Relationship_Rq = "rq"
	GetRelatedConceptsOptions_Relationship_Ru = "ru"
	GetRelatedConceptsOptions_Relationship_Sib = "sib"
	GetRelatedConceptsOptions_Relationship_Siblings = "siblings"
	GetRelatedConceptsOptions_Relationship_Sy = "sy"
	GetRelatedConceptsOptions_Relationship_Synonym = "synonym"
	GetRelatedConceptsOptions_Relationship_Xr = "xr"
)

// NewGetRelatedConceptsOptions : Instantiate GetRelatedConceptsOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetRelatedConceptsOptions(corpus string, nameOrID string, relationship string) *GetRelatedConceptsOptions {
	return &GetRelatedConceptsOptions{
		Corpus: core.StringPtr(corpus),
		NameOrID: core.StringPtr(nameOrID),
		Relationship: core.StringPtr(relationship),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetRelatedConceptsOptions) SetCorpus(corpus string) *GetRelatedConceptsOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetNameOrID : Allow user to set NameOrID
func (options *GetRelatedConceptsOptions) SetNameOrID(nameOrID string) *GetRelatedConceptsOptions {
	options.NameOrID = core.StringPtr(nameOrID)
	return options
}

// SetRelationship : Allow user to set Relationship
func (options *GetRelatedConceptsOptions) SetRelationship(relationship string) *GetRelatedConceptsOptions {
	options.Relationship = core.StringPtr(relationship)
	return options
}

// SetOntology : Allow user to set Ontology
func (options *GetRelatedConceptsOptions) SetOntology(ontology string) *GetRelatedConceptsOptions {
	options.Ontology = core.StringPtr(ontology)
	return options
}

// SetRelationshipAttributes : Allow user to set RelationshipAttributes
func (options *GetRelatedConceptsOptions) SetRelationshipAttributes(relationshipAttributes []string) *GetRelatedConceptsOptions {
	options.RelationshipAttributes = relationshipAttributes
	return options
}

// SetSources : Allow user to set Sources
func (options *GetRelatedConceptsOptions) SetSources(sources []string) *GetRelatedConceptsOptions {
	options.Sources = sources
	return options
}

// SetRecursive : Allow user to set Recursive
func (options *GetRelatedConceptsOptions) SetRecursive(recursive bool) *GetRelatedConceptsOptions {
	options.Recursive = core.BoolPtr(recursive)
	return options
}

// SetTreeLayout : Allow user to set TreeLayout
func (options *GetRelatedConceptsOptions) SetTreeLayout(treeLayout bool) *GetRelatedConceptsOptions {
	options.TreeLayout = core.BoolPtr(treeLayout)
	return options
}

// SetMaxDepth : Allow user to set MaxDepth
func (options *GetRelatedConceptsOptions) SetMaxDepth(maxDepth int64) *GetRelatedConceptsOptions {
	options.MaxDepth = core.Int64Ptr(maxDepth)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRelatedConceptsOptions) SetHeaders(param map[string]string) *GetRelatedConceptsOptions {
	options.Headers = param
	return options
}

// GetSearchMatchesOptions : The GetSearchMatches options.
type GetSearchMatchesOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Document ID (e.g, 7014026).
	DocumentID *string `json:"document_id" validate:"required"`

	// Minimum score .0 to 1.0.
	MinScore *float64 `json:"min_score" validate:"required"`

	// cui[,rank,[type]] - Example: "C0030567,10". The rank is an optional value from 0 to 10 (defalut is 10). Special rank
	// values: 0=omit, 10=require. Related concepts can also be included by appending, '-PAR' (parents), '-CHD' (children),
	// or '-SIB' (siblings) to the CUI (eg., to include all children of C0030567: 'C0030567-CHD')).  The type may
	// explicitly select a semanic type for a concept.  If no type is specified, a default type is selected.
	Cuis []string `json:"cuis,omitempty"`

	// Case insensitive text searches.
	Text []string `json:"text,omitempty"`

	// Highlight all text spans matching these semantic types.  Semantic types for the corpus can be found using the
	// /v1/corpora/{corpus}/types method.
	Types []string `json:"types,omitempty"`

	// Highlight all text spans matching these attributes.  An attribute may also specify a range value (e.g.,
	// age:years:65-100) or  a string value (e.g., gender:female).  The attribute may be qualified with one or more
	// qualifiers (e.g., Treated,Severe>>diabetes)  An attribute may target a specific CUI.  (e.g., C0003864::disease).
	Attributes []string `json:"attributes,omitempty"`

	// Highlight all text spans matching these values.  e.g., age:years:within:65-100 or gender:female  a string value
	// (e.g., gender:female).
	Values []string `json:"values,omitempty"`

	// Highlight all text spans matching these NLU relations.  e.g., druggroup,treat,indication.
	NluRelations []string `json:"nlu_relations,omitempty"`

	// Limit the number of matching passages per search concept/search term (1 to 250).  Default is 50.
	Limit *int64 `json:"_limit,omitempty"`

	// HTML tag used to highlight search concepts found in the text.  Default is '&ltb&gt'.
	SearchTagBegin *string `json:"search_tag_begin,omitempty"`

	// HTML tag used to highlight search concepts found in the text.  Default is '&lt/b&gt'.
	SearchTagEnd *string `json:"search_tag_end,omitempty"`

	// HTML tag used to highlight related concepts found in the text.
	RelatedTagBegin *string `json:"related_tag_begin,omitempty"`

	// HTML tag used to highlight related concepts found in the text.
	RelatedTagEnd *string `json:"related_tag_end,omitempty"`

	// Comma separated list of fields to return:  passages, annotations, highlightedTitle, highlightedAbstract,
	// highlightedBody, highlightedSections.
	Fields *string `json:"_fields,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSearchMatchesOptions : Instantiate GetSearchMatchesOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetSearchMatchesOptions(corpus string, documentID string, minScore float64) *GetSearchMatchesOptions {
	return &GetSearchMatchesOptions{
		Corpus: core.StringPtr(corpus),
		DocumentID: core.StringPtr(documentID),
		MinScore: core.Float64Ptr(minScore),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetSearchMatchesOptions) SetCorpus(corpus string) *GetSearchMatchesOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetDocumentID : Allow user to set DocumentID
func (options *GetSearchMatchesOptions) SetDocumentID(documentID string) *GetSearchMatchesOptions {
	options.DocumentID = core.StringPtr(documentID)
	return options
}

// SetMinScore : Allow user to set MinScore
func (options *GetSearchMatchesOptions) SetMinScore(minScore float64) *GetSearchMatchesOptions {
	options.MinScore = core.Float64Ptr(minScore)
	return options
}

// SetCuis : Allow user to set Cuis
func (options *GetSearchMatchesOptions) SetCuis(cuis []string) *GetSearchMatchesOptions {
	options.Cuis = cuis
	return options
}

// SetText : Allow user to set Text
func (options *GetSearchMatchesOptions) SetText(text []string) *GetSearchMatchesOptions {
	options.Text = text
	return options
}

// SetTypes : Allow user to set Types
func (options *GetSearchMatchesOptions) SetTypes(types []string) *GetSearchMatchesOptions {
	options.Types = types
	return options
}

// SetAttributes : Allow user to set Attributes
func (options *GetSearchMatchesOptions) SetAttributes(attributes []string) *GetSearchMatchesOptions {
	options.Attributes = attributes
	return options
}

// SetValues : Allow user to set Values
func (options *GetSearchMatchesOptions) SetValues(values []string) *GetSearchMatchesOptions {
	options.Values = values
	return options
}

// SetNluRelations : Allow user to set NluRelations
func (options *GetSearchMatchesOptions) SetNluRelations(nluRelations []string) *GetSearchMatchesOptions {
	options.NluRelations = nluRelations
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetSearchMatchesOptions) SetLimit(limit int64) *GetSearchMatchesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetSearchTagBegin : Allow user to set SearchTagBegin
func (options *GetSearchMatchesOptions) SetSearchTagBegin(searchTagBegin string) *GetSearchMatchesOptions {
	options.SearchTagBegin = core.StringPtr(searchTagBegin)
	return options
}

// SetSearchTagEnd : Allow user to set SearchTagEnd
func (options *GetSearchMatchesOptions) SetSearchTagEnd(searchTagEnd string) *GetSearchMatchesOptions {
	options.SearchTagEnd = core.StringPtr(searchTagEnd)
	return options
}

// SetRelatedTagBegin : Allow user to set RelatedTagBegin
func (options *GetSearchMatchesOptions) SetRelatedTagBegin(relatedTagBegin string) *GetSearchMatchesOptions {
	options.RelatedTagBegin = core.StringPtr(relatedTagBegin)
	return options
}

// SetRelatedTagEnd : Allow user to set RelatedTagEnd
func (options *GetSearchMatchesOptions) SetRelatedTagEnd(relatedTagEnd string) *GetSearchMatchesOptions {
	options.RelatedTagEnd = core.StringPtr(relatedTagEnd)
	return options
}

// SetFields : Allow user to set Fields
func (options *GetSearchMatchesOptions) SetFields(fields string) *GetSearchMatchesOptions {
	options.Fields = core.StringPtr(fields)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSearchMatchesOptions) SetHeaders(param map[string]string) *GetSearchMatchesOptions {
	options.Headers = param
	return options
}

// GetSimilarConceptsOptions : The GetSimilarConcepts options.
type GetSimilarConceptsOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Preferred name or concept ID.
	NameOrID *string `json:"name_or_id" validate:"required"`

	// Return similar concepts from any of these ontologites.
	ReturnOntologies []string `json:"return_ontologies" validate:"required"`

	// The ontology that defines the cui.
	Ontology *string `json:"ontology,omitempty"`

	// Number of possible concepts to return. Default is 250.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSimilarConceptsOptions : Instantiate GetSimilarConceptsOptions
func (*InsightsForMedicalLiteratureServiceV1) NewGetSimilarConceptsOptions(corpus string, nameOrID string, returnOntologies []string) *GetSimilarConceptsOptions {
	return &GetSimilarConceptsOptions{
		Corpus: core.StringPtr(corpus),
		NameOrID: core.StringPtr(nameOrID),
		ReturnOntologies: returnOntologies,
	}
}

// SetCorpus : Allow user to set Corpus
func (options *GetSimilarConceptsOptions) SetCorpus(corpus string) *GetSimilarConceptsOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetNameOrID : Allow user to set NameOrID
func (options *GetSimilarConceptsOptions) SetNameOrID(nameOrID string) *GetSimilarConceptsOptions {
	options.NameOrID = core.StringPtr(nameOrID)
	return options
}

// SetReturnOntologies : Allow user to set ReturnOntologies
func (options *GetSimilarConceptsOptions) SetReturnOntologies(returnOntologies []string) *GetSimilarConceptsOptions {
	options.ReturnOntologies = returnOntologies
	return options
}

// SetOntology : Allow user to set Ontology
func (options *GetSimilarConceptsOptions) SetOntology(ontology string) *GetSimilarConceptsOptions {
	options.Ontology = core.StringPtr(ontology)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetSimilarConceptsOptions) SetLimit(limit int64) *GetSimilarConceptsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSimilarConceptsOptions) SetHeaders(param map[string]string) *GetSimilarConceptsOptions {
	options.Headers = param
	return options
}

// Message : Object representing repository message.
type Message struct {
	// Message semantic type.
	MessageType *string `json:"messageType,omitempty"`

	// Message link.
	URL *string `json:"url,omitempty"`

	// Message request.
	Request interface{} `json:"request,omitempty"`

	// Request headers.
	Headers []string `json:"headers,omitempty"`

	// Message status.
	Status *int64 `json:"status,omitempty"`

	// Message response.
	Response interface{} `json:"response,omitempty"`
}

// Constants associated with the Message.MessageType property.
// Message semantic type.
const (
	Message_MessageType_ElasticSearch = "elastic_search"
	Message_MessageType_ExpandedRequest = "expanded_request"
)


// UnmarshalMessage unmarshals an instance of Message from the specified map of raw messages.
func UnmarshalMessage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Message)
	err = core.UnmarshalPrimitive(m, "messageType", &obj.MessageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "request", &obj.Request)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "headers", &obj.Headers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "response", &obj.Response)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetadataFields : MetadataFields struct
type MetadataFields struct {
	// Corpus name.
	Corpus *string `json:"corpus,omitempty"`

	// Corpus description.
	CorpusDescription *string `json:"corpusDescription,omitempty"`

	// Metadata fields.
	Fields map[string][]string `json:"fields,omitempty"`
}


// UnmarshalMetadataFields unmarshals an instance of MetadataFields from the specified map of raw messages.
func UnmarshalMetadataFields(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetadataFields)
	err = core.UnmarshalPrimitive(m, "corpus", &obj.Corpus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "corpusDescription", &obj.CorpusDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fields", &obj.Fields)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MonitorCorpusOptions : The MonitorCorpus options.
type MonitorCorpusOptions struct {
	// Apikey with read only permissions for monitoring.
	Apikey *string `json:"apikey" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMonitorCorpusOptions : Instantiate MonitorCorpusOptions
func (*InsightsForMedicalLiteratureServiceV1) NewMonitorCorpusOptions(apikey string) *MonitorCorpusOptions {
	return &MonitorCorpusOptions{
		Apikey: core.StringPtr(apikey),
	}
}

// SetApikey : Allow user to set Apikey
func (options *MonitorCorpusOptions) SetApikey(apikey string) *MonitorCorpusOptions {
	options.Apikey = core.StringPtr(apikey)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *MonitorCorpusOptions) SetHeaders(param map[string]string) *MonitorCorpusOptions {
	options.Headers = param
	return options
}

// PassagesModel : PassagesModel struct
type PassagesModel struct {
	// Document passages.
	AllPassages []Passage `json:"allPassages,omitempty"`

	// Search term to passages.
	SearchTermToPassages map[string][]Passage `json:"searchTermToPassages,omitempty"`
}


// UnmarshalPassagesModel unmarshals an instance of PassagesModel from the specified map of raw messages.
func UnmarshalPassagesModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PassagesModel)
	err = core.UnmarshalModel(m, "allPassages", &obj.AllPassages, UnmarshalPassage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "searchTermToPassages", &obj.SearchTermToPassages, UnmarshalPassage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PossbileValues : PossbileValues struct
type PossbileValues struct {
	DisplayValue *string `json:"displayValue,omitempty"`

	Value *string `json:"value,omitempty"`
}


// UnmarshalPossbileValues unmarshals an instance of PossbileValues from the specified map of raw messages.
func UnmarshalPossbileValues(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PossbileValues)
	err = core.UnmarshalPrimitive(m, "displayValue", &obj.DisplayValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RankedDocLinks : RankedDocLinks struct
type RankedDocLinks struct {
	// Links for search matches.
	HrefSearchMatches *string `json:"hrefSearchMatches,omitempty"`

	// Links for categorized search matches.
	HrefCategories *string `json:"hrefCategories,omitempty"`
}


// UnmarshalRankedDocLinks unmarshals an instance of RankedDocLinks from the specified map of raw messages.
func UnmarshalRankedDocLinks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RankedDocLinks)
	err = core.UnmarshalPrimitive(m, "hrefSearchMatches", &obj.HrefSearchMatches)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hrefCategories", &obj.HrefCategories)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchOptions : The Search options.
type SearchOptions struct {
	// Corpus name.
	Corpus *string `json:"corpus" validate:"required"`

	// Define the query using JSON.<script>function setSearchApiBody(e,file) {$.get(file).done(function(text) {
	// $(e.target).closest('td').prev('td').find('textarea').val(text);})}</script><p>Try sample queries:<ul
	// style='margin-top: 0; list-style:none; padding-left:0;'><li><span style='text-decoration: underline; cursor:
	// pointer;' onclick='setSearchApiBody(event, "conceptSearch.json")'><a>Concept Search</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event,
	// "conceptsAggregation.json")'><a>Concepts Aggregation</a></span></li><li><span style='text-decoration: underline;
	// cursor: pointer;' onclick='setSearchApiBody(event, "keywordSearch.json")'><a>Keyword Search</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event,
	// "attributeSearch.json")'><a>Attribute Search</a></span></li><li><span style='text-decoration: underline; cursor:
	// pointer;' onclick='setSearchApiBody(event, "types.json")'><a>Get Semantic Types</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event,
	// "attributes.json")'><a>Attributes</a></span></li><li><span style='text-decoration: underline; cursor: pointer;'
	// onclick='setSearchApiBody(event, "attributeTypeahead.json")'><a>Attribute Typeahead</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event,
	// "conceptTypeahead.json")'><a>Concept Typeahead</a></span></li><li><span style='text-decoration: underline; cursor:
	// pointer;' onclick='setSearchApiBody(event, "passages.json")'><a>Passages</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event,
	// "authorSearch.json")'><a>Author Search</a></span></li><li><span style='text-decoration: underline; cursor: pointer;'
	// onclick='setSearchApiBody(event, "authorsAggregation.json")'><a>Author Aggregation</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event,
	// "authorTypeahead.json")'><a>Author Typeahead</a></span></li><li><span style='text-decoration: underline; cursor:
	// pointer;' onclick='setSearchApiBody(event, "titleTypeahead.json")'><a>Title Typeahead</a></span></li><li><span
	// style='text-decoration: underline; cursor: pointer;' onclick='setSearchApiBody(event, "dateHistogram.json")'><a>Date
	// Histogram</a></span></li><li><span style='text-decoration: underline; cursor: pointer;'
	// onclick='setSearchApiBody(event, "dateRange.json")'><a>Date Range</a></span></li><li><span style='text-decoration:
	// underline; cursor: pointer;' onclick='setSearchApiBody(event, "advancedSearch.json")'><a>Advanced
	// Search</a></span></li></ul>.
	Returns *ReturnsModel `json:"returns" validate:"required"`

	// Query model
	Query *QueryModel `json:"query,omitempty"`

	// Verbose output.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSearchOptions : Instantiate SearchOptions
func (*InsightsForMedicalLiteratureServiceV1) NewSearchOptions(corpus string, returns *ReturnsModel) *SearchOptions {
	return &SearchOptions{
		Corpus: core.StringPtr(corpus),
		Returns: returns,
//		Body: core.StringPtr(body),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *SearchOptions) SetCorpus(corpus string) *SearchOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetQuery : Allow user to set Query model
func (options *SearchOptions) SetQuery(query *QueryModel) *SearchOptions {
	options.Query = query
	return options
}

// SetReturns : Allow user to set Returns model
func (options *SearchOptions) SetReturns(returns *ReturnsModel) *SearchOptions {
	options.Returns = returns
	return options
}

// SetVerbose : Allow user to set Verbose
func (options *SearchOptions) SetVerbose(verbose bool) *SearchOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SearchOptions) SetHeaders(param map[string]string) *SearchOptions {
	options.Headers = param
	return options
}

// SetCorpusConfigOptions : The SetCorpusConfig options.
type SetCorpusConfigOptions struct {
	// Repository connection userid.
	UserName *string `json:"userName,omitempty"`

	// Repository connection password.
	Password *string `json:"password,omitempty"`

	// Repository connection URI.
	CorpusURI *string `json:"corpusURI,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCorpusConfigOptions : Instantiate SetCorpusConfigOptions
func (*InsightsForMedicalLiteratureServiceV1) NewSetCorpusConfigOptions() *SetCorpusConfigOptions {
	return &SetCorpusConfigOptions{}
}

// SetUserName : Allow user to set UserName
func (options *SetCorpusConfigOptions) SetUserName(userName string) *SetCorpusConfigOptions {
	options.UserName = core.StringPtr(userName)
	return options
}

// SetPassword : Allow user to set Password
func (options *SetCorpusConfigOptions) SetPassword(password string) *SetCorpusConfigOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetCorpusURI : Allow user to set CorpusURI
func (options *SetCorpusConfigOptions) SetCorpusURI(corpusURI string) *SetCorpusConfigOptions {
	options.CorpusURI = core.StringPtr(corpusURI)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCorpusConfigOptions) SetHeaders(param map[string]string) *SetCorpusConfigOptions {
	options.Headers = param
	return options
}

// SetCorpusSchemaOptions : The SetCorpusSchema options.
type SetCorpusSchemaOptions struct {
	// Input and Output field names.
	EnrichmentTargets []interface{} `json:"enrichmentTargets,omitempty"`

	// Metadata field names.
	MetadataFields []interface{} `json:"metadataFields,omitempty"`

	// Corpus name.
	CorpusName *string `json:"corpusName,omitempty"`

	// Reference indices.
	References map[string]interface{} `json:"references,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCorpusSchemaOptions : Instantiate SetCorpusSchemaOptions
func (*InsightsForMedicalLiteratureServiceV1) NewSetCorpusSchemaOptions() *SetCorpusSchemaOptions {
	return &SetCorpusSchemaOptions{}
}

// SetEnrichmentTargets : Allow user to set EnrichmentTargets
func (options *SetCorpusSchemaOptions) SetEnrichmentTargets(enrichmentTargets []interface{}) *SetCorpusSchemaOptions {
	options.EnrichmentTargets = enrichmentTargets
	return options
}

// SetMetadataFields : Allow user to set MetadataFields
func (options *SetCorpusSchemaOptions) SetMetadataFields(metadataFields []interface{}) *SetCorpusSchemaOptions {
	options.MetadataFields = metadataFields
	return options
}

// SetCorpusName : Allow user to set CorpusName
func (options *SetCorpusSchemaOptions) SetCorpusName(corpusName string) *SetCorpusSchemaOptions {
	options.CorpusName = core.StringPtr(corpusName)
	return options
}

// SetReferences : Allow user to set References
func (options *SetCorpusSchemaOptions) SetReferences(references map[string]interface{}) *SetCorpusSchemaOptions {
	options.References = references
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCorpusSchemaOptions) SetHeaders(param map[string]string) *SetCorpusSchemaOptions {
	options.Headers = param
	return options
}

// StringBuilder : StringBuilder struct
type StringBuilder struct {
}


// UnmarshalStringBuilder unmarshals an instance of StringBuilder from the specified map of raw messages.
func UnmarshalStringBuilder(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StringBuilder)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TypeaheadOptions : The Typeahead options.
type TypeaheadOptions struct {
	// Comma-separated corpora names.
	Corpus *string `json:"corpus" validate:"required"`

	// Query string.
	Query *string `json:"query" validate:"required"`

	// Include suggestions belonging to the selected ontology(ies).
	Ontologies []string `json:"ontologies,omitempty"`

	// Include or exclude suggestions belonging to one of these types.  Types can be found using /v1/corpora/{corpus}/types
	// method.  Defaults to all.
	Types []string `json:"types,omitempty"`

	// Select concepts belonging to disorders, drugs or genes.
	Category *string `json:"category,omitempty"`

	// Verbose output.  Include hit counts and relationship counts for each concept.
	Verbose *bool `json:"verbose,omitempty"`

	// Maximum number of suggestions to return.
	Limit *int64 `json:"_limit,omitempty"`

	// Maximum hit (document) count for suggested concepts. Default is 500000.  High hit count concepts tend to be very
	// broad (e.g, Disease) and result in longer search times.
	MaxHitCount *int64 `json:"max_hit_count,omitempty"`

	// Remove duplicate concepts.
	NoDuplicates *bool `json:"no_duplicates,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the TypeaheadOptions.Ontologies property.
const (
	TypeaheadOptions_Ontologies_Concepts = "concepts"
	TypeaheadOptions_Ontologies_Mesh = "mesh"
)

// Constants associated with the TypeaheadOptions.Category property.
// Select concepts belonging to disorders, drugs or genes.
const (
	TypeaheadOptions_Category_Disorders = "disorders"
	TypeaheadOptions_Category_Drugs = "drugs"
	TypeaheadOptions_Category_Genes = "genes"
)

// NewTypeaheadOptions : Instantiate TypeaheadOptions
func (*InsightsForMedicalLiteratureServiceV1) NewTypeaheadOptions(corpus string, query string) *TypeaheadOptions {
	return &TypeaheadOptions{
		Corpus: core.StringPtr(corpus),
		Query: core.StringPtr(query),
	}
}

// SetCorpus : Allow user to set Corpus
func (options *TypeaheadOptions) SetCorpus(corpus string) *TypeaheadOptions {
	options.Corpus = core.StringPtr(corpus)
	return options
}

// SetQuery : Allow user to set Query
func (options *TypeaheadOptions) SetQuery(query string) *TypeaheadOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetOntologies : Allow user to set Ontologies
func (options *TypeaheadOptions) SetOntologies(ontologies []string) *TypeaheadOptions {
	options.Ontologies = ontologies
	return options
}

// SetTypes : Allow user to set Types
func (options *TypeaheadOptions) SetTypes(types []string) *TypeaheadOptions {
	options.Types = types
	return options
}

// SetCategory : Allow user to set Category
func (options *TypeaheadOptions) SetCategory(category string) *TypeaheadOptions {
	options.Category = core.StringPtr(category)
	return options
}

// SetVerbose : Allow user to set Verbose
func (options *TypeaheadOptions) SetVerbose(verbose bool) *TypeaheadOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetLimit : Allow user to set Limit
func (options *TypeaheadOptions) SetLimit(limit int64) *TypeaheadOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetMaxHitCount : Allow user to set MaxHitCount
func (options *TypeaheadOptions) SetMaxHitCount(maxHitCount int64) *TypeaheadOptions {
	options.MaxHitCount = core.Int64Ptr(maxHitCount)
	return options
}

// SetNoDuplicates : Allow user to set NoDuplicates
func (options *TypeaheadOptions) SetNoDuplicates(noDuplicates bool) *TypeaheadOptions {
	options.NoDuplicates = core.BoolPtr(noDuplicates)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *TypeaheadOptions) SetHeaders(param map[string]string) *TypeaheadOptions {
	options.Headers = param
	return options
}

// AggregationModel : Model for field aggregations.
type AggregationModel struct {
	// Name of the aggregation.
	Name *string `json:"name,omitempty"`

	// Corpus frequency of the aggregation.
	DocumentCount *int64 `json:"documentCount,omitempty"`
}


// UnmarshalAggregationModel unmarshals an instance of AggregationModel from the specified map of raw messages.
func UnmarshalAggregationModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AggregationModel)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "documentCount", &obj.DocumentCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnnotationModel : Model for congntive asset annotations.
type AnnotationModel struct {
	// Unique identifer of annotation.
	UniqueID *int64 `json:"uniqueId,omitempty"`

	// List of identifiers associated with annotation.
	StickyIds []int64 `json:"stickyIds,omitempty"`

	// Source ontology of annotation.
	Ontology *string `json:"ontology,omitempty"`

	// Document section for annotation.
	Section *string `json:"section,omitempty"`

	// Ontology provide normalized name of annotation.
	PreferredName *string `json:"preferredName,omitempty"`

	// Ontology provided unique identifier of annotation.
	Cui *string `json:"cui,omitempty"`

	// Attribute identifier of annotation.
	AttributeID *string `json:"attributeId,omitempty"`

	// Qualifier for attribute annotation.
	Qualifiers []string `json:"qualifiers,omitempty"`

	// Ontology provided semantic type of annotation.
	Type *string `json:"type,omitempty"`

	// Whether the annotation is a negated span.
	Negated *bool `json:"negated,omitempty"`

	// Whether the annotation is a hypothetical span.
	Hypothetical *bool `json:"hypothetical,omitempty"`

	// Unit of measure for attribute value annotation.
	Unit *string `json:"unit,omitempty"`

	// Minumum value for attribute value annotation.
	MinValue *string `json:"minValue,omitempty"`

	// Maximum value for attribute value annotation.
	MaxValue *string `json:"maxValue,omitempty"`

	// Mathematical operator for attribute value annotation.
	Operator *string `json:"operator,omitempty"`

	// Ontology type of source relation annotation.
	NluSourceType *string `json:"nluSourceType,omitempty"`

	// Relation name for annotation.
	NluRelation *string `json:"nluRelation,omitempty"`

	// Ontology type of target relation annotation.
	NluTargetType *string `json:"nluTargetType,omitempty"`

	NluEntityIndex *string `json:"nluEntityIndex,omitempty"`

	NluMentionIndex *string `json:"nluMentionIndex,omitempty"`

	NluRelationID *string `json:"nluRelationId,omitempty"`

	NluSide *string `json:"nluSide,omitempty"`

	// Starting offset of annotation.
	Begin *int64 `json:"begin,omitempty"`

	// Ending offset of annotation.
	End *int64 `json:"end,omitempty"`

	// Relevancy score of annotation.
	Score *float32 `json:"score,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`

	Features map[string]string `json:"features,omitempty"`

	// Number of times artifact is mentioned in the corpus.
	Hits *int64 `json:"hits,omitempty"`
}


// UnmarshalAnnotationModel unmarshals an instance of AnnotationModel from the specified map of raw messages.
func UnmarshalAnnotationModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnnotationModel)
	err = core.UnmarshalPrimitive(m, "uniqueId", &obj.UniqueID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stickyIds", &obj.StickyIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ontology", &obj.Ontology)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "section", &obj.Section)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attributeId", &obj.AttributeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "qualifiers", &obj.Qualifiers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minValue", &obj.MinValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxValue", &obj.MaxValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluSourceType", &obj.NluSourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluRelation", &obj.NluRelation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluTargetType", &obj.NluTargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluEntityIndex", &obj.NluEntityIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluMentionIndex", &obj.NluMentionIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluRelationId", &obj.NluRelationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluSide", &obj.NluSide)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hits", &obj.Hits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ArtifactModel : Model for ontology artifact.
type ArtifactModel struct {
	// Ontology provided unique identifier for artifact.
	Cui *string `json:"cui,omitempty"`

	// Source ontology for artifact.
	Ontology *string `json:"ontology,omitempty"`

	// Ontology provided normalized name for artifact.
	PreferredName *string `json:"preferredName,omitempty"`

	// Ontology provided alternative name for artifact.
	AlternativeName *string `json:"alternativeName,omitempty"`

	// Ontology semantic type for artifact.
	SemanticType *string `json:"semanticType,omitempty"`

	// Search weight assigned to artifact.
	Rank *int64 `json:"rank,omitempty"`

	// Number of corpus documents artifact was found in.
	HitCount *int64 `json:"hitCount,omitempty"`

	// Relevance score for artifact.
	Score *float32 `json:"score,omitempty"`

	// List of artifact synonyms.
	SurfaceForms []string `json:"surfaceForms,omitempty"`
}


// UnmarshalArtifactModel unmarshals an instance of ArtifactModel from the specified map of raw messages.
func UnmarshalArtifactModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ArtifactModel)
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ontology", &obj.Ontology)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternativeName", &obj.AlternativeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semanticType", &obj.SemanticType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rank", &obj.Rank)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hitCount", &obj.HitCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "surfaceForms", &obj.SurfaceForms)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Attribute : Object representing an attribute artifact.
type Attribute struct {
	// Unique identifier for attribute artifact.
	AttributeID *string `json:"attributeId,omitempty"`

	// Display name for attribute artifact.
	DisplayName *string `json:"displayName,omitempty"`

	// Corpus frequency for attribute artifact.
	Count *int64 `json:"count,omitempty"`
}


// UnmarshalAttribute unmarshals an instance of Attribute from the specified map of raw messages.
func UnmarshalAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attribute)
	err = core.UnmarshalPrimitive(m, "attributeId", &obj.AttributeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "displayName", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Backend : Object representing repository response.
type Backend struct {
	// Repository messages.
	Messages []Message `json:"messages,omitempty"`
}


// UnmarshalBackend unmarshals an instance of Backend from the specified map of raw messages.
func UnmarshalBackend(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backend)
	err = core.UnmarshalModel(m, "messages", &obj.Messages, UnmarshalMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BooleanOperands : Object representingn boolean operands search criteria.
type BooleanOperands struct {
	// Boolean search condition.
	BoolExpression *string `json:"boolExpression,omitempty"`

	// Ontology artifacts supporing boolean search condition.
	BoolOperands []BoolOperand `json:"boolOperands,omitempty"`
}


// UnmarshalBooleanOperands unmarshals an instance of BooleanOperands from the specified map of raw messages.
func UnmarshalBooleanOperands(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BooleanOperands)
	err = core.UnmarshalPrimitive(m, "boolExpression", &obj.BoolExpression)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "boolOperands", &obj.BoolOperands, UnmarshalBoolOperand)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoriesModel : Model representing ontology categories.
type CategoriesModel struct {
	// License for corpus.
	ModelLicense *string `json:"modelLicense,omitempty"`

	HighlightedTitle *string`json:"highlightedTitle,omitempty"`

	HighlightedAbstract *string `json:"highlightedAbstract,omitempty"`

	HighlightedBody *string `json:"highlightedBody,omitempty"`

	// Document sections with annotation tags.
	HighlightedSections map[string]string `json:"highlightedSections,omitempty"`

	// Document passages with annotation tags.
	Passages map[string]MatchingPassages `json:"passages,omitempty"`

	// List of document annotations.
	Annotations map[string]AnnotationModel `json:"annotations,omitempty"`
}


// UnmarshalCategoriesModel unmarshals an instance of CategoriesModel from the specified map of raw messages.
func UnmarshalCategoriesModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoriesModel)
	err = core.UnmarshalPrimitive(m, "modelLicense", &obj.ModelLicense)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedTitle", &obj.HighlightedTitle)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedAbstract", &obj.HighlightedAbstract)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedBody", &obj.HighlightedBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedSections", &obj.HighlightedSections)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalMatchingPassages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotations", &obj.Annotations, UnmarshalAnnotationModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CommonDataModel : Model representing common data across annotations.
type CommonDataModel struct {
	// Object of all ontology artifacts found in the document.
	Unstructured []UnstructuredModel `json:"unstructured,omitempty"`
}


// UnmarshalCommonDataModel unmarshals an instance of CommonDataModel from the specified map of raw messages.
func UnmarshalCommonDataModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CommonDataModel)
	err = core.UnmarshalModel(m, "unstructured", &obj.Unstructured, UnmarshalUnstructuredModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Concept : Object reprensting an ontology artifact.
type Concept struct {
	// Ontology for artifact in search results.
	Ontology *string `json:"ontology,omitempty"`

	// Unique identifier for ontolgoy artifact in search results.
	Cui *string `json:"cui,omitempty"`

	// Ontology defined semantic type for artifact in search results.
	PreferredName *string `json:"preferredName,omitempty"`

	// Ontology defined normalized name for artifact in search results.
	AlternativeName *string `json:"alternativeName,omitempty"`

	// Ontology defined alternative name for artifact in search results.
	SemanticType *string `json:"semanticType,omitempty"`

	// Corpus frequency of artifact.
	Count *int64 `json:"count,omitempty"`

	// Corpus frequency of artifact.
	HitCount *int64 `json:"hitCount,omitempty"`

	// Relevancy score of artifact in search results.
	Score *float32 `json:"score,omitempty"`

	// Corpus frequency count.
	Parents *Count `json:"parents,omitempty"`

	// Corpus frequency count.
	Children *Count `json:"children,omitempty"`

	// Corpus frequency count.
	Siblings *Count `json:"siblings,omitempty"`

	// Corpus frequency count.
	Related *Count `json:"related,omitempty"`

	// Document identifiers for artifacts in search results.
	DocumentIds []string `json:"documentIds,omitempty"`

	// attribute data type for artifact in search results.
	DataType *string `json:"dataType,omitempty"`

	// Attribute value unit for artifact.
	Unit *string `json:"unit,omitempty"`

	// Attribute value operator for artifact.
	Operator *string `json:"operator,omitempty"`

	// Minimum value for attribute artifact.
	MinValue *string `json:"minValue,omitempty"`

	// Maximum value for attribute artifact.
	MaxValue *string `json:"maxValue,omitempty"`

	// Source vocabulary of arttifact.
	Vocab *string `json:"vocab,omitempty"`

	// Artifact properties.
	Properties []string `json:"properties,omitempty"`
}


// UnmarshalConcept unmarshals an instance of Concept from the specified map of raw messages.
func UnmarshalConcept(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Concept)
	err = core.UnmarshalPrimitive(m, "ontology", &obj.Ontology)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternativeName", &obj.AlternativeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semanticType", &obj.SemanticType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hitCount", &obj.HitCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parents", &obj.Parents, UnmarshalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "children", &obj.Children, UnmarshalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "siblings", &obj.Siblings, UnmarshalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "related", &obj.Related, UnmarshalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "documentIds", &obj.DocumentIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dataType", &obj.DataType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minValue", &obj.MinValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxValue", &obj.MaxValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vocab", &obj.Vocab)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "properties", &obj.Properties)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptInfoModel : Model representing ontology annotations.
type ConceptInfoModel struct {
	// Ontology provided unique identifier for artifact.
	Cui *string `json:"cui,omitempty"`

	// Source onology of artifact.
	Ontology *string `json:"ontology,omitempty"`

	// Ontology defined normalized name for artifact.
	PreferredName *string `json:"preferredName,omitempty"`

	// Ontology defined semanic types for artifact.
	SemanticTypes []string `json:"semanticTypes,omitempty"`

	// Ontology defined synonyms for artifact.
	SurfaceForms []string `json:"surfaceForms,omitempty"`

	// Ontology provided definition for artifact.
	Definition *string `json:"definition,omitempty"`

	// Whether the artifact has parent artifacts in the ontology.
	HasParents *bool `json:"hasParents,omitempty"`

	// Whether the artifact has child artifacts in the ontology.
	HasChildren *bool `json:"hasChildren,omitempty"`

	// Whether the artifact has sibling artifacts in the ontology.
	HasSiblings *bool `json:"hasSiblings,omitempty"`
}


// UnmarshalConceptInfoModel unmarshals an instance of ConceptInfoModel from the specified map of raw messages.
func UnmarshalConceptInfoModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptInfoModel)
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ontology", &obj.Ontology)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semanticTypes", &obj.SemanticTypes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "surfaceForms", &obj.SurfaceForms)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "definition", &obj.Definition)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hasParents", &obj.HasParents)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hasChildren", &obj.HasChildren)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hasSiblings", &obj.HasSiblings)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptListModel : List of ontolgoy artifacts.
type ConceptListModel struct {
	// List of ontology artifacts.
	Concepts []ArtifactModel `json:"concepts,omitempty"`
}


// UnmarshalConceptListModel unmarshals an instance of ConceptListModel from the specified map of raw messages.
func UnmarshalConceptListModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptListModel)
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalArtifactModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptModel : Model representing an ontology annotation.
type ConceptModel struct {
	// Service generated unique identifier of ontology artifact.
	UniqueID *int64 `json:"uniqueId,omitempty"`

	// Identifiers associated with artifact unique identifier.
	StickyIds []int64 `json:"stickyIds,omitempty"`

	// Document section where artifact was found.
	Section *string `json:"section,omitempty"`

	// Ontology semantic type for artifact (if applicable).
	Type *string `json:"type,omitempty"`

	// Staring offset of artifact in document section.
	Begin *int64 `json:"begin,omitempty"`

	// Ending offset of artifact in document section.
	End *int64 `json:"end,omitempty"`

	// Actual document section text artifact represents.
	CoveredText *string `json:"coveredText,omitempty"`

	// Ontology defined unique identifier of artifact.
	Cui *string `json:"cui,omitempty"`

	// Ontology defined normalized name of artifact.
	PreferredName *string `json:"preferredName,omitempty"`

	// Ontology providing the artifact.
	Source *string `json:"source,omitempty"`

	// Whether span represented by artifact is negated.
	Negated *bool `json:"negated,omitempty"`

	// Whether span represented by artifact is hypothetical.
	Hypothetical *bool `json:"hypothetical,omitempty"`

	// Time based offset of artifact in a video transcript (if applicable).
	Timestamp *int64 `json:"timestamp,omitempty"`

	// Identifier of attribute where artifact is defined (if applicable).
	AttributeID *string `json:"attributeId,omitempty"`

	// List of qualifers defined for an attribute artifact.
	Qualifiers []string `json:"qualifiers,omitempty"`

	// Unit of measure for attribute defined artifact (if applicable).
	Unit *string `json:"unit,omitempty"`

	// Starting range value for attribute artifact (if applicable).
	MinValue *string `json:"minValue,omitempty"`

	// Ending range value for attribute artifact (if applicable).
	MaxValue *string `json:"maxValue,omitempty"`

	// Mathmatical operator for attribute artifact (if applicable).
	Operator *string `json:"operator,omitempty"`

	// List of additional artifact features.
	Features map[string]string `json:"features,omitempty"`

	// Model coreference chain to which artifact belongs.
	NluEntityIndex *string `json:"nluEntityIndex,omitempty"`

	// Artifact position in Model coreference chain.
	NluMentionIndex *string `json:"nluMentionIndex,omitempty"`

	// Relation unique identifier artifact is associated with.
	NluRelationID *string `json:"nluRelationId,omitempty"`

	// Whether artifact is a source or target of a relationship.
	NluSide *string `json:"nluSide,omitempty"`

	// Model type for artifact when the source of a relationship.
	NluSourceType *string `json:"nluSourceType,omitempty"`

	// Name of the realtion an artifact is associated with.
	NluRelation *string `json:"nluRelation,omitempty"`

	// Model type for artifact when the target of a relationship.
	NluTargetType *string `json:"nluTargetType,omitempty"`

	// Number of times artifact is mentioned in the corpus.
	Hits *int64 `json:"hits,omitempty"`
}


// UnmarshalConceptModel unmarshals an instance of ConceptModel from the specified map of raw messages.
func UnmarshalConceptModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptModel)
	err = core.UnmarshalPrimitive(m, "uniqueId", &obj.UniqueID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stickyIds", &obj.StickyIds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "section", &obj.Section)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "coveredText", &obj.CoveredText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attributeId", &obj.AttributeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "qualifiers", &obj.Qualifiers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "minValue", &obj.MinValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "maxValue", &obj.MaxValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluEntityIndex", &obj.NluEntityIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluMentionIndex", &obj.NluMentionIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluRelationId", &obj.NluRelationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluSide", &obj.NluSide)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluSourceType", &obj.NluSourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluRelation", &obj.NluRelation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nluTargetType", &obj.NluTargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hits", &obj.Hits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CorporaConfig : Model respresenting configured corpora.
type CorporaConfig struct {
	// List of corpora found in the instance.
	Corpora []CorpusModel `json:"corpora,omitempty"`
}


// UnmarshalCorporaConfig unmarshals an instance of CorporaConfig from the specified map of raw messages.
func UnmarshalCorporaConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CorporaConfig)
	err = core.UnmarshalModel(m, "corpora", &obj.Corpora, UnmarshalCorpusModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CorpusInfoModel : Object representing a configured corpus document count.
type CorpusInfoModel struct {
	// Number of documents.
	DocumentCount *int64 `json:"documentCount,omitempty"`

	// Ontologies found in the corpus.
	Providers []CorpusProvider `json:"providers,omitempty"`
}


// UnmarshalCorpusInfoModel unmarshals an instance of CorpusInfoModel from the specified map of raw messages.
func UnmarshalCorpusInfoModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CorpusInfoModel)
	err = core.UnmarshalPrimitive(m, "documentCount", &obj.DocumentCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "providers", &obj.Providers, UnmarshalCorpusProvider)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CorpusModel : Object representing a configured corpus.
type CorpusModel struct {
	// Name of the corpus.
	CorpusName *string `json:"corpusName,omitempty"`

	// Ontologies found in the corpus.
	Ontologies []string `json:"ontologies,omitempty"`

	// Descriptive name of the corpus.
	DescriptiveName *string `json:"descriptiveName,omitempty"`

	// BVT status of the corpus.
	Bvt *bool `json:"bvt,omitempty"`

	// Repository location of the corpus.
	ElasticsearchIndex *string `json:"elasticsearchIndex,omitempty"`
}


// UnmarshalCorpusModel unmarshals an instance of CorpusModel from the specified map of raw messages.
func UnmarshalCorpusModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CorpusModel)
	err = core.UnmarshalPrimitive(m, "corpusName", &obj.CorpusName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ontologies", &obj.Ontologies)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "descriptiveName", &obj.DescriptiveName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bvt", &obj.Bvt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "elasticsearchIndex", &obj.ElasticsearchIndex)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CorpusProvider: Object representing a corpus provider.
type CorpusProvider struct {
	// Name of the corpus.
	DocumentCount *int64 `json:"documentCount,omitempty"`

	// Ontologies found in the corpus.
	Name *string `json:"name,omitempty"`
}

// UnmarshalCorpusProvider unmarshals an instance of CorpusProvider from the specified map of raw messages.
func UnmarshalCorpusProvider(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CorpusProvider)
	err = core.UnmarshalPrimitive(m, "documentCount", &obj.DocumentCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}


// Count : Corpus frequency count.
type Count struct {
	// Number of documents for artifact result.
	Count *int64 `json:"count,omitempty"`

	// Number of documents for artifact result.
	Hits *int64 `json:"hits,omitempty"`
}


// UnmarshalCount unmarshals an instance of Count from the specified map of raw messages.
func UnmarshalCount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Count)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hits", &obj.Hits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DataModel : Model representing ontology artifacts.
type DataModel struct {
	// List of ontolgy artifacts found in the document.
	Concepts []ConceptModel `json:"concepts,omitempty"`

	// List of ontolgy attribute value artifacts found in the document.
	AttributeValues []ConceptModel `json:"attributeValues,omitempty"`

	// List of ontology relations found in the document.
	Relations []RelationModel `json:"relations,omitempty"`

	Mesh []ConceptModel `json:"mesh,omitempty"`

	Text []ConceptModel `json:"text,omitempty"`
}


// UnmarshalDataModel unmarshals an instance of DataModel from the specified map of raw messages.
func UnmarshalDataModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataModel)
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConceptModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributeValues", &obj.AttributeValues, UnmarshalConceptModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalRelationModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "mesh", &obj.Mesh, UnmarshalConceptModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "text", &obj.Text, UnmarshalConceptModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EntryModel : Object representing a passage.
type EntryModel struct {
	// Unique identifier of passage.
	ID *string `json:"id,omitempty"`

	// Whether passage is a negated span.
	Negated *bool `json:"negated,omitempty"`

	// List of sentences within passage.
	Sentences []SentenceModel `json:"sentences,omitempty"`
}


// UnmarshalEntryModel unmarshals an instance of EntryModel from the specified map of raw messages.
func UnmarshalEntryModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EntryModel)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentences", &obj.Sentences, UnmarshalSentenceModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FieldOptions : Supported options for the field.
type FieldOptions struct {
	// List of supported options.
	Supports []string `json:"supports,omitempty"`
}


// UnmarshalFieldOptions unmarshals an instance of FieldOptions from the specified map of raw messages.
func UnmarshalFieldOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FieldOptions)
	err = core.UnmarshalPrimitive(m, "supports", &obj.Supports)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HistogramData : histogram data.
type HistogramData struct {
	// Date associated with result.
	Date *string `json:"date,omitempty"`

	// Number of documents for date range result.
	Hits *int64 `json:"hits,omitempty"`
}


// UnmarshalHistogramData unmarshals an instance of HistogramData from the specified map of raw messages.
func UnmarshalHistogramData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HistogramData)
	err = core.UnmarshalPrimitive(m, "date", &obj.Date)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hits", &obj.Hits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HitCount : Corpus frequency of artifact.
type HitCount struct {
	// Corpus frequency of artifact.
	HitCount *int64 `json:"hitCount,omitempty"`
}


// UnmarshalHitCount unmarshals an instance of HitCount from the specified map of raw messages.
func UnmarshalHitCount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HitCount)
	err = core.UnmarshalPrimitive(m, "hitCount", &obj.HitCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MatchEntry : Object representing a match entry.
type MatchEntry struct {
	// Whether match is a negated span.
	Negated *bool `json:"negated,omitempty"`

	// Relevancy score of the match.
	Score *float32 `json:"score,omitempty"`

	// List of sentences within the matched text.
	Sentences []SentenceModel `json:"sentences,omitempty"`

	// Unique identifier of the match.
	ID *string `json:"id,omitempty"`
}


// UnmarshalMatchEntry unmarshals an instance of MatchEntry from the specified map of raw messages.
func UnmarshalMatchEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MatchEntry)
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentences", &obj.Sentences, UnmarshalSentenceModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MatchingPassages : Model for matching passages of search criteria.
type MatchingPassages struct {
	// identifier.
	Cui *string `json:"cui,omitempty"`

	// Whether passages is negatedc.
	Negated *bool `json:"negated,omitempty"`

	// Match score
	Score *float64 `json:"score,omitempty"`

	// Matching sentences.
	Sentences []SentenceModel `json:"sentences,omitempty"`
}

// UnmarshalMatchingPassages unmarshal an isntance of MatchingPassages from the raw result.
func UnmarshalMatchingPassages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MatchingPassages)
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sentences", &obj.Sentences, UnmarshalSentenceModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MetadataModel : Model for document metadata.
type MetadataModel struct {
	// List of document fields in the corpus.
	Fields map[string]FieldOptions `json:"fields,omitempty"`

	// List of fields that where enriched.
	SectionFieldNames []string `json:"sectionFieldNames,omitempty"`

	// List of fields enriched with attributes.
	AttrSectionFieldNames []string `json:"attrSectionFieldNames,omitempty"`

	// List of fields enriched with attribute qualifiers.
	QualifierSectionFieldNames []string `json:"qualifierSectionFieldNames,omitempty"`

	// List of fields with MeSH annotations.
	MeshSectionFieldNames []string `json:"meshSectionFieldNames,omitempty"`

	FieldIndexMap map[string]string `json:"fieldIndexMap,omitempty"`
}


// UnmarshalMetadataModel unmarshals an instance of MetadataModel from the specified map of raw messages.
func UnmarshalMetadataModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetadataModel)
	err = core.UnmarshalModel(m, "fields", &obj.Fields, UnmarshalFieldOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionFieldNames", &obj.SectionFieldNames)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attrSectionFieldNames", &obj.AttrSectionFieldNames)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "qualifierSectionFieldNames", &obj.QualifierSectionFieldNames)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "meshSectionFieldNames", &obj.MeshSectionFieldNames)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fieldIndexMap", &obj.FieldIndexMap)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Passage : Object representing a document passage.
type Passage struct {
	// Document section for passage.
	DocumentSection *string `json:"documentSection,omitempty"`

	Text *StringBuilder `json:"text,omitempty"`

	// Timestamp of passage in video transcript.
	Timestamp *int64 `json:"timestamp,omitempty"`

	// Preferred name for highlighted text span.
	PreferredName *string `json:"preferredName,omitempty"`
}


// UnmarshalPassage unmarshals an instance of Passage from the specified map of raw messages.
func UnmarshalPassage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Passage)
	err = core.UnmarshalPrimitive(m, "documentSection", &obj.DocumentSection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "text", &obj.Text, UnmarshalStringBuilder)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PassagesModel: Object represnting document passages.
//type PassagesModel struct {
	// Passage entry.
//	Entry Entry `json:"entry,omitempty`
//}

// UnmarshalPassagesModel unmarshals an instance of PassagesModel from the specified map of raw messages.
//func UnmarshalPassagesModel(m map[string]json.RawMessage, result interface{}) (err error) {
//	obj := new(PassagesModel)/
//	err = core.UnmarshalModel(m, "entry", &obj.Text, UnmarshalEntry)
//	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
//	return
//}

// Qualifer : Object representing an artifact qualifier.
type Qualifer struct {
	// Unique identifier for attribute qualifier.
	ID *string `json:"id,omitempty"`

	// Name of attribute qualifier.
	Name *string `json:"name,omitempty"`
}


// UnmarshalQualifer unmarshals an instance of Qualifer from the specified map of raw messages.
func UnmarshalQualifer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Qualifer)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// QueryConcept : Objecy representing a concept query condition.
type QueryConcept struct {
	// BoolOperand portion of the expression for this condition.
	BoolOperand *string `json:"boolOperand,omitempty"`

	// concept identifier.
	Cui *string `json:"cui,omitempty"`

	// List of related concepts to include
	IncludeRelated []string `json:"includeRelated,omitempty"`

	// Whether to target negated instances of the concept
	Negated *bool `json:"negated,omitempty"`

	// Ontology for concept
	Ontology *string `json:"ontology,omitempty"`

	// Rank of concept
	Rank *string `json:"rank,omitempty"`

	// Sections to search for concept
	Sections []string `json:"section,omitempty"`

	// Semantic type for concept
	SemanticType *string `json:"semanticType,omitempty"`

	// Text for text-based search
	Text *string `json:"text,omitempty"`
}

// NewQueryModel instantiates new instance of QueryModel
func (*InsightsForMedicalLiteratureServiceV1) NewQueryConcept() *QueryConcept {
	return &QueryConcept{}
}

// Sets the search expression
func (query *QueryConcept) SetBoolOperand(operand string) *QueryConcept {
	query.BoolOperand = core.StringPtr(operand)
	return query
}

// Sets the concept identifier
func (query *QueryConcept) SetCui(cuii string) *QueryConcept {
	query.Cui = core.StringPtr(cui)
	return query
}

// Sets the included relations
func (query *QueryConcept) SetInludeRelation(relations []string) *QueryConcept {
	query.IncludeRelated = relations
	return query
}

// Sets the negated flag
func (query *QueryConcept) SetNegated(negated bool) *QueryConcept {
	query.Negated = core.BoolPtr(negated)
	return query
}

// Sets the source ontology
func (query *QueryConcept) SetOntology(ontology string) *QueryConcept {
	query.Ontology = core.StringPtr(ontology)
	return query
}

// Sets the section list
func (query *QueryConcept) SetSections(sections []string) *QueryConcept {
	query.Sections = sections
	return query
}

// Sets the semantic type
func (query *QueryConcept) SetSemanticType(semType string) *QueryConcept {
	query.SemanticType = core.StringPtr(semType)
	return query
}

// Sets the text search
func (query *QueryConcept) SetText(text string) *QueryConcept {
	query.Text = core.StringPtr(text)
	return query
}



// QueryModel : Object representing a query condition.
type QueryModel struct {
	// BoolExpression representing the entrire query.
	BoolExpression *string `json:"boolExpression,omitempty"`

	// BoolRegexp a boolean regular expression.
	BoolRegexp map[string]string `json:"boolRegexp,omitempty"`

	// BoolTerm qeury condition.
	BoolTerm map[string]string `json:"boolTerm,omitempty"`

	// Concepts list of ontology based query conditions.
	Concepts []QueryConcept `json:"concepts,omitempty"`

	// Scrolling cursor identifier.
	CursorId *string `json:"cursordId,omitempty"`

	// Date range query condition.
	DateRange map[string]Range `json:"dateRange,omitempty"`

	// Regular expression for value condition.
	Regexp []map[string]string `json:"regexp,omitempty"`

	// Boost factofr for title search emphasis.
	Tille *TitleBoost `json;"title,omitempty"`
}

// NewQueryModel instantiates new instance of QueryModel
func (*InsightsForMedicalLiteratureServiceV1) NewQueryModel() *QueryModel {
	return &QueryModel{}
}

// Sets the search expression
func (query *QueryModel) SetBoolExpression(expression string) *QueryModel {
	query.BoolExpression = core.StringPtr(expression)
	return query
}

// Sets the boolean regular expression condition.
func (query *QueryModel) SetBoolRegexp(regexp []map[string]string) *QueryModel {
	query.BoolRegexp = regexp
	return query
}

// Sets the booleam term condition.
func (query *QueryModel) SetBoolTerm(term map[string]string) *QueryModel {
	query.BoolTerm = term
	return query
}

// Sets the scrolling cursor identifier.
func (query *QueryModel) SetCursorId(cursor string) *QueryModel {
	query.CursorId = core.StringPtr(cursor)
	return query
}

// Sets the concept conditions.
func (query *QueryModel) SetConcepts(concepts []Concepts) *QueryModel {
	query.Concepts = concepts
	return query
}

// Sets the date range
func (query *QueryModel) SetDateRange(ranges []map[string]Range) *QueryModel {
	query.DateRange = ranges
	return query
}

// Sets the booleam term condition.
func (query *QueryModel) SetRegexp(exp string) *QueryModel {
	query.Regexp = core.StringPtr(exp)
	return query
}

// Sets the booleam term condition.
func (query *QueryModel) SetTitleBoost(booster *Title) *QueryModel {
	query.Title = booster
	return query
}

// UnmarshalQueryModel unmarshals a QueryModel instance from the raw mwessage.
func UnmarshalQueryModel(m map[string]json.RawMessage, result interface{}} (err error) {
	obj := new(QueryModel)
	err = core.UnmarshalPrimitive(m, "boolExpression", &obj.BoolExpression)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "boolRegexp", &obj.BoolRegexp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "boolTerm", &obj.BoolTerm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalQueryConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cursorId", &obj.CursorId)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dateRange", &obj.DateRange, UnmarshalRange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "regexp", &obj.Regexp)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "title", &obj.Title, UnmarshalTitle)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Range : object representing a date range.
type Range struct {
	// Begging of range.
	Begin *string `json:"begin,omitempty"`

	// Ending of range.
	End *string `json:"end,omitempty"`
}

// NewRange instantiates new instance of QueryModel.
func (*InsightsForMedicalLiteratureServiceV1) NewRange() *Range {
	return &Range{}
}
.
// Sets the range beginning value
func (options *Range) SetBegin(begin string) *Range {
	options.Begin = core.StringPtr(begin)
	return options
}

// Sets the range ending value.
func (options *Range) SetBEnd(end string) *Range {
	options.End = core.StringPtr(end)
	return options
}

// UnmarshalRange unmarshals a Range instance from the raw mwessage.
func UnmarshalRange(m map[string]json.RawMessage, result interface{}} (err error) {
	obj := new(Range)
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RangeModel : Object representing an attribute value range.
type RangeModel struct {
	// Operator assigned to attribute value.
	Operator *string `json:"operator,omitempty"`

	// Minimum value assigned to attribute value.
	Min *string `json:"min,omitempty"`

	// Maximum value assigned to attribute value.
	Max *string `json:"max,omitempty"`

	// Corpus frequency for attribute value.
	Count *int64 `json:"count,omitempty"`
}


// UnmarshalRangeModel unmarshals an instance of RangeModel from the specified map of raw messages.
func UnmarshalRangeModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RangeModel)
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min", &obj.Min)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max", &obj.Max)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RankedDocument : Object representing a document search result.
type RankedDocument struct {
	// Document identifier.
	DocumentID *string `json:"documentId,omitempty"`

	// Document title.
	Title *string `json:"title,omitempty"`

	// Additional document fields.
	Metadata map[string][]string `json:"metadata,omitempty"`

	// Document section.
	SectionName *string `json:"sectionName,omitempty"`

	// Document section identifier.
	SectionID *string `json:"sectionId,omitempty"`

	// Document corpus.
	Corpus *string `json:"corpus,omitempty"`

	Links *RankedDocLinks `json:"links,omitempty"`

	Passages *PassagesModel `json:"passages,omitempty"`

	// Document annotations.
	Annotations map[string]AnnotationModel `json:"annotations,omitempty"`
}


// UnmarshalRankedDocument unmarshals an instance of RankedDocument from the specified map of raw messages.
func UnmarshalRankedDocument(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RankedDocument)
	err = core.UnmarshalPrimitive(m, "documentId", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionName", &obj.SectionName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionId", &obj.SectionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "corpus", &obj.Corpus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "links", &obj.Links, UnmarshalRankedDocLinks)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalPassagesModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotations", &obj.Annotations, UnmarshalAnnotationModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelatedConceptModel : Model for concept ontology relation.
type RelatedConceptModel struct {
	// Ontology provided unique identifier for artifact.
	Cui *string `json:"cui,omitempty"`

	// Source ontology for artifact.
	Ontology *string `json:"ontology,omitempty"`

	// Ontology provided normalized name for artifact.
	PreferredName *string `json:"preferredName,omitempty"`

	// Ontology provided alternative name for artifact.
	AlternativeName *string `json:"alternativeName,omitempty"`

	// Ontology semantic type for artifact.
	SemanticType *string `json:"semanticType,omitempty"`

	// Search weight assigned to artifact.
	Rank *int64 `json:"rank,omitempty"`

	// Number of corpus documents artifact was found in.
	HitCount *int64 `json:"hitCount,omitempty"`

	// Relevance score for artifact.
	Score *float32 `json:"score,omitempty"`

	// List of artifact synonyms.
	SurfaceForms []string `json:"surfaceForms,omitempty"`

	// List of artifacts for the relation.
	NextConcepts []RelatedConceptModel `json:"nextConcepts,omitempty"`
}


// UnmarshalRelatedConceptModel unmarshals an instance of RelatedConceptModel from the specified map of raw messages.
func UnmarshalRelatedConceptModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelatedConceptModel)
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ontology", &obj.Ontology)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "alternativeName", &obj.AlternativeName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semanticType", &obj.SemanticType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rank", &obj.Rank)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hitCount", &obj.HitCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "surfaceForms", &obj.SurfaceForms)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nextConcepts", &obj.NextConcepts, UnmarshalRelatedConceptModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelatedConceptsModel : Model for concept ontology relations.
type RelatedConceptsModel struct {
	// List of artifacts for the relation.
	Concepts []RelatedConceptModel `json:"concepts,omitempty"`
}


// UnmarshalRelatedConceptsModel unmarshals an instance of RelatedConceptsModel from the specified map of raw messages.
func UnmarshalRelatedConceptsModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelatedConceptsModel)
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalRelatedConceptModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RelationModel : Object representing an ontology relation.
type RelationModel struct {
	// Relation unique identifier artifact is associated with.
	RelationID *string `json:"relationId,omitempty"`

	// Name of the realtion an artifact is associated with.
	Relation *string `json:"relation,omitempty"`

	// Objeft representing a document text span.
	Source *TextSpan `json:"source,omitempty"`

	// Objeft representing a document text span.
	Target *TextSpan `json:"target,omitempty"`
}


// UnmarshalRelationModel unmarshals an instance of RelationModel from the specified map of raw messages.
func UnmarshalRelationModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RelationModel)
	err = core.UnmarshalPrimitive(m, "relationId", &obj.RelationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relation", &obj.Relation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "source", &obj.Source, UnmarshalTextSpan)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTextSpan)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchMatchesModel : Object representing a corpus search match.
type SearchMatchesModel struct {
	// Unique identifier for matched document in corpus.
	ExternalID *string `json:"externalId,omitempty"`

	// Unique identifier for matched document in corpus.
	DocumentID *string `json:"documentId,omitempty"`

	// Unique identifier for matched document parent in corpus.
	ParentDocumentID *string `json:"parentDocumentId,omitempty"`

	// Publication name for matched document in corpus.
	PublicationName *string `json:"publicationName,omitempty"`

	// Publication date for matched document in corpus.
	PublicationDate *string `json:"publicationDate,omitempty"`

	// Publication URL for matched document in corpus.
	PublicationURL *string `json:"publicationURL,omitempty"`

	// Authors of matched document in corpus.
	Authors []string `json:"authors,omitempty"`

	// Title of matched document in corpus.
	Title *string `json:"title,omitempty"`

	// Usage license for matched document in corpus.
	MedlineLicense *string `json:"medlineLicense,omitempty"`

	// Pubmed link for matched document in corpus.
	HrefPubMed *string `json:"hrefPubMed,omitempty"`

	HrefPmc *string `json:"hrefPmc,omitempty"`

	HrefDoi *string `json:"hrefDoi,omitempty"`

	// Link to PDF for matched document in corpus.
	PdfURL *string `json:"pdfUrl,omitempty"`

	// Link to sourc origin for matched document in corpus.
	ReferenceURL *string `json:"referenceUrl,omitempty"`

	HighlightedTitle *string `json:"highlightedTitle,omitempty"`

	HighlightedAbstract *string `json:"highlightedAbstract,omitempty"`

	HighlightedBody *string `json:"highlightedBody,omitempty"`

	// Matched document sections with annotation tags.
	HighlightedSections map[string]string `json:"highlightedSections,omitempty"`

	// Matched document passages with annotation tags.
	Passages map[string]MatchingPassages `json:"passages,omitempty"`

	// Matched document annotations.
	Annotations map[string]AnnotationModel `json:"annotations,omitempty"`
}


// UnmarshalSearchMatchesModel unmarshals an instance of SearchMatchesModel from the specified map of raw messages.
func UnmarshalSearchMatchesModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SearchMatchesModel)
	err = core.UnmarshalPrimitive(m, "externalId", &obj.ExternalID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "documentId", &obj.DocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parentDocumentId", &obj.ParentDocumentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publicationName", &obj.PublicationName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publicationDate", &obj.PublicationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publicationURL", &obj.PublicationURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "authors", &obj.Authors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "medlineLicense", &obj.MedlineLicense)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hrefPubMed", &obj.HrefPubMed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hrefPmc", &obj.HrefPmc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hrefDoi", &obj.HrefDoi)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pdfUrl", &obj.PdfURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "referenceUrl", &obj.ReferenceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedTitle", &obj.HighlightedTitle)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedAbstract", &obj.HighlightedAbstract)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedBody", &obj.HighlightedBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "highlightedSections", &obj.HighlightedSections)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passages", &obj.Passages, UnmarshalMatchingPassages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotations", &obj.Annotations, UnmarshalAnnotationModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchModel : Model for search criteria.
type SearchModel struct {
	// Link.
	Href *string `json:"href,omitempty"`

	// Page number.
	PageNumber *int64 `json:"pageNumber,omitempty"`

	// Search result limit.
	GetLimit *int64 `json:"get_limit,omitempty"`

	// Total number of search matches in the corpus.
	TotalDocumentCount *int64 `json:"totalDocumentCount,omitempty"`

	// Ontology artifact results from search.
	Concepts []Concept `json:"concepts,omitempty"`

	// Ontology semantic types.
	Types []string `json:"types,omitempty"`

	// Attribute artifact results from search.
	Attributes []Attribute `json:"attributes,omitempty"`

	// Attribute artifact value results from search.
	Values []Concept `json:"values,omitempty"`

	// Attribute value range results from search.
	Ranges map[string]RangeModel `json:"ranges,omitempty"`

	// Type-ahead suggestion results in search.
	Typeahead []Concept `json:"typeahead,omitempty"`

	// Aggregate result targets in search.
	Aggregations map[string][]AggregationModel `json:"aggregations,omitempty"`

	// Date range of results from search.
	DateHistograms map[string][]HistogramData `json:"dateHistograms,omitempty"`

	// Attribute qualifier results from search.
	Qualifiers []Qualifer `json:"qualifiers,omitempty"`

	// Object representing repository response.
	Backend *Backend `json:"backend,omitempty"`

	// Search expression that includes all levels of criteria expression.
	ExpandedQuery interface{} `json:"expandedQuery,omitempty"`

	// Object representingn boolean operands search criteria.
	ParsedBoolExpression *BooleanOperands `json:"parsedBoolExpression,omitempty"`

	// Whether ontolgoy artifacts were provided in search conditions.
	ConceptsExist map[string]int64 `json:"conceptsExist,omitempty"`

	CursorID *string `json:"cursorId,omitempty"`

	Vocabs []string `json:"vocabs,omitempty"`

	// Annotations returned for the document.
	Annotations map[string]CommonDataModel `json:"annotations,omitempty"`

	Metadata *MetadataFields `json:"metadata,omitempty"`

	// Documents returned from search.
	Documents []RankedDocument `json:"documents,omitempty"`

	SubQueries []SearchModel `json:"subQueries,omitempty"`
}


// UnmarshalSearchModel unmarshals an instance of SearchModel from the specified map of raw messages.
func UnmarshalSearchModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SearchModel)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pageNumber", &obj.PageNumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "get_limit", &obj.GetLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "totalDocumentCount", &obj.TotalDocumentCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "types", &obj.Types)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalAttribute)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ranges", &obj.Ranges, UnmarshalRangeModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "typeahead", &obj.Typeahead, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "aggregations", &obj.Aggregations, UnmarshalAggregationModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dateHistograms", &obj.DateHistograms, UnmarshalHistogramData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "qualifiers", &obj.Qualifiers, UnmarshalQualifer)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "backend", &obj.Backend, UnmarshalBackend)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expandedQuery", &obj.ExpandedQuery)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "parsedBoolExpression", &obj.ParsedBoolExpression, UnmarshalBooleanOperands)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "conceptsExist", &obj.ConceptsExist)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cursorId", &obj.CursorID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vocabs", &obj.Vocabs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotations", &obj.Annotations, UnmarshalCommonDataModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadataFields)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "documents", &obj.Documents, UnmarshalRankedDocument)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subQueries", &obj.SubQueries, UnmarshalSearchModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SentenceModel : Object representing a document sentence.
type SentenceModel struct {
	// Document section for sentence.
	DocumentSection *string `json:"documentSection,omitempty"`

	Text *string `json:"text,omitempty"`

	// Starting sentence offset.
	Begin *int64 `json:"begin,omitempty"`

	// Ending sentence offset.
	End *int64 `json:"end,omitempty"`

	// Timestamp of sentence in video transcript.
	Timestamp *int64 `json:"timestamp,omitempty"`
}


// UnmarshalSentenceModel unmarshals an instance of SentenceModel from the specified map of raw messages.
func UnmarshalSentenceModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SentenceModel)
	err = core.UnmarshalPrimitive(m, "documentSection", &obj.DocumentSection)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceStatus : Object representing service runtime status.
type ServiceStatus struct {
	// scurrent service state.
	ServiceState *string `json:"serviceState,omitempty"`

	// service state details.
	StateDetails *string `json:"stateDetails,omitempty"`
}

// Constants associated with the ServiceStatus.ServiceState property.
// scurrent service state.
const (
	ServiceStatus_ServiceState_Error = "ERROR"
	ServiceStatus_ServiceState_Ok = "OK"
	ServiceStatus_ServiceState_Warning = "WARNING"
)


// UnmarshalServiceStatus unmarshals an instance of ServiceStatus from the specified map of raw messages.
func UnmarshalServiceStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceStatus)
	err = core.UnmarshalPrimitive(m, "serviceState", &obj.ServiceState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stateDetails", &obj.StateDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TextSpan : Objeft representing a document text span.
type TextSpan struct {
	// Document section where artifact was found.
	Section *string `json:"section,omitempty"`

	// Start of text span.
	Begin *int64 `json:"begin,omitempty"`

	// End of text span.
	End *int64 `json:"end,omitempty"`

	// Covered text span.
	CoveredText *string `json:"coveredText,omitempty"`

	// Documemnt provider.
	Source *string `json:"source,omitempty"`

	// Text span type.
	Type *string `json:"type,omitempty"`
}


// UnmarshalTextSpan unmarshals an instance of TextSpan from the specified map of raw messages.
func UnmarshalTextSpan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TextSpan)
	err = core.UnmarshalPrimitive(m, "section", &obj.Section)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "begin", &obj.Begin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "coveredText", &obj.CoveredText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Title : Model for setting title boost factor.
type Title struct {
	// Boost factor
	Boost *int64 `json:"boost,omitempty"`
}

// UnmarshalTitle unmarshals an instance of Title from the specified map of raw messages.
func UnmarshalTitle(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Title)
	err = core.UnmarshalPrimitive(m, "boost", &obj.Boost)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnstructuredModel : Model representing unstructed text.
type UnstructuredModel struct {
	// Text of the document.
	Text *string `json:"text,omitempty"`

	// Model representing ontology artifacts.
	Data *DataModel `json:"data,omitempty"`
}


// UnmarshalUnstructuredModel unmarshals an instance of UnstructuredModel from the specified map of raw messages.
func UnmarshalUnstructuredModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UnstructuredModel)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data", &obj.Data, UnmarshalDataModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
