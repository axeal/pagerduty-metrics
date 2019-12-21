package main

import (
	"os"
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
)

func main() {
	var opts pagerduty.ListEscalationPoliciesOptions
	var authtoken = os.Getenv("PAGERDUTY_TOKEN")
	client := pagerduty.NewClient(authtoken)
	eps, err := client.ListEscalationPolicies(opts)
	if err != nil {
		panic(err)
	}
	for _, p := range eps.EscalationPolicies {
		fmt.Println(p.Name)
	}
}