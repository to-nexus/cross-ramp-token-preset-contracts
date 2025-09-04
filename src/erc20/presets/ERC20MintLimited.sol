// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base, ERC20PeriodMintLimit} from "../extensions/ERC20PeriodMintLimit.sol";

contract ERC20MintLimited is ERC20PeriodMintLimit {
    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes memory extensionData
    ) ERC20Base(owner, forges, name, symbol, decimals) ERC20PeriodMintLimit(extensionData) {}

    function mint(address to, uint256 amount) public override {
        ERC20PeriodMintLimit.mint(to, amount);
    }
}

import {ERC20BasePreset} from "../ERC20Base.sol";

contract ERC20MintLimitedPreset is ERC20BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC20MintLimited).creationCode;
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

        return abi.encodePacked(code(), abi.encode(owner, forges, name, symbol, decimals, data));
    }
}
