package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mustInt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("두 개의 정수를 입력하세요:")

		if scanner.Scan() {
			input := scanner.Text()
			fmt.Printf("입력값: %s\n", input)
			values := strings.Fields(input)

			if len(values) != 2 {
				fmt.Println("정확히 두 개의 값을 입력해야 합니다.")
				continue
			}

			a, err1 := strconv.Atoi(values[0])
			b, err2 := strconv.Atoi(values[1])

			if err1 == nil && err2 == nil {
				fmt.Printf("입력받은 두 정수: %d, %d\n", a, b)
				break // 올바른 입력을 받았으므로 루프를 종료
			} else {
				fmt.Println("올바른 정수가 아닙니다. 다시 시도해주세요.")
			}
		} else {
			fmt.Println("입력을 읽는 중 오류가 발생했습니다.")
			return
		}
	}
}

func clearStream() {
	stdin := bufio.NewReader(os.Stdin)

	var a int = 0
	var b int = 0

	n, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
