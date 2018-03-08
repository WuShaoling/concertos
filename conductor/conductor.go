package conductor

import (
	"github.com/concertos/conductor/module/restapi"
	"github.com/concertos/conductor/module/scheduler"
)

type Conductor struct {
	restApi  *restapi.RestApi
	schedule *scheduler.Scheduler
}
