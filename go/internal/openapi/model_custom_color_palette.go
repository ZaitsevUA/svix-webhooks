/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CustomColorPalette struct for CustomColorPalette
type CustomColorPalette struct {
	BackgroundHover NullableString `json:"backgroundHover,omitempty"`
	BackgroundPrimary NullableString `json:"backgroundPrimary,omitempty"`
	BackgroundSecondary NullableString `json:"backgroundSecondary,omitempty"`
	InteractiveAccent NullableString `json:"interactiveAccent,omitempty"`
	TextDanger NullableString `json:"textDanger,omitempty"`
	TextPrimary NullableString `json:"textPrimary,omitempty"`
}

// NewCustomColorPalette instantiates a new CustomColorPalette object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomColorPalette() *CustomColorPalette {
	this := CustomColorPalette{}
	return &this
}

// NewCustomColorPaletteWithDefaults instantiates a new CustomColorPalette object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomColorPaletteWithDefaults() *CustomColorPalette {
	this := CustomColorPalette{}
	return &this
}

// GetBackgroundHover returns the BackgroundHover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetBackgroundHover() string {
	if o == nil || o.BackgroundHover.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackgroundHover.Get()
}

// GetBackgroundHoverOk returns a tuple with the BackgroundHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetBackgroundHoverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackgroundHover.Get(), o.BackgroundHover.IsSet()
}

// HasBackgroundHover returns a boolean if a field has been set.
func (o *CustomColorPalette) HasBackgroundHover() bool {
	if o != nil && o.BackgroundHover.IsSet() {
		return true
	}

	return false
}

// SetBackgroundHover gets a reference to the given NullableString and assigns it to the BackgroundHover field.
func (o *CustomColorPalette) SetBackgroundHover(v string) {
	o.BackgroundHover.Set(&v)
}
// SetBackgroundHoverNil sets the value for BackgroundHover to be an explicit nil
func (o *CustomColorPalette) SetBackgroundHoverNil() {
	o.BackgroundHover.Set(nil)
}

// UnsetBackgroundHover ensures that no value is present for BackgroundHover, not even an explicit nil
func (o *CustomColorPalette) UnsetBackgroundHover() {
	o.BackgroundHover.Unset()
}

// GetBackgroundPrimary returns the BackgroundPrimary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetBackgroundPrimary() string {
	if o == nil || o.BackgroundPrimary.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackgroundPrimary.Get()
}

// GetBackgroundPrimaryOk returns a tuple with the BackgroundPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetBackgroundPrimaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackgroundPrimary.Get(), o.BackgroundPrimary.IsSet()
}

// HasBackgroundPrimary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasBackgroundPrimary() bool {
	if o != nil && o.BackgroundPrimary.IsSet() {
		return true
	}

	return false
}

// SetBackgroundPrimary gets a reference to the given NullableString and assigns it to the BackgroundPrimary field.
func (o *CustomColorPalette) SetBackgroundPrimary(v string) {
	o.BackgroundPrimary.Set(&v)
}
// SetBackgroundPrimaryNil sets the value for BackgroundPrimary to be an explicit nil
func (o *CustomColorPalette) SetBackgroundPrimaryNil() {
	o.BackgroundPrimary.Set(nil)
}

// UnsetBackgroundPrimary ensures that no value is present for BackgroundPrimary, not even an explicit nil
func (o *CustomColorPalette) UnsetBackgroundPrimary() {
	o.BackgroundPrimary.Unset()
}

// GetBackgroundSecondary returns the BackgroundSecondary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetBackgroundSecondary() string {
	if o == nil || o.BackgroundSecondary.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackgroundSecondary.Get()
}

// GetBackgroundSecondaryOk returns a tuple with the BackgroundSecondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetBackgroundSecondaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackgroundSecondary.Get(), o.BackgroundSecondary.IsSet()
}

// HasBackgroundSecondary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasBackgroundSecondary() bool {
	if o != nil && o.BackgroundSecondary.IsSet() {
		return true
	}

	return false
}

// SetBackgroundSecondary gets a reference to the given NullableString and assigns it to the BackgroundSecondary field.
func (o *CustomColorPalette) SetBackgroundSecondary(v string) {
	o.BackgroundSecondary.Set(&v)
}
// SetBackgroundSecondaryNil sets the value for BackgroundSecondary to be an explicit nil
func (o *CustomColorPalette) SetBackgroundSecondaryNil() {
	o.BackgroundSecondary.Set(nil)
}

// UnsetBackgroundSecondary ensures that no value is present for BackgroundSecondary, not even an explicit nil
func (o *CustomColorPalette) UnsetBackgroundSecondary() {
	o.BackgroundSecondary.Unset()
}

// GetInteractiveAccent returns the InteractiveAccent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetInteractiveAccent() string {
	if o == nil || o.InteractiveAccent.Get() == nil {
		var ret string
		return ret
	}
	return *o.InteractiveAccent.Get()
}

