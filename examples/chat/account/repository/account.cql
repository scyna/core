DROP KEYSPACE IF EXISTS chat_account;
CREATE KEYSPACE chat_account WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS chat_account.account;
CREATE TABLE chat_account.account
(
    id           bigint PRIMARY KEY,
    email        text,
    name         text,
    password     text,
    avatar       text,
    dob          text,
    gender       smallint,
    tel          text,
);
CREATE INDEX ON chat_account.account (email);

DROP TABLE IF EXISTS chat_account.event_store;
CREATE TABLE chat_account.event_store
(
    event_id        bigint,
    entity_id       bigint,
    channel         text,
    data            blob,
    PRIMARY KEY(event_id)
);
