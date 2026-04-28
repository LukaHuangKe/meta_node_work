// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/** 
 * 创建一个完整的用户管理系统，实现以下功能：
用户注册（包含name、email）
更新个人资料
存款功能（payable）
查询用户信息
获取所有用户列表
分批查询用户
限制最多1000个用户
 */
contract UserManagementSystem {
    // :定义User结构体
    struct User {
        // name, email, balance, registeredAt, exists
        string name;
        string email;
        uint balance;
        uint registeredAt;
        bool exists;
    }
    
    // 定义数据存储
    mapping(address => User) public users;
    address[] public userAddresses;
    uint256 public userCount;
    uint256 public constant MAX_USERS = 1000;

    event UserRegistered(address indexed user, string name);
    event UserUpdated(address indexed user);
    event Deposit(address indexed user, uint amount);

    
    // 实现注册功能
    function register(string memory name, string memory email) public {
        // 检查是否已注册
        require(!users[msg.sender].exists, "User already registered");
        // 检查是否达到上限
        require(userCount <= MAX_USERS, "User count reach limit");
        // 创建用户
        users[msg.sender] = User({
            name: name,
            email: email,
            balance: 0,
            registeredAt: block.timestamp,
            exists: true
        });
        // 添加到列表
        userAddresses.push(msg.sender);
        // 更新计数
        userCount++;
        emit UserRegistered(msg.sender, name);
    }
    
    // 更新个人信息
    function updateProfile(string memory name, string memory email) public {
        // 检查是否已注册
        require(users[msg.sender].exists, "User not registered yet");
        // 更新个人信息
        users[msg.sender].name = name;
        users[msg.sender].email = email;
        emit UserUpdated(msg.sender);
    }

    // 存款
    function deposit(uint amount) public payable{
        // 检查是否已注册
        require(users[msg.sender].exists, "User not registered yet");
        // 更新余额
        users[msg.sender].balance += amount;
        emit Deposit(msg.sender, amount);
    }

    function getUserInfo(address user) public view returns (User memory) {
        require(users[msg.sender].exists, "User not registered yet");
        return users[user];
    }

    function getAllUserInfo() public view returns (User[] memory) {
        User[] memory res = new User[](userAddresses.length);
        for(uint i = 0; i < userAddresses.length; i++){
            res[i] = users[userAddresses[i]];
        }
        return res;
    }

    function getUsersByRange(uint start, uint end) public view returns (User[] memory) {
        require(start < end, "Invalid range");
        require(end <= userAddresses.length, "out of length");

        uint len = end - start;
        User[] memory res = new User[](len);
        for(uint i = 0; i < len; i++){
            res[i] = users[userAddresses[start + i]];
        }
                   
        return res;
    }
}