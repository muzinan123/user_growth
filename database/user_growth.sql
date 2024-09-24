/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_coin_detail`;

CREATE TABLE `tb_coin_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'User ID',
  `task_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Task ID',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT 'Points, positive for rewards, negative for penalties',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Modification time',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_coin_task`;

CREATE TABLE `tb_coin_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `task` varchar(255) NOT NULL DEFAULT '' COMMENT 'Task name, must be unique',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT 'Points, positive for rewards, negative for penalties, 0 requires external call to pass value',
  `limit` int(11) NOT NULL DEFAULT '0' COMMENT 'Daily limit, default 0 for no limit',
  `start` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Effective start time',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Modification time',
  `sys_status` int(11) NOT NULL DEFAULT '0' COMMENT 'Status, default 0 for active, 1 for deleted',
  PRIMARY KEY (`id`),
  UNIQUE KEY `task` (`task`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_coin_user`;

CREATE TABLE `tb_coin_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'User ID',
  `coins` int(11) NOT NULL DEFAULT '0' COMMENT 'Total points',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Modification time',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_grade_info`;

CREATE TABLE `tb_grade_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT 'Grade name',
  `description` varchar(3000) NOT NULL COMMENT 'Grade description',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT 'Maximum growth value for this grade',
  `expired` int(11) NOT NULL DEFAULT '0' COMMENT 'Validity period in days, default 0 for never expire',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Modification time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_grade_privilege`;

CREATE TABLE `tb_grade_privilege` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `grade_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Grade ID',
  `product` varchar(255) NOT NULL DEFAULT '' COMMENT 'Product',
  `function` varchar(255) NOT NULL DEFAULT '' COMMENT 'Function',
  `description` varchar(3000) NOT NULL DEFAULT '' COMMENT 'Description',
  `expired` int(11) NOT NULL DEFAULT '0' COMMENT 'Validity period in days, default 0 for never expire',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Modification time',
  `sys_status` int(11) NOT NULL DEFAULT '0' COMMENT 'Status, default 0 for active, 1 for deleted',
  PRIMARY KEY (`id`),
  KEY `grade_id` (`grade_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_grade_user`;

CREATE TABLE `tb_grade_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT 'User ID',
  `grade_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Grade ID',
  `expired` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Expiration time',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT 'Growth value',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Modification time',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
