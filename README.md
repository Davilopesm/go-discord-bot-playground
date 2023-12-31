# Go Discord Bot Playground

## Introduction

Simple Discord bot written in Go to play around with building Discord bots and GO lang.

Idea from: <https://codingchallenges.fyi/challenges/challenge-discord>

## Features

The bot currently supports the following commands:

- `!mybot bored`: Retrieves a random activity suggestion from the Bored API (https://www.boredapi.com/) and sends it to the current channel.

- `!mybot monster`: Retrieves a random monster name from the MMORPG Tibia (saved in monsters.json) and sends it to the current channel.

- `!mybot monsters`: Lists all monsters from the MMORPG Tibia that have been added to the bot (saved in monsters.json).

- `!mybot addmonster <monsterName>`: Adds the specified `monsterName` to the bot's internal set of monsters.

## Usage

To use the bot, you will need to create a Discord application and obtain the OAuth2 client ID and client secret. You can then configure the bot to use these credentials in the `.env` file by adding the APP_TOKEN.

Once the bot is configured, you can start it by running the following command:

```bash
go run main.go
```
