pragma solidity 0.4.24;

import {Editable} from "./editable.sol";
import {Content} from "./content.sol";
import {Container} from "./container.sol";
import {KMSSpace} from "./kms_space.sol";
import {UserSpace} from "./user_space.sol";
import {AccessIndexor} from "./access_indexor.sol";
import {Publishable} from "./publishable.sol";

/* -- Revision history --
BaseContent20190221101600ML: First versioned released
BaseContent20190301121900ML: Adds support for getAccessInfo, to replace getAccessCharge (not deprecated yet)
BaseContent20190315175100ML: Migrated to 0.4.24
BaseContent20190321122100ML: accessRequest returns requestID, removed ml_hash from access_complete event
BaseContent20190510151500ML: creation via ContentSpace factory, modified getAccessInfo API
BaseContent20190510151500ML: creation via ContentSpace factory, modified getAccessInfo API
BaseContent20190522154000SS: Changed hash bytes32 to string
BaseContent20190528193400ML: Modified to support non-library containers
BaseContent20190605203200ML: Splits publish and confirm logic
BaseContent20190724203300ML: Enforces access rights in access request
BaseContent20190801141600ML: Fixes the access rights grant for paid content
BaseContent20191029161700ML: Removed debug statements for accessRequest
BaseContent20191107151400ML: Differentiates onlyEditor from onlyOwner
*/


