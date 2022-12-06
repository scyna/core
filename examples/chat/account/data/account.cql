DROP KEYSPACE IF EXISTS msg_account;
CREATE KEYSPACE msg_account WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS msg_account.account;
CREATE TABLE msg_account.user
(
    id           bigint PRIMARY KEY,
    email        text,
    name         text,
    password     text,
);
CREATE INDEX ON msg_account.account (email);