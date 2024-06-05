CREATE TABLE `stores` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `phone` varchar(20) UNIQUE NOT NULL,
  `address` varchar(255) UNIQUE,
  `email` varchar(255) UNIQUE,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
);

--bun:split

CREATE TABLE `books` (
  `store_id` integer NOT NULL,
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `title` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
);

--bun:split

CREATE TABLE `authors` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `book_id` integer NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
);

--bun:split

ALTER TABLE `books` ADD FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`);

--bun:split

ALTER TABLE `books` ADD FOREIGN KEY (`id`) REFERENCES `authors` (`book_id`);
