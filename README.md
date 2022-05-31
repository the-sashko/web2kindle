# web2kindle

This application converts a web page text content into a .mobi e-book and send it to a Kindle device (via "Sent to Kindle" service)

## How to install

### Preparing

#### Email

1. You must create email in "Send to Kindle" service for your device
2. You must have additional email in service that provides access via SMTP protocol
3. You must add external email to trusted list in "Send to Kindle" service
4. In some cases "Send to Kindle" service doesn't accept email that sent from this application directly, you can resolve this issue by adding rule of redirecting email from external/self-hosted email service to "Send to Kindle" email

#### Telegram bots

1. You must create main Telegram bot for sending URLs to this application
2. You also can create additional Telegram bot for error logs

### Installation

1. `git clone https://github.com/the-sashko/web2kindle.git web2kindle`
2. `cd web2kindle`
3. `cd config`
4. `cp config_sample.json config.json`
5. `cp credentials_sample.json credentials.json`
6. Set up your data in config.json and credentials.json files. If you have no additional Telegram bot for logs, you must use credentials of your main Telegram bot instead credentials of Telegram bot for logs

#### Install with docker

1. `cd <PATH_TO_APPLICATION>`
2. `./scripts/docker/run.sh`

#### Install without docker
1. Install Golang (version 1.18+)
2. Install Calibre (https://calibre-ebook.com/download)

## How to use

### Run application

#### Run with docker container

1. `docker up web2kindle`

#### Run without docker

1. `cd <PATH_TO_APPLICATION>`
2. `./scripts/run.sh -m <MODE>`
3. Where `<MODE>` is running mode (default, test or loop)
4. *Default* - it is a single run mode. Application will process new messages from Telegram and exit. You can use this mode in cron jobs
5. *Test* - it is a test run mode. Application will process url from config and exit
6. *Loop* - it is a deamon run mode. Application will run in infinite loop

### Send URL to Telegram bot
1. Just send url as a message to Telegram bot
