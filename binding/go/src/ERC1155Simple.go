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

// IMulticall3Call is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call struct {
	Target   common.Address
	CallData []byte
}

// IMulticall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// IMulticall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// IMulticall3Result is an auto generated low-level Go binding around an user-defined struct.
type IMulticall3Result struct {
	Success    bool
	ReturnData []byte
}

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// ERC1155SimpleMetaData contains all meta data concerning the ERC1155Simple contract.
var ERC1155SimpleMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"burnFromBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"before\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"current\",\"type\":\"string\"}],\"name\":\"ERC1155BaseSetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC1155Simple",
	Bin: "0x608060405234801561000f575f5ffd5b50604051613caf380380613caf83398101604081905261002e916104c9565b84848484848085858162015180816001600160a01b03811661006a57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100935f826101c2565b505050506100c77faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c836101c260201b60201c565b5080515f5b81811015610106576100fe60048483815181106100eb576100eb61060a565b60200260200101516101d660201b60201c565b6001016100cc565b505050506101198161024e60201b60201c565b5082515f036101445760405163e1dea2ef60e01b8152636e616d6560e01b6004820152602401610061565b81515f036101705760405163e1dea2ef60e01b8152651cde5b589bdb60d21b6004820152602401610061565b80515f036101995760405163e1dea2ef60e01b81526275726960e81b6004820152602401610061565b60096101a584826106a2565b50600a6101b283826106a2565b505050505050505050505061075c565b5f6101cd838361025a565b90505b92915050565b6001600160a01b0381166102075760405163e1dea2ef60e01b815264666f72676560d81b6004820152602401610061565b610211828261028d565b1561024a576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b610257816102a1565b50565b5f8061026684846102ad565b905080156101cd575f848152600360205260409020610285908461028d565b509392505050565b5f6101cd836001600160a01b038416610313565b600861024a82826106a2565b5f82610309575f6102c66002546001600160a01b031690565b6001600160a01b0316146102ed57604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6101cd838361035f565b5f81815260018301602052604081205461035857508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556101d0565b505f6101d0565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610358575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556103b73390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d0565b80516001600160a01b0381168114610415575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156104565761045661041a565b604052919050565b5f82601f83011261046d575f5ffd5b81516001600160401b038111156104865761048661041a565b610499601f8201601f191660200161042e565b8181528460208386010111156104ad575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a086880312156104dd575f5ffd5b6104e6866103ff565b60208701519095506001600160401b03811115610501575f5ffd5b8601601f81018813610511575f5ffd5b80516001600160401b0381111561052a5761052a61041a565b8060051b61053a6020820161042e565b9182526020818401810192908101908b841115610555575f5ffd5b6020850194505b8385101561057e5761056d856103ff565b82526020948501949091019061055c565b60408b0151909850935050506001600160401b03821115905061059f575f5ffd5b6105ab8882890161045e565b606088015190945090506001600160401b038111156105c8575f5ffd5b6105d48882890161045e565b608088015190935090506001600160401b038111156105f1575f5ffd5b6105fd8882890161045e565b9150509295509295909350565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061063257607f821691505b60208210810361065057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561069d57805f5260205f20601f840160051c8101602085101561067b5750805b601f840160051c820191505b8181101561069a575f8155600101610687565b50505b505050565b81516001600160401b038111156106bb576106bb61041a565b6106cf816106c9845461061e565b84610656565b6020601f821160018114610701575f83156106ea5750848201515b5f19600385901b1c1916600184901b17845561069a565b5f84815260208120601f198516915b828110156107305787850151825560209485019460019092019101610710565b508482101561074d57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b613546806107695f395ff3fe608060405234801561000f575f5ffd5b50600436106102b6575f3560e01c80639010d07c11610171578063cc8463c8116100d2578063e08ba4bb11610088578063ec87621c1161006e578063ec87621c14610617578063f242432a1461063e578063fe2df3e814610651575f5ffd5b8063e08ba4bb146105f1578063e985e9c514610604575f5ffd5b8063cf6eefb7116100b8578063cf6eefb71461058a578063d547741f146105d6578063d602b9fd146105e9575f5ffd5b8063cc8463c81461057a578063cefc142914610582575f5ffd5b8063a217fddf11610127578063a3246ad31161010d578063a3246ad31461054c578063a40283d51461055f578063ca15c87314610567575f5ffd5b8063a217fddf14610532578063a22cb46514610539575f5ffd5b806395d89b411161015757806395d89b41146104f05780639ca92df9146104f8578063a1eda53c1461050b575f5ffd5b80639010d07c1461049a57806391d14854146104ad575f5ffd5b80632f2ff15d1161021b578063649a5ec7116101d157806384ef8ffc116101b757806384ef8ffc1461044057806387f453531461047f5780638da5cb5b14610492575f5ffd5b8063649a5ec71461041a578063731133e91461042d575f5ffd5b80634e1273f4116102015780634e1273f4146103d25780635c4e62c4146103f2578063634e93da14610407575f5ffd5b80632f2ff15d146103ac57806336568abe146103bf575f5ffd5b80630e89341c116102705780631f7fdffa116102565780631f7fdffa14610364578063248a9ca3146103775780632eb2c2d614610399575f5ffd5b80630e89341c1461033e578063124d91e514610351575f5ffd5b8063022d63fb116102a0578063022d63fb1461030357806306fdde031461031f5780630aa6220b14610334575f5ffd5b8062fdd58e146102ba57806301ffc9a7146102e0575b5f5ffd5b6102cd6102c8366004612a68565b610664565b6040519081526020015b60405180910390f35b6102f36102ee366004612abd565b61069d565b60405190151581526020016102d7565b620697805b60405165ffffffffffff90911681526020016102d7565b610327610769565b6040516102d79190612b24565b61033c6107f9565b005b61032761034c366004612b36565b61080e565b61033c61035f366004612b4d565b610819565b61033c610372366004612ccb565b610829565b6102cd610385366004612b36565b5f9081526020819052604090206001015490565b61033c6103a7366004612dfa565b6108d0565b61033c6103ba366004612ea9565b6108dd565b61033c6103cd366004612ea9565b6108eb565b6103e56103e0366004612ed3565b6108f5565b6040516102d79190612fd0565b6103fa610901565b6040516102d79190612fe2565b61033c61041536600461303a565b610912565b61033c610428366004613053565b610925565b61033c61043b366004613078565b610938565b60025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102d7565b6102f361048d36600461303a565b61098c565b61045a610998565b61045a6104a83660046130be565b6109b8565b6102f36104bb366004612ea9565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6103276109cf565b61045a610506366004612b36565b6109de565b6105136109ea565b6040805165ffffffffffff9384168152929091166020830152016102d7565b6102cd5f81565b61033c6105473660046130ed565b610a64565b6103fa61055a366004612b36565b610a6e565b6102cd610a87565b6102cd610575366004612b36565b610a92565b610308610aa8565b61033c610b45565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102d7565b61033c6105e4366004612ea9565b610ba1565b61033c610bab565b61033c6105ff366004613115565b610bbd565b6102f3610612366004613189565b610bc8565b6102cd7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61033c61064c3660046131b1565b610c04565b61033c61065f366004613205565b610c11565b5f81815260066020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091528120545b90505b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061072f57507fffffffff0000000000000000000000000000000000000000000000000000000082167f9ea8d94d00000000000000000000000000000000000000000000000000000000145b8061075a57507fffffffff000000000000000000000000000000000000000000000000000000008216155b80610697575061069782610ca2565b60606009805461077890613284565b80601f01602080910402602001604051908101604052809291908181526020018280546107a490613284565b80156107ef5780601f106107c6576101008083540402835291602001916107ef565b820191905f5260205f20905b8154815290600101906020018083116107d257829003601f168201915b5050505050905090565b5f61080381610d43565b61080b610d4d565b50565b606061069782610d59565b610824838383610deb565b505050565b6108323361098c565b61088857335b6040517f25a9dbc100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015b60405180910390fd5b6108c985858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610e7892505050565b5050505050565b6108c98585858585610eda565b6108e78282610f74565b5050565b6108e78282610fb5565b606061069483836110ba565b606061090d600461119e565b905090565b5f61091c81610d43565b6108e7826111b1565b5f61092f81610d43565b6108e782611230565b6109413361098c565b61094b5733610838565b6108c985858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061129f92505050565b5f610697600483611320565b5f61090d60025473ffffffffffffffffffffffffffffffffffffffff1690565b5f828152600360205260408120610694908361134e565b6060600a805461077890613284565b5f61069760048361134e565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610a2c57504265ffffffffffff821610155b610a37575f5f610a5c565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b6108e78282611359565b5f8181526003602052604090206060906106979061119e565b5f61090d6004611364565b5f81815260036020526040812061069790611364565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610ae957504265ffffffffffff8216105b610b1b576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610b3f565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610b99576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161087f565b61080b61136d565b6108e7828261145e565b5f610bb581610d43565b61080b61149f565b6108248383836114a9565b73ffffffffffffffffffffffffffffffffffffffff8083165f90815260076020908152604080832093851683529290529081205460ff16610694565b6108c98585858585611513565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610c3b81610d43565b612a3882610c4b576115a5610c4f565b6115f85b9050835f5b81811015610c9957610c916004888884818110610c7357610c736132d5565b9050602002016020810190610c88919061303a565b8563ffffffff16565b600101610c54565b50505050505050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167fd9b67a26000000000000000000000000000000000000000000000000000000001480610d3457507fffffffff0000000000000000000000000000000000000000000000000000000082167f0e89341c00000000000000000000000000000000000000000000000000000000145b806106975750610697826116ba565b61080b81336116c4565b610d575f5f611749565b565b606060088054610d6890613284565b80601f0160208091040260200160405190810160405280929190818152602001828054610d9490613284565b8015610ddf5780601f10610db657610100808354040283529160200191610ddf565b820191905f5260205f20905b815481529060010190602001808311610dc257829003601f168201915b50505050509050919050565b73ffffffffffffffffffffffffffffffffffffffff8316610e3a576040517f01a835140000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b604080516001808252602082018590528183019081526060820184905260a082019092525f608082018181529192916108c9918791859085906118a2565b73ffffffffffffffffffffffffffffffffffffffff8416610ec7576040517f57f447ce0000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b610ed45f858585856118a2565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff86168114801590610f085750610f068682610bc8565b155b15610f5f576040517fe237d92200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528716602482015260440161087f565b610f6c86868686866118af565b505050505050565b81610fab576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108e7828261195a565b81158015610fdd575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b156110b05760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580611031575065ffffffffffff8116155b8061104457504265ffffffffffff821610155b15611085576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161087f565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b6108e7828261197e565b6060815183511461110457815183516040517f5b0599910000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161087f565b5f835167ffffffffffffffff81111561111f5761111f612b7d565b604051908082528060200260200182016040528015611148578160200160208202803683370190505b5090505f5b84518110156111965760208082028601015161117190602080840287010151610664565b828281518110611183576111836132d5565b602090810291909101015260010161114d565b509392505050565b60605f6111aa836119d7565b9392505050565b5f6111ba610aa8565b6111c342611a2f565b6111cd919061332f565b90506111d98282611a7e565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61123a82611b19565b61124342611a2f565b61124d919061332f565b90506112598282611749565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff84166112ee576040517f57f447ce0000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b60408051600180825260208201869052818301908152606082018590526080820190925290610f6c5f878484876118a2565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610694565b5f6106948383611b60565b6108e7338383611b86565b5f610697825490565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff168015806113bd57504265ffffffffffff821610155b156113fe576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161087f565b6114265f61142160025473ffffffffffffffffffffffffffffffffffffffff1690565b611b91565b506114315f83611b9c565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b81611495576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108e78282611ba7565b610d575f5f611a7e565b73ffffffffffffffffffffffffffffffffffffffff83166114f8576040517f01a835140000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b610824835f848460405180602001604052805f8152506118a2565b3373ffffffffffffffffffffffffffffffffffffffff86168114801590611541575061153f8682610bc8565b155b15611598576040517fe237d92200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80831660048301528716602482015260440161087f565b610f6c8686868686611bcb565b6115af8282611c9b565b156108e75760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff8116611667576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f726765000000000000000000000000000000000000000000000000000000600482015260240161087f565b6116718282611cbc565b156108e75760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f61069782611cdd565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166108e7576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161087f565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801561181d574265ffffffffffff821610156117f4576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a0100000000000000000000000000000000000000000000000000000291909117905561181d565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b6108c98585858585611ce7565b73ffffffffffffffffffffffffffffffffffffffff84166118fe576040517f57f447ce0000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b73ffffffffffffffffffffffffffffffffffffffff851661194d576040517f01a835140000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b6108c985858585856118a2565b5f8281526020819052604090206001015461197481610d43565b610ed48383611b9c565b73ffffffffffffffffffffffffffffffffffffffff811633146119cd576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108248282611b91565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610ddf57602002820191905f5260205f20905b815481526020019060010190808311611a105750505050509050919050565b5f65ffffffffffff821115611a7a576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161087f565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015610824576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611b23610aa8565b90508065ffffffffffff168365ffffffffffff1611611b4b57611b46838261334d565b6111aa565b6111aa65ffffffffffff841662069780611d47565b5f825f018281548110611b7557611b756132d5565b905f5260205f200154905092915050565b610824838383611d56565b5f6106948383611e3c565b5f6106948383611e67565b5f82815260208190526040902060010154611bc181610d43565b610ed48383611b91565b73ffffffffffffffffffffffffffffffffffffffff8416611c1a576040517f57f447ce0000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b73ffffffffffffffffffffffffffffffffffffffff8516611c69576040517f01a835140000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b60408051600180825260208201869052818301908152606082018590526080820190925290610c9987878484876118a2565b5f6106948373ffffffffffffffffffffffffffffffffffffffff8416611e92565b5f6106948373ffffffffffffffffffffffffffffffffffffffff8416611f75565b5f61069782611fc1565b611cf385858585612016565b73ffffffffffffffffffffffffffffffffffffffff8416156108c95782513390600103611d395760208481015190840151611d32838989858589612022565b5050610f6c565b610f6c818787878787612211565b5f828218828410028218610694565b73ffffffffffffffffffffffffffffffffffffffff8216611da5576040517fced3e1000000000000000000000000000000000000000000000000000000000081525f600482015260240161087f565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526007602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b5f5f611e4884846123a0565b90508015610694575f8481526003602052604090206111969084611c9b565b5f5f611e738484612401565b90508015610694575f8481526003602052604090206111969084611cbc565b5f8181526001830160205260408120548015611f6c575f611eb460018361336b565b85549091505f90611ec79060019061336b565b9050808214611f26575f865f018281548110611ee557611ee56132d5565b905f5260205f200154905080875f018481548110611f0557611f056132d5565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611f3757611f3761337e565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610697565b5f915050610697565b5f818152600183016020526040812054611fba57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610697565b505f610697565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f0000000000000000000000000000000000000000000000000000000014806106975750610697826124bf565b610ed484848484612514565b73ffffffffffffffffffffffffffffffffffffffff84163b15610f6c576040517ff23a6e6100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063f23a6e619061209990899089908890889088906004016133ab565b6020604051808303815f875af19250505080156120f1575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526120ee9181019061340c565b60015b61217e573d80801561211e576040519150601f19603f3d011682016040523d82523d5f602084013e612123565b606091505b5080515f03612176576040517f57f447ce00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260240161087f565b805160208201fd5b7fffffffff0000000000000000000000000000000000000000000000000000000081167ff23a6e610000000000000000000000000000000000000000000000000000000014610c99576040517f57f447ce00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260240161087f565b73ffffffffffffffffffffffffffffffffffffffff84163b15610f6c576040517fbc197c8100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063bc197c81906122889089908990889088908890600401613427565b6020604051808303815f875af19250505080156122e0575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526122dd9181019061340c565b60015b61230d573d80801561211e576040519150601f19603f3d011682016040523d82523d5f602084013e612123565b7fffffffff0000000000000000000000000000000000000000000000000000000081167fbc197c810000000000000000000000000000000000000000000000000000000014610c99576040517f57f447ce00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8616600482015260240161087f565b5f821580156123c9575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b156123f757600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b61069483836127f7565b5f826124b5575f61242760025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612474576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b61069483836128b0565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f314987860000000000000000000000000000000000000000000000000000000014806106975750610697826129a2565b805182511461255c57815181516040517f5b0599910000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161087f565b335f5b83518110156126cb5760208181028581018201519085019091015173ffffffffffffffffffffffffffffffffffffffff881615612661575f82815260066020908152604080832073ffffffffffffffffffffffffffffffffffffffff8c1684529091529020548181101561262c576040517f03dee4c500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a16600482015260248101829052604481018390526064810184905260840161087f565b5f83815260066020908152604080832073ffffffffffffffffffffffffffffffffffffffff8d16845290915290209082900390555b73ffffffffffffffffffffffffffffffffffffffff8716156126c1575f82815260066020908152604080832073ffffffffffffffffffffffffffffffffffffffff8b168452909152812080548392906126bb9084906134a3565b90915550505b505060010161255f565b5082516001036127725760208301515f9060208401519091508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f628585604051612763929190918252602082015260400190565b60405180910390a450506108c9565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb86866040516127e89291906134b6565b60405180910390a45050505050565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615611fba575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610697565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16611fba575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556129403390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610697565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061069757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610697565b610d576134e3565b803573ffffffffffffffffffffffffffffffffffffffff81168114612a63575f5ffd5b919050565b5f5f60408385031215612a79575f5ffd5b612a8283612a40565b946020939093013593505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461080b575f5ffd5b5f60208284031215612acd575f5ffd5b813561069481612a90565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6106946020830184612ad8565b5f60208284031215612b46575f5ffd5b5035919050565b5f5f5f60608486031215612b5f575f5ffd5b612b6884612a40565b95602085013595506040909401359392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612bf157612bf1612b7d565b604052919050565b5f67ffffffffffffffff821115612c1257612c12612b7d565b5060051b60200190565b5f82601f830112612c2b575f5ffd5b8135612c3e612c3982612bf9565b612baa565b8082825260208201915060208360051b860101925085831115612c5f575f5ffd5b602085015b83811015612c7c578035835260209283019201612c64565b5095945050505050565b5f5f83601f840112612c96575f5ffd5b50813567ffffffffffffffff811115612cad575f5ffd5b602083019150836020828501011115612cc4575f5ffd5b9250929050565b5f5f5f5f5f60808688031215612cdf575f5ffd5b612ce886612a40565b9450602086013567ffffffffffffffff811115612d03575f5ffd5b612d0f88828901612c1c565b945050604086013567ffffffffffffffff811115612d2b575f5ffd5b612d3788828901612c1c565b935050606086013567ffffffffffffffff811115612d53575f5ffd5b612d5f88828901612c86565b969995985093965092949392505050565b5f82601f830112612d7f575f5ffd5b813567ffffffffffffffff811115612d9957612d99612b7d565b612dca60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612baa565b818152846020838601011115612dde575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f60a08688031215612e0e575f5ffd5b612e1786612a40565b9450612e2560208701612a40565b9350604086013567ffffffffffffffff811115612e40575f5ffd5b612e4c88828901612c1c565b935050606086013567ffffffffffffffff811115612e68575f5ffd5b612e7488828901612c1c565b925050608086013567ffffffffffffffff811115612e90575f5ffd5b612e9c88828901612d70565b9150509295509295909350565b5f5f60408385031215612eba575f5ffd5b82359150612eca60208401612a40565b90509250929050565b5f5f60408385031215612ee4575f5ffd5b823567ffffffffffffffff811115612efa575f5ffd5b8301601f81018513612f0a575f5ffd5b8035612f18612c3982612bf9565b8082825260208201915060208360051b850101925087831115612f39575f5ffd5b6020840193505b82841015612f6257612f5184612a40565b825260209384019390910190612f40565b9450505050602083013567ffffffffffffffff811115612f80575f5ffd5b612f8c85828601612c1c565b9150509250929050565b5f8151808452602084019350602083015f5b82811015612fc6578151865260209586019590910190600101612fa8565b5093949350505050565b602081525f6106946020830184612f96565b602080825282518282018190525f918401906040840190835b8181101561302f57835173ffffffffffffffffffffffffffffffffffffffff16835260209384019390920191600101612ffb565b509095945050505050565b5f6020828403121561304a575f5ffd5b61069482612a40565b5f60208284031215613063575f5ffd5b813565ffffffffffff81168114610694575f5ffd5b5f5f5f5f5f6080868803121561308c575f5ffd5b61309586612a40565b94506020860135935060408601359250606086013567ffffffffffffffff811115612d53575f5ffd5b5f5f604083850312156130cf575f5ffd5b50508035926020909101359150565b80358015158114612a63575f5ffd5b5f5f604083850312156130fe575f5ffd5b61310783612a40565b9150612eca602084016130de565b5f5f5f60608486031215613127575f5ffd5b61313084612a40565b9250602084013567ffffffffffffffff81111561314b575f5ffd5b61315786828701612c1c565b925050604084013567ffffffffffffffff811115613173575f5ffd5b61317f86828701612c1c565b9150509250925092565b5f5f6040838503121561319a575f5ffd5b6131a383612a40565b9150612eca60208401612a40565b5f5f5f5f5f60a086880312156131c5575f5ffd5b6131ce86612a40565b94506131dc60208701612a40565b93506040860135925060608601359150608086013567ffffffffffffffff811115612e90575f5ffd5b5f5f5f60408486031215613217575f5ffd5b833567ffffffffffffffff81111561322d575f5ffd5b8401601f8101861361323d575f5ffd5b803567ffffffffffffffff811115613253575f5ffd5b8660208260051b8401011115613267575f5ffd5b60209182019450925061327b9085016130de565b90509250925092565b600181811c9082168061329857607f821691505b6020821081036132cf577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff818116838216019081111561069757610697613302565b65ffffffffffff828116828216039081111561069757610697613302565b8181038181111561069757610697613302565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015282606082015260a060808201525f61340160a0830184612ad8565b979650505050505050565b5f6020828403121561341c575f5ffd5b815161069481612a90565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015260a060408201525f61347160a0830186612f96565b82810360608401526134838186612f96565b905082810360808401526134978185612ad8565b98975050505050505050565b8082018082111561069757610697613302565b604081525f6134c86040830185612f96565b82810360208401526134da8185612f96565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212205645717a3f5c48fcfecd501a9a9082e6382562702f93444930f31de4cc4fe81c64736f6c634300081c0033",
}

