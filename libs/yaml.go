package libs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Pwd struct {
	Password int `yaml:"password"`
}

type Env struct {
	GinMode string   `yaml:"gin_mode"`
	Authors []string `yaml:"authors"`
	Age     int      `yaml:"age"`
	Dev     Pwd      `yaml:"dev"`
	Test    Pwd      `yaml:"test"`
}

func readYaml(path string) (env Env) {
	content, err := os.ReadFile(path)
	checkError(err)

	err = yaml.Unmarshal(content, &env)
	checkError(err)

	fmt.Println(err, env)
	return env
}

func writeYaml(path string, env Env) {
	data, err := yaml.Marshal(env)
	checkError(err)

	err = os.WriteFile(path, data, 0777)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	path := "config/conf.yaml"

	// 读取yaml
	env := readYaml(path)

	// 修改值
	env.Age = 18
	env.Test.Password = 88

	// 写入yaml
	writeYaml(path, env)
}
