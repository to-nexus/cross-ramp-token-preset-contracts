// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {TokenFactoryImpl} from "../src/TokenFactoryImpl.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Test, console} from "forge-std-1.10.0/src//Test.sol";

import {ERC20Capped, ERC20CappedPreset} from "../src/erc20/presets/ERC20Capped.sol";
import {
    ERC20CappedInitialSupply,
    ERC20CappedInitialSupplyPreset
} from "../src/erc20/presets/ERC20CappedInitialSupply.sol";
import {
    ERC20CappedInitialSupplyMintLimited,
    ERC20CappedInitialSupplyMintLimitedPreset
} from "../src/erc20/presets/ERC20CappedInitialSupplyMintLimited.sol";
import {
    ERC20CappedInitialSupplyMultiMintLimited,
    ERC20CappedInitialSupplyMultiMintLimitedPreset
} from "../src/erc20/presets/ERC20CappedInitialSupplyMultiMintLimited.sol";
import {ERC20CappedMintLimited, ERC20CappedMintLimitedPreset} from "../src/erc20/presets/ERC20CappedMintLimited.sol";
import {
    ERC20CappedMultiMintLimited,
    ERC20CappedMultiMintLimitedPreset
} from "../src/erc20/presets/ERC20CappedMultiMintLimited.sol";
import {ERC20DeflationaryFixed, ERC20DeflationaryFixedPreset} from "../src/erc20/presets/ERC20DeflationaryFixed.sol";
import {ERC20Fixed, ERC20FixedPreset} from "../src/erc20/presets/ERC20Fixed.sol";
import {ERC20Mintable, ERC20MintablePreset} from "../src/erc20/presets/ERC20Mintable.sol";

import {ERC20MintLimited, ERC20MintLimitedPreset} from "../src/erc20/presets/ERC20MintLimited.sol";
import {ERC20MultiMintLimited, ERC20MultiMintLimitedPreset} from "../src/erc20/presets/ERC20MultiMintLimited.sol";

import {ERC721AutoIncr, ERC721AutoIncrPreset} from "../src/erc721/presets/ERC721AutoIncr.sol";
import {ERC721Simple, ERC721SimplePreset} from "../src/erc721/presets/ERC721Simple.sol";

import {ERC1155Simple, ERC1155SimplePreset} from "../src/erc1155/presets/ERC1155Simple.sol";

