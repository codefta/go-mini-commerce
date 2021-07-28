SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 1;

START TRANSACTION;
SET time_zone = "+00:00";

--
-- Database: `mini_commerce`
--

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `product_name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `price` int NOT NULL,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `category_name` varchar(100) NOT NULL
);

CREATE TABLE `product_category` (
  `product_id` int,
  `category_id` int
);

ALTER TABLE `product_category` ADD CONSTRAINT product_fk_id FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE;
ALTER TABLE `product_category` ADD CONSTRAINT category_fk_id FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE;

COMMIT;
