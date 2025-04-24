-- +goose Up
-- +goose StatementBegin
BEGIN TRANSACTION;

-- Insert roles
INSERT INTO
    ms_role (name, slug)
VALUES
    ('Admin', 'admin'),
    ('Officer', 'officer'),
    ('Buyer', 'buyer');

-- Insert permissions
INSERT INTO
    ms_permission (name)
VALUES
    ('admin_access'),
    ('officer_access'),
    ('buyer_access');

-- Assign permissions to roles
INSERT INTO
    piv_role_permission (ms_role_id, ms_permission_id)
VALUES
    (1, 1),
    (2, 2),
    (3, 3);

INSERT INTO
    ms_user (username, employee_id, password)
VALUES
    (
        '1234567',
        1,
        '$2a$12$5AUsEjCFAA19zIgZCrBQO.Ut2mSW5poQoSe8OtQm2SJ0r2IbfQoeK' --1234567
    );

-- Assign roles to users
INSERT INTO
    piv_user_role (ms_user_id, ms_role_id)
VALUES
    (1, 2);

COMMIT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
BEGIN TRANSACTION;

DELETE FROM piv_role_permission;

DELETE FROM piv_user_role;

DELETE FROM ms_permission;

DELETE FROM ms_role;

DELETE FROM ms_user;

COMMIT;

-- +goose StatementEnd
