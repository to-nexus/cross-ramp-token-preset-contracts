// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {ERC721Simple} from "../../src/erc721/presets/ERC721Simple.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC721SimpleTest is Test {
    ERC721Simple public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    string constant NAME = "Simple NFT";
    string constant SYMBOL = "SNFT";
    string constant BASE_TOKEN_URI = "https://api.example.com/metadata/";

    function setUp() public {
        owner = makeAddr("owner");
        forge1 = makeAddr("forge1");
        forge2 = makeAddr("forge2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        token = new ERC721Simple(owner, forges, NAME, SYMBOL, BASE_TOKEN_URI);
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.owner(), owner);
        // Check if forge1 and forge2 have forge role
        assertTrue(token.isForge(forge1));
        assertTrue(token.isForge(forge2));
    }

    function test_mint_by_forge() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        uint256 mintedTokenId = token.mint(user1, tokenId, data);

        assertEq(mintedTokenId, tokenId);
        assertEq(token.ownerOf(tokenId), user1);
        assertEq(token.balanceOf(user1), 1);
    }

    function test_mint_multiple_tokens() public {
        uint256 tokenId1 = 1;
        uint256 tokenId2 = 2;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId1, data);

        vm.prank(forge2);
        token.mint(user2, tokenId2, data);

        assertEq(token.ownerOf(tokenId1), user1);
        assertEq(token.ownerOf(tokenId2), user2);
        assertEq(token.balanceOf(user1), 1);
        assertEq(token.balanceOf(user2), 1);
    }

    function test_mint_fails_for_non_forge() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(user1);
        vm.expectRevert();
        token.mint(user1, tokenId, data);
    }

    function test_mint_fails_for_duplicate_token_id() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user2, tokenId, data);
    }

    function test_burn_from_by_owner() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        vm.prank(user1);
        token.burnFrom(user1, tokenId);

        vm.expectRevert();
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(user1), 0);
    }

    function test_burn_from_by_approved() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        vm.prank(user1);
        token.approve(user2, tokenId);

        vm.prank(user2);
        token.burnFrom(user1, tokenId);

        vm.expectRevert();
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(user1), 0);
    }

    function test_burn_from_fails_for_unauthorized() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        vm.prank(user2);
        vm.expectRevert();
        token.burnFrom(user1, tokenId);
    }

    function test_token_uri() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        string memory expectedUri = string(abi.encodePacked(BASE_TOKEN_URI, "1"));
        assertEq(token.tokenURI(tokenId), expectedUri);
    }

    function test_transfer() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        vm.prank(user1);
        token.transferFrom(user1, user2, tokenId);

        assertEq(token.ownerOf(tokenId), user2);
        assertEq(token.balanceOf(user1), 0);
        assertEq(token.balanceOf(user2), 1);
    }

    function test_approve() public {
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, data);

        vm.prank(user1);
        token.approve(user2, tokenId);

        assertEq(token.getApproved(tokenId), user2);
    }

    function test_set_approval_for_all() public {
        vm.prank(user1);
        token.setApprovalForAll(user2, true);

        assertTrue(token.isApprovedForAll(user1, user2));

        vm.prank(user1);
        token.setApprovalForAll(user2, false);

        assertFalse(token.isApprovedForAll(user1, user2));
    }

    function test_supports_interface() public view {
        // ERC721
        assertTrue(token.supportsInterface(0x80ac58cd));
        // ERC721Metadata
        assertTrue(token.supportsInterface(0x5b5e139f));
        // ERC165
        assertTrue(token.supportsInterface(0x01ffc9a7));
    }

    function test_owner_can_add_forge() public {
        address newForge = makeAddr("newForge");

        vm.prank(owner);
        address[] memory forges = new address[](1);
        forges[0] = newForge;
        token.setForges(forges, true);

        assertTrue(token.isForge(newForge));

        // Test that new forge can mint
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(newForge);
        token.mint(user1, tokenId, data);

        assertEq(token.ownerOf(tokenId), user1);
    }

    function test_owner_can_remove_forge() public {
        vm.prank(owner);
        address[] memory forges = new address[](1);
        forges[0] = forge1;
        token.setForges(forges, false);

        assertFalse(token.isForge(forge1));

        // Test that removed forge cannot mint
        uint256 tokenId = 1;
        bytes memory data = "";

        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, tokenId, data);
    }

    function test_non_owner_cannot_add_forge() public {
        address newForge = makeAddr("newForge");

        address[] memory forges = new address[](1);
        forges[0] = newForge;

        vm.prank(user1);
        vm.expectRevert();
        token.setForges(forges, true);
    }
}
