// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RoleManagement {
    // 定义角色枚举
    enum Role { None, User, Admin, Owner }
    
    // 存储用户角色
    mapping(address => Role) public roles;
    
    address public owner;

    event RoleAssigned(address indexed user, Role newRole);
    event RoleRevoked(address indexed user);
    
    constructor() {
        owner = msg.sender;
        roles[msg.sender] = Role.Owner;
        emit RoleAssigned(msg.sender, Role.Owner);
    }
    
    // 定义modifier
    modifier onlyOwner() {
        // 检查是否为Owner
        require(
            roles[msg.sender] == Role.Owner,
            "Not authorized"
        );
        _;  // 函数体会插入到这里
    }
    
    modifier onlyAdmin() {
        // 检查是否为Admin或Owner
        require(
            roles[msg.sender] == Role.Owner || roles[msg.sender] == Role.Admin,
            "Not authorized"
        );
        _;
    }
    
    // 实现功能函数
    function addAdmin(address user) public onlyOwner {
        require(user != address(0), "Invalid user");
        require(roles[user] != Role.Owner, "User already Owner, can't change");
        roles[user] = Role.Admin;
        emit RoleAssigned(user, Role.Admin);
    }
    
    function addUser(address user) public onlyAdmin {
        // Admin添加User
        require(user != address(0), "Invalid user");
        require(roles[user] == Role.None, "User already exists, can't change");
        roles[user] = Role.User;
        emit RoleAssigned(user, Role.User);
    }

    function revokeRole(address user) public onlyOwner {
        require(user != owner, "Can't revoked owner role");
        delete roles[user];
        emit RoleRevoked(user);
    }
    
    function getRole(address user) public view returns (Role) {
        // 查询角色
        return roles[user];
    }
}