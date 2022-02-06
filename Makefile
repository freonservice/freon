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
	GOOS=linux GOARCH=amd64 go build ${GO_BUILD_ARGS} -o ${PROJECTNAME} ./cmd/${PROJECTNAME}

genny-generate:
	genny -in=./internal/srv/frontend/error.go -out=./internal/srv/frontend/gen.error.go gen "HealthCheck=Login,LogoutUser,RegUser,CreateLocalization,UserMe,Info,ListLocalization,DeleteLocalization,CreateIdentifier,ListIdentifiers,DeleteIdentifier,CreateCategory,ListCategories,DeleteCategory,UpdateCategory,UpdateIdentifier,CreateTranslation,ListTranslations,DeleteTranslation,UpdateTranslation,StatusTranslation,UserChangePassword,UserChangeProfile,ListUser,Statistic,ListTranslationFiles,DeleteTranslationFile,Version,SettingTranslation,SettingStorage,AutoTranslation,SupportedLanguages"

mockgen-create:
	mockgen -package=app -source=./internal/app/app.go -destination=./internal/app/mock.go Appl,Auth,Repo,Password,SettingRepo
	mockgen -package=app -source=./internal/storage/storage.go -destination=./internal/app/storage.mock.go Storage

proto-generate:
	protoc -I./proto_files -I./vendor --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import,require_unimplemented_servers=false ./proto_files/const.proto
	protoc -I./proto_files -I./vendor --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import,require_unimplemented_servers=false ./proto_files/localization.proto
	protoc -I./proto_files -I./vendor --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import,require_unimplemented_servers=false ./proto_files/translation.proto
	protoc -I./proto_files -I./vendor --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import,require_unimplemented_servers=false ./proto_files/translation_file.proto
	protoc -I./proto_files -I./vendor --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import,require_unimplemented_servers=false ./proto_files/freon_service.proto

reform-generate:
	reform internal/dao