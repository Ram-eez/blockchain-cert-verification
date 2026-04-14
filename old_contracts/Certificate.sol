// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract CertificateSystem {

    address public admin;

    // Constructor → sets deployer as admin
    constructor() {
        admin = msg.sender;
    }

    // Certificate structure
    struct Certificate {
        bytes32 certHash;
        uint256 issueDate;
        bool isValid;
    }

    // Storage (certificate ID → Certificate)
    mapping(string => Certificate) private certificates;

    // Access control modifier
    modifier onlyAdmin() {
        require(msg.sender == admin, "Not authorized");
        _;
    }

    // 🔹 ISSUE CERTIFICATE
    function issueCertificate(
        string memory _certId,
        string memory _data
    ) public onlyAdmin {

        require(certificates[_certId].issueDate == 0, "Already exists");

        bytes32 hash = keccak256(abi.encodePacked(_data));

        certificates[_certId] = Certificate(
            hash,
            block.timestamp,
            true
        );
    }

    // 🔹 VERIFY CERTIFICATE
    function verifyCertificate(
        string memory _certId,
        string memory _data
    ) public view returns (bool) {

        bytes32 hash = keccak256(abi.encodePacked(_data));

        return certificates[_certId].certHash == hash &&
               certificates[_certId].isValid;
    }

    // 🔹 REVOKE CERTIFICATE
    function revokeCertificate(string memory _certId)
        public onlyAdmin
    {
        require(certificates[_certId].isValid == true, "Already revoked");
        certificates[_certId].isValid = false;
    }

    // 🔹 OPTIONAL: Get certificate details (for frontend)
    function getCertificate(string memory _certId)
        public view returns (bytes32, uint256, bool)
    {
        Certificate memory cert = certificates[_certId];
        return (cert.certHash, cert.issueDate, cert.isValid);
    }
}