init

```mysql
CREATE TABLE `check` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) NOT NULL,
  `check_type` varchar(255) NOT NULL,
  `check_value` varchar(255) DEFAULT NULL,
  `check_success` char(1) NOT NULL DEFAULT 'n',
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

