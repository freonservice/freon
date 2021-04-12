package frontend

import (
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
)

func (srv *server) statistic(params op.StatisticParams, session *app.UserSession) op.StatisticResponder {
	ctx, log := fromRequest(params.HTTPRequest, session)
	stat, err := srv.app.GetStatistic(ctx)
	switch {
	default:
		return errStatistic(log, err, codeInternal)
	case err == nil:
		return op.NewStatisticOK().WithPayload(apiStatistic(stat))
	}
}
