package keeper

import (
	"context"

	"github.com/amparks100/registry/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
)

func (k msgServer) SendIbcPermissionUpdate(goCtx context.Context, msg *types.MsgSendIbcPermissionUpdate) (*types.MsgSendIbcPermissionUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcPermissionUpdatePacketData

	packet.PublicKey = msg.PublicKey
	packet.Role = msg.Role
	packet.Sender = msg.Creator
	packet.DsnpId = msg.DsnpId

	// Transmit the packet
	err := k.TransmitIbcPermissionUpdatePacket(
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

	return &types.MsgSendIbcPermissionUpdateResponse{}, nil
}
