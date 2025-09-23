// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.4.0/proxy/ERC1967/ERC1967Proxy.sol";
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

import {ERC721AutoIncr, ERC721AutoIncrPreset} from "../src/erc721/presets/ERC721AutoIncr.sol";
import {ERC721Simple, ERC721SimplePreset} from "../src/erc721/presets/ERC721Simple.sol";

import {ERC1155Simple, ERC1155SimplePreset} from "../src/erc1155/presets/ERC1155Simple.sol";

contract DeployScript is Script {
    function deployLogics() external {
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

        address erc721AutoIncr = address(new ERC721AutoIncrPreset());
        address erc721Simple = address(new ERC721SimplePreset());

        address erc1155Simple = address(new ERC1155SimplePreset());

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

        console.log("address private erc721AutoIncr =", erc721AutoIncr, ";");
        console.log("address private erc721Simple =", erc721Simple, ";");

        console.log("address private erc1155Simple =", erc1155Simple, ";");
    }

    function initializeTokenFactory(
        address tokenFactoryImpl,
        address owner,
        address[] memory deployers,
        address[] memory erc20Impls,
        address[] memory erc721Impls,
        address[] memory erc1155Impls
    ) external {
        vm.startBroadcast();
        address tokenFactory = address(
            new ERC1967Proxy(
                tokenFactoryImpl,
                abi.encodeCall(TokenFactoryImpl.initialize, (owner, deployers, erc20Impls, erc721Impls, erc1155Impls))
            )
        );
        vm.stopBroadcast();
        console.log("address private tokenFactory =", tokenFactory, ";");
    }
}
