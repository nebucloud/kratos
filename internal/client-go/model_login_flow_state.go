/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.

API version:
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// LoginFlowState The experimental state represents the state of a login flow. This field is EXPERIMENTAL and subject to change!
type LoginFlowState string

// List of loginFlowState
const (
	LOGINFLOWSTATE_CHOOSE_METHOD    LoginFlowState = "choose_method"
	LOGINFLOWSTATE_SENT_EMAIL       LoginFlowState = "sent_email"
	LOGINFLOWSTATE_PASSED_CHALLENGE LoginFlowState = "passed_challenge"
)

// All allowed values of LoginFlowState enum
var AllowedLoginFlowStateEnumValues = []LoginFlowState{
	"choose_method",
	"sent_email",
	"passed_challenge",
}

func (v *LoginFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoginFlowState(value)
	for _, existing := range AllowedLoginFlowStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoginFlowState", value)
}

// NewLoginFlowStateFromValue returns a pointer to a valid LoginFlowState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoginFlowStateFromValue(v string) (*LoginFlowState, error) {
	ev := LoginFlowState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LoginFlowState: valid values are %v", v, AllowedLoginFlowStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoginFlowState) IsValid() bool {
	for _, existing := range AllowedLoginFlowStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to loginFlowState value
func (v LoginFlowState) Ptr() *LoginFlowState {
	return &v
}

type NullableLoginFlowState struct {
	value *LoginFlowState
	isSet bool
}

func (v NullableLoginFlowState) Get() *LoginFlowState {
	return v.value
}

func (v *NullableLoginFlowState) Set(val *LoginFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginFlowState(val *LoginFlowState) *NullableLoginFlowState {
	return &NullableLoginFlowState{value: val, isSet: true}
}

func (v NullableLoginFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
