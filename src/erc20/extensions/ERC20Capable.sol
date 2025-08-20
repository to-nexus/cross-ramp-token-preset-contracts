// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20Capable is ERC20Base {
    /**
     * @dev Total supply cap has been exceeded.
     */
    error ERC20Capable__ERC20ExceededCap(uint256 increasedSupply, uint256 cap);

    // keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.ERC20Capable")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20CapableStorageLocation =
        0x6124d57edb60d866286eaee21ef1ce5eb21f42d9a55fe900ab9d9e1f5aa87200;

    function __ERC20Capable_init(uint256 cap_) internal onlyInitializing {
        if (cap_ == 0) {
            revert TokenBase__NullInput("cap");
        }
        assembly {
            sstore(ERC20CapableStorageLocation, cap_)
        }
    }

    function remainingSupply() external view returns (uint256) {
        uint256 currentSupply = totalSupply();
        uint256 maxSupply = cap();
        unchecked {
            return maxSupply > currentSupply ? maxSupply - currentSupply : 0;
        }
    }

    function cap() public view returns (uint256) {
        uint256 cap_;
        assembly {
            cap_ := sload(ERC20CapableStorageLocation)
        }
        return cap_;
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
