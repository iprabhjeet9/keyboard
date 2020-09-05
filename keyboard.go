package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func reader() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, err

}

func GetString() (string, error) {
	input, err := reader()
	return input, err
}

func GetFloat() (float64, error) {
	input, err := reader()
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func GetInt() (int, error) {
	input, err := reader()
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return number, nil
}
