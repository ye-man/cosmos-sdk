package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	// ModuleName is the module name constant used in many places
	ModuleName = "distribution"

	// StoreKey is the store key string for distribution
	StoreKey = ModuleName

	// TStoreKey is the transient store key for distribution
	TStoreKey = "transient_" + ModuleName

	// RouterKey is the message route for distribution
	RouterKey = ModuleName

	// QuerierRoute is the querier route for distribution
	QuerierRoute = ModuleName
)

// ModuleAddress distribution module account address
var ModuleAddress = sdk.AccAddress(crypto.AddressHash([]byte(ModuleName)))
