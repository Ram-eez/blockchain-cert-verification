// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract CertificateRegistry {

    address public owner;
    mapping(address => bool) public authorizedIssuers;

    struct Certificate {
        string  recipientName;
        string  courseName;
        string  grade;
        uint256 issueDate;
        address issuedBy;
        bool    isValid;
        bool    exists;
    }

    mapping(bytes32 => Certificate) private certificates;

    event IssuerAdded(address indexed issuer);
    event IssuerRemoved(address indexed issuer);
    event CertificateIssued(bytes32 indexed certificateId, string recipientName, string courseName, address indexed issuedBy, uint256 issueDate);
    event CertificateRevoked(bytes32 indexed certificateId, address indexed revokedBy, uint256 revokeDate);

    modifier onlyOwner() {
        require(msg.sender == owner, "Access denied: Not owner");
        _;
    }

    modifier onlyAuthorized() {
        require(authorizedIssuers[msg.sender] || msg.sender == owner, "Access denied: Not an authorized issuer");
        _;
    }

    modifier certificateMustExist(bytes32 _id) {
        require(certificates[_id].exists, "Certificate does not exist");
        _;
    }

    constructor() {
        owner = msg.sender;
        authorizedIssuers[msg.sender] = true;
    }

    function addIssuer(address _issuer) external onlyOwner {
        require(_issuer != address(0), "Invalid address");
        require(!authorizedIssuers[_issuer], "Already an authorized issuer");
        authorizedIssuers[_issuer] = true;
        emit IssuerAdded(_issuer);
    }

    function removeIssuer(address _issuer) external onlyOwner {
        require(_issuer != owner, "Cannot remove owner");
        require(authorizedIssuers[_issuer], "Not an authorized issuer");
        authorizedIssuers[_issuer] = false;
        emit IssuerRemoved(_issuer);
    }

    function issueCertificate(
        string memory _recipientName,
        string memory _courseName,
        string memory _grade,
        string memory _uniqueSeed
    ) external onlyAuthorized returns (bytes32) {
        bytes32 certId = generateCertificateId(_uniqueSeed);
        require(!certificates[certId].exists, "Certificate ID already exists");
        certificates[certId] = Certificate({
            recipientName : _recipientName,
            courseName    : _courseName,
            grade         : _grade,
            issueDate     : block.timestamp,
            issuedBy      : msg.sender,
            isValid       : true,
            exists        : true
        });
        emit CertificateIssued(certId, _recipientName, _courseName, msg.sender, block.timestamp);
        return certId;
    }

    function revokeCertificate(bytes32 _certificateId)
        external
        onlyAuthorized
        certificateMustExist(_certificateId)
    {
        require(certificates[_certificateId].isValid, "Certificate already revoked");
        certificates[_certificateId].isValid = false;
        emit CertificateRevoked(_certificateId, msg.sender, block.timestamp);
    }

    function verifyCertificate(bytes32 _certificateId)
        external
        view
        certificateMustExist(_certificateId)
        returns (
            string memory recipientName,
            string memory courseName,
            string memory grade,
            uint256 issueDate,
            address issuedBy,
            bool isValid
        )
    {
        Certificate memory cert = certificates[_certificateId];
        return (cert.recipientName, cert.courseName, cert.grade, cert.issueDate, cert.issuedBy, cert.isValid);
    }

    function isCertificateValid(bytes32 _certificateId) external view returns (bool) {
        if (!certificates[_certificateId].exists) return false;
        return certificates[_certificateId].isValid;
    }

    function isAuthorized(address _addr) external view returns (bool) {
        return authorizedIssuers[_addr];
    }

    function generateCertificateId(string memory _seed) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_seed));
    }
}