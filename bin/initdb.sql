CREATE TABLE `comments`
(
    `id`        varchar(64) NOT NULL,
    `video_id`  varchar(64) DEFAULT NULL,
    `author_id` int(11) unsigned DEFAULT NULL,
    `content`   text,
    `time`      datetime    DEFAULT current_timestamp COMMENT '评论时间\r\n',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `sessions`
(
    `session_id` varchar(64) NOT NULL,
    `TTL`        int(64) unsigned DEFAULT NULL,
    `login_name` varchar(64) DEFAULT NULL,
    PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT,
    `login_name` varchar(64) NOT NULL DEFAULT '',
    `pwd`        text        NOT NULL,
    PRIMARY KEY (`id`, `login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `video_del_rec`
(
    `video_id` varchar(64) NOT NULL,
    PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `video_info`
(
    `id`            varchar(64) NOT NULL,
    `author_id`     int(10) unsigned NOT NULL COMMENT '用户id',
    `name`          text        NOT NULL COMMENT '视频名称',
    `display_ctime` text        NOT NULL COMMENT '显示的创建时间',
    `create_time`   datetime DEFAULT NULL COMMENT '入库时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;