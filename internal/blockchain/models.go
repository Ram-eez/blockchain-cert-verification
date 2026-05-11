package blockchain

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	ID               uuid.UUID
	InstituteID      uuid.UUID
	CertificateHash  string
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	BlockchainTxHash string
	BlockchainStatus string
	IsRevoked        bool
	IssuedAt         time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CreateCertificateParams struct {
	InstituteID      uuid.UUID
	CertificateHash  string
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	BlockchainTxHash string
	IssuedAt         time.Time
}
