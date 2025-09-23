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

// ERC20CappedMetaData contains all meta data concerning the ERC20Capped contract.
var ERC20CappedMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"extensionData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"increasedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"ERC20Capable__ERC20ExceededCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20Capped",
	Bin: "0x6101a0604052348015610010575f5ffd5b50604051613a2e380380613a2e83398101604081905261002f916106b0565b8086868686868280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b03160361009457604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100bd5f826102b8565b505050506100f17faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c836102b860201b60201c565b5080515f5b818110156101305761012860048483815181106101155761011561078a565b60200260200101516102cc60201b60201c565b6001016100f6565b5050505081600990816101439190610822565b50600a6101508282610822565b506101609150839050600b610344565b6101205261016f81600c610344565b61014052815160208084019190912060e052815190820120610100524660a0526101fb60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361022f5760405163e1dea2ef60e01b8152636e616d6560e01b600482015260240161008b565b81515f0361025b5760405163e1dea2ef60e01b8152651cde5b589bdb60d21b600482015260240161008b565b60ff1661016052505082515f925061027d9150830160209081019084016108dc565b9050805f036102a75760405163e1dea2ef60e01b81526206361760ec1b600482015260240161008b565b610180525061094b95505050505050565b5f6102c38383610374565b90505b92915050565b6001600160a01b0381166102fd5760405163e1dea2ef60e01b815264666f72676560d81b600482015260240161008b565b61030782826103a7565b15610340576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f60208351101561035f57610358836103bb565b90506102c6565b8161036a8482610822565b5060ff90506102c6565b5f8061038084846103f8565b905080156102c3575f84815260036020526040902061039f90846103a7565b509392505050565b5f6102c3836001600160a01b03841661045e565b5f5f829050601f815111156103e5578260405163305a27a960e01b815260040161008b91906108f3565b80516103f082610928565b179392505050565b5f82610454575f6104116002546001600160a01b031690565b6001600160a01b03161461043857604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6102c383836104aa565b5f8181526001830160205260408120546104a357508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556102c6565b505f6102c6565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166104a3575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556105023390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016102c6565b80516001600160a01b0381168114610560575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156105a1576105a1610565565b604052919050565b5f82601f8301126105b8575f5ffd5b81516001600160401b038111156105d1576105d1610565565b8060051b6105e160208201610579565b918252602081850181019290810190868411156105fc575f5ffd5b6020860192505b83831015610625576106148361054a565b825260209283019290910190610603565b9695505050505050565b5f82601f83011261063e575f5ffd5b8151602083015f806001600160401b0384111561065d5761065d610565565b50601f8301601f191660200161067281610579565b915050828152858383011115610686575f5ffd5b8282602083015e5f92810160200192909252509392505050565b805160ff81168114610560575f5ffd5b5f5f5f5f5f5f60c087890312156106c5575f5ffd5b6106ce8761054a565b60208801519096506001600160401b038111156106e9575f5ffd5b6106f589828a016105a9565b604089015190965090506001600160401b03811115610712575f5ffd5b61071e89828a0161062f565b606089015190955090506001600160401b0381111561073b575f5ffd5b61074789828a0161062f565b935050610756608088016106a0565b60a08801519092506001600160401b03811115610771575f5ffd5b61077d89828a0161062f565b9150509295509295509295565b634e487b7160e01b5f52603260045260245ffd5b600181811c908216806107b257607f821691505b6020821081036107d057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561081d57805f5260205f20601f840160051c810160208510156107fb5750805b601f840160051c820191505b8181101561081a575f8155600101610807565b50505b505050565b81516001600160401b0381111561083b5761083b610565565b61084f81610849845461079e565b846107d6565b6020601f821160018114610881575f831561086a5750848201515b5f19600385901b1c1916600184901b17845561081a565b5f84815260208120601f198516915b828110156108b05787850151825560209485019460019092019101610890565b50848210156108cd57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f602082840312156108ec575f5ffd5b5051919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156107d0575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516130706109be5f395f81816103ee01528181610e890152611b9d01525f6103c201525f61150c01525f6114df01525f61119b01525f61117301525f6110ce01525f6110f801525f61112201526130705ff3fe608060405234801561000f575f5ffd5b50600436106102e3575f3560e01c806387f4535311610187578063ca15c873116100dd578063d547741f11610093578063dd62ed3e1161006e578063dd62ed3e1461069b578063ec87621c146106ae578063fe2df3e8146106d5575f5ffd5b8063d547741f14610678578063d602b9fd1461068b578063da0239a614610693575f5ffd5b8063cefc1429116100c3578063cefc142914610611578063cf6eefb714610619578063d505accf14610665575f5ffd5b8063ca15c873146105f6578063cc8463c814610609575f5ffd5b80639ca92df91161013d578063a3246ad311610118578063a3246ad3146105c8578063a40283d5146105db578063a9059cbb146105e3575f5ffd5b80639ca92df914610587578063a1eda53c1461059a578063a217fddf146105c1575f5ffd5b80639010d07c1161016d5780639010d07c1461052957806391d148541461053c57806395d89b411461057f575f5ffd5b806387f453531461050e5780638da5cb5b14610521575f5ffd5b80633644e5151161023c578063649a5ec7116101f25780637ecebe00116101cd5780637ecebe00146104a157806384b0196e146104b457806384ef8ffc146104cf575f5ffd5b8063649a5ec71461046857806370a082311461047b57806379cc67901461048e575f5ffd5b806340c10f191161022257806340c10f191461042d5780635c4e62c414610440578063634e93da14610455575f5ffd5b80633644e5151461041257806336568abe1461041a575f5ffd5b806318160ddd1161029c5780632f2ff15d116102775780632f2ff15d146103a8578063313ce567146103bb578063355274ea146103ec575f5ffd5b806318160ddd1461035d57806323b872dd14610373578063248a9ca314610386575f5ffd5b806306fdde03116102cc57806306fdde031461032b578063095ea7b3146103405780630aa6220b14610353575f5ffd5b806301ffc9a7146102e7578063022d63fb1461030f575b5f5ffd5b6102fa6102f5366004612aae565b6106e8565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff9091168152602001610306565b6103336108ea565b6040516103069190612b39565b6102fa61034e366004612b73565b6108f9565b61035b61090b565b005b610365610920565b604051908152602001610306565b6102fa610381366004612b9b565b61092a565b610365610394366004612bd5565b5f9081526020819052604090206001015490565b61035b6103b6366004612bec565b61093e565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610306565b7f0000000000000000000000000000000000000000000000000000000000000000610365565b61036561094c565b61035b610428366004612bec565b610955565b61035b61043b366004612b73565b61095f565b6104486109af565b6040516103069190612c16565b61035b610463366004612c6e565b6109bb565b61035b610476366004612c87565b6109ce565b610365610489366004612c6e565b6109e1565b61035b61049c366004612b73565b610a0b565b6103656104af366004612c6e565b610a3d565b6104bc610a47565b6040516103069796959493929190612cac565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610306565b6102fa61051c366004612c6e565b610aa5565b6104e9610ab1565b6104e9610537366004612d6b565b610ad1565b6102fa61054a366004612bec565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610333610ae8565b6104e9610595366004612bd5565b610af2565b6105a2610afe565b6040805165ffffffffffff938416815292909116602083015201610306565b6103655f81565b6104486105d6366004612bd5565b610b78565b610365610b91565b6102fa6105f1366004612b73565b610b9c565b610365610604366004612bd5565b610ba7565b610314610bbd565b61035b610c5a565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff16602083015201610306565b61035b610673366004612d8b565b610cb6565b61035b610686366004612bec565b610e5f565b61035b610e69565b610365610e7b565b6103656106a9366004612df8565b610ec0565b6103657faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61035b6106e3366004612e20565b610ef9565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061077a57507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b806107c657507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b8061081257507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b8061083d57507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061088957507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b806108d557507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b806108e457506108e482610f8a565b92915050565b60606108f4610f94565b905090565b5f6109048383611024565b9392505050565b5f6109158161103b565b61091d611045565b50565b5f6108f460085490565b5f610936848484611051565b949350505050565b6109488282611074565b5050565b5f6108f46110b5565b61094882826111eb565b61096833610aa5565b6109a5576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61094882826112f0565b60606108f4600461134a565b5f6109c58161103b565b61094882611356565b5f6109d88161103b565b610948826113d5565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600660205260408120546108e4565b73ffffffffffffffffffffffffffffffffffffffff82163314610a3357610a33823383611444565b6109488282611454565b5f6108e4826114ae565b5f6060805f5f5f6060610a586114d8565b610a60611505565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f6108e4600483611532565b5f6108f460025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206109049083611560565b60606108f461156b565b5f6108e4600483611560565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610b4057504265ffffffffffff821610155b610b4b575f5f610b70565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f8181526003602052604090206060906108e49061134a565b5f6108f4600461157a565b5f6109048383611583565b5f8181526003602052604081206108e49061157a565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610bfe57504265ffffffffffff8216105b610c30576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610c54565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610cae576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161099c565b61091d611590565b83421115610cf3576040517f627913020000000000000000000000000000000000000000000000000000000081526004810185905260240161099c565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610d4b8c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610db282611681565b90505f610dc1828787876116c8565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e48576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b16602482015260440161099c565b610e538a8a8a6116f4565b50505050505050505050565b6109488282611701565b5f610e738161103b565b61091d611742565b5f5f610e85610920565b90507f0000000000000000000000000000000000000000000000000000000000000000818111610eb5575f610eb9565b8181035b9250505090565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152600760209081526040808320938516835292905290812054610904565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610f238161103b565b612aa682610f335761174c610f37565b61179f5b9050835f5b81811015610f8157610f796004888884818110610f5b57610f5b612ea6565b9050602002016020810190610f709190612c6e565b8563ffffffff16565b600101610f3c565b50505050505050565b5f6108e482611861565b606060098054610fa390612ed3565b80601f0160208091040260200160405190810160405280929190818152602001828054610fcf90612ed3565b801561101a5780601f10610ff15761010080835404028352916020019161101a565b820191905f5260205f20905b815481529060010190602001808311610ffd57829003601f168201915b5050505050905090565b5f336110318185856116f4565b5060019392505050565b61091d813361186b565b61104f5f5f6118f0565b565b5f3361105e858285611444565b611069858585611a49565b506001949350505050565b816110ab576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109488282611af2565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561111a57507f000000000000000000000000000000000000000000000000000000000000000046145b1561114457507f000000000000000000000000000000000000000000000000000000000000000090565b6108f4604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b81158015611213575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156112e65760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611267575065ffffffffffff8116155b8061127a57504265ffffffffffff821610155b156112bb576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161099c565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6109488282611b1c565b73ffffffffffffffffffffffffffffffffffffffff821661133f576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161099c565b6109485f8383611b75565b60605f61090483611c12565b5f61135f610bbd565b61136842611c6b565b6113729190612f51565b905061137e8282611cba565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f6113df82611d55565b6113e842611c6b565b6113f29190612f51565b90506113fe82826118f0565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b61144f838383611d9c565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166114a3576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161099c565b610948825f83611b75565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d60205260408120546108e4565b60606108f47f0000000000000000000000000000000000000000000000000000000000000000600b611e3f565b60606108f47f0000000000000000000000000000000000000000000000000000000000000000600c611e3f565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610904565b5f6109048383611ee8565b6060600a8054610fa390612ed3565b5f6108e4825490565b5f33611031818585611a49565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168015806115e057504265ffffffffffff821610155b15611621576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161099c565b6116495f61164460025473ffffffffffffffffffffffffffffffffffffffff1690565b611f0e565b506116545f83611f19565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f6108e461168d6110b5565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f6116d888888888611f24565b9250925092506116e88282612017565b50909695505050505050565b61144f838383600161211a565b81611738576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109488282612126565b61104f5f5f611cba565b611756828261214a565b156109485760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff811661180e576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f726765000000000000000000000000000000000000000000000000000000600482015260240161099c565b611818828261216b565b156109485760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f6108e48261218c565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610948576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161099c565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680156119c4574265ffffffffffff8216101561199b576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a010000000000000000000000000000000000000000000000000000029190911790556119c4565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff8316611a98576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f600482015260240161099c565b73ffffffffffffffffffffffffffffffffffffffff8216611ae7576040517fec442f050000000000000000000000000000000000000000000000000000000081525f600482015260240161099c565b61144f838383611b75565b5f82815260208190526040902060010154611b0c8161103b565b611b168383611f19565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611b6b576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61144f8282611f0e565b611b808383836121e1565b73ffffffffffffffffffffffffffffffffffffffff831661144f577f00000000000000000000000000000000000000000000000000000000000000005f611bc5610920565b905081811115611c0b576040517f168a7f6e000000000000000000000000000000000000000000000000000000008152600481018290526024810183905260440161099c565b5050505050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611c5f57602002820191905f5260205f20905b815481526020019060010190808311611c4b575b50505050509050919050565b5f65ffffffffffff821115611cb6576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161099c565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff88161717909355900416801561144f576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611d5f610bbd565b90508065ffffffffffff168365ffffffffffff1611611d8757611d828382612f6f565b610904565b61090465ffffffffffff8416620697806121ec565b5f611da78484610ec0565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015611b165781811015611e31576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161099c565b611b1684848484035f61211a565b606060ff8314611e5957611e52836121fb565b90506108e4565b818054611e6590612ed3565b80601f0160208091040260200160405190810160405280929190818152602001828054611e9190612ed3565b8015611edc5780601f10611eb357610100808354040283529160200191611edc565b820191905f5260205f20905b815481529060010190602001808311611ebf57829003601f168201915b505050505090506108e4565b5f825f018281548110611efd57611efd612ea6565b905f5260205f200154905092915050565b5f6109048383612238565b5f610904838361226b565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611f5d57505f9150600390508261200d565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611fae573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661200457505f92506001915082905061200d565b92505f91508190505b9450945094915050565b5f82600381111561202a5761202a612f8d565b03612033575050565b600182600381111561204757612047612f8d565b0361207e576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600282600381111561209257612092612f8d565b036120cc576040517ffce698f70000000000000000000000000000000000000000000000000000000081526004810182905260240161099c565b60038260038111156120e0576120e0612f8d565b03610948576040517fd78bce0c0000000000000000000000000000000000000000000000000000000081526004810182905260240161099c565b611b1684848484612296565b5f828152602081905260409020600101546121408161103b565b611b168383611f0e565b5f6109048373ffffffffffffffffffffffffffffffffffffffff84166123db565b5f6109048373ffffffffffffffffffffffffffffffffffffffff84166124be565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806108e457506108e48261250a565b61144f83838361255f565b5f828218828410028218610904565b60605f61220783612706565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f6122448484612746565b90508015610904575f848152600360205260409020612263908461214a565b509392505050565b5f5f61227784846127a7565b90508015610904575f848152600360205260409020612263908461216b565b73ffffffffffffffffffffffffffffffffffffffff84166122e5576040517fe602df050000000000000000000000000000000000000000000000000000000081525f600482015260240161099c565b73ffffffffffffffffffffffffffffffffffffffff8316612334576040517f94280d620000000000000000000000000000000000000000000000000000000081525f600482015260240161099c565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526007602090815260408083209387168352929052208290558015611b16578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516123cd91815260200190565b60405180910390a350505050565b5f81815260018301602052604081205480156124b5575f6123fd600183612fba565b85549091505f9061241090600190612fba565b905080821461246f575f865f01828154811061242e5761242e612ea6565b905f5260205f200154905080875f01848154811061244e5761244e612ea6565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061248057612480612fcd565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506108e4565b5f9150506108e4565b5f81815260018301602052604081205461250357508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556108e4565b505f6108e4565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806108e457506108e482612865565b73ffffffffffffffffffffffffffffffffffffffff8316612596578060085f82825461258b9190612ffa565b909155506126469050565b73ffffffffffffffffffffffffffffffffffffffff83165f908152600660205260409020548181101561261b576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161099c565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661266f5760088054829003905561269a565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516126f991815260200190565b60405180910390a3505050565b5f60ff8216601f8111156108e4576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8215801561276f575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561279d57600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61090483836128fb565b5f8261285b575f6127cd60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161461281a576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b61090483836129b4565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108e457507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146108e4565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615612503575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016108e4565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16612503575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612a443390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016108e4565b61104f61300d565b5f60208284031215612abe575f5ffd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610904575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109046020830184612aed565b803573ffffffffffffffffffffffffffffffffffffffff81168114612b6e575f5ffd5b919050565b5f5f60408385031215612b84575f5ffd5b612b8d83612b4b565b946020939093013593505050565b5f5f5f60608486031215612bad575f5ffd5b612bb684612b4b565b9250612bc460208501612b4b565b929592945050506040919091013590565b5f60208284031215612be5575f5ffd5b5035919050565b5f5f60408385031215612bfd575f5ffd5b82359150612c0d60208401612b4b565b90509250929050565b602080825282518282018190525f918401906040840190835b81811015612c6357835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612c2f565b509095945050505050565b5f60208284031215612c7e575f5ffd5b61090482612b4b565b5f60208284031215612c97575f5ffd5b813565ffffffffffff81168114610904575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f612ce660e0830189612aed565b8281036040840152612cf88189612aed565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015612d5a578351835260209384019390920191600101612d3c565b50909b9a5050505050505050505050565b5f5f60408385031215612d7c575f5ffd5b50508035926020909101359150565b5f5f5f5f5f5f5f60e0888a031215612da1575f5ffd5b612daa88612b4b565b9650612db860208901612b4b565b95506040880135945060608801359350608088013560ff81168114612ddb575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215612e09575f5ffd5b612e1283612b4b565b9150612c0d60208401612b4b565b5f5f5f60408486031215612e32575f5ffd5b833567ffffffffffffffff811115612e48575f5ffd5b8401601f81018613612e58575f5ffd5b803567ffffffffffffffff811115612e6e575f5ffd5b8660208260051b8401011115612e82575f5ffd5b6020918201945092508401358015158114612e9b575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c90821680612ee757607f821691505b602082108103612f1e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff81811683821601908111156108e4576108e4612f24565b65ffffffffffff82811682821603908111156108e4576108e4612f24565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b818103818111156108e4576108e4612f24565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b808201808211156108e4576108e4612f24565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea264697066735822122012d4a055a0feb589c57e367d4691ce1e2dccf7ab12679b16e2e277f4d803158664736f6c634300081c0033",
}

