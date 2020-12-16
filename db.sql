-- MariaDB dump 10.17  Distrib 10.5.4-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: cp_dev
-- ------------------------------------------------------
-- Server version	10.5.4-MariaDB-1:10.5.4+maria~bionic

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `pin`
--

DROP TABLE IF EXISTS `pin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `workbook_id` int(10) unsigned NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `title` varchar(256) NOT NULL,
  `description` varchar(256) NOT NULL,
  `filters` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`filters`)),
  `visualization_flag` tinyint(1) DEFAULT 0,
  `context` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`context`)),
  `creation_date` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id` (`id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 WITH SYSTEM VERSIONING;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pin`
--

LOCK TABLES `pin` WRITE;
/*!40000 ALTER TABLE `pin` DISABLE KEYS */;
INSERT INTO `pin` VALUES (3,1,'user_1','title_1','desc_1','{\"name\":\"Filter-1\", \"param_1\":1}',0,'{\"name\":\"Context-1\", \"param-1\":2}','2020-11-29 18:30:01'),(4,2,'user_1','title_2','desc_2','{\"name\":\"Filter-2\", \"param_1\":2}',0,'{\"name\":\"Context-2\", \"param-2\":2}','2020-11-29 18:30:02'),(5,3,'user_1','title_3','desc_3','{\"name\":\"Filter-3\", \"param_3\":2}',0,'{\"name\":\"Context-3\", \"param-3\":2}','2020-11-29 18:30:03'),(6,4,'user_1','title_4','desc_4','{\"name\":\"Filter-4\", \"param_4\":2}',0,'{\"name\":\"Context-4\", \"param-4\":2}','2020-11-29 18:30:04'),(7,5,'user_1','title_5','desc_5','{\"name\":\"Filter-5\", \"param_5\":2}',0,'{\"name\":\"Context-5\", \"param-5\":2}','2020-11-29 18:30:05');
/*!40000 ALTER TABLE `pin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `template`
--

DROP TABLE IF EXISTS `template`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `template` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  `definition` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`definition`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 WITH SYSTEM VERSIONING;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `template`
--

LOCK TABLES `template` WRITE;
/*!40000 ALTER TABLE `template` DISABLE KEYS */;
INSERT INTO `template` VALUES (1,'template-1','{\"name\":\"Definition-1\", \"param-1\":2}'),(2,'template-2','{\"name\":\"Definition-2\", \"param-1\":5}'),(3,'template-3','{\"name\":\"Definition-3\", \"param-1\":7}'),(4,'template-4','{\"name\":\"Definition-4\", \"param-1\":8}'),(5,'template-5','{\"name\":\"Definition-5\", \"param-1\":9}');
/*!40000 ALTER TABLE `template` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `workbook_id` int(10) unsigned NOT NULL,
  `role` varchar(256) NOT NULL,
  PRIMARY KEY (`id`,`workbook_id`),
  KEY `id` (`id`,`workbook_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 WITH SYSTEM VERSIONING;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('user_1',1,'Approver'),('user_1',2,'Approver'),('user_1',5,'Planner');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `workbook`
--

DROP TABLE IF EXISTS `workbook`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `workbook` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `template_id` tinyint(4) DEFAULT NULL,
  `da_dataset_id` int(10) unsigned NOT NULL,
  `scope` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`scope`)),
  `last_modified_by` varchar(36) NOT NULL,
  `last_modified` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `status` enum('PLANNING','REVIEWED','PUBLISHED','REJECTED') DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id` (`id`,`template_id`,`da_dataset_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 WITH SYSTEM VERSIONING;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workbook`
--

LOCK TABLES `workbook` WRITE;
/*!40000 ALTER TABLE `workbook` DISABLE KEYS */;
INSERT INTO `workbook` VALUES (1,1,1,'{\"name\":\"Scope-1\", \"param-1\":3}','user_1','2020-11-29 18:30:01','PLANNING'),(2,2,2,'{\"name\":\"Scope-2\", \"param-1\":13}','user_1','2020-11-29 18:30:02','PLANNING'),(3,3,3,'{\"name\":\"Scope-3\", \"param-1\":33}','user_1','2020-11-29 18:30:03','REVIEWED'),(4,4,4,'{\"name\":\"Scope-4\", \"param-1\":34}','user_1','2020-11-29 18:30:04','REVIEWED'),(5,5,5,'{\"name\":\"Scope-5\", \"param-1\":35}','user_1','2020-11-29 18:30:05','PUBLISHED');
/*!40000 ALTER TABLE `workbook` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `workbook_comments`
--

DROP TABLE IF EXISTS `workbook_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `workbook_comments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `workbook_id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL,
  `comment` varchar(256) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `id` (`id`,`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 WITH SYSTEM VERSIONING;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `workbook_comments`
--

LOCK TABLES `workbook_comments` WRITE;
/*!40000 ALTER TABLE `workbook_comments` DISABLE KEYS */;
INSERT INTO `workbook_comments` VALUES (1,'1','user_1','comment-1'),(2,'1','user_1','comment-2'),(3,'1','user_1','comment-3'),(4,'1','user_1','comment-4'),(5,'1','user_1','comment-5'),(6,'2','user_1','comment-6'),(7,'2','user_1','comment-7');
/*!40000 ALTER TABLE `workbook_comments` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-12-04  0:29:10
