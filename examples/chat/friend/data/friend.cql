DROP KEYSPACE IF EXISTS ex_friend;
CREATE KEYSPACE ex WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS ex_friend.account;
CREATE TABLE ex_friend.account
(
);

DROP TABLE IF EXISTS ex_friend.has_friend;
CREATE TABLE ex_friend.has_friend
(
    user_id        bigint,
    friend_id      bigint,
    PRIMARY KEY(user_id, friend_id)
);
