package kernel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWeather(city string) map[string]interface{} {
	var data string
	// 取 redis
	data = getWeatherByRedis(city)

	if data != "" {
		fmt.Println("取得 Redis 資料")
		return format(data)
	}

	// 無 redis 資料，則取 DB
	data = getWeatherByDB(city)
	if data != "" {
		fmt.Println("取得 DB 資料")

		setWeatherToRedis(city, data)

		return format(data)
	}

	// 無 DB 資料，則取 API
	data = getWeatherByAPI(city)
	fmt.Println("取得 API 資料")

	setWeatherToRedis(city, data)
	setWeatherToDB(city, data)

	return format(data)
}

// 調整輸出資料格式
func format(content string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(content), &result)

	return result
}

func getWeatherByAPI(city string) string {
	url := "http://weather.json.tw/api?region=" + city

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	return string(body)
}

func getWeatherByDB(city string) string {
	// 建立 DB 連線
	db := DB{}
	db.CreateConn()
	defer db.Conn.Close()

	// 進行查詢
	result := Weather{}
	db.Conn.Where("city = ?", city).First(&result)

	return result.City
}

func getWeatherByRedis(city string) string {
	// 建立 Redis 連線
	r := Redis{}
	r.CreateConn()
	defer r.Conn.Close()

	// 取 Redis 資料
	return r.Get(city)
}

func setWeatherToDB(city string, content string) {
	// 建立 DB 連線
	db := DB{}
	db.CreateConn()
	defer db.Conn.Close()

	new := Weather{
		City:    city,
		Content: content,
	}

	// 寫入 DB
	db.Conn.Create(&new)
}

func setWeatherToRedis(city string, content string) {
	// 建立 Redis 連線
	r := Redis{}
	r.CreateConn()
	defer r.Conn.Close()

	// 寫入 Redis
	r.Set(city, content)
}
