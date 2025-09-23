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

// ERC20CappedMultiMintLimitedMetaData contains all meta data concerning the ERC20CappedMultiMintLimited contract.
var ERC20CappedMultiMintLimitedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes[2]\",\"name\":\"extensionData\",\"type\":\"bytes[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableMintCapacities\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintPerPeriods\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodConfigs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodStartTimes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"updateMintLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"oldLimits\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"}],\"name\":\"MintLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"availableCapacity\",\"type\":\"uint256\"}],\"name\":\"PeriodStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__ExceedsPeriodLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20PeriodsMintLimit__InvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC20PeriodsMintLimit__InvalidLimitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodManager__InvalidDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offsetSeconds\",\"type\":\"int256\"}],\"name\":\"PeriodManager__InvalidOffset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20CappedMultiMintLimited",
	Bin: "0x6101a0604052348015610010575f5ffd5b50604051614b7d380380614b7d83398101604081905261002f91610ab5565b8060016020020151815f602002015187878787878280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b0316036100a257604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100cb5f82610549565b505050506100ff7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361054960201b60201c565b5080515f5b8181101561013e57610136600484838151811061012357610123610b8f565b602002602001015161055d60201b60201c565b600101610104565b5050505081600990816101519190610c27565b50600a61015e8282610c27565b5061016e9150839050600b6105d5565b6101205261017d81600c6105d5565b61014052815160208084019190912060e052815190820120610100524660a05261020960e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361023d5760405163e1dea2ef60e01b8152636e616d6560e01b6004820152602401610099565b81515f036102695760405163e1dea2ef60e01b8152651cde5b589bdb60d21b6004820152602401610099565b60ff1661016052505082515f925061028b915083016020908101908401610ce1565b9050805f036102b55760405163e1dea2ef60e01b81526206361760ec1b6004820152602401610099565b610180525080515f90819081906102d59085016020908101908601610d53565b805192955090935091508015806102ed575083518114155b806102f9575082518114155b1561031757604051633fffd6c160e01b815260040160405180910390fd5b5f80805b8381101561046f575f5f88838151811061033757610337610b8f565b602002602001015187848151811061035157610351610b8f565b602002602001015191509150815f1480610369575080155b156103a95760405163e1dea2ef60e01b81527f6c696d697473206f72206475726174696f6e73000000000000000000000000006004820152602401610099565b84821115806103b85750838111155b156103d95760405163360b969560e21b815260048101849052602401610099565b600f60405180604001604052806103f58561060560201b60201c565b6001600160801b0316815260200161042b8b878151811061041857610418610b8f565b602002602001015161063c60201b60201c565b600f0b90528154600181810184555f938452602093849020835194909301516001600160801b03908116600160801b0294169390931791015591945092500161031b565b50600e8390558351610488906010906020870190610846565b50826001600160401b038111156104a1576104a16108b0565b6040519080825280602002602001820160405280156104ca578160200160208202803683370190505b5080516104df91601191602090910190610846565b50826001600160401b038111156104f8576104f86108b0565b604051908082528060200260200182016040528015610521578160200160208202803683370190505b50805161053691601291602090910190610846565b5050505050505050505050505050610e91565b5f6105548383610670565b90505b92915050565b6001600160a01b03811661058e5760405163e1dea2ef60e01b815264666f72676560d81b6004820152602401610099565b61059882826106a3565b156105d1576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f6020835110156105f0576105e9836106b7565b9050610557565b816105fb8482610c27565b5060ff9050610557565b5f6001600160801b03821115610638576040516306dfcc6560e41b81526080600482015260248101839052604401610099565b5090565b80600f81900b811461066b5760405163327269a760e01b81526080600482015260248101839052604401610099565b919050565b5f8061067c84846106f4565b90508015610554575f84815260036020526040902061069b90846106a3565b509392505050565b5f610554836001600160a01b03841661075a565b5f5f829050601f815111156106e1578260405163305a27a960e01b81526004016100999190610e39565b80516106ec82610e6e565b179392505050565b5f82610750575f61070d6002546001600160a01b031690565b6001600160a01b03161461073457604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b61055483836107a6565b5f81815260018301602052604081205461079f57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610557565b505f610557565b5f828152602081815260408083206001600160a01b038516845290915281205460ff1661079f575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556107fe3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610557565b828054828255905f5260205f2090810192821561087f579160200282015b8281111561087f578251825591602001919060010190610864565b506106389291505b80821115610638575f8155600101610887565b80516001600160a01b038116811461066b575f5ffd5b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156108e6576108e66108b0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610914576109146108b0565b604052919050565b5f6001600160401b03821115610934576109346108b0565b5060051b60200190565b5f82601f83011261094d575f5ffd5b815161096061095b8261091c565b6108ec565b8082825260208201915060208360051b860101925085831115610981575f5ffd5b602085015b838110156109a5576109978161089a565b835260209283019201610986565b5095945050505050565b5f806001600160401b038411156109c8576109c86108b0565b50601f8301601f19166020016109dd816108ec565b9150508281528383830111156109f1575f5ffd5b8282602083015e5f602084830101529392505050565b5f82601f830112610a16575f5ffd5b610554838351602085016109af565b805160ff8116811461066b575f5ffd5b5f82601f830112610a44575f5ffd5b610a4c6108c4565b806040840185811115610a5d575f5ffd5b845b81811015610aaa5780516001600160401b03811115610a7c575f5ffd5b8601601f81018813610a8c575f5ffd5b610a9b888251602084016109af565b85525060209384019301610a5f565b509095945050505050565b5f5f5f5f5f5f60c08789031215610aca575f5ffd5b610ad38761089a565b60208801519096506001600160401b03811115610aee575f5ffd5b610afa89828a0161093e565b604089015190965090506001600160401b03811115610b17575f5ffd5b610b2389828a01610a07565b606089015190955090506001600160401b03811115610b40575f5ffd5b610b4c89828a01610a07565b935050610b5b60808801610a25565b60a08801519092506001600160401b03811115610b76575f5ffd5b610b8289828a01610a35565b9150509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b600181811c90821680610bb757607f821691505b602082108103610bd557634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610c2257805f5260205f20601f840160051c81016020851015610c005750805b601f840160051c820191505b81811015610c1f575f8155600101610c0c565b50505b505050565b81516001600160401b03811115610c4057610c406108b0565b610c5481610c4e8454610ba3565b84610bdb565b6020601f821160018114610c86575f8315610c6f5750848201515b5f19600385901b1c1916600184901b178455610c1f565b5f84815260208120601f198516915b82811015610cb55787850151825560209485019460019092019101610c95565b5084821015610cd257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215610cf1575f5ffd5b5051919050565b5f82601f830112610d07575f5ffd5b8151610d1561095b8261091c565b8082825260208201915060208360051b860101925085831115610d36575f5ffd5b602085015b838110156109a5578051835260209283019201610d3b565b5f5f5f60608486031215610d65575f5ffd5b83516001600160401b03811115610d7a575f5ffd5b610d8686828701610cf8565b602086015190945090506001600160401b03811115610da3575f5ffd5b8401601f81018613610db3575f5ffd5b8051610dc161095b8261091c565b8082825260208201915060208360051b850101925088831115610de2575f5ffd5b6020840193505b82841015610e04578351825260209384019390910190610de9565b6040880151909550925050506001600160401b03811115610e23575f5ffd5b610e2f86828701610cf8565b9150509250925092565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610bd5575f1960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516101605161018051613c79610f045f395f818161043d0152818161148c0152612b4501525f61041101525f611cb701525f611c8a01525f6117a001525f61177801525f6116d301525f6116fd01525f6117270152613c795ff3fe608060405234801561000f575f5ffd5b5060043610610319575f3560e01c806387f45353116101a8578063ca15c873116100f3578063d505accf1161009e578063da0239a611610079578063da0239a61461071b578063dd62ed3e14610723578063ec87621c14610736578063fe2df3e81461075d575f5ffd5b8063d505accf146106ed578063d547741f14610700578063d602b9fd14610713575f5ffd5b8063cdd760b9116100ce578063cdd760b914610683578063cefc142914610699578063cf6eefb7146106a1575f5ffd5b8063ca15c87314610660578063cc8463c814610673578063ccf6d7241461067b575f5ffd5b8063a1eda53c11610153578063a40283d51161012e578063a40283d51461063d578063a9059cbb14610645578063c3ac4c2d14610658575f5ffd5b8063a1eda53c146105fc578063a217fddf14610623578063a3246ad31461062a575f5ffd5b806391d148541161018357806391d148541461059e57806395d89b41146105e15780639ca92df9146105e9575f5ffd5b806387f45353146105705780638da5cb5b146105835780639010d07c1461058b575f5ffd5b80633644e51511610268578063649a5ec7116102135780637ecebe00116101ee5780637ecebe001461050357806384b0196e1461051657806384ef8ffc14610531575f5ffd5b8063649a5ec7146104ca57806370a08231146104dd57806379cc6790146104f0575f5ffd5b806340c10f191161024357806340c10f191461048f5780635c4e62c4146104a2578063634e93da146104b7575f5ffd5b80633644e5151461046157806336568abe1461046957806337e653ea1461047c575f5ffd5b806318160ddd116102c85780632f2ff15d116102a35780632f2ff15d146103f7578063313ce5671461040a578063355274ea1461043b575f5ffd5b806318160ddd146103ac57806323b872dd146103c2578063248a9ca3146103d5575f5ffd5b806306fdde03116102f857806306fdde031461037a578063095ea7b31461038f5780630aa6220b146103a2575f5ffd5b806286ced81461031d57806301ffc9a71461033b578063022d63fb1461035e575b5f5ffd5b610325610770565b604051610332919061350b565b60405180910390f35b61034e61034936600461351d565b610815565b6040519015158152602001610332565b620697805b60405165ffffffffffff9091168152602001610332565b610382610a17565b60405161033291906135a8565b61034e61039d3660046135e2565b610a26565b6103aa610a38565b005b6103b4610a4d565b604051908152602001610332565b61034e6103d036600461360a565b610a57565b6103b46103e3366004613644565b5f9081526020819052604090206001015490565b6103aa61040536600461365b565b610a6b565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610332565b7f00000000000000000000000000000000000000000000000000000000000000006103b4565b6103b4610a79565b6103aa61047736600461365b565b610a82565b6103aa61048a3660046136cd565b610a8c565b6103aa61049d3660046135e2565b610c7a565b6104aa610c84565b604051610332919061370c565b6103aa6104c5366004613764565b610c90565b6103aa6104d836600461377d565b610ca3565b6103b46104eb366004613764565b610cb6565b6103aa6104fe3660046135e2565b610ce0565b6103b4610511366004613764565b610d12565b61051e610d1c565b60405161033297969594939291906137a2565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610332565b61034e61057e366004613764565b610d7a565b61054b610d86565b61054b610599366004613838565b610da6565b61034e6105ac36600461365b565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610382610dbd565b61054b6105f7366004613644565b610dc7565b610604610dd3565b6040805165ffffffffffff938416815292909116602083015201610332565b6103b45f81565b6104aa610638366004613644565b610e4d565b6103b4610e66565b61034e6106533660046135e2565b610e71565b610325610e7c565b6103b461066e366004613644565b610ffa565b610363611010565b6103256110ad565b61068b611103565b604051610332929190613858565b6103aa61125d565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610332565b6103aa6106fb3660046138a4565b6112b9565b6103aa61070e36600461365b565b611462565b6103aa61146c565b6103b461147e565b6103b4610731366004613911565b6114c3565b6103b47faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b6103aa61076b366004613939565b6114fc565b600e546060905f8167ffffffffffffffff81111561079057610790613990565b6040519080825280602002602001820160405280156107b9578160200160208202803683370190505b5090505f5b8281101561080e576107e9600f82815481106107dc576107dc6139bd565b905f5260205f200161158d565b8282815181106107fb576107fb6139bd565b60209081029190910101526001016107be565b5092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a042910000000000000000000000000000000000000000000000000000000014806108a757507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b806108f357507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b8061093f57507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b8061096a57507fffffffff000000000000000000000000000000000000000000000000000000008216155b806109b657507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b80610a0257507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b80610a115750610a1182611598565b92915050565b6060610a216115a2565b905090565b5f610a318383611629565b9392505050565b5f610a4281611640565b610a4a61164a565b50565b5f610a2160085490565b5f610a63848484611656565b949350505050565b610a758282611679565b5050565b5f610a216116ba565b610a7582826117f0565b5f610a9681611640565b815f819003610af8576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d697473000000000000000000000000000000000000000000000060048201526024015b60405180910390fd5b600e548114610b33576040517f3fffd6c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805b82811015610c29575f868683818110610b5157610b516139bd565b905060200201359050805f03610bb5576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f6e65774c696d69747300000000000000000000000000000000000000000000006004820152602401610aef565b828111158015610be557507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114155b15610c1f576040517fd82e5a5400000000000000000000000000000000000000000000000000000000815260048101839052602401610aef565b9150600101610b36565b507f89be4a9eb7586334d719b1c87a0fd1e8306235335bf257d957c53aa7b22fcb7460108686604051610c5e939291906139ea565b60405180910390a1610c7260108686613470565b505050505050565b610a7582826118f5565b6060610a216004611af5565b5f610c9a81611640565b610a7582611b01565b5f610cad81611640565b610a7582611b80565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260066020526040812054610a11565b73ffffffffffffffffffffffffffffffffffffffff82163314610d0857610d08823383611bef565b610a758282611bff565b5f610a1182611c59565b5f6060805f5f5f6060610d2d611c83565b610d35611cb0565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f610a11600483611cdd565b5f610a2160025473ffffffffffffffffffffffffffffffffffffffff1690565b5f828152600360205260408120610a319083611d0b565b6060610a21611d16565b5f610a11600483611d0b565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610e1557504265ffffffffffff821610155b610e20575f5f610e45565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f818152600360205260409020606090610a1190611af5565b5f610a216004611d25565b5f610a318383611d2e565b60605f610e87610770565b90505f6011805480602002602001604051908101604052809291908181526020018280548015610ed457602002820191905f5260205f20905b815481526020019060010190808311610ec0575b505050505090505f600e5490505f8167ffffffffffffffff811115610efb57610efb613990565b604051908082528060200260200182016040528015610f24578160200160208202803683370190505b5090505f5b82811015610ff1575f848281518110610f4457610f446139bd565b60200260200101519050858281518110610f6057610f606139bd565b60200260200101518103610fad5760128281548110610f8157610f816139bd565b905f5260205f200154838381518110610f9c57610f9c6139bd565b602002602001018181525050610fe8565b60108281548110610fc057610fc06139bd565b905f5260205f200154838381518110610fdb57610fdb6139bd565b6020026020010181815250505b50600101610f29565b50949350505050565b5f818152600360205260408120610a1190611d25565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801515801561105157504265ffffffffffff8216105b611083576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff166110a7565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b606060108054806020026020016040519081016040528092919081815260200182805480156110f957602002820191905f5260205f20905b8154815260200190600101908083116110e5575b5050505050905090565b600e5460609081905f8167ffffffffffffffff81111561112557611125613990565b60405190808252806020026020018201604052801561114e578160200160208202803683370190505b5090505f8267ffffffffffffffff81111561116b5761116b613990565b604051908082528060200260200182016040528015611194578160200160208202803683370190505b5090505f5b8381101561125257600f81815481106111b4576111b46139bd565b5f9182526020909120015483516fffffffffffffffffffffffffffffffff909116908490839081106111e8576111e86139bd565b602002602001018181525050600f8181548110611207576112076139bd565b5f918252602090912001548251700100000000000000000000000000000000909104600f0b9083908390811061123f5761123f6139bd565b6020908102919091010152600101611199565b509094909350915050565b60015473ffffffffffffffffffffffffffffffffffffffff163381146112b1576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610aef565b610a4a611d3b565b834211156112f6576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610aef565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861134e8c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f6113b582611e2c565b90505f6113c482878787611e73565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461144b576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610aef565b6114568a8a8a611e9f565b50505050505050505050565b610a758282611eac565b5f61147681611640565b610a4a611eed565b5f5f611488610a4d565b90507f00000000000000000000000000000000000000000000000000000000000000008181116114b8575f6114bc565b8181035b9250505090565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152600760209081526040808320938516835292905290812054610a31565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c61152681611640565b6134b58261153657611ef761153a565b611f4a5b9050835f5b818110156115845761157c600488888481811061155e5761155e6139bd565b90506020020160208101906115739190613764565b8563ffffffff16565b60010161153f565b50505050505050565b5f610a11824261200c565b5f610a1182612188565b6060600980546115b190613a79565b80601f01602080910402602001604051908101604052809291908181526020018280546115dd90613a79565b80156110f95780601f106115ff576101008083540402835291602001916110f9565b820191905f5260205f20905b81548152906001019060200180831161160b57509395945050505050565b5f33611636818585611e9f565b5060019392505050565b610a4a8133612192565b6116545f5f612217565b565b5f33611663858285611bef565b61166e858585612370565b506001949350505050565b816116b0576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a758282612419565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561171f57507f000000000000000000000000000000000000000000000000000000000000000046145b1561174957507f000000000000000000000000000000000000000000000000000000000000000090565b610a21604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b81158015611818575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156118eb5760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168115158061186c575065ffffffffffff8116155b8061187f57504265ffffffffffff821610155b156118c0576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610aef565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b610a75828261243d565b5f6118fe610770565b600e549091505f5b81811015611ae4575f5f5f60128481548110611924576119246139bd565b905f5260205f20015460118581548110611940576119406139bd565b905f5260205f20015487868151811061195b5761195b6139bd565b60200260200101519250925092508082146119ee575f60108581548110611984576119846139bd565b905f5260205f200154905081601186815481106119a3576119a36139bd565b905f5260205f200181905550809350817ff9c0c5a7cf394a1a1fc99bbc29c65c560ead7b51cdccdacb7f392071f67c13d1826040516119e491815260200190565b60405180910390a2505b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8303611a39578260128581548110611a2957611a296139bd565b5f91825260209091200155611ad6565b86831015611ab457600f8481548110611a5457611a546139bd565b5f918252602090912001546040517faace2cbe0000000000000000000000000000000000000000000000000000000081526fffffffffffffffffffffffffffffffff90911660048201526024810188905260448101849052606401610aef565b86830360128581548110611aca57611aca6139bd565b5f918252602090912001555b505050806001019050611906565b50611aef8484612496565b50505050565b60605f610a31836124e1565b5f611b0a611010565b611b134261253a565b611b1d9190613af7565b9050611b298282612589565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f611b8a82612624565b611b934261253a565b611b9d9190613af7565b9050611ba98282612217565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b611bfa83838361266b565b505050565b73ffffffffffffffffffffffffffffffffffffffff8216611c4e576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610aef565b610a75825f8361270e565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d6020526040812054610a11565b6060610a217f0000000000000000000000000000000000000000000000000000000000000000600b612719565b6060610a217f0000000000000000000000000000000000000000000000000000000000000000600c612719565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610a31565b5f610a3183836127bb565b6060600a80546115b190613a79565b5f610a11825490565b5f33611636818585612370565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16801580611d8b57504265ffffffffffff821610155b15611dcc576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610aef565b611df45f611def60025473ffffffffffffffffffffffffffffffffffffffff1690565b6127e1565b50611dff5f836127ec565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f610a11611e386116ba565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f611e83888888886127f7565b925092509250611e9382826128ea565b50909695505050505050565b611bfa83838360016129ed565b81611ee3576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a7582826129f9565b6116545f5f612589565b611f018282612a1d565b15610a755760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611fb9576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610aef565b611fc38282612a3e565b15610a755760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b81545f906fffffffffffffffffffffffffffffffff16810361205a576040517f28632f2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82546fffffffffffffffffffffffffffffffff16828161207c5761207c613b15565b8454919006909203917001000000000000000000000000000000009004600f0b5f036120a9575080610a11565b82545f700100000000000000000000000000000000909104600f0b13156120f45782546120ed907001000000000000000000000000000000009004600f0b83613b42565b9050610a11565b82545f90612118907001000000000000000000000000000000009004600f0b613b55565b9050808310156121765783546040517ff41373a600000000000000000000000000000000000000000000000000000000815260048101859052700100000000000000000000000000000000909104600f0b6024820152604401610aef565b6121808184613b8b565b915050610a11565b5f610a1182612a5f565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610a75576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610aef565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680156122eb574265ffffffffffff821610156122c2576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a010000000000000000000000000000000000000000000000000000029190911790556122eb565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff83166123bf576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610aef565b73ffffffffffffffffffffffffffffffffffffffff821661240e576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610aef565b611bfa83838361270e565b5f8281526020819052604090206001015461243381611640565b611aef83836127ec565b73ffffffffffffffffffffffffffffffffffffffff8116331461248c576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bfa82826127e1565b61249f33610d7a565b6124d7576040517f25a9dbc1000000000000000000000000000000000000000000000000000000008152336004820152602401610aef565b610a758282612ab4565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561252e57602002820191905f5260205f20905b81548152602001906001019080831161251a575b50505050509050919050565b5f65ffffffffffff821115612585576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610aef565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015611bfa576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f61262e611010565b90508065ffffffffffff168365ffffffffffff1611612656576126518382613b9e565b610a31565b610a3165ffffffffffff841662069780612b0e565b5f61267684846114c3565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015611aef5781811015612700576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610aef565b611aef84848484035f6129ed565b611bfa838383612b1d565b606060ff831461272c576120ed83612bba565b81805461273890613a79565b80601f016020809104026020016040519081016040528092919081815260200182805461276490613a79565b80156127af5780601f10612786576101008083540402835291602001916127af565b820191905f5260205f20905b81548152906001019060200180831161279257829003601f168201915b50505050509050610a11565b5f825f0182815481106127d0576127d06139bd565b905f5260205f200154905092915050565b5f610a318383612bf7565b5f610a318383612c2a565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561283057505f915060039050826128e0565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612881573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166128d757505f9250600191508290506128e0565b92505f91508190505b9450945094915050565b5f8260038111156128fd576128fd613bbc565b03612906575050565b600182600381111561291a5761291a613bbc565b03612951576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561296557612965613bbc565b0361299f576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610aef565b60038260038111156129b3576129b3613bbc565b03610a75576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610aef565b611aef84848484612c55565b5f82815260208190526040902060010154612a1381611640565b611aef83836127e1565b5f610a318373ffffffffffffffffffffffffffffffffffffffff8416612d9a565b5f610a318373ffffffffffffffffffffffffffffffffffffffff8416612e7d565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f000000000000000000000000000000000000000000000000000000001480610a115750610a1182612ec9565b73ffffffffffffffffffffffffffffffffffffffff8216612b03576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610aef565b610a755f838361270e565b5f828218828410028218610a31565b612b28838383612f1e565b73ffffffffffffffffffffffffffffffffffffffff8316611bfa577f00000000000000000000000000000000000000000000000000000000000000005f612b6d610a4d565b905081811115612bb3576040517f168a7f6e0000000000000000000000000000000000000000000000000000000081526004810182905260248101839052604401610aef565b5050505050565b60605f612bc683612f29565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f612c038484612f69565b90508015610a31575f848152600360205260409020612c229084612a1d565b509392505050565b5f5f612c368484612fca565b90508015610a31575f848152600360205260409020612c229084612a3e565b73ffffffffffffffffffffffffffffffffffffffff8416612ca4576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610aef565b73ffffffffffffffffffffffffffffffffffffffff8316612cf3576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610aef565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526007602090815260408083209387168352929052208290558015611aef578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051612d8c91815260200190565b60405180910390a350505050565b5f8181526001830160205260408120548015612e74575f612dbc600183613b8b565b85549091505f90612dcf90600190613b8b565b9050808214612e2e575f865f018281548110612ded57612ded6139bd565b905f5260205f200154905080875f018481548110612e0d57612e0d6139bd565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080612e3f57612e3f613be9565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610a11565b5f915050610a11565b5f818152600183016020526040812054612ec257508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610a11565b505f610a11565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f31498786000000000000000000000000000000000000000000000000000000001480610a115750610a1182613088565b611bfa83838361311e565b5f60ff8216601f811115610a11576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82158015612f92575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b15612fc057600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b610a3183836132c5565b5f8261307e575f612ff060025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461303d576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b610a31838361337e565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610a1157507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610a11565b73ffffffffffffffffffffffffffffffffffffffff8316613155578060085f82825461314a9190613b42565b909155506132059050565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260066020526040902054818110156131da576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610aef565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661322e57600880548290039055613259565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516132b891815260200190565b60405180910390a3505050565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615612ec2575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610a11565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16612ec2575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561340e3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610a11565b828054828255905f5260205f209081019282156134a9579160200282015b828111156134a957823582559160200191906001019061348e565b506125859291506134bd565b611654613c16565b5b80821115612585575f81556001016134be565b5f8151808452602084019350602083015f5b828110156135015781518652602095860195909101906001016134e3565b5093949350505050565b602081525f610a3160208301846134d1565b5f6020828403121561352d575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610a31575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610a31602083018461355c565b803573ffffffffffffffffffffffffffffffffffffffff811681146135dd575f5ffd5b919050565b5f5f604083850312156135f3575f5ffd5b6135fc836135ba565b946020939093013593505050565b5f5f5f6060848603121561361c575f5ffd5b613625846135ba565b9250613633602085016135ba565b929592945050506040919091013590565b5f60208284031215613654575f5ffd5b5035919050565b5f5f6040838503121561366c575f5ffd5b8235915061367c602084016135ba565b90509250929050565b5f5f83601f840112613695575f5ffd5b50813567ffffffffffffffff8111156136ac575f5ffd5b6020830191508360208260051b85010111156136c6575f5ffd5b9250929050565b5f5f602083850312156136de575f5ffd5b823567ffffffffffffffff8111156136f4575f5ffd5b61370085828601613685565b90969095509350505050565b602080825282518282018190525f918401906040840190835b8181101561375957835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101613725565b509095945050505050565b5f60208284031215613774575f5ffd5b610a31826135ba565b5f6020828403121561378d575f5ffd5b813565ffffffffffff81168114610a31575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f6137dc60e083018961355c565b82810360408401526137ee818961355c565b905086606084015273ffffffffffffffffffffffffffffffffffffffff861660808401528460a084015282810360c084015261382a81856134d1565b9a9950505050505050505050565b5f5f60408385031215613849575f5ffd5b50508035926020909101359150565b604081525f61386a60408301856134d1565b82810360208401528084518083526020830191506020860192505f5b81811015611e93578351835260209384019390920191600101613886565b5f5f5f5f5f5f5f60e0888a0312156138ba575f5ffd5b6138c3886135ba565b96506138d1602089016135ba565b95506040880135945060608801359350608088013560ff811681146138f4575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215613922575f5ffd5b61392b836135ba565b915061367c602084016135ba565b5f5f5f6040848603121561394b575f5ffd5b833567ffffffffffffffff811115613961575f5ffd5b61396d86828701613685565b90945092505060208401358015158114613985575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b604080825284549082018190525f8581526020812090916060840190835b81811015613a26578354835260019384019360209093019201613a08565b505083810360208501528481527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff851115613a5f575f5ffd5b8460051b9150818660208301370160200195945050505050565b600181811c90821680613a8d57607f821691505b602082108103613ac4577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff8181168382160190811115610a1157610a11613aca565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b80820180821115610a1157610a11613aca565b5f7f80000000000000000000000000000000000000000000000000000000000000008203613b8557613b85613aca565b505f0390565b81810381811115610a1157610a11613aca565b65ffffffffffff8281168282160390811115610a1157610a11613aca565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212207c22909f39927f83586ac52e2e49f0a61d90a4975e1ff231ebcb7d5d1eeb749264736f6c634300081c0033",
}

