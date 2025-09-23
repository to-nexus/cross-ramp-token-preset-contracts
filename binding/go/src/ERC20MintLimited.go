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

// ERC20MintLimitedMetaData contains all meta data concerning the ERC20MintLimited contract.
var ERC20MintLimitedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extensionData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"updateMintLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimits\",\"type\":\"uint256\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodManager__InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offsetSeconds\",\"type\":\"int256\"}],\"name\":\"PeriodManager__InvalidOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20MintLimited",
	Bin: "0x610180604052348015610010575f5ffd5b50604051613e09380380613e0983398101604081905261002f9161078a565b8086868686868280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b03160361009457604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100bd5f8261032c565b505050506100f17faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361032c60201b60201c565b5080515f5b8181101561013057610128600484838151811061011557610115610864565b602002602001015161034060201b60201c565b6001016100f6565b50505050816009908161014391906108fc565b50600a61015082826108fc565b506101609150839050600b6103b8565b6101205261016f81600c6103b8565b61014052815160208084019190912060e052815190820120610100524660a0526101fb60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361022f5760405163e1dea2ef60e01b8152636e616d6560e01b600482015260240161008b565b81515f0361025b5760405163e1dea2ef60e01b8152651cde5b589bdb60d21b600482015260240161008b565b60ff1661016052505082515f9250829150819061028190850160209081019086016109b6565b925092509250825f1480610293575080155b156102c55760405163e1dea2ef60e01b81526e1b1a5b5a5d081bdc881c195c9a5bd9608a1b600482015260240161008b565b60405180604001604052806102df856103e860201b60201c565b6001600160801b031681526020016102f68461041f565b600f0b905280516020909101516001600160801b03908116600160801b02911617600e5560105550610a39975050505050505050565b5f6103378383610453565b90505b92915050565b6001600160a01b0381166103715760405163e1dea2ef60e01b815264666f72676560d81b600482015260240161008b565b61037b8282610486565b156103b4576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f6020835110156103d3576103cc8361049a565b905061033a565b816103de84826108fc565b5060ff905061033a565b5f6001600160801b0382111561041b576040516306dfcc6560e41b8152608060048201526024810183905260440161008b565b5090565b80600f81900b811461044e5760405163327269a760e01b8152608060048201526024810183905260440161008b565b919050565b5f8061045f84846104d7565b90508015610337575f84815260036020526040902061047e9084610486565b509392505050565b5f610337836001600160a01b03841661053d565b5f5f829050601f815111156104c4578260405163305a27a960e01b815260040161008b91906109e1565b80516104cf82610a16565b179392505050565b5f82610533575f6104f06002546001600160a01b031690565b6001600160a01b03161461051757604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6103378383610589565b5f81815260018301602052604081205461058257508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561033a565b505f61033a565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610582575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556105e13390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161033a565b80516001600160a01b038116811461044e575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561067b5761067b61063f565b604052919050565b5f82601f830112610692575f5ffd5b81516001600160401b038111156106ab576106ab61063f565b8060051b6106bb60208201610653565b918252602081850181019290810190868411156106d6575f5ffd5b6020860192505b838310156106ff576106ee83610629565b8252602092830192909101906106dd565b9695505050505050565b5f82601f830112610718575f5ffd5b8151602083015f806001600160401b038411156107375761073761063f565b50601f8301601f191660200161074c81610653565b915050828152858383011115610760575f5ffd5b8282602083015e5f92810160200192909252509392505050565b805160ff8116811461044e575f5ffd5b5f5f5f5f5f5f60c0878903121561079f575f5ffd5b6107a887610629565b60208801519096506001600160401b038111156107c3575f5ffd5b6107cf89828a01610683565b604089015190965090506001600160401b038111156107ec575f5ffd5b6107f889828a01610709565b606089015190955090506001600160401b03811115610815575f5ffd5b61082189828a01610709565b9350506108306080880161077a565b60a08801519092506001600160401b0381111561084b575f5ffd5b61085789828a01610709565b9150509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061088c57607f821691505b6020821081036108aa57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156108f757805f5260205f20601f840160051c810160208510156108d55750805b601f840160051c820191505b818110156108f4575f81556001016108e1565b50505b505050565b81516001600160401b038111156109155761091561063f565b610929816109238454610878565b846108b0565b6020601f82116001811461095b575f83156109445750848201515b5f19600385901b1c1916600184901b1784556108f4565b5f84815260208120601f198516915b8281101561098a578785015182556020948501946001909201910161096a565b50848210156109a757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f5f5f606084860312156109c8575f5ffd5b5050815160208301516040909301519094929350919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156108aa575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051613375610a945f395f6103eb01525f61162601525f6115f901525f61124d01525f61122501525f61118001525f6111aa01525f6111d401526133755ff3fe608060405234801561000f575f5ffd5b5060043610610304575f3560e01c806384ef8ffc1161019d578063bdf7acce116100e8578063d547741f11610093578063e01d55c51161006e578063e01d55c5146106f8578063ec87621c1461070b578063fe2df3e814610732575f5ffd5b8063d547741f146106ca578063d602b9fd146106dd578063dd62ed3e146106e5575f5ffd5b8063cefc1429116100c3578063cefc142914610663578063cf6eefb71461066b578063d505accf146106b7575f5ffd5b8063bdf7acce14610640578063ca15c87314610648578063cc8463c81461065b575f5ffd5b80639ca92df911610148578063a3246ad311610123578063a3246ad314610612578063a40283d514610625578063a9059cbb1461062d575f5ffd5b80639ca92df9146105d1578063a1eda53c146105e4578063a217fddf1461060b575f5ffd5b80639010d07c116101785780639010d07c1461057357806391d148541461058657806395d89b41146105c9575f5ffd5b806384ef8ffc1461051957806387f45353146105585780638da5cb5b1461056b575f5ffd5b80633644e5151161025d578063634e93da1161020857806379cc6790116101e357806379cc6790146104d85780637ecebe00146104eb57806384b0196e146104fe575f5ffd5b8063634e93da1461049f578063649a5ec7146104b257806370a08231146104c5575f5ffd5b8063436ba62611610238578063436ba626146104435780635068b4de1461044b5780635c4e62c41461048a575f5ffd5b80633644e5151461041557806336568abe1461041d57806340c10f1914610430575f5ffd5b806318160ddd116102bd5780632d9ed80d116102985780632d9ed80d146103c95780632f2ff15d146103d1578063313ce567146103e4575f5ffd5b806318160ddd1461037e57806323b872dd14610394578063248a9ca3146103a7575f5ffd5b806306fdde03116102ed57806306fdde031461034c578063095ea7b3146103615780630aa6220b14610374575f5ffd5b806301ffc9a714610308578063022d63fb14610330575b5f5ffd5b61031b610316366004612d50565b610745565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff9091168152602001610327565b610354610947565b6040516103279190612ddb565b61031b61036f366004612e15565b610956565b61037c610968565b005b61038661097d565b604051908152602001610327565b61031b6103a2366004612e3d565b610987565b6103866103b5366004612e77565b5f9081526020819052604090206001015490565b601054610386565b61037c6103df366004612e8e565b61099b565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610327565b6103866109a9565b61037c61042b366004612e8e565b6109b2565b61037c61043e366004612e15565b6109bc565b6103866109c6565b600e54604080516fffffffffffffffffffffffffffffffff83168152700100000000000000000000000000000000909204600f0b602083015201610327565b6104926109ef565b6040516103279190612eb8565b61037c6104ad366004612f10565b6109fb565b61037c6104c0366004612f29565b610a0e565b6103866104d3366004612f10565b610a21565b61037c6104e6366004612e15565b610a4b565b6103866104f9366004612f10565b610a7d565b610506610a87565b6040516103279796959493929190612f4e565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610327565b61031b610566366004612f10565b610ae5565b610533610af1565b61053361058136600461300d565b610b11565b61031b610594366004612e8e565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610354610b28565b6105336105df366004612e77565b610b32565b6105ec610b3e565b6040805165ffffffffffff938416815292909116602083015201610327565b6103865f81565b610492610620366004612e77565b610bb8565b610386610bd1565b61031b61063b366004612e15565b610bdc565b610386610be7565b610386610656366004612e77565b610bf2565b610335610c08565b61037c610ca5565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610327565b61037c6106c536600461302d565b610d06565b61037c6106d8366004612e8e565b610eaf565b61037c610eb9565b6103866106f336600461309a565b610ecb565b61037c610706366004612e77565b610f04565b6103867faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61037c6107403660046130c2565b610fab565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a042910000000000000000000000000000000000000000000000000000000014806107d757507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b8061082357507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b8061086f57507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b8061089a57507fffffffff000000000000000000000000000000000000000000000000000000008216155b806108e657507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b8061093257507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b8061094157506109418261103c565b92915050565b6060610951611046565b905090565b5f61096183836110d6565b9392505050565b5f610972816110ed565b61097a6110f7565b50565b5f61095160085490565b5f610993848484611103565b949350505050565b6109a58282611126565b5050565b5f610951611167565b6109a5828261129d565b6109a582826113a2565b5f6109dd600f54600e61145690919063ffffffff16565b6109e8575060115490565b5060105490565b60606109516004611469565b5f610a05816110ed565b6109a582611475565b5f610a18816110ed565b6109a5826114f4565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260066020526040812054610941565b73ffffffffffffffffffffffffffffffffffffffff82163314610a7357610a73823383611563565b6109a5828261156e565b5f610941826115c8565b5f6060805f5f5f6060610a986115f2565b610aa061161f565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f61094160048361164c565b5f61095160025473ffffffffffffffffffffffffffffffffffffffff1690565b5f828152600360205260408120610961908361167a565b6060610951611685565b5f61094160048361167a565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610b8057504265ffffffffffff821610155b610b8b575f5f610bb0565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f81815260036020526040902060609061094190611469565b5f6109516004611694565b5f610961838361169d565b5f610951600e6116aa565b5f81815260036020526040812061094190611694565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610c4957504265ffffffffffff8216105b610c7b576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610c9f565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610cfe576040517fc22c80220000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61097a6116b5565b83421115610d43576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610cf5565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610d9b8c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610e02826117a6565b90505f610e11828787876117ed565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e98576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610cf5565b610ea38a8a8a611819565b50505050505050505050565b6109a58282611826565b5f610ec3816110ed565b61097a611867565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152600760209081526040808320938516835292905290812054610961565b5f610f0e816110ed565b815f03610f69576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69740000000000000000000000000000000000000000000000006004820152602401610cf5565b60105460408051918252602082018490527f864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e482910160405180910390a150601055565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610fd5816110ed565b612d4882610fe557611871610fe9565b6118c45b9050835f5b818110156110335761102b600488888481811061100d5761100d613148565b90506020020160208101906110229190612f10565b8563ffffffff16565b600101610fee565b50505050505050565b5f61094182611986565b60606009805461105590613175565b80601f016020809104026020016040519081016040528092919081815260200182805461108190613175565b80156110cc5780601f106110a3576101008083540402835291602001916110cc565b820191905f5260205f20905b8154815290600101906020018083116110af57829003601f168201915b5050505050905090565b5f336110e3818585611819565b5060019392505050565b61097a8133611990565b6111015f5f611a15565b565b5f33611110858285611563565b61111b858585611b6e565b506001949350505050565b8161115d576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109a58282611c17565b5f3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156111cc57507f000000000000000000000000000000000000000000000000000000000000000046145b156111f657507f000000000000000000000000000000000000000000000000000000000000000090565b610951604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b811580156112c5575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156113985760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611319575065ffffffffffff8116155b8061132c57504265ffffffffffff821610155b1561136d576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610cf5565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6109a58282611c41565b6011545f6113b0600e6116aa565b9050600f5481146113fc57600f81905560105460405181815290925081907ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d19060200160405180910390a25b5081811015611441576040517f156d32c50000000000000000000000000000000000000000000000000000000081526004810183905260248101829052604401610cf5565b8181036011556114518383611c9a565b505050565b5f81611461846116aa565b119392505050565b60605f61096183611ce5565b5f61147e610c08565b61148742611d3e565b61149191906131f3565b905061149d8282611d8d565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6114fe82611e28565b61150742611d3e565b61151191906131f3565b905061151d8282611a15565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b611451838383611e6f565b73ffffffffffffffffffffffffffffffffffffffff82166115bd576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610cf5565b6109a5825f83611f12565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d6020526040812054610941565b60606109517f0000000000000000000000000000000000000000000000000000000000000000600b611f1d565b60606109517f0000000000000000000000000000000000000000000000000000000000000000600c611f1d565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610961565b5f6109618383611fc6565b6060600a805461105590613175565b5f610941825490565b5f336110e3818585611b6e565b5f6109418242611fec565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061170557504265ffffffffffff821610155b15611746576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610cf5565b61176e5f61176960025473ffffffffffffffffffffffffffffffffffffffff1690565b612161565b506117795f8361216c565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f6109416117b2611167565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f6117fd88888888612177565b92509250925061180d828261226a565b50909695505050505050565b611451838383600161236d565b8161185d576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109a58282612379565b6111015f5f611d8d565b61187b828261239d565b156109a55760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611933576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610cf5565b61193d82826123be565b156109a55760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f610941826123df565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166109a5576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610cf5565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611ae9574265ffffffffffff82161015611ac0576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611ae9565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff8316611bbd576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610cf5565b73ffffffffffffffffffffffffffffffffffffffff8216611c0c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610cf5565b611451838383611f12565b5f82815260208190526040902060010154611c31816110ed565b611c3b838361216c565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611c90576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114518282612161565b611ca333610ae5565b611cdb576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610cf5565b6109a58282612434565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611d3257602002820191905f5260205f20905b815481526020019060010190808311611d1e575b50505050509050919050565b5f65ffffffffffff821115611d89576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610cf5565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611451576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611e32610c08565b90508065ffffffffffff168365ffffffffffff1611611e5a57611e558382613211565b610961565b61096165ffffffffffff84166206978061248e565b5f611e7a8484610ecb565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015611c3b5781811015611f04576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610cf5565b611c3b84848484035f61236d565b61145183838361249d565b606060ff8314611f3757611f3083612644565b9050610941565b818054611f4390613175565b80601f0160208091040260200160405190810160405280929190818152602001828054611f6f90613175565b8015611fba5780601f10611f9157610100808354040283529160200191611fba565b820191905f5260205f20905b815481529060010190602001808311611f9d57829003601f168201915b50505050509050610941565b5f825f018281548110611fdb57611fdb613148565b905f5260205f200154905092915050565b81545f906fffffffffffffffffffffffffffffffff16810361203a576040517f28632f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82546fffffffffffffffffffffffffffffffff16828161205c5761205c61322f565b8454919006909203917001000000000000000000000000000000009004600f0b5f03612089575080610941565b82545f700100000000000000000000000000000000909104600f0b13156120cd578254611f30907001000000000000000000000000000000009004600f0b8361325c565b82545f906120f1907001000000000000000000000000000000009004600f0b61326f565b90508083101561214f5783546040517ff41373a600000000000000000000000000000000000000000000000000000000815260048101859052700100000000000000000000000000000000909104600f0b6024820152604401610cf5565b61215981846132a5565b915050610941565b5f6109618383612681565b5f61096183836126b4565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156121b057505f91506003905082612260565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612201573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661225757505f925060019150829050612260565b92505f91508190505b9450945094915050565b5f82600381111561227d5761227d6132b8565b03612286575050565b600182600381111561229a5761229a6132b8565b036122d1576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156122e5576122e56132b8565b0361231f576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610cf5565b6003826003811115612333576123336132b8565b036109a5576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610cf5565b611c3b848484846126df565b5f82815260208190526040902060010154612393816110ed565b611c3b8383612161565b5f6109618373ffffffffffffffffffffffffffffffffffffffff8416612824565b5f6109618373ffffffffffffffffffffffffffffffffffffffff8416612907565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610941575061094182612953565b73ffffffffffffffffffffffffffffffffffffffff8216612483576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610cf5565b6109a55f8383611f12565b5f828218828410028218610961565b73ffffffffffffffffffffffffffffffffffffffff83166124d4578060085f8282546124c9919061325c565b909155506125849050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526006602052604090205481811015612559576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610cf5565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166125ad576008805482900390556125d8565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161263791815260200190565b60405180910390a3505050565b60605f612650836129a8565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f61268d84846129e8565b90508015610961575f8481526003602052604090206126ac908461239d565b509392505050565b5f5f6126c08484612a49565b90508015610961575f8481526003602052604090206126ac90846123be565b73ffffffffffffffffffffffffffffffffffffffff841661272e576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610cf5565b73ffffffffffffffffffffffffffffffffffffffff831661277d576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610cf5565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526007602090815260408083209387168352929052208290558015611c3b578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161281691815260200190565b60405180910390a350505050565b5f81815260018301602052604081205480156128fe575f6128466001836132a5565b85549091505f90612859906001906132a5565b90508082146128b8575f865f01828154811061287757612877613148565b905f5260205f200154905080875f01848154811061289757612897613148565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806128c9576128c96132e5565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610941565b5f915050610941565b5f81815260018301602052604081205461294c57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610941565b505f610941565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610941575061094182612b07565b5f60ff8216601f811115610941576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82158015612a11575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b15612a3f57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6109618383612b9d565b5f82612afd575f612a6f60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612abc576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6109618383612c56565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061094157507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610941565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff161561294c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610941565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1661294c575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612ce63390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610941565b611101613312565b5f60208284031215612d60575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610961575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109616020830184612d8f565b803573ffffffffffffffffffffffffffffffffffffffff81168114612e10575f5ffd5b919050565b5f5f60408385031215612e26575f5ffd5b612e2f83612ded565b946020939093013593505050565b5f5f5f60608486031215612e4f575f5ffd5b612e5884612ded565b9250612e6660208501612ded565b929592945050506040919091013590565b5f60208284031215612e87575f5ffd5b5035919050565b5f5f60408385031215612e9f575f5ffd5b82359150612eaf60208401612ded565b90509250929050565b602080825282518282018190525f918401906040840190835b81811015612f0557835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612ed1565b509095945050505050565b5f60208284031215612f20575f5ffd5b61096182612ded565b5f60208284031215612f39575f5ffd5b813565ffffffffffff81168114610961575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f612f8860e0830189612d8f565b8281036040840152612f9a8189612d8f565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015612ffc578351835260209384019390920191600101612fde565b50909b9a5050505050505050505050565b5f5f6040838503121561301e575f5ffd5b50508035926020909101359150565b5f5f5f5f5f5f5f60e0888a031215613043575f5ffd5b61304c88612ded565b965061305a60208901612ded565b95506040880135945060608801359350608088013560ff8116811461307d575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156130ab575f5ffd5b6130b483612ded565b9150612eaf60208401612ded565b5f5f5f604084860312156130d4575f5ffd5b833567ffffffffffffffff8111156130ea575f5ffd5b8401601f810186136130fa575f5ffd5b803567ffffffffffffffff811115613110575f5ffd5b8660208260051b8401011115613124575f5ffd5b602091820194509250840135801515811461313d575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c9082168061318957607f821691505b6020821081036131c0577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff8181168382160190811115610941576109416131c6565b65ffffffffffff8281168282160390811115610941576109416131c6565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b80820180821115610941576109416131c6565b5f7f8000000000000000000000000000000000000000000000000000000000000000820361329f5761329f6131c6565b505f0390565b81810381811115610941576109416131c6565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220865ccafd340555b5fd9ac8fa0399ecac2c2611309c04a699c1eb5479a8bd819864736f6c634300081c0033",
}

