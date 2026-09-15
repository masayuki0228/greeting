// Package greeting は挨拶文を作るための小さなパッケージです。
package greeting

import "fmt"

// Hello は name 宛ての挨拶文を返します。
// name が空の場合は "World" を使います。
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// Goodbye は name 宛ての別れの挨拶文を返します。
// name が空の場合は "World" を使います。
func Goodbye(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Goodbye, %s!", name)
}
