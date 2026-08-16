CREATE TABLE IF NOT EXISTS businesses (
  id uuid DEFAULT gen_random_uuid (),
  name varchar NOT NULL,
  slug varchar UNIQUE NOT NULL,
  timezone varchar NOT NULL,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