// GetInteractiveAccentOk returns a tuple with the InteractiveAccent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetInteractiveAccentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InteractiveAccent.Get(), o.InteractiveAccent.IsSet()
}

// HasInteractiveAccent returns a boolean if a field has been set.
func (o *CustomColorPalette) HasInteractiveAccent() bool {
	if o != nil && o.InteractiveAccent.IsSet() {
		return true
	}

	return false
}

// SetInteractiveAccent gets a reference to the given NullableString and assigns it to the InteractiveAccent field.
func (o *CustomColorPalette) SetInteractiveAccent(v string) {
	o.InteractiveAccent.Set(&v)
}
// SetInteractiveAccentNil sets the value for InteractiveAccent to be an explicit nil
func (o *CustomColorPalette) SetInteractiveAccentNil() {
	o.InteractiveAccent.Set(nil)
}

// UnsetInteractiveAccent ensures that no value is present for InteractiveAccent, not even an explicit nil
func (o *CustomColorPalette) UnsetInteractiveAccent() {
	o.InteractiveAccent.Unset()
}

// GetTextDanger returns the TextDanger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetTextDanger() string {
	if o == nil || o.TextDanger.Get() == nil {
		var ret string
		return ret
	}
	return *o.TextDanger.Get()
}

// GetTextDangerOk returns a tuple with the TextDanger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetTextDangerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TextDanger.Get(), o.TextDanger.IsSet()
}

// HasTextDanger returns a boolean if a field has been set.
func (o *CustomColorPalette) HasTextDanger() bool {
	if o != nil && o.TextDanger.IsSet() {
		return true
	}

	return false
}

// SetTextDanger gets a reference to the given NullableString and assigns it to the TextDanger field.
func (o *CustomColorPalette) SetTextDanger(v string) {
	o.TextDanger.Set(&v)
}
// SetTextDangerNil sets the value for TextDanger to be an explicit nil
func (o *CustomColorPalette) SetTextDangerNil() {
	o.TextDanger.Set(nil)
}

// UnsetTextDanger ensures that no value is present for TextDanger, not even an explicit nil
func (o *CustomColorPalette) UnsetTextDanger() {
	o.TextDanger.Unset()
}

// GetTextPrimary returns the TextPrimary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorPalette) GetTextPrimary() string {
	if o == nil || o.TextPrimary.Get() == nil {
		var ret string
		return ret
	}
	return *o.TextPrimary.Get()
}

// GetTextPrimaryOk returns a tuple with the TextPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorPalette) GetTextPrimaryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TextPrimary.Get(), o.TextPrimary.IsSet()
}

// HasTextPrimary returns a boolean if a field has been set.
func (o *CustomColorPalette) HasTextPrimary() bool {
	if o != nil && o.TextPrimary.IsSet() {
		return true
	}

	return false
}

// SetTextPrimary gets a reference to the given NullableString and assigns it to the TextPrimary field.
func (o *CustomColorPalette) SetTextPrimary(v string) {
	o.TextPrimary.Set(&v)
}
// SetTextPrimaryNil sets the value for TextPrimary to be an explicit nil
func (o *CustomColorPalette) SetTextPrimaryNil() {
	o.TextPrimary.Set(nil)
}

// UnsetTextPrimary ensures that no value is present for TextPrimary, not even an explicit nil
func (o *CustomColorPalette) UnsetTextPrimary() {
	o.TextPrimary.Unset()
}

func (o CustomColorPalette) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackgroundHover.IsSet() {
		toSerialize["backgroundHover"] = o.BackgroundHover.Get()
	}
	if o.BackgroundPrimary.IsSet() {
		toSerialize["backgroundPrimary"] = o.BackgroundPrimary.Get()
	}
	if o.BackgroundSecondary.IsSet() {
		toSerialize["backgroundSecondary"] = o.BackgroundSecondary.Get()
	}
	if o.InteractiveAccent.IsSet() {
		toSerialize["interactiveAccent"] = o.InteractiveAccent.Get()
	}
	if o.TextDanger.IsSet() {
		toSerialize["textDanger"] = o.TextDanger.Get()
	}
	if o.TextPrimary.IsSet() {
		toSerialize["textPrimary"] = o.TextPrimary.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCustomColorPalette struct {
	value *CustomColorPalette
	isSet bool
}

func (v NullableCustomColorPalette) Get() *CustomColorPalette {
	return v.value
}

func (v *NullableCustomColorPalette) Set(val *CustomColorPalette) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomColorPalette) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomColorPalette) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomColorPalette(val *CustomColorPalette) *NullableCustomColorPalette {
	return &NullableCustomColorPalette{value: val, isSet: true}
}

func (v NullableCustomColorPalette) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomColorPalette) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


