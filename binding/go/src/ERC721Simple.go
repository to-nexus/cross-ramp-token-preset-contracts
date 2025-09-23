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

// ERC721SimpleMetaData contains all meta data concerning the ERC721Simple contract.
var ERC721SimpleMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"forges\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"beginDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDefaultAdminTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"changeDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultAdminDelayIncreaseWait\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"forgeByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forgeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"isForge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDefaultAdminDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollbackDefaultAdminDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"forges_\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setForges\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminDelayChangeCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"effectSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminDelayChangeScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultAdminTransferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"acceptSchedule\",\"type\":\"uint48\"}],\"name\":\"DefaultAdminTransferScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"before\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"current\",\"type\":\"string\"}],\"name\":\"ERC721BaseSetBaseURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forge\",\"type\":\"address\"}],\"name\":\"ForgeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"schedule\",\"type\":\"uint48\"}],\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"}],\"name\":\"AccessControlInvalidDefaultAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"field\",\"type\":\"bytes32\"}],\"name\":\"TokenBase__NullInput\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TokenBase__OnlyForge\",\"type\":\"error\"}]",
	ID:  "ERC721Simple",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161374638038061374683398101604081905261002e916104ff565b8484848484828286868162015180816001600160a01b03811661006b57604051636116401160e11b81525f60048201526024015b60405180910390fd5b600180546001600160d01b0316600160d01b65ffffffffffff8516021790556100945f826101ca565b505050506100c87faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c836101ca60201b60201c565b5080515f5b81811015610107576100ff60048483815181106100ec576100ec610640565b60200260200101516101de60201b60201c565b6001016100cd565b50505050816006908161011a91906106d8565b50600761012782826106d8565b50505082515f036101545760405163e1dea2ef60e01b8152636e616d6560e01b6004820152602401610062565b81515f036101805760405163e1dea2ef60e01b8152651cde5b589bdb60d21b6004820152602401610062565b80515f036101b25760405163e1dea2ef60e01b81526b62617365546f6b656e55524960a01b6004820152602401610062565b6101bb81610256565b5050505050505050505061085a565b5f6101d5838361029c565b90505b92915050565b6001600160a01b03811661020f5760405163e1dea2ef60e01b815264666f72676560d81b6004820152602401610062565b61021982826102cf565b15610252576040516001600160a01b038216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25b5050565b7ffcd3ec38ee6a55750e2501220e2c259a8309f45d101943d2f1c9a40ecfd6b3a6600c826040516102889291906107c0565b60405180910390a1600c61025282826106d8565b5f806102a884846102e3565b905080156101d5575f8481526003602052604090206102c790846102cf565b509392505050565b5f6101d5836001600160a01b038416610349565b5f8261033f575f6102fc6002546001600160a01b031690565b6001600160a01b03161461032357604051631fe1e13d60e11b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0384161790555b6101d58383610395565b5f81815260018301602052604081205461038e57508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556101d8565b505f6101d8565b5f828152602081815260408083206001600160a01b038516845290915281205460ff1661038e575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556103ed3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d8565b80516001600160a01b038116811461044b575f5ffd5b919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561048c5761048c610450565b604052919050565b5f82601f8301126104a3575f5ffd5b81516001600160401b038111156104bc576104bc610450565b6104cf601f8201601f1916602001610464565b8181528460208386010111156104e3575f5ffd5b8160208501602083015e5f918101602001919091529392505050565b5f5f5f5f5f60a08688031215610513575f5ffd5b61051c86610435565b60208701519095506001600160401b03811115610537575f5ffd5b8601601f81018813610547575f5ffd5b80516001600160401b0381111561056057610560610450565b8060051b61057060208201610464565b9182526020818401810192908101908b84111561058b575f5ffd5b6020850194505b838510156105b4576105a385610435565b825260209485019490910190610592565b60408b0151909850935050506001600160401b0382111590506105d5575f5ffd5b6105e188828901610494565b606088015190945090506001600160401b038111156105fe575f5ffd5b61060a88828901610494565b608088015190935090506001600160401b03811115610627575f5ffd5b61063388828901610494565b9150509295509295909350565b634e487b7160e01b5f52603260045260245ffd5b600181811c9082168061066857607f821691505b60208210810361068657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156106d357805f5260205f20601f840160051c810160208510156106b15750805b601f840160051c820191505b818110156106d0575f81556001016106bd565b50505b505050565b81516001600160401b038111156106f1576106f1610450565b610705816106ff8454610654565b8461068c565b6020601f821160018114610737575f83156107205750848201515b5f19600385901b1c1916600184901b1784556106d0565b5f84815260208120601f198516915b828110156107665787850151825560209485019460019092019101610746565b508482101561078357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b604081525f5f84546107d181610654565b806040860152600182165f81146107ef576001811461080b5761083c565b60ff1983166060870152606082151560051b870101935061083c565b875f5260205f205f5b8381101561083357815488820160600152600190910190602001610814565b87016060019450505b50505082810360208401526108518185610792565b95945050505050565b612edf806108675f395ff3fe608060405234801561000f575f5ffd5b50600436106102c2575f3560e01c80639010d07c1161017c578063b88d4fde116100dd578063cf6eefb711610093578063e985e9c51161006e578063e985e9c51461062d578063ec87621c14610640578063fe2df3e814610667575f5ffd5b8063cf6eefb7146105c6578063d547741f14610612578063d602b9fd14610625575f5ffd5b8063ca15c873116100c3578063ca15c873146105a3578063cc8463c8146105b6578063cefc1429146105be575f5ffd5b8063b88d4fde1461057d578063c87b56dd14610590575f5ffd5b8063a1eda53c11610132578063a22cb46511610118578063a22cb4651461054f578063a3246ad314610562578063a40283d514610575575f5ffd5b8063a1eda53c14610521578063a217fddf14610548575f5ffd5b806394d008ef1161016257806394d008ef146104f357806395d89b41146105065780639ca92df91461050e575f5ffd5b80639010d07c1461049d57806391d14854146104b0575f5ffd5b806342842e0e1161022657806370a08231116101dc57806384ef8ffc116101c257806384ef8ffc1461046457806387f45353146104825780638da5cb5b14610495575f5ffd5b806370a082311461043e57806379cc679014610451575f5ffd5b8063634e93da1161020c578063634e93da146104055780636352211e14610418578063649a5ec71461042b575f5ffd5b806342842e0e146103dd5780635c4e62c4146103f0575f5ffd5b80630aa6220b1161027b578063248a9ca311610261578063248a9ca3146103875780632f2ff15d146103b757806336568abe146103ca575f5ffd5b80630aa6220b1461036c57806323b872dd14610374575f5ffd5b806306fdde03116102ab57806306fdde031461030a578063081812fc1461031f578063095ea7b314610357575f5ffd5b806301ffc9a7146102c6578063022d63fb146102ee575b5f5ffd5b6102d96102d4366004612831565b61067a565b60405190151581526020015b60405180910390f35b620697805b60405165ffffffffffff90911681526020016102e5565b610312610821565b6040516102e59190612898565b61033261032d3660046128aa565b610830565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102e5565b61036a6103653660046128e9565b61083a565b005b61036a610848565b61036a610382366004612911565b61085d565b6103a96103953660046128aa565b5f9081526020819052604090206001015490565b6040519081526020016102e5565b61036a6103c536600461294b565b61086d565b61036a6103d836600461294b565b610877565b61036a6103eb366004612911565b610881565b6103f861089b565b6040516102e59190612975565b61036a6104133660046129cd565b6108a7565b6103326104263660046128aa565b6108ba565b61036a6104393660046129e6565b6108c4565b6103a961044c3660046129cd565b6108d7565b61036a61045f3660046128e9565b6108e1565b60025473ffffffffffffffffffffffffffffffffffffffff16610332565b6102d96104903660046129cd565b6108f5565b610332610901565b6103326104ab366004612a0b565b610921565b6102d96104be36600461294b565b5f9182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6103a9610501366004612b20565b61093f565b610312610998565b61033261051c3660046128aa565b6109a2565b6105296109ae565b6040805165ffffffffffff9384168152929091166020830152016102e5565b6103a95f81565b61036a61055d366004612b82565b610a28565b6103f86105703660046128aa565b610a32565b6103a9610a4b565b61036a61058b366004612baa565b610a56565b61031261059e3660046128aa565b610a68565b6103a96105b13660046128aa565b610a73565b6102f3610a89565b61036a610b26565b6001546040805173ffffffffffffffffffffffffffffffffffffffff831681527401000000000000000000000000000000000000000090920465ffffffffffff166020830152016102e5565b61036a61062036600461294b565b610b82565b61036a610b8c565b6102d961063b366004612c0e565b610b9e565b6103a97faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c81565b61036a610675366004612c36565b610bda565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167ff0a0429100000000000000000000000000000000000000000000000000000000148061070c57507fffffffff0000000000000000000000000000000000000000000000000000000082167fed1c6f7f00000000000000000000000000000000000000000000000000000000145b8061075857507fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd00000000000000000000000000000000000000000000000000000000145b806107a457507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b806107cf57507fffffffff000000000000000000000000000000000000000000000000000000008216155b8061081b57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b606061082b610c6b565b905090565b5f61081b82610cfb565b6108448282610d0f565b5050565b5f61085281610d1a565b61085a610d24565b50565b610868838383610d30565b505050565b6108448282610e19565b6108448282610e5a565b61086883838360405180602001604052805f815250610a56565b606061082b6004610f5f565b5f6108b181610d1a565b61084482610f6b565b5f61081b82610fea565b5f6108ce81610d1a565b61084482610ff4565b5f61081b82611063565b6108ec8233836110db565b610844816110e6565b5f61081b600483611144565b5f61082b60025473ffffffffffffffffffffffffffffffffffffffff1690565b5f8281526003602052604081206109389083611172565b9392505050565b5f610949336108f5565b610986576040517f25a9dbc10000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b610990848461117d565b509092915050565b606061082b61122a565b5f61081b600483611172565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff1680151580156109f057504265ffffffffffff821610155b6109fb575f5f610a20565b60025474010000000000000000000000000000000000000000900465ffffffffffff16815b915091509091565b6108448282611239565b5f81815260036020526040902060609061081b90610f5f565b5f61082b6004611244565b610a628484848461124d565b50505050565b606061081b82611265565b5f81815260036020526040812061081b90611244565b6002545f907a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015158015610aca57504265ffffffffffff8216105b610afc576001547a010000000000000000000000000000000000000000000000000000900465ffffffffffff16610b20565b60025474010000000000000000000000000000000000000000900465ffffffffffff165b91505090565b60015473ffffffffffffffffffffffffffffffffffffffff16338114610b7a576040517fc22c802200000000000000000000000000000000000000000000000000000000815233600482015260240161097d565b61085a6112c9565b61084482826113ba565b5f610b9681610d1a565b61085a6113fb565b73ffffffffffffffffffffffffffffffffffffffff8083165f908152600b6020908152604080832093851683529290529081205460ff16610938565b7faf290d8680820aad922855f39b306097b20e28774d6c1ad35a20325630c3a02c610c0481610d1a565b6127fc82610c1457611405610c18565b6114585b9050835f5b81811015610c6257610c5a6004888884818110610c3c57610c3c612cb5565b9050602002016020810190610c5191906129cd565b8563ffffffff16565b600101610c1d565b50505050505050565b606060068054610c7a90612ce2565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca690612ce2565b8015610cf15780601f10610cc857610100808354040283529160200191610cf1565b820191905f5260205f20905b815481529060010190602001808311610cd457829003601f168201915b5050505050905090565b5f610d058261151a565b5061081b82611577565b6108448282336115a0565b61085a81336115ad565b610d2e5f5f611632565b565b73ffffffffffffffffffffffffffffffffffffffff8216610d7f576040517f64a0ae920000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b5f610d8b83833361178b565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a62576040517f64283d7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8086166004830152602482018490528216604482015260640161097d565b81610e50576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610844828261179f565b81158015610e82575060025473ffffffffffffffffffffffffffffffffffffffff8281169116145b15610f555760015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1681151580610ed6575065ffffffffffff8116155b80610ee957504265ffffffffffff821610155b15610f2a576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161097d565b5050600180547fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff1690555b61084482826117c3565b60605f6109388361181c565b5f610f74610a89565b610f7d42611875565b610f879190612d60565b9050610f9382826118c4565b60405165ffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff8316907f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed69060200160405180910390a25050565b5f61081b8261151a565b5f610ffe8261195f565b61100742611875565b6110119190612d60565b905061101d8282611632565b6040805165ffffffffffff8085168252831660208201527ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b910160405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff82166110b3576040517f89c62b640000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b5073ffffffffffffffffffffffffffffffffffffffff165f9081526009602052604090205490565b6108688383836119a6565b5f6110f25f835f61178b565b905073ffffffffffffffffffffffffffffffffffffffff8116610844576040517f7e2732890000000000000000000000000000000000000000000000000000000081526004810183905260240161097d565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610938565b5f6109388383611a56565b73ffffffffffffffffffffffffffffffffffffffff82166111cc576040517f64a0ae920000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b5f6111d883835f61178b565b905073ffffffffffffffffffffffffffffffffffffffff811615610868576040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081525f600482015260240161097d565b606060078054610c7a90612ce2565b610844338383611a7c565b5f61081b825490565b61125884848461085d565b610a623385858585611a87565b60606112708261151a565b505f61127a611c7d565b90505f8151116112985760405180602001604052805f815250610938565b806112a284611c8c565b6040516020016112b3929190612d95565b6040516020818303038152906040529392505050565b60015473ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900465ffffffffffff1680158061131957504265ffffffffffff821610155b1561135a576040517f19ca5ebb00000000000000000000000000000000000000000000000000000000815265ffffffffffff8216600482015260240161097d565b6113825f61137d60025473ffffffffffffffffffffffffffffffffffffffff1690565b611d48565b5061138d5f83611d53565b5050600180547fffffffffffff000000000000000000000000000000000000000000000000000016905550565b816113f1576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108448282611d5e565b610d2e5f5f6118c4565b61140f8282611d82565b156108445760405173ffffffffffffffffffffffffffffffffffffffff8216907f24de35b14425bf91d8515f01462b1ca88c604c24504db133d962403bd3ebf15f905f90a25050565b73ffffffffffffffffffffffffffffffffffffffff81166114c7576040517fe1dea2ef0000000000000000000000000000000000000000000000000000000081527f666f726765000000000000000000000000000000000000000000000000000000600482015260240161097d565b6114d18282611da3565b156108445760405173ffffffffffffffffffffffffffffffffffffffff8216907f239043ab1d209fe1fa727e5e8803207aa1fffadcb28f962777c853c7c0b986db905f90a25050565b5f5f61152583611dc4565b905073ffffffffffffffffffffffffffffffffffffffff811661081b576040517f7e2732890000000000000000000000000000000000000000000000000000000081526004810184905260240161097d565b5f818152600a602052604081205473ffffffffffffffffffffffffffffffffffffffff1661081b565b6108688383836001611ded565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610844576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161097d565b6002547a010000000000000000000000000000000000000000000000000000900465ffffffffffff168015611706574265ffffffffffff821610156116dd576002546001805479ffffffffffffffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000090920465ffffffffffff167a01000000000000000000000000000000000000000000000000000002919091179055611706565b6040517f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5905f90a15b506002805473ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000065ffffffffffff9485160279ffffffffffffffffffffffffffffffffffffffffffffffffffff16177a0100000000000000000000000000000000000000000000000000009290931691909102919091179055565b5f611797848484611df9565b949350505050565b5f828152602081905260409020600101546117b981610d1a565b610a628383611d53565b73ffffffffffffffffffffffffffffffffffffffff81163314611812576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108688282611d48565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561186957602002820191905f5260205f20905b815481526020019060010190808311611855575b50505050509050919050565b5f65ffffffffffff8211156118c0576040517f6dfcc650000000000000000000000000000000000000000000000000000000008152603060048201526024810183905260440161097d565b5090565b600180547401000000000000000000000000000000000000000065ffffffffffff84811682027fffffffffffff0000000000000000000000000000000000000000000000000000841673ffffffffffffffffffffffffffffffffffffffff881617179093559004168015610868576040517f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109905f90a1505050565b5f5f611969610a89565b90508065ffffffffffff168365ffffffffffff16116119915761198c8382612da9565b610938565b61093865ffffffffffff841662069780611f6b565b6119b1838383611f7a565b6108685773ffffffffffffffffffffffffffffffffffffffff8316611a05576040517f7e2732890000000000000000000000000000000000000000000000000000000081526004810182905260240161097d565b6040517f177e802f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024810182905260440161097d565b5f825f018281548110611a6b57611a6b612cb5565b905f5260205f200154905092915050565b610868838383611f86565b73ffffffffffffffffffffffffffffffffffffffff83163b15611c76576040517f150b7a0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063150b7a0290611afc908890889087908790600401612dc7565b6020604051808303815f875af1925050508015611b54575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611b5191810190612e21565b60015b611be1573d808015611b81576040519150601f19603f3d011682016040523d82523d5f602084013e611b86565b606091505b5080515f03611bd9576040517f64a0ae9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161097d565b805160208201fd5b7fffffffff0000000000000000000000000000000000000000000000000000000081167f150b7a020000000000000000000000000000000000000000000000000000000014611c74576040517f64a0ae9200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161097d565b505b5050505050565b6060600c8054610c7a90612ce2565b60605f611c9883612082565b60010190505f8167ffffffffffffffff811115611cb757611cb7612a2b565b6040519080825280601f01601f191660200182016040528015611ce1576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084611ceb57509392505050565b5f6109388383612163565b5f6109388383612196565b5f82815260208190526040902060010154611d7881610d1a565b610a628383611d48565b5f6109388373ffffffffffffffffffffffffffffffffffffffff84166121c1565b5f6109388373ffffffffffffffffffffffffffffffffffffffff84166122a4565b5f8181526008602052604081205473ffffffffffffffffffffffffffffffffffffffff1661081b565b610a62848484846122f0565b5f5f611e0484611dc4565b905073ffffffffffffffffffffffffffffffffffffffff831615611e2d57611e2d8184866110db565b73ffffffffffffffffffffffffffffffffffffffff811615611ea057611e555f855f5f611ded565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260096020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190555b73ffffffffffffffffffffffffffffffffffffffff851615611ee85773ffffffffffffffffffffffffffffffffffffffff85165f908152600960205260409020805460010190555b5f8481526008602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b5f828218828410028218610938565b5f61179784848461248d565b73ffffffffffffffffffffffffffffffffffffffff8216611feb576040517f5b08ba1800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161097d565b73ffffffffffffffffffffffffffffffffffffffff8381165f818152600b602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106120ca577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106120f6576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061211457662386f26fc10000830492506010015b6305f5e100831061212c576305f5e100830492506008015b612710831061214057612710830492506004015b60648310612152576064830492506002015b600a831061081b5760010192915050565b5f5f61216f8484612532565b90508015610938575f84815260036020526040902061218e9084611d82565b509392505050565b5f5f6121a28484612593565b90508015610938575f84815260036020526040902061218e9084611da3565b5f818152600183016020526040812054801561229b575f6121e3600183612e3c565b85549091505f906121f690600190612e3c565b9050808214612255575f865f01828154811061221457612214612cb5565b905f5260205f200154905080875f01848154811061223457612234612cb5565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061226657612266612e4f565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061081b565b5f91505061081b565b5f8181526001830160205260408120546122e957508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561081b565b505f61081b565b8080612311575073ffffffffffffffffffffffffffffffffffffffff821615155b15612439575f6123208461151a565b905073ffffffffffffffffffffffffffffffffffffffff83161580159061237357508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561238657506123848184610b9e565b155b156123d5576040517fa9fbf51f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161097d565b811561243757838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f908152600a6020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b5f73ffffffffffffffffffffffffffffffffffffffff83161580159061179757508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806124ed57506124ed8484610b9e565b8061179757508273ffffffffffffffffffffffffffffffffffffffff1661251383611577565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b5f8215801561255b575060025473ffffffffffffffffffffffffffffffffffffffff8381169116145b1561258957600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555b6109388383612651565b5f82612647575f6125b960025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614612606576040517f3fc3c27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790555b610938838361270a565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16156122e9575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161081b565b5f8281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff166122e9575f8381526020818152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561279a3390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161081b565b610d2e612e7c565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461085a575f5ffd5b5f60208284031215612841575f5ffd5b813561093881612804565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f610938602083018461284c565b5f602082840312156128ba575f5ffd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146128e4575f5ffd5b919050565b5f5f604083850312156128fa575f5ffd5b612903836128c1565b946020939093013593505050565b5f5f5f60608486031215612923575f5ffd5b61292c846128c1565b925061293a602085016128c1565b929592945050506040919091013590565b5f5f6040838503121561295c575f5ffd5b8235915061296c602084016128c1565b90509250929050565b602080825282518282018190525f918401906040840190835b818110156129c257835173ffffffffffffffffffffffffffffffffffffffff1683526020938401939092019160010161298e565b509095945050505050565b5f602082840312156129dd575f5ffd5b610938826128c1565b5f602082840312156129f6575f5ffd5b813565ffffffffffff81168114610938575f5ffd5b5f5f60408385031215612a1c575f5ffd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112612a67575f5ffd5b813567ffffffffffffffff811115612a8157612a81612a2b565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff82111715612aed57612aed612a2b565b604052818152838201602001851015612b04575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f60608486031215612b32575f5ffd5b612b3b846128c1565b925060208401359150604084013567ffffffffffffffff811115612b5d575f5ffd5b612b6986828701612a58565b9150509250925092565b803580151581146128e4575f5ffd5b5f5f60408385031215612b93575f5ffd5b612b9c836128c1565b915061296c60208401612b73565b5f5f5f5f60808587031215612bbd575f5ffd5b612bc6856128c1565b9350612bd4602086016128c1565b925060408501359150606085013567ffffffffffffffff811115612bf6575f5ffd5b612c0287828801612a58565b91505092959194509250565b5f5f60408385031215612c1f575f5ffd5b612c28836128c1565b915061296c602084016128c1565b5f5f5f60408486031215612c48575f5ffd5b833567ffffffffffffffff811115612c5e575f5ffd5b8401601f81018613612c6e575f5ffd5b803567ffffffffffffffff811115612c84575f5ffd5b8660208260051b8401011115612c98575f5ffd5b602091820194509250612cac908501612b73565b90509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b600181811c90821680612cf657607f821691505b602082108103612d2d577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b65ffffffffffff818116838216019081111561081b5761081b612d33565b5f81518060208401855e5f93019283525090919050565b5f611797612da38386612d7e565b84612d7e565b65ffffffffffff828116828216039081111561081b5761081b612d33565b73ffffffffffffffffffffffffffffffffffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152826040820152608060608201525f612e17608083018461284c565b9695505050505050565b5f60208284031215612e31575f5ffd5b815161093881612804565b8181038181111561081b5761081b612d33565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220595a8d3ce1ee907ee7d448b5809301f50af267df456fe2ea94022cb8c11692b464736f6c634300081c0033",
}

