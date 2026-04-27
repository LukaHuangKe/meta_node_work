// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TypeConversionPractice {
    // 任务1：安全的uint256转uint8
    function safeConvertToUint8(uint256 value) public pure returns (uint8) {
        // 添加范围检查
        require(value <= type(uint8).max, 'Value is out of range');
        return uint8(value);
    }

    // 任务2：字符串比较
    function compareStrings(string memory a, string memory b) 
    public pure returns (bool) 
    {
        // 实现字符串比较
        // 提示：使用keccak256哈希函数
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }

    // 任务3：零地址检查
    function isZeroAddress(address addr) public pure returns (bool) {
        // 检查是否为零地址
        return addr == address(0);
        // address(0) = 0x0000000000000000000000000000000000000000
    }
}