// ERC20CappedMultiMintLimited is an auto generated Go binding around an Ethereum contract.
type ERC20CappedMultiMintLimited struct {
	abi abi.ABI
}

// NewERC20CappedMultiMintLimited creates a new instance of ERC20CappedMultiMintLimited.
func NewERC20CappedMultiMintLimited() *ERC20CappedMultiMintLimited {
	parsed, err := ERC20CappedMultiMintLimitedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20CappedMultiMintLimited{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20CappedMultiMintLimited) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals, bytes[2] extensionData) returns()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8, extensionData [2][]byte) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("", owner, forges, name, symbol, decimals, extensionData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackMANAGERROLE() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("allowance", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("approve", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackAvailableMintCapacities() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("availableMintCapacities")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAvailableMintCapacities is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc3ac4c2d.
//
// Solidity: function availableMintCapacities() view returns(uint256[])
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAvailableMintCapacities(data []byte) ([]*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("availableMintCapacities", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("balanceOf", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackCap() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("cap")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackCap(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("cap", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackDecimals() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("decimals", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackDefaultAdmin() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("defaultAdmin", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackDefaultAdminDelay() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("defaultAdminDelay", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackEip712Domain() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("eip712Domain", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("forgeByIndex", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackForgeCount() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("forgeCount", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackForges() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("forges", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("getRoleAdmin", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("getRoleMember", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("getRoleMemberCount", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("getRoleMembers", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("hasRole", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("isForge", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackMaxMintPerPeriods() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("maxMintPerPeriods")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxMintPerPeriods is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xccf6d724.
//
// Solidity: function maxMintPerPeriods() view returns(uint256[])
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackMaxMintPerPeriods(data []byte) ([]*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("maxMintPerPeriods", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackName() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackName(data []byte) (string, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("name", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackNonces(owner common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("nonces", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackOwner() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("owner", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("pendingDefaultAdminDelay", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackPeriodConfigs() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("periodConfigs")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodConfigs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcdd760b9.
//
// Solidity: function periodConfigs() view returns(uint256[], int256[])
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPeriodConfigs(data []byte) (PeriodConfigsOutput, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("periodConfigs", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackPeriodStartTimes() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("periodStartTimes")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPeriodStartTimes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0086ced8.
//
// Solidity: function periodStartTimes() view returns(uint256[])
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPeriodStartTimes(data []byte) ([]*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("periodStartTimes", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemainingSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackRemainingSupply() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("remainingSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemainingSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackRemainingSupply(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("remainingSupply", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("supportsInterface", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackSymbol() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("symbol", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackTotalSupply() []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("totalSupply", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("transfer", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20CappedMultiMintLimited.abi.Unpack("transferFrom", data)
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) PackUpdateMintLimits(newLimits []*big.Int) []byte {
	enc, err := eRC20CappedMultiMintLimited.abi.Pack("updateMintLimits", newLimits)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC20CappedMultiMintLimitedApproval represents a Approval event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedApproval) ContractEventName() string {
	return ERC20CappedMultiMintLimitedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackApprovalEvent(log *types.Log) (*ERC20CappedMultiMintLimitedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedApproval)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20CappedMultiMintLimitedDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20CappedMultiMintLimitedDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20CappedMultiMintLimitedDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20CappedMultiMintLimitedDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedEIP712DomainChanged) ContractEventName() string {
	return ERC20CappedMultiMintLimitedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedForgeAdded represents a ForgeAdded event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedForgeAdded) ContractEventName() string {
	return ERC20CappedMultiMintLimitedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackForgeAddedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedForgeRemoved represents a ForgeRemoved event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedForgeRemoved) ContractEventName() string {
	return ERC20CappedMultiMintLimitedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackForgeRemovedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedMintLimitUpdated represents a MintLimitUpdated event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedMintLimitUpdated struct {
	OldLimits []*big.Int
	NewLimits []*big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedMintLimitUpdatedEventName = "MintLimitUpdated"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedMintLimitUpdated) ContractEventName() string {
	return ERC20CappedMultiMintLimitedMintLimitUpdatedEventName
}

// UnpackMintLimitUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackMintLimitUpdatedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedMintLimitUpdated, error) {
	event := "MintLimitUpdated"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedMintLimitUpdated)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedPeriodStarted represents a PeriodStarted event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedPeriodStarted struct {
	PeriodStart       *big.Int
	AvailableCapacity *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedPeriodStartedEventName = "PeriodStarted"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedPeriodStarted) ContractEventName() string {
	return ERC20CappedMultiMintLimitedPeriodStartedEventName
}

// UnpackPeriodStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPeriodStartedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedPeriodStarted, error) {
	event := "PeriodStarted"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedPeriodStarted)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedRoleAdminChanged) ContractEventName() string {
	return ERC20CappedMultiMintLimitedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedRoleGranted represents a RoleGranted event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedRoleGranted) ContractEventName() string {
	return ERC20CappedMultiMintLimitedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackRoleGrantedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedRoleRevoked represents a RoleRevoked event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedRoleRevoked) ContractEventName() string {
	return ERC20CappedMultiMintLimitedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackRoleRevokedEvent(log *types.Log) (*ERC20CappedMultiMintLimitedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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

// ERC20CappedMultiMintLimitedTransfer represents a Transfer event raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedMultiMintLimitedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20CappedMultiMintLimitedTransfer) ContractEventName() string {
	return ERC20CappedMultiMintLimitedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackTransferEvent(log *types.Log) (*ERC20CappedMultiMintLimitedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20CappedMultiMintLimited.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedMultiMintLimitedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20CappedMultiMintLimited.abi.Events[event].Inputs {
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
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20CapableERC20ExceededCap"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20CapableERC20ExceededCapError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20PeriodsMintLimitExceedsPeriodLimit"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20PeriodsMintLimitExceedsPeriodLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20PeriodsMintLimitInvalidLength"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20PeriodsMintLimitInvalidLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC20PeriodsMintLimitInvalidLimitData"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC20PeriodsMintLimitInvalidLimitDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["PeriodManagerInvalidDuration"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackPeriodManagerInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["PeriodManagerInvalidOffset"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackPeriodManagerInvalidOffsetError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["SafeCastOverflowedIntDowncast"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackSafeCastOverflowedIntDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20CappedMultiMintLimited.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20CappedMultiMintLimited.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20CappedMultiMintLimitedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20CappedMultiMintLimitedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20CappedMultiMintLimitedAccessControlBadConfirmation, error) {
	out := new(ERC20CappedMultiMintLimitedAccessControlBadConfirmation)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20CappedMultiMintLimitedAccessControlEnforcedDefaultAdminRules)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20CappedMultiMintLimitedAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20CappedMultiMintLimitedAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20CappedMultiMintLimitedAccessControlInvalidDefaultAdmin)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20CappedMultiMintLimitedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20CappedMultiMintLimitedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20CappedMultiMintLimitedAccessControlUnauthorizedAccount)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20CappedMultiMintLimitedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20CappedMultiMintLimitedECDSAInvalidSignature, error) {
	out := new(ERC20CappedMultiMintLimitedECDSAInvalidSignature)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20CappedMultiMintLimitedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20CappedMultiMintLimitedECDSAInvalidSignatureLength, error) {
	out := new(ERC20CappedMultiMintLimitedECDSAInvalidSignatureLength)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20CappedMultiMintLimitedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20CappedMultiMintLimitedECDSAInvalidSignatureS, error) {
	out := new(ERC20CappedMultiMintLimitedECDSAInvalidSignatureS)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20CapableERC20ExceededCap represents a ERC20Capable__ERC20ExceededCap error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20CapableERC20ExceededCap struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func ERC20CappedMultiMintLimitedERC20CapableERC20ExceededCapErrorID() common.Hash {
	return common.HexToHash("0x168a7f6ea7d992fe05deb1d140d19ff46ba8920881431f8d796b5ceade4790fa")
}

// UnpackERC20CapableERC20ExceededCapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20CapableERC20ExceededCapError(raw []byte) (*ERC20CappedMultiMintLimitedERC20CapableERC20ExceededCap, error) {
	out := new(ERC20CappedMultiMintLimitedERC20CapableERC20ExceededCap)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20CapableERC20ExceededCap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20CappedMultiMintLimitedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20CappedMultiMintLimitedERC20InsufficientAllowance, error) {
	out := new(ERC20CappedMultiMintLimitedERC20InsufficientAllowance)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20CappedMultiMintLimitedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20CappedMultiMintLimitedERC20InsufficientBalance, error) {
	out := new(ERC20CappedMultiMintLimitedERC20InsufficientBalance)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20CappedMultiMintLimitedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20InvalidApproverError(raw []byte) (*ERC20CappedMultiMintLimitedERC20InvalidApprover, error) {
	out := new(ERC20CappedMultiMintLimitedERC20InvalidApprover)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20CappedMultiMintLimitedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20CappedMultiMintLimitedERC20InvalidReceiver, error) {
	out := new(ERC20CappedMultiMintLimitedERC20InvalidReceiver)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20CappedMultiMintLimitedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20InvalidSenderError(raw []byte) (*ERC20CappedMultiMintLimitedERC20InvalidSender, error) {
	out := new(ERC20CappedMultiMintLimitedERC20InvalidSender)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20CappedMultiMintLimitedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20CappedMultiMintLimitedERC20InvalidSpender, error) {
	out := new(ERC20CappedMultiMintLimitedERC20InvalidSpender)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit represents a ERC20PeriodsMintLimit__ExceedsPeriodLimit error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit struct {
	Period    *big.Int
	Requested *big.Int
	Available *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available)
func ERC20CappedMultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimitErrorID() common.Hash {
	return common.HexToHash("0xaace2cbe23069663ea57043f9018ae5639c94c4b7537987827b85a4a671f9687")
}

// UnpackERC20PeriodsMintLimitExceedsPeriodLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20PeriodsMintLimitExceedsPeriodLimitError(raw []byte) (*ERC20CappedMultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit, error) {
	out := new(ERC20CappedMultiMintLimitedERC20PeriodsMintLimitExceedsPeriodLimit)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitExceedsPeriodLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLength represents a ERC20PeriodsMintLimit__InvalidLength error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLength()
func ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLengthErrorID() common.Hash {
	return common.HexToHash("0x3fffd6c1aac13cb22e11380a50735d4cdde9620742e341a28fbb7c0113c01dbe")
}

// UnpackERC20PeriodsMintLimitInvalidLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLength()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20PeriodsMintLimitInvalidLengthError(raw []byte) (*ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLength, error) {
	out := new(ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLength)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitInvalidLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLimitData represents a ERC20PeriodsMintLimit__InvalidLimitData error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLimitData struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index)
func ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLimitDataErrorID() common.Hash {
	return common.HexToHash("0xd82e5a5493b676ee930d3b301ed5da6339364f86d1828cf9de0b2e2baa8068b3")
}

// UnpackERC20PeriodsMintLimitInvalidLimitDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC20PeriodsMintLimitInvalidLimitDataError(raw []byte) (*ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLimitData, error) {
	out := new(ERC20CappedMultiMintLimitedERC20PeriodsMintLimitInvalidLimitData)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC20PeriodsMintLimitInvalidLimitData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20CappedMultiMintLimitedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20CappedMultiMintLimitedERC2612ExpiredSignature, error) {
	out := new(ERC20CappedMultiMintLimitedERC2612ExpiredSignature)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20CappedMultiMintLimitedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20CappedMultiMintLimitedERC2612InvalidSigner, error) {
	out := new(ERC20CappedMultiMintLimitedERC2612InvalidSigner)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20CappedMultiMintLimitedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackInvalidAccountNonceError(raw []byte) (*ERC20CappedMultiMintLimitedInvalidAccountNonce, error) {
	out := new(ERC20CappedMultiMintLimitedInvalidAccountNonce)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedInvalidShortString represents a InvalidShortString error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20CappedMultiMintLimitedInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackInvalidShortStringError(raw []byte) (*ERC20CappedMultiMintLimitedInvalidShortString, error) {
	out := new(ERC20CappedMultiMintLimitedInvalidShortString)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedPeriodManagerInvalidDuration represents a PeriodManager__InvalidDuration error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedPeriodManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidDuration()
func ERC20CappedMultiMintLimitedPeriodManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x28632f2e6184743eb990383edd4904c645090260b538d05e23af6d26cb637330")
}

// UnpackPeriodManagerInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidDuration()
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPeriodManagerInvalidDurationError(raw []byte) (*ERC20CappedMultiMintLimitedPeriodManagerInvalidDuration, error) {
	out := new(ERC20CappedMultiMintLimitedPeriodManagerInvalidDuration)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedPeriodManagerInvalidOffset represents a PeriodManager__InvalidOffset error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedPeriodManagerInvalidOffset struct {
	Timestamp     *big.Int
	OffsetSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func ERC20CappedMultiMintLimitedPeriodManagerInvalidOffsetErrorID() common.Hash {
	return common.HexToHash("0xf41373a6f64e5d167ce7bcc91a19d341958fd006a060f85330cfc8aa087401d6")
}

// UnpackPeriodManagerInvalidOffsetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PeriodManager__InvalidOffset(uint256 timestamp, int256 offsetSeconds)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackPeriodManagerInvalidOffsetError(raw []byte) (*ERC20CappedMultiMintLimitedPeriodManagerInvalidOffset, error) {
	out := new(ERC20CappedMultiMintLimitedPeriodManagerInvalidOffset)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "PeriodManagerInvalidOffset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedSafeCastOverflowedIntDowncast represents a SafeCastOverflowedIntDowncast error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedSafeCastOverflowedIntDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func ERC20CappedMultiMintLimitedSafeCastOverflowedIntDowncastErrorID() common.Hash {
	return common.HexToHash("0x327269a7f29c3c5436f42eeed1c1adf0d4d525f36360483f4e83ab79e98f9089")
}

// UnpackSafeCastOverflowedIntDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedIntDowncast(uint8 bits, int256 value)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackSafeCastOverflowedIntDowncastError(raw []byte) (*ERC20CappedMultiMintLimitedSafeCastOverflowedIntDowncast, error) {
	out := new(ERC20CappedMultiMintLimitedSafeCastOverflowedIntDowncast)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedIntDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20CappedMultiMintLimitedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20CappedMultiMintLimitedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20CappedMultiMintLimitedSafeCastOverflowedUintDowncast)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedStringTooLong represents a StringTooLong error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20CappedMultiMintLimitedStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackStringTooLongError(raw []byte) (*ERC20CappedMultiMintLimitedStringTooLong, error) {
	out := new(ERC20CappedMultiMintLimitedStringTooLong)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20CappedMultiMintLimitedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackTokenBaseNullInputError(raw []byte) (*ERC20CappedMultiMintLimitedTokenBaseNullInput, error) {
	out := new(ERC20CappedMultiMintLimitedTokenBaseNullInput)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedMultiMintLimitedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20CappedMultiMintLimited contract.
type ERC20CappedMultiMintLimitedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20CappedMultiMintLimitedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20CappedMultiMintLimited *ERC20CappedMultiMintLimited) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20CappedMultiMintLimitedTokenBaseOnlyForge, error) {
	out := new(ERC20CappedMultiMintLimitedTokenBaseOnlyForge)
	if err := eRC20CappedMultiMintLimited.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
