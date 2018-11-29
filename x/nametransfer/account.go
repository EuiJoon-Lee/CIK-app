package nametransfer

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
)

type NameAccount struct {
	Name          string         `json:"name"`
	Address       sdk.AccAddress `json:"address"`
	Coins         sdk.Coins      `json:"coins"`
	PubKey        crypto.PubKey  `json:"public_key"`
	AccountNumber int64          `json:"account_number"`
	Sequence      int64          `json:"sequence"`
}

// Prototype function for NameAccount
func ProtoNameAccount() auth.Account {
	return &NameAccount{}
}

func NewNameAccountWithAddress(addr sdk.AccAddress) NameAccount {
	return NameAccount{
		Address: addr,
	}
}

// Implements sdk.Account.
func (acc NameAccount) GetAddress() sdk.AccAddress {
	return acc.Address
}

// Implements sdk.Account.
func (acc *NameAccount) SetAddress(addr sdk.AccAddress) error {
	if len(acc.Address) != 0 {
		return errors.New("cannot override NameAccount address")
	}
	acc.Address = addr
	return nil
}

// Implements sdk.Account.
func (acc NameAccount) GetPubKey() crypto.PubKey {
	return acc.PubKey
}

// Implements sdk.Account.
func (acc *NameAccount) SetPubKey(pubKey crypto.PubKey) error {
	acc.PubKey = pubKey
	return nil
}

// Implements sdk.Account.
func (acc *NameAccount) GetCoins() sdk.Coins {
	return acc.Coins
}

// Implements sdk.Account.
func (acc *NameAccount) SetCoins(coins sdk.Coins) error {
	acc.Coins = coins
	return nil
}

// Implements Account
func (acc *NameAccount) GetAccountNumber() int64 {
	return acc.AccountNumber
}

// Implements Account
func (acc *NameAccount) SetAccountNumber(accNumber int64) error {
	acc.AccountNumber = accNumber
	return nil
}

// Implements sdk.Account.
func (acc *NameAccount) GetSequence() int64 {
	return acc.Sequence
}

// Implements sdk.Account.
func (acc *NameAccount) SetSequence(seq int64) error {
	acc.Sequence = seq
	return nil
}

func (acc *NameAccount) SetName(name string) {
	acc.Name = name
}

func (acc *NameAccount) GetName() string {
	return acc.Name
}
