package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "!mybot"

type Activity struct {
	Activity string `json:"activity"`
}

func handleMonsters(session *discordgo.Session, message *discordgo.MessageCreate, messageArguments []string) {
	// Switch case looking to weird here as I do not wanna add more funcs

	// Get a monster as embbed message
	if strings.ToLower(messageArguments[1]) == "monster" {
		fileContent, err := os.ReadFile("monsters.json")
		if err != nil {
			log.Fatal("Error reading file")
		}

		var monsterNames []string
		err = json.Unmarshal(fileContent, &monsterNames)
		if err != nil {
			log.Fatal("Error decoding JSON:", err)
		}

		author := discordgo.MessageEmbedAuthor{
			Name: "CipSoft",
			URL:  "https://tibia.com",
		}
		embed := discordgo.MessageEmbed{
			Title:  monsterNames[rand.Intn(len(monsterNames))],
			Author: &author,
		}

		if _, err := session.ChannelMessageSendEmbed(message.ChannelID, &embed); err != nil {
			log.Println("Error sending embed message:", err)
		}
	}

	// Get all monsters
	if strings.ToLower(messageArguments[1]) == "monsters" {
		fileContent, err := os.ReadFile("monsters.json")
		if err != nil {
			log.Fatal("Error reading file")
		}

		var data []string
		err = json.Unmarshal(fileContent, &data)
		if err != nil {
			log.Fatal("Error decoding JSON:", err)
		}

		var items strings.Builder
		for _, element := range data {
			items.WriteString("\n" + element)
		}

		if _, err := session.ChannelMessageSend(message.ChannelID, items.String()); err != nil {
			log.Fatal("Error sending message:", err)
		}
	}

	// Create a monster
	if strings.ToLower(messageArguments[1]) == "addmonster" {
		monster := messageArguments[2]
		fileContent, err := os.ReadFile("monsters.json")
		if err != nil {
			log.Fatal("Error reading file")
		}

		var data []string
		err = json.Unmarshal(fileContent, &data)
		if err != nil {
			log.Fatal("Error decoding JSON:", err)
		}

		data = append(data, monster)

		dataBytes, err := json.Marshal(data)
		if err != nil {
			log.Fatal("Error decoding JSON:", err)
		}
		_ = os.WriteFile("monsters.json", dataBytes, 0644)

		if _, err := session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Monster %s added", monster)); err != nil {
			log.Fatal("Error sending message:", err)
		}
	}
}

func main() {
	godotenv.Load()
	token := os.Getenv("APP_TOKEN")

	discordConnection, failedConnection := discordgo.New("Bot " + token)

	if failedConnection != nil {
		log.Fatal(failedConnection)
	}

	discordConnection.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}

		messageArguments := strings.Split(message.Content, " ")

		if messageArguments[0] != prefix {
			return
		}

		// Consume API and transform response
		if strings.ToLower(messageArguments[1]) == "bored" {

			resp, err := http.Get("https://www.boredapi.com/api/activity")
			if err != nil {
				log.Fatal("Error getting ideas", err)
			}
			defer resp.Body.Close()

			var activity Activity
			err = json.NewDecoder(resp.Body).Decode(&activity)
			if err != nil {
				log.Fatal("Error decoding JSON:", err)
			}

			if _, err := session.ChannelMessageSend(message.ChannelID, activity.Activity); err != nil {
				log.Fatal("Error sending message:", err)
			}
		}

		// Monster Domain
		if strings.Contains(strings.ToLower(messageArguments[1]), "monster") {
			handleMonsters(session, message, messageArguments)
		}

	})

	discordConnection.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	failedConnection = discordConnection.Open()
	if failedConnection != nil {
		log.Fatal(failedConnection)
	}
	defer discordConnection.Close()

	fmt.Println("Up and running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
