// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

enum TokenType {
    ERC20,
    ERC721,
    ERC1155
}

interface ITokenFactory {
    // ========== 이벤트 ==========

    event TokenDeployed(
        address owner, TokenType indexed tokenType, address indexed tokenAddress, address indexed presetAddress
    );

    event PresetLogicSet(TokenType indexed tokenType, address indexed presetAddress);

    event PresetLogicRemoved(TokenType indexed tokenType, address indexed presetAddress);

    // ========== 토큰 배포 함수 ==========

    function deployERC20(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes memory data,
        address preset
    ) external returns (address tokenAddress);

    function deployERC721(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        string memory baseTokenURI,
        bytes memory data,
        address preset
    ) external returns (address tokenAddress);

    function deployERC1155(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        string memory uri,
        bytes memory data,
        address preset
    ) external returns (address tokenAddress);

    // ========== 프리셋 관리 함수 ==========

    function setPresets(TokenType tokenType, address[] calldata presets, bool add) external;

    // ========== 조회 함수 ==========

    function getPresets(TokenType tokenType) external view returns (address[] memory);
}
