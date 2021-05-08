package oauthserver

//go:generate go run ../../../../scripts/gencode/gencode.go --component server  --config ../../../../config/api.oauth.yml  --package oauthserver --out-dir . --out ./oauthserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel
