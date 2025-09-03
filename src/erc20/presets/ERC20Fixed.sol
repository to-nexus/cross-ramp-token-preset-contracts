// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base, ERC20InitialSupply} from "../extensions/ERC20InitialSupply.sol";

contract ERC20Fixed is ERC20InitialSupply {
    error ERC20Fixed__MintingNotAllowed();
    error ERC20Fixed__BurningNotAllowed();

    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        uint256 initialSupply,
        address initialRecipient
    ) ERC20Base(owner, forges, name, symbol, decimals) ERC20InitialSupply(initialSupply, initialRecipient) {}

    function mint(address, uint256) public pure override {
        revert ERC20Fixed__MintingNotAllowed();
    }

    function burnFrom(address, uint256) public pure override {
        revert ERC20Fixed__BurningNotAllowed();
    }
}

import {ERC20BasePreset} from "../ERC20Base.sol";

contract ERC20FixedPreset is ERC20BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC20Fixed).creationCode;
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

        (uint256 initialSupply, address initialRecipient) = abi.decode(data, (uint256, address));
        return
            abi.encodePacked(code(), abi.encode(owner, forges, name, symbol, decimals, initialSupply, initialRecipient));
    }
}
