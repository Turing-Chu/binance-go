/*
Binance Public Spot API

OpenAPI Specifications for the Binance Public Spot API  API documents:   - [https://github.com/binance/binance-spot-api-docs](https://github.com/binance/binance-spot-api-docs)   - [https://binance-docs.github.io/apidocs/spot/en](https://binance-docs.github.io/apidocs/spot/en)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// InlineResponse20031 struct for InlineResponse20031
type InlineResponse20031 struct {
	Id string `json:"id"`
}

// NewInlineResponse20031 instantiates a new InlineResponse20031 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20031(id string) *InlineResponse20031 {
	this := InlineResponse20031{}
	this.Id = id
	return &this
}

// NewInlineResponse20031WithDefaults instantiates a new InlineResponse20031 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20031WithDefaults() *InlineResponse20031 {
	this := InlineResponse20031{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse20031) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20031) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse20031) SetId(v string) {
	o.Id = v
}

func (o InlineResponse20031) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20031 struct {
	value *InlineResponse20031
	isSet bool
}

func (v NullableInlineResponse20031) Get() *InlineResponse20031 {
	return v.value
}

func (v *NullableInlineResponse20031) Set(val *InlineResponse20031) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20031) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20031) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20031(val *InlineResponse20031) *NullableInlineResponse20031 {
	return &NullableInlineResponse20031{value: val, isSet: true}
}

func (v NullableInlineResponse20031) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20031) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

