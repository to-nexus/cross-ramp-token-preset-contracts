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

    /// @custom:storage-location erc7201:cross.ramp.storage.erc20.ERC20PeriodMintLimit
    struct ERC20PeriodMintLimitStorage {
        PeriodManager.PeriodConfig period;
        uint256 periodStartTime; // The block.timestamp when the current period started
        uint256 limit; // The maximum amount that can be minted in a period
        uint256 periodCapacity; // The remaining capacity for the current period
    }

    // keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.ERC20PeriodMintLimit")) - 1)) & ~bytes32(uint256(0xff))
    bytes32 private constant ERC20PeriodMintLimitStorageLocation =
        0xbfe38164f4301e1b504c75c70e5c430bf30509ee0f9b81dc0d9a573797e93f00;

    function _getERC20PeriodMintLimitStorage() private pure returns (ERC20PeriodMintLimitStorage storage $) {
        assembly {
            $.slot := ERC20PeriodMintLimitStorageLocation
        }
    }

    function __ERC20PeriodMintLimit_init(uint256 duration, int256 offsetSeconds, uint256 limit)
        internal
        onlyInitializing
    {
        if (duration == 0 || limit == 0) {
            revert TokenBase__NullInput("limit or period");
        }

        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();
        $.period = PeriodManager.PeriodConfig(SafeCast.toUint128(duration), SafeCast.toInt128(offsetSeconds));
        $.limit = limit;
    }

    function mint(address to, uint256 amount) public virtual override {
        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();

        uint256 periodCapacity = $.periodCapacity;
        {
            uint256 currentPeriodStart = $.period.getCurrentPeriodStart();
            // Check if the period has started
            if (currentPeriodStart != $.periodStartTime) {
                // Initialize the period start time if not set
                $.periodStartTime = currentPeriodStart;
                periodCapacity = $.limit;

                emit PeriodStarted(currentPeriodStart, periodCapacity);
            }
        }

        // Check available capacity
        if (periodCapacity < amount) {
            revert ERC20PeriodMintLimit__ExceedsPeriodLimit(amount, periodCapacity);
        }
        // Update available capacity
        unchecked {
            $.periodCapacity = periodCapacity - amount;
        }

        // Mint the tokens
        super.mint(to, amount);
    }

    function periodConfig() external view returns (uint256, int256) {
        PeriodManager.PeriodConfig storage period = _getERC20PeriodMintLimitStorage().period;
        return (period.duration, period.offsetSeconds);
    }

    function maxMintPerPeriod() external view returns (uint256) {
        return _getERC20PeriodMintLimitStorage().limit;
    }

    function availableMintCapacity() external view returns (uint256) {
        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();
        return $.period.isNewPeriod($.periodStartTime) ? $.limit : $.periodCapacity;
    }

    function periodStartTime() public view returns (uint256) {
        return _getERC20PeriodMintLimitStorage().period.getCurrentPeriodStart();
    }

    function updateMintLimit(uint256 newLimit) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newLimit == 0) revert TokenBase__NullInput("newLimit");

        ERC20PeriodMintLimitStorage storage $ = _getERC20PeriodMintLimitStorage();
        // Update limit
        emit MintLimitUpdated($.limit, newLimit);
        $.limit = newLimit;
    }
}
