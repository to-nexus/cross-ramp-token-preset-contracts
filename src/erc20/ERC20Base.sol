// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC20} from "@openzeppelin-contracts-5.4.0/token/ERC20/ERC20.sol";
import {ERC20Permit} from "@openzeppelin-contracts-5.4.0/token/ERC20/extensions/ERC20Permit.sol";
import {IERC20Forge} from "../interfaces/IERC20Forge.sol";
import {TokenBase} from "../TokenBase.sol";

abstract contract ERC20Base is TokenBase, IERC20Forge, ERC20, ERC20Permit {
    uint8 private immutable _decimals;

    constructor(address owner, address[] memory forges, string memory name_, string memory symbol_, uint8 decimals_)
        TokenBase(owner, forges)
        ERC20(name_, symbol_)
        ERC20Permit(name_)
    {
        if (bytes(name_).length == 0) revert TokenBase__NullInput("name");
        if (bytes(symbol_).length == 0) revert TokenBase__NullInput("symbol");
        _decimals = decimals_;
    }

    function mint(address to, uint256 amount) public virtual override onlyForge {
        _mint(to, amount);
    }

    function burnFrom(address from, uint256 amount) public virtual override {
        _spendAllowance(from, _msgSender(), amount);
        _burn(from, amount);
    }

    ////////////////////
    // override ERC20 //
    ////////////////////

    function name() public view virtual override returns (string memory) {
        return ERC20.name();
    }

    function symbol() public view virtual override returns (string memory) {
        return ERC20.symbol();
    }

    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    function totalSupply() public view virtual override returns (uint256) {
        return ERC20.totalSupply();
    }

    function balanceOf(address account) public view virtual override returns (uint256) {
        return ERC20.balanceOf(account);
    }

    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return ERC20.allowance(owner, spender);
    }

    function transfer(address to, uint256 value) public virtual override returns (bool) {
        return ERC20.transfer(to, value);
    }

    function approve(address spender, uint256 value) public virtual override returns (bool) {
        return ERC20.approve(spender, value);
    }

    function transferFrom(address from, address to, uint256 value) public virtual override returns (bool) {
        return ERC20.transferFrom(from, to, value);
    }

    function _update(address from, address to, uint256 value) internal virtual override {
        ERC20._update(from, to, value);
    }

    function _approve(address owner, address spender, uint256 value, bool emitEvent) internal virtual override {
        ERC20._approve(owner, spender, value, emitEvent);
    }

    function _spendAllowance(address owner, address spender, uint256 value) internal virtual override {
        ERC20._spendAllowance(owner, spender, value);
    }

    //////////////////////////
    // override ERC20Permit //
    //////////////////////////

    function nonces(address owner) public view virtual override returns (uint256) {
        return ERC20Permit.nonces(owner);
    }

    ////////////////////////
    // override TokenBase //
    ////////////////////////

    function supportsInterface(bytes4 interfaceId) public view virtual override(TokenBase) returns (bool) {
        return interfaceId == type(IPreset).interfaceId || interfaceId == type(IERC20Forge).interfaceId
            || interfaceId == type(IERC20).interfaceId || interfaceId == type(IERC20Metadata).interfaceId
            || interfaceId == type(IERC20Errors).interfaceId || interfaceId == type(IERC20Permit).interfaceId
            || interfaceId == type(IERC5267).interfaceId || TokenBase.supportsInterface(interfaceId);
    }
}

import {IPreset} from "../interfaces/IPreset.sol";
import {IERC20} from "@openzeppelin-contracts-5.4.0/token/ERC20/IERC20.sol";
import {IERC20Permit} from "@openzeppelin-contracts-5.4.0/token/erc20/extensions/IERC20Permit.sol";
import {IERC20Metadata} from "@openzeppelin-contracts-5.4.0/interfaces/IERC20Metadata.sol";
import {IERC20Errors} from "@openzeppelin-contracts-5.4.0/interfaces/draft-IERC6093.sol";
import {IERC5267} from "@openzeppelin-contracts-5.4.0/interfaces/IERC5267.sol";

abstract contract ERC20BasePreset is IPreset {
    function code() external pure virtual override returns (bytes memory);

    function deployCode(bytes memory initialData) external pure virtual override returns (bytes memory);

    function supportsInterface(bytes4 interfaceId) external view virtual override returns (bool) {
        return interfaceId == type(IPreset).interfaceId || interfaceId == type(IERC20Forge).interfaceId
            || interfaceId == type(IERC20).interfaceId || interfaceId == type(IERC20Metadata).interfaceId
            || interfaceId == type(IERC20Errors).interfaceId || interfaceId == type(IERC20Permit).interfaceId
            || interfaceId == type(IERC5267).interfaceId || interfaceId == 0x01ffc9a7; // ERC165
    }
}
