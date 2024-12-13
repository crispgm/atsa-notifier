# ATSA Notifier

[![Go](https://github.com/crispgm/atsa-notifier/actions/workflows/build.yml/badge.svg)](https://github.com/crispgm/atsa-notifier/actions/workflows/build.yml)

Nofitier of foosball tournament made for ATSA organizers.

## Features

- Sync matches in progress from Kickertool Live.
- Notify with multiple methods (TTS, Discord Webhook, and Feishu Webhook).
- Save data automatically in local storage.

## Usage

1. Download the [latest version](https://github.com/crispgm/atsa-notifier/releases)
2. Unzip `atsa-notifier.zip`
3. Configure at `./conf/conf.yml`
4. Download `players.csv` from ATSA's AirTable and put it under `./data`
5. Run:
   ```shell
   ./atsa-notifier
   ```

## Limitations

- ATSA Notifier requires modern browers, because we implement ATSA Notifier with the latest JavaScript API.
- Chrome is highly recommended, because the built-in voice synthesizers of Chrome are really good.

## Dev

Run:

```shell
go run cmd/atsa-notifier/main.go
```

## License

Copyright (c) David Zhang, 2024. [MIT License](/LICENSE).
