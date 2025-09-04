// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {
    AccessControlUpgradeable,
    IAccessControl
} from "@openzeppelin-contracts-upgradeable-5.4.0/access/AccessControlUpgradeable.sol";
import {AccessControlDefaultAdminRulesUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.4.0/access/extensions/AccessControlDefaultAdminRulesUpgradeable.sol";
import {AccessControlEnumerableUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.4.0/access/extensions/AccessControlEnumerableUpgradeable.sol";

abstract contract BaseAccessControlUpgradeable is
    IAccessControl,
    AccessControlUpgradeable,
    AccessControlDefaultAdminRulesUpgradeable,
    AccessControlEnumerableUpgradeable
{
    function __BaseAccessControl_init(address _owner) internal onlyInitializing {
        __BaseAccessControl_init_unchained(_owner);
    }

    function __BaseAccessControl_init_unchained(address _owner) internal onlyInitializing {
        __AccessControl_init();
        __AccessControlDefaultAdminRules_init(86400, _owner);
        __AccessControlEnumerable_init();
    }

    //////////////
    // override //
    //////////////

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable, AccessControlEnumerableUpgradeable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function grantRole(bytes32 role, address account)
        public
        virtual
        override(IAccessControl, AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable)
    {
        super.grantRole(role, account);
    }

    function revokeRole(bytes32 role, address account)
        public
        virtual
        override(IAccessControl, AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable)
    {
        super.revokeRole(role, account);
    }

    function renounceRole(bytes32 role, address account)
        public
        virtual
        override(IAccessControl, AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable)
    {
        super.renounceRole(role, account);
    }

    function _grantRole(bytes32 role, address account)
        internal
        virtual
        override(AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable, AccessControlEnumerableUpgradeable)
        returns (bool)
    {
        return super._grantRole(role, account);
    }

    function _revokeRole(bytes32 role, address account)
        internal
        virtual
        override(AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable, AccessControlEnumerableUpgradeable)
        returns (bool)
    {
        return super._revokeRole(role, account);
    }

    function _setRoleAdmin(bytes32 role, bytes32 adminRole)
        internal
        virtual
        override(AccessControlUpgradeable, AccessControlDefaultAdminRulesUpgradeable)
    {
        super._setRoleAdmin(role, adminRole);
    }
}