// ERC1155Simple is an auto generated Go binding around an Ethereum contract.
type ERC1155Simple struct {
	abi abi.ABI
}

// NewERC1155Simple creates a new instance of ERC1155Simple.
func NewERC1155Simple() *ERC1155Simple {
	parsed, err := ERC1155SimpleMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC1155Simple{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC1155Simple) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, string uri) returns()
func (eRC1155Simple *ERC1155Simple) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, uri string) []byte {
	enc, err := eRC1155Simple.abi.Pack("", owner, forges, name, symbol, uri)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC1155Simple *ERC1155Simple) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC1155Simple.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC1155Simple *ERC1155Simple) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC1155Simple.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC1155Simple *ERC1155Simple) PackMANAGERROLE() []byte {
	enc, err := eRC1155Simple.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC1155Simple *ERC1155Simple) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC1155Simple.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC1155Simple *ERC1155Simple) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC1155Simple.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (eRC1155Simple *ERC1155Simple) PackBalanceOf(account common.Address, id *big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("balanceOf", account, id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (eRC1155Simple *ERC1155Simple) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC1155Simple.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackBalanceOfBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (eRC1155Simple *ERC1155Simple) PackBalanceOfBatch(accounts []common.Address, ids []*big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("balanceOfBatch", accounts, ids)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOfBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (eRC1155Simple *ERC1155Simple) UnpackBalanceOfBatch(data []byte) ([]*big.Int, error) {
	out, err := eRC1155Simple.abi.Unpack("balanceOfBatch", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackBeginDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x634e93da.
//
// Solidity: function beginDefaultAdminTransfer(address newAdmin) returns()
func (eRC1155Simple *ERC1155Simple) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x124d91e5.
//
// Solidity: function burnFrom(address from, uint256 tokenID, uint256 amount) returns()
func (eRC1155Simple *ERC1155Simple) PackBurnFrom(from common.Address, tokenID *big.Int, amount *big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("burnFrom", from, tokenID, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFromBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe08ba4bb.
//
// Solidity: function burnFromBatch(address from, uint256[] tokenIDs, uint256[] amounts) returns()
func (eRC1155Simple *ERC1155Simple) PackBurnFromBatch(from common.Address, tokenIDs []*big.Int, amounts []*big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("burnFromBatch", from, tokenIDs, amounts)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC1155Simple *ERC1155Simple) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC1155Simple.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC1155Simple *ERC1155Simple) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC1155Simple *ERC1155Simple) PackDefaultAdmin() []byte {
	enc, err := eRC1155Simple.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC1155Simple.abi.Unpack("defaultAdmin", data)
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
func (eRC1155Simple *ERC1155Simple) PackDefaultAdminDelay() []byte {
	enc, err := eRC1155Simple.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC1155Simple.abi.Unpack("defaultAdminDelay", data)
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
func (eRC1155Simple *ERC1155Simple) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC1155Simple.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC1155Simple.abi.Unpack("defaultAdminDelayIncreaseWait", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackForgeByIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC1155Simple *ERC1155Simple) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC1155Simple *ERC1155Simple) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC1155Simple.abi.Unpack("forgeByIndex", data)
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
func (eRC1155Simple *ERC1155Simple) PackForgeCount() []byte {
	enc, err := eRC1155Simple.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC1155Simple *ERC1155Simple) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC1155Simple.abi.Unpack("forgeCount", data)
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
func (eRC1155Simple *ERC1155Simple) PackForges() []byte {
	enc, err := eRC1155Simple.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC1155Simple *ERC1155Simple) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC1155Simple.abi.Unpack("forges", data)
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
func (eRC1155Simple *ERC1155Simple) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC1155Simple *ERC1155Simple) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC1155Simple.abi.Unpack("getRoleAdmin", data)
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
func (eRC1155Simple *ERC1155Simple) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC1155Simple *ERC1155Simple) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC1155Simple.abi.Unpack("getRoleMember", data)
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
func (eRC1155Simple *ERC1155Simple) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC1155Simple *ERC1155Simple) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC1155Simple.abi.Unpack("getRoleMemberCount", data)
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
func (eRC1155Simple *ERC1155Simple) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC1155Simple *ERC1155Simple) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC1155Simple.abi.Unpack("getRoleMembers", data)
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
func (eRC1155Simple *ERC1155Simple) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC1155Simple *ERC1155Simple) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC1155Simple *ERC1155Simple) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC1155Simple.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (eRC1155Simple *ERC1155Simple) PackIsApprovedForAll(account common.Address, operator common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("isApprovedForAll", account, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (eRC1155Simple *ERC1155Simple) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := eRC1155Simple.abi.Unpack("isApprovedForAll", data)
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
func (eRC1155Simple *ERC1155Simple) PackIsForge(forge common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC1155Simple *ERC1155Simple) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC1155Simple.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x731133e9.
//
// Solidity: function mint(address to, uint256 tokenID, uint256 amount, bytes data) returns()
func (eRC1155Simple *ERC1155Simple) PackMint(to common.Address, tokenID *big.Int, amount *big.Int, data []byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("mint", to, tokenID, amount, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMintBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] tokenIDs, uint256[] amounts, bytes data) returns()
func (eRC1155Simple *ERC1155Simple) PackMintBatch(to common.Address, tokenIDs []*big.Int, amounts []*big.Int, data []byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("mintBatch", to, tokenIDs, amounts, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC1155Simple *ERC1155Simple) PackName() []byte {
	enc, err := eRC1155Simple.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC1155Simple *ERC1155Simple) UnpackName(data []byte) (string, error) {
	out, err := eRC1155Simple.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC1155Simple *ERC1155Simple) PackOwner() []byte {
	enc, err := eRC1155Simple.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC1155Simple *ERC1155Simple) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC1155Simple.abi.Unpack("owner", data)
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
func (eRC1155Simple *ERC1155Simple) PackPendingDefaultAdmin() []byte {
	enc, err := eRC1155Simple.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC1155Simple *ERC1155Simple) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC1155Simple.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC1155Simple *ERC1155Simple) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC1155Simple.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC1155Simple *ERC1155Simple) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC1155Simple.abi.Unpack("pendingDefaultAdminDelay", data)
	outstruct := new(PendingDefaultAdminDelayOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.NewDelay = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Schedule = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackRenounceRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (eRC1155Simple *ERC1155Simple) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC1155Simple *ERC1155Simple) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC1155Simple.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC1155Simple *ERC1155Simple) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC1155Simple.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeBatchTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (eRC1155Simple *ERC1155Simple) PackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (eRC1155Simple *ERC1155Simple) PackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("safeTransferFrom", from, to, id, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (eRC1155Simple *ERC1155Simple) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := eRC1155Simple.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC1155Simple *ERC1155Simple) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC1155Simple.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC1155Simple *ERC1155Simple) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC1155Simple.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC1155Simple *ERC1155Simple) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC1155Simple.abi.Unpack("supportsInterface", data)
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
func (eRC1155Simple *ERC1155Simple) PackSymbol() []byte {
	enc, err := eRC1155Simple.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC1155Simple *ERC1155Simple) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC1155Simple.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackUri is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (eRC1155Simple *ERC1155Simple) PackUri(id *big.Int) []byte {
	enc, err := eRC1155Simple.abi.Pack("uri", id)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUri is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (eRC1155Simple *ERC1155Simple) UnpackUri(data []byte) (string, error) {
	out, err := eRC1155Simple.abi.Unpack("uri", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// ERC1155SimpleApprovalForAll represents a ApprovalForAll event raised by the ERC1155Simple contract.
type ERC1155SimpleApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleApprovalForAll) ContractEventName() string {
	return ERC1155SimpleApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (eRC1155Simple *ERC1155Simple) UnpackApprovalForAllEvent(log *types.Log) (*ERC1155SimpleApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleApprovalForAll)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC1155Simple contract.
type ERC1155SimpleDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC1155SimpleDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC1155SimpleDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC1155Simple contract.
type ERC1155SimpleDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC1155SimpleDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC1155SimpleDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC1155Simple contract.
type ERC1155SimpleDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC1155SimpleDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC1155SimpleDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC1155Simple contract.
type ERC1155SimpleDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC1155SimpleDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC1155Simple *ERC1155Simple) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC1155SimpleDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleERC1155BaseSetBaseURI represents a ERC1155BaseSetBaseURI event raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155BaseSetBaseURI struct {
	Before  string
	Current string
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleERC1155BaseSetBaseURIEventName = "ERC1155BaseSetBaseURI"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleERC1155BaseSetBaseURI) ContractEventName() string {
	return ERC1155SimpleERC1155BaseSetBaseURIEventName
}

// UnpackERC1155BaseSetBaseURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC1155BaseSetBaseURI(string before, string current)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155BaseSetBaseURIEvent(log *types.Log) (*ERC1155SimpleERC1155BaseSetBaseURI, error) {
	event := "ERC1155BaseSetBaseURI"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleERC1155BaseSetBaseURI)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleForgeAdded represents a ForgeAdded event raised by the ERC1155Simple contract.
type ERC1155SimpleForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleForgeAdded) ContractEventName() string {
	return ERC1155SimpleForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC1155Simple *ERC1155Simple) UnpackForgeAddedEvent(log *types.Log) (*ERC1155SimpleForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleForgeRemoved represents a ForgeRemoved event raised by the ERC1155Simple contract.
type ERC1155SimpleForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleForgeRemoved) ContractEventName() string {
	return ERC1155SimpleForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC1155Simple *ERC1155Simple) UnpackForgeRemovedEvent(log *types.Log) (*ERC1155SimpleForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleRoleAdminChanged represents a RoleAdminChanged event raised by the ERC1155Simple contract.
type ERC1155SimpleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleRoleAdminChanged) ContractEventName() string {
	return ERC1155SimpleRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC1155Simple *ERC1155Simple) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC1155SimpleRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleRoleGranted represents a RoleGranted event raised by the ERC1155Simple contract.
