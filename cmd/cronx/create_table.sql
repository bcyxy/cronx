/* 任务对象 */
CREATE TABLE job_conf_obj (
    id           bigint(20)  NOT NULL AUTO_INCREMENT,
    type         varchar(32) NOT NULL,
    obj_group    varchar(32) NOT NULL,
    obj_key      varchar(32) NOT NULL,
    version      varchar(64) NOT NULL,
    flags        varchar(64) NOT NULL,
    create_time  datetime    NOT NULL,
    update_time  datetime    NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uidx1 (obj_group, obj_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 任务动作 */
CREATE TABLE job_conf_action (
    id           bigint(20)   NOT NULL AUTO_INCREMENT,
    act_key      varchar(32)  NOT NULL,
    obj_version  varchar(64)  NOT NULL,
    act_type     varchar(32)  NOT NULL,
    act_param    varchar(64)  NOT NULL,
    create_time  datetime     NOT NULL,
    update_time  datetime     NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uidx1 (act_key, obj_version)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 任务计划 */
CREATE TABLE job_conf_sch (
    id           bigint(20)  NOT NULL AUTO_INCREMENT,
    act_key      varchar(32) NOT NULL,
    obj_type     varchar(32) NOT NULL,
    obj_scope    varchar(32) NOT NULL,
    cycle        int(10)     NOT NULL,
    pri          int(10)     NOT NULL,
    create_time  datetime    NOT NULL,
    update_time  datetime    NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 系统状态 */
CREATE TABLE sys_status (
    id           bigint(20)   NOT NULL AUTO_INCREMENT,
    state_type   varchar(64)  NOT NULL,
    state_key    varchar(128) NOT NULL,
    state_val    varchar(128) NOT NULL,
    create_time  datetime     NOT NULL,
    update_time  datetime     NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY uidx1 (state_type, state_key)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
