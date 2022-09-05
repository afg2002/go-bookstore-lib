-- phpMyAdmin SQL Dump
-- version 5.1.1deb5ubuntu1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 05, 2022 at 09:40 PM
-- Server version: 8.0.30-0ubuntu0.22.04.1
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `perpustakaan`
--

-- --------------------------------------------------------

--
-- Table structure for table `buku`
--

CREATE TABLE `buku` (
  `id_buku` int NOT NULL,
  `cover_buku` blob NOT NULL,
  `judul` varchar(60) NOT NULL,
  `harga` int NOT NULL,
  `pengarang` varchar(40) NOT NULL,
  `kategori` varchar(30) NOT NULL,
  `penerbit` varchar(40) NOT NULL,
  `tahun` year NOT NULL,
  `stok` int NOT NULL,
  `deskripsi` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `buku`
--

INSERT INTO `buku` (`id_buku`, `cover_buku`, `judul`, `harga`, `pengarang`, `kategori`, `penerbit`, `tahun`, `stok`, `deskripsi`) VALUES
(31, 0x393738363032343234363934355f4c6175742d4265726365726974612e6a7067, 'Laut Bercerita', 75000, 'Leila S. Chudori', 'Fiksi-Sastra', 'Kepustakaan Populer Gramedia', 2017, 200, 'Laut Bercerita, novel terbaru Leila S. Chudori, bertutur tentang kisah keluarga yang kehilangan, sekumpulan sahabat yang merasakan kekosongan di dada, sekelompok orang yang gemar menyiksa dan lancar berkhianat, sejumlah keluarga yang mencari kejelasan makam anaknya, dan tentang cinta yang  tak akan luntu'),
(32, 0x393738363032303532333331365f4d656c616e676b61685f55565f53706f745f52342d312e6a7067, 'Melangkah', 65100, 'Js. Khairen', 'Novel', 'Gramedia Widiasarana Indonesia', 2020, 50, 'Listrik padam di seluruh Jawa dan Bali secara misterius! Ancaman nyata kekuatan baru yang hendak menaklukkan Nusantara. \r\n\r\nSaat yang sama, empat sahabat mendarat di Sumba, hanya untuk mendapati nasib ratusan juta manusia ada di tangan mereka! Empat mahasiswa ekonomi ini, harus bertarung melawan pasukan berkuda yang bisa melontarkan listrik! Semua dipersulit oleh seorang buronan tingkat tinggi bertopeng pahlawan yang punya rencana mengerikan. \r\n\r\nTernyata pesan arwah nenek moyang itu benar-benar terwujud. “Akan datang kegelapan yang berderap, bersama ribuan kuda raksasa di kala malam. Mereka bangun setelah sekian lama, untuk menghancurkan Nusantara. Seorang lelaki dan seorang perempuan ditakdirkan membaurkan air di lautan dan api di pegunungan. Menyatukan tanah yang menghujam, dan udara yang terhampar.” \r\n\r\nKisah tentang persahabatan, tentang jurang ego anak dan orangtua, tentang menyeimbangkan logika dan perasaan. Juga tentang melangkah menuju masa depan. Bahwa, apa pun yang menjadi luka masa lalu, biarlah mengering bersama waktu.'),
(33, 0x393738363032343831353530395f42656c616e746172612d312e6a7067, 'Belantara', 105000, 'Cixin Liu', 'Novel', 'Kepustakaan Populer Gramedia', 2022, 30, 'Apa yang bakal terjadi bila manusia tahu Bumi akan diinvasi alien empat abad lagi? \r\n\r\nSesudah mengetahui keberadaan Bumi, peradaban Trisurya mengirimkan armada penyerbu, dan pengintai berupa proton cerdas—sofon—yang bisa mengetahui semua informasi di Bumi kecuali apa yang ada di dalam pikiran manusia. Itulah dasar Proyek Penghadap Tembok, di mana sejumlah ahli siasat ditugasi untuk membuat strategi dalam kepala mereka sendiri tanpa bisa diketahui Trisurya. Sementara itu, Bumi harus membangun armada antariksa, tapi apakah manusia bisa mengatasi perpecahan antarnegara dan antarideologi untuk melakukannya? \r\n\r\nInilah langkah peradaban manusia dalam persiapan untuk Perang Terakhir.'),
(34, 0x393738363233303032393738335f4a756a757473756b616973656e5f352e6a7067, 'Jujutsu Kaisen 05', 32000, 'Gege Akutami', 'Manga', 'Elex Media Komputindo', 2022, 10, 'Program pertukaran dengan Akademi Jujutsu Kyoto dimulai. Pihak yang duluan menyingkirkan jurei tingkat 2 di area pertandinganlah yang akan jadi pemenangnya. Todo yang gemar berkelahi segera menyerang pihak Tokyo! Saat Todo dan Itadori saling berhadapan, anak-anak Kyoto yang lain ikut mengepung Itadori dengan niat untuk membunuhnya...!?'),
(35, 0x436f7665725f446570616e5f4a756a757473755f4b616973656e5f302e6a7067, 'Jujutsu Kaisen 0', 45000, 'Gege Akutami', 'Manga', 'Elex Media Komputindo', 2022, 40, 'Yuta Okkotsu seorang siswa SMA yang menginginkan hukuman mati untuk dirinya sendiri. Dia menderita karena Rika - roh pendendam yang menghantuinya. Namun, Satoru Gojo, seorang guru di “Akademi Jujutsu” - sekolah para shaman, meyakinkan Okkotsu untuk pindah ke sekolah tersebut. Inilah prekuel dari JUJUTSU KAISEN!');

-- --------------------------------------------------------

--
-- Table structure for table `cart`
--

CREATE TABLE `cart` (
  `id_cart` int NOT NULL,
  `id_user` int NOT NULL,
  `id_buku` int NOT NULL,
  `total` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `cart`
--

INSERT INTO `cart` (`id_cart`, `id_user`, `id_buku`, `total`) VALUES
(102, 18, 35, 1);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id_user` int NOT NULL,
  `email` varchar(40) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nama` varchar(40) NOT NULL,
  `role` enum('admin','anggota') NOT NULL,
  `jenis_kelamin` enum('L','P') NOT NULL,
  `no_telp` varchar(15) NOT NULL,
  `alamat` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id_user`, `email`, `password`, `nama`, `role`, `jenis_kelamin`, `no_telp`, `alamat`) VALUES
(2, 'afg2002@gmail.com', '$2a$10$.BAtHzmqWzrE8REWlyiPBeYzzLhjpQ7oLL69GxA9TAXE2mOGmOX1y', 'agan', 'admin', 'L', '+6285156283645', 'test'),
(18, 'rafisya@gmail.com', '$2a$10$BBiOjvHgT0SiDV.I/YPRQ.b0sVo2vRmk4A1afNCpDdkzpPLz70Mbe', 'Rafisya', 'admin', 'P', '+62111', '11');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `buku`
--
ALTER TABLE `buku`
  ADD PRIMARY KEY (`id_buku`);

--
-- Indexes for table `cart`
--
ALTER TABLE `cart`
  ADD PRIMARY KEY (`id_cart`),
  ADD KEY `id_buku` (`id_buku`),
  ADD KEY `id_user` (`id_user`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `buku`
--
ALTER TABLE `buku`
  MODIFY `id_buku` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT for table `cart`
--
ALTER TABLE `cart`
  MODIFY `id_cart` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=103;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id_user` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `cart`
--
ALTER TABLE `cart`
  ADD CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`) ON DELETE CASCADE,
  ADD CONSTRAINT `cart_ibfk_2` FOREIGN KEY (`id_buku`) REFERENCES `buku` (`id_buku`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
