//SPDX-License-Identifier: MIT
pragma solidity ^0.7.1;

import {Content} from "./content.sol";
import {BaseContent} from "./base_content.sol";
import {BaseAccessControlGroup} from "./base_access_control_group.sol";


//
// Following features are implemented:
//      - ensure that only owner and accessor registered in a specified group can view the content
//      - ensure that non-owner can only download the package within the availability window
//


/* -- Revision history --
AvlDelivery20190404103300ML: First versioned released
AvlDelivery20190510152500ML: updated for new runAccessInfo API
AvlDelivery20200211163600ML: updated to comform to authV3
*/


contract AvailsDelivery is Content {

    struct AvailabilityData {
        bool clearedSD; // cleared for SD
        bool clearedHD; // cleared for HD
        uint startDate; // Available on
        uint endDate; // Unavailable on
        address region; // Territory
    }

    mapping(address => address) public accessors;  // Indexed using OTT Delivery BaseContent address
    mapping(address => AvailabilityData) public availability; // Indexed using OTT Delivery BaseContent address
    
    constructor(){
        version ="AvlDelivery20200211163600ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    }


    function setAvailability(address payable content, bool sd,bool hd, uint start, uint end, address region ) public {
        BaseContent c = BaseContent(content);
        require(msg.sender == c.owner());
        uint startDate = start;
        if (startDate == 0) {
            startDate = block.timestamp;
        }
        availability[content] = AvailabilityData(sd, hd, startDate, end, region);
    }

    function setAccessorGroup(address payable content, address accessor) public {
        BaseContent c = BaseContent(content);
        require(msg.sender == c.owner());
        accessors[content] = accessor;
    }

    function validateAvailability(address payable accessor) public view {
        require(hasAccess(msg.sender, accessor));
        require(isAvailable(msg.sender, accessor) == 0);
    }

    function hasAccess(address payable content, address payable accessor) public view returns (bool) {
        address payable accessorGroup = payable(accessors[content]);
        if (accessorGroup == address(0x0)) {
            return true;
        }
        BaseAccessControlGroup group = BaseAccessControlGroup(accessorGroup);
        return group.hasAccess(accessor);
    }

    function isAvailable(address content, address payable accessor) public view returns (uint8) {
        uint8 available = 0;
        AvailabilityData storage avail = availability[content];
        address payable region = payable(avail.region);

        if (region != address(0x0)) {
            BaseAccessControlGroup group = BaseAccessControlGroup(region);
            if (group.hasAccess(accessor) == false) {
                available = 10;
                //emit Log("Wrong territory");
            }
        }
        if (avail.startDate > block.timestamp) {
            available = 20;
            //emit Log("Not yet available");
        }
        if ((avail.endDate != 0) && (avail.endDate < block.timestamp)) {
            available = 30;
            //emit Log("Not longer available");
        }
        return available;
    }


    function runAccessInfo(
        bytes32[] memory, /*customValues*/
        address[] memory, /*stakeholders*/
        address payable accessor
    )
    public view override returns (uint8, uint8, uint8, uint256)
    {
        if (!hasAccess(msg.sender,accessor)) {
            return (0, 100, 0, 0);
        }
        uint8 availCode = isAvailable(msg.sender, accessor);
        if (availCode != 0) {
            return (0, availCode, 0, 0);
        }
        return (DEFAULT_SEE + DEFAULT_ACCESS, 0, 0, 0);
    }


    function runAccess(
        bytes32, /*request_ID*/
        bytes32[] memory, /*custom_values*/
        address[] memory, /*stake_holders*/
        address payable accessor
    )
    public payable returns(uint)
    {
        validateAvailability(accessor);
        return 0;
    }

}
