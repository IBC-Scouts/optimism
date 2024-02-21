// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IBCCrossDomainMessengerMetaData contains all meta data concerning the IBCCrossDomainMessenger contract.
var IBCCrossDomainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_CALLDATA_OVERHEAD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_CALL_OVERHEAD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_CONSTANT_OVERHEAD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_GAS_CHECK_BUFFER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_RESERVED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGas\",\"inputs\":[{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"failedMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1CrossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"messageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayMessage\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendMessage\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"successfulMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xDomainMessageSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FailedRelayedMessage\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RelayedMessage\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SentMessage\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SentMessage\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"messageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SentMessageExtension1\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001e3b38038062001e3b833981016040819052620000349162000233565b6001600160a01b0381166080526200004b62000052565b5062000265565b600054600160a81b900460ff16158080156200007b57506000546001600160a01b90910460ff16105b80620000b2575062000098306200019f60201b620012fd1760201c565b158015620000b25750600054600160a01b900460ff166001145b6200011b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff60a01b1916600160a01b179055801562000149576000805460ff60a81b1916600160a81b1790555b62000153620001ae565b80156200019c576000805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b03163b151590565b600054600160a81b900460ff166200021d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840162000112565b60cc80546001600160a01b03191661dead179055565b6000602082840312156200024657600080fd5b81516001600160a01b03811681146200025e57600080fd5b9392505050565b608051611ba56200029660003960008181610318015281816103790152818161046a0152610c540152611ba56000f3fe60806040526004361061015f5760003560e01c80638129fc1c116100c0578063a711986911610074578063b28ade2511610059578063b28ade25146103cd578063d764ad0b146103ed578063ecc704281461040057600080fd5b8063a71198691461036a578063b1b1b2091461039d57600080fd5b80638cbeeef2116100a55780638cbeeef2146101fe5780639fce812c14610306578063a4e7f8bd1461033a57600080fd5b80638129fc1c146102da57806383a74074146102ef57600080fd5b80634c1d6a69116101175780635644cfdf116100fc5780635644cfdf1461026a5780635c975abb146102805780636e296e45146102a057600080fd5b80634c1d6a69146101fe57806354fd4d501461021457600080fd5b80632828d7e8116101485780632828d7e8146101ac5780633dbb202b146101c15780633f827a5a146101d657600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b50610179604081565b6101d46101cf3660046116d7565b610465565b005b3480156101e257600080fd5b506101eb600181565b60405161ffff909116815260200161018e565b34801561020a57600080fd5b50610179619c4081565b34801561022057600080fd5b5061025d6040518060400160405280600581526020017f312e382e3000000000000000000000000000000000000000000000000000000081525081565b60405161018e91906117a7565b34801561027657600080fd5b5061017961138881565b34801561028c57600080fd5b5060005b604051901515815260200161018e565b3480156102ac57600080fd5b506102b56106c9565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b3480156102e657600080fd5b506101d46107b5565b3480156102fb57600080fd5b5061017962030d4081565b34801561031257600080fd5b506102b57f000000000000000000000000000000000000000000000000000000000000000081565b34801561034657600080fd5b506102906103553660046117c1565b60ce6020526000908152604090205460ff1681565b34801561037657600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102b5565b3480156103a957600080fd5b506102906103b83660046117c1565b60cb6020526000908152604090205460ff1681565b3480156103d957600080fd5b506101796103e83660046117da565b6109b2565b6101d46103fb36600461182e565b610a20565b34801561040c57600080fd5b5061045760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60405190815260200161018e565b61059e7f00000000000000000000000000000000000000000000000000000000000000006104948585856109b2565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061050060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161051c97969594939291906118f9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611319565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561062360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610635959493929190611958565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610798576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000547501000000000000000000000000000000000000000000900460ff1615808015610800575060005460017401000000000000000000000000000000000000000090910460ff16105b806108325750303b158015610832575060005474010000000000000000000000000000000000000000900460ff166001145b6108be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161078f565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000179055801561094457600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b61094c611371565b80156109af57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000611388619c4080603f6109ce604063ffffffff88166119d5565b6109d89190611a05565b6109e36010886119d5565b6109f09062030d40611a53565b6109fa9190611a53565b610a049190611a53565b610a0e9190611a53565b610a189190611a53565b949350505050565b60f087901c60028110610adb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a40161078f565b8061ffff16600003610bd0576000610b2c878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f925061144a915050565b600081815260cb602052604090205490915060ff1615610bce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161078f565b505b6000610c16898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061146992505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610cae57853414610c8a57610c8a611a7f565b600081815260ce602052604090205460ff1615610ca957610ca9611a7f565b610e00565b3415610d62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161078f565b600081815260ce602052604090205460ff16610e00576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161078f565b610e098761148c565b15610ebc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161078f565b600081815260cb602052604090205460ff1615610f5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161078f565b610f7c85610f6d611388619c40611a53565b67ffffffffffffffff166114e1565b1580610fa2575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156110bb57600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016110b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161078f565b50506112f4565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061114c88619c405a61110f9190611aae565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114ff92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156111e357600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26112f0565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016112f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161078f565b5050505b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b8373ffffffffffffffffffffffffffffffffffffffff167fecf079dc36711b69eae2d3ea1b148e593e45d8f663834aac87bab9ae231346f584848460405161136393929190611ac5565b60405180910390a250505050565b6000547501000000000000000000000000000000000000000000900460ff1661141c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161078f565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b600061145885858585611519565b805190602001209050949350505050565b60006114798787878787876115b2565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82163014806114db575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016115329493929190611af7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b60608686868686866040516024016115cf96959493929190611b41565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461167557600080fd5b919050565b60008083601f84011261168c57600080fd5b50813567ffffffffffffffff8111156116a457600080fd5b6020830191508360208285010111156116bc57600080fd5b9250929050565b803563ffffffff8116811461167557600080fd5b600080600080606085870312156116ed57600080fd5b6116f685611651565b9350602085013567ffffffffffffffff81111561171257600080fd5b61171e8782880161167a565b90945092506117319050604086016116c3565b905092959194509250565b6000815180845260005b8181101561176257602081850181015186830182015201611746565b81811115611774576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006117ba602083018461173c565b9392505050565b6000602082840312156117d357600080fd5b5035919050565b6000806000604084860312156117ef57600080fd5b833567ffffffffffffffff81111561180657600080fd5b6118128682870161167a565b90945092506118259050602085016116c3565b90509250925092565b600080600080600080600060c0888a03121561184957600080fd5b8735965061185960208901611651565b955061186760408901611651565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561189157600080fd5b61189d8a828b0161167a565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261194b60c0830184866118b0565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff861681526080602082015260006119886080830186886118b0565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516818304811182151516156119fc576119fc6119a6565b02949350505050565b600067ffffffffffffffff80841680611a47577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611a7657611a766119a6565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611ac057611ac06119a6565b500390565b67ffffffffffffffff84168152826020820152606060408201526000611aee606083018461173c565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611b30608083018561173c565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611b8c60c083018461173c565b9897505050505050505056fea164736f6c634300080f000a",
}

// IBCCrossDomainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IBCCrossDomainMessengerMetaData.ABI instead.
var IBCCrossDomainMessengerABI = IBCCrossDomainMessengerMetaData.ABI

// IBCCrossDomainMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IBCCrossDomainMessengerMetaData.Bin instead.
var IBCCrossDomainMessengerBin = IBCCrossDomainMessengerMetaData.Bin

// DeployIBCCrossDomainMessenger deploys a new Ethereum contract, binding an instance of IBCCrossDomainMessenger to it.
func DeployIBCCrossDomainMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, _l1CrossDomainMessenger common.Address) (common.Address, *types.Transaction, *IBCCrossDomainMessenger, error) {
	parsed, err := IBCCrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IBCCrossDomainMessengerBin), backend, _l1CrossDomainMessenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IBCCrossDomainMessenger{IBCCrossDomainMessengerCaller: IBCCrossDomainMessengerCaller{contract: contract}, IBCCrossDomainMessengerTransactor: IBCCrossDomainMessengerTransactor{contract: contract}, IBCCrossDomainMessengerFilterer: IBCCrossDomainMessengerFilterer{contract: contract}}, nil
}

// IBCCrossDomainMessenger is an auto generated Go binding around an Ethereum contract.
type IBCCrossDomainMessenger struct {
	IBCCrossDomainMessengerCaller     // Read-only binding to the contract
	IBCCrossDomainMessengerTransactor // Write-only binding to the contract
	IBCCrossDomainMessengerFilterer   // Log filterer for contract events
}

// IBCCrossDomainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBCCrossDomainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCCrossDomainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBCCrossDomainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCCrossDomainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBCCrossDomainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCCrossDomainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBCCrossDomainMessengerSession struct {
	Contract     *IBCCrossDomainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBCCrossDomainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBCCrossDomainMessengerCallerSession struct {
	Contract *IBCCrossDomainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IBCCrossDomainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBCCrossDomainMessengerTransactorSession struct {
	Contract     *IBCCrossDomainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IBCCrossDomainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBCCrossDomainMessengerRaw struct {
	Contract *IBCCrossDomainMessenger // Generic contract binding to access the raw methods on
}

// IBCCrossDomainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBCCrossDomainMessengerCallerRaw struct {
	Contract *IBCCrossDomainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IBCCrossDomainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBCCrossDomainMessengerTransactorRaw struct {
	Contract *IBCCrossDomainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBCCrossDomainMessenger creates a new instance of IBCCrossDomainMessenger, bound to a specific deployed contract.
func NewIBCCrossDomainMessenger(address common.Address, backend bind.ContractBackend) (*IBCCrossDomainMessenger, error) {
	contract, err := bindIBCCrossDomainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessenger{IBCCrossDomainMessengerCaller: IBCCrossDomainMessengerCaller{contract: contract}, IBCCrossDomainMessengerTransactor: IBCCrossDomainMessengerTransactor{contract: contract}, IBCCrossDomainMessengerFilterer: IBCCrossDomainMessengerFilterer{contract: contract}}, nil
}

// NewIBCCrossDomainMessengerCaller creates a new read-only instance of IBCCrossDomainMessenger, bound to a specific deployed contract.
func NewIBCCrossDomainMessengerCaller(address common.Address, caller bind.ContractCaller) (*IBCCrossDomainMessengerCaller, error) {
	contract, err := bindIBCCrossDomainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerCaller{contract: contract}, nil
}

// NewIBCCrossDomainMessengerTransactor creates a new write-only instance of IBCCrossDomainMessenger, bound to a specific deployed contract.
func NewIBCCrossDomainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IBCCrossDomainMessengerTransactor, error) {
	contract, err := bindIBCCrossDomainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerTransactor{contract: contract}, nil
}

// NewIBCCrossDomainMessengerFilterer creates a new log filterer instance of IBCCrossDomainMessenger, bound to a specific deployed contract.
func NewIBCCrossDomainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IBCCrossDomainMessengerFilterer, error) {
	contract, err := bindIBCCrossDomainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerFilterer{contract: contract}, nil
}

// bindIBCCrossDomainMessenger binds a generic wrapper to an already deployed contract.
func bindIBCCrossDomainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBCCrossDomainMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBCCrossDomainMessenger.Contract.IBCCrossDomainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.IBCCrossDomainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.IBCCrossDomainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBCCrossDomainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) MESSAGEVERSION() (uint16, error) {
	return _IBCCrossDomainMessenger.Contract.MESSAGEVERSION(&_IBCCrossDomainMessenger.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) MESSAGEVERSION() (uint16, error) {
	return _IBCCrossDomainMessenger.Contract.MESSAGEVERSION(&_IBCCrossDomainMessenger.CallOpts)
}

// MINGASCALLDATAOVERHEAD is a free data retrieval call binding the contract method 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) MINGASCALLDATAOVERHEAD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "MIN_GAS_CALLDATA_OVERHEAD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINGASCALLDATAOVERHEAD is a free data retrieval call binding the contract method 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) MINGASCALLDATAOVERHEAD() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.MINGASCALLDATAOVERHEAD(&_IBCCrossDomainMessenger.CallOpts)
}

