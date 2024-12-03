// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakeV3factory

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
	_ = abi.ConvertType
)

// PancakeV3factoryMetaData contains all meta data concerning the PancakeV3factory contract.
var PancakeV3factoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"FeeAmountExtraInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"SetLmPoolDeployer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"WhiteListAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacingExtraInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lmPoolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setFeeAmountExtraInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol0\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol1\",\"type\":\"uint32\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lmPool\",\"type\":\"address\"}],\"name\":\"setLmPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"setLmPoolDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"setWhiteListAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100f55760003560e01c80637e8435e6116100975780638da5cb5b116100665780638da5cb5b146103a55780638ff38e80146103ad578063a1671295146103df578063e4a86a9914610428576100f5565b80637e8435e6146102c357806380d6a7921461030a57806388e8006d1461033d5780638a7c195f1461037a576100f5565b806322afcccb116100d357806322afcccb146101dc5780633119049a1461021557806343db87da1461021d5780635e492ac8146102bb576100f5565b806311ff5e8d146100fa57806313af4035146101375780631698ee821461016a575b600080fd5b6101356004803603604081101561011057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610463565b005b6101356004803603602081101561014d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610590565b6101b36004803603606081101561018057600080fd5b50803573ffffffffffffffffffffffffffffffffffffffff908116916020810135909116906040013562ffffff166106a3565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101fe600480360360208110156101f257600080fd5b503562ffffff166106dc565b6040805160029290920b8252519081900360200190f35b6101b36106f1565b61027a6004803603608081101561023357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135909116906fffffffffffffffffffffffffffffffff60408201358116916060013516610715565b60405180836fffffffffffffffffffffffffffffffff168152602001826fffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b6101b361086b565b610135600480360360608110156102d957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169063ffffffff60208201358116916040013516610887565b6101356004803603602081101561032057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166109a5565b61035f6004803603602081101561035357600080fd5b503562ffffff16610a9a565b60408051921515835290151560208301528051918290030190f35b6101356004803603604081101561039057600080fd5b5062ffffff813516906020013560020b610ab8565b6101b3610cc3565b610135600480360360608110156103c357600080fd5b5062ffffff813516906020810135151590604001351515610cdf565b6101b3600480360360608110156103f557600080fd5b50803573ffffffffffffffffffffffffffffffffffffffff908116916020810135909116906040013562ffffff16610e50565b6101356004803603604081101561043e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001351515611234565b60005473ffffffffffffffffffffffffffffffffffffffff163314806104a0575060055473ffffffffffffffffffffffffffffffffffffffff1633145b61050b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4e6f74206f776e6572206f72204c4d20706f6f6c206465706c6f796572000000604482015290519081900360640190fd5b8173ffffffffffffffffffffffffffffffffffffffff1663cc7e7fa2826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561057457600080fd5b505af1158015610588573d6000803e3d6000fd5b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461061657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff808516939216917fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c91a3600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600260209081526000938452604080852082529284528284209052825290205473ffffffffffffffffffffffffffffffffffffffff1681565b60016020526000908152604090205460020b81565b7f00000000000000000000000041ff9aa7e16b8b1a8a8dc4f0efacd93d02d071c981565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461079e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b604080517f85b6672900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526fffffffffffffffffffffffffffffffff8088166024840152861660448301528251908916926385b6672992606480820193918290030181600087803b15801561082b57600080fd5b505af115801561083f573d6000803e3d6000fd5b505050506040513d604081101561085557600080fd5b5080516020909101519097909650945050505050565b60055473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff16331461090d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b604080517fb0d0d21100000000000000000000000000000000000000000000000000000000815263ffffffff808516600483015283166024820152905173ffffffffffffffffffffffffffffffffffffffff85169163b0d0d21191604480830192600092919082900301818387803b15801561098857600080fd5b505af115801561099c573d6000803e3d6000fd5b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a2b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca40790600090a250565b60036020526000908152604090205460ff8082169161010090041682565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b3e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b620f42408262ffffff1610610b5257600080fd5b60008160020b138015610b6957506140008160020b125b610b7257600080fd5b62ffffff8216600090815260016020526040902054600290810b900b15610b9857600080fd5b62ffffff828116600081815260016020818152604080842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000016600289900b9788161790558051808201825284815280830193845285855260039092528084209151825493517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909416901515177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010093151593909302929092179055517fc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc9190a3604080516000815260016020820152815162ffffff8516927fed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88928290030190a25050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff163314610d6557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b62ffffff8316600090815260016020526040902054600290810b900b610d8a57600080fd5b604080518082018252831515808252831515602080840182815262ffffff89166000818152600384528790209551865492511515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9115157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009094169390931716919091179094558451928352820152825191927fed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d8892918290030190a2505050565b60008273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610e8b57600080fd5b6000808473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1610610ec8578486610ecb565b85855b909250905073ffffffffffffffffffffffffffffffffffffffff8216610ef057600080fd5b62ffffff8416600090815260016020908152604080832054600383529281902081518083019092525460ff8082161515835261010090910416151591810191909152600291820b9182900b15801590610f4a575080602001515b610fb557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f666565206973206e6f7420617661696c61626c65207965740000000000000000604482015290519081900360640190fd5b805115611024573360009081526004602052604090205460ff16611024576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806113e16032913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff84811660009081526002602090815260408083208785168452825280832062ffffff8b168452909152902054161561107057600080fd5b604080517ffad5359f00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8681166024830152858116604483015262ffffff89166064830152600285900b608483015291517f00000000000000000000000041ff9aa7e16b8b1a8a8dc4f0efacd93d02d071c99092169163fad5359f9160a4808201926020929091908290030181600087803b15801561112657600080fd5b505af115801561113a573d6000803e3d6000fd5b505050506040513d602081101561115057600080fd5b505173ffffffffffffffffffffffffffffffffffffffff80861660008181526002602081815260408084208a871680865290835281852062ffffff8f168087529084528286208054988a167fffffffffffffffffffffffff0000000000000000000000000000000000000000998a1681179091558287528585528387208888528552838720828852855295839020805490981686179097558151938a900b8452918301939093528251959a5093947f783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118929181900390910190a4505050509392505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146112ba57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f74206f776e65720000000000000000000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604090205460ff161515811515141561135557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f7374617465206e6f74206368616e676500000000000000000000000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526004602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016851515908117909155825190815291517faec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a29281900390910190a2505056fe757365722073686f756c6420626520696e20746865207768697465206c69737420666f722074686973206665652074696572a164736f6c6343000706000a",
}

