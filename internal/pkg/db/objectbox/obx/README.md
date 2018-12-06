### How to generate bindings
Quite a few manual changes are still required

* before generation, change bson.ObjectId to string in the code manually (so that the generator picks it up)

```bash
obxmodels=internal/pkg/db/objectbox/obxmodels
alias objectbox-gogen="go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen/ -persist ${obxmodels}/objectbox-model.json"
objectbox-gogen -source pkg/models/reading.go
objectbox-gogen -source pkg/models/event.go

mv pkg/models/*obx.go ${obxmodels}/

for f in ${obxmodels}/*.go; do sed -i 's/package models/package obx/g' "$f"; done

```

* manually add casting between `string` & `bson.ObjectId` (just those that are necessary to build)
* add `. "github.com/edgexfoundry/edgex-go/pkg/models"` import to each binding
*  