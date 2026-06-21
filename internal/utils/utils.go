package utils

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

type CertificateCard struct {
	RecipientName    string
	CourseName       string
	Grade            string
	IssuingAuthority string
	IssuedAt         string

	CertificateHash string
	TransactionHash string

	VerifyURL      string
	TransactionURL string

	QRCode []byte
}

func GenerateCertificateCard(req CertificateCard) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetTitle(
		"Certificate Issuance Receipt",
		false,
	)

	pdf.SetAuthor(
		"Blockchain Certificate Verification System",
		false,
	)

	pdf.SetMargins(
		15,
		15,
		15,
	)

	pdf.AddPage()

	// title
	pdf.SetFont(
		"Arial",
		"B",
		20,
	)

	pdf.Cell(
		0,
		10,
		"Certificate Issued Successfully",
	)

	pdf.Ln(12)

	// subtitle
	pdf.SetTextColor(
		0,
		128,
		0,
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		0,
		8,
		"Recorded on Ethereum Sepolia",
	)

	pdf.SetTextColor(
		0,
		0,
		0,
	)

	pdf.Ln(12)

	// qr code
	if len(req.QRCode) > 0 {

		opt := gofpdf.ImageOptions{
			ImageType: "PNG",
		}

		pdf.RegisterImageOptionsReader(
			"qr",
			opt,
			bytes.NewReader(
				req.QRCode,
			),
		)

		pdf.ImageOptions(
			"qr",
			155,
			18,
			40,
			40,
			false,
			opt,
			0,
			"",
		)
	}

	// recipient
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		50,
		8,
		"Recipient",
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		0,
		8,
		req.RecipientName,
	)

	pdf.Ln(8)

	// course
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		50,
		8,
		"Course",
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		0,
		8,
		req.CourseName,
	)

	pdf.Ln(8)

	// grade
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		50,
		8,
		"Grade",
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		0,
		8,
		req.Grade,
	)

	pdf.Ln(8)

	// issuer
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		50,
		8,
		"Issued By",
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		0,
		8,
		req.IssuingAuthority,
	)

	pdf.Ln(8)

	// issue date
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		50,
		8,
		"Issued At",
	)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		0,
		8,
		req.IssuedAt,
	)

	pdf.Ln(15)

	// certificate hash
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		0,
		8,
		"Certificate Hash",
	)

	pdf.Ln(8)

	pdf.SetFont(
		"Courier",
		"",
		9,
	)

	pdf.MultiCell(
		0,
		5,
		req.CertificateHash,
		"",
		"",
		false,
	)

	pdf.Ln(5)

	// transaction hash
	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		0,
		8,
		"Blockchain Transaction Hash",
	)

	pdf.Ln(8)

	pdf.SetFont(
		"Courier",
		"",
		9,
	)

	pdf.MultiCell(
		0,
		5,
		req.TransactionHash,
		"",
		"",
		false,
	)

	pdf.Ln(10)

	// verification url
	pdf.SetTextColor(
		0,
		0,
		0,
	)

	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		0,
		8,
		"Verification URL",
	)

	pdf.Ln(8)

	pdf.SetTextColor(
		0,
		0,
		255,
	)

	pdf.WriteLinkString(
		5,
		req.VerifyURL,
		req.VerifyURL,
	)

	pdf.Ln(10)

	// transaction url
	pdf.SetTextColor(
		0,
		0,
		0,
	)

	pdf.SetFont(
		"Arial",
		"B",
		12,
	)

	pdf.Cell(
		0,
		8,
		"Transaction URL",
	)

	pdf.Ln(8)

	pdf.SetTextColor(
		0,
		0,
		255,
	)

	pdf.WriteLinkString(
		5,
		"View on Etherscan",
		req.TransactionURL,
	)

	pdf.SetTextColor(
		0,
		0,
		0,
	)

	var buf bytes.Buffer

	err := pdf.Output(
		&buf,
	)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
