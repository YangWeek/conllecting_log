package sysinfo

//import (
//	"fmt"
//	client "github.com/influxdata/influxdb1-client/v2"
//	"github.com/shirou/gopsutil/cpu"
//	"github.com/shirou/gopsutil/disk"
//	"github.com/shirou/gopsutil/load"
//	"github.com/shirou/gopsutil/mem"
//
//	"log"
//	"time"
//)
//
//var (
//	cli client.Client
//)
//
//func connInflux() (err error) {
//	cli, err = client.NewHTTPClient(client.HTTPConfig{
//		Addr:     "http://127.0.0.1:8086",
//		Username: "admin",
//		Password: "",
//	})
//	return
//}
//
//// insert
//func writesPoints(percent float64) {
//	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
//		Database:  "monitor",
//		Precision: "s", //精度，默认ns
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	tags := map[string]string{"cpu": "cpu"}
//	fields := map[string]interface{}{
//		"cpu_percent": percent,
//	}
//
//	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
//	if err != nil {
//		log.Fatal(err)
//	}
//	bp.AddPoint(pt)
//	err = cli.Write(bp)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("insert success")
//}
//
//func getCpuInfo() {
//	percent, _ := cpu.Percent(time.Second, false)
//	fmt.Printf("cpu percent is %v\n", percent)
//	writesPoints(percent[0])
//}
//
//func getCpuLoad() {
//	info, _ := load.Avg()
//	fmt.Printf("%v\n", info)
//}
//
//// mem info
//func getMemInfo() {
//	memInfo, _ := mem.VirtualMemory()
//	fmt.Printf("mem info:%v\n", memInfo)
//}
//
//// disk info
//func getDiskInfo() {
//	parts, err := disk.Partitions(true)
//	if err != nil {
//		fmt.Printf("get Partitions failed, err:%v\n", err)
//		return
//	}
//	for _, part := range parts {
//		fmt.Printf("part:%v\n", part.String())
//		diskInfo, _ := disk.Usage(part.Mountpoint)
//		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
//	}
//
//	ioStat, _ := disk.IOCounters()
//	for k, v := range ioStat {
//		fmt.Printf("%v:%v\n", k, v)
//	}
//}
//
//func main() {
//	err := connInflux()
//	if err != nil {
//		fmt.Printf("conn influxdb failed, err is %v\n", err)
//		return
//	}
//	ticker := time.Tick(time.Second)
//	for {
//		select {
//		case <-ticker:
//			getCpuInfo()
//		}
//	}
//}
