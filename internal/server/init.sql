create table if not exists `user` (
    `id` int(11) not null auto_increment,
    `name` varchar(255) not null,
    `email` varchar(255) not null,
    `password` varchar(255) not null,
    `created_at` datetime not null,
    `updated_at` datetime not null,
    primary key (`id`),
    unique key `email` (`email`)
) engine=InnoDB default charset=utf8mb4;



INSERT INTO USERS (name, email, password, created_at, updated_at) VALUES ('admin', 'admin@localhost', 'admin', NOW(), NOW());