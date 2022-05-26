package contract

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/poanetwork/tokenbridge-monitor/config"
	"github.com/poanetwork/tokenbridge-monitor/contract/abi"
	"github.com/poanetwork/tokenbridge-monitor/contract/bridgeabi"
	"github.com/poanetwork/tokenbridge-monitor/ethclient"
)

type BridgeContract struct {
	*Contract
}

func NewBridgeContract(client ethclient.Client, addr common.Address, mode config.BridgeMode) *BridgeContract {
	return &BridgeContract{NewContract(client, addr, getBridgeABI(mode))}
}

func getBridgeABI(mode config.BridgeMode) abi.ABI {
	if mode == config.BridgeModeErcToNative {
		return bridgeabi.ErcToNativeABI
	}
	return bridgeabi.ArbitraryMessageABI
}

func (c *BridgeContract) ValidatorContractAddress(ctx context.Context) (common.Address, error) {
	res, err := c.Call(ctx, "validatorContract")
	if err != nil {
		return common.Address{}, fmt.Errorf("cannot obtain validator contract address: %w", err)
	}
	return common.BytesToAddress(res), nil
}
