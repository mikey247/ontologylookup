# Ontology Lookup Service

This is a Go project that fetches ontology details from the EBI Ontology Lookup Service.

## Project Structure

go.mod main.go result.json service/ ontology_lookup.go


## How to Run

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Run the `main.go` file with the ontologyID and format flags:

```sh
 go run cmd/ontologylookup/main.go -ontologyID=efo -format=machine
```

The package is also published pkg.go.dev on https://pkg.go.dev/github.com/mikey247/ontologylookup

## Functionality
The main function of this project is defined in the FetchOntologyDetails function in the main.go file. This function fetches ontology details for a given ontology ID and format. If the format is "machine", the details are written to a result.json file. Otherwise, the details are printed to the console.

The FetchOntologyDetails function uses the OntologyLookup service defined in the service/ontology_lookup.go file to fetch the ontology details from the EBI Ontology Lookup Service.

## Output
The output of this project is either a result.json file containing the ontology details or the ontology details printed to the console, depending on the format specified when running the script.
