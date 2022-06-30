// 데이터 타입 : 문자열(1)
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	//  - 큰 따움표 "", 백스쿼트 ``,
	//  - golang : 문자 char 타입 제공하지 않음 -> rune(int32 에 대한 Alias)로 문자 코드 값으로 표현
	//  - 문자 : '' 로 작성
	//  - 자주 사용하는 escape : \\, \', \", \a(콘솔별), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)

	// 예제 1.
	var str1 string = "c:\\Dev\\go_exercise\\src\\" // -> c:\Dev\go_exercise\src\
	str2 := `c:\Dev\go_exercise\src`

	fmt.Println("str1 : ", str1)
	fmt.Println("str2 : ", str2)

	// 예제 2.
	var str3 string = "Hello World!"
	var str4 string = "안녕하세요" // Len, Length
	var str5 string = "\ud55c\uae00"

	fmt.Println("str3 : ", str3)
	fmt.Println("str4 : ", str4)
	fmt.Println("str5 : ", str5)

	// 예제 3. len 함수 - 길이, 즉 Btye 수 구하는 것, 한글 한글자 3Bytes
	fmt.Println("len(\"Hello World!\") : ", len(str3))
	fmt.Println("len(\"안녕하세요\") : ", len(str4))

	// 예제 4. utf8.RuneCountInString 함수 - 실제 길이
	fmt.Println("utf8.RuneCountInString(\"Hello World!\") : ", utf8.RuneCountInString(str3))
	fmt.Println("utf8.RuneCountInString(\"안녕하세요\") : ", utf8.RuneCountInString(str4))
	fmt.Println("len([]rune(\"안녕하세요\")) : ", len([]rune(str4)))
}
