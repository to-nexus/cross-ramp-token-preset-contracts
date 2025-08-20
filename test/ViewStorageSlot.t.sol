// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "forge-std-1.10.0/src/Test.sol";

contract ERC20CappedTest is Test {
    function setUp() external {}

    function test_slot_token_base() external pure {
        bytes32 slot =
            keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.TokenBase")) - 1)) & ~bytes32(uint256(0xff));
        console.log("TokenBase storage slot:");
        console.logBytes32(slot);
    }

    function test_slot_erc20_decimals() external pure {
        bytes32 slot =
            keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.decimals")) - 1)) & ~bytes32(uint256(0xff));
        console.log("ERC20 decimals storage slot:");
        console.logBytes32(slot);
    }

    function test_slot_erc20_capable() external pure {
        bytes32 slot = keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.ERC20Capable")) - 1))
            & ~bytes32(uint256(0xff));
        console.log("ERC20Capable storage slot:");
        console.logBytes32(slot);
    }

    function test_slot_erc20_period_mint_limit() external pure {
        bytes32 slot = keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.ERC20PeriodMintLimit")) - 1))
            & ~bytes32(uint256(0xff));
        console.log("ERC20PeriodMintLimit storage slot:");
        console.logBytes32(slot);
    }

    function test_slot_erc20_periods_mint_limit() external pure {
        bytes32 slot = keccak256(abi.encode(uint256(keccak256("cross.ramp.storage.erc20.ERC20PeriodsMintLimit")) - 1))
            & ~bytes32(uint256(0xff));
        console.log("ERC20PeriodsMintLimit storage slot:");
        console.logBytes32(slot);
    }
}
