// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import "forge-std-1.10.0/src/Test.sol";
import {ERC1967Proxy} from "@openzeppelin-contracts-5.4.0/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC20Fixed} from "../src/erc20/presets/ERC20Fixed.sol";
import {TokenBase} from "../src/TokenBase.sol";

contract ERC20FixedTest is Test {
    ERC20Fixed token;
    address owner = address(0x1);
    address manager = address(0x2);
    address recipient = address(0x3);
    string name = "FixedToken";
    string symbol = "FIX";
    uint8 decimals = 18;
    uint256 initialSupply = 1000 ether;

    function setUp() public {
        address tokenImpl = address(new ERC20Fixed());
        bytes memory data;
        address proxy = address(
            new ERC1967Proxy(
                tokenImpl,
                abi.encodeCall(
                    ERC20Fixed.initialize, (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
                )
            )
        );
        token = ERC20Fixed(proxy);
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

    function test_mint_reverts() public {
        vm.prank(owner);
        vm.expectRevert(ERC20Fixed.ERC20Fixed__MintingNotAllowed.selector);
        token.mint(recipient, 1 ether);
    }

    function test_burn_reverts() public {
        vm.expectRevert(ERC20Fixed.ERC20Fixed__BurningNotAllowed.selector);
        token.burnFrom(recipient, 1 ether);
    }

    function test_initialize_reverts_on_zero_supply() public {
        address tokenImpl = address(new ERC20Fixed());
        initialSupply = 0;
        bytes memory data = abi.encode(uint256(0));
        vm.expectRevert(abi.encodeWithSelector(TokenBase.TokenBase__NullInput.selector, bytes32("initialSupply")));
        new ERC1967Proxy(
            tokenImpl,
            abi.encodeCall(
                ERC20Fixed.initialize, (owner, manager, name, symbol, decimals, initialSupply, recipient, data)
            )
        );
    }

    function test_transfer() public {
        vm.prank(recipient);
        token.transfer(address(0x4), 100 ether);
        assertEq(token.balanceOf(address(0x4)), 100 ether);
    }
}
