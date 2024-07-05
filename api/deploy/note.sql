create database note;
use note;

-- 修改后的 note 表
create table note
(
    id         bigint not null auto_increment,
    title      varchar(100),
    file_id    int,
    created_at datetime        default now(),
    creator_id int    not null default 0,
    is_deleted tinyint         default 0,
    index idx_note_title (title),
    index idx_note_creator_id (creator_id),
    primary key (id)

) charset utf8mb4
  collate utf8mb4_unicode_ci;

-- 修改后的 file 表
create table file
(
    id           bigint not null auto_increment,
    file_name    varchar(100),
    storage_id   varchar(100) comment '存储服务中的坐标',
    storage_type int comment '文件支持多种存储方式',
    created_at   datetime        default now(),
    creator_id   int    not null default 0,
    upload_type  int comment '文件支持多种上传方式',
    is_deleted   tinyint         default 0,
    index idx_file_creator_id (creator_id),
    index idx_file_storage_id (storage_id),
    primary key (id)
) charset utf8mb4
  collate utf8mb4_unicode_ci;