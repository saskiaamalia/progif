CREATE TABLE `Schedule` (
  `ID` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `Tanggal` varchar(30) NOT NULL,
  `Kegiatan` varchar(30) NOT NULL,
  `Tempat` varchar(30) NOT NULL,
  `Keterangan` varchar(30) NOT NULL,
  `PIC` varchar(30) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;