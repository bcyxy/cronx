INSERT INTO job_conf_obj
(type, obj_group, obj_key, version, flags, create_time, update_time) VALUES
('DEV', '', '127.0.0.1', 'LINUX', '', NOW(), NOW()),
('DEV', '', '127.0.0.2', 'LINUX', '', NOW(), NOW());

INSERT INTO job_conf_action
(act_key, obj_version, act_type, act_param, create_time, update_time) VALUES
('mem_rate', 'LINUX', 'SSH', '{"cmd":"free", "reg":"Mem: +(\d+) +(\d+)"}', NOW(), NOW());

INSERT INTO job_conf_sch
(act_key, obj_type, obj_scope, cycle, pri, create_time, update_time) VALUES
('mem_rate', 'DEV', 'ALL', 10, 20, NOW(), NOW());
