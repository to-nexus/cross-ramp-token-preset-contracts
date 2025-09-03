// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20Capable is ERC20Base {
    uint256 private immutable _cap;
    /**
     * @dev Total supply cap has been exceeded.
     */

    error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap);

    constructor(uint256 cap_) {
        if (cap_ == 0) {
            revert TokenBase__NullInput("cap");
        }
        _cap = cap_;
    }

    function remainingSupply() external view returns (uint256) {
        uint256 currentSupply = totalSupply();
        uint256 maxSupply = cap();
        unchecked {
            return maxSupply > currentSupply ? maxSupply - currentSupply : 0;
        }
    }

    function cap() public view returns (uint256) {
        return _cap;
    }

    function _update(address from, address to, uint256 value) internal virtual override {
        super._update(from, to, value);

        if (from == address(0)) {
            uint256 maxSupply = cap();
            uint256 supply = totalSupply();
            if (supply > maxSupply) {
                revert ERC20Capable__ERC20ExceededCap(supply, maxSupply);
            }
        }
    }
}
