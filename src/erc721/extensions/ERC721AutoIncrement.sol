// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC721Base} from "../ERC721Base.sol";

abstract contract ERC721AutoIncrement is ERC721Base {
    uint256 private _currentTokenId;

    function mint(address to, uint256, bytes memory data) external override onlyForge returns (uint256) {
        uint256 tokenID = ++_currentTokenId;
        _mint(to, tokenID);
        return tokenID;
    }
}
