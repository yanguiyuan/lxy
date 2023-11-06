CREATE DATABASE lxy DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE lxy;
CREATE TABLE IF NOT EXISTS words(
    `id`  BIGINT PRIMARY KEY  COMMENT 'ID',
    `value` VARCHAR(50) NOT NULL COMMENT '流行语名称',
    `paraphrase` TEXT COMMENT '释义',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS tag_types(
      `id`  BIGINT PRIMARY KEY  COMMENT 'ID',
     `value` VARCHAR(50) NOT NULL COMMENT '标签类型名称',
     `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS tags(
    `id`  BIGINT PRIMARY KEY  COMMENT 'ID',
    `tag_type_id` BIGINT NOT NULL,
    `value` VARCHAR(50) NOT NULL COMMENT '标签名称',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tag_type_id) REFERENCES tag_types(id)
);

CREATE TABLE IF NOT EXISTS word_tags(
    `word_id` BIGINT NOT NULL,
    `tag_id` BIGINT NOT NULL,
    FOREIGN KEY (word_id) REFERENCES words(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);
CREATE TABLE IF NOT EXISTS baidu_hots(
  `word_id` BIGINT NOT NULL,
  `start_time` TIMESTAMP NOT NULL ,
  `end_time` TIMESTAMP NOT NULL ,
  `search_avg` int4,
  `search_yoy` int4,
  `search_qoq` int4,
  `search_pc_avg` int4,
  `search_pc_yoy` int4,
  `search_pc_qoq` int4,
  `search_wise_avg` int4,
  `search_wise_yoy` int4,
  `search_wise_qoq` int4,
  `feed_avg` int4,
  `feed_yoy` int4,
  `feed_qoq` int4
);