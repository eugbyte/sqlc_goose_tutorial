-- +goose Up
ALTER TABLE public.book
    ADD COLUMN blurb text;

-- +goose Down
ALTER TABLE public.book
    DROP COLUMN blurb;
