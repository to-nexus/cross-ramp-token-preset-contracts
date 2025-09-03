// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {SafeCast} from "@openzeppelin-contracts-5.4.0/utils/math/SafeCast.sol";

import {ERC20Base} from "../ERC20Base.sol";
import {PeriodManager} from "../../libraries/PeriodManager.sol";

abstract contract ERC20PeriodMintLimit is ERC20Base {
    using PeriodManager for PeriodManager.PeriodConfig;

    error ERC20PeriodMintLimit__InvalidLength();
    error ERC20PeriodMintLimit__InvalidLimitData();
    error ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256 requested, uint256 available);

    event PeriodStarted(uint256 indexed periodStart, uint256 availableCapacity);
    event MintLimitUpdated(uint256 oldLimits, uint256 newLimits);

    PeriodManager.PeriodConfig private _period;
    uint256 private _periodStartTime; // The block.timestamp when the current period started
    uint256 private _limit; // The maximum amount that can be minted in a period
    uint256 private _periodCapacity; // The remaining capacity for the current period

    constructor(uint256 duration, int256 offsetSeconds, uint256 limit) {
        if (duration == 0 || limit == 0) {
            revert TokenBase__NullInput("limit or period");
        }

        _period = PeriodManager.PeriodConfig(SafeCast.toUint128(duration), SafeCast.toInt128(offsetSeconds));
        _limit = limit;
    }

    function mint(address to, uint256 amount) public virtual override {
        uint256 periodCapacity = _periodCapacity;
        {
            uint256 currentPeriodStart = _period.getCurrentPeriodStart();
            // Check if the period has started
            if (currentPeriodStart != _periodStartTime) {
                // Initialize the period start time if not set
                _periodStartTime = currentPeriodStart;
                periodCapacity = _limit;

                emit PeriodStarted(currentPeriodStart, periodCapacity);
            }
        }

        // Check available capacity
        if (periodCapacity < amount) {
            revert ERC20PeriodMintLimit__ExceedsPeriodLimit(amount, periodCapacity);
        }
        // Update available capacity
        unchecked {
            _periodCapacity = periodCapacity - amount;
        }

        // Mint the tokens
        super.mint(to, amount);
    }

    function periodConfig() external view returns (uint256, int256) {
        return (_period.duration, _period.offsetSeconds);
    }

    function maxMintPerPeriod() external view returns (uint256) {
        return _limit;
    }

    function availableMintCapacity() external view returns (uint256) {
        return _period.isNewPeriod(_periodStartTime) ? _limit : _periodCapacity;
    }

    function periodStartTime() public view returns (uint256) {
        return _period.getCurrentPeriodStart();
    }

    function updateMintLimit(uint256 newLimit) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newLimit == 0) revert TokenBase__NullInput("newLimit");

        // Update limit
        emit MintLimitUpdated(_limit, newLimit);
        _limit = newLimit;
    }
}
