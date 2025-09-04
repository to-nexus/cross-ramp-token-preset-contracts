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

// ERC20CappedMintLimitedMetaData contains all meta data concerning the ERC20CappedMintLimited contract.
var ERC20CappedMintLimitedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes[2]\",\"name\":\"extensionData\",\"type\":\"bytes[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"updateMintLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLimits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimits\",\"type\":\"uint256\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodManager__InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offsetSeconds\",\"type\":\"int256\"}],\"name\":\"PeriodManager__InvalidOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20CappedMintLimited",
	Bin: "0x6101a0604052348015610010575f5ffd5b5060405161406a38038061406a83398101604081905261002f9161088b565b8060016020020151815f602002015187878787878280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b0316036100a257604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100cb5f82610380565b505050506100ff7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361038060201b60201c565b5080515f5b8181101561013e57610136600484838151811061012357610123610965565b602002602001015161039460201b60201c565b600101610104565b50505050816009908161015191906109fd565b50600a61015e82826109fd565b5061016e9150839050600b61040c565b6101205261017d81600c61040c565b61014052815160208084019190912060e052815190820120610100524660a05261020960e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361023d5760405163e1dea2ef60e01b8152636e616d6560e01b6004820152602401610099565b81515f036102695760405163e1dea2ef60e01b8152651cde5b589bdb60d21b6004820152602401610099565b60ff1661016052505082515f925061028b915083016020908101908401610ab7565b9050805f036102b55760405163e1dea2ef60e01b81526206361760ec1b6004820152602401610099565b610180525080515f90819081906102d59085016020908101908601610ace565b925092509250825f14806102e7575080155b156103195760405163e1dea2ef60e01b81526e1b1a5b5a5d081bdc881c195c9a5bd9608a1b6004820152602401610099565b60405180604001604052806103338561043c60201b60201c565b6001600160801b0316815260200161034a84610473565b600f0b905280516020909101516001600160801b03908116600160801b02911617600e5560105550610b51975050505050505050565b5f61038b83836104a7565b90505b92915050565b6001600160a01b0381166103c55760405163e1dea2ef60e01b815264666f72676560d81b6004820152602401610099565b6103cf82826104da565b15610408576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f60208351101561042757610420836104ee565b905061038e565b8161043284826109fd565b5060ff905061038e565b5f6001600160801b0382111561046f576040516306dfcc6560e41b81526080600482015260248101839052604401610099565b5090565b80600f81900b81146104a25760405163327269a760e01b81526080600482015260248101839052604401610099565b919050565b5f806104b3848461052b565b9050801561038b575f8481526003602052604090206104d290846104da565b509392505050565b5f61038b836001600160a01b038416610591565b5f5f829050601f81511115610518578260405163305a27a960e01b81526004016100999190610af9565b805161052382610b2e565b179392505050565b5f82610587575f6105446002546001600160a01b031690565b6001600160a01b03161461056b57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b61038b83836105dd565b5f8181526001830160205260408120546105d657508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561038e565b505f61038e565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166105d6575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556106353390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161038e565b80516001600160a01b03811681146104a2575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156106c9576106c9610693565b60405290565b604051601f8201601f191681016001600160401b03811182821017156106f7576106f7610693565b604052919050565b5f82601f83011261070e575f5ffd5b81516001600160401b0381111561072757610727610693565b8060051b610737602082016106cf565b91825260208185018101929081019086841115610752575f5ffd5b6020860192505b8383101561077b5761076a8361067d565b825260209283019290910190610759565b9695505050505050565b5f806001600160401b0384111561079e5761079e610693565b50601f8301601f19166020016107b3816106cf565b9150508281528383830111156107c7575f5ffd5b8282602083015e5f602084830101529392505050565b5f82601f8301126107ec575f5ffd5b61038b83835160208501610785565b805160ff811681146104a2575f5ffd5b5f82601f83011261081a575f5ffd5b6108226106a7565b806040840185811115610833575f5ffd5b845b818110156108805780516001600160401b03811115610852575f5ffd5b8601601f81018813610862575f5ffd5b61087188825160208401610785565b85525060209384019301610835565b509095945050505050565b5f5f5f5f5f5f60c087890312156108a0575f5ffd5b6108a98761067d565b60208801519096506001600160401b038111156108c4575f5ffd5b6108d089828a016106ff565b604089015190965090506001600160401b038111156108ed575f5ffd5b6108f989828a016107dd565b606089015190955090506001600160401b03811115610916575f5ffd5b61092289828a016107dd565b935050610931608088016107fb565b60a08801519092506001600160401b0381111561094c575f5ffd5b61095889828a0161080b565b9150509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061098d57607f821691505b6020821081036109ab57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156109f857805f5260205f20601f840160051c810160208510156109d65750805b601f840160051c820191505b818110156109f5575f81556001016109e2565b50505b505050565b81516001600160401b03811115610a1657610a16610693565b610a2a81610a248454610979565b846109b1565b6020601f821160018114610a5c575f8315610a455750848201515b5f19600385901b1c1916600184901b1784556109f5565b5f84815260208120601f198516915b82811015610a8b5787850151825560209485019460019092019101610a6b565b5084821015610aa857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215610ac7575f5ffd5b5051919050565b5f5f5f60608486031215610ae0575f5ffd5b5050815160208301516040909301519094929350919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156109ab575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516134a6610bc45f395f818161042d01528181610f1d015261254e01525f61040101525f6116af01525f61168201525f6112d601525f6112ae01525f61120901525f61123301525f61125d01526134a65ff3fe608060405234801561000f575f5ffd5b506004361061031a575f3560e01c806384ef8ffc116101a8578063bdf7acce116100f3578063d547741f1161009e578063dd62ed3e11610079578063dd62ed3e14610729578063e01d55c51461073c578063ec87621c1461074f578063fe2df3e814610776575f5ffd5b8063d547741f14610706578063d602b9fd14610719578063da0239a614610721575f5ffd5b8063cefc1429116100ce578063cefc14291461069f578063cf6eefb7146106a7578063d505accf146106f3575f5ffd5b8063bdf7acce1461067c578063ca15c87314610684578063cc8463c814610697575f5ffd5b80639ca92df911610153578063a3246ad31161012e578063a3246ad31461064e578063a40283d514610661578063a9059cbb14610669575f5ffd5b80639ca92df91461060d578063a1eda53c14610620578063a217fddf14610647575f5ffd5b80639010d07c116101835780639010d07c146105af57806391d14854146105c257806395d89b4114610605575f5ffd5b806384ef8ffc1461055557806387f45353146105945780638da5cb5b146105a7575f5ffd5b80633644e51511610268578063634e93da1161021357806379cc6790116101ee57806379cc6790146105145780637ecebe001461052757806384b0196e1461053a575f5ffd5b8063634e93da146104db578063649a5ec7146104ee57806370a0823114610501575f5ffd5b8063436ba62611610243578063436ba6261461047f5780635068b4de146104875780635c4e62c4146104c6575f5ffd5b80633644e5151461045157806336568abe1461045957806340c10f191461046c575f5ffd5b806323b872dd116102c85780632f2ff15d116102a35780632f2ff15d146103e7578063313ce567146103fa578063355274ea1461042b575f5ffd5b806323b872dd146103aa578063248a9ca3146103bd5780632d9ed80d146103df575f5ffd5b8063095ea7b3116102f8578063095ea7b3146103775780630aa6220b1461038a57806318160ddd14610394575f5ffd5b806301ffc9a71461031e578063022d63fb1461034657806306fdde0314610362575b5f5ffd5b61033161032c366004612e81565b610789565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff909116815260200161033d565b61036a61098b565b60405161033d9190612f0c565b610331610385366004612f46565b61099a565b6103926109ac565b005b61039c6109c1565b60405190815260200161033d565b6103316103b8366004612f6e565b6109cb565b61039c6103cb366004612fa8565b5f9081526020819052604090206001015490565b60105461039c565b6103926103f5366004612fbf565b6109df565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161033d565b7f000000000000000000000000000000000000000000000000000000000000000061039c565b61039c6109ed565b610392610467366004612fbf565b6109f6565b61039261047a366004612f46565b610a00565b61039c610a0a565b600e54604080516fffffffffffffffffffffffffffffffff83168152700100000000000000000000000000000000909204600f0b60208301520161033d565b6104ce610a33565b60405161033d9190612fe9565b6103926104e9366004613041565b610a3f565b6103926104fc36600461305a565b610a52565b61039c61050f366004613041565b610a65565b610392610522366004612f46565b610a8f565b61039c610535366004613041565b610ac1565b610542610acb565b60405161033d979695949392919061307f565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161033d565b6103316105a2366004613041565b610b29565b61056f610b35565b61056f6105bd36600461313e565b610b55565b6103316105d0366004612fbf565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61036a610b6c565b61056f61061b366004612fa8565b610b76565b610628610b82565b6040805165ffffffffffff93841681529290911660208301520161033d565b61039c5f81565b6104ce61065c366004612fa8565b610bfc565b61039c610c15565b610331610677366004612f46565b610c20565b61039c610c2b565b61039c610692366004612fa8565b610c36565b61034b610c4c565b610392610ce9565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff1660208301520161033d565b61039261070136600461315e565b610d4a565b610392610714366004612fbf565b610ef3565b610392610efd565b61039c610f0f565b61039c6107373660046131cb565b610f54565b61039261074a366004612fa8565b610f8d565b61039c7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b6103926107843660046131f3565b611034565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061081b57507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b8061086757507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b806108b357507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b806108de57507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061092a57507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b8061097657507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b806109855750610985826110c5565b92915050565b60606109956110cf565b905090565b5f6109a5838361115f565b9392505050565b5f6109b681611176565b6109be611180565b50565b5f61099560085490565b5f6109d784848461118c565b949350505050565b6109e982826111af565b5050565b5f6109956111f0565b6109e98282611326565b6109e9828261142b565b5f610a21600f54600e6114df90919063ffffffff16565b610a2c575060115490565b5060105490565b606061099560046114f2565b5f610a4981611176565b6109e9826114fe565b5f610a5c81611176565b6109e98261157d565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260066020526040812054610985565b73ffffffffffffffffffffffffffffffffffffffff82163314610ab757610ab78233836115ec565b6109e982826115f7565b5f61098582611651565b5f6060805f5f5f6060610adc61167b565b610ae46116a8565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f6109856004836116d5565b5f61099560025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206109a59083611703565b606061099561170e565b5f610985600483611703565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610bc457504265ffffffffffff821610155b610bcf575f5f610bf4565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f818152600360205260409020606090610985906114f2565b5f610995600461171d565b5f6109a58383611726565b5f610995600e611733565b5f8181526003602052604081206109859061171d565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610c8d57504265ffffffffffff8216105b610cbf576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610ce3565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610d42576040517fc22c80220000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6109be61173e565b83421115610d87576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610d39565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610ddf8c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610e468261182f565b90505f610e5582878787611876565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610edc576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610d39565b610ee78a8a8a6118a2565b50505050505050505050565b6109e982826118af565b5f610f0781611176565b6109be6118f0565b5f5f610f196109c1565b90507f0000000000000000000000000000000000000000000000000000000000000000818111610f49575f610f4d565b8181035b9250505090565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081526007602090815260408083209385168352929052908120546109a5565b5f610f9781611176565b815f03610ff2576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69740000000000000000000000000000000000000000000000006004820152602401610d39565b60105460408051918252602082018490527f864790bdf9878a0378c6fc2b0ce53bf74ca13b901bc97a1cb94aa88f1600e482910160405180910390a150601055565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61105e81611176565b612e798261106e576118fa611072565b61194d5b9050835f5b818110156110bc576110b4600488888481811061109657611096613279565b90506020020160208101906110ab9190613041565b8563ffffffff16565b600101611077565b50505050505050565b5f61098582611a0f565b6060600980546110de906132a6565b80601f016020809104026020016040519081016040528092919081815260200182805461110a906132a6565b80156111555780601f1061112c57610100808354040283529160200191611155565b820191905f5260205f20905b81548152906001019060200180831161113857829003601f168201915b5050505050905090565b5f3361116c8185856118a2565b5060019392505050565b6109be8133611a19565b61118a5f5f611a9e565b565b5f336111998582856115ec565b6111a4858585611bf7565b506001949350505050565b816111e6576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109e98282611ca0565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561125557507f000000000000000000000000000000000000000000000000000000000000000046145b1561127f57507f000000000000000000000000000000000000000000000000000000000000000090565b610995604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b8115801561134e575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156114215760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806113a2575065ffffffffffff8116155b806113b557504265ffffffffffff821610155b156113f6576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610d39565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6109e98282611cca565b6011545f611439600e611733565b9050600f54811461148557600f81905560105460405181815290925081907ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d19060200160405180910390a25b50818110156114ca576040517f156d32c50000000000000000000000000000000000000000000000000000000081526004810183905260248101829052604401610d39565b8181036011556114da8383611d23565b505050565b5f816114ea84611733565b119392505050565b60605f6109a583611d6e565b5f611507610c4c565b61151042611dc7565b61151a9190613324565b90506115268282611e16565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61158782611eb1565b61159042611dc7565b61159a9190613324565b90506115a68282611a9e565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b6114da838383611ef8565b73ffffffffffffffffffffffffffffffffffffffff8216611646576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610d39565b6109e9825f83611f9b565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d6020526040812054610985565b60606109957f0000000000000000000000000000000000000000000000000000000000000000600b611fa6565b60606109957f0000000000000000000000000000000000000000000000000000000000000000600c611fa6565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156109a5565b5f6109a5838361204f565b6060600a80546110de906132a6565b5f610985825490565b5f3361116c818585611bf7565b5f6109858242612075565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061178e57504265ffffffffffff821610155b156117cf576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610d39565b6117f75f6117f260025473ffffffffffffffffffffffffffffffffffffffff1690565b6121ea565b506118025f836121f5565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f61098561183b6111f0565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f61188688888888612200565b92509250925061189682826122f3565b50909695505050505050565b6114da83838360016123f6565b816118e6576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109e98282612402565b61118a5f5f611e16565b6119048282612426565b156109e95760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff81166119bc576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610d39565b6119c68282612447565b156109e95760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f61098582612468565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166109e9576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610d39565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611b72574265ffffffffffff82161015611b49576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611b72565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff8316611c46576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610d39565b73ffffffffffffffffffffffffffffffffffffffff8216611c95576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610d39565b6114da838383611f9b565b5f82815260208190526040902060010154611cba81611176565b611cc483836121f5565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611d19576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114da82826121ea565b611d2c33610b29565b611d64576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610d39565b6109e982826124bd565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611dbb57602002820191905f5260205f20905b815481526020019060010190808311611da7575b50505050509050919050565b5f65ffffffffffff821115611e12576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610d39565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff8816171790935590041680156114da576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611ebb610c4c565b90508065ffffffffffff168365ffffffffffff1611611ee357611ede8382613342565b6109a5565b6109a565ffffffffffff841662069780612517565b5f611f038484610f54565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015611cc45781811015611f8d576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610d39565b611cc484848484035f6123f6565b6114da838383612526565b606060ff8314611fc057611fb9836125c3565b9050610985565b818054611fcc906132a6565b80601f0160208091040260200160405190810160405280929190818152602001828054611ff8906132a6565b80156120435780601f1061201a57610100808354040283529160200191612043565b820191905f5260205f20905b81548152906001019060200180831161202657829003601f168201915b50505050509050610985565b5f825f01828154811061206457612064613279565b905f5260205f200154905092915050565b81545f906fffffffffffffffffffffffffffffffff1681036120c3576040517f28632f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82546fffffffffffffffffffffffffffffffff1682816120e5576120e5613360565b8454919006909203917001000000000000000000000000000000009004600f0b5f03612112575080610985565b82545f700100000000000000000000000000000000909104600f0b1315612156578254611fb9907001000000000000000000000000000000009004600f0b8361338d565b82545f9061217a907001000000000000000000000000000000009004600f0b6133a0565b9050808310156121d85783546040517ff41373a600000000000000000000000000000000000000000000000000000000815260048101859052700100000000000000000000000000000000909104600f0b6024820152604401610d39565b6121e281846133d6565b915050610985565b5f6109a58383612600565b5f6109a58383612633565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561223957505f915060039050826122e9565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561228a573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166122e057505f9250600191508290506122e9565b92505f91508190505b9450945094915050565b5f826003811115612306576123066133e9565b0361230f575050565b6001826003811115612323576123236133e9565b0361235a576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561236e5761236e6133e9565b036123a8576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610d39565b60038260038111156123bc576123bc6133e9565b036109e9576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610d39565b611cc48484848461265e565b5f8281526020819052604090206001015461241c81611176565b611cc483836121ea565b5f6109a58373ffffffffffffffffffffffffffffffffffffffff84166127a3565b5f6109a58373ffffffffffffffffffffffffffffffffffffffff8416612886565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806109855750610985826128d2565b73ffffffffffffffffffffffffffffffffffffffff821661250c576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610d39565b6109e95f8383611f9b565b5f8282188284100282186109a5565b612531838383612927565b73ffffffffffffffffffffffffffffffffffffffff83166114da577f00000000000000000000000000000000000000000000000000000000000000005f6125766109c1565b9050818111156125bc576040517f168a7f6e0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610d39565b5050505050565b60605f6125cf83612932565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f61260c8484612972565b905080156109a5575f84815260036020526040902061262b9084612426565b509392505050565b5f5f61263f84846129d3565b905080156109a5575f84815260036020526040902061262b9084612447565b73ffffffffffffffffffffffffffffffffffffffff84166126ad576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610d39565b73ffffffffffffffffffffffffffffffffffffffff83166126fc576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610d39565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526007602090815260408083209387168352929052208290558015611cc4578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161279591815260200190565b60405180910390a350505050565b5f818152600183016020526040812054801561287d575f6127c56001836133d6565b85549091505f906127d8906001906133d6565b9050808214612837575f865f0182815481106127f6576127f6613279565b905f5260205f200154905080875f01848154811061281657612816613279565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061284857612848613416565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610985565b5f915050610985565b5f8181526001830160205260408120546128cb57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610985565b505f610985565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610985575061098582612a91565b6114da838383612b27565b5f60ff8216601f811115610985576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8215801561299b575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b156129c957600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6109a58383612cce565b5f82612a87575f6129f960025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612a46576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6109a58383612d87565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061098557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610985565b73ffffffffffffffffffffffffffffffffffffffff8316612b5e578060085f828254612b53919061338d565b90915550612c0e9050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526006602052604090205481811015612be3576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610d39565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216612c3757600880548290039055612c62565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051612cc191815260200190565b60405180910390a3505050565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16156128cb575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610985565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff166128cb575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612e173390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610985565b61118a613443565b5f60208284031215612e91575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146109a5575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109a56020830184612ec0565b803573ffffffffffffffffffffffffffffffffffffffff81168114612f41575f5ffd5b919050565b5f5f60408385031215612f57575f5ffd5b612f6083612f1e565b946020939093013593505050565b5f5f5f60608486031215612f80575f5ffd5b612f8984612f1e565b9250612f9760208501612f1e565b929592945050506040919091013590565b5f60208284031215612fb8575f5ffd5b5035919050565b5f5f60408385031215612fd0575f5ffd5b82359150612fe060208401612f1e565b90509250929050565b602080825282518282018190525f918401906040840190835b8181101561303657835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613002565b509095945050505050565b5f60208284031215613051575f5ffd5b6109a582612f1e565b5f6020828403121561306a575f5ffd5b813565ffffffffffff811681146109a5575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6130b960e0830189612ec0565b82810360408401526130cb8189612ec0565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b8181101561312d57835183526020938401939092019160010161310f565b50909b9a5050505050505050505050565b5f5f6040838503121561314f575f5ffd5b50508035926020909101359150565b5f5f5f5f5f5f5f60e0888a031215613174575f5ffd5b61317d88612f1e565b965061318b60208901612f1e565b95506040880135945060608801359350608088013560ff811681146131ae575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156131dc575f5ffd5b6131e583612f1e565b9150612fe060208401612f1e565b5f5f5f60408486031215613205575f5ffd5b833567ffffffffffffffff81111561321b575f5ffd5b8401601f8101861361322b575f5ffd5b803567ffffffffffffffff811115613241575f5ffd5b8660208260051b8401011115613255575f5ffd5b602091820194509250840135801515811461326e575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c908216806132ba57607f821691505b6020821081036132f1577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff8181168382160190811115610985576109856132f7565b65ffffffffffff8281168282160390811115610985576109856132f7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b80820180821115610985576109856132f7565b5f7f800000000000000000000000000000000000000000000000000000000000000082036133d0576133d06132f7565b505f0390565b81810381811115610985576109856132f7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220c38665220bddf396d7096ad189bbe0038a00c5a94d7866a360ca603392db012b64736f6c634300081c0033",
}

