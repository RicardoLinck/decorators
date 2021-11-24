package cmd

import (
	"log"
	"os"

	"github.com/RicardoLinck/decorators/cache"
	"github.com/RicardoLinck/decorators/service"
)

type runner interface {
	GetData(keys ...string) error
}

type defaultRunner struct {
	url string
}

func (d *defaultRunner) GetData(keys ...string) error {
	c := service.NewClient(d.url)

	cc := cache.NewCachedDataGetter(c)
	for _, k := range keys {
		log.Print(cc.GetData(k))
	}

	return nil
}

type dryRunner struct {
	runner
}

func (d *dryRunner) GetData(keys ...string) error {
	log.Default().SetOutput(os.Stdout)
	return d.runner.GetData(keys...)
}

type fileRunner struct {
	runner
	filePath string
}

func (f *fileRunner) GetData(keys ...string) error {
	file, err := os.Create(f.filePath)
	if err != nil {
		return err
	}

	log.Default().SetOutput(file)
	return f.runner.GetData(keys...)
}
