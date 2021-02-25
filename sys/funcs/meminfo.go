package funcs

import (
	"github.com/didi/nightingale/src/common/dataobj"
	"github.com/shirou/gopsutil/mem"
	"github.com/toolkits/pkg/logger"
)

func memInfo() (*mem.VirtualMemoryStat, error) {
	meminfo, err := mem.VirtualMemory()
	return meminfo, err
}

func MemMetrics() []*dataobj.MetricValue {
	meminfo, err := memInfo()
	if err != nil {
		logger.Error("get mem info failed: ", err)
		return nil
	}

	return []*dataobj.MetricValue{
		GaugeValue("memory.bytes.total", meminfo.Total),
		GaugeValue("memory.bytes.used", meminfo.Used),
		GaugeValue("memory.bytes.free", meminfo.Available),
		GaugeValue("mmemoryem.bytes.used.percent", meminfo.UsedPercent),
	}

}
