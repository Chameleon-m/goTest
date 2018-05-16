package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"log"
	"github.com/json-iterator/go"
	"encoding/xml"
	"time"
	"os"
	"os/signal"
	"context"
	"encoding/csv"
	"fmt"
	"reflect"
	"github.com/Chameleon-m/hotellook/models"
	"github.com/Chameleon-m/hotellook/gates"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Env struct {
	db models.Datastore
}

func main() {

	db, err := models.NewDB("root:root@/hotellook?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	env := &Env{db}

	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	//router := gin.New()
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/api/v1/hotels")
	{
		v1.GET("/", env.fetchAllHotels)
		v1.POST("/gate1", env.loadJson)
		v1.POST("/gate2", env.loadCsv)
		v1.POST("/gate3", env.loadXml)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
		//ReadTimeout:    10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func (env *Env) fetchAllHotels(c *gin.Context) {

	hotels, err := env.db.HotelsFind()
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": hotels})
}

func (env *Env) loadJson(c *gin.Context) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	defer c.Request.Body.Close()
	dec := json.NewDecoder(c.Request.Body)
	// Можно было просто отлавливать ошибку и нужным спосом сообщать что изменился формат
	// но я пошел другим путём :)
	// dec.DisallowUnknownFields()
	// env.db.Begin()
	for {
		// todo изучить jsoniter и декодить по его канонам.

		var raw jsoniter.RawMessage

		if err := dec.Decode(&raw); err == io.EOF {
			break
		} else if err != nil {
			//env.db.Rollback()
			panic(err)
		}

		var gate gates.Gate1

		if err := json.Unmarshal(raw, &gate); err != nil {
			//env.db.Rollback()
			panic(err)
		}

		if err := json.Unmarshal(raw, &gate.X); err != nil {
			//env.db.Rollback()
			panic(err)
		}

		// можно и одним Unmarshal и потом приводить
		//if n, ok := gate.X["latitude"].(float64); ok {
		//	gate.Latitude = float64(n)
		//	delete(gate.X, "latitude")
		//}

		// todo сделать для полей с типом struct
		t := reflect.TypeOf(gate)
		for i := 0; i < t.NumField(); i++ {
			delete(gate.X, t.Field(i).Tag.Get("json"))
		}

		if len(gate.X) > 0 {
			// что-то делаем, (пишем в лог, отправляем алерт, кидаем в мониторинг)
			// или приводим к нужным типам и куда-то применяем.
			fmt.Printf("%+v\n", gate.X)
		}

		env.db.HotelSave(gate.GetHotel(env.db))
	}
	//env.db.Commit()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (env *Env) loadCsv(c *gin.Context) {
	defer c.Request.Body.Close()
	dec := csv.NewReader(c.Request.Body)
	//dec.FieldsPerRecord = -1
	//dec.ReuseRecord = true
	//dec.Comma = ';'
	//dec.Comment = '#'
	//var record []string
	nRecords := 0
	for {
		record, err := dec.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		var gate gates.Gate2

		nRecords++
		if nRecords == 1 {
			// todo запускат ьв отдельной горутине, отлавливать панику там
			if diff := gate.GetDiffFields(record); len(diff) > 0 {
				// что-то делаем, (пишем в лог, отправляем алерт, кидаем в мониторинг)
				log.Panicf("CSV diff fields %+v\n", diff)
			}
			continue
		}

		gate.Assign(record)

		env.db.HotelSave(gate.GetHotel(env.db))
	}

	c.Status(http.StatusOK)
}

func (env *Env) loadXml(c *gin.Context) {
	defer c.Request.Body.Close()
	dec := xml.NewDecoder(c.Request.Body);
	for {

		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.XML(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": fmt.Sprintf("decoder.Token() failed with '%s'\n", err),
			})
			break
		}

		switch v := t.(type) {

		case xml.StartElement:

			if v.Name.Local == "hotel" {
				var gate gates.Gate3
				if err = dec.DecodeElement(&gate, &v); err != nil {
					c.XML(http.StatusInternalServerError, gin.H{
						"status":  http.StatusInternalServerError,
						"message": fmt.Sprintf("decoder.DecodeElement() failed with '%s'\n", err),
					})
					break
				}

				if len(gate.X) > 0 {
					// что-то делаем, (пишем в лог, отправляем алерт, кидаем в мониторинг)
					// или приводим к нужным типам и куда-то применяем.
					for _, v := range gate.X {
						fmt.Printf("%+v\n", v)
					}
				}

				env.db.HotelSave(gate.GetHotel(env.db))
			}

		}
	}

	c.XML(http.StatusOK, gin.H{"status": http.StatusOK})
}
