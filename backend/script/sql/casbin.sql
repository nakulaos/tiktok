-- tiktok.casbin_rule definition
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
                               `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
                               `ptype` varchar(32) NOT NULL DEFAULT '',
                               `v0` varchar(255) NOT NULL DEFAULT '',
                               `v1` varchar(255) NOT NULL DEFAULT '',
                               `v2` varchar(255) NOT NULL DEFAULT '',
                               `v3` varchar(255) NOT NULL DEFAULT '',
                               `v4` varchar(255) NOT NULL DEFAULT '',
                               `v5` varchar(255) NOT NULL DEFAULT '',
                               PRIMARY KEY (`id`) USING BTREE,
                            INDEX `idx_casbin_rule` (`ptype`,`v0`,`v1`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
        ('p','visitor','/v1/user/register','POST','allow','',''),
        ('p','visitor','/v1/user/login/username','POST','allow','',''),
        ('p','visitor','/v1/user/login/email','POST','allow','',''),
        ('p','visitor','/v1/user/login/phone','POST','allow','',''),
        ('p','normalUser','/v1/user/info','POST','allow','',''),
        ('g','normalUser','visitor','','','',''),
        ('g','admin','normalUser','','','','');


INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
       ('p','normalUser','/v1/relation/action','POST','allow','',''),
       ('p','normalUser','/v1/relation/favoriteList','GET','allow','',''),
       ('p','normalUser','/v1/relation/followerList','GET','allow','',''),
       ('p','normalUser','/v1/relation/friendList','GET','allow','','');

INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
      ('p','visitor','/v1/user/upload','POST','allow','','');


INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
    ('p','normalUser','/v1/user/upload','POST','allow','',''),
    ('p','normalUser','/v1/feed/create','POST','allow','','');


INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
    ('p','normalUser','/v1/favorite/action','POST','allow','',''),
    ('p','normalUser','/v1/favorite/list','GET','allow','',''),
    ('p','normalUser','/v1/star/action','POST','allow','',''),
    ('p','normalUser','/v1/star/list','GET','allow','','');


INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
    ('p','normalUser','/v1/feed/CategoryVideosList','GET','allow','',''),
    ('p','normalUser','/v1/feed/UserVideosList','POST','allow','',''),
    ('p','normalUser','/v1/feed/VideosList','POST','allow','',''),
    ('p','normalUser','/v1/feed/deleteViedo','POST','allow','',''),
    ('p','normalUser','/v1/feed/duration','POST','allow','',''),
    ('p','normalUser','/v1/feed/history','GET','allow','',''),
    ('p','normalUser','/v1/feed/neighbors','GET','allow','',''),
    ('p','normalUser','/v1/feed/populars','POST','allow','',''),
    ('p','normalUser','/v1/feed/recommends','POST','allow','',''),
    ('p','normalUser','/v1/feed/searcheEs','POST','allow','',''),
    ('p','normalUser','/v1/feed/videoinfo','GET','allow','','');

INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
    ('p','normalUser','/v1/comment/action','POST','allow','',''),
    ('p','normalUser','/v1/comment/list','GET','allow','',''),
    ('p','visitor','/v1/danmu/action','POST','allow','',''),
    ('p','visitor','/v1/danmu/list','GET','allow','','');



INSERT INTO casbin_rule (ptype, v0, v1, v2, v3, v4, v5) VALUES
 ('p','normalUser','/v1/live/start','GET','allow','',''),
  ('p','normalUser','/v1/live/list','GET','allow','','');


