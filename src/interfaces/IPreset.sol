// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IPreset {
    function code() external pure returns (bytes memory);
    function deployCode(bytes memory initialData) external pure returns (bytes memory);
    function supportInterface(bytes4 interfaceId) external view returns (bool);
}
