package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"

	yaml "gopkg.in/yaml.v2"
)

type Player struct {
	Name       string  `yaml:"Name"`
	Wins       int     `yaml:"Wins"`
	Losses     int     `yaml:"Losses"`
	Draws      int     `yaml:"Draws"`
	WinPercent float64 `yaml:"WinPercent"`
}

type Session struct {
	Date    string   `yaml:"Date"`
	Winner  Player   `yaml:"Winner"`
	Players []Player `yaml:"Players"`
	Games   []Game   `yaml:"Games"`
}

type Game struct {
	Winner Player `yaml:"Winner"`
	Date   string `yaml:"Date"`
	Note   string `yaml:"Note"`
}

type Config struct {
	Bucket string `yaml:"bucket"`
	Host   string `yaml:"host"`
}

func (c *Config) LoadConfig() *Config {
	configFile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		log.Printf("Unable to open config file %v", err)
	}

	err = yaml.Unmarshal(configFile, c)

	if err != nil {
		log.Fatalf("Can't parse file %v", err)
	}

	return c
}

func DownloadFile(bucket string, object string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}

	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %w", object, err)
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	return data, nil
}

func SaveSession(bucket string, object string, record []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Printf("storage.NewClient: %+v", err)
	}

	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)

	defer wc.Close()
	wc.ContentType = "text/plain"

	_, e := wc.Write(record)
	if e != nil {
		fmt.Printf("Can't write to bucket %+v", e)
	}
}
