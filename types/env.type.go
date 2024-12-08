package types

type GoEnv string

const (
	GoEnvDev        GoEnv = "development"
	GoEnvStaging    GoEnv = "staging"
	GoEnvProduction GoEnv = "production"
)

func (env GoEnv) String() string {
	return string(env)
}

func ParseGoEnv(env string) GoEnv {
	switch env {
	case "production":
		return GoEnvProduction
	case "staging":
		return GoEnvStaging
	default:
		return GoEnvDev
	}
}
