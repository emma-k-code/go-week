package main

import (
	"encoding/json"
	"fmt"
)

type UserData struct {
	Ret []struct {
		ID               int    `json:"id"`
		ParentID         int    `json:"parent_id"`
		Parent           int    `json:"parent"`
		AllParents       []int  `json:"all_parents"`
		Domain           int    `json:"domain"`
		LastLogin        string `json:"last_login"`
		Currency         string `json:"currency"`
		PasswordExpireAt string `json:"password_expire_at"`
		PasswordReset    bool   `json:"password_reset"`
		LastBank         int    `json:"last_bank"`
		Username         string `json:"username"`
		Block            bool   `json:"block"`
		Bankrupt         bool   `json:"bankrupt"`
		Cash             struct {
			ID          int     `json:"id"`
			UserID      int     `json:"user_id"`
			Balance     float64 `json:"balance"`
			PreSub      int     `json:"pre_sub"`
			PreAdd      int     `json:"pre_add"`
			Currency    string  `json:"currency"`
			LastEntryAt string  `json:"last_entry_at"`
		} `json:"cash"`
		CashFake    interface{} `json:"cash_fake"`
		Card        interface{} `json:"card"`
		EnabledCard interface{} `json:"enabled_card"`
	} `json:"ret"`
	Result  string `json:"result"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}

type Q2Data struct {
	IP   string
	Host string
	Port int
}

func main() {
	Q1()
	Q2()
}

func Q1() {
	data := UserData{}
	err := json.Unmarshal(getApiResult(), &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Q1")
	fmt.Printf("id: %d\n", data.Ret[0].ID)
	fmt.Printf("domain: %d\n", data.Ret[0].Domain)
	fmt.Printf("username: %s\n", data.Ret[0].Username)
	fmt.Printf("cash: %+v\n", data.Ret[0].Cash)
	fmt.Println()
}

func Q2() {
	/* Q2-1
	IP:192.17.55.3	Host:docs.google.com 		Port:80
	IP:192.17.33.17 Host:ts-phpadmin.vir999.com Port:80
	IP:192.17.99.52 Host:jsonviewer.stack.hu 	Port:7711
	*/
	data := []Q2Data{
		Q2Data{IP: "192.17.55.3", Host: "docs.google.com", Port: 80},
		Q2Data{IP: "192.17.33.17", Host: "ts-phpadmin.vir999.com", Port: 80},
		Q2Data{IP: "192.17.99.52", Host: "jsonviewer.stack.hu", Port: 7711},
	}

	fmt.Println("Q2-1")
	fmt.Printf("%+v\n", data)
	fmt.Println()

	/* Q2-2
	IP:177.18.2.33 Host:github.com Port:80
	*/
	data = append(data, Q2Data{IP: "177.18.2.33", Host: "github.com", Port: 80})

	fmt.Println("Q2-2")
	fmt.Printf("%+v\n", data)
	fmt.Println()

	// Q2-3
	updateSetting(&data[2])

	fmt.Println("Q2-3")
	fmt.Printf("%+v\n", data)
	fmt.Println()

}

func updateSetting(data *Q2Data) {
	data.Port = 80
}

func getApiResult() []byte {
	userData := `{"ret":[{"id":148500286,"parent_id":1414663,"parent":1414663,"all_parents":[1414663,59735,58971,24367,6],"domain":6,"last_login":"2018-06-12T09:17:51+0800","currency":"CNY","password_expire_at":"2020-09-04T18:19:02+0800","password_reset":false,"last_bank":54,"username":"wesley01","block":false,"bankrupt":false,"cash":{"id":71814701,"user_id":148500286,"balance":4923642.0545,"pre_sub":0,"pre_add":0,"currency":"CNY","last_entry_at":"20180612155900"},"cash_fake":null,"card":null,"enabled_card":null}],"result":"ok","profile":{"execution_time":14,"server_name":"ipl-web01.rd5.qa"}}`

	return []byte(userData)
}
