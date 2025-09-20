package auth

import "github.com/ExtraProjects860/Project-Device-Mobile/config"

func jwtKey() []byte {
	return []byte(config.GetEnv().API.JwtKey)
}

func refreshKey() []byte {
	return []byte(config.GetEnv().API.RefreshKey)
}
