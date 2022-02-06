package frontend

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/freonservice/freon/api/openapi/frontend/restapi"
	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
	"github.com/freonservice/freon/pkg/def"
	"github.com/freonservice/freon/pkg/netx"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/powerman/structlog"
	"github.com/sebest/xff"
)

type (
	Ctx = context.Context
	Log = *structlog.Logger

	Config struct {
		Addr     netx.Addr
		BasePath string
	}

	server struct {
		app app.Appl
		cfg Config
	}
)

func NewServer(auth app.Auth, appl app.Appl, cfg Config) (*restapi.Server, error) { //nolint:funlen
	srv := &server{
		app: appl,
		cfg: cfg,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("load embedded swagger spec: %w", err)
	}
	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := op.NewFreonFrontendAPI(swaggerSpec)
	api.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf
	api.JWTBearerAuth = auth.IsAuthorized

	api.HealthCheckHandler = op.HealthCheckHandlerFunc(srv.HealthCheck)
	api.LoginHandler = op.LoginHandlerFunc(srv.authorize)
	api.LogoutUserHandler = op.LogoutUserHandlerFunc(srv.logout)
	api.RegUserHandler = op.RegUserHandlerFunc(srv.regUser)
	api.UserMeHandler = op.UserMeHandlerFunc(srv.userMe)
	api.InfoHandler = op.InfoHandlerFunc(srv.info)
	api.UserChangePasswordHandler = op.UserChangePasswordHandlerFunc(srv.userChangePassword)
	api.UserChangeProfileHandler = op.UserChangeProfileHandlerFunc(srv.userChangeProfile)
	api.CreateLocalizationHandler = op.CreateLocalizationHandlerFunc(srv.createLocalization)
	api.ListLocalizationHandler = op.ListLocalizationHandlerFunc(srv.listLocalization)
	api.DeleteLocalizationHandler = op.DeleteLocalizationHandlerFunc(srv.deleteLocalization)
	api.CreateIdentifierHandler = op.CreateIdentifierHandlerFunc(srv.createIdentifier)
	api.ListIdentifiersHandler = op.ListIdentifiersHandlerFunc(srv.listIdentifiers)
	api.DeleteIdentifierHandler = op.DeleteIdentifierHandlerFunc(srv.deleteIdentifier)
	api.CreateCategoryHandler = op.CreateCategoryHandlerFunc(srv.createCategory)
	api.ListCategoriesHandler = op.ListCategoriesHandlerFunc(srv.listCategories)
	api.DeleteCategoryHandler = op.DeleteCategoryHandlerFunc(srv.deleteCategory)
	api.UpdateCategoryHandler = op.UpdateCategoryHandlerFunc(srv.updateCategory)
	api.UpdateIdentifierHandler = op.UpdateIdentifierHandlerFunc(srv.updateIdentifier)
	api.CreateTranslationHandler = op.CreateTranslationHandlerFunc(srv.createTranslation)
	api.UpdateTranslationHandler = op.UpdateTranslationHandlerFunc(srv.updateTranslation)
	api.DeleteTranslationHandler = op.DeleteTranslationHandlerFunc(srv.deleteTranslation)
	api.ListTranslationsHandler = op.ListTranslationsHandlerFunc(srv.listTranslations)
	api.StatusTranslationHandler = op.StatusTranslationHandlerFunc(srv.statusTranslation)
	api.ListUserHandler = op.ListUserHandlerFunc(srv.listUser)
	api.StatisticHandler = op.StatisticHandlerFunc(srv.statistic)
	api.CreateTranslationFilesHandler = op.CreateTranslationFilesHandlerFunc(srv.createTranslationFile)
	api.ListTranslationFilesHandler = op.ListTranslationFilesHandlerFunc(srv.listTranslationFiles)
	api.DeleteTranslationFileHandler = op.DeleteTranslationFileHandlerFunc(srv.deleteTranslationFile)
	api.VersionHandler = op.VersionHandlerFunc(srv.version)
	api.SettingsHandler = op.SettingsHandlerFunc(srv.settings)
	api.SettingTranslationHandler = op.SettingTranslationHandlerFunc(srv.settingTranslation)
	api.SettingStorageHandler = op.SettingStorageHandlerFunc(srv.settingStorage)
	api.SupportedLanguagesHandler = op.SupportedLanguagesHandlerFunc(srv.supportedLanguages)
	api.AutoTranslationHandler = op.AutoTranslationHandlerFunc(srv.autoTranslation)

	server := restapi.NewServer(api)
	server.Host = cfg.Addr.Host()
	server.Port = cfg.Addr.Port()

	api.UseSwaggerUI()
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(cfg.BasePath)
		accesslog := makeAccessLog()
		return noCache(xffmw.Handler(logger(recovery(accesslog(
			middleware.Spec(cfg.BasePath, restapi.FlatSwaggerJSON, cors(handler)))))))
	}

	middlewares := func(handler http.Handler) http.Handler {
		return handler
	}
	server.SetHandler(globalMiddlewares(api.Serve(middlewares)))

	log := structlog.New()
	log.Info("Frontend OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	return server, nil
}

func fromRequest(r *http.Request, auth *app.UserSession) (Ctx, Log) {
	ctx := r.Context()
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ctx = def.NewContextWithRemoteIP(ctx, remoteIP)
	var userID int64
	if auth != nil {
		userID = auth.UserID
	}
	log := structlog.FromContext(ctx, nil).SetDefaultKeyvals(
		def.LogUserID, userID,
		def.LogFunc, def.LogHandler,
		structlog.KeyTime, structlog.Auto,
		structlog.KeySource, structlog.Auto,
	)
	return ctx, log
}
