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

// ERC721AutoIncrMetaData contains all meta data concerning the ERC721AutoIncr contract.
var ERC721AutoIncrMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"before\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"current\",\"type\":\"string\"}],\"name\":\"ERC721BaseSetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC721AutoIncr",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161378b38038061378b83398101604081905261002e916104ff565b8484848484828286868162015180816001600160a01b03811661006b57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100945f826101ca565b505050506100c87faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c836101ca60201b60201c565b5080515f5b81811015610107576100ff60048483815181106100ec576100ec610640565b60200260200101516101de60201b60201c565b6001016100cd565b50505050816006908161011a91906106d8565b50600761012782826106d8565b50505082515f036101545760405163e1dea2ef60e01b8152636e616d6560e01b6004820152602401610062565b81515f036101805760405163e1dea2ef60e01b8152651cde5b589bdb60d21b6004820152602401610062565b80515f036101b25760405163e1dea2ef60e01b81526b62617365546f6b656e55524960a01b6004820152602401610062565b6101bb81610256565b5050505050505050505061085a565b5f6101d5838361029c565b90505b92915050565b6001600160a01b03811661020f5760405163e1dea2ef60e01b815264666f72676560d81b6004820152602401610062565b61021982826102cf565b15610252576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b7ffcd3ec38ee6a55750e2501220e2c259a8309f45d101943d2f1c9a40ecfd6b3a6600c826040516102889291906107c0565b60405180910390a1600c61025282826106d8565b5f806102a884846102e3565b905080156101d5575f8481526003602052604090206102c790846102cf565b509392505050565b5f6101d5836001600160a01b038416610349565b5f8261033f575f6102fc6002546001600160a01b031690565b6001600160a01b03161461032357604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6101d58383610395565b5f81815260018301602052604081205461038e57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556101d8565b505f6101d8565b5f828152602081815260408083206001600160a01b038516845290915281205460ff1661038e575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556103ed3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d8565b80516001600160a01b038116811461044b575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561048c5761048c610450565b604052919050565b5f82601f8301126104a3575f5ffd5b81516001600160401b038111156104bc576104bc610450565b6104cf601f8201601f1916602001610464565b8181528460208386010111156104e3575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a08688031215610513575f5ffd5b61051c86610435565b60208701519095506001600160401b03811115610537575f5ffd5b8601601f81018813610547575f5ffd5b80516001600160401b0381111561056057610560610450565b8060051b61057060208201610464565b9182526020818401810192908101908b84111561058b575f5ffd5b6020850194505b838510156105b4576105a385610435565b825260209485019490910190610592565b60408b0151909850935050506001600160401b0382111590506105d5575f5ffd5b6105e188828901610494565b606088015190945090506001600160401b038111156105fe575f5ffd5b61060a88828901610494565b608088015190935090506001600160401b03811115610627575f5ffd5b61063388828901610494565b9150509295509295909350565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061066857607f821691505b60208210810361068657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156106d357805f5260205f20601f840160051c810160208510156106b15750805b601f840160051c820191505b818110156106d0575f81556001016106bd565b50505b505050565b81516001600160401b038111156106f1576106f1610450565b610705816106ff8454610654565b8461068c565b6020601f821160018114610737575f83156107205750848201515b5f19600385901b1c1916600184901b1784556106d0565b5f84815260208120601f198516915b828110156107665787850151825560209485019460019092019101610746565b508482101561078357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f5f84546107d181610654565b806040860152600182165f81146107ef576001811461080b5761083c565b60ff1983166060870152606082151560051b870101935061083c565b875f5260205f205f5b8381101561083357815488820160600152600190910190602001610814565b87016060019450505b50505082810360208401526108518185610792565b95945050505050565b612f24806108675f395ff3fe608060405234801561000f575f5ffd5b50600436106102c2575f3560e01c80639010d07c1161017c578063b88d4fde116100dd578063cf6eefb711610093578063e985e9c51161006e578063e985e9c51461062d578063ec87621c14610640578063fe2df3e814610667575f5ffd5b8063cf6eefb7146105c6578063d547741f14610612578063d602b9fd14610625575f5ffd5b8063ca15c873116100c3578063ca15c873146105a3578063cc8463c8146105b6578063cefc1429146105be575f5ffd5b8063b88d4fde1461057d578063c87b56dd14610590575f5ffd5b8063a1eda53c11610132578063a22cb46511610118578063a22cb4651461054f578063a3246ad314610562578063a40283d514610575575f5ffd5b8063a1eda53c14610521578063a217fddf14610548575f5ffd5b806394d008ef1161016257806394d008ef146104f357806395d89b41146105065780639ca92df91461050e575f5ffd5b80639010d07c1461049d57806391d14854146104b0575f5ffd5b806342842e0e1161022657806370a08231116101dc57806384ef8ffc116101c257806384ef8ffc1461046457806387f45353146104825780638da5cb5b14610495575f5ffd5b806370a082311461043e57806379cc679014610451575f5ffd5b8063634e93da1161020c578063634e93da146104055780636352211e14610418578063649a5ec71461042b575f5ffd5b806342842e0e146103dd5780635c4e62c4146103f0575f5ffd5b80630aa6220b1161027b578063248a9ca311610261578063248a9ca3146103875780632f2ff15d146103b757806336568abe146103ca575f5ffd5b80630aa6220b1461036c57806323b872dd14610374575f5ffd5b806306fdde03116102ab57806306fdde031461030a578063081812fc1461031f578063095ea7b314610357575f5ffd5b806301ffc9a7146102c6578063022d63fb146102ee575b5f5ffd5b6102d96102d436600461283f565b61067a565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016102e5565b610312610821565b6040516102e591906128a6565b61033261032d3660046128b8565b610830565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102e5565b61036a6103653660046128f7565b61083a565b005b61036a610848565b61036a61038236600461291f565b61085d565b6103a96103953660046128b8565b5f9081526020819052604090206001015490565b6040519081526020016102e5565b61036a6103c5366004612959565b61086d565b61036a6103d8366004612959565b610877565b61036a6103eb36600461291f565b610881565b6103f861089b565b6040516102e59190612983565b61036a6104133660046129db565b6108a7565b6103326104263660046128b8565b6108ba565b61036a6104393660046129f4565b6108c4565b6103a961044c3660046129db565b6108d7565b61036a61045f3660046128f7565b6108e1565b60025473ffffffffffffffffffffffffffffffffffffffff16610332565b6102d96104903660046129db565b6108f5565b610332610901565b6103326104ab366004612a19565b610921565b6102d96104be366004612959565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6103a9610501366004612b2e565b61093f565b6103126109ae565b61033261051c3660046128b8565b6109b8565b6105296109c4565b6040805165ffffffffffff9384168152929091166020830152016102e5565b6103a95f81565b61036a61055d366004612b90565b610a3e565b6103f86105703660046128b8565b610a48565b6103a9610a61565b61036a61058b366004612bb8565b610a6c565b61031261059e3660046128b8565b610a7e565b6103a96105b13660046128b8565b610a89565b6102f3610a9f565b61036a610b3c565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102e5565b61036a610620366004612959565b610b98565b61036a610ba2565b6102d961063b366004612c1c565b610bb4565b6103a97faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61036a610675366004612c44565b610bf0565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061070c57507fffffffff0000000000000000000000000000000000000000000000000000000082167fed1c6f7f00000000000000000000000000000000000000000000000000000000145b8061075857507fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd00000000000000000000000000000000000000000000000000000000145b806107a457507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b806107cf57507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061081b57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b606061082b610c81565b905090565b5f61081b82610d11565b6108448282610d25565b5050565b5f61085281610d30565b61085a610d3a565b50565b610868838383610d46565b505050565b6108448282610e2f565b6108448282610e70565b61086883838360405180602001604052805f815250610a6c565b606061082b6004610f75565b5f6108b181610d30565b61084482610f81565b5f61081b82611000565b5f6108ce81610d30565b6108448261100a565b5f61081b82611079565b6108ec8233836110f1565b610844816110fc565b5f61081b60048361115a565b5f61082b60025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206109389083611188565b9392505050565b5f610949336108f5565b610986576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b5f600d5f815461099590612cf0565b918290555090506109a68582611193565b949350505050565b606061082b611240565b5f61081b600483611188565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610a0657504265ffffffffffff821610155b610a11575f5f610a36565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b610844828261124f565b5f81815260036020526040902060609061081b90610f75565b5f61082b600461125a565b610a7884848484611263565b50505050565b606061081b8261127b565b5f81815260036020526040812061081b9061125a565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610ae057504265ffffffffffff8216105b610b12576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610b36565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610b90576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161097d565b61085a6112df565b61084482826113d0565b5f610bac81610d30565b61085a611411565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152600b6020908152604080832093851683529290529081205460ff16610938565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610c1a81610d30565b61280a82610c2a5761141b610c2e565b61146e5b9050835f5b81811015610c7857610c706004888884818110610c5257610c52612d27565b9050602002016020810190610c6791906129db565b8563ffffffff16565b600101610c33565b50505050505050565b606060068054610c9090612d54565b80601f0160208091040260200160405190810160405280929190818152602001828054610cbc90612d54565b8015610d075780601f10610cde57610100808354040283529160200191610d07565b820191905f5260205f20905b815481529060010190602001808311610cea57829003601f168201915b5050505050905090565b5f610d1b82611530565b5061081b8261158d565b6108448282336115b6565b61085a81336115c3565b610d445f5f611648565b565b73ffffffffffffffffffffffffffffffffffffffff8216610d95576040517f64a0ae920000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b5f610da18383336117a1565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a78576040517f64283d7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8086166004830152602482018490528216604482015260640161097d565b81610e66576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61084482826117ad565b81158015610e98575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610f6b5760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610eec575065ffffffffffff8116155b80610eff57504265ffffffffffff821610155b15610f40576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161097d565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b61084482826117d1565b60605f6109388361182a565b5f610f8a610a9f565b610f9342611883565b610f9d9190612da5565b9050610fa982826118d2565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61081b82611530565b5f6110148261196d565b61101d42611883565b6110279190612da5565b90506110338282611648565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff82166110c9576040517f89c62b640000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b5073ffffffffffffffffffffffffffffffffffffffff165f9081526009602052604090205490565b6108688383836119b4565b5f6111085f835f6117a1565b905073ffffffffffffffffffffffffffffffffffffffff8116610844576040517f7e2732890000000000000000000000000000000000000000000000000000000081526004810183905260240161097d565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610938565b5f6109388383611a64565b73ffffffffffffffffffffffffffffffffffffffff82166111e2576040517f64a0ae920000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b5f6111ee83835f6117a1565b905073ffffffffffffffffffffffffffffffffffffffff811615610868576040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b606060078054610c9090612d54565b610844338383611a8a565b5f61081b825490565b61126e84848461085d565b610a783385858585611a95565b606061128682611530565b505f611290611c8b565b90505f8151116112ae5760405180602001604052805f815250610938565b806112b884611c9a565b6040516020016112c9929190612dda565b6040516020818303038152906040529392505050565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061132f57504265ffffffffffff821610155b15611370576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161097d565b6113985f61139360025473ffffffffffffffffffffffffffffffffffffffff1690565b611d56565b506113a35f83611d61565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b81611407576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108448282611d6c565b610d445f5f6118d2565b6114258282611d90565b156108445760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff81166114dd576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f726765000000000000000000000000000000000000000000000000000000600482015260240161097d565b6114e78282611db1565b156108445760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f5f61153b83611dd2565b905073ffffffffffffffffffffffffffffffffffffffff811661081b576040517f7e2732890000000000000000000000000000000000000000000000000000000081526004810184905260240161097d565b5f818152600a602052604081205473ffffffffffffffffffffffffffffffffffffffff1661081b565b6108688383836001611dfb565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610844576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161097d565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16801561171c574265ffffffffffff821610156116f3576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a0100000000000000000000000000000000000000000000000000000291909117905561171c565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f6109a6848484611e07565b5f828152602081905260409020600101546117c781610d30565b610a788383611d61565b73ffffffffffffffffffffffffffffffffffffffff81163314611820576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108688282611d56565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561187757602002820191905f5260205f20905b815481526020019060010190808311611863575b50505050509050919050565b5f65ffffffffffff8211156118ce576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161097d565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015610868576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611977610a9f565b90508065ffffffffffff168365ffffffffffff161161199f5761199a8382612dee565b610938565b61093865ffffffffffff841662069780611f79565b6119bf838383611f88565b6108685773ffffffffffffffffffffffffffffffffffffffff8316611a13576040517f7e2732890000000000000000000000000000000000000000000000000000000081526004810182905260240161097d565b6040517f177e802f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024810182905260440161097d565b5f825f018281548110611a7957611a79612d27565b905f5260205f200154905092915050565b610868838383611f94565b73ffffffffffffffffffffffffffffffffffffffff83163b15611c84576040517f150b7a0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063150b7a0290611b0a908890889087908790600401612e0c565b6020604051808303815f875af1925050508015611b62575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611b5f91810190612e66565b60015b611bef573d808015611b8f576040519150601f19603f3d011682016040523d82523d5f602084013e611b94565b606091505b5080515f03611be7576040517f64a0ae9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161097d565b805160208201fd5b7fffffffff0000000000000000000000000000000000000000000000000000000081167f150b7a020000000000000000000000000000000000000000000000000000000014611c82576040517f64a0ae9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161097d565b505b5050505050565b6060600c8054610c9090612d54565b60605f611ca683612090565b60010190505f8167ffffffffffffffff811115611cc557611cc5612a39565b6040519080825280601f01601f191660200182016040528015611cef576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084611cf957509392505050565b5f6109388383612171565b5f61093883836121a4565b5f82815260208190526040902060010154611d8681610d30565b610a788383611d56565b5f6109388373ffffffffffffffffffffffffffffffffffffffff84166121cf565b5f6109388373ffffffffffffffffffffffffffffffffffffffff84166122b2565b5f8181526008602052604081205473ffffffffffffffffffffffffffffffffffffffff1661081b565b610a78848484846122fe565b5f5f611e1284611dd2565b905073ffffffffffffffffffffffffffffffffffffffff831615611e3b57611e3b8184866110f1565b73ffffffffffffffffffffffffffffffffffffffff811615611eae57611e635f855f5f611dfb565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260096020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190555b73ffffffffffffffffffffffffffffffffffffffff851615611ef65773ffffffffffffffffffffffffffffffffffffffff85165f908152600960205260409020805460010190555b5f8481526008602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b5f828218828410028218610938565b5f6109a684848461249b565b73ffffffffffffffffffffffffffffffffffffffff8216611ff9576040517f5b08ba1800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161097d565b73ffffffffffffffffffffffffffffffffffffffff8381165f818152600b602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106120d8577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310612104576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061212257662386f26fc10000830492506010015b6305f5e100831061213a576305f5e100830492506008015b612710831061214e57612710830492506004015b60648310612160576064830492506002015b600a831061081b5760010192915050565b5f5f61217d8484612540565b90508015610938575f84815260036020526040902061219c9084611d90565b509392505050565b5f5f6121b084846125a1565b90508015610938575f84815260036020526040902061219c9084611db1565b5f81815260018301602052604081205480156122a9575f6121f1600183612e81565b85549091505f9061220490600190612e81565b9050808214612263575f865f01828154811061222257612222612d27565b905f5260205f200154905080875f01848154811061224257612242612d27565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061227457612274612e94565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061081b565b5f91505061081b565b5f8181526001830160205260408120546122f757508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561081b565b505f61081b565b808061231f575073ffffffffffffffffffffffffffffffffffffffff821615155b15612447575f61232e84611530565b905073ffffffffffffffffffffffffffffffffffffffff83161580159061238157508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561239457506123928184610bb4565b155b156123e3576040517fa9fbf51f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161097d565b811561244557838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f908152600a6020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f73ffffffffffffffffffffffffffffffffffffffff8316158015906109a657508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806124fb57506124fb8484610bb4565b806109a657508273ffffffffffffffffffffffffffffffffffffffff166125218361158d565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b5f82158015612569575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561259757600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b610938838361265f565b5f82612655575f6125c760025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612614576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b6109388383612718565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16156122f7575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161081b565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff166122f7575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556127a83390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161081b565b610d44612ec1565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461085a575f5ffd5b5f6020828403121561284f575f5ffd5b813561093881612812565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610938602083018461285a565b5f602082840312156128c8575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146128f2575f5ffd5b919050565b5f5f60408385031215612908575f5ffd5b612911836128cf565b946020939093013593505050565b5f5f5f60608486031215612931575f5ffd5b61293a846128cf565b9250612948602085016128cf565b929592945050506040919091013590565b5f5f6040838503121561296a575f5ffd5b8235915061297a602084016128cf565b90509250929050565b602080825282518282018190525f918401906040840190835b818110156129d057835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161299c565b509095945050505050565b5f602082840312156129eb575f5ffd5b610938826128cf565b5f60208284031215612a04575f5ffd5b813565ffffffffffff81168114610938575f5ffd5b5f5f60408385031215612a2a575f5ffd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612a75575f5ffd5b813567ffffffffffffffff811115612a8f57612a8f612a39565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612afb57612afb612a39565b604052818152838201602001851015612b12575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f60608486031215612b40575f5ffd5b612b49846128cf565b925060208401359150604084013567ffffffffffffffff811115612b6b575f5ffd5b612b7786828701612a66565b9150509250925092565b803580151581146128f2575f5ffd5b5f5f60408385031215612ba1575f5ffd5b612baa836128cf565b915061297a60208401612b81565b5f5f5f5f60808587031215612bcb575f5ffd5b612bd4856128cf565b9350612be2602086016128cf565b925060408501359150606085013567ffffffffffffffff811115612c04575f5ffd5b612c1087828801612a66565b91505092959194509250565b5f5f60408385031215612c2d575f5ffd5b612c36836128cf565b915061297a602084016128cf565b5f5f5f60408486031215612c56575f5ffd5b833567ffffffffffffffff811115612c6c575f5ffd5b8401601f81018613612c7c575f5ffd5b803567ffffffffffffffff811115612c92575f5ffd5b8660208260051b8401011115612ca6575f5ffd5b602091820194509250612cba908501612b81565b90509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612d2057612d20612cc3565b5060010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c90821680612d6857607f821691505b602082108103612d9f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b65ffffffffffff818116838216019081111561081b5761081b612cc3565b5f81518060208401855e5f93019283525090919050565b5f6109a6612de88386612dc3565b84612dc3565b65ffffffffffff828116828216039081111561081b5761081b612cc3565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152826040820152608060608201525f612e5c608083018461285a565b9695505050505050565b5f60208284031215612e76575f5ffd5b815161093881612812565b8181038181111561081b5761081b612cc3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220ebe8583a26b5e200eef9ea65d137d3d269169bded5a5bbf617f147f5ed983d9064736f6c634300081c0033",
}

