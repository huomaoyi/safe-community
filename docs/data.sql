-- 用户信息表
create table user_account
(
  id              int auto_increment
    primary key,
  created_at      datetime     default CURRENT_TIMESTAMP null comment '创建时间',
  updated_at      datetime     default CURRENT_TIMESTAMP null comment '更新时间',
  real_name       varchar(128) default ''                null comment '真实姓名',
  alia_name       varchar(128)                           not null comment '昵称',
  phone           varchar(20)  default ''                null comment '电话号码',
  email           varchar(128) default ''                null comment '邮箱',
  community       varchar(256) default ''                null comment '小区名称',
  building_number int          default 0                 null comment '小区楼号',
  building_uint   int          default 0                 null comment '楼单元号',
  house_number    int          default 0                 null comment '房间号',
  province        varchar(128) default ''                null comment '省份',
  city            varchar(128) default ''                null comment '城市',
  constraint user_account_alia_name_uindex
    unique (alia_name)
);

-- 用户体温表
create table user_temperature
(
  id          int auto_increment primary key,
  created_at  datetime       default CURRENT_TIMESTAMP null comment '创建时间',
  updated_at  datetime       default CURRENT_TIMESTAMP null comment '更新时间',
  alia_name   varchar(128)   default ''                null comment '用户昵称',
  temperature decimal(32, 4) default 0.0000            null comment '用户体温'
)
  comment '用户温度表';






