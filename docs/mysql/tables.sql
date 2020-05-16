CREATE SCHEMA `vshare` DEFAULT CHARACTER SET utf8mb4;
CREATE TABLE `vshare`.`t_user`
(
    `id`          INT      NOT NULL,
    `uid`         INT      NOT NULL AUTO_INCREMENT,
    `status`      INT      NOT NULL DEFAULT 0,
    `register_at` DATETIME NOT NULL,
    `updated_at`  DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `index2` (`uid` ASC)
);
CREATE TABLE `vshare`.`t_user_access`
(
    `uid`         INT          NOT NULL,
    `access_type` INT          NOT NULL COMMENT '1-微信登陆',
    `wx_openid`   VARCHAR(255) NULL,
    `wx_unionid`  VARCHAR(255) NULL,
    PRIMARY KEY (`uid`)
);

CREATE TABLE `vshare`.`t_room`
(
    `rid`  VARCHAR(255) NOT NULL COMMENT '房间id',
    `host` VARCHAR(45)  NOT NULL COMMENT '房主uid',
    `vid`  VARCHAR(45)  NOT NULL COMMENT '视频vid',
    PRIMARY KEY (`rid`)
);

CREATE TABLE `vshare`.`t_video`
(
    `vid`       VARCHAR(32)  NOT NULL,
    `file_path` VARCHAR(255) NULL,
    `filename`  VARCHAR(255) NULL,
    PRIMARY KEY (`vid`)
);
