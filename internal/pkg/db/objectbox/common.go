package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/pkg/models"
)

func onCreate(base *models.BaseObject) {
	if base.Created == 0 {
		base.Created = db.MakeTimestamp()
	}
}

func onUpdate(base *models.BaseObject) {
	base.Modified = db.MakeTimestamp()
}
