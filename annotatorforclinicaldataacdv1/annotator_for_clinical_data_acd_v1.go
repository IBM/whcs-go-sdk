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

// Package annotatorforclinicaldataacdv1 : Operations and models for the AnnotatorForClinicalDataAcdV1 service
package annotatorforclinicaldataacdv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/whcs-go-sdk/common"
	"io"
	"reflect"
)

// AnnotatorForClinicalDataAcdV1 : Natural Language Processing (NLP) service featuring a set of medical domain
// annotators for use in detecting entities and medical concepts from unstructured data. Multiple annotators may be
// invoked from a single request.
//
// Version: 1.0.0 2020-04-20T09:06:17Z
type AnnotatorForClinicalDataAcdV1 struct {
	Service *core.BaseService

	// The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.
	Version *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://annotator-for-clinical-data-acd.cloud.ibm.com/services/clinical_data_annotator/api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ACD_SERVICE"


// AnnotatorForClinicalDataAcdV1Options : Service options
type AnnotatorForClinicalDataAcdV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.
	Version *string `validate:"required"`
}

// NewAnnotatorForClinicalDataAcdV1UsingExternalConfig : constructs an instance of AnnotatorForClinicalDataAcdV1 with passed in options and external configuration.
func NewAnnotatorForClinicalDataAcdV1UsingExternalConfig(options *AnnotatorForClinicalDataAcdV1Options) (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	annotatorForClinicalDataAcd, err = NewAnnotatorForClinicalDataAcdV1(options)
	if err != nil {
		return
	}

	err = annotatorForClinicalDataAcd.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = annotatorForClinicalDataAcd.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAnnotatorForClinicalDataAcdV1 : constructs an instance of AnnotatorForClinicalDataAcdV1 with passed in options.
func NewAnnotatorForClinicalDataAcdV1(options *AnnotatorForClinicalDataAcdV1Options) (service *AnnotatorForClinicalDataAcdV1, err error) {
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

	service = &AnnotatorForClinicalDataAcdV1{
		Service: baseService,
		Version: options.Version,
	}

	return
}

// SetServiceURL sets the service URL
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) SetServiceURL(url string) error {
	return annotatorForClinicalDataAcd.Service.SetServiceURL(url)
}

// GetProfiles : Get list of available persisted profiles
// Returns a summary including ID and description of the available persisted profiles.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetProfiles(getProfilesOptions *GetProfilesOptions) (result map[string] AcdProfile, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getProfilesOptions, "getProfilesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/profiles"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProfilesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetProfiles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdProfile)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// CreateProfile : Persist a new profile
// This API persists a new profile.  A profile is identified by an ID.  This ID can optionally be specified as part of
// the request body when invoking <b>POST /v1/analyze</b> API.  A profile contains annotator configuration information
// that will be applied to the annotators specified in the annotator flow.<p>If a caller would choose to have the ID of
// the new profile generated on their behalf, then in the request body the "id" field of the profile definition should
// be an empty string ("").  The auto-generated ID would be a normalized form of the "name" field from the profile
// definition.<p><b>Sample Profile #1</b><br>A profile definition that configures the 'concept_detection' annotator to
// use the UMLS umls.latest library.<br><pre>{<br>  "id": "acd_profile_cd_umls_latest",<br>  "name": "Profile for the
// latest Concept Detection UMLS Library",<br>  "description": "Provides configurations for running Concept Detection
// with the latest UMLS library",<br>  "annotators": [<br>    {<br>      "name": "concept_detection",<br>
// "parameters": {<br>         "libraries": ["umls.latest"]<br>       }<br>    }<br>  ]<br>}</pre><p><b>Sample Profile
// #2</b><br>A profile definition that configures the 'concept_detection' annotator to exclude any annotations where the
// semantic type does not equal 'neop'.<br><pre>{<br>  "id": "acd_profile_cd_neop_only",<br>  "name": "Profile for
// Concept Detection neop Semantic Type",<br>  "description": "Concept Detection configuration fitler to exclude
// annotations where semantic type does not equal 'neop'.",<br>  "annotators": [<br>    {<br>       "name":
// "concept_detection",<br>       "configurations": [<br>         {<br>           "filter": {<br>             "target":
// "unstructured.data.concepts",<br>             "condition": {<br>                "type": "match",<br>
// "field": "semanticType",<br>                "values": [<br>                   "neop"<br>                 ],<br>
//          "not": false,<br>                "caseInsensitive": false,<br>                "operator": "equals"<br>
//        }<br>            }<br>         }<br>       ]<br>    }<br>  ]<br>}</pre>.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) CreateProfile(createProfileOptions *CreateProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProfileOptions, "createProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProfileOptions, "createProfileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/profiles"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "CreateProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	body := make(map[string]interface{})
	if createProfileOptions.ID != nil {
		body["id"] = createProfileOptions.ID
	}
	if createProfileOptions.Name != nil {
		body["name"] = createProfileOptions.Name
	}
	if createProfileOptions.Description != nil {
		body["description"] = createProfileOptions.Description
	}
	if createProfileOptions.PublishedDate != nil {
		body["publishedDate"] = createProfileOptions.PublishedDate
	}
	if createProfileOptions.Publish != nil {
		body["publish"] = createProfileOptions.Publish
	}
	if createProfileOptions.Version != nil {
		body["version"] = createProfileOptions.Version
	}
	if createProfileOptions.CartridgeID != nil {
		body["cartridgeId"] = createProfileOptions.CartridgeID
	}
	if createProfileOptions.Annotators != nil {
		body["annotators"] = createProfileOptions.Annotators
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
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

	var result *AcdProfile
        if rawResponse != nil { 
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdProfile)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// GetProfile : Get details of a specific profile
// Using the specified profile ID, retrieves the profile definition.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetProfilesByID(getProfilesByIdOptions *GetProfilesByIdOptions) (result *AcdProfile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProfilesByIdOptions, "getProfilesByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProfilesByIdOptions, "getProfilesByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/profiles"}
	pathParameters := []string{*getProfilesByIdOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProfilesByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetProfilesById")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdProfile)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// UpdateProfile : Update a persisted profile definition
// Using the specified Profile ID, updates the profile definition.  This is a complete replacement of the existing
// profile definition using the JSON object provided in the request body.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) UpdateProfile(updateProfileOptions *UpdateProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProfileOptions, "updateProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProfileOptions, "updateProfileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/profiles"}
	pathParameters := []string{*updateProfileOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "UpdateProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	body := make(map[string]interface{})
	if updateProfileOptions.NewID != nil {
		body["id"] = updateProfileOptions.NewID
	}
	if updateProfileOptions.NewName != nil {
		body["name"] = updateProfileOptions.NewName
	}
	if updateProfileOptions.NewDescription != nil {
		body["description"] = updateProfileOptions.NewDescription
	}
	if updateProfileOptions.NewPublishedDate != nil {
		body["publishedDate"] = updateProfileOptions.NewPublishedDate
	}
	if updateProfileOptions.NewPublish != nil {
		body["publish"] = updateProfileOptions.NewPublish
	}
	if updateProfileOptions.NewVersion != nil {
		body["version"] = updateProfileOptions.NewVersion
	}
	if updateProfileOptions.NewCartridgeID != nil {
		body["cartridgeId"] = updateProfileOptions.NewCartridgeID
	}
	if updateProfileOptions.NewAnnotators != nil {
		body["annotators"] = updateProfileOptions.NewAnnotators
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
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

	var result *AcdProfile
        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdProfile)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// DeleteProfile : Delete a persisted profile
// Using the specified profile ID, deletes the profile from the list of persisted profiles.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) DeleteProfile(deleteProfileOptions *DeleteProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProfileOptions, "deleteProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProfileOptions, "deleteProfileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/profiles"}
	pathParameters := []string{*deleteProfileOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "DeleteProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

        var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

	var result *AcdProfile
        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdProfile)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// GetFlows : Get list of available persisted flows
// Returns a summary including ID and description of the available persisted flows.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetFlows(getFlowsOptions *GetFlowsOptions) (result map[string]AcdFlow, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getFlowsOptions, "getFlowsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/flows"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFlowsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetFlows")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdFlow)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// CreateFlows : Persist a new flow definition
// This API persists a new flow.  A flow is identified by an ID.  This ID can optionally be specified as part of the
// request body when invoking <b>POST /v1/analyze</b> API.  A flow definition contains a list one or more annotators,
// and optionally can include annotator configuration, a flow ID, and/or flow sequence.<p>If a caller would choose to
// have the ID of the new flow generated on their behalf, then in the request body the "id" field of the flow definition
// should be an empty string ("").  The auto-generated ID would be a normalized form of the "name" field from the flow
// definition.<p><p><b>Sample Flow #1</b><br>A flow definition that includes two annotators.<br><pre>{<br>  "id":
// "flow_simple",<br>  "name": "flow simple",<br>  "description": "A simple flow with two annotators",<br>
// "annotatorFlows": [<br>      {<br>       "flow": {<br>          "elements": [<br>             {<br>
// "annotator": {<br>                   "name": "concept_detection"<br>                }<br>             },<br>
//    {<br>               "annotator": {<br>                   "name": "symptom_disease"<br>                }<br>
//      }<br>           ],<br>       "async": false<br>        }<br>      }<br>   ]<br>}</pre><p><b>Sample Flow
// #2</b><br>A flow definition that includes the 'concept_detection' annotator and configuration details for the
// 'concept_detection' annotator.<br><pre>{<br>  "id": "flow_concept_detection_exclude_non_neop",<br>  "name": "flow
// concept detection exclude non neop",<br>  "description": "A flow excluding detected concepts that do not have 'neop'
// semantic type",<br>  "annotatorFlows": [<br>      {<br>       "flow": {<br>          "elements": [<br>
// {<br>               "annotator": {<br>                   "name": "concept_detection",<br>
// "configurations": [<br>                      {<br>                        "filter": {<br>
// "target": "unstructured.data.concepts",<br>                           "condition": {<br>
// "type": "match",<br>                              "field": "semanticType",<br>                              "values":
// [<br>                                 "neop"<br>                                ],<br>
// "not": false,<br>                              "caseInsensitive": false,<br>                              "operator":
// "equals"<br>                            }<br>                         }<br>                      }<br>
//     ]<br>                 }<br>              }<br>         ],<br>       "async": false<br>        }<br>      }<br>
// ]<br>}</pre>.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) CreateFlows(createFlowsOptions *CreateFlowsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createFlowsOptions, "createFlowsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createFlowsOptions, "createFlowsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/flows"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createFlowsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "CreateFlows")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	body := make(map[string]interface{})
	if createFlowsOptions.ID != nil {
		body["id"] = createFlowsOptions.ID
	}
	if createFlowsOptions.Name != nil {
		body["name"] = createFlowsOptions.Name
	}
	if createFlowsOptions.Description != nil {
		body["description"] = createFlowsOptions.Description
	}
	if createFlowsOptions.PublishedDate != nil {
		body["publishedDate"] = createFlowsOptions.PublishedDate
	}
	if createFlowsOptions.Publish != nil {
		body["publish"] = createFlowsOptions.Publish
	}
	if createFlowsOptions.Version != nil {
		body["version"] = createFlowsOptions.Version
	}
	if createFlowsOptions.CartridgeID != nil {
		body["cartridgeId"] = createFlowsOptions.CartridgeID
	}
	if createFlowsOptions.AnnotatorFlows != nil {
		body["annotatorFlows"] = createFlowsOptions.AnnotatorFlows
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
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

	var result *AcdFlow
        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdFlow)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// GetFlowsByID : Get details of a specific flow
// Using the specified Flow ID, retrieves the flow definition.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetFlowsByID(getFlowsByIdOptions *GetFlowsByIdOptions) (result *AcdFlow, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getFlowsByIdOptions, "getFlowsByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getFlowsByIdOptions, "getFlowsByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/flows"}
	pathParameters := []string{*getFlowsByIdOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFlowsByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetFlowsByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdFlow)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// UpdateFlows : Update a persisted flow definition
// Using the specified Flow ID, updates the persisted flow definition.  This is a complete replacement of the existing
// flow definition using the JSON object provided in the request body.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) UpdateFlows(updateFlowsOptions *UpdateFlowsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateFlowsOptions, "updateFlowsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateFlowsOptions, "updateFlowsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/flows"}
	pathParameters := []string{*updateFlowsOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateFlowsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "UpdateFlows")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	body := make(map[string]interface{})
	if updateFlowsOptions.NewID != nil {
		body["id"] = updateFlowsOptions.NewID
	}
	if updateFlowsOptions.NewName != nil {
		body["name"] = updateFlowsOptions.NewName
	}
	if updateFlowsOptions.NewDescription != nil {
		body["description"] = updateFlowsOptions.NewDescription
	}
	if updateFlowsOptions.NewPublishedDate != nil {
		body["publishedDate"] = updateFlowsOptions.NewPublishedDate
	}
	if updateFlowsOptions.NewPublish != nil {
		body["publish"] = updateFlowsOptions.NewPublish
	}
	if updateFlowsOptions.NewVersion != nil {
		body["version"] = updateFlowsOptions.NewVersion
	}
	if updateFlowsOptions.NewCartridgeID != nil {
		body["cartridgeId"] = updateFlowsOptions.NewCartridgeID
	}
	if updateFlowsOptions.NewAnnotatorFlows != nil {
		body["annotatorFlows"] = updateFlowsOptions.NewAnnotatorFlows
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
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

	var result *AcdFlow
        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdFlow)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// DeleteFlows : Delete a persisted flow
// Using the specified Flow ID, deletes the flow from the list of persisted flows.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) DeleteFlows(deleteFlowsOptions *DeleteFlowsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteFlowsOptions, "deleteFlowsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteFlowsOptions, "deleteFlowsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/flows"}
	pathParameters := []string{*deleteFlowsOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteFlowsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "DeleteFlows")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

        var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

	var result *AcdFlow
        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdFlow)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// RunPipeline : Detect entities & relations from unstructured data
// <p>This API accepts a JSON request model featuring both the unstructured data to be analyzed as well as the desired
// annotator flow.<p/><p><b>Annotator Chaining</b><br/>Sample request invoking both the concept_detection and
// symptom_disease annotators asynchronously. This sample request references configurations via a profile id. Profiles
// define configurations that can be referenced within a request. Profile is optional. A default profile is used if no
// profile id is available in the annotator flow. The default profile contains the parameters for the concept detection
// and the attribute detection. An empty profile can be used if absolutely no parameters are attached to any annotators.
// See <a href=".." target="_blank">documentation</a> for more information. </p><pre>{<br/>  "annotatorFlows": [<br/>
// {<br/>      "profile" : "default_profile_v1.0", <br/>      "flow": {<br/>        "elements": [<br/>          {<br/>
//          "annotator": {<br/>              "name": "concept_detection"<br/>            }<br/>          },<br/>
//  {<br/>            "annotator": {<br/>              "name": "symptom_disease"<br/>             }<br/>          }<br/>
//        ],<br/>        "async": false<br/>      }<br/>    }<br/>  ],<br/>  "unstructured": [<br/>    {<br/>
// "text": "Patient has lung cancer, but did not smoke. She may consider chemotherapy as part of a treatment plan."<br/>
//    }<br/>  ]<br/>}<br/></pre><p><b>Annotation Filtering</b><br/>Sample request invoking concept_detection with a
// filter defined to exclude any annotations detected from concept_detection where the semanticType field does not equal
// "neop".</p><pre>{<br/>  "annotatorFlows": [<br/>    {<br/>      "flow": {<br/>        "elements": [<br/>
// {<br/>            "annotator": {<br/>              "name": "concept_detection",<br/>              "configurations":
// [<br/>                {<br/>                  "filter": {<br/>                     "target":
// "unstructured.data.concepts",<br/>                     "condition": {<br/>                        "type":
// "match",<br/>                        "field": "semanticType",<br/>                        "values": [<br/>
//                "neop"<br/>                         ],<br/>                        "not": false,<br/>
//       "caseInsensitive": false,<br/>                        "operator": "equals"<br/>                     }<br/>
//             }<br/>                }<br/>              ]<br/>            }<br/>          }<br/>        ],<br/>
// "async": false<br/>      }<br/>    }<br/>  ],<br/>  "unstructured": [<br/>    {<br/>      "text": "Patient has lung
// cancer, but did not smoke. She may consider chemotherapy as part of a treatment plan."<br/>    }<br/>
// ]<br/>}<br/></pre><p><b>Annotators that support annotation filtering:</b> allergy, bathing_assistance, cancer,
// concept_detection, dressing_assistance, eating_assistance, ejection_fraction, lab_value, medication, named_entities,
// procedure, seeing_assistance, smoking, symptom_disease, toileting_assistance,
// walking_assistance.</p><hr/><p><b>Annotation Augmentation</b><br/>Sample request invoking the cancer annotator and
// providing a whitelist entry for a new custom surface form: "lungcancer".</p><pre>{<br/> "annotatorFlows": [<br/>
// {<br/>     "flow": {<br/>       "elements": [<br/>          {<br/>           "annotator": {<br/>             "name":
// "cancer",<br/>             "configurations": [<br/>                {<br/>                 "whitelist": {<br/>
//           "name": "cancer",<br/>                   "entries": [<br/>                      {<br/>
// "surfaceForms": [<br/>                   "lungcancer"<br/>                ],<br/>               "features": {<br/>
//                "normalizedName": "lung cancer",<br/>                   "hccCode": "9",<br/>
// "icd10Code": "C34.9",<br/>                   "ccsCode": "19",<br/>                   "icd9Code": "162.9",<br/>
//            "conceptId": "93880001"<br/>                }<br/>                      }<br/>                    ]<br/>
//                }<br/>                }<br/>              ]<br/>            }<br/>          }<br/>        ],<br/>
//  "async": false<br/>      }<br/>    }<br/>  ],<br/> "unstructured": [<br/>    {<br/>     "text": "The patient was
// diagnosed with lungcancer, on Dec 23, 2011."<br/>    }<br/>  ]<br/>}<br/></pre><b>Annotators that support annotation
// augmentation:</b> allergy, bathing_assistance, cancer, dressing_assistance, eating_assistance, ejection_fraction,
// lab_value, medication, named_entities, procedure, seeing_assistance, smoking, symptom_disease, toileting_assistance,
// walking_assistance.<br/>.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) RunPipeline(runPipelineOptions *RunPipelineOptions) (result *AnalyticFlowBeanInput, response *core.DetailedResponse, err error) {

	err = core.ValidateNotNil(runPipelineOptions, "runPipelineOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runPipelineOptions, "runPipelineOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/analyze"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range runPipelineOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "RunPipeline")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))
	if runPipelineOptions.DebugTextRestore != nil {
		builder.AddQuery("debug_text_restore", fmt.Sprint(*runPipelineOptions.DebugTextRestore))
	}
	if runPipelineOptions.ReturnAnalyzedText != nil {
		builder.AddQuery("return_analyzed_text", fmt.Sprint(*runPipelineOptions.ReturnAnalyzedText))
	}

	body := make(map[string]interface{})
	if runPipelineOptions.Unstructured != nil {
		body["unstructured"] = runPipelineOptions.Unstructured
	}
	if runPipelineOptions.AnnotatorFlows != nil {
		body["annotatorFlows"] = runPipelineOptions.AnnotatorFlows
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
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticFlowBeanInput)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// RunPipelineWithFlow : analyze with a pre-specified flow
// <p>This API accepts a flow identifier as well as a <emph>TEXT</emph> or a <emph>JSON</emph> request model featuring
// the unstructured text to be analyzed. <p/><p><b>JSON request model with unstructured text </b></p><pre>{<br/>
// "unstructured": [<br/>    {<br/>      "text": "Patient has lung cancer, but did not smoke. She may consider
// chemotherapy as part of a treatment plan."<br/>    }<br/>  ]<br/>}<br/></pre><p><b>JSON request model with existing
// annotations </b><br/></p><pre>{<br> "unstructured": [<br>    {<br>      "text": "Patient will not start on cisplatin
// 80mg on 1/1/2018. Patient is also diabetic.",<br>      "data": {<br>        "concepts": [<br>          {<br>
//   "cui": "C0030705",<br>            "preferredName": "Patients",<br>            "semanticType": "podg",<br>
//  "source": "umls",<br>            "sourceVersion": "2017AA",<br>            "type":
// "umls.PatientOrDisabledGroup",<br>            "begin": 0,<br>            "end": 7,<br>            "coveredText":
// "Patient"<br>          }<br> ]<br>      }  <br>    } <br> ]<br>}<br></pre>.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) RunPipelineWithFlow(runPipelineWithFlowOptions *RunPipelineWithFlowOptions) (result *AnalyticFlowBeanInput, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(runPipelineWithFlowOptions, "runPipelineWithFlowOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(runPipelineWithFlowOptions, "runPipelineWithFlowOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/analyze"}
	pathParameters := []string{*runPipelineWithFlowOptions.FlowID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range runPipelineWithFlowOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "RunPipelineWithFlow")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))
	builder.AddQuery("return_analyzed_text", fmt.Sprint(*runPipelineWithFlowOptions.ReturnAnalyzedText))
	if runPipelineWithFlowOptions.DebugTextRestore != nil {
		builder.AddQuery("debug_text_restore", fmt.Sprint(*runPipelineWithFlowOptions.DebugTextRestore))
	}

	if (runPipelineWithFlowOptions.AnalyticFlowBeanInput != nil && runPipelineWithFlowOptions.AnalyticFlowBeanInput.Unstructured != nil) {
		contentType := "application/json"
		if runPipelineWithFlowOptions.ContentType != nil {
			contentType = *runPipelineWithFlowOptions.ContentType
		}
		builder.AddHeader("Content-Type", fmt.Sprint(contentType))
		_, err = builder.SetBodyContentJSON(runPipelineWithFlowOptions.AnalyticFlowBeanInput)
		if err != nil {
			return
		}
	} else if (runPipelineWithFlowOptions.Body != nil) {
		contentType := "text/plain"
		if runPipelineWithFlowOptions.ContentType != nil {
			contentType = *runPipelineWithFlowOptions.ContentType
		}
		builder.AddHeader("Content-Type", fmt.Sprint(contentType))
		_, err = builder.SetBodyContent(contentType, nil, nil, runPipelineWithFlowOptions.Body)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("Text must be supplied to be analyzed")
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnalyticFlowBeanInput)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// GetAnnotators : Get list of available annotators
// Get list of available annotators that can be leveraged to detect information from unstructured data. One or more
// annnotators can be leveraged within a single request to the service.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetAnnotators(getAnnotatorsOptions *GetAnnotatorsOptions) (result map[string]AnnotatorDescription, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAnnotatorsOptions, "getAnnotatorsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/annotators"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAnnotatorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetAnnotators")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnnotatorDescription)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// GetAnnotatorsByID : Get details of a specific annotator
// Get details of an annotator that can be used to detect information from unstructured data.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetAnnotatorsByID(getAnnotatorsByIdOptions *GetAnnotatorsByIdOptions) (result *AnnotatorDescription, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAnnotatorsByIdOptions, "getAnnotatorsByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAnnotatorsByIdOptions, "getAnnotatorsByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/annotators"}
	pathParameters := []string{*getAnnotatorsByIdOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAnnotatorsByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetAnnotatorsByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAnnotatorDescription)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// DeleteUserSpecificArtifacts : Delete tenant specific artifacts
// Delete tenant specific artifacts.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) DeleteUserSpecificArtifacts(deleteUserSpecificArtifactsOptions *DeleteUserSpecificArtifactsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteUserSpecificArtifactsOptions, "deleteUserSpecificArtifactsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/user_data"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserSpecificArtifactsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "DeleteUserSpecificArtifacts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = annotatorForClinicalDataAcd.Service.Request(request, nil)

	return
}

