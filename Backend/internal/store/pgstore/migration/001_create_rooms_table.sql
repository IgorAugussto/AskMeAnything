-- Write your migrate up statements here

CREATE TABLE IF NOT EXIST rooms(
  "id"     uuid           PRIMARY KEY  NOT NULL  DEFAULT  gen_ramdom_uuid(),
  "theme"  VARCHAR(255)                NOT NULL  
)

---- create above / drop below ----

DROP TABLE IF EXIST rooms;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
