// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ERC20FixedMetaData contains all meta data concerning the ERC20Fixed contract.
var ERC20FixedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extensionData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20Fixed__BurningNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20Fixed__MintingNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20Fixed",
	Bin: "0x610180604052348015610010575f5ffd5b50604051613a21380380613a2183398101604081905261002f91610879565b8086868686868280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b03160361009457604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100bd5f82610309565b505050506100f17faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361030960201b60201c565b5080515f5b8181101561013057610128600484838151811061011557610115610953565b602002602001015161031d60201b60201c565b6001016100f6565b50505050816009908161014391906109ea565b50600a61015082826109ea565b506101609150839050600b610395565b6101205261016f81600c610395565b61014052815160208084019190912060e052815190820120610100524660a0526101fb60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361022f5760405163e1dea2ef60e01b8152636e616d6560e01b600482015260240161008b565b81515f0361025b5760405163e1dea2ef60e01b8152651cde5b589bdb60d21b600482015260240161008b565b60ff1661016052505082515f925082915061027f9084016020908101908501610aa4565b91509150815f036102b55760405163e1dea2ef60e01b81526c696e697469616c537570706c7960981b600482015260240161008b565b6001600160a01b0381166102f15760405163e1dea2ef60e01b81526f1a5b9a5d1a585b149958da5c1a595b9d60821b600482015260240161008b565b6102fb81836103c5565b505050505050505050610b49565b5f61031483836103f9565b90505b92915050565b6001600160a01b03811661034e5760405163e1dea2ef60e01b815264666f72676560d81b600482015260240161008b565b610358828261042c565b15610391576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f6020835110156103b0576103a983610440565b9050610317565b816103bb84826109ea565b5060ff9050610317565b6001600160a01b0382166103ee5760405163ec442f0560e01b81525f600482015260240161008b565b6103915f838361047d565b5f80610405848461048d565b90508015610314575f848152600360205260409020610424908461042c565b509392505050565b5f610314836001600160a01b0384166104f3565b5f5f829050601f8151111561046a578260405163305a27a960e01b815260040161008b9190610ad2565b805161047582610b07565b179392505050565b61048883838361053f565b505050565b5f826104e9575f6104a66002546001600160a01b031690565b6001600160a01b0316146104cd57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6103148383610665565b5f81815260018301602052604081205461053857508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610317565b505f610317565b6001600160a01b038316610569578060085f82825461055e9190610b2a565b909155506105d99050565b6001600160a01b0383165f90815260066020526040902054818110156105bb5760405163391434e360e21b81526001600160a01b0385166004820152602481018290526044810183905260640161008b565b6001600160a01b0384165f9081526006602052604090209082900390555b6001600160a01b0382166105f557600880548290039055610613565b6001600160a01b0382165f9081526006602052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161065891815260200190565b60405180910390a3505050565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610538575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556106bd3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610317565b6001600160a01b0381168114610719575f5ffd5b50565b805161072781610705565b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156107685761076861072c565b604052919050565b5f82601f83011261077f575f5ffd5b81516001600160401b038111156107985761079861072c565b8060051b6107a860208201610740565b918252602081850181019290810190868411156107c3575f5ffd5b6020860192505b838310156107ee5782516107dd81610705565b8252602092830192909101906107ca565b9695505050505050565b5f82601f830112610807575f5ffd5b8151602083015f806001600160401b038411156108265761082661072c565b50601f8301601f191660200161083b81610740565b91505082815285838301111561084f575f5ffd5b8282602083015e5f92810160200192909252509392505050565b805160ff81168114610727575f5ffd5b5f5f5f5f5f5f60c0878903121561088e575f5ffd5b6108978761071c565b60208801519096506001600160401b038111156108b2575f5ffd5b6108be89828a01610770565b604089015190965090506001600160401b038111156108db575f5ffd5b6108e789828a016107f8565b606089015190955090506001600160401b03811115610904575f5ffd5b61091089828a016107f8565b93505061091f60808801610869565b60a08801519092506001600160401b0381111561093a575f5ffd5b61094689828a016107f8565b9150509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061097b57607f821691505b60208210810361099957634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561048857805f5260205f20601f840160051c810160208510156109c45750805b601f840160051c820191505b818110156109e3575f81556001016109d0565b5050505050565b81516001600160401b03811115610a0357610a0361072c565b610a1781610a118454610967565b8461099f565b6020601f821160018114610a49575f8315610a325750848201515b5f19600385901b1c1916600184901b1784556109e3565b5f84815260208120601f198516915b82811015610a785787850151825560209485019460019092019101610a58565b5084821015610a9557868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f5f60408385031215610ab5575f5ffd5b825191506020830151610ac781610705565b809150509250929050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610999575f1960209190910360031b1b16919050565b8082018082111561031757634e487b7160e01b5f52601160045260245ffd5b60805160a05160c05160e05161010051610120516101405161016051612e7d610ba45f395f6103ac01525f6113a601525f61137901525f6110f901525f6110d101525f61102c01525f61105601525f6110800152612e7d5ff3fe608060405234801561000f575f5ffd5b50600436106102cd575f3560e01c806387f453531161017c578063a9059cbb116100dd578063d505accf11610093578063dd62ed3e1161006e578063dd62ed3e14610657578063ec87621c1461066a578063fe2df3e814610691575f5ffd5b8063d505accf14610629578063d547741f1461063c578063d602b9fd1461064f575f5ffd5b8063cc8463c8116100c3578063cc8463c8146105cd578063cefc1429146105d5578063cf6eefb7146105dd575f5ffd5b8063a9059cbb146105a7578063ca15c873146105ba575f5ffd5b80639ca92df911610132578063a217fddf11610118578063a217fddf14610585578063a3246ad31461058c578063a40283d51461059f575f5ffd5b80639ca92df91461054b578063a1eda53c1461055e575f5ffd5b80639010d07c116101625780639010d07c146104ed57806391d148541461050057806395d89b4114610543575f5ffd5b806387f45353146104d25780638da5cb5b146104e5575f5ffd5b80633644e51511610231578063649a5ec7116101e75780637ecebe00116101c25780637ecebe001461046557806384b0196e1461047857806384ef8ffc14610493575f5ffd5b8063649a5ec71461042c57806370a082311461043f57806379cc679014610452575f5ffd5b806340c10f191161021757806340c10f19146103f15780635c4e62c414610404578063634e93da14610419575f5ffd5b80633644e515146103d657806336568abe146103de575f5ffd5b806318160ddd11610286578063248a9ca31161026c578063248a9ca3146103705780632f2ff15d14610392578063313ce567146103a5575f5ffd5b806318160ddd1461034757806323b872dd1461035d575f5ffd5b806306fdde03116102b657806306fdde0314610315578063095ea7b31461032a5780630aa6220b1461033d575f5ffd5b806301ffc9a7146102d1578063022d63fb146102f9575b5f5ffd5b6102e46102df3660046128bb565b6106a4565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016102f0565b61031d6108a6565b6040516102f09190612946565b6102e4610338366004612980565b6108b5565b6103456108c7565b005b61034f6108dc565b6040519081526020016102f0565b6102e461036b3660046129a8565b6108e6565b61034f61037e3660046129e2565b5f9081526020819052604090206001015490565b6103456103a03660046129f9565b6108fa565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016102f0565b61034f610908565b6103456103ec3660046129f9565b610911565b6103456103ff366004612980565b61091b565b61040c61094d565b6040516102f09190612a23565b610345610427366004612a7b565b610959565b61034561043a366004612a94565b61096c565b61034f61044d366004612a7b565b61097f565b610345610460366004612980565b6109a9565b61034f610473366004612a7b565b6109db565b6104806109e5565b6040516102f09796959493929190612ab9565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102f0565b6102e46104e0366004612a7b565b610a43565b6104ad610a4f565b6104ad6104fb366004612b78565b610a6f565b6102e461050e3660046129f9565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61031d610a86565b6104ad6105593660046129e2565b610a90565b610566610a9c565b6040805165ffffffffffff9384168152929091166020830152016102f0565b61034f5f81565b61040c61059a3660046129e2565b610b16565b61034f610b2f565b6102e46105b5366004612980565b610b3a565b61034f6105c83660046129e2565b610b45565b6102fe610b5b565b610345610bf8565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102f0565b610345610637366004612b98565b610c59565b61034561064a3660046129f9565b610e02565b610345610e0c565b61034f610665366004612c05565b610e1e565b61034f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61034561069f366004612c2d565b610e57565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061073657507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b8061078257507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b806107ce57507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b806107f957507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061084557507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b8061089157507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b806108a057506108a082610ee8565b92915050565b60606108b0610ef2565b905090565b5f6108c08383610f82565b9392505050565b5f6108d181610f99565b6108d9610fa3565b50565b5f6108b060085490565b5f6108f2848484610faf565b949350505050565b6109048282610fd2565b5050565b5f6108b0611013565b6109048282611149565b6040517f21b8298300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60606108b0600461124e565b5f61096381610f99565b6109048261125a565b5f61097681610f99565b610904826112d9565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600660205260408120546108a0565b6040517fd342df9000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6108a082611348565b5f6060805f5f5f60606109f6611372565b6109fe61139f565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f6108a06004836113cc565b5f6108b060025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206108c090836113fa565b60606108b0611405565b5f6108a06004836113fa565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610ade57504265ffffffffffff821610155b610ae9575f5f610b0e565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f8181526003602052604090206060906108a09061124e565b5f6108b06004611414565b5f6108c0838361141d565b5f8181526003602052604081206108a090611414565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610b9c57504265ffffffffffff8216105b610bce576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610bf2565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610c51576040517fc22c80220000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6108d961142a565b83421115610c96576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610c48565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610cee8c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610d558261151b565b90505f610d6482878787611562565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610deb576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610c48565b610df68a8a8a61158e565b50505050505050505050565b61090482826115a0565b5f610e1681610f99565b6108d96115e1565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081526007602090815260408083209385168352929052908120546108c0565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610e8181610f99565b6128b382610e91576115eb610e95565b61163e5b9050835f5b81811015610edf57610ed76004888884818110610eb957610eb9612cb3565b9050602002016020810190610ece9190612a7b565b8563ffffffff16565b600101610e9a565b50505050505050565b5f6108a082611700565b606060098054610f0190612ce0565b80601f0160208091040260200160405190810160405280929190818152602001828054610f2d90612ce0565b8015610f785780601f10610f4f57610100808354040283529160200191610f78565b820191905f5260205f20905b815481529060010190602001808311610f5b57829003601f168201915b5050505050905090565b5f33610f8f81858561158e565b5060019392505050565b6108d9813361170a565b610fad5f5f61178f565b565b5f33610fbc8582856118e8565b610fc78585856118f3565b506001949350505050565b81611009576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610904828261199c565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561107857507f000000000000000000000000000000000000000000000000000000000000000046145b156110a257507f000000000000000000000000000000000000000000000000000000000000000090565b6108b0604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b81158015611171575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156112445760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806111c5575065ffffffffffff8116155b806111d857504265ffffffffffff821610155b15611219576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610c48565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b61090482826119c6565b60605f6108c083611a1f565b5f611263610b5b565b61126c42611a78565b6112769190612d5e565b90506112828282611ac7565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6112e382611b62565b6112ec42611a78565b6112f69190612d5e565b9050611302828261178f565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d60205260408120546108a0565b60606108b07f0000000000000000000000000000000000000000000000000000000000000000600b611ba9565b60606108b07f0000000000000000000000000000000000000000000000000000000000000000600c611ba9565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156108c0565b5f6108c08383611c52565b6060600a8054610f0190612ce0565b5f6108a0825490565b5f33610f8f8185856118f3565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061147a57504265ffffffffffff821610155b156114bb576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610c48565b6114e35f6114de60025473ffffffffffffffffffffffffffffffffffffffff1690565b611c78565b506114ee5f83611c83565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f6108a0611527611013565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f61157288888888611c8e565b9250925092506115828282611d81565b50909695505050505050565b61159b8383836001611e84565b505050565b816115d7576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109048282611e90565b610fad5f5f611ac7565b6115f58282611eb4565b156109045760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff81166116ad576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610c48565b6116b78282611ed5565b156109045760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f6108a082611ef6565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610904576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610c48565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611863574265ffffffffffff8216101561183a576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611863565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b61159b838383611f4b565b73ffffffffffffffffffffffffffffffffffffffff8316611942576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610c48565b73ffffffffffffffffffffffffffffffffffffffff8216611991576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610c48565b61159b838383611fee565b5f828152602081905260409020600101546119b681610f99565b6119c08383611c83565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611a15576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61159b8282611c78565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611a6c57602002820191905f5260205f20905b815481526020019060010190808311611a58575b50505050509050919050565b5f65ffffffffffff821115611ac3576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610c48565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff88161717909355900416801561159b576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611b6c610b5b565b90508065ffffffffffff168365ffffffffffff1611611b9457611b8f8382612d7c565b6108c0565b6108c065ffffffffffff841662069780611ff9565b606060ff8314611bc357611bbc83612008565b90506108a0565b818054611bcf90612ce0565b80601f0160208091040260200160405190810160405280929190818152602001828054611bfb90612ce0565b8015611c465780601f10611c1d57610100808354040283529160200191611c46565b820191905f5260205f20905b815481529060010190602001808311611c2957829003601f168201915b505050505090506108a0565b5f825f018281548110611c6757611c67612cb3565b905f5260205f200154905092915050565b5f6108c08383612045565b5f6108c08383612078565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611cc757505f91506003905082611d77565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611d18573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611d6e57505f925060019150829050611d77565b92505f91508190505b9450945094915050565b5f826003811115611d9457611d94612d9a565b03611d9d575050565b6001826003811115611db157611db1612d9a565b03611de8576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115611dfc57611dfc612d9a565b03611e36576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610c48565b6003826003811115611e4a57611e4a612d9a565b03610904576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610c48565b6119c0848484846120a3565b5f82815260208190526040902060010154611eaa81610f99565b6119c08383611c78565b5f6108c08373ffffffffffffffffffffffffffffffffffffffff84166121e8565b5f6108c08373ffffffffffffffffffffffffffffffffffffffff84166122cb565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806108a057506108a082612317565b5f611f568484610e1e565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156119c05781811015611fe0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610c48565b6119c084848484035f611e84565b61159b83838361236c565b5f8282188284100282186108c0565b60605f61201483612513565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f6120518484612553565b905080156108c0575f8481526003602052604090206120709084611eb4565b509392505050565b5f5f61208484846125b4565b905080156108c0575f8481526003602052604090206120709084611ed5565b73ffffffffffffffffffffffffffffffffffffffff84166120f2576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610c48565b73ffffffffffffffffffffffffffffffffffffffff8316612141576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610c48565b73ffffffffffffffffffffffffffffffffffffffff8085165f90815260076020908152604080832093871683529290522082905580156119c0578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516121da91815260200190565b60405180910390a350505050565b5f81815260018301602052604081205480156122c2575f61220a600183612dc7565b85549091505f9061221d90600190612dc7565b905080821461227c575f865f01828154811061223b5761223b612cb3565b905f5260205f200154905080875f01848154811061225b5761225b612cb3565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061228d5761228d612dda565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506108a0565b5f9150506108a0565b5f81815260018301602052604081205461231057508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556108a0565b505f6108a0565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806108a057506108a082612672565b73ffffffffffffffffffffffffffffffffffffffff83166123a3578060085f8282546123989190612e07565b909155506124539050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526006602052604090205481811015612428576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610c48565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661247c576008805482900390556124a7565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161250691815260200190565b60405180910390a3505050565b5f60ff8216601f8111156108a0576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8215801561257c575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b156125aa57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6108c08383612708565b5f82612668575f6125da60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612627576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6108c083836127c1565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108a057507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146108a0565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615612310575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016108a0565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16612310575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556128513390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016108a0565b610fad612e1a565b5f602082840312156128cb575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146108c0575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6108c060208301846128fa565b803573ffffffffffffffffffffffffffffffffffffffff8116811461297b575f5ffd5b919050565b5f5f60408385031215612991575f5ffd5b61299a83612958565b946020939093013593505050565b5f5f5f606084860312156129ba575f5ffd5b6129c384612958565b92506129d160208501612958565b929592945050506040919091013590565b5f602082840312156129f2575f5ffd5b5035919050565b5f5f60408385031215612a0a575f5ffd5b82359150612a1a60208401612958565b90509250929050565b602080825282518282018190525f918401906040840190835b81811015612a7057835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612a3c565b509095945050505050565b5f60208284031215612a8b575f5ffd5b6108c082612958565b5f60208284031215612aa4575f5ffd5b813565ffffffffffff811681146108c0575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f612af360e08301896128fa565b8281036040840152612b0581896128fa565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015612b67578351835260209384019390920191600101612b49565b50909b9a5050505050505050505050565b5f5f60408385031215612b89575f5ffd5b50508035926020909101359150565b5f5f5f5f5f5f5f60e0888a031215612bae575f5ffd5b612bb788612958565b9650612bc560208901612958565b95506040880135945060608801359350608088013560ff81168114612be8575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215612c16575f5ffd5b612c1f83612958565b9150612a1a60208401612958565b5f5f5f60408486031215612c3f575f5ffd5b833567ffffffffffffffff811115612c55575f5ffd5b8401601f81018613612c65575f5ffd5b803567ffffffffffffffff811115612c7b575f5ffd5b8660208260051b8401011115612c8f575f5ffd5b6020918201945092508401358015158114612ca8575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c90821680612cf457607f821691505b602082108103612d2b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff81811683821601908111156108a0576108a0612d31565b65ffffffffffff82811682821603908111156108a0576108a0612d31565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b818103818111156108a0576108a0612d31565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b808201808211156108a0576108a0612d31565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212201e4c5ba7947c603788313152bdf6ae85cca7e287ec26d24a37c207fecd1d93bd64736f6c634300081c0033",
}

