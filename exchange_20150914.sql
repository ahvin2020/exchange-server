-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.6.17 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             9.2.0.4947
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping database structure for exchange
CREATE DATABASE IF NOT EXISTS `exchange` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `exchange`;


-- Dumping structure for table exchange.countries
CREATE TABLE IF NOT EXISTS `countries` (
  `id` int(3) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(60) DEFAULT NULL,
  `iso_alpha2` varchar(2) DEFAULT NULL,
  `iso_alpha3` varchar(3) DEFAULT NULL,
  `iso_numeric` int(11) DEFAULT NULL,
  `currency_code` char(3) DEFAULT NULL,
  `currency_name` varchar(32) DEFAULT NULL,
  `currrency_symbol` varchar(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=240 DEFAULT CHARSET=utf8;

-- Dumping data for table exchange.countries: ~239 rows (approximately)
DELETE FROM `countries`;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
INSERT INTO `countries` (`id`, `name`, `iso_alpha2`, `iso_alpha3`, `iso_numeric`, `currency_code`, `currency_name`, `currrency_symbol`) VALUES
	(1, 'Afghanistan', 'AF', 'AFG', 4, 'AFN', 'Afghani', '؋'),
	(2, 'Albania', 'AL', 'ALB', 8, 'ALL', 'Lek', 'Lek'),
	(3, 'Algeria', 'DZ', 'DZA', 12, 'DZD', 'Dinar', NULL),
	(4, 'American Samoa', 'AS', 'ASM', 16, 'USD', 'Dollar', '$'),
	(5, 'Andorra', 'AD', 'AND', 20, 'EUR', 'Euro', '€'),
	(6, 'Angola', 'AO', 'AGO', 24, 'AOA', 'Kwanza', 'Kz'),
	(7, 'Anguilla', 'AI', 'AIA', 660, 'XCD', 'Dollar', '$'),
	(8, 'Antarctica', 'AQ', 'ATA', 10, '', '', NULL),
	(9, 'Antigua and Barbuda', 'AG', 'ATG', 28, 'XCD', 'Dollar', '$'),
	(10, 'Argentina', 'AR', 'ARG', 32, 'ARS', 'Peso', '$'),
	(11, 'Armenia', 'AM', 'ARM', 51, 'AMD', 'Dram', NULL),
	(12, 'Aruba', 'AW', 'ABW', 533, 'AWG', 'Guilder', 'ƒ'),
	(13, 'Australia', 'AU', 'AUS', 36, 'AUD', 'Dollar', '$'),
	(14, 'Austria', 'AT', 'AUT', 40, 'EUR', 'Euro', '€'),
	(15, 'Azerbaijan', 'AZ', 'AZE', 31, 'AZN', 'Manat', 'ман'),
	(16, 'Bahamas', 'BS', 'BHS', 44, 'BSD', 'Dollar', '$'),
	(17, 'Bahrain', 'BH', 'BHR', 48, 'BHD', 'Dinar', NULL),
	(18, 'Bangladesh', 'BD', 'BGD', 50, 'BDT', 'Taka', NULL),
	(19, 'Barbados', 'BB', 'BRB', 52, 'BBD', 'Dollar', '$'),
	(20, 'Belarus', 'BY', 'BLR', 112, 'BYR', 'Ruble', 'p.'),
	(21, 'Belgium', 'BE', 'BEL', 56, 'EUR', 'Euro', '€'),
	(22, 'Belize', 'BZ', 'BLZ', 84, 'BZD', 'Dollar', 'BZ$'),
	(23, 'Benin', 'BJ', 'BEN', 204, 'XOF', 'Franc', NULL),
	(24, 'Bermuda', 'BM', 'BMU', 60, 'BMD', 'Dollar', '$'),
	(25, 'Bhutan', 'BT', 'BTN', 64, 'BTN', 'Ngultrum', NULL),
	(26, 'Bolivia', 'BO', 'BOL', 68, 'BOB', 'Boliviano', '$b'),
	(27, 'Bosnia and Herzegovina', 'BA', 'BIH', 70, 'BAM', 'Marka', 'KM'),
	(28, 'Botswana', 'BW', 'BWA', 72, 'BWP', 'Pula', 'P'),
	(29, 'Bouvet Island', 'BV', 'BVT', 74, 'NOK', 'Krone', 'kr'),
	(30, 'Brazil', 'BR', 'BRA', 76, 'BRL', 'Real', 'R$'),
	(31, 'British Indian Ocean Territory', 'IO', 'IOT', 86, 'USD', 'Dollar', '$'),
	(32, 'British Virgin Islands', 'VG', 'VGB', 92, 'USD', 'Dollar', '$'),
	(33, 'Brunei', 'BN', 'BRN', 96, 'BND', 'Dollar', '$'),
	(34, 'Bulgaria', 'BG', 'BGR', 100, 'BGN', 'Lev', 'лв'),
	(35, 'Burkina Faso', 'BF', 'BFA', 854, 'XOF', 'Franc', NULL),
	(36, 'Burundi', 'BI', 'BDI', 108, 'BIF', 'Franc', NULL),
	(37, 'Cambodia', 'KH', 'KHM', 116, 'KHR', 'Riels', '៛'),
	(38, 'Cameroon', 'CM', 'CMR', 120, 'XAF', 'Franc', 'FCF'),
	(39, 'Canada', 'CA', 'CAN', 124, 'CAD', 'Dollar', '$'),
	(40, 'Cape Verde', 'CV', 'CPV', 132, 'CVE', 'Escudo', NULL),
	(41, 'Cayman Islands', 'KY', 'CYM', 136, 'KYD', 'Dollar', '$'),
	(42, 'Central African Republic', 'CF', 'CAF', 140, 'XAF', 'Franc', 'FCF'),
	(43, 'Chad', 'TD', 'TCD', 148, 'XAF', 'Franc', NULL),
	(44, 'Chile', 'CL', 'CHL', 152, 'CLP', 'Peso', NULL),
	(45, 'China', 'CN', 'CHN', 156, 'CNY', 'Yuan Renminbi', '¥'),
	(46, 'Christmas Island', 'CX', 'CXR', 162, 'AUD', 'Dollar', '$'),
	(47, 'Cocos Islands', 'CC', 'CCK', 166, 'AUD', 'Dollar', '$'),
	(48, 'Colombia', 'CO', 'COL', 170, 'COP', 'Peso', '$'),
	(49, 'Comoros', 'KM', 'COM', 174, 'KMF', 'Franc', NULL),
	(50, 'Cook Islands', 'CK', 'COK', 184, 'NZD', 'Dollar', '$'),
	(51, 'Costa Rica', 'CR', 'CRI', 188, 'CRC', 'Colon', '₡'),
	(52, 'Croatia', 'HR', 'HRV', 191, 'HRK', 'Kuna', 'kn'),
	(53, 'Cuba', 'CU', 'CUB', 192, 'CUP', 'Peso', '₱'),
	(54, 'Cyprus', 'CY', 'CYP', 196, 'CYP', 'Pound', NULL),
	(55, 'Czech Republic', 'CZ', 'CZE', 203, 'CZK', 'Koruna', 'Kč'),
	(56, 'Democratic Republic of the Congo', 'CD', 'COD', 180, 'CDF', 'Franc', NULL),
	(57, 'Denmark', 'DK', 'DNK', 208, 'DKK', 'Krone', 'kr'),
	(58, 'Djibouti', 'DJ', 'DJI', 262, 'DJF', 'Franc', NULL),
	(59, 'Dominica', 'DM', 'DMA', 212, 'XCD', 'Dollar', '$'),
	(60, 'Dominican Republic', 'DO', 'DOM', 214, 'DOP', 'Peso', 'RD$'),
	(61, 'East Timor', 'TL', 'TLS', 626, 'USD', 'Dollar', '$'),
	(62, 'Ecuador', 'EC', 'ECU', 218, 'USD', 'Dollar', '$'),
	(63, 'Egypt', 'EG', 'EGY', 818, 'EGP', 'Pound', '£'),
	(64, 'El Salvador', 'SV', 'SLV', 222, 'SVC', 'Colone', '$'),
	(65, 'Equatorial Guinea', 'GQ', 'GNQ', 226, 'XAF', 'Franc', 'FCF'),
	(66, 'Eritrea', 'ER', 'ERI', 232, 'ERN', 'Nakfa', 'Nfk'),
	(67, 'Estonia', 'EE', 'EST', 233, 'EEK', 'Kroon', 'kr'),
	(68, 'Ethiopia', 'ET', 'ETH', 231, 'ETB', 'Birr', NULL),
	(69, 'Falkland Islands', 'FK', 'FLK', 238, 'FKP', 'Pound', '£'),
	(70, 'Faroe Islands', 'FO', 'FRO', 234, 'DKK', 'Krone', 'kr'),
	(71, 'Fiji', 'FJ', 'FJI', 242, 'FJD', 'Dollar', '$'),
	(72, 'Finland', 'FI', 'FIN', 246, 'EUR', 'Euro', '€'),
	(73, 'France', 'FR', 'FRA', 250, 'EUR', 'Euro', '€'),
	(74, 'French Guiana', 'GF', 'GUF', 254, 'EUR', 'Euro', '€'),
	(75, 'French Polynesia', 'PF', 'PYF', 258, 'XPF', 'Franc', NULL),
	(76, 'French Southern Territories', 'TF', 'ATF', 260, 'EUR', 'Euro  ', '€'),
	(77, 'Gabon', 'GA', 'GAB', 266, 'XAF', 'Franc', 'FCF'),
	(78, 'Gambia', 'GM', 'GMB', 270, 'GMD', 'Dalasi', 'D'),
	(79, 'Georgia', 'GE', 'GEO', 268, 'GEL', 'Lari', NULL),
	(80, 'Germany', 'DE', 'DEU', 276, 'EUR', 'Euro', '€'),
	(81, 'Ghana', 'GH', 'GHA', 288, 'GHC', 'Cedi', '¢'),
	(82, 'Gibraltar', 'GI', 'GIB', 292, 'GIP', 'Pound', '£'),
	(83, 'Greece', 'GR', 'GRC', 300, 'EUR', 'Euro', '€'),
	(84, 'Greenland', 'GL', 'GRL', 304, 'DKK', 'Krone', 'kr'),
	(85, 'Grenada', 'GD', 'GRD', 308, 'XCD', 'Dollar', '$'),
	(86, 'Guadeloupe', 'GP', 'GLP', 312, 'EUR', 'Euro', '€'),
	(87, 'Guam', 'GU', 'GUM', 316, 'USD', 'Dollar', '$'),
	(88, 'Guatemala', 'GT', 'GTM', 320, 'GTQ', 'Quetzal', 'Q'),
	(89, 'Guinea', 'GN', 'GIN', 324, 'GNF', 'Franc', NULL),
	(90, 'Guinea-Bissau', 'GW', 'GNB', 624, 'XOF', 'Franc', NULL),
	(91, 'Guyana', 'GY', 'GUY', 328, 'GYD', 'Dollar', '$'),
	(92, 'Haiti', 'HT', 'HTI', 332, 'HTG', 'Gourde', 'G'),
	(93, 'Heard Island and McDonald Islands', 'HM', 'HMD', 334, 'AUD', 'Dollar', '$'),
	(94, 'Honduras', 'HN', 'HND', 340, 'HNL', 'Lempira', 'L'),
	(95, 'Hong Kong', 'HK', 'HKG', 344, 'HKD', 'Dollar', '$'),
	(96, 'Hungary', 'HU', 'HUN', 348, 'HUF', 'Forint', 'Ft'),
	(97, 'Iceland', 'IS', 'ISL', 352, 'ISK', 'Krona', 'kr'),
	(98, 'India', 'IN', 'IND', 356, 'INR', 'Rupee', '₹'),
	(99, 'Indonesia', 'ID', 'IDN', 360, 'IDR', 'Rupiah', 'Rp'),
	(100, 'Iran', 'IR', 'IRN', 364, 'IRR', 'Rial', '﷼'),
	(101, 'Iraq', 'IQ', 'IRQ', 368, 'IQD', 'Dinar', NULL),
	(102, 'Ireland', 'IE', 'IRL', 372, 'EUR', 'Euro', '€'),
	(103, 'Israel', 'IL', 'ISR', 376, 'ILS', 'Shekel', '₪'),
	(104, 'Italy', 'IT', 'ITA', 380, 'EUR', 'Euro', '€'),
	(105, 'Ivory Coast', 'CI', 'CIV', 384, 'XOF', 'Franc', NULL),
	(106, 'Jamaica', 'JM', 'JAM', 388, 'JMD', 'Dollar', '$'),
	(107, 'Japan', 'JP', 'JPN', 392, 'JPY', 'Yen', '¥'),
	(108, 'Jordan', 'JO', 'JOR', 400, 'JOD', 'Dinar', NULL),
	(109, 'Kazakhstan', 'KZ', 'KAZ', 398, 'KZT', 'Tenge', 'лв'),
	(110, 'Kenya', 'KE', 'KEN', 404, 'KES', 'Shilling', NULL),
	(111, 'Kiribati', 'KI', 'KIR', 296, 'AUD', 'Dollar', '$'),
	(112, 'Kuwait', 'KW', 'KWT', 414, 'KWD', 'Dinar', NULL),
	(113, 'Kyrgyzstan', 'KG', 'KGZ', 417, 'KGS', 'Som', 'лв'),
	(114, 'Laos', 'LA', 'LAO', 418, 'LAK', 'Kip', '₭'),
	(115, 'Latvia', 'LV', 'LVA', 428, 'LVL', 'Lat', 'Ls'),
	(116, 'Lebanon', 'LB', 'LBN', 422, 'LBP', 'Pound', '£'),
	(117, 'Lesotho', 'LS', 'LSO', 426, 'LSL', 'Loti', 'L'),
	(118, 'Liberia', 'LR', 'LBR', 430, 'LRD', 'Dollar', '$'),
	(119, 'Libya', 'LY', 'LBY', 434, 'LYD', 'Dinar', NULL),
	(120, 'Liechtenstein', 'LI', 'LIE', 438, 'CHF', 'Franc', 'CHF'),
	(121, 'Lithuania', 'LT', 'LTU', 440, 'LTL', 'Litas', 'Lt'),
	(122, 'Luxembourg', 'LU', 'LUX', 442, 'EUR', 'Euro', '€'),
	(123, 'Macao', 'MO', 'MAC', 446, 'MOP', 'Pataca', 'MOP'),
	(124, 'Macedonia', 'MK', 'MKD', 807, 'MKD', 'Denar', 'ден'),
	(125, 'Madagascar', 'MG', 'MDG', 450, 'MGA', 'Ariary', NULL),
	(126, 'Malawi', 'MW', 'MWI', 454, 'MWK', 'Kwacha', 'MK'),
	(127, 'Malaysia', 'MY', 'MYS', 458, 'MYR', 'Ringgit', 'RM'),
	(128, 'Maldives', 'MV', 'MDV', 462, 'MVR', 'Rufiyaa', 'Rf'),
	(129, 'Mali', 'ML', 'MLI', 466, 'XOF', 'Franc', NULL),
	(130, 'Malta', 'MT', 'MLT', 470, 'MTL', 'Lira', NULL),
	(131, 'Marshall Islands', 'MH', 'MHL', 584, 'USD', 'Dollar', '$'),
	(132, 'Martinique', 'MQ', 'MTQ', 474, 'EUR', 'Euro', '€'),
	(133, 'Mauritania', 'MR', 'MRT', 478, 'MRO', 'Ouguiya', 'UM'),
	(134, 'Mauritius', 'MU', 'MUS', 480, 'MUR', 'Rupee', '₨'),
	(135, 'Mayotte', 'YT', 'MYT', 175, 'EUR', 'Euro', '€'),
	(136, 'Mexico', 'MX', 'MEX', 484, 'MXN', 'Peso', '$'),
	(137, 'Micronesia', 'FM', 'FSM', 583, 'USD', 'Dollar', '$'),
	(138, 'Moldova', 'MD', 'MDA', 498, 'MDL', 'Leu', NULL),
	(139, 'Monaco', 'MC', 'MCO', 492, 'EUR', 'Euro', '€'),
	(140, 'Mongolia', 'MN', 'MNG', 496, 'MNT', 'Tugrik', '₮'),
	(141, 'Montserrat', 'MS', 'MSR', 500, 'XCD', 'Dollar', '$'),
	(142, 'Morocco', 'MA', 'MAR', 504, 'MAD', 'Dirham', NULL),
	(143, 'Mozambique', 'MZ', 'MOZ', 508, 'MZN', 'Meticail', 'MT'),
	(144, 'Myanmar', 'MM', 'MMR', 104, 'MMK', 'Kyat', 'K'),
	(145, 'Namibia', 'NA', 'NAM', 516, 'NAD', 'Dollar', '$'),
	(146, 'Nauru', 'NR', 'NRU', 520, 'AUD', 'Dollar', '$'),
	(147, 'Nepal', 'NP', 'NPL', 524, 'NPR', 'Rupee', '₨'),
	(148, 'Netherlands', 'NL', 'NLD', 528, 'EUR', 'Euro', '€'),
	(149, 'Netherlands Antilles', 'AN', 'ANT', 530, 'ANG', 'Guilder', 'ƒ'),
	(150, 'New Caledonia', 'NC', 'NCL', 540, 'XPF', 'Franc', NULL),
	(151, 'New Zealand', 'NZ', 'NZL', 554, 'NZD', 'Dollar', '$'),
	(152, 'Nicaragua', 'NI', 'NIC', 558, 'NIO', 'Cordoba', 'C$'),
	(153, 'Niger', 'NE', 'NER', 562, 'XOF', 'Franc', NULL),
	(154, 'Nigeria', 'NG', 'NGA', 566, 'NGN', 'Naira', '₦'),
	(155, 'Niue', 'NU', 'NIU', 570, 'NZD', 'Dollar', '$'),
	(156, 'Norfolk Island', 'NF', 'NFK', 574, 'AUD', 'Dollar', '$'),
	(157, 'North Korea', 'KP', 'PRK', 408, 'KPW', 'Won', '₩'),
	(158, 'Northern Mariana Islands', 'MP', 'MNP', 580, 'USD', 'Dollar', '$'),
	(159, 'Norway', 'NO', 'NOR', 578, 'NOK', 'Krone', 'kr'),
	(160, 'Oman', 'OM', 'OMN', 512, 'OMR', 'Rial', '﷼'),
	(161, 'Pakistan', 'PK', 'PAK', 586, 'PKR', 'Rupee', '₨'),
	(162, 'Palau', 'PW', 'PLW', 585, 'USD', 'Dollar', '$'),
	(163, 'Palestinian Territory', 'PS', 'PSE', 275, 'ILS', 'Shekel', '₪'),
	(164, 'Panama', 'PA', 'PAN', 591, 'PAB', 'Balboa', 'B/.'),
	(165, 'Papua New Guinea', 'PG', 'PNG', 598, 'PGK', 'Kina', NULL),
	(166, 'Paraguay', 'PY', 'PRY', 600, 'PYG', 'Guarani', 'Gs'),
	(167, 'Peru', 'PE', 'PER', 604, 'PEN', 'Sol', 'S/.'),
	(168, 'Philippines', 'PH', 'PHL', 608, 'PHP', 'Peso', 'Php'),
	(169, 'Pitcairn', 'PN', 'PCN', 612, 'NZD', 'Dollar', '$'),
	(170, 'Poland', 'PL', 'POL', 616, 'PLN', 'Zloty', 'zł'),
	(171, 'Portugal', 'PT', 'PRT', 620, 'EUR', 'Euro', '€'),
	(172, 'Puerto Rico', 'PR', 'PRI', 630, 'USD', 'Dollar', '$'),
	(173, 'Qatar', 'QA', 'QAT', 634, 'QAR', 'Rial', '﷼'),
	(174, 'Republic of the Congo', 'CG', 'COG', 178, 'XAF', 'Franc', 'FCF'),
	(175, 'Reunion', 'RE', 'REU', 638, 'EUR', 'Euro', '€'),
	(176, 'Romania', 'RO', 'ROU', 642, 'RON', 'Leu', 'lei'),
	(177, 'Russia', 'RU', 'RUS', 643, 'RUB', 'Ruble', 'руб'),
	(178, 'Rwanda', 'RW', 'RWA', 646, 'RWF', 'Franc', NULL),
	(179, 'Saint Helena', 'SH', 'SHN', 654, 'SHP', 'Pound', '£'),
	(180, 'Saint Kitts and Nevis', 'KN', 'KNA', 659, 'XCD', 'Dollar', '$'),
	(181, 'Saint Lucia', 'LC', 'LCA', 662, 'XCD', 'Dollar', '$'),
	(182, 'Saint Pierre and Miquelon', 'PM', 'SPM', 666, 'EUR', 'Euro', '€'),
	(183, 'Saint Vincent and the Grenadines', 'VC', 'VCT', 670, 'XCD', 'Dollar', '$'),
	(184, 'Samoa', 'WS', 'WSM', 882, 'WST', 'Tala', 'WS$'),
	(185, 'San Marino', 'SM', 'SMR', 674, 'EUR', 'Euro', '€'),
	(186, 'Sao Tome and Principe', 'ST', 'STP', 678, 'STD', 'Dobra', 'Db'),
	(187, 'Saudi Arabia', 'SA', 'SAU', 682, 'SAR', 'Rial', '﷼'),
	(188, 'Senegal', 'SN', 'SEN', 686, 'XOF', 'Franc', NULL),
	(189, 'Serbia and Montenegro', 'CS', 'SCG', 891, 'RSD', 'Dinar', 'Дин'),
	(190, 'Seychelles', 'SC', 'SYC', 690, 'SCR', 'Rupee', '₨'),
	(191, 'Sierra Leone', 'SL', 'SLE', 694, 'SLL', 'Leone', 'Le'),
	(192, 'Singapore', 'SG', 'SGP', 702, 'SGD', 'Dollar', '$'),
	(193, 'Slovakia', 'SK', 'SVK', 703, 'SKK', 'Koruna', 'Sk'),
	(194, 'Slovenia', 'SI', 'SVN', 705, 'EUR', 'Euro', '€'),
	(195, 'Solomon Islands', 'SB', 'SLB', 90, 'SBD', 'Dollar', '$'),
	(196, 'Somalia', 'SO', 'SOM', 706, 'SOS', 'Shilling', 'S'),
	(197, 'South Africa', 'ZA', 'ZAF', 710, 'ZAR', 'Rand', 'R'),
	(198, 'South Georgia and the South Sandwich Islands', 'GS', 'SGS', 239, 'GBP', 'Pound', '£'),
	(199, 'South Korea', 'KR', 'KOR', 410, 'KRW', 'Won', '₩'),
	(200, 'Spain', 'ES', 'ESP', 724, 'EUR', 'Euro', '€'),
	(201, 'Sri Lanka', 'LK', 'LKA', 144, 'LKR', 'Rupee', '₨'),
	(202, 'Sudan', 'SD', 'SDN', 736, 'SDD', 'Dinar', NULL),
	(203, 'Suriname', 'SR', 'SUR', 740, 'SRD', 'Dollar', '$'),
	(204, 'Svalbard and Jan Mayen', 'SJ', 'SJM', 744, 'NOK', 'Krone', 'kr'),
	(205, 'Swaziland', 'SZ', 'SWZ', 748, 'SZL', 'Lilangeni', NULL),
	(206, 'Sweden', 'SE', 'SWE', 752, 'SEK', 'Krona', 'kr'),
	(207, 'Switzerland', 'CH', 'CHE', 756, 'CHF', 'Franc', 'CHF'),
	(208, 'Syria', 'SY', 'SYR', 760, 'SYP', 'Pound', '£'),
	(209, 'Taiwan', 'TW', 'TWN', 158, 'TWD', 'Dollar', 'NT$'),
	(210, 'Tajikistan', 'TJ', 'TJK', 762, 'TJS', 'Somoni', NULL),
	(211, 'Tanzania', 'TZ', 'TZA', 834, 'TZS', 'Shilling', NULL),
	(212, 'Thailand', 'TH', 'THA', 764, 'THB', 'Baht', '฿'),
	(213, 'Togo', 'TG', 'TGO', 768, 'XOF', 'Franc', NULL),
	(214, 'Tokelau', 'TK', 'TKL', 772, 'NZD', 'Dollar', '$'),
	(215, 'Tonga', 'TO', 'TON', 776, 'TOP', 'Pa\'anga', 'T$'),
	(216, 'Trinidad and Tobago', 'TT', 'TTO', 780, 'TTD', 'Dollar', 'TT$'),
	(217, 'Tunisia', 'TN', 'TUN', 788, 'TND', 'Dinar', NULL),
	(218, 'Turkey', 'TR', 'TUR', 792, 'TRY', 'Lira', 'YTL'),
	(219, 'Turkmenistan', 'TM', 'TKM', 795, 'TMM', 'Manat', 'm'),
	(220, 'Turks and Caicos Islands', 'TC', 'TCA', 796, 'USD', 'Dollar', '$'),
	(221, 'Tuvalu', 'TV', 'TUV', 798, 'AUD', 'Dollar', '$'),
	(222, 'U.S. Virgin Islands', 'VI', 'VIR', 850, 'USD', 'Dollar', '$'),
	(223, 'Uganda', 'UG', 'UGA', 800, 'UGX', 'Shilling', NULL),
	(224, 'Ukraine', 'UA', 'UKR', 804, 'UAH', 'Hryvnia', '₴'),
	(225, 'United Arab Emirates', 'AE', 'ARE', 784, 'AED', 'Dirham', NULL),
	(226, 'United Kingdom', 'GB', 'GBR', 826, 'GBP', 'Pound', '£'),
	(227, 'United States', 'US', 'USA', 840, 'USD', 'Dollar', '$'),
	(228, 'United States Minor Outlying Islands', 'UM', 'UMI', 581, 'USD', 'Dollar ', '$'),
	(229, 'Uruguay', 'UY', 'URY', 858, 'UYU', 'Peso', '$U'),
	(230, 'Uzbekistan', 'UZ', 'UZB', 860, 'UZS', 'Som', 'лв'),
	(231, 'Vanuatu', 'VU', 'VUT', 548, 'VUV', 'Vatu', 'Vt'),
	(232, 'Vatican', 'VA', 'VAT', 336, 'EUR', 'Euro', '€'),
	(233, 'Venezuela', 'VE', 'VEN', 862, 'VEF', 'Bolivar', 'Bs'),
	(234, 'Vietnam', 'VN', 'VNM', 704, 'VND', 'Dong', '₫'),
	(235, 'Wallis and Futuna', 'WF', 'WLF', 876, 'XPF', 'Franc', NULL),
	(236, 'Western Sahara', 'EH', 'ESH', 732, 'MAD', 'Dirham', NULL),
	(237, 'Yemen', 'YE', 'YEM', 887, 'YER', 'Rial', '﷼'),
	(238, 'Zambia', 'ZM', 'ZMB', 894, 'ZMK', 'Kwacha', 'ZK'),
	(239, 'Zimbabwe', 'ZW', 'ZWE', 716, 'ZWD', 'Dollar', 'Z$');
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;


-- Dumping structure for table exchange.reset_password_tickets
CREATE TABLE IF NOT EXISTS `reset_password_tickets` (
  `user_id` bigint(10) unsigned NOT NULL,
  `token_hash` varchar(100) NOT NULL,
  `token_expiry` datetime NOT NULL,
  `token_used` tinyint(1) NOT NULL,
  UNIQUE KEY `username` (`user_id`),
  KEY `token_hash` (`token_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table exchange.reset_password_tickets: ~1 rows (approximately)
DELETE FROM `reset_password_tickets`;
/*!40000 ALTER TABLE `reset_password_tickets` DISABLE KEYS */;
INSERT INTO `reset_password_tickets` (`user_id`, `token_hash`, `token_expiry`, `token_used`) VALUES
	(1, '5627ef8abd54ed4895cc00d2709a1fb5', '2015-08-31 20:47:24', 1);
/*!40000 ALTER TABLE `reset_password_tickets` ENABLE KEYS */;


-- Dumping structure for table exchange.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `gender` varchar(1) NOT NULL,
  `birthday` date NOT NULL,
  `country_id` int(3) unsigned NOT NULL,
  `bio` text NOT NULL,
  `profile_pic` varchar(50) NOT NULL,
  `token` varchar(50) NOT NULL,
  `token_expiry` datetime NOT NULL,
  `create_ip` varchar(15) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_users_countries` (`country_id`),
  KEY `username` (`username`),
  KEY `token` (`token`),
  CONSTRAINT `FK_users_countries` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- Dumping data for table exchange.users: ~3 rows (approximately)
DELETE FROM `users`;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `email`, `password`, `first_name`, `last_name`, `gender`, `birthday`, `country_id`, `bio`, `profile_pic`, `token`, `token_expiry`, `create_ip`, `created`) VALUES
	(1, 'admin', 'ahvin2020@gmail.com', '$2a$10$0n2hlMmFNhAnkL6NVEvINONtNLp5BdfLG2n/RArdwUWqsGV.azjW.', 'Admin', 'Tan', 'O', '1987-12-17', 192, 'pikachu', 'admin_5sx0hzoPTZyG03NK.jpg', '16k1ZLRZx01odTVCnqpdgWS81i0KXGxJ', '2015-09-14 20:29:54', '123.4.4', '2014-12-24 00:00:00'),
	(2, 'ahvin2020', 'ahvin2020@gmail.com', '$2a$10$0n2hlMmFNhAnkL6NVEvINONtNLp5BdfLG2n/RArdwUWqsGV.azjW.', 'Ahvin', 'Tan', 'O', '1987-12-17', 192, 'pikachu', 'admin_5sx0hzoPTZyG03NK.jpg', 'AZ3ww4hIHPIi4mOUkIYVWeKAyF0e2k38', '2015-09-14 22:24:42', '123.4.4', '2014-12-24 00:00:00'),
	(3, 'test', 'ahvin2020@gmail.com', '$2a$10$0n2hlMmFNhAnkL6NVEvINONtNLp5BdfLG2n/RArdwUWqsGV.azjW.', 'Ahvin', 'Tan', 'O', '1987-12-17', 192, 'pikachu', 'admin_5sx0hzoPTZyG03NK.jpg', 'AZ3ww4hIHPIi4mOUkIYVWeKAyF0e2k38', '2015-09-14 20:33:39', '123.4.4', '2014-12-24 00:00:00');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;


-- Dumping structure for table exchange.user_currencies
CREATE TABLE IF NOT EXISTS `user_currencies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `sell_amount` decimal(10,2) NOT NULL,
  `sell_currency` char(3) NOT NULL,
  `buy_amount` decimal(10,2) NOT NULL,
  `buy_currency` char(3) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

-- Dumping data for table exchange.user_currencies: ~6 rows (approximately)
DELETE FROM `user_currencies`;
/*!40000 ALTER TABLE `user_currencies` DISABLE KEYS */;
INSERT INTO `user_currencies` (`id`, `user_id`, `sell_amount`, `sell_currency`, `buy_amount`, `buy_currency`) VALUES
	(1, 1, 5656.00, 'sel', 4545.00, 'buy'),
	(2, 1, 5000.00, 'MYR', 333.00, 'SGD'),
	(6, 0, 22.00, 'bb', 12212.00, 'aa'),
	(8, 0, 44.00, 'dre', 33.00, 'se'),
	(9, 0, 2.00, '34', 1.00, '2'),
	(10, 1, 3.00, '4', 1.00, '2');
/*!40000 ALTER TABLE `user_currencies` ENABLE KEYS */;


-- Dumping structure for table exchange.user_follows
CREATE TABLE IF NOT EXISTS `user_follows` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `follower_user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_follower_user_id` (`user_id`,`follower_user_id`),
  KEY `user_id` (`user_id`),
  KEY `follower_user_id` (`follower_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;

-- Dumping data for table exchange.user_follows: ~3 rows (approximately)
DELETE FROM `user_follows`;
/*!40000 ALTER TABLE `user_follows` DISABLE KEYS */;
INSERT INTO `user_follows` (`id`, `user_id`, `follower_user_id`) VALUES
	(10, 1, 2),
	(7, 1, 3),
	(6, 3, 1);
/*!40000 ALTER TABLE `user_follows` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
