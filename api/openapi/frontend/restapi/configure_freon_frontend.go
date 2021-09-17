// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/freonservice/freon/api/openapi/frontend/restapi/op"
	"github.com/freonservice/freon/internal/app"
)


func configureFlags(api *op.FreonFrontendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *op.FreonFrontendAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.JWTBearerAuth == nil {
		api.JWTBearerAuth = func(token string) (*app.UserSession, error) {
			return nil, errors.NotImplemented("api key auth (JWTBearer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.CreateCategoryHandler == nil {
		api.CreateCategoryHandler = op.CreateCategoryHandlerFunc(func(params op.CreateCategoryParams, principal *app.UserSession) op.CreateCategoryResponder {
			return op.CreateCategoryNotImplemented()
		})
	}
	if api.CreateIdentifierHandler == nil {
		api.CreateIdentifierHandler = op.CreateIdentifierHandlerFunc(func(params op.CreateIdentifierParams, principal *app.UserSession) op.CreateIdentifierResponder {
			return op.CreateIdentifierNotImplemented()
		})
	}
	if api.CreateLocalizationHandler == nil {
		api.CreateLocalizationHandler = op.CreateLocalizationHandlerFunc(func(params op.CreateLocalizationParams, principal *app.UserSession) op.CreateLocalizationResponder {
			return op.CreateLocalizationNotImplemented()
		})
	}
	if api.CreateTranslationHandler == nil {
		api.CreateTranslationHandler = op.CreateTranslationHandlerFunc(func(params op.CreateTranslationParams, principal *app.UserSession) op.CreateTranslationResponder {
			return op.CreateTranslationNotImplemented()
		})
	}
	if api.CreateTranslationFilesHandler == nil {
		api.CreateTranslationFilesHandler = op.CreateTranslationFilesHandlerFunc(func(params op.CreateTranslationFilesParams, principal *app.UserSession) op.CreateTranslationFilesResponder {
			return op.CreateTranslationFilesNotImplemented()
		})
	}
	if api.DeleteCategoryHandler == nil {
		api.DeleteCategoryHandler = op.DeleteCategoryHandlerFunc(func(params op.DeleteCategoryParams, principal *app.UserSession) op.DeleteCategoryResponder {
			return op.DeleteCategoryNotImplemented()
		})
	}
	if api.DeleteIdentifierHandler == nil {
		api.DeleteIdentifierHandler = op.DeleteIdentifierHandlerFunc(func(params op.DeleteIdentifierParams, principal *app.UserSession) op.DeleteIdentifierResponder {
			return op.DeleteIdentifierNotImplemented()
		})
	}
	if api.DeleteLocalizationHandler == nil {
		api.DeleteLocalizationHandler = op.DeleteLocalizationHandlerFunc(func(params op.DeleteLocalizationParams, principal *app.UserSession) op.DeleteLocalizationResponder {
			return op.DeleteLocalizationNotImplemented()
		})
	}
	if api.DeleteTranslationHandler == nil {
		api.DeleteTranslationHandler = op.DeleteTranslationHandlerFunc(func(params op.DeleteTranslationParams, principal *app.UserSession) op.DeleteTranslationResponder {
			return op.DeleteTranslationNotImplemented()
		})
	}
	if api.DeleteTranslationFileHandler == nil {
		api.DeleteTranslationFileHandler = op.DeleteTranslationFileHandlerFunc(func(params op.DeleteTranslationFileParams, principal *app.UserSession) op.DeleteTranslationFileResponder {
			return op.DeleteTranslationFileNotImplemented()
		})
	}
	if api.HealthCheckHandler == nil {
		api.HealthCheckHandler = op.HealthCheckHandlerFunc(func(params op.HealthCheckParams) op.HealthCheckResponder {
			return op.HealthCheckNotImplemented()
		})
	}
	if api.InfoHandler == nil {
		api.InfoHandler = op.InfoHandlerFunc(func(params op.InfoParams, principal *app.UserSession) op.InfoResponder {
			return op.InfoNotImplemented()
		})
	}
	if api.ListCategoriesHandler == nil {
		api.ListCategoriesHandler = op.ListCategoriesHandlerFunc(func(params op.ListCategoriesParams, principal *app.UserSession) op.ListCategoriesResponder {
			return op.ListCategoriesNotImplemented()
		})
	}
	if api.ListIdentifiersHandler == nil {
		api.ListIdentifiersHandler = op.ListIdentifiersHandlerFunc(func(params op.ListIdentifiersParams, principal *app.UserSession) op.ListIdentifiersResponder {
			return op.ListIdentifiersNotImplemented()
		})
	}
	if api.ListLocalizationHandler == nil {
		api.ListLocalizationHandler = op.ListLocalizationHandlerFunc(func(params op.ListLocalizationParams, principal *app.UserSession) op.ListLocalizationResponder {
			return op.ListLocalizationNotImplemented()
		})
	}
	if api.ListTranslationFilesHandler == nil {
		api.ListTranslationFilesHandler = op.ListTranslationFilesHandlerFunc(func(params op.ListTranslationFilesParams, principal *app.UserSession) op.ListTranslationFilesResponder {
			return op.ListTranslationFilesNotImplemented()
		})
	}
	if api.ListTranslationsHandler == nil {
		api.ListTranslationsHandler = op.ListTranslationsHandlerFunc(func(params op.ListTranslationsParams, principal *app.UserSession) op.ListTranslationsResponder {
			return op.ListTranslationsNotImplemented()
		})
	}
	if api.ListUserHandler == nil {
		api.ListUserHandler = op.ListUserHandlerFunc(func(params op.ListUserParams, principal *app.UserSession) op.ListUserResponder {
			return op.ListUserNotImplemented()
		})
	}
	if api.LoginHandler == nil {
		api.LoginHandler = op.LoginHandlerFunc(func(params op.LoginParams) op.LoginResponder {
			return op.LoginNotImplemented()
		})
	}
	if api.LogoutUserHandler == nil {
		api.LogoutUserHandler = op.LogoutUserHandlerFunc(func(params op.LogoutUserParams, principal *app.UserSession) op.LogoutUserResponder {
			return op.LogoutUserNotImplemented()
		})
	}
	if api.RegUserHandler == nil {
		api.RegUserHandler = op.RegUserHandlerFunc(func(params op.RegUserParams, principal *app.UserSession) op.RegUserResponder {
			return op.RegUserNotImplemented()
		})
	}
	if api.StatisticHandler == nil {
		api.StatisticHandler = op.StatisticHandlerFunc(func(params op.StatisticParams, principal *app.UserSession) op.StatisticResponder {
			return op.StatisticNotImplemented()
		})
	}
	if api.StatusTranslationHandler == nil {
		api.StatusTranslationHandler = op.StatusTranslationHandlerFunc(func(params op.StatusTranslationParams, principal *app.UserSession) op.StatusTranslationResponder {
			return op.StatusTranslationNotImplemented()
		})
	}
	if api.UpdateCategoryHandler == nil {
		api.UpdateCategoryHandler = op.UpdateCategoryHandlerFunc(func(params op.UpdateCategoryParams, principal *app.UserSession) op.UpdateCategoryResponder {
			return op.UpdateCategoryNotImplemented()
		})
	}
	if api.UpdateIdentifierHandler == nil {
		api.UpdateIdentifierHandler = op.UpdateIdentifierHandlerFunc(func(params op.UpdateIdentifierParams, principal *app.UserSession) op.UpdateIdentifierResponder {
			return op.UpdateIdentifierNotImplemented()
		})
	}
	if api.UpdateTranslationHandler == nil {
		api.UpdateTranslationHandler = op.UpdateTranslationHandlerFunc(func(params op.UpdateTranslationParams, principal *app.UserSession) op.UpdateTranslationResponder {
			return op.UpdateTranslationNotImplemented()
		})
	}
	if api.UserChangePasswordHandler == nil {
		api.UserChangePasswordHandler = op.UserChangePasswordHandlerFunc(func(params op.UserChangePasswordParams, principal *app.UserSession) op.UserChangePasswordResponder {
			return op.UserChangePasswordNotImplemented()
		})
	}
	if api.UserChangeProfileHandler == nil {
		api.UserChangeProfileHandler = op.UserChangeProfileHandlerFunc(func(params op.UserChangeProfileParams, principal *app.UserSession) op.UserChangeProfileResponder {
			return op.UserChangeProfileNotImplemented()
		})
	}
	if api.UserChangeStatusHandler == nil {
		api.UserChangeStatusHandler = op.UserChangeStatusHandlerFunc(func(params op.UserChangeStatusParams, principal *app.UserSession) op.UserChangeStatusResponder {
			return op.UserChangeStatusNotImplemented()
		})
	}
	if api.UserMeHandler == nil {
		api.UserMeHandler = op.UserMeHandlerFunc(func(params op.UserMeParams, principal *app.UserSession) op.UserMeResponder {
			return op.UserMeNotImplemented()
		})
	}
	if api.VersionHandler == nil {
		api.VersionHandler = op.VersionHandlerFunc(func(params op.VersionParams, principal *app.UserSession) op.VersionResponder {
			return op.VersionNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
