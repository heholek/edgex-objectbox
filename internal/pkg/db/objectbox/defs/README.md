These files are used to generate bindings. 
They are based on "github.com/edgexfoundry/go-mod-core-contracts/models" with a few changes/additions, mostly:
* indexes
* converters  

### How to generate bindings
```bash
cd edgex-objectbox
internal/pkg/db/objectbox/defs/generate.sh
```

### Event
`Event` is handled through a wrapper and can't be generated automatically. To update it:
* uncomment the `#obxmodels=internal/pkg/db/objectbox/defs/correlation #source` line in `generate.sh` and execute
* format the generated file `models/defs/event.obx.go` (after fixing the rel import error),
* compare `defs/correlation/event.obx.go` to `obx/event.obx.go` and merge manually

### Command
`Command` is using a wrapper adding DeviceId, compare with the updated command.obx.go with the original before commiting 
and update the imports, `Load()` and `Flatten()`.
 