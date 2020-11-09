package bivrost

import (
	"github.com/koinworks/asgard-heimdal/libs/logger"
	"github.com/koinworks/asgard-heimdal/libs/serror"
)

func (ox *bvObject) initRoute() {
	ox.Service.Get("/", ox.welcome)

	err := ox.Service.Docs("docs/api-docs.yaml")
	if err != nil {
		logger.Warn(serror.NewFromErrorc(err, "Failed to register route api docs"))
	}

	// TODO: declare bivrost route here
}
