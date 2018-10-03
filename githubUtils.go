// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2018, The 2ACoin Developers
//
// Please see the included LICENSE file for more information.
//

package main

import (
	"github.com/mcuadros/go-version"
)

type release struct {
	TagName string `json:"tag_name"`
	URL     string `json:"html_url"`
}

const url = "https://api.github.com/repos/2acoin/2acoin-wallet-go/releases/latest"

func checkIfNewReleaseAvailableOnGithub(currentVersion string) (newVersion string, urlNewVersion string) {

	latestRelease := new(release)
	getJSONFromHTTPRequest(url, latestRelease)

	if version.CompareSimple(latestRelease.TagName, currentVersion) == 1 {
		return latestRelease.TagName, latestRelease.URL
	}

	return "", ""
}
