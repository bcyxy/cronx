/* 任务对象
|type    |obj_group |obj_key         |version       |flags  |
|--------|----------|----------------|--------------|-------|
|NET_DEV |BJYZ      |192.168.11.22   |HUAWEI:CE1234 |EVR,L1 |
|NET_SEG |BJYZ      |192.168.11.0/24 |              |       |*/
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

/* 任务动作
|act_key |obj_version |act_type |act_param              |
|--------|------------|---------|-----------------------|
|ping    |            |PING_SEG |                       |
|if_flow |HUAWEI:CE*  |SNMPWALK |1.3.6.1.2.1.31.1.1.1.6 |*/
CREATE TABLE `job_conf_action` (
    id           bigint(20)   NOT NULL AUTO_INCREMENT,
    act_key      varchar(32)  NOT NULL,
    obj_version  varchar(64)  NOT NULL,
    act_type     varchar(32)  NOT NULL,
    act_param    varchar(64)  NOT NULL,
    create_time  datetime     NOT NULL,
    update_time  datetime     NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 任务计划
|act_key |obj_type  |obj_val |interval |pri |
|--------|----------|--------|---------|----|
|ping    |ALL       |        |60       |20  |
|if_flow |ALL       |        |30       |20  |
|if_flow |OBJ_GROUP |BJYZ    |15       |30  |*/
CREATE TABLE `job_conf_sch` (
    id           bigint(20)  NOT NULL AUTO_INCREMENT,
    act_key      varchar(32) NOT NULL,
    obj_type     varchar(32) NOT NULL,
    obj_val      varchar(32) NOT NULL,
    interval     int(10)     NOT NULL,
    pri          int(10)     NOT NULL,
    create_time  datetime    NOT NULL,
    update_time  datetime    NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 系统状态 */
CREATE TABLE `sys_status` (
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT,
    `create_time`  datetime    NOT NULL,
    `update_time`  datetime    NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;"
