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

// StacksList stacks list
//
// swagger:model stacksList
type StacksList struct {

	// activate stacks
	ActivateStacks []*CommonStack `json:"activate stacks"`

	// curated stacks
	CuratedStacks []*CommonStack `json:"curated stacks"`

	// kabanero stacks
	KabaneroStacks []*KabaneroStack `json:"kabanero stacks"`

	// new curated stacks
	NewCuratedStacks []*CommonStack `json:"new curated stacks"`

	// obsolete stacks
	ObsoleteStacks []*CommonStack `json:"obsolete stacks"`

	// repositories
	Repositories []*StacksListRepositoriesItems0 `json:"repositories"`
}

// Validate validates this stacks list
func (m *StacksList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivateStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCuratedStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKabaneroStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewCuratedStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObsoleteStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StacksList) validateActivateStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivateStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.ActivateStacks); i++ {
		if swag.IsZero(m.ActivateStacks[i]) { // not required
			continue
		}

		if m.ActivateStacks[i] != nil {
			if err := m.ActivateStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activate stacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksList) validateCuratedStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.CuratedStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.CuratedStacks); i++ {
		if swag.IsZero(m.CuratedStacks[i]) { // not required
			continue
		}

		if m.CuratedStacks[i] != nil {
			if err := m.CuratedStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("curated stacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksList) validateKabaneroStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.KabaneroStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.KabaneroStacks); i++ {
		if swag.IsZero(m.KabaneroStacks[i]) { // not required
			continue
		}

		if m.KabaneroStacks[i] != nil {
			if err := m.KabaneroStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kabanero stacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksList) validateNewCuratedStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.NewCuratedStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.NewCuratedStacks); i++ {
		if swag.IsZero(m.NewCuratedStacks[i]) { // not required
			continue
		}

		if m.NewCuratedStacks[i] != nil {
			if err := m.NewCuratedStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("new curated stacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksList) validateObsoleteStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.ObsoleteStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.ObsoleteStacks); i++ {
		if swag.IsZero(m.ObsoleteStacks[i]) { // not required
			continue
		}

		if m.ObsoleteStacks[i] != nil {
			if err := m.ObsoleteStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("obsolete stacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StacksList) validateRepositories(formats strfmt.Registry) error {

	if swag.IsZero(m.Repositories) { // not required
		return nil
	}

	for i := 0; i < len(m.Repositories); i++ {
		if swag.IsZero(m.Repositories[i]) { // not required
			continue
		}

		if m.Repositories[i] != nil {
			if err := m.Repositories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("repositories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StacksList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksList) UnmarshalBinary(b []byte) error {
	var res StacksList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StacksListRepositoriesItems0 stacks list repositories items0
//
// swagger:model StacksListRepositoriesItems0
type StacksListRepositoriesItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this stacks list repositories items0
func (m *StacksListRepositoriesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StacksListRepositoriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StacksListRepositoriesItems0) UnmarshalBinary(b []byte) error {
	var res StacksListRepositoriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
