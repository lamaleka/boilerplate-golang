-- +goose Up
-- +goose StatementBegin
CREATE TABLE ms_role (
    id INT IDENTITY (1, 1) PRIMARY KEY,
    uuid UNIQUEIDENTIFIER DEFAULT NEWSEQUENTIALID () UNIQUE NOT NULL,
    slug NVARCHAR (50) UNIQUE NOT NULL,
    name NVARCHAR (50) UNIQUE NOT NULL,
    created_at DATETIME NOT NULL DEFAULT GETDATE (),
    updated_at DATETIME,
    deleted_at DATETIME,
);

CREATE TABLE ms_permission (
    id INT IDENTITY (1, 1) PRIMARY KEY,
    uuid UNIQUEIDENTIFIER DEFAULT NEWSEQUENTIALID () UNIQUE NOT NULL,
    name NVARCHAR (50) UNIQUE NOT NULL,
    created_at DATETIME NOT NULL DEFAULT GETDATE (),
    updated_at DATETIME,
    deleted_at DATETIME,
);

CREATE TABLE ms_user (
    id INT IDENTITY (1, 1) PRIMARY KEY,
    uuid UNIQUEIDENTIFIER DEFAULT NEWSEQUENTIALID () UNIQUE NOT NULL,
    username CHAR(7) UNIQUE NOT NULL,
    employee_id INT UNIQUE NOT NULL,
    password NVARCHAR (255) NOT NULL,
    is_active BIT NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT GETDATE (),
    updated_at DATETIME,
    deleted_at DATETIME,
    CONSTRAINT fk_ms_user_employee FOREIGN KEY (employee_id) REFERENCES ms_employee (id),
);

CREATE TABLE piv_user_role (
    ms_user_id INT NOT NULL,
    ms_role_id INT NOT NULL,
    PRIMARY KEY (ms_user_id, ms_role_id),
    FOREIGN KEY (ms_user_id) REFERENCES ms_user (id) ON DELETE CASCADE,
    FOREIGN KEY (ms_role_id) REFERENCES ms_role (id) ON DELETE CASCADE
);

CREATE TABLE piv_role_permission (
    ms_role_id INT NOT NULL,
    ms_permission_id INT NOT NULL,
    PRIMARY KEY (ms_role_id, ms_permission_id),
    FOREIGN KEY (ms_role_id) REFERENCES ms_role (id) ON DELETE CASCADE,
    FOREIGN KEY (ms_permission_id) REFERENCES ms_permission (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE piv_role_permission;

DROP TABLE piv_user_role;

DROP TABLE ms_permission;

DROP TABLE ms_role;

DROP TABLE ms_user;

-- +goose StatementEnd
