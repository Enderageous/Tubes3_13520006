-- --------------------------------------------------------
-- Host:                         remotemysql.com
-- Server version:               8.0.13-4 - Percona Server (GPL), Release '4', Revision 'f0a32b8'
-- Server OS:                    debian-linux-gnu
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for 82kHLszeSh
CREATE DATABASE IF NOT EXISTS `82kHLszeSh` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `82kHLszeSh`;

-- Dumping structure for table 82kHLszeSh.disease
CREATE TABLE IF NOT EXISTS `disease` (
  `disease_id` int(11) NOT NULL AUTO_INCREMENT,
  `disease_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `dna_sequence` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`disease_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Dumping data for table 82kHLszeSh.disease: ~9 rows (approximately)
/*!40000 ALTER TABLE `disease` DISABLE KEYS */;
INSERT INTO `disease` (`disease_id`, `disease_name`, `dna_sequence`) VALUES
	(14, 'Sugma', 'ACACACACCGCCGCCCC'),
	(15, 'takut tambah dewasa', 'AAAAAA'),
	(16, 'Sukma', 'CGCCGC'),
	(17, 'COVID', 'ATGATGATAGATG'),
	(18, 'SARS', 'ATGATGTGTAGTAG'),
	(19, 'H5N1', 'ATGATGTGTAGTAG'),
	(20, 'HIV', 'CGATTAGCTAGCTAGCTAGCTAGCATCGATCAGATCAG'),
	(21, 'SleepDeprived', 'CGATTAGCTAGCTAGCTAGCTAGCATCGATCAGATCAG'),
	(22, 'Flu', 'CGATTAGCTAGCTAGCTAGCTAGCATCGATCAGATCAG');
/*!40000 ALTER TABLE `disease` ENABLE KEYS */;

-- Dumping structure for table 82kHLszeSh.prediction
CREATE TABLE IF NOT EXISTS `prediction` (
  `prediction_id` int(11) NOT NULL AUTO_INCREMENT,
  `prediction_date` date NOT NULL,
  `patient_name` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
  `dna_sequence` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `disease_id` int(11) NOT NULL,
  `result` tinyint(1) NOT NULL,
  `accuracy` float NOT NULL,
  PRIMARY KEY (`prediction_id`),
  KEY `disease_id` (`disease_id`),
  CONSTRAINT `prediction_ibfk_1` FOREIGN KEY (`disease_id`) REFERENCES `disease` (`disease_id`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- Dumping data for table 82kHLszeSh.prediction: ~16 rows (approximately)
/*!40000 ALTER TABLE `prediction` DISABLE KEYS */;
INSERT INTO `prediction` (`prediction_id`, `prediction_date`, `patient_name`, `dna_sequence`, `disease_id`, `result`, `accuracy`) VALUES
	(25, '2022-04-29', 'sy', 'CCCCCAAAAAAGGAT', 15, 1, 1),
	(29, '2022-04-29', 'ayaya', 'ACACACACCGCCGCCCC', 15, 0, 1),
	(30, '2022-04-29', 'ayaya', 'ACACACACCGCCGCCCC', 15, 0, 1),
	(31, '2022-04-29', 'ayaya', 'ACACACACCGCCGCCCC', 15, 0, 1),
	(32, '2022-04-27', 'ayaya', 'ACACACACCGCCGCCCC', 15, 0, 1),
	(33, '2022-04-29', 'ayaya', 'ACACACACCGCCGCCCC', 15, 0, 1),
	(34, '2022-04-29', 'ayaya', 'ACACACACCGCCGCCCC', 15, 0, 1),
	(37, '2022-04-29', 'ayaya', 'CGCCGC', 16, 1, 1),
	(38, '2022-04-29', 'adalah sy frfrfr', 'AATATTATATATAACTGGAGAT', 15, 0, 1),
	(39, '2022-04-29', 'uiui', 'ATGATGATAGATG', 16, 0, 1),
	(40, '2022-04-29', 'tes1', 'GTAATGATGAT', 15, 0, 1),
	(41, '2022-04-29', 'Vionie', 'ATGATGATAGATG', 17, 1, 1),
	(42, '2022-04-29', 'Ziel', 'GTAATGATGAT', 18, 0, 1),
	(43, '2022-04-29', 'AJI', 'CGCTCGCGCGCTATATATACGATTAGCTAGCTAGCTAGCTAGCATCGATCAGATCAGATATATATATCGCGCGCATCGATCGA', 20, 1, 1),
	(44, '2022-04-29', 'AJI', 'CGCTCGCGCGCTATATATACGATTAGCTAGCGTAGCTAGCTAGCATCGATCAGATCAGATATATATATCGCGCGCATCGATCGA', 20, 0, 1),
	(45, '2022-04-29', 'AJI', 'CGCTCGCGCGCTATATATACGATTAGCTAGCTAGCTAGCTAGCATCGATCAGATCAGATATATATATCGCGCGCATCGATCGA', 21, 1, 1);
/*!40000 ALTER TABLE `prediction` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
