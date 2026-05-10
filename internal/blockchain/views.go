package blockchain

type IssueCertificateRequest struct {
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
