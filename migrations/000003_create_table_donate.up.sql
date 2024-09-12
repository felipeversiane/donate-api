CREATE TABLE IF NOT EXISTS donations (
    id UUID PRIMARY KEY,
    amount NUMERIC(10, 2) NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    payment_method VARCHAR(50) NOT NULL,
    installment_number INT NOT NULL DEFAULT 0,
    donation_type VARCHAR(50) NOT NULL,
    period INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
    CONSTRAINT chk_donation_installments CHECK (
        (donation_type = 'one_time' AND installment_number = 0) OR
        (donation_type = 'recurring' AND payment_method = 'credit_card' AND installment_number > 0)
    ),
    CONSTRAINT chk_donation_period CHECK (
        (donation_type = 'one_time' AND period = 0) OR
        (donation_type = 'recurring' AND period > 0)
    )
);