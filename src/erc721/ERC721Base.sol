// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC721} from "@openzeppelin-contracts-5.4.0/token/ERC721/ERC721.sol";

import {TokenBase} from "../TokenBase.sol";
import {IERC721Forge} from "../interfaces/IERC721Forge.sol";

abstract contract ERC721Base is TokenBase, IERC721Forge, ERC721 {
    event ERC721BaseSetBaseURI(string before, string current);

    string private _baseTokenURI;

    constructor(
        address owner,
        address[] memory forges_,
        string memory name_,
        string memory symbol_,
        string memory baseTokenURI_
    ) TokenBase(owner, forges_) ERC721(name_, symbol_) {
        if (bytes(name_).length == 0) revert TokenBase__NullInput("name");
        if (bytes(symbol_).length == 0) revert TokenBase__NullInput("symbol");
        if (bytes(baseTokenURI_).length == 0) revert TokenBase__NullInput("baseTokenURI");
        _setBaseURI(baseTokenURI_);
    }

    function mint(address to, uint256 tokenID, bytes memory data)
        external
        virtual
        override
        onlyForge
        returns (uint256)
    {
        _mint(to, tokenID);
        return tokenID;
    }

    function burnFrom(address from, uint256 tokenID) external virtual override {
        if (from != _requireOwned(tokenID)) revert IERC721Errors.ERC721InvalidOwner(from);
        _checkAuthorized(from, msg.sender, tokenID);
        _burn(tokenID);
    }

    function _baseURI() internal view override returns (string memory) {
        return _baseTokenURI;
    }

    function _setBaseURI(string memory baseTokenURI) internal {
        emit ERC721BaseSetBaseURI(_baseTokenURI, baseTokenURI);
        _baseTokenURI = baseTokenURI;
    }

    function balanceOf(address owner) public view virtual override returns (uint256) {
        return ERC721.balanceOf(owner);
    }

    function ownerOf(uint256 tokenId) public view virtual override returns (address) {
        return ERC721.ownerOf(tokenId);
    }

    function name() public view virtual override returns (string memory) {
        return ERC721.name();
    }

    function symbol() public view virtual override returns (string memory) {
        return ERC721.symbol();
    }

    function tokenURI(uint256 tokenId) public view virtual override returns (string memory) {
        return ERC721.tokenURI(tokenId);
    }

    function approve(address to, uint256 tokenId) public virtual override {
        ERC721.approve(to, tokenId);
    }

    function getApproved(uint256 tokenId) public view virtual override returns (address) {
        return ERC721.getApproved(tokenId);
    }

    function setApprovalForAll(address operator, bool approved) public virtual override {
        ERC721.setApprovalForAll(operator, approved);
    }

    function isApprovedForAll(address owner, address operator) public view virtual override returns (bool) {
        return ERC721.isApprovedForAll(owner, operator);
    }

    function transferFrom(address from, address to, uint256 tokenId) public virtual override {
        ERC721.transferFrom(from, to, tokenId);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public virtual override {
        ERC721.safeTransferFrom(from, to, tokenId, data);
    }

    function _ownerOf(uint256 tokenId) internal view virtual override returns (address) {
        return ERC721._ownerOf(tokenId);
    }

    function _getApproved(uint256 tokenId) internal view virtual override returns (address) {
        return ERC721._getApproved(tokenId);
    }

    function _isAuthorized(address owner, address spender, uint256 tokenId)
        internal
        view
        virtual
        override
        returns (bool)
    {
        return ERC721._isAuthorized(owner, spender, tokenId);
    }

    function _checkAuthorized(address owner, address spender, uint256 tokenId) internal view virtual override {
        ERC721._checkAuthorized(owner, spender, tokenId);
    }

    function _increaseBalance(address account, uint128 value) internal virtual override {
        ERC721._increaseBalance(account, value);
    }

    function _update(address to, uint256 tokenId, address auth) internal virtual override returns (address) {
        return ERC721._update(to, tokenId, auth);
    }

    function _safeMint(address to, uint256 tokenId, bytes memory data) internal virtual override {
        ERC721._safeMint(to, tokenId, data);
    }

    function _safeTransfer(address from, address to, uint256 tokenId, bytes memory data) internal virtual override {
        ERC721._safeTransfer(from, to, tokenId, data);
    }

    function _approve(address to, uint256 tokenId, address auth, bool emitEvent) internal virtual override {
        ERC721._approve(to, tokenId, auth, emitEvent);
    }

    function _setApprovalForAll(address owner, address operator, bool approved) internal virtual override {
        ERC721._setApprovalForAll(owner, operator, approved);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(TokenBase, ERC721) returns (bool) {
        return interfaceId == type(IPreset).interfaceId || interfaceId == type(IERC721Forge).interfaceId
            || interfaceId == type(IERC721).interfaceId || interfaceId == type(IERC721Metadata).interfaceId
            || interfaceId == type(IERC721Errors).interfaceId || interfaceId == 0x01ffc9a7; // ERC165
    }
}

import {IPreset} from "../interfaces/IPreset.sol";
import {IERC721Errors} from "@openzeppelin-contracts-5.4.0/interfaces/draft-IERC6093.sol";
import {IERC721} from "@openzeppelin-contracts-5.4.0/token/ERC721/IERC721.sol";
import {IERC721Metadata} from "@openzeppelin-contracts-5.4.0/token/ERC721/extensions/IERC721Metadata.sol";

abstract contract ERC721BasePreset is IPreset {
    function code() external pure virtual override returns (bytes memory);

    function deployCode(bytes memory initialData) external pure virtual override returns (bytes memory);

    function supportsInterface(bytes4 interfaceId) external view virtual override returns (bool) {
        return interfaceId == type(IPreset).interfaceId || interfaceId == type(IERC721Forge).interfaceId
            || interfaceId == type(IERC721).interfaceId || interfaceId == type(IERC721Metadata).interfaceId
            || interfaceId == type(IERC721Errors).interfaceId || interfaceId == 0x01ffc9a7; // ERC165
    }
}
