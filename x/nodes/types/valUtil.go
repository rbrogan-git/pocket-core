package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	"net/url"
	"strconv"
	"strings"
)

// Validators is a collection of Validator
type Validators []Validator

func (v Validators) String() (out string) {
	for _, val := range v {
		out += val.String() + "\n\n"
	}
	return strings.TrimSpace(out)
}

// String returns a human readable string representation of a validator.
func (v Validator) String() string {
	return fmt.Sprintf("Address:\t\t%s\nPublic Key:\t\t%s\nJailed:\t\t\t%v\nStatus:\t\t\t%s\nTokens:\t\t\t%s\n"+
		"ServiceURL:\t\t%s\nChains:\t\t\t%v\nUnstaking Completion Time:\t\t%v"+
		"\n----\n",
		v.Address, v.PublicKey.RawString(), v.Jailed, v.Status, v.StakedTokens, v.ServiceURL, v.Chains, v.UnstakingCompletionTime,
	)
}

// Returns the proto endcoding of a validator
func MarshalValidator(cdc *codec.ProtoCodec, validator Validator) ([]byte, error) {

	return cdc.MarshalBinaryLengthPrefixed(&ValidatorProto{
		Address:                 validator.Address,
		PublicKey:               validator.PublicKey.RawString(),
		Jailed:                  validator.Jailed,
		Status:                  validator.Status,
		Chains:                  validator.Chains,
		ServiceURL:              validator.ServiceURL,
		StakedTokens:            validator.StakedTokens,
		UnstakingCompletionTime: validator.UnstakingCompletionTime,
	})
}

// MUST decode the validator from the bytes
func UnmarshalValidator(cdc *codec.ProtoCodec, valBytes []byte) (v Validator, err error) {
	validator, err := UnmarshalProtoValidator(cdc, valBytes)
	if err != nil {
		return
	}
	pubkey, err := crypto.NewPublicKey(validator.PublicKey)
	if err != nil {
		return
	}
	return Validator{
		Address:                 validator.Address,
		PublicKey:               pubkey,
		Jailed:                  validator.Jailed,
		Status:                  validator.Status,
		Chains:                  validator.Chains,
		ServiceURL:              validator.ServiceURL,
		StakedTokens:            validator.StakedTokens,
		UnstakingCompletionTime: validator.UnstakingCompletionTime,
	}, nil
}

func UnmarshalProtoValidator(cdc *codec.ProtoCodec, valBytes []byte) (v ValidatorProto, err error) {
	var validator ValidatorProto
	err = cdc.UnmarshalBinaryLengthPrefixed(valBytes, &validator)
	if err != nil {
		return
	}
	return v, nil
}

// Marshals struct into JSON
func (v Validators) JSON() (out []byte, err error) {
	// each element should be a JSON
	return json.Marshal(v)
}

// MarshalJSON marshals the validator to JSON using Hex
func (v Validator) MarshalJSON() ([]byte, error) {
	return ModuleCdc.MarshalJSON(ValidatorProto{
		Address:                 v.Address,
		PublicKey:               v.PublicKey.RawString(),
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		ServiceURL:              v.ServiceURL,
		Chains:                  v.Chains,
		StakedTokens:            v.StakedTokens,
		UnstakingCompletionTime: v.UnstakingCompletionTime,
	})
}

// UnmarshalJSON unmarshals the validator from JSON using Hex
func (v *Validator) UnmarshalJSON(data []byte) error {
	bv := &ValidatorProto{}
	if err := ModuleCdc.UnmarshalJSON(data, bv); err != nil {
		return err
	}
	publicKey, err := crypto.NewPublicKey(bv.PublicKey)
	if err != nil {
		return err
	}
	*v = Validator{
		Address:                 bv.Address,
		PublicKey:               publicKey,
		Jailed:                  bv.Jailed,
		Chains:                  bv.Chains,
		ServiceURL:              bv.ServiceURL,
		StakedTokens:            bv.StakedTokens,
		Status:                  bv.Status,
		UnstakingCompletionTime: bv.UnstakingCompletionTime,
	}
	return nil
}

// TODO shared code among modules below

const (
	httpsPrefix = "https://"
	httpPrefix  = "http://"
	colon       = ":"
	period      = "."
)

func ValidateServiceURL(u string) sdk.Error {
	u = strings.ToLower(u)
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return ErrInvalidServiceURL(ModuleName, err)
	}
	if u[:8] != httpsPrefix && u[:7] != httpPrefix {
		return ErrInvalidServiceURL(ModuleName, fmt.Errorf("invalid url prefix"))
	}
	temp := strings.Split(u, colon)
	if len(temp) != 3 {
		return ErrInvalidServiceURL(ModuleName, fmt.Errorf("needs :port"))
	}
	port, err := strconv.Atoi(temp[2])
	if err != nil {
		return ErrInvalidServiceURL(ModuleName, fmt.Errorf("invalid port, cant convert to integer"))
	}
	if port > 65535 || port < 0 {
		return ErrInvalidServiceURL(ModuleName, fmt.Errorf("invalid port, out of valid port range"))
	}
	if !strings.Contains(u, period) {
		return ErrInvalidServiceURL(ModuleName, fmt.Errorf("must contain one '.'"))
	}
	return nil
}

const (
	NetworkIdentifierLength = 2
)

func ValidateNetworkIdentifier(chain string) sdk.Error {
	// decode string into bz
	h, err := hex.DecodeString(chain)
	if err != nil {
		return ErrInvalidNetworkIdentifier(ModuleName, err)
	}
	// ensure length isn't 0
	if len(h) == 0 {
		return ErrInvalidNetworkIdentifier(ModuleName, fmt.Errorf("net id is empty"))
	}
	// ensure length
	if len(h) > NetworkIdentifierLength {
		return ErrInvalidNetworkIdentifier(ModuleName, fmt.Errorf("net id length is > %d", NetworkIdentifierLength))
	}
	return nil
}

func (v ValidatorProto) IsStaked() bool             { return v.GetStatus().Equal(sdk.Staked) }
func (v ValidatorProto) IsUnstaked() bool           { return v.GetStatus().Equal(sdk.Unstaked) }
func (v ValidatorProto) IsUnstaking() bool          { return v.GetStatus().Equal(sdk.Unstaking) }
func (v ValidatorProto) IsJailed() bool             { return v.Jailed }
func (v ValidatorProto) GetStatus() sdk.StakeStatus { return v.Status }
func (v ValidatorProto) GetAddress() sdk.Address    { return v.Address }
func (v ValidatorProto) GetTokens() sdk.Int         { return v.StakedTokens }

// MarshalJSON marshals the validator to JSON using Hex
func (v ValidatorProto) FromProto() Validator {
	pubkey, _ := crypto.NewPublicKey(v.PublicKey)
	return Validator{
		Address:                 v.Address,
		PublicKey:               pubkey,
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		ServiceURL:              v.ServiceURL,
		Chains:                  v.Chains,
		StakedTokens:            v.StakedTokens,
		UnstakingCompletionTime: v.UnstakingCompletionTime,
	}
}

// MarshalJSON marshals the validator to JSON using Hex
func (v Validator) ToProto() ValidatorProto {
	return ValidatorProto{
		Address:                 v.Address,
		PublicKey:               v.PublicKey.RawString(),
		Jailed:                  v.Jailed,
		Status:                  v.Status,
		ServiceURL:              v.ServiceURL,
		Chains:                  v.Chains,
		StakedTokens:            v.StakedTokens,
		UnstakingCompletionTime: v.UnstakingCompletionTime,
	}
}