// CartridgesGet : Get list of available deployment status
// Returns a summary including ID and status of the available deployments.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) CartridgesGet(cartridgesGetOptions *CartridgesGetOptions) (result *ListAcdCartridges, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(cartridgesGetOptions, "cartridgesGetOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/cartridges"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range cartridgesGetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "CartridgesGet")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListAcdCartridges)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// CartridgesPostMultipart : Create a cartridge deployment
// Create a cartridge deployment from a cartridge archive file.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) CartridgesPostMultipart(cartridgesPostMultipartOptions *CartridgesPostMultipartOptions) (result *DeployCartridgeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cartridgesPostMultipartOptions, "cartridgesPostMultipartOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cartridgesPostMultipartOptions, "cartridgesPostMultipartOptions")
	if err != nil {
		return
	}
	if cartridgesPostMultipartOptions.ArchiveFile == nil {
		err = fmt.Errorf("Archive File must be supplied")
		return
	}

	pathSegments := []string{"v1/cartridges"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range cartridgesPostMultipartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "CartridgesPostMultipart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	if cartridgesPostMultipartOptions.ArchiveFile != nil {
		builder.AddFormData("archive_file", "filename",
			core.StringNilMapper(cartridgesPostMultipartOptions.ArchiveFileContentType), cartridgesPostMultipartOptions.ArchiveFile)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}
	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeployCartridgeResponse)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// CartridgesPutMultipart : Create a cartridge deployment
// Update a cartridge deployment from a cartridge archive file.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) CartridgesPutMultipart(cartridgesPutMultipartOptions *CartridgesPutMultipartOptions) (result *DeployCartridgeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cartridgesPutMultipartOptions, "cartridgesPutMultipartOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cartridgesPutMultipartOptions, "cartridgesPutMultipartOptions")
	if err != nil {
		return
	}
	if cartridgesPutMultipartOptions.ArchiveFile == nil {
		err = fmt.Errorf("Archive File must be supplied")
		return
	}

	pathSegments := []string{"v1/cartridges"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range cartridgesPutMultipartOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "CartridgesPutMultipart")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	if cartridgesPutMultipartOptions.ArchiveFile != nil {
		builder.AddFormData("archive_file", "filename",
			core.StringNilMapper(cartridgesPutMultipartOptions.ArchiveFileContentType), cartridgesPutMultipartOptions.ArchiveFile)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}
	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeployCartridgeResponse)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// CartridgesGetID : Get details of a specific deployment
// Using the specified Catridge ID, retrieves the deployment status.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) CartridgesGetID(cartridgesGetIdOptions *CartridgesGetIdOptions) (result *AcdCartridges, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cartridgesGetIdOptions, "cartridgesGetIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cartridgesGetIdOptions, "cartridgesGetIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/cartridges"}
	pathParameters := []string{*cartridgesGetIdOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range cartridgesGetIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "CartridgesGetID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcdCartridges)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// DeployCartridge : Deploy a cartridge
// Deploy a cartridge from a cartridge archive file.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) DeployCartridge(deployCartridgeOptions *DeployCartridgeOptions) (result *DeployCartridgeResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deployCartridgeOptions, "deployCartridgeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deployCartridgeOptions, "deployCartridgeOptions")
	if err != nil {
		return
	}
	if deployCartridgeOptions.ArchiveFile == nil {
		err = fmt.Errorf("Archive File must be supplied")
		return
	}

	pathSegments := []string{"v1/deploy"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deployCartridgeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "DeployCartridge")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("version", fmt.Sprint(*annotatorForClinicalDataAcd.Version))
	if deployCartridgeOptions.Update != nil {
		builder.AddQuery("update", fmt.Sprint(*deployCartridgeOptions.Update))
	}

	if deployCartridgeOptions.ArchiveFile != nil {
		builder.AddFormData("archive_file", "filename",
			core.StringNilMapper(deployCartridgeOptions.ArchiveFileContentType), deployCartridgeOptions.ArchiveFile)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeployCartridgeResponse)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// GetHealthCheckStatus : Determine if service is running correctly
// This resource differs from /status in that it will will always return a 500 error if the service state is not OK.
// This makes it simpler for service front ends (such as Datapower) to detect a failed service.
func (annotatorForClinicalDataAcd *AnnotatorForClinicalDataAcdV1) GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions) (result *ServiceStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHealthCheckStatusOptions, "getHealthCheckStatusOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/status/health_check"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(annotatorForClinicalDataAcd.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHealthCheckStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("annotator_for_clinical_data_acd", "V1", "GetHealthCheckStatus")
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
	response, err = annotatorForClinicalDataAcd.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}

        if rawResponse != nil {
	        err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceStatus)
	        if err != nil {
		        return
	        }
	        response.Result = result
        }

	return
}

// AcdCartridges : AcdCartridges struct
type AcdCartridges struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusCode *int64 `json:"statusCode,omitempty"`

	StatusLocation *string `json:"statusLocation,omitempty"`

	StartTime *string `json:"startTime,omitempty"`

	EndTime *string `json:"endTime,omitempty"`

	Duration *string `json:"duration,omitempty"`

	CorrelationID *string `json:"correlationId,omitempty"`

	ArtifactResponseCode *int64 `json:"artifactResponseCode,omitempty"`

	ArtifactResponse []ServiceError `json:"artifactResponse,omitempty"`
}

