package pd

import (
	"log"
	"github.com/PagerDuty/go-pagerduty"
)

// Client wraps the PagerDuty Client
type Client struct {
	PagerDutyClient pagerduty.Client
}

// NewClient creates a new PagerDuty API Client
func NewClient (authtoken string) *Client {
	return &Client{
		PagerDutyClient: *pagerduty.NewClient(authtoken),
	}
}

func (c *Client) ListIncidents() ([]pagerduty.Incident, error) {

	var opts pagerduty.ListIncidentsOptions
	
	response, err := c.PagerDutyClient.ListIncidents(opts)
	if err != nil {
		return nil, err
	}
	return response.Incidents, nil
}

func (c *Client) ListAllIncidents() ([]pagerduty.Incident, error) {

	var opts pagerduty.ListIncidentsOptions
	opts.Limit = 100
	opts.Offset = 0
	opts.DateRange = "all"
	var incidents []pagerduty.Incident
	_, err := c.listIncidentsRecursively(&opts, &incidents)
	if err != nil {
		return nil, err
	}
	return incidents, nil
}

	func (c *Client) listIncidentsRecursively(opts *pagerduty.ListIncidentsOptions, incidents *[]pagerduty.Incident) (*[]pagerduty.Incident, error) {	
	response, err := c.PagerDutyClient.ListIncidents(*opts)
	if err != nil {
		return nil, err
	}
	log.Printf("Offset: %d, Limit: %d, Length: %d", opts.Offset, opts.Limit, len(response.Incidents))
	*incidents = append(*incidents, response.Incidents...)
	if response.More {
		opts.Offset += opts.Limit
		return c.listIncidentsRecursively(opts, incidents)
	}
	return incidents, nil
}