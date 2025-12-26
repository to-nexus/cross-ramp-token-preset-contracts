// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {ERC20DeflationaryFixed} from "../../src/erc20/presets/ERC20DeflationaryFixed.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC20DeflationaryFixedTest is Test {
    ERC20DeflationaryFixed public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;
    address public initialRecipient;

    uint256 constant INITIAL_SUPPLY = 1000000 * 10 ** 18; // 1M tokens
    string constant NAME = "Deflationary Fixed Token";
    string constant SYMBOL = "DEFLAT";
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
        token = new ERC20DeflationaryFixed(
            owner, forges, NAME, SYMBOL, DECIMALS, abi.encode(INITIAL_SUPPLY, initialRecipient)
        );
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.decimals(), DECIMALS);
        assertEq(token.totalSupply(), INITIAL_SUPPLY);
        assertEq(token.balanceOf(initialRecipient), INITIAL_SUPPLY);
        assertEq(token.owner(), owner);
    }

    function test_revert_when_mint_attempted_by_forge() public {
        vm.prank(forge1);
        vm.expectRevert(ERC20DeflationaryFixed.ERC20DeflationaryFixed__MintingNotAllowed.selector);
        token.mint(user1, 1000);
    }

    function test_revert_when_mint_attempted_by_owner() public {
        vm.prank(owner);
        vm.expectRevert(ERC20DeflationaryFixed.ERC20DeflationaryFixed__MintingNotAllowed.selector);
        token.mint(user1, 1000);
    }

    function test_revert_when_mint_attempted_by_user() public {
        vm.prank(user1);
        vm.expectRevert(ERC20DeflationaryFixed.ERC20DeflationaryFixed__MintingNotAllowed.selector);
        token.mint(user1, 1000);
    }

    function test_self_burn() public {
        uint256 burnAmount = 1000 * 10 ** 18;

        // Initial recipient burns their own tokens
        vm.prank(initialRecipient);
        token.burnFrom(initialRecipient, burnAmount);

        assertEq(token.balanceOf(initialRecipient), INITIAL_SUPPLY - burnAmount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount);
    }

    function test_burn_with_allowance() public {
        uint256 amount = 1000 * 10 ** 18;

        // Transfer some tokens to user1 first
        vm.prank(initialRecipient);
        token.transfer(user1, amount);

        // Approve user2 to burn from user1
        vm.prank(user1);
        token.approve(user2, amount);

        // User2 burns from user1
        vm.prank(user2);
        token.burnFrom(user1, amount / 2);

        assertEq(token.balanceOf(user1), amount / 2);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - amount / 2);
        assertEq(token.allowance(user1, user2), amount / 2);
    }

    function test_burn_reduces_total_supply() public {
        uint256 burnAmount1 = 100000 * 10 ** 18;
        uint256 burnAmount2 = 200000 * 10 ** 18;

        // First burn
        vm.prank(initialRecipient);
        token.burnFrom(initialRecipient, burnAmount1);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount1);

        // Second burn
        vm.prank(initialRecipient);
        token.burnFrom(initialRecipient, burnAmount2);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount1 - burnAmount2);
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

    function test_burn_without_allowance_reverts() public {
        uint256 amount = 1000 * 10 ** 18;

        // Transfer some tokens to user1 first
        vm.prank(initialRecipient);
        token.transfer(user1, amount);

        // User2 tries to burn from user1 without allowance - should revert
        vm.prank(user2);
        vm.expectRevert();
        token.burnFrom(user1, amount);
    }

    function test_total_supply_only_decreases() public {
        uint256 burnAmount = 500000 * 10 ** 18;

        // Burn some tokens
        vm.prank(initialRecipient);
        token.burnFrom(initialRecipient, burnAmount);

        // Total supply should be less than initial
        assertLt(token.totalSupply(), INITIAL_SUPPLY);
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount);

        // Try to mint (should fail)
        vm.prank(forge1);
        vm.expectRevert(ERC20DeflationaryFixed.ERC20DeflationaryFixed__MintingNotAllowed.selector);
        token.mint(user1, 1);

        // Total supply remains the same after failed mint attempt
        assertEq(token.totalSupply(), INITIAL_SUPPLY - burnAmount);
    }
}

