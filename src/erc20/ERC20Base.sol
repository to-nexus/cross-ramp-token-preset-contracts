// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20Upgradeable} from "@openzeppelin-contracts-upgradeable-5.4.0/token/ERC20/ERC20Upgradeable.sol";
import {ERC20PermitUpgradeable} from
    "@openzeppelin-contracts-upgradeable-5.4.0/token/ERC20/extensions/ERC20PermitUpgradeable.sol";
import {IERC20Forge} from "../interfaces/IERC20Forge.sol";
import {TokenBase} from "../TokenBase.sol";

abstract contract ERC20Base is TokenBase, IERC20Forge, ERC20Upgradeable, ERC20PermitUpgradeable {
    // keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.decimals")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20DecimalsStorageLocation =
        0x487e9ce4507767927f0d69bc1e3e8a725bb861e44fe4728e56521e66cd47da00;

    constructor() {
        _disableInitializers();
    }

    function initialize(
        address owner,
        address manager,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory
    ) external virtual override initializer {
        __ERC20Base_init(owner, manager, _name, _symbol, _decimals, initialSupply, initialRecipient);
    }

    function __ERC20Base_init(
        address owner,
        address manager,
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        uint256 initialSupply,
        address initialRecipient
    ) internal onlyInitializing {
        if (bytes(_name).length == 0) revert TokenBase__NullInput("name");
        if (bytes(_symbol).length == 0) revert TokenBase__NullInput("symbol");

        __TokenBase_init(owner, manager);

        __ERC20_init(_name, _symbol);
        __ERC20Permit_init(_name);
        __ERC20Base_init_unchained(_decimals, initialSupply, initialRecipient);
    }

    function __ERC20Base_init_unchained(uint8 _decimals, uint256 initialSupply, address initialRecipient)
        private
        onlyInitializing
    {
        assembly {
            sstore(ERC20DecimalsStorageLocation, _decimals)
        }
        if (initialSupply != 0) {
            if (initialRecipient == address(0)) revert TokenBase__NullInput("initialRecipient");
            ERC20Upgradeable._update(address(0), initialRecipient, initialSupply);
        }
    }

    function mint(address to, uint256 amount) public virtual override onlyForge {
        _mint(to, amount);
    }

    function burnFrom(address from, uint256 amount) public virtual override {
        _spendAllowance(from, _msgSender(), amount);
        _burn(from, amount);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IERC20Forge).interfaceId || super.supportsInterface(interfaceId);
    }

    ///////////////////////////////
    // override ERC20Upgradeable //
    ///////////////////////////////

    function name() public view virtual override returns (string memory) {
        return ERC20Upgradeable.name();
    }

    function symbol() public view virtual override returns (string memory) {
        return ERC20Upgradeable.symbol();
    }

    function decimals() public view virtual override returns (uint8) {
        uint8 _decimals;
        assembly {
            _decimals := sload(ERC20DecimalsStorageLocation)
        }
        return _decimals;
    }

    function totalSupply() public view virtual override returns (uint256) {
        return ERC20Upgradeable.totalSupply();
    }

    function balanceOf(address account) public view virtual override returns (uint256) {
        return ERC20Upgradeable.balanceOf(account);
    }

    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return ERC20Upgradeable.allowance(owner, spender);
    }

    function transfer(address to, uint256 value) public virtual override returns (bool) {
        return ERC20Upgradeable.transfer(to, value);
    }

    function approve(address spender, uint256 value) public virtual override returns (bool) {
        return ERC20Upgradeable.approve(spender, value);
    }

    function transferFrom(address from, address to, uint256 value) public virtual override returns (bool) {
        return ERC20Upgradeable.transferFrom(from, to, value);
    }

    function _update(address from, address to, uint256 value) internal virtual override {
        ERC20Upgradeable._update(from, to, value);
    }

    function _approve(address owner, address spender, uint256 value, bool emitEvent) internal virtual override {
        ERC20Upgradeable._approve(owner, spender, value, emitEvent);
    }

    function _spendAllowance(address owner, address spender, uint256 value) internal virtual override {
        ERC20Upgradeable._spendAllowance(owner, spender, value);
    }

    /////////////////////////////////////
    // override ERC20PermitUpgradeable //
    /////////////////////////////////////

    function nonces(address owner) public view virtual override returns (uint256) {
        return ERC20PermitUpgradeable.nonces(owner);
    }
}
