# reQuery
reQuery is an ACI moquery clone that performs remote queries and queries against the backup file. This tool is based on the [goACI](https://github.com/brightpuddle/goaci) library.


## Getting Started

Download the latest release from [the releases tab](https://github.com/brightpuddle/requery/releases).

requery is similar to moquery with a few key differences. Since requery runs remotely it always requires a "target" which is the positional argument in the examples above. The target can be an APIC hostname or a `.tar.gz` configuration backup from the APIC.

Query parameters, e.g. `-x` and `-f` only work against the live configuration. For filtering results from the backup, use the Unix CLI tools, e.g. grep, egrep, awk, etc.


```
requery $ requery apic -u admin -d uni/tn-infra
Password: Total count: 1

# fvTenant.attributes
{
  "annotation": "",
  "childAction": "",
  "descr": "",
  "dn": "uni/tn-infra",
  "extMngdBy": "",
  "lcOwn": "local",
  "modTs": "2018-01-10T18:07:57.861+00:00",
  "monPolDn": "uni/tn-common/monepg-default",
  "name": "infra",
  "nameAlias": "",
  "ownerKey": "",
  "ownerTag": "",
  "status": "",
  "uid": "0"
}

requery $ requery ~/src/tmp/config.tar.gz -d uni/tn-infra
Total count: 1

# fvTenant.attributes
{
  "annotation": "",
  "descr": "",
  "dn": "uni/tn-infra",
  "name": "infra",
  "nameAlias": "",
  "ownerKey": "",
  "ownerTag": ""
}

requery $ requery ~/src/tmp/config.tar.gz -c fvBD
Total count: 4

# fvBD.attributes
{
  "OptimizeWanBandwidth": "no",
  "annotation": "",
  "arpFlood": "no",
  "descr": "",
  "dn": "uni/tn-infra/BD-ave-ctrl",
  "epClear": "no",
  "epMoveDetectMode": "",
  "intersiteBumTrafficAllow": "no",
  "intersiteL2Stretch": "no",
  "ipLearning": "yes",
  "limitIpLearnToSubnets": "yes",
  "llAddr": "::",
  "mac": "00:22:BD:F8:19:FF",
  "mcastAllow": "no",
  "multiDstPktAct": "bd-flood",
  "name": "ave-ctrl",
  "nameAlias": "",
  "ownerKey": "",
  "ownerTag": "",
  "type": "regular",
  "unicastRoute": "yes",
  "unkMacUcastAct": "proxy",
  "unkMcastAct": "flood",
  "vmac": "not-applicable"
}
...
```

## Differences from moquery

### Filtering data
moquery uses a unique filtering syntax for the `-f` filter option, e.g. `fv.BD.name=="my-tenant"`. reQuery uses the same query syntax as the API, e.g. `eq(fvBD.name,"my-tenant")`, so the queries are interchangable with other documentation, other tools, and Visore.

### Output
reQuery always outputs in JSON; however, the data structure is flattened and is roughly similar in appearance to moquery.

### CLI arguments

#### Not implemented:

`-a --attrs` - reQuery always displays all attributes. Use external tools like awk and grep to limit results to config only.

`-o --output` - reQuery only outputs JSON.

`-p --port` - Just add the port to the hostname/IP, e.g. `10.0.0.1:443`

#### Renamed:
`-c --klass` - This was renamed to `-c --class`.


#### Unique to requery:
`-m --mode` - By default the mode is determined by the extention, i.e. .tar.gz is a backup file. If this doesn't apply, use the mode option to specify `http` or `backup`.


## Feedback
Feedback and/or pull requests welcome.
