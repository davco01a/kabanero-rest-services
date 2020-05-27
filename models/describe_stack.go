// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeStack describe stack
//
// swagger:model describeStack
type DescribeStack struct {

	// digest check
	DigestCheck string `json:"digest check,omitempty"`

	// git repo url
	GitRepoURL string `json:"git repo url,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// image digest
	ImageDigest string `json:"image digest,omitempty"`

	// kabanero digest
	KabaneroDigest string `json:"kabanero digest,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project
	Project string `json:"project,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this describe stack
func (m *DescribeStack) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeStack) UnmarshalBinary(b []byte) error {
	var res DescribeStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
