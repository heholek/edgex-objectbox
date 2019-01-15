These files are used to generate bindings. 
They are based on "github.com/edgexfoundry/edgex-go/pkg/models" with a few changes/additions, mostly:
* indexes
* converters  

### How to generate bindings
```bash
cd edgex-go
obxbindings=internal/pkg/db/objectbox/obx
obxmodels=internal/pkg/db/objectbox/models
alias objectbox-gogen="go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen/ -byValue -persist ${obxbindings}/objectbox-model.json"
objectbox-gogen -source ${obxmodels}/reading.go
objectbox-gogen -source ${obxmodels}/event.go
objectbox-gogen -source ${obxmodels}/value-descriptor.go

for f in ${obxmodels}/*obx.go; do sed -i 's/import (/import (\n\t. "github.com\/edgexfoundry\/edgex-go\/pkg\/models"/g' "$f"; done
mv ${obxmodels}/*obx.go ${obxbindings}/
for f in ${obxbindings}/*.go; do sed -i 's/package models/package obx/g' "$f"; done
```

### TODOs
* check/fix the licenses of the files
* indexes
* relations