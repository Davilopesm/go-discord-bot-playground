package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "!mybot"

func main() {
	godotenv.Load()
	token := os.Getenv("APP_TOKEN")

	discordConnection, failedConnection := discordgo.New("Bot " + token)

	if failedConnection != nil {
		log.Fatal(failedConnection)
	}

	discordConnection.AddHandler(func (session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}

		messageArguments := strings.Split(message.Content, " ")

		if messageArguments[0] != prefix {
			return
		}


		if messageArguments[1] == "Hello" {
			session.ChannelMessageSend(message.ChannelID, "World!")
		}


		if strings.ToLower(messageArguments[1]) == "monster" {
			monsterNames := []string{
				"Gloom Wolf", "Scarab", "Demonic Crawler", "Poison Spider", "Undead Warrior", "Grim Reaper", "Death Worm", "Efreet", "Lava Lurcher", "Serpent Spawn",
				"Blood Beast", "Banshee", "Ghoul", "Skeleton Warrior", "Giant Rat", "Ice Golem", "Fire Elemental", "Demonic Hydra", "Dragon Hatchling", "Dwarven Sentinel",
				"Gargoyle", "Lich", "Orcish Brute", "Pirate Captain", "Troll", "Skeleton", "Vampire", "War Wolf",
				"Barbed Creeper", "Blood Crab", "Dark Torturer", "Efreeti Fire Elemental", "Giant Spider", "Grim Spectre", "Lava Lurker", "Serpent Spawn",
				"Carrion Bird", "Spectre", "Deathbringer", "Demon", "Fleshcrawler",
			}

			author := discordgo.MessageEmbedAuthor{
				Name: "CipSoft",
				URL: "https://tibia.com",
			}
			embed := discordgo.MessageEmbed{
				Title: monsterNames[rand.Intn(len(monsterNames))],
				Author: &author,
			}

			if _, err := session.ChannelMessageSendEmbed(message.ChannelID, &embed); err != nil {
        log.Println("Error sending embed message:", err)
    	}
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