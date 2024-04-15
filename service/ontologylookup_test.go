package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchOntologyDetails(t *testing.T) {
	// Setup a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/expectedPath" {
			t.Errorf("Expected to request '/expectedPath', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"config": map[string]interface{}{
				"title":       "Example Ontology",
				"description": "A sample ontology",
			},
			"numberOfTerms": 100,
			"status":        "active",
		})
	}))
	defer ts.Close()

	ol := OntologyLookup{BaseURL: ts.URL}
	details, err := ol.FetchOntologyDetails("expectedPath")
	if err != nil {
		t.Fatal("Failed to fetch ontology details:", err)
	}

	if details.OntologyTitle != "Example Ontology" {
		t.Errorf("Expected title 'Example Ontology', got '%s'", details.OntologyTitle)
	}
}
