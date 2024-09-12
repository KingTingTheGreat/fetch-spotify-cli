package main

import (
	"fetch-spotify-cli/cnsts"
	"fetch-spotify-cli/env"
	"fetch-spotify-cli/spotify"
	"fetch-spotify-cli/util"
	"fetch-spotify-cli/web"
	"os"
	"strings"
)

func topTrackFromWeb(args []string) (string, string) {
	flag := args[0]

	if flag != cnsts.WEB_FLAG && flag != cnsts.WEB_FLAG_SHORT {
		util.EndWithErr("invalid flag")
	}

	var id string
	if len(args) > 1 {
		id = args[1]
	} else {
		id = env.LoadIdEnvVar()
	}

	if id == "" {
		util.EndWithErr("missing web id")
	}

	songData := web.TopSong(id)
	songParts := strings.Split(songData, cnsts.DELIM)

	if len(songParts) != 2 {
		util.EndWithErr(songData)
	}

	trackText := songParts[0]
	imageUrl := songParts[1]

	return trackText, imageUrl
}

func topTrackFromLocal() (string, string) {
	env.LoadSpotifyEnvVars()

	if spotify.SpotifyVars.ClientID == "" || spotify.SpotifyVars.ClientSecret == "" {
		util.EndWithErr("Missing client ID or client secret")
	}

	if spotify.SpotifyVars.AccessToken == "" || spotify.SpotifyVars.RefreshToken == "" {
		spotify.InitializeSpotifyAccess()
	}

	topTrack := spotify.TopSong()

	trackText := topTrack.Name + " - " + topTrack.Artists[0].Name

	return trackText, topTrack.Album.Images[0].Url
}

func main() {
	var trackText, imageUrl string
	args := os.Args[1:]
	if len(args) > 0 {
		trackText, imageUrl = topTrackFromWeb(args)
	} else {
		trackText, imageUrl = topTrackFromLocal()
	}

	ansiImage := util.AnsiImage(imageUrl)

	os.Stdout.WriteString(util.WriteOutputToFile(ansiImage, trackText))
}
