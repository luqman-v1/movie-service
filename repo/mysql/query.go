package mysql

const (
	QueryUpSertMovie = `
        INSERT INTO 
		movies(
			title, 
			year,
			imdb_id,
			type,
			poster
			)
		VALUES( ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		imdb_id = VALUES(imdb_id); 
        `
)
