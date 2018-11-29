package nametransfer

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	accountKeeper auth.AccountKeeper
	coinKeeper    bank.Keeper

	namesStoreKey sdk.StoreKey

	cdc *codec.Codec
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(accountKeeper auth.AccountKeeper, coinKeeper bank.Keeper, namesStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		accountKeeper: accountKeeper,
		coinKeeper:    coinKeeper,
		namesStoreKey: namesStoreKey,
		cdc:           cdc,
	}
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, owner sdk.AccAddress) sdk.Result {
	store := ctx.KVStore(k.namesStoreKey)
	account := k.accountKeeper.GetAccount(ctx, owner)
	switch account := account.(type) {
	case *NameAccount:
		account.SetName(name)
		k.accountKeeper.SetAccount(ctx, account)

		store.Set([]byte(name), account.GetAddress())
		return sdk.Result{}
	default:
		errMsg := fmt.Sprintf("Unrecognized account type")
		return sdk.ErrUnknownRequest(errMsg).Result()
	}

	return sdk.ErrUnknownRequest("Unexpected error").Result()
}
