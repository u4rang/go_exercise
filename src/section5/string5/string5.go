// 데이터 타입 : 문자열 연산(3)
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// 예제 3. 조합(결합)
	str1 := "Oracle has consistently taken the approach that storing and managing data in a converged database is more efficient and productive than breaking up into multiple single-use engines - which inevitably results in data integrity, " +
		"consistency and security issues. " +
		"Simply put, a converged database is a multi-model, multi-tenant, multi-workload database. " +
		"Oracle Database fully supports multiple data models and access methods, simplifies consolidation while ensuring isolation, " +
		"and excels in typical database workload use cases - both operational and analytical. Click the image below for a video introduction to Oracle's converged database."
	str2 := "This is summary for Oracle..."

	// 결합 : 일반연산
	fmt.Println("str1 + str2 : ", str1+str2)

	// 결합 : Join
	strSet := []string{} // 슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)
	fmt.Println("strings.Join(strSet, \"\") : ", strings.Join(strSet, ""))

}
