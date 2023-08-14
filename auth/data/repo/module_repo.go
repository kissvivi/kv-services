package repo

import (
	"kv-services/auth/data"
)

type ModuleRepo struct {
	data.AuthRepo[data.User]
}