// ERC20Fixed is an auto generated Go binding around an Ethereum contract.
type ERC20Fixed struct {
	abi abi.ABI
}

// NewERC20Fixed creates a new instance of ERC20Fixed.
func NewERC20Fixed() *ERC20Fixed {
	parsed, err := ERC20FixedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Fixed{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Fixed) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals, bytes extensionData) returns()
func (eRC20Fixed *ERC20Fixed) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8, extensionData []byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("", owner, forges, name, symbol, decimals, extensionData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20Fixed.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Fixed.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackDOMAINSEPARATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Fixed.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Fixed.abi.Unpack("DOMAIN_SEPARATOR", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackMANAGERROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) PackMANAGERROLE() []byte {
	enc, err := eRC20Fixed.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Fixed.abi.Unpack("MANAGER_ROLE", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackAcceptDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcefc1429.
//
// Solidity: function acceptDefaultAdminTransfer() returns()
func (eRC20Fixed *ERC20Fixed) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20Fixed.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBeginDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (eRC20Fixed *ERC20Fixed) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address , uint256 ) pure returns()
func (eRC20Fixed *ERC20Fixed) PackBurnFrom(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("burnFrom", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20Fixed *ERC20Fixed) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20Fixed.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC20Fixed *ERC20Fixed) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Fixed *ERC20Fixed) PackDecimals() []byte {
	enc, err := eRC20Fixed.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Fixed *ERC20Fixed) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Fixed.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, err
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20Fixed *ERC20Fixed) PackDefaultAdmin() []byte {
	enc, err := eRC20Fixed.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("defaultAdmin", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20Fixed *ERC20Fixed) PackDefaultAdminDelay() []byte {
	enc, err := eRC20Fixed.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("defaultAdminDelay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDefaultAdminDelayIncreaseWait is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20Fixed *ERC20Fixed) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20Fixed.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("defaultAdminDelayIncreaseWait", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Fixed *ERC20Fixed) PackEip712Domain() []byte {
	enc, err := eRC20Fixed.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Fixed *ERC20Fixed) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Fixed.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Fixed *ERC20Fixed) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Fixed *ERC20Fixed) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("forgeByIndex", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackForgeCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Fixed *ERC20Fixed) PackForgeCount() []byte {
	enc, err := eRC20Fixed.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("forgeCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Fixed *ERC20Fixed) PackForges() []byte {
	enc, err := eRC20Fixed.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Fixed *ERC20Fixed) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("forges", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20Fixed *ERC20Fixed) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20Fixed.abi.Unpack("getRoleAdmin", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackGetRoleMember is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20Fixed *ERC20Fixed) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20Fixed *ERC20Fixed) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("getRoleMember", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetRoleMemberCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("getRoleMemberCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetRoleMembers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20Fixed *ERC20Fixed) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20Fixed *ERC20Fixed) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("getRoleMembers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGrantRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (eRC20Fixed *ERC20Fixed) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Fixed *ERC20Fixed) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsForge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Fixed *ERC20Fixed) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address , uint256 ) pure returns()
