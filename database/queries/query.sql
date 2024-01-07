-- @formatter:off

-- name: GetNote :one
select * from note where id = ? limit 1;

-- name: ListNotes :many
select *
from note
where
    coalesce(cast(@filter as text), '') = '' or
    (name like @filter or content like @filter)
order by id
limit @limit
offset @offset;

-- name: CreateNote :one
insert into note (
  name, content, created_at, updated_at
) values (
  ?, ?, current_timestamp, current_timestamp
)
returning *;

-- name: UpdateNote :one
update note
set
  name = ?,
  content = ?,
  updated_at = current_timestamp
where id = ?
returning *;

-- name: DeleteNote :exec
delete from note where id = ?;
