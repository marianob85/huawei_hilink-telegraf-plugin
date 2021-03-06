# Huawei API Input plugin

This input plugin will gather huawei api data

### Configuration:

```
[[inputs.huawei_api]]
	# List of devices
	DevicesAddress = ["http://192.168.8.1/"] # require
	
	# Collection of data to gather
	ApiUrl = ["api/device/signal", "api/device/information"] # require
```

### Tags:

- address
- api_url

### Example Output:

```
$ ./telegraf --config telegraf.conf --input-filter huawei_api --test
* Plugin: inputs.huawei_api, Collection 1
> huawei_api,address=http://192.168.2.1,api_url=api/device/signal,host=Mariano-PC ecio="",psatt="1",pci="90",cell_id="57732867",rsrp="-94dBm",sinr="2dB",mode="7",sc="",rsrq="-9dB",rssi="-63dBm",rscp="" 1496497917000000000
> huawei_api,address=http://192.168.8.1,api_url=api/device/signal,host=Mariano-PC rsrp="-89dBm",sinr="8dB",ecio="",psatt="1",lte_bandinfo="",sc="",cell_id="57732867",rsrq="-6dB",rssi="-65dBm",rscp="",mode="7",lte_bandwidth="",pci="90" 1496497917000000000
$ ./telegraf --config telegraf.conf --input-filter huawei_hilink_api --test
* Plugin: inputs.huawei_hilink_api, Collection 1
> huawei_api,address=http://192.168.8.1,host=debian,api_url=api/device/signal,DeviceName=E3372 rsrq=-7i,rsrp=-93i,sinr=7i,rssi=-69i,pci=90i 1498032147000000000
> huawei_api,api_url=api/device/signal,DeviceName=E3272,address=http://192.168.2.1,host=debian pci=90i,rsrq=-9i,sinr=9i,rsrp=-92i,rssi=-63i 1498032147000000000
> huawei_api,api_url=api/monitoring/status,DeviceName=E3372,address=http://192.168.8.1,host=debian CurrentNetworkType=19i 1498032147000000000
> huawei_api,api_url=api/monitoring/status,DeviceName=E3272,address=http://192.168.2.1,host=debian CurrentNetworkType=19i 1498032147000000000
```
