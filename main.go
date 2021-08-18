package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func VersionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}

func compareVersion(version1 string, version2 string) int {
	versionA := strings.Split(version1, ".")
	versionB := strings.Split(version2, ".")

	for i := len(versionA); i < 4; i++ {
		versionA = append(versionA, "0")
	}
	fmt.Println(versionA)
	for i := len(versionB); i < 4; i++ {
		versionB = append(versionB, "0")
	}
	fmt.Println(versionB)
	for i := 0; i < 4; i++ {
		version1, _ := strconv.Atoi(versionA[i])
		version2, _ := strconv.Atoi(versionB[i])
		if version1 == version2 {
			continue
		} else if version1 > version2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func main() {
	fmt.Println(strings.TrimPrefix("/home/gavin/Documents/go/src/dataStructure/temp/chartdemo", "/home/gavin/Documents/go/src/dataStructure/temp/"))
	fmt.Println(getPrePath("/home/gavin/Documents/go/src/dataStructure/temp/chartdemo"))
}

//输入：./data/catalog/mychart/chartdemo/chartdemo-0.1.0.tgz
//返回：./data/catalog/mychart/chartdemo/
func getPrePath(fullPath string) (string, string) {
	comma := strings.LastIndex(fullPath, "/")
	return fullPath[:comma], fullPath[comma+1:]
}

// 读取文件夹内所有文件内容，忽略文件夹
func ReadFiles(src string) map[string]string {
	mapp := make(map[string]string)
	var paths []string
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, path := range paths {
		a, _ := ioutil.ReadFile(path)
		mapp[path] = string(a)
	}
	return mapp
}

//解压tgz压缩包
func DeCompress(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

//执行命令行，返回结果
func execCommand(cmd string) (*bytes.Buffer, error) {
	c := exec.Command("/bin/bash", "-c", cmd)
	var out bytes.Buffer
	c.Stdout = &out
	if err := c.Run(); err != nil {
		return nil, err
	}
	return &out, nil
}

//value.yaml配置文件对应的结构体
type Value struct {
	ReplicaCount       int                    `yaml:"replicaCount"`
	Image              Image                  `yaml:"image"`
	ImagePullSecrets   []string               `yaml:"imagePullSecrets"`
	NameOverride       string                 `yaml:"nameOverride"`
	FullnameOverride   string                 `yaml:"fullnameOverride"`
	ServiceAccount     ServiceAccount         `yaml:"serviceAccount"`
	PodAnnotations     map[string]interface{} `yaml:"podAnnotations"`
	PodSecurityContext map[string]interface{} `yaml:"podSecurityContext"`
	SecurityContextmap map[string]interface{} `yaml:"securityContext"`
	Service            Service                `yaml:"service"`
	Ingress            Ingress                `yaml:"ingress"`
	Resources          map[string]interface{} `yaml:"resources"`
	Autoscaling        Autoscaling            `yaml:"autoscaling"`
	NodeSelector       map[string]interface{} `yaml:"nodeSelector"`
	Tolerations        []string               `yaml:"tolerations"`
	Affinity           map[string]interface{} `yaml:"affinity"`
}

type Image struct {
	Repository string `yaml:"repository"`
	PullPolicy string `yaml:"pullPolicy"`
	Tag        string `yaml:"tag"`
}

type ServiceAccount struct {
	Create      string                 `yaml:"create"`
	Annotations map[string]interface{} `yaml:"annotations"`
	Name        string                 `yaml:"name"`
}

type Service struct {
	Type string `yaml:"type"`
	Port int32  `yaml:"port"`
}

type Ingress struct {
	Enabled     bool                   `yaml:"enabled"`
	ClassName   string                 `yaml:"className"`
	Annotations map[string]interface{} `yaml:"annotations"`
	Hosts       []Host                 `yaml:"hosts"`
	Tls         []string               `yaml:"tls"`
}

type Host struct {
	Host  string `yaml:"host"`
	Paths []Path `yaml:"paths"`
}

type Path struct {
	Path     string `yaml:"path"`
	PathType string `yaml:"pathType"`
}

type Autoscaling struct {
	Enabled                        bool `yaml:"enabled"`
	MinReplicas                    int  `yaml:"minReplicas"`
	MaxReplicas                    int  `yaml:"maxReplicas"`
	TargetCPUUtilizationPercentage int  `yaml:"targetCPUUtilizationPercentage"`
}

// type Image struct {
// 	Repository string `yaml:"repository" json:"repository"`
// 	PullPolicy string `yaml:"pullPolicy" json:"pullPolicy"`
// 	Tag        string `yaml:"tag" json:"tag"`
// }

// type SiteConfig struct {
// 	HttpPort  int
// 	HttpsOn   bool
// 	Domain    string
// 	HttpsPort int
// }

// type NginxConfig struct {
// 	Port   int
// 	Logsrc string
// 	src    string
// }
// type ConfigEngine struct {
// 	data map[interface{}]interface{}
// }

// // 将ymal文件中的内容进行加载
// func (c *ConfigEngine) Load(src string) error {
// 	ext := c.guessFileType(src)
// 	if ext == "" {
// 		return fmt.Errorf("cant not load %s", src)
// 	}
// 	return c.loadFromYaml(src)
// }

// // 更新到yaml文件中，全量覆盖
// func (c *ConfigEngine) Update(src string) error {
// 	err := c.writeToYaml(src)
// 	if err != nil {
// 		return fmt.Errorf("cant not update %s,%s", src, err)
// 	}
// 	return nil
// }

// //判断配置文件名是否为yaml格式
// func (c *ConfigEngine) guessFileType(src string) string {
// 	s := strings.Split(src, ".")
// 	ext := s[len(s)-1]
// 	switch ext {
// 	case "yaml", "yml":
// 		return "yaml"
// 	}
// 	return ""
// }

// // 将配置yaml文件中的进行加载
// func (c *ConfigEngine) loadFromYaml(src string) error {
// 	yamlS, readErr := ioutil.ReadFile(src)
// 	if readErr != nil {
// 		return readErr
// 	}
// 	// yaml解析的时候c.data如果没有被初始化，会自动为你做初始化
// 	return yaml.Unmarshal(yamlS, &c.data)
// }

// func (c *ConfigEngine) writeToYaml(src string) error {
// 	data, err := yaml.Marshal(c.data) // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
// 	if err != nil {
// 		return errors.New("can not update " + src + " config")
// 	}
// 	return ioutil.WriteFile(src, data, 0777)
// }

// // 从配置文件中获取值
// func (c *ConfigEngine) Get(name string) interface{} {
// 	src := strings.Split(name, ".")
// 	data := c.data
// 	for key, value := range src {
// 		v, ok := data[value]
// 		if !ok {
// 			break
// 		}
// 		if (key + 1) == len(src) {
// 			return v
// 		}
// 		if reflect.TypeOf(v).String() == "map[interface {}]interface {}" {
// 			data = v.(map[interface{}]interface{})
// 		}
// 	}
// 	return nil
// }

// // 从配置文件中获取string类型的值
// func (c *ConfigEngine) GetString(name string) string {
// 	value := c.Get(name)
// 	switch value := value.(type) {
// 	case string:
// 		return value
// 	case bool, float64, int:
// 		return fmt.Sprint(value)
// 	default:
// 		return ""
// 	}
// }

// // 从配置文件中获取int类型的值
// func (c *ConfigEngine) GetInt(name string) int {
// 	value := c.Get(name)
// 	switch value := value.(type) {
// 	case string:
// 		i, _ := strconv.Atoi(value)
// 		return i
// 	case int:
// 		return value
// 	case bool:
// 		if value {
// 			return 1
// 		}
// 		return 0
// 	case float64:
// 		return int(value)
// 	default:
// 		return 0
// 	}
// }

// // 从配置文件中获取bool类型的值
// func (c *ConfigEngine) GetBool(name string) bool {
// 	value := c.Get(name)
// 	switch value := value.(type) {
// 	case string:
// 		str, _ := strconv.ParseBool(value)
// 		return str
// 	case int:
// 		if value != 0 {
// 			return true
// 		}
// 		return false
// 	case bool:
// 		return value
// 	case float64:
// 		if value != 0.0 {
// 			return true
// 		}
// 		return false
// 	default:
// 		return false
// 	}
// }

// // 从配置文件中获取Float64类型的值
// func (c *ConfigEngine) GetFloat64(name string) float64 {
// 	value := c.Get(name)
// 	switch value := value.(type) {
// 	case string:
// 		str, _ := strconv.ParseFloat(value, 64)
// 		return str
// 	case int:
// 		return float64(value)
// 	case bool:
// 		if value {
// 			return float64(1)
// 		}
// 		return float64(0)
// 	case float64:
// 		return value
// 	default:
// 		return float64(0)
// 	}
// }

// // 从配置文件中获取Struct类型的值,这里的struct是你自己定义的根据配置文件
// func (c *ConfigEngine) GetStruct(name string, s interface{}) interface{} {
// 	d := c.Get(name)
// 	switch d.(type) {
// 	case string:
// 		c.setField(s, name, d)
// 	case map[interface{}]interface{}:
// 		c.mapToStruct(d.(map[interface{}]interface{}), s)
// 	}
// 	return s
// }

// // name在yaml中是小写的，但是我们的结构体首字母是大写，需要转换一下
// func (c *ConfigEngine) mapToStruct(m map[interface{}]interface{}, s interface{}) interface{} {
// 	for key, value := range m {
// 		switch key.(type) {
// 		case string:
// 			c.setField(s, Ucfirst(key.(string)), value)
// 		}
// 	}
// 	return s
// }

// // 这部分代码是重点，需要多看看
// func (c *ConfigEngine) setField(obj interface{}, name string, value interface{}) error {
// 	// reflect.Indirect 返回value对应的值
// 	structValue := reflect.Indirect(reflect.ValueOf(obj))
// 	structFieldValue := structValue.FieldByName(name)

// 	// isValid 显示的测试一个空指针
// 	if !structFieldValue.IsValid() {
// 		return fmt.Errorf("No such field: %s in obj", name)
// 	}

// 	// CanSet判断值是否可以被更改
// 	if !structFieldValue.CanSet() {
// 		return fmt.Errorf("Cannot set %s field value", name)
// 	}

// 	structFieldType := structFieldValue.Type() // 获取要更改值的类型
// 	val := reflect.ValueOf(value)              //需要更改的值
// 	//设置值
// 	if structFieldType.Kind() == reflect.Struct && val.Kind() == reflect.Map { //处理复合对象
// 		vint := val.Interface()
// 		switch vint.(type) {
// 		case map[interface{}]interface{}:
// 			for key, value := range vint.(map[interface{}]interface{}) {
// 				c.setField(structFieldValue.Addr().Interface(), key.(string), value)
// 			}
// 		case map[string]interface{}:
// 			for key, value := range vint.(map[string]interface{}) {
// 				c.setField(structFieldValue.Addr().Interface(), key, value)
// 			}
// 		}
// 	} else { //处理对象
// 		if structFieldType != val.Type() { //key与value类型不符合
// 			return errors.New("Provided value type didn't match obj field type")
// 		}
// 		structFieldValue.Set(val)
// 	}
// 	return nil
// }

// //首字母改成大写
// func Ucfirst(str string) string {
// 	for i, v := range str {
// 		return string(unicode.ToUpper(v)) + str[i+1:]
// 	}
// 	return ""
// }
