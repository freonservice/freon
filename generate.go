package freon

//go:generate protoc @protoc.cfg ./proto_files/const.proto
//go:generate protoc @protoc.cfg ./proto_files/localization.proto
//go:generate protoc @protoc.cfg ./proto_files/translation.proto
//go:generate protoc @protoc.cfg ./proto_files/translation_file.proto

//go:generate protoc @protoc.cfg ./proto_files/freon_service.proto

//go:generate reform internal/dao
