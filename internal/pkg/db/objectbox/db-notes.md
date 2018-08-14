
DB generalization issues:

* IDs are bson.ObjectId or string
* Configuration: how about a map (string->string)? 

Misc:

* Benchmarking issue: number of elements added
* It's Event."ID" but usually it's "Id"
* count should be long
* Client DB: Values should not be passed by value
* Should we really return values by value?