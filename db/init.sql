USE app;

CREATE TABLE
    `user` (
            `id` VARCHAR(100) NOT NULL COMMENT 'ユーザー ID',
            `name` VARCHAR(100) NOT NULL COMMENT 'ユーザー名',
            `password` VARCHAR(100) NOT NULL COMMENT 'パスワード',
            `created_at` DATETIME NOT NULL COMMENT '作成日時',
            `updated_at` DATETIME NOT NULL COMMENT '更新日時',
            PRIMARY KEY (`id`)
) COMMENT = 'ユーザー';

CREATE TABLE
    `books` (
            `id` VARCHAR(100) NOT NULL COMMENT 'コーヒー ID',
            `title` VARCHAR(100) COMMENT 'コーヒー名',
            `created_at` DATETIME NOT NULL COMMENT '作成日時',
            `updated_at` DATETIME NOT NULL COMMENT '更新日時',
            PRIMARY KEY (`id`)
) COMMENT = 'コーヒー';


INSERT INTO `user`(id, name, password, created_at, updated_at)
VALUES ('1', 'name1', 'password', NOW(), NOW());

INSERT INTO `user`(d, name, password, created_at, updated_at)
VALUES ('2', 'name2', 'password', NOW(), NOW());

INSERT INTO `user`(id, name, password, created_at, updated_at)
VALUES ('3', 'name3', 'password', NOW(), NOW());

INSERT INTO `books` (id, title, created_at, updated_at)
VALUES ('1', 'title1', NOW(), NOW());

INSERT INTO `books` (id, title, created_at, updated_at)
VALUES ('2', 'title2', NOW(), NOW());