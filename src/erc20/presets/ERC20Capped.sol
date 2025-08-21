// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base, ERC20Capable} from "../extensions/ERC20Capable.sol";

contract ERC20Capped is ERC20Capable {
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data
    ) external override initializer {
        {
            uint256 cap_ = abi.decode(data, (uint256));
            __ERC20Capable_init(cap_);
        }
        __ERC20Base_init(_owner, manager, name, symbol, decimals, initialSupply, initialRecipient);
    }
}
