package env

import (
	"fetch-spotify-cli/cnsts"
	"fetch-spotify-cli/types"
	"fetch-spotify-cli/util"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var env_file = ""

func WriteToEnvFile(clientId, clientSecret, accessToken, refreshToken string) {
	envVars := map[string]string{
		cnsts.ENV_CLIENT_ID:     clientId,
		cnsts.ENV_CLIENT_SECRET: clientSecret,
		cnsts.ENV_ACCESS_TOKEN:  accessToken,
		cnsts.ENV_REFRESH_TOKEN: refreshToken,
		cnsts.ENV_WEB_ID:        os.Getenv(cnsts.ENV_WEB_ID),
	}

	godotenv.Write(envVars, env_file)
}

func LoadSpotifyEnvVars() types.SpotifyVars {
	basePath := util.BasePath()

	err := godotenv.Load(basePath + "\\.env")
	if err != nil {
		err = godotenv.Load(basePath + "/.env")
		if err != nil {
			util.EndWithErr("cannot load .env file spotify")
		} else {
			env_file = basePath + "/.env"
		}
	} else {
		env_file = basePath + "\\.env"
	}

	return types.SpotifyVars{
		ClientID:     os.Getenv(cnsts.ENV_CLIENT_ID),
		ClientSecret: os.Getenv(cnsts.ENV_CLIENT_SECRET),
		AccessToken:  os.Getenv(cnsts.ENV_ACCESS_TOKEN),
		RefreshToken: os.Getenv(cnsts.ENV_REFRESH_TOKEN),
		State:        randomString(),
	}
}

func LoadIdEnvVar() string {
	basePath := util.BasePath()

	err := godotenv.Load(basePath + "\\.env")
	if err != nil {
		err = godotenv.Load(basePath + "/.env")
		if err != nil {
			util.EndWithErr("cannot load .env file id")
		}
	}

	return strings.TrimSpace(os.Getenv(cnsts.ENV_WEB_ID))
}
