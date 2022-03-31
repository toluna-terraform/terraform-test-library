package coverage

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func WriteCovergeFiles(t *testing.T, c *terraform.Options, moduleName string) {
	if _, err := os.Stat("reports"); os.IsNotExist(err) {
		os.MkdirAll("reports", 0700) // Create your file
	}
	log.Println("Writing Generated resources for coverage verification.")
	file := []byte(terraform.RunTerraformCommand(t, c, "state", "list"))
	_ = ioutil.WriteFile("resource_list.txt", file, 0644)
	resource_file, err := os.Open("resource_list.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(resource_file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	resource_file.Close()
	os.Remove("cover.go")
	os.Remove("reports/cover.out")
	f, err := os.OpenFile("cover.go",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("package test\n\nimport \"log\"\n\nfunc check_cover(s string) {\n\t\tswitch {\n"); err != nil {
		log.Fatalln(err)
	}
	coverFile, err := os.OpenFile("reports/cover.out",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := coverFile.WriteString("mode: set\n"); err != nil {
		log.Fatalln(err)
	}
	line := 7
	for _, eachline := range txtlines {
		if !strings.Contains(eachline, "data.") {
			if err != nil {
				log.Fatalln(err)
			}
			eachline = strings.Replace(eachline, "\"", "'", -1)
			resource_name := strings.Split(eachline, ".")
			rn := fmt.Sprintf("case s == \"%s.%s\":\n\t\tlog.Println(\"check coverage\")", resource_name[len(resource_name)-2], resource_name[len(resource_name)-1])
			if _, err := f.WriteString(rn + "\n"); err != nil {
				log.Fatalln(err)
			}
			cn := fmt.Sprintf("%s/cover.go:%d.1,%d.0 1 0\n", moduleName, line, line+1)
			if _, err := coverFile.WriteString(cn); err != nil {
				log.Fatalln(err)
			}
			line = line + 2
		}
	}
	defer os.Remove("resource_list.txt")
	defer f.Close()
	if _, err := f.WriteString("\t}\n}"); err != nil {
		log.Fatalln(err)
	}
}

func MarkAsCovered(s string, moduleName string) {
	f, err := os.Open("cover.go")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	lineNum := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), s) {

			input, err := ioutil.ReadFile("reports/cover.out")
			if err != nil {
				log.Fatalln(err)
			}

			lines := strings.Split(string(input), "\n")

			for i, line := range lines {
				codeline := fmt.Sprintf("%d.1", lineNum)
				if strings.Contains(line, codeline) {
					cn := fmt.Sprintf("%s/cover.go:%d.1,%d.0 1 1", moduleName, lineNum, lineNum+1)
					lines[i] = cn
				}
			}
			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile("reports/cover.out", []byte(output), 0644)
			if err != nil {
				log.Fatalln(err)
			}

		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
