// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KabaneroStack kabanero stack
//
// swagger:model KabaneroStack
type KabaneroStack struct {

	// name
	Name string `json:"name,omitempty"`

	// status
	Status []*KabaneroStackStatusItems0 `json:"status"`
}

// Validate validates this kabanero stack
func (m *KabaneroStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KabaneroStack) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	for i := 0; i < len(m.Status); i++ {
		if swag.IsZero(m.Status[i]) { // not required
			continue
		}

		if m.Status[i] != nil {
			if err := m.Status[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KabaneroStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KabaneroStack) UnmarshalBinary(b []byte) error {
	var res KabaneroStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KabaneroStackStatusItems0 kabanero stack status items0
//
// swagger:model KabaneroStackStatusItems0
type KabaneroStackStatusItems0 struct {

	// image name
	ImageName string `json:"Image name,omitempty"`

	// digest check
	DigestCheck string `json:"digest check,omitempty"`

	// image digest
	ImageDigest string `json:"image digest,omitempty"`

	// kabanero digest
	KabaneroDigest string `json:"kabanero digest,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this kabanero stack status items0
func (m *KabaneroStackStatusItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KabaneroStackStatusItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KabaneroStackStatusItems0) UnmarshalBinary(b []byte) error {
	var res KabaneroStackStatusItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
