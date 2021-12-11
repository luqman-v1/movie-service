CREATE TABLE `movies` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `title` varchar(255) NOT NULL,
                          `year` varchar(255) NOT NULL,
                          `imdb_id` varchar(255) NOT NULL,
                          `type` varchar(255) NOT NULL,
                          `poster` varchar(255) NOT NULL,
                          `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `movie_imdb_id_unique` (`imdb_id`),
                          KEY `movies_id_index` (`id`) USING BTREE,
                          KEY `movies_imdb_id_index` (`imdb_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1041 DEFAULT CHARSET=utf8;