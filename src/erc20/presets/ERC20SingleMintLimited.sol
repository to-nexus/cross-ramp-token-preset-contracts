// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";
import {ERC20Capable} from "../extensions/ERC20Capable.sol";
import {ERC20PeriodMintLimit} from "../extensions/ERC20PeriodMintLimit.sol";

contract ERC20SingleMintLimited is ERC20Base, ERC20Capable, ERC20PeriodMintLimit {
    function initialize(
        address owner,
        address manager,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient,
        bytes memory data
    ) external override initializer {
        {
            uint256 cap_;
            (cap_, data) = abi.decode(data, (uint256, bytes));
            __ERC20Capable_init(cap_);
        }
        {
            uint256 duration;
            int256 offsetSeconds;
            uint256 limit;
            (duration, offsetSeconds, limit) = abi.decode(data, (uint256, int256, uint256));
            __ERC20PeriodMintLimit_init(duration, offsetSeconds, limit);
        }
        __ERC20Base_init(owner, manager, name, symbol, decimals, initialSupply, initialRecipient);
    }

    function mint(address to, uint256 amount) public override(ERC20Base, ERC20PeriodMintLimit) {
        ERC20PeriodMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}
