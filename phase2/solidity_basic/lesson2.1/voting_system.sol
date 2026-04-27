// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
使用enum定义投票选项：Yes, No, Abstain
使用mapping记录每个地址的投票
使用uint统计每个选项的票数
实现投票和查询功能
 */
contract VotingSystem {
    enum VoteOption {
        Yes,
        No,
        Abstain
    }

    mapping(address => VoteOption) public voteMap;
    mapping(address => bool) public hasVotedMap;
    uint256 public yesVotes;
    uint256 public noVotes;
    uint256 public abstainVotes;
    
    // 还需要定义一个event
    event Voted(address indexed voter, VoteOption vote);


    function vote(VoteOption _vote) public {
        // 实现投票逻辑
        // - 检查是否已投票
        require(!hasVotedMap[msg.sender], "You have voted");
        // - 记录投票
        voteMap[msg.sender] = _vote;
        hasVotedMap[msg.sender] = true;
        
        // - 更新计数
        if (_vote == VoteOption.Yes) {
            yesVotes++;
        } else if (_vote == VoteOption.No) {
            noVotes++;
        } else if (_vote == VoteOption.Abstain) {
            abstainVotes++;
        }

        emit Voted(msg.sender, _vote);
    }
    // 查询函数
    function getResults() public view returns (uint, uint, uint) {
        return (yesVotes, noVotes, abstainVotes);
    }

    function getMyVote() public view returns (VoteOption) {
        require(hasVotedMap[msg.sender], "You haven't voted");
        return voteMap[msg.sender];
    }
}