package httpserver

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"wylzabc/ethagent/ethcli"
)

type InHospitalData struct {
	TimeStamp      string `json:"timeStamp"`
	HospitalId     string `json:"hospitalId"`
	DevId          string `json:"devId"`
	PersonId       string `json:"personId"`
	PersonVeinData string `json:"personVeinData"`
	PersonPicData  string `json:"personPicData"`
	AdditionalInfo string `json:"additionalInfo"`
}

func InitRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("test", getTest)
	router.GET("/testput", getTestPut)
	router.GET("/testget", getTestGet)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/rwtest.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rwtest.html", gin.H{
			"title": "Posts",
		})
	})
	router.POST("/inhospitaldata/putdata", postInHostpitalDataPutData)
	router.GET("/inhospitaldata/:personid", getInHostpitalDataGetDataByPersonid)
	router.GET("/inhospitaldata/:personid/:date", getInHostpitalDataGetDataByPersonidAndDate)

	return router
}

func getTest(c *gin.Context) {
	c.String(http.StatusOK, viper.GetString("log.errfilename"))
}

func getTestPut(c *gin.Context) {
	picfilebytes, _ := ioutil.ReadFile("./test.jpg")
	picdata := base64.StdEncoding.EncodeToString(picfilebytes)
	ethcli.PutData("abcd", "20180808080808", "123", "123", "", picdata, "")
	c.String(http.StatusOK, "test ok")
}

func getTestGet(c *gin.Context) {
	data, _ := ethcli.GetDataByPersonidAndDate("abcd", "20180808080808")
	var personData []InHospitalData
	err := json.Unmarshal([]byte(data), &personData)
	if err != nil {
		c.String(http.StatusOK, "test err")
	} else {
		ddd, _ := base64.StdEncoding.DecodeString(personData[0].PersonPicData)
		_ = ioutil.WriteFile("./output.jpg", ddd, 0666)
		c.String(http.StatusOK, personData[0].PersonId)
	}
}

func postInHostpitalDataPutData(c *gin.Context) {
	personid := c.PostForm("personid")
	timestamp := c.PostForm("timestamp")
	hospitalid := c.PostForm("hospitalid")
	devid := c.PostForm("devid")
	personveindata := c.PostForm("personveindata")
	personpicdata := c.PostForm("personpicdata")
	additionalinfo := c.PostForm("additionalinfo")

	ethcli.PutData(personid, timestamp, hospitalid, devid, personveindata, personpicdata, additionalinfo)
	//fmt.Println(personid, timestamp, hospitalid, devid, personveindata, personpicdata, additionalinfo)
	c.String(http.StatusOK, "保存成功")
	//c.Redirect(http.StatusMovedPermanently, "/rwtest.html")
}

func getInHostpitalDataGetDataByPersonid(c *gin.Context) {
	personid := c.Param("personid")
	data, _ := ethcli.GetDataByPersonid(personid)
	c.String(http.StatusOK, data)
}

func getInHostpitalDataGetDataByPersonidAndDate(c *gin.Context) {
	personid := c.Param("personid")
	date := c.Param("date")
	data, _ := ethcli.GetDataByPersonidAndDate(personid, date)
	c.String(http.StatusOK, data)
}

func Run() error {
	err := ethcli.Init()
	if err != nil {
		fmt.Printf("ethcli init error:%v\n", err)
		return err
	}

	router := InitRouter()

	go router.Run(fmt.Sprintf(":%d", 8888))

	return nil
}
