// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC20Fixed} from "../../src/erc20/presets/ERC20Fixed.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC20FixedTest is Test {
    ERC20Fixed public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;
    address public initialRecipient;

    uint256 constant INITIAL_SUPPLY = 1000000 * 10 ** 18; // 1M tokens
    string constant NAME = "Fixed Token";
    string constant SYMBOL = "FIXED";
    uint8 constant DECIMALS = 18;

    function setUp() public {
        owner = makeAddr("owner");
        forge1 = makeAddr("forge1");
        forge2 = makeAddr("forge2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        initialRecipient = makeAddr("initialRecipient");

        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        token = new ERC20Fixed(owner, forges, NAME, SYMBOL, DECIMALS, INITIAL_SUPPLY, initialRecipient);
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.decimals(), DECIMALS);
        assertEq(token.totalSupply(), INITIAL_SUPPLY);
        assertEq(token.balanceOf(initialRecipient), INITIAL_SUPPLY);
        assertEq(token.owner(), owner);
    }

    function test_revert_when_mint_attempted() public {
        vm.prank(forge1);
        vm.expectRevert(ERC20Fixed.ERC20Fixed__MintingNotAllowed.selector);
        token.mint(user1, 1000);
    }

    function test_revert_when_burnFrom_attempted() public {
        uint256 amount = 1000 * 10 ** 18;

        // Transfer some tokens to user1 first
        vm.prank(initialRecipient);
        token.transfer(user1, amount);

        // Approve user2 to burn from user1
        vm.prank(user1);
        token.approve(user2, amount);

        // Try to burn - should revert
        vm.prank(user2);
        vm.expectRevert(ERC20Fixed.ERC20Fixed__BurningNotAllowed.selector);
        token.burnFrom(user1, amount);
    }

    function test_transfer() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(initialRecipient);
        token.transfer(user1, amount);

        assertEq(token.balanceOf(initialRecipient), INITIAL_SUPPLY - amount);
        assertEq(token.balanceOf(user1), amount);
    }

    function test_approve() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(initialRecipient);
        token.approve(user1, amount);

        assertEq(token.allowance(initialRecipient, user1), amount);
    }

    function test_transferFrom() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(initialRecipient);
        token.approve(user1, amount);

        vm.prank(user1);
        token.transferFrom(initialRecipient, user2, amount);

        assertEq(token.balanceOf(initialRecipient), INITIAL_SUPPLY - amount);
        assertEq(token.balanceOf(user2), amount);
        assertEq(token.allowance(initialRecipient, user1), 0);
    }

    function test_totalSupply_remains_fixed() public {
        // Transfer tokens around
        vm.prank(initialRecipient);
        token.transfer(user1, 500000 * 10 ** 18);

        vm.prank(user1);
        token.transfer(user2, 100000 * 10 ** 18);

        // Total supply should remain the same
        assertEq(token.totalSupply(), INITIAL_SUPPLY);
    }

    function test_non_forge_cannot_mint() public {
        vm.prank(owner);
        vm.expectRevert(ERC20Fixed.ERC20Fixed__MintingNotAllowed.selector);
        token.mint(user1, 1000);

        vm.prank(user1);
        vm.expectRevert(ERC20Fixed.ERC20Fixed__MintingNotAllowed.selector);
        token.mint(user1, 1000);
    }
}
