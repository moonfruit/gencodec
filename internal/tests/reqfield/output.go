// Code generated by github.com/moonfruit/gencodec. DO NOT EDIT.

package reqfield

import (
	"encoding/json"
	"errors"
)

// MarshalJSON marshals as JSON.
func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Required    int `gencodec:"required"`
		NotRequired int `gencodec:"required" json:"-"`
	}
	var enc X
	enc.Required = x.Required
	enc.NotRequired = x.NotRequired
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Required    *int `gencodec:"required"`
		NotRequired *int `gencodec:"required" json:"-"`
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Required == nil {
		return errors.New("missing required field 'required' for X")
	}
	x.Required = *dec.Required
	if dec.NotRequired != nil {
		x.NotRequired = *dec.NotRequired
	}
	return nil
}
