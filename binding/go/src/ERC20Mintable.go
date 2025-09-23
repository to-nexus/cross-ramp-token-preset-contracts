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

// ERC20MintableMetaData contains all meta data concerning the ERC20Mintable contract.
var ERC20MintableMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC20Mintable",
	Bin: "0x610180604052348015610010575f5ffd5b5060405161385138038061385183398101604081905261002f916105dc565b84848484848280604051806040016040528060018152602001603160f81b815250858589898162015180815f6001600160a01b0316816001600160a01b03160361009357604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100bc5f82610270565b505050506100f07faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c8361027060201b60201c565b5080515f5b8181101561012f57610127600484838151811061011457610114610702565b602002602001015161028460201b60201c565b6001016100f5565b505050508160099081610142919061079a565b50600a61014f828261079a565b5061015f9150839050600b6102fc565b6101205261016e81600c6102fc565b61014052815160208084019190912060e052815190820120610100524660a0526101fa60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0525082515f0361022e5760405163e1dea2ef60e01b8152636e616d6560e01b600482015260240161008a565b81515f0361025a5760405163e1dea2ef60e01b8152651cde5b589bdb60d21b600482015260240161008a565b60ff1661016052506108ac975050505050505050565b5f61027b838361032c565b90505b92915050565b6001600160a01b0381166102b55760405163e1dea2ef60e01b815264666f72676560d81b600482015260240161008a565b6102bf828261035f565b156102f8576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b5f6020835110156103175761031083610373565b905061027e565b81610322848261079a565b5060ff905061027e565b5f8061033884846103b0565b9050801561027b575f848152600360205260409020610357908461035f565b509392505050565b5f61027b836001600160a01b038416610416565b5f5f829050601f8151111561039d578260405163305a27a960e01b815260040161008a9190610854565b80516103a882610889565b179392505050565b5f8261040c575f6103c96002546001600160a01b031690565b6001600160a01b0316146103f057604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b61027b8383610462565b5f81815260018301602052604081205461045b57508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561027e565b505f61027e565b5f828152602081815260408083206001600160a01b038516845290915281205460ff1661045b575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556104ba3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161027e565b80516001600160a01b0381168114610518575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156105595761055961051d565b604052919050565b5f82601f830112610570575f5ffd5b81516001600160401b038111156105895761058961051d565b61059c601f8201601f1916602001610531565b8181528460208386010111156105b0575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b805160ff81168114610518575f5ffd5b5f5f5f5f5f60a086880312156105f0575f5ffd5b6105f986610502565b60208701519095506001600160401b03811115610614575f5ffd5b8601601f81018813610624575f5ffd5b80516001600160401b0381111561063d5761063d61051d565b8060051b61064d60208201610531565b9182526020818401810192908101908b841115610668575f5ffd5b6020850194505b838510156106915761068085610502565b82526020948501949091019061066f565b60408b0151909850935050506001600160401b0382111590506106b2575f5ffd5b6106be88828901610561565b606088015190945090506001600160401b038111156106db575f5ffd5b6106e788828901610561565b9250506106f6608087016105cc565b90509295509295909350565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061072a57607f821691505b60208210810361074857634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561079557805f5260205f20601f840160051c810160208510156107735750805b601f840160051c820191505b81811015610792575f815560010161077f565b50505b505050565b81516001600160401b038111156107b3576107b361051d565b6107c7816107c18454610716565b8461074e565b6020601f8211600181146107f9575f83156107e25750848201515b5f19600385901b1c1916600184901b178455610792565b5f84815260208120601f198516915b828110156108285787850151825560209485019460019092019101610808565b508482101561084557868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610748575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051612f4a6109075f395f6103ac01525f61148301525f61145601525f61111201525f6110ea01525f61104501525f61106f01525f6110990152612f4a5ff3fe608060405234801561000f575f5ffd5b50600436106102cd575f3560e01c806387f453531161017c578063a9059cbb116100dd578063d505accf11610093578063dd62ed3e1161006e578063dd62ed3e14610657578063ec87621c1461066a578063fe2df3e814610691575f5ffd5b8063d505accf14610629578063d547741f1461063c578063d602b9fd1461064f575f5ffd5b8063cc8463c8116100c3578063cc8463c8146105cd578063cefc1429146105d5578063cf6eefb7146105dd575f5ffd5b8063a9059cbb146105a7578063ca15c873146105ba575f5ffd5b80639ca92df911610132578063a217fddf11610118578063a217fddf14610585578063a3246ad31461058c578063a40283d51461059f575f5ffd5b80639ca92df91461054b578063a1eda53c1461055e575f5ffd5b80639010d07c116101625780639010d07c146104ed57806391d148541461050057806395d89b4114610543575f5ffd5b806387f45353146104d25780638da5cb5b146104e5575f5ffd5b80633644e51511610231578063649a5ec7116101e75780637ecebe00116101c25780637ecebe001461046557806384b0196e1461047857806384ef8ffc14610493575f5ffd5b8063649a5ec71461042c57806370a082311461043f57806379cc679014610452575f5ffd5b806340c10f191161021757806340c10f19146103f15780635c4e62c414610404578063634e93da14610419575f5ffd5b80633644e515146103d657806336568abe146103de575f5ffd5b806318160ddd11610286578063248a9ca31161026c578063248a9ca3146103705780632f2ff15d14610392578063313ce567146103a5575f5ffd5b806318160ddd1461034757806323b872dd1461035d575f5ffd5b806306fdde03116102b657806306fdde0314610315578063095ea7b31461032a5780630aa6220b1461033d575f5ffd5b806301ffc9a7146102d1578063022d63fb146102f9575b5f5ffd5b6102e46102df366004612988565b6106a4565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016102f0565b61031d6108a6565b6040516102f09190612a13565b6102e4610338366004612a4d565b6108b5565b6103456108c7565b005b61034f6108dc565b6040519081526020016102f0565b6102e461036b366004612a75565b6108e6565b61034f61037e366004612aaf565b5f9081526020819052604090206001015490565b6103456103a0366004612ac6565b6108fa565b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016102f0565b61034f610908565b6103456103ec366004612ac6565b610911565b6103456103ff366004612a4d565b61091b565b61040c61096b565b6040516102f09190612af0565b610345610427366004612b48565b610977565b61034561043a366004612b61565b61098a565b61034f61044d366004612b48565b61099d565b610345610460366004612a4d565b6109c7565b61034f610473366004612b48565b6109f9565b610480610a03565b6040516102f09796959493929190612b86565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102f0565b6102e46104e0366004612b48565b610a61565b6104ad610a6d565b6104ad6104fb366004612c45565b610a8d565b6102e461050e366004612ac6565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61031d610aa4565b6104ad610559366004612aaf565b610aae565b610566610aba565b6040805165ffffffffffff9384168152929091166020830152016102f0565b61034f5f81565b61040c61059a366004612aaf565b610b34565b61034f610b4d565b6102e46105b5366004612a4d565b610b58565b61034f6105c8366004612aaf565b610b63565b6102fe610b79565b610345610c16565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102f0565b610345610637366004612c65565b610c72565b61034561064a366004612ac6565b610e1b565b610345610e25565b61034f610665366004612cd2565b610e37565b61034f7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61034561069f366004612cfa565b610e70565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061073657507fffffffff0000000000000000000000000000000000000000000000000000000082167f390d688900000000000000000000000000000000000000000000000000000000145b8061078257507fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000145b806107ce57507fffffffff0000000000000000000000000000000000000000000000000000000082167fa219a02500000000000000000000000000000000000000000000000000000000145b806107f957507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061084557507fffffffff0000000000000000000000000000000000000000000000000000000082167f9d8ff7da00000000000000000000000000000000000000000000000000000000145b8061089157507fffffffff0000000000000000000000000000000000000000000000000000000082167f84b0196e00000000000000000000000000000000000000000000000000000000145b806108a057506108a082610f01565b92915050565b60606108b0610f0b565b905090565b5f6108c08383610f9b565b9392505050565b5f6108d181610fb2565b6108d9610fbc565b50565b5f6108b060085490565b5f6108f2848484610fc8565b949350505050565b6109048282610feb565b5050565b5f6108b061102c565b6109048282611162565b61092433610a61565b610961576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6109048282611267565b60606108b060046112c1565b5f61098181610fb2565b610904826112cd565b5f61099481610fb2565b6109048261134c565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600660205260408120546108a0565b73ffffffffffffffffffffffffffffffffffffffff821633146109ef576109ef8233836113bb565b61090482826113cb565b5f6108a082611425565b5f6060805f5f5f6060610a1461144f565b610a1c61147c565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f6108a06004836114a9565b5f6108b060025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206108c090836114d7565b60606108b06114e2565b5f6108a06004836114d7565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610afc57504265ffffffffffff821610155b610b07575f5f610b2c565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b5f8181526003602052604090206060906108a0906112c1565b5f6108b060046114f1565b5f6108c083836114fa565b5f8181526003602052604081206108a0906114f1565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610bba57504265ffffffffffff8216105b610bec576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610c10565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610c6a576040517fc22c8022000000000000000000000000000000000000000000000000000000008152336004820152602401610958565b6108d9611507565b83421115610caf576040517f6279130200000000000000000000000000000000000000000000000000000000815260048101859052602401610958565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610d078c73ffffffffffffffffffffffffffffffffffffffff165f908152600d6020526040902080546001810190915590565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610d6e826115f8565b90505f610d7d8287878761163f565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610e04576040517f4b800e4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528b166024820152604401610958565b610e0f8a8a8a61166b565b50505050505050505050565b6109048282611678565b5f610e2f81610fb2565b6108d96116b9565b73ffffffffffffffffffffffffffffffffffffffff8083165f9081526007602090815260408083209385168352929052908120546108c0565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610e9a81610fb2565b61298082610eaa576116c3610eae565b6117165b9050835f5b81811015610ef857610ef06004888884818110610ed257610ed2612d80565b9050602002016020810190610ee79190612b48565b8563ffffffff16565b600101610eb3565b50505050505050565b5f6108a0826117d8565b606060098054610f1a90612dad565b80601f0160208091040260200160405190810160405280929190818152602001828054610f4690612dad565b8015610f915780601f10610f6857610100808354040283529160200191610f91565b820191905f5260205f20905b815481529060010190602001808311610f7457829003601f168201915b5050505050905090565b5f33610fa881858561166b565b5060019392505050565b6108d981336117e2565b610fc65f5f611867565b565b5f33610fd58582856113bb565b610fe08585856119c0565b506001949350505050565b81611022576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109048282611a69565b5f3073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561109157507f000000000000000000000000000000000000000000000000000000000000000046145b156110bb57507f000000000000000000000000000000000000000000000000000000000000000090565b6108b0604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b8115801561118a575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b1561125d5760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff16811515806111de575065ffffffffffff8116155b806111f157504265ffffffffffff821610155b15611232576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610958565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6109048282611a93565b73ffffffffffffffffffffffffffffffffffffffff82166112b6576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610958565b6109045f8383611aec565b60605f6108c083611af7565b5f6112d6610b79565b6112df42611b50565b6112e99190612e2b565b90506112f58282611b9f565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61135682611c3a565b61135f42611b50565b6113699190612e2b565b90506113758282611867565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b6113c6838383611c81565b505050565b73ffffffffffffffffffffffffffffffffffffffff821661141a576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610958565b610904825f83611aec565b73ffffffffffffffffffffffffffffffffffffffff81165f908152600d60205260408120546108a0565b60606108b07f0000000000000000000000000000000000000000000000000000000000000000600b611d24565b60606108b07f0000000000000000000000000000000000000000000000000000000000000000600c611d24565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260018301602052604081205415156108c0565b5f6108c08383611dcd565b6060600a8054610f1a90612dad565b5f6108a0825490565b5f33610fa88185856119c0565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061155757504265ffffffffffff821610155b15611598576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff82166004820152602401610958565b6115c05f6115bb60025473ffffffffffffffffffffffffffffffffffffffff1690565b611df3565b506115cb5f83611dfe565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b5f6108a061160461102c565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f5f5f5f61164f88888888611e09565b92509250925061165f8282611efc565b50909695505050505050565b6113c68383836001611fff565b816116af576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610904828261200b565b610fc65f5f611b9f565b6116cd828261202f565b156109045760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611785576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f7267650000000000000000000000000000000000000000000000000000006004820152602401610958565b61178f8282612050565b156109045760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f6108a082612071565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610904576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260248101839052604401610958565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801561193b574265ffffffffffff82161015611912576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a0100000000000000000000000000000000000000000000000000000291909117905561193b565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b73ffffffffffffffffffffffffffffffffffffffff8316611a0f576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610958565b73ffffffffffffffffffffffffffffffffffffffff8216611a5e576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610958565b6113c6838383611aec565b5f82815260208190526040902060010154611a8381610fb2565b611a8d8383611dfe565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611ae2576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113c68282611df3565b6113c68383836120c6565b6060815f01805480602002602001604051908101604052809291908181526020018280548015611b4457602002820191905f5260205f20905b815481526020019060010190808311611b30575b50505050509050919050565b5f65ffffffffffff821115611b9b576040517f6dfcc6500000000000000000000000000000000000000000000000000000000081526030600482015260248101839052604401610958565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff8816171790935590041680156113c6576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611c44610b79565b90508065ffffffffffff168365ffffffffffff1611611c6c57611c678382612e49565b6108c0565b6108c065ffffffffffff84166206978061226d565b5f611c8c8484610e37565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015611a8d5781811015611d16576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610958565b611a8d84848484035f611fff565b606060ff8314611d3e57611d378361227c565b90506108a0565b818054611d4a90612dad565b80601f0160208091040260200160405190810160405280929190818152602001828054611d7690612dad565b8015611dc15780601f10611d9857610100808354040283529160200191611dc1565b820191905f5260205f20905b815481529060010190602001808311611da457829003601f168201915b505050505090506108a0565b5f825f018281548110611de257611de2612d80565b905f5260205f200154905092915050565b5f6108c083836122b9565b5f6108c083836122ec565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611e4257505f91506003905082611ef2565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611e93573d5f5f3e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611ee957505f925060019150829050611ef2565b92505f91508190505b9450945094915050565b5f826003811115611f0f57611f0f612e67565b03611f18575050565b6001826003811115611f2c57611f2c612e67565b03611f63576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115611f7757611f77612e67565b03611fb1576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610958565b6003826003811115611fc557611fc5612e67565b03610904576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610958565b611a8d84848484612317565b5f8281526020819052604090206001015461202581610fb2565b611a8d8383611df3565b5f6108c08373ffffffffffffffffffffffffffffffffffffffff841661245c565b5f6108c08373ffffffffffffffffffffffffffffffffffffffff841661253f565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806108a057506108a08261258b565b73ffffffffffffffffffffffffffffffffffffffff83166120fd578060085f8282546120f29190612e94565b909155506121ad9050565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526006602052604090205481811015612182576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610958565b73ffffffffffffffffffffffffffffffffffffffff84165f9081526006602052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166121d657600880548290039055612201565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526006602052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161226091815260200190565b60405180910390a3505050565b5f8282188284100282186108c0565b60605f612288836125e0565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f5f6122c58484612620565b905080156108c0575f8481526003602052604090206122e4908461202f565b509392505050565b5f5f6122f88484612681565b905080156108c0575f8481526003602052604090206122e49084612050565b73ffffffffffffffffffffffffffffffffffffffff8416612366576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610958565b73ffffffffffffffffffffffffffffffffffffffff83166123b5576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610958565b73ffffffffffffffffffffffffffffffffffffffff8085165f9081526007602090815260408083209387168352929052208290558015611a8d578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161244e91815260200190565b60405180910390a350505050565b5f8181526001830160205260408120548015612536575f61247e600183612ea7565b85549091505f9061249190600190612ea7565b90508082146124f0575f865f0182815481106124af576124af612d80565b905f5260205f200154905080875f0184815481106124cf576124cf612d80565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061250157612501612eba565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506108a0565b5f9150506108a0565b5f81815260018301602052604081205461258457508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556108a0565b505f6108a0565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806108a057506108a08261273f565b5f60ff8216601f8111156108a0576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f82158015612649575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561267757600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6108c083836127d5565b5f82612735575f6126a760025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16146126f4576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6108c0838361288e565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806108a057507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146108a0565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615612584575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016108a0565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16612584575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561291e3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016108a0565b610fc6612ee7565b5f60208284031215612998575f5ffd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146108c0575f5ffd5b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6108c060208301846129c7565b803573ffffffffffffffffffffffffffffffffffffffff81168114612a48575f5ffd5b919050565b5f5f60408385031215612a5e575f5ffd5b612a6783612a25565b946020939093013593505050565b5f5f5f60608486031215612a87575f5ffd5b612a9084612a25565b9250612a9e60208501612a25565b929592945050506040919091013590565b5f60208284031215612abf575f5ffd5b5035919050565b5f5f60408385031215612ad7575f5ffd5b82359150612ae760208401612a25565b90509250929050565b602080825282518282018190525f918401906040840190835b81811015612b3d57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612b09565b509095945050505050565b5f60208284031215612b58575f5ffd5b6108c082612a25565b5f60208284031215612b71575f5ffd5b813565ffffffffffff811681146108c0575f5ffd5b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f612bc060e08301896129c7565b8281036040840152612bd281896129c7565b6060840188905273ffffffffffffffffffffffffffffffffffffffff8716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015612c34578351835260209384019390920191600101612c16565b50909b9a5050505050505050505050565b5f5f60408385031215612c56575f5ffd5b50508035926020909101359150565b5f5f5f5f5f5f5f60e0888a031215612c7b575f5ffd5b612c8488612a25565b9650612c9260208901612a25565b95506040880135945060608801359350608088013560ff81168114612cb5575f5ffd5b9699959850939692959460a0840135945060c09093013592915050565b5f5f60408385031215612ce3575f5ffd5b612cec83612a25565b9150612ae760208401612a25565b5f5f5f60408486031215612d0c575f5ffd5b833567ffffffffffffffff811115612d22575f5ffd5b8401601f81018613612d32575f5ffd5b803567ffffffffffffffff811115612d48575f5ffd5b8660208260051b8401011115612d5c575f5ffd5b6020918201945092508401358015158114612d75575f5ffd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c90821680612dc157607f821691505b602082108103612df8577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff81811683821601908111156108a0576108a0612dfe565b65ffffffffffff82811682821603908111156108a0576108a0612dfe565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b808201808211156108a0576108a0612dfe565b818103818111156108a0576108a0612dfe565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220320ad15919beb71ddf0c9d525b168380addfd2995fcf605d324342333e44e88e64736f6c634300081c0033",
}

