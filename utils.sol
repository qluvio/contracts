pragma solidity >0.4.24;

import {Ownable} from "./ownable.sol";

library Utils {

    bytes4 constant checkV3Sig = bytes4(keccak256("versionAPI()"));
    function checkV3Contract(address _checkAddr) internal view returns (bool) {
        bool success;
        bytes memory data;
        (success, data) = _checkAddr.staticcall(abi.encodeWithSelector(checkV3Sig));
        if (!success || data.length == 0 || data[0] == 0x0) {
            return false;
        }
        return true;
    }
}

contract Verifier {
    function recoverAddr(bytes32 msgHash, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        return ecrecover(msgHash, v, r, s);
    }

    function isSigned(address _addr, bytes32 msgHash, uint8 v, bytes32 r, bytes32 s) public pure returns (bool) {
        return ecrecover(msgHash, v, r, s) == _addr;
    }
}