contract TokenFactoryTest is Test {
    TokenFactoryImpl public factory;

    address public erc20CappedPreset;
    address public erc20CappedInitialSupplyPreset;
    address public erc20CappedInitialSupplyMintLimitedPreset;
    address public erc20CappedInitialSupplyMultiMintLimitedPreset;
    address public erc20CappedMintLimitedPreset;
    address public erc20CappedMultiMintLimitedPreset;
    address public erc20DeflationaryFixedPreset;
    address public erc20FixedPreset;
    address public erc20MintablePreset;
    address public erc20MintLimitedPreset;
    address public erc20MultiMintLimitedPreset;

    address public erc721SimplePreset;
    address public erc721AutoIncrPreset;

    address public erc1155SimplePreset;

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
        erc20CappedInitialSupplyPreset = address(new ERC20CappedInitialSupplyPreset());
        erc20CappedInitialSupplyMintLimitedPreset = address(new ERC20CappedInitialSupplyMintLimitedPreset());
        erc20CappedInitialSupplyMultiMintLimitedPreset = address(new ERC20CappedInitialSupplyMultiMintLimitedPreset());
        erc20CappedMintLimitedPreset = address(new ERC20CappedMintLimitedPreset());
        erc20CappedMultiMintLimitedPreset = address(new ERC20CappedMultiMintLimitedPreset());
        erc20DeflationaryFixedPreset = address(new ERC20DeflationaryFixedPreset());
        erc20FixedPreset = address(new ERC20FixedPreset());
        erc20MintablePreset = address(new ERC20MintablePreset());
        erc20MintLimitedPreset = address(new ERC20MintLimitedPreset());
        erc20MultiMintLimitedPreset = address(new ERC20MultiMintLimitedPreset());

        erc721SimplePreset = address(new ERC721SimplePreset());
        erc721AutoIncrPreset = address(new ERC721AutoIncrPreset());

        erc1155SimplePreset = address(new ERC1155SimplePreset());

        address[] memory erc20Impls = new address[](11);
        erc20Impls[0] = erc20CappedPreset;
        erc20Impls[1] = erc20CappedInitialSupplyPreset;
        erc20Impls[2] = erc20CappedInitialSupplyMintLimitedPreset;
        erc20Impls[3] = erc20CappedInitialSupplyMultiMintLimitedPreset;
        erc20Impls[4] = erc20CappedMintLimitedPreset;
        erc20Impls[5] = erc20CappedMultiMintLimitedPreset;
        erc20Impls[6] = erc20DeflationaryFixedPreset;
        erc20Impls[7] = erc20FixedPreset;
        erc20Impls[8] = erc20MintablePreset;
        erc20Impls[9] = erc20MintLimitedPreset;
        erc20Impls[10] = erc20MultiMintLimitedPreset;

        address[] memory erc721Impls = new address[](2);
        erc721Impls[0] = erc721SimplePreset;
        erc721Impls[1] = erc721AutoIncrPreset;

        address[] memory erc1155Impls = new address[](1);
        erc1155Impls[0] = erc1155SimplePreset;

        address[] memory emptyAddrs;
        address factoryImpl = address(new TokenFactoryImpl());
        address proxy = address(
            new ERC1967Proxy(
                factoryImpl,
                abi.encodeCall(TokenFactoryImpl.initialize, (owner, emptyAddrs, erc20Impls, erc721Impls, erc1155Impls))
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

    function test_create_erc20_capped_initial_supply() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        uint256 cap = 1000000 * 10 ** 18; // 1M cap
        uint256 initialSupply = 500000 * 10 ** 18; // 500
        address initialRecipient = makeAddr("initialRecipient");

        vm.prank(owner);
        bytes[2] memory extensionData = [
            abi.encode(cap), // 1M cap
            abi.encode(initialSupply, initialRecipient) // 500k initial supply to initialRecipient
        ];
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Token Initial Supply",
            "CIP",
            18,
            abi.encode(extensionData),
            erc20CappedInitialSupplyPreset
        );
        ERC20Capped token = ERC20Capped(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Capped Token Initial Supply");
        assertEq(token.symbol(), "CIP");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), initialSupply);
        assertEq(token.cap(), cap);
        assertEq(token.balanceOf(initialRecipient), initialSupply);
    }

    function test_create_erc20_capped_initial_supply_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        uint256 cap = 1000000 * 10 ** 18; // 1M cap
        uint256 initialSupply = 500000 * 10 ** 18; // 500
        address initialRecipient = makeAddr("initialRecipient");

        vm.prank(owner);
        bytes[3] memory extensionData = [
            abi.encode(cap), // 1M cap
            abi.encode(initialSupply, initialRecipient), // 500k initial supply to initialRecipient
            abi.encode(86400, 0, 100000 * 10 ** 18) // 1 day duration, 0 offset, 100k limit
        ];
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Initial Supply Mint Limited Token",
            "CSML",
            18,
            abi.encode(extensionData),
            erc20CappedInitialSupplyMintLimitedPreset
        );
        ERC20CappedInitialSupplyMintLimited token = ERC20CappedInitialSupplyMintLimited(tokenAddress);

        vm.warp(block.timestamp + 86400);
        assertEq(token.owner(), owner);
        assertEq(token.name(), "Capped Initial Supply Mint Limited Token");
        assertEq(token.symbol(), "CSML");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), initialSupply);
        assertEq(token.cap(), cap);
        assertEq(token.availableMintCapacity(), 100000 * 10 ** 18);
        assertEq(token.balanceOf(initialRecipient), initialSupply);
    }

    function test_create_erc20_capped_initial_supply_multi_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        uint256 cap = 1000000 * 10 ** 18; // 1M cap
        uint256 initialSupply = 500000 * 10 ** 18; // 500
        address initialRecipient = makeAddr("initialRecipient");
        uint256[] memory durations = new uint256[](3);
        int256[] memory offsetSeconds = new int256[](3);
        uint256[] memory limits = new uint256[](3);

        durations[0] = 3600; // 1 hour
        durations[1] = 86400; // 1 day
        durations[2] = 604800; // 1 week

        offsetSeconds[0] = 0; // No offset
        offsetSeconds[1] = 0; // No offset
        offsetSeconds[2] = 0; // No offset

        limits[0] = 1000 * 10 ** 18; // 1K tokens per hour
        limits[1] = 10000 * 10 ** 18; // 10K tokens per day
        limits[2] = 50000 * 10 ** 18; // 100K tokens per week

        vm.prank(owner);
        bytes[3] memory extensionData = [
            abi.encode(cap), // 1M cap
            abi.encode(initialSupply, initialRecipient), // 500k initial supply to initialRecipient
            abi.encode(durations, offsetSeconds, limits) // durations, offsets, limits
        ];
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Mint Limited Token",
            "CAPPEDLIM",
            18,
            abi.encode(extensionData),
            erc20CappedInitialSupplyMultiMintLimitedPreset
        );
        ERC20CappedInitialSupplyMultiMintLimited token = ERC20CappedInitialSupplyMultiMintLimited(tokenAddress);

        vm.warp(block.timestamp + 86400);
        assertEq(token.owner(), owner);
        assertEq(token.name(), "Capped Mint Limited Token");
        assertEq(token.symbol(), "CAPPEDLIM");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), initialSupply);
        assertEq(token.cap(), 1000000 * 10 ** 18);
        assertEq(token.balanceOf(initialRecipient), initialSupply);
        uint256[] memory maxMintPerPeriods = token.maxMintPerPeriods();
        uint256 length = maxMintPerPeriods.length;
        assertEq(length, 3);
        for (uint256 i = 0; i < length; i++) {
            assertEq(maxMintPerPeriods[i], limits[i]);
        }
        (uint256[] memory _durations, int256[] memory _offsetSeconds) = token.periodConfigs();
        assertEq(_durations.length, durations.length);
        assertEq(_offsetSeconds.length, offsetSeconds.length);
        for (uint256 i = 0; i < durations.length; i++) {
            assertEq(_durations[i], durations[i]);
            assertEq(_offsetSeconds[i], offsetSeconds[i]);
        }
    }

    function test_create_erc20_capped_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        bytes[2] memory extensionData = [
            abi.encode(1000000 * 10 ** 18), // 1M cap
            abi.encode(86400, 0, 100000 * 10 ** 18) // 1 day duration, 0 offset, 100k limit
        ];
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Mint Limited Token",
            "CAPPEDLIM",
            18,
            abi.encode(extensionData),
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

    function test_create_erc20_capped_multi_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        uint256 cap = 1000000 * 10 ** 18; // 1M cap
        uint256[] memory durations = new uint256[](3);
        int256[] memory offsetSeconds = new int256[](3);
        uint256[] memory limits = new uint256[](3);

        durations[0] = 3600; // 1 hour
        durations[1] = 86400; // 1 day
        durations[2] = 604800; // 1 week

        offsetSeconds[0] = 0; // No offset
        offsetSeconds[1] = 0; // No offset
        offsetSeconds[2] = 0; // No offset

        limits[0] = 1000 * 10 ** 18; // 1K tokens per hour
        limits[1] = 10000 * 10 ** 18; // 10K tokens per day
        limits[2] = 50000 * 10 ** 18; // 100K tokens per week

        vm.prank(owner);
        bytes[2] memory extensionData = [
            abi.encode(cap), // 1M cap
            abi.encode(durations, offsetSeconds, limits) // durations, offsets, limits
        ];
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Capped Mint Limited Token",
            "CAPPEDLIM",
            18,
            abi.encode(extensionData),
            erc20CappedMultiMintLimitedPreset
        );
        ERC20CappedMultiMintLimited token = ERC20CappedMultiMintLimited(tokenAddress);

        vm.warp(block.timestamp + 86400);
        assertEq(token.owner(), owner);
        assertEq(token.name(), "Capped Mint Limited Token");
        assertEq(token.symbol(), "CAPPEDLIM");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 0);
        assertEq(token.cap(), 1000000 * 10 ** 18);
        uint256[] memory maxMintPerPeriods = token.maxMintPerPeriods();
        uint256 length = maxMintPerPeriods.length;
        assertEq(length, 3);
        for (uint256 i = 0; i < length; i++) {
            assertEq(maxMintPerPeriods[i], limits[i]);
        }
        (uint256[] memory _durations, int256[] memory _offsetSeconds) = token.periodConfigs();
        assertEq(_durations.length, durations.length);
        assertEq(_offsetSeconds.length, offsetSeconds.length);
        for (uint256 i = 0; i < durations.length; i++) {
            assertEq(_durations[i], durations[i]);
            assertEq(_offsetSeconds[i], offsetSeconds[i]);
        }
    }

    function test_create_erc20_deflationary_fixed() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        uint256 initialSupply = 1000000 * 10 ** 18;

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Deflationary Fixed Token",
            "DEFLAT",
            18,
            abi.encode(initialSupply, user1), // 1M initial supply to user1
            erc20DeflationaryFixedPreset
        );
        ERC20DeflationaryFixed token = ERC20DeflationaryFixed(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Deflationary Fixed Token");
        assertEq(token.symbol(), "DEFLAT");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), initialSupply);
        assertEq(token.balanceOf(user1), initialSupply);

        // Verify minting is not allowed
        vm.prank(forge1);
        vm.expectRevert(ERC20DeflationaryFixed.ERC20DeflationaryFixed__MintingNotAllowed.selector);
        token.mint(user2, 1000);

        // Verify burning works
        uint256 burnAmount = 100000 * 10 ** 18;
        vm.prank(user1);
        token.burnFrom(user1, burnAmount);
        assertEq(token.totalSupply(), initialSupply - burnAmount);
        assertEq(token.balanceOf(user1), initialSupply - burnAmount);
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

    function test_create_erc20_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Mint Limited Token",
            "LIM",
            18,
            abi.encode(86400, 0, 100000 * 10 ** 18), // 1 day duration, 0 offset, 100k limit
            erc20MintLimitedPreset
        );
        ERC20MintLimited token = ERC20MintLimited(tokenAddress);

        vm.warp(block.timestamp + 86400);
        assertEq(token.owner(), owner);
        assertEq(token.name(), "Mint Limited Token");
        assertEq(token.symbol(), "LIM");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 0);
        assertEq(token.maxMintPerPeriod(), 100000 * 10 ** 18);
        (uint256 duration, int256 offsetSeconds) = token.periodConfig();
        assertEq(duration, 86400);
        assertEq(offsetSeconds, 0);
    }

    function test_create_erc20_multi_mint_limited() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        uint256[] memory durations = new uint256[](3);
        int256[] memory offsetSeconds = new int256[](3);
        uint256[] memory limits = new uint256[](3);

        durations[0] = 3600; // 1 hour
        durations[1] = 86400; // 1 day
        durations[2] = 604800; // 1 week

        offsetSeconds[0] = 0; // No offset
        offsetSeconds[1] = 0; // No offset
        offsetSeconds[2] = 0; // No offset

        limits[0] = 1000 * 10 ** 18; // 1K tokens per hour
        limits[1] = 10000 * 10 ** 18; // 10K tokens per day
        limits[2] = 50000 * 10 ** 18; // 100K tokens per week

        vm.prank(owner);
        address tokenAddress = factory.deployERC20(
            owner,
            forges,
            "Mint Limited Token",
            "MLIM",
            18,
            abi.encode(durations, offsetSeconds, limits), // 1M cap, 1 day duration, 0 offset, 100k limit
            erc20MultiMintLimitedPreset
        );
        ERC20MultiMintLimited token = ERC20MultiMintLimited(tokenAddress);

        vm.warp(block.timestamp + 86400);
        assertEq(token.owner(), owner);
        assertEq(token.name(), "Mint Limited Token");
        assertEq(token.symbol(), "MLIM");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), 0);
        uint256[] memory maxMintPerPeriods = token.maxMintPerPeriods();
        uint256 length = maxMintPerPeriods.length;
        assertEq(length, 3);
        for (uint256 i = 0; i < length; i++) {
            assertEq(maxMintPerPeriods[i], limits[i]);
        }
        (uint256[] memory _durations, int256[] memory _offsetSeconds) = token.periodConfigs();
        assertEq(_durations.length, durations.length);
        assertEq(_offsetSeconds.length, offsetSeconds.length);
        for (uint256 i = 0; i < durations.length; i++) {
            assertEq(_durations[i], durations[i]);
            assertEq(_offsetSeconds[i], offsetSeconds[i]);
        }
    }

    function test_create_erc721_simple() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC721(
            owner,
            forges,
            "Simple NFT",
            "SNFT",
            "https://example.com/metadata/",
            abi.encode(), // no extra params
            erc721SimplePreset
        );
        ERC721Simple token = ERC721Simple(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Simple NFT");
        assertEq(token.symbol(), "SNFT");
    }

    function test_create_erc721_auto_incr() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC721(
            owner,
            forges,
            "AutoIncr TokenID NFT",
            "AITN",
            "https://example.com/metadata/",
            abi.encode(), // no extra params
            erc721AutoIncrPreset
        );
        ERC721AutoIncr token = ERC721AutoIncr(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "AutoIncr TokenID NFT");
        assertEq(token.symbol(), "AITN");
    }

    function test_create_erc1155_simple() public {
        address[] memory forges = new address[](2);
        forges[0] = forge1;
        forges[1] = forge2;

        vm.prank(owner);
        address tokenAddress = factory.deployERC1155(
            owner,
            forges,
            "Simple 1155",
            "S1155",
            "https://example.com/metadata/{id}.json",
            abi.encode(), // no extra params
            erc1155SimplePreset
        );
        ERC1155Simple token = ERC1155Simple(tokenAddress);

        assertEq(token.owner(), owner);
        assertEq(token.name(), "Simple 1155");
        assertEq(token.symbol(), "S1155");
    }
}
