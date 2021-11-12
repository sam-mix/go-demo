package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
			/*
				// 过滤指定格式
				ok := strings.HasSuffix(fi.Name(), ".go")
				if ok {
					files = append(files, dirPth+PthSep+fi.Name())
				}

			*/
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

//获取目录里面的文件名与expr正则匹配的文件
func GetAllFilesMatch(dirPth string, expr string) (files []string, err error) {
	var reg *regexp.Regexp = nil

	if expr != "" {
		//先编译正则 提高匹配效率
		reg, err = regexp.Compile(expr)
		if err != nil {
			return nil, err
		}
	}

	return getAllFilesMatchInter(reg, dirPth)
}

//获取指定目录下的所有文件,包含子目录下的文件
func getAllFilesMatchInter(reg *regexp.Regexp, dirPth string) (files []string, err error) {
	var dirs []string

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
		} else {
			if reg != nil {
				// 过滤指定格式
				ok := reg.MatchString(fi.Name())
				if ok {
					files = append(files, dirPth+PthSep+fi.Name())
				}
			} else {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, err := getAllFilesMatchInter(reg, table)
		if err != nil {
			return nil, err
		}

		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func changeKey(fileName string) {
	buf, _ := ioutil.ReadFile(fileName)
	//log.Printf("%v", string(buf))
	if !bytes.Contains(buf, []byte("OCI")) {
		return
	}

	ok, _ := regexp.Match("OCI[a-zA-Z]", buf)
	if !ok {
		return
	}

	fmt.Printf("修改文件[%s]\n", fileName)
	buf = bytes.ReplaceAll(buf, []byte("\nOCI"), []byte("\nHT_DPI"))
	buf = bytes.ReplaceAll(buf, []byte("\rOCI"), []byte("\rHT_DPI"))
	/*
		buf = bytes.ReplaceAll(buf, []byte(";OCI"), []byte(";HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("\"OCI"), []byte("\"HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("/OCI"), []byte("/HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("@OCI"), []byte("@HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("\nOCI"), []byte("\nHT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("!OCI"), []byte("!HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("(OCI"), []byte("(HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("=OCI"), []byte("=HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte(" OCI"), []byte(" HT_DPI"))
		buf = bytes.ReplaceAll(buf, []byte("\tOCI"), []byte("\tHT_DPI"))
	*/
	ioutil.WriteFile(fileName, buf, os.ModePerm)
}

func changeKey1(fileName string) {
	buf, _ := ioutil.ReadFile(fileName)
	//log.Printf("%v", string(buf))
	if !bytes.Contains(buf, []byte("ocifetchinto")) {
		return
	}

	fmt.Printf("修改文件[%s]\n", fileName)
	buf = bytes.ReplaceAll(buf, []byte("ocifetchinto"), []byte("HT_DPIFetchInto"))

	ioutil.WriteFile(fileName, buf, os.ModePerm)
}

func changeKey2(fileName string) {
	buf, _ := ioutil.ReadFile(fileName)
	//log.Printf("%v", string(buf))
	if !bytes.Contains(buf, []byte("HT_DPIFetchinto")) {
		return
	}

	fmt.Printf("修改文件[%s]\n", fileName)
	buf = bytes.ReplaceAll(buf, []byte("HT_DPIFetchinto"), []byte("HT_DPIFetchInto"))

	ioutil.WriteFile(fileName, buf, os.ModePerm)
}

func changeKey3(fileName string) {
	buf, _ := ioutil.ReadFile(fileName)
	//log.Printf("%v", string(buf))
	if !bytes.Contains(buf, []byte("HT_DPIexecute")) {
		return
	}

	fmt.Printf("修改文件[%s]\n", fileName)
	buf = bytes.ReplaceAll(buf, []byte("HT_DPIexecute"), []byte("HT_DPIExecute"))

	ioutil.WriteFile(fileName, buf, os.ModePerm)
}

/*
ocibindbyname(
ociBindByName(
OCIBindByName(
OCICommit(
ocidefinebyname(
OCIExecute(
ocifetch(
ocifetchstatement(
OCIFreeStatement(
ocilogoff(
OCILogoff(
OCIRollback(
ocisetprefetch(
*/
func changeKey4(fileName string) {
	buf, _ := ioutil.ReadFile(fileName)
	//log.Printf("%v", string(buf))
	if !bytes.Contains(buf, []byte("oci")) {
		return
	}

	fmt.Printf("修改文件[%s]\n", fileName)
	buf = bytes.ReplaceAll(buf, []byte("ocibindbyname"), []byte("HT_DPIBindByName"))
	buf = bytes.ReplaceAll(buf, []byte("ociBindByName"), []byte("HT_DPIBindByName"))
	buf = bytes.ReplaceAll(buf, []byte("ocidefinebyname"), []byte("HT_DPIDefineByName"))
	buf = bytes.ReplaceAll(buf, []byte("ocifetch"), []byte("HT_DPIFetch"))
	buf = bytes.ReplaceAll(buf, []byte("ocifetchstatement"), []byte("HT_DPIFetchStatement"))
	buf = bytes.ReplaceAll(buf, []byte("ocilogoff"), []byte("HT_DPILogoff"))
	buf = bytes.ReplaceAll(buf, []byte("ocisetprefetch"), []byte("HT_DPISetPreFetch"))

	ioutil.WriteFile(fileName, buf, os.ModePerm)
}

func main() {
	files, dirs, _ := GetFilesAndDirs(`C:\inc_chk`)

	for _, dir := range dirs {
		fmt.Printf("获取的文件夹为[%s]\n", dir)
	}

	for _, dir := range files {
		fmt.Printf("获取的文件为[%s]\n", dir)
	}

	xfiles, err := GetAllFilesMatch(`C:\inc_chk_dpi`, `.*\.php$`)
	if err != nil {
		log.Printf("err = %v", err)
		os.Exit(1)
	}

	for _, file := range xfiles {
		//fmt.Printf("获取的文件为[%s]\n", file)
		changeKey(file)
		/*
			changeKey4(file)
			changeKey1(file)
			changeKey2(file)
			changeKey3(file)

		*/
	}
}