// UnmarshalAcdCartridges unmarshals an instance of AcdCartridges from the specified map of raw messages.
func UnmarshalAcdCartridges(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AcdCartridges)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "statusCode", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "statusLocation", &obj.StatusLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "startTime", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endTime", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duration", &obj.Duration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "correlationId", &obj.CorrelationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "artifactResponseCode", &obj.ArtifactResponseCode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "artifactResponse", &obj.ArtifactResponse, UnmarshalServiceError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListAcdCartridges : ListAcdCartridges struct
type ListAcdCartridges struct {
	// List of cartridges.
	Cartridges []AcdCartridges `json:"cartridges,omitempty"`
}

// UnmarshalListAcdCartridges unmarshals an instance of ListAcdCartridges from the specified map of raw messages.
func UnmarshalListAcdCartridges(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAcdCartridges)
	err = core.UnmarshalModel(m, "cartridges", &obj.Cartridges, UnmarshalAcdCartridges)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AcdFlow : AcdFlow struct
type AcdFlow struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	PublishedDate *string `json:"publishedDate,omitempty"`

	Publish *bool `json:"publish,omitempty"`

	Version *string `json:"version,omitempty"`

	CartridgeID *string `json:"cartridgeId,omitempty"`

	AnnotatorFlows []AnnotatorFlow `json:"annotatorFlows,omitempty"`
}

// UnmarshalAcdFlow unmarshals an instance of AcdFlow from the specified map of raw messages.
func UnmarshalAcdFlow(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AcdFlow)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publishedDate", &obj.PublishedDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publish", &obj.Publish)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cartridgeId", &obj.CartridgeID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotatorFlows", &obj.AnnotatorFlows, UnmarshalAnnotatorFlow)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AcdProfile : AcdProfile struct
type AcdProfile struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	PublishedDate *string `json:"publishedDate,omitempty"`

	Publish *bool `json:"publish,omitempty"`

	Version *string `json:"version,omitempty"`

	CartridgeID *string `json:"cartridgeId,omitempty"`

	Annotators []Annotator `json:"annotators,omitempty"`
}

// UnmarshalAcdProfile unmarshals an instance of AcdProfile from the specified map of raw messages.
func UnmarshalAcdProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AcdProfile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publishedDate", &obj.PublishedDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publish", &obj.Publish)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cartridgeId", &obj.CartridgeID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotators", &obj.Annotators, UnmarshalAnnotator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AdverseEvent : AdverseEvent insight model struct
type AdverseEvent struct {
	Score *float64 `json:"score,omitempty"`

	AllergyScore *float64 `json:"allergyScore,omitempty"`

	Usage *Usage `json:"usage,omitempty"`
}

// UnmarshalAdverseEvent unmarshals an instance of AdverseEvent from the specific raw message.
func UnmarshalAdverseEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AdverseEvent)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allergyScore", &obj.AllergyScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnalyticFlowBeanInput : AnalyticFlowBeanInput struct
type AnalyticFlowBeanInput struct {
	Unstructured []UnstructuredContainer `json:"unstructured,omitempty"`

	AnnotatorFlows []AnnotatorFlow `json:"annotatorFlows,omitempty"`
}

// NewAnalyticFlowBeanInput : Instantiate AnalyticFlowBeanInput
func (*AnnotatorForClinicalDataAcdV1) NewAnalyticFlowBeanInput() *AnalyticFlowBeanInput {
	return &AnalyticFlowBeanInput{}
}

// SetUnstructured : Allow user to set Unstructured
func (input *AnalyticFlowBeanInput) SetUnstructured(unstructuredContainer []UnstructuredContainer) {
	input.Unstructured = unstructuredContainer
}

// UnmarshalAnalyticFlowBeanInput unmarshals an instance of AnalyticFlowBeanInput from the specified map of raw messages.
func UnmarshalAnalyticFlowBeanInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnalyticFlowBeanInput)
	err = core.UnmarshalModel(m, "unstructured", &obj.Unstructured, UnmarshalUnstructuredContainer)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "annotatorFlows", &obj.AnnotatorFlows, UnmarshalAnnotatorFlow)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Annotation : Annotation struct
type Annotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`
}

// UnmarshalAnnotation unmarshals an instance of Annotation from the specified map of raw messages.
func UnmarshalAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Annotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Annotator : Annotator struct
type Annotator struct {
	Name *string `json:"name" validate:"required"`

	Parameters map[string][]string `json:"parameters,omitempty"`

	Configurations []ConfigurationEntity `json:"configurations,omitempty"`
}

// NewAnnotator : Instantiate Annotator (Generic Model Constructor)
func (*AnnotatorForClinicalDataAcdV1) NewAnnotator(name string) (model *Annotator, err error) {
	model = &Annotator{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAnnotator unmarshals an instance of Annotator from the specified map of raw messages.
func UnmarshalAnnotator(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Annotator)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "parameters", &obj.Parameters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configurations", &obj.Configurations, UnmarshalConfigurationEntity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnnotatorDescription : AnnotatorDescription struct
type AnnotatorDescription struct {
	Description *string `json:"description,omitempty"`

	Version *string `json:"version,omitempty"`
}

// UnmarshalAnnotatorDescription unmarshals an instance of AnnotatorDescription from the specified map of raw messages.
func UnmarshalAnnotatorDescription(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnnotatorDescription)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AnnotatorFlow : AnnotatorFlow struct
type AnnotatorFlow struct {
	Profile *string `json:"profile,omitempty"`

	Flow *Flow `json:"flow" validate:"required"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	Data map[string][]Entity `json:"data,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`

	GlobalConfigurations []ConfigurationEntity `json:"globalConfigurations,omitempty"`

	Uid *int64 `json:"uid,omitempty"`
}

// NewAnnotatorFlow : Instantiate AnnotatorFlow (Generic Model Constructor)
func (*AnnotatorForClinicalDataAcdV1) NewAnnotatorFlow(flow *Flow) (model *AnnotatorFlow, err error) {
	model = &AnnotatorFlow{
		Flow: flow,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAnnotatorFlow unmarshals an instance of AnnotatorFlow from the specified map of raw messages.
func UnmarshalAnnotatorFlow(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AnnotatorFlow)
	err = core.UnmarshalPrimitive(m, "profile", &obj.Profile)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "flow", &obj.Flow, UnmarshalFlow)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data", &obj.Data, UnmarshalEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "globalConfigurations", &obj.GlobalConfigurations, UnmarshalConfigurationEntity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.Uid)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssistanceAnnotation : AssistanceAnnotation struct
type AssistanceAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Modality *string `json:"modality,omitempty"`

	PrimaryActionNormalizedName *string `json:"primaryActionNormalizedName,omitempty"`

	PrimaryActionSurfaceForm *string `json:"primaryActionSurfaceForm,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`
}

// UnmarshalAssistanceAnnotation unmarshals an instance of AssistanceAnnotation from the specified map of raw messages.
func UnmarshalAssistanceAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssistanceAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modality", &obj.Modality)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primaryActionNormalizedName", &obj.PrimaryActionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primaryActionSurfaceForm", &obj.PrimaryActionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttributeValueAnnotation : AttributeValueAnnotation struct
type AttributeValueAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Concept *Concept `json:"concept,omitempty"`

	DisambiguationData *Disambiguation `json:"disambiguationData,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	CptCode *string `json:"cptCode,omitempty"`

	Icd9Code *string `json:"icd9Code,omitempty"`

	Icd10Code *string `json:"icd10Code,omitempty"`

	LoincId *string `json:"loincId,omitempty"`

	MeshId *string `json:"meshId,omitempty"`

	Name *string `json:"name,omitempty"`

	NciCode *string `json:"nciCode,omitempty"`

	PreferredName *string `json:"preferredName,omitempty"`

	RxNormId *string `json:"rxNormId,omitempty"`

	SnomedConceptId *string `json:"snomedConceptId,omitempty"`

	Source *string `json:"source,omitempty"`

	SourceVersion *string `json:"sourceVersion,omitempty"`

	Values []AttributeValueEntry `json:"values,omitempty"`

	Vocabs *string `json:"vocabs,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`

	CcsCode *string `json:"ccsCode,omitempty"`

	HccCode *string `json:"hccCode,omitempty"`

	RuleId *string `json:"ruleId,omitempty"`

	DerivedFrom []Concept `json:"derivedFrom,omitempty"`

	Temporal []Temporal `json:"temporal,omitempty"`

        EvidenceSpans []Reference `json:"evidenceSpans,omitempty"`
}

// UnmarshalAttributeValueAnnotation unmarshals an instance of AttributeValueAnnotation from the specified map of raw messages.
func UnmarshalAttributeValueAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttributeValueAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concept", &obj.Concept, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguationData", &obj.DisambiguationData, UnmarshalDisambiguation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cptCode", &obj.CptCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icd9Code", &obj.Icd9Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icd10Code", &obj.Icd10Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "loincId", &obj.LoincId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "meshId", &obj.MeshId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nciCode", &obj.NciCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rxNormId", &obj.RxNormId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "snomedConceptId", &obj.SnomedConceptId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sourceVersion", &obj.SourceVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "values", &obj.Values, UnmarshalAttributeValueEntry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vocabs", &obj.Vocabs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ccsCode", &obj.CcsCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hccCode", &obj.HccCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ruleId", &obj.RuleId)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "derivedFrom", &obj.DerivedFrom, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "temporal", &obj.Temporal, UnmarshalTemporal)
	if err != nil {
		return
	}
        err = core.UnmarshalModel(m, "evidenceSpans", &obj.EvidenceSpans, UnmarshalReference)
        if err != nil {
                return
        }
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttributeValueEntry : AttributeValueEntry struct
type AttributeValueEntry struct {
	// Value of entry
	Value *string `json:"value,omitempty"`

	// Unit of measure
	Unit *string `json:"unit,omitempty"`

	// Frequency of value
	Frequency *string `json:"frequency,omitempty"`

	// Duration of value
	Duration *string `json:"duration,omitempty"`

	// Dimension of value
	Dimension *string `json:"dimension,omitempty"`
}

// UnmarshalAttributeValueEntry unmarshals an instance of AttributeValueEntry from the specified map of raw messages.
func UnmarshalAttributeValueEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttributeValueEntry)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "frequency", &obj.Frequency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duration", &obj.Duration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dimension", &obj.Dimension)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CancerAnnotation : CancerAnnotation struct
type CancerAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Cui *string `json:"cui,omitempty"`

	Modality *string `json:"modality,omitempty"`

	Cancer []map[string]interface{} `json:"cancer,omitempty"`

	DisambiguationData *Disambiguation `json:"disambiguationData,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`
}

// UnmarshalCancerAnnotation unmarshals an instance of CancerAnnotation from the specified map of raw messages.
func UnmarshalCancerAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CancerAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modality", &obj.Modality)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguationData", &obj.DisambiguationData, UnmarshalDisambiguation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cancer", &obj.Cancer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CartridgesGetIdOptions : The CartridgesGetID options.
type CartridgesGetIdOptions struct {
	// Cartridge ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCartridgesGetIdOptions : Instantiate CartridgesGetIdOptions
func (*AnnotatorForClinicalDataAcdV1) NewCartridgesGetIdOptions(id string) *CartridgesGetIdOptions {
	return &CartridgesGetIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *CartridgesGetIdOptions) SetID(id string) *CartridgesGetIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CartridgesGetIdOptions) SetHeaders(param map[string]string) *CartridgesGetIdOptions {
	options.Headers = param
	return options
}

// CartridgesGetOptions : The CartridgesGet options.
type CartridgesGetOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCartridgesGetOptions : Instantiate CartridgesGetOptions
func (*AnnotatorForClinicalDataAcdV1) NewCartridgesGetOptions() *CartridgesGetOptions {
	return &CartridgesGetOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *CartridgesGetOptions) SetHeaders(param map[string]string) *CartridgesGetOptions {
	options.Headers = param
	return options
}

