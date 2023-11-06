create table words(
    id bigint primary key ,
    value varchar(50) not null unique ,
--     has_derive bool comment '是否有衍生',
    paraphrase text comment '释义',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table tags(
    id bigint primary key ,
    tag_type_id bigint not null ,
    value varchar(50) not null unique ,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table tag_types(
    id bigint primary key ,
    value varchar(50) not null unique ,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table word_tags(
    id bigint primary key ,
    word_id bigint not null,
    tag_id bigint not null ,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- 创建一个words触发器
CREATE TRIGGER words_update_timestamp
    AFTER UPDATE ON words
    FOR EACH ROW
BEGIN
    UPDATE words
    SET update_time = DATETIME('now')
    WHERE id = NEW.id;
END;

-- 创建一个tags触发器
CREATE TRIGGER tags_update_timestamp
    AFTER UPDATE ON tags
    FOR EACH ROW
BEGIN
    UPDATE tags
    SET update_time = DATETIME('now')
    WHERE id = NEW.id;
END;