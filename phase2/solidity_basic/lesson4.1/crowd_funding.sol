// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Crowdfunding {
    enum State { Fundraising, Success, Failed, PaidOut }
    
    State public currentState = State.Fundraising;
    
    address public creator;
    uint public goal;
    uint public deadline;
    uint public totalFunded;
    
    mapping(address => uint) public contributions;
    
    constructor(uint _goal, uint _duration) {
        creator = msg.sender;
        goal = _goal;
        deadline = block.timestamp + _duration;
    }
    
    modifier inState(State expected) {
        require(currentState == expected, "Invalid state");
        _;
    }
    
    // 贡献资金（只在Fundraising状态）
    function contribute() 
        public payable inState(State.Fundraising) 
    {
        require(block.timestamp < deadline, "Campaign ended");
        require(msg.value > 0, "Must send ETH");
        
        contributions[msg.sender] += msg.value;
        totalFunded += msg.value;
    }
    
    // 检查目标（deadline后调用）
    function checkGoalReached() public inState(State.Fundraising) {
        require(block.timestamp >= deadline, "Campaign not ended yet");
        
        if (totalFunded >= goal) {
            currentState = State.Success;
        } else {
            currentState = State.Failed;
        }
    }
    
    // 创建者提取资金（成功后）
    function payout() public inState(State.Success) {
        require(msg.sender == creator, "Only creator can payout");
        
        currentState = State.PaidOut;
        payable(creator).transfer(address(this).balance);
    }
    
    // 退款（失败后）
    function refund() public inState(State.Failed) {
        uint amount = contributions[msg.sender];
        require(amount > 0, "No contribution to refund");
        
        contributions[msg.sender] = 0;
        payable(msg.sender).transfer(amount);
    }
}