// CartridgesPostMultipartOptions : The CartridgesPostMultipart options.
type CartridgesPostMultipartOptions struct {
	// Cartridge archive file.
	ArchiveFile io.ReadCloser `json:"archive_file,omitempty"`

	// The content type of archiveFile.
	ArchiveFileContentType *string `json:"archive_file_content_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCartridgesPostMultipartOptions : Instantiate CartridgesPostMultipartOptions
func (*AnnotatorForClinicalDataAcdV1) NewCartridgesPostMultipartOptions() *CartridgesPostMultipartOptions {
	return &CartridgesPostMultipartOptions{}
}

// SetArchiveFile : Allow user to set ArchiveFile
func (options *CartridgesPostMultipartOptions) SetArchiveFile(archiveFile io.ReadCloser) *CartridgesPostMultipartOptions {
	options.ArchiveFile = archiveFile
	return options
}

// SetArchiveFileContentType : Allow user to set ArchiveFileContentType
func (options *CartridgesPostMultipartOptions) SetArchiveFileContentType(archiveFileContentType string) *CartridgesPostMultipartOptions {
	options.ArchiveFileContentType = core.StringPtr(archiveFileContentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CartridgesPostMultipartOptions) SetHeaders(param map[string]string) *CartridgesPostMultipartOptions {
	options.Headers = param
	return options
}

// CartridgesPutMultipartOptions : The CartridgesPutMultipart options.
type CartridgesPutMultipartOptions struct {
	// Cartridge archive file.
	ArchiveFile io.ReadCloser `json:"archive_file,omitempty"`

	// The content type of archiveFile.
	ArchiveFileContentType *string `json:"archive_file_content_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCartridgesPutMultipartOptions : Instantiate CartridgesPutMultipartOptions
func (*AnnotatorForClinicalDataAcdV1) NewCartridgesPutMultipartOptions() *CartridgesPutMultipartOptions {
	return &CartridgesPutMultipartOptions{}
}

// SetArchiveFile : Allow user to set ArchiveFile
func (options *CartridgesPutMultipartOptions) SetArchiveFile(archiveFile io.ReadCloser) *CartridgesPutMultipartOptions {
	options.ArchiveFile = archiveFile
	return options
}

// SetArchiveFileContentType : Allow user to set ArchiveFileContentType
func (options *CartridgesPutMultipartOptions) SetArchiveFileContentType(archiveFileContentType string) *CartridgesPutMultipartOptions {
	options.ArchiveFileContentType = core.StringPtr(archiveFileContentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CartridgesPutMultipartOptions) SetHeaders(param map[string]string) *CartridgesPutMultipartOptions {
	options.Headers = param
	return options
}

// Concept : Concept struct
type Concept struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Cui *string `json:"cui,omitempty"`

	DisambiguationData *Disambiguation `json:"disambiguationData,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	CptCode *string `json:"cptCode,omitempty"`

	Icd9Code *string `json:"icd9Code,omitempty"`

	Icd10Code *string `json:"icd10Code,omitempty"`

	LoincId *string `json:"loincId,omitempty"`

	MeshId *string `json:"meshId,omitempty"`

	Name *string `json:"name,omitempty"`

	NciCode *string `json:"nciCode,omitempty"`

	PreferredName *string `json:"preferredName,omitempty"`

	RxNormId *string `json:"rxNormId,omitempty"`

	SnomedConceptId *string `json:"snomedConceptId,omitempty"`

	Source *string `json:"source,omitempty"`

	SourceVersion *string `json:"sourceVersion,omitempty"`

	SemanticType *string `json:"semanticType,omitempty"`

	Vocabs *string `json:"vocabs,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`

	RuleId *string `json:"ruleId,omitempty"`

	DerivedFrom []Concept `json:"derivedFrom,omitempty"`

	Temporal []Temporal `json:"temporal,omitempty"`

        SelectionLabel *string `json:"selectionLabel,omitempty"`

        ValueIndex *int64 `json:"valueIndex,omitempty"`
}

// UnmarshalConcept unmarshals an instance of Concept from the specified map of raw messages.
func UnmarshalConcept(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Concept)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguationData", &obj.DisambiguationData, UnmarshalDisambiguation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cptCode", &obj.CptCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icd9Code", &obj.Icd9Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icd10Code", &obj.Icd10Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "loincId", &obj.LoincId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "meshId", &obj.MeshId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nciCode", &obj.NciCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rxNormId", &obj.RxNormId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "snomedConceptId", &obj.SnomedConceptId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sourceVersion", &obj.SourceVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semanticType", &obj.SemanticType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vocabs", &obj.Vocabs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ruleId", &obj.RuleId)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "derivedFrom", &obj.DerivedFrom, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "temporal", &obj.Temporal, UnmarshalTemporal)
	if err != nil {
		return
	}
        err = core.UnmarshalPrimitive(m, "selectionLabel", &obj.SelectionLabel)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "valueIndex", &obj.ValueIndex)
        if err != nil {
                return
        }
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConceptValue : ConceptValue struct
type ConceptValue struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Cui *string `json:"cui,omitempty"`

	Dimension *string `json:"dimension,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	PreferredName *string `json:"preferredName,omitempty"`

	Trigger *string `json:"trigger,omitempty"`

	Value *string `json:"value,omitempty"`

	Source *string `json:"source,omitempty"`

	RuleId *string `json:"ruleId,omitempty"`

	DerivedFrom []Concept `json:"derivedFrom,omitempty"`

	Unit *string `json:"unit,omitempty"`

	Values []map[string]interface{} `json:"values,omitempty"`

	RangeBegin *string `json:"rangeBegin,omitempty"`

	RangeEnd *string `json:"rangeEnd,omitempty"`
}

// UnmarshalConceptValue unmarshals an instance of ConceptValue from the specified map of raw messages.
func UnmarshalConceptValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConceptValue)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dimension", &obj.Dimension)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferredName", &obj.PreferredName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trigger", &obj.Trigger)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ruleId", &obj.RuleId)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "derivedFrom", &obj.DerivedFrom, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values", &obj.Values)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rangeBegin", &obj.RangeBegin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rangeEnd", &obj.RangeEnd)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationEntity : ConfigurationEntity struct
type ConfigurationEntity struct {
	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	Uid *int64 `json:"uid,omitempty"`

	Mergeid *int64 `json:"mergeid,omitempty"`
}

// UnmarshalConfigurationEntity unmarshals an instance of ConfigurationEntity from the specified map of raw messages.
func UnmarshalConfigurationEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationEntity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.Uid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mergeid", &obj.Mergeid)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateFlowsOptions : The CreateFlows options.
type CreateFlowsOptions struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	PublishedDate *string `json:"publishedDate,omitempty"`

	Publish *bool `json:"publish,omitempty"`

	Version *string `json:"version,omitempty"`

	CartridgeID *string `json:"cartridgeId,omitempty"`

	AnnotatorFlows []AnnotatorFlow `json:"annotatorFlows,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateFlowsOptions : Instantiate CreateFlowsOptions
func (*AnnotatorForClinicalDataAcdV1) NewCreateFlowsOptions() *CreateFlowsOptions {
	return &CreateFlowsOptions{}
}

// SetID : Allow user to set ID
func (options *CreateFlowsOptions) SetID(id string) *CreateFlowsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *CreateFlowsOptions) SetName(name string) *CreateFlowsOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateFlowsOptions) SetDescription(description string) *CreateFlowsOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetPublishedDate : Allow user to set PublishedDate
func (options *CreateFlowsOptions) SetPublishedDate(publishedDate string) *CreateFlowsOptions {
	options.PublishedDate = core.StringPtr(publishedDate)
	return options
}

// SetPublish : Allow user to set Publish
func (options *CreateFlowsOptions) SetPublish(publish bool) *CreateFlowsOptions {
	options.Publish = core.BoolPtr(publish)
	return options
}

// SetVersion : Allow user to set Version
func (options *CreateFlowsOptions) SetVersion(version string) *CreateFlowsOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetCartridgeID : Allow user to set CartridgeID
func (options *CreateFlowsOptions) SetCartridgeID(cartridgeID string) *CreateFlowsOptions {
	options.CartridgeID = core.StringPtr(cartridgeID)
	return options
}

// SetAnnotatorFlows : Allow user to set AnnotatorFlows
func (options *CreateFlowsOptions) SetAnnotatorFlows(annotatorFlows []AnnotatorFlow) *CreateFlowsOptions {
	options.AnnotatorFlows = annotatorFlows
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateFlowsOptions) SetHeaders(param map[string]string) *CreateFlowsOptions {
	options.Headers = param
	return options
}

// CreateProfileOptions : The CreateProfile options.
type CreateProfileOptions struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	PublishedDate *string `json:"publishedDate,omitempty"`

	Publish *bool `json:"publish,omitempty"`

	Version *string `json:"version,omitempty"`

	CartridgeID *string `json:"cartridgeId,omitempty"`

	Annotators []Annotator `json:"annotators,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProfileOptions : Instantiate CreateProfileOptions
func (*AnnotatorForClinicalDataAcdV1) NewCreateProfileOptions() *CreateProfileOptions {
	return &CreateProfileOptions{}
}

// SetID : Allow user to set ID
func (options *CreateProfileOptions) SetID(id string) *CreateProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *CreateProfileOptions) SetName(name string) *CreateProfileOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateProfileOptions) SetDescription(description string) *CreateProfileOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetPublishedDate : Allow user to set PublishedDate
func (options *CreateProfileOptions) SetPublishedDate(publishedDate string) *CreateProfileOptions {
	options.PublishedDate = core.StringPtr(publishedDate)
	return options
}

// SetPublish : Allow user to set Publish
func (options *CreateProfileOptions) SetPublish(publish bool) *CreateProfileOptions {
	options.Publish = core.BoolPtr(publish)
	return options
}

// SetVersion : Allow user to set Version
func (options *CreateProfileOptions) SetVersion(version string) *CreateProfileOptions {
	options.Version = core.StringPtr(version)
	return options
}

// SetCartridgeID : Allow user to set CartridgeID
func (options *CreateProfileOptions) SetCartridgeID(cartridgeID string) *CreateProfileOptions {
	options.CartridgeID = core.StringPtr(cartridgeID)
	return options
}

// SetAnnotators : Allow user to set Annotators
func (options *CreateProfileOptions) SetAnnotators(annotators []Annotator) *CreateProfileOptions {
	options.Annotators = annotators
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProfileOptions) SetHeaders(param map[string]string) *CreateProfileOptions {
	options.Headers = param
	return options
}

// DeleteFlowsOptions : The DeleteFlows options.
type DeleteFlowsOptions struct {
	// Flow ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteFlowsOptions : Instantiate DeleteFlowsOptions
func (*AnnotatorForClinicalDataAcdV1) NewDeleteFlowsOptions(id string) *DeleteFlowsOptions {
	return &DeleteFlowsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteFlowsOptions) SetID(id string) *DeleteFlowsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteFlowsOptions) SetHeaders(param map[string]string) *DeleteFlowsOptions {
	options.Headers = param
	return options
}

// DeleteProfileOptions : The DeleteProfile options.
type DeleteProfileOptions struct {
	// Profile ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProfileOptions : Instantiate DeleteProfileOptions
func (*AnnotatorForClinicalDataAcdV1) NewDeleteProfileOptions(id string) *DeleteProfileOptions {
	return &DeleteProfileOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteProfileOptions) SetID(id string) *DeleteProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProfileOptions) SetHeaders(param map[string]string) *DeleteProfileOptions {
	options.Headers = param
	return options
}

// DeleteUserSpecificArtifactsOptions : The DeleteUserSpecificArtifacts options.
type DeleteUserSpecificArtifactsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserSpecificArtifactsOptions : Instantiate DeleteUserSpecificArtifactsOptions
func (*AnnotatorForClinicalDataAcdV1) NewDeleteUserSpecificArtifactsOptions() *DeleteUserSpecificArtifactsOptions {
	return &DeleteUserSpecificArtifactsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserSpecificArtifactsOptions) SetHeaders(param map[string]string) *DeleteUserSpecificArtifactsOptions {
	options.Headers = param
	return options
}

// DeployCartridgeOptions : The DeployCartridge options.
type DeployCartridgeOptions struct {
	// Cartridge archive file.
	ArchiveFile io.ReadCloser `json:"archive_file,omitempty"`

	// The content type of archiveFile.
	ArchiveFileContentType *string `json:"archive_file_content_type,omitempty"`

	// Update resources if they already exist.
	Update *bool `json:"update,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeployCartridgeOptions : Instantiate DeployCartridgeOptions
func (*AnnotatorForClinicalDataAcdV1) NewDeployCartridgeOptions() *DeployCartridgeOptions {
	return &DeployCartridgeOptions{}
}

// SetArchiveFile : Allow user to set ArchiveFile
func (options *DeployCartridgeOptions) SetArchiveFile(archiveFile io.ReadCloser) *DeployCartridgeOptions {
	options.ArchiveFile = archiveFile
	return options
}

// SetArchiveFileContentType : Allow user to set ArchiveFileContentType
func (options *DeployCartridgeOptions) SetArchiveFileContentType(archiveFileContentType string) *DeployCartridgeOptions {
	options.ArchiveFileContentType = core.StringPtr(archiveFileContentType)
	return options
}

// SetUpdate : Allow user to set Update
func (options *DeployCartridgeOptions) SetUpdate(update bool) *DeployCartridgeOptions {
	options.Update = core.BoolPtr(update)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeployCartridgeOptions) SetHeaders(param map[string]string) *DeployCartridgeOptions {
	options.Headers = param
	return options
}

// DeployCartridgeResponse : DeployCartridgeResponse struct
type DeployCartridgeResponse struct {
  // Deploy cartridge HTTP code
	Code *int64 `json:"code,omitempty"`

  // Error response body
	ArtifactResponse []ServiceError `json:"artifactResponse,omitempty"`

  // Deploy cartridge response code
	ArtifactResponseCode *int64 `json:"artifactResponseCode,omitempty"`

	// Deploy cartridge correlation ID
  CorrelationId *string `json:"correlationId,omitempty"`

	// Description
  Description *string `json:"description,omitempty"`

  // List of flow IDs
	Flows []string `json:"flows,omitempty"`

	// List of profile IDs
	Profiles []string `json:"profiles,omitempty"`

	// ID of cartridge
  ID *string `json:"id,omitempty"`

	// Name of cartridge
  Name *string `json:"name,omitempty"`

	// Publish date
  PublishDate *string `json:"publishDate,omitempty"`

	// Start time
	StartTime *string `json:"startTime,omitempty"`

	// Status of deployed cartridge
  Status *string `json:"status,omitempty"`

	// HTTP status code
	StatusCode *int64 `json:"statusCode,omitempty"`

	// Location
  StatusLocation *string `json:"statusLocation,omitempty"`

	// Version
  Version *string `json:"version,omitempty"`
}

// UnmarshalDeployCartridgeResponse unmarshals an instance of DeployCartridgeResponse from the specified map of raw messages.
func UnmarshalDeployCartridgeResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeployCartridgeResponse)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "artifactResponse", &obj.ArtifactResponse, UnmarshalServiceError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "artifactResponseCode", &obj.ArtifactResponseCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "correlationId", &obj.CorrelationId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "flows", &obj.Flows)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profiles", &obj.Profiles)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publishDate", &obj.PublishDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "statusCode", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "statusLocation", &obj.StatusLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//DiagnosisInsight : DiagnosisInsight struct
type DiagnosisInsight struct {
	Usage *Usage `json:"usage,omitempty"`

	SuspectedScore *float64 `json:"suspectedScore,omitempty"`

	SymptomScore *float64 `json:"symptomScore,omitempty"`

	TraumaScore *float64 `json:"traumaScore,omitempty"`

	FamilyHistoryScore *float64 `json:"familyHistoryScore,omitempty"`

	Modifiers *DiagnosisModifier `json:"modifiers,omitempty"`
}

// UnmarshalDiagnosisInsight unmarshals an instance of DiagnosisInsight from the specified map of raw messages.
func UnmarshalDiagnosisInsight(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiagnosisInsight)
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "suspectedScore", &obj.SuspectedScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "symptomScore", &obj.SymptomScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "traumaScore", &obj.TraumaScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "familyHistoryScore", &obj.FamilyHistoryScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "modifiers", &obj.Modifiers, UnmarshalDiagnosisModifier)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DiagnosisModifier : Diagnosis Modifier Insight struct.
type DiagnosisModifier struct {
	AssociatedProcedures []Evidence `json:"associatedProcedures,omitempty"`

	Sites []Site `json:"sites,omitempty"`
}

// UnmarshalDiagnosisModifier unmarshal an instance of DiagnosisModifier from the specified raw message.
func UnmarshalDiagnosisModifier(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DiagnosisModifier)
	err = core.UnmarshalModel(m, "associatedProcedures", &obj.AssociatedProcedures, UnmarshalEvidence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sites", &obj.Sites, UnmarshalSite)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Disambiguation : Disambiguation struct
type Disambiguation struct {
	Validity *string `json:"validity,omitempty"`
}

// UnmarshalConfigurationEntity unmarshals an instance of ConfigurationEntity from the specified map of raw messages.
func UnmarshalDisambiguation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Disambiguation)
	err = core.UnmarshalPrimitive(m, "validity", &obj.Validity)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DoseChangedEvent : DoseChangedEvent insight model struct
type DoseChangedEvent struct {
	Score *float64 `json:"score,omitempty"`

	AllergyScore *float64 `json:"allergyScore,omitempty"`

	Usage *Usage `json:"usage,omitempty"`
}

