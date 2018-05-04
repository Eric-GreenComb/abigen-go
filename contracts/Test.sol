pragma solidity ^0.4.11;

contract Test {

    string public value;

    function ModifyValue(string val) {
        value = val;
    }
}