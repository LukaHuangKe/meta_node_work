// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

// 模板合约（只部署一次）
contract TokenImplementation {
    string public name;
    string public symbol;
    address public creator;
    uint256 public totalSupply;
    
    mapping(address => uint256) public balances;
    
    /**
     * @notice 初始化函数（替代构造函数）
     * @dev Clone不能使用构造函数，所以用初始化函数
     */
    function initialize(
        string memory _name,
        string memory _symbol,
        uint256 _supply
    ) public {
        require(creator == address(0), "Already initialized");
        name = _name;
        symbol = _symbol;
        creator = msg.sender;
        totalSupply = _supply;
        balances[msg.sender] = _supply;
    }
    
    function transfer(address to, uint256 amount) public {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        balances[msg.sender] -= amount;
        balances[to] += amount;
    }
}

// Clone工厂合约
contract CloneFactory {
    // 模板合约地址
    address public implementation;
    
    // 记录所有克隆的地址
    address[] public clones;
    
    // 记录每个用户创建的克隆
    mapping(address => address[]) public userClones;
    
    event CloneCreated(address indexed cloneAddress, address indexed creator);
    
    /**
     * @notice 构造函数：部署模板合约
     * @dev 模板合约只部署一次
     */
    constructor() {
        implementation = address(new TokenImplementation());
    }
    
    /**
     * @notice 创建克隆
     * @param name 代币名称
     * @param symbol 代币符号
     * @param initialSupply 初始供应量
     * @return 克隆合约地址
     * @dev 使用create2创建确定性地址的克隆
     */
    function createClone(
        string memory name,
        string memory symbol,
        uint256 initialSupply
    ) public returns (address) {
        // 使用create2创建克隆（需要实现最小代理合约）
        // 这里简化示例，实际需要使用EIP-1167标准
        bytes memory bytecode = getCloneBytecode();
        bytes32 salt = keccak256(abi.encodePacked(msg.sender, clones.length));
        
        address clone;
        assembly {
            clone := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
        }
        
        // 初始化克隆
        TokenImplementation(clone).initialize(name, symbol, initialSupply);
        
        // 记录克隆地址
        clones.push(clone);
        userClones[msg.sender].push(clone);
        
        emit CloneCreated(clone, msg.sender);
        return clone;
    }
    
    /**
     * @notice 获取克隆字节码（EIP-1167最小代理）
     * @dev 这是EIP-1167标准的最小代理合约字节码
     */
    function getCloneBytecode() internal view returns (bytes memory) {
        // EIP-1167最小代理合约字节码
        // 实际实现需要使用OpenZeppelin的Clones库
        return abi.encodePacked(
            hex"3d602d80600a3d3981f3363d3d373d3d3d363d73",
            implementation,
            hex"5af43d82803e903d91602b57fd5bf3"
        );
    }
}