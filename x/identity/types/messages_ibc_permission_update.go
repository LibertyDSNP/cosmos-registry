package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIbcPermissionUpdate{}

func NewMsgSendIbcPermissionUpdate(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	dsnpId string,
	publicKey string,
	role string,
) *MsgSendIbcPermissionUpdate {
	return &MsgSendIbcPermissionUpdate{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		PublicKey:        publicKey,
		Role:             role,
		DsnpId:           dsnpId,
	}
}

func (msg *MsgSendIbcPermissionUpdate) Route() string {
	return RouterKey
}

func (msg *MsgSendIbcPermissionUpdate) Type() string {
	return "SendIbcPermissionUpdate"
}

func (msg *MsgSendIbcPermissionUpdate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendIbcPermissionUpdate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendIbcPermissionUpdate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
