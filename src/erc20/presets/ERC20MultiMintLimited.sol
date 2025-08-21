// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";
import {ERC20Capable} from "../extensions/ERC20Capable.sol";
import {ERC20PeriodsMintLimit} from "../extensions/ERC20PeriodsMintLimit.sol";

contract ERC20MultiMintLimited is ERC20Base, ERC20Capable, ERC20PeriodsMintLimit {
    constructor() {
        _disableInitializers();
    }

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
            uint256[] memory durations;
            int256[] memory offsetSeconds;
            uint256[] memory limits;
            (durations, offsetSeconds, limits) = abi.decode(data, (uint256[], int256[], uint256[]));
            __ERC20PeriodsMintLimit_init(durations, offsetSeconds, limits);
        }
        __ERC20Base_init(owner, manager, name, symbol, decimals, initialSupply, initialRecipient);
    }

    function mint(address to, uint256 amount) public override(ERC20Base, ERC20PeriodsMintLimit) {
        ERC20PeriodsMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}
