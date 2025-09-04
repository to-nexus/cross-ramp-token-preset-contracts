// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {SafeCast} from "@openzeppelin-contracts-5.4.0/utils/math/SafeCast.sol";

import {PeriodManager} from "../../libraries/PeriodManager.sol";
import {ERC20Base} from "../ERC20Base.sol";

abstract contract ERC20PeriodsMintLimit is ERC20Base {
    using PeriodManager for PeriodManager.PeriodConfig;

    error ERC20PeriodsMintLimit__InvalidLength();
    error ERC20PeriodsMintLimit__InvalidLimitData(uint256 index);
    error ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256 period, uint256 requested, uint256 available);

    event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity);
    event MintLimitUpdated(uint256[] oldLimits, uint256[] newLimits);

    uint256 private _length;
    PeriodManager.PeriodConfig[] private _periods;
    uint256[] private _limits;
    uint256[] private _periodStartTimes;
    uint256[] private _periodCapacities;

    constructor(bytes memory data) {
        (uint256[] memory durations, int256[] memory offsetSeconds, uint256[] memory limits) =
            abi.decode(data, (uint256[], int256[], uint256[]));
        uint256 length = limits.length;
        if (length == 0 || length != durations.length || length != offsetSeconds.length) {
            revert ERC20PeriodsMintLimit__InvalidLength();
        }

        uint256 minDuration = 0;
        uint256 minLimit = 0;
        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                (uint256 duration, uint256 limit) = (durations[i], limits[i]);
                if (duration == 0 || limit == 0) {
                    revert TokenBase__NullInput("limits or durations");
                }
                if (duration <= minDuration || limit <= minLimit) {
                    revert ERC20PeriodsMintLimit__InvalidLimitData(i);
                }
                _periods.push(
                    PeriodManager.PeriodConfig(SafeCast.toUint128(duration), SafeCast.toInt128(offsetSeconds[i]))
                );
                (minDuration, minLimit) = (duration, limit);
            }
        }

        _length = length;
        _limits = limits;
        _periodStartTimes = new uint256[](length);
        _periodCapacities = new uint256[](length);
    }

    function mint(address to, uint256 amount) public virtual override {
        uint256[] memory currentPeriodStartTimes = periodStartTimes();
        uint256 length = _length;

        unchecked {
            for (uint256 i = 0; i < length; ++i) {
                (uint256 periodCapacity, uint256 periodStartTime, uint256 currentPeriodStartTime) =
                    (_periodCapacities[i], _periodStartTimes[i], currentPeriodStartTimes[i]);

                // Check if the period has started
                if (periodStartTime != currentPeriodStartTime) {
                    // Initialize the period start time if not set
                    uint256 limit = _limits[i];
                    _periodStartTimes[i] = currentPeriodStartTime;
                    periodCapacity = limit;

                    emit PeriodStarted(currentPeriodStartTime, limit);
                }

                if (periodCapacity == type(uint256).max) {
                    // If the period capacity is set to max, it means no limit for this period
                    _periodCapacities[i] = periodCapacity;
                } else {
                    // Check available capacity
                    if (periodCapacity < amount) {
                        revert ERC20PeriodsMintLimit__ExceedsPeriodLimit(_periods[i].duration, amount, periodCapacity);
                    }
                    // Update available capacity
                    _periodCapacities[i] = (periodCapacity - amount);
                }
            }
        }

        // Mint the tokens
        super.mint(to, amount);
    }

    function periodConfigs() external view returns (uint256[] memory, int256[] memory) {
        uint256 length = _length;
        uint256[] memory durations = new uint256[](length);
        int256[] memory offsets = new int256[](length);
        for (uint256 i = 0; i < length;) {
            unchecked {
                durations[i] = _periods[i].duration;
                offsets[i] = _periods[i].offsetSeconds;
                ++i;
            }
        }
        return (durations, offsets);
    }

    function maxMintPerPeriods() external view returns (uint256[] memory) {
        return _limits;
    }

    function availableMintCapacities() external view returns (uint256[] memory) {
        uint256[] memory currentPeriodStartTimes = periodStartTimes();
        uint256[] memory periodStartTimes_ = _periodStartTimes;

        uint256 length = _length;
        uint256[] memory capacities = new uint256[](length);
        for (uint256 i = 0; i < length;) {
            uint256 _periodStart = periodStartTimes_[i];
            if (_periodStart == currentPeriodStartTimes[i]) {
                capacities[i] = _periodCapacities[i];
            } else {
                capacities[i] = _limits[i];
            }
            unchecked {
                ++i;
            }
        }
        return capacities;
    }

    function periodStartTimes() public view returns (uint256[] memory) {
        uint256 length = _length;
        uint256[] memory startTimes = new uint256[](length);
        for (uint256 i = 0; i < length;) {
            startTimes[i] = _periods[i].getCurrentPeriodStart();
            unchecked {
                ++i;
            }
        }
        return startTimes;
    }

    function updateMintLimits(uint256[] calldata newLimits) external onlyRole(DEFAULT_ADMIN_ROLE) {
        // Validate new limit
        uint256 length = newLimits.length;
        if (length == 0) revert TokenBase__NullInput("newLimits");

        if (length != _length) {
            revert ERC20PeriodsMintLimit__InvalidLength();
        }
        uint256 minLimit = 0;
        for (uint256 i = 0; i < length;) {
            uint256 newLimit = newLimits[i];
            if (newLimit == 0) revert TokenBase__NullInput("newLimits");
            if (newLimit <= minLimit && newLimit != type(uint256).max) {
                revert ERC20PeriodsMintLimit__InvalidLimitData(i);
            }
            minLimit = newLimit;
            unchecked {
                ++i;
            }
        }

        // Update limit
        emit MintLimitUpdated(_limits, newLimits);
        _limits = newLimits;
    }
}