func (eRC20Fixed *ERC20Fixed) PackMint(arg0 common.Address, arg1 *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("mint", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Fixed *ERC20Fixed) PackName() []byte {
	enc, err := eRC20Fixed.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Fixed *ERC20Fixed) UnpackName(data []byte) (string, error) {
	out, err := eRC20Fixed.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackNonces is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("nonces", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Fixed *ERC20Fixed) PackOwner() []byte {
	enc, err := eRC20Fixed.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Fixed *ERC20Fixed) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20Fixed.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPendingDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20Fixed *ERC20Fixed) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20Fixed.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20Fixed *ERC20Fixed) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20Fixed.abi.Unpack("pendingDefaultAdmin", data)
	outstruct := new(PendingDefaultAdminOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPendingDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20Fixed *ERC20Fixed) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20Fixed.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20Fixed *ERC20Fixed) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20Fixed.abi.Unpack("pendingDefaultAdminDelay", data)
	outstruct := new(PendingDefaultAdminDelayOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewDelay = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20Fixed *ERC20Fixed) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (eRC20Fixed *ERC20Fixed) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20Fixed *ERC20Fixed) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Fixed.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20Fixed *ERC20Fixed) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20Fixed.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20Fixed *ERC20Fixed) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20Fixed.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Fixed *ERC20Fixed) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20Fixed.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Fixed *ERC20Fixed) PackSymbol() []byte {
	enc, err := eRC20Fixed.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Fixed *ERC20Fixed) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Fixed.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Fixed *ERC20Fixed) PackTotalSupply() []byte {
	enc, err := eRC20Fixed.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Fixed *ERC20Fixed) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Fixed.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Fixed.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Fixed *ERC20Fixed) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Fixed.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// ERC20FixedApproval represents a Approval event raised by the ERC20Fixed contract.
