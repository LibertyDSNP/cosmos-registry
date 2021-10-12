package types

// ValidateBasic is used for validating the packet
func (p IbcPermissionUpdatePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcPermissionUpdatePacketData) GetBytes() ([]byte, error) {
	var modulePacket IdentityPacketData

	modulePacket.Packet = &IdentityPacketData_IbcPermissionUpdatePacket{&p}

	return modulePacket.Marshal()
}
