CREATE TABLE folders (
  id SERIAL,
  parent_id INT,
  name VARCHAR(60) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified_at TIMESTAMP NOT NULL,
  deleted BOOL NOT NULL DEFAULT FALSE,
  PRIMARY KEY(id),
  CONSTRAINT fk_parent
    FOREIGN KEY(parent_id)
      REFERENCES folders(id)
)