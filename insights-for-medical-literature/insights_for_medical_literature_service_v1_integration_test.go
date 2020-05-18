package insightsformedicalliteratureservicev1_test

import (
	"github.com/IBM/go-sdk-core/v3/core"
	insightsformedicalliteratureservicev1 "github.com/IBM/whcs-go-sdk/insights-for-medical-literature"
	"github.com/joho/godotenv"
	"os"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const configFile = "iml-service.env"

var (
	IML          *insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1
	err          error
	configLoaded bool = false
	version      string
	serviceUrl   string
	iamUrl       string
	apikey       string
	disableSsl   bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration not found, skipping....")
	}
}

var _ = Describe(`InsightsForMedicalLiteratureV1`, func() {
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

		if apikey is nil {
			IML, err = insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:    serviceUrl,
				Version: core.StringPtr(version),
				Authenticator: &core.NoauthAuthenticator{
				},
			})
			IML.DisableSSLVerification(true)
		} else {
			IML, err = insightsformedicalliteratureservicev1.NewInsightsForMedicalLiteratureServiceV1(&insightsformedicalliteratureservicev1.InsightsForMedicalLiteratureServiceV1Options{
				URL:    serviceUrl,
				Version: core.StringPtr(version),
				Authenticator: &core.IamAuthenticator{
					ApiKey: apikey,
					URL: iamUrl,
					DisableSSLVerification: disableSsl,
				},
			})
		}
		Expect(err).To(BeNil())
