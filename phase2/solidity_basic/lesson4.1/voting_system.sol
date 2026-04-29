// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VotingSystem {
    struct Proposal {
        string description;
        uint voteCount;
        uint deadline;
        bool exists;
    }
    
    address public owner;
    uint public proposalCount;
    
    mapping(uint => Proposal) public proposals;
    mapping(uint => mapping(address => bool)) public hasVoted;
    
    event ProposalCreated(uint indexed proposalId, string description, uint durationDays);
    event Voted(uint indexed proposalId, address indexed voter);

    constructor() {
        owner = msg.sender;
    }
    
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;  // 下划线表示函数体的位置
    }

    // 实现创建提案
    function createProposal(string memory description, uint durationDays) 
        public onlyOwner
    {
        // 检查权限
        // 验证参数
        require(bytes(description).length > 0, "Empty description");
        require(durationDays >= 1 && durationDays <= 30, "Invalid duration");

        // 创建提案
        uint proposalId = proposalCount++;
        uint deadLine = block.timestamp + (durationDays * 1 days);
        proposals[proposalId] = Proposal({
            description: description,
            voteCount: 0,
            deadline: deadLine,
            exists: true
        });
        emit ProposalCreated(proposalId, description, durationDays);
    }
    
    // 实现投票
    function vote(uint proposalId) public {
        // 检查提案存在
        require(proposals[proposalId].exists, "Proposal does not exist");
        // 检查是否已投票
        require(!hasVoted[proposalId][msg.sender], "Already voted");
        // 检查是否已截止
        require(block.timestamp < proposals[proposalId].deadline, "Proposal deadline has passed");
        
        // 执行投票
        hasVoted[proposalId][msg.sender] = true;
        proposals[proposalId].voteCount++;
        emit Voted(proposalId, msg.sender);
    }
    
    // 获取获胜提案
    function getWinner() public view returns (uint) {
        // 检查是否有提案
        require(proposalCount > 0, "No proposals created");

        // 遍历所有提案
        uint maxVoteCount = 0;
        uint winnerProposalId = 0;
        // 找出票数最多的
        for(uint i = 0; i < proposalCount; i++){
            if(proposals[i].voteCount > maxVoteCount){
                maxVoteCount = proposals[i].voteCount;
                winnerProposalId = i;
            }
        }
        return winnerProposalId;
    }
}