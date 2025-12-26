// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC721Base} from "../ERC721Base.sol";

contract ERC721Simple is ERC721Base {
    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        string memory baseTokenURI
    ) ERC721Base(owner, forges, name, symbol, baseTokenURI) {}
}

import {ERC721BasePreset} from "../ERC721Base.sol";

contract ERC721SimplePreset is ERC721BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC721Simple).creationCode;
    }

    function deployCode(bytes memory initialData) external pure override returns (bytes memory) {
        (
            address owner,
            address[] memory forges,
            string memory name,
            string memory symbol,
            string memory baseTokenURI,
        ) = abi.decode(initialData, (address, address[], string, string, string, bytes));

        return abi.encodePacked(code(), abi.encode(owner, forges, name, symbol, baseTokenURI));
    }
}
