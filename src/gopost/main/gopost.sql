/* sqlite dialect */
create table posts (
    title varchar(500) not null default 'New Post',
    slug varchar(255),
    post_date text,
    body text
);
