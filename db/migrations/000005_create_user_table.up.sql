CREATE TABLE users
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email      text
        constraint users_email_checker_lowercase_ck
            check (email = lower(email)),
    password   text,
    identity   text,
    name       text
)