DROP KEYSPACE IF EXISTS chat_account;
CREATE KEYSPACE msg_account WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS chat_account.account;
CREATE TABLE chat_account.account
(
    id           bigint PRIMARY KEY,
    email        text,
    name         text,
    password     text,
);
CREATE INDEX ON chat_account.account (email);