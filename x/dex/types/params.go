package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMarkets = []byte("Markets")
	// TODO: Determine the default value
	DefaultMarkets = "{\"wbtc-usdc\": {\"fees\": {\"maker\": 0, \"taker\": 0}}, \"eth-usdc\": {\"fees\": {\"maker\": 0, \"taker\": 0}}}"
)

type MarketsParsed map[string]MarketParsed
type MarketParsed struct {
	Fees MarketsParsedFees
}
type MarketsParsedFees struct {
	Maker float32
	Taker float32
}

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	markets string,
) Params {
	return Params{
		Markets: markets,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMarkets,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMarkets, &p.Markets, validateMarkets),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMarkets(p.Markets); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMarkets validates the Markets param
func validateMarkets(v interface{}) error {
	markets, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = markets

	return nil
}
