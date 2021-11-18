# EpicGamesWebScraper

EpicGamesWebScraper is a CLI Go app that visits Epic games store and sends information about the free weekly game via Slack.





## Getting Started
### Running program

* The simplest way is to download the binary from [GitHub Releases](https://github.com/ahmedmansourxyz/EpicGamesWebScraper/releases) and run it (only windows). 

### Building from source

Requirements:

- [Go 1.17 or newer](https://golang.org/dl/)

```bash
$ git clone "https://github.com/ahmedmansourxyz/EpicGamesWebScraper.git"
$ cd EpicGamesWebScraper/
$ go mod init github.com/ahmedmansourxyz/EpicGamesWebScraper
$ go mod tidy
$ go build
```


### Executing program
* For a UNIX Operating system:
```bash
$ ./EpicGamesWebScraper
```
* For Windows:
```cmd
 ./EpicGamesWebScraper.exe
```
## Using program
Requirments:
- [Generating an incoming webhook URL by following Slack API documentation](https://api.slack.com/messaging/webhooks)<br>


Now you can simply execute the program, enter your webhook URL and finally specify the rate in which the slack message is sent.

***Please note that the format of input for this rate should be like the following: An integer followed by the unit of time. <br> Exmaples: 12 sec, 5 d, 30 min***	

## Screenshot of successful execution of program

![ScreenShot](https://i.postimg.cc/Nj9HkxGm/Example.png)


## Author

Ahmed Mansour

## Version History

* 0.1
    * Initial Release

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