// ERC20Mintable is an auto generated Go binding around an Ethereum contract.
type ERC20Mintable struct {
	abi abi.ABI
}

// NewERC20Mintable creates a new instance of ERC20Mintable.
func NewERC20Mintable() *ERC20Mintable {
	parsed, err := ERC20MintableMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC20Mintable{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC20Mintable) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, uint8 decimals) returns()
func (eRC20Mintable *ERC20Mintable) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, decimals uint8) []byte {
	enc, err := eRC20Mintable.abi.Pack("", owner, forges, name, symbol, decimals)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC20Mintable.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Mintable.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC20Mintable *ERC20Mintable) PackDOMAINSEPARATOR() []byte {
	enc, err := eRC20Mintable.abi.Pack("DOMAIN_SEPARATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDOMAINSEPARATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) UnpackDOMAINSEPARATOR(data []byte) ([32]byte, error) {
	out, err := eRC20Mintable.abi.Unpack("DOMAIN_SEPARATOR", data)
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
func (eRC20Mintable *ERC20Mintable) PackMANAGERROLE() []byte {
	enc, err := eRC20Mintable.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC20Mintable.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC20Mintable *ERC20Mintable) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC20Mintable.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("allowance", data)
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
func (eRC20Mintable *ERC20Mintable) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackApprove(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("approve", data)
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
func (eRC20Mintable *ERC20Mintable) PackBalanceOf(account common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("balanceOf", data)
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
func (eRC20Mintable *ERC20Mintable) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 amount) returns()
func (eRC20Mintable *ERC20Mintable) PackBurnFrom(from common.Address, amount *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("burnFrom", from, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC20Mintable *ERC20Mintable) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC20Mintable.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC20Mintable *ERC20Mintable) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Mintable *ERC20Mintable) PackDecimals() []byte {
	enc, err := eRC20Mintable.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (eRC20Mintable *ERC20Mintable) UnpackDecimals(data []byte) (uint8, error) {
	out, err := eRC20Mintable.abi.Unpack("decimals", data)
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
func (eRC20Mintable *ERC20Mintable) PackDefaultAdmin() []byte {
	enc, err := eRC20Mintable.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("defaultAdmin", data)
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
func (eRC20Mintable *ERC20Mintable) PackDefaultAdminDelay() []byte {
	enc, err := eRC20Mintable.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("defaultAdminDelay", data)
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
func (eRC20Mintable *ERC20Mintable) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC20Mintable.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC20Mintable *ERC20Mintable) PackEip712Domain() []byte {
	enc, err := eRC20Mintable.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (eRC20Mintable *ERC20Mintable) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := eRC20Mintable.abi.Unpack("eip712Domain", data)
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
func (eRC20Mintable *ERC20Mintable) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC20Mintable *ERC20Mintable) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("forgeByIndex", data)
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
func (eRC20Mintable *ERC20Mintable) PackForgeCount() []byte {
	enc, err := eRC20Mintable.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("forgeCount", data)
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
func (eRC20Mintable *ERC20Mintable) PackForges() []byte {
	enc, err := eRC20Mintable.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC20Mintable *ERC20Mintable) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("forges", data)
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
func (eRC20Mintable *ERC20Mintable) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC20Mintable *ERC20Mintable) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC20Mintable.abi.Unpack("getRoleAdmin", data)
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
func (eRC20Mintable *ERC20Mintable) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC20Mintable *ERC20Mintable) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("getRoleMember", data)
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
func (eRC20Mintable *ERC20Mintable) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("getRoleMemberCount", data)
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
func (eRC20Mintable *ERC20Mintable) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC20Mintable *ERC20Mintable) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("getRoleMembers", data)
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
func (eRC20Mintable *ERC20Mintable) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Mintable *ERC20Mintable) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("hasRole", data)
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
func (eRC20Mintable *ERC20Mintable) PackIsForge(forge common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("isForge", data)
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
func (eRC20Mintable *ERC20Mintable) PackMint(to common.Address, amount *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("mint", to, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Mintable *ERC20Mintable) PackName() []byte {
	enc, err := eRC20Mintable.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC20Mintable *ERC20Mintable) UnpackName(data []byte) (string, error) {
	out, err := eRC20Mintable.abi.Unpack("name", data)
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
func (eRC20Mintable *ERC20Mintable) PackNonces(owner common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("nonces", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNonces is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackNonces(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("nonces", data)
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
func (eRC20Mintable *ERC20Mintable) PackOwner() []byte {
	enc, err := eRC20Mintable.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC20Mintable *ERC20Mintable) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC20Mintable.abi.Unpack("owner", data)
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
func (eRC20Mintable *ERC20Mintable) PackPendingDefaultAdmin() []byte {
	enc, err := eRC20Mintable.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC20Mintable *ERC20Mintable) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC20Mintable.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC20Mintable *ERC20Mintable) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC20Mintable.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC20Mintable *ERC20Mintable) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC20Mintable.abi.Unpack("pendingDefaultAdminDelay", data)
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
func (eRC20Mintable *ERC20Mintable) PackPermit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("permit", owner, spender, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (eRC20Mintable *ERC20Mintable) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC20Mintable *ERC20Mintable) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC20Mintable.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC20Mintable *ERC20Mintable) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC20Mintable.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC20Mintable *ERC20Mintable) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC20Mintable.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Mintable *ERC20Mintable) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC20Mintable.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("supportsInterface", data)
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
func (eRC20Mintable *ERC20Mintable) PackSymbol() []byte {
	enc, err := eRC20Mintable.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC20Mintable *ERC20Mintable) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC20Mintable.abi.Unpack("symbol", data)
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
func (eRC20Mintable *ERC20Mintable) PackTotalSupply() []byte {
	enc, err := eRC20Mintable.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (eRC20Mintable *ERC20Mintable) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := eRC20Mintable.abi.Unpack("totalSupply", data)
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
func (eRC20Mintable *ERC20Mintable) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackTransfer(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("transfer", data)
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
func (eRC20Mintable *ERC20Mintable) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := eRC20Mintable.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (eRC20Mintable *ERC20Mintable) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := eRC20Mintable.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// ERC20MintableApproval represents a Approval event raised by the ERC20Mintable contract.
type ERC20MintableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintableApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC20MintableApproval) ContractEventName() string {
	return ERC20MintableApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (eRC20Mintable *ERC20Mintable) UnpackApprovalEvent(log *types.Log) (*ERC20MintableApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableApproval)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC20Mintable contract.
type ERC20MintableDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintableDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20MintableDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC20MintableDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC20MintableDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC20Mintable contract.
type ERC20MintableDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MintableDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20MintableDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC20MintableDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC20MintableDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC20Mintable contract.
type ERC20MintableDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintableDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC20MintableDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC20MintableDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC20MintableDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC20Mintable contract.
type ERC20MintableDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC20MintableDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC20MintableDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC20MintableDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC20Mintable *ERC20Mintable) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC20MintableDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableEIP712DomainChanged represents a EIP712DomainChanged event raised by the ERC20Mintable contract.
type ERC20MintableEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC20MintableEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MintableEIP712DomainChanged) ContractEventName() string {
	return ERC20MintableEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (eRC20Mintable *ERC20Mintable) UnpackEIP712DomainChangedEvent(log *types.Log) (*ERC20MintableEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableForgeAdded represents a ForgeAdded event raised by the ERC20Mintable contract.
type ERC20MintableForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintableForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC20MintableForgeAdded) ContractEventName() string {
	return ERC20MintableForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC20Mintable *ERC20Mintable) UnpackForgeAddedEvent(log *types.Log) (*ERC20MintableForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableForgeRemoved represents a ForgeRemoved event raised by the ERC20Mintable contract.
type ERC20MintableForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintableForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC20MintableForgeRemoved) ContractEventName() string {
	return ERC20MintableForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC20Mintable *ERC20Mintable) UnpackForgeRemovedEvent(log *types.Log) (*ERC20MintableForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableRoleAdminChanged represents a RoleAdminChanged event raised by the ERC20Mintable contract.
type ERC20MintableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC20MintableRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC20MintableRoleAdminChanged) ContractEventName() string {
	return ERC20MintableRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC20Mintable *ERC20Mintable) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC20MintableRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableRoleGranted represents a RoleGranted event raised by the ERC20Mintable contract.
type ERC20MintableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintableRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC20MintableRoleGranted) ContractEventName() string {
	return ERC20MintableRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Mintable *ERC20Mintable) UnpackRoleGrantedEvent(log *types.Log) (*ERC20MintableRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableRoleRevoked represents a RoleRevoked event raised by the ERC20Mintable contract.
type ERC20MintableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC20MintableRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC20MintableRoleRevoked) ContractEventName() string {
	return ERC20MintableRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC20Mintable *ERC20Mintable) UnpackRoleRevokedEvent(log *types.Log) (*ERC20MintableRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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

// ERC20MintableTransfer represents a Transfer event raised by the ERC20Mintable contract.
type ERC20MintableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC20MintableTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC20MintableTransfer) ContractEventName() string {
	return ERC20MintableTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (eRC20Mintable *ERC20Mintable) UnpackTransferEvent(log *types.Log) (*ERC20MintableTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC20Mintable.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC20MintableTransfer)
	if len(log.Data) > 0 {
		if err := eRC20Mintable.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC20Mintable.abi.Events[event].Inputs {
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
func (eRC20Mintable *ERC20Mintable) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InsufficientAllowance"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InsufficientAllowanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidApprover"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidSender"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC20InvalidSpender"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC20InvalidSpenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC2612ExpiredSignature"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC2612ExpiredSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["ERC2612InvalidSigner"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackERC2612InvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["InvalidAccountNonce"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackInvalidAccountNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC20Mintable.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC20Mintable.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC20MintableAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC20Mintable contract.
type ERC20MintableAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC20MintableAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC20Mintable *ERC20Mintable) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC20MintableAccessControlBadConfirmation, error) {
	out := new(ERC20MintableAccessControlBadConfirmation)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC20Mintable contract.
type ERC20MintableAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC20MintableAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC20Mintable *ERC20Mintable) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC20MintableAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC20MintableAccessControlEnforcedDefaultAdminDelay)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC20Mintable contract.
type ERC20MintableAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC20MintableAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC20Mintable *ERC20Mintable) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC20MintableAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC20MintableAccessControlEnforcedDefaultAdminRules)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC20Mintable contract.
type ERC20MintableAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC20MintableAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC20Mintable *ERC20Mintable) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC20MintableAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC20MintableAccessControlInvalidDefaultAdmin)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC20Mintable contract.
type ERC20MintableAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC20MintableAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC20Mintable *ERC20Mintable) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC20MintableAccessControlUnauthorizedAccount, error) {
	out := new(ERC20MintableAccessControlUnauthorizedAccount)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the ERC20Mintable contract.
type ERC20MintableECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func ERC20MintableECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (eRC20Mintable *ERC20Mintable) UnpackECDSAInvalidSignatureError(raw []byte) (*ERC20MintableECDSAInvalidSignature, error) {
	out := new(ERC20MintableECDSAInvalidSignature)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the ERC20Mintable contract.
type ERC20MintableECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func ERC20MintableECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (eRC20Mintable *ERC20Mintable) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*ERC20MintableECDSAInvalidSignatureLength, error) {
	out := new(ERC20MintableECDSAInvalidSignatureLength)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the ERC20Mintable contract.
type ERC20MintableECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func ERC20MintableECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (eRC20Mintable *ERC20Mintable) UnpackECDSAInvalidSignatureSError(raw []byte) (*ERC20MintableECDSAInvalidSignatureS, error) {
	out := new(ERC20MintableECDSAInvalidSignatureS)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InsufficientAllowance represents a ERC20InsufficientAllowance error raised by the ERC20Mintable contract.
type ERC20MintableERC20InsufficientAllowance struct {
	Spender   common.Address
	Allowance *big.Int
	Needed    *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func ERC20MintableERC20InsufficientAllowanceErrorID() common.Hash {
	return common.HexToHash("0xfb8f41b23e99d2101d86da76cdfa87dd51c82ed07d3cb62cbc473e469dbc75c3")
}

// UnpackERC20InsufficientAllowanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InsufficientAllowanceError(raw []byte) (*ERC20MintableERC20InsufficientAllowance, error) {
	out := new(ERC20MintableERC20InsufficientAllowance)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InsufficientAllowance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InsufficientBalance represents a ERC20InsufficientBalance error raised by the ERC20Mintable contract.
type ERC20MintableERC20InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func ERC20MintableERC20InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xe450d38cd8d9f7d95077d567d60ed49c7254716e6ad08fc9872816c97e0ffec6")
}

// UnpackERC20InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InsufficientBalanceError(raw []byte) (*ERC20MintableERC20InsufficientBalance, error) {
	out := new(ERC20MintableERC20InsufficientBalance)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidApprover represents a ERC20InvalidApprover error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidApprover(address approver)
func ERC20MintableERC20InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xe602df05cc75712490294c6c104ab7c17f4030363910a7a2626411c6d3118847")
}

// UnpackERC20InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidApprover(address approver)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidApproverError(raw []byte) (*ERC20MintableERC20InvalidApprover, error) {
	out := new(ERC20MintableERC20InvalidApprover)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidReceiver represents a ERC20InvalidReceiver error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func ERC20MintableERC20InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0xec442f055133b72f3b2f9f0bb351c406b178527de2040a7d1feb4e058771f613")
}

// UnpackERC20InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidReceiver(address receiver)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidReceiverError(raw []byte) (*ERC20MintableERC20InvalidReceiver, error) {
	out := new(ERC20MintableERC20InvalidReceiver)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidSender represents a ERC20InvalidSender error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSender(address sender)
func ERC20MintableERC20InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x96c6fd1edd0cd6ef7ff0ecc0facdf53148dc0048b57fe58af65755250a7a96bd")
}

