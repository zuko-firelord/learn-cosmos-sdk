package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateData{}, "chat/CreateData", nil)
	cdc.RegisterConcrete(&MsgUpdateData{}, "chat/UpdateData", nil)
	cdc.RegisterConcrete(&MsgDeleteData{}, "chat/DeleteData", nil)
	cdc.RegisterConcrete(&MsgSendMsgSpace{}, "chat/SendMsgSpace", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateData{},
		&MsgUpdateData{},
		&MsgDeleteData{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendMsgSpace{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
