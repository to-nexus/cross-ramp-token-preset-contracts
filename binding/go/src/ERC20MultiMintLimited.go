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

// ERC20MultiMintLimitedMetaData contains all meta data concerning the ERC20MultiMintLimited contract.
var ERC20MultiMintLimitedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extensionData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacities\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriods\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodConfigs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartTimes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"updateMintLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"oldLimits\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodsMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodManager__InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offsetSeconds\",\"type\":\"int256\"}],\"name\":\"PeriodManager__InvalidOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20MultiMintLimited",
	Bin: "0x610180604052348015610010575f5ffd5b5060405161491c38038061491c83398101604081905261002f916109b4565b8086868686868280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b03160361009457604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100bd5f826104f5565b505050506100f17faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c836104f560201b60201c565b5080515f5b8181101561013057610128600484838151811061011557610115610a8e565b602002602001015161050960201b60201c565b6001016100f6565b5050505081600990816101439190610b26565b50600a6101508282610b26565b506101609150839050600b610581565b6101205261016f81600c610581565b61014052815160208084019190912060e052815190820120610100524660a0526101fb60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361022f5760405163e1dea2ef60e01b8152636e616d6560e01b600482015260240161008b565b81515f0361025b5760405163e1dea2ef60e01b8152651cde5b589bdb60d21b600482015260240161008b565b60ff1661016052505082515f925082915081906102819085016020908101908601610c3b565b80519295509093509150801580610299575083518114155b806102a5575082518114155b156102c357604051633fffd6c160e01b815260040160405180910390fd5b5f80805b8381101561041b575f5f8883815181106102e3576102e3610a8e565b60200260200101518784815181106102fd576102fd610a8e565b602002602001015191509150815f1480610315575080155b156103555760405163e1dea2ef60e01b81527f6c696d697473206f72206475726174696f6e7300000000000000000000000000600482015260240161008b565b84821115806103645750838111155b156103855760405163360b969560e21b81526004810184905260240161008b565b600f60405180604001604052806103a1856105b160201b60201c565b6001600160801b031681526020016103d78b87815181106103c4576103c4610a8e565b60200260200101516105e860201b60201c565b600f0b90528154600181810184555f938452602093849020835194909301516001600160801b03908116600160801b029416939093179101559194509250016102c7565b50600e83905583516104349060109060208701906107f2565b50826001600160401b0381111561044d5761044d61085c565b604051908082528060200260200182016040528015610476578160200160208202803683370190505b50805161048b916011916020909101906107f2565b50826001600160401b038111156104a4576104a461085c565b6040519080825280602002602001820160405280156104cd578160200160208202803683370190505b5080516104e2916012916020909101906107f2565b5050505050505050505050505050610d79565b5f610500838361061c565b90505b92915050565b6001600160a01b03811661053a5760405163e1dea2ef60e01b815264666f72676560d81b600482015260240161008b565b610544828261064f565b1561057d576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f60208351101561059c5761059583610663565b9050610503565b816105a78482610b26565b5060ff9050610503565b5f6001600160801b038211156105e4576040516306dfcc6560e41b8152608060048201526024810183905260440161008b565b5090565b80600f81900b81146106175760405163327269a760e01b8152608060048201526024810183905260440161008b565b919050565b5f8061062884846106a0565b90508015610500575f848152600360205260409020610647908461064f565b509392505050565b5f610500836001600160a01b038416610706565b5f5f829050601f8151111561068d578260405163305a27a960e01b815260040161008b9190610d21565b805161069882610d56565b179392505050565b5f826106fc575f6106b96002546001600160a01b031690565b6001600160a01b0316146106e057604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6105008383610752565b5f81815260018301602052604081205461074b57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610503565b505f610503565b5f828152602081815260408083206001600160a01b038516845290915281205460ff1661074b575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556107aa3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610503565b828054828255905f5260205f2090810192821561082b579160200282015b8281111561082b578251825591602001919060010190610810565b506105e49291505b808211156105e4575f8155600101610833565b80516001600160a01b0381168114610617575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156108985761089861085c565b604052919050565b5f6001600160401b038211156108b8576108b861085c565b5060051b60200190565b5f82601f8301126108d1575f5ffd5b81516108e46108df826108a0565b610870565b8082825260208201915060208360051b860101925085831115610905575f5ffd5b602085015b838110156109295761091b81610846565b83526020928301920161090a565b5095945050505050565b5f82601f830112610942575f5ffd5b8151602083015f806001600160401b038411156109615761096161085c565b50601f8301601f191660200161097681610870565b91505082815285838301111561098a575f5ffd5b8282602083015e5f92810160200192909252509392505050565b805160ff81168114610617575f5ffd5b5f5f5f5f5f5f60c087890312156109c9575f5ffd5b6109d287610846565b60208801519096506001600160401b038111156109ed575f5ffd5b6109f989828a016108c2565b604089015190965090506001600160401b03811115610a16575f5ffd5b610a2289828a01610933565b606089015190955090506001600160401b03811115610a3f575f5ffd5b610a4b89828a01610933565b935050610a5a608088016109a4565b60a08801519092506001600160401b03811115610a75575f5ffd5b610a8189828a01610933565b9150509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610ab657607f821691505b602082108103610ad457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610b2157805f5260205f20601f840160051c81016020851015610aff5750805b601f840160051c820191505b81811015610b1e575f8155600101610b0b565b50505b505050565b81516001600160401b03811115610b3f57610b3f61085c565b610b5381610b4d8454610aa2565b84610ada565b6020601f821160018114610b85575f8315610b6e5750848201515b5f19600385901b1c1916600184901b178455610b1e565b5f84815260208120601f198516915b82811015610bb45787850151825560209485019460019092019101610b94565b5084821015610bd157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82601f830112610bef575f5ffd5b8151610bfd6108df826108a0565b8082825260208201915060208360051b860101925085831115610c1e575f5ffd5b602085015b83811015610929578051835260209283019201610c23565b5f5f5f60608486031215610c4d575f5ffd5b83516001600160401b03811115610c62575f5ffd5b610c6e86828701610be0565b602086015190945090506001600160401b03811115610c8b575f5ffd5b8401601f81018613610c9b575f5ffd5b8051610ca96108df826108a0565b8082825260208201915060208360051b850101925088831115610cca575f5ffd5b6020840193505b82841015610cec578351825260209384019390910190610cd1565b6040880151909550925050506001600160401b03811115610d0b575f5ffd5b610d1786828701610be0565b9150509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610ad4575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051613b48610dd45f395f6103fb01525f611c2e01525f611c0101525f61171701525f6116ef01525f61164a01525f61167401525f61169e0152613b485ff3fe608060405234801561000f575f5ffd5b5060043610610303575f3560e01c806387f453531161019d578063ca15c873116100e8578063d505accf11610093578063dd62ed3e1161006e578063dd62ed3e146106df578063ec87621c146106f2578063fe2df3e814610719575f5ffd5b8063d505accf146106b1578063d547741f146106c4578063d602b9fd146106d7575f5ffd5b8063cdd760b9116100c3578063cdd760b914610647578063cefc14291461065d578063cf6eefb714610665575f5ffd5b8063ca15c87314610624578063cc8463c814610637578063ccf6d7241461063f575f5ffd5b8063a1eda53c11610148578063a40283d511610123578063a40283d514610601578063a9059cbb14610609578063c3ac4c2d1461061c575f5ffd5b8063a1eda53c146105c0578063a217fddf146105e7578063a3246ad3146105ee575f5ffd5b806391d148541161017857806391d148541461056257806395d89b41146105a55780639ca92df9146105ad575f5ffd5b806387f45353146105345780638da5cb5b146105475780639010d07c1461054f575f5ffd5b80633644e5151161025d578063649a5ec7116102085780637ecebe00116101e35780637ecebe00146104c757806384b0196e146104da57806384ef8ffc146104f5575f5ffd5b8063649a5ec71461048e57806370a08231146104a157806379cc6790146104b4575f5ffd5b806340c10f191161023857806340c10f19146104535780635c4e62c414610466578063634e93da1461047b575f5ffd5b80633644e5151461042557806336568abe1461042d57806337e653ea14610440575f5ffd5b80630aa6220b116102bd578063248a9ca311610298578063248a9ca3146103bf5780632f2ff15d146103e1578063313ce567146103f4575f5ffd5b80630aa6220b1461038c57806318160ddd1461039657806323b872dd146103ac575f5ffd5b8063022d63fb116102ed578063022d63fb1461034857806306fdde0314610364578063095ea7b314610379575f5ffd5b806286ced81461030757806301ffc9a714610325575b5f5ffd5b61030f61072c565b60405161031c91906133da565b60405180910390f35b6103386103333660046133ec565b6107d1565b604051901515815260200161031c565b620697805b60405165ffffffffffff909116815260200161031c565b61036c6109d3565b60405161031c9190613477565b6103386103873660046134b1565b6109e2565b6103946109f4565b005b61039e610a09565b60405190815260200161031c565b6103386103ba3660046134d9565b610a13565b61039e6103cd366004613513565b5f9081526020819052604090206001015490565b6103946103ef36600461352a565b610a27565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161031c565b61039e610a35565b61039461043b36600461352a565b610a3e565b61039461044e36600461359c565b610a48565b6103946104613660046134b1565b610c36565b61046e610c40565b60405161031c91906135db565b610394610489366004613633565b610c4c565b61039461049c36600461364c565b610c5f565b61039e6104af366004613633565b610c72565b6103946104c23660046134b1565b610c9c565b61039e6104d5366004613633565b610cce565b6104e2610cd8565b60405161031c9796959493929190613671565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161031c565b610338610542366004613633565b610d36565b61050f610d42565b61050f61055d366004613707565b610d62565b61033861057036600461352a565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61036c610d79565b61050f6105bb366004613513565b610d83565b6105c8610d8f565b6040805165ffffffffffff93841681529290911660208301520161031c565b61039e5f81565b61046e6105fc366004613513565b610e09565b61039e610e22565b6103386106173660046134b1565b610e2d565b61030f610e38565b61039e610632366004613513565b610fb6565b61034d610fcc565b61030f611069565b61064f6110bf565b60405161031c929190613727565b610394611219565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161031c565b6103946106bf366004613773565b611275565b6103946106d236600461352a565b61141e565b610394611428565b61039e6106ed3660046137e0565b61143a565b61039e7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b610394610727366004613808565b611473565b600e546060905f8167ffffffffffffffff81111561074c5761074c61385f565b604051908082528060200260200182016040528015610775578160200160208202803683370190505b5090505f5b828110156107ca576107a5600f82815481106107985761079861388c565b905f5260205f2001611504565b8282815181106107b7576107b761388c565b602090810291909101015260010161077a565b5092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061086357507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b806108af57507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b806108fb57507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b8061092657507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061097257507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b806109be57507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b806109cd57506109cd8261150f565b92915050565b60606109dd611519565b905090565b5f6109ed83836115a0565b9392505050565b5f6109fe816115b7565b610a066115c1565b50565b5f6109dd60085490565b5f610a1f8484846115cd565b949350505050565b610a3182826115f0565b5050565b5f6109dd611631565b610a318282611767565b5f610a52816115b7565b815f819003610ab4576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d697473000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b600e548114610aef576040517f3fffd6c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805b82811015610be5575f868683818110610b0d57610b0d61388c565b905060200201359050805f03610b71576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69747300000000000000000000000000000000000000000000006004820152602401610aab565b828111158015610ba157507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114155b15610bdb576040517fd82e5a5400000000000000000000000000000000000000000000000000000000815260048101839052602401610aab565b9150600101610af2565b507f89be4a9eb7586334d719b1c87a0fd1e8306235335bf257d957c53aa7b22fcb7460108686604051610c1a939291906138b9565b60405180910390a1610c2e6010868661333f565b505050505050565b610a31828261186c565b60606109dd6004611a6c565b5f610c56816115b7565b610a3182611a78565b5f610c69816115b7565b610a3182611af7565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600660205260408120546109cd565b73ffffffffffffffffffffffffffffffffffffffff82163314610cc457610cc4823383611b66565b610a318282611b76565b5f6109cd82611bd0565b5f6060805f5f5f6060610ce9611bfa565b610cf1611c27565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f6109cd600483611c54565b5f6109dd60025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206109ed9083611c82565b60606109dd611c8d565b5f6109cd600483611c82565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610dd157504265ffffffffffff821610155b610ddc575f5f610e01565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f8181526003602052604090206060906109cd90611a6c565b5f6109dd6004611c9c565b5f6109ed8383611ca5565b60605f610e4361072c565b90505f6011805480602002602001604051908101604052809291908181526020018280548015610e9057602002820191905f5260205f20905b815481526020019060010190808311610e7c575b505050505090505f600e5490505f8167ffffffffffffffff811115610eb757610eb761385f565b604051908082528060200260200182016040528015610ee0578160200160208202803683370190505b5090505f5b82811015610fad575f848281518110610f0057610f0061388c565b60200260200101519050858281518110610f1c57610f1c61388c565b60200260200101518103610f695760128281548110610f3d57610f3d61388c565b905f5260205f200154838381518110610f5857610f5861388c565b602002602001018181525050610fa4565b60108281548110610f7c57610f7c61388c565b905f5260205f200154838381518110610f9757610f9761388c565b6020026020010181815250505b50600101610ee5565b50949350505050565b5f8181526003602052604081206109cd90611c9c565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801515801561100d57504265ffffffffffff8216105b61103f576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16611063565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b606060108054806020026020016040519081016040528092919081815260200182805480156110b557602002820191905f5260205f20905b8154815260200190600101908083116110a1575b5050505050905090565b600e5460609081905f8167ffffffffffffffff8111156110e1576110e161385f565b60405190808252806020026020018201604052801561110a578160200160208202803683370190505b5090505f8267ffffffffffffffff8111156111275761112761385f565b604051908082528060200260200182016040528015611150578160200160208202803683370190505b5090505f5b8381101561120e57600f81815481106111705761117061388c565b5f9182526020909120015483516fffffffffffffffffffffffffffffffff909116908490839081106111a4576111a461388c565b602002602001018181525050600f81815481106111c3576111c361388c565b5f918252602090912001548251700100000000000000000000000000000000909104600f0b908390839081106111fb576111fb61388c565b6020908102919091010152600101611155565b509094909350915050565b60015473ffffffffffffffffffffffffffffffffffffffff1633811461126d576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610aab565b610a06611cb2565b834211156112b2576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610aab565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861130a8c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61137182611da3565b90505f61138082878787611dea565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611407576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610aab565b6114128a8a8a611e16565b50505050505050505050565b610a318282611e23565b5f611432816115b7565b610a06611e64565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081526007602090815260408083209385168352929052908120546109ed565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61149d816115b7565b613384826114ad57611e6e6114b1565b611ec15b9050835f5b818110156114fb576114f360048888848181106114d5576114d561388c565b90506020020160208101906114ea9190613633565b8563ffffffff16565b6001016114b6565b50505050505050565b5f6109cd8242611f83565b5f6109cd826120ff565b60606009805461152890613948565b80601f016020809104026020016040519081016040528092919081815260200182805461155490613948565b80156110b55780601f10611576576101008083540402835291602001916110b5565b820191905f5260205f20905b81548152906001019060200180831161158257509395945050505050565b5f336115ad818585611e16565b5060019392505050565b610a068133612109565b6115cb5f5f61218e565b565b5f336115da858285611b66565b6115e58585856122e7565b506001949350505050565b81611627576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a318282612390565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561169657507f000000000000000000000000000000000000000000000000000000000000000046145b156116c057507f000000000000000000000000000000000000000000000000000000000000000090565b6109dd604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b8115801561178f575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156118625760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806117e3575065ffffffffffff8116155b806117f657504265ffffffffffff821610155b15611837576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610aab565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b610a3182826123b4565b5f61187561072c565b600e549091505f5b81811015611a5b575f5f5f6012848154811061189b5761189b61388c565b905f5260205f200154601185815481106118b7576118b761388c565b905f5260205f2001548786815181106118d2576118d261388c565b6020026020010151925092509250808214611965575f601085815481106118fb576118fb61388c565b905f5260205f2001549050816011868154811061191a5761191a61388c565b905f5260205f200181905550809350817ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d18260405161195b91815260200190565b60405180910390a2505b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83036119b05782601285815481106119a0576119a061388c565b5f91825260209091200155611a4d565b86831015611a2b57600f84815481106119cb576119cb61388c565b5f918252602090912001546040517faace2cbe0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201526024810188905260448101849052606401610aab565b86830360128581548110611a4157611a4161388c565b5f918252602090912001555b50505080600101905061187d565b50611a66848461240d565b50505050565b60605f6109ed83612458565b5f611a81610fcc565b611a8a426124b1565b611a9491906139c6565b9050611aa08282612500565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f611b018261259b565b611b0a426124b1565b611b1491906139c6565b9050611b20828261218e565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b611b718383836125e2565b505050565b73ffffffffffffffffffffffffffffffffffffffff8216611bc5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610aab565b610a31825f83612685565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d60205260408120546109cd565b60606109dd7f0000000000000000000000000000000000000000000000000000000000000000600b612690565b60606109dd7f0000000000000000000000000000000000000000000000000000000000000000600c612690565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156109ed565b5f6109ed8383612732565b6060600a805461152890613948565b5f6109cd825490565b5f336115ad8185856122e7565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580611d0257504265ffffffffffff821610155b15611d43576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610aab565b611d6b5f611d6660025473ffffffffffffffffffffffffffffffffffffffff1690565b612758565b50611d765f83612763565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f6109cd611daf611631565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f611dfa8888888861276e565b925092509250611e0a8282612861565b50909695505050505050565b611b718383836001612964565b81611e5a576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a318282612970565b6115cb5f5f612500565b611e788282612994565b15610a315760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611f30576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610aab565b611f3a82826129b5565b15610a315760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b81545f906fffffffffffffffffffffffffffffffff168103611fd1576040517f28632f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82546fffffffffffffffffffffffffffffffff168281611ff357611ff36139e4565b8454919006909203917001000000000000000000000000000000009004600f0b5f036120205750806109cd565b82545f700100000000000000000000000000000000909104600f0b131561206b578254612064907001000000000000000000000000000000009004600f0b83613a11565b90506109cd565b82545f9061208f907001000000000000000000000000000000009004600f0b613a24565b9050808310156120ed5783546040517ff41373a600000000000000000000000000000000000000000000000000000000815260048101859052700100000000000000000000000000000000909104600f0b6024820152604401610aab565b6120f78184613a5a565b9150506109cd565b5f6109cd826129d6565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610a31576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610aab565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015612262574265ffffffffffff82161015612239576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055612262565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff8316612336576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610aab565b73ffffffffffffffffffffffffffffffffffffffff8216612385576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610aab565b611b71838383612685565b5f828152602081905260409020600101546123aa816115b7565b611a668383612763565b73ffffffffffffffffffffffffffffffffffffffff81163314612403576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b718282612758565b61241633610d36565b61244e576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610aab565b610a318282612a2b565b6060815f018054806020026020016040519081016040528092919081815260200182805480156124a557602002820191905f5260205f20905b815481526020019060010190808311612491575b50505050509050919050565b5f65ffffffffffff8211156124fc576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610aab565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611b71576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f6125a5610fcc565b90508065ffffffffffff168365ffffffffffff16116125cd576125c88382613a6d565b6109ed565b6109ed65ffffffffffff841662069780612a85565b5f6125ed848461143a565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015611a665781811015612677576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610aab565b611a6684848484035f612964565b611b71838383612a94565b606060ff83146126a35761206483612c3b565b8180546126af90613948565b80601f01602080910402602001604051908101604052809291908181526020018280546126db90613948565b80156127265780601f106126fd57610100808354040283529160200191612726565b820191905f5260205f20905b81548152906001019060200180831161270957829003601f168201915b505050505090506109cd565b5f825f0182815481106127475761274761388c565b905f5260205f200154905092915050565b5f6109ed8383612c78565b5f6109ed8383612cab565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156127a757505f91506003905082612857565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156127f8573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661284e57505f925060019150829050612857565b92505f91508190505b9450945094915050565b5f82600381111561287457612874613a8b565b0361287d575050565b600182600381111561289157612891613a8b565b036128c8576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156128dc576128dc613a8b565b03612916576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610aab565b600382600381111561292a5761292a613a8b565b03610a31576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610aab565b611a6684848484612cd6565b5f8281526020819052604090206001015461298a816115b7565b611a668383612758565b5f6109ed8373ffffffffffffffffffffffffffffffffffffffff8416612e1b565b5f6109ed8373ffffffffffffffffffffffffffffffffffffffff8416612efe565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806109cd57506109cd82612f4a565b73ffffffffffffffffffffffffffffffffffffffff8216612a7a576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610aab565b610a315f8383612685565b5f8282188284100282186109ed565b73ffffffffffffffffffffffffffffffffffffffff8316612acb578060085f828254612ac09190613a11565b90915550612b7b9050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526006602052604090205481811015612b50576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610aab565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216612ba457600880548290039055612bcf565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051612c2e91815260200190565b60405180910390a3505050565b60605f612c4783612f9f565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f612c848484612fdf565b905080156109ed575f848152600360205260409020612ca39084612994565b509392505050565b5f5f612cb78484613040565b905080156109ed575f848152600360205260409020612ca390846129b5565b73ffffffffffffffffffffffffffffffffffffffff8416612d25576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610aab565b73ffffffffffffffffffffffffffffffffffffffff8316612d74576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610aab565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526007602090815260408083209387168352929052208290558015611a66578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051612e0d91815260200190565b60405180910390a350505050565b5f8181526001830160205260408120548015612ef5575f612e3d600183613a5a565b85549091505f90612e5090600190613a5a565b9050808214612eaf575f865f018281548110612e6e57612e6e61388c565b905f5260205f200154905080875f018481548110612e8e57612e8e61388c565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612ec057612ec0613ab8565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506109cd565b5f9150506109cd565b5f818152600183016020526040812054612f4357508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556109cd565b505f6109cd565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806109cd57506109cd826130fe565b5f60ff8216601f8111156109cd576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82158015613008575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561303657600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6109ed8383613194565b5f826130f4575f61306660025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16146130b3576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6109ed838361324d565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806109cd57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146109cd565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615612f43575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016109cd565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16612f43575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556132dd3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016109cd565b828054828255905f5260205f20908101928215613378579160200282015b8281111561337857823582559160200191906001019061335d565b506124fc92915061338c565b6115cb613ae5565b5b808211156124fc575f815560010161338d565b5f8151808452602084019350602083015f5b828110156133d05781518652602095860195909101906001016133b2565b5093949350505050565b602081525f6109ed60208301846133a0565b5f602082840312156133fc575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146109ed575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109ed602083018461342b565b803573ffffffffffffffffffffffffffffffffffffffff811681146134ac575f5ffd5b919050565b5f5f604083850312156134c2575f5ffd5b6134cb83613489565b946020939093013593505050565b5f5f5f606084860312156134eb575f5ffd5b6134f484613489565b925061350260208501613489565b929592945050506040919091013590565b5f60208284031215613523575f5ffd5b5035919050565b5f5f6040838503121561353b575f5ffd5b8235915061354b60208401613489565b90509250929050565b5f5f83601f840112613564575f5ffd5b50813567ffffffffffffffff81111561357b575f5ffd5b6020830191508360208260051b8501011115613595575f5ffd5b9250929050565b5f5f602083850312156135ad575f5ffd5b823567ffffffffffffffff8111156135c3575f5ffd5b6135cf85828601613554565b90969095509350505050565b602080825282518282018190525f918401906040840190835b8181101561362857835173ffffffffffffffffffffffffffffffffffffffff168352602093840193909201916001016135f4565b509095945050505050565b5f60208284031215613643575f5ffd5b6109ed82613489565b5f6020828403121561365c575f5ffd5b813565ffffffffffff811681146109ed575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6136ab60e083018961342b565b82810360408401526136bd818961342b565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c08401526136f981856133a0565b9a9950505050505050505050565b5f5f60408385031215613718575f5ffd5b50508035926020909101359150565b604081525f61373960408301856133a0565b82810360208401528084518083526020830191506020860192505f5b81811015611e0a578351835260209384019390920191600101613755565b5f5f5f5f5f5f5f60e0888a031215613789575f5ffd5b61379288613489565b96506137a060208901613489565b95506040880135945060608801359350608088013560ff811681146137c3575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156137f1575f5ffd5b6137fa83613489565b915061354b60208401613489565b5f5f5f6040848603121561381a575f5ffd5b833567ffffffffffffffff811115613830575f5ffd5b61383c86828701613554565b90945092505060208401358015158114613854575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b604080825284549082018190525f8581526020812090916060840190835b818110156138f55783548352600193840193602090930192016138d7565b505083810360208501528481527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85111561392e575f5ffd5b8460051b9150818660208301370160200195945050505050565b600181811c9082168061395c57607f821691505b602082108103613993577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff81811683821601908111156109cd576109cd613999565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b808201808211156109cd576109cd613999565b5f7f80000000000000000000000000000000000000000000000000000000000000008203613a5457613a54613999565b505f0390565b818103818111156109cd576109cd613999565b65ffffffffffff82811682821603908111156109cd576109cd613999565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220c7ba5fee38a6b7d2ab4387ced37d5d91ed867e229d5bd298d649ad0657cc46ed64736f6c634300081c0033",
}

