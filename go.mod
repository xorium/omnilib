module github.com/xorium/omnilib

go 1.15

require (
	github.com/andybalholm/brotli v1.0.1 // indirect
	github.com/google/jsonapi v0.0.0-20201022225600-f822737867f6
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/savsgio/gotils v0.0.0-20210120114113-f9d780dcbd93
	github.com/stretchr/testify v1.6.1
	github.com/valyala/fasthttp v1.19.0
	gitlab.omnicube.ru/omnicube/omninanage v0.0.0-00010101000000-000000000000
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0

)

replace gitlab.omnicube.ru/omnicube/omninanage => ../omninanage
