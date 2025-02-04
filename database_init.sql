CREATE TABLE `public_keys` (
	`room` VARCHAR(30) NOT NULL COLLATE 'utf8mb4_0900_ai_ci',
	`user` VARCHAR(30) NOT NULL COLLATE 'utf8mb4_0900_ai_ci',
	`key` BINARY(64) NOT NULL,
	PRIMARY KEY (`room`, `user`) USING BTREE
)
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB
;
