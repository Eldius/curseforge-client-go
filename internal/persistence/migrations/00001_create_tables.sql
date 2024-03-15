-- +migrate Up

CREATE TABLE IF NOT EXISTS mod_info (
    id NUMBER PRIMARY KEY
    , name VARCHAR(255)
    , url VARCHAR(255)
    , source_url VARCHAR(255)
    , wiki_url VARCHAR(255)
    , description TEXT
    , class_id NUMBER
    , game_id NUMBER
    , versions JSON
    , categories JSON
    , authors JSON
);

CREATE TABLE IF NOT EXISTS mod_category (
    id NUMBER PRIMARY KEY
    , class_id NUMBER
    , name VARCHAR(255)
    , url VARCHAR(255)
    , icon_url VARCHAR(255)
    , last_modified DATETIME
    , is_class bool
    , game_id NUMBER
    , parent_id NUMBER
);

-- +migrate Down

drop table IF EXISTS mod_info;
drop table IF EXISTS mod_category;
