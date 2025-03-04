package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type Fliter struct {
	Address   common.Address `json:"address,omitempty"`
	FromBlock CallMsgBigInt  `json:"fromBlock,omitempty"`
	ToBlock   CallMsgBigInt  `json:"toBlock,omitempty"`
	Topics    []string       `json:"topics,omitempty"`
}

type Event struct {
	Address         common.Address `json:"address"`
	BlockHash       common.Hash    `json:"blockHash"`
	BlockNumber     string         `json:"blockNumber"`
	Topics          []string       `json:"topics,omitempty"`
	TransactionHash common.Hash    `json:"transactionHash"`
	Data            string         `json:"data,omitempty"`
}
