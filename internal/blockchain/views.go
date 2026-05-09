package blockchain

type IssueCertificateRequest struct {
	RecipientName string
	CourseName    string
	Grade         string
	UniqueSeed    string
}

type VerifyCertificateResponse struct {
	RecipientName string
	CourseName    string
	Grade         string
	IssueDate     uint64
	IssuedBy      string
	IsValid       bool
}
