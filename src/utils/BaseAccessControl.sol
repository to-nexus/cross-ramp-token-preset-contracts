// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {AccessControl, IAccessControl} from "@openzeppelin-contracts-5.4.0/access/AccessControl.sol";
import {AccessControlDefaultAdminRules} from
    "@openzeppelin-contracts-5.4.0/access/extensions/AccessControlDefaultAdminRules.sol";
import {AccessControlEnumerable} from "@openzeppelin-contracts-5.4.0/access/extensions/AccessControlEnumerable.sol";

abstract contract BaseAccessControl is
    IAccessControl,
    AccessControl,
    AccessControlDefaultAdminRules,
    AccessControlEnumerable
{
    constructor(address _owner) AccessControlDefaultAdminRules(86400, _owner) {}

    //////////////
    // override //
    //////////////

    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControl, AccessControlDefaultAdminRules, AccessControlEnumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function grantRole(bytes32 role, address account)
        public
        virtual
        override(IAccessControl, AccessControl, AccessControlDefaultAdminRules)
    {
        super.grantRole(role, account);
    }

    function revokeRole(bytes32 role, address account)
        public
        virtual
        override(IAccessControl, AccessControl, AccessControlDefaultAdminRules)
    {
        super.revokeRole(role, account);
    }

    function renounceRole(bytes32 role, address account)
        public
        virtual
        override(IAccessControl, AccessControl, AccessControlDefaultAdminRules)
    {
        super.renounceRole(role, account);
    }

    function _grantRole(bytes32 role, address account)
        internal
        virtual
        override(AccessControl, AccessControlDefaultAdminRules, AccessControlEnumerable)
        returns (bool)
    {
        return super._grantRole(role, account);
    }

    function _revokeRole(bytes32 role, address account)
        internal
        virtual
        override(AccessControl, AccessControlDefaultAdminRules, AccessControlEnumerable)
        returns (bool)
    {
        return super._revokeRole(role, account);
    }

    function _setRoleAdmin(bytes32 role, bytes32 adminRole)
        internal
        virtual
        override(AccessControl, AccessControlDefaultAdminRules)
    {
        super._setRoleAdmin(role, adminRole);
    }
}
