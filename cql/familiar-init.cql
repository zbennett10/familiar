CREATE KEYSPACE IF NOT EXISTS  familiar
WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

CREATE TABLE familiar.user (
    id uuid,
    email text,
    join_date timeuuid,
    firstname text,
    lastname text,
    PRIMARY KEY ((id, email), join_date)
) WITH CLUSTERING ORDER BY (join_date ASC);

CREATE TABLE familiar.note (
    id uuid,
    user_id uuid,
    content text,
    tags list<text>,
    create_date timeuuid,
    modified_date timeuuid,
    PRIMARY KEY ((id, user_id), tags, modified_date, create_date, content)
) WITH CLUSTERING ORDER BY (modified_date ASC);

CREATE TABLE familiar.note_attachment (
    id uuid,
    upload_date timeuuid,
    note_id uuid,
    chunk_order int,
    PRIMARY KEY ((note_id, chunk_order), upload_date)
);

CREATE TABLE familiar.note_attachment_data (
    chunk_id uuid,
    chunk blob,
    PRIMARY KEY (chunk_id)
);
