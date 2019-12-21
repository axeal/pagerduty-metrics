package main

import (
	"os"
	"log"
	"fmt"
	"encoding/json"
	"github.com/axeal/pagerduty-metrics/pd"
)

func main() {

	pdToken := os.Getenv("PAGERDUTY_TOKEN")
	pdClient := pd.NewClient(pdToken)
	incidents, err := pdClient.ListAllIncidents()
	if err != nil {
		log.Fatalf("Error listing incidents: %s", err)
	}
	for _, incident := range incidents {
		prettyJSON, err := json.MarshalIndent(incident, "", "    ")
		if err != nil {
			log.Fatal("Failed to generate json", err)
		}
		fmt.Printf("%s\n", string(prettyJSON))
	}
	fmt.Printf("Number of elements in list: %d", len(incidents))

}