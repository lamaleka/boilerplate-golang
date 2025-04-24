-- +goose Up
-- +goose StatementBegin
CREATE TABLE ms_employee (
    id INT IDENTITY (1, 1) PRIMARY KEY,
    uuid UNIQUEIDENTIFIER DEFAULT NEWSEQUENTIALID () UNIQUE NOT NULL,
    user_id INT NULL,
    nama VARCHAR(100) NOT NULL CHECK (nama <> ''),
    badge CHAR(7) NOT NULL UNIQUE CHECK (badge <> ''),
    dept_id VARCHAR(12) CHECK (
        dept_id IS NULL
        OR dept_id <> ''
    ),
    dept_title VARCHAR(255) NOT NULL CHECK (dept_title <> ''),
    email VARCHAR(255) CHECK (
        email IS NULL
        OR email <> ''
    ),
    pos_id VARCHAR(8) CHECK (
        pos_id IS NULL
        OR pos_id <> ''
    ),
    pos_title VARCHAR(255) NOT NULL CHECK (pos_title <> ''),
    employee_type TINYINT DEFAULT 1 NOT NULL CHECK (employee_type IN (1, 2, 3)),
    is_active BIT NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT GETDATE (),
    updated_at DATETIME,
    deleted_at DATETIME
);

CREATE INDEX idx_ms_employee_badge ON ms_employee (badge);

CREATE INDEX idx_ms_employee_created_at ON ms_employee (created_at DESC);

CREATE UNIQUE INDEX idx_ms_employee_user_id ON ms_employee (user_id)
WHERE
    user_id IS NOT NULL;

BEGIN TRANSACTION;

INSERT INTO
    ms_employee (
        user_id,
        nama,
        badge,
        dept_id,
        dept_title,
        email,
        pos_id,
        pos_title,
        employee_type
    )
VALUES
    (
        1,
        'Lorem',
        '1234567',
        '123456712345',
        'Lorem Department',
        'lorem@gmail.com',
        '12345678',
        'Lorem Position',
        1
    );

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ms_employee;

-- +goose StatementEnd
