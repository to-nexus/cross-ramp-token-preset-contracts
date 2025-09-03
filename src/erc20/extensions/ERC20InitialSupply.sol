// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20InitialSupply is ERC20Base {
    constructor(uint256 initialSupply, address initialRecipient) {
        if (initialSupply == 0) revert TokenBase__NullInput("initialSupply");
        if (initialRecipient == address(0)) revert TokenBase__NullInput("initialRecipient");
        _mint(initialRecipient, initialSupply);
    }
}