//		Expect(err).ToNot(BeNil())
		Expect(IML.Service.Options.URL).To(Not(Equal("")))
	})

	Describe(`GetCorporaConfig(getCorporaConfigOptions *GetCorporaConfigOptions)`, func() {
		It(`Successfully get corpora info`, func() {
			shouldSkipTest()

			getDocumentsOptions := IML.NewGetCorporaConfigOptions()
			result, detailedResponse, err := IML.GetCorporaConfig(getDocumentsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			numCorpora := len(result.Corpora)
			corporaCounter := 0
			for corporaCounter < numCorpora {
				Expect(*result.Corpora[corporaCounter].CorpusName).ToNot(BeNil())
				Expect(*result.Corpora[corporaCounter].DescriptiveName).ToNot(BeNil())
				ontologies := result.Corpora[corporaCounter].Ontologies
				numOntologies := len(ontologies)
				ontologyCounter := 0
				for ontologyCounter < numOntologies {
					Expect(result.Corpora[corporaCounter].Ontologies[ontologyCounter]).ToNot(BeNil())
					ontologyCounter++
				}
				corporaCounter++
			}
		})
	})

	Describe(`GetCorpusConfig(getCorpusConfigOptions *GetCorpusConfigOptions)`, func() {
		It(`Successfully get corpous config info`, func() {
			shouldSkipTest()

			getCorpusConfigOptions := IML.NewGetCorpusConfigOptions("ctgov")
			result, detailedResponse, err := IML.GetCorpusConfig(getCorpusConfigOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.Corpora[0].CorpusName).ToNot(BeNil())
			Expect(*result.Corpora[0].DescriptiveName).ToNot(BeNil())
			ontologies := result.Corpora[0].Ontologies
			numOntologies := len(ontologies)
			ontologyCounter := 0
			for ontologyCounter < numOntologies {
				Expect(result.Corpora[0].Ontologies[ontologyCounter]).ToNot(BeNil())
				ontologyCounter++
			}
		})

	})

	Describe(`GetHealthCheckStatus(getHealthCheckStatusOptions *GetHealthCheckStatusOptions)`, func() {
		It(`Successfully get service status`, func() {
			healthCheckOptions := IML.NewGetHealthCheckStatusOptions()
			result, detailedResponse, err := IML.GetHealthCheckStatus(healthCheckOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result.ServiceState).To(Equal(core.StringPtr("OK")))
		})
	})

	Describe(`GetConcepts(getConceptsOptions *GetConceptsOptions)`, func() {
		It(`Successfully get concept info`, func() {
			preferredName := []string{"heart"}
			getConceptsOptions := IML.NewGetConceptsOptions("ctgov")
			getConceptsOptions.SetPreferredNames(preferredName)
			result, detailedResponse, err := IML.GetConcepts(getConceptsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(*result.Concepts[0].Cui).ToNot(BeNil())
			Expect(*result.Concepts[0].PreferredName).ToNot(BeNil())
		})
	})

	Describe(`GetCuiInfo(getCuiInfoOptions *GetCuiInfoOptions`, func() {
		It(`Successfully get concept details`, func() {
			getCuiInfoOptions := IML.NewGetCuiInfoOptions("ctgov", "C0018787")
			result, detailedResponse, err := IML.GetCuiInfo(getCuiInfoOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(*result.PreferredName).ToNot(BeNil())
			Expect(*result.Definition).ToNot(BeNil())
			numTypes := len(result.SemanticTypes)
			typeCounter := 0
			for typeCounter < numTypes {
				Expect(result.SemanticTypes[typeCounter]).ToNot(BeNil())
				typeCounter++
			}
			numSForms := len(result.SurfaceForms)
			formsCounter := 0
			for formsCounter < numSForms {
				Expect(result.SurfaceForms[formsCounter]).ToNot(BeNil())
				formsCounter++
			}
		})
	})

	Describe(`GetHitCount(getHitCountOptions *GetHitCountOptions`, func() {
		It(`Successfully hit a concepts corpus concentration`, func() {
			getHitCountOptions := IML.NewGetHitCountOptions("ctgov", "C0018787")
			result, detailedResponse, err := IML.GetHitCount(getHitCountOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(result).ToNot(BeZero())
		})
	})

	Describe(`GetRelatedConcepts(getRelatedConceptsOptions *GetRelatedConceptsOptions`, func() {
		It(`Successfully retrieve related concepts`, func() {
			getRelatedConceptsOptions := IML.NewGetRelatedConceptsOptions("medline", "C0018787", "chd")
			result, detailedResponse, err := IML.GetRelatedConcepts(getRelatedConceptsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetSimilarConcepts(getSimilarConceptsOptions *GetSimilarConceptsOptions)`, func() {
		It(`Successly retrieve similar concepts`, func() {
			getRelatedConceptsOptions := IML.NewGetRelatedConceptsOptions("medline", "C0018787", "chd")
			result, detailedResponse, err := IML.GetRelatedConcepts(getRelatedConceptsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
		})
	})

	Describe(`GetDocuments(getDocumentsOptions *GetDocumentsOptions)`, func() {
		It(`Successfully get document count info`, func() {
			shouldSkipTest()

			getDocumentsOptions := IML.NewGetDocumentsOptions("ctgov")
			result, detailedResponse, err := IML.GetDocuments(getDocumentsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())
			Expect(*result.DocumentCount).ToNot(BeZero())
			numProviders := len(result.Providers)
			providerCounter := 0
			for providerCounter < numProviders {
				provider := result.Providers[providerCounter]
				Expect(*provider.Name).ToNot(BeNil())
				Expect(*provider.DocumentCount).ToNot(BeNil())
				providerCounter++
			}
		})
	})

	Describe(`GetDocumentAnnotations(getDocumentAnnotationsOptions *GetDocumentAnnotationsOptions`, func() {
		It(`Successfully retrieve document annotations`, func() {
			getDocumentAnnotationsOptions := IML.NewGetDocumentAnnotationsOptions("ctgov", "NCT00796159", "title")
			detailedResponse, err := IML.GetDocumentAnnotations(getDocumentAnnotationsOptions)
			Expect(err).To(BeNil())
			Expect(detailedResponse).ToNot(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
//			Expect(detailedResponse.Result).ToNot(BeNil())
//			unstructured := detailedResponse.Result.Unstructured[0]
//			data := unstructured.Data
//			Expect(data.Text).ToNot(BeNil())
//			attributes := data.AttributeValues
//			concepts := data.Concepts
//			mesh := data.Mesh

//			if len(attributes) > 0 {
//				numAttrs := len(attributes)
//				attrCounter := 0
//				for attrCounter > numAttrs {
//					attribute := attributes[attrCounter]
//					Expect(attribute.Text).ToNot(BeNil())
//					Expect(attribute.AM.UniqueID).ToNot(BeNil())
//					attrCounter++
//				}
//			}

//			if len(concepts) > 0 {
//				numConcepts := len(concepts)
//				conceptCounter := 0
//				for conceptCounter > numConcepts {
//					concept := attributes[conceptCounter]
//					Expect(concept.Text).ToNot(BeNil())
//					Expect(concept.AM.UniqueID).ToNot(BeNil())
//					conceptCounter++
//				}
//			}

//			if len(mesh) > 0 {
//				numMesh := len(attributes)
//				meshCounter := 0
//				for meshCounter > numMesh {
//					mesh := attributes[meshCounter]
//					Expect(mesh.Text).ToNot(BeNil())
//					Expect(mesh.AM.UniqueID).ToNot(BeNil())
//					meshCounter++
//				}
//			}
		})

		Describe(`GetDocumentCategories(getDocumentCategoriesOptions *GetDocumentCategoriesOptions`, func() {
			It(`Successfully retrieve document annotations for category`, func() {
				getDocumentCategoriesOptions := IML.NewGetDocumentCategoriesOptions("ctgov", "NCT00796159")
				result, detailedResponse, err := IML.GetDocumentCategories(getDocumentCategoriesOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
				Expect(*result.HighlightedTitle).ToNot(BeNil())
				Expect(*result.HighlightedAbstract).ToNot(BeNil())
				Expect(*result.HighlightedBody).ToNot(BeNil())
			})
		})

		Describe(`GetDocumentMultipleCategories(getDocumentMultipleCategoriesOptions *GetDocumentMultipleCategoriesOptions)`, func() {
			It(`Successfully retrieve annotations for multiple categories`, func() {
				category := insightsformedicalliteratureservicev1.Category{}
				category.Name = core.StringPtr(insightsformedicalliteratureservicev1.GetDocumentCategoriesOptions_Category_Drugs)
				category.Category = core.StringPtr(insightsformedicalliteratureservicev1.GetDocumentCategoriesOptions_Category_Drugs)
				categories := []insightsformedicalliteratureservicev1.Category{category}

				getDocuemntMultipeCategoriesOptions := IML.NewGetDocumentMultipleCategoriesOptions("ctgov", "NCT00796159", categories)
				getDocuemntMultipeCategoriesOptions.SetCategories(categories)
				result, detailedResponse, err := IML.GetDocumentMultipleCategories(getDocuemntMultipeCategoriesOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
			})
		})

		Describe(`GetSearchMatches(getSearchMatchesOptions *GetSearchMatchesOptions)`, func() {
			It(`Successfully retrieve search matches`, func() {
				getSearchMatchesOptions := IML.NewGetSearchMatchesOptions("ctgov", "NCT00796159", .5)
				getSearchMatchesOptions.SetCuis([]string{"C0018787"})
				result, detailedResponse, err := IML.GetSearchMatches(getSearchMatchesOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
				Expect(*result.DocumentID).To(Equal("NCT00796159"))
			})
		})

		Describe(`Search(searchOptions *SearchOptions)`, func() {
			It(`Successfully search the corpus`, func() {
				body := `{"query": { "concepts": [{ "ontology": "concepts", "cui": "C0018787", "rank": "10"}]}, "returns": { "documents": { "limit": "10", "offset": 0}}}`
				getSearchOptions := IML.NewSearchOptions("ctgov", body)
				result, detailedResponse, err := IML.Search(getSearchOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
				Expect(*result.TotalDocumentCount).ToNot(BeZero())
			})
		})

		Describe(`GetFields(getFieldsOptions *GetFieldsOptions)`, func() {
			It(`Successfully retrieve corpus fields`, func() {
				getFieldsOptions := IML.NewGetFieldsOptions("ctgov")
				result, detailedResponse, err := IML.GetFields(getFieldsOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
				Expect(result.Fields).ToNot(BeNil())
			})
		})

		Describe(`Typeahead(typeaheadOptions *TypeaheadOptions)`, func() {
			It(`Successfully retrieve mathcing suggestions`, func() {
				getTypeaheadOptions := IML.NewTypeaheadOptions("ctgov", "hear")
				result, detailedResponse, err := IML.Typeahead(getTypeaheadOptions)
				Expect(err).To(BeNil())
				Expect(detailedResponse).ToNot(BeNil())
				Expect(detailedResponse.StatusCode).To(Equal(200))
				Expect(result).ToNot(BeNil())
				count := len(result.Concepts)
				counter := 0
				for counter < count {
					artifact := result.Concepts[counter]
					Expect(*artifact.Cui).ToNot(BeNil())
					Expect(*artifact.PreferredName).ToNot(BeNil())
					Expect(*artifact.SemanticType).ToNot(BeNil())
					Expect(*artifact.HitCount).ToNot(BeNil())
					counter++
				}
			})
		})
	})
})
