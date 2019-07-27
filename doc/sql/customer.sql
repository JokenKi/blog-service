CREATE TABLE `customer` (
`id`  bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键' ,
`name`  varchar(64) NOT NULL DEFAULT '' COMMENT '用户名' ,
`nick_name`  varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称' ,
`passwd`  varchar(32) NOT NULL DEFAULT '' COMMENT '登陆密码' ,
`salt`  smallint(5) UNSIGNED NULL COMMENT '盐值' ,
`phone`  varchar(11) NULL DEFAULT '' COMMENT '手机号' ,
`time_create`  bigint(20) UNSIGNED NOT NULL DEFAULT 0 ,
`time_update`  bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间' ,
`time_latest_login`  bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最近登陆时间' ,
PRIMARY KEY (`id`)
)
;

INSERT INTO `blog`.`customer` (`id`, `name`, `nick_name`, `passwd`, `salt`, `phone`, `time_create`, `time_update`, `time_latest_login`) VALUES ('1', 'admin', '超级管理员', '11aad30a7652bd870b7cecb0909477c1', '1234', '18612802880', '1557733831', '1557733831', '0');
