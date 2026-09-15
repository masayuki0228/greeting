// Package greeting は挨拶文を作るための小さなパッケージです。
package greeting

import "fmt"

// Hello は name 宛ての挨拶文を lang の言語で返します。
// lang が "ja" なら日本語、それ以外は英語です。
// name が空の場合は "World" を使います。
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "ja" {
		return fmt.Sprintf("こんにちは、%sさん!", name)
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

// GoodMorning は name 宛ての朝の挨拶文を返します。
// name が空の場合は "World" を使います。
func GoodMorning(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Good morning, %s!", name)
}
