create table if not exists virgo_user (
      id Int64,
      name String,
      age Int32,
      gender Int8,
      create_at DateTime,
      update_at DateTime,
      delete_flag Int8
) engine = MergeTree
    PRIMARY KEY id
    ORDER BY id ;