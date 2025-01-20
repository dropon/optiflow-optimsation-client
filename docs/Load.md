# Load

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | **string** | Indicates the specific dimension of the load, such as its volume, weight, or size. | 
**Value** | **float64** | Represents the numeric value associated with the load&#39;s dimension. This value could be the actual measurement or quantity of the load. | 

## Methods

### NewLoad

`func NewLoad(dimension string, value float64, ) *Load`

NewLoad instantiates a new Load object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadWithDefaults

`func NewLoadWithDefaults() *Load`

NewLoadWithDefaults instantiates a new Load object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *Load) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *Load) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *Load) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetValue

`func (o *Load) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Load) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Load) SetValue(v float64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


