DROP KEYSPACE IF EXISTS ex_customer;
CREATE KEYSPACE ex_customer WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

DROP TABLE IF EXISTS ex_customer.customer;
CREATE TABLE ex_customer.customer
(
    id           text PRIMARY KEY,
    name         text,
    identity     text,
    email        text,
    phone        text,
    dob          text,
    gender       text,
);
CREATE INDEX ON ex_customer.customer (identity);
