-- Enable UUID extension if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table: ipos
CREATE TABLE IF NOT EXISTS ipos (
    id VARCHAR(64) PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    short_name VARCHAR(100),
    symbol VARCHAR(32),
    isin VARCHAR(32),
    category VARCHAR(32) DEFAULT 'Mainboard',
    ipo_type VARCHAR(32) DEFAULT 'Book Built',
    exchange VARCHAR(32) DEFAULT 'NSE/BSE',
    issue_size VARCHAR(64),
    price_band_low NUMERIC(12,2) DEFAULT 0,
    price_band_high NUMERIC(12,2) DEFAULT 0,
    price_band VARCHAR(64),
    lot_size INTEGER DEFAULT 1,
    minimum_investment VARCHAR(64),
    sector VARCHAR(100),
    open_date TIMESTAMPTZ,
    close_date TIMESTAMPTZ,
    allotment_date TIMESTAMPTZ,
    listing_date TIMESTAMPTZ,
    status VARCHAR(32) NOT NULL,
    gmp NUMERIC(10,2) DEFAULT 0,
    subscription NUMERIC(10,2) DEFAULT 0,
    radar_score NUMERIC(5,2) DEFAULT 0,
    listing_gain NUMERIC(6,2) DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ipos_status ON ipos(status);
CREATE INDEX IF NOT EXISTS idx_ipos_symbol ON ipos(symbol);
CREATE INDEX IF NOT EXISTS idx_ipos_open_date ON ipos(open_date);
CREATE INDEX IF NOT EXISTS idx_ipos_close_date ON ipos(close_date);

-- Table: ipo_subscriptions
CREATE TABLE IF NOT EXISTS ipo_subscriptions (
    id BIGSERIAL PRIMARY KEY,
    ipo_id VARCHAR(64) NOT NULL REFERENCES ipos(id) ON DELETE CASCADE,
    timestamp TIMESTAMPTZ NOT NULL,
    retail NUMERIC(8,2) DEFAULT 0,
    qib NUMERIC(8,2) DEFAULT 0,
    nii NUMERIC(8,2) DEFAULT 0,
    employee NUMERIC(8,2) DEFAULT 0,
    shareholder NUMERIC(8,2) DEFAULT 0,
    total NUMERIC(8,2) DEFAULT 0,
    applications BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_ipo_subscriptions_ipo_id ON ipo_subscriptions(ipo_id);
CREATE INDEX IF NOT EXISTS idx_ipo_subscriptions_timestamp ON ipo_subscriptions(timestamp);

-- Table: gmp_histories
CREATE TABLE IF NOT EXISTS gmp_histories (
    id BIGSERIAL PRIMARY KEY,
    ipo_id VARCHAR(64) NOT NULL REFERENCES ipos(id) ON DELETE CASCADE,
    gmp NUMERIC(10,2) NOT NULL,
    source VARCHAR(100) NOT NULL,
    disclaimer VARCHAR(255) DEFAULT 'Unofficial market indicator',
    recorded_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_gmp_histories_ipo_id ON gmp_histories(ipo_id);
CREATE INDEX IF NOT EXISTS idx_gmp_histories_recorded_at ON gmp_histories(recorded_at);

-- Table: listing_performances
CREATE TABLE IF NOT EXISTS listing_performances (
    id BIGSERIAL PRIMARY KEY,
    ipo_id VARCHAR(64) UNIQUE NOT NULL REFERENCES ipos(id) ON DELETE CASCADE,
    issue_price NUMERIC(12,2) NOT NULL,
    listing_price NUMERIC(12,2) NOT NULL,
    listing_gain_percent NUMERIC(8,2) NOT NULL,
    current_price NUMERIC(12,2),
    current_gain_percent NUMERIC(8,2),
    recorded_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