// ERC20MintLimited is an auto generated Go binding around an Ethereum contract.
type ERC20MintLimited struct {
	abi abi.ABI
}

// NewERC20MintLimited creates a new instance of ERC20MintLimited.
func NewERC20MintLimited() *ERC20MintLimited {
	parsed, err := ERC20MintLimitedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20MintLimited{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20MintLimited) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals, bytes extensionData) returns()
func (eRC20MintLimited *ERC20MintLimited) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8, extensionData []byte) []byte {
	enc, err := eRC20MintLimited.abi.Pack("", owner, forges, name, symbol, decimals, extensionData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20MintLimited *ERC20MintLimited) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20MintLimited.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20MintLimited *ERC20MintLimited) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20MintLimited.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20MintLimited.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20MintLimited *ERC20MintLimited) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20MintLimited.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackMANAGERROLE() []byte {
	enc, err := eRC20MintLimited.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20MintLimited *ERC20MintLimited) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20MintLimited.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20MintLimited.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("allowance", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20MintLimited *ERC20MintLimited) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20MintLimited.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackAvailableMintCapacity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x436ba626.
//
// Solidity: function availableMintCapacity() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) PackAvailableMintCapacity() []byte {
	enc, err := eRC20MintLimited.abi.Pack("availableMintCapacity")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAvailableMintCapacity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x436ba626.
