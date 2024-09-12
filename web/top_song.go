package web

import (
	"fetch-spotify-cli/util"
	"io"
	"net/http"
)

func TopSong(id string) string {
	req, err := http.NewRequest("GET", "https://top-fetch.vercel.app/api/track?id="+id, nil)
	if err != nil {
		util.EndWithErr("cannot create request to TopFetch web server")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		util.EndWithErr("cannot send request to TopFetch web server")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		util.EndWithErr("cannot read response body from TopFetch web server")
	}

	return string(body)
}
