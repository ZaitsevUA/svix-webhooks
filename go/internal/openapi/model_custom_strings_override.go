/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomStringsOverride struct for CustomStringsOverride
type CustomStringsOverride struct {
	ChannelsMany NullableString `json:"channelsMany,omitempty"`
	ChannelsOne NullableString `json:"channelsOne,omitempty"`
}

// NewCustomStringsOverride instantiates a new CustomStringsOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomStringsOverride() *CustomStringsOverride {
	this := CustomStringsOverride{}
	return &this
}

// NewCustomStringsOverrideWithDefaults instantiates a new CustomStringsOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomStringsOverrideWithDefaults() *CustomStringsOverride {
	this := CustomStringsOverride{}
	return &this
}

// GetChannelsMany returns the ChannelsMany field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomStringsOverride) GetChannelsMany() string {
	if o == nil || o.ChannelsMany.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChannelsMany.Get()
}

// GetChannelsManyOk returns a tuple with the ChannelsMany field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomStringsOverride) GetChannelsManyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChannelsMany.Get(), o.ChannelsMany.IsSet()
}

// HasChannelsMany returns a boolean if a field has been set.
func (o *CustomStringsOverride) HasChannelsMany() bool {
	if o != nil && o.ChannelsMany.IsSet() {
		return true
	}

	return false
}

// SetChannelsMany gets a reference to the given NullableString and assigns it to the ChannelsMany field.
func (o *CustomStringsOverride) SetChannelsMany(v string) {
	o.ChannelsMany.Set(&v)
}
// SetChannelsManyNil sets the value for ChannelsMany to be an explicit nil
func (o *CustomStringsOverride) SetChannelsManyNil() {
	o.ChannelsMany.Set(nil)
}

// UnsetChannelsMany ensures that no value is present for ChannelsMany, not even an explicit nil
func (o *CustomStringsOverride) UnsetChannelsMany() {
	o.ChannelsMany.Unset()
}

// GetChannelsOne returns the ChannelsOne field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomStringsOverride) GetChannelsOne() string {
	if o == nil || o.ChannelsOne.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChannelsOne.Get()
}

// GetChannelsOneOk returns a tuple with the ChannelsOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomStringsOverride) GetChannelsOneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChannelsOne.Get(), o.ChannelsOne.IsSet()
}

// HasChannelsOne returns a boolean if a field has been set.
func (o *CustomStringsOverride) HasChannelsOne() bool {
	if o != nil && o.ChannelsOne.IsSet() {
		return true
	}

	return false
}

// SetChannelsOne gets a reference to the given NullableString and assigns it to the ChannelsOne field.
func (o *CustomStringsOverride) SetChannelsOne(v string) {
	o.ChannelsOne.Set(&v)
}
// SetChannelsOneNil sets the value for ChannelsOne to be an explicit nil
func (o *CustomStringsOverride) SetChannelsOneNil() {
	o.ChannelsOne.Set(nil)
}

// UnsetChannelsOne ensures that no value is present for ChannelsOne, not even an explicit nil
func (o *CustomStringsOverride) UnsetChannelsOne() {
	o.ChannelsOne.Unset()
}

func (o CustomStringsOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelsMany.IsSet() {
		toSerialize["channelsMany"] = o.ChannelsMany.Get()
	}
	if o.ChannelsOne.IsSet() {
		toSerialize["channelsOne"] = o.ChannelsOne.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCustomStringsOverride struct {
	value *CustomStringsOverride
	isSet bool
}

func (v NullableCustomStringsOverride) Get() *CustomStringsOverride {
	return v.value
}

func (v *NullableCustomStringsOverride) Set(val *CustomStringsOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomStringsOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomStringsOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomStringsOverride(val *CustomStringsOverride) *NullableCustomStringsOverride {
	return &NullableCustomStringsOverride{value: val, isSet: true}
}

func (v NullableCustomStringsOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomStringsOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