type ERC20FixedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20FixedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20FixedApproval) ContractEventName() string {
	return ERC20FixedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Fixed *ERC20Fixed) UnpackApprovalEvent(log *types.Log) (*ERC20FixedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedApproval)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20Fixed contract.
type ERC20FixedDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20FixedDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20FixedDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20FixedDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20FixedDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20Fixed contract.
type ERC20FixedDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20FixedDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20FixedDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20FixedDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20FixedDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20Fixed contract.
type ERC20FixedDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20FixedDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20FixedDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20FixedDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20FixedDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20Fixed contract.
type ERC20FixedDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20FixedDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20FixedDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20FixedDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20Fixed *ERC20Fixed) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20FixedDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Fixed contract.
type ERC20FixedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20FixedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20FixedEIP712DomainChanged) ContractEventName() string {
	return ERC20FixedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Fixed *ERC20Fixed) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20FixedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedForgeAdded represents a ForgeAdded event raised by the ERC20Fixed contract.
type ERC20FixedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20FixedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20FixedForgeAdded) ContractEventName() string {
	return ERC20FixedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20Fixed *ERC20Fixed) UnpackForgeAddedEvent(log *types.Log) (*ERC20FixedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedForgeRemoved represents a ForgeRemoved event raised by the ERC20Fixed contract.
type ERC20FixedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20FixedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20FixedForgeRemoved) ContractEventName() string {
	return ERC20FixedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20Fixed *ERC20Fixed) UnpackForgeRemovedEvent(log *types.Log) (*ERC20FixedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20Fixed contract.
type ERC20FixedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20FixedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20FixedRoleAdminChanged) ContractEventName() string {
	return ERC20FixedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20Fixed *ERC20Fixed) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20FixedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedRoleGranted represents a RoleGranted event raised by the ERC20Fixed contract.
