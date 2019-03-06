These files are used to generate bindings. 
They are based on "github.com/edgexfoundry/go-mod-core-contracts/models" with a few changes/additions, mostly:
* indexes
* converters  

### How to generate bindings
```bash
cd edgex-go
internal/pkg/db/objectbox/models/generate.sh
```

### TODOs
* check/fix the licenses of the files
* indexes
* relations