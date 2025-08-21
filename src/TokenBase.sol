// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {UUPSUpgradeable} from "@openzeppelin-contracts-upgradeable-5.4.0/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-contracts-upgradeable-5.4.0/access/OwnableUpgradeable.sol";
import {AccessControlUpgradeable} from "@openzeppelin-contracts-upgradeable-5.4.0/access/AccessControlUpgradeable.sol";
import {EnumerableSet} from "@openzeppelin-contracts-5.4.0/utils/structs/EnumerableSet.sol";

abstract contract TokenBase is UUPSUpgradeable, AccessControlUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER");

    error TokenBase__NullInput(bytes32 field);
    error TokenBase__OnlyForge(address caller);

    event ForgeAdded(address indexed forge);
    event ForgeRemoved(address indexed forge);

    // keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.TokenBase")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant TokenBaseStorageLocation =
        0x6d8cdc32fd80446b76af2a0b4f9cf165aa632ecb729677fa8eb4cec7fe4fa300;

    function _getForgesStorage() private pure returns (EnumerableSet.AddressSet storage $) {
        assembly {
            $.slot := TokenBaseStorageLocation
        }
    }

    constructor() {
        _disableInitializers();
    }

    function __TokenBase_init(address _owner, address _manager) internal onlyInitializing {
        __AccessControl_init();
        __UUPSUpgradeable_init();

        if (_owner == address(0)) revert TokenBase__NullInput("owner");
        _grantRole(DEFAULT_ADMIN_ROLE, _owner);
        _grantRole(MANAGER_ROLE, _owner);

        if (_manager != address(0) && _manager != _owner) {
            _grantRole(MANAGER_ROLE, _manager);
        }

        _setRoleAdmin(MANAGER_ROLE, DEFAULT_ADMIN_ROLE);
    }

    modifier onlyForge() {
        if (!isForge(_msgSender())) revert TokenBase__OnlyForge(_msgSender());
        _;
    }

    function forges() external view returns (address[] memory) {
        return _getForgesStorage().values();
    }

    function forgeCount() external view returns (uint256) {
        return _getForgesStorage().length();
    }

    function forgeByIndex(uint256 index) external view returns (address) {
        return _getForgesStorage().at(index);
    }

    function isForge(address forge) public view returns (bool) {
        return _getForgesStorage().contains(forge);
    }

    function setForges(address[] calldata _forges, bool add) external onlyRole(MANAGER_ROLE) {
        function (EnumerableSet.AddressSet storage, address) fn = add ? _addForge : _removeForge;

        EnumerableSet.AddressSet storage $ = _getForgesStorage();
        unchecked {
            uint256 length = _forges.length;
            for (uint256 i = 0; i < length; ++i) {
                fn($, _forges[i]);
            }
        }
    }

    function _addForge(EnumerableSet.AddressSet storage $, address forge) private {
        if (forge == address(0)) revert TokenBase__NullInput("forge");
        if ($.add(forge)) {
            emit ForgeAdded(forge);
        }
    }

    function _removeForge(EnumerableSet.AddressSet storage $, address forge) private {
        if ($.remove(forge)) {
            emit ForgeRemoved(forge);
        }
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return AccessControlUpgradeable.supportsInterface(interfaceId);
    }

    function _authorizeUpgrade(address) internal override onlyRole(DEFAULT_ADMIN_ROLE) {}
}
