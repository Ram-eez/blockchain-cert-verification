-- =========================================
-- ENABLE EXTENSIONS
-- =========================================

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- =========================================
-- ENUMS
-- =========================================

CREATE TYPE blockchain_status AS ENUM (
    'pending',
    'confirmed',
    'failed',
    'revoked'
);

-- =========================================
-- INSTITUTES
-- =========================================

CREATE TABLE institutes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,

    email TEXT NOT NULL UNIQUE,

    password_hash TEXT NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT true,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_institutes_email
ON institutes(email);

-- =========================================
-- CERTIFICATES
-- =========================================

CREATE TABLE certificates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    institute_id UUID NOT NULL
        REFERENCES institutes(id)
        ON DELETE CASCADE,

    certificate_hash VARCHAR(66) NOT NULL UNIQUE,

    recipient_name TEXT NOT NULL,

    course_name TEXT NOT NULL,

    grade TEXT NOT NULL,

    issuing_authority TEXT NOT NULL,

    blockchain_tx_hash VARCHAR(66) NOT NULL,

    blockchain_status blockchain_status NOT NULL DEFAULT 'pending',

    is_revoked BOOLEAN NOT NULL DEFAULT false,

    issued_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- =========================================
-- INDEXES
-- =========================================

CREATE UNIQUE INDEX idx_certificates_hash
ON certificates(certificate_hash);

CREATE INDEX idx_certificates_institute_id
ON certificates(institute_id);

CREATE INDEX idx_certificates_tx_hash
ON certificates(blockchain_tx_hash);

CREATE INDEX idx_certificates_status
ON certificates(blockchain_status);