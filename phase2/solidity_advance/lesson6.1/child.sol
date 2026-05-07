// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract GrandParent {
    function identify() public virtual returns (string memory) {
        return "GrandParent";
    }
}

contract Parent1 is GrandParent {
    function identify() public virtual override returns (string memory) {
        return "Parent1";
    }
}

contract Parent2 is GrandParent {
    function identify() public virtual override returns (string memory) {
        return "Parent2";
    }
}

contract Child is Parent1, Parent2 {
    // 遇到这种情况，子合约必须显式重写冲突的函数，并明确告诉编译器你要用哪个，不然编译报错
    function identify() public pure override(Parent1, Parent2) returns (string memory) {
        return "Child";
    }
}