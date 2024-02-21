// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";
import { Predeploys } from "src/libraries/Predeploys.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { ISemver } from "src/universal/ISemver.sol";
import { L2ToL1MessagePasser } from "src/L2/L2ToL1MessagePasser.sol";
import { Constants } from "src/libraries/Constants.sol";

/// @custom:proxied
/// @custom:predeploy TODO: fill in address
/// @title IBCCrossDomainMessenger
/// @notice The IBCCrossDomainMessenger is a high-level interface for message passing between the EVM
///         an ABCI state machine.
contract IBCCrossDomainMessenger is CrossDomainMessenger, ISemver {
    /// @custom:semver 1.8.0
    string public constant version = "1.8.0";

    /// @notice Emitted when a contract wants to send a msg to the ABCI state machine.
    /// @param to         Address that the message is directed to.
    /// @param gasLimit   Gas limit for execution of the message.
    /// @param value      Value associated with the message.
    /// @param data       ABI encoded message to be executed.
    event SentMessage(address indexed to, uint256 gasLimit, uint256 value, bytes data);


    /// @notice Constructs the IBCCrossDomainMessenger contract.
    /// @param _l1CrossDomainMessenger Address of the L1CrossDomainMessenger contract.
    /// TODO: necessary?
    constructor(address _l1CrossDomainMessenger) CrossDomainMessenger(_l1CrossDomainMessenger) {
        initialize();
    }

    /// @notice Initializer.
    function initialize() public initializer {
        __CrossDomainMessenger_init();
    }

    /// @custom:legacy
    /// @notice Legacy getter for the remote messenger.
    ///         Use otherMessenger going forward.
    /// @return Address of the L1CrossDomainMessenger contract.
    function l1CrossDomainMessenger() public view returns (address) {
        return OTHER_MESSENGER;
    }

    /// @inheritdoc CrossDomainMessenger
    function _sendMessage(address _to, uint64 _gasLimit, uint256 _value, bytes memory _data) internal override {
        emit SentMessage(_to, _gasLimit, _value, _data);
    }

    /// @inheritdoc CrossDomainMessenger
    function _isOtherMessenger() internal view override returns (bool) {
        return AddressAliasHelper.undoL1ToL2Alias(msg.sender) == OTHER_MESSENGER;
    }

    /// @inheritdoc CrossDomainMessenger
    function _isUnsafeTarget(address _target) internal view override returns (bool) {
        return _target == address(this) || _target == address(Predeploys.L2_TO_L1_MESSAGE_PASSER);
    }
}
