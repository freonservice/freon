package dal

const (
	sqlSelectActiveUserSession = `
		SELECT s.id, s.user_id, u.status 
		FROM user_sessions as s LEFT JOIN users u on s.user_id = u.id 
		WHERE s.token = $1 AND s.active = true LIMIT 1
	`
	// sqlSelectIdentifiers = `
	//	SELECT id, name, description, example_text, category_id FROM identifiers
	// `
	// sqlSelectIdentifiers = `
	//	SELECT id, name, description, example_text, category_id FROM identifiers ORDER BY id DESC
	// `
	sqlSelectCategories         = `SELECT id, name FROM categories ORDER BY id DESC`
	sqlSelectLocalizationListID = `SELECT id FROM localizations ORDER BY id DESC`
	sqlSelectIdentifierListID   = `SELECT id FROM identifiers WHERE status = $1 ORDER BY id DESC`
	sqlSelectUsers              = `SELECT id, uuid_id, email, first_name, second_name, status, role, created_at,
	updated_at FROM users ORDER BY id DESC`

	sqlUpdateUserSession           = `UPDATE user_sessions SET active = false WHERE token = $1`
	sqlUpdateNameCategory          = `UPDATE categories SET name = $1, updated_at = NOW() WHERE id = $2`
	sqlUpdateHideStatusTranslation = `UPDATE translations SET status = $1, updated_at = NOW() WHERE id = $2`

	sqlDeleteIdentifier   = `DELETE FROM identifiers WHERE id = $1`
	sqlDeleteLocalization = `DELETE FROM localizations WHERE id = $1`
	sqlDeleteCategory     = `DELETE FROM categories WHERE id = $1`
	sqlDeleteTranslation  = `DELETE FROM translations WHERE id = $1`

	sqlStatCountCategories    = `SELECT COUNT(*) FROM categories`
	sqlStatCountUsers         = `SELECT COUNT(*) FROM users WHERE status = $1`
	sqlStatCountLocalizations = `SELECT COUNT(*) FROM localizations WHERE status = $1`
	sqlStatCountIdentifiers   = `SELECT COUNT(*) FROM identifiers WHERE status = $1`
	sqlStatTranslations       = `
	SELECT COUNT(t.id) AS c, l.lang_name FROM translations AS t 
	JOIN localizations AS l ON t.localization_id=l.id 
	WHERE t.status = $1  
	GROUP BY l.lang_name
	`
)