// UnpackERC20InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSender(address sender)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidSenderError(raw []byte) (*ERC20MintableERC20InvalidSender, error) {
	out := new(ERC20MintableERC20InvalidSender)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC20InvalidSpender represents a ERC20InvalidSpender error raised by the ERC20Mintable contract.
type ERC20MintableERC20InvalidSpender struct {
	Spender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC20InvalidSpender(address spender)
func ERC20MintableERC20InvalidSpenderErrorID() common.Hash {
	return common.HexToHash("0x94280d62c347d8d9f4d59a76ea321452406db88df38e0c9da304f58b57b373a2")
}

// UnpackERC20InvalidSpenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC20InvalidSpender(address spender)
func (eRC20Mintable *ERC20Mintable) UnpackERC20InvalidSpenderError(raw []byte) (*ERC20MintableERC20InvalidSpender, error) {
	out := new(ERC20MintableERC20InvalidSpender)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC20InvalidSpender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC2612ExpiredSignature represents a ERC2612ExpiredSignature error raised by the ERC20Mintable contract.
type ERC20MintableERC2612ExpiredSignature struct {
	Deadline *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func ERC20MintableERC2612ExpiredSignatureErrorID() common.Hash {
	return common.HexToHash("0x627913023c184eaad13735d5a3d2657ae76ec9a872a70e0fc57522ef1a114d58")
}

// UnpackERC2612ExpiredSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612ExpiredSignature(uint256 deadline)
func (eRC20Mintable *ERC20Mintable) UnpackERC2612ExpiredSignatureError(raw []byte) (*ERC20MintableERC2612ExpiredSignature, error) {
	out := new(ERC20MintableERC2612ExpiredSignature)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC2612ExpiredSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableERC2612InvalidSigner represents a ERC2612InvalidSigner error raised by the ERC20Mintable contract.
type ERC20MintableERC2612InvalidSigner struct {
	Signer common.Address
	Owner  common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func ERC20MintableERC2612InvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x4b800e463b323b1d856edf9dec70329a639d13874a57f5c28219ff57128756db")
}

// UnpackERC2612InvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC2612InvalidSigner(address signer, address owner)
func (eRC20Mintable *ERC20Mintable) UnpackERC2612InvalidSignerError(raw []byte) (*ERC20MintableERC2612InvalidSigner, error) {
	out := new(ERC20MintableERC2612InvalidSigner)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "ERC2612InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableInvalidAccountNonce represents a InvalidAccountNonce error raised by the ERC20Mintable contract.
type ERC20MintableInvalidAccountNonce struct {
	Account      common.Address
	CurrentNonce *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func ERC20MintableInvalidAccountNonceErrorID() common.Hash {
	return common.HexToHash("0x752d88c0de02638abf10e8e31861e4c68dc1f3a1e7d840e580683f2c282bfc7a")
}

// UnpackInvalidAccountNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAccountNonce(address account, uint256 currentNonce)
func (eRC20Mintable *ERC20Mintable) UnpackInvalidAccountNonceError(raw []byte) (*ERC20MintableInvalidAccountNonce, error) {
	out := new(ERC20MintableInvalidAccountNonce)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "InvalidAccountNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableInvalidShortString represents a InvalidShortString error raised by the ERC20Mintable contract.
type ERC20MintableInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func ERC20MintableInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (eRC20Mintable *ERC20Mintable) UnpackInvalidShortStringError(raw []byte) (*ERC20MintableInvalidShortString, error) {
	out := new(ERC20MintableInvalidShortString)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC20Mintable contract.
type ERC20MintableSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC20MintableSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC20Mintable *ERC20Mintable) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC20MintableSafeCastOverflowedUintDowncast, error) {
	out := new(ERC20MintableSafeCastOverflowedUintDowncast)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableStringTooLong represents a StringTooLong error raised by the ERC20Mintable contract.
type ERC20MintableStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func ERC20MintableStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (eRC20Mintable *ERC20Mintable) UnpackStringTooLongError(raw []byte) (*ERC20MintableStringTooLong, error) {
	out := new(ERC20MintableStringTooLong)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC20Mintable contract.
type ERC20MintableTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC20MintableTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC20Mintable *ERC20Mintable) UnpackTokenBaseNullInputError(raw []byte) (*ERC20MintableTokenBaseNullInput, error) {
	out := new(ERC20MintableTokenBaseNullInput)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC20MintableTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC20Mintable contract.
type ERC20MintableTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC20MintableTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC20Mintable *ERC20Mintable) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC20MintableTokenBaseOnlyForge, error) {
	out := new(ERC20MintableTokenBaseOnlyForge)
	if err := eRC20Mintable.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
