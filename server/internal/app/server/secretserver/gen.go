package secretserver

//go:generate go run ../../../../scripts/gencode/gencode.go --component server  --config ../../../../config/api.secret.yml  --package secretserver --out-dir . --out ./secretserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/secretmodel
