package cnsts

const (
	REDIRECT_URI      = "http://localhost:8080/callback"
	AUTH_URL          = "https://accounts.spotify.com/authorize"
	TOKEN_URL         = "https://accounts.spotify.com/api/token"
	SCOPE             = "user-top-read"
	TOP_TRACKS_URL    = "https://api.spotify.com/v1/me/top/tracks?time_range=short_term"
	TOP_ARTISTS_URL   = "https://api.spotify.com/v1/me/top/artists?time_range=short_term"
	ANSI_RESET        = "\x1b[0m"
	ANSI_FOREGROUND   = "\x1b[38;2;"
	CHAR              = "█"
	FONT_ASPECT_RATIO = 0.46
	DIM               = 44
	OUTPUT_FILE_NAME  = "output.txt"
	ENV_CLIENT_ID     = "SPOTIFY_CLIENT_ID"
	ENV_CLIENT_SECRET = "SPOTIFY_CLIENT_SECRET"
	ENV_ACCESS_TOKEN  = "SPOTIFY_ACCESS_TOKEN"
	ENV_REFRESH_TOKEN = "SPOTIFY_REFRESH_TOKEN"
	ENV_WEB_ID        = "WEB_ID"
	WEB_FLAG          = "--web"
	WEB_FLAG_SHORT    = "-w"
	DELIM             = "\x1d"
)
