-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Aug 26, 2025 at 12:00 PM
-- Server version: 8.0.30
-- PHP Version: 8.2.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_rest_api`
--

-- --------------------------------------------------------

--
-- Table structure for table `authors`
--

CREATE TABLE `authors` (
  `id` bigint UNSIGNED NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `authors`
--

INSERT INTO `authors` (`id`, `name`, `gender`, `email`, `age`, `created_at`, `updated_at`) VALUES
(5, 'Fadil', 'L', 'fadil@gmail.com', 21, '2025-08-25 16:31:28.652', '2025-08-26 18:40:48.091'),
(8, 'Ahmad', 'L', 'ahmad@gmail.com', 31, '2025-08-25 16:39:12.063', '2025-08-25 16:39:12.063'),
(9, 'Budi', 'L', 'budi@gmail.com', 29, '2025-08-25 16:49:15.931', '2025-08-25 16:49:15.931'),
(10, 'Ani', 'P', 'ani@gmail.com', 19, '2025-08-25 16:50:25.501', '2025-08-25 16:50:25.501'),
(11, 'Akmal Rahim', 'L', 'akmalrahim@gmail.com', 25, '2025-08-25 21:52:05.270', '2025-08-25 21:52:05.270');

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `id` bigint UNSIGNED NOT NULL,
  `author_id` bigint UNSIGNED DEFAULT NULL,
  `title` varchar(120) NOT NULL,
  `description` text NOT NULL,
  `published_year` varchar(4) NOT NULL,
  `published` varchar(120) NOT NULL,
  `pages` varchar(5) NOT NULL,
  `isbn` varchar(70) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `books`
--

INSERT INTO `books` (`id`, `author_id`, `title`, `description`, `published_year`, `published`, `pages`, `isbn`, `created_at`, `updated_at`) VALUES
(1, 8, 'Database Design Best Practices', 'Panduan desain database relasional yang efisien dan aman, lengkap dengan contoh.', '2019', 'DB Experts Press', '180', 'ISBN-5555-6666', '2025-08-25 21:58:29.768', '2025-08-26 18:42:08.433'),
(2, 11, 'Python for Data Science', 'Panduan praktis Python untuk analisis data, machine learning, dan visualisasi.', '2021', 'DataTech Publishing', '300', 'ISBN-1111-2222', '2025-08-25 22:01:45.139', '2025-08-25 22:01:45.139'),
(4, 11, 'Mastering RESTful API with Go', 'Panduan lengkap membangun RESTful API menggunakan bahasa Go, mencakup praktik terbaik dan studi kasus.', '2021', 'Go Dev Publisher', '250', 'ISBN-7777-8888', '2025-08-26 18:43:10.740', '2025-08-26 18:43:10.740'),
(5, 10, 'Clean Code for Web Development', 'Prinsip menulis kode bersih untuk pengembangan aplikasi web yang mudah dibaca, dipelihara, dan scalable.', '2020', 'TechCraft Publishing', '320', 'ISBN-9999-1111', '2025-08-26 18:43:23.816', '2025-08-26 18:43:23.816'),
(7, 11, 'Blockchain Technology Explained', 'Penjelasan teknologi blockchain dari dasar hingga penerapan nyata di industri.', '2021', 'CryptoTech Books', '275', 'ISBN-9999-0000', '2025-08-26 18:54:30.754', '2025-08-26 18:54:30.754'),
(8, 9, 'Agile Project Management', 'Metodologi agile untuk mengelola proyek perangkat lunak yang adaptif dan cepat.', '2017', 'AgileWorks Publishing', '195', 'ISBN-3333-4444', '2025-08-26 18:54:54.125', '2025-08-26 18:54:54.125'),
(9, 5, 'Kubernetes for DevOps Engineers', 'anduan komprehensif untuk mengelola container dengan Kubernetes bagi DevOps.', '2020', 'DL Labs Press', '310', 'ISBN-5555-6666', '2025-08-26 18:55:03.658', '2025-08-26 18:55:59.232');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `authors`
--
ALTER TABLE `authors`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_books_author` (`author_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `authors`
--
ALTER TABLE `authors`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `books`
--
ALTER TABLE `books`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `books`
--
ALTER TABLE `books`
  ADD CONSTRAINT `fk_books_author` FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
