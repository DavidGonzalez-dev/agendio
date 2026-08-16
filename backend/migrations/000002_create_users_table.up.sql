-- 1. Crear el tipo ENUM primero
-- Nota: Postgres no soporta "IF NOT EXISTS" en CREATE TYPE,
-- por lo que se usa un bloque DO para hacerlo de forma segura:
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE role AS ENUM ('admin', 'client');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar NOT NULL,
    email varchar UNIQUE NOT NULL,
    password_hash varchar NOT NULL,
    role role NOT NULL,
    business_id uuid REFERENCES businesses(id) ON DELETE CASCADE,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
);