// ERC20CappedMintLimited is an auto generated Go binding around an Ethereum contract.
type ERC20CappedMintLimited struct {
	abi abi.ABI
}

// NewERC20CappedMintLimited creates a new instance of ERC20CappedMintLimited.
func NewERC20CappedMintLimited() *ERC20CappedMintLimited {
	parsed, err := ERC20CappedMintLimitedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20CappedMintLimited{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20CappedMintLimited) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals, bytes[2] extensionData) returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8, extensionData [2][]byte) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("", owner, forges, name, symbol, decimals, extensionData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackMANAGERROLE() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("allowance", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("approve", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackAvailableMintCapacity() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("availableMintCapacity")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAvailableMintCapacity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x436ba626.
//
// Solidity: function availableMintCapacity() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAvailableMintCapacity(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("availableMintCapacity", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("balanceOf", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackCap() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("cap")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackCap(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("cap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackDecimals() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("decimals", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackDefaultAdmin() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("defaultAdmin", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackDefaultAdminDelay() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("defaultAdminDelay", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackEip712Domain() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("eip712Domain", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("forgeByIndex", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackForgeCount() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("forgeCount", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackForges() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("forges", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("getRoleAdmin", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("getRoleMember", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("getRoleMemberCount", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("getRoleMembers", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("hasRole", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("isForge", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackMaxMintPerPeriod() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("maxMintPerPeriod")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriod is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d9ed80d.
//
// Solidity: function maxMintPerPeriod() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackMaxMintPerPeriod(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("maxMintPerPeriod", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackName() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackName(data []byte) (string, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("name", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackNonces(owner common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("nonces", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackOwner() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("owner", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("pendingDefaultAdminDelay", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackPeriodConfig() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("periodConfig")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5068b4de.
//
// Solidity: function periodConfig() view returns(uint256, int256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPeriodConfig(data []byte) (PeriodConfigOutput, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("periodConfig", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackPeriodStartTime() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("periodStartTime")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartTime is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbdf7acce.
//
// Solidity: function periodStartTime() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPeriodStartTime(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("periodStartTime", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemainingSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackRemainingSupply() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("remainingSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemainingSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackRemainingSupply(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("remainingSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("supportsInterface", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackSymbol() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("symbol", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackTotalSupply() []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("totalSupply", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("transfer", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20CappedMintLimited.abi.Unpack("transferFrom", data)
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) PackUpdateMintLimit(newLimit *big.Int) []byte {
	enc, err := eRC20CappedMintLimited.abi.Pack("updateMintLimit", newLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20CappedMintLimitedApproval represents a Approval event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedApproval) ContractEventName() string {
	return ERC20CappedMintLimitedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackApprovalEvent(log *types.Log) (*ERC20CappedMintLimitedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedApproval)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20CappedMintLimitedDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20CappedMintLimitedDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20CappedMintLimitedDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20CappedMintLimitedDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20CappedMintLimitedDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20CappedMintLimitedDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20CappedMintLimitedDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20CappedMintLimitedDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedEIP712DomainChanged) ContractEventName() string {
	return ERC20CappedMintLimitedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20CappedMintLimitedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedForgeAdded represents a ForgeAdded event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedForgeAdded) ContractEventName() string {
	return ERC20CappedMintLimitedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackForgeAddedEvent(log *types.Log) (*ERC20CappedMintLimitedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedForgeRemoved represents a ForgeRemoved event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedForgeRemoved) ContractEventName() string {
	return ERC20CappedMintLimitedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackForgeRemovedEvent(log *types.Log) (*ERC20CappedMintLimitedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedMintLimitUpdated represents a MintLimitUpdated event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedMintLimitUpdated struct {
	OldLimits *big.Int
	NewLimits *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedMintLimitUpdatedEventName = "MintLimitUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedMintLimitUpdated) ContractEventName() string {
	return ERC20CappedMintLimitedMintLimitUpdatedEventName
}

// UnpackMintLimitUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintLimitUpdated(uint256 oldLimits, uint256 newLimits)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackMintLimitUpdatedEvent(log *types.Log) (*ERC20CappedMintLimitedMintLimitUpdated, error) {
	event := "MintLimitUpdated"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedMintLimitUpdated)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedPeriodStarted represents a PeriodStarted event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedPeriodStarted struct {
	PeriodStart       *big.Int
	AvailableCapacity *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedPeriodStartedEventName = "PeriodStarted"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedPeriodStarted) ContractEventName() string {
	return ERC20CappedMintLimitedPeriodStartedEventName
}

// UnpackPeriodStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPeriodStartedEvent(log *types.Log) (*ERC20CappedMintLimitedPeriodStarted, error) {
	event := "PeriodStarted"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedPeriodStarted)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedRoleAdminChanged) ContractEventName() string {
	return ERC20CappedMintLimitedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20CappedMintLimitedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedRoleGranted represents a RoleGranted event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedRoleGranted) ContractEventName() string {
	return ERC20CappedMintLimitedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackRoleGrantedEvent(log *types.Log) (*ERC20CappedMintLimitedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedRoleRevoked represents a RoleRevoked event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedRoleRevoked) ContractEventName() string {
	return ERC20CappedMintLimitedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackRoleRevokedEvent(log *types.Log) (*ERC20CappedMintLimitedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMintLimitedTransfer represents a Transfer event raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMintLimitedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMintLimitedTransfer) ContractEventName() string {
	return ERC20CappedMintLimitedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackTransferEvent(log *types.Log) (*ERC20CappedMintLimitedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20CappedMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMintLimitedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMintLimited.abi.Events[event].Inputs {
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
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20CapableERC20ExceededCap"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20CapableERC20ExceededCapError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20PeriodMintLimitExceedsPeriodLimit"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20PeriodMintLimitExceedsPeriodLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20PeriodMintLimitInvalidLength"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20PeriodMintLimitInvalidLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC20PeriodMintLimitInvalidLimitData"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC20PeriodMintLimitInvalidLimitDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["PeriodManagerInvalidDuration"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackPeriodManagerInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["PeriodManagerInvalidOffset"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackPeriodManagerInvalidOffsetError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMintLimited.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20CappedMintLimited.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20CappedMintLimitedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20CappedMintLimitedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20CappedMintLimitedAccessControlBadConfirmation, error) {
	out := new(ERC20CappedMintLimitedAccessControlBadConfirmation)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20CappedMintLimitedAccessControlEnforcedDefaultAdminRules)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20CappedMintLimitedAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20CappedMintLimitedAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20CappedMintLimitedAccessControlInvalidDefaultAdmin)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20CappedMintLimitedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20CappedMintLimitedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20CappedMintLimitedAccessControlUnauthorizedAccount)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20CappedMintLimitedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20CappedMintLimitedECDSAInvalidSignature, error) {
	out := new(ERC20CappedMintLimitedECDSAInvalidSignature)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20CappedMintLimitedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20CappedMintLimitedECDSAInvalidSignatureLength, error) {
	out := new(ERC20CappedMintLimitedECDSAInvalidSignatureLength)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20CappedMintLimitedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20CappedMintLimitedECDSAInvalidSignatureS, error) {
	out := new(ERC20CappedMintLimitedECDSAInvalidSignatureS)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20CapableERC20ExceededCap represents a ERC20Capable__ERC20ExceededCap error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20CapableERC20ExceededCap struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func ERC20CappedMintLimitedERC20CapableERC20ExceededCapErrorID() common.Hash {
	return common.HexToHash("0x168a7f6ea7d992fe05deb1d140d19ff46ba8920881431f8d796b5ceade4790fa")
}

// UnpackERC20CapableERC20ExceededCapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20CapableERC20ExceededCapError(raw []byte) (*ERC20CappedMintLimitedERC20CapableERC20ExceededCap, error) {
	out := new(ERC20CappedMintLimitedERC20CapableERC20ExceededCap)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20CapableERC20ExceededCap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20CappedMintLimitedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20CappedMintLimitedERC20InsufficientAllowance, error) {
	out := new(ERC20CappedMintLimitedERC20InsufficientAllowance)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20CappedMintLimitedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20CappedMintLimitedERC20InsufficientBalance, error) {
	out := new(ERC20CappedMintLimitedERC20InsufficientBalance)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20CappedMintLimitedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20InvalidApproverError(raw []byte) (*ERC20CappedMintLimitedERC20InvalidApprover, error) {
	out := new(ERC20CappedMintLimitedERC20InvalidApprover)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20CappedMintLimitedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20CappedMintLimitedERC20InvalidReceiver, error) {
	out := new(ERC20CappedMintLimitedERC20InvalidReceiver)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20CappedMintLimitedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20InvalidSenderError(raw []byte) (*ERC20CappedMintLimitedERC20InvalidSender, error) {
	out := new(ERC20CappedMintLimitedERC20InvalidSender)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20CappedMintLimitedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20CappedMintLimitedERC20InvalidSpender, error) {
	out := new(ERC20CappedMintLimitedERC20InvalidSpender)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20PeriodMintLimitExceedsPeriodLimit represents a ERC20PeriodMintLimit__ExceedsPeriodLimit error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20PeriodMintLimitExceedsPeriodLimit struct {
	Requested *big.Int
	Available *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available)
func ERC20CappedMintLimitedERC20PeriodMintLimitExceedsPeriodLimitErrorID() common.Hash {
	return common.HexToHash("0x156d32c58749d9c638f5b3f8e359a7ee3d1aa2d1c7d31a6301ce10260bfd45b4")
}

// UnpackERC20PeriodMintLimitExceedsPeriodLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20PeriodMintLimitExceedsPeriodLimitError(raw []byte) (*ERC20CappedMintLimitedERC20PeriodMintLimitExceedsPeriodLimit, error) {
	out := new(ERC20CappedMintLimitedERC20PeriodMintLimitExceedsPeriodLimit)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitExceedsPeriodLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLength represents a ERC20PeriodMintLimit__InvalidLength error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLength()
func ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLengthErrorID() common.Hash {
	return common.HexToHash("0xf38583ae0b0acd19cabaa28c7b388b3f7e6260d4a4adbf38b03fb876a11643f0")
}

// UnpackERC20PeriodMintLimitInvalidLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLength()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20PeriodMintLimitInvalidLengthError(raw []byte) (*ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLength, error) {
	out := new(ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLength)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitInvalidLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLimitData represents a ERC20PeriodMintLimit__InvalidLimitData error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLimitData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLimitData()
func ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLimitDataErrorID() common.Hash {
	return common.HexToHash("0x436572d08cf45a9b922f4897fbf93bb0e84720cc3edcbc7328ffcb075921f9a8")
}

// UnpackERC20PeriodMintLimitInvalidLimitDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodMintLimit__InvalidLimitData()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC20PeriodMintLimitInvalidLimitDataError(raw []byte) (*ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLimitData, error) {
	out := new(ERC20CappedMintLimitedERC20PeriodMintLimitInvalidLimitData)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodMintLimitInvalidLimitData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20CappedMintLimitedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20CappedMintLimitedERC2612ExpiredSignature, error) {
	out := new(ERC20CappedMintLimitedERC2612ExpiredSignature)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20CappedMintLimitedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20CappedMintLimitedERC2612InvalidSigner, error) {
	out := new(ERC20CappedMintLimitedERC2612InvalidSigner)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20CappedMintLimitedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackInvalidAccountNonceError(raw []byte) (*ERC20CappedMintLimitedInvalidAccountNonce, error) {
	out := new(ERC20CappedMintLimitedInvalidAccountNonce)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedInvalidShortString represents a InvalidShortString error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20CappedMintLimitedInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackInvalidShortStringError(raw []byte) (*ERC20CappedMintLimitedInvalidShortString, error) {
	out := new(ERC20CappedMintLimitedInvalidShortString)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedPeriodManagerInvalidDuration represents a PeriodManager__InvalidDuration error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedPeriodManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidDuration()
func ERC20CappedMintLimitedPeriodManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x28632f2e6184743eb990383edd4904c645090260b538d05e23af6d26cb637330")
}

// UnpackPeriodManagerInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidDuration()
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPeriodManagerInvalidDurationError(raw []byte) (*ERC20CappedMintLimitedPeriodManagerInvalidDuration, error) {
	out := new(ERC20CappedMintLimitedPeriodManagerInvalidDuration)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedPeriodManagerInvalidOffset represents a PeriodManager__InvalidOffset error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedPeriodManagerInvalidOffset struct {
	Timestamp     *big.Int
	OffsetSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func ERC20CappedMintLimitedPeriodManagerInvalidOffsetErrorID() common.Hash {
	return common.HexToHash("0xf41373a6f64e5d167ce7bcc91a19d341958fd006a060f85330cfc8aa087401d6")
}

// UnpackPeriodManagerInvalidOffsetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackPeriodManagerInvalidOffsetError(raw []byte) (*ERC20CappedMintLimitedPeriodManagerInvalidOffset, error) {
	out := new(ERC20CappedMintLimitedPeriodManagerInvalidOffset)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidOffset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func ERC20CappedMintLimitedSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*ERC20CappedMintLimitedSafeCastOverflowedIntDowncast, error) {
	out := new(ERC20CappedMintLimitedSafeCastOverflowedIntDowncast)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20CappedMintLimitedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20CappedMintLimitedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20CappedMintLimitedSafeCastOverflowedUintDowncast)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedStringTooLong represents a StringTooLong error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20CappedMintLimitedStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackStringTooLongError(raw []byte) (*ERC20CappedMintLimitedStringTooLong, error) {
	out := new(ERC20CappedMintLimitedStringTooLong)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20CappedMintLimitedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackTokenBaseNullInputError(raw []byte) (*ERC20CappedMintLimitedTokenBaseNullInput, error) {
	out := new(ERC20CappedMintLimitedTokenBaseNullInput)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMintLimitedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20CappedMintLimited contract.
type ERC20CappedMintLimitedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20CappedMintLimitedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20CappedMintLimited *ERC20CappedMintLimited) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20CappedMintLimitedTokenBaseOnlyForge, error) {
	out := new(ERC20CappedMintLimitedTokenBaseOnlyForge)
	if err := eRC20CappedMintLimited.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