//
// Solidity: function availableMintCapacity() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackAvailableMintCapacity(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("availableMintCapacity", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("balanceOf", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20MintLimited *ERC20MintLimited) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20MintLimited *ERC20MintLimited) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20MintLimited.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC20MintLimited *ERC20MintLimited) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MintLimited *ERC20MintLimited) PackDecimals() []byte {
	enc, err := eRC20MintLimited.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MintLimited *ERC20MintLimited) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20MintLimited.abi.Unpack("decimals", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackDefaultAdmin() []byte {
	enc, err := eRC20MintLimited.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20MintLimited.abi.Unpack("defaultAdmin", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackDefaultAdminDelay() []byte {
	enc, err := eRC20MintLimited.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("defaultAdminDelay", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20MintLimited.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackEip712Domain() []byte {
	enc, err := eRC20MintLimited.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20MintLimited *ERC20MintLimited) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20MintLimited.abi.Unpack("eip712Domain", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20MintLimited *ERC20MintLimited) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20MintLimited.abi.Unpack("forgeByIndex", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackForgeCount() []byte {
	enc, err := eRC20MintLimited.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("forgeCount", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackForges() []byte {
	enc, err := eRC20MintLimited.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20MintLimited *ERC20MintLimited) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20MintLimited.abi.Unpack("forges", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20MintLimited.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20MintLimited *ERC20MintLimited) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20MintLimited.abi.Unpack("getRoleAdmin", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20MintLimited *ERC20MintLimited) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20MintLimited.abi.Unpack("getRoleMember", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20MintLimited.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("getRoleMemberCount", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20MintLimited.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20MintLimited *ERC20MintLimited) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20MintLimited.abi.Unpack("getRoleMembers", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20MintLimited *ERC20MintLimited) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20MintLimited *ERC20MintLimited) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20MintLimited.abi.Unpack("hasRole", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20MintLimited *ERC20MintLimited) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20MintLimited.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMaxMintPerPeriod is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) PackMaxMintPerPeriod() []byte {
	enc, err := eRC20MintLimited.abi.Pack("maxMintPerPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackMaxMintPerPeriod(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("maxMintPerPeriod", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (eRC20MintLimited *ERC20MintLimited) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MintLimited *ERC20MintLimited) PackName() []byte {
	enc, err := eRC20MintLimited.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MintLimited *ERC20MintLimited) UnpackName(data []byte) (string, error) {
	out, err := eRC20MintLimited.abi.Unpack("name", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackNonces(owner common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("nonces", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackOwner() []byte {
	enc, err := eRC20MintLimited.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20MintLimited *ERC20MintLimited) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20MintLimited.abi.Unpack("owner", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20MintLimited.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20MintLimited *ERC20MintLimited) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20MintLimited.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20MintLimited.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20MintLimited *ERC20MintLimited) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20MintLimited.abi.Unpack("pendingDefaultAdminDelay", data)
	outstruct := new(PendingDefaultAdminDelayOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewDelay = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPeriodConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5068b4de.
//
// Solidity: function periodConfig() view returns(uint256, int256)
func (eRC20MintLimited *ERC20MintLimited) PackPeriodConfig() []byte {
	enc, err := eRC20MintLimited.abi.Pack("periodConfig")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5068b4de.
//
// Solidity: function periodConfig() view returns(uint256, int256)
func (eRC20MintLimited *ERC20MintLimited) UnpackPeriodConfig(data []byte) (PeriodConfigOutput, error) {
	out, err := eRC20MintLimited.abi.Unpack("periodConfig", data)
	outstruct := new(PeriodConfigOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPeriodStartTime is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbdf7acce.
//
// Solidity: function periodStartTime() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) PackPeriodStartTime() []byte {
	enc, err := eRC20MintLimited.abi.Pack("periodStartTime")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbdf7acce.
//
// Solidity: function periodStartTime() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackPeriodStartTime(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("periodStartTime", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20MintLimited *ERC20MintLimited) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20MintLimited.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (eRC20MintLimited *ERC20MintLimited) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20MintLimited *ERC20MintLimited) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MintLimited.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20MintLimited *ERC20MintLimited) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20MintLimited.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20MintLimited *ERC20MintLimited) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20MintLimited.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MintLimited *ERC20MintLimited) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20MintLimited.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MintLimited *ERC20MintLimited) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20MintLimited.abi.Unpack("supportsInterface", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackSymbol() []byte {
	enc, err := eRC20MintLimited.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20MintLimited *ERC20MintLimited) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20MintLimited.abi.Unpack("symbol", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackTotalSupply() []byte {
	enc, err := eRC20MintLimited.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20MintLimited *ERC20MintLimited) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20MintLimited.abi.Unpack("totalSupply", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20MintLimited *ERC20MintLimited) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20MintLimited.abi.Unpack("transfer", data)
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
func (eRC20MintLimited *ERC20MintLimited) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20MintLimited *ERC20MintLimited) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20MintLimited.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpdateMintLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe01d55c5.
//
// Solidity: function updateMintLimit(uint256 newLimit) returns()
func (eRC20MintLimited *ERC20MintLimited) PackUpdateMintLimit(newLimit *big.Int) []byte {
	enc, err := eRC20MintLimited.abi.Pack("updateMintLimit", newLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20MintLimitedApproval represents a Approval event raised by the ERC20MintLimited contract.
type ERC20MintLimitedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedApproval) ContractEventName() string {
	return ERC20MintLimitedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20MintLimited *ERC20MintLimited) UnpackApprovalEvent(log *types.Log) (*ERC20MintLimitedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedApproval)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20MintLimited contract.
type ERC20MintLimitedDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20MintLimitedDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20MintLimitedDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20MintLimited contract.
type ERC20MintLimitedDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20MintLimitedDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20MintLimitedDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20MintLimited contract.
type ERC20MintLimitedDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20MintLimitedDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20MintLimitedDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20MintLimited contract.
type ERC20MintLimitedDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20MintLimitedDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20MintLimited *ERC20MintLimited) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20MintLimitedDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20MintLimited contract.
type ERC20MintLimitedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedEIP712DomainChanged) ContractEventName() string {
	return ERC20MintLimitedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20MintLimited *ERC20MintLimited) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20MintLimitedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedForgeAdded represents a ForgeAdded event raised by the ERC20MintLimited contract.
type ERC20MintLimitedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedForgeAdded) ContractEventName() string {
	return ERC20MintLimitedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20MintLimited *ERC20MintLimited) UnpackForgeAddedEvent(log *types.Log) (*ERC20MintLimitedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedForgeRemoved represents a ForgeRemoved event raised by the ERC20MintLimited contract.
type ERC20MintLimitedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedForgeRemoved) ContractEventName() string {
	return ERC20MintLimitedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20MintLimited *ERC20MintLimited) UnpackForgeRemovedEvent(log *types.Log) (*ERC20MintLimitedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedMintLimitUpdated represents a MintLimitUpdated event raised by the ERC20MintLimited contract.
type ERC20MintLimitedMintLimitUpdated struct {
	OldLimits *big.Int
	NewLimits *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedMintLimitUpdatedEventName = "MintLimitUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedMintLimitUpdated) ContractEventName() string {
	return ERC20MintLimitedMintLimitUpdatedEventName
}

// UnpackMintLimitUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintLimitUpdated(uint256 oldLimits, uint256 newLimits)
func (eRC20MintLimited *ERC20MintLimited) UnpackMintLimitUpdatedEvent(log *types.Log) (*ERC20MintLimitedMintLimitUpdated, error) {
	event := "MintLimitUpdated"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedMintLimitUpdated)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedPeriodStarted represents a PeriodStarted event raised by the ERC20MintLimited contract.
type ERC20MintLimitedPeriodStarted struct {
	PeriodStart       *big.Int
	AvailableCapacity *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedPeriodStartedEventName = "PeriodStarted"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedPeriodStarted) ContractEventName() string {
	return ERC20MintLimitedPeriodStartedEventName
}

// UnpackPeriodStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity)
func (eRC20MintLimited *ERC20MintLimited) UnpackPeriodStartedEvent(log *types.Log) (*ERC20MintLimitedPeriodStarted, error) {
	event := "PeriodStarted"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedPeriodStarted)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20MintLimited contract.
type ERC20MintLimitedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedRoleAdminChanged) ContractEventName() string {
	return ERC20MintLimitedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20MintLimited *ERC20MintLimited) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20MintLimitedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedRoleGranted represents a RoleGranted event raised by the ERC20MintLimited contract.
type ERC20MintLimitedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedRoleGranted) ContractEventName() string {
	return ERC20MintLimitedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20MintLimited *ERC20MintLimited) UnpackRoleGrantedEvent(log *types.Log) (*ERC20MintLimitedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedRoleRevoked represents a RoleRevoked event raised by the ERC20MintLimited contract.
type ERC20MintLimitedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedRoleRevoked) ContractEventName() string {
	return ERC20MintLimitedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20MintLimited *ERC20MintLimited) UnpackRoleRevokedEvent(log *types.Log) (*ERC20MintLimitedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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

// ERC20MintLimitedTransfer represents a Transfer event raised by the ERC20MintLimited contract.
type ERC20MintLimitedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintLimitedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20MintLimitedTransfer) ContractEventName() string {
	return ERC20MintLimitedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20MintLimited *ERC20MintLimited) UnpackTransferEvent(log *types.Log) (*ERC20MintLimitedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20MintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintLimitedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20MintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MintLimited.abi.Events[event].Inputs {
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
func (eRC20MintLimited *ERC20MintLimited) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20PeriodMintLimitExceedsPeriodLimit"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20PeriodMintLimitExceedsPeriodLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20PeriodMintLimitInvalidLength"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20PeriodMintLimitInvalidLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC20PeriodMintLimitInvalidLimitData"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC20PeriodMintLimitInvalidLimitDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["PeriodManagerInvalidDuration"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackPeriodManagerInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["PeriodManagerInvalidOffset"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackPeriodManagerInvalidOffsetError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MintLimited.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20MintLimited.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20MintLimitedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20MintLimited contract.
type ERC20MintLimitedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20MintLimitedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20MintLimited *ERC20MintLimited) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20MintLimitedAccessControlBadConfirmation, error) {
	out := new(ERC20MintLimitedAccessControlBadConfirmation)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20MintLimited contract.
type ERC20MintLimitedAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20MintLimitedAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20MintLimited *ERC20MintLimited) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20MintLimitedAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20MintLimitedAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20MintLimited contract.
type ERC20MintLimitedAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20MintLimitedAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20MintLimited *ERC20MintLimited) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20MintLimitedAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20MintLimitedAccessControlEnforcedDefaultAdminRules)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20MintLimited contract.
type ERC20MintLimitedAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20MintLimitedAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20MintLimited *ERC20MintLimited) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20MintLimitedAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20MintLimitedAccessControlInvalidDefaultAdmin)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20MintLimited contract.
type ERC20MintLimitedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20MintLimitedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20MintLimited *ERC20MintLimited) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20MintLimitedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20MintLimitedAccessControlUnauthorizedAccount)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20MintLimited contract.
type ERC20MintLimitedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20MintLimitedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20MintLimited *ERC20MintLimited) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20MintLimitedECDSAInvalidSignature, error) {
	out := new(ERC20MintLimitedECDSAInvalidSignature)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20MintLimited contract.
type ERC20MintLimitedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20MintLimitedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20MintLimited *ERC20MintLimited) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20MintLimitedECDSAInvalidSignatureLength, error) {
	out := new(ERC20MintLimitedECDSAInvalidSignatureLength)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20MintLimited contract.
type ERC20MintLimitedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20MintLimitedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20MintLimited *ERC20MintLimited) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20MintLimitedECDSAInvalidSignatureS, error) {
	out := new(ERC20MintLimitedECDSAInvalidSignatureS)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20MintLimitedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20MintLimitedERC20InsufficientAllowance, error) {
	out := new(ERC20MintLimitedERC20InsufficientAllowance)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20MintLimitedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20MintLimitedERC20InsufficientBalance, error) {
	out := new(ERC20MintLimitedERC20InsufficientBalance)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20MintLimitedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20InvalidApproverError(raw []byte) (*ERC20MintLimitedERC20InvalidApprover, error) {
	out := new(ERC20MintLimitedERC20InvalidApprover)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20MintLimitedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20MintLimitedERC20InvalidReceiver, error) {
	out := new(ERC20MintLimitedERC20InvalidReceiver)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20MintLimitedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20InvalidSenderError(raw []byte) (*ERC20MintLimitedERC20InvalidSender, error) {
	out := new(ERC20MintLimitedERC20InvalidSender)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20MintLimitedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20MintLimitedERC20InvalidSpender, error) {
	out := new(ERC20MintLimitedERC20InvalidSpender)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20PeriodMintLimitExceedsPeriodLimit represents a ERC20PeriodMintLimit__ExceedsPeriodLimit error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20PeriodMintLimitExceedsPeriodLimit struct {
	Requested *big.Int
	Available *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available)
func ERC20MintLimitedERC20PeriodMintLimitExceedsPeriodLimitErrorID() common.Hash {
	return common.HexToHash("0x156d32c58749d9c638f5b3f8e359a7ee3d1aa2d1c7d31a6301ce10260bfd45b4")
}

// UnpackERC20PeriodMintLimitExceedsPeriodLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20PeriodMintLimitExceedsPeriodLimitError(raw []byte) (*ERC20MintLimitedERC20PeriodMintLimitExceedsPeriodLimit, error) {
	out := new(ERC20MintLimitedERC20PeriodMintLimitExceedsPeriodLimit)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitExceedsPeriodLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20PeriodMintLimitInvalidLength represents a ERC20PeriodMintLimit__InvalidLength error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20PeriodMintLimitInvalidLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLength()
func ERC20MintLimitedERC20PeriodMintLimitInvalidLengthErrorID() common.Hash {
	return common.HexToHash("0xf38583ae0b0acd19cabaa28c7b388b3f7e6260d4a4adbf38b03fb876a11643f0")
}

// UnpackERC20PeriodMintLimitInvalidLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLength()
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20PeriodMintLimitInvalidLengthError(raw []byte) (*ERC20MintLimitedERC20PeriodMintLimitInvalidLength, error) {
	out := new(ERC20MintLimitedERC20PeriodMintLimitInvalidLength)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitInvalidLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC20PeriodMintLimitInvalidLimitData represents a ERC20PeriodMintLimit__InvalidLimitData error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC20PeriodMintLimitInvalidLimitData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLimitData()
func ERC20MintLimitedERC20PeriodMintLimitInvalidLimitDataErrorID() common.Hash {
	return common.HexToHash("0x436572d08cf45a9b922f4897fbf93bb0e84720cc3edcbc7328ffcb075921f9a8")
}

// UnpackERC20PeriodMintLimitInvalidLimitDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLimitData()
func (eRC20MintLimited *ERC20MintLimited) UnpackERC20PeriodMintLimitInvalidLimitDataError(raw []byte) (*ERC20MintLimitedERC20PeriodMintLimitInvalidLimitData, error) {
	out := new(ERC20MintLimitedERC20PeriodMintLimitInvalidLimitData)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitInvalidLimitData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20MintLimitedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20MintLimitedERC2612ExpiredSignature, error) {
	out := new(ERC20MintLimitedERC2612ExpiredSignature)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20MintLimited contract.
type ERC20MintLimitedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20MintLimitedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20MintLimited *ERC20MintLimited) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20MintLimitedERC2612InvalidSigner, error) {
	out := new(ERC20MintLimitedERC2612InvalidSigner)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20MintLimited contract.
type ERC20MintLimitedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20MintLimitedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20MintLimited *ERC20MintLimited) UnpackInvalidAccountNonceError(raw []byte) (*ERC20MintLimitedInvalidAccountNonce, error) {
	out := new(ERC20MintLimitedInvalidAccountNonce)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedInvalidShortString represents a InvalidShortString error raised by the ERC20MintLimited contract.
type ERC20MintLimitedInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20MintLimitedInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20MintLimited *ERC20MintLimited) UnpackInvalidShortStringError(raw []byte) (*ERC20MintLimitedInvalidShortString, error) {
	out := new(ERC20MintLimitedInvalidShortString)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedPeriodManagerInvalidDuration represents a PeriodManager__InvalidDuration error raised by the ERC20MintLimited contract.
type ERC20MintLimitedPeriodManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidDuration()
func ERC20MintLimitedPeriodManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x28632f2e6184743eb990383edd4904c645090260b538d05e23af6d26cb637330")
}

// UnpackPeriodManagerInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidDuration()
func (eRC20MintLimited *ERC20MintLimited) UnpackPeriodManagerInvalidDurationError(raw []byte) (*ERC20MintLimitedPeriodManagerInvalidDuration, error) {
	out := new(ERC20MintLimitedPeriodManagerInvalidDuration)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedPeriodManagerInvalidOffset represents a PeriodManager__InvalidOffset error raised by the ERC20MintLimited contract.
type ERC20MintLimitedPeriodManagerInvalidOffset struct {
	Timestamp     *big.Int
	OffsetSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func ERC20MintLimitedPeriodManagerInvalidOffsetErrorID() common.Hash {
	return common.HexToHash("0xf41373a6f64e5d167ce7bcc91a19d341958fd006a060f85330cfc8aa087401d6")
}

// UnpackPeriodManagerInvalidOffsetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func (eRC20MintLimited *ERC20MintLimited) UnpackPeriodManagerInvalidOffsetError(raw []byte) (*ERC20MintLimitedPeriodManagerInvalidOffset, error) {
	out := new(ERC20MintLimitedPeriodManagerInvalidOffset)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidOffset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the ERC20MintLimited contract.
type ERC20MintLimitedSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func ERC20MintLimitedSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (eRC20MintLimited *ERC20MintLimited) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*ERC20MintLimitedSafeCastOverflowedIntDowncast, error) {
	out := new(ERC20MintLimitedSafeCastOverflowedIntDowncast)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20MintLimited contract.
type ERC20MintLimitedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20MintLimitedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20MintLimited *ERC20MintLimited) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20MintLimitedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20MintLimitedSafeCastOverflowedUintDowncast)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedStringTooLong represents a StringTooLong error raised by the ERC20MintLimited contract.
type ERC20MintLimitedStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20MintLimitedStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20MintLimited *ERC20MintLimited) UnpackStringTooLongError(raw []byte) (*ERC20MintLimitedStringTooLong, error) {
	out := new(ERC20MintLimitedStringTooLong)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20MintLimited contract.
type ERC20MintLimitedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20MintLimitedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20MintLimited *ERC20MintLimited) UnpackTokenBaseNullInputError(raw []byte) (*ERC20MintLimitedTokenBaseNullInput, error) {
	out := new(ERC20MintLimitedTokenBaseNullInput)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintLimitedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20MintLimited contract.
type ERC20MintLimitedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20MintLimitedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20MintLimited *ERC20MintLimited) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20MintLimitedTokenBaseOnlyForge, error) {
	out := new(ERC20MintLimitedTokenBaseOnlyForge)
	if err := eRC20MintLimited.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
