package types

// ValidateBasic is used for validating the packet
func (p IbcRegistrationPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcRegistrationPacketData) GetBytes() ([]byte, error) {
	var modulePacket IdentityPacketData

	modulePacket.Packet = &IdentityPacketData_IbcRegistrationPacket{&p}

	return modulePacket.Marshal()
}
