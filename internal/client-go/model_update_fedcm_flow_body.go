/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.
 *
 * API version:
 * Contact: office@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateFedcmFlowBody struct for UpdateFedcmFlowBody
type UpdateFedcmFlowBody struct {
	// CSRFToken is the anti-CSRF token.
	CsrfToken string `json:"csrf_token"`
	// Nonce is the nonce that was used in the `navigator.credentials.get` call. If specified, it must match the `nonce` claim in the token.
	Nonce *string `json:"nonce,omitempty"`
	// Token contains the result of `navigator.credentials.get`.
	Token string `json:"token"`
}

// NewUpdateFedcmFlowBody instantiates a new UpdateFedcmFlowBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFedcmFlowBody(csrfToken string, token string) *UpdateFedcmFlowBody {
	this := UpdateFedcmFlowBody{}
	this.CsrfToken = csrfToken
	this.Token = token
	return &this
}

// NewUpdateFedcmFlowBodyWithDefaults instantiates a new UpdateFedcmFlowBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFedcmFlowBodyWithDefaults() *UpdateFedcmFlowBody {
	this := UpdateFedcmFlowBody{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value
func (o *UpdateFedcmFlowBody) GetCsrfToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value
// and a boolean to check if the value has been set.
func (o *UpdateFedcmFlowBody) GetCsrfTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsrfToken, true
}

// SetCsrfToken sets field value
func (o *UpdateFedcmFlowBody) SetCsrfToken(v string) {
	o.CsrfToken = v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *UpdateFedcmFlowBody) GetNonce() string {
	if o == nil || o.Nonce == nil {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFedcmFlowBody) GetNonceOk() (*string, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *UpdateFedcmFlowBody) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *UpdateFedcmFlowBody) SetNonce(v string) {
	o.Nonce = &v
}

// GetToken returns the Token field value
func (o *UpdateFedcmFlowBody) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateFedcmFlowBody) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateFedcmFlowBody) SetToken(v string) {
	o.Token = v
}

func (o UpdateFedcmFlowBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFedcmFlowBody struct {
	value *UpdateFedcmFlowBody
	isSet bool
}

func (v NullableUpdateFedcmFlowBody) Get() *UpdateFedcmFlowBody {
	return v.value
}

func (v *NullableUpdateFedcmFlowBody) Set(val *UpdateFedcmFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFedcmFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFedcmFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFedcmFlowBody(val *UpdateFedcmFlowBody) *NullableUpdateFedcmFlowBody {
	return &NullableUpdateFedcmFlowBody{value: val, isSet: true}
}

func (v NullableUpdateFedcmFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFedcmFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
