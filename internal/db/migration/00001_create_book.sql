-- +goose Up
CREATE TABLE public.book
(
    id    BIGSERIAL PRIMARY KEY,
    title text NOT NULL,
    author text NOT NULL
);

-- +goose Down
DROP TABLE public.book;
