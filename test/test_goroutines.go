package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//INSERT INTO `unetmap_host`
//(`id`,`ip`,`name`,`mac`,`description`,`unattendclass_id`,`type_id`,`network_id`,`department_id`,`location_id`,`noconf`,`disabled`,`seen`,`switch_id`,`port`,`next_server_id`,`license_id`,`boot_stage`,`install_linux`,`boot_httpxe`,`localboot_type`,`product_key`)
//VALUES(NULL,168430091,'test1','ffffffff','test',NULL,1,NULL,640,4662,0,0,NULL,NULL,NULL,NULL,NULL,0,0,0,0,NULL)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//я это делаю, чтобы провнерить блядский git которым я не умею пользоваться
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	/*file, err := os.Create("INSERT.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file1, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file1.Close()*/
	lines, err := readLines("1.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var ip = 167801091
	for i, line := range lines {
		text := ("(NULL," + strconv.Itoa(ip) + ",'geocol-" + strconv.Itoa(i+2) + "'" + "," + "'" + string(line) + "'" + ",' ',NULL,1,1641,992,8242,0,0,NULL,NULL,NULL,NULL,NULL,0,0,0,0,NULL),")
		fmt.Println(text)
		ip++

	}
}
