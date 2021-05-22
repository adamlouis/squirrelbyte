package kvserver

//go:generate go run ../../../../scripts/gencode/gencode.go --component server  --config ../../../../config/api.kv.yml  --package kvserver --out-dir . --out ./kvserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel
