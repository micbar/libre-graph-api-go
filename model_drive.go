/*
Open Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
)

// Drive The user's HomeDrive. Read-only.
type Drive struct {
	BaseItem
	// Describes the type of drive represented by this resource. Values are \"personal\" for users home spaces, \"project\" or \"share\". Read-only.
	DriveType *string `json:"driveType,omitempty"`
	Owner *IdentitySet `json:"owner,omitempty"`
	Quota *Quota `json:"quota,omitempty"`
	// All items contained in the drive. Read-only. Nullable.
	Items *[]DriveItem `json:"items,omitempty"`
	Root *DriveItem `json:"root,omitempty"`
}

// NewDrive instantiates a new Drive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrive() *Drive {
	this := Drive{}
	return &this
}

// NewDriveWithDefaults instantiates a new Drive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriveWithDefaults() *Drive {
	this := Drive{}
	return &this
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *Drive) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *Drive) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *Drive) SetDriveType(v string) {
	o.DriveType = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Drive) GetOwner() IdentitySet {
	if o == nil || o.Owner == nil {
		var ret IdentitySet
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetOwnerOk() (*IdentitySet, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Drive) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given IdentitySet and assigns it to the Owner field.
func (o *Drive) SetOwner(v IdentitySet) {
	o.Owner = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *Drive) GetQuota() Quota {
	if o == nil || o.Quota == nil {
		var ret Quota
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetQuotaOk() (*Quota, bool) {
	if o == nil || o.Quota == nil {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *Drive) HasQuota() bool {
	if o != nil && o.Quota != nil {
		return true
	}

	return false
}

// SetQuota gets a reference to the given Quota and assigns it to the Quota field.
func (o *Drive) SetQuota(v Quota) {
	o.Quota = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Drive) GetItems() []DriveItem {
	if o == nil || o.Items == nil {
		var ret []DriveItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetItemsOk() (*[]DriveItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Drive) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DriveItem and assigns it to the Items field.
func (o *Drive) SetItems(v []DriveItem) {
	o.Items = &v
}

// GetRoot returns the Root field value if set, zero value otherwise.
func (o *Drive) GetRoot() DriveItem {
	if o == nil || o.Root == nil {
		var ret DriveItem
		return ret
	}
	return *o.Root
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Drive) GetRootOk() (*DriveItem, bool) {
	if o == nil || o.Root == nil {
		return nil, false
	}
	return o.Root, true
}

// HasRoot returns a boolean if a field has been set.
func (o *Drive) HasRoot() bool {
	if o != nil && o.Root != nil {
		return true
	}

	return false
}

// SetRoot gets a reference to the given DriveItem and assigns it to the Root field.
func (o *Drive) SetRoot(v DriveItem) {
	o.Root = &v
}

func (o Drive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseItem, errBaseItem := json.Marshal(o.BaseItem)
	if errBaseItem != nil {
		return []byte{}, errBaseItem
	}
	errBaseItem = json.Unmarshal([]byte(serializedBaseItem), &toSerialize)
	if errBaseItem != nil {
		return []byte{}, errBaseItem
	}
	if o.DriveType != nil {
		toSerialize["driveType"] = o.DriveType
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Quota != nil {
		toSerialize["quota"] = o.Quota
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Root != nil {
		toSerialize["root"] = o.Root
	}
	return json.Marshal(toSerialize)
}

type NullableDrive struct {
	value *Drive
	isSet bool
}

func (v NullableDrive) Get() *Drive {
	return v.value
}

func (v *NullableDrive) Set(val *Drive) {
	v.value = val
	v.isSet = true
}

func (v NullableDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrive(val *Drive) *NullableDrive {
	return &NullableDrive{value: val, isSet: true}
}

func (v NullableDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

