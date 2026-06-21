package blockchain

import (
	"time"

	"github.com/google/uuid"
)

type IssueCertificateRequest struct {
	InstituteID      uuid.UUID
	PdfHash          [32]byte
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
}

type IssueCertificateResponse struct {
	QRCodeBase64     string
	PDFBase64        string
	VerifyURL        string
	CertificateHash  string
	BlockchainTxHash string
	TransactionURL   string
	RecipientName    string
	CourseName       string
	Grade            string
	IssuedAt         time.Time
	IssuingAuthority string
}

type VerifyCertificateResponse struct {
	Exists           bool
	IsValid          bool
	IsRevoked        bool
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssueDate        uint64
	IssuedAt         time.Time
	CertificateHash  string
	BlockchainTxHash string
	VerifyURL        string
	TransactionURL   string
}

type LoginRequest struct {
	Email    string
	Password string
}

// admin
type CreateInstitutionRequest struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type CreateInstitutionResponse struct {
	Name     string
	Email    string
	Password string
}
