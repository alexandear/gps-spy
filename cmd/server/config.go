package server

const (
	defaultBuntdbPath       = ":memory:"
	defaultSpecnomeryServer = "specnomery.local"
)

var config Config

type Config struct {
	buntdbPath       string
	specnomeryServer string
}

func init() {
	config.buntdbPath = defaultBuntdbPath
	config.specnomeryServer = defaultSpecnomeryServer
}
