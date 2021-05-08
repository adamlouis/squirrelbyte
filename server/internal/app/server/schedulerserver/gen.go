package schedulerserver

//go:generate go run ../../../../scripts/gencode/gencode.go --component server  --config ../../../../config/api.scheduler.yml  --package schedulerserver --out-dir . --out ./schedulerserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel
