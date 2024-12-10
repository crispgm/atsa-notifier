# ATSA Notifier

## Concepts

- Tournament Channel: Private channel created for each tournament. e.g. `HK-ATSA50-2024-12`. For Feishu, a group will be created similarily.
- Tournament Bot: Unified bot for all tournament channels.
- Control Panel: Setup and control the bot.
- Announcements: Call players with voices.
- Message Notifications:
    - Feishu: for tournaments in China.
    - Discord: for tournaments in rest of world.

## Features

### Profile

New fields needed in Player's Profile:
- Discord Account: e.g. `username`
- User Language: `En-US`, `Zh-CN`, `Zh-TW`, `Zh-HK`, `Ja-JP`, or `Zh-SG`

### Match Notifications

- Three batches notification.
- For Feishu, Buzz at the 1st time, Buzz with text message at the 2nd time, and finally Buzz with phone call.

```plain
@ATSA-Notifier:
    [Announcement]
    @Tsoi Yu Hang @Harrod Ho 🆚 @David Zhang @Johan Hannerstal
    Open Double Qualification at Table 3

--- 30 seconds later
@ATSA-Notifier:
    [Announcement] Second Call
    @Tsoi Yu Hang @Harrod Ho 🆚 @David Zhang @Johan Hannerstal
    Open Double Qualification at Table 3

--- 2 minutes later
@ATSA-Notifier:
    [Announcement] Final Call
    @Tsoi Yu Hang @Harrod Ho 🆚 @David Zhang @Johan Hannerstal
    Open Double Qualification at Table 3
```

### Announcement

Make announcements via Web Speech API.

## Dev

### Deployment

ATSA Notifier runs locally.

Run:
```shell
go run cmd/atsa-notifier/main.go
```

### SDK

- Discord Bot
    - Doc <https://discord.com/developers/docs/intro>
    - SDK <https://github.com/bwmarrin/discordgo>
- Feishu Bot
    - Doc <https://open.feishu.cn/>
