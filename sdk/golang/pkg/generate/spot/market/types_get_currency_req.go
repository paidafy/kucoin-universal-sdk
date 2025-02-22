// Code generated by Kucoin Universal SDK Generator; DO NOT EDIT.

package market

// GetCurrencyReq struct for GetCurrencyReq
type GetCurrencyReq struct {
	// Path parameter, Currency
	Currency *string `json:"currency,omitempty" path:"currency" url:"-"`
	// Support for querying the chain of currency, e.g. The available value for USDT are OMNI, ERC20, TRC20. This only apply for multi-chain currency, and there is no need for single chain currency.
	Chain *string `json:"chain,omitempty" url:"chain,omitempty"`
}

// NewGetCurrencyReq instantiates a new GetCurrencyReq object
// This constructor will assign default values to properties that have it defined
func NewGetCurrencyReq() *GetCurrencyReq {
	this := GetCurrencyReq{}
	return &this
}

// NewGetCurrencyReqWithDefaults instantiates a new GetCurrencyReq object
// This constructor will only assign default values to properties that have it defined,
func NewGetCurrencyReqWithDefaults() *GetCurrencyReq {
	this := GetCurrencyReq{}
	return &this
}

func (o *GetCurrencyReq) ToMap() map[string]interface{} {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["chain"] = o.Chain
	return toSerialize
}

type GetCurrencyReqBuilder struct {
	obj *GetCurrencyReq
}

func NewGetCurrencyReqBuilder() *GetCurrencyReqBuilder {
	return &GetCurrencyReqBuilder{obj: NewGetCurrencyReqWithDefaults()}
}

// Path parameter, Currency
func (builder *GetCurrencyReqBuilder) SetCurrency(value string) *GetCurrencyReqBuilder {
	builder.obj.Currency = &value
	return builder
}

// Support for querying the chain of currency, e.g. The available value for USDT are OMNI, ERC20, TRC20. This only apply for multi-chain currency, and there is no need for single chain currency.
func (builder *GetCurrencyReqBuilder) SetChain(value string) *GetCurrencyReqBuilder {
	builder.obj.Chain = &value
	return builder
}

func (builder *GetCurrencyReqBuilder) Build() *GetCurrencyReq {
	return builder.obj
}
