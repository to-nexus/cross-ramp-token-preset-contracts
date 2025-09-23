// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {ERC20Base} from "../ERC20Base.sol";
import {ERC20Capable} from "../extensions/ERC20Capable.sol";
import {ERC20InitialSupply} from "../extensions/ERC20InitialSupply.sol";

contract ERC20CappedInitialSupply is ERC20Base, ERC20Capable, ERC20InitialSupply {
    constructor(
        address owner,
        address[] memory forges,
        string memory name,
        string memory symbol,
        uint8 decimals,
        bytes[2] memory extensionData
    )
        ERC20Base(owner, forges, name, symbol, decimals)
        ERC20Capable(extensionData[0])
        ERC20InitialSupply(extensionData[1])
    {}

    function _update(address from, address to, uint256 value) internal override(ERC20Base, ERC20Capable) {
        super._update(from, to, value);
    }
}

import {ERC20BasePreset} from "../ERC20Base.sol";
import {Test, console} from "forge-std-1.10.0/src/Test.sol";

contract ERC20CappedInitialSupplyPreset is ERC20BasePreset, Test {
    function code() public pure override returns (bytes memory) {
        return type(ERC20CappedInitialSupply).creationCode;
    }

    function deployCode(bytes memory initialData) external pure override returns (bytes memory) {
        (
            address owner,
            address[] memory forges,
            string memory name,
            string memory symbol,
            uint8 decimals,
            bytes memory data
        ) = abi.decode(initialData, (address, address[], string, string, uint8, bytes));
        return abi.encodePacked(code(), abi.encode(owner, forges, name, symbol, decimals, abi.decode(data, (bytes[2]))));
    }
}
