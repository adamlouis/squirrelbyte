package model

//go:generate go run ../../scripts/gencode/gencode.go --component model --config ../../config/api.document.yml  --package documentmodel --out ./documentmodel/model.gen.go
//go:generate go run ../../scripts/gencode/gencode.go --component model --config ../../config/api.job.yml       --package jobmodel --out ./jobmodel/model.gen.go
//go:generate go run ../../scripts/gencode/gencode.go --component model --config ../../config/api.scheduler.yml --package schedulermodel --out ./schedulermodel/model.gen.go
//go:generate go run ../../scripts/gencode/gencode.go --component model --config ../../config/api.secret.yml    --package secretmodel --out ./secretmodel/model.gen.go
//go:generate go run ../../scripts/gencode/gencode.go --component model --config ../../config/api.oauth.yml     --package oauthmodel --out ./oauthmodel/model.gen.go
