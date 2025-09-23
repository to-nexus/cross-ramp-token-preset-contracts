// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC1155Base} from "../ERC1155Base.sol";

contract ERC1155Simple is ERC1155Base {
    constructor(address owner, address[] memory forges, string memory name, string memory symbol, string memory uri)
        ERC1155Base(owner, forges, name, symbol, uri)
    {}
}

import {ERC1155BasePreset} from "../ERC1155Base.sol";

contract ERC1155SimplePreset is ERC1155BasePreset {
    function code() public pure override returns (bytes memory) {
        return type(ERC1155Simple).creationCode;
    }

    function deployCode(bytes memory initialData) external pure override returns (bytes memory) {
        (address owner, address[] memory forges, string memory name, string memory symbol, string memory uri,) =
            abi.decode(initialData, (address, address[], string, string, string, bytes));

        return abi.encodePacked(code(), abi.encode(owner, forges, name, symbol, uri));
    }
}
