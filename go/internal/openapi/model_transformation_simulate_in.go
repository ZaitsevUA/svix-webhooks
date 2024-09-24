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

// TransformationSimulateIn struct for TransformationSimulateIn
type TransformationSimulateIn struct {
	Channels *[]string `json:"channels,omitempty"`
	Code string `json:"code"`
	// The event type's name
	EventType string `json:"eventType"`
	Payload map[string]interface{} `json:"payload"`
}

// NewTransformationSimulateIn instantiates a new TransformationSimulateIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationSimulateIn(code string, eventType string, payload map[string]interface{}) *TransformationSimulateIn {
	this := TransformationSimulateIn{}
	this.Code = code
	this.EventType = eventType
	this.Payload = payload
	return &this
}

// NewTransformationSimulateInWithDefaults instantiates a new TransformationSimulateIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationSimulateInWithDefaults() *TransformationSimulateIn {
	this := TransformationSimulateIn{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransformationSimulateIn) GetChannels() []string {
	if o == nil || o.Channels == nil {
		var ret []string
		return ret
	}
	return *o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransformationSimulateIn) GetChannelsOk() (*[]string, bool) {
	if o == nil || o.Channels == nil {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *TransformationSimulateIn) HasChannels() bool {
	if o != nil && o.Channels != nil {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []string and assigns it to the Channels field.
func (o *TransformationSimulateIn) SetChannels(v []string) {
	o.Channels = &v
}

// GetCode returns the Code field value
func (o *TransformationSimulateIn) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TransformationSimulateIn) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TransformationSimulateIn) SetCode(v string) {
	o.Code = v
}

// GetEventType returns the EventType field value
func (o *TransformationSimulateIn) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransformationSimulateIn) GetEventTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransformationSimulateIn) SetEventType(v string) {
	o.EventType = v
}

// GetPayload returns the Payload field value
func (o *TransformationSimulateIn) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *TransformationSimulateIn) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *TransformationSimulateIn) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

func (o TransformationSimulateIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableTransformationSimulateIn struct {
	value *TransformationSimulateIn
	isSet bool
}

func (v NullableTransformationSimulateIn) Get() *TransformationSimulateIn {
	return v.value
}

func (v *NullableTransformationSimulateIn) Set(val *TransformationSimulateIn) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationSimulateIn) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationSimulateIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationSimulateIn(val *TransformationSimulateIn) *NullableTransformationSimulateIn {
	return &NullableTransformationSimulateIn{value: val, isSet: true}
}

func (v NullableTransformationSimulateIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationSimulateIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


