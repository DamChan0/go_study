package main

const (
	Level1 int = iota + 2 // 2
	Level2                // 3
	Level3                // 4
)

func Scaler(a int, b int, level int) int { // 매개변수 와 반환값을 정의한다와

	result := (a + b) * level
	return result
}
