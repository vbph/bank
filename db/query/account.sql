-- name: CreateAccount :one
insert into accounts (email, password, balance)
values ($1, $2, $3)
returning *;

-- name: ReadAccount :one
select * from accounts
where email = $1
limit 1;

-- name: ChangePassword :one
update accounts
set password = $2
where id = $1
returning *;

-- name: UpdateBalance :one
update accounts
set balance = balance + sqlc.arg(amount)
where id = sqlc.arg(id)
returning *;

-- name: DeleteAccount :one
delete from accounts
where id = $1
returning *;
