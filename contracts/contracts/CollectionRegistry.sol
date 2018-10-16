pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./AppRegistry.sol";
import "./SchemaRegistry.sol";

contract CollectionRegistry {
    using SafeMath for uint256;

    struct Collection {
        bytes32 appId;
        bytes32 schemaId;
        IncentivePolicy policy;
        mapping (bytes32 => bool) auth;
    }

    struct IncentivePolicy {
        // calculate with ETH. ex) 0.35ETH == 0.35%
        uint256 self;
        uint256 owner;
    }

    AppRegistry appReg;
    SchemaRegistry schemaReg;
    mapping (bytes32 => Collection) reg;

    constructor(
        AppRegistry _appReg,
        SchemaRegistry _schemaReg
    ) public {
        appReg = _appReg;
        schemaReg = _schemaReg;
    }

    function newCollection(bytes32 _appId, bytes32 _schemaId, uint256 _ratio) internal view returns (Collection memory) {
        require(schemaReg.check(_schemaId), "invalid schema");
        require(check(_schemaId), "collection already exists");
        return Collection({
            appId: _appId,
            schemaId: _schemaId,
            policy: IncentivePolicy({
                self: _ratio,
                owner: uint256(100 ether).sub(_ratio)
            })
        });
    }

    function register(
        bytes32 _appId, 
        bytes32 _schemaId, 
        uint256 _ratio
    ) public {
        require(appReg.checkOwner(_appId, msg.sender), "only owner can transfer ownership");
        reg[_schemaId] = newCollection(_appId, _schemaId, _ratio);
    }

    function unregister(bytes32 _id) public {
        require(appReg.checkOwner(reg[_id].appId, msg.sender), "only owner can transfer ownership");
        delete reg[_id];
    }

    function allow(bytes32 _id) public {
        reg[_id].auth[msg.sender] = true;
    }

    function deny(bytes32 _id) public {
        delete reg[_id].auth[msg.sender] = false;
    }

    function _get(bytes32 _id) internal view returns (Collection storage) {
        return reg[_id];
    }

    function get(bytes32 _id) public view returns (bytes32, bytes32) {
        Collection storage collection = _get(_id);
        return (
            collection.appId,
            collection.schemaId
        );
    }

    function check(bytes32 _id) public view returns (bool) {
        return (reg[_id].appId != bytes32(0x0));
    }

    function checkAllowed(bytes32 _id, bytes32 _uid) public view returns (bool) {
        return reg[_id].auth[_uid];
    }
}