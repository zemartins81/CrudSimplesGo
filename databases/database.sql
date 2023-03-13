CREATE TABLE `names` (
    `id` int(6) NOT NULL UNIQUE AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `email` varchar(50) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;