-- name: BookChapterById :one
SELECT *
FROM book_chapters
WHERE id = $1;

-- name: BookChaptersByBookGroupId :many
SELECT *
FROM book_chapters
WHERE book_group_id = $1
ORDER BY id
OFFSET $2 ROWS FETCH FIRST $3 ROWS ONLY;

-- name: UpdateBookChapter :exec
UPDATE book_chapters
SET chapter_number=$2,
    name=$3,
    text_content=$4
WHERE id = $1;

-- name: InsertBookChapter :one
INSERT INTO book_chapters(chapter_number, name, text_content, type, book_group_id, owner_id)
VALUES (@chapter_number, @name, @text_content, @type, @book_group_id, @owner_id)
RETURNING *;

-- name: DeleteBookChapterById :exec
DELETE
FROM book_chapters
WHERE id = $1;

-- name: DeleteBookChapterByBookGroupId :exec
DELETE
FROM book_chapters
WHERE book_group_id = $1;

-- name: GetBookGroupChapters :many
SELECT book_chapters.chapter_number,
       book_chapters.name,
       book_chapters.id as chapterId,
       book_chapters.date_created,
       u.id             as userId,
       u.user_name,
       coalesce(sum(bcv.count), 0) as totalView
FROM book_chapters
         JOIN book_groups bg on book_chapters.book_group_id = bg.id
         JOIN users u on book_chapters.owner_id = u.id
         JOIN book_chapter_views bcv on book_chapters.id = bcv.book_chapter_id
WHERE bg.id = $1
GROUP BY book_chapters.id, u.id;

-- name: GetBookChapterOwner :one
SELECT users.id, users.user_name
FROM users
         JOIN book_chapters bc on users.id = bc.owner_id
WHERE bc.id = $1;

