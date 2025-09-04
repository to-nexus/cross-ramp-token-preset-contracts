// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base, ERC20PeriodsMintLimit} from "../extensions/ERC20PeriodsMintLimit.sol";

contract ERC20MultiMintLimited is ERC20PeriodsMintLimit {
    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256[] memory durations,
        int256[] memory offsetSeconds,
        uint256[] memory limits
    ) ERC20Base(owner, forges, name, symbol, decimals) ERC20PeriodsMintLimit(durations, offsetSeconds, limits) {}

    function mint(address to, uint256 amount) public override {
        ERC20PeriodsMintLimit.mint(to, amount);
    }
}

import {ERC20BasePreset} from "../ERC20Base.sol";

contract ERC20MultiMintLimitedPreset is ERC20BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC20MultiMintLimited).creationCode;
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

        (uint256[] memory durations, int256[] memory offsetSeconds, uint256[] memory limits) =
            abi.decode(data, (uint256[], int256[], uint256[]));

        return abi.encodePacked(
            code(), abi.encode(owner, forges, name, symbol, decimals, durations, offsetSeconds, limits)
        );
    }
}