// UnmarshalDoseChangedEvent unmarshals an instance of DoseChangedEvent from the specific raw message.
func UnmarshalDoseChangedEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DoseChangedEvent)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allergyScore", &obj.AllergyScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EjectionFractionAnnotation : EjectionFractionAnnotation struct
type EjectionFractionAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	IsRange *string `json:"isRange,omitempty"`

	FirstValue *string `json:"firstValue,omitempty"`

	SecondValue *string `json:"secondValue,omitempty"`

	EfAlphabeticValueNormalizedName *string `json:"efAlphabeticnValueNormalizedName,omitempty"`

	EfAlphabeticValueSurfaceForm *string `json:"efAlphabeticValueSurfaceForm,omitempty"`

	EfSuffixNormalizedName *string `json:"efSuffixNormalizedName,omitempty"`

	EfSuffixSurfaceForm *string `json:"efSuffixSurfaceForm,omitempty"`

	EfTermNormalizedName *string `json:"efTermNormalizedName,omitempty"`

	EfTermSurfaceForm *string `json:"efTermSurfaceForm,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`
}

// UnmarshalEjectionFractionAnnotation unmarshals an instance of EjectionFractionAnnotation from the specified map of raw messages.
func UnmarshalEjectionFractionAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EjectionFractionAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "isRange", &obj.IsRange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "firstValue", &obj.FirstValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secondValue", &obj.SecondValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "efAlphabeticValueNormalizedName", &obj.EfAlphabeticValueNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "efAlphabeticValueSurfaceForm", &obj.EfAlphabeticValueSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "efSuffixNormalizedName", &obj.EfAlphabeticValueNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "efSuffixSurfaceForm", &obj.EfAlphabeticValueSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "efTermNormalizedName", &obj.EfAlphabeticValueNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "efTermSurfaceForm", &obj.EfAlphabeticValueSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Entity : Entity struct
type Entity struct {
	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	Uid *int64 `json:"uid,omitempty"`

	Mergeid *int64 `json:"mergeid,omitempty"`
}

// UnmarshalEntity unmarshals an instance of Entity from the specified map of raw messages.
func UnmarshalEntity(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Entity)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.Uid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "mergeid", &obj.Mergeid)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Evidence : Insight Model Data Evidence struct
type Evidence struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`
}

// UnmarshalEvidence unmarshals an instance of Evidence from the specified map of raw messages.
func UnmarshalEvidence(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Evidence)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Flow : Flow struct
type Flow struct {
	Elements []FlowEntry `json:"elements,omitempty"`

	Async *bool `json:"async,omitempty"`
}

// NewFlow : Instantiate Flow (Generic Model Constructor)
func (*AnnotatorForClinicalDataAcdV1) NewFlow(elements []FlowEntry, async *bool) (model *Flow, err error) {
	model = &Flow{
		Elements: elements,
		Async:    async,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFlow unmarshals an instance of Flow from the specified map of raw messages.
func UnmarshalFlow(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Flow)
	err = core.UnmarshalModel(m, "elements", &obj.Elements, UnmarshalFlowEntry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "async", &obj.Async)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FlowEntry : FlowEntry struct
type FlowEntry struct {
	Annotator *Annotator `json:"annotator,omitempty"`
}

// NewFlow : Instantiate Flow (Generic Model Constructor)
func (*AnnotatorForClinicalDataAcdV1) NewFlowEntry(annotator *Annotator) (model *FlowEntry, err error) {
	model = &FlowEntry{
		Annotator: annotator,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFlowEntry unmarshals an instance of FlowEntry from the specified map of raw messages.
func UnmarshalFlowEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FlowEntry)
	err = core.UnmarshalModel(m, "annotator", &obj.Annotator, UnmarshalAnnotator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAnnotatorsByIdOptions : The GetAnnotatorsByID options.
type GetAnnotatorsByIdOptions struct {
	// The ID the Service API was registered under.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAnnotatorsByIdOptions : Instantiate GetAnnotatorsByIdOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetAnnotatorsByIdOptions(id string) *GetAnnotatorsByIdOptions {
	return &GetAnnotatorsByIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetAnnotatorsByIdOptions) SetID(id string) *GetAnnotatorsByIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAnnotatorsByIdOptions) SetHeaders(param map[string]string) *GetAnnotatorsByIdOptions {
	options.Headers = param
	return options
}

// GetAnnotatorsOptions : The GetAnnotators options.
type GetAnnotatorsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAnnotatorsOptions : Instantiate GetAnnotatorsOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetAnnotatorsOptions() *GetAnnotatorsOptions {
	return &GetAnnotatorsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAnnotatorsOptions) SetHeaders(param map[string]string) *GetAnnotatorsOptions {
	options.Headers = param
	return options
}

// GetFlowsByIdOptions : The GetFlowsByID options.
type GetFlowsByIdOptions struct {
	// Flow ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFlowsByIdOptions : Instantiate GetFlowsByIdOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetFlowsByIdOptions(id string) *GetFlowsByIdOptions {
	return &GetFlowsByIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetFlowsByIdOptions) SetID(id string) *GetFlowsByIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetFlowsByIdOptions) SetHeaders(param map[string]string) *GetFlowsByIdOptions {
	options.Headers = param
	return options
}

// GetFlowsOptions : The GetFlows options.
type GetFlowsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFlowsOptions : Instantiate GetFlowsOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetFlowsOptions() *GetFlowsOptions {
	return &GetFlowsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetFlowsOptions) SetHeaders(param map[string]string) *GetFlowsOptions {
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
	GetHealthCheckStatusOptions_Format_Xml  = "xml"
)

// NewGetHealthCheckStatusOptions : Instantiate GetHealthCheckStatusOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetHealthCheckStatusOptions() *GetHealthCheckStatusOptions {
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

// GetProfileOptions : The GetProfile options.
type GetProfilesByIdOptions struct {
	// Profile ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProfileOptions : Instantiate GetProfileOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetProfilesByIdOptions(id string) *GetProfilesByIdOptions {
	return &GetProfilesByIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetProfilesByIdOptions) SetID(id string) *GetProfilesByIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetProfilesByIdOptions) SetHeaders(param map[string]string) *GetProfilesByIdOptions {
	options.Headers = param
	return options
}

// GetProfilesOptions : The GetProfiles options.
type GetProfilesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProfilesOptions : Instantiate GetProfilesOptions
func (*AnnotatorForClinicalDataAcdV1) NewGetProfilesOptions() *GetProfilesOptions {
	return &GetProfilesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetProfilesOptions) SetHeaders(param map[string]string) *GetProfilesOptions {
	options.Headers = param
	return options
}

// InsightModelAlcoholUsage : InsightModel Alcohol Usage struct
type InsightModelAlcoholUsage struct {
        UseScore *float64 `json:"useScore,omitempty"`

        NoneScore *float64 `json:"noneScore,omitempty"`

        UnknownScore *float64 `json:"unknownScore,omitempty"`

        DiscussedScore *float64 `json:"discussedScore,omitempty"`
}

// UnmarshalInsightModelAlcoholUsage unmarshals an instance of InsightModelAlcoholUsage from the specified raw message.
func UnmarshalInsightModelAlcoholUsage(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelAlcoholUsage)
        err = core.UnmarshalPrimitive(m, "useScore", &obj.UseScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "noneScore", &obj.NoneScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "unknownScore", &obj.UnknownScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "discussedScore", &obj.DiscussedScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelAlcoholUseQualifier : InsightModel Alcohol Use Qualifier struct
type InsightModelAlcoholUseQualifier struct {
        LightScore *float64 `json:"lightScore,omitempty"`

        ModerateScore *float64 `json:"moderateScore,omitempty"`

        HeavyScore *float64 `json:"heavyScore,omitempty"`

        AbuseScore *float64 `json:"abuseScore,omitempty"`
}

// UnmarshalInsightModelAlcoholUseQualifier unmarshals an instance of InsightModelAlcoholUseQualifier from the specified raw message.
func UnmarshalInsightModelAlcoholUseQualifier(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelAlcoholUseQualifier)
        err = core.UnmarshalPrimitive(m, "lightScore", &obj.LightScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "moderateScore", &obj.ModerateScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "heavyScore", &obj.HeavyScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "abuseScore", &obj.AbuseScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelAlcoholUseStatus : InsightModel Alcohol Use Status struct
type InsightModelAlcoholUseStatus struct {
        StoppedScore *float64 `json:"stoppedScore,omitempty"`

        NeverScore *float64 `json:"neverScore,omitempty"`
}

// UnmarshalInsightModelAlcoholUseStatus unmarshals an instance of InsightModelAlcoholUseStatus from the specified raw message.
func UnmarshalInsightModelAlcoholUseStatus(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelAlcoholUseStatus)
        err = core.UnmarshalPrimitive(m, "stoppedScore", &obj.StoppedScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "neverScore", &obj.NeverScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelAlcohol : InsightModel Alcohol struct
type InsightModelAlcohol struct {
        Usage *InsightModelAlcoholUsage `json:"usage,omitempty"`

        UseStatus *InsightModelAlcoholUseStatus `json:"useStatus,omitempty"`

        UseQualifier *InsightModelAlcoholUseQualifier `json:"useQualifier,omitempty"`

        ExposureScore *float64 `json:"exposureScore,omitempty"`

        NonPatientScore *float64 `json:"nonPatientScore,omitempty"`

        TreatmentScore *float64 `json:"treatmentScore,omitempty"`
}

// UnmarshalInsightModelAlcohol unmarshals an instance of InsightModelAlcohol from the specified raw message.
func UnmarshalInsightModelAlcohol(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelAlcohol)
        err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalInsightModelAlcoholUsage)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "useStatus", &obj.UseStatus, UnmarshalInsightModelAlcoholUseStatus)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "useQualifier", &obj.UseQualifier, UnmarshalInsightModelAlcoholUseQualifier)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "exposureScore", &obj.ExposureScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "nonPatientScore", &obj.NonPatientScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "treatmentScore", &obj.TreatmentScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelIllicitDrugUseStatus : InsightModel Illicit Drug Use Status struct
type InsightModelIllicitDrugUseStatus struct {
        CurrentScore *float64 `json:"currentScore,omitempty"`
       
        StoppedScore *float64 `json:"stoppedScore,omitempty"`

        NeverScore *float64 `json:"neverScore,omitempty"`

        ComplianceScore *float64 `json:"complianceScore,omitempty"`
}

// UnmarshalInsightModelIllicitDrugUseStatus unmarshals an instance of InsightModelIllicitDrugUseStatus from the specified raw message.
func UnmarshalInsightModelIllicitDrugUseStatus(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelIllicitDrugUseStatus)
        err = core.UnmarshalPrimitive(m, "currentScore", &obj.CurrentScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "stoppedScore", &obj.StoppedScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "neverScore", &obj.NeverScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "complianceScore", &obj.ComplianceScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelIllicitDrugUseQualifier : InsightModel Illicit Drug Use Qualifier struct
type InsightModelIllicitDrugUseQualifier struct {
        LightScore *float64 `json:"lightScore,omitempty"`

        ModerateScore *float64 `json:"moderateScore,omitempty"`

        HeavyScore *float64 `json:"heavyScore,omitempty"`
}

// UnmarshalInsightModelIllicitDrugUseQualifier unmarshals an instance of InsightModelIllicitDrugUseQualifier from the specified raw message.
func UnmarshalInsightModelIllicitDrugUseQualifier(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelIllicitDrugUseQualifier)
        err = core.UnmarshalPrimitive(m, "lightScore", &obj.LightScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "moderateScore", &obj.ModerateScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "heavyScore", &obj.HeavyScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelIllicitDrugUseDimension : InsightModel Illicit Drug Use Dimension struct
type InsightModelIllicitDrugUseDimension struct {
        AbuseScore *float64 `json:"abuseScore,omitempty"`

        MedicalScore *float64 `json:"medicalScore,omitempty"`
}

// UnmarshalInsightModelIllicitDrugUseDimension unmarshals an instance of InsightModelIllicitDrugUseDimension from the specified raw message.
func UnmarshalInsightModelIllicitDrugUseDimension(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelIllicitDrugUseDimension)
        err = core.UnmarshalPrimitive(m, "abuseScore", &obj.AbuseScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "medicalScore", &obj.MedicalScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelIllicitDrugUsage : InsightModel Illicit Drug Usage struct
type InsightModelIllicitDrugUsage struct {
        UseScore *float64 `json:"useScore,omitempty"`
        
        NoneScore *float64 `json:"noneScore,omitempty"`
        
        UnknownScore *float64 `json:"unknownScore,omitempty"`
        
        DiscussedScore *float64 `json:"discussedScore,omitempty"`

        TreatmentScore *float64 `json:"treatmentScore,omitempty"`
}

// UnmarshalInsightModelIllicitDrugUsage unmarshals an instance of InsightModelIllicitDrugUsage from the specified raw message.
func UnmarshalInsightModelIllicitDrugUsage(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelIllicitDrugUsage)
        err = core.UnmarshalPrimitive(m, "useScore", &obj.UseScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "noneScore", &obj.NoneScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "unknownScore", &obj.UnknownScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "discussedScore", &obj.DiscussedScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "treatmentScore", &obj.TreatmentScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelIllicitDrug : InsightModel Illicit Drug struct
type InsightModelIllicitDrug struct {
        Usage *InsightModelIllicitDrugUsage `json:"usage,omitempty"`

        UseStatus *InsightModelIllicitDrugUseStatus `json:"useStatus,omitempty"`

        UseQualifier *InsightModelIllicitDrugUseQualifier `json:"useQualifier,omitempty"`

        UseDimension *InsightModelIllicitDrugUseDimension `json:"useDimension,omitempty"`

        ExposureScore *float64 `json:"exposureScore,omitempty"`

        NonPatientScore *float64 `json:"nonPatientScore,omitempty"`

        TreatmentScore *float64 `json:"treatmentScore,omitempty"`
}

// UnmarshalInsightModelIllicitDrug unmarshals an instance of InsightModelIllicitDrug from the specified raw message.
func UnmarshalInsightModelIllicitDrug(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelIllicitDrug)
        err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalInsightModelIllicitDrugUsage)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "useStatus", &obj.UseStatus, UnmarshalInsightModelIllicitDrugUseStatus)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "useQualifier", &obj.UseQualifier, UnmarshalInsightModelIllicitDrugUseQualifier)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "useDimension", &obj.UseDimension, UnmarshalInsightModelIllicitDrugUseDimension)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "exposureScore", &obj.ExposureScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "nonPatientScore", &obj.NonPatientScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "treatmentScore", &obj.TreatmentScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelTobaccoUseStatus : InsightModel Tobacco Use Status struct
type InsightModelTobaccoUseStatus struct {
        CurrentScore *float64 `json:"currentScore,omitempty"`

        StoppedScore *float64 `json:"stoppedScore,omitempty"`

        NeverScore *float64 `json:"neverScore,omitempty"`
}

// UnmarshalInsightModelTobaccoUseStatus unmarshals an instance of InsightModelTobaccoUseStatus from the specified raw message.
func UnmarshalInsightModelTobaccoUseStatus(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelTobaccoUseStatus)
        err = core.UnmarshalPrimitive(m, "currentScore", &obj.CurrentScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "stoppedScore", &obj.StoppedScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "neverScore", &obj.NeverScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelTobaccoUsage : InsightModel Tobacco Usage struct
type InsightModelTobaccoUsage struct {
        UseScore *float64 `json:"useScore,omitempty"`

        NoneScore *float64 `json:"noneScore,omitempty"`

        UnknownScore *float64 `json:"unknownScore,omitempty"`

        DiscussedScore *float64 `json:"discussedScore,omitempty"`
}

// UnmarshalInsightModelTobaccoUsage unmarshals an instance of InsightModelTobaccoUsage from the specified raw message.
func UnmarshalInsightModelTobaccoUsage(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelTobaccoUsage)
        err = core.UnmarshalPrimitive(m, "useScore", &obj.UseScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "noneScore", &obj.NoneScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "unknownScore", &obj.UnknownScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "discussedScore", &obj.DiscussedScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelTobacco : InsightModel Tobacco struct
type InsightModelTobacco struct {
        Usage *InsightModelTobaccoUsage `json:"usage,omitempty"`

        UseStatus *InsightModelTobaccoUseStatus `json:"useStatus,omitempty"`

        ExposureScore *float64 `json:"exposureScore,omitempty"`

        FamilyHistoryScore *float64 `json:"familyHistoryScore,omitempty"`

        NonPatientScore *float64 `json:"nonPatientScore,omitempty"`

        TreatmentScore *float64 `json:"treatmentScore,omitempty"`
}

// UnmarshalInsightModelTobacco unmarshals an instance of InsightModelTobacco from the specified raw message.
func UnmarshalInsightModelTobacco(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelTobacco)
        err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalInsightModelTobaccoUsage)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "useStatus", &obj.UseStatus, UnmarshalInsightModelTobaccoUseStatus)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "exposureScore", &obj.ExposureScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "familyHistoryScore", &obj.ExposureScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "nonPatientScore", &obj.NonPatientScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "treatmentScore", &obj.TreatmentScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelSubstanceAbuseTreatment : InsightModel Substance Abuse Treatment struct
type InsightModelSubstanceAbuseTreatment struct {
        DiscussedScore *float64 `json:"discussedScore,omitempty"`

        ComplianceScore *float64 `json:"complianceScore,omitempty"`
}

// UnmarshalInsightModelSubstanceAbuseTreatment unmarshals an instance of InsightModelSubstanceAbuseTreatment from the specified raw message.
func UnmarshalInsightModelSubstanceAbuseTreatment(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelSubstanceAbuseTreatment)
        err = core.UnmarshalPrimitive(m, "discussedScore", &obj.DiscussedScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "complianceScore", &obj.ComplianceScore)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModelSubstanceAbuse : InsightModel Substance Abuse struct
type InsightModelSubstanceAbuse struct {
        TreatmentScore *float64 `json:"treatmentScore,omitempty"`

        NonPatientScore *float64 `json:"nonPatientScore,omitempty"`

        Treatment *InsightModelSubstanceAbuseTreatment `json:"treatment,omitempty"`
}

// UnmarshalInsightModelSubstanceAbuse unmarshals an instance of InsightModelSubstanceAbuse from the specified raw message.
func UnmarshalInsightModelSubstanceAbuse(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(InsightModelSubstanceAbuse)
        err = core.UnmarshalPrimitive(m, "treatmentScore", &obj.TreatmentScore)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "nonPatientScore", &obj.NonPatientScore)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "treatment", &obj.Treatment, UnmarshalInsightModelSubstanceAbuseTreatment)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// InsightModel : InsightModel struct
type InsightModel struct {
	Procedure *ProcedureInsight `json:"procedure,omitempty"`

	Diagnosis *DiagnosisInsight `json:"diagnosis,omitempty"`

	Medication *MedicationInsight `json:"medication,omitempty"`

	Normality *NormalityInsight `json:"normality,omitempty"`

        Tobacco *InsightModelTobacco `json:"tobacco,omitempty"`

        Alcohol *InsightModelAlcohol `json:"alcohol,omitempty"`

        IllicitDrug *InsightModelIllicitDrug `json:"illicitDrug,omitempty"`

        SubstanceAbuse *InsightModelSubstanceAbuse `json:"substance,omitempty"`
}

// UnmarshalInsightModel unmarshals an instance of InsightModel from the specified raw message.
func UnmarshalInsightModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InsightModel)
	err = core.UnmarshalModel(m, "procedure", &obj.Procedure, UnmarshalProcedureInsight)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "diagnosis", &obj.Diagnosis, UnmarshalDiagnosisInsight)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "medication", &obj.Medication, UnmarshalMedicationInsight)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "normality", &obj.Normality, UnmarshalNormalityInsight)
	if err != nil {
		return
	}
        err = core.UnmarshalModel(m, "tobacco", &obj.Tobacco, UnmarshalInsightModelTobacco)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "alcohol", &obj.Alcohol, UnmarshalInsightModelAlcohol)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "tobacco", &obj.Tobacco, UnmarshalInsightModelTobacco)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "illicitDrug", &obj.IllicitDrug, UnmarshalInsightModelIllicitDrug)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "substance", &obj.SubstanceAbuse, UnmarshalInsightModelSubstanceAbuse)
        if err != nil {
                return
        }
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LabValueAnnotation : LabValueAnnotation struct
type LabValueAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	DateInMilliseconds *string `json:"dateInMilliseconds,omitempty"`

	LabValue *string `json:"labValue,omitempty"`

	LabTypeNormalizedName *string `json:"labTypeNormalizedName,omitempty"`

	LabTypeSurfaceForm *string `json:"labTypeSurfaceForm,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`
}

