// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";
import {ERC20Capable} from "../extensions/ERC20Capable.sol";
import {ERC20InitialSupply} from "../extensions/ERC20InitialSupply.sol";
import {ERC20PeriodMintLimit} from "../extensions/ERC20PeriodMintLimit.sol";

contract ERC20CappedInitialSupplyMintLimited is ERC20Base, ERC20Capable, ERC20InitialSupply, ERC20PeriodMintLimit {
    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 cap,
        uint256 initialSupply,
        address initialRecipient,
        uint256 duration,
        int256 offsetSeconds,
        uint256 limit
    )
        ERC20Base(owner, forges, name, symbol, decimals)
        ERC20Capable(cap)
        ERC20InitialSupply(initialSupply, initialRecipient)
        ERC20PeriodMintLimit(duration, offsetSeconds, limit)
    {}

    function mint(address to, uint256 amount) public override(ERC20Base, ERC20PeriodMintLimit) {
        ERC20PeriodMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}

import {ERC20BasePreset} from "../ERC20Base.sol";

contract ERC20CappedInitialSupplyMintLimitedPreset is ERC20BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC20CappedInitialSupplyMintLimited).creationCode;
    }

    function deployCode(bytes memory initialData) external pure override returns (bytes memory) {
        (
            address owner,
            address[] memory forges,
            string memory name,
            string memory symbol,
            uint8 decimals,
            bytes memory data
        ) = abi.decode(initialData, (address, address[], string, string, uint8, bytes));

        (
            uint256 cap,
            uint256 initialSupply,
            address initialRecipient,
            uint256 duration,
            int256 offsetSeconds,
            uint256 limit
        ) = abi.decode(data, (uint256, uint256, address, uint256, int256, uint256));

        return abi.encodePacked(
            code(),
            abi.encode(
                owner,
                forges,
                name,
                symbol,
                decimals,
                cap,
                initialSupply,
                initialRecipient,
                duration,
                offsetSeconds,
                limit
            )
        );
    }
}
