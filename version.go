package updater

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
	latest "github.com/tcnksm/go-latest"
)

type VersionInput struct {
	OptoutEnvVar   *string
	URL            string
	Os             string
	App            string
	CurrentVersion string
}

type VersionOutput struct {
	Outdated       bool
	CurrentVersion string
	Message        string
	LatestVersion  string
}

func NewInput(app, url, currentVersion string, optoutEnvVar *string) *VersionInput {
	i := &VersionInput{
		App:            app,
		Os:             runtime.GOOS,
		CurrentVersion: currentVersion,
		URL:            url,
		OptoutEnvVar:   optoutEnvVar,
	}

	return i
}

func (i *VersionInput) isOptOut() bool {
	return i.OptoutEnvVar != nil && getEnv(*i.OptoutEnvVar, "") == "false"
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func CheckVersion(input *VersionInput) (*VersionOutput, error) {

	if input == nil {
		return nil, errors.New("ErrNoInput")
	}

	if input.isOptOut() {
		return nil, errors.New("User-OptOut-Version-Check")
	}

	json := &latest.JSON{
		URL: fmt.Sprintf("%s/version?app=%s&current_version=%s&os=%s", input.URL, input.App, input.CurrentVersion, input.Os),
	}

	res, err := latest.Check(json, input.CurrentVersion)

	if err != nil {
		log.WithError(err).Debug("failed fetching latest version check")
		return nil, err
	}

	return &VersionOutput{Outdated: res.Outdated, CurrentVersion: input.CurrentVersion, Message: res.Meta.Message, LatestVersion: res.Current}, nil

}
