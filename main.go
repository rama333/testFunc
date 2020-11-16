package main

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/gomodule/redigo/redis"
	"github.com/tiaguinho/gosoap"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var(
	a=0
)

type Sent struct {
	Date string `json:"send_date,omitempty"`
	Id int `json:"id,omitempty"`
	Sms_id int `json:"sms_id,omitempty"`
	Sms_text string `json:"sms_text,omitempty"`
	Source_addr string `json:"source_addr,omitempty"`
	Dest_addr string `json:"dest_addr,omitempty"`
	Delivery_time string `json:"delivery_time,omitempty"`
	//Dest_addr_now int64 `json:"dest_addr_now,omitempty"`
	Sequence int64 `json:"sequence,omitempty"`
}

type Test struct {
	Num string `json:"num"`
}


func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.1.3:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func ping(c redis.Conn) error {
	// Send PING command to Redis
	pong, err := c.Do("PING")
	if err != nil {
		return err
	}

	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

func getCode() (string) {
	charSet := "abcdedfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	rand.Seed(time.Now().Unix())

	var code strings.Builder

	for i := 0; i < 4; i++ {
		code.WriteString(string(charSet[rand.Intn(len(charSet)-1)]))
	}

	return code.String()
}

type GetIPLocationResponse struct {

    Result int `xml:"Result"`
	Description string `xml:"Description"`
	SubscriberID string `xml:"SubscriberID"`
	Services string `xml:"Services"`
	Quotas string `xml:"Quotas"`
	IpAddress string `xml:"IpAddress"`
	MacAddress string `xml:"MacAddress"`
}


var (
	r GetIPLocationResponse
)

func main()  {


	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient("http://192.168.100.113:7301/WS4PortalEJBBean/WS4PortalWS?wsdl", httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}


	log.Println(soap.HeaderName)

	// Use gosoap.ArrayParams to support fixed position params
	params := gosoap.Params{
		"namespace" : "ws4p.irbis",
		"ipPort": "4555",
		"ipAddress": "8.8.8.8",
	}

	res, err := soap.Call("WS4PortalWS", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}


	log.Println(string(res.Body))

	//res.Unmarshal(&r)
	//
	//// GetIpLocationResult will be a string. We need to parse it to XML
	//result := GetIPLocationResult{}
	//err = xml.Unmarshal([]byte(r.GetIPLocationResult), &result)
	//if err != nil {
	//	log.Fatalf("xml.Unmarshal error: %s", err)
	//}
	//
	//if result.Country != "US" {
	//	log.Fatalf("error: %+v", r)
	//}
	//
	//log.Println("Country: ", result.Country)
	//log.Println("State: ", result.State)



	// if everything went well res.ID should have its
	// value set with the one returned by the service.




	//
	//pool := newPool()
	//conn := pool.Get()
	//
	//err := ping(conn)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//defer conn.Close()
	//
	//t1 := time.Now()
	//for i := 0; i < 10000; i++ {
	//	//	s:= "test" + strconv.Itoa(i)
	//	if _, err := conn.Do("SET", i, "123"); err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//
	//fmt.Println("ok")
	//
	//
	//t := time.Now()
	//values, err := redis.Int64(conn.Do("GET", 100))
	//if err != nil {
	//	fmt.Println("value rr")
	//	log.Fatal(err)
	//
	//}
	//fmt.Println(t1)
	//fmt.Println(t)
	//
	//fmt.Println(values)
	//



	//test:= Test{Num: "test"}


	//if _, err := conn.Do("SET", "graph", "grap"); err != nil {
	//	log.Fatal(err)
	//}







	//timeq := time.Now().Unix()
	//
	//time.Sleep(time.Second*3)
	//
	//fmt.Println(timeq)
	//
	//timeq = time.Now().Unix()
	//
	//fmt.Println(timeq)
	//
	//
	//
	//defer func() {
	//	if r:=recover();r != nil{
	//		fmt.Println(r)
	//
	//		fmt.Println(a+a)
	//	}
	//}()
	//
	//
	//for i := 0; i < 10; i++ {
	//
	//	a=i
	//	if (i==5){
	//
	//		panic("panic")
	//	}
	//
	//}









	//fmt.Println(time.Now().Second())
	//
	//timeNow := time.Now()
	//
	//fmt.Println(timeNow.Second())
	//
	//time.Sleep(time.Second*70)
	//
	//timeNowTemp := time.Now()
	//
	//
	//fmt.Println(timeNowTemp.Second())
	//fmt.Println(timeNow.Second())
	//
	//
	//fmt.Println((timeNowTemp.Sub(timeNow)).Seconds())

	//valueArgs := make([]string, 0)
	//
	//valueArgs = append(valueArgs, "post.Column1")
	//valueArgs = append(valueArgs, "post.Column2")
	//valueArgs = append(valueArgs, "post.Column3")
	//
	//stmt := fmt.Sprintf("INSERT INTO my_sample_table (column1, column2, column3) VALUES %s",
	//	strings.Join(valueArgs, ","))
	//
	//fmt.Println(stmt)

	//jsonByteArray := []byte(`{"send_date": "2020-11-05 15:27:11", "id": 296250426, "sms_id": 19, "sms_text": "Важность: ПредупреждениеСервис: \n Data Base мониторингМетрика: Analytic Сообщение: Used FRA space 81 percent(s)Статус: ОткрытиеДата: 02.11.2020 04:36:29", "source_addr": "79393929146", "dest_addr": "79047174347", "delivery_time": "201105152708012+", "sequence": 2518}`)
	//
	//
	//var send Sent
	//err := json.Unmarshal(jsonByteArray, &send)
	//if (err != nil){
	//	log.Println(err)
	//}
	//log.Println(send)


	//fmt.Println(parseDate("201020112435"))
	//fmt.Println(parseDate("201020160945012+"))
	//
	//des_adr, _ := strconv.ParseInt("79393048431", 10, 64)
	//
	//fmt.Println(des_adr)


}

func parseDate(date string)  (time.Time) {

	YY,MM,DD,hh,mm,ss := date[0:2], date[2:4], date[4:6],date[6:8], date[8:10], date[10:12]
	tempDate := strconv.Itoa(time.Now().Year())[0:2] + YY + "-" + MM + "-" + DD + " " + hh + ":" + mm + ":" + ss

	fmt.Println(tempDate)

	layout := "2006-01-02 15:04:05"

	d, err := time.Parse(layout,tempDate)
	if err != nil{
		fmt.Println(err)
	}

	t := d.Format(layout)

	fmt.Println("ok", t)



	return d
}
