// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1155Simple} from "../../src/erc1155/presets/ERC1155Simple.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC1155SimpleTest is Test {
    ERC1155Simple public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    string constant NAME = "Simple Multi Token";
    string constant SYMBOL = "SMT";
    string constant URI = "https://api.example.com/metadata/{id}.json";

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
        token = new ERC1155Simple(owner, forges, NAME, SYMBOL, URI);
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.owner(), owner);
        assertEq(token.uri(0), URI);
        // Check if forge1 and forge2 have forge role
        assertTrue(token.isForge(forge1));
        assertTrue(token.isForge(forge2));
    }

    function test_mint_single_token() public {
        uint256 tokenId = 1;
        uint256 amount = 100;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, amount, data);

        assertEq(token.balanceOf(user1, tokenId), amount);
    }

    function test_mint_multiple_tokens_same_user() public {
        uint256 tokenId1 = 1;
        uint256 tokenId2 = 2;
        uint256 amount1 = 100;
        uint256 amount2 = 200;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId1, amount1, data);

        vm.prank(forge2);
        token.mint(user1, tokenId2, amount2, data);

        assertEq(token.balanceOf(user1, tokenId1), amount1);
        assertEq(token.balanceOf(user1, tokenId2), amount2);
    }

    function test_mint_same_token_different_users() public {
        uint256 tokenId = 1;
        uint256 amount = 100;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, amount, data);

        vm.prank(forge2);
        token.mint(user2, tokenId, amount, data);

        assertEq(token.balanceOf(user1, tokenId), amount);
        assertEq(token.balanceOf(user2, tokenId), amount);
    }

    function test_mint_batch() public {
        uint256[] memory tokenIds = new uint256[](3);
        tokenIds[0] = 1;
        tokenIds[1] = 2;
        tokenIds[2] = 3;

        uint256[] memory amounts = new uint256[](3);
        amounts[0] = 100;
        amounts[1] = 200;
        amounts[2] = 300;

        bytes memory data = "";

        vm.prank(forge1);
        token.mintBatch(user1, tokenIds, amounts, data);

        assertEq(token.balanceOf(user1, tokenIds[0]), amounts[0]);
        assertEq(token.balanceOf(user1, tokenIds[1]), amounts[1]);
        assertEq(token.balanceOf(user1, tokenIds[2]), amounts[2]);
    }

    function test_mint_fails_for_non_forge() public {
        uint256 tokenId = 1;
        uint256 amount = 100;
        bytes memory data = "";

        vm.prank(user1);
        vm.expectRevert();
        token.mint(user1, tokenId, amount, data);
    }

    function test_burn_from() public {
        uint256 tokenId = 1;
        uint256 mintAmount = 100;
        uint256 burnAmount = 30;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, mintAmount, data);

        vm.prank(user1);
        token.burnFrom(user1, tokenId, burnAmount);

        assertEq(token.balanceOf(user1, tokenId), mintAmount - burnAmount);
    }

    function test_burn_from_batch() public {
        uint256[] memory tokenIds = new uint256[](2);
        tokenIds[0] = 1;
        tokenIds[1] = 2;

        uint256[] memory mintAmounts = new uint256[](2);
        mintAmounts[0] = 100;
        mintAmounts[1] = 200;

        uint256[] memory burnAmounts = new uint256[](2);
        burnAmounts[0] = 30;
        burnAmounts[1] = 50;

        bytes memory data = "";

        vm.prank(forge1);
        token.mintBatch(user1, tokenIds, mintAmounts, data);

        vm.prank(user1);
        token.burnFromBatch(user1, tokenIds, burnAmounts);

        assertEq(token.balanceOf(user1, tokenIds[0]), mintAmounts[0] - burnAmounts[0]);
        assertEq(token.balanceOf(user1, tokenIds[1]), mintAmounts[1] - burnAmounts[1]);
    }

    function test_safe_transfer_from() public {
        uint256 tokenId = 1;
        uint256 amount = 100;
        uint256 transferAmount = 30;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, amount, data);

        vm.prank(user1);
        token.safeTransferFrom(user1, user2, tokenId, transferAmount, data);

        assertEq(token.balanceOf(user1, tokenId), amount - transferAmount);
        assertEq(token.balanceOf(user2, tokenId), transferAmount);
    }

    function test_safe_batch_transfer_from() public {
        uint256[] memory tokenIds = new uint256[](2);
        tokenIds[0] = 1;
        tokenIds[1] = 2;

        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 100;
        amounts[1] = 200;

        uint256[] memory transferAmounts = new uint256[](2);
        transferAmounts[0] = 30;
        transferAmounts[1] = 50;

        bytes memory data = "";

        vm.prank(forge1);
        token.mintBatch(user1, tokenIds, amounts, data);

        vm.prank(user1);
        token.safeBatchTransferFrom(user1, user2, tokenIds, transferAmounts, data);

        assertEq(token.balanceOf(user1, tokenIds[0]), amounts[0] - transferAmounts[0]);
        assertEq(token.balanceOf(user1, tokenIds[1]), amounts[1] - transferAmounts[1]);
        assertEq(token.balanceOf(user2, tokenIds[0]), transferAmounts[0]);
        assertEq(token.balanceOf(user2, tokenIds[1]), transferAmounts[1]);
    }

    function test_approval_for_all() public {
        vm.prank(user1);
        token.setApprovalForAll(user2, true);

        assertTrue(token.isApprovedForAll(user1, user2));

        vm.prank(user1);
        token.setApprovalForAll(user2, false);

        assertFalse(token.isApprovedForAll(user1, user2));
    }

    function test_transfer_with_approval() public {
        uint256 tokenId = 1;
        uint256 amount = 100;
        uint256 transferAmount = 30;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, amount, data);

        vm.prank(user1);
        token.setApprovalForAll(user2, true);

        vm.prank(user2);
        token.safeTransferFrom(user1, user2, tokenId, transferAmount, data);

        assertEq(token.balanceOf(user1, tokenId), amount - transferAmount);
        assertEq(token.balanceOf(user2, tokenId), transferAmount);
    }

    function test_transfer_fails_without_approval() public {
        uint256 tokenId = 1;
        uint256 amount = 100;
        uint256 transferAmount = 30;
        bytes memory data = "";

        vm.prank(forge1);
        token.mint(user1, tokenId, amount, data);

        vm.prank(user2);
        vm.expectRevert();
        token.safeTransferFrom(user1, user2, tokenId, transferAmount, data);
    }

    function test_balance_of_batch() public {
        uint256[] memory tokenIds = new uint256[](3);
        tokenIds[0] = 1;
        tokenIds[1] = 2;
        tokenIds[2] = 3;

        uint256[] memory amounts = new uint256[](3);
        amounts[0] = 100;
        amounts[1] = 200;
        amounts[2] = 300;

        address[] memory accounts = new address[](3);
        accounts[0] = user1;
        accounts[1] = user1;
        accounts[2] = user1;

        bytes memory data = "";

        vm.prank(forge1);
        token.mintBatch(user1, tokenIds, amounts, data);

        uint256[] memory balances = token.balanceOfBatch(accounts, tokenIds);

        assertEq(balances[0], amounts[0]);
        assertEq(balances[1], amounts[1]);
        assertEq(balances[2], amounts[2]);
    }

    function test_supports_interface() public view {
        // ERC1155
        assertTrue(token.supportsInterface(0xd9b67a26));
        // ERC1155MetadataURI
        assertTrue(token.supportsInterface(0x0e89341c));
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
        uint256 amount = 100;
        bytes memory data = "";

        vm.prank(newForge);
        token.mint(user1, tokenId, amount, data);

        assertEq(token.balanceOf(user1, tokenId), amount);
    }

    function test_owner_can_remove_forge() public {
        vm.prank(owner);
        address[] memory forges = new address[](1);
        forges[0] = forge1;
        token.setForges(forges, false);

        assertFalse(token.isForge(forge1));

        // Test that removed forge cannot mint
        uint256 tokenId = 1;
        uint256 amount = 100;
        bytes memory data = "";

        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, tokenId, amount, data);
    }

    function test_non_owner_cannot_add_forge() public {
        address newForge = makeAddr("newForge");

        address[] memory forges = new address[](1);
        forges[0] = newForge;

        vm.prank(user1);
        vm.expectRevert();
        token.setForges(forges, true);
    }

    function test_uri() public view {
        assertEq(token.uri(1), URI);
        assertEq(token.uri(999), URI);
    }
}