type ERC20FixedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20FixedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20FixedRoleGranted) ContractEventName() string {
	return ERC20FixedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Fixed *ERC20Fixed) UnpackRoleGrantedEvent(log *types.Log) (*ERC20FixedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedRoleRevoked represents a RoleRevoked event raised by the ERC20Fixed contract.
type ERC20FixedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20FixedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20FixedRoleRevoked) ContractEventName() string {
	return ERC20FixedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Fixed *ERC20Fixed) UnpackRoleRevokedEvent(log *types.Log) (*ERC20FixedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ERC20FixedTransfer represents a Transfer event raised by the ERC20Fixed contract.
type ERC20FixedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20FixedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20FixedTransfer) ContractEventName() string {
	return ERC20FixedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Fixed *ERC20Fixed) UnpackTransferEvent(log *types.Log) (*ERC20FixedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20Fixed.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20FixedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Fixed.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Fixed.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (eRC20Fixed *ERC20Fixed) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20FixedBurningNotAllowed"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20FixedBurningNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20FixedMintingNotAllowed"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20FixedMintingNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Fixed.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20Fixed.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20FixedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20Fixed contract.
type ERC20FixedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20FixedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20Fixed *ERC20Fixed) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20FixedAccessControlBadConfirmation, error) {
	out := new(ERC20FixedAccessControlBadConfirmation)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20Fixed contract.
type ERC20FixedAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20FixedAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20Fixed *ERC20Fixed) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20FixedAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20FixedAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20Fixed contract.
type ERC20FixedAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20FixedAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20Fixed *ERC20Fixed) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20FixedAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20FixedAccessControlEnforcedDefaultAdminRules)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20Fixed contract.
type ERC20FixedAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20FixedAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20Fixed *ERC20Fixed) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20FixedAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20FixedAccessControlInvalidDefaultAdmin)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20Fixed contract.
type ERC20FixedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20FixedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20Fixed *ERC20Fixed) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20FixedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20FixedAccessControlUnauthorizedAccount)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Fixed contract.
type ERC20FixedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20FixedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Fixed *ERC20Fixed) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20FixedECDSAInvalidSignature, error) {
	out := new(ERC20FixedECDSAInvalidSignature)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Fixed contract.
type ERC20FixedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20FixedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Fixed *ERC20Fixed) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20FixedECDSAInvalidSignatureLength, error) {
	out := new(ERC20FixedECDSAInvalidSignatureLength)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Fixed contract.
type ERC20FixedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20FixedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Fixed *ERC20Fixed) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20FixedECDSAInvalidSignatureS, error) {
	out := new(ERC20FixedECDSAInvalidSignatureS)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20FixedBurningNotAllowed represents a ERC20Fixed__BurningNotAllowed error raised by the ERC20Fixed contract.
type ERC20FixedERC20FixedBurningNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Fixed__BurningNotAllowed()
func ERC20FixedERC20FixedBurningNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xd342df9019e4f1569194e8182eed24aea86c9425b082b66fd407767649e53ed5")
}

