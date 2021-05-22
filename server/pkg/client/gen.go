package model

//go:generate go run ../../scripts/gencode/gencode.go --component client --config ../../config/api.document.yml --package documentclient --out ./documentclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/documentmodel
//go:generate go run ../../scripts/gencode/gencode.go --component client --config ../../config/api.job.yml --package jobclient --out ./jobclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/jobmodel
//go:generate go run ../../scripts/gencode/gencode.go --component client --config ../../config/api.scheduler.yml --package schedulerclient --out ./schedulerclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/schedulermodel
//go:generate go run ../../scripts/gencode/gencode.go --component client --config ../../config/api.kv.yml --package kvclient --out ./kvclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/kvmodel
//go:generate go run ../../scripts/gencode/gencode.go --component client --config ../../config/api.oauth.yml --package oauthclient --out ./oauthclient/client.gen.go --model-package github.com/adamlouis/squirrelbyte/server/pkg/model/oauthmodel
