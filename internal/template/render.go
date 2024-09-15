package template

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/toolsascode/gomodeler/internal/helpers"
)

type Data struct {
	Env map[string]any
}

func Run() {
	var (
		templateFileList = getTemplateFileList()
		envData          = renderENV()
	)

	log.Debugf("Environments: %#v", envData)
	for i, v := range templateFileList {
		renderFile(envData, v, getOutputFile(i))
	}

}

func renderFile(data *Data, filename string, outputFile *string) {

	var (
		fileExtName = ""
	)

	if strings.Count(filename, ".") > 1 {
		fileExtNameArray := strings.Split(filename, ".")
		fileExtNameLen := len(fileExtNameArray) - 1
		fileExtName = strings.Replace(filename, fmt.Sprintf(".%s", fileExtNameArray[fileExtNameLen]), "", 1)
	}

	var t = template.Must(template.New(path.Base(filename)).Funcs(helpers.GetFuncMap()).ParseFiles([]string{filename}...))

	if err := generateFile(t, data, fileExtName, outputFile); err != nil {
		log.Fatal(err)
	}

}

func renderENV() *Data {

	var (
		envs    = helpers.ConvertFlagToKeyValue(viper.GetStringSlice("env"))
		envFile = viper.GetStringSlice("envfile")
		// storeKey  = make(map[string]interface{})
		extension = "yaml"
	)

	for _, v := range envFile {

		if strings.Contains(v, "json") {
			extension = "json"
		}

		f, err := os.Open(v)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		// storeKey = map[string]interface{}{
		// 	"environments": f,
		// }

		viper.SetConfigType(extension)
		_ = viper.MergeConfig(f)

	}

	_ = viper.MergeConfigMap(envs)

	data := &Data{
		Env: viper.GetViper().AllSettings(),
	}

	return data
}

func getOutputFile(i int) *string {
	var (
		outputFiles = viper.GetStringSlice("output.files")
		outputPath  = viper.GetString("output.path")
	)

	if len(outputFiles) > 0 {

		var (
			lenTpl = len(getTemplateFileList())
			lenOut = len(outputFiles)
		)

		if lenTpl != lenOut {
			log.Fatalf("The total number of templates is different from the total output. TPL: %d and Output: %d", lenTpl, lenOut)
			return nil
		}

		return &outputFiles[i]
	}

	if outputPath == "" {
		// outputPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		outputPath = "."
	}

	return &outputPath

}

func getTemplateFileList() []string {

	var (
		templateFiles    = viper.GetStringSlice("template.files")
		templateFileList = helpers.ReadFilesDir(viper.GetString("template.path"))
	)

	return append(templateFileList, templateFiles...)

}

// generateFile
// Args:
// <filename> output filename
func generateFile(t *template.Template, data *Data, filename string, outputFile *string) error {

	var (
		f            *os.File
		_, fileSplit = filepath.Split(filename)
		file         = filepath.Join(*outputFile, fileSplit)
	)

	if _, err := os.Stat(*outputFile); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(*outputFile, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.Create(file)
	if err != nil {
		log.Panic(err)
		return err
	}

	log.Debugf("File created: %s", file)

	if err := t.Execute(f, data.Env); err != nil {
		log.Panic(err)
		return err
	}

	err = f.Close()
	if err != nil {
		log.Panic(err)
		return err
	}

	return err
}
