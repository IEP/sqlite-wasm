create table note (
    id integer primary key,
    name text not null ,
    content text,
    created_at datetime,
    updated_at datetime
);
