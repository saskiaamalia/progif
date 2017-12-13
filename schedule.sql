-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 13, 2017 at 08:50 AM
-- Server version: 10.1.28-MariaDB
-- PHP Version: 7.1.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `jadwal`
--

-- --------------------------------------------------------

--
-- Table structure for table `schedule`
--

CREATE TABLE `schedule` (
  `ID` int(10) NOT NULL,
  `Tanggal` date NOT NULL,
  `Kegiatan` varchar(100) NOT NULL,
  `Tempat` varchar(100) NOT NULL,
  `Keterangan` varchar(100) NOT NULL,
  `PIC` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `schedule`
--

INSERT INTO `schedule` (`ID`, `Tanggal`, `Kegiatan`, `Tempat`, `Keterangan`, `PIC`) VALUES
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(0, '0000-00-00', '', '', '', ''),
(1001, '2018-01-04', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1002, '2018-01-05', 'Syukuran Wisuda', 'Kantin Barak ITB', 'Acara internal dalam rangka merayakan wisudawan oktober', 'Ida\r'),
(1003, '2018-01-05', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1004, '2018-01-06', 'Sidang Anggota', 'Sekretariat MBWG', 'Untuk menentukan keberlanjutan unit', 'Qodri\r'),
(1005, '2018-01-08', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1006, '2018-01-09', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1007, '2018-01-10', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1008, '2018-01-11', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1009, '2018-01-12', 'Hearing Caketum', 'Sekretariat MBWG', 'Salah satu proses pemilihan ca-formatur baru', 'Nate\r'),
(1010, '2018-01-13', 'Latihan intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1011, '2018-01-14', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1012, '2018-01-15', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1013, '2018-01-16', 'Sosialisasi Pelantikan', 'Sekretariat MBWG', 'Sosialisasi untuk kegiatan pelantikan anggota baru', 'Zen\r'),
(1014, '2018-01-17', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1015, '2018-01-18', 'Hearing Caketum', 'Sekretariat MBWG', 'Salah satu proses pemilihan ca-formatur baru', 'Nate\r'),
(1016, '2018-01-19', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1017, '2018-01-21', 'Survei Pelantikan', 'Sekretariat MBWG', 'Survey lapangan untuk tempat pelantikan anggota baru', 'Zen\r'),
(1018, '2018-01-22', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1019, '2018-01-23', 'Sosialisasi Pelantikan', 'Sekretariat MBWG', 'Sosialisasi untuk kegiatan pelantikan anggota baru', 'Zen\r'),
(1020, '2018-01-24', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1021, '2018-01-25', 'Pemilihan Caketum', 'Sekretariat MBWG', 'Salah satu proses pemilihan ca-formatur baru', 'Nate\r'),
(1022, '2018-01-27', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1023, '2018-01-28', 'Survei Pelantikan', 'Sekretariat MBWG', 'Survey lapangan untuk tempat pelantikan anggota baru', 'Zen\r'),
(1024, '2018-01-29', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1025, '2018-01-30', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1026, '2018-01-31', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1027, '2018-02-01', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria\r'),
(1028, '2018-02-02', 'Penampilan Perdana Calon Anggota Baru', 'Lapangan Basket ITB', 'Penampilan perdana calon anggota baru', 'Zhafira\r'),
(1029, '2018-02-03', 'Pelantikan Anggota Baru', 'TBA', 'Proses pengangkatan calon anggota baru menjadi anggota biasa', 'Zen\r'),
(1030, '2018-02-04', 'Pelantikan Anggota Baru', 'TBA', 'Proses pengangkatan calon anggota baru menjadi anggota biasa', 'Zen'),
(0, '2017-12-08', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-08', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-08', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-08', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-09', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-09', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-09', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria'),
(0, '2017-12-10', 'Latihan Intensif', 'Lapangan Basket ITB', 'Latihan yang dilakukan untuk calon anggota baru', 'Satria');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
