package jobserver

//go:generate go run ../../../../scripts/gencode/gencode.go --component server  --config ../../../../config/api.job.yml  --package jobserver --out-dir . --out ./jobserver.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel
