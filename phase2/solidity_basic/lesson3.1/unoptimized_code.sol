// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract UnoptimizedCode {
    uint[] public data;
    
    function process(uint[] memory values) public {
        for(uint i = 0; i < values.length; i++) {
            if(values[i] > 10) {
                data.push(values[i]);
            }
        }
    }

    function processOptimize(uint[] calldata values) public {
        uint len = values.length;

        uint count = 0;
        uint[] memory temp = new uint[](len);

        for(uint i = 0; i < len; i++) {
            if(values[i] > 10) {
                temp[count] = values[i];
                count++;
            }
        }


        for(uint i = 0; i < count; i++) {
            data.push(temp[i]);
        }
    }
}