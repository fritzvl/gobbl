package gbl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleIntegration struct {
}

func (ci *ConsoleIntegration) GenericRequest(c *Context) (GenericRequest, error) {
	genericRequest := GenericRequest{
		Text: c.RawRequest.(string),
	}

	return genericRequest, nil
}

func (ci *ConsoleIntegration) User(c *Context) (User, error) {
	user := User{
		ID:        "consoleid",
		FirstName: "John",
		LastName:  "Smith",
		Email:     "john.smith@dummymail.com",
	}

	return user, nil
}

func (ci *ConsoleIntegration) Respond(c *Context) (*interface{}, error) {
	for _, msg := range c.R.(*BasicResponse).messages {
		fmt.Printf("[bot] %s\n", msg)
	}

	return nil, nil
}

func (ci *ConsoleIntegration) Listen(bot *Bot) {
	reader := bufio.NewReader(os.Stdin)

	var input = ""
	var err error

	for input != "exit()" {
		func() {
			fmt.Print("> ")
			input, err = reader.ReadString('\n')
			if err != nil {
				panic(err)
			}

			input = strings.TrimSpace(input)

			inputCtx := InputContext{
				RawRequest:  strings.TrimSpace(input),
				Integration: ci,
				Response:    &BasicResponse{},
			}

			_, err = bot.Execute(&inputCtx)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
			}
		}()
	}

}
