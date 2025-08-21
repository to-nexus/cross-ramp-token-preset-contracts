// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "forge-std-1.10.0/src/Test.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.4.0/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC20PeriodsMintLimit, ERC20MultiMintLimited} from "../src/erc20/presets/ERC20MultiMintLimited.sol";

contract ERC20MultiMintLimitedTest is Test {
    ERC20MultiMintLimited token;
    address owner = address(0x1);
    address manager = address(0x2);
    address recipient = address(0x3);
    string name = "MultiMintLimitedToken";
    string symbol = "MML";
    uint8 decimals = 18;
    uint256 initialSupply = 1000 ether;
    uint256 cap = 2000 ether;
    uint256[] durations = [uint256(1 days), 7 days];
    int256[] offsets = [int256(0), 0];
    uint256[] limits = [100 ether, 500 ether];

    function setUp() public {
        vm.warp(durations[1] + 1);
        address tokenImpl = address(new ERC20MultiMintLimited());
        bytes memory data = abi.encode(cap, abi.encode(durations, offsets, limits));
        address proxy = address(
            new ERC1967Proxy(
                tokenImpl,
                abi.encodeCall(
                    ERC20MultiMintLimited.initialize,
                    (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
                )
            )
        );
        token = ERC20MultiMintLimited(proxy);
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
        address tokenImpl = address(new ERC20MultiMintLimited());
        bytes memory data = hex"1234";
        vm.expectRevert();
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20MultiMintLimited.initialize,
                (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_initialize_reverts_on_cap_too_low() public {
        address tokenImpl = address(new ERC20MultiMintLimited());
        bytes memory data = abi.encode(uint256(initialSupply - 1), abi.encode(durations, offsets, limits));
        vm.expectRevert(
            abi.encodeWithSignature("ERC20Capable__ERC20ExceededCap(uint256,uint256)", initialSupply, initialSupply - 1)
        );
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20MultiMintLimited.initialize,
                (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_mint_within_limit() public {
        vm.prank(owner);
        token.mint(recipient, limits[0]);
        assertEq(token.balanceOf(recipient), initialSupply + limits[0]);
    }

    function test_mint_over_limit_reverts() public {
        vm.prank(owner);
        token.mint(recipient, limits[0]);
        vm.expectRevert(
            abi.encodeWithSignature(
                "ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256,uint256,uint256)", durations[0], 1, 0
            )
        );
        vm.prank(owner);
        token.mint(recipient, 1);
    }

    function test_mint_next_period() public {
        vm.startPrank(owner);
        token.mint(recipient, limits[0]);
        vm.warp(block.timestamp + durations[0]);
        token.mint(recipient, limits[0]);
        vm.stopPrank();
    }

    function test_mint_over_limit_reverts_2() public {
        vm.startPrank(owner);
        for (uint256 i = 0; i < 5; i++) {
            token.mint(recipient, limits[0]);
            vm.warp(block.timestamp + durations[0]);
        }
        for (uint256 i = 0; i < 2; i++) {
            vm.expectRevert(
                abi.encodeWithSignature(
                    "ERC20PeriodsMintLimit__ExceedsPeriodLimit(uint256,uint256,uint256)", durations[1], 1, 0
                )
            );
            token.mint(recipient, 1);
            vm.warp(block.timestamp + durations[0]);
        }
        // success
        token.mint(recipient, 1);
        vm.stopPrank();
    }

    function test_transfer() public {
        vm.prank(recipient);
        token.transfer(address(0x4), 100 ether);
        assertEq(token.balanceOf(address(0x4)), 100 ether);
    }
}
