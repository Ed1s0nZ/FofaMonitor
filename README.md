# FofaMonitor  
通过Fofa引擎实现对外网暴露资产进行监控，如果有新增资产则更新到数据库，并通过企业微信机器人每日上报匹配结果、存活站点。  

建表（之前用于假网站追踪，所以起了fake_web作为表名，根据自己需要更改，并修改sql.go代码中的表名）：
```   
CREATE TABLE `fake_web` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(512) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `port` varchar(512) NOT NULL DEFAULT '' COMMENT '端口',
  `domain` varchar(512) NOT NULL DEFAULT '' COMMENT '域名',
  `country` varchar(512) NOT NULL DEFAULT '' COMMENT '国家代码',
  `title` varchar(512) NOT NULL DEFAULT '' COMMENT '网站标题',
  `protocol` varchar(1000) NOT NULL DEFAULT '' COMMENT '协议名',
  `country_name` varchar(1000) NOT NULL DEFAULT '' COMMENT '国家名',
  `region` varchar(1000) NOT NULL DEFAULT '' COMMENT '区域',
  `city` varchar(1000) NOT NULL DEFAULT '' COMMENT '城市',
  `longitude` varchar(1000) NOT NULL DEFAULT '' COMMENT '地理位置 经度',
  `latitude` varchar(1000) NOT NULL DEFAULT '' COMMENT '地理位置 纬度',
  `as_number` varchar(1000) NOT NULL DEFAULT '' COMMENT 'asn编号',
  `as_organization` varchar(1000) NOT NULL DEFAULT '' COMMENT 'asn组织',
  `host` varchar(1000) NOT NULL DEFAULT '' COMMENT '主机名',
  `os` varchar(1000) NOT NULL DEFAULT '' COMMENT '操作系统',
  `server` varchar(512) NOT NULL DEFAULT '' COMMENT '网站server',
  `icp` varchar(512) NOT NULL DEFAULT '' COMMENT 'icp备案号',
  `jarm` varchar(512) NOT NULL DEFAULT '' COMMENT 'jarm 指纹',
  `header` varchar(512) NOT NULL DEFAULT '' COMMENT '网站header',
  `banner` varchar(512) NOT NULL DEFAULT '' COMMENT '协议 banner',
  `cert` varchar(512) NOT NULL DEFAULT '' COMMENT '证书',
  `base_protocol` varchar(100) NOT NULL DEFAULT '' COMMENT '基础协议，比如tcp/udp',
  `link` varchar(100) NOT NULL DEFAULT '' COMMENT '资产的URL链接',
  `body` varchar(100) NOT NULL DEFAULT '' COMMENT '网站正文内容',
  `hashvalue` varchar(100) NOT NULL COMMENT '唯一hash值',
  `cname` varchar(100) NOT NULL DEFAULT '' COMMENT '域名cname',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3499 DEFAULT CHARSET=utf8mb4;
```
