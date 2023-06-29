# Automaic Link-Sender Bot

### This project is for practicing Go standard library and SQLite

![Gopher](https://github.com/markraiter/bot/blob/main/16530727.png)

This bot can save the link, provided by the sender and send random link by request.

This bot realized on telegram API but can be transferred to any other messanger if add the appropriate client.  All other business logic will remain unchanged.
___

To start this bot:
1. Go to telegram @BotFather and create your bot
2. Follow the instructions, provided by @BotFather and get the `token` for your bot
3. Then start the app with 
``` bash
go run cmd/main.go -telegram-bot-token 'YOUR TOKEN'
```

