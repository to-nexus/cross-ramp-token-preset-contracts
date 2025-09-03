// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.4.0/proxy/utils/UUPSUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.4.0/utils/structs/EnumerableSet.sol";
import {Create2} from "@openzeppelin-contracts-5.4.0/utils/Create2.sol";
import {IERC165} from "@openzeppelin-contracts-5.4.0/interfaces/IERC165.sol";

import {IERC20Forge} from "./interfaces/IERC20Forge.sol";
import {IERC721Forge} from "./interfaces/IERC721Forge.sol";
import {IERC1155Forge} from "./interfaces/IERC1155Forge.sol";
import {IPreset} from "./interfaces/IPreset.sol";

import {TokenType, ITokenFactory} from "./interfaces/ITokenFactory.sol";
import {BaseAccessControlUpgradeable} from "./utils/BaseAccessControlUpgradeable.sol";

contract TokenFactoryImpl is ITokenFactory, BaseAccessControlUpgradeable, UUPSUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    error TokenFactory__ZeroAddress(bytes32 field);
    error TokenFactory__InvalidLogic(TokenType, address);
    error TokenFactory__DeployFailed(TokenType, address);

    bytes32 public constant DEPLOYER_ROLE = keccak256("DEPLOYER");

    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.TokenFactory.erc20Impls")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TOKEN_FACTORY_ERC20_STORAGE_LOCATION =
        0xd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f00;
    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.TokenFactory.erc721Impls")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TOKEN_FACTORY_ERC721_STORAGE_LOCATION =
        0xd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f00;
    // keccak256(abi.encode(uint256(keccak256("cross.storage.forge.TokenFactory.erc1155Impls")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TOKEN_FACTORY_ERC1155_STORAGE_LOCATION =
        0xd9ce7ba49891d0b3422c928b675902668a780038fd864217ee35722e549a9f00;

    function _getTokenFactoryERC20Storage() private pure returns (EnumerableSet.AddressSet storage $) {
        assembly {
            $.slot := TOKEN_FACTORY_ERC20_STORAGE_LOCATION
        }
    }

    function _getTokenFactoryERC721Storage() private pure returns (EnumerableSet.AddressSet storage $) {
        assembly {
            $.slot := TOKEN_FACTORY_ERC721_STORAGE_LOCATION
        }
    }

    function _getTokenFactoryERC1155Storage() private pure returns (EnumerableSet.AddressSet storage $) {
        assembly {
            $.slot := TOKEN_FACTORY_ERC1155_STORAGE_LOCATION
        }
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    function initialize(
        address owner,
        address[] memory erc20Impls,
        address[] memory erc721Impls,
        address[] memory erc1155Impls
    ) external initializer {
        __BaseAccessControl_init(owner);
        __UUPSUpgradeable_init();

        _grantRole(DEPLOYER_ROLE, owner);
        _setImpls(TokenType.ERC20, type(IERC20Forge).interfaceId, _getTokenFactoryERC20Storage(), erc20Impls, true);
        _setImpls(TokenType.ERC721, type(IERC721Forge).interfaceId, _getTokenFactoryERC721Storage(), erc721Impls, true);
        _setImpls(
            TokenType.ERC1155, type(IERC1155Forge).interfaceId, _getTokenFactoryERC1155Storage(), erc1155Impls, true
        );
    }

    function getPresets(TokenType tokenType) external view returns (address[] memory) {
        if (tokenType == TokenType.ERC20) {
            return _getTokenFactoryERC20Storage().values();
        } else if (tokenType == TokenType.ERC721) {
            return _getTokenFactoryERC721Storage().values();
        } else {
            return _getTokenFactoryERC1155Storage().values();
        }
    }

    function deployERC20(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes memory data,
        address preset
    ) external override onlyRole(DEPLOYER_ROLE) returns (address token) {
        if (!_getTokenFactoryERC20Storage().contains(preset)) {
            revert TokenFactory__InvalidLogic(TokenType.ERC20, preset);
        }

        bytes memory initialData = abi.encode(owner, forges, name, symbol, decimals, data);
        token = Create2.deploy(0, keccak256(abi.encodePacked(name, symbol)), IPreset(preset).deployCode(initialData));
        if (token == address(0)) revert TokenFactory__DeployFailed(TokenType.ERC20, preset);

        emit TokenDeployed(owner, TokenType.ERC20, token, preset);
    }

    function deployERC721(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory data,
        address preset
    ) external onlyRole(DEPLOYER_ROLE) returns (address token) {
        if (!_getTokenFactoryERC721Storage().contains(preset)) {
            revert TokenFactory__InvalidLogic(TokenType.ERC721, preset);
        }

        bytes memory initialData = abi.encode(owner, forges, name, symbol, baseTokenURI, data);
        token = Create2.deploy(0, keccak256(abi.encodePacked(name, symbol)), IPreset(preset).deployCode(initialData));
        if (token == address(0)) revert TokenFactory__DeployFailed(TokenType.ERC721, preset);

        emit TokenDeployed(owner, TokenType.ERC721, token, preset);
    }

    function deployERC1155(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        string memory uri,
        bytes memory data,
        address preset
    ) external onlyRole(DEPLOYER_ROLE) returns (address token) {
        if (!_getTokenFactoryERC1155Storage().contains(preset)) {
            revert TokenFactory__InvalidLogic(TokenType.ERC1155, preset);
        }

        bytes memory initialData = abi.encode(owner, forges, name, symbol, uri, data);
        token = Create2.deploy(0, keccak256(abi.encodePacked(name, symbol)), IPreset(preset).deployCode(initialData));
        if (token == address(0)) revert TokenFactory__DeployFailed(TokenType.ERC1155, preset);

        emit TokenDeployed(owner, TokenType.ERC1155, token, preset);
    }

    function setPresets(TokenType tokenType, address[] calldata presets, bool add)
        external
        onlyRole(DEFAULT_ADMIN_ROLE)
    {
        EnumerableSet.AddressSet storage _impls;
        bytes4 interfaceId;
        if (tokenType == TokenType.ERC20) {
            (_impls, interfaceId) = (_getTokenFactoryERC20Storage(), type(IERC20Forge).interfaceId);
        } else if (tokenType == TokenType.ERC721) {
            (_impls, interfaceId) = (_getTokenFactoryERC721Storage(), type(IERC721Forge).interfaceId);
        } else {
            (_impls, interfaceId) = (_getTokenFactoryERC1155Storage(), type(IERC1155Forge).interfaceId);
        }
        _setImpls(tokenType, interfaceId, _impls, presets, add);
    }

    function _setImpls(
        TokenType tokeType,
        bytes4 interfaceId,
        EnumerableSet.AddressSet storage _impls,
        address[] memory impls,
        bool add
    ) private {
        uint256 length = impls.length;
        unchecked {
            if (add) {
                for (uint256 i = 0; i < length; ++i) {
                    address impl = impls[i];
                    if (impl == address(0)) {
                        revert TokenFactory__ZeroAddress("impls");
                    }
                    if (!IERC165(impl).supportsInterface(interfaceId)) {
                        revert TokenFactory__InvalidLogic(tokeType, impl);
                    }
                    if (_impls.add(impl)) {
                        emit PresetLogicSet(tokeType, impl);
                    }
                }
            } else {
                for (uint256 i = 0; i < length; ++i) {
                    address impl = impls[i];
                    if (_impls.remove(impl)) {
                        emit PresetLogicRemoved(tokeType, impl);
                    }
                }
            }
        }
    }

    function _authorizeUpgrade(address) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
