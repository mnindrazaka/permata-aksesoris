-- MySQL dump 10.13  Distrib 8.0.30, for macos12 (x86_64)
--
-- Host: localhost    Database: permata_aksesoris
-- ------------------------------------------------------
-- Server version	8.0.29

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
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `serial` varchar(50) NOT NULL,
  `title` varchar(100) NOT NULL,
  `slug` varchar(100) NOT NULL,
  `icon` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`serial`),
  UNIQUE KEY `slug_UNIQUE` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES ('CAT-1de488c4-fe14-4b3d-ac17-23dd8fe6383f','Pencil','pencil',''),('CAT-55f49826-0869-4475-8342-def5cf59f8a7','tas','tas',''),('CAT-6dc8366b-2499-4dcd-a50b-6bfa6ab7f949','pita-pita','pita-pita',''),('CAT-bde3d484-04ff-48cb-b339-a0c582684372','jam tangan','jam-tangan',''),('CAT-DZsZDebw','kacamata','kacamata',NULL);
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_images`
--

DROP TABLE IF EXISTS `product_images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product_images` (
  `serial` varchar(50) NOT NULL,
  `url` varchar(250) NOT NULL,
  `product_serial` varchar(50) NOT NULL,
  PRIMARY KEY (`serial`),
  KEY `product_serial_idx` (`product_serial`),
  CONSTRAINT `product_serial` FOREIGN KEY (`product_serial`) REFERENCES `products` (`serial`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_images`
--

LOCK TABLES `product_images` WRITE;
/*!40000 ALTER TABLE `product_images` DISABLE KEYS */;
INSERT INTO `product_images` VALUES ('','https://cdf.orami.co.id/unsafe/cdn-cas.orami.co.id/parenting/images/ciri-frame-kacamata-yang-bagus.width-800.jpegquality-80.jpg','PRD-cd5bf8d6-416e-4e17-b8c3-25dcd24f68c0'),('IMG-06ac1651-e9be-4420-a178-11d2b3c55515','https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//97/MTA-2485348/trinity-optima-production_trinity-optima-production-naura-kacamata-sunnies-love-pink-merchandise_full02.jpg','PROD-sLehZ9N4'),('IMG-23b21466-229e-43bd-9d34-0c08ce5689e3','https://www.static-src.com/wcsstore/Indraprastha/images/catalog/full//97/MTA-2485348/trinity-optima-production_trinity-optima-production-naura-kacamata-sunnies-love-pink-merchandise_full02.jpg','PROD-sLehZ9N4');
/*!40000 ALTER TABLE `product_images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `serial` varchar(50) NOT NULL,
  `title` varchar(100) NOT NULL,
  `slug` varchar(100) NOT NULL,
  `thumbnail` varchar(250) DEFAULT NULL,
  `description` longtext,
  `category_serial` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`serial`),
  UNIQUE KEY `slug_UNIQUE` (`slug`),
  KEY `category_serial_idx` (`category_serial`),
  CONSTRAINT `category_serial` FOREIGN KEY (`category_serial`) REFERENCES `categories` (`serial`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES ('PRD-04324f7b-c120-4172-bd82-6c744bf5fbab','kacamata hitam sangat','kacamata-hitam-sangat','https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/19/5c501909-2134-4942-a316-2802656bda11.jpg','Ini kacamata hitam silau','CAT-DZsZDebw'),('PRD-0b78cfdc-4cc4-44be-a67e-9b1f71a785cc','kacamata hitam banget','kacamata-hitam-banget','https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/19/5c501909-2134-4942-a316-2802656bda11.jpg','Ini kacamata hitam silau','CAT-DZsZDebw'),('PRD-141b3563-874e-4d16-9d29-4272de539e12','kacamata hitam silau sekali','kacamata-hitam-silau-sekali','https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/19/5c501909-2134-4942-a316-2802656bda11.jpg','Ini kacamata hitam silau','CAT-DZsZDebw'),('PRD-24e8f8b3-eccc-418d-9632-374da92de970','kacamata hitam gelap','kacamata-hitam-gelap','https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/19/5c501909-2134-4942-a316-2802656bda11.jpg','Ini kacamata hitam silau','CAT-DZsZDebw'),('PRD-28e14c99-911e-4844-b51e-e77fc5353c6d','kacamata hitam dark','kacamata-hitam-dark','https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/19/5c501909-2134-4942-a316-2802656bda11.jpg','Ini kacamata hitam silau','CAT-DZsZDebw'),('PRD-7a5dcbb1-8b49-43e4-9b07-406963e436a7','Kacamata Terkini','kacamata-terkini','http://img.priceza.co.id/img2/2130004/0001/2130004-20220719222117-21964356142526946.jpg','Ini kacamata bukan kacamata biasa, tapi kacamata sangat gaul sekali','CAT-DZsZDebw'),('PRD-7c5e22da-a48b-4346-affc-da2bd1db3f72','kacamata hitam legam','kacamata-hitam-legam','https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/19/5c501909-2134-4942-a316-2802656bda11.jpg','Ini kacamata hitam silau','CAT-DZsZDebw'),('PRD-acbd5c41-5807-414c-a9b8-bbe6c84c6899','Kacamata Modern','kacamata-modern','http://img.priceza.co.id/img2/2130004/0001/2130004-20220719222117-21964356142526946.jpg','Ini kacamata bukan kacamata biasa, tapi kacamata sangat gaul sekali','CAT-DZsZDebw'),('PRD-b730b694-6f8f-4df1-9664-6c5f35f1128f','Kacamata Sangat Modern','kacamata-sangat-modern','http://img.priceza.co.id/img2/2130004/0001/2130004-20220719222117-21964356142526946.jpg','Ini kacamata bukan kacamata biasa, tapi kacamata sangat gaul sekali','CAT-DZsZDebw'),('PRD-cd5bf8d6-416e-4e17-b8c3-25dcd24f68c0','Kacamata zaman now','kacamata-zaman-now','http://img.priceza.co.id/img2/2130004/0001/2130004-20220719222117-21964356142526946.jpg','Ini kacamata bukan kacamata biasa, tapi kacamata sangat gaul sekali','CAT-DZsZDebw'),('PROD-sLehZ9N4','Kacamata Gaul Sekali','kacamata-gaul-sekali','http://img.priceza.co.id/img2/2130004/0001/2130004-20220719222117-21964356142526946.jpg','Ini kacamata bukan kacamata biasa, tapi kacamata sangat gaul sekali','CAT-DZsZDebw');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `serial` varchar(50) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `password` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`serial`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('USR-vow5CJS4','mnindrazaka@gmail.com','mnindrazaka');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-01-10 22:47:55
