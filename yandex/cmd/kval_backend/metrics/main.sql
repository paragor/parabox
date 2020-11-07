create table qrtz_log
(
    id                integer primary key autoincrement not null,
    job_name          varchar(80)                       not null,
    job_group         varchar(80)                       not null,
    trigger_fire_time timestamp with time zone          not null,
    job_finished_time timestamp with time zone,
    job_status        varchar(200),
    host_name         varchar(80)                       not null,
    constraint con_name_group_fire_time unique (
                                                job_name, job_group, trigger_fire_time, host_name
        )
);

insert into qrtz_log (id, job_name, job_group, trigger_fire_time, job_finished_time, job_status, host_name)
values (1, 'jobA', 'G1', '2020-09-01 08:00:00', NULL, NULL, 'host1'),
       (2, 'jobA', 'G1', '2020-09-01 07:00:00', '2020-09-01 07:00:03', 'OK', 'host1'),
       (3, 'jobA', 'G1', '2020-09-01 06:00:00', '2020-09-01 06:00:04', 'OK', 'host2'),
       (4, 'jobB', 'G1', '2020-09-01 08:00:00', NULL, 'ERR', 'host1'),
       (5, 'jobB', 'G1', '2020-09-01 07:00:00', '2020-09-01 07:00:03', 'OK', 'host1'),
       (6, 'jobB', 'G1', '2020-09-01 06:00:00', '2020-09-01 06:00:04', 'OK', 'host1'),
       (7, 'jobC', 'G1', '2020-09-01 08:00:00', '2020-09-01 08:00:01', 'OK', 'host3'),
       (8, 'jobC', 'G1', '2020-09-01 07:00:00', NULL, NULL, 'host2'),
       (9, 'jobC', 'G1', '2020-09-01 06:00:00', '2020-09-01 06:00:04', 'OK', 'host1'),
       (10, 'jobD', 'G1', '2020-09-01 09:59:59', NULL, NULL, 'host1'),
       (11, 'jobD', 'G1', '2020-06-01 07:00:00', '2020-06-01 10:00:00', 'OK', 'host1'),
       (12, 'jobD', 'G1', '2020-05-01 06:00:00', '2020-05-01 10:00:00', 'OK', 'host1')



