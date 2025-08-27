-- =============================================================================
-- library_go - esquema inicial
-- PostgreSQL 14+
-- =============================================================================

-- Opcional: crear un esquema dedicado
CREATE SCHEMA IF NOT EXISTS library;
SET search_path TO library, public;

-- =============================================================================
-- Tabla: authors
-- =============================================================================
CREATE TABLE IF NOT EXISTS authors (
    id            UUID PRIMARY KEY,
    first_name    TEXT NOT NULL,
    last_name     TEXT NOT NULL,
    bio           TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_authors_last_name ON authors (last_name);

-- Trigger simple para updated_at (aplicable a varias tablas)
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
 END;
$$ LANGUAGE plpgsql;

-- =============================================================================
-- Tabla: books
-- =============================================================================
CREATE TABLE IF NOT EXISTS books (
    id            	UUID PRIMARY KEY,
    title           TEXT NOT NULL,
    isbn            TEXT UNIQUE,         -- opcional       
    description     TEXT,
    published_year  INT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_books_title ON books USING GIN (to_tsvector('simple', title));

CREATE TRIGGER trg_books_updated_at
BEFORE UPDATE ON books
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- =============================================================================
-- Tabla pivote: book_authors (N:N)
-- =============================================================================
CREATE TABLE IF NOT EXISTS book_authors (
    book_id   UUID NOT NULL,
    author_id UUID NOT NULL,
    CONSTRAINT pk_book_authors PRIMARY KEY (book_id, author_id),
    CONSTRAINT fk_book_authors_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    CONSTRAINT fk_book_authors_author FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);

-- =============================================================================
-- Tabla: members (socios)
-- =============================================================================
CREATE TABLE IF NOT EXISTS members (
    id          UUID PRIMARY KEY,
    full_name   TEXT NOT NULL,
    email       TEXT UNIQUE,
    phone       TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_members_name ON members USING GIN (to_tsvector('simple', full_name));

CREATE TRIGGER trg_members_updated_at
BEFORE UPDATE ON members
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- =============================================================================
-- Tabla: book_copies (copias físicas)
-- =============================================================================
CREATE TABLE IF NOT EXISTS book_copies (
    id          UUID PRIMARY KEY,
    book_id     UUID NOT NULL,
    barcode     TEXT UNIQUE,             -- opcional: código/etiqueta de la copia        
    condition   TEXT,                    -- estado físico
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,  -- copia habilitada para préstamo
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_book_copies_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_copies_book_id ON book_copies (book_id);

-- =============================================================================
-- Tabla: loans (préstamos)
-- Regla clave: solo un préstamo activo por copia (retornada_at IS NULL)
-- =============================================================================
CREATE TABLE IF NOT EXISTS loans (
	id           UUID PRIMARY KEY,
    member_id    UUID NOT NULL,
    copy_id      UUID NOT NULL,
    loaned_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    due_date     DATE NOT NULL,
    returned_at  TIMESTAMPTZ,           -- NULL = aún prestado
    fine_cents   INT NOT NULL DEFAULT 0, -- multa calculada al devolver (si aplica)
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_loans_member FOREIGN KEY (member_id) REFERENCES members(id) ON DELETE RESTRICT,
    CONSTRAINT fk_loans_copy FOREIGN KEY (copy_id) REFERENCES book_copies(id) ON DELETE RESTRICT
    --CONSTRAINT chk_due_date_future CHECK (due_date >= (loaned_at AT TIME ZONE 'UTC')::date),
    --CONSTRAINT chk_fine_non_negative CHECK (fine_cents >= 0)
);

CREATE INDEX IF NOT EXISTS idx_loans_member_id ON loans (member_id);
CREATE INDEX IF NOT EXISTS idx_loans_copy_id ON loans (copy_id);
CREATE INDEX IF NOT EXISTS idx_loans_active ON loans (copy_id) WHERE returned_at IS NULL;

-- Único préstamo activo por copia:
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_indexes
    WHERE schemaname = 'library' AND indexname = 'uniq_active_loan_per_copy'
  ) THEN
    CREATE UNIQUE INDEX uniq_active_loan_per_copy
    ON loans (copy_id)
    WHERE returned_at IS NULL;
  END IF;
END$$;

-- =============================================================================
-- Tabla: reservations (reservas por libro)
-- Regla clave: una reserva activa por miembro/libro (cancelada/atendida = no activa)
-- =============================================================================
CREATE TABLE IF NOT EXISTS reservations (
    id               UUID PRIMARY KEY,
    member_id          UUID NOT NULL,
    book_id            UUID NOT NULL,
    reserved_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    canceled_at        TIMESTAMPTZ,        -- NULL = activa
    fulfilled_loan_id  UUID, 			-- si se convirtió en préstamo
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_reservations_member FOREIGN KEY (member_id) REFERENCES members(id) ON DELETE RESTRICT,
    CONSTRAINT fk_reservations_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE RESTRICT,
    CONSTRAINT fk_reservations_fulfilled_loan FOREIGN KEY (fulfilled_loan_id) REFERENCES loans(id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_reservations_book_id ON reservations (book_id);
CREATE INDEX IF NOT EXISTS idx_reservations_member_id ON reservations (member_id);

-- Única reserva ACTIVA por miembro/libro:
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_indexes
    WHERE schemaname = 'library' AND indexname = 'uniq_active_reservation_per_member_book'
  ) THEN
    CREATE UNIQUE INDEX uniq_active_reservation_per_member_book
    ON reservations (member_id, book_id)
    WHERE canceled_at IS NULL AND fulfilled_loan_id IS NULL;
  END IF;
END$$;

-- =============================================================================
-- Reglas/Checks adicionales
-- =============================================================================
ALTER TABLE loans
  ADD CONSTRAINT chk_due_date_future CHECK (due_date >= (loaned_at AT TIME ZONE 'UTC')::date);

ALTER TABLE loans
  ADD CONSTRAINT chk_fine_non_negative CHECK (fine_cents >= 0);
  
-- limpiar la dba

-- Rollback del esquema inicial
/*
	SET search_path TO library, public;

	DROP INDEX IF EXISTS uniq_active_reservation_per_member_book;
	DROP INDEX IF EXISTS uniq_active_loan_per_copy;
	
	DROP TABLE IF EXISTS reservations;
	DROP TABLE IF EXISTS loans;
	DROP TABLE IF EXISTS book_copies;
	DROP TABLE IF EXISTS book_authors;
	DROP TABLE IF EXISTS members;
	DROP TABLE IF EXISTS books;
	DROP TABLE IF EXISTS authors;

-- Función utilitaria
	DROP FUNCTION IF EXISTS set_updated_at();

-- (Opcional) eliminar el esquema
-- DROP SCHEMA IF EXISTS library CASCADE;
*/
