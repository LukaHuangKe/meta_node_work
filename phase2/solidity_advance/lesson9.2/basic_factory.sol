// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

// 简单的代币合约
contract SimpleToken {
    string public name;
    string public symbol;
    address public creator;
    uint256 public totalSupply;

    mapping(address => uint256) public balances;

    constructor(string memory _name, string memory _symbol, uint256 _supply){
        name = _name;
        symbol = _symbol;
        totalSupply = _supply;
        balances[msg.sender] = _supply;
    }

    function transfer(address to, uint256 amount) public {
        require(amount > 0, "Invalid amount");
        require(to != address(0), "Invalid address");

        balances[msg.sender] -= amount;
        balances[to] += amount;
    }
}

contract TokenFactory {
    // 记录所有创建的合约
    SimpleToken[] public tokens;
    // 记录每个用户创建的代币
    mapping(address => address[]) public userTokens;

    // 事件：记录代币创建
    event TokenCreated(
        address indexed tokenAddress,
        string name,
        string symbol,
        address indexed creator
    );

    function createToken(string memory name, string memory symbol, uint256 initialSupply) public returns (address) {
        SimpleToken newToken = new SimpleToken(name, symbol, initialSupply);
        tokens.push(newToken);
        // address(newToken)只是类型转换，不涉及存储读写或合约创建，几乎不消耗 Gas，大约 3-5 Gas
        userTokens[msg.sender].push(address(newToken));

        emit TokenCreated(address(newToken), name, symbol, msg.sender);
        return address(newToken);
    }

    function getTokenCount() public view returns (uint256) {
        return tokens.length;
    }

    function getUserTokens() public view returns (address[] memory) {
        return userTokens[msg.sender];
    }
}