// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "forge-std-1.10.0/src/Test.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.4.0/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC20SingleMintLimited} from "../src/erc20/presets/ERC20SingleMintLimited.sol";

contract ERC20SingleMintLimitedTest is Test {
    ERC20SingleMintLimited token;
    address owner = address(0x1);
    address manager = address(0x2);
    address recipient = address(0x3);
    string name = "SingleMintLimitedToken";
    string symbol = "SML";
    uint8 decimals = 18;
    uint256 initialSupply = 1000 ether;
    uint256 cap = 2000 ether;
    uint256 duration = 1 days;
    int256 offset = 0;
    uint256 limit = 500 ether;

    function setUp() public {
        vm.warp(duration + 1);
        address tokenImpl = address(new ERC20SingleMintLimited());
        bytes memory data = abi.encode(cap, duration, offset, limit);
        address proxy = address(
            new ERC1967Proxy(
                tokenImpl,
                abi.encodeCall(
                    ERC20SingleMintLimited.initialize,
                    (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
                )
            )
        );
        token = ERC20SingleMintLimited(proxy);
        address[] memory forges = new address[](1);
        forges[0] = owner;
        vm.prank(owner);
        token.setForges(forges, true);
    }

    function test_initialize_set_values() public {
        assertEq(token.name(), name);
        assertEq(token.symbol(), symbol);
        assertEq(token.decimals(), decimals);
        assertEq(token.balanceOf(recipient), initialSupply);
    }

    function test_initialize_reverts_on_invalid_data() public {
        address tokenImpl = address(new ERC20SingleMintLimited());
        bytes memory data = hex"1234";
        vm.expectRevert(abi.encodeWithSignature("ERC20SingleMintLimited__InvalidInitialData()"));
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20SingleMintLimited.initialize,
                (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_initialize_reverts_on_cap_too_low() public {
        address tokenImpl = address(new ERC20SingleMintLimited());
        bytes memory data = abi.encode(initialSupply - 1, duration, offset, limit);
        vm.expectRevert(
            abi.encodeWithSelector(
                ERC20SingleMintLimited.ERC20SingleMintLimited__CapTooLow.selector, initialSupply - 1, initialSupply
            )
        );
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20SingleMintLimited.initialize,
                (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_mint_within_limit() public {
        vm.prank(owner);
        token.mint(recipient, limit);
        assertEq(token.balanceOf(recipient), initialSupply + limit);
    }

    function test_mint_over_limit_reverts() public {
        vm.startPrank(owner);
        token.mint(recipient, limit);
        vm.expectRevert(abi.encodeWithSignature("ERC20PeriodMintLimit__ExceedsPeriodLimit(uint256,uint256)", 1, 0));
        token.mint(recipient, 1);
        vm.stopPrank();
    }

    function test_mint_next_period() public {
        vm.startPrank(owner);
        token.mint(recipient, limit);
        vm.warp(block.timestamp + duration);
        token.mint(recipient, limit);
        vm.stopPrank();
    }

    function test_transfer() public {
        vm.prank(recipient);
        token.transfer(address(0x4), 100 ether);
        assertEq(token.balanceOf(address(0x4)), 100 ether);
    }
}
