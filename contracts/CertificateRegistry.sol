// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract CertificateRegistry {

    address public owner;

    struct Certificate {
        bytes32 pdfHash;
        string recipientName;
        string courseName;
        string grade;
        string issuingAuthority;
        uint256 issueDate;
        bool isValid;
        bool exists;
    }

    // pdf hash => certificate
    mapping(bytes32 => Certificate) private certificates;

    event CertificateIssued(
        bytes32 indexed pdfHash,
        string recipientName,
        string courseName,
        string issuingAuthority,
        uint256 issueDate
    );

    event CertificateRevoked(
        bytes32 indexed pdfHash,
        uint256 revokeDate
    );

    modifier onlyOwner() {
        require(msg.sender == owner, "Access denied");
        _;
    }

    modifier certificateExists(bytes32 _pdfHash) {
        require(certificates[_pdfHash].exists, "Certificate does not exist");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    // =========================
    // ISSUE CERTIFICATE
    // =========================

    function issueCertificate(
        bytes32 _pdfHash,
        string memory _recipientName,
        string memory _courseName,
        string memory _grade,
        string memory _issuingAuthority
    ) external onlyOwner {

        require(!certificates[_pdfHash].exists, "Certificate already exists");

        certificates[_pdfHash] = Certificate({
            pdfHash: _pdfHash,
            recipientName: _recipientName,
            courseName: _courseName,
            grade: _grade,
            issuingAuthority: _issuingAuthority,
            issueDate: block.timestamp,
            isValid: true,
            exists: true
        });

        emit CertificateIssued(
            _pdfHash,
            _recipientName,
            _courseName,
            _issuingAuthority,
            block.timestamp
        );
    }

    // =========================
    // VERIFY CERTIFICATE
    // =========================

    function verifyCertificate(bytes32 _pdfHash)
        external
        view
        certificateExists(_pdfHash)
        returns (
            string memory recipientName,
            string memory courseName,
            string memory grade,
            string memory issuingAuthority,
            uint256 issueDate,
            bool isValid
        )
    {
        Certificate memory cert = certificates[_pdfHash];

        return (
            cert.recipientName,
            cert.courseName,
            cert.grade,
            cert.issuingAuthority,
            cert.issueDate,
            cert.isValid
        );
    }

    // =========================
    // REVOKE CERTIFICATE
    // =========================

    function revokeCertificate(bytes32 _pdfHash)
        external
        onlyOwner
        certificateExists(_pdfHash)
    {
        require(
            certificates[_pdfHash].isValid,
            "Certificate already revoked"
        );

        certificates[_pdfHash].isValid = false;

        emit CertificateRevoked(
            _pdfHash,
            block.timestamp
        );
    }

    // =========================
    // QUICK VALIDITY CHECK
    // =========================

    function isCertificateValid(bytes32 _pdfHash)
        external
        view
        returns (bool)
    {
        if (!certificates[_pdfHash].exists) {
            return false;
        }

        return certificates[_pdfHash].isValid;
    }
}