contract BaseContent is Editable, Publishable {

    bytes32 public version ="BaseContent20191107151400ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    address public contentType;
    address public addressKMS;
    address public contentContractAddress;
    address public libraryAddress;

    uint256 public accessCharge;
    //bool refundable;

    bytes32 public constant STATUS_PUBLISHED = "Published";
    bytes32 public constant STATUS_DRAFT = "Draft";
    bytes32 public constant STATUS_REVIEW = "Draft in review";

    uint8 public visibility = 0;
    uint8 public constant CAN_SEE = 1;
    uint8 public constant CAN_ACCESS = 10;
    uint8 public constant CAN_EDIT = 100;

    struct RequestData {
        address originator; // client address requesting
        uint256 amountPaid; // number of token received
        int8 status; //0 access requested, 1 access granted, -1 access refused, 2 access completed, -2 access error
        uint256 settled; //Amount of the escrowed money (amountPaid) that has been settled (paid to owner or refunded)
    }
   /* //ML: removing for now as this enables the editing of field that are not meant to be edited
      //   this method should be limited to object that are in the process of being duplicated
    function migrate(address _contentType,
            address _addressKMS,
            address _contentContractAddress,
            // address _libraryAddress,
            uint256 _accessCharge,
            int _statusCode,
            // uint256 _requestID,
            uint8 _visibility,
            string _objectHash,
            string _versionHashes
        ) public onlyOwner {

        contentType = _contentType;
        addressKMS = _addressKMS;
        contentContractAddress = _contentContractAddress;
        // libraryAddress = _libraryAddress; // TODO: set by library factory method?

        accessCharge = _accessCharge;
        statusCode = _statusCode;
        // requestID = _requestID;
        visibility = _visibility;

        super.migrate(_objectHash, _versionHashes);

        return;
    }
    */

    // TODO: remove?
    mapping(uint256 => RequestData) public requestMap;

    event ContentObjectCreate(address containingLibrary);
    event SetContentType(address contentType, address contentContractAddress);

    event LogPayment(uint256 requestID, string label, address payee, uint256 amount);
    event AccessGrant(uint256 requestID, bool access_granted, string reKey, string encryptedAESKey);
    event AccessRequestValue(bytes32 customValue);
    event AccessRequestStakeholder(address stakeholder);
    event SetContentContract(address contentContractAddress);

    event SetAccessCharge(uint256 accessCharge);
    event GetAccessCharge(uint8 level, uint256 accessCharge);
    event InsufficientFunds(uint256 accessCharge, uint256 amountProvided);
    event Publish(bool requestStatus, int statusCode, string objectHash);

    // Debug events
    event InvokeCustomPreHook(address custom_contract);
    event ReturnCustomHook(address custom_contract, uint256 result);
    event InvokeCustomPostHook(address custom_contract);

    modifier onlyFromLibrary() {
        require(msg.sender == libraryAddress);
        _;
    }

    constructor(address content_space, address lib, address content_type) public payable {
        contentSpace = content_space;
        libraryAddress = lib;
        statusCode = -1;
        contentType = content_type;
        visibility = CAN_ACCESS; //default could be made a function of the library.

        //get custom contract address associated with content_type from hash
        /*
        BaseLibrary lib = BaseLibrary(libraryAddress);
        contentContractAddress = lib.contentTypeContracts(content_type);
        addressKMS = lib.addressKMS();
        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            require(c.runCreate() == 0);
        }
        */
        emit ContentObjectCreate(libraryAddress);
    }


    function setVisibility(uint8 visibility_code) public onlyEditor { //debatable whether this should be kept onlyOwner as visibility influence who can edit
        visibility = visibility_code;
    }

    function statusDescription() public constant returns (bytes32) {
        return statusCodeDescription(statusCode);
    }

    function statusCodeDescription(int status_code) public constant returns (bytes32) {
        bytes32 codeDescription = 0x0;
        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            codeDescription = c.runDescribeStatus(status_code);
        }
        if (codeDescription != 0x0) {
            return codeDescription;
        }
        if (status_code == 0) {
            return STATUS_PUBLISHED;
        }
        if (status_code < 0) {
            return  STATUS_DRAFT;
        }
        if (status_code > 0) {
            return STATUS_REVIEW;
        }
    }

    /* should be deprecated
    function setContentType(address content_type) public onlyOwner {
        contentType = content_type;
        //get custom contract address associated with content_type from hash
        BaseLibrary lib = BaseLibrary(libraryAddress);
        contentContractAddress = lib.contentTypeContracts(content_type);
        addressKMS = lib.addressKMS();
        if (contentContractAddress != 0x0) {

        }
        emit SetContentType(content_type, contentContractAddress);
    }

    */
    function setStatusCode(int status_code) public returns (int) {
        if ((tx.origin == owner) && ((status_code < 0) || ((status_code > 0) && (statusCode < 0)))) {

            //Owner can revert content to draft mode regardless of status (debatable for published, viewable content)
            //  and owner can move status from draft to in review
            statusCode = status_code;
        }

        if (msg.sender == libraryAddress) {
            // Library has full authority to change status, it is assumed entitlement is done there
            statusCode = status_code;
        }

        emit SetStatusCode(statusCode);
        return statusCode;
    }

    function setAddressKMS(address address_KMS) public onlyEditor {
        addressKMS = address_KMS;
    }

    function getKMSInfo(bytes prefix) public view returns (string, string) {
        KMSSpace contentSpaceObj = KMSSpace(contentSpace);
        if (addressKMS == 0x0 || contentSpaceObj.checkKMSAddr(addressKMS) == 0) {
            return ("", "");
        }
        return contentSpaceObj.getKMSInfo(contentSpaceObj.getKMSID(addressKMS), prefix);
    }

    //Owner can change this, unless the contract they are already set it prevent them to do so.
    function setContentContractAddress(address addr) public onlyEditor {
        Content c;
        if (contentContractAddress != 0x0) {
            c = Content(contentContractAddress);
            uint killStatus = c.runKill();
            require(killStatus == 0);
        }
        contentContractAddress = addr;
        if (addr != 0x0) {
            c = Content(addr);
            uint createStatus = c.runCreate();
            require(createStatus == 0);
        }
        emit SetContentContract(contentContractAddress);
    }

    // Visibility codes
    //      0   -> visible
    //      10  -> content not published
    //      100 -> calculation of price exceeds specified cap (accessCharge)
    //      255 -> unset
    // Access codes
    //      0   -> paid for
    //      10  -> content not published
    //      100 -> content available if paid for
    //      255 -> unset

    function getWIPAccessInfo(address accessor) private view returns (uint8, uint8, uint256) {
        if ((accessor == owner) || (visibility >= CAN_EDIT) ){
            return (0, 0, accessCharge);
        }
<<<<<<< HEAD
        UserSpace contentSpaceObj = UserSpace(contentSpace);
        address userWallet = contentSpaceObj.userWallets(tx.origin);
=======
        BaseContentSpace contentSpaceObj = BaseContentSpace(contentSpace);
        address userWallet = contentSpaceObj.userWallets(accessor);
>>>>>>> dev-access-req-redux
        if (userWallet != 0x0) {
            AccessIndexor wallet = AccessIndexor(userWallet);
            if (wallet.checkContentObjectRights(address(this), wallet.TYPE_EDIT()) == true) {
                return (0, 0, accessCharge);
            }
        }
        if (Container(libraryAddress).canReview(accessor) == true) { //special case of pre-publish review
            return (0, 0, accessCharge);
        }
        return (10, 10, accessCharge);
    }

    function getCustomInfo(uint8 level, bytes32[] custom_values, address[] stakeholders) public view returns (uint8, uint8, uint256) {
        uint256 levelAccessCharge = accessCharge;
        uint8 visibilityCode = (visibility >= CAN_SEE) ? 0 : 255;
        uint8 accessCode = (visibility >= CAN_ACCESS) ? 0 :255;
        if (contentContractAddress != 0x0) {
            uint8 customMask;
            uint8 customVisibility;
            uint8 customAccess;
            uint256 customCharge;
            Content c = Content(contentContractAddress);
            (customMask, customVisibility, customAccess, customCharge) = c.runAccessInfo(level, custom_values, stakeholders);
            if (customCharge > accessCharge) {
                visibilityCode = 100;
            } else {
                if ((customMask & c.DEFAULT_SEE()) == 0) {
                    visibilityCode = customVisibility;
                }
                if ((customMask & c.DEFAULT_ACCESS()) == 0) {
                    accessCode = customAccess;
                }
                if ((customMask & c.DEFAULT_CHARGE()) == 0) {
                    levelAccessCharge = customCharge;
                }
            }
        }
        return (visibilityCode, accessCode, levelAccessCharge);
    }

    function getAccessInfo(address accessor, uint8 level, bytes32[] custom_values, address[] stakeholders) public view returns (uint8, uint8, uint256) {

        if (statusCode != 0) {
            return getWIPAccessInfo(accessor); //broken out to reduce complexity (compiler failed)
        }
        uint256 levelAccessCharge;
        uint8 visibilityCode;
        uint8 accessCode;
        (visibilityCode, accessCode, levelAccessCharge) = getCustomInfo( level, custom_values, stakeholders);//broken out to reduce complexity (compiler failed)

        if ((visibilityCode == 255) || (accessCode == 255) ) {
            UserSpace contentSpaceObj = UserSpace(contentSpace);
            address userWallet = contentSpaceObj.userWallets(accessor);
            if (userWallet != 0x0) {
                AccessIndexor wallet = AccessIndexor(accessor);
                if (visibilityCode == 255) { //No custom calculations
                    if (wallet.checkContentObjectRights(address(this), wallet.TYPE_SEE()) == true) {
                        visibilityCode = 0;
                    }
                }
                if (visibilityCode == 0) { //if content is not visible, no point in checking if it is accessible
                    if (accessCode == 255) {
                        if (wallet.checkContentObjectRights(address(this), wallet.TYPE_ACCESS()) == true) {
                            accessCode = 0;
                        } else {
                            accessCode = 100; //content accessible if paid for
                        }
                    }
                }
            }
        }
        return (visibilityCode, accessCode, levelAccessCharge);
    }

    function setAccessCharge(uint256 charge) public onlyEditor returns (uint256) {
        accessCharge = charge;
        emit SetAccessCharge(accessCharge);
        return accessCharge;
    }

    function canEdit() public view returns (bool) {
        UserSpace contentSpaceObj = UserSpace(contentSpace);
        address walletAddress = contentSpaceObj.userWallets(tx.origin);
        AccessIndexor wallet = AccessIndexor(walletAddress);
        return wallet.checkContentObjectRights(address(this), wallet.TYPE_EDIT());
    }

    /*
    function canPublish() public view returns (bool) {
        return (canEdit() || msg.sender == libraryAddress);
    }

    function canCommit() public view returns (bool) {
        return canEdit();
    }*/

    function canConfirm() public view returns (bool) {
        Container lib = Container(libraryAddress);
        return lib.canNodePublish(msg.sender);
    }

    // override from Editable
    function parentAddress() public view returns (address) {
        return libraryAddress;
    }

    // TODO: why payable?
    function publish() public payable returns (bool) {
        bool submitStatus = Container(libraryAddress).publish(address(this));
        // Log event
        emit Publish(submitStatus, statusCode, objectHash); // TODO: confirm?
        return submitStatus;
    }

    function updateStatus(int status_code) public returns (int) {
        // require((tx.origin == owner) || (msg.sender == owner) || (msg.sender == libraryAddress));
        require(canEdit() || msg.sender == libraryAddress);
        int newStatusCode;
        if (contentContractAddress == 0x0) {
            if (canEdit() && ((status_code == -1) || (status_code == 1))) {
                newStatusCode = status_code; //owner can change status back to draft or to in-review
            } else if ((msg.sender == libraryAddress) && (statusCode >= 0)) {
                newStatusCode = status_code; //library can change status of content in review to any status
            }
        } else {
            Content c = Content(contentContractAddress);
            newStatusCode = c.runStatusChange(status_code);
        }
        statusCode = newStatusCode;
        emit SetStatusCode(statusCode);
        return statusCode;
    }


    //this function allows custom content contract to call makeRequestPayment
    function processRequestPayment(uint256 request_ID, address payee, string label, uint256 amount) public returns (bool) {
        require((contentContractAddress != 0x0) && (msg.sender == contentContractAddress));
        return makeRequestPayment(request_ID, payee, label, amount);
    }

    function makeRequestPayment(uint256 request_ID, address payee, string label, uint256 amount) private returns (bool) {
        RequestData storage r = requestMap[request_ID];
        if ((r.settled + amount) <= r.amountPaid) {
            payee.transfer(amount);
            r.settled = r.settled + amount;
            emit LogPayment(request_ID, label, payee, amount);
        }
        return true;
    }

    /*
    event DbgAccess(
        uint256 charged,
        uint received,
        uint converted,
        bool enough
    );
    event DbgAccessCode(uint8 code);
    */


    function updateRequest() public {
        if (contentContractAddress == 0x0) {
            super.updateRequest();
        } else {
            Content c = Content(contentContractAddress);
            uint editCode = c.runEdit();
            if (editCode == 100) {
                super.updateRequest();
            } else {
                require(editCode == 0);
                emit UpdateRequest(objectHash);
            }
        }
    }

    function accessRequestContext(
        bytes32 requestNonce,
        bytes32 contextHash,
        address accessor,
        uint256 request_timestamp
    ) public payable returns (bytes32) {
        bytes32[] memory emptyVals;
        address[] memory emptyAddrs;
        return accessRequestInternal(requestNonce, 0, "", emptyVals, emptyAddrs, contextHash, accessor, request_timestamp);
    }

    event AccessRequest(
        bytes32 requestNonce,
        // uint8 level, TODO: WHY?
        // string contentHash, TODO: WHY?
        // string pkeRequestor, TODO: not used
        address libraryAddress, // likely will need for tenancy - but could be something else ...?
        string pkeAFGH,
        bytes32 contextHash,
        address accessor,           // I've called this 'userAddress' - don't care too much but ideally it's the same name wherever it means the same thing!
        uint256 request_timestamp // always millis - either set from context or with blockchain now()
    );

    //  level - the security group for which the access request is for
    //  pkeRequestor - ethereum public key of the requestor (ECIES)
    //  pkeAFGH - ephemeral public key of the requestor (AFGH)
    //  customValues - an array of custom values used to convey additional information
    //  stakeholders - an array of additional address used to provide additional relevant addresses
    function accessRequest(
        bytes32 requestNonce,
        uint8 level,
        string pkeAFGH,
        bytes32[] custom_values,
        address[] stakeholders
    )
        public payable returns (bytes32) {
        return accessRequestInternal(requestNonce, level, pkeAFGH, custom_values, stakeholders, 0x0, msg.sender, now * 1000);
    }

    function validateAccess(address accessor, uint8 level, bytes32[] custom_values, address[] stakeholders) internal {
        uint256 requiredFund;
        uint8 visibilityCode;
        uint8 accessCode;

        (visibilityCode, accessCode, requiredFund) = getAccessInfo(accessor, level, custom_values, stakeholders);

        if (accessCode == 100) { //Check if request is funded, except if user is owner or has paid already
            require(msg.value >= uint(requiredFund));
            //emit DbgAccess(requiredFund, msg.value, uint(requiredFund), (msg.value >= uint(requiredFund)));
            setPaidRights();
            accessCode = 0;
        }
        require(accessCode == 0);
    }

    function accessRequestInternal(
        bytes32 requestNonce,
        uint8 level,
        string pke_AFGH,
        bytes32[] custom_values,
        address[] stakeholders,
        bytes32 contextHash,
        address accessor,
        uint256 request_timestamp
    )
    internal returns (bytes32) {

        validateAccess(accessor, level, custom_values, stakeholders);

        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            uint result = c.runAccess(requestNonce, level, custom_values, stakeholders);
            require(result == 0);
        }

        // Raise Event
        emit AccessRequest(requestNonce, libraryAddress, pke_AFGH, contextHash, accessor, request_timestamp);

        // Logs custom key/value pairs
        uint256 i;
        for (i = 0; i < custom_values.length; i++) {
            if (custom_values[i] != 0x0) {
                emit AccessRequestValue(custom_values[i]);
            }
        }
        for (i = 0; i < stakeholders.length; i++) {
            if (stakeholders[i] != 0x0) {
                emit AccessRequestStakeholder(stakeholders[i]);
            }
        }

        return requestNonce;
    }

    //The rekey provided is encrypted with the pkeRequestor
    // if reKey is blank, then access was denied
    function accessGrant(
        uint256 request_ID,
        bool access_granted,
        string re_key,
        string encrypted_AES_key
    )
        public returns (bool)
    {
        require((msg.sender == owner) || (msg.sender == addressKMS));

        RequestData storage r = requestMap[request_ID];
        require(r.originator != 0x0);
        bool result = access_granted;
        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            result = (c.runGrant(request_ID, access_granted) == 0);
        } else { //default behavior is settlement upon access grant
            if (r.settled < r.amountPaid) {
                if (access_granted == false) {
                    //escrowed fund to be refunded to accessor
                    makeRequestPayment(request_ID, r.originator, "access declined", r.amountPaid - r.settled);
                } else {
                    //escrowed fund to be paid to owner
                    makeRequestPayment(request_ID, owner, "owner payment", r.amountPaid - r.settled);
                }
            }
        }
        if (result == true) {
            r.status = 1;
            emit AccessGrant(request_ID, true, re_key, encrypted_AES_key);
        } else {
            r.status = -1;
            emit AccessGrant(request_ID, false, "", "");
        }
        return result;
    }

    // FROM will be context address?
    event AccessComplete(
        bytes32 requestID,
        uint256 scorePct,
        bool customContractResult,
        address libraryAddress,
        bytes32 contextHash,
        address accessor,           // I've called this 'userAddress' - don't care too much but ideally it's the same name wherever it means the same thing!
        uint256 request_timestamp   // always millis - either set from context or with blockchain now()
    );

    function accessCompleteContext(
        bytes32 requestID,
        uint256 score_pct,
        bytes32 contextHash,
        address accessor,
        uint256 request_timestamp
        ) public payable returns (bool) {

        bool success = (score_pct != 0);
        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            uint256 result = uint256(c.runFinalize(requestID, score_pct));
            success = (result == 0);
        }
        emit AccessComplete(requestID, score_pct, success, libraryAddress, contextHash, accessor, request_timestamp);
        return success;
    }

    // sender passes the quality score as pct of best possible (converted to 1-100 scale)
    // the fabric provides to this access,
    // hash of the version of the ML-computed segment matrix used (stored as a'part'),
    // and a maximum credit to a client based on failed SLA (quality)
    // score and hash are recorded in an event
    // contract has astate variable creditBack
    // a credit refund is calculated for the caller (if )
    // will true up the charge based on the quality score by transferring tokens back
    // to the sender
    //
    // add a state variable in the contract indicating whether to credit back based on quality score
    function accessComplete(bytes32 request_ID, uint256 score_pct, bytes32 ml_out_hash) public payable returns (bool) {
        require(ml_out_hash == ml_out_hash); //placeholder for verification of signature
        bool success = (score_pct != 0);
        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            uint256 result = uint256(c.runFinalize(request_ID, score_pct));
            success = (result == 0);
        }
        // record to event
        emit AccessComplete(request_ID, score_pct, success, libraryAddress, 0x0, msg.sender, now * 1000);
        return success;
    }

    function kill() public onlyFromLibrary {
        if (contentContractAddress != 0x0) {
            Content c = Content(contentContractAddress);
            uint canKill = c.runKill();
            require((canKill == 0) || (canKill == 100));
            if (canKill == 100) {
                c.kill();
            }
        }
        super.kill();
    }

    function setPaidRights() private {
        UserSpace contentSpaceObj = UserSpace(contentSpace);
        address walletAddress = contentSpaceObj.getAccessWallet();
        AccessIndexor indexor = AccessIndexor(walletAddress);
        indexor.setAccessRights();
    }

    function setRights(address stakeholder, uint8 access_type, uint8 access) public {
        UserSpace contentSpaceObj = UserSpace(contentSpace);
        address walletAddress = contentSpaceObj.userWallets(stakeholder);
        if (walletAddress == 0x0){
            //stakeholder is not a user (hence group or wallet)
            setGroupRights(stakeholder, access_type, access);
        } else {
            setGroupRights(walletAddress, access_type, access);
        }
    }

    function setGroupRights(address group, uint8 access_type, uint8 access) public {
        AccessIndexor indexor = AccessIndexor(group);
        indexor.setContentObjectRights(address(this), access_type, access);
    }

}
