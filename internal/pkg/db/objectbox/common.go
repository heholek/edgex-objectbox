package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/objectbox/objectbox-go/objectbox"
)

func onCreate(base *models.BaseObject) {
	if base.Created == 0 {
		base.Created = db.MakeTimestamp()
	}
}

func onUpdate(base *models.BaseObject) {
	base.Modified = db.MakeTimestamp()
}

// NOTE because there's no StringVector.ContainsAny() query, we build it using OR
func stringVectorContainsAny(sv *objectbox.PropertyStringVector, items []string, caseSensitive bool) objectbox.Condition {
	var conditions = make([]objectbox.Condition, len(items))
	for k, str := range items {
		conditions[k] = sv.Contains(str, caseSensitive)
	}

	return objectbox.MatchAny(conditions...)
}
