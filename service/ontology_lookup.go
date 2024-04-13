package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OntologyLookup struct {
	BaseURL string
}

type OntologyDetails struct {
	OntologyTitle       string `json:"ontology_title"`
	OntologyDescription string `json:"ontology_description"`
	TermCount           int    `json:"term_count"`
	Status              string `json:"status"`
}

func (ol *OntologyLookup) FetchOntologyDetails(ontologyID string) (*OntologyDetails, error) {
	fmt.Printf("Fetching details for ontology with ID: %s\n", ontologyID)
	url := fmt.Sprintf("%s%s", ol.BaseURL, ontologyID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// fmt.Printf("Response status code: %d\n", response.StatusCode)

	if response.StatusCode == http.StatusOK {
		var data map[string]interface{}
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("Data: %v\n", data)

		ontologyTitle := data["config"].(map[string]interface{})["title"].(string)
		ontologyDescription := data["config"].(map[string]interface{})["description"].(string)
		termCount := int(data["numberOfTerms"].(float64))
		status := data["status"].(string)

		return &OntologyDetails{
			OntologyTitle:       ontologyTitle,
			OntologyDescription: ontologyDescription,
			TermCount:           termCount,
			Status:              status,
		}, nil

	} else {
		return nil, fmt.Errorf("Ontology not found: %d", response.StatusCode)
	}
}
