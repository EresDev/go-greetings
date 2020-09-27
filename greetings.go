package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error){
	
	if name == "" {
		return "", errors.New("empty name")
	}
	
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hello %v",
		"Great to see you %v",
		"Hey %v",
	}
	
	return formats[rand.Intn(len(formats))]
}