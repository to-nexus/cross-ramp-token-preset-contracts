// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Script, console} from "forge-std-1.10.0/src/Script.sol";

import {TokenFactoryImpl} from "../src/TokenFactoryImpl.sol";
import {ERC20Capped, ERC20CappedPreset} from "../src/erc20/presets/ERC20Capped.sol";
import {
    ERC20CappedInitialSupply, ERC20CappedInitialSupplyPreset
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
import {ERC20Fixed, ERC20FixedPreset} from "../src/erc20/presets/ERC20Fixed.sol";
import {ERC20Mintable, ERC20MintablePreset} from "../src/erc20/presets/ERC20Mintable.sol";

import {ERC20MintLimited, ERC20MintLimitedPreset} from "../src/erc20/presets/ERC20MintLimited.sol";
import {ERC20MultiMintLimited, ERC20MultiMintLimitedPreset} from "../src/erc20/presets/ERC20MultiMintLimited.sol";

contract DeployScript is Script {
    address private constant OWNER = 0x26e8D58B2f3279D45f98D736942E4fcC9700581f;

    // deployRampLogics
    function deployRampLogics() external {
        vm.startBroadcast();
        address tokenFactoryImpl = address(new TokenFactoryImpl());

        address erc20CappedPreset = address(new ERC20CappedPreset());
        address erc20CappedInitialSupplyPreset = address(new ERC20CappedInitialSupplyPreset());
        address erc20CappedInitialSupplyMintLimitedPreset = address(new ERC20CappedInitialSupplyMintLimitedPreset());
        address erc20CappedInitialSupplyMultiMintLimitedPreset =
            address(new ERC20CappedInitialSupplyMultiMintLimitedPreset());
        address erc20CappedMintLimitedPreset = address(new ERC20CappedMintLimitedPreset());
        address erc20CappedMultiMintLimitedPreset = address(new ERC20CappedMultiMintLimitedPreset());
        address erc20FixedPreset = address(new ERC20FixedPreset());
        address erc20MintablePreset = address(new ERC20MintablePreset());
        address erc20MintLimitedPreset = address(new ERC20MintLimitedPreset());
        address erc20MultiMintLimitedPreset = address(new ERC20MultiMintLimitedPreset());
        vm.stopBroadcast();
        console.log("address private tokenFactoryImpl =", tokenFactoryImpl, ";");
        console.log("address private erc20CappedPreset =", erc20CappedPreset, ";");
        console.log("address private erc20CappedInitialSupplyPreset =", erc20CappedInitialSupplyPreset, ";");
        console.log(
            "address private erc20CappedInitialSupplyMintLimitedPreset =",
            erc20CappedInitialSupplyMintLimitedPreset,
            ";"
        );
        console.log(
            "address private erc20CappedInitialSupplyMultiMintLimitedPreset =",
            erc20CappedInitialSupplyMultiMintLimitedPreset,
            ";"
        );
        console.log("address private erc20CappedMintLimitedPreset =", erc20CappedMintLimitedPreset, ";");
        console.log("address private erc20CappedMultiMintLimitedPreset =", erc20CappedMultiMintLimitedPreset, ";");
        console.log("address private erc20FixedPreset =", erc20FixedPreset, ";");
        console.log("address private erc20MintablePreset =", erc20MintablePreset, ";");
        console.log("address private erc20MintLimitedPreset =", erc20MintLimitedPreset, ";");
        console.log("address private erc20MultiMintLimitedPreset =", erc20MultiMintLimitedPreset, ";");
    }

    function initializeTokenFactory(address owner, address tokenFactoryImpl, address[] memory erc20Impls) external {
        address[] memory empty;
        vm.startBroadcast();
        address tokenFactory = address(
            new ERC1967Proxy(
                tokenFactoryImpl, abi.encodeCall(TokenFactoryImpl.initialize, (owner, empty, erc20Impls, empty, empty))
            )
        );
        vm.stopBroadcast();
        console.log("address private tokenFactory =", tokenFactory, ";");
    }
}
