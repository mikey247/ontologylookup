package main

import (
	"flag"

	"github.com/mikey247/ontologylookup"
)

func main() {
	ontologyID := flag.String("ontologyID", "efo", "The ID of the ontology to fetch details for")
	format := flag.String("format", "machine", "The format to return the ontology details in")

	flag.Parse()

	ontologylookup.FetchOntologyDetails(*ontologyID, *format)
}