// PancakeV3factoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeV3factoryMetaData.ABI instead.
var PancakeV3factoryABI = PancakeV3factoryMetaData.ABI

// PancakeV3factoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PancakeV3factoryMetaData.Bin instead.
var PancakeV3factoryBin = PancakeV3factoryMetaData.Bin

// DeployPancakeV3factory deploys a new Ethereum contract, binding an instance of PancakeV3factory to it.
func DeployPancakeV3factory(auth *bind.TransactOpts, backend bind.ContractBackend, _poolDeployer common.Address) (common.Address, *types.Transaction, *PancakeV3factory, error) {
	parsed, err := PancakeV3factoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PancakeV3factoryBin), backend, _poolDeployer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PancakeV3factory{PancakeV3factoryCaller: PancakeV3factoryCaller{contract: contract}, PancakeV3factoryTransactor: PancakeV3factoryTransactor{contract: contract}, PancakeV3factoryFilterer: PancakeV3factoryFilterer{contract: contract}}, nil
}

// PancakeV3factory is an auto generated Go binding around an Ethereum contract.
type PancakeV3factory struct {
	PancakeV3factoryCaller     // Read-only binding to the contract
	PancakeV3factoryTransactor // Write-only binding to the contract
	PancakeV3factoryFilterer   // Log filterer for contract events
}

