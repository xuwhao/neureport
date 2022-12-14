package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/neucn/neugo"
	"github.com/xuwhao/neureport/config"
)

const REPORT_URL = "https://e-report.neu.edu.cn/notes/create"
const REPORT_API_URL = "https://e-report.neu.edu.cn/api/notes"

func main() {

	// check arguments
	fpath, fname := "./", "config.json"
	if len(os.Args) == 2 && (strings.Compare(os.Args[1], "-h") == 0 || strings.Compare(os.Args[1], "--help") == 0) {
		log.Fatalf("Usage:\n neureport -fpath <path-without-filename> -fname <filename>\n")
	} else if len(os.Args) != 5 && len(os.Args) != 1 {
		log.Fatalf("Usage:\n neureport -fpath <path-without-filename> -fname <filename>\n")
	} else if len(os.Args) == 5 {
		for i := 1; i < len(os.Args); i += 2 {
			if strings.Compare(os.Args[i], "-fpath") == 0 {
				fpath = os.Args[i+1]
			} else if strings.Compare(os.Args[i], "-fname") == 0 {
				fname = os.Args[i+1]
			}
		}
	}

	// get config instance
	instance := config.GetInstance(fpath, fname)
	log.Printf("Starting report use Student ID [%s]...\n", instance.StudentID)
	client := neugo.NewSession()
	err := neugo.Use(client).WithAuth(instance.StudentID, instance.Password).Login(neugo.CAS)
	if err != nil {
		log.Fatalf("neugo login failed: %s\n", err.Error())
	}

	reported := isReported(client)
	if reported {
		fmt.Printf("You have been reported today.\n")
		os.Exit(0)
	}

	// Get token
	req, err := http.NewRequest("GET", REPORT_URL, nil)
	if err != nil {
		log.Fatalf("Create Get request failed: %s\n", err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Get report page failed, maybe not in campus network, error: %s\n", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("Incorrent status code [%d]", resp.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Create document failed: %s", err.Error())
	}
	inputs := doc.Find("input[name=_token]")
	attrList := inputs.Nodes[0].Attr
	var token string
	for _, attr := range attrList {
		if attr.Key == "value" {
			token = attr.Val
		}
	}

	// build params
	instance.SetToken(token)
	params := instance.Info
	log.Println("Send report request with params: ", params)
	paramData, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("Convert params to json failed, error: %s\n", err.Error())
	}
	fmt.Println(string(paramData))

	// make report request
	req, err = http.NewRequest("POST", REPORT_API_URL, bytes.NewReader(paramData))
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	// do report
	reportRes, err := client.Do(req)
	if err != nil {
		log.Fatalf("Do report failed, error: %s\n", err.Error())
	}

	if reportRes.StatusCode != 201 {
		log.Fatalf("Do report failed, status code: %d\n", reportRes.StatusCode)
	}

	// check result
	reported = isReported(client)
	if !reported {
		log.Fatalf("Report failed, please do it manually.\n")
	}
	// req, err = http.NewRequest("GET", REPORT_API_URL, nil)
	// req.Header.Add("Accept", "*/*")
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	// reportRes, err = client.Do(req)
	// if err != nil {
	// 	log.Fatalf("Get report result failed, error: %s\n", err.Error())
	// }
	// if reportRes.StatusCode != 200 {
	// 	log.Fatalf("Get report result failed, status code: %d\n", reportRes.StatusCode)
	// }

	// defer reportRes.Body.Close()
	// body, err := ioutil.ReadAll(reportRes.Body)
	// if err != nil {
	// 	log.Fatalf("Read report result body failed: %s", err.Error())
	// }

	// // convert result to struct
	// var result map[string][]Reported
	// err = json.Unmarshal(body, &result)
	// if err != nil {
	// 	log.Fatalf("Parse report result body failed, error: %s\n", err.Error())
	// }

	// // get result array
	// reportedList := result["data"]
	// date := time.Now().Format("2006-01-02")

	// if strings.Compare(reportedList[len(reportedList)-1].CreatedOn, date) != 0 {
	// 	log.Fatalf("Report failed, please do it manually.\n")
	// }

	log.Printf("Report successfually, See you tomorrow!")
	os.Exit(0)
}

func isReported(client *http.Client) bool {
	req, err := http.NewRequest("GET", REPORT_API_URL, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	reportRes, err := client.Do(req)
	if err != nil {
		log.Fatalf("Get report result failed, error: %s\n", err.Error())
	}
	if reportRes.StatusCode != 200 {
		log.Fatalf("Get report result failed, status code: %d\n", reportRes.StatusCode)
	}

	defer reportRes.Body.Close()
	body, err := ioutil.ReadAll(reportRes.Body)
	if err != nil {
		log.Fatalf("Read report result body failed: %s", err.Error())
	}

	// convert result to struct
	var result map[string][]Reported
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Parse report result body failed, error: %s\n", err.Error())
	}

	// get result array
	reportedList := result["data"]
	date := time.Now().Format("2006-01-02")

	if strings.Compare(reportedList[len(reportedList)-1].CreatedOn, date) != 0 {
		return false
	}
	return true
}

type Reported struct {
	CreatedOn                  string      `json:"created_on"`
	AddressArea                interface{} `json:"address_area"`
	XingchengxinxiGuojia       string      `json:"xingchengxinxi_guojia"`
	XingchengxinxiShengfen     string      `json:"xingchengxinxi_shengfen"`
	XingchengxinxiChengshi     string      `json:"xingchengxinxi_chengshi"`
	XingchengxinxiQuxian       interface{} `json:"xingchengxinxi_quxian"`
	XingchengxinxiXiangxidizhi interface{} `json:"xingchengxinxi_xiangxidizhi"`
	Credits                    int         `json:"credits"`
}