type ERC1155SimpleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleRoleGranted) ContractEventName() string {
	return ERC1155SimpleRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC1155Simple *ERC1155Simple) UnpackRoleGrantedEvent(log *types.Log) (*ERC1155SimpleRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleRoleRevoked represents a RoleRevoked event raised by the ERC1155Simple contract.
type ERC1155SimpleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleRoleRevoked) ContractEventName() string {
	return ERC1155SimpleRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC1155Simple *ERC1155Simple) UnpackRoleRevokedEvent(log *types.Log) (*ERC1155SimpleRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleTransferBatch represents a TransferBatch event raised by the ERC1155Simple contract.
type ERC1155SimpleTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleTransferBatchEventName = "TransferBatch"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleTransferBatch) ContractEventName() string {
	return ERC1155SimpleTransferBatchEventName
}

// UnpackTransferBatchEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (eRC1155Simple *ERC1155Simple) UnpackTransferBatchEvent(log *types.Log) (*ERC1155SimpleTransferBatch, error) {
	event := "TransferBatch"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleTransferBatch)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleTransferSingle represents a TransferSingle event raised by the ERC1155Simple contract.
type ERC1155SimpleTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleTransferSingleEventName = "TransferSingle"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleTransferSingle) ContractEventName() string {
	return ERC1155SimpleTransferSingleEventName
}

// UnpackTransferSingleEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (eRC1155Simple *ERC1155Simple) UnpackTransferSingleEvent(log *types.Log) (*ERC1155SimpleTransferSingle, error) {
	event := "TransferSingle"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleTransferSingle)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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

// ERC1155SimpleURI represents a URI event raised by the ERC1155Simple contract.
type ERC1155SimpleURI struct {
	Value string
	Id    *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC1155SimpleURIEventName = "URI"

// ContractEventName returns the user-defined event name.
func (ERC1155SimpleURI) ContractEventName() string {
	return ERC1155SimpleURIEventName
}

// UnpackURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event URI(string value, uint256 indexed id)
func (eRC1155Simple *ERC1155Simple) UnpackURIEvent(log *types.Log) (*ERC1155SimpleURI, error) {
	event := "URI"
	if log.Topics[0] != eRC1155Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC1155SimpleURI)
	if len(log.Data) > 0 {
		if err := eRC1155Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC1155Simple.abi.Events[event].Inputs {
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
func (eRC1155Simple *ERC1155Simple) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155InsufficientBalance"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155InsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155InvalidApprover"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155InvalidArrayLength"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155InvalidArrayLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155InvalidOperator"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155InvalidSender"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["ERC1155MissingApprovalForAll"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackERC1155MissingApprovalForAllError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC1155Simple.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC1155Simple.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC1155SimpleAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC1155Simple contract.
type ERC1155SimpleAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC1155SimpleAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC1155Simple *ERC1155Simple) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC1155SimpleAccessControlBadConfirmation, error) {
	out := new(ERC1155SimpleAccessControlBadConfirmation)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC1155Simple contract.
type ERC1155SimpleAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC1155SimpleAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC1155Simple *ERC1155Simple) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC1155SimpleAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC1155SimpleAccessControlEnforcedDefaultAdminDelay)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC1155Simple contract.
type ERC1155SimpleAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC1155SimpleAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC1155Simple *ERC1155Simple) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC1155SimpleAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC1155SimpleAccessControlEnforcedDefaultAdminRules)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC1155Simple contract.
type ERC1155SimpleAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC1155SimpleAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC1155Simple *ERC1155Simple) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC1155SimpleAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC1155SimpleAccessControlInvalidDefaultAdmin)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC1155Simple contract.
type ERC1155SimpleAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC1155SimpleAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC1155Simple *ERC1155Simple) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC1155SimpleAccessControlUnauthorizedAccount, error) {
	out := new(ERC1155SimpleAccessControlUnauthorizedAccount)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155InsufficientBalance represents a ERC1155InsufficientBalance error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155InsufficientBalance struct {
	Sender  common.Address
	Balance *big.Int
	Needed  *big.Int
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func ERC1155SimpleERC1155InsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0x03dee4c573c982787b5f3537d6323ffaca9d864448aa6bd828ada9e5d0837036")
}

// UnpackERC1155InsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InsufficientBalance(address sender, uint256 balance, uint256 needed, uint256 tokenId)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155InsufficientBalanceError(raw []byte) (*ERC1155SimpleERC1155InsufficientBalance, error) {
	out := new(ERC1155SimpleERC1155InsufficientBalance)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155InvalidApprover represents a ERC1155InvalidApprover error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func ERC1155SimpleERC1155InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0x3e31884e33c33ce0039d1905e3c252950ae3b95240f36d4fff81f5ff6752ef99")
}

