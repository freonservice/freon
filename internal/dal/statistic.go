package dal

import (
	"database/sql"

	"github.com/freonservice/freon/internal/dao"
	api "github.com/freonservice/freon/pkg/freonApi"
)

// count localizations
// count identifiers
// count users
// count completed translations by localizations
// count categories

func (r *Repo) GetStatistic(ctx Ctx) (*dao.Statistic, error) {
	var stat dao.Statistic
	_ = r.DB.QueryRowContext(ctx, sqlStatCountCategories).Scan(&stat.CountCategories)
	_ = r.DB.QueryRowContext(ctx, sqlStatCountUsers, api.Status_ACTIVE).Scan(&stat.CountUsers)
	_ = r.DB.QueryRowContext(ctx, sqlStatCountIdentifiers, api.Status_ACTIVE).Scan(&stat.CountIdentifiers)
	_ = r.DB.QueryRowContext(ctx, sqlStatCountLocalizations, api.Status_ACTIVE).Scan(&stat.CountLocalizations)

	rows, err := r.DB.QueryContext(ctx, sqlStatTranslations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*dao.StatTranslation
	for rows.Next() {
		var entity dao.StatTranslation
		if err = rows.Scan(&entity.Fulfilled, &entity.LangName); err != nil {
			break
		}
		stats = append(stats, &entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	stat.StatTranslations = stats
	return &stat, nil
}
