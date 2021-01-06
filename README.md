# Huawei API Input plugin

This input plugin will gather huawei hilink modem data

### Install Instructions 

To integrate with telegraf, extend the telegraf.conf using the following example
```
[[inputs.execd]]
   command = ["/path/to/huawei_hilink-telegraf-plugin", "-config", "/path/to/huawei_hilink-telegraf-plugin.config"]
   signal = "STDIN"
```

### Configuration:

```
[[inputs.huawei_hilink]]
	# List of devices
	DevicesAddress = ["http://192.168.8.1"] # require
	[[inputs.huawei_hilink.module]]
		Url = "api/device/signal"
		Fields = ["pci", "rsrq", "rsrp", "rssi", "sinr" ]
	[[inputs.huawei_hilink.module]]
		Url = "api/monitoring/status"
		Fields = [ "CurrentNetworkType", "CurrentNetworkTypeEx", "ConnectionStatus", "maxsignal", "SignalIcon"]
	[[inputs.huawei_hilink.module]]
		Url = "api/device/information"
		Fields = ["DeviceName"]
		Is_tag = true
	[[inputs.huawei_hilink.module]]	
		Url = "api/monitoring/traffic-statistics"
		Fields = ["TotalUpload", "TotalDownload" ]
	[[inputs.huawei_hilink.module]]
		Url = "api/monitoring/month_statistics"
		Fields = ["CurrentMonthDownload", "CurrentMonthUpload"]
```

### Tags:

- address
- url

### Example Output:

```
$ ./telegraf --config telegraf.conf --input-filter huawei_hilink --test
* Plugin: inputs.huawei_hilink, Collection 1
> huawei_hilink,address=http://192.168.2.1,url=api/device/signal,host=Mariano-PC ecio="",psatt="1",pci="90",cell_id="57732867",rsrp="-94dBm",sinr="2dB",mode="7",sc="",rsrq="-9dB",rssi="-63dBm",rscp="" 1496497917000000000
> huawei_hilink,address=http://192.168.8.1,url=api/device/signal,host=Mariano-PC rsrp="-89dBm",sinr="8dB",ecio="",psatt="1",lte_bandinfo="",sc="",cell_id="57732867",rsrq="-6dB",rssi="-65dBm",rscp="",mode="7",lte_bandwidth="",pci="90" 1496497917000000000
$ ./telegraf --config telegraf.conf --input-filter huawei_hilink_api --test
* Plugin: inputs.huawei_hilink_api, Collection 1
> huawei_hilink,address=http://192.168.8.1,host=debian,url=api/device/signal,DeviceName=E3372 rsrq=-7i,rsrp=-93i,sinr=7i,rssi=-69i,pci=90i 1498032147000000000
> huawei_hilink,url=api/device/signal,DeviceName=E3272,address=http://192.168.2.1,host=debian pci=90i,rsrq=-9i,sinr=9i,rsrp=-92i,rssi=-63i 1498032147000000000
> huawei_hilink,url=api/monitoring/status,DeviceName=E3372,address=http://192.168.8.1,host=debian CurrentNetworkType=19i 1498032147000000000
> huawei_hilink,url=api/monitoring/status,DeviceName=E3272,address=http://192.168.2.1,host=debian CurrentNetworkType=19i 1498032147000000000
```
