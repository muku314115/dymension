package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	lockuptypes "github.com/osmosis-labs/osmosis/v15/x/lockup/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
}

// TransferKeeper defines the expected transfer keeper
type TransferKeeper interface {
	HasDenomTrace(ctx sdk.Context, denomTraceHash tmbytes.HexBytes) bool
}

// ChannelKeeper defines the expected IBC channel keeper
type ChannelKeeper interface {
	GetChannelClientState(ctx sdk.Context, portID, channelID string) (string, exported.ClientState, error)
}

type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
}

type LockupKeeper interface {
	HasLock(ctx sdk.Context, owner sdk.AccAddress, denom string, duration time.Duration) bool
	CreateLock(ctx sdk.Context, owner sdk.AccAddress, coins sdk.Coins, duration time.Duration) (lockuptypes.PeriodLock, error)
	AddToExistingLock(ctx sdk.Context, owner sdk.AccAddress, coin sdk.Coin, duration time.Duration) (uint64, error)
}
