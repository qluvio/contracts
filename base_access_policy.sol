pragma solidity 0.4.24;

import "./ownable.sol";
import "./meta_object.sol";
import "./lib_accessmanager.sol";
import "./editable.sol";

contract BaseAccessPolicy is MetaObject, AccessIndexor, Editable {

    string public description;

    // conditions can be stored as part of metadata - i.e. in MetaObject ???

    string public constant EFFECT_ALLOW = "allow";
    string public constant EFFECT_DENY = "deny";
    string public effect;

    // TODO: actually, maybe this goes in content space ...?
    event AccessPolicyCreated(address space);
    event AccessPolicyDeleted(address space); // TODO: kill method?

    constructor(address _contentSpace)  public payable {
        contentSpace = _contentSpace; // TODO: why isn't content space in constructor of Ownable?
        effect = EFFECT_ALLOW; // default is allow
    }

    int public constant OP_ADD_SUBJECT = 1;
    int public constant OP_REMOVE_SUBJECT = 2;
    int public constant OP_ADD_RESOURCE = 3;
    int public constant OP_REMOVE_RESOURCE = 4;
    int public constant OP_ADD_ACTION = 5;
    int public constant OP_REMOVE_ACTION = 6;
    int public constant OP_SET_EFFECT = 7;

    // same as AccessIndexor ...?
    uint8 public constant CATEGORY_MIN = 1;
    uint8 public constant CATEGORY_CONTENT_OBJECT = 1;
    uint8 public constant CATEGORY_GROUP = 2;
    uint8 public constant CATEGORY_LIBRARY = 3;
    uint8 public constant CATEGORY_CONTENT_TYPE = 4;
    uint8 public constant CATEGORY_CONTRACT = 5;
    uint8 public constant CATEGORY_POLICY = 6;
    uint8 public constant CATEGORY_SPACE = 7;
    uint8 public constant CATEGORY_MAX = 7;

    uint8 public constant SUBJECT_ACCOUNT = 100;
    uint8 public constant SUBJECT_GROUP = 101;

    uint8 public constant REL_EFFECT = 200;
    uint8 public constant REL_ACTION = 201;

    event AccessPolicyChanged(address space, int op, bytes32 relTarget, uint8 relType);

    function setEffect(string _newEffect) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        require(keccak256(abi.encodePacked(_newEffect)) == keccak256(abi.encodePacked(EFFECT_ALLOW)) ||
        keccak256(abi.encodePacked(_newEffect)) == keccak256(abi.encodePacked(EFFECT_DENY)));
        effect = _newEffect;
        bytes32 result;
        assembly {
            result := mload(add(_newEffect, 32))
        }
        emit AccessPolicyChanged(contentSpace, OP_SET_EFFECT, result, REL_EFFECT);
        return true;
    }

    function addUser(address _subject) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        emit AccessPolicyChanged(contentSpace, OP_ADD_SUBJECT, bytes32(_subject), SUBJECT_ACCOUNT);
        return true;
    }

    function removeUser(address _subject) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        emit AccessPolicyChanged(contentSpace, OP_REMOVE_SUBJECT, bytes32(_subject), SUBJECT_ACCOUNT);
        return true;
    }

    function addGroup(address _subject) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        emit AccessPolicyChanged(contentSpace, OP_ADD_SUBJECT, bytes32(_subject), SUBJECT_GROUP);
        return true;
    }

    function removeGroup(address _subject) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        emit AccessPolicyChanged(contentSpace, OP_REMOVE_SUBJECT, bytes32(_subject), SUBJECT_GROUP);
        return true;
    }

    function addResource(address _resource, uint8 _category) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        require(_category >= CATEGORY_MIN && _category <= CATEGORY_MAX);
        emit AccessPolicyChanged(contentSpace, OP_ADD_RESOURCE, bytes32(_resource), _category);
        return true;
    }

    function removeResource(address _resource, uint8 _category) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        require(_category >= CATEGORY_MIN && _category <= CATEGORY_MAX);
        emit AccessPolicyChanged(contentSpace, OP_REMOVE_RESOURCE, bytes32(_resource), _category);
        return true;
    }

    function addAction(string _action) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        bytes memory testAction = bytes(_action);
        require(testAction.length > 0 && testAction.length <= 32);
        bytes32 result;
        assembly {
            result := mload(add(_action, 32))
        }
        emit AccessPolicyChanged(contentSpace, OP_ADD_ACTION, result, REL_ACTION);
        return true;
    }

    function removeAction(string _action) public returns (bool) {
        require(msg.sender == owner || AccessManager.isAllowed(msg.sender, address(this), "edit"));
        bytes memory testAction = bytes(_action);
        require(testAction.length > 0 && testAction.length <= 32);
        bytes32 result;
        assembly {
            result := mload(add(_action, 32))
        }
        emit AccessPolicyChanged(contentSpace, OP_REMOVE_ACTION, result, REL_ACTION);
        return true;
    }
}