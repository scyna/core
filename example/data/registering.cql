DROP KEYSPACE IF EXISTS registering;
CREATE KEYSPACE registering WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS registering.registration;
CREATE TABLE registering.registration
(
    email        text PRIMARY KEY,
    id           bigint,
    name         text,
    password     text,
    otp          text,
    expired      timestamp,
    email_count  counter
);
CREATE INDEX ON registering.registration(id);

DROP TABLE IF EXISTS registering.completed;
CREATE TABLE registering.completed
(
    email        text PRIMARY KEY,
    state        text /*ACTIVE, BLOCKED, DELETED*/
);