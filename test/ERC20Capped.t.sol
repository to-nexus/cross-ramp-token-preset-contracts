// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "forge-std-1.10.0/src/Test.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.4.0/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC20Capped} from "../src/erc20/presets/ERC20Capped.sol";

contract ERC20CappedTest is Test {
    ERC20Capped token;
    address owner = address(0x1);
    address manager = address(0x2);
    address recipient = address(0x3);
    string name = "CappedToken";
    string symbol = "CAP";
    uint8 decimals = 18;
    uint256 initialSupply = 1000 ether;
    uint256 cap = 2000 ether;

    function setUp() public {
        address tokenImpl = address(new ERC20Capped());
        bytes memory data = abi.encode(cap);
        address proxy = address(
            new ERC1967Proxy(
                tokenImpl,
                abi.encodeCall(
                    ERC20Capped.initialize, (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
                )
            )
        );
        token = ERC20Capped(proxy);
    }

    function test_initialize_set_values() public {
        assertEq(token.name(), name);
        assertEq(token.symbol(), symbol);
        assertEq(token.decimals(), decimals);
        assertEq(token.balanceOf(recipient), initialSupply);
    }

    function test_initialize_reverts_on_invalid_cap() public {
        address tokenImpl = address(new ERC20Capped());
        bytes memory data = hex"1234"; // Invalid data;
        vm.expectRevert(ERC20Capped.ERC20Capped__InvalidCapData.selector);
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20Capped.initialize, (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_initialize_reverts_on_cap_too_low() public {
        address tokenImpl = address(new ERC20Capped());
        bytes memory data = abi.encode(initialSupply - 1);
        vm.expectRevert(
            abi.encodeWithSelector(ERC20Capped.ERC20Capped__CapTooLow.selector, initialSupply - 1, initialSupply)
        );
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20Capped.initialize, (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_transfer() public {
        vm.prank(recipient);
        token.transfer(address(0x4), 100 ether);
        assertEq(token.balanceOf(address(0x4)), 100 ether);
    }
}
