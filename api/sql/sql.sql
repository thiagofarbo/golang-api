CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS user;

CREATE TABLE user(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES user(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES user(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE posts(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES user(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;