// UnmarshalLabValueAnnotation unmarshals an instance of LabValueAnnotation from the specified map of raw messages.
func UnmarshalLabValueAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LabValueAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dateInMilliseconds", &obj.DateInMilliseconds)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labValue", &obj.LabValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labValuenNormalizedName", &obj.LabTypeNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labValueSurfaceForm", &obj.LabTypeSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MedicationAnnotation : MedicationAnnotation struct
type MedicationAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Drug []map[string]interface{} `json:"drug,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`

	Temporal []Temporal `json:"temporal,omitempty"`

        DisambiguationData *Disambiguation `json:"disambiguationData,omitempty"`
}

// UnmarshalMedicationAnnotation unmarshals an instance of MedicationAnnotation from the specified map of raw messages.
func UnmarshalMedicationAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MedicationAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "drug", &obj.Drug)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "temporal", &obj.Temporal, UnmarshalTemporal)
	if err != nil {
		return
	}
        err = core.UnmarshalModel(m, "disambiguationData", &obj.DisambiguationData, UnmarshalDisambiguation)
        if err != nil {
                return
        }
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MedicationInsight : MedicationInsight struct
type MedicationInsight struct {
	Usage *Usage `json:"usage,omitempty"`

	StartedEvent *StartedEvent `json:"startedEvent,omitempty"`

	StoppedEvent *StoppedEvent `json:"stoppedEvent,omitempty"`

	DoseChangedEvent *DoseChangedEvent `json:"doseChangedEvent,omitempty"`

	AdverseEvent *AdverseEvent `json:"adverseEvent,omitempty"`
}

// UnmarshalMedicationInsight unmarshals an instance of MedicationInsight from the specific raw message.
func UnmarshalMedicationInsight(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MedicationInsight)
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "startedEvent", &obj.StartedEvent, UnmarshalStartedEvent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "stoppedEvent", &obj.StoppedEvent, UnmarshalStoppedEvent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "doseChangedEvent", &obj.DoseChangedEvent, UnmarshalDoseChangedEvent)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "adverseEvent", &obj.AdverseEvent, UnmarshalAdverseEvent)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NegatedSpan : NegatedSpan struct
type NegatedSpan struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Trigger []map[string]interface{} `json:"trigger,omitempty"`
}

// UnmarshalNegatedSpan unmarshals an instance of NegatedSpan from the specified map of raw messages.
func UnmarshalNegatedSpan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NegatedSpan)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dateInMilliseconds", &obj.Trigger)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NluEntities : NluEntities struct
type NluEntities struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	Type *string `json:"type,omitempty"`

	Source *string `json:"source,omitempty"`

	Relevance *float64 `json:"relevance,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`
}

// UnmarshalNluEntities unmarshals an instance of NluEntities from the specified map of raw messages.
func UnmarshalNluEntities(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NluEntities)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relevance", &obj.Relevance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Node : Node struct
type NodeAnnotation struct {
	Entity *NodeEntityAnnotation `json:"entity,omitempty"`
}

// UnmarshalNode unmarshals an instance of Node from the specified map of raw messages.
func UnmarshalNodeAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NodeAnnotation)
	err = core.UnmarshalModel(m, "entity", &obj.Entity, UnmarshalNodeEntityAnnotation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NodeEntity : NodeEntity struct
type NodeEntityAnnotation struct {
	UID *int64 `json:"uid,omitempty"`
}

// UnmarshalNodeEntity unmarshals an instance of NodeEntity from the specified map of raw messages.
func UnmarshalNodeEntityAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NodeEntityAnnotation)
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NormalityInsight : NormalityInsight struct.
type NormalityInsight struct {
	NormalityUsage *NormalityUsage `json:"usage,omitempty"`

	Evidence []Evidence `json:"evidence,omitempty"`

	DirectlyAffectedScore *float64 `json:"directlyAffectedScore,omitempty"`
}

// UnmarshalNormalityInsight unmarshal an instance of NormalityInsight from the specified raw message.
func UnmarshalNormalityInsight(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NormalityInsight)
	err = core.UnmarshalModel(m, "usage", &obj.NormalityUsage, UnmarshalNormalityUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "evidence", &obj.Evidence, UnmarshalEvidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "directlyAffectedScore", &obj.DirectlyAffectedScore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NormalityUsage : NormalityUsage insight model struct
type NormalityUsage struct {
	NormalScore *float64 `json:"normalScore,omitempty"`

	AbnormalScore *float64 `json:"abnormalScore,omitempty"`

	UnknownScore *float64 `json:"unknownScore,omitempty"`

	NonFindingScore *float64 `json:"nonFindingScore,omitempty"`
}

// UnmarshalNormalityUsage unmarshals an instance of NormalityUsage from the specific raw message.
func UnmarshalNormalityUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NormalityUsage)
	err = core.UnmarshalPrimitive(m, "normalScore", &obj.NormalScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "abnormalScore", &obj.AbnormalScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unknownScore", &obj.UnknownScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "nonFindingScoreScore", &obj.NonFindingScore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProcedureAnnotation : ProcedureAnnotation struct
type ProcedureAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	DateInMilliseconds *string `json:"dateInMilliseconds,omitempty"`

	DisambiguationData *Disambiguation `json:"disambiguationData,omitempty"`

	ProcedureNormalizedName *string `json:"procedureNormalizedName,omitempty"`

	ProcedureSurfaceForm *string `json:"procedureSurfaceForm,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	SnomedConceptId *string `json:"snomedConceptId,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`

	Temporal []Temporal `json:"temporal,omitempty"`
}

// UnmarshalProcedcureAnnotation unmarshals an instance of ProcedureAnnotation from the specified map of raw messages.
func UnmarshalProcedureAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProcedureAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dateInMilliseconds", &obj.DateInMilliseconds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguationData", &obj.DisambiguationData, UnmarshalDisambiguation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "procedureNormalizedName", &obj.ProcedureNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "procedureSurfaceForm", &obj.ProcedureSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "snomedConceptId", &obj.SnomedConceptId)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "temporal", &obj.Temporal, UnmarshalTemporal)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProcedureInsight : ProcedureInsight struct.
type ProcedureInsight struct {
	Task *Task `json:"task,omitempty"`

	Type *Type `json:"type,omitempty"`

	Usage *Usage `json:"usage,omitempty"`

	Modifiers *ProcedureModifier `json:"modifiers,omitempty"`
}

// UnmarshalProcedureInsight unmarshal an instance of ProcedureInsight from the specified raw message.
func UnmarshalProcedureInsight(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProcedureInsight)
	err = core.UnmarshalModel(m, "task", &obj.Task, UnmarshalTask)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "type", &obj.Type, UnmarshalType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "modifiers", &obj.Modifiers, UnmarshalProcedureModifier)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProcedureModifier : Procedure Modifier Insight struct.
type ProcedureModifier struct {
	AssociatedDiagnosis []Evidence `json:"associatedDiagnosis,omitempty"`

	Sites []Site `json:"sites,omitempty"`
}

// UnmarshalProcedureModifier unmarshal an instance of ProcedureModifier from the specified raw message.
func UnmarshalProcedureModifier(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProcedureModifier)
	err = core.UnmarshalModel(m, "associatedDiagnosis", &obj.AssociatedDiagnosis, UnmarshalEvidence)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sites", &obj.Sites, UnmarshalSite)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Reference : Reference struct.
type Reference struct {
        UID *int64 `json:"uid,omitempty"`

        SelectionLabel *string `json:"selectionLabel,omitempty"`

        ValueIndex *int64 `json:"valueIndex,omitempty"`
}

// UnmarshalReference unmarshal an instance of Reference from the specified raw message.
func UnmarshalReference(m map[string]json.RawMessage, result interface{}) (err error) {
        obj := new(Reference)
        err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "selectionLabel", &obj.SelectionLabel)
        if err != nil {
                return
        }
        err = core.UnmarshalPrimitive(m, "valueIndex", &obj.ValueIndex)
        if err != nil {
                return
        }
        reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
        return
}

// Relations : Relations struct.
type Relations struct {
	Source *string `json:"source,omitempty"`

	Score *float64 `json:"score,omitempty"`

	Nodes []NodeAnnotation `json:"nodes,omitempty"`

	Type *string `json:"type,omitempty"`
}

// UnmarshalRelations unmarshal an instance of Relations from the specified raw message.
func UnmarshalRelations(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Relations)
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nodes", &obj.Nodes, UnmarshalNodeAnnotation)
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

