package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:           PortID,
		RegistrationList: []Registration{},
		DelegationList:   []Delegation{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in registration
	registrationIndexMap := make(map[string]struct{})

	for _, elem := range gs.RegistrationList {
		index := string(RegistrationKey(elem.Index))
		if _, ok := registrationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for registration")
		}
		registrationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in delegations
	delegationsIndexMap := make(map[string]struct{})

	for _, elem := range gs.DelegationList {
		index := string(DelegationKey(elem.Index))
		if _, ok := delegationsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for delegations")
		}
		delegationsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