// ERC20Capped is an auto generated Go binding around an Ethereum contract.
type ERC20Capped struct {
	abi abi.ABI
}

// NewERC20Capped creates a new instance of ERC20Capped.
func NewERC20Capped() *ERC20Capped {
	parsed, err := ERC20CappedMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Capped{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Capped) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals, bytes extensionData) returns()
func (eRC20Capped *ERC20Capped) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8, extensionData []byte) []byte {
	enc, err := eRC20Capped.abi.Pack("", owner, forges, name, symbol, decimals, extensionData)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Capped *ERC20Capped) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20Capped.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20Capped *ERC20Capped) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Capped.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20Capped *ERC20Capped) PackMANAGERROLE() []byte {
	enc, err := eRC20Capped.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20Capped *ERC20Capped) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20Capped.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Capped *ERC20Capped) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("allowance", data)
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
func (eRC20Capped *ERC20Capped) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Capped *ERC20Capped) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("approve", data)
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
func (eRC20Capped *ERC20Capped) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("balanceOf", data)
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
func (eRC20Capped *ERC20Capped) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20Capped *ERC20Capped) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20Capped *ERC20Capped) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20Capped.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20Capped *ERC20Capped) PackCap() []byte {
	enc, err := eRC20Capped.abi.Pack("cap")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackCap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackCap(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("cap", data)
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
func (eRC20Capped *ERC20Capped) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Capped *ERC20Capped) PackDecimals() []byte {
	enc, err := eRC20Capped.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Capped *ERC20Capped) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Capped.abi.Unpack("decimals", data)
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
func (eRC20Capped *ERC20Capped) PackDefaultAdmin() []byte {
	enc, err := eRC20Capped.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20Capped *ERC20Capped) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("defaultAdmin", data)
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
func (eRC20Capped *ERC20Capped) PackDefaultAdminDelay() []byte {
	enc, err := eRC20Capped.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20Capped *ERC20Capped) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("defaultAdminDelay", data)
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
func (eRC20Capped *ERC20Capped) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20Capped.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20Capped *ERC20Capped) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC20Capped *ERC20Capped) PackEip712Domain() []byte {
	enc, err := eRC20Capped.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Capped *ERC20Capped) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Capped.abi.Unpack("eip712Domain", data)
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
func (eRC20Capped *ERC20Capped) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Capped *ERC20Capped) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("forgeByIndex", data)
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
func (eRC20Capped *ERC20Capped) PackForgeCount() []byte {
	enc, err := eRC20Capped.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("forgeCount", data)
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
func (eRC20Capped *ERC20Capped) PackForges() []byte {
	enc, err := eRC20Capped.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Capped *ERC20Capped) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("forges", data)
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
func (eRC20Capped *ERC20Capped) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20Capped *ERC20Capped) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20Capped.abi.Unpack("getRoleAdmin", data)
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
func (eRC20Capped *ERC20Capped) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20Capped *ERC20Capped) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("getRoleMember", data)
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
func (eRC20Capped *ERC20Capped) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("getRoleMemberCount", data)
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
func (eRC20Capped *ERC20Capped) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20Capped *ERC20Capped) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("getRoleMembers", data)
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
func (eRC20Capped *ERC20Capped) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Capped *ERC20Capped) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Capped *ERC20Capped) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("hasRole", data)
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
func (eRC20Capped *ERC20Capped) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Capped *ERC20Capped) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (eRC20Capped *ERC20Capped) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Capped *ERC20Capped) PackName() []byte {
	enc, err := eRC20Capped.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Capped *ERC20Capped) UnpackName(data []byte) (string, error) {
	out, err := eRC20Capped.abi.Unpack("name", data)
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
func (eRC20Capped *ERC20Capped) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("nonces", data)
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
func (eRC20Capped *ERC20Capped) PackOwner() []byte {
	enc, err := eRC20Capped.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Capped *ERC20Capped) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20Capped.abi.Unpack("owner", data)
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
func (eRC20Capped *ERC20Capped) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20Capped.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// PendingDefaultAdminOutput serves as a container for the return parameters of contract
// method PendingDefaultAdmin.
type PendingDefaultAdminOutput struct {
	NewAdmin common.Address
	Schedule *big.Int
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20Capped *ERC20Capped) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20Capped.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC20Capped *ERC20Capped) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20Capped.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PendingDefaultAdminDelayOutput serves as a container for the return parameters of contract
// method PendingDefaultAdminDelay.
type PendingDefaultAdminDelayOutput struct {
	NewDelay *big.Int
	Schedule *big.Int
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20Capped *ERC20Capped) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20Capped.abi.Unpack("pendingDefaultAdminDelay", data)
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
func (eRC20Capped *ERC20Capped) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRemainingSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20Capped *ERC20Capped) PackRemainingSupply() []byte {
	enc, err := eRC20Capped.abi.Pack("remainingSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemainingSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackRemainingSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("remainingSupply", data)
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
func (eRC20Capped *ERC20Capped) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20Capped *ERC20Capped) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Capped.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20Capped *ERC20Capped) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20Capped.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20Capped *ERC20Capped) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20Capped.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Capped *ERC20Capped) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20Capped.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Capped *ERC20Capped) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("supportsInterface", data)
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
func (eRC20Capped *ERC20Capped) PackSymbol() []byte {
	enc, err := eRC20Capped.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Capped *ERC20Capped) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Capped.abi.Unpack("symbol", data)
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
func (eRC20Capped *ERC20Capped) PackTotalSupply() []byte {
	enc, err := eRC20Capped.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Capped *ERC20Capped) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Capped.abi.Unpack("totalSupply", data)
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
func (eRC20Capped *ERC20Capped) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Capped *ERC20Capped) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("transfer", data)
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
func (eRC20Capped *ERC20Capped) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Capped.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Capped *ERC20Capped) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Capped.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// ERC20CappedApproval represents a Approval event raised by the ERC20Capped contract.
type ERC20CappedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20CappedApproval) ContractEventName() string {
	return ERC20CappedApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Capped *ERC20Capped) UnpackApprovalEvent(log *types.Log) (*ERC20CappedApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedApproval)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20Capped contract.
