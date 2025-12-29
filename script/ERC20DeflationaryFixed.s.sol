// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std-1.10.0/src/Script.sol";

import {TokenFactoryImpl} from "../src/TokenFactoryImpl.sol";
import {ERC20DeflationaryFixedPreset} from "../src/erc20/presets/ERC20DeflationaryFixed.sol";
import {TokenType} from "../src/interfaces/ITokenFactory.sol";

contract ERC20DeflationaryFixedScript is Script {
    function deployERC20DeflationaryFixedPreset() external {
        vm.startBroadcast();
        address erc20DeflationaryFixedCode = address(new ERC20DeflationaryFixedPreset());
        vm.stopBroadcast();
        console.log("ERC20DeflationaryFixedPreset deployed to:", erc20DeflationaryFixedCode);
    }

    function registerERC20DeflationaryFixedPreset(address tokenFactory, address erc20DeflationaryFixedCode) external {
        address[] memory presets = new address[](1);
        presets[0] = erc20DeflationaryFixedCode;

        vm.startBroadcast();
        TokenFactoryImpl(tokenFactory).setPresets(TokenType.ERC20, presets, true);
        vm.stopBroadcast();
    }
}
