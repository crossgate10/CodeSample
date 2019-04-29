package regexSample

import (
"bufio"
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
"log"
"os"
"regexp"
"strings"
)

const (
	dbHost = "YourHost"
	dbUser = "UserName"
	dbPass = "Password"
	dbName = "YourDbName"
)

func main() {
	// Initialize connection.
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", dbUser, dbPass, dbHost, dbName)
	db, _ := sql.Open("mysql", connectionString)
	defer db.Close()

	path := "TxtFilePath"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jsonStr := scanner.Text()

		// Use online regex tool to test pattern. e.q. https://regex101.com/
		r, _ := regexp.Compile("\\(\\'(.*)\\', \\[(.*)\\], \\((.*)\\)\\)")
		submatch := r.FindSubmatch([]byte(jsonStr))
		//fmt.Println(string(submatch[1]))
		//fmt.Println(string(submatch[2]))
		//fmt.Println(string(submatch[3]))

		// match group processing.
		contents := strings.Split(string(submatch[3]), ",")
		name := peel(contents[0])

		// Output or write to Db
		query := fmt.Sprintf("Query with %s", name)

		//log.Print(query)
		db.Exec(query)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Print("Done.")
}

func peel(s string) string{
	ns := strings.TrimSpace(s)
	ns = strings.Replace(ns, "'", "", -1)
	ns = strings.TrimSpace(ns)
	return ns
}