// UnpackERC20FixedBurningNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Fixed__BurningNotAllowed()
func (eRC20Fixed *ERC20Fixed) UnpackERC20FixedBurningNotAllowedError(raw []byte) (*ERC20FixedERC20FixedBurningNotAllowed, error) {
	out := new(ERC20FixedERC20FixedBurningNotAllowed)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20FixedBurningNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20FixedMintingNotAllowed represents a ERC20Fixed__MintingNotAllowed error raised by the ERC20Fixed contract.
type ERC20FixedERC20FixedMintingNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Fixed__MintingNotAllowed()
func ERC20FixedERC20FixedMintingNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x21b82983de03f0653d31e4c2f78db5241001ff57ffdbc14aa9d7a41d9c814b1e")
}

// UnpackERC20FixedMintingNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Fixed__MintingNotAllowed()
func (eRC20Fixed *ERC20Fixed) UnpackERC20FixedMintingNotAllowedError(raw []byte) (*ERC20FixedERC20FixedMintingNotAllowed, error) {
	out := new(ERC20FixedERC20FixedMintingNotAllowed)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20FixedMintingNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Fixed contract.
type ERC20FixedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20FixedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20FixedERC20InsufficientAllowance, error) {
	out := new(ERC20FixedERC20InsufficientAllowance)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Fixed contract.
type ERC20FixedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20FixedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20FixedERC20InsufficientBalance, error) {
	out := new(ERC20FixedERC20InsufficientBalance)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20FixedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidApproverError(raw []byte) (*ERC20FixedERC20InvalidApprover, error) {
	out := new(ERC20FixedERC20InvalidApprover)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20FixedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20FixedERC20InvalidReceiver, error) {
	out := new(ERC20FixedERC20InvalidReceiver)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20FixedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidSenderError(raw []byte) (*ERC20FixedERC20InvalidSender, error) {
	out := new(ERC20FixedERC20InvalidSender)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Fixed contract.
type ERC20FixedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20FixedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Fixed *ERC20Fixed) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20FixedERC20InvalidSpender, error) {
	out := new(ERC20FixedERC20InvalidSpender)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Fixed contract.
type ERC20FixedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20FixedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Fixed *ERC20Fixed) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20FixedERC2612ExpiredSignature, error) {
	out := new(ERC20FixedERC2612ExpiredSignature)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Fixed contract.
type ERC20FixedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20FixedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Fixed *ERC20Fixed) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20FixedERC2612InvalidSigner, error) {
	out := new(ERC20FixedERC2612InvalidSigner)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Fixed contract.
type ERC20FixedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20FixedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Fixed *ERC20Fixed) UnpackInvalidAccountNonceError(raw []byte) (*ERC20FixedInvalidAccountNonce, error) {
	out := new(ERC20FixedInvalidAccountNonce)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedInvalidShortString represents a InvalidShortString error raised by the ERC20Fixed contract.
type ERC20FixedInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20FixedInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20Fixed *ERC20Fixed) UnpackInvalidShortStringError(raw []byte) (*ERC20FixedInvalidShortString, error) {
	out := new(ERC20FixedInvalidShortString)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20Fixed contract.
type ERC20FixedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20FixedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20Fixed *ERC20Fixed) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20FixedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20FixedSafeCastOverflowedUintDowncast)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedStringTooLong represents a StringTooLong error raised by the ERC20Fixed contract.
type ERC20FixedStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20FixedStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20Fixed *ERC20Fixed) UnpackStringTooLongError(raw []byte) (*ERC20FixedStringTooLong, error) {
	out := new(ERC20FixedStringTooLong)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20Fixed contract.
type ERC20FixedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20FixedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20Fixed *ERC20Fixed) UnpackTokenBaseNullInputError(raw []byte) (*ERC20FixedTokenBaseNullInput, error) {
	out := new(ERC20FixedTokenBaseNullInput)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20FixedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20Fixed contract.
type ERC20FixedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20FixedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20Fixed *ERC20Fixed) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20FixedTokenBaseOnlyForge, error) {
	out := new(ERC20FixedTokenBaseOnlyForge)
	if err := eRC20Fixed.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
