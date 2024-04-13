package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mikey247/ontologylookup/service"
)

func FetchOntologyDetails(ontologyID string, format string) error {
	ol := service.OntologyLookup{
		BaseURL: "https://www.ebi.ac.uk/ols4/api/ontologies/",
	}

	details, err := ol.FetchOntologyDetails(ontologyID)

	if err != nil {
		fmt.Printf("Error fetching ontology details: %v\n", err)
		return err
	}

	if format == "machine" {
		file, err := os.Create("result.json")
		if err != nil {
			fmt.Printf("Error creating result.json file: %v\n", err)
			return err
		}
		defer file.Close()

		jsonData, err := json.Marshal(details)
		if err != nil {
			fmt.Printf("Error marshaling ontology details to JSON: %v\n", err)
			return err
		}

		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Printf("Error writing ontology details to result.json file: %v\n", err)
			return err
		}

		fmt.Println("Ontology details written to result.json file")

		return nil
	}

	fmt.Printf("Ontology Title: %s\n", details.OntologyTitle)
	fmt.Printf("Ontology Description: %s\n", details.OntologyDescription)
	fmt.Printf("Term Count: %d\n", details.TermCount)
	fmt.Printf("Status: %s\n", details.Status)

	return nil
}

func main() {
	ontologyID := flag.String("ontologyID", "efo", "The ID of the ontology to fetch details for")
	format := flag.String("format", "json", "The format to return the ontology details in")

	flag.Parse()

	FetchOntologyDetails(*ontologyID, *format)
}
