package repo

import "kv-services/auth/data"

type UserRepo struct {
	data.AuthRepo[data.User]
}