// ERC721Simple is an auto generated Go binding around an Ethereum contract.
type ERC721Simple struct {
	abi abi.ABI
}

// NewERC721Simple creates a new instance of ERC721Simple.
func NewERC721Simple() *ERC721Simple {
	parsed, err := ERC721SimpleMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ERC721Simple{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ERC721Simple) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address owner, address[] forges, string name, string symbol, string baseTokenURI) returns()
func (eRC721Simple *ERC721Simple) PackConstructor(owner common.Address, forges []common.Address, name string, symbol string, baseTokenURI string) []byte {
	enc, err := eRC721Simple.abi.Pack("", owner, forges, name, symbol, baseTokenURI)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDEFAULTADMINROLE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC721Simple *ERC721Simple) PackDEFAULTADMINROLE() []byte {
	enc, err := eRC721Simple.abi.Pack("DEFAULT_ADMIN_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDEFAULTADMINROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (eRC721Simple *ERC721Simple) UnpackDEFAULTADMINROLE(data []byte) ([32]byte, error) {
	out, err := eRC721Simple.abi.Unpack("DEFAULT_ADMIN_ROLE", data)
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
func (eRC721Simple *ERC721Simple) PackMANAGERROLE() []byte {
	enc, err := eRC721Simple.abi.Pack("MANAGER_ROLE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMANAGERROLE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec87621c.
//
// Solidity: function MANAGER_ROLE() view returns(bytes32)
func (eRC721Simple *ERC721Simple) UnpackMANAGERROLE(data []byte) ([32]byte, error) {
	out, err := eRC721Simple.abi.Unpack("MANAGER_ROLE", data)
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
func (eRC721Simple *ERC721Simple) PackAcceptDefaultAdminTransfer() []byte {
	enc, err := eRC721Simple.abi.Pack("acceptDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (eRC721Simple *ERC721Simple) PackApprove(to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("approve", to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721Simple *ERC721Simple) PackBalanceOf(owner common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("balanceOf", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (eRC721Simple *ERC721Simple) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := eRC721Simple.abi.Unpack("balanceOf", data)
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
func (eRC721Simple *ERC721Simple) PackBeginDefaultAdminTransfer(newAdmin common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("beginDefaultAdminTransfer", newAdmin)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBurnFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 tokenID) returns()
func (eRC721Simple *ERC721Simple) PackBurnFrom(from common.Address, tokenID *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("burnFrom", from, tokenID)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelDefaultAdminTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd602b9fd.
//
// Solidity: function cancelDefaultAdminTransfer() returns()
func (eRC721Simple *ERC721Simple) PackCancelDefaultAdminTransfer() []byte {
	enc, err := eRC721Simple.abi.Pack("cancelDefaultAdminTransfer")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackChangeDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x649a5ec7.
//
// Solidity: function changeDefaultAdminDelay(uint48 newDelay) returns()
func (eRC721Simple *ERC721Simple) PackChangeDefaultAdminDelay(newDelay *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("changeDefaultAdminDelay", newDelay)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDefaultAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC721Simple *ERC721Simple) PackDefaultAdmin() []byte {
	enc, err := eRC721Simple.abi.Pack("defaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84ef8ffc.
//
// Solidity: function defaultAdmin() view returns(address)
func (eRC721Simple *ERC721Simple) UnpackDefaultAdmin(data []byte) (common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("defaultAdmin", data)
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
func (eRC721Simple *ERC721Simple) PackDefaultAdminDelay() []byte {
	enc, err := eRC721Simple.abi.Pack("defaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc8463c8.
//
// Solidity: function defaultAdminDelay() view returns(uint48)
func (eRC721Simple *ERC721Simple) UnpackDefaultAdminDelay(data []byte) (*big.Int, error) {
	out, err := eRC721Simple.abi.Unpack("defaultAdminDelay", data)
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
func (eRC721Simple *ERC721Simple) PackDefaultAdminDelayIncreaseWait() []byte {
	enc, err := eRC721Simple.abi.Pack("defaultAdminDelayIncreaseWait")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDefaultAdminDelayIncreaseWait is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x022d63fb.
//
// Solidity: function defaultAdminDelayIncreaseWait() view returns(uint48)
func (eRC721Simple *ERC721Simple) UnpackDefaultAdminDelayIncreaseWait(data []byte) (*big.Int, error) {
	out, err := eRC721Simple.abi.Unpack("defaultAdminDelayIncreaseWait", data)
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
func (eRC721Simple *ERC721Simple) PackForgeByIndex(index *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("forgeByIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeByIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ca92df9.
//
// Solidity: function forgeByIndex(uint256 index) view returns(address)
func (eRC721Simple *ERC721Simple) UnpackForgeByIndex(data []byte) (common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("forgeByIndex", data)
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
func (eRC721Simple *ERC721Simple) PackForgeCount() []byte {
	enc, err := eRC721Simple.abi.Pack("forgeCount")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForgeCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa40283d5.
//
// Solidity: function forgeCount() view returns(uint256)
func (eRC721Simple *ERC721Simple) UnpackForgeCount(data []byte) (*big.Int, error) {
	out, err := eRC721Simple.abi.Unpack("forgeCount", data)
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
func (eRC721Simple *ERC721Simple) PackForges() []byte {
	enc, err := eRC721Simple.abi.Pack("forges")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackForges is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c4e62c4.
//
// Solidity: function forges() view returns(address[])
func (eRC721Simple *ERC721Simple) UnpackForges(data []byte) ([]common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("forges", data)
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
func (eRC721Simple *ERC721Simple) PackGetApproved(tokenId *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("getApproved", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (eRC721Simple *ERC721Simple) UnpackGetApproved(data []byte) (common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("getApproved", data)
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
func (eRC721Simple *ERC721Simple) PackGetRoleAdmin(role [32]byte) []byte {
	enc, err := eRC721Simple.abi.Pack("getRoleAdmin", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (eRC721Simple *ERC721Simple) UnpackGetRoleAdmin(data []byte) ([32]byte, error) {
	out, err := eRC721Simple.abi.Unpack("getRoleAdmin", data)
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
func (eRC721Simple *ERC721Simple) PackGetRoleMember(role [32]byte, index *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("getRoleMember", role, index)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMember is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (eRC721Simple *ERC721Simple) UnpackGetRoleMember(data []byte) (common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("getRoleMember", data)
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
func (eRC721Simple *ERC721Simple) PackGetRoleMemberCount(role [32]byte) []byte {
	enc, err := eRC721Simple.abi.Pack("getRoleMemberCount", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMemberCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (eRC721Simple *ERC721Simple) UnpackGetRoleMemberCount(data []byte) (*big.Int, error) {
	out, err := eRC721Simple.abi.Unpack("getRoleMemberCount", data)
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
func (eRC721Simple *ERC721Simple) PackGetRoleMembers(role [32]byte) []byte {
	enc, err := eRC721Simple.abi.Pack("getRoleMembers", role)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetRoleMembers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (eRC721Simple *ERC721Simple) UnpackGetRoleMembers(data []byte) ([]common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("getRoleMembers", data)
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
func (eRC721Simple *ERC721Simple) PackGrantRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("grantRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHasRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC721Simple *ERC721Simple) PackHasRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("hasRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHasRole is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (eRC721Simple *ERC721Simple) UnpackHasRole(data []byte) (bool, error) {
	out, err := eRC721Simple.abi.Unpack("hasRole", data)
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
func (eRC721Simple *ERC721Simple) PackIsApprovedForAll(owner common.Address, operator common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("isApprovedForAll", owner, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (eRC721Simple *ERC721Simple) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := eRC721Simple.abi.Unpack("isApprovedForAll", data)
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
func (eRC721Simple *ERC721Simple) PackIsForge(forge common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("isForge", forge)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsForge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87f45353.
//
// Solidity: function isForge(address forge) view returns(bool)
func (eRC721Simple *ERC721Simple) UnpackIsForge(data []byte) (bool, error) {
	out, err := eRC721Simple.abi.Unpack("isForge", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94d008ef.
//
// Solidity: function mint(address to, uint256 tokenID, bytes data) returns(uint256)
func (eRC721Simple *ERC721Simple) PackMint(to common.Address, tokenID *big.Int, data []byte) []byte {
	enc, err := eRC721Simple.abi.Pack("mint", to, tokenID, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94d008ef.
//
// Solidity: function mint(address to, uint256 tokenID, bytes data) returns(uint256)
func (eRC721Simple *ERC721Simple) UnpackMint(data []byte) (*big.Int, error) {
	out, err := eRC721Simple.abi.Unpack("mint", data)
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
func (eRC721Simple *ERC721Simple) PackName() []byte {
	enc, err := eRC721Simple.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (eRC721Simple *ERC721Simple) UnpackName(data []byte) (string, error) {
	out, err := eRC721Simple.abi.Unpack("name", data)
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
func (eRC721Simple *ERC721Simple) PackOwner() []byte {
	enc, err := eRC721Simple.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (eRC721Simple *ERC721Simple) UnpackOwner(data []byte) (common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("owner", data)
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
func (eRC721Simple *ERC721Simple) PackOwnerOf(tokenId *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("ownerOf", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwnerOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (eRC721Simple *ERC721Simple) UnpackOwnerOf(data []byte) (common.Address, error) {
	out, err := eRC721Simple.abi.Unpack("ownerOf", data)
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
func (eRC721Simple *ERC721Simple) PackPendingDefaultAdmin() []byte {
	enc, err := eRC721Simple.abi.Pack("pendingDefaultAdmin")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdmin is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6eefb7.
//
// Solidity: function pendingDefaultAdmin() view returns(address newAdmin, uint48 schedule)
func (eRC721Simple *ERC721Simple) UnpackPendingDefaultAdmin(data []byte) (PendingDefaultAdminOutput, error) {
	out, err := eRC721Simple.abi.Unpack("pendingDefaultAdmin", data)
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
func (eRC721Simple *ERC721Simple) PackPendingDefaultAdminDelay() []byte {
	enc, err := eRC721Simple.abi.Pack("pendingDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPendingDefaultAdminDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1eda53c.
//
// Solidity: function pendingDefaultAdminDelay() view returns(uint48 newDelay, uint48 schedule)
func (eRC721Simple *ERC721Simple) UnpackPendingDefaultAdminDelay(data []byte) (PendingDefaultAdminDelayOutput, error) {
	out, err := eRC721Simple.abi.Unpack("pendingDefaultAdminDelay", data)
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
func (eRC721Simple *ERC721Simple) PackRenounceRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("renounceRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRevokeRole is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (eRC721Simple *ERC721Simple) PackRevokeRole(role [32]byte, account common.Address) []byte {
	enc, err := eRC721Simple.abi.Pack("revokeRole", role, account)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRollbackDefaultAdminDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0aa6220b.
//
// Solidity: function rollbackDefaultAdminDelay() returns()
func (eRC721Simple *ERC721Simple) PackRollbackDefaultAdminDelay() []byte {
	enc, err := eRC721Simple.abi.Pack("rollbackDefaultAdminDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (eRC721Simple *ERC721Simple) PackSafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("safeTransferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeTransferFrom0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (eRC721Simple *ERC721Simple) PackSafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) []byte {
	enc, err := eRC721Simple.abi.Pack("safeTransferFrom0", from, to, tokenId, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (eRC721Simple *ERC721Simple) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := eRC721Simple.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetForges is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe2df3e8.
//
// Solidity: function setForges(address[] forges_, bool add) returns()
func (eRC721Simple *ERC721Simple) PackSetForges(forges []common.Address, add bool) []byte {
	enc, err := eRC721Simple.abi.Pack("setForges", forges, add)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721Simple *ERC721Simple) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := eRC721Simple.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (eRC721Simple *ERC721Simple) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := eRC721Simple.abi.Unpack("supportsInterface", data)
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
func (eRC721Simple *ERC721Simple) PackSymbol() []byte {
	enc, err := eRC721Simple.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (eRC721Simple *ERC721Simple) UnpackSymbol(data []byte) (string, error) {
	out, err := eRC721Simple.abi.Unpack("symbol", data)
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
func (eRC721Simple *ERC721Simple) PackTokenURI(tokenId *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("tokenURI", tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenURI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (eRC721Simple *ERC721Simple) UnpackTokenURI(data []byte) (string, error) {
	out, err := eRC721Simple.abi.Unpack("tokenURI", data)
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
func (eRC721Simple *ERC721Simple) PackTransferFrom(from common.Address, to common.Address, tokenId *big.Int) []byte {
	enc, err := eRC721Simple.abi.Pack("transferFrom", from, to, tokenId)
	if err != nil {
		panic(err)
	}
	return enc
}

// ERC721SimpleApproval represents a Approval event raised by the ERC721Simple contract.
type ERC721SimpleApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleApproval) ContractEventName() string {
	return ERC721SimpleApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (eRC721Simple *ERC721Simple) UnpackApprovalEvent(log *types.Log) (*ERC721SimpleApproval, error) {
	event := "Approval"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleApproval)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleApprovalForAll represents a ApprovalForAll event raised by the ERC721Simple contract.
type ERC721SimpleApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleApprovalForAll) ContractEventName() string {
	return ERC721SimpleApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (eRC721Simple *ERC721Simple) UnpackApprovalForAllEvent(log *types.Log) (*ERC721SimpleApprovalForAll, error) {
	event := "ApprovalForAll"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleApprovalForAll)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleDefaultAdminDelayChangeCanceled represents a DefaultAdminDelayChangeCanceled event raised by the ERC721Simple contract.
type ERC721SimpleDefaultAdminDelayChangeCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleDefaultAdminDelayChangeCanceledEventName = "DefaultAdminDelayChangeCanceled"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleDefaultAdminDelayChangeCanceled) ContractEventName() string {
	return ERC721SimpleDefaultAdminDelayChangeCanceledEventName
}

// UnpackDefaultAdminDelayChangeCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeCanceled()
func (eRC721Simple *ERC721Simple) UnpackDefaultAdminDelayChangeCanceledEvent(log *types.Log) (*ERC721SimpleDefaultAdminDelayChangeCanceled, error) {
	event := "DefaultAdminDelayChangeCanceled"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleDefaultAdminDelayChangeCanceled)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleDefaultAdminDelayChangeScheduled represents a DefaultAdminDelayChangeScheduled event raised by the ERC721Simple contract.
type ERC721SimpleDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleDefaultAdminDelayChangeScheduledEventName = "DefaultAdminDelayChangeScheduled"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleDefaultAdminDelayChangeScheduled) ContractEventName() string {
	return ERC721SimpleDefaultAdminDelayChangeScheduledEventName
}

// UnpackDefaultAdminDelayChangeScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminDelayChangeScheduled(uint48 newDelay, uint48 effectSchedule)
func (eRC721Simple *ERC721Simple) UnpackDefaultAdminDelayChangeScheduledEvent(log *types.Log) (*ERC721SimpleDefaultAdminDelayChangeScheduled, error) {
	event := "DefaultAdminDelayChangeScheduled"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleDefaultAdminDelayChangeScheduled)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleDefaultAdminTransferCanceled represents a DefaultAdminTransferCanceled event raised by the ERC721Simple contract.
type ERC721SimpleDefaultAdminTransferCanceled struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleDefaultAdminTransferCanceledEventName = "DefaultAdminTransferCanceled"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleDefaultAdminTransferCanceled) ContractEventName() string {
	return ERC721SimpleDefaultAdminTransferCanceledEventName
}

// UnpackDefaultAdminTransferCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferCanceled()
func (eRC721Simple *ERC721Simple) UnpackDefaultAdminTransferCanceledEvent(log *types.Log) (*ERC721SimpleDefaultAdminTransferCanceled, error) {
	event := "DefaultAdminTransferCanceled"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleDefaultAdminTransferCanceled)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleDefaultAdminTransferScheduled represents a DefaultAdminTransferScheduled event raised by the ERC721Simple contract.
type ERC721SimpleDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleDefaultAdminTransferScheduledEventName = "DefaultAdminTransferScheduled"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleDefaultAdminTransferScheduled) ContractEventName() string {
	return ERC721SimpleDefaultAdminTransferScheduledEventName
}

// UnpackDefaultAdminTransferScheduledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultAdminTransferScheduled(address indexed newAdmin, uint48 acceptSchedule)
func (eRC721Simple *ERC721Simple) UnpackDefaultAdminTransferScheduledEvent(log *types.Log) (*ERC721SimpleDefaultAdminTransferScheduled, error) {
	event := "DefaultAdminTransferScheduled"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleDefaultAdminTransferScheduled)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleERC721BaseSetBaseURI represents a ERC721BaseSetBaseURI event raised by the ERC721Simple contract.
type ERC721SimpleERC721BaseSetBaseURI struct {
	Before  string
	Current string
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleERC721BaseSetBaseURIEventName = "ERC721BaseSetBaseURI"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleERC721BaseSetBaseURI) ContractEventName() string {
	return ERC721SimpleERC721BaseSetBaseURIEventName
}

// UnpackERC721BaseSetBaseURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721BaseSetBaseURI(string before, string current)
func (eRC721Simple *ERC721Simple) UnpackERC721BaseSetBaseURIEvent(log *types.Log) (*ERC721SimpleERC721BaseSetBaseURI, error) {
	event := "ERC721BaseSetBaseURI"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleERC721BaseSetBaseURI)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleForgeAdded represents a ForgeAdded event raised by the ERC721Simple contract.
type ERC721SimpleForgeAdded struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleForgeAddedEventName = "ForgeAdded"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleForgeAdded) ContractEventName() string {
	return ERC721SimpleForgeAddedEventName
}

// UnpackForgeAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeAdded(address indexed forge)
func (eRC721Simple *ERC721Simple) UnpackForgeAddedEvent(log *types.Log) (*ERC721SimpleForgeAdded, error) {
	event := "ForgeAdded"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleForgeAdded)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleForgeRemoved represents a ForgeRemoved event raised by the ERC721Simple contract.
type ERC721SimpleForgeRemoved struct {
	Forge common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleForgeRemovedEventName = "ForgeRemoved"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleForgeRemoved) ContractEventName() string {
	return ERC721SimpleForgeRemovedEventName
}

// UnpackForgeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ForgeRemoved(address indexed forge)
func (eRC721Simple *ERC721Simple) UnpackForgeRemovedEvent(log *types.Log) (*ERC721SimpleForgeRemoved, error) {
	event := "ForgeRemoved"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleForgeRemoved)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleRoleAdminChanged represents a RoleAdminChanged event raised by the ERC721Simple contract.
type ERC721SimpleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleRoleAdminChangedEventName = "RoleAdminChanged"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleRoleAdminChanged) ContractEventName() string {
	return ERC721SimpleRoleAdminChangedEventName
}

// UnpackRoleAdminChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (eRC721Simple *ERC721Simple) UnpackRoleAdminChangedEvent(log *types.Log) (*ERC721SimpleRoleAdminChanged, error) {
	event := "RoleAdminChanged"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleRoleAdminChanged)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleRoleGranted represents a RoleGranted event raised by the ERC721Simple contract.
type ERC721SimpleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleRoleGrantedEventName = "RoleGranted"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleRoleGranted) ContractEventName() string {
	return ERC721SimpleRoleGrantedEventName
}

// UnpackRoleGrantedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC721Simple *ERC721Simple) UnpackRoleGrantedEvent(log *types.Log) (*ERC721SimpleRoleGranted, error) {
	event := "RoleGranted"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleRoleGranted)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleRoleRevoked represents a RoleRevoked event raised by the ERC721Simple contract.
type ERC721SimpleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleRoleRevokedEventName = "RoleRevoked"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleRoleRevoked) ContractEventName() string {
	return ERC721SimpleRoleRevokedEventName
}

// UnpackRoleRevokedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (eRC721Simple *ERC721Simple) UnpackRoleRevokedEvent(log *types.Log) (*ERC721SimpleRoleRevoked, error) {
	event := "RoleRevoked"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleRoleRevoked)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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

// ERC721SimpleTransfer represents a Transfer event raised by the ERC721Simple contract.
type ERC721SimpleTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const ERC721SimpleTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (ERC721SimpleTransfer) ContractEventName() string {
	return ERC721SimpleTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (eRC721Simple *ERC721Simple) UnpackTransferEvent(log *types.Log) (*ERC721SimpleTransfer, error) {
	event := "Transfer"
	if log.Topics[0] != eRC721Simple.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ERC721SimpleTransfer)
	if len(log.Data) > 0 {
		if err := eRC721Simple.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range eRC721Simple.abi.Events[event].Inputs {
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
func (eRC721Simple *ERC721Simple) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["AccessControlBadConfirmation"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackAccessControlBadConfirmationError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["AccessControlEnforcedDefaultAdminDelay"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackAccessControlEnforcedDefaultAdminDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["AccessControlEnforcedDefaultAdminRules"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackAccessControlEnforcedDefaultAdminRulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["AccessControlInvalidDefaultAdmin"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackAccessControlInvalidDefaultAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["AccessControlUnauthorizedAccount"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackAccessControlUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721IncorrectOwner"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721IncorrectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721InsufficientApproval"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721InsufficientApprovalError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721InvalidApprover"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721InvalidApproverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721InvalidOperator"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721InvalidOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721InvalidOwner"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721InvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721InvalidReceiver"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721InvalidReceiverError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721InvalidSender"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721InvalidSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["ERC721NonexistentToken"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackERC721NonexistentTokenError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["TokenBaseNullInput"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackTokenBaseNullInputError(raw[4:])
	}
	if bytes.Equal(raw[:4], eRC721Simple.abi.Errors["TokenBaseOnlyForge"].ID.Bytes()[:4]) {
		return eRC721Simple.UnpackTokenBaseOnlyForgeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ERC721SimpleAccessControlBadConfirmation represents a AccessControlBadConfirmation error raised by the ERC721Simple contract.
type ERC721SimpleAccessControlBadConfirmation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlBadConfirmation()
func ERC721SimpleAccessControlBadConfirmationErrorID() common.Hash {
	return common.HexToHash("0x6697b23232a647058342c0724fe7c415cab25915b54e5dbc03f233173d37b41c")
}

// UnpackAccessControlBadConfirmationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlBadConfirmation()
func (eRC721Simple *ERC721Simple) UnpackAccessControlBadConfirmationError(raw []byte) (*ERC721SimpleAccessControlBadConfirmation, error) {
	out := new(ERC721SimpleAccessControlBadConfirmation)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "AccessControlBadConfirmation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleAccessControlEnforcedDefaultAdminDelay represents a AccessControlEnforcedDefaultAdminDelay error raised by the ERC721Simple contract.
type ERC721SimpleAccessControlEnforcedDefaultAdminDelay struct {
	Schedule *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func ERC721SimpleAccessControlEnforcedDefaultAdminDelayErrorID() common.Hash {
	return common.HexToHash("0x19ca5ebb8fb33f00e502c9392eddab1501674629178bf69b853cf037aaf4bb5d")
}

// UnpackAccessControlEnforcedDefaultAdminDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminDelay(uint48 schedule)
func (eRC721Simple *ERC721Simple) UnpackAccessControlEnforcedDefaultAdminDelayError(raw []byte) (*ERC721SimpleAccessControlEnforcedDefaultAdminDelay, error) {
	out := new(ERC721SimpleAccessControlEnforcedDefaultAdminDelay)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleAccessControlEnforcedDefaultAdminRules represents a AccessControlEnforcedDefaultAdminRules error raised by the ERC721Simple contract.
type ERC721SimpleAccessControlEnforcedDefaultAdminRules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func ERC721SimpleAccessControlEnforcedDefaultAdminRulesErrorID() common.Hash {
	return common.HexToHash("0x3fc3c27ae3db78c81b8f6e685172134623efa268ee8cd8d54be38ad2a74fc13b")
}

// UnpackAccessControlEnforcedDefaultAdminRulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlEnforcedDefaultAdminRules()
func (eRC721Simple *ERC721Simple) UnpackAccessControlEnforcedDefaultAdminRulesError(raw []byte) (*ERC721SimpleAccessControlEnforcedDefaultAdminRules, error) {
	out := new(ERC721SimpleAccessControlEnforcedDefaultAdminRules)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "AccessControlEnforcedDefaultAdminRules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleAccessControlInvalidDefaultAdmin represents a AccessControlInvalidDefaultAdmin error raised by the ERC721Simple contract.
type ERC721SimpleAccessControlInvalidDefaultAdmin struct {
	DefaultAdmin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func ERC721SimpleAccessControlInvalidDefaultAdminErrorID() common.Hash {
	return common.HexToHash("0xc22c8022f2a840d6b6a9f113407715f5bbd4e88c1b0dd9434dc00700ba609ed4")
}

// UnpackAccessControlInvalidDefaultAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlInvalidDefaultAdmin(address defaultAdmin)
func (eRC721Simple *ERC721Simple) UnpackAccessControlInvalidDefaultAdminError(raw []byte) (*ERC721SimpleAccessControlInvalidDefaultAdmin, error) {
	out := new(ERC721SimpleAccessControlInvalidDefaultAdmin)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "AccessControlInvalidDefaultAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleAccessControlUnauthorizedAccount represents a AccessControlUnauthorizedAccount error raised by the ERC721Simple contract.
type ERC721SimpleAccessControlUnauthorizedAccount struct {
	Account    common.Address
	NeededRole [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func ERC721SimpleAccessControlUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0xe2517d3fbfae6f8515ef5ff1ccedc3933ab0cbbda0b492c06eb54ad10ef03b3e")
}

// UnpackAccessControlUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessControlUnauthorizedAccount(address account, bytes32 neededRole)
func (eRC721Simple *ERC721Simple) UnpackAccessControlUnauthorizedAccountError(raw []byte) (*ERC721SimpleAccessControlUnauthorizedAccount, error) {
	out := new(ERC721SimpleAccessControlUnauthorizedAccount)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "AccessControlUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721IncorrectOwner represents a ERC721IncorrectOwner error raised by the ERC721Simple contract.
type ERC721SimpleERC721IncorrectOwner struct {
	Sender  common.Address
	TokenId *big.Int
	Owner   common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func ERC721SimpleERC721IncorrectOwnerErrorID() common.Hash {
	return common.HexToHash("0x64283d7b313c8117c125f736876fa2b4e90ea3831a4716dfdb87d2f540e26289")
}

// UnpackERC721IncorrectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721IncorrectOwner(address sender, uint256 tokenId, address owner)
func (eRC721Simple *ERC721Simple) UnpackERC721IncorrectOwnerError(raw []byte) (*ERC721SimpleERC721IncorrectOwner, error) {
	out := new(ERC721SimpleERC721IncorrectOwner)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721IncorrectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721InsufficientApproval represents a ERC721InsufficientApproval error raised by the ERC721Simple contract.
type ERC721SimpleERC721InsufficientApproval struct {
	Operator common.Address
	TokenId  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func ERC721SimpleERC721InsufficientApprovalErrorID() common.Hash {
	return common.HexToHash("0x177e802f6f313bc89797ecace66d6d29ab4719cbaaacbb87367264048b1eb861")
}

// UnpackERC721InsufficientApprovalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InsufficientApproval(address operator, uint256 tokenId)
func (eRC721Simple *ERC721Simple) UnpackERC721InsufficientApprovalError(raw []byte) (*ERC721SimpleERC721InsufficientApproval, error) {
	out := new(ERC721SimpleERC721InsufficientApproval)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721InsufficientApproval", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721InvalidApprover represents a ERC721InvalidApprover error raised by the ERC721Simple contract.
type ERC721SimpleERC721InvalidApprover struct {
	Approver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidApprover(address approver)
func ERC721SimpleERC721InvalidApproverErrorID() common.Hash {
	return common.HexToHash("0xa9fbf51f86b8e03595d59dc726bb10c329bb24f62589be276d8dd193ca0b69ea")
}

// UnpackERC721InvalidApproverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidApprover(address approver)
func (eRC721Simple *ERC721Simple) UnpackERC721InvalidApproverError(raw []byte) (*ERC721SimpleERC721InvalidApprover, error) {
	out := new(ERC721SimpleERC721InvalidApprover)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721InvalidApprover", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721InvalidOperator represents a ERC721InvalidOperator error raised by the ERC721Simple contract.
type ERC721SimpleERC721InvalidOperator struct {
	Operator common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOperator(address operator)
func ERC721SimpleERC721InvalidOperatorErrorID() common.Hash {
	return common.HexToHash("0x5b08ba185e8f577075361f3a3555a6580a227ce22734dcc979c1aeadf894658b")
}

// UnpackERC721InvalidOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOperator(address operator)
func (eRC721Simple *ERC721Simple) UnpackERC721InvalidOperatorError(raw []byte) (*ERC721SimpleERC721InvalidOperator, error) {
	out := new(ERC721SimpleERC721InvalidOperator)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721InvalidOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721InvalidOwner represents a ERC721InvalidOwner error raised by the ERC721Simple contract.
type ERC721SimpleERC721InvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidOwner(address owner)
func ERC721SimpleERC721InvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x89c62b6479af2e623826dcc39c5133061d35b66d72de92833401dd2fd6567480")
}

// UnpackERC721InvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidOwner(address owner)
func (eRC721Simple *ERC721Simple) UnpackERC721InvalidOwnerError(raw []byte) (*ERC721SimpleERC721InvalidOwner, error) {
	out := new(ERC721SimpleERC721InvalidOwner)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721InvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721InvalidReceiver represents a ERC721InvalidReceiver error raised by the ERC721Simple contract.
type ERC721SimpleERC721InvalidReceiver struct {
	Receiver common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func ERC721SimpleERC721InvalidReceiverErrorID() common.Hash {
	return common.HexToHash("0x64a0ae9278f805eaf991dcd18ca78756d280b7508b764ef1b255c55845c11df9")
}

// UnpackERC721InvalidReceiverError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidReceiver(address receiver)
func (eRC721Simple *ERC721Simple) UnpackERC721InvalidReceiverError(raw []byte) (*ERC721SimpleERC721InvalidReceiver, error) {
	out := new(ERC721SimpleERC721InvalidReceiver)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721InvalidReceiver", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721InvalidSender represents a ERC721InvalidSender error raised by the ERC721Simple contract.
type ERC721SimpleERC721InvalidSender struct {
	Sender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721InvalidSender(address sender)
func ERC721SimpleERC721InvalidSenderErrorID() common.Hash {
	return common.HexToHash("0x73c6ac6e10798e95d99e1f130d923eb40193ecb8d094ec3dce93292564eb3b17")
}

// UnpackERC721InvalidSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721InvalidSender(address sender)
func (eRC721Simple *ERC721Simple) UnpackERC721InvalidSenderError(raw []byte) (*ERC721SimpleERC721InvalidSender, error) {
	out := new(ERC721SimpleERC721InvalidSender)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721InvalidSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleERC721NonexistentToken represents a ERC721NonexistentToken error raised by the ERC721Simple contract.
type ERC721SimpleERC721NonexistentToken struct {
	TokenId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func ERC721SimpleERC721NonexistentTokenErrorID() common.Hash {
	return common.HexToHash("0x7e273289a3a9ef6670f06df7dca227856fc925e956db96980692764a8bc734d7")
}

// UnpackERC721NonexistentTokenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC721NonexistentToken(uint256 tokenId)
func (eRC721Simple *ERC721Simple) UnpackERC721NonexistentTokenError(raw []byte) (*ERC721SimpleERC721NonexistentToken, error) {
	out := new(ERC721SimpleERC721NonexistentToken)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "ERC721NonexistentToken", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the ERC721Simple contract.
type ERC721SimpleSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func ERC721SimpleSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (eRC721Simple *ERC721Simple) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*ERC721SimpleSafeCastOverflowedUintDowncast, error) {
	out := new(ERC721SimpleSafeCastOverflowedUintDowncast)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleTokenBaseNullInput represents a TokenBase__NullInput error raised by the ERC721Simple contract.
type ERC721SimpleTokenBaseNullInput struct {
	Field [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func ERC721SimpleTokenBaseNullInputErrorID() common.Hash {
	return common.HexToHash("0xe1dea2ef34b70efdd13a051f03d858db14b7868b90c2bf12109dd86bd06a6021")
}

// UnpackTokenBaseNullInputError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__NullInput(bytes32 field)
func (eRC721Simple *ERC721Simple) UnpackTokenBaseNullInputError(raw []byte) (*ERC721SimpleTokenBaseNullInput, error) {
	out := new(ERC721SimpleTokenBaseNullInput)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "TokenBaseNullInput", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ERC721SimpleTokenBaseOnlyForge represents a TokenBase__OnlyForge error raised by the ERC721Simple contract.
type ERC721SimpleTokenBaseOnlyForge struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func ERC721SimpleTokenBaseOnlyForgeErrorID() common.Hash {
	return common.HexToHash("0x25a9dbc17c85688aaa4e1275eaf8f2a0822a5bf72b075c1ddfff945b8481dc8b")
}

// UnpackTokenBaseOnlyForgeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenBase__OnlyForge(address caller)
func (eRC721Simple *ERC721Simple) UnpackTokenBaseOnlyForgeError(raw []byte) (*ERC721SimpleTokenBaseOnlyForge, error) {
	out := new(ERC721SimpleTokenBaseOnlyForge)
	if err := eRC721Simple.abi.UnpackIntoInterface(out, "TokenBaseOnlyForge", raw); err != nil {
		return nil, err
	}
	return out, nil
}
