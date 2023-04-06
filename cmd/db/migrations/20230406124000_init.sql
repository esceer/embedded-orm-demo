-- migrate:up
CREATE TABLE entities (
  id varchar(128) PRIMARY KEY,
  description varchar(255) not null
);

-- migrate:down
DROP TABLE IF EXISTS entities;
