CREATE KEYSPACE oauth WITH REPLICATION = {'class': 'SimpleStrategy','replication_factor':1}

describe KEYSPACES;

USE oauth;

describle tables;

CREATE TABLE access_tokens( access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);