// MINGASCALLDATAOVERHEAD is a free data retrieval call binding the contract method 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) MINGASCALLDATAOVERHEAD() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.MINGASCALLDATAOVERHEAD(&_IBCCrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADDENOMINATOR is a free data retrieval call binding the contract method 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) MINGASDYNAMICOVERHEADDENOMINATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINGASDYNAMICOVERHEADDENOMINATOR is a free data retrieval call binding the contract method 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) MINGASDYNAMICOVERHEADDENOMINATOR() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADDENOMINATOR(&_IBCCrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADDENOMINATOR is a free data retrieval call binding the contract method 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) MINGASDYNAMICOVERHEADDENOMINATOR() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADDENOMINATOR(&_IBCCrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADNUMERATOR is a free data retrieval call binding the contract method 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) MINGASDYNAMICOVERHEADNUMERATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINGASDYNAMICOVERHEADNUMERATOR is a free data retrieval call binding the contract method 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) MINGASDYNAMICOVERHEADNUMERATOR() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADNUMERATOR(&_IBCCrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADNUMERATOR is a free data retrieval call binding the contract method 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) MINGASDYNAMICOVERHEADNUMERATOR() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADNUMERATOR(&_IBCCrossDomainMessenger.CallOpts)
}

// OTHERMESSENGER is a free data retrieval call binding the contract method 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) OTHERMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "OTHER_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERMESSENGER is a free data retrieval call binding the contract method 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) OTHERMESSENGER() (common.Address, error) {
	return _IBCCrossDomainMessenger.Contract.OTHERMESSENGER(&_IBCCrossDomainMessenger.CallOpts)
}

// OTHERMESSENGER is a free data retrieval call binding the contract method 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) OTHERMESSENGER() (common.Address, error) {
	return _IBCCrossDomainMessenger.Contract.OTHERMESSENGER(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYCALLOVERHEAD is a free data retrieval call binding the contract method 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) RELAYCALLOVERHEAD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "RELAY_CALL_OVERHEAD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYCALLOVERHEAD is a free data retrieval call binding the contract method 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) RELAYCALLOVERHEAD() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYCALLOVERHEAD(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYCALLOVERHEAD is a free data retrieval call binding the contract method 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) RELAYCALLOVERHEAD() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYCALLOVERHEAD(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYCONSTANTOVERHEAD is a free data retrieval call binding the contract method 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) RELAYCONSTANTOVERHEAD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "RELAY_CONSTANT_OVERHEAD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYCONSTANTOVERHEAD is a free data retrieval call binding the contract method 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) RELAYCONSTANTOVERHEAD() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYCONSTANTOVERHEAD(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYCONSTANTOVERHEAD is a free data retrieval call binding the contract method 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) RELAYCONSTANTOVERHEAD() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYCONSTANTOVERHEAD(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYGASCHECKBUFFER is a free data retrieval call binding the contract method 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) RELAYGASCHECKBUFFER(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "RELAY_GAS_CHECK_BUFFER")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYGASCHECKBUFFER is a free data retrieval call binding the contract method 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) RELAYGASCHECKBUFFER() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYGASCHECKBUFFER(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYGASCHECKBUFFER is a free data retrieval call binding the contract method 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) RELAYGASCHECKBUFFER() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYGASCHECKBUFFER(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYRESERVEDGAS is a free data retrieval call binding the contract method 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) RELAYRESERVEDGAS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "RELAY_RESERVED_GAS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYRESERVEDGAS is a free data retrieval call binding the contract method 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) RELAYRESERVEDGAS() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYRESERVEDGAS(&_IBCCrossDomainMessenger.CallOpts)
}

// RELAYRESERVEDGAS is a free data retrieval call binding the contract method 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) RELAYRESERVEDGAS() (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.RELAYRESERVEDGAS(&_IBCCrossDomainMessenger.CallOpts)
}

// BaseGas is a free data retrieval call binding the contract method 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) BaseGas(opts *bind.CallOpts, _message []byte, _minGasLimit uint32) (uint64, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "baseGas", _message, _minGasLimit)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseGas is a free data retrieval call binding the contract method 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) BaseGas(_message []byte, _minGasLimit uint32) (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.BaseGas(&_IBCCrossDomainMessenger.CallOpts, _message, _minGasLimit)
}

