DROP KEYSPACE IF EXISTS chat_friend;
CREATE KEYSPACE chat_friend WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS chat_friend.account;
CREATE TABLE chat_friend.account
(
    id              bigint,
    email           text,
    name            text,
    avatar          text,
);

DROP TABLE IF EXISTS chat_friend.has_friend;
CREATE TABLE chat_friend.has_friend
(
    account_id      bigint,
    friend_id       bigint,
    PRIMARY KEY(user_id, friend_id)
);
