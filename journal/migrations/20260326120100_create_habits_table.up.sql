CREATE TABLE habits
(
    uuid       UUID PRIMARY KEY,
    user_uuid  UUID        NOT NULL,
    icon       INTEGER     NOT NULL,
    name       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
