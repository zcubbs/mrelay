CREATE TABLE IF NOT EXISTS email (
  id INTEGER PRIMARY KEY AUTOINCREMENT, -- For SQLite compatibility
  account VARCHAR(255),
  from_address TEXT NOT NULL,
  to_address TEXT NOT NULL,
  subject TEXT NOT NULL,
  body TEXT NOT NULL,
  -- TZ Timestamp
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- TIMESTAMP without TZ for compatibility
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  sent_at TIMESTAMP,
  errors TEXT
);
