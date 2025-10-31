# OrderConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RespectedSequences** | Pointer to [**[]RespectedOrderSequence**](RespectedOrderSequence.md) | A list of sequences that must be respected when scheduling routes. Orders belonging to a category that occurs earlier in the sequence must be delivered in the route before an order belonging to a category later in the sequence can be picked up. | [optional] [default to []]
**LoadingIncompatibilities** | Pointer to [**[]OrderLoadingIncompatibilityConstraint**](OrderLoadingIncompatibilityConstraint.md) | A list of constraints that prevent orders to be loaded or unloaded while other orders are loaded in the vehicle or in a compartment of a vehicle. | [optional] [default to []]

## Methods

### NewOrderConstraints

`func NewOrderConstraints() *OrderConstraints`

NewOrderConstraints instantiates a new OrderConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderConstraintsWithDefaults

`func NewOrderConstraintsWithDefaults() *OrderConstraints`

NewOrderConstraintsWithDefaults instantiates a new OrderConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRespectedSequences

`func (o *OrderConstraints) GetRespectedSequences() []RespectedOrderSequence`

GetRespectedSequences returns the RespectedSequences field if non-nil, zero value otherwise.

### GetRespectedSequencesOk

`func (o *OrderConstraints) GetRespectedSequencesOk() (*[]RespectedOrderSequence, bool)`

GetRespectedSequencesOk returns a tuple with the RespectedSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectedSequences

`func (o *OrderConstraints) SetRespectedSequences(v []RespectedOrderSequence)`

SetRespectedSequences sets RespectedSequences field to given value.

### HasRespectedSequences

`func (o *OrderConstraints) HasRespectedSequences() bool`

HasRespectedSequences returns a boolean if a field has been set.

### GetLoadingIncompatibilities

`func (o *OrderConstraints) GetLoadingIncompatibilities() []OrderLoadingIncompatibilityConstraint`

GetLoadingIncompatibilities returns the LoadingIncompatibilities field if non-nil, zero value otherwise.

### GetLoadingIncompatibilitiesOk

`func (o *OrderConstraints) GetLoadingIncompatibilitiesOk() (*[]OrderLoadingIncompatibilityConstraint, bool)`

GetLoadingIncompatibilitiesOk returns a tuple with the LoadingIncompatibilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadingIncompatibilities

`func (o *OrderConstraints) SetLoadingIncompatibilities(v []OrderLoadingIncompatibilityConstraint)`

SetLoadingIncompatibilities sets LoadingIncompatibilities field to given value.

### HasLoadingIncompatibilities

`func (o *OrderConstraints) HasLoadingIncompatibilities() bool`

HasLoadingIncompatibilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


