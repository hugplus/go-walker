package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/utils/http_util"
)

func GetIP(c *gin.Context) string {
	ClientIP := c.ClientIP()
	//fmt.Println("ClientIP:", ClientIP)
	RemoteIP := c.RemoteIP()
	//fmt.Println("RemoteIP:", RemoteIP)
	ip := c.Request.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = c.Request.Header.Get("X-real-ip")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if RemoteIP != "127.0.0.1" {
		ip = RemoteIP
	}
	if ClientIP != "127.0.0.1" {
		ip = ClientIP
	}
	return ip
}

type IPLocation struct {
	Code int            `json:"code"` //返回码 200成功
	Msg  string         `json:"msg"`  //返回消息
	Data IPLocationData `json:"data"`
}

type IPLocationData struct {
	AreaCode       string   `json:"area_code"`       //: "320311",
	Province       string   `json:"province"`        //省: "江苏",
	City           string   `json:"city"`            //: "徐州",
	District       string   `json:"district"`        //: "丰县",
	CityCode       string   `json:"city_code"`       //: "0516",
	Continent      string   `json:"continent"`       //: "亚洲",
	Country        string   `json:"country"`         //: "中国",
	CountryCode    string   `json:"country_code"`    //: "CN",
	CountryEnglish string   `json:"country_english"` //: "",
	Elevation      string   `json:"elevation"`       //: "40",
	Ip             string   `json:"ip"`              //: "114.234.76.140",
	Isp            string   `json:"isp"`             //: "电信",
	Latitude       string   `json:"latitude"`        //: "34.227883",
	LocalTime      string   `json:"local_time"`      //: "2023-08-02 14:36",
	Longitude      string   `json:"longitude"`       //: "117.213995",
	MultiStreet    []Street `json:"multi_street"`
	Street         string   `json:"street"`          //: "解放路168号",
	Version        string   `json:"version"`         //: "V4",
	WeatherStation string   `json:"weather_station"` //: "CHXX0437",
	ZipCode        string   `json:"zip_code"`        //: "221006"
}

type Street struct {
	Lng          string `json:"lng"`           //经度: "116.60833",
	Lat          string `json:"lat"`           //纬度: "34.701533",
	Province     string `json:"province"`      //省: "江苏",
	City         string `json:"city"`          //: "徐州",
	District     string `json:"district"`      //: "丰县",
	Street       string `json:"street"`        //: "解放路168号",
	StreetNumber string `json:"street_number"` //: "解放路168号"
}

func GetLocationByIp(secretKey, ip string, location *IPLocationData) error {
	url := "https://api.ipdatacloud.com/v2/query?ip=" + ip + "&key=" + secretKey
	client := http_util.HTTPClient{}
	data, err := client.Get(url)
	if err != nil {
		return err
	}
	var ipd IPLocation
	if err := json.Unmarshal(data, &ipd); err != nil {
		return err
	}
	if ipd.Code != 200 {
		return errors.New(fmt.Sprintf("获取出错 code:{%d} msg{%s}", ipd.Code, ipd.Msg))
	}
	*location = ipd.Data
	return nil
}
