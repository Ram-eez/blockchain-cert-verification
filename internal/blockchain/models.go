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

type Institute struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type CreateInstituteModel struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
}