// UnpackERC1155InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidApprover(address approver)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155InvalidApproverError(raw []byte) (*ERC1155SimpleERC1155InvalidApprover, error) {
	out := new(ERC1155SimpleERC1155InvalidApprover)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155InvalidArrayLength represents a ERC1155InvalidArrayLength error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155InvalidArrayLength struct {
	IdsLength    *big.Int
	ValuesLength *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func ERC1155SimpleERC1155InvalidArrayLengthErrorID() common.Hash {
	return common.HexToHash("0x5b0599913619cfa5633692652638ed25cafcd079c9beae8c251b12c23dcc83f2")
}

// UnpackERC1155InvalidArrayLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidArrayLength(uint256 idsLength, uint256 valuesLength)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155InvalidArrayLengthError(raw []byte) (*ERC1155SimpleERC1155InvalidArrayLength, error) {
	out := new(ERC1155SimpleERC1155InvalidArrayLength)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155InvalidArrayLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155InvalidOperator represents a ERC1155InvalidOperator error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func ERC1155SimpleERC1155InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0xced3e10010b9d2aa24827119d0db4a8feec73aea48b4b3e470d8a9f3ff723569")
}

// UnpackERC1155InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidOperator(address operator)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155InvalidOperatorError(raw []byte) (*ERC1155SimpleERC1155InvalidOperator, error) {
	out := new(ERC1155SimpleERC1155InvalidOperator)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155InvalidReceiver represents a ERC1155InvalidReceiver error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func ERC1155SimpleERC1155InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x57f447ceed621d9e134e26de5772c88799abb7322ce2a87f95dce247d47105c6")
}

