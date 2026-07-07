// SPDX-License-Identifier: LGPL-3.0-or-later
// Copyright 2026 BOOSTRY Co., Ltd.
pragma solidity ^0.8.30;

/// @title ValidatorSet
/// @notice Reference validator contract for QBFT contract-based validator selection.
/// @dev
/// This contract exposes `getValidators()` as the canonical read interface used by
/// QBFT contract-based validator selection. The returned validator set excludes
/// validators that have put themselves into maintenance mode, while governance
/// voting power remains based on the full registered validator set.
contract ValidatorSet {
    /// @notice Governance operation type for validator set proposals.
    /// @dev Values are part of the proposal ID preimage. Do not reorder them in upgrades.
    enum Operation {
        Add,
        Remove
    }

    /// @dev Tracks votes for a single operation/candidate pair.
    /// The nested `voted` mapping prevents duplicate votes by the same validator.
    struct Proposal {
        uint256 votes;
        bool executed;
        mapping(address => bool) voted;
    }

    /// @notice Emitted when a candidate reaches quorum and is added as a validator.
    /// @param validator The validator address added to the registered validator set.
    event ValidatorAdded(address indexed validator);

    /// @notice Emitted when a candidate reaches quorum and is removed from the validator set.
    /// @param validator The validator address removed from the registered validator set.
    event ValidatorRemoved(address indexed validator);

    /// @notice Emitted whenever a validator casts a governance vote.
    /// @param proposalId Unique ID derived from the operation and candidate.
    /// @param operation The requested validator set operation.
    /// @param candidate The address proposed for addition or removal.
    /// @param voter The validator that cast the vote.
    /// @param votes Current number of votes recorded for the proposal after this vote.
    /// @param quorum Number of votes required for the proposal to execute.
    event ValidatorVote(
        bytes32 indexed proposalId,
        Operation indexed operation,
        address indexed candidate,
        address voter,
        uint256 votes,
        uint256 quorum
    );

    /// @notice Emitted when a validator enters maintenance mode.
    /// @param validator The validator excluded from `getValidators()` while in maintenance.
    event MaintenanceEntered(address indexed validator);

    /// @notice Emitted when a validator exits maintenance mode.
    /// @param validator The validator included again in `getValidators()`.
    event MaintenanceExited(address indexed validator);

    address[] private validators;
    mapping(address => bool) private validatorStatus;
    mapping(address => uint256) private validatorIndexes;
    mapping(address => bool) private maintenanceStatus;
    uint256 private maintenanceCount;
    mapping(bytes32 => Proposal) private proposals;

    modifier onlyValidator() {
        require(validatorStatus[msg.sender], "ValidatorSet: sender is not validator");
        _;
    }

    /// @notice Deploys the validator set with the initial registered validators.
    /// @dev
    /// The initial list must be non-empty, contain no zero address, and contain no
    /// duplicate addresses. All initial validators are active at deployment.
    /// @param initialValidators Initial registered validator addresses.
    constructor(address[] memory initialValidators) {
        require(initialValidators.length > 0, "ValidatorSet: empty validators");
        for (uint256 i = 0; i < initialValidators.length; i++) {
            _addValidator(initialValidators[i]);
        }
    }

    /// @notice Returns active validators for QBFT proposer selection and quorum calculation.
    /// @dev
    /// This is the compatibility method expected by the existing QBFT contract mode.
    /// Validators in maintenance are deliberately excluded from this returned list.
    /// The returned order follows the registered validator order, skipping maintenance entries.
    /// @return Active validator addresses.
    function getValidators() external view returns (address[] memory) {
        uint256 activeCount = validators.length - maintenanceCount;
        address[] memory activeValidators = new address[](activeCount);
        uint256 index = 0;

        for (uint256 i = 0; i < validators.length; i++) {
            address validator = validators[i];
            if (!maintenanceStatus[validator]) {
                activeValidators[index] = validator;
                index++;
            }
        }

        return activeValidators;
    }

    /// @notice Returns the full registered validator set.
    /// @dev Includes validators currently in maintenance.
    /// @return All registered validator addresses in registration order.
    function getAllValidators() external view returns (address[] memory) {
        address[] memory allValidators = new address[](validators.length);
        for (uint256 i = 0; i < validators.length; i++) {
            allValidators[i] = validators[i];
        }
        return allValidators;
    }

    /// @notice Checks whether an address is a registered validator.
    /// @param validator Address to check.
    /// @return True if the address is registered as a validator.
    function isValidator(address validator) external view returns (bool) {
        return validatorStatus[validator];
    }

    /// @notice Checks whether a validator is currently in maintenance.
    /// @param validator Address to check.
    /// @return True if the address is flagged as in maintenance.
    function isInMaintenance(address validator) external view returns (bool) {
        return maintenanceStatus[validator];
    }

    /// @notice Returns the tolerated Byzantine fault count `f`.
    /// @dev Calculated from the full registered validator count as `(N - 1) / 3`.
    /// @return Number of faulty validators tolerated by the current registered set size.
    function faultTolerance() public view returns (uint256) {
        return (validators.length - 1) / 3;
    }

    /// @notice Returns the validator governance quorum.
    /// @dev Uses the QBFT-style threshold `2f + 1`, where `f` is based on registered validators.
    /// @return Number of votes required to execute add/remove proposals.
    function quorumSize() public view returns (uint256) {
        return (2 * faultTolerance()) + 1;
    }

    /// @notice Returns the number of validators currently in maintenance.
    /// @return Current maintenance validator count.
    function maintenanceSize() external view returns (uint256) {
        return maintenanceCount;
    }

    /// @notice Computes the proposal ID for a validator set operation.
    /// @dev The same operation/candidate pair maps to the same proposal for all voters.
    /// @param operation Validator set operation type.
    /// @param candidate Candidate address proposed for addition or removal.
    /// @return Unique proposal ID.
    function proposalId(Operation operation, address candidate) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(operation, candidate));
    }

    /// @notice Returns the number of votes recorded for a proposal.
    /// @param operation Validator set operation type.
    /// @param candidate Candidate address proposed for addition or removal.
    /// @return Number of votes recorded for the proposal.
    function proposalVotes(Operation operation, address candidate) external view returns (uint256) {
        return proposals[proposalId(operation, candidate)].votes;
    }

    /// @notice Checks whether a proposal has already executed.
    /// @param operation Validator set operation type.
    /// @param candidate Candidate address proposed for addition or removal.
    /// @return True if the proposal reached quorum and executed.
    function proposalExecuted(Operation operation, address candidate) external view returns (bool) {
        return proposals[proposalId(operation, candidate)].executed;
    }

    /// @notice Checks whether a validator has voted for a proposal.
    /// @param operation Validator set operation type.
    /// @param candidate Candidate address proposed for addition or removal.
    /// @param voter Validator address to check.
    /// @return True if `voter` has already voted for the proposal.
    function hasVoted(Operation operation, address candidate, address voter) external view returns (bool) {
        return proposals[proposalId(operation, candidate)].voted[voter];
    }

    /// @notice Votes to add a new validator.
    /// @dev
    /// Only registered validators may vote. The candidate must not already be a validator.
    /// When the proposal reaches `quorumSize()`, the candidate is added immediately.
    /// @param candidate Address proposed as a new validator.
    function voteAddValidator(address candidate) external onlyValidator {
        require(candidate != address(0), "ValidatorSet: zero candidate");
        require(!validatorStatus[candidate], "ValidatorSet: candidate is validator");

        _vote(Operation.Add, candidate);
    }

    /// @notice Votes to remove an existing validator.
    /// @dev
    /// Only registered validators may vote. The candidate must currently be registered.
    /// When the proposal reaches `quorumSize()`, the candidate is removed immediately.
    /// If the candidate is in maintenance, the maintenance count is decremented as part of removal.
    /// @param candidate Existing validator proposed for removal.
    function voteRemoveValidator(address candidate) external onlyValidator {
        require(validatorStatus[candidate], "ValidatorSet: candidate is not validator");
        require(validators.length > 1, "ValidatorSet: cannot remove last validator");

        _vote(Operation.Remove, candidate);
    }

    /// @notice Puts the caller into maintenance mode.
    /// @dev
    /// Maintenance mode excludes the caller from `getValidators()` but does not remove
    /// governance voting rights. At most `faultTolerance()` validators may be in
    /// maintenance at the same time.
    function enterMaintenance() external onlyValidator {
        require(!maintenanceStatus[msg.sender], "ValidatorSet: already in maintenance");
        require(maintenanceCount < faultTolerance(), "ValidatorSet: too many validators in maintenance");

        maintenanceStatus[msg.sender] = true;
        maintenanceCount++;

        emit MaintenanceEntered(msg.sender);
    }

    /// @notice Removes the caller from maintenance mode.
    /// @dev The caller becomes visible through `getValidators()` again after this call.
    function exitMaintenance() external onlyValidator {
        require(maintenanceStatus[msg.sender], "ValidatorSet: not in maintenance");

        maintenanceStatus[msg.sender] = false;
        maintenanceCount--;

        emit MaintenanceExited(msg.sender);
    }

    /// @dev Records a validator vote and executes the proposal immediately once quorum is reached.
    /// Reverts if the proposal already executed or the caller already voted for it.
    function _vote(Operation operation, address candidate) private {
        bytes32 id = proposalId(operation, candidate);
        Proposal storage proposal = proposals[id];
        require(!proposal.executed, "ValidatorSet: proposal executed");
        require(!proposal.voted[msg.sender], "ValidatorSet: already voted");

        proposal.voted[msg.sender] = true;
        proposal.votes++;

        uint256 quorum = quorumSize();
        emit ValidatorVote(id, operation, candidate, msg.sender, proposal.votes, quorum);

        if (proposal.votes >= quorum) {
            proposal.executed = true;
            if (operation == Operation.Add) {
                _addValidator(candidate);
                emit ValidatorAdded(candidate);
            } else {
                _removeValidator(candidate);
                emit ValidatorRemoved(candidate);
            }
        }
    }

    /// @dev Adds a validator to the registered set and preserves registration order.
    /// Reverts on zero address or duplicate registration.
    function _addValidator(address validator) private {
        require(validator != address(0), "ValidatorSet: zero validator");
        require(!validatorStatus[validator], "ValidatorSet: duplicate validator");

        validatorStatus[validator] = true;
        validatorIndexes[validator] = validators.length;
        validators.push(validator);
    }

    /// @dev Removes a validator from the registered set while preserving the order of remaining validators.
    /// Clears maintenance state first so `maintenanceCount` remains consistent.
    function _removeValidator(address validator) private {
        require(validatorStatus[validator], "ValidatorSet: validator not found");

        if (maintenanceStatus[validator]) {
            maintenanceStatus[validator] = false;
            maintenanceCount--;
        }

        uint256 index = validatorIndexes[validator];
        for (uint256 i = index; i < validators.length - 1; i++) {
            address nextValidator = validators[i + 1];
            validators[i] = nextValidator;
            validatorIndexes[nextValidator] = i;
        }

        validators.pop();
        delete validatorIndexes[validator];
        delete validatorStatus[validator];
    }
}
