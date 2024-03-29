package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// a map that associates each of the named people with a message
func Hellos(names []string) (map[string]string, error) {
	//a map to associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())

}
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome yo!",
		"Hail %v! Well met yo!",
	}
	return formats[rand.Intn(len(formats))]

}
