package auth

import (
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// SaveYamlConfigToken generates a yaml token config
func SaveYamlConfigToken(c *Token, fileName string) error {

	bytes, err := yaml.Marshal(c)

	if err != nil {
		return err
	}

	ensureDir(fileName)

	return ioutil.WriteFile(fileName, bytes, 0644)
}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}
