-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 08, 2022 at 01:50 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kancha`
--

-- --------------------------------------------------------

--
-- Table structure for table `answers`
--

CREATE TABLE `answers` (
  `Ansid` int(11) NOT NULL,
  `Anstext` text NOT NULL,
  `AnsmemberID` int(11) NOT NULL,
  `QuestID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `answers`
--

INSERT INTO `answers` (`Ansid`, `Anstext`, `AnsmemberID`, `QuestID`) VALUES
(2, 'ไม่รู้ครับ', 1, 3),
(3, 'ไม่ทราบฮะ', 1, 3),
(5, 'ไม่รู้ๆๆๆๆๆๆๆๆๆๆๆๆๆๆ', 1, 3),
(6, 'testqqqqqq', 1, 4),
(7, 'ถามกูเกิ้ลเลย', 1, 3),
(8, 'ดูดตั้งแต่ตอนไหนหรอค่ะ', 2, 6),
(9, 'ไม่บอกๆ', 1, 6),
(11, 'test4', 1, 3),
(12, 'ไม่รู้ๆๆๆๆ', 3, 3),
(13, 'testtest5', 3, 3),
(14, 'test6', 1, 3),
(15, 'test7', 1, 3);

-- --------------------------------------------------------

--
-- Table structure for table `member`
--

CREATE TABLE `member` (
  `Memid` int(11) NOT NULL,
  `Fullname` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `Pwd` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `member`
--

INSERT INTO `member` (`Memid`, `Fullname`, `Email`, `Pwd`) VALUES
(1, 'khajornsak krongyut', 'khajornsak.kr.63@ubu.ac.th', '123456789'),
(2, 'sirikanya sukhum', 'sirikanya.su.63@ubu.ac.th', '12345678'),
(3, 'pipat unjit', 'pipat@gmail.com', '12345');

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `QuestID` int(11) NOT NULL,
  `QuestTitle` varchar(255) NOT NULL,
  `QuestDetail` text NOT NULL,
  `QuestMemberID` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`QuestID`, `QuestTitle`, `QuestDetail`, `QuestMemberID`) VALUES
(3, 'แต่ละสายพันธุ์ปลูกยังไงบ้างครับ', 'ผมไม่มีพื้นฐานความรู้เรื่องนี้มาก่อน', 1),
(4, 'test', 'test', 1),
(5, 'test2', 'test2', 1),
(6, 'ปลูกกัญชายังไงฮะ', 'ปลูกยังไงผมไม่รู้ ผมรู้แต่การดูด', 1),
(12, 'test3', '', 1),
(13, 'test4', 'test4', 1),
(14, 'test5', 'test5', 1),
(15, 'dfgdf', 'dfgdg', 1),
(16, 'sdfssdf', 'sdfsdfdds', 1),
(17, 'diw', 'diw', 1),
(18, 'พันธุ์ไรน่าปลูก', 'ช่วยบอกทีๆๆๆๆๆๆๆๆ', 1),
(19, 'พันธุ์ไรน่าปลูกสุด', 'บอกมาาาาาาาาาาาาาาาาาาา', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `answers`
--
ALTER TABLE `answers`
  ADD PRIMARY KEY (`Ansid`),
  ADD KEY `AnsmemberID` (`AnsmemberID`),
  ADD KEY `QuestID` (`QuestID`);

--
-- Indexes for table `member`
--
ALTER TABLE `member`
  ADD PRIMARY KEY (`Memid`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`QuestID`),
  ADD KEY `QuestMemberID` (`QuestMemberID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `answers`
--
ALTER TABLE `answers`
  MODIFY `Ansid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `member`
--
ALTER TABLE `member`
  MODIFY `Memid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `QuestID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `answers`
--
ALTER TABLE `answers`
  ADD CONSTRAINT `answers_ibfk_1` FOREIGN KEY (`AnsmemberID`) REFERENCES `member` (`Memid`),
  ADD CONSTRAINT `answers_ibfk_2` FOREIGN KEY (`QuestID`) REFERENCES `questions` (`QuestID`);

--
-- Constraints for table `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `questions_ibfk_1` FOREIGN KEY (`QuestMemberID`) REFERENCES `member` (`Memid`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
