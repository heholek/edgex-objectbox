package objectbox

import (
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/objectbox-go/objectbox"
	"strings"
)

func onCreate(base *models.Timestamps) {
	if base.Created == 0 {
		base.Created = db.MakeTimestamp()
	}
}

func onUpdate(base *models.Timestamps) {
	base.Modified = db.MakeTimestamp()
}

// NOTE because there's no StringVector.ContainsAny() query, we build it using OR
func stringVectorContainsAny(sv *objectbox.PropertyStringVector, items []string, caseSensitive bool) objectbox.Condition {
	var conditions = make([]objectbox.Condition, len(items))
	for k, str := range items {
		conditions[k] = sv.Contains(str, caseSensitive)
	}

	return objectbox.Any(conditions...)
}

func mapError(err error) error {
	if err == nil {
		return nil
	}

	if strings.HasPrefix(err.Error(), "Unique constraint") &&
		strings.Contains(err.Error(), "would be violated by putting entity") {
		return db.ErrNotUnique
	}

	return err
}
