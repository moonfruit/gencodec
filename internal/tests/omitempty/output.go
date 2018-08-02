// Code generated by github.com/moonfruit/gencodec. DO NOT EDIT.

package omitempty

import (
	"encoding/json"
)

var _ = (*Xo)(nil)

// MarshalJSON marshals as JSON.
func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Int replacedInt `json:",omitempty"`
	}
	var enc X
	enc.Int = replacedInt(x.Int)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Int *replacedInt `json:",omitempty"`
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = int(*dec.Int)
	}
	return nil
}

// MarshalYAML marshals as YAML.
func (x X) MarshalYAML() (interface{}, error) {
	type X struct {
		Int replacedInt `json:",omitempty"`
	}
	var enc X
	enc.Int = replacedInt(x.Int)
	return &enc, nil
}

// UnmarshalYAML unmarshals from YAML.
func (x *X) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type X struct {
		Int *replacedInt `json:",omitempty"`
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = int(*dec.Int)
	}
	return nil
}

// MarshalTOML marshals as TOML.
func (x X) MarshalTOML() (interface{}, error) {
	type X struct {
		Int replacedInt `json:",omitempty"`
	}
	var enc X
	enc.Int = replacedInt(x.Int)
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (x *X) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type X struct {
		Int *replacedInt `json:",omitempty"`
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = int(*dec.Int)
	}
	return nil
}
