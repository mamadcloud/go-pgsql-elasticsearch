CREATE TABLE delta_update (
  id serial not null primary key,
  item_id int,
  action varchar(100),
  table_name varchar(150),
  timestamp timestamp
)