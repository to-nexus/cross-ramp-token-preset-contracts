// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC20CappedMintLimited} from "../../src/erc20/presets/ERC20CappedMintLimited.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC20CappedMintLimitedTest is Test {
    ERC20CappedMintLimited public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    uint256 constant CAP = 10000000 * 10 ** 18; // 10M tokens cap
    uint256 constant DURATION = 86400; // 1 day in seconds
    int256 constant OFFSET_SECONDS = 0; // No offset
    uint256 constant LIMIT = 100000 * 10 ** 18; // 100K tokens per period
    string constant NAME = "Mint Limited Token";
    string constant SYMBOL = "MLIMIT";
    uint8 constant DECIMALS = 18;

    function setUp() public {
        vm.warp(86401); // Move past any potential time-based restrictions
        owner = makeAddr("owner");
        forge1 = makeAddr("forge1");
        forge2 = makeAddr("forge2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        bytes[2] memory extensionData = [abi.encode(CAP), abi.encode(DURATION, OFFSET_SECONDS, LIMIT)];
        token = new ERC20CappedMintLimited(owner, forges, NAME, SYMBOL, DECIMALS, extensionData);
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.decimals(), DECIMALS);
        assertEq(token.cap(), CAP);
        assertEq(token.totalSupply(), 0);
        assertEq(token.owner(), owner);
    }

    function test_mint_within_limit() public {
        uint256 amount = 50000 * 10 ** 18; // Half of the limit

        vm.prank(forge1);
        token.mint(user1, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.totalSupply(), amount);
    }

    function test_mint_upto_limit() public {
        vm.prank(forge1);
        token.mint(user1, LIMIT);

        assertEq(token.balanceOf(user1), LIMIT);
        assertEq(token.totalSupply(), LIMIT);
    }

    function test_revert_when_mint_exceeds_limit() public {
        uint256 amount = LIMIT + 1;

        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, amount);
    }

    function test_revert_when_mint_exceeds_limit_in_multiple_calls() public {
        uint256 amount1 = 60000 * 10 ** 18;
        uint256 amount2 = 50000 * 10 ** 18; // This would exceed the limit

        vm.prank(forge1);
        token.mint(user1, amount1);

        vm.prank(forge2);
        vm.expectRevert();
        token.mint(user2, amount2);
    }

    function test_mint_in_new_period() public {
        uint256 amount = LIMIT;

        // Mint in first period
        vm.prank(forge1);
        token.mint(user1, amount);

        // Move to next period
        vm.warp(block.timestamp + DURATION + 1);

        // Should be able to mint again
        vm.prank(forge2);
        token.mint(user2, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(user2), amount);
        assertEq(token.totalSupply(), amount * 2);
    }

    function test_revert_when_mint_exceeds_cap() public {
        // Mint close to cap over multiple periods
        uint256 periodsNeeded = CAP / LIMIT;

        for (uint256 i = 0; i < periodsNeeded; i++) {
            vm.warp(block.timestamp + DURATION * i + 1);
            vm.prank(forge1);
            token.mint(user1, LIMIT);
        }

        // Try to mint more - should revert due to cap
        vm.warp(block.timestamp + DURATION + 1);
        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, 1);
    }

    function test_revert_when_non_forge_mints() public {
        vm.prank(user1);
        vm.expectRevert();
        token.mint(user1, 1000);

        vm.prank(owner);
        vm.expectRevert();
        token.mint(user1, 1000);
    }

    function test_transfer() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, amount);

        vm.prank(user1);
        token.transfer(user2, amount / 2);

        assertEq(token.balanceOf(user1), amount / 2);
        assertEq(token.balanceOf(user2), amount / 2);
    }

    function test_burnFrom() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, amount);

        vm.prank(user1);
        token.approve(user2, amount);

        vm.prank(user2);
        token.burnFrom(user1, amount / 2);

        assertEq(token.balanceOf(user1), amount / 2);
        assertEq(token.totalSupply(), amount / 2);
    }

    function test_multiple_mints_same_period() public {
        uint256 amount1 = 30000 * 10 ** 18;
        uint256 amount2 = 40000 * 10 ** 18;
        uint256 amount3 = 30000 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, amount1);

        vm.prank(forge2);
        token.mint(user2, amount2);

        vm.prank(forge1);
        token.mint(user1, amount3);

        assertEq(token.balanceOf(user1), amount1 + amount3);
        assertEq(token.balanceOf(user2), amount2);
        assertEq(token.totalSupply(), amount1 + amount2 + amount3);
    }

    function test_period_limit_reset() public {
        uint256 amount = 80000 * 10 ** 18;

        // Mint in first period
        vm.prank(forge1);
        token.mint(user1, amount);

        // Try to mint more in same period - should fail
        vm.prank(forge2);
        vm.expectRevert();
        token.mint(user2, 30000 * 10 ** 18);

        // Move to next period
        vm.warp(block.timestamp + DURATION + 1);

        // Should be able to mint again
        vm.prank(forge2);
        token.mint(user2, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(user2), amount);
    }
}
