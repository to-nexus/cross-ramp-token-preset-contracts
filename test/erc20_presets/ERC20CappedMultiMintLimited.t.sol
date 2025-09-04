// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC20CappedMultiMintLimited} from "../../src/erc20/presets/ERC20CappedMultiMintLimited.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC20CappedMultiMintLimitedTest is Test {
    ERC20CappedMultiMintLimited public token;
    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    uint256 constant CAP = 10000000 * 10 ** 18; // 10M tokens cap
    string constant NAME = "Multi Mint Limited Token";
    string constant SYMBOL = "MMLT";
    uint8 constant DECIMALS = 18;

    // Period configurations
    uint256[] public durations;
    int256[] public offsetSeconds;
    uint256[] public limits;

    event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity);
    event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits);

    function setUp() public {
        vm.warp(604800);
        owner = makeAddr("owner");
        forge1 = makeAddr("forge1");
        forge2 = makeAddr("forge2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        // Setup period configurations
        // Period 1: 1 hour, 1000 tokens limit
        // Period 2: 1 day, 10000 tokens limit
        // Period 3: 1 week, 50000 tokens limit
        durations = new uint256[](3);
        durations[0] = 3600; // 1 hour
        durations[1] = 86400; // 1 day
        durations[2] = 604800; // 1 week

        offsetSeconds = new int256[](3);
        offsetSeconds[0] = 0; // No offset
        offsetSeconds[1] = 0; // No offset
        offsetSeconds[2] = 0; // No offset

        limits = new uint256[](3);
        limits[0] = 1000 * 10 ** 18; // 1K tokens per hour
        limits[1] = 10000 * 10 ** 18; // 10K tokens per day
        limits[2] = 50000 * 10 ** 18; // 100K tokens per week

        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        bytes[2] memory extensionData = [abi.encode(CAP), abi.encode(durations, offsetSeconds, limits)];
        token = new ERC20CappedMultiMintLimited(owner, forges, NAME, SYMBOL, DECIMALS, extensionData);
    }

    // ========== Initial State Tests ==========

    function test_initial_state() public {
        assertEq(token.name(), NAME);
        assertEq(token.symbol(), SYMBOL);
        assertEq(token.decimals(), DECIMALS);
        assertEq(token.cap(), CAP);
        assertEq(token.totalSupply(), 0);
        assertEq(token.owner(), owner);

        // Check period configurations
        (uint256[] memory returnedDurations, int256[] memory returnedOffsets) = token.periodConfigs();
        assertEq(returnedDurations.length, 3);
        assertEq(returnedOffsets.length, 3);
        assertEq(returnedDurations[0], durations[0]);
        assertEq(returnedDurations[1], durations[1]);
        assertEq(returnedDurations[2], durations[2]);

        // Check limits
        uint256[] memory maxLimits = token.maxMintPerPeriods();
        assertEq(maxLimits.length, 3);
        assertEq(maxLimits[0], limits[0]);
        assertEq(maxLimits[1], limits[1]);
        assertEq(maxLimits[2], limits[2]);

        // Check initial available capacities (should equal limits)
        uint256[] memory capacities = token.availableMintCapacities();
        assertEq(capacities.length, 3);
        assertEq(capacities[0], limits[0]);
        assertEq(capacities[1], limits[1]);
        assertEq(capacities[2], limits[2]);
    }

    function test_revert_when_invalid_constructor_params() public {
        address[] memory forges = new address[](1);
        forges[0] = forge1;

        // Test with mismatched array lengths
        uint256[] memory invalidDurations = new uint256[](2);
        invalidDurations[0] = 3600;
        invalidDurations[1] = 86400;

        vm.prank(owner);
        vm.expectRevert();
        bytes[2] memory extensionData = [abi.encode(CAP), abi.encode(invalidDurations, offsetSeconds, limits)];
        new ERC20CappedMultiMintLimited(owner, forges, NAME, SYMBOL, DECIMALS, extensionData);
    }

    function test_revert_when_zero_limits() public {
        address[] memory forges = new address[](1);
        forges[0] = forge1;

        uint256[] memory zeroLimits = new uint256[](3);
        zeroLimits[0] = 0; // Zero limit should revert
        zeroLimits[1] = 10000 * 10 ** 18;
        zeroLimits[2] = 50000 * 10 ** 18;

        vm.prank(owner);
        vm.expectRevert();
        bytes[2] memory extensionData = [abi.encode(CAP), abi.encode(durations, offsetSeconds, zeroLimits)];
        new ERC20CappedMultiMintLimited(owner, forges, NAME, SYMBOL, DECIMALS, extensionData);
    }

    // ========== Minting Tests ==========

    function test_mint_within_limits() public {
        uint256 amount = 500 * 10 ** 18; // Half of hourly limit

        vm.prank(forge1);
        token.mint(user1, amount);

        assertEq(token.balanceOf(user1), amount);
        assertEq(token.totalSupply(), amount);

        // Check available capacities
        uint256[] memory capacities = token.availableMintCapacities();
        assertEq(capacities[0], limits[0] - amount); // Hourly reduced
        assertEq(capacities[1], limits[1] - amount); // Daily reduced
        assertEq(capacities[2], limits[2] - amount); // Weekly reduced
    }

    function test_mint_upto_hourly_limit() public {
        vm.prank(forge1);
        token.mint(user1, limits[0]); // Mint full hourly limit

        assertEq(token.balanceOf(user1), limits[0]);

        // Check capacities
        uint256[] memory capacities = token.availableMintCapacities();
        assertEq(capacities[0], 0); // Hourly exhausted
        assertEq(capacities[1], limits[1] - limits[0]); // Daily reduced
        assertEq(capacities[2], limits[2] - limits[0]); // Weekly reduced
    }

    function test_revert_when_exceeds_hourly_limit() public {
        uint256 amount = limits[0] + 1; // Exceed hourly limit

        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, amount);
    }

    function test_revert_when_exceeds_daily_limit() public {
        // First mint most of daily limit
        {
            uint256 count = limits[1] / limits[0];
            for (uint256 i = 0; i < count; i++) {
                vm.prank(forge1);
                token.mint(user1, limits[0]);
                vm.warp(block.timestamp + 3601); // Move to next hour
            }
        }

        // Try to mint more than remaining daily capacity
        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, 1); // Would exceed daily limit
    }

    function test_revert_when_exceeds_weekly_limit() public {
        // Mint most of weekly limit in multiple transactions
        {
            uint256 count = limits[2] / limits[1];
            for (uint256 i = 0; i < count; i++) {
                uint256 startTime = block.timestamp;
                // Mint full daily limit
                {
                    uint256 dailyCount = limits[1] / limits[0];
                    for (uint256 j = 0; j < dailyCount; j++) {
                        vm.prank(forge1);
                        token.mint(user1, limits[0]);
                        vm.warp(block.timestamp + 3601); // Move to next hour
                    }
                }
                vm.warp(startTime + 86401); // Move to next day
            }
        }

        // Try to mint more than remaining weekly capacity
        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, 1); // Would exceed weekly limit
    }

    function test_mint_in_new_hourly_period() public {
        // Mint full hourly limit
        vm.prank(forge1);
        token.mint(user1, limits[0]);

        // Move to next hour
        vm.warp(block.timestamp + 3601);

        // Should be able to mint again (hourly limit reset)
        vm.prank(forge1);
        token.mint(user2, limits[0]);

        assertEq(token.balanceOf(user1), limits[0]);
        assertEq(token.balanceOf(user2), limits[0]);
        assertEq(token.totalSupply(), limits[0] * 2);
    }

    function test_mint_in_new_daily_period() public {
        // Mint full daily limit
        {
            uint256 count = limits[1] / limits[0];
            for (uint256 i = 0; i < count; i++) {
                vm.prank(forge1);
                token.mint(user1, limits[0]);
                vm.warp(block.timestamp + 3601); // Move to next hour
            }
        }

        // Move to next day
        vm.warp(block.timestamp + 86401);

        // Should be able to mint again (daily limit reset)
        {
            // Mint full daily limit
            uint256 count = limits[1] / limits[0];
            for (uint256 i = 0; i < count; i++) {
                vm.prank(forge1);
                token.mint(user2, limits[0]);
                vm.warp(block.timestamp + 3601); // Move to next hour
            }
        }

        assertEq(token.balanceOf(user1), limits[1]);
        assertEq(token.balanceOf(user2), limits[1]);
        assertEq(token.totalSupply(), limits[1] * 2);
    }

    function test_mint_in_new_weekly_period() public {
        // Mint full weekly limit
        {
            uint256 count = limits[2] / limits[1];
            for (uint256 i = 0; i < count; i++) {
                // Mint full daily limit
                {
                    uint256 dailyCount = limits[1] / limits[0];
                    for (uint256 j = 0; j < dailyCount; j++) {
                        vm.prank(forge1);
                        token.mint(user1, limits[0]);
                        vm.warp(block.timestamp + 3601); // Move to next hour
                    }
                }
                vm.warp(block.timestamp + 86401); // Move to next day
            }
        }

        // Move to next week
        vm.warp(block.timestamp + 604801);

        // Should be able to mint again (weekly limit reset)
        {
            uint256 count = limits[2] / limits[1];
            for (uint256 i = 0; i < count; i++) {
                // Mint full daily limit
                {
                    uint256 dailyCount = limits[1] / limits[0];
                    for (uint256 j = 0; j < dailyCount; j++) {
                        vm.prank(forge1);
                        token.mint(user2, limits[0]);
                        vm.warp(block.timestamp + 3601); // Move to next hour
                    }
                }
                vm.warp(block.timestamp + 86401); // Move to next day
            }
        }

        assertEq(token.balanceOf(user1), limits[2]);
        assertEq(token.balanceOf(user2), limits[2]);
        assertEq(token.totalSupply(), limits[2] * 2);
    }

    function test_multiple_mints_same_period() public {
        uint256 amount1 = 300 * 10 ** 18;
        uint256 amount2 = 400 * 10 ** 18;
        uint256 amount3 = 200 * 10 ** 18;

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

    // ========== Period Management Tests ==========

    function test_period_start_times() public {
        uint256[] memory startTimes = token.periodStartTimes();
        assertEq(startTimes.length, 3);

        // All periods should start at current block timestamp (aligned to period)
        for (uint256 i = 0; i < startTimes.length; i++) {
            assertTrue(startTimes[i] <= block.timestamp);
        }
    }

    function test_available_capacities_after_mint() public {
        uint256 mintAmount = 750 * 10 ** 18;

        vm.prank(forge1);
        token.mint(user1, mintAmount);

        uint256[] memory capacities = token.availableMintCapacities();
        assertEq(capacities[0], limits[0] - mintAmount); // Hourly
        assertEq(capacities[1], limits[1] - mintAmount); // Daily
        assertEq(capacities[2], limits[2] - mintAmount); // Weekly
    }

    function test_capacities_reset_after_period_change() public {
        uint256 mintAmount = 500 * 10 ** 18;

        // Mint some tokens
        vm.prank(forge1);
        token.mint(user1, mintAmount);

        // Move to next hour (hourly period resets)
        vm.warp(block.timestamp + 3601);

        uint256[] memory capacities = token.availableMintCapacities();
        assertEq(capacities[0], limits[0]); // Hourly reset
        assertEq(capacities[1], limits[1] - mintAmount); // Daily still reduced
        assertEq(capacities[2], limits[2] - mintAmount); // Weekly still reduced
    }

    // ========== Limit Update Tests ==========

    function test_update_mint_limits() public {
        uint256[] memory newLimits = new uint256[](3);
        newLimits[0] = 2000 * 10 ** 18; // Double hourly limit
        newLimits[1] = 20000 * 10 ** 18; // Double daily limit
        newLimits[2] = 200000 * 10 ** 18; // Double weekly limit

        vm.expectEmit(false, false, false, true);
        emit MintLimitUpdated(limits, newLimits);

        vm.prank(owner);
        token.updateMintLimits(newLimits);

        uint256[] memory updatedLimits = token.maxMintPerPeriods();
        assertEq(updatedLimits[0], newLimits[0]);
        assertEq(updatedLimits[1], newLimits[1]);
        assertEq(updatedLimits[2], newLimits[2]);
    }

    function test_revert_when_update_limits_invalid_length() public {
        uint256[] memory invalidLimits = new uint256[](2); // Wrong length
        invalidLimits[0] = 2000 * 10 ** 18;
        invalidLimits[1] = 20000 * 10 ** 18;

        vm.prank(owner);
        vm.expectRevert();
        token.updateMintLimits(invalidLimits);
    }

    function test_revert_when_update_limits_with_zero() public {
        uint256[] memory invalidLimits = new uint256[](3);
        invalidLimits[0] = 2000 * 10 ** 18;
        invalidLimits[1] = 0; // Zero should revert
        invalidLimits[2] = 200000 * 10 ** 18;

        vm.prank(owner);
        vm.expectRevert();
        token.updateMintLimits(invalidLimits);
    }

    function test_revert_when_update_limits_non_decreasing() public {
        uint256[] memory invalidLimits = new uint256[](3);
        invalidLimits[0] = 20000 * 10 ** 18; // Higher than daily
        invalidLimits[1] = 10000 * 10 ** 18; // Lower than hourly
        invalidLimits[2] = 200000 * 10 ** 18;

        vm.prank(owner);
        vm.expectRevert();
        token.updateMintLimits(invalidLimits);
    }

    function test_revert_when_update_limits_without_admin_role() public {
        uint256[] memory newLimits = new uint256[](3);
        newLimits[0] = 2000 * 10 ** 18;
        newLimits[1] = 20000 * 10 ** 18;
        newLimits[2] = 200000 * 10 ** 18;

        vm.prank(forge1);
        vm.expectRevert();
        token.updateMintLimits(newLimits);
    }

    function test_update_limits_with_max_value() public {
        uint256[] memory newLimits = new uint256[](3);
        newLimits[0] = 1000 * 10 ** 18;
        newLimits[1] = 10000 * 10 ** 18;
        newLimits[2] = type(uint256).max; // No limit for weekly

        vm.prank(owner);
        token.updateMintLimits(newLimits);

        uint256[] memory updatedLimits = token.maxMintPerPeriods();
        assertEq(updatedLimits[2], type(uint256).max);
    }

    // ========== Access Control Tests ==========

    function test_revert_when_non_forge_mints() public {
        vm.prank(user1);
        vm.expectRevert();
        token.mint(user1, 1000);

        vm.prank(owner);
        vm.expectRevert();
        token.mint(user1, 1000);
    }

    // ========== Token Function Tests ==========

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
        assertEq(token.totalSupply(), amount / 2);
    }

    // ========== Complex Scenario Tests ==========

    function test_complex_minting_scenario() public {
        // Scenario: Multiple mints across different periods

        // 1. Mint 800 tokens (within all limits)
        vm.prank(forge1);
        token.mint(user1, 800 * 10 ** 18);

        // 2. Move to next hour and mint 900 (within all limits)
        vm.warp(block.timestamp + 3601);
        vm.prank(forge1);
        token.mint(user2, 900 * 10 ** 18);

        // 3. Move to next hour and mint 1000 (within all limits)
        vm.warp(block.timestamp + 3601);
        vm.prank(forge1);
        token.mint(user1, 1000 * 10 ** 18);

        // 4. Move to next hour and mint 1001 (exceed hourly limit, should revert)
        vm.warp(block.timestamp + 3601);
        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, 1001 * 10 ** 18);
    }

    function test_limit_update_during_operation() public {
        // Mint some tokens
        vm.prank(forge1);
        token.mint(user1, 500 * 10 ** 18);

        // Update limits
        uint256[] memory newLimits = new uint256[](3);
        newLimits[0] = 1500 * 10 ** 18; // Increase hourly
        newLimits[1] = 15000 * 10 ** 18; // Increase daily
        newLimits[2] = 150000 * 10 ** 18; // Increase weekly

        vm.prank(owner);
        token.updateMintLimits(newLimits);

        vm.warp(block.timestamp + 3601); // Move to next hour
        // Should be able to mint more with new limits
        vm.prank(forge1);
        token.mint(user2, 1000 * 10 ** 18); // This would have failed with old hourly limit

        assertEq(token.totalSupply(), 1500 * 10 ** 18);
    }

    function test_revert_when_mint_exceeds_cap() public {
        // Mint close to cap over multiple periods
        uint256 limitCount = CAP / limits[2];
        for (uint256 i = 0; i < limitCount; i++) {
            uint256 startTime = block.timestamp;
            uint256 count = limits[2] / limits[1];
            for (uint256 j = 0; j < count; j++) {
                // Mint full daily limit
                uint256 startTime1 = block.timestamp;
                uint256 dailyCount = limits[1] / limits[0];
                for (uint256 k = 0; k < dailyCount; k++) {
                    vm.prank(forge1);
                    token.mint(user1, limits[0]);
                    vm.warp(block.timestamp + 3601); // Move to next hour
                }
                vm.warp(startTime1 + 86401); // Move to next day
            }
            vm.warp(startTime + 604800); // Move to next day
        }
        // Try to mint more - should revert due to cap
        vm.prank(forge1);
        vm.expectRevert();
        token.mint(user1, 1);
    }
}
