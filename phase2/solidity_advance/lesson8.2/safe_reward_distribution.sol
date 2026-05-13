// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract SafeRewardDistribution {
    mapping(address => uint256) public rewards;
    mapping(address => uint256) public claimDeadline;

    uint256 public constant CLAIM_PERIOD = 30 days;

    function setReward(address user, uint256 amount) external {
        require(amount > 0, "Invalid reward amount");
        require(user != address(0), "Invalid user address");
        rewards[user] = amount;
        claimDeadline[user] = block.timestamp + CLAIM_PERIOD;
    }

    function claimReward(address user) external {
        uint256 amount = rewards[user];
        // check
        require(amount > 0, "No reward to claim");
        require(block.timestamp >= claimDeadline[user], "Claim deadline not reached");

        // effect
        rewards[user] = 0;

        // 转账 interaction
        (bool success,) = user.call{value:amount}("");
        require(success, "Transfer failed");
    }
}