// UnpackERC1155InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidReceiver(address receiver)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155InvalidReceiverError(raw []byte) (*ERC1155SimpleERC1155InvalidReceiver, error) {
	out := new(ERC1155SimpleERC1155InvalidReceiver)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155InvalidSender represents a ERC1155InvalidSender error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155InvalidSender(address sender)
func ERC1155SimpleERC1155InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x01a83514e94b34009110b75cac6742ba33bd7c62f18a3616bafea52855d3b175")
}

// UnpackERC1155InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155InvalidSender(address sender)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155InvalidSenderError(raw []byte) (*ERC1155SimpleERC1155InvalidSender, error) {
	out := new(ERC1155SimpleERC1155InvalidSender)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleERC1155MissingApprovalForAll represents a ERC1155MissingApprovalForAll error raised by the ERC1155Simple contract.
type ERC1155SimpleERC1155MissingApprovalForAll struct {
	Operator common.Address
	Owner    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func ERC1155SimpleERC1155MissingApprovalForAllErrorID() common.Hash {
	return common.HexToHash("0xe237d922be9fac42efeaaaffb42cc6b57e0ff95d94a1b74daeff69adc7657754")
}

// UnpackERC1155MissingApprovalForAllError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1155MissingApprovalForAll(address operator, address owner)
func (eRC1155Simple *ERC1155Simple) UnpackERC1155MissingApprovalForAllError(raw []byte) (*ERC1155SimpleERC1155MissingApprovalForAll, error) {
	out := new(ERC1155SimpleERC1155MissingApprovalForAll)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "ERC1155MissingApprovalForAll", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC1155Simple contract.
type ERC1155SimpleSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC1155SimpleSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC1155Simple *ERC1155Simple) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC1155SimpleSafeCastOverflowedUintDowncast, error) {
	out := new(ERC1155SimpleSafeCastOverflowedUintDowncast)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC1155Simple contract.
type ERC1155SimpleTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC1155SimpleTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC1155Simple *ERC1155Simple) UnpackTokenBaseNullInputError(raw []byte) (*ERC1155SimpleTokenBaseNullInput, error) {
	out := new(ERC1155SimpleTokenBaseNullInput)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC1155SimpleTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC1155Simple contract.
type ERC1155SimpleTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC1155SimpleTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC1155Simple *ERC1155Simple) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC1155SimpleTokenBaseOnlyForge, error) {
	out := new(ERC1155SimpleTokenBaseOnlyForge)
	if err := eRC1155Simple.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
