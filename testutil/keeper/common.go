package keeper

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func (kf KeeperFactory) ParamsKeeper() paramskeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	transientStoreKey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	kf.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, kf.DB)
	kf.StateStore.MountStoreWithDB(transientStoreKey, sdk.StoreTypeTransient, kf.DB)

	paramsKeeper := paramskeeper.NewKeeper(kf.EncodingConfig.Marshaler, kf.EncodingConfig.Amino, storeKey, transientStoreKey)

	return paramsKeeper
}

func (kf KeeperFactory) AccountKeeper(paramsKeeper paramskeeper.Keeper, maccPerms map[string][]string) authkeeper.AccountKeeper {
	storeKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	kf.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, kf.DB)

	subspace := paramsKeeper.Subspace(authtypes.ModuleName)
	accountKeeper := authkeeper.NewAccountKeeper(
		kf.EncodingConfig.Marshaler, storeKey, subspace, authtypes.ProtoBaseAccount, maccPerms,
	)

	return accountKeeper
}

func (kf KeeperFactory) BankKeeper(paramsKeeper paramskeeper.Keeper, accountKeeper authkeeper.AccountKeeper, blockedMaccAddresses map[string]bool) bankkeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	kf.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, kf.DB)

	subspace := paramsKeeper.Subspace(banktypes.ModuleName)
	bankKeeper := bankkeeper.NewBaseKeeper(
		kf.EncodingConfig.Marshaler, storeKey, accountKeeper, subspace, blockedMaccAddresses,
	)

	return bankKeeper
}

func (kf KeeperFactory) CapabilityKeeper() capabilitykeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(capabilitytypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(capabilitytypes.MemStoreKey)
	kf.StateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, kf.DB)
	kf.StateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)

	capabilityKeeper := capabilitykeeper.NewKeeper(
		kf.EncodingConfig.Marshaler, storeKey, memStoreKey,
	)

	return *capabilityKeeper
}
