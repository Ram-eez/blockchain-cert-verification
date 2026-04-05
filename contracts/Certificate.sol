// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Certificate {
    address public admin;

    struct Cert {
        uint256 timestamp;
        bool valid;
    }

    mapping(bytes32 => Cert) public certificates;

    constructor() {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        require(msg.sender == admin, "Not authorized");
        _;
    }

    function issueCertificate(bytes32 hash) public onlyAdmin {
        require(certificates[hash].timestamp == 0, "Already issued");

        certificates[hash] = Cert(block.timestamp, true);
    }

    function revokeCertificate(bytes32 hash) public onlyAdmin {
        require(certificates[hash].valid, "Not valid");

        certificates[hash].valid = false;
    }

    function verifyCertificate(bytes32 hash) public view returns (bool, uint256) {
        Cert memory cert = certificates[hash];
        return (cert.valid, cert.timestamp);
    }
}