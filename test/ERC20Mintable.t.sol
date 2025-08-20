// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "forge-std-1.10.0/src/Test.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.4.0/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC20Mintable} from "../src/erc20/presets/ERC20Mintable.sol";

contract ERC20MintableTest is Test {
    ERC20Mintable token;
    address owner = address(0x1);
    address manager = address(0x2);
    address recipient = address(0x3);
    string name = "MintableToken";
    string symbol = "MINT";
    uint8 decimals = 18;
    uint256 initialSupply = 1000 ether;

    function setUp() public {
        address tokenImpl = address(new ERC20Mintable());
        bytes memory data;
        address proxy = address(
            new ERC1967Proxy(
                tokenImpl,
                abi.encodeCall(
                    ERC20Mintable.initialize, (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
                )
            )
        );
        token = ERC20Mintable(proxy);
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

    function test_mint() public {
        vm.prank(owner);
        token.mint(recipient, 1 ether);
        assertEq(token.balanceOf(recipient), initialSupply + 1 ether);
    }

    function test_transfer() public {
        vm.prank(recipient);
        token.transfer(address(0x4), 100 ether);
        assertEq(token.balanceOf(address(0x4)), 100 ether);
    }
}
