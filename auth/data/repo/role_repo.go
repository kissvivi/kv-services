package repo

import "kv-services/auth/data"

type RoleRepo struct {
	data.AuthRepo[data.Role]
}