// PancakeV3factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeV3factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeV3factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeV3factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeV3factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeV3factorySession struct {
	Contract     *PancakeV3factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeV3factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeV3factoryCallerSession struct {
	Contract *PancakeV3factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PancakeV3factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeV3factoryTransactorSession struct {
	Contract     *PancakeV3factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PancakeV3factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeV3factoryRaw struct {
	Contract *PancakeV3factory // Generic contract binding to access the raw methods on
}

// PancakeV3factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeV3factoryCallerRaw struct {
	Contract *PancakeV3factoryCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeV3factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeV3factoryTransactorRaw struct {
	Contract *PancakeV3factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeV3factory creates a new instance of PancakeV3factory, bound to a specific deployed contract.
func NewPancakeV3factory(address common.Address, backend bind.ContractBackend) (*PancakeV3factory, error) {
	contract, err := bindPancakeV3factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factory{PancakeV3factoryCaller: PancakeV3factoryCaller{contract: contract}, PancakeV3factoryTransactor: PancakeV3factoryTransactor{contract: contract}, PancakeV3factoryFilterer: PancakeV3factoryFilterer{contract: contract}}, nil
}

// NewPancakeV3factoryCaller creates a new read-only instance of PancakeV3factory, bound to a specific deployed contract.
func NewPancakeV3factoryCaller(address common.Address, caller bind.ContractCaller) (*PancakeV3factoryCaller, error) {
	contract, err := bindPancakeV3factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryCaller{contract: contract}, nil
}

// NewPancakeV3factoryTransactor creates a new write-only instance of PancakeV3factory, bound to a specific deployed contract.
func NewPancakeV3factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeV3factoryTransactor, error) {
	contract, err := bindPancakeV3factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryTransactor{contract: contract}, nil
}

// NewPancakeV3factoryFilterer creates a new log filterer instance of PancakeV3factory, bound to a specific deployed contract.
func NewPancakeV3factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeV3factoryFilterer, error) {
	contract, err := bindPancakeV3factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryFilterer{contract: contract}, nil
}

// bindPancakeV3factory binds a generic wrapper to an already deployed contract.
func bindPancakeV3factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeV3factoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV3factory *PancakeV3factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV3factory.Contract.PancakeV3factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV3factory *PancakeV3factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.PancakeV3factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV3factory *PancakeV3factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.PancakeV3factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeV3factory *PancakeV3factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeV3factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeV3factory *PancakeV3factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeV3factory *PancakeV3factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.contract.Transact(opts, method, params...)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3factory *PancakeV3factoryCaller) FeeAmountTickSpacing(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeV3factory.contract.Call(opts, &out, "feeAmountTickSpacing", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3factory *PancakeV3factorySession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _PancakeV3factory.Contract.FeeAmountTickSpacing(&_PancakeV3factory.CallOpts, arg0)
}

// FeeAmountTickSpacing is a free data retrieval call binding the contract method 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (_PancakeV3factory *PancakeV3factoryCallerSession) FeeAmountTickSpacing(arg0 *big.Int) (*big.Int, error) {
	return _PancakeV3factory.Contract.FeeAmountTickSpacing(&_PancakeV3factory.CallOpts, arg0)
}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3factory *PancakeV3factoryCaller) FeeAmountTickSpacingExtraInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	var out []interface{}
	err := _PancakeV3factory.contract.Call(opts, &out, "feeAmountTickSpacingExtraInfo", arg0)

	outstruct := new(struct {
		WhitelistRequested bool
		Enabled            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WhitelistRequested = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3factory *PancakeV3factorySession) FeeAmountTickSpacingExtraInfo(arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	return _PancakeV3factory.Contract.FeeAmountTickSpacingExtraInfo(&_PancakeV3factory.CallOpts, arg0)
}

// FeeAmountTickSpacingExtraInfo is a free data retrieval call binding the contract method 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (_PancakeV3factory *PancakeV3factoryCallerSession) FeeAmountTickSpacingExtraInfo(arg0 *big.Int) (struct {
	WhitelistRequested bool
	Enabled            bool
}, error) {
	return _PancakeV3factory.Contract.FeeAmountTickSpacingExtraInfo(&_PancakeV3factory.CallOpts, arg0)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3factory *PancakeV3factoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3factory.contract.Call(opts, &out, "getPool", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3factory *PancakeV3factorySession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _PancakeV3factory.Contract.GetPool(&_PancakeV3factory.CallOpts, arg0, arg1, arg2)
}

// GetPool is a free data retrieval call binding the contract method 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (_PancakeV3factory *PancakeV3factoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (common.Address, error) {
	return _PancakeV3factory.Contract.GetPool(&_PancakeV3factory.CallOpts, arg0, arg1, arg2)
}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3factory *PancakeV3factoryCaller) LmPoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3factory.contract.Call(opts, &out, "lmPoolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3factory *PancakeV3factorySession) LmPoolDeployer() (common.Address, error) {
	return _PancakeV3factory.Contract.LmPoolDeployer(&_PancakeV3factory.CallOpts)
}

// LmPoolDeployer is a free data retrieval call binding the contract method 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (_PancakeV3factory *PancakeV3factoryCallerSession) LmPoolDeployer() (common.Address, error) {
	return _PancakeV3factory.Contract.LmPoolDeployer(&_PancakeV3factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3factory *PancakeV3factoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3factory *PancakeV3factorySession) Owner() (common.Address, error) {
	return _PancakeV3factory.Contract.Owner(&_PancakeV3factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeV3factory *PancakeV3factoryCallerSession) Owner() (common.Address, error) {
	return _PancakeV3factory.Contract.Owner(&_PancakeV3factory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3factory *PancakeV3factoryCaller) PoolDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeV3factory.contract.Call(opts, &out, "poolDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3factory *PancakeV3factorySession) PoolDeployer() (common.Address, error) {
	return _PancakeV3factory.Contract.PoolDeployer(&_PancakeV3factory.CallOpts)
}

// PoolDeployer is a free data retrieval call binding the contract method 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (_PancakeV3factory *PancakeV3factoryCallerSession) PoolDeployer() (common.Address, error) {
	return _PancakeV3factory.Contract.PoolDeployer(&_PancakeV3factory.CallOpts)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3factory *PancakeV3factoryTransactor) CollectProtocol(opts *bind.TransactOpts, pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "collectProtocol", pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3factory *PancakeV3factorySession) CollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.CollectProtocol(&_PancakeV3factory.TransactOpts, pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_PancakeV3factory *PancakeV3factoryTransactorSession) CollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.CollectProtocol(&_PancakeV3factory.TransactOpts, pool, recipient, amount0Requested, amount1Requested)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3factory *PancakeV3factoryTransactor) CreatePool(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "createPool", tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3factory *PancakeV3factorySession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.CreatePool(&_PancakeV3factory.TransactOpts, tokenA, tokenB, fee)
}

// CreatePool is a paid mutator transaction binding the contract method 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (_PancakeV3factory *PancakeV3factoryTransactorSession) CreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.CreatePool(&_PancakeV3factory.TransactOpts, tokenA, tokenB, fee)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) EnableFeeAmount(opts *bind.TransactOpts, fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "enableFeeAmount", fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3factory *PancakeV3factorySession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.EnableFeeAmount(&_PancakeV3factory.TransactOpts, fee, tickSpacing)
}

// EnableFeeAmount is a paid mutator transaction binding the contract method 0x8a7c195f.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) EnableFeeAmount(fee *big.Int, tickSpacing *big.Int) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.EnableFeeAmount(&_PancakeV3factory.TransactOpts, fee, tickSpacing)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) SetFeeAmountExtraInfo(opts *bind.TransactOpts, fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "setFeeAmountExtraInfo", fee, whitelistRequested, enabled)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3factory *PancakeV3factorySession) SetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetFeeAmountExtraInfo(&_PancakeV3factory.TransactOpts, fee, whitelistRequested, enabled)
}

// SetFeeAmountExtraInfo is a paid mutator transaction binding the contract method 0x8ff38e80.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) SetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetFeeAmountExtraInfo(&_PancakeV3factory.TransactOpts, fee, whitelistRequested, enabled)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) SetFeeProtocol(opts *bind.TransactOpts, pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "setFeeProtocol", pool, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3factory *PancakeV3factorySession) SetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetFeeProtocol(&_PancakeV3factory.TransactOpts, pool, feeProtocol0, feeProtocol1)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7e8435e6.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) SetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetFeeProtocol(&_PancakeV3factory.TransactOpts, pool, feeProtocol0, feeProtocol1)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) SetLmPool(opts *bind.TransactOpts, pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "setLmPool", pool, lmPool)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3factory *PancakeV3factorySession) SetLmPool(pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetLmPool(&_PancakeV3factory.TransactOpts, pool, lmPool)
}

// SetLmPool is a paid mutator transaction binding the contract method 0x11ff5e8d.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) SetLmPool(pool common.Address, lmPool common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetLmPool(&_PancakeV3factory.TransactOpts, pool, lmPool)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) SetLmPoolDeployer(opts *bind.TransactOpts, _lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "setLmPoolDeployer", _lmPoolDeployer)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3factory *PancakeV3factorySession) SetLmPoolDeployer(_lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetLmPoolDeployer(&_PancakeV3factory.TransactOpts, _lmPoolDeployer)
}

