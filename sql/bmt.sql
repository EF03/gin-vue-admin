-- auto-generated definition
create table bmt_activity
(
    id         bigint auto_increment
        primary key,
    created_at datetime(3)  null,
    updated_at datetime(3)  null,
    deleted_at datetime(3)  null,
    name       varchar(50)  not null comment '活动主题',
    start_time datetime(3)  null comment '活动开始时间',
    end_time   datetime(3)  null comment '活动结束时间',
    content    varchar(250) null comment '活动内容',
    cost       int          null comment '活动花费',
    person_num smallint     null comment '活动人数'
)
    comment '羽球活動';

create index idx_bmt_activity_deleted_at
    on bmt_activity (deleted_at);


alter table bmt_period
    add hash1 varchar(250) null comment '雜湊1';
alter table bmt_period
    add hash2 varchar(250) null comment '雜湊2';

-- auto-generated definition
create table bmt_period
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3) null,
    updated_at    datetime(3) null,
    deleted_at    datetime(3) null,
    bmt_id        mediumint   not null comment '羽球活动主键',
    red_user_id1  smallint    null comment '红对队员1',
    red_user_id2  smallint    null comment '红对队员2',
    blue_user_id1 smallint    null comment '蓝对队员1',
    blue_user_id2 smallint    null comment '蓝对队员2',
    red_score     smallint    null comment '红对分数',
    blue_score    smallint    null comment '藍隊分數',
    winner        tinyint     null comment '胜负：1-红对,2-蓝队',
    hash1         varchar(250) null comment '雜湊1',
    hash2         varchar(250) null comment '雜湊2'
);

create index idx_bmt_period_deleted_at
    on bmt_period (deleted_at);

-- auto-generated definition
create table bmt_user
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    bmt_id     mediumint   null comment '羽球活动id',
    uid        smallint    null comment '參加者',
    times      smallint    null comment '参予小节次数',
    constraint bmt_user_uqi_index
        unique (bmt_id, uid)
);

create index idx_bmt_user_deleted_at
    on bmt_user (deleted_at);