// RunPipelineOptions : The RunPipeline options.
type RunPipelineOptions struct {
	Unstructured []UnstructuredContainer `json:"unstructured" validate:"required"`

	AnnotatorFlows []AnnotatorFlow `json:"annotatorFlows" validate:"required"`

	// If true, any ReplaceTextChange annotations will be left in the container and the modified text before restoring to
	// original form will stored in the metadata that is returned.  Otherwise, these annotations and modified text will be
	// removed from the container.
	DebugTextRestore *bool `json:"debug_text_restore,omitempty"`

	// Set this to true to show the analyzed text in the response.
	ReturnAnalyzedText *bool `json:"return_analyzed_text,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunPipelineOptions : Instantiate RunPipelineOptions
func (*AnnotatorForClinicalDataAcdV1) NewRunPipelineOptions() *RunPipelineOptions {
	return &RunPipelineOptions{}
}

// SetUnstructured : Allow user to set Unstructured
func (options *RunPipelineOptions) SetUnstructured(unstructured []UnstructuredContainer) *RunPipelineOptions {
	options.Unstructured = unstructured
	return options
}

// SetAnnotatorFlows : Allow user to set AnnotatorFlows
func (options *RunPipelineOptions) SetAnnotatorFlows(annotatorFlows []AnnotatorFlow) *RunPipelineOptions {
	options.AnnotatorFlows = annotatorFlows
	return options
}

// SetDebugTextRestore : Allow user to set DebugTextRestore
func (options *RunPipelineOptions) SetDebugTextRestore(debugTextRestore bool) *RunPipelineOptions {
	options.DebugTextRestore = core.BoolPtr(debugTextRestore)
	return options
}

// SetReturnAnalyzedText : Allow user to set ReturnAnalyzedText
func (options *RunPipelineOptions) SetReturnAnalyzedText(returnAnalyzedText bool) *RunPipelineOptions {
	options.ReturnAnalyzedText = core.BoolPtr(returnAnalyzedText)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RunPipelineOptions) SetHeaders(param map[string]string) *RunPipelineOptions {
	options.Headers = param
	return options
}

// RunPipelineWithFlowOptions : The RunPipelineWithFlow options.
type RunPipelineWithFlowOptions struct {
	// flow identifier.
	FlowID *string `json:"flow_id" validate:"required"`

	// Set this to true to show the analyzed text in the response.
	ReturnAnalyzedText *bool `json:"return_analyzed_text" validate:"required"`

	// Input request data in TEXT or JSON format .
	Body *string `json:"body,omitempty"`

	// Input request data in TEXT or JSON format .
	AnalyticFlowBeanInput *AnalyticFlowBeanInput `json:"AnalyticFlowBeanInput,omitempty"`

	// The type of the input. A character encoding can be specified by including a `charset` parameter. For example,
	// 'text/plain;charset=utf-8'.
	ContentType *string `json:"Content-Type,omitempty"`

	// If true, any ReplaceTextChange annotations will be left in the container and the modified text before restoring to
	// original form will be returned in the metadata.  Otherwise, these annotations and modified text will be removed from
	// the container.
	DebugTextRestore *bool `json:"debug_text_restore,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRunPipelineWithFlowOptions : Instantiate RunPipelineWithFlowOptions
func (*AnnotatorForClinicalDataAcdV1) NewRunPipelineWithFlowOptions(flowID string, returnAnalyzedText bool) *RunPipelineWithFlowOptions {
	return &RunPipelineWithFlowOptions{
		FlowID:             core.StringPtr(flowID),
		ReturnAnalyzedText: core.BoolPtr(returnAnalyzedText),
	}
}

// SetFlowID : Allow user to set FlowID
func (options *RunPipelineWithFlowOptions) SetFlowID(flowID string) *RunPipelineWithFlowOptions {
	options.FlowID = core.StringPtr(flowID)
	return options
}

// SetReturnAnalyzedText : Allow user to set ReturnAnalyzedText
func (options *RunPipelineWithFlowOptions) SetReturnAnalyzedText(returnAnalyzedText bool) *RunPipelineWithFlowOptions {
	options.ReturnAnalyzedText = core.BoolPtr(returnAnalyzedText)
	return options
}

// SetAnalyticFlowBeanInput : Allow user to set AnalyticFlowBeanInput
func (options *RunPipelineWithFlowOptions) SetAnalyticFlowBeanInput(analyticFlowBeanInput *AnalyticFlowBeanInput) *RunPipelineWithFlowOptions {
	options.AnalyticFlowBeanInput = analyticFlowBeanInput
	return options
}

// SetBody : Allow user to set Body
func (options *RunPipelineWithFlowOptions) SetBody(body string) *RunPipelineWithFlowOptions {
	options.Body = core.StringPtr(body)
	return options
}

// SetContentType : Allow user to set ContentType
func (options *RunPipelineWithFlowOptions) SetContentType(contentType string) *RunPipelineWithFlowOptions {
	options.ContentType = core.StringPtr(contentType)
	return options
}

// SetDebugTextRestore : Allow user to set DebugTextRestore
func (options *RunPipelineWithFlowOptions) SetDebugTextRestore(debugTextRestore bool) *RunPipelineWithFlowOptions {
	options.DebugTextRestore = core.BoolPtr(debugTextRestore)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RunPipelineWithFlowOptions) SetHeaders(param map[string]string) *RunPipelineWithFlowOptions {
	options.Headers = param
	return options
}

// SectionAnnotation : SectionAnnotation struct
type SectionAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	SectionType *string `json:"sectionType,omitempty"`

	Type *string `json:"type,omitempty"`

	Trigger *SectionTrigger `json:"trigger,omitempty"`
}

// UnmarshalSectionAnnotation unmarshals an instance of SectionAnnotation from the specified map of raw messages.
func UnmarshalSectionAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SectionAnnotation)
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
	err = core.UnmarshalPrimitive(m, "sectionType", &obj.SectionType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "trigger", &obj.Trigger, UnmarshalSectionTrigger)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SectionTrigger : SectionTrigger struct
type SectionTrigger struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	Source *string `json:"source,omitempty"`

	Type *string `json:"type,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`
}

// UnmarshalSectionTrigger unmarshals an instance of SectionTrigger from the specified map of raw messages.
func UnmarshalSectionTrigger(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SectionTrigger)
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
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Site : Insight Model Data Site struct
type Site struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	Type *string `json:"type,omitempty"`
}

// UnmarshalSite unmarshals an instance of Site from the specified map of raw messages.
func UnmarshalSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Site)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SmokingAnnotation : SmokingAnnotation struct
type SmokingAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	Modality *string `json:"modality,omitempty"`

	Current *string `json:"current,omitempty"`

	Participation *string `json:"participation,omitempty"`

	SmokeTermNormalizedName *string `json:"smokeTermNormalizedName,omitempty"`

	SmokeTermSurfaceForm *string `json:"smokeTermSurfaceForm,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`
}

// UnmarshalSmokingAnnotation unmarshals an instance of SmokingAnnotation from the specified map of raw messages.
func UnmarshalSmokingAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SmokingAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modality", &obj.Modality)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current", &obj.Current)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "participation", &obj.Participation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "smokeTermNormalizedName", &obj.SmokeTermNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "smoketermSurfaceForm", &obj.SmokeTermSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpellCorrectedText : SpellCorrectedText struct
type SpellCorrectedText struct {
	CorrectedText string `json:"correctedText,omitempty"`

	DebugText string `json:"debugText,omitempty"`
}

// UnmarshalSpellCorrectedText unmarshals an instance of SpellCorrectedText from the specified map of raw messages.
func UnmarshalSpellCorrectedText(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpellCorrectedText)
	err = core.UnmarshalPrimitive(m, "correctedText", &obj.CorrectedText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "debugText", &obj.DebugText)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpellingCorrection : SpellingCorrection struct
type SpellingCorrection struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

// UnmarshalSpellingCorrection unmarshals an instance of SpellingCorrection from the specified map of raw messages.
func UnmarshalSpellingCorrection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpellingCorrection)
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
	err = core.UnmarshalModel(m, "suggestions", &obj.Suggestions, UnmarshalSuggestion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StartedEvent : StartedEvent insight model struct
type StartedEvent struct {
	Score *float64 `json:"score,omitempty"`

	AllergyScore *float64 `json:"allergyScore,omitempty"`

	Usage *Usage `json:"usage,omitempty"`
}

// UnmarshalStartedEvent unmarshals an instance of StartedEvent from the specific raw message.
func UnmarshalStartedEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StartedEvent)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allergyScore", &obj.AllergyScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StoppedEvent : StoppedEvent insight model struct
type StoppedEvent struct {
	Score *float64 `json:"score,omitempty"`

	AllergyScore *float64 `json:"allergyScore,omitempty"`

	Usage *Usage `json:"usage,omitempty"`
}

// UnmarshalStoppedEvent unmarshals an instance of StoppedEvent from the specific raw message.
func UnmarshalStoppedEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StoppedEvent)
	err = core.UnmarshalPrimitive(m, "score", &obj.Score)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allergyScore", &obj.AllergyScore)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Suggestion : Suggestion struct
type Suggestion struct {
	Applied *bool `json:"applied,omitempty"`

	Confidence *float64 `json:"confidence,omitempty"`

	SemTypes []string `json:"semTypes.omitempty"`

	Text *string `json:"text,omitempty"`
}

// UnmarshalSuggesation unmarshals an instance of Suggestion from the specified map of raw messages.
func UnmarshalSuggestion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Suggestion)
	err = core.UnmarshalPrimitive(m, "applied", &obj.Applied)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "confidence", &obj.Confidence)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "semTypes", &obj.SemTypes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SymptomDiseaseAnnotation : SymptomDiseaseAnnotation struct
type SymptomDiseaseAnnotation struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	UID *int64 `json:"uid,omitempty"`

	Hypothetical *bool `json:"hypothetical,omitempty"`

	Negated *bool `json:"negated,omitempty"`

	DateInMilliseconds *string `json:"dateInMilliseconds,omitempty"`

	DisambiguationData *Disambiguation `json:"disambiguationData,omitempty"`

	CcsCode *string `json:"ccsCode,omitempty"`

	Cui *string `json:"cui,omitempty"`

	HccCode *string `json:"hccCode,omitempty"`

	Icd9Code *string `json:"icd9Code,omitempty"`

	Icd10Code *string `json:"icd10Code,omitempty"`

	Modality *string `json:"modality,omitempty"`

	SnomedConceptId *string `json:"snomedConceptId,omitempty"`

	SymptomDiseaseNormalizedName *string `json:"symptomDieaseNormalizedName,omitempty"`

	SymptomDiseaseSurfaceForm *string `json:"symptomDieaseSurfaceForm,omitempty"`

	SectionNormalizedName *string `json:"sectionNormalizedName,omitempty"`

	SectionSurfaceForm *string `json:"sectionSurfaceForm,omitempty"`

	InsightModelData *InsightModel `json:"insightModelData,omitempty"`

	Temporal []Temporal `json:"temporal,omitempty"`
}

// UnmarshalSymptomDiseaseAnnotation unmarshals an instance of SymptomDiseaseAnnotaiton from the specified map of raw messages.
func UnmarshalSymptomDiseaseAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SymptomDiseaseAnnotation)
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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.UID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hypothetical", &obj.Hypothetical)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "negated", &obj.Negated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dateInMilliseconds", &obj.DateInMilliseconds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "disambiguationData", &obj.DisambiguationData, UnmarshalDisambiguation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ccsCode", &obj.CcsCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cui", &obj.Cui)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hccCode", &obj.HccCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icd9Code", &obj.Icd9Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icd10Code", &obj.Icd10Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modality", &obj.Modality)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "snomedConceptId", &obj.SnomedConceptId)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "symptomDiseaseNormalizedName", &obj.SymptomDiseaseNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "symptomDieaseSurfaceForm", &obj.SymptomDiseaseSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionNormalizedName", &obj.SectionNormalizedName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sectionSurfaceForm", &obj.SectionSurfaceForm)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "insightModelData", &obj.InsightModelData, UnmarshalInsightModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "temporal", &obj.Temporal, UnmarshalTemporal)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Task : Task insight model struct
type Task struct {
	TherapeuticScore *float64 `json:"therapeuticScore,omitempty"`

	DiagnosticScore *float64 `json:"diagnosticScore,omitempty"`

	LabTestScore *float64 `json:"labTestScore,omitempty"`

	SurgicalTaskScore *float64 `json:"surgicalTaskScore,omitempty"`

	ClinicalAssessmenttScore *float64 `json:"clinicalAssessmentScore,omitempty"`
}

