# reQuery
reQuery is an ACI moquery clone that performs remote queries and queries against the backup file. This tool is based on the [goACI](https://github.com/brightpuddle/goaci) library.


## Getting Started

## Differences from moquery

### Filtering data
moquery uses a unique filtering syntax for the `-f` filter option, e.g. `fv.BD.name=="my-tenant"`. reQuery uses the same query syntax as the API, e.g. `eq(fvBD.name,"my-tenant")`, so the queries are interchangable with the documentation, other tools, and Visore.

### Output
reQuery always outputs in JSON; however, the data structure is flattened and is similar to moquery.

### CLI arguments

#### Not implemented:

`-a --attrs` - reQuery always displays all attributes. Use external tools like awk and grep to limit results to config only.

`-o --output` - reQuery only outputs JSON.

`-p --port` - Just add the port to the hostname/IP, e.g. `10.0.0.1:443`

#### Renamed:
`-c --klass` - This was renamed to `-c --class`.


#### Unique to requery:
`-m --mode` - By default the mode is determined by the extention, i.e. .tar.gz is a backup file. If this doesn't apply, use the mode option to specify `http` or `backup`.