// SetLmPoolDeployer is a paid mutator transaction binding the contract method 0x80d6a792.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) SetLmPoolDeployer(_lmPoolDeployer common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetLmPoolDeployer(&_PancakeV3factory.TransactOpts, _lmPoolDeployer)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3factory *PancakeV3factorySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetOwner(&_PancakeV3factory.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetOwner(&_PancakeV3factory.TransactOpts, _owner)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3factory *PancakeV3factoryTransactor) SetWhiteListAddress(opts *bind.TransactOpts, user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3factory.contract.Transact(opts, "setWhiteListAddress", user, verified)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3factory *PancakeV3factorySession) SetWhiteListAddress(user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetWhiteListAddress(&_PancakeV3factory.TransactOpts, user, verified)
}

// SetWhiteListAddress is a paid mutator transaction binding the contract method 0xe4a86a99.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (_PancakeV3factory *PancakeV3factoryTransactorSession) SetWhiteListAddress(user common.Address, verified bool) (*types.Transaction, error) {
	return _PancakeV3factory.Contract.SetWhiteListAddress(&_PancakeV3factory.TransactOpts, user, verified)
}

// PancakeV3factoryFeeAmountEnabledIterator is returned from FilterFeeAmountEnabled and is used to iterate over the raw logs and unpacked data for FeeAmountEnabled events raised by the PancakeV3factory contract.
type PancakeV3factoryFeeAmountEnabledIterator struct {
	Event *PancakeV3factoryFeeAmountEnabled // Event containing the contract specifics and raw log

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
func (it *PancakeV3factoryFeeAmountEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3factoryFeeAmountEnabled)
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
		it.Event = new(PancakeV3factoryFeeAmountEnabled)
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
func (it *PancakeV3factoryFeeAmountEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3factoryFeeAmountEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3factoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the PancakeV3factory contract.
type PancakeV3factoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountEnabled is a free log retrieval operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3factory *PancakeV3factoryFilterer) FilterFeeAmountEnabled(opts *bind.FilterOpts, fee []*big.Int, tickSpacing []*big.Int) (*PancakeV3factoryFeeAmountEnabledIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _PancakeV3factory.contract.FilterLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryFeeAmountEnabledIterator{contract: _PancakeV3factory.contract, event: "FeeAmountEnabled", logs: logs, sub: sub}, nil
}

// WatchFeeAmountEnabled is a free log subscription operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3factory *PancakeV3factoryFilterer) WatchFeeAmountEnabled(opts *bind.WatchOpts, sink chan<- *PancakeV3factoryFeeAmountEnabled, fee []*big.Int, tickSpacing []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var tickSpacingRule []interface{}
	for _, tickSpacingItem := range tickSpacing {
		tickSpacingRule = append(tickSpacingRule, tickSpacingItem)
	}

	logs, sub, err := _PancakeV3factory.contract.WatchLogs(opts, "FeeAmountEnabled", feeRule, tickSpacingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3factoryFeeAmountEnabled)
				if err := _PancakeV3factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
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

// ParseFeeAmountEnabled is a log parse operation binding the contract event 0xc66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (_PancakeV3factory *PancakeV3factoryFilterer) ParseFeeAmountEnabled(log types.Log) (*PancakeV3factoryFeeAmountEnabled, error) {
	event := new(PancakeV3factoryFeeAmountEnabled)
	if err := _PancakeV3factory.contract.UnpackLog(event, "FeeAmountEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3factoryFeeAmountExtraInfoUpdatedIterator is returned from FilterFeeAmountExtraInfoUpdated and is used to iterate over the raw logs and unpacked data for FeeAmountExtraInfoUpdated events raised by the PancakeV3factory contract.
type PancakeV3factoryFeeAmountExtraInfoUpdatedIterator struct {
	Event *PancakeV3factoryFeeAmountExtraInfoUpdated // Event containing the contract specifics and raw log

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
func (it *PancakeV3factoryFeeAmountExtraInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3factoryFeeAmountExtraInfoUpdated)
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
		it.Event = new(PancakeV3factoryFeeAmountExtraInfoUpdated)
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
func (it *PancakeV3factoryFeeAmountExtraInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3factoryFeeAmountExtraInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3factoryFeeAmountExtraInfoUpdated represents a FeeAmountExtraInfoUpdated event raised by the PancakeV3factory contract.
type PancakeV3factoryFeeAmountExtraInfoUpdated struct {
	Fee                *big.Int
	WhitelistRequested bool
	Enabled            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFeeAmountExtraInfoUpdated is a free log retrieval operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3factory *PancakeV3factoryFilterer) FilterFeeAmountExtraInfoUpdated(opts *bind.FilterOpts, fee []*big.Int) (*PancakeV3factoryFeeAmountExtraInfoUpdatedIterator, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3factory.contract.FilterLogs(opts, "FeeAmountExtraInfoUpdated", feeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryFeeAmountExtraInfoUpdatedIterator{contract: _PancakeV3factory.contract, event: "FeeAmountExtraInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeAmountExtraInfoUpdated is a free log subscription operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3factory *PancakeV3factoryFilterer) WatchFeeAmountExtraInfoUpdated(opts *bind.WatchOpts, sink chan<- *PancakeV3factoryFeeAmountExtraInfoUpdated, fee []*big.Int) (event.Subscription, error) {

	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3factory.contract.WatchLogs(opts, "FeeAmountExtraInfoUpdated", feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3factoryFeeAmountExtraInfoUpdated)
				if err := _PancakeV3factory.contract.UnpackLog(event, "FeeAmountExtraInfoUpdated", log); err != nil {
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

// ParseFeeAmountExtraInfoUpdated is a log parse operation binding the contract event 0xed85b616dbfbc54d0f1180a7bd0f6e3bb645b269b234e7a9edcc269ef1443d88.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (_PancakeV3factory *PancakeV3factoryFilterer) ParseFeeAmountExtraInfoUpdated(log types.Log) (*PancakeV3factoryFeeAmountExtraInfoUpdated, error) {
	event := new(PancakeV3factoryFeeAmountExtraInfoUpdated)
	if err := _PancakeV3factory.contract.UnpackLog(event, "FeeAmountExtraInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3factoryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the PancakeV3factory contract.
type PancakeV3factoryOwnerChangedIterator struct {
	Event *PancakeV3factoryOwnerChanged // Event containing the contract specifics and raw log

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
func (it *PancakeV3factoryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3factoryOwnerChanged)
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
		it.Event = new(PancakeV3factoryOwnerChanged)
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
func (it *PancakeV3factoryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3factoryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3factoryOwnerChanged represents a OwnerChanged event raised by the PancakeV3factory contract.
type PancakeV3factoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3factory *PancakeV3factoryFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PancakeV3factoryOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeV3factory.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryOwnerChangedIterator{contract: _PancakeV3factory.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3factory *PancakeV3factoryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *PancakeV3factoryOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeV3factory.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3factoryOwnerChanged)
				if err := _PancakeV3factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_PancakeV3factory *PancakeV3factoryFilterer) ParseOwnerChanged(log types.Log) (*PancakeV3factoryOwnerChanged, error) {
	event := new(PancakeV3factoryOwnerChanged)
	if err := _PancakeV3factory.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3factoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the PancakeV3factory contract.
type PancakeV3factoryPoolCreatedIterator struct {
	Event *PancakeV3factoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *PancakeV3factoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3factoryPoolCreated)
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
		it.Event = new(PancakeV3factoryPoolCreated)
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
func (it *PancakeV3factoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3factoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3factoryPoolCreated represents a PoolCreated event raised by the PancakeV3factory contract.
type PancakeV3factoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3factory *PancakeV3factoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address, fee []*big.Int) (*PancakeV3factoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3factory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryPoolCreatedIterator{contract: _PancakeV3factory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3factory *PancakeV3factoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *PancakeV3factoryPoolCreated, token0 []common.Address, token1 []common.Address, fee []*big.Int) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _PancakeV3factory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3factoryPoolCreated)
				if err := _PancakeV3factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (_PancakeV3factory *PancakeV3factoryFilterer) ParsePoolCreated(log types.Log) (*PancakeV3factoryPoolCreated, error) {
	event := new(PancakeV3factoryPoolCreated)
	if err := _PancakeV3factory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3factorySetLmPoolDeployerIterator is returned from FilterSetLmPoolDeployer and is used to iterate over the raw logs and unpacked data for SetLmPoolDeployer events raised by the PancakeV3factory contract.
type PancakeV3factorySetLmPoolDeployerIterator struct {
	Event *PancakeV3factorySetLmPoolDeployer // Event containing the contract specifics and raw log

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
func (it *PancakeV3factorySetLmPoolDeployerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3factorySetLmPoolDeployer)
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
		it.Event = new(PancakeV3factorySetLmPoolDeployer)
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
func (it *PancakeV3factorySetLmPoolDeployerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3factorySetLmPoolDeployerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3factorySetLmPoolDeployer represents a SetLmPoolDeployer event raised by the PancakeV3factory contract.
type PancakeV3factorySetLmPoolDeployer struct {
	LmPoolDeployer common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetLmPoolDeployer is a free log retrieval operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3factory *PancakeV3factoryFilterer) FilterSetLmPoolDeployer(opts *bind.FilterOpts, lmPoolDeployer []common.Address) (*PancakeV3factorySetLmPoolDeployerIterator, error) {

	var lmPoolDeployerRule []interface{}
	for _, lmPoolDeployerItem := range lmPoolDeployer {
		lmPoolDeployerRule = append(lmPoolDeployerRule, lmPoolDeployerItem)
	}

	logs, sub, err := _PancakeV3factory.contract.FilterLogs(opts, "SetLmPoolDeployer", lmPoolDeployerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factorySetLmPoolDeployerIterator{contract: _PancakeV3factory.contract, event: "SetLmPoolDeployer", logs: logs, sub: sub}, nil
}

// WatchSetLmPoolDeployer is a free log subscription operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3factory *PancakeV3factoryFilterer) WatchSetLmPoolDeployer(opts *bind.WatchOpts, sink chan<- *PancakeV3factorySetLmPoolDeployer, lmPoolDeployer []common.Address) (event.Subscription, error) {

	var lmPoolDeployerRule []interface{}
	for _, lmPoolDeployerItem := range lmPoolDeployer {
		lmPoolDeployerRule = append(lmPoolDeployerRule, lmPoolDeployerItem)
	}

	logs, sub, err := _PancakeV3factory.contract.WatchLogs(opts, "SetLmPoolDeployer", lmPoolDeployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3factorySetLmPoolDeployer)
				if err := _PancakeV3factory.contract.UnpackLog(event, "SetLmPoolDeployer", log); err != nil {
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

// ParseSetLmPoolDeployer is a log parse operation binding the contract event 0x4c912280cda47bed324de14f601d3f125a98254671772f3f1f491e50fa0ca407.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (_PancakeV3factory *PancakeV3factoryFilterer) ParseSetLmPoolDeployer(log types.Log) (*PancakeV3factorySetLmPoolDeployer, error) {
	event := new(PancakeV3factorySetLmPoolDeployer)
	if err := _PancakeV3factory.contract.UnpackLog(event, "SetLmPoolDeployer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeV3factoryWhiteListAddedIterator is returned from FilterWhiteListAdded and is used to iterate over the raw logs and unpacked data for WhiteListAdded events raised by the PancakeV3factory contract.
type PancakeV3factoryWhiteListAddedIterator struct {
	Event *PancakeV3factoryWhiteListAdded // Event containing the contract specifics and raw log

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
func (it *PancakeV3factoryWhiteListAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeV3factoryWhiteListAdded)
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
		it.Event = new(PancakeV3factoryWhiteListAdded)
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
func (it *PancakeV3factoryWhiteListAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeV3factoryWhiteListAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeV3factoryWhiteListAdded represents a WhiteListAdded event raised by the PancakeV3factory contract.
type PancakeV3factoryWhiteListAdded struct {
	User     common.Address
	Verified bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhiteListAdded is a free log retrieval operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3factory *PancakeV3factoryFilterer) FilterWhiteListAdded(opts *bind.FilterOpts, user []common.Address) (*PancakeV3factoryWhiteListAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeV3factory.contract.FilterLogs(opts, "WhiteListAdded", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeV3factoryWhiteListAddedIterator{contract: _PancakeV3factory.contract, event: "WhiteListAdded", logs: logs, sub: sub}, nil
}

// WatchWhiteListAdded is a free log subscription operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3factory *PancakeV3factoryFilterer) WatchWhiteListAdded(opts *bind.WatchOpts, sink chan<- *PancakeV3factoryWhiteListAdded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeV3factory.contract.WatchLogs(opts, "WhiteListAdded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeV3factoryWhiteListAdded)
				if err := _PancakeV3factory.contract.UnpackLog(event, "WhiteListAdded", log); err != nil {
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

// ParseWhiteListAdded is a log parse operation binding the contract event 0xaec42ac7f1bb8651906ae6522f50a19429e124e8ea678ef59fd27750759288a2.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (_PancakeV3factory *PancakeV3factoryFilterer) ParseWhiteListAdded(log types.Log) (*PancakeV3factoryWhiteListAdded, error) {
	event := new(PancakeV3factoryWhiteListAdded)
	if err := _PancakeV3factory.contract.UnpackLog(event, "WhiteListAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
