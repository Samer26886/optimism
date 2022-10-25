// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":27850,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":27853,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":28464,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":27722,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":27842,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1876,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1878,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1880,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcher\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063f2fde38b1161005b578063f2fde38b146101a5578063f45e65d8146101b8578063ffa1ad74146101c157600080fd5b80638da5cb5b14610161578063935f029e1461017f578063c4d66de81461019257600080fd5b80630c18c162116100b25780630c18c1621461012d57806354fd4d5014610144578063715018a61461015957600080fd5b8063025a3a29146100ce578063058a8e5f14610118575b600080fd5b6067546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61012b6101263660046109a9565b6101c9565b005b61013660655481565b60405190815260200161010f565b61014c610292565b60405161010f9190610a60565b61012b610335565b60335473ffffffffffffffffffffffffffffffffffffffff166100ee565b61012b61018d366004610a73565b610349565b61012b6101a03660046109a9565b6103e2565b61012b6101b33660046109a9565b610583565b61013660665481565b610136600081565b6101d161063a565b606780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811790915560408051602081019290925260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290506000807f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516102869190610a60565b60405180910390a35050565b60606102bd7f00000000000000000000000000000000000000000000000000000000000000006106bb565b6102e67f00000000000000000000000000000000000000000000000000000000000000006106bb565b61030f7f00000000000000000000000000000000000000000000000000000000000000006106bb565b60405160200161032193929190610a95565b604051602081830303815290604052905090565b61033d61063a565b61034760006107f8565b565b61035161063a565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516103d59190610a60565b60405180910390a3505050565b600054610100900460ff16158080156104025750600054600160ff909116105b8061041c5750303b15801561041c575060005460ff166001145b6104ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561050b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61051361086f565b61051c82610583565b801561057f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b61058b61063a565b73ffffffffffffffffffffffffffffffffffffffff811661062e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104a4565b610637816107f8565b50565b60335473ffffffffffffffffffffffffffffffffffffffff163314610347576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104a4565b6060816000036106fe57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610728578061071281610b3a565b91506107219050600a83610ba1565b9150610702565b60008167ffffffffffffffff81111561074357610743610bb5565b6040519080825280601f01601f19166020018201604052801561076d576020820181803683370190505b5090505b84156107f057610782600183610be4565b915061078f600a86610bfb565b61079a906030610c0f565b60f81b8183815181106107af576107af610c27565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506107e9600a86610ba1565b9450610771565b949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610906576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a4565b610347600054610100900460ff166109a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104a4565b610347336107f8565b6000602082840312156109bb57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146109df57600080fd5b9392505050565b60005b83811015610a015781810151838201526020016109e9565b83811115610a10576000848401525b50505050565b60008151808452610a2e8160208601602086016109e6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006109df6020830184610a16565b60008060408385031215610a8657600080fd5b50508035926020909101359150565b60008451610aa78184602089016109e6565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610ae3816001850160208a016109e6565b60019201918201528351610afe8160028401602088016109e6565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b6b57610b6b610b0b565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610bb057610bb0610b72565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610bf657610bf6610b0b565b500390565b600082610c0a57610c0a610b72565b500690565b60008219821115610c2257610c22610b0b565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
