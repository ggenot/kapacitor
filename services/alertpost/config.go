package alertpost

import (
	"net/url"

	"github.com/pkg/errors"
)

type Config struct {
	Endpoint string            `toml:"endpoint" override:"endpoint"`
	URL      string            `toml:"url" override:"url"`
	Headers  map[string]string `toml:"headers" override:"headers,redact"`
}

func (c Config) Validate() error {
	if c.Endpoint == "" {
		return errors.New("must specify endpoint name")
	}

	if c.URL == "" {
		return errors.New("must specify url")
	}

	if _, err := url.Parse(c.URL); err != nil {
		return errors.Wrapf(err, "invalid URL %q", c.URL)
	}

	return nil
}
