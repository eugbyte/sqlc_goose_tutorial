-- name: GetBook :one
SELECT *
FROM public.book
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT *
FROM public.book
WHERE (sqlc.narg('title')::text IS NULL OR title ILIKE '%' || sqlc.narg('title')::text || '%')
  AND (sqlc.narg('author')::text IS NULL OR author ILIKE '%' || sqlc.narg('author')::text || '%')
ORDER BY title;

-- name: CreateBook :one
INSERT INTO public.book (title, author)
VALUES ($1, $2) RETURNING *;

-- name: UpdateBook :exec
UPDATE public.book
set title  = $2,
    author = $3
WHERE id = $1;

-- name: DeleteBook :exec
DELETE
FROM public.book
WHERE id = $1;

-- name: DeleteAllBooks :execrows
DELETE
FROM public.book;