// BaseGas is a free data retrieval call binding the contract method 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) BaseGas(_message []byte, _minGasLimit uint32) (uint64, error) {
	return _IBCCrossDomainMessenger.Contract.BaseGas(&_IBCCrossDomainMessenger.CallOpts, _message, _minGasLimit)
}

// FailedMessages is a free data retrieval call binding the contract method 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) FailedMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "failedMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) FailedMessages(arg0 [32]byte) (bool, error) {
	return _IBCCrossDomainMessenger.Contract.FailedMessages(&_IBCCrossDomainMessenger.CallOpts, arg0)
}

// FailedMessages is a free data retrieval call binding the contract method 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) FailedMessages(arg0 [32]byte) (bool, error) {
	return _IBCCrossDomainMessenger.Contract.FailedMessages(&_IBCCrossDomainMessenger.CallOpts, arg0)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) L1CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "l1CrossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _IBCCrossDomainMessenger.Contract.L1CrossDomainMessenger(&_IBCCrossDomainMessenger.CallOpts)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _IBCCrossDomainMessenger.Contract.L1CrossDomainMessenger(&_IBCCrossDomainMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) MessageNonce() (*big.Int, error) {
	return _IBCCrossDomainMessenger.Contract.MessageNonce(&_IBCCrossDomainMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _IBCCrossDomainMessenger.Contract.MessageNonce(&_IBCCrossDomainMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) Paused() (bool, error) {
	return _IBCCrossDomainMessenger.Contract.Paused(&_IBCCrossDomainMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) Paused() (bool, error) {
	return _IBCCrossDomainMessenger.Contract.Paused(&_IBCCrossDomainMessenger.CallOpts)
}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) SuccessfulMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "successfulMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) SuccessfulMessages(arg0 [32]byte) (bool, error) {
	return _IBCCrossDomainMessenger.Contract.SuccessfulMessages(&_IBCCrossDomainMessenger.CallOpts, arg0)
}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) SuccessfulMessages(arg0 [32]byte) (bool, error) {
	return _IBCCrossDomainMessenger.Contract.SuccessfulMessages(&_IBCCrossDomainMessenger.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) Version() (string, error) {
	return _IBCCrossDomainMessenger.Contract.Version(&_IBCCrossDomainMessenger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) Version() (string, error) {
	return _IBCCrossDomainMessenger.Contract.Version(&_IBCCrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBCCrossDomainMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _IBCCrossDomainMessenger.Contract.XDomainMessageSender(&_IBCCrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _IBCCrossDomainMessenger.Contract.XDomainMessageSender(&_IBCCrossDomainMessenger.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) Initialize() (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.Initialize(&_IBCCrossDomainMessenger.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactorSession) Initialize() (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.Initialize(&_IBCCrossDomainMessenger.TransactOpts)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd764ad0b.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactor) RelayMessage(opts *bind.TransactOpts, _nonce *big.Int, _sender common.Address, _target common.Address, _value *big.Int, _minGasLimit *big.Int, _message []byte) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.contract.Transact(opts, "relayMessage", _nonce, _sender, _target, _value, _minGasLimit, _message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd764ad0b.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) RelayMessage(_nonce *big.Int, _sender common.Address, _target common.Address, _value *big.Int, _minGasLimit *big.Int, _message []byte) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.RelayMessage(&_IBCCrossDomainMessenger.TransactOpts, _nonce, _sender, _target, _value, _minGasLimit, _message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd764ad0b.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactorSession) RelayMessage(_nonce *big.Int, _sender common.Address, _target common.Address, _value *big.Int, _minGasLimit *big.Int, _message []byte) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.RelayMessage(&_IBCCrossDomainMessenger.TransactOpts, _nonce, _sender, _target, _value, _minGasLimit, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactor) SendMessage(opts *bind.TransactOpts, _target common.Address, _message []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.contract.Transact(opts, "sendMessage", _target, _message, _minGasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerSession) SendMessage(_target common.Address, _message []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.SendMessage(&_IBCCrossDomainMessenger.TransactOpts, _target, _message, _minGasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerTransactorSession) SendMessage(_target common.Address, _message []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _IBCCrossDomainMessenger.Contract.SendMessage(&_IBCCrossDomainMessenger.TransactOpts, _target, _message, _minGasLimit)
}

// IBCCrossDomainMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerFailedRelayedMessageIterator struct {
	Event *IBCCrossDomainMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCCrossDomainMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCCrossDomainMessengerFailedRelayedMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCCrossDomainMessengerFailedRelayedMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCCrossDomainMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCCrossDomainMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCCrossDomainMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerFailedRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*IBCCrossDomainMessengerFailedRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerFailedRelayedMessageIterator{contract: _IBCCrossDomainMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *IBCCrossDomainMessengerFailedRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCCrossDomainMessengerFailedRelayedMessage)
				if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*IBCCrossDomainMessengerFailedRelayedMessage, error) {
	event := new(IBCCrossDomainMessengerFailedRelayedMessage)
	if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCCrossDomainMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerInitializedIterator struct {
	Event *IBCCrossDomainMessengerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCCrossDomainMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCCrossDomainMessengerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCCrossDomainMessengerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCCrossDomainMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCCrossDomainMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCCrossDomainMessengerInitialized represents a Initialized event raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*IBCCrossDomainMessengerInitializedIterator, error) {

	logs, sub, err := _IBCCrossDomainMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerInitializedIterator{contract: _IBCCrossDomainMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IBCCrossDomainMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _IBCCrossDomainMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCCrossDomainMessengerInitialized)
				if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) ParseInitialized(log types.Log) (*IBCCrossDomainMessengerInitialized, error) {
	event := new(IBCCrossDomainMessengerInitialized)
	if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCCrossDomainMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerRelayedMessageIterator struct {
	Event *IBCCrossDomainMessengerRelayedMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCCrossDomainMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCCrossDomainMessengerRelayedMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCCrossDomainMessengerRelayedMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCCrossDomainMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCCrossDomainMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCCrossDomainMessengerRelayedMessage represents a RelayedMessage event raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*IBCCrossDomainMessengerRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.FilterLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerRelayedMessageIterator{contract: _IBCCrossDomainMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *IBCCrossDomainMessengerRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.WatchLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCCrossDomainMessengerRelayedMessage)
				if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) ParseRelayedMessage(log types.Log) (*IBCCrossDomainMessengerRelayedMessage, error) {
	event := new(IBCCrossDomainMessengerRelayedMessage)
	if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCCrossDomainMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerSentMessageIterator struct {
	Event *IBCCrossDomainMessengerSentMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCCrossDomainMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCCrossDomainMessengerSentMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCCrossDomainMessengerSentMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCCrossDomainMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCCrossDomainMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCCrossDomainMessengerSentMessage represents a SentMessage event raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerSentMessage struct {
	To       common.Address
	GasLimit *big.Int
	Value    *big.Int
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0xecf079dc36711b69eae2d3ea1b148e593e45d8f663834aac87bab9ae231346f5.
//
// Solidity: event SentMessage(address indexed to, uint256 gasLimit, uint256 value, bytes data)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, to []common.Address) (*IBCCrossDomainMessengerSentMessageIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.FilterLogs(opts, "SentMessage", toRule)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerSentMessageIterator{contract: _IBCCrossDomainMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0xecf079dc36711b69eae2d3ea1b148e593e45d8f663834aac87bab9ae231346f5.
//
// Solidity: event SentMessage(address indexed to, uint256 gasLimit, uint256 value, bytes data)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *IBCCrossDomainMessengerSentMessage, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.WatchLogs(opts, "SentMessage", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCCrossDomainMessengerSentMessage)
				if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSentMessage is a log parse operation binding the contract event 0xecf079dc36711b69eae2d3ea1b148e593e45d8f663834aac87bab9ae231346f5.
//
// Solidity: event SentMessage(address indexed to, uint256 gasLimit, uint256 value, bytes data)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) ParseSentMessage(log types.Log) (*IBCCrossDomainMessengerSentMessage, error) {
	event := new(IBCCrossDomainMessengerSentMessage)
	if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCCrossDomainMessengerSentMessage0Iterator is returned from FilterSentMessage0 and is used to iterate over the raw logs and unpacked data for SentMessage0 events raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerSentMessage0Iterator struct {
	Event *IBCCrossDomainMessengerSentMessage0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCCrossDomainMessengerSentMessage0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCCrossDomainMessengerSentMessage0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCCrossDomainMessengerSentMessage0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCCrossDomainMessengerSentMessage0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCCrossDomainMessengerSentMessage0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCCrossDomainMessengerSentMessage0 represents a SentMessage0 event raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerSentMessage0 struct {
	Target       common.Address
	Sender       common.Address
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage0 is a free log retrieval operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) FilterSentMessage0(opts *bind.FilterOpts, target []common.Address) (*IBCCrossDomainMessengerSentMessage0Iterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.FilterLogs(opts, "SentMessage0", targetRule)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerSentMessage0Iterator{contract: _IBCCrossDomainMessenger.contract, event: "SentMessage0", logs: logs, sub: sub}, nil
}

// WatchSentMessage0 is a free log subscription operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) WatchSentMessage0(opts *bind.WatchOpts, sink chan<- *IBCCrossDomainMessengerSentMessage0, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.WatchLogs(opts, "SentMessage0", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCCrossDomainMessengerSentMessage0)
				if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "SentMessage0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSentMessage0 is a log parse operation binding the contract event 0xcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) ParseSentMessage0(log types.Log) (*IBCCrossDomainMessengerSentMessage0, error) {
	event := new(IBCCrossDomainMessengerSentMessage0)
	if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "SentMessage0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBCCrossDomainMessengerSentMessageExtension1Iterator is returned from FilterSentMessageExtension1 and is used to iterate over the raw logs and unpacked data for SentMessageExtension1 events raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerSentMessageExtension1Iterator struct {
	Event *IBCCrossDomainMessengerSentMessageExtension1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IBCCrossDomainMessengerSentMessageExtension1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBCCrossDomainMessengerSentMessageExtension1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IBCCrossDomainMessengerSentMessageExtension1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IBCCrossDomainMessengerSentMessageExtension1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBCCrossDomainMessengerSentMessageExtension1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBCCrossDomainMessengerSentMessageExtension1 represents a SentMessageExtension1 event raised by the IBCCrossDomainMessenger contract.
type IBCCrossDomainMessengerSentMessageExtension1 struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSentMessageExtension1 is a free log retrieval operation binding the contract event 0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546.
//
// Solidity: event SentMessageExtension1(address indexed sender, uint256 value)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) FilterSentMessageExtension1(opts *bind.FilterOpts, sender []common.Address) (*IBCCrossDomainMessengerSentMessageExtension1Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.FilterLogs(opts, "SentMessageExtension1", senderRule)
	if err != nil {
		return nil, err
	}
	return &IBCCrossDomainMessengerSentMessageExtension1Iterator{contract: _IBCCrossDomainMessenger.contract, event: "SentMessageExtension1", logs: logs, sub: sub}, nil
}

// WatchSentMessageExtension1 is a free log subscription operation binding the contract event 0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546.
//
// Solidity: event SentMessageExtension1(address indexed sender, uint256 value)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) WatchSentMessageExtension1(opts *bind.WatchOpts, sink chan<- *IBCCrossDomainMessengerSentMessageExtension1, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IBCCrossDomainMessenger.contract.WatchLogs(opts, "SentMessageExtension1", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBCCrossDomainMessengerSentMessageExtension1)
				if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "SentMessageExtension1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSentMessageExtension1 is a log parse operation binding the contract event 0x8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d546.
//
// Solidity: event SentMessageExtension1(address indexed sender, uint256 value)
func (_IBCCrossDomainMessenger *IBCCrossDomainMessengerFilterer) ParseSentMessageExtension1(log types.Log) (*IBCCrossDomainMessengerSentMessageExtension1, error) {
	event := new(IBCCrossDomainMessengerSentMessageExtension1)
	if err := _IBCCrossDomainMessenger.contract.UnpackLog(event, "SentMessageExtension1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
