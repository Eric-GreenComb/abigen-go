pragma solidity ^0.4.4;

contract Test {

    string public value;

    function ModifyValue(string val) {
        value = val;
    }
}