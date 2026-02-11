-- +goose Up

create table feed_follows (
    id uuid primary key,
    feed_id UUID not null references feeds(id) on delete cascade,
    user_id UUID not null references users(id) on delete cascade,
    unique(feed_id, user_id)
);


-- +goose Down
drop table feed_follows;
