-- name: GetLinks :many
select id, url_code, destination from links;

-- name: CreateLink :one
insert into links (url_code, destination) values ($1, $2) returning *;

-- name: FindLinkDestinationByCode :one
select links.destination from links where links.url_code = $1 ;

-- name: DeleteLink :one
delete from links where id = $1 returning *;
