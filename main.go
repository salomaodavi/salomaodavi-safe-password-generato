package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Conjuntos de caracteres
const (
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars  = "0123456789"
	specialChars = "!@#$%&*()-_=+[]{}<>?/|"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🔐 Safe Password Generator - GoLang 🚀")
	fmt.Println("--------------------------------------")

	// Input: tamanho da senha
	length := getInputInt("➡️  Informe o tamanho da senha: ")

	// Input: quais caracteres incluir
	includeLower := getInputBool("➡️  Incluir letras minúsculas? (s/n): ")
	includeUpper := getInputBool("➡️  Incluir letras maiúsculas? (s/n): ")
	includeNumbers := getInputBool("➡️  Incluir números? (s/n): ")
	includeSpecial := getInputBool("➡️  Incluir caracteres especiais? (s/n): ")

	// Verificar se pelo menos uma opção foi escolhida
	if !includeLower && !includeUpper && !includeNumbers && !includeSpecial {
		fmt.Println("❌ Você deve selecionar pelo menos uma opção de caractere.")
		return
	}

	// Gerar senha
	password := generatePassword(length, includeLower, includeUpper, includeNumbers, includeSpecial)

	fmt.Printf("\n🔑 Senha gerada: %s\n", password)

	// Perguntar se quer salvar em arquivo
	save := getInputBool("💾 Deseja salvar essa senha no arquivo passwords.txt? (s/n): ")
	if save {
		saveToFile(password)
		fmt.Println("✅ Senha salva em passwords.txt")
	}
}

// Função para gerar senha
func generatePassword(length int, lower, upper, number, special bool) string {
	allChars := ""
	if lower {
		allChars += lowerChars
	}
	if upper {
		allChars += upperChars
	}
	if number {
		allChars += numberChars
	}
	if special {
		allChars += specialChars
	}

	chars := []rune(allChars)
	password := make([]rune, length)

	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

// Função para pegar input inteiro
func getInputInt(prompt string) int {
	var value int
	fmt.Print(prompt)
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println("❌ Entrada inválida. Tente novamente.")
		os.Exit(1)
	}
	return value
}

// Função para pegar input sim/não
func getInputBool(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	return input == "s" || input == "sim"
}

// Função para salvar senha no arquivo
func saveToFile(password string) {
	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Erro ao salvar no arquivo:", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	_, err = file.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, password))
	if err != nil {
		fmt.Println("❌ Erro ao escrever no arquivo:", err)
	}
}