// ERC20MultiMintLimited is an auto generated Go binding around an Ethereum contract.
type ERC20MultiMintLimited struct {
	abi abi.ABI
}

// NewERC20MultiMintLimited creates a new instance of ERC20MultiMintLimited.
func NewERC20MultiMintLimited() *ERC20MultiMintLimited {
	parsed, err := ERC20MultiMintLimitedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20MultiMintLimited{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20MultiMintLimited) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals, bytes extensionData) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8, extensionData []byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("", owner, forges, name, symbol, decimals, extensionData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMANAGERROLE() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("allowance", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackAvailableMintCapacities is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3ac4c2d.
//
// Solidity: function availableMintCapacities() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackAvailableMintCapacities() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("availableMintCapacities")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAvailableMintCapacities is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc3ac4c2d.
//
// Solidity: function availableMintCapacities() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAvailableMintCapacities(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("availableMintCapacities", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("balanceOf", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDecimals() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("decimals", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDefaultAdmin() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("defaultAdmin", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDefaultAdminDelay() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("defaultAdminDelay", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackEip712Domain() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("eip712Domain", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("forgeByIndex", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackForgeCount() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("forgeCount", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackForges() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("forges", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("getRoleAdmin", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("getRoleMember", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("getRoleMemberCount", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("getRoleMembers", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("hasRole", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMaxMintPerPeriods is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xccf6d724.
//
// Solidity: function maxMintPerPeriods() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMaxMintPerPeriods() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("maxMintPerPeriods")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriods is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xccf6d724.
//
// Solidity: function maxMintPerPeriods() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMaxMintPerPeriods(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("maxMintPerPeriods", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackName() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackName(data []byte) (string, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("name", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackNonces(owner common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("nonces", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackOwner() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("owner", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("pendingDefaultAdminDelay", data)
	outstruct := new(PendingDefaultAdminDelayOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewDelay = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackPeriodConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdd760b9.
//
// Solidity: function periodConfigs() view returns(uint256[], int256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPeriodConfigs() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("periodConfigs")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodConfigs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcdd760b9.
//
// Solidity: function periodConfigs() view returns(uint256[], int256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodConfigs(data []byte) (PeriodConfigsOutput, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("periodConfigs", data)
	outstruct := new(PeriodConfigsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackPeriodStartTimes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0086ced8.
//
// Solidity: function periodStartTimes() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPeriodStartTimes() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("periodStartTimes")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartTimes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0086ced8.
//
// Solidity: function periodStartTimes() view returns(uint256[])
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodStartTimes(data []byte) ([]*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("periodStartTimes", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("supportsInterface", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackSymbol() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("symbol", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackTotalSupply() []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("totalSupply", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("transfer", data)
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20MultiMintLimited.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackUpdateMintLimits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37e653ea.
//
// Solidity: function updateMintLimits(uint256[] newLimits) returns()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) PackUpdateMintLimits(newLimits []*big.Int) []byte {
	enc, err := eRC20MultiMintLimited.abi.Pack("updateMintLimits", newLimits)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20MultiMintLimitedApproval represents a Approval event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedApproval) ContractEventName() string {
	return ERC20MultiMintLimitedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackApprovalEvent(log *types.Log) (*ERC20MultiMintLimitedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedApproval)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20MultiMintLimitedDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20MultiMintLimitedDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20MultiMintLimitedDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20MultiMintLimitedDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20MultiMintLimitedDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20MultiMintLimitedDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20MultiMintLimitedDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20MultiMintLimitedDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedEIP712DomainChanged) ContractEventName() string {
	return ERC20MultiMintLimitedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20MultiMintLimitedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedForgeAdded represents a ForgeAdded event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedForgeAdded) ContractEventName() string {
	return ERC20MultiMintLimitedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeAddedEvent(log *types.Log) (*ERC20MultiMintLimitedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedForgeRemoved represents a ForgeRemoved event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedForgeRemoved) ContractEventName() string {
	return ERC20MultiMintLimitedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackForgeRemovedEvent(log *types.Log) (*ERC20MultiMintLimitedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedMintLimitUpdated represents a MintLimitUpdated event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedMintLimitUpdated struct {
	OldLimits []*big.Int
	NewLimits []*big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedMintLimitUpdatedEventName = "MintLimitUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedMintLimitUpdated) ContractEventName() string {
	return ERC20MultiMintLimitedMintLimitUpdatedEventName
}

// UnpackMintLimitUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackMintLimitUpdatedEvent(log *types.Log) (*ERC20MultiMintLimitedMintLimitUpdated, error) {
	event := "MintLimitUpdated"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedMintLimitUpdated)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedPeriodStarted represents a PeriodStarted event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedPeriodStarted struct {
	PeriodStart       *big.Int
	AvailableCapacity *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedPeriodStartedEventName = "PeriodStarted"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedPeriodStarted) ContractEventName() string {
	return ERC20MultiMintLimitedPeriodStartedEventName
}

// UnpackPeriodStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodStartedEvent(log *types.Log) (*ERC20MultiMintLimitedPeriodStarted, error) {
	event := "PeriodStarted"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedPeriodStarted)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedRoleAdminChanged) ContractEventName() string {
	return ERC20MultiMintLimitedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20MultiMintLimitedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedRoleGranted represents a RoleGranted event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedRoleGranted) ContractEventName() string {
	return ERC20MultiMintLimitedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRoleGrantedEvent(log *types.Log) (*ERC20MultiMintLimitedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedRoleRevoked represents a RoleRevoked event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedRoleRevoked) ContractEventName() string {
	return ERC20MultiMintLimitedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackRoleRevokedEvent(log *types.Log) (*ERC20MultiMintLimitedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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

// ERC20MultiMintLimitedTransfer represents a Transfer event raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MultiMintLimitedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20MultiMintLimitedTransfer) ContractEventName() string {
	return ERC20MultiMintLimitedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTransferEvent(log *types.Log) (*ERC20MultiMintLimitedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20MultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MultiMintLimitedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20MultiMintLimited.abi.Events[event].Inputs {
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
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20PeriodsMintLimitExceedsPeriodLimit"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20PeriodsMintLimitExceedsPeriodLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20PeriodsMintLimitInvalidLength"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20PeriodsMintLimitInvalidLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC20PeriodsMintLimitInvalidLimitData"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC20PeriodsMintLimitInvalidLimitDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["PeriodManagerInvalidDuration"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackPeriodManagerInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["PeriodManagerInvalidOffset"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackPeriodManagerInvalidOffsetError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20MultiMintLimited.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20MultiMintLimited.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20MultiMintLimitedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20MultiMintLimitedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20MultiMintLimitedAccessControlBadConfirmation, error) {
	out := new(ERC20MultiMintLimitedAccessControlBadConfirmation)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20MultiMintLimitedAccessControlEnforcedDefaultAdminRules)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20MultiMintLimitedAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20MultiMintLimitedAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20MultiMintLimitedAccessControlInvalidDefaultAdmin)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20MultiMintLimitedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20MultiMintLimitedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20MultiMintLimitedAccessControlUnauthorizedAccount)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20MultiMintLimitedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20MultiMintLimitedECDSAInvalidSignature, error) {
	out := new(ERC20MultiMintLimitedECDSAInvalidSignature)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20MultiMintLimitedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20MultiMintLimitedECDSAInvalidSignatureLength, error) {
	out := new(ERC20MultiMintLimitedECDSAInvalidSignatureLength)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20MultiMintLimitedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20MultiMintLimitedECDSAInvalidSignatureS, error) {
	out := new(ERC20MultiMintLimitedECDSAInvalidSignatureS)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20MultiMintLimitedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20MultiMintLimitedERC20InsufficientAllowance, error) {
	out := new(ERC20MultiMintLimitedERC20InsufficientAllowance)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20MultiMintLimitedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20MultiMintLimitedERC20InsufficientBalance, error) {
	out := new(ERC20MultiMintLimitedERC20InsufficientBalance)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20MultiMintLimitedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidApproverError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidApprover, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidApprover)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20MultiMintLimitedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidReceiver, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidReceiver)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20MultiMintLimitedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidSenderError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidSender, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidSender)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20MultiMintLimitedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20MultiMintLimitedERC20InvalidSpender, error) {
	out := new(ERC20MultiMintLimitedERC20InvalidSpender)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit represents a ERC20PeriodsMintLimit__ExceedsPeriodLimit error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit struct {
	Period    *big.Int
	Requested *big.Int
	Available *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available)
func ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimitErrorID() common.Hash {
	return common.HexToHash("0xaace2cbe23069663ea57043f9018ae5639c94c4b7537987827b85a4a671f9687")
}

// UnpackERC20PeriodsMintLimitExceedsPeriodLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20PeriodsMintLimitExceedsPeriodLimitError(raw []byte) (*ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit, error) {
	out := new(ERC20MultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitExceedsPeriodLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength represents a ERC20PeriodsMintLimit__InvalidLength error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLength()
func ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLengthErrorID() common.Hash {
	return common.HexToHash("0x3fffd6c1aac13cb22e11380a50735d4cdde9620742e341a28fbb7c0113c01dbe")
}

// UnpackERC20PeriodsMintLimitInvalidLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLength()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20PeriodsMintLimitInvalidLengthError(raw []byte) (*ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength, error) {
	out := new(ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLength)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitInvalidLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData represents a ERC20PeriodsMintLimit__InvalidLimitData error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index)
func ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitDataErrorID() common.Hash {
	return common.HexToHash("0xd82e5a5493b676ee930d3b301ed5da6339364f86d1828cf9de0b2e2baa8068b3")
}

// UnpackERC20PeriodsMintLimitInvalidLimitDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC20PeriodsMintLimitInvalidLimitDataError(raw []byte) (*ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData, error) {
	out := new(ERC20MultiMintLimitedERC20PeriodsMintLimitInvalidLimitData)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitInvalidLimitData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20MultiMintLimitedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20MultiMintLimitedERC2612ExpiredSignature, error) {
	out := new(ERC20MultiMintLimitedERC2612ExpiredSignature)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20MultiMintLimitedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20MultiMintLimitedERC2612InvalidSigner, error) {
	out := new(ERC20MultiMintLimitedERC2612InvalidSigner)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20MultiMintLimitedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackInvalidAccountNonceError(raw []byte) (*ERC20MultiMintLimitedInvalidAccountNonce, error) {
	out := new(ERC20MultiMintLimitedInvalidAccountNonce)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedInvalidShortString represents a InvalidShortString error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20MultiMintLimitedInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackInvalidShortStringError(raw []byte) (*ERC20MultiMintLimitedInvalidShortString, error) {
	out := new(ERC20MultiMintLimitedInvalidShortString)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedPeriodManagerInvalidDuration represents a PeriodManager__InvalidDuration error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedPeriodManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidDuration()
func ERC20MultiMintLimitedPeriodManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x28632f2e6184743eb990383edd4904c645090260b538d05e23af6d26cb637330")
}

// UnpackPeriodManagerInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidDuration()
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodManagerInvalidDurationError(raw []byte) (*ERC20MultiMintLimitedPeriodManagerInvalidDuration, error) {
	out := new(ERC20MultiMintLimitedPeriodManagerInvalidDuration)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedPeriodManagerInvalidOffset represents a PeriodManager__InvalidOffset error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedPeriodManagerInvalidOffset struct {
	Timestamp     *big.Int
	OffsetSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func ERC20MultiMintLimitedPeriodManagerInvalidOffsetErrorID() common.Hash {
	return common.HexToHash("0xf41373a6f64e5d167ce7bcc91a19d341958fd006a060f85330cfc8aa087401d6")
}

// UnpackPeriodManagerInvalidOffsetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackPeriodManagerInvalidOffsetError(raw []byte) (*ERC20MultiMintLimitedPeriodManagerInvalidOffset, error) {
	out := new(ERC20MultiMintLimitedPeriodManagerInvalidOffset)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidOffset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func ERC20MultiMintLimitedSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*ERC20MultiMintLimitedSafeCastOverflowedIntDowncast, error) {
	out := new(ERC20MultiMintLimitedSafeCastOverflowedIntDowncast)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20MultiMintLimitedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20MultiMintLimitedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20MultiMintLimitedSafeCastOverflowedUintDowncast)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedStringTooLong represents a StringTooLong error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20MultiMintLimitedStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackStringTooLongError(raw []byte) (*ERC20MultiMintLimitedStringTooLong, error) {
	out := new(ERC20MultiMintLimitedStringTooLong)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20MultiMintLimitedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTokenBaseNullInputError(raw []byte) (*ERC20MultiMintLimitedTokenBaseNullInput, error) {
	out := new(ERC20MultiMintLimitedTokenBaseNullInput)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MultiMintLimitedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20MultiMintLimited contract.
type ERC20MultiMintLimitedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20MultiMintLimitedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20MultiMintLimited *ERC20MultiMintLimited) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20MultiMintLimitedTokenBaseOnlyForge, error) {
	out := new(ERC20MultiMintLimitedTokenBaseOnlyForge)
	if err := eRC20MultiMintLimited.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
