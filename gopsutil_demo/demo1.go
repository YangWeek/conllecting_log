package sysinfo

//import (
//	"fmt"
//	"log"
//	"time"
//
//	client "github.com/influxdata/influxdb1-client/v2"
//)
//
//// influxdb demo
//
//func connInflux() client.Client {
//	cli, err := client.NewHTTPClient(client.HTTPConfig{
//		Addr:     "http://127.0.0.1:8086",
//		Username: "admin",
//		Password: "",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	return cli
//}
//
//// query
//func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
//	q := client.Query{
//		Command:  cmd,
//		Database: "tets",
//	}
//	if response, err := cli.Query(q); err == nil {
//		if response.Error() != nil {
//			return res, response.Error()
//		}
//		res = response.Results
//	} else {
//		return res, err
//	}
//	return res, nil
//}
//
//// insert
//func writesPoints(cli client.Client) {
//	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
//		Database:  "tets",
//		Precision: "s", //精度，默认ns
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	tags := map[string]string{"cpu": "ih-cpu"}
//	fields := map[string]interface{}{
//		"idle":   201.1,
//		"system": 43.3,
//		"user":   86.6,
//	}
//
//	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
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
//func main() {
//	conn := connInflux()
//	fmt.Println(conn)
//
//	// insert
//	writesPoints(conn)
//
//	// 获取10条数据并展示
//	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
//	res, err := queryDB(conn, qs)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, row := range res[0].Series[0].Values {
//		for j, value := range row {
//			log.Printf("j:%d value:%v\n", j, value)
//		}
//	}
//}