// ERC721AutoIncr is an auto generated Go binding around an Ethereum contract.
type ERC721AutoIncr struct {
	abi abi.ABI
}

// NewERC721AutoIncr creates a new instance of ERC721AutoIncr.
func NewERC721AutoIncr() *ERC721AutoIncr {
	parsed, err := ERC721AutoIncrMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721AutoIncr{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721AutoIncr) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, string baseTokenURI) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, baseTokenURI string) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("", owner, forges, name, symbol, baseTokenURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC721AutoIncr *ERC721AutoIncr) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC721AutoIncr.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackMANAGERROLE() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC721AutoIncr.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721AutoIncr *ERC721AutoIncr) PackBalanceOf(owner common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC721AutoIncr.abi.Unpack("balanceOf", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 tokenID) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackBurnFrom(from common.Address, tokenID *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("burnFrom", from, tokenID)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) PackDefaultAdmin() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("defaultAdmin", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackDefaultAdminDelay() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC721AutoIncr.abi.Unpack("defaultAdminDelay", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC721AutoIncr.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("forgeByIndex", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackForgeCount() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC721AutoIncr.abi.Unpack("forgeCount", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackForges() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC721AutoIncr *ERC721AutoIncr) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("forges", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, err
}

// PackGetApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("getApproved", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetRoleAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC721AutoIncr *ERC721AutoIncr) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC721AutoIncr.abi.Unpack("getRoleAdmin", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("getRoleMember", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC721AutoIncr.abi.Unpack("getRoleMemberCount", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC721AutoIncr *ERC721AutoIncr) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("getRoleMembers", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC721AutoIncr.abi.Unpack("hasRole", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := eRC721AutoIncr.abi.Unpack("isApprovedForAll", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackIsForge(forge common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC721AutoIncr.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94d008ef.
//
// Solidity: function mint(address to, uint256 , bytes data) returns(uint256)
func (eRC721AutoIncr *ERC721AutoIncr) PackMint(to common.Address, arg1 *big.Int, data []byte) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("mint", to, arg1, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94d008ef.
//
// Solidity: function mint(address to, uint256 , bytes data) returns(uint256)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackMint(data []byte) (*big.Int, error) {
	out, err := eRC721AutoIncr.abi.Unpack("mint", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721AutoIncr *ERC721AutoIncr) PackName() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackName(data []byte) (string, error) {
	out, err := eRC721AutoIncr.abi.Unpack("name", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackOwner() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackOwnerOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := eRC721AutoIncr.abi.Unpack("ownerOf", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackPendingDefaultAdmin() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC721AutoIncr.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC721AutoIncr.abi.Unpack("pendingDefaultAdminDelay", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC721AutoIncr.abi.Unpack("supportsInterface", data)
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
func (eRC721AutoIncr *ERC721AutoIncr) PackSymbol() []byte {
	enc, err := eRC721AutoIncr.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC721AutoIncr.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTokenURI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721AutoIncr *ERC721AutoIncr) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackTokenURI(data []byte) (string, error) {
	out, err := eRC721AutoIncr.abi.Unpack("tokenURI", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721AutoIncr *ERC721AutoIncr) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721AutoIncr.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC721AutoIncrApproval represents a Approval event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrApproval) ContractEventName() string {
	return ERC721AutoIncrApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackApprovalEvent(log *types.Log) (*ERC721AutoIncrApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrApproval)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrApprovalForAll represents a ApprovalForAll event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrApprovalForAll) ContractEventName() string {
	return ERC721AutoIncrApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackApprovalForAllEvent(log *types.Log) (*ERC721AutoIncrApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrApprovalForAll)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC721AutoIncrDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC721AutoIncrDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC721AutoIncrDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC721AutoIncrDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC721AutoIncrDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC721AutoIncrDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC721AutoIncrDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC721AutoIncrDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrERC721BaseSetBaseURI represents a ERC721BaseSetBaseURI event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721BaseSetBaseURI struct {
	Before  string
	Current string
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrERC721BaseSetBaseURIEventName = "ERC721BaseSetBaseURI"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrERC721BaseSetBaseURI) ContractEventName() string {
	return ERC721AutoIncrERC721BaseSetBaseURIEventName
}

// UnpackERC721BaseSetBaseURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721BaseSetBaseURI(string before, string current)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721BaseSetBaseURIEvent(log *types.Log) (*ERC721AutoIncrERC721BaseSetBaseURI, error) {
	event := "ERC721BaseSetBaseURI"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrERC721BaseSetBaseURI)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrForgeAdded represents a ForgeAdded event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrForgeAdded) ContractEventName() string {
	return ERC721AutoIncrForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackForgeAddedEvent(log *types.Log) (*ERC721AutoIncrForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrForgeRemoved represents a ForgeRemoved event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrForgeRemoved) ContractEventName() string {
	return ERC721AutoIncrForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackForgeRemovedEvent(log *types.Log) (*ERC721AutoIncrForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrRoleAdminChanged represents a RoleAdminChanged event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrRoleAdminChanged) ContractEventName() string {
	return ERC721AutoIncrRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC721AutoIncrRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrRoleGranted represents a RoleGranted event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrRoleGranted) ContractEventName() string {
	return ERC721AutoIncrRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackRoleGrantedEvent(log *types.Log) (*ERC721AutoIncrRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrRoleRevoked represents a RoleRevoked event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrRoleRevoked) ContractEventName() string {
	return ERC721AutoIncrRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackRoleRevokedEvent(log *types.Log) (*ERC721AutoIncrRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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

// ERC721AutoIncrTransfer represents a Transfer event raised by the ERC721AutoIncr contract.
type ERC721AutoIncrTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721AutoIncrTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC721AutoIncrTransfer) ContractEventName() string {
	return ERC721AutoIncrTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackTransferEvent(log *types.Log) (*ERC721AutoIncrTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC721AutoIncr.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721AutoIncrTransfer)
	if len(log.Data) > 0 {
		if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721AutoIncr.abi.Events[event].Inputs {
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
func (eRC721AutoIncr *ERC721AutoIncr) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackERC721NonexistentTokenError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721AutoIncr.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC721AutoIncr.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC721AutoIncrAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC721AutoIncrAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC721AutoIncr *ERC721AutoIncr) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC721AutoIncrAccessControlBadConfirmation, error) {
	out := new(ERC721AutoIncrAccessControlBadConfirmation)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC721AutoIncrAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC721AutoIncrAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC721AutoIncrAccessControlEnforcedDefaultAdminDelay)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC721AutoIncrAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC721AutoIncr *ERC721AutoIncr) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC721AutoIncrAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC721AutoIncrAccessControlEnforcedDefaultAdminRules)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC721AutoIncrAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC721AutoIncrAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC721AutoIncrAccessControlInvalidDefaultAdmin)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC721AutoIncrAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC721AutoIncrAccessControlUnauthorizedAccount, error) {
	out := new(ERC721AutoIncrAccessControlUnauthorizedAccount)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func ERC721AutoIncrERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721IncorrectOwnerError(raw []byte) (*ERC721AutoIncrERC721IncorrectOwner, error) {
	out := new(ERC721AutoIncrERC721IncorrectOwner)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func ERC721AutoIncrERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721InsufficientApprovalError(raw []byte) (*ERC721AutoIncrERC721InsufficientApproval, error) {
	out := new(ERC721AutoIncrERC721InsufficientApproval)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721InvalidApprover represents a ERC721InvalidApprover error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func ERC721AutoIncrERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721InvalidApproverError(raw []byte) (*ERC721AutoIncrERC721InvalidApprover, error) {
	out := new(ERC721AutoIncrERC721InvalidApprover)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721InvalidOperator represents a ERC721InvalidOperator error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func ERC721AutoIncrERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721InvalidOperatorError(raw []byte) (*ERC721AutoIncrERC721InvalidOperator, error) {
	out := new(ERC721AutoIncrERC721InvalidOperator)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721InvalidOwner represents a ERC721InvalidOwner error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func ERC721AutoIncrERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721InvalidOwnerError(raw []byte) (*ERC721AutoIncrERC721InvalidOwner, error) {
	out := new(ERC721AutoIncrERC721InvalidOwner)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func ERC721AutoIncrERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721InvalidReceiverError(raw []byte) (*ERC721AutoIncrERC721InvalidReceiver, error) {
	out := new(ERC721AutoIncrERC721InvalidReceiver)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721InvalidSender represents a ERC721InvalidSender error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func ERC721AutoIncrERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721InvalidSenderError(raw []byte) (*ERC721AutoIncrERC721InvalidSender, error) {
	out := new(ERC721AutoIncrERC721InvalidSender)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrERC721NonexistentToken represents a ERC721NonexistentToken error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func ERC721AutoIncrERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackERC721NonexistentTokenError(raw []byte) (*ERC721AutoIncrERC721NonexistentToken, error) {
	out := new(ERC721AutoIncrERC721NonexistentToken)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC721AutoIncrSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC721AutoIncrSafeCastOverflowedUintDowncast, error) {
	out := new(ERC721AutoIncrSafeCastOverflowedUintDowncast)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC721AutoIncrTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackTokenBaseNullInputError(raw []byte) (*ERC721AutoIncrTokenBaseNullInput, error) {
	out := new(ERC721AutoIncrTokenBaseNullInput)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721AutoIncrTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC721AutoIncr contract.
type ERC721AutoIncrTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC721AutoIncrTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC721AutoIncr *ERC721AutoIncr) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC721AutoIncrTokenBaseOnlyForge, error) {
	out := new(ERC721AutoIncrTokenBaseOnlyForge)
	if err := eRC721AutoIncr.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
