SET GLOBAL general_log = 1;
SET GLOBAL general_log_file = "/tmp/mysql_general.log";

CREATE DATABASE playerdata;

USE playerdata;

CREATE TABLE players (
    id INT AUTO_INCREMENT NOT NULL,
    fullname VARCHAR(128) NOT NULL,
    nickname VARCHAR(128) NOT NULL,
    kd DECIMAL(5, 2) NOT NULL,
    team VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO
    players (nickname, fullname, kd, team)
VALUES
    ('silen', 'Jan Kotarba', 1.69, "mongo"),
    ('nico', 'Andrew Null', 1.88, "apple"),
    ('xmal', 'Meg Trust', 3.23, "mongo"),
    ('nicodoz', 'David Bend', 0.23, "apple");