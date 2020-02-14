/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// SpecialModelName struct for SpecialModelName
type SpecialModelName struct {
	SpecialPropertyName *int64 `json:"$special[property.name],omitempty"`
}

// GetSpecialPropertyName returns the SpecialPropertyName field value if set, zero value otherwise.
func (o *SpecialModelName) GetSpecialPropertyName() int64 {
	if o == nil || o.SpecialPropertyName == nil {
		var ret int64
		return ret
	}
	return *o.SpecialPropertyName
}

// GetSpecialPropertyNameOk returns a tuple with the SpecialPropertyName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *SpecialModelName) GetSpecialPropertyNameOk() (int64, bool) {
	if o == nil || o.SpecialPropertyName == nil {
		var ret int64
		return ret, false
	}
	return *o.SpecialPropertyName, true
}

// HasSpecialPropertyName returns a boolean if a field has been set.
func (o *SpecialModelName) HasSpecialPropertyName() bool {
	if o != nil && o.SpecialPropertyName != nil {
		return true
	}

	return false
}

// SetSpecialPropertyName gets a reference to the given int64 and assigns it to the SpecialPropertyName field.
func (o *SpecialModelName) SetSpecialPropertyName(v int64) {
	o.SpecialPropertyName = &v
}

type NullableSpecialModelName struct {
	Value SpecialModelName
	ExplicitNull bool
}

func (v NullableSpecialModelName) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableSpecialModelName) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}