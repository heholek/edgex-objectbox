These files are used to generate bindings. 
They are based on "github.com/edgexfoundry/go-mod-core-contracts/models" with a few changes/additions, mostly:
* indexes
* converters  

### How to generate bindings
```bash
cd edgex-objectbox
internal/pkg/db/objectbox/models/generate.sh
```

Additionally, `Event` is handled through a wrapper and can't be generated automatically.
To update it:
* uncomment the `#obxmodels=internal/pkg/db/objectbox/models/correlation #source` line in `generate.sh` and execute
* format the generated file `models/correlation/event.obx.go` (after fixing the rel import error),
* compare `models/correlation/event.obx.go` to `obx/event.obx.go` and merge manually

### TODOs
* check/fix the licenses of the files
* indexes
* relations