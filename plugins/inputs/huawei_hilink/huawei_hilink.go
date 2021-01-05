package huawei_hilink

import (
	//"fmt"
	"errors"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/knq/hilink"
	"regexp"
	"strconv"
	"sync"
)

type HuaweiHilink struct {
	DevicesAddress []string

	Module []Module
}

func (s *HuaweiHilink) Description() string {
	return "Return huawei hilink data"
}

const sampleConfig = `
	# List of devices
	DevicesAddress = ["http://192.168.2.1", "http://192.168.8.1"] # require
	
	[[inputs.huawei_hilink.module]]
		Url = "api/device/signal"
		Fields = ["pci", "rsrq", "rsrp", "rssi", "sinr" ]
	[[inputs.huawei_hilink.module]]
		Url = "api/monitoring/status"
		Fields = [ "CurrentNetworkType" ]
	[[inputs.huawei_hilink.module]]
		Url = "api/device/information"
		Fields = ["DeviceName"]
		Is_tag = true
`

func (s *HuaweiHilink) SampleConfig() string {
	return sampleConfig
}

type Module struct {
	// Name will be the name of the measurement.
	Url string

	Fields []string

	Is_tag bool
}

func (m *Module) ContainField(key string) bool {
	for _, value := range m.Fields {
		if value == key {
			return true
		}
	}

	return false
}

func (m *Module) Parse(data interface{}) (interface{}, error) {
	dataRegex := regexp.MustCompile(`(-?\d+)\D*`)
	stats := dataRegex.FindStringSubmatch(data.(string))

	if len(stats) == 2 {
		value, err := strconv.Atoi(stats[1])
		if err != nil {
			return nil, err
		} else {
			return value, nil
		}
	}

	return nil, errors.New("Can't parse: " + data.(string))
}

func (p *HuaweiHilink) GetFields(data map[string]interface{}, module Module) (map[string]interface{}, map[string]string, error) {
	fields := make(map[string]interface{})
	tags := make(map[string]string)
	for k, v := range data {
		if module.ContainField(k) {
			if module.Is_tag {
				tags[k] = v.(string)
			} else {
				value, err := module.Parse(v.(string))
				if err == nil {
					fields[k] = value
				}
			}
		}
	}

	return fields, tags, nil
}

func (p *HuaweiHilink) ConcatTags(tag map[string]string, tags []map[string]string) map[string]string {

	ret := make(map[string]string)
	for k, v := range tag {
		ret[k] = v
	}

	for _, m := range tags {
		for k, v := range m {
			ret[k] = v
		}
	}

	return ret
}

type Table struct {
	Fields map[string]interface{}
	Tags   map[string]string
}

func (p *HuaweiHilink) Gather(acc telegraf.Accumulator) error {
	var wg_address sync.WaitGroup
	var wg_url sync.WaitGroup
	for _, address := range p.DevicesAddress {
		wg_address.Add(1)
		go func(address string) {
			defer wg_address.Done()
			opts := []hilink.Option{
				hilink.URL(address),
			}
			// create client
			client, err := hilink.NewClient(opts...)
			if err != nil {
				acc.AddError(err)
			} else {
				tables := make([]Table, 0)
				tags := make([]map[string]string, 0)

				for _, module := range p.Module {
					wg_url.Add(1)
					go func(module Module, address string, tableCollections *[]Table, tagCollection *[]map[string]string) {
						defer wg_url.Done()
						data, err := client.Do(module.Url, nil)
						if err != nil {
							acc.AddError(err)
						} else {
							field, tag, err := p.GetFields(data, module)
							if err != nil {
								acc.AddError(err)
							} else {
								*tagCollection = append(*tagCollection, tag)
								table := Table{Fields: make(map[string]interface{}), Tags: make(map[string]string)}
								table.Tags["url"] = module.Url
								table.Fields = field
								*tableCollections = append(*tableCollections, table)
							}
						}
					}(module, address, &tables, &tags)
				}
				wg_url.Wait()

				for _, table := range tables {
					tag := p.ConcatTags(table.Tags, tags)
					tag["address"] = address
					acc.AddFields("huawei", table.Fields, tag)
				}
			}
		}(address)
	}

	wg_address.Wait()

	return nil
}

func init() {
	inputs.Add("huawei_hilink", func() telegraf.Input {
		return &HuaweiHilink{}
	})
}
