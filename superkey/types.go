package superkey

import (
	sourcesapi "github.com/lindgrenj6/sources-api-client-go"
)

// CreateRequest - struct representing a request for a superkey
type CreateRequest struct {
	TenantID        string            `json:"tenant_id"`
	SourceID        string            `json:"source_id"`
	ApplicationID   string            `json:"application_id"`
	ApplicationType string            `json:"application_type"`
	SuperKey        string            `json:"super_key"`
	Provider        string            `json:"provider"`
	Extra           map[string]string `json:"extra"`
	SuperKeySteps   []Step            `json:"superkey_steps"`
}

// Step - struct representing a step for SuperKey
type Step struct {
	Step          int               `json:"step"`
	Name          string            `json:"name"`
	Payload       string            `json:"payload"`
	Substitutions map[string]string `json:"substitutions"`
}

// DestroyRequest - struct representing a teardown request for an application
// created through superkey
type DestroyRequest struct {
	TenantID       string                       `json:"tenant_id"`
	SuperKey       string                       `json:"super_key"`
	GUID           string                       `json:"guid"`
	Provider       string                       `json:"provider"`
	StepsCompleted map[string]map[string]string `json:"steps_completed"`
	SuperKeySteps  []Step                       `json:"superkey_steps"`
}

// App - represents an application that can be posted to sources after being
// populated
type App struct {
	SourceID    string                                      `json:"source_id"`
	Extra       map[string]interface{}                      `json:"extra"`
	AuthPayload sourcesapi.BulkCreatePayloadAuthentications `json:"authentication_payload"`
}

// ForgedApplication - struct to hold the output of a superkey
// create_application request
type ForgedApplication struct {
	Product        *App
	StepsCompleted map[string]map[string]string
	Request        *CreateRequest
	Client         Provider
	GUID           string
}

// Provider the interface for all of the superkey providers currently just a
// single method is needed (ForgeApplication)
type Provider interface {
	ForgeApplication(*CreateRequest) (*ForgedApplication, error)
	TearDown(*ForgedApplication) []error
}
