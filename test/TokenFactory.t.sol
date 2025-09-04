// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {TokenFactoryImpl} from "../src/TokenFactoryImpl.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Test, console} from "forge-std-1.10.0/src//Test.sol";

import {ERC20Capped, ERC20CappedPreset} from "../src/erc20/presets/ERC20Capped.sol";
import {ERC20CappedMintLimited, ERC20CappedMintLimitedPreset} from "../src/erc20/presets/ERC20CappedMintLimited.sol";
import {ERC20Fixed, ERC20FixedPreset} from "../src/erc20/presets/ERC20Fixed.sol";
import {ERC20Mintable, ERC20MintablePreset} from "../src/erc20/presets/ERC20Mintable.sol";

contract TokenFactoryTest is Test {
    TokenFactoryImpl public factory;

    address public erc20CappedPreset;
    address public erc20CappedMintLimitedPreset;
    address public erc20FixedPreset;
    address public erc20MintablePreset;

    address public owner;
    address public forge1;
    address public forge2;
    address public user1;
    address public user2;

    function setUp() public {
        owner = makeAddr("owner");
        forge1 = makeAddr("forge1");
        forge2 = makeAddr("forge2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        erc20CappedPreset = address(new ERC20CappedPreset());
        erc20CappedMintLimitedPreset = address(new ERC20CappedMintLimitedPreset());
        erc20FixedPreset = address(new ERC20FixedPreset());
        erc20MintablePreset = address(new ERC20MintablePreset());

        address[] memory erc20Impls = new address[](4);
        erc20Impls[0] = erc20CappedPreset;
        erc20Impls[1] = erc20CappedMintLimitedPreset;
        erc20Impls[2] = erc20FixedPreset;
        erc20Impls[3] = erc20MintablePreset;
        address[] memory emptyAddrs;

        address factoryImpl = address(new TokenFactoryImpl());
        address proxy = address(
            new ERC1967Proxy(
                factoryImpl,
                abi.encodeCall(TokenFactoryImpl.initialize, (owner, emptyAddrs, erc20Impls, emptyAddrs, emptyAddrs))
            )
        );
        factory = TokenFactoryImpl(proxy);
    }

    function test_create_erc20_capped() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Token",
            "CAPPED",
            18,
            abi.encode(1000000 * 10 ** 18), // 1M cap
            erc20CappedPreset
        );
        ERC20Capped token = ERC20Capped(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Capped Token");
        assertEq(token.symbol(), "CAPPED");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 0);
        assertEq(token.cap(), 1000000 * 10 ** 18);
    }

    function test_create_erc20_capped_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Mint Limited Token",
            "CAPPEDLIM",
            18,
            abi.encode(1000000 * 10 ** 18, 86400, 0, 100000 * 10 ** 18), // 1M cap, 1 day duration, 0 offset, 100k limit
            erc20CappedMintLimitedPreset
        );
        ERC20CappedMintLimited token = ERC20CappedMintLimited(tokenAddress);

        vm.warp(block.timestamp + 86400);
        assertEq(token.owner(), owner);
        assertEq(token.name(), "Capped Mint Limited Token");
        assertEq(token.symbol(), "CAPPEDLIM");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 0);
        assertEq(token.cap(), 1000000 * 10 ** 18);
        assertEq(token.availableMintCapacity(), 100000 * 10 ** 18);
    }

    function test_create_erc20_fixed() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Fixed Token",
            "FIXED",
            18,
            abi.encode(1000000 * 10 ** 18, user1), // 1M initial supply to user1
            erc20FixedPreset
        );
        ERC20Fixed token = ERC20Fixed(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Fixed Token");
        assertEq(token.symbol(), "FIXED");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 1000000 * 10 ** 18);
        assertEq(token.balanceOf(user1), 1000000 * 10 ** 18);
    }

    function test_create_erc20_mintable() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Mintable Token",
            "MINT",
            18,
            abi.encode(), // no extra params
            erc20MintablePreset
        );
        ERC20Mintable token = ERC20Mintable(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Mintable Token");
        assertEq(token.symbol(), "MINT");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 0);
    }
}
