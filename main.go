package main

import (
	"bufio"
	"fmt"
	"hotel-check-in/database"
	"hotel-check-in/entities"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
	var username string
	var password string
	var userAlreadyExists bool
	var selectedHotel entities.Hotel
	var option string

	loadingMessage := color.New(color.FgGreen, color.Bold).SprintFunc()
	welcomeMessage := color.New(color.FgBlue).SprintFunc()
	userMessage := color.New(color.BgWhite, color.FgBlack, color.Bold).SprintFunc()

	fmt.Println(loadingMessage("Conectando..."))
	time.Sleep(5 * time.Second)
	fmt.Println(welcomeMessage("Comece a utilizar nosso novo app!"))

	fmt.Print("Já possui cadastro? (s/n): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userAlreadyExists = scanner.Text() == "s"

	if userAlreadyExists {
		fmt.Println(welcomeMessage("Bem-vindo de volta!"))
	} else {
		fmt.Println(welcomeMessage("Informe seu nome de usuário:"))

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		username = scanner.Text()

		fmt.Println(welcomeMessage("Informe sua senha:"))
		passwordScanner := bufio.NewScanner(os.Stdin)
		passwordScanner.Scan()
		password = passwordScanner.Text()

		fmt.Println(loadingMessage("Cadastrando..."))
		time.Sleep(5 * time.Second)

		Person := entities.Person{
			ID:       uuid.NewString(),
			Name:     username,
			Password: password,
		}

		fmt.Println(userMessage("Seja bem-vindo! ", Person.Name))

		for index, hotel := range database.MockHotelDatabase {
			fmt.Println(index+1, "-", hotel.Name)
		}

		fmt.Println("Escolha o hotel que deseja fazer seu check-in: (1-5)")

		optionScanner := bufio.NewScanner(os.Stdin)
		optionScanner.Scan()
		option = optionScanner.Text()

		optionValue, err := strconv.Atoi(option)

		if err != nil {
			panic(err)
		}

		for index, hotel := range database.MockHotelDatabase {
			if optionValue == index+1 {
				selectedHotel = hotel
			}
		}

		fmt.Println(loadingMessage("Realizando check-in para ", selectedHotel.Name, " ..."))
		time.Sleep(5 * time.Second)

		CheckIn := entities.CheckIn{
			ID:      uuid.NewString(),
			UserId:  Person.ID,
			HotelId: selectedHotel.ID,
		}

		fmt.Println(loadingMessage("Check-in realizado com sucesso! Seu código de check-in é: ", CheckIn.ID))

	}
}
