package github.com/hantmac/juejinxiaoceToMobi/processData

import (
	"encoding/json"
	"fmt"
	"github.com/Ares002/mobi"
	"github.com/sirupsen/logrus"
	"github.com/hantmac/juejinxiaoceToMobi/config"
	"github.com/hantmac/juejinxiaoceToMobi/httpClient"
	"github.com/hantmac/juejinxiaoceToMobi/model"
	"net/url"
	"os"
)

type BookStruct struct {
	Title   string
	Content string
	Author string
}

func GetSectionsFromJueJin() ([]string, error) {
	headers := map[string]string{
		"Content-type": "application/json",
	}
	getUrl := config.Cfg().GetUrl
	jResb := model.JueJinResponse{}
	jReqb := &model.JueJinRequestBody{}
	jReqb.ID = config.Cfg().ID
	jReqb.ClientID = config.Cfg().ClientID
	jReqb.Src = config.Cfg().Src
	jReqb.Token = config.Cfg().Token
	jReqb.UID = config.Cfg().UID
	var rq = url.Values{}
	rq.Add("id", jReqb.ID)
	rq.Add("client_id", jReqb.ClientID)
	rq.Add("src", jReqb.Src)
	rq.Add("uid", jReqb.UID)
	rq.Add("token", jReqb.Token)
	reqBody := rq.Encode()

	jResbJson, _, err := httpClient.DoRequest("GET", getUrl+"?"+reqBody, headers, nil, 20)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jResbJson, &jResb)
	if err != nil {
		logrus.WithError(err).Fatal("unmarshal cmsResponse failed")
		return nil, err
	}

	fmt.Println(jResb)

	return jResb.D.Section, nil
}

func (book *BookStruct) GetSectionData(sectionId string) {
	headers := map[string]string{
		"Content-type": "application/json",
	}
	sectionUrl := config.Cfg().GetSectionUrl

	sectionRes := model.GetSectionResponse{}
	jReq := &model.JueJinRequestBody{}
	jReq.ID = config.Cfg().ID
	jReq.ClientID = config.Cfg().ClientID
	jReq.Src = config.Cfg().Src
	jReq.Token = config.Cfg().Token
	jReq.UID = config.Cfg().UID
	var rq = url.Values{}
	rq.Add("id", jReq.ID)
	rq.Add("client_id", jReq.ClientID)
	rq.Add("src", jReq.Src)
	rq.Add("uid", jReq.UID)
	rq.Add("token", jReq.Token)
	rq.Add("sectionId", sectionId)
	reqBody := rq.Encode()
	jResbJson, _, err := httpClient.DoRequest("GET", sectionUrl+"?"+reqBody, headers, nil, 20)
	if err != nil {

	}

	err = json.Unmarshal(jResbJson, &sectionRes)
	if err != nil {
		logrus.WithError(err).Fatal("unmarshal cmsResponse failed")
	}

	book.Title = sectionRes.D.Title
	book.Content = sectionRes.D.Content

}

func DownloadAndConvert() {
	book := &BookStruct{}
	m, err := mobi.NewWriter(config.Cfg().Title+".mobi")
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(config.Cfg().Title, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
	} else {
		fmt.Printf("mkdir success!\n")
	}

	m.Title(config.Cfg().Title)
	m.Compression(mobi.CompressionNone) // LZ77 compression is also possible using  mobi.CompressionPalmDoc
	// Meta data
	m.NewExthRecord(mobi.EXTH_DOCTYPE, "mobi")
	m.NewExthRecord(mobi.EXTH_AUTHOR, "github.com/hantmac")
	sectionSlc, err := GetSectionsFromJueJin()
	if err != nil {
		logrus.WithError(err).Fatal("Get section from juejin failed")
	}

	for _, sectionId := range sectionSlc {
		book.GetSectionData(sectionId)
		f,err := os.Create("./"+config.Cfg().Title+"/"+book.Title+".md")

		defer f.Close()

		if err !=nil {
			fmt.Println(err.Error())
			} else {
			_,err = f.Write([]byte(book.Content))
			fmt.Println(book.Title + "已写入并转换成md")
			}

		m.NewChapter(book.Title, []byte(book.Content))
		fmt.Println(book.Title + "已下载并转换成mobi")
	}
	m.Write()
	fmt.Println("下载md,mobi&转换完毕")
}
