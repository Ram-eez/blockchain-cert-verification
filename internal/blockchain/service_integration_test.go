package blockchain

import (
	"blockchain/internal/config"
	"context"
	"crypto/rand"
	"testing"
	"time"
)

func TestIssueCertificateIntegration(t *testing.T) {
	svc := newTestService(t)

	req := IssueCertificateRequest{
		PdfHash:          randomHash(t),
		RecipientName:    "Alice Example",
		CourseName:       "Blockchain 101",
		Grade:            "A",
		IssuingAuthority: "Example University",
	}

	if err := svc.IssueCertificate(context.Background(), req); err != nil {
		t.Fatalf("issue certificate failed: %v", err)
	}

	resp := waitForVerify(t, svc, req.PdfHash, boolPtr(true))
	if resp.RecipientName != req.RecipientName {
		t.Fatalf("recipient mismatch: got %q want %q", resp.RecipientName, req.RecipientName)
	}
	if resp.CourseName != req.CourseName {
		t.Fatalf("course mismatch: got %q want %q", resp.CourseName, req.CourseName)
	}
	if resp.Grade != req.Grade {
		t.Fatalf("grade mismatch: got %q want %q", resp.Grade, req.Grade)
	}
	if resp.IssuingAuthority != req.IssuingAuthority {
		t.Fatalf("issuing authority mismatch: got %q want %q", resp.IssuingAuthority, req.IssuingAuthority)
	}
	if resp.IssueDate == 0 {
		t.Fatalf("expected non-zero issue date")
	}
}

func TestVerifyCertificateIntegration(t *testing.T) {
	svc := newTestService(t)

	req := IssueCertificateRequest{
		PdfHash:          randomHash(t),
		RecipientName:    "Bob Example",
		CourseName:       "Distributed Systems",
		Grade:            "B+",
		IssuingAuthority: "Example University",
	}

	if err := svc.IssueCertificate(context.Background(), req); err != nil {
		t.Fatalf("issue certificate failed: %v", err)
	}

	resp := waitForVerify(t, svc, req.PdfHash, boolPtr(true))
	if !resp.IsValid {
		t.Fatalf("expected certificate to be valid")
	}
}

func TestRevokeCertificateIntegration(t *testing.T) {
	svc := newTestService(t)

	req := IssueCertificateRequest{
		PdfHash:          randomHash(t),
		RecipientName:    "Carol Example",
		CourseName:       "Smart Contracts",
		Grade:            "A-",
		IssuingAuthority: "Example University",
	}

	if err := svc.IssueCertificate(context.Background(), req); err != nil {
		t.Fatalf("issue certificate failed: %v", err)
	}

	_ = waitForVerify(t, svc, req.PdfHash, boolPtr(true))

	if err := svc.RevokeCertificate(context.Background(), req.PdfHash); err != nil {
		t.Fatalf("revoke certificate failed: %v", err)
	}

	resp := waitForVerify(t, svc, req.PdfHash, boolPtr(false))
	if resp.IsValid {
		t.Fatalf("expected certificate to be revoked")
	}
}

func newTestService(t *testing.T) BlockChainService {
	t.Helper()

	cfg := config.LoadConfig()
	return NewBlockChainService(cfg, nil)
}

func randomHash(t *testing.T) [32]byte {
	t.Helper()

	var hash [32]byte
	if _, err := rand.Read(hash[:]); err != nil {
		t.Fatalf("failed to generate hash: %v", err)
	}

	return hash
}

func waitForVerify(t *testing.T, svc BlockChainService, certificateID [32]byte, wantValid *bool) *VerifyCertificateResponse {
	t.Helper()

	timeout := 2 * time.Minute
	if deadline, ok := t.Deadline(); ok {
		remaining := time.Until(deadline) - 5*time.Second
		if remaining <= 0 {
			t.Skip("test timeout too short for chain confirmation; increase go test -timeout")
		}
		timeout = remaining
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	var lastErr error
	for {
		select {
		case <-ctx.Done():
			if lastErr != nil {
				t.Fatalf("timeout waiting for verify: %v (try go test -timeout 2m)", lastErr)
			}
			t.Fatalf("timeout waiting for verify (try go test -timeout 2m)")
		case <-ticker.C:
			resp, err := svc.VerifyCertificate(context.Background(), certificateID)
			if err != nil {
				lastErr = err
				continue
			}
			if wantValid != nil && resp.IsValid != *wantValid {
				lastErr = nil
				continue
			}
			return resp
		}
	}
}

func boolPtr(v bool) *bool {
	return &v
}
