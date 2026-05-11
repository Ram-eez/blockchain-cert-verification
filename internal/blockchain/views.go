package blockchain

import "github.com/google/uuid"

type IssueCertificateRequest struct {
	InstituteID      uuid.UUID
	PdfHash          [32]byte
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
}

type VerifyCertificateResponse struct {
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssueDate        uint64
	IsValid          bool
}
