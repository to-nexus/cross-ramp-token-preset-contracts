// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC1155} from "@openzeppelin-contracts-5.4.0/token/ERC1155/ERC1155.sol";

import {TokenBase} from "../TokenBase.sol";
import {IERC1155Forge} from "../interfaces/IERC1155Forge.sol";

abstract contract ERC1155Base is TokenBase, IERC1155Forge, ERC1155 {
    event ERC1155BaseSetBaseURI(string before, string current);

    string private _name;
    string private _symbol;

    constructor(address owner, address[] memory forges_, string memory name_, string memory symbol_, string memory uri_)
        TokenBase(owner, forges_)
        ERC1155(uri_)
    {
        if (bytes(name_).length == 0) revert TokenBase__NullInput("name");
        if (bytes(symbol_).length == 0) revert TokenBase__NullInput("symbol");
        if (bytes(uri_).length == 0) revert TokenBase__NullInput("uri");
        _name = name_;
        _symbol = symbol_;
    }

    function mint(address to, uint256 tokenID, uint256 amount, bytes calldata data) external onlyForge {
        _mint(to, tokenID, amount, data);
    }

    function mintBatch(address to, uint256[] memory tokenIDs, uint256[] memory amounts, bytes calldata data)
        external
        onlyForge
    {
        _mintBatch(to, tokenIDs, amounts, data);
    }

    function burnFrom(address from, uint256 tokenID, uint256 amount) external {
        _burn(from, tokenID, amount);
    }

    function burnFromBatch(address from, uint256[] memory tokenIDs, uint256[] memory amounts) external {
        _burnBatch(from, tokenIDs, amounts);
    }

    function name() public view virtual returns (string memory) {
        return _name;
    }

    function symbol() public view virtual returns (string memory) {
        return _symbol;
    }

    function uri(uint256 id) public view virtual override returns (string memory) {
        return ERC1155.uri(id);
    }

    function balanceOf(address account, uint256 id) public view virtual override returns (uint256) {
        return ERC1155.balanceOf(account, id);
    }

    function balanceOfBatch(address[] memory accounts, uint256[] memory ids)
        public
        view
        virtual
        override
        returns (uint256[] memory)
    {
        return ERC1155.balanceOfBatch(accounts, ids);
    }

    function setApprovalForAll(address operator, bool approved) public virtual override {
        ERC1155.setApprovalForAll(operator, approved);
    }

    function isApprovedForAll(address account, address operator) public view virtual override returns (bool) {
        return ERC1155.isApprovedForAll(account, operator);
    }

    function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes memory data)
        public
        virtual
        override
    {
        ERC1155.safeTransferFrom(from, to, id, value, data);
    }

    function safeBatchTransferFrom(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory values,
        bytes memory data
    ) public virtual override {
        ERC1155.safeBatchTransferFrom(from, to, ids, values, data);
    }

    function _update(address from, address to, uint256[] memory ids, uint256[] memory values)
        internal
        virtual
        override
    {
        ERC1155._update(from, to, ids, values);
    }

    function _updateWithAcceptanceCheck(
        address from,
        address to,
        uint256[] memory ids,
        uint256[] memory values,
        bytes memory data
    ) internal virtual override {
        ERC1155._updateWithAcceptanceCheck(from, to, ids, values, data);
    }

    function _setURI(string memory newuri) internal virtual override {
        ERC1155._setURI(newuri);
    }

    function _setApprovalForAll(address owner, address operator, bool approved) internal virtual override {
        ERC1155._setApprovalForAll(owner, operator, approved);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override(TokenBase, ERC1155) returns (bool) {
        return interfaceId == type(IPreset).interfaceId || interfaceId == type(IERC1155Forge).interfaceId
            || interfaceId == type(IERC1155Errors).interfaceId || super.supportsInterface(interfaceId);
    }
}

import {IPreset} from "../interfaces/IPreset.sol";
import {IERC1155Errors} from "@openzeppelin-contracts-5.4.0/interfaces/draft-IERC6093.sol";
import {IERC1155} from "@openzeppelin-contracts-5.4.0/token/ERC1155/IERC1155.sol";
import {IERC1155MetadataURI} from "@openzeppelin-contracts-5.4.0/token/ERC1155/extensions/IERC1155MetadataURI.sol";

abstract contract ERC1155BasePreset is IPreset {
    function code() external pure virtual override returns (bytes memory);

    function deployCode(bytes memory initialData) external pure virtual override returns (bytes memory);

    function supportsInterface(bytes4 interfaceId) external view virtual override returns (bool) {
        return interfaceId == type(IPreset).interfaceId || interfaceId == type(IERC1155Forge).interfaceId
            || interfaceId == type(IERC1155).interfaceId || interfaceId == type(IERC1155MetadataURI).interfaceId
            || interfaceId == type(IERC1155Errors).interfaceId || interfaceId == 0x01ffc9a7; // ERC165
    }
}
