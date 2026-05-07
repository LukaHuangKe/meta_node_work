// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

library AdvancedMath {
    // 平方根（Newton-Raphson方法）
    function sqrt(uint256 x) internal pure returns (uint256) {
        if (x == 0) return 0;
        
        uint256 z = (x + 1) / 2;
        uint256 y = x;
        
        while (z < y) {
            y = z;
            z = (x / z + z) / 2;
        }
        
        return y;
    }
    
    // 实现最大公约数
    function gcd(uint256 a, uint256 b) internal pure returns (uint256) {
        // 提示：使用辗转相除法
        while (b != 0) {
            uint256 temp = b;
            b = a % b;
            a = temp;
        }
        return a;
    }
    
    // 幂运算（快速幂算法）
    function power(uint256 base, uint256 exponent) 
        internal pure returns (uint256) 
    {
        if (exponent == 0) return 1;
        
        uint256 result = 1;
        uint256 currentBase = base;
        
        while (exponent > 0) {
            if (exponent % 2 == 1) {
                result *= currentBase;
            }
            currentBase *= currentBase;
            exponent /= 2;
        }
        
        return result;
    }
}