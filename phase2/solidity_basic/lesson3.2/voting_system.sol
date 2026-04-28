// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/** 
 * 创建一个提案投票系统：

定义Proposal结构体（包含voters的mapping）
支持创建提案
支持投票（每人只能投一次）
查询提案信息
获取获胜提案
 */
contract VotingSystem {
    struct Proposal {
        string description;
        uint256 voteCount;
        uint256 deadline;
        bool executed;
        mapping(address => bool) voters;
    }

    mapping(uint => Proposal) private proposals;
    uint256 public proposalCount;
    
    event ProposalCreated(uint256 indexed proposalId, string description);
    event Voted(uint256 indexed proposalId, address indexed voter);

    function createProposal(
        string memory description,
        uint256 duration
    ) public returns (uint256) {
        require(duration > 0, "Duration must be greater than 0");
        uint256 proposalId = proposalCount++;
        
        // 先获取 storage 引用，然后逐个赋值
        Proposal storage proposal = proposals[proposalId];
        proposal.description = description;
        proposal.voteCount = 0;
        proposal.deadline = block.timestamp + duration;
        proposal.executed = false;
        
        emit ProposalCreated(proposalId, description);
        return proposalId;
    }

    function vote(uint256 proposalId) public {
        require(proposalId < proposalCount, "Invalid proposal ID");

        Proposal storage p = proposals[proposalId];
        require(!p.voters[msg.sender], "Voter has already voted");
        require(block.timestamp < p.deadline, "Voting period has ended");

        p.voters[msg.sender] = true;
        p.voteCount++;
        emit Voted(proposalId, msg.sender);
    }

    function hasVoted(
        uint256 proposalId,
        address voter
    ) public view returns (bool) {
        require(proposalId < proposalCount, "Invalid proposal ID");
        return proposals[proposalId].voters[voter];
    }

    // mapping不能出现在返回值里
    function getProposalInfo(uint256 proposalId) public view returns (
        string memory description,
        uint256 voteCount,
        uint256 deadline,
        bool executed
    ){
        require(proposalId < proposalCount, "Invalid proposal ID");
        Proposal storage p = proposals[proposalId];
        return (
            p.description,
            p.voteCount,
            p.deadline,
            p.executed
        );
    }

    function getWinningProposal() public view returns (uint256 winningProposalId) {
        uint256 maxVoteCount = 0;
        for(uint256 i = 0; i < proposalCount; i++){
            if(proposals[i].voteCount > maxVoteCount){
                maxVoteCount = proposals[i].voteCount;
                winningProposalId = i;
            }
        }
        return winningProposalId;
    }
}