package development

import (
	"fmt"
	"regexp"
)

func SqlToStruct(sql string) (result string) {

	pattern := `([a-z A-Z])\s*([a-z A-Z])\s*(?<tableName>)\s*`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(sql)
	fmt.Println(match)
	return
}
