// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

// 使用CEI模式修复
contract SecurityVaultCEI {
    mapping(address => uint256) public balances;

    function deposit() external payable{
        balances[msg.sender] += msg.value;
    }

    function withdraw() external {
        //check
        uint256 amount = balances[msg.sender];
        require(amount > 0, "Insufficient balance");

        //effect
        balances[msg.sender] = 0;

        //interaction
        (bool success,) = msg.sender.call{value:amount}("");
        require(success, "Transfer failed");
    }
}

// 使用锁修复
contract SecurityVauleLock {
    mapping(address => uint256) public balances;
    bool private locked;

    modifier noReentrant() {
        require(!locked, "No reentrancy");
        locked = true;
        _;
        locked = false;
    }

    function deposit() external payable{
        balances[msg.sender] += msg.value;
    }

    function withdraw() external noReentrant {
        //check
        uint256 amount = balances[msg.sender];
        require(amount > 0, "Insufficient balance");

        //effect
        balances[msg.sender] = 0;

        //interaction
        (bool success,) = msg.sender.call{value:amount}("");
        require(success, "Transfer failed");
    }
}