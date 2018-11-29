package nametransfer

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgSetName struct {
	Name  string
	Owner sdk.AccAddress
}

func NewMsgSetName(name string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Owner: owner,
	}
}

// Name Implements Msg.
func (msg MsgSetName) Route() string { return "nametransfer" }

// Type Implements Msg.
func (msg MsgSetName) Type() string { return "set_name" }

// ValdateBasic Implements Msg.
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetName) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
