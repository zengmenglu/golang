
CREATE DATABASE if not exists gorm_example;

USE gorm_example;

CREATE TABLE kv_tbl(
    k VARCHAR(1024) NOT NULL,
    v VARCHAR(1024) NOT NULL,
    PRIMARY KEY (k))ENGINE=InnoDB DEFAULT CHARSET=utf8;