package keeper

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/amparks100/registry/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// TransmitIbcRegistrationPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcRegistrationPacket(
	ctx sdk.Context,
	packetData types.IbcRegistrationPacketData,
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

// OnRecvIbcRegistrationPacket processes packet reception
func (k Keeper) OnRecvIbcRegistrationPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcRegistrationPacketData) (packetAck types.IbcRegistrationPacketAck, err error) {
	fmt.Println("Hello receiving registration")

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	//verify handle doesnt already exist
	if k.DoesHandleExist(ctx, data.Handle) {
		return packetAck, sdkerrors.Wrap(sdkerrors.Error{}, "handle already exists")
	}

	newDsnpId := k.GetCurrentDsnpId(ctx) + 1

	//create a new registration (default to owner)
	k.SetRegistration(ctx, types.Registration{
		Index:  strconv.FormatUint(newDsnpId, 10),
		DsnpId: newDsnpId,
		Handle: data.Handle,
	})
	//update the dsnpid iterator
	k.SetDsnpId(ctx, newDsnpId)
	//create delegation with address and default role
	initDelegation := types.Delegation{
		DsnpId:  newDsnpId,
		Address: data.Sender,
		Role:    0,
	}
	k.AppendDelegation(ctx, initDelegation)

	return packetAck, nil
}

// OnAcknowledgementIbcRegistrationPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcRegistrationPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcRegistrationPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcRegistrationPacketAck

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

// OnTimeoutIbcRegistrationPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcRegistrationPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcRegistrationPacketData) error {

	// TODO: packet timeout logic

	return nil
}
