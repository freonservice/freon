PROJECTNAME := "freon"
PACKAGENAME := $(shell go list -m -f '{{.Path}}')
GITVERSION := $(shell git describe --dirty --always --tags --long)
DATE := $(shell date -u '+%Y-%m-%d-%H:%M UTC')
GO_BUILD_ARGS = \
  -gcflags "all=-trimpath=$(shell dirname $(shell pwd))" \
  -asmflags "all=-trimpath=$(shell dirname $(shell pwd))" \
  -ldflags " \
    -X '${PACKAGENAME}/pkg/version.gitVersion=${GITVERSION}' \
    -X '${PACKAGENAME}/pkg/version.buildDate=${DATE}' \
  " \

build-freon:
	go build ${GO_BUILD_ARGS} -o ${PROJECTNAME} ./cmd/${PROJECTNAME}

genny-generate:
	genny -in=./internal/srv/frontend/error.go -out=./internal/srv/frontend/gen.error.go gen "HealthCheck=Login,LogoutUser,RegUser,CreateLocalization,UserMe,Info,ListLocalization,DeleteLocalization,CreateIdentifier,ListIdentifiers,DeleteIdentifier,CreateCategory,ListCategories,DeleteCategory,UpdateCategory,UpdateIdentifier,CreateTranslation,ListTranslations,DeleteTranslation,UpdateTranslation,StatusTranslation,UserChangePassword,UserChangeProfile,ListUser,Statistic,ListTranslationFiles,DeleteTranslationFile,Version"

mockgen-create:
	mockgen -package=app -source=./internal/app/app.go -destination=./internal/app/mock.go Appl,Auth,Repo,Password
