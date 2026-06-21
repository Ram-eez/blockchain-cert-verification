package blockchain

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository interface {
	CreateCertificate(ctx context.Context, params CreateCertificateParams) error
	RevokeCertificate(ctx context.Context, certificateHash string) error
	GetInstituteByEmail(ctx context.Context, email string) (*Institute, error)
	GetCertificateByHash(ctx context.Context, hash string) (*Certificate, error)
}

type projectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) ProjectRepository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) CreateCertificate(ctx context.Context, params CreateCertificateParams) error {
	query := `INSERT INTO certificates(institute_id, certificate_hash, recipient_name, course_name, grade, issuing_authority, blockchain_tx_hash, issued_at)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.Exec(ctx, query, params.InstituteID, params.CertificateHash, params.RecipientName, params.CourseName, params.Grade, params.IssuingAuthority, params.BlockchainTxHash, params.IssuedAt)

	return err
}

func (r *projectRepository) RevokeCertificate(ctx context.Context, certificateHash string) error {
	query := `
		UPDATE certificates SET is_revoked = true, blockchain_status = 'revoked', updated_at = now()
		WHERE certificate_hash = $1`

	_, err := r.db.Exec(ctx, query, certificateHash)

	return err
}

func (r *projectRepository) GetInstituteByEmail(ctx context.Context, email string) (*Institute, error) {
	query := `SELECT id,name,email,password_hash,is_active,created_at,updated_at FROM institutes WHERE email=$1`

	var institute Institute

	err := r.db.QueryRow(ctx, query, email).Scan(
		&institute.ID,
		&institute.Name,
		&institute.Email,
		&institute.PasswordHash,
		&institute.IsActive,
		&institute.CreatedAt,
		&institute.UpdatedAt,
	)

	return &institute, err
}
func (r *projectRepository) GetCertificateByHash(ctx context.Context, hash string) (*Certificate, error) {
	query := `SELECT id, institute_id, certificate_hash, recipient_name, course_name, grade, issuing_authority, blockchain_tx_hash, blockchain_status, is_revoked, issued_at, created_at, updated_at FROM certificates WHERE certificate_hash = $1`

	var cert Certificate

	err := r.db.QueryRow(ctx, query, hash).Scan(&cert.ID, &cert.InstituteID, &cert.CertificateHash, &cert.RecipientName, &cert.CourseName, &cert.Grade, &cert.IssuingAuthority, &cert.BlockchainTxHash, &cert.BlockchainStatus, &cert.IsRevoked, &cert.IssuedAt, &cert.CreatedAt, &cert.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &cert, nil
}
