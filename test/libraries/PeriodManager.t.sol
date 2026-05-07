// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {PeriodManager} from "../../src/libraries/PeriodManager.sol";
import {Test} from "forge-std-1.10.0/src/Test.sol";

contract PeriodManagerHarness {
    using PeriodManager for PeriodManager.PeriodConfig;

    PeriodManager.PeriodConfig public config;

    constructor(uint128 duration, int128 offsetSeconds) {
        config.duration = duration;
        config.offsetSeconds = offsetSeconds;
    }

    function getPeriodStartForTime(uint256 timestamp) external view returns (uint256) {
        return config.getPeriodStartForTime(timestamp);
    }
}

contract PeriodManagerTest is Test {
    uint128 constant DURATION = 86400;

    // ─── Positive offset ───────────────────────────────────────────────
    // duration=86400, offset=+3600
    // Periods: [3600,90000), [90000,176400), ...

    function test_PositiveOffset_BoundaryAtPeriodStart() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        assertEq(h.getPeriodStartForTime(3600), 3600, "t=3600 is Day0 start");
    }

    function test_PositiveOffset_MidPeriod() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        assertEq(h.getPeriodStartForTime(50000), 3600, "t=50000 mid-period returns Day0 start");
    }

    function test_PositiveOffset_BoundaryAtUTCMidnight() public {
        // t=86400 is UTC midnight — still BEFORE real boundary (90000)
        // BUG: old code returns 90000 (future period), correct returns 3600 (Day0)
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        assertEq(h.getPeriodStartForTime(86400), 3600, "t=86400 (UTC midnight) is still Day0");
    }

    function test_PositiveOffset_BoundaryJustBeforeRealEdge() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        assertEq(h.getPeriodStartForTime(89999), 3600, "t=89999 (1s before Day1) is still Day0");
    }

    function test_PositiveOffset_BoundaryAtRealEdge() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        assertEq(h.getPeriodStartForTime(90000), 90000, "t=90000 starts Day1");
    }

    function test_PositiveOffset_InNextPeriod() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        // t=172800 is in [90000,176400) → Day1 start = 90000
        assertEq(h.getPeriodStartForTime(172800), 90000, "t=172800 is in Day1");
    }

    function test_PositiveOffset_TimestampBelowOffsetReverts() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        vm.expectRevert(
            abi.encodeWithSelector(PeriodManager.PeriodManager__InvalidOffset.selector, uint256(3599), int256(3600))
        );
        h.getPeriodStartForTime(3599);
    }

    // ─── Invariant: periodStart <= t < periodStart + duration ──────────

    function test_PositiveOffset_InvariantReturnLETimestamp() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 3600);
        uint256[5] memory times = [uint256(3600), 50000, 86400, 89999, 90000];
        for (uint256 i = 0; i < times.length; i++) {
            uint256 start = h.getPeriodStartForTime(times[i]);
            assertLe(start, times[i], "period start must be <= timestamp");
            assertGt(start + DURATION, times[i], "timestamp must be within period");
        }
    }

    // ─── Negative offset ───────────────────────────────────────────────
    // duration=86400, offset=-3600
    // Periods: [82800,169200), [169200,255600), ...

    function test_NegativeOffset_BoundaryAtPeriodStart() public {
        // BUG: old code: floor(82800)=0, 0<3600 → revert. Correct: returns 82800
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        assertEq(h.getPeriodStartForTime(82800), 82800, "t=82800 is Day1 period start");
    }

    function test_NegativeOffset_BoundaryAtUTCMidnight() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        assertEq(h.getPeriodStartForTime(86400), 82800, "t=86400 (UTC midnight) is in Day1");
    }

    function test_NegativeOffset_JustBeforeRealEdge() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        assertEq(h.getPeriodStartForTime(169199), 82800, "t=169199 is still Day1");
    }

    function test_NegativeOffset_BoundaryAtRealEdge() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        assertEq(h.getPeriodStartForTime(169200), 169200, "t=169200 starts Day2");
    }

    function test_NegativeOffset_InSecondPeriod() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        assertEq(h.getPeriodStartForTime(172799), 169200, "t=172799 is in Day2");
    }

    function test_NegativeOffset_SmallTimestampReverts() public {
        // t=82799 → shifted=86399, floor=0 < 3600 → revert(82799, -3600)
        // BUG: old code reverts with (0, -3600) due to modified local var
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        vm.expectRevert(
            abi.encodeWithSelector(PeriodManager.PeriodManager__InvalidOffset.selector, uint256(82799), int256(-3600))
        );
        h.getPeriodStartForTime(82799);
    }

    // ─── Invariant: negative offset ─────────────────────────────────────

    function test_NegativeOffset_InvariantReturnLETimestamp() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, -3600);
        uint256[4] memory times = [uint256(82800), 86400, 169199, 169200];
        for (uint256 i = 0; i < times.length; i++) {
            uint256 start = h.getPeriodStartForTime(times[i]);
            assertLe(start, times[i], "period start must be <= timestamp");
            assertGt(start + DURATION, times[i], "timestamp must be within period");
        }
    }

    // ─── Zero offset (regression baseline) ─────────────────────────────

    function test_ZeroOffset_RegressionBaseline() public {
        PeriodManagerHarness h = new PeriodManagerHarness(DURATION, 0);
        assertEq(h.getPeriodStartForTime(0), 0, "t=0 zero-offset");
        assertEq(h.getPeriodStartForTime(86399), 0, "t=86399 is in period starting at 0");
        assertEq(h.getPeriodStartForTime(86400), 86400, "t=86400 starts new period");
        assertEq(h.getPeriodStartForTime(172800), 172800, "t=172800 zero-offset");
    }
}
