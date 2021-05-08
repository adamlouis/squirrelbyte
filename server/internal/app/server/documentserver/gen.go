package documentserver

//go:generate go run ../../../../scripts/gencode/gencode.go --component server  --config ../../../../config/api.document.yml  --package documentserver --out-dir . --out ./documentserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel
