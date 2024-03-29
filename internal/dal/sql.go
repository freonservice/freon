package dal

const (
	sqlSelectActiveUserSession = `
		SELECT s.id, s.user_id, u.status 
		FROM user_sessions as s LEFT JOIN users u on s.user_id = u.id 
		WHERE s.token = $1 AND s.active = true LIMIT 1
	`
	sqlSelectCategories         = `SELECT id, name FROM categories ORDER BY id DESC`
	sqlSelectLocalizationListID = `SELECT id FROM localizations ORDER BY id DESC`
	sqlSelectIdentifierListID   = `SELECT id FROM identifiers WHERE status = $1 ORDER BY id DESC`
	sqlSelectUsers              = `SELECT id, uuid_id, email, first_name, second_name, status, role, created_at,
	updated_at FROM users ORDER BY id DESC`
	sqlSelectTranslationByID = `
	SELECT t.singular, t.plural, t.status, t.identifier_id, t.localization_id, l.locale FROM translations AS t 
	JOIN localizations AS l ON t.localization_id=l.id  
	WHERE t.id = $1   
	ORDER BY t.id DESC LIMIT 1`

	sqlUpdateUserSession       = `UPDATE user_sessions SET active = false WHERE token = $1`
	sqlUpdateNameCategory      = `UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2`
	sqlUpdateStatusTranslation = `UPDATE translations SET status = $1, updated_at = NOW() WHERE id = $2`
	sqlUpdateTranslation       = `
		UPDATE translations SET status = $1, singular = $2, plural = $3, updated_at = NOW() 
		WHERE localization_id = $4 AND identifier_id = $5
	`

	sqlDeleteIdentifier      = `DELETE FROM identifiers WHERE id = $1`
	sqlDeleteLocalization    = `DELETE FROM localizations WHERE id = $1`
	sqlDeleteCategory        = `DELETE FROM categories WHERE id = $1`
	sqlDeleteTranslation     = `DELETE FROM translations WHERE id = $1`
	sqlDeleteTranslationFile = `DELETE FROM translation_files WHERE id = $1`

	sqlStatCountCategories    = `SELECT COUNT(*) FROM categories`
	sqlStatCountUsers         = `SELECT COUNT(*) FROM users WHERE status = $1`
	sqlStatCountLocalizations = `SELECT COUNT(*) FROM localizations WHERE status = $1`
	sqlStatCountIdentifiers   = `SELECT COUNT(*) FROM identifiers WHERE status = $1`
	sqlStatTranslations       = `
	SELECT COUNT(CASE WHEN t.status = 0 THEN 1 END) AS f,
	l.lang_name FROM translations AS t  
	JOIN localizations AS l ON t.localization_id=l.id GROUP BY l.lang_name`
	sqlSelectTranslation = `
	SELECT t.singular, t.plural FROM translations AS t 
	JOIN localizations AS l ON t.localization_id=l.id  
	JOIN identifiers AS i ON t.identifier_id=i.id  
	WHERE l.locale = $1 AND i.name = $2   
	ORDER BY t.id DESC LIMIT 1`
)
