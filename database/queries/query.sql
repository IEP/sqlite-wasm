-- @formatter:off

-- name: GetNote :one
select * from note where id = ? limit 1;

-- name: ListNotes :many
select *
from note
order by id
limit ?
offset ?;

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
