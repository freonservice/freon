package dal

import (
	"database/sql"

	"github.com/freonservice/freon/internal/dao"
	"github.com/freonservice/freon/pkg/api"
)

// count localizations
// count identifiers
// count users
// count completed translations by localizations
// count categories

func (r *r) GetStatistic(ctx Ctx) (*dao.Statistic, error) {
	var stat dao.Statistic
	_ = r.DB.QueryRowContext(ctx, sqlStatCountCategories).Scan(&stat.CountCategories)
	_ = r.DB.QueryRowContext(ctx, sqlStatCountUsers, api.Status_ACTIVE).Scan(&stat.CountUsers)
	_ = r.DB.QueryRowContext(ctx, sqlStatCountIdentifiers, api.Status_ACTIVE).Scan(&stat.CountIdentifiers)
	_ = r.DB.QueryRowContext(ctx, sqlStatCountLocalizations, api.Status_ACTIVE).Scan(&stat.CountLocalizations)

	rows, err := r.DB.QueryContext(ctx, sqlStatTranslations, api.Status_ACTIVE)
	if err != nil {
		return nil, err
	} else if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []*dao.StatTranslation
	for rows.Next() {
		var entity dao.StatTranslation
		if err = rows.Scan(&entity.Count, &entity.LangName); err != nil {
			break
		}
		stats = append(stats, &entity)
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	stat.StatTranslations = stats

	return &stat, nil
}
