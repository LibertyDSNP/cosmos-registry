package keeper

import (
	"errors"
	"strconv"

	"github.com/amparks100/registry/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// TransmitIbcPermissionUpdatePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcPermissionUpdatePacket(
	ctx sdk.Context,
	packetData types.IbcPermissionUpdatePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIbcPermissionUpdatePacket processes packet reception
func (k Keeper) OnRecvIbcPermissionUpdatePacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPermissionUpdatePacketData) (packetAck types.IbcPermissionUpdatePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic
	//get permission for sender and verify
	dsnpId, _ := strconv.ParseUint(data.DsnpId, 10, 64)
	delegations := k.GetAllDelegationsForDsnpId(ctx, dsnpId)
	if len(delegations) == 0 {
		ctx.Logger().Error("Cant find registration")
		return packetAck, sdkerrors.Wrap(sdkerrors.Error{}, "user does not exist")
	}
	senderPermissionIndex := -1
	permissionToUpdateIndex := -1
	ctx.Logger().Info("TEST")
	for i := range delegations {
		if delegations[i].Address == data.Sender {
			senderPermissionIndex = i
		}
		if delegations[i].Address == data.PublicKey {
			permissionToUpdateIndex = i
		}
	}
	allowed := types.DoesRoleHavePermission(uint8(delegations[senderPermissionIndex].Role), types.PERMISSION_DELEGATEADD)
	if !allowed {
		ctx.Logger().Error("Sender not allowed to change permission")
		return packetAck, sdkerrors.Wrap(sdkerrors.Error{}, "Unauthorized")
	}

	u, err := strconv.ParseUint(data.Role, 10, 32)
	if err != nil {
		ctx.Logger().Error("cant parse int", "role", data.Role)
		return packetAck, sdkerrors.Wrap(sdkerrors.Error{}, "Cant parse role input")
	}
	if permissionToUpdateIndex >= 0 {
		delegations[permissionToUpdateIndex].Role = uint32(u)
		k.SetDelegation(ctx, delegations[permissionToUpdateIndex])
	} else {
		k.AppendDelegation(ctx, types.Delegation{
			DsnpId:  dsnpId,
			Address: data.PublicKey,
			Role:    uint32(u),
		})
	}

	return packetAck, nil
}

// OnAcknowledgementIbcPermissionUpdatePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcPermissionUpdatePacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPermissionUpdatePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcPermissionUpdatePacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcPermissionUpdatePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcPermissionUpdatePacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPermissionUpdatePacketData) error {

	// TODO: packet timeout logic

	return nil
}
