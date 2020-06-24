-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: live_code_db
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `m_dosen`
--

DROP TABLE IF EXISTS `m_dosen`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_dosen` (
  `kode_dosen` varchar(15) NOT NULL,
  `nama_dosen` varchar(45) NOT NULL,
  PRIMARY KEY (`kode_dosen`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_dosen`
--

LOCK TABLES `m_dosen` WRITE;
/*!40000 ALTER TABLE `m_dosen` DISABLE KEYS */;
INSERT INTO `m_dosen` VALUES ('20200624115458','Mas Edo'),('20200624150554','Kang Prana'),('20200624152719','Maulana Ibrahim');
/*!40000 ALTER TABLE `m_dosen` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_mahasiswa`
--

DROP TABLE IF EXISTS `m_mahasiswa`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_mahasiswa` (
  `nim` varchar(15) NOT NULL,
  `nama_mahasiswa` varchar(45) NOT NULL,
  `jenis_kelamin` int NOT NULL,
  `alamat` text NOT NULL,
  PRIMARY KEY (`nim`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_mahasiswa`
--

LOCK TABLES `m_mahasiswa` WRITE;
/*!40000 ALTER TABLE `m_mahasiswa` DISABLE KEYS */;
INSERT INTO `m_mahasiswa` VALUES ('20200624124703','Maulana Ibrahim',1,'Cirebon'),('20200624131307','Ibrahim Update',1,'Cirebon'),('20200624140746','Ibrahim Maulana',1,'Cirebon'),('20200624144000','Ibrahim Lagi',1,'Cirebon');
/*!40000 ALTER TABLE `m_mahasiswa` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_mahasiswa_matkul`
--

DROP TABLE IF EXISTS `m_mahasiswa_matkul`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_mahasiswa_matkul` (
  `nim` varchar(15) NOT NULL,
  `kode_matkul` varchar(15) NOT NULL,
  `nilai` varchar(2) DEFAULT NULL,
  `nilaiNumber` int DEFAULT '0',
  PRIMARY KEY (`nim`,`kode_matkul`),
  KEY `fk_m_mahasiswa_has_m_mata_kuliah_m_mata_kuliah1_idx` (`kode_matkul`),
  KEY `fk_m_mahasiswa_has_m_mata_kuliah_m_mahasiswa1_idx` (`nim`),
  CONSTRAINT `fk_m_mahasiswa_has_m_mata_kuliah_m_mahasiswa1` FOREIGN KEY (`nim`) REFERENCES `m_mahasiswa` (`nim`),
  CONSTRAINT `fk_m_mahasiswa_has_m_mata_kuliah_m_mata_kuliah1` FOREIGN KEY (`kode_matkul`) REFERENCES `m_mata_kuliah` (`kode_matkul`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_mahasiswa_matkul`
--

LOCK TABLES `m_mahasiswa_matkul` WRITE;
/*!40000 ALTER TABLE `m_mahasiswa_matkul` DISABLE KEYS */;
INSERT INTO `m_mahasiswa_matkul` VALUES ('20200624124703','20200624115620','B',89),('20200624124703','20200624115811','A',0),('20200624131307','20200624115811','B',0),('20200624144000','20200624152554','C',75);
/*!40000 ALTER TABLE `m_mahasiswa_matkul` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `m_mata_kuliah`
--

DROP TABLE IF EXISTS `m_mata_kuliah`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `m_mata_kuliah` (
  `kode_matkul` varchar(15) NOT NULL,
  `nama_matkul` varchar(45) NOT NULL,
  `kode_dosen` varchar(15) NOT NULL,
  PRIMARY KEY (`kode_matkul`),
  KEY `fk_m_mata_kuliah_m_dosen_idx` (`kode_dosen`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `m_mata_kuliah`
--

LOCK TABLES `m_mata_kuliah` WRITE;
/*!40000 ALTER TABLE `m_mata_kuliah` DISABLE KEYS */;
INSERT INTO `m_mata_kuliah` VALUES ('20200624115620','Dasar Program','20200624115458'),('20200624115811','Bahasa','20200624150554'),('20200624152554','Golang DB','20200624150554'),('20200624152708','Golang DB Mysql','20200624115458');
/*!40000 ALTER TABLE `m_mata_kuliah` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-06-24 15:46:02
