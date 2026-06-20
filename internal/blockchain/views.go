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
	QRCode           []byte
	VerifyURL        string
	CertificateHash  string
	BlockchainTxHash string
	RecipientName    string
	CourseName       string
	Grade            string
	IssuedAt         time.Time
}

type VerifyCertificateResponse struct {
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssueDate        uint64
	IsValid          bool
}

type LoginRequest struct {
	Email    string
	Password string
}
