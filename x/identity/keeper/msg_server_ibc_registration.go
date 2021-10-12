package keeper

import (
	"context"

	"github.com/amparks100/registry/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
)

func (k msgServer) SendIbcRegistration(goCtx context.Context, msg *types.MsgSendIbcRegistration) (*types.MsgSendIbcRegistrationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcRegistrationPacketData

	packet.Handle = msg.Handle
	packet.Sender = msg.Creator

	// Transmit the packet
	err := k.TransmitIbcRegistrationPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcRegistrationResponse{}, nil
}
