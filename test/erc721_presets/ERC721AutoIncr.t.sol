// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {ERC721AutoIncr} from "../../src/erc721/presets/ERC721AutoIncr.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC721AutoIncrTest is Test {
    ERC721AutoIncr public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    string constant NAME = "AutoIncrement NFT";
    string constant SYMBOL = "AINFT";
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
        token = new ERC721AutoIncr(owner, forges, NAME, SYMBOL, BASE_TOKEN_URI);
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.owner(), owner);
        // Check if forge1 and forge2 have forge role
        assertTrue(token.isForge(forge1));
        assertTrue(token.isForge(forge2));
    }

    function test_mint_auto_increment_starts_from_1() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 mintedTokenId = token.mint(user1, 999, data); // input tokenId is ignored

        assertEq(mintedTokenId, 1); // Should be 1 regardless of input
        assertEq(token.ownerOf(1), user1);
        assertEq(token.balanceOf(user1), 1);
    }

    function test_mint_auto_increment_sequential() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId1 = token.mint(user1, 100, data); // input ignored

        vm.prank(forge2);
        uint256 tokenId2 = token.mint(user2, 200, data); // input ignored

        vm.prank(forge1);
        uint256 tokenId3 = token.mint(user1, 300, data); // input ignored

        assertEq(tokenId1, 1);
        assertEq(tokenId2, 2);
        assertEq(tokenId3, 3);

        assertEq(token.ownerOf(1), user1);
        assertEq(token.ownerOf(2), user2);
        assertEq(token.ownerOf(3), user1);
        assertEq(token.balanceOf(user1), 2);
        assertEq(token.balanceOf(user2), 1);
    }

    function test_mint_multiple_tokens_to_same_user() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId1 = token.mint(user1, 0, data);

        vm.prank(forge1);
        uint256 tokenId2 = token.mint(user1, 0, data);

        vm.prank(forge1);
        uint256 tokenId3 = token.mint(user1, 0, data);

        assertEq(tokenId1, 1);
        assertEq(tokenId2, 2);
        assertEq(tokenId3, 3);
        assertEq(token.balanceOf(user1), 3);
    }

    function test_mint_fails_for_non_forge() public {
        bytes memory data = "";

        vm.prank(user1);
        vm.expectRevert();
        token.mint(user1, 1, data);
    }

    function test_no_duplicate_token_ids() public {
        bytes memory data = "";

        // Mint 10 tokens and verify they're all unique and sequential
        for (uint256 i = 1; i <= 10; i++) {
            vm.prank(forge1);
            uint256 tokenId = token.mint(user1, 999, data); // input always ignored
            assertEq(tokenId, i);
            assertEq(token.ownerOf(i), user1);
        }

        assertEq(token.balanceOf(user1), 10);
    }

    function test_burn_from_by_owner() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId = token.mint(user1, 0, data);

        vm.prank(user1);
        token.burnFrom(user1, tokenId);

        vm.expectRevert();
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(user1), 0);
    }

    function test_burn_from_by_approved() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId = token.mint(user1, 0, data);

        vm.prank(user1);
        token.approve(user2, tokenId);

        vm.prank(user2);
        token.burnFrom(user1, tokenId);

        vm.expectRevert();
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(user1), 0);
    }

    function test_burn_from_fails_for_unauthorized() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId = token.mint(user1, 0, data);

        vm.prank(user2);
        vm.expectRevert();
        token.burnFrom(user1, tokenId);
    }

    function test_token_uri() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId = token.mint(user1, 0, data);

        string memory expectedUri = string(abi.encodePacked(BASE_TOKEN_URI, "1"));
        assertEq(token.tokenURI(tokenId), expectedUri);
    }

    function test_transfer() public {
        bytes memory data = "";

        vm.prank(forge1);
        uint256 tokenId = token.mint(user1, 0, data);

        vm.prank(user1);
        token.transferFrom(user1, user2, tokenId);

        assertEq(token.ownerOf(tokenId), user2);
        assertEq(token.balanceOf(user1), 0);
        assertEq(token.balanceOf(user2), 1);
    }

    function test_mint_after_burn_continues_increment() public {
        bytes memory data = "";

        // Mint tokens 1, 2, 3
        vm.prank(forge1);
        uint256 tokenId1 = token.mint(user1, 0, data);
        vm.prank(forge1);
        uint256 tokenId2 = token.mint(user1, 0, data);
        vm.prank(forge1);
        uint256 tokenId3 = token.mint(user1, 0, data);

        assertEq(tokenId1, 1);
        assertEq(tokenId2, 2);
        assertEq(tokenId3, 3);

        // Burn token 2
        vm.prank(user1);
        token.burnFrom(user1, 2);

        // Mint new token should be 4 (not reusing burned tokenId)
        vm.prank(forge1);
        uint256 tokenId4 = token.mint(user2, 0, data);
        assertEq(tokenId4, 4);

        assertEq(token.balanceOf(user1), 2); // has tokens 1 and 3
        assertEq(token.balanceOf(user2), 1); // has token 4
    }

    function test_mint_with_different_input_values() public {
        bytes memory data = "";

        // Test that input tokenId is completely ignored
        vm.prank(forge1);
        uint256 tokenId1 = token.mint(user1, type(uint256).max, data);

        vm.prank(forge1);
        uint256 tokenId2 = token.mint(user1, 0, data);

        vm.prank(forge1);
        uint256 tokenId3 = token.mint(user1, 12345, data);

        assertEq(tokenId1, 1);
        assertEq(tokenId2, 2);
        assertEq(tokenId3, 3);
    }
}