// UnmarshalTask unmarshals an instance of Task from the specific raw message.
func UnmarshalTask(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Task)
	err = core.UnmarshalPrimitive(m, "therapeuticScore", &obj.TherapeuticScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "diagnosticScore", &obj.DiagnosticScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labTestScore", &obj.LabTestScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "surgicalTaskScore", &obj.SurgicalTaskScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "clinicalAssessmentScore", &obj.ClinicalAssessmenttScore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Temporal : Temporal struct
type Temporal struct {
	Begin *int64 `json:"begin,omitempty"`

	End *int64 `json:"end,omitempty"`

	CoveredText *string `json:"coveredText,omitempty"`

	TemporalType map[string]float64 `json:"temporalType,omitempty"`

	RelationTypes map[string]float64 `json:"relationTypes,omitempty"`
}

// UnmarshalTemporal unmarshals an instance of Temporal from the specified map of raw messages.
func UnmarshalTemporal(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Temporal)
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
	err = core.UnmarshalPrimitive(m, "temporalType", &obj.TemporalType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "relationTypes", &obj.RelationTypes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Type : Type insight model struct
type Type struct {
	DeviceScore *float64 `json:"deviceScore,omitempty"`

	MaterialScore *float64 `json:"materialScore,omitempty"`

	MedicationScore *float64 `json:"medicationScore,omitempty"`

	ProcedureScore *float64 `json:"procedureScore,omitempty"`

	ConditionManagementScore *float64 `json:"conditionManagementScore,omitempty"`
}

// UnmarshalType unmarshals an instance of Type from the specific raw message.
func UnmarshalType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Type)
	err = core.UnmarshalPrimitive(m, "deviceScore", &obj.DeviceScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "materialScore", &obj.MaterialScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "medicationScore", &obj.MedicationScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "procedureScore", &obj.ProcedureScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "conditionManagementScore", &obj.ConditionManagementScore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnstructuredContainer : UnstructuredContainer struct
type UnstructuredContainer struct {
	Text *string `json:"text,omitempty"`

	ID *string `json:"id,omitempty"`

	Type *string `json:"type,omitempty"`

	Data *ContainerAnnotation `json:"data,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`

	Uid *int64 `json:"uid,omitempty"`
}

// NewUnstructuredContainer : Instantiate UnstructuredContainer
func (*AnnotatorForClinicalDataAcdV1) NewUnstructuredContainer() *UnstructuredContainer {
	return &UnstructuredContainer{}
}

// UnmarshalUnstructuredContainer unmarshals an instance of UnstructuredContainer from the specified map of raw messages.
func UnmarshalUnstructuredContainer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UnstructuredContainer)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "data", &obj.Data, UnmarshalContainerAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uid", &obj.Uid)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetText : sets the container text
func (container *UnstructuredContainer) SetText(text string) {
	container.Text = core.StringPtr(text)
}

// SetData : sets the container data
func (container *UnstructuredContainer) SetData(data *ContainerAnnotation) {
	container.Data = data
}

// ContainerAnnotation : ContainerAnnotation struct
type ContainerAnnotation struct {

	AllergyInd []Annotation `json:"allergyInd,omitempty"`

	AllergyMedicationInd []Annotation  `json:"AllergyMedicationInd,omitempty"`

	AttributeValues []AttributeValueAnnotation `json:"attributeValues,omitempty"`

	BathingAssistanceInd []AssistanceAnnotation `json:"BathingAssistanceInd,omitempty"`

	IcaCancerDiagnosisInd []CancerAnnotation `json:"IcaCancerDiagnosisInd,omitempty"`

	Concepts []Concept `json:"concepts,omitempty"`

	ConceptValues []ConceptValue `json:"conceptValues,omitempty"`

	EjectionFractionInd []EjectionFractionAnnotation `json:"ejectionFractionInd,omitempty"`

	HypotheticalSpans []Annotation `json:"hypotheticalSpans,omitempty"`

	LabValueInd []LabValueAnnotation `json:"LabValueInd,omitempty"`

	MedicationInd []MedicationAnnotation `json:"MedicationInd,omitempty"`

	EmailAddressInd []Annotation `json:"EmailAddressInd,omitempty"`

	PersonInd []Annotation `json:"PersonInd,omitempty"`

	US_PhoneNumberInd []Annotation `json:"US_PhoneNumberInd,omitempty"`

	MedicalInstitutionInd []Annotation `json:"MedicalInstitutionInd,omitempty"`

	OrganizationInd []Annotation `json:"OrganizationInd,omitempty"`

	NegatedSpans []NegatedSpan `json:"negatedSpans,omitempty"`

	ProcedureInd []ProcedureAnnotation `json:"ProcedureInd,omitempty"`

	SeeingAssistanceInd []AssistanceAnnotation `json:"SeeingAssistanceInd,omitempty"`

	SmokingInd []SmokingAnnotation `json:"SmokingInd,omitempty"`

	SymptomDiseaseInd []SymptomDiseaseAnnotation `json:"SymptomDiseaseInd,omitempty"`

	ToiletingAssistanceInd []AssistanceAnnotation `json:"ToiletingAssistanceInd,omitempty"`

	WalkingAssistanceInd []AssistanceAnnotation `json:"WalkingAssistanceInd,omitempty"`

	Sections []SectionAnnotation `json:"sections,omitempty"`

	NluEntities []NluEntities `json:"nluEntities,omitempty"`

	Relations []Relations `json:"relations,omitempty"`

	SpellingCorrections []SpellingCorrection `json:"spellingCorrections,omitempty"`

	SpellCorrectedText []SpellCorrectedText `json:"spellCorrectedText,omitempty"`

	TemporalSpans []Temporal `json:"temporalSpans,omitempty"`

        Lines []Annotation `json:"lines,omitempty"`

        Sentences []Annotation `json:"sentences,omitempty"`

        Paragraphs []Annotation `json:"paragraphs,omitempty"`
}

// NewContainerAnnotation : Instantiate ContainerAnnotation
func (*AnnotatorForClinicalDataAcdV1) NewContainerAnnotation() *ContainerAnnotation {
	return &ContainerAnnotation{}
}

// UnmarshalContainerAnnotation unmarshals an instance of ContainerAnnotation from the specified map of raw messages.
func UnmarshalContainerAnnotation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ContainerAnnotation)
	err = core.UnmarshalModel(m, "AllergyInd", &obj.AllergyInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "AllergyMedicationInd", &obj.AllergyMedicationInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attributeValues", &obj.AttributeValues, UnmarshalAttributeValueAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "BathingAssistanceInd", &obj.BathingAssistanceInd, UnmarshalAssistanceAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "IcaCancerDiagnosisInd", &obj.IcaCancerDiagnosisInd, UnmarshalCancerAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "concepts", &obj.Concepts, UnmarshalConcept)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "conceptValues", &obj.ConceptValues, UnmarshalConceptValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "EjectionFractionInd", &obj.EjectionFractionInd, UnmarshalEjectionFractionAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "hypotheticalSpans", &obj.HypotheticalSpans, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "LabValueInd", &obj.LabValueInd, UnmarshalLabValueAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "MedicationInd", &obj.MedicationInd, UnmarshalMedicationAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "EmailAddressInd", &obj.EmailAddressInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "PersonInd", &obj.PersonInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "US_PhoneNumberInd", &obj.US_PhoneNumberInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "MedicalInstitutionInd", &obj.MedicalInstitutionInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "OrganizationInd", &obj.OrganizationInd, UnmarshalAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "negatedSpans", &obj.NegatedSpans, UnmarshalNegatedSpan)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ProcedureInd", &obj.ProcedureInd, UnmarshalProcedureAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "SeeingAssistanceInd", &obj.SeeingAssistanceInd, UnmarshalAssistanceAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "SmokingInd", &obj.SmokingInd, UnmarshalSmokingAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "SymptomDiseaseInd", &obj.SymptomDiseaseInd, UnmarshalSymptomDiseaseAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ToiletingAssistanceInd", &obj.ToiletingAssistanceInd, UnmarshalAssistanceAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "WalkingAssistanceInd", &obj.WalkingAssistanceInd, UnmarshalAssistanceAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sections", &obj.Sections, UnmarshalSectionAnnotation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nluEntities", &obj.NluEntities, UnmarshalNluEntities)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "relations", &obj.Relations, UnmarshalRelations)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spellingCorrections", &obj.SpellingCorrections, UnmarshalSpellingCorrection)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spellCorrectedText", &obj.SpellCorrectedText, UnmarshalSpellCorrectedText)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "temporalSpans", &obj.TemporalSpans, UnmarshalTemporal)
	if err != nil {
		return
	}
        err = core.UnmarshalModel(m, "Lines", &obj.Lines, UnmarshalAnnotation)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "Sentences", &obj.Sentences, UnmarshalAnnotation)
        if err != nil {
                return
        }
        err = core.UnmarshalModel(m, "Paragraphs", &obj.Paragraphs, UnmarshalAnnotation)
        if err != nil {
                return
        }
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateFlowsOptions : The UpdateFlows options.
type UpdateFlowsOptions struct {
	// Flow ID.
	ID *string `json:"id" validate:"required"`

	NewID *string `json:"new_id,omitempty"`

	NewName *string `json:"new_name,omitempty"`

	NewDescription *string `json:"new_description,omitempty"`

	NewPublishedDate *string `json:"new_publishedDate,omitempty"`

	NewPublish *bool `json:"new_publish,omitempty"`

	NewVersion *string `json:"new_version,omitempty"`

	NewCartridgeID *string `json:"new_cartridgeId,omitempty"`

	NewAnnotatorFlows []AnnotatorFlow `json:"new_annotatorFlows,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateFlowsOptions : Instantiate UpdateFlowsOptions
func (*AnnotatorForClinicalDataAcdV1) NewUpdateFlowsOptions(id string) *UpdateFlowsOptions {
	return &UpdateFlowsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateFlowsOptions) SetID(id string) *UpdateFlowsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetNewID : Allow user to set NewID
func (options *UpdateFlowsOptions) SetNewID(newID string) *UpdateFlowsOptions {
	options.NewID = core.StringPtr(newID)
	return options
}

// SetNewName : Allow user to set NewName
func (options *UpdateFlowsOptions) SetNewName(newName string) *UpdateFlowsOptions {
	options.NewName = core.StringPtr(newName)
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateFlowsOptions) SetNewDescription(newDescription string) *UpdateFlowsOptions {
	options.NewDescription = core.StringPtr(newDescription)
	return options
}

// SetNewPublishedDate : Allow user to set NewPublishedDate
func (options *UpdateFlowsOptions) SetNewPublishedDate(newPublishedDate string) *UpdateFlowsOptions {
	options.NewPublishedDate = core.StringPtr(newPublishedDate)
	return options
}

// SetNewPublish : Allow user to set NewPublish
func (options *UpdateFlowsOptions) SetNewPublish(newPublish bool) *UpdateFlowsOptions {
	options.NewPublish = core.BoolPtr(newPublish)
	return options
}

// SetNewVersion : Allow user to set NewVersion
func (options *UpdateFlowsOptions) SetNewVersion(newVersion string) *UpdateFlowsOptions {
	options.NewVersion = core.StringPtr(newVersion)
	return options
}

// SetNewCartridgeID : Allow user to set NewCartridgeID
func (options *UpdateFlowsOptions) SetNewCartridgeID(newCartridgeID string) *UpdateFlowsOptions {
	options.NewCartridgeID = core.StringPtr(newCartridgeID)
	return options
}

// SetNewAnnotatorFlows : Allow user to set NewAnnotatorFlows
func (options *UpdateFlowsOptions) SetNewAnnotatorFlows(newAnnotatorFlows []AnnotatorFlow) *UpdateFlowsOptions {
	options.NewAnnotatorFlows = newAnnotatorFlows
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateFlowsOptions) SetHeaders(param map[string]string) *UpdateFlowsOptions {
	options.Headers = param
	return options
}

// UpdateProfileOptions : The UpdateProfile options.
type UpdateProfileOptions struct {
	// Profile ID.
	ID *string `json:"id" validate:"required"`

	NewID *string `json:"new_id,omitempty"`

	NewName *string `json:"new_name,omitempty"`

	NewDescription *string `json:"new_description,omitempty"`

	NewPublishedDate *string `json:"new_publishedDate,omitempty"`

	NewPublish *bool `json:"new_publish,omitempty"`

	NewVersion *string `json:"new_version,omitempty"`

	NewCartridgeID *string `json:"new_cartridgeId,omitempty"`

	NewAnnotators []Annotator `json:"new_annotators,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProfileOptions : Instantiate UpdateProfileOptions
func (*AnnotatorForClinicalDataAcdV1) NewUpdateProfileOptions(id string) *UpdateProfileOptions {
	return &UpdateProfileOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateProfileOptions) SetID(id string) *UpdateProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetNewID : Allow user to set NewID
func (options *UpdateProfileOptions) SetNewID(newID string) *UpdateProfileOptions {
	options.NewID = core.StringPtr(newID)
	return options
}

// SetNewName : Allow user to set NewName
func (options *UpdateProfileOptions) SetNewName(newName string) *UpdateProfileOptions {
	options.NewName = core.StringPtr(newName)
	return options
}

// SetNewDescription : Allow user to set NewDescription
func (options *UpdateProfileOptions) SetNewDescription(newDescription string) *UpdateProfileOptions {
	options.NewDescription = core.StringPtr(newDescription)
	return options
}

// SetNewPublishedDate : Allow user to set NewPublishedDate
func (options *UpdateProfileOptions) SetNewPublishedDate(newPublishedDate string) *UpdateProfileOptions {
	options.NewPublishedDate = core.StringPtr(newPublishedDate)
	return options
}

// SetNewPublish : Allow user to set NewPublish
func (options *UpdateProfileOptions) SetNewPublish(newPublish bool) *UpdateProfileOptions {
	options.NewPublish = core.BoolPtr(newPublish)
	return options
}

// SetNewVersion : Allow user to set NewVersion
func (options *UpdateProfileOptions) SetNewVersion(newVersion string) *UpdateProfileOptions {
	options.NewVersion = core.StringPtr(newVersion)
	return options
}

// SetNewCartridgeID : Allow user to set NewCartridgeID
func (options *UpdateProfileOptions) SetNewCartridgeID(newCartridgeID string) *UpdateProfileOptions {
	options.NewCartridgeID = core.StringPtr(newCartridgeID)
	return options
}

// SetNewAnnotators : Allow user to set NewAnnotators
func (options *UpdateProfileOptions) SetNewAnnotators(newAnnotators []Annotator) *UpdateProfileOptions {
	options.NewAnnotators = newAnnotators
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProfileOptions) SetHeaders(param map[string]string) *UpdateProfileOptions {
	options.Headers = param
	return options
}

// Usage : Usage insight model struct
type Usage struct {
	ExplicitScore *float64 `json:"explicitScore,omitempty"`

	PatientReportedScore *float64 `json:"patientReportedScore,omitempty"`

	DiscussedScore *float64 `json:"discussedScore,omitempty"`

	TakenScore *float64 `json:"takenScore,omitempty"`

	LabMeasurementScore *float64 `json:"labMeasurementScore,omitempty"`
}

// UnmarshalUsage unmarshals an instance of Usage from the specific raw message.
func UnmarshalUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Usage)
	err = core.UnmarshalPrimitive(m, "explicitScore", &obj.ExplicitScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "patientReportedScore", &obj.PatientReportedScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "discussedScore", &obj.DiscussedScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "takenScore", &obj.TakenScore)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labmeasurementScore", &obj.LabMeasurementScore)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceError : Object representing an HTTP response with an error.
type ServiceError struct {
	// respone code.
	Code *int64 `json:"code,omitempty"`

	// response error message.
	Message *string `json:"message,omitempty"`

	// error severity level.
	Level *string `json:"level,omitempty"`

	// error description.
	Description *string `json:"description,omitempty"`

	// additional error information.
	MoreInfo *string `json:"moreInfo,omitempty"`

	// error message correlation identifier.
	CorrelationID *string `json:"correlationId,omitempty"`

	Artifact *string `json:"artifact,omitempty"`

	Href *string `json:"href,omitempty"`
}

// Constants associated with the ServiceError.Level property.
// error severity level.
const (
	ServiceError_Level_Error   = "ERROR"
	ServiceError_Level_Info    = "INFO"
	ServiceError_Level_Warning = "WARNING"
)

// UnmarshalServiceError unmarshals an instance of ServiceError from the specified map of raw messages.
func UnmarshalServiceError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceError)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "level", &obj.Level)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "moreInfo", &obj.MoreInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "correlationId", &obj.CorrelationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "artifact", &obj.Artifact)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
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
	ServiceStatus_ServiceState_Error   = "ERROR"
	ServiceStatus_ServiceState_Ok      = "OK"
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
