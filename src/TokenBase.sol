// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {EnumerableSet} from "@openzeppelin-contracts-5.4.0/utils/structs/EnumerableSet.sol";

import {BaseAccessControl} from "./utils/BaseAccessControl.sol";

abstract contract TokenBase is BaseAccessControl {
    using EnumerableSet for EnumerableSet.AddressSet;

    error TokenBase__NullInput(bytes32 field);
    error TokenBase__OnlyForge(address caller);

    event ForgeAdded(address indexed forge);
    event ForgeRemoved(address indexed forge);

    bytes32 public constant MANAGER_ROLE = keccak256("MANAGER");
    EnumerableSet.AddressSet private _forges;

    constructor(address owner, address[] memory forges_) BaseAccessControl(owner) {
        _grantRole(MANAGER_ROLE, owner);
        unchecked {
            uint256 length = forges_.length;
            for (uint256 i = 0; i < length; ++i) {
                _addForge(_forges, forges_[i]);
            }
        }
    }

    modifier onlyForge() {
        if (!isForge(_msgSender())) revert TokenBase__OnlyForge(_msgSender());
        _;
    }

    function forges() external view returns (address[] memory) {
        return _forges.values();
    }

    function forgeCount() external view returns (uint256) {
        return _forges.length();
    }

    function forgeByIndex(uint256 index) external view returns (address) {
        return _forges.at(index);
    }

    function isForge(address forge) public view returns (bool) {
        return _forges.contains(forge);
    }

    function setForges(address[] calldata forges_, bool add) external onlyRole(MANAGER_ROLE) {
        function(EnumerableSet.AddressSet storage, address) fn = add ? _addForge : _removeForge;

        unchecked {
            uint256 length = forges_.length;
            for (uint256 i = 0; i < length; ++i) {
                fn(_forges, forges_[i]);
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
        return BaseAccessControl.supportsInterface(interfaceId);
    }
}
