// Package issue1168 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/vulados/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package issue1168

import (
	"encoding/json"
	"fmt"
)

// ProblemDetails defines model for ProblemDetails.
type ProblemDetails struct {
	// Detail A human readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`

	// Instance An absolute URI that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced.
	Instance *string `json:"instance,omitempty"`

	// Status The HTTP status code generated by the origin server for this occurrence of the problem.
	Status *int32 `json:"status,omitempty"`

	// Title A short, summary of the problem type. Written in english and readable for engineers (usually not suited for non technical stakeholders and not localized); example: Service Unavailable
	Title *string `json:"title,omitempty"`

	// Type An absolute URI that identifies the problem type.  When dereferenced, it SHOULD provide human-readable documentation for the problem type (e.g., using HTML).
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Misc400Error defines model for Misc400Error.
type Misc400Error = ProblemDetails

// Misc404Error defines model for Misc404Error.
type Misc404Error = ProblemDetails

// Getter for additional properties for ProblemDetails. Returns the specified
// element and whether it was found
func (a ProblemDetails) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ProblemDetails
func (a *ProblemDetails) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ProblemDetails to handle AdditionalProperties
func (a *ProblemDetails) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["detail"]; found {
		err = json.Unmarshal(raw, &a.Detail)
		if err != nil {
			return fmt.Errorf("error reading 'detail': %w", err)
		}
		delete(object, "detail")
	}

	if raw, found := object["instance"]; found {
		err = json.Unmarshal(raw, &a.Instance)
		if err != nil {
			return fmt.Errorf("error reading 'instance': %w", err)
		}
		delete(object, "instance")
	}

	if raw, found := object["status"]; found {
		err = json.Unmarshal(raw, &a.Status)
		if err != nil {
			return fmt.Errorf("error reading 'status': %w", err)
		}
		delete(object, "status")
	}

	if raw, found := object["title"]; found {
		err = json.Unmarshal(raw, &a.Title)
		if err != nil {
			return fmt.Errorf("error reading 'title': %w", err)
		}
		delete(object, "title")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ProblemDetails to handle AdditionalProperties
func (a ProblemDetails) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Detail != nil {
		object["detail"], err = json.Marshal(a.Detail)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'detail': %w", err)
		}
	}

	if a.Instance != nil {
		object["instance"], err = json.Marshal(a.Instance)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'instance': %w", err)
		}
	}

	if a.Status != nil {
		object["status"], err = json.Marshal(a.Status)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'status': %w", err)
		}
	}

	if a.Title != nil {
		object["title"], err = json.Marshal(a.Title)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'title': %w", err)
		}
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
