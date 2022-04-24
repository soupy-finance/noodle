package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyChainContracts     = []byte("ChainContracts")
	DefaultChainContracts = map[string]string{"ethereum": "0x0"}
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	chainContracts map[string]string,
) Params {
	return Params{
		ChainContracts: chainContracts,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultChainContracts,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyChainContracts, &p.ChainContracts, validateChainContracts),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateChainContracts(p.ChainContracts); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateChainContracts validates the ChainContracts param
func validateChainContracts(v interface{}) error {
	chainContracts, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = chainContracts

	return nil
}
