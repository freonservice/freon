package api

//go:generate rm -rf frontend
//go:generate swagger generate server --api-package op --spec=swagger_frontend.yml --server-package frontend/restapi --model-package frontend/model --strict-responders --strict-additional-properties --keep-spec-order --principal github.com/freonservice/freon/internal/app.UserSession --exclude-main
//go:generate find frontend/restapi -maxdepth 1 -name "configure_*.go" -exec sed -i -e "/go:generate/d" {} ;
