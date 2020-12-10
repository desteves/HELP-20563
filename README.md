# HELP-20563

Examples on how to update arrays and nested arrays in Golang. Currently it uses bson.* but can be easily adopted to used structs.

The updated fields are the emojis.

- Example 1 uses the positional array operator.
- Example 2 uses the positional array operator with reference to a specific field
- Example 3 uses array filters

## Sample Run

```
export MONGODB_ATLAS_URI="mongodb+srv://foo:PASSWORD@demoatlascluster.bwkgk.mongodb.net/tutorial?retryWrites=true&w=majority"
go run main

2020/12/10 07:29:55 Example One Is [{Key:deviceID Value:example_one} {Key:preferences Value:[[{Key:K Value:setting-1} {Key:V Value:1234}] [{Key:K Value:setting-2} {Key:V Value:hello}] [{Key:K Value:setting-3} {Key:V Value:true}]]}]
2020/12/10 07:29:56 Example One Updated map[_id:ObjectID("5fd222d35cb6242bb4df3352") deviceID:example_one preferences:[map[K:setting-1 V:1234] 😍😍😍 map[K:setting-3 V:true]]]
2020/12/10 07:29:56 Example Two Is [{Key:deviceID Value:example_two} {Key:preferences Value:[[{Key:K Value:setting-1} {Key:V Value:270172701727017}] [{Key:K Value:setting-2} {Key:V Value:[{Key:setting-2a Value:GMT-7} {Key:setting-2b Value:1234}]}]]}]
2020/12/10 07:29:57 Example Two Updated map[_id:ObjectID("5fd222d45cb6242bb4df3353") deviceID:example_two preferences:[map[K:setting-1 V:270172701727017] map[K:setting-2 V:map[setting-2a:GMT-7 setting-2b:1234] setting-2b:🤩🤩🤩]]]
2020/12/10 07:29:57 Example three Is [{Key:deviceID Value:example_three} {Key:preferences Value:[[{Key:K Value:setting-1} {Key:V Value:270172701727017}] [{Key:K Value:setting-2} {Key:V Value:[{Key:setting-2a Value:helloworld} {Key:setting-2b Value:270172701727017}]}] [{Key:K Value:setting-3} {Key:V Value:[{Key:setting-3a Value:foobarfoobar} {Key:setting-3b Value:[[{Key:K Value:k1} {Key:V Value:-2701727017}] [{Key:K Value:k2} {Key:V Value:true}] [{Key:K Value:k3} {Key:V Value:false}]]}]}]]}]
2020/12/10 07:29:57 options_three  &{Registry:<nil> Filters:[map[myFirstFilter.K:setting-3] map[mySecondFilter.K:k1]]}
2020/12/10 07:29:57 Example three Updated map[_id:ObjectID("5fd222d55cb6242bb4df3354") deviceID:example_three preferences:[map[K:setting-1 V:270172701727017] map[K:setting-2 V:map[setting-2a:helloworld setting-2b:270172701727017]] map[K:setting-3 V:map[setting-3a:foobarfoobar setting-3b:[map[K:🔥🔥🔥 V:-2701727017] map[K:k2 V:true] map[K:k3 V:false]]]]]]
```