type ERC20CappedDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20CappedDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20Capped *ERC20Capped) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20CappedDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20Capped contract.
type ERC20CappedDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20CappedDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20Capped *ERC20Capped) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20CappedDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20Capped contract.
type ERC20CappedDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20CappedDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20Capped *ERC20Capped) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20CappedDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20Capped contract.
type ERC20CappedDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20CappedDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20CappedDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20CappedDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20Capped *ERC20Capped) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20CappedDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Capped contract.
type ERC20CappedEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20CappedEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedEIP712DomainChanged) ContractEventName() string {
	return ERC20CappedEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Capped *ERC20Capped) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20CappedEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedForgeAdded represents a ForgeAdded event raised by the ERC20Capped contract.
type ERC20CappedForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20CappedForgeAdded) ContractEventName() string {
	return ERC20CappedForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20Capped *ERC20Capped) UnpackForgeAddedEvent(log *types.Log) (*ERC20CappedForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedForgeRemoved represents a ForgeRemoved event raised by the ERC20Capped contract.
type ERC20CappedForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20CappedForgeRemoved) ContractEventName() string {
	return ERC20CappedForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20Capped *ERC20Capped) UnpackForgeRemovedEvent(log *types.Log) (*ERC20CappedForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20Capped contract.
type ERC20CappedRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20CappedRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20CappedRoleAdminChanged) ContractEventName() string {
	return ERC20CappedRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20Capped *ERC20Capped) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20CappedRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedRoleGranted represents a RoleGranted event raised by the ERC20Capped contract.
