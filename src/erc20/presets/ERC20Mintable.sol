// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";

contract ERC20Mintable is ERC20Base {
    constructor(address owner, address[] memory forges, string memory name, string memory symbol, uint8 decimals)
        ERC20Base(owner, forges, name, symbol, decimals)
    {}
}

import {ERC20BasePreset} from "../ERC20Base.sol";

contract ERC20MintablePreset is ERC20BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC20Mintable).creationCode;
    }

    function deployCode(bytes memory initialData) external pure override returns (bytes memory) {
        (address owner, address[] memory forges, string memory name, string memory symbol, uint8 decimals,) =
            abi.decode(initialData, (address, address[], string, string, uint8, bytes));

        return abi.encodePacked(code(), abi.encode(owner, forges, name, symbol, decimals));
    }
}
