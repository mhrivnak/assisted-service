// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageCreateParams image create params
//
// swagger:model image-create-params
type ImageCreateParams struct {

	// SSH public key for debugging the installation.
	SSHPublicKey string `json:"ssh_public_key,omitempty"`
}

// Validate validates this image create params
func (m *ImageCreateParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageCreateParams) UnmarshalBinary(b []byte) error {
	var res ImageCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
