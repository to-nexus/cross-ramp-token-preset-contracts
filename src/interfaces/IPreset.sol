// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

interface IPreset is IERC165 {
    function code() external pure returns (bytes memory);
    function deployCode(bytes memory initialData) external pure returns (bytes memory);
}
