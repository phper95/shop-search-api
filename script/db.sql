CREATE DATABASE shop CHARACTER SET utf8 COLLATE utf8_general_ci;
use shop;
CREATE TABLE product(
                        id int auto_increment primary key,
                        name varchar(120) NOT NULL default '',
                        price decimal(10,2)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;