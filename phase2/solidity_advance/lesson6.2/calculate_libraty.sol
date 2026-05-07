// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

// 定义一个数学运算库
library MathOperations {
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        return a + b;
    }
    
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "Subtraction underflow");
        return a - b;
    }
    
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) return 0;
        uint256 c = a * b;
        require(c / a == b, "Multiplication overflow");
        return c;
    }
}

// 使用库的合约
contract Calculator {
    using MathOperations for uint256;
    function calculate(uint256 x, uint256 y) public pure returns (uint256) {
        return MathOperations.add(x, y);
    }

    function testZero() public pure {
        uint256 result = uint256(0).add(0);
        assert(result == 0);
    }
}