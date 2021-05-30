/*
 * Svix
 *
 * The Svix server API documentation
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ListResponseEndpointOut struct for ListResponseEndpointOut
type ListResponseEndpointOut struct {
	Data []EndpointOut `json:"data"`
	Iterator *string `json:"iterator,omitempty"`
	Done bool `json:"done"`
}

// NewListResponseEndpointOut instantiates a new ListResponseEndpointOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListResponseEndpointOut(data []EndpointOut, done bool, ) *ListResponseEndpointOut {
	this := ListResponseEndpointOut{}
	this.Data = data
	this.Done = done
	return &this
}

// NewListResponseEndpointOutWithDefaults instantiates a new ListResponseEndpointOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListResponseEndpointOutWithDefaults() *ListResponseEndpointOut {
	this := ListResponseEndpointOut{}
	return &this
}

// GetData returns the Data field value
func (o *ListResponseEndpointOut) GetData() []EndpointOut {
	if o == nil  {
		var ret []EndpointOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListResponseEndpointOut) GetDataOk() (*[]EndpointOut, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListResponseEndpointOut) SetData(v []EndpointOut) {
	o.Data = v
}

// GetIterator returns the Iterator field value if set, zero value otherwise.
func (o *ListResponseEndpointOut) GetIterator() string {
	if o == nil || o.Iterator == nil {
		var ret string
		return ret
	}
	return *o.Iterator
}

// GetIteratorOk returns a tuple with the Iterator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListResponseEndpointOut) GetIteratorOk() (*string, bool) {
	if o == nil || o.Iterator == nil {
		return nil, false
	}
	return o.Iterator, true
}

// HasIterator returns a boolean if a field has been set.
func (o *ListResponseEndpointOut) HasIterator() bool {
	if o != nil && o.Iterator != nil {
		return true
	}

	return false
}

// SetIterator gets a reference to the given string and assigns it to the Iterator field.
func (o *ListResponseEndpointOut) SetIterator(v string) {
	o.Iterator = &v
}

// GetDone returns the Done field value
func (o *ListResponseEndpointOut) GetDone() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *ListResponseEndpointOut) GetDoneOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *ListResponseEndpointOut) SetDone(v bool) {
	o.Done = v
}

func (o ListResponseEndpointOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Iterator != nil {
		toSerialize["iterator"] = o.Iterator
	}
	if true {
		toSerialize["done"] = o.Done
	}
	return json.Marshal(toSerialize)
}

type NullableListResponseEndpointOut struct {
	value *ListResponseEndpointOut
	isSet bool
}

func (v NullableListResponseEndpointOut) Get() *ListResponseEndpointOut {
	return v.value
}

func (v *NullableListResponseEndpointOut) Set(val *ListResponseEndpointOut) {
	v.value = val
	v.isSet = true
}

func (v NullableListResponseEndpointOut) IsSet() bool {
	return v.isSet
}

func (v *NullableListResponseEndpointOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListResponseEndpointOut(val *ListResponseEndpointOut) *NullableListResponseEndpointOut {
	return &NullableListResponseEndpointOut{value: val, isSet: true}
}

func (v NullableListResponseEndpointOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListResponseEndpointOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

