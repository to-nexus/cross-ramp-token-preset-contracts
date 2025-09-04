// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC20CappedInitialSupply} from "../../src/erc20/presets/ERC20CappedInitialSupply.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC20CappedInitialSupplyTest is Test {
    ERC20CappedInitialSupply public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    uint256 public INITIAL_SUPPLY = 1000000 * 10 ** 18; // 1M tokens initial supply
    address public INITIAL_RECIPIENT;

    string constant NAME = "Mintable Token";
    string constant SYMBOL = "MINT";
    uint8 constant DECIMALS = 18;
    uint256 constant CAP = 1000000000 * 10 ** 18; // 1B tokens cap

    function setUp() public {
        owner = makeAddr("owner");
        forge1 = makeAddr("forge1");
        forge2 = makeAddr("forge2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        INITIAL_RECIPIENT = makeAddr("initial_recipient");

        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        bytes[2] memory extensionData = [abi.encode(CAP), abi.encode(INITIAL_SUPPLY, INITIAL_RECIPIENT)];
        token = new ERC20CappedInitialSupply(owner, forges, NAME, SYMBOL, DECIMALS, extensionData);
    }

    function test_initial_state() public view {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.decimals(), DECIMALS);
        assertEq(token.totalSupply(), INITIAL_SUPPLY);
        assertEq(token.owner(), owner);
        assertEq(token.cap(), CAP);
        assertEq(token.balanceOf(INITIAL_RECIPIENT), INITIAL_SUPPLY);
    }

    function test_mint() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY + amount);
    }

    function test_mint_multiple_times() public {
        uint256 amount1 = 1000 * 10 ** 18;
        uint256 amount2 = 2000 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, amount1);

        vm.prank(forge2);
        token.mint(user2, amount2);

        assertEq(token.balanceOf(user1), amount1);
        assertEq(token.balanceOf(user2), amount2);
        assertEq(token.totalSupply(), INITIAL_SUPPLY + amount1 + amount2);
    }

    function test_revert_when_mint_exceeds_cap() public {
        vm.prank(forge1);
        token.mint(user1, CAP - INITIAL_SUPPLY);

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

    function test_approve() public {
        uint256 amount = 1000 * 10 ** 18;

        vm.prank(user1);
        token.approve(user2, amount);

        assertEq(token.allowance(user1, user2), amount);
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
        assertEq(token.totalSupply(), INITIAL_SUPPLY + (amount / 2));
    }

    function test_mint_and_burn_cycle() public {
        uint256 amount = 1000 * 10 ** 18;

        // Mint tokens
        vm.prank(forge1);
        token.mint(user1, amount);

        // User1 approves user2 to burn
        vm.prank(user1);
        token.approve(user2, amount);

        // Burn half the tokens
        vm.prank(user2);
        token.burnFrom(user1, amount / 2);

        // Mint more tokens
        vm.prank(forge2);
        token.mint(user1, amount);

        assertEq(token.balanceOf(user1), amount + amount / 2);
        assertEq(token.totalSupply(), INITIAL_SUPPLY + (amount + amount / 2));
    }

    function test_multiple_forges_mint() public {
        uint256 amount = 500 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, amount);

        vm.prank(forge2);
        token.mint(user2, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(user2), amount);
        assertEq(token.totalSupply(), INITIAL_SUPPLY + (amount * 2));
    }
}
