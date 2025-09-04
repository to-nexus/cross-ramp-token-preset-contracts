// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";
import {ERC20Capable} from "../extensions/ERC20Capable.sol";
import {ERC20InitialSupply} from "../extensions/ERC20InitialSupply.sol";
import {ERC20PeriodsMintLimit} from "../extensions/ERC20PeriodsMintLimit.sol";

contract ERC20CappedInitialSupplyMultiMintLimited is
    ERC20Base,
    ERC20Capable,
    ERC20InitialSupply,
    ERC20PeriodsMintLimit
{
    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 cap,
        uint256 initialSupply,
        address initialRecipient,
        uint256[] memory durations,
        int256[] memory offsetSeconds,
        uint256[] memory limits
    )
        ERC20Base(owner, forges, name, symbol, decimals)
        ERC20Capable(cap)
        ERC20InitialSupply(1, address(1))
        ERC20PeriodsMintLimit(durations, offsetSeconds, limits)
    {}

    function mint(address to, uint256 amount) public override(ERC20Base, ERC20PeriodsMintLimit) {
        ERC20PeriodsMintLimit.mint(to, amount);
    }

    function _update(address from, address to, uint256 value) internal override(ERC20Base, ERC20Capable) {
        ERC20Capable._update(from, to, value);
    }
}

import {ERC20BasePreset} from "../ERC20Base.sol";

contract ERC20CappedInitialSupplyMultiMintLimitedPreset is ERC20BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC20CappedInitialSupplyMultiMintLimited).creationCode;
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
            uint256[] memory durations,
            int256[] memory offsetSeconds,
            uint256[] memory limits
        ) = abi.decode(data, (uint256, uint256, address, uint256[], int256[], uint256[]));

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
                durations,
                offsetSeconds,
                limits
            )
        );
    }
}