type ERC20CappedRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20CappedRoleGranted) ContractEventName() string {
	return ERC20CappedRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Capped *ERC20Capped) UnpackRoleGrantedEvent(log *types.Log) (*ERC20CappedRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedRoleRevoked represents a RoleRevoked event raised by the ERC20Capped contract.
type ERC20CappedRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20CappedRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20CappedRoleRevoked) ContractEventName() string {
	return ERC20CappedRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Capped *ERC20Capped) UnpackRoleRevokedEvent(log *types.Log) (*ERC20CappedRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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

// ERC20CappedTransfer represents a Transfer event raised by the ERC20Capped contract.
type ERC20CappedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20CappedTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20CappedTransfer) ContractEventName() string {
	return ERC20CappedTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Capped *ERC20Capped) UnpackTransferEvent(log *types.Log) (*ERC20CappedTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20Capped.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20CappedTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Capped.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Capped.abi.Events[event].Inputs {
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
func (eRC20Capped *ERC20Capped) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20CapableERC20ExceededCap"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20CapableERC20ExceededCapError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Capped.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20Capped.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20CappedAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20Capped contract.
type ERC20CappedAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20CappedAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20Capped *ERC20Capped) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20CappedAccessControlBadConfirmation, error) {
	out := new(ERC20CappedAccessControlBadConfirmation)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20Capped contract.
type ERC20CappedAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20CappedAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20Capped *ERC20Capped) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20CappedAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20CappedAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20Capped contract.
type ERC20CappedAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20CappedAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20Capped *ERC20Capped) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20CappedAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20CappedAccessControlEnforcedDefaultAdminRules)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20Capped contract.
type ERC20CappedAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20CappedAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20Capped *ERC20Capped) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20CappedAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20CappedAccessControlInvalidDefaultAdmin)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20Capped contract.
type ERC20CappedAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20CappedAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20Capped *ERC20Capped) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20CappedAccessControlUnauthorizedAccount, error) {
	out := new(ERC20CappedAccessControlUnauthorizedAccount)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Capped contract.
type ERC20CappedECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20CappedECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Capped *ERC20Capped) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20CappedECDSAInvalidSignature, error) {
	out := new(ERC20CappedECDSAInvalidSignature)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Capped contract.
type ERC20CappedECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20CappedECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Capped *ERC20Capped) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20CappedECDSAInvalidSignatureLength, error) {
	out := new(ERC20CappedECDSAInvalidSignatureLength)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Capped contract.
type ERC20CappedECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20CappedECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Capped *ERC20Capped) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20CappedECDSAInvalidSignatureS, error) {
	out := new(ERC20CappedECDSAInvalidSignatureS)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20CapableERC20ExceededCap represents a ERC20Capable__ERC20ExceededCap error raised by the ERC20Capped contract.
type ERC20CappedERC20CapableERC20ExceededCap struct {
	IncreasedSupply *big.Int
	Cap             *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func ERC20CappedERC20CapableERC20ExceededCapErrorID() common.Hash {
	return common.HexToHash("0x168a7f6ea7d992fe05deb1d140d19ff46ba8920881431f8d796b5ceade4790fa")
}

// UnpackERC20CapableERC20ExceededCapError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap)
func (eRC20Capped *ERC20Capped) UnpackERC20CapableERC20ExceededCapError(raw []byte) (*ERC20CappedERC20CapableERC20ExceededCap, error) {
	out := new(ERC20CappedERC20CapableERC20ExceededCap)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20CapableERC20ExceededCap", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Capped contract.
type ERC20CappedERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20CappedERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Capped *ERC20Capped) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20CappedERC20InsufficientAllowance, error) {
	out := new(ERC20CappedERC20InsufficientAllowance)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Capped contract.
type ERC20CappedERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20CappedERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Capped *ERC20Capped) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20CappedERC20InsufficientBalance, error) {
	out := new(ERC20CappedERC20InsufficientBalance)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20CappedERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidApproverError(raw []byte) (*ERC20CappedERC20InvalidApprover, error) {
	out := new(ERC20CappedERC20InvalidApprover)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20CappedERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20CappedERC20InvalidReceiver, error) {
	out := new(ERC20CappedERC20InvalidReceiver)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20CappedERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidSenderError(raw []byte) (*ERC20CappedERC20InvalidSender, error) {
	out := new(ERC20CappedERC20InvalidSender)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Capped contract.
type ERC20CappedERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20CappedERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Capped *ERC20Capped) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20CappedERC20InvalidSpender, error) {
	out := new(ERC20CappedERC20InvalidSpender)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Capped contract.
type ERC20CappedERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20CappedERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Capped *ERC20Capped) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20CappedERC2612ExpiredSignature, error) {
	out := new(ERC20CappedERC2612ExpiredSignature)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Capped contract.
type ERC20CappedERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20CappedERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Capped *ERC20Capped) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20CappedERC2612InvalidSigner, error) {
	out := new(ERC20CappedERC2612InvalidSigner)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Capped contract.
type ERC20CappedInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20CappedInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Capped *ERC20Capped) UnpackInvalidAccountNonceError(raw []byte) (*ERC20CappedInvalidAccountNonce, error) {
	out := new(ERC20CappedInvalidAccountNonce)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedInvalidShortString represents a InvalidShortString error raised by the ERC20Capped contract.
type ERC20CappedInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20CappedInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20Capped *ERC20Capped) UnpackInvalidShortStringError(raw []byte) (*ERC20CappedInvalidShortString, error) {
	out := new(ERC20CappedInvalidShortString)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20Capped contract.
type ERC20CappedSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20CappedSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20Capped *ERC20Capped) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20CappedSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20CappedSafeCastOverflowedUintDowncast)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedStringTooLong represents a StringTooLong error raised by the ERC20Capped contract.
type ERC20CappedStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20CappedStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20Capped *ERC20Capped) UnpackStringTooLongError(raw []byte) (*ERC20CappedStringTooLong, error) {
	out := new(ERC20CappedStringTooLong)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20Capped contract.
type ERC20CappedTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20CappedTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20Capped *ERC20Capped) UnpackTokenBaseNullInputError(raw []byte) (*ERC20CappedTokenBaseNullInput, error) {
	out := new(ERC20CappedTokenBaseNullInput)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20CappedTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20Capped contract.
type ERC20CappedTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20CappedTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20Capped *ERC20Capped) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20CappedTokenBaseOnlyForge, error) {
	out := new(ERC20CappedTokenBaseOnlyForge)
	if err := eRC20Capped.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
