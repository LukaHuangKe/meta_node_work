// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TodoList {
    struct Todo {
        string todoName;
        bool completed;
    }

    mapping(address => Todo[]) private userTodos;
    uint public constant MAX_TODO_SIZE = 100;

    event TodoAdded(address indexed user, uint index, string todoName);
    event TodoCompleted(address indexed user, uint index);
    event TodoDeleted(address indexed user, uint index);

    // 添加待办
    function addTodo(string memory task) public {
        require(bytes(task).length > 0, "Task cannot be empty");
        require(userTodos[msg.sender].length < MAX_TODO_SIZE, "User has reached the maximum todo limit");
        userTodos[msg.sender].push(Todo({
            todoName: string(task), 
            completed: false
        }));
        emit TodoAdded(msg.sender, userTodos[msg.sender].length - 1, task);
    }

    // 标记为完成
    function completeTodo(uint index) public {
        require(index < userTodos[msg.sender].length, "Invalid index range");
        userTodos[msg.sender][index].completed = true;
        emit TodoCompleted(msg.sender, index);
    }

    // 删除待办（快速删除，不保序）
    function deleteTodo(uint index) public{
        require(index < userTodos[msg.sender].length, "Invalid index range");
        userTodos[msg.sender][index] = userTodos[msg.sender][userTodos[msg.sender].length - 1];
        userTodos[msg.sender].pop();
        emit TodoDeleted(msg.sender, index);
    }

    // 获取所有待办。返回的数据是从 storage （ userTodos 是 storage）复制一份到 memory，
    // 且Solidity 规定：从函数返回的复杂类型（数组、struct）必须在 memory 中
    function getAllTodos() public view returns (Todo[] memory) {
        return userTodos[msg.sender];
    }

    // 获取待办数量
    function getTodoLength() public view returns (uint) {
        return userTodos[msg.sender].length;
    }

    // 获取未完成的待办
    function getUnCompletedTodos() public view returns (Todo[] memory) {
        Todo[] memory allTodos = getAllTodos();
        uint count = 0;
        for(uint i = 0; i < allTodos.length; i++){
            if(!allTodos[i].completed){
                count++;
            }
        }

        Todo[] memory res = new Todo[](count);
        uint index = 0;
        for(uint i = 0; i < allTodos.length; i++){
            if(!allTodos[i].completed){
                res[index] = allTodos[i];
                index++;
            }
        }
        return res;
    }
}