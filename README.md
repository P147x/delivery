# Delivery-CLI
Simple delivery tracker, working with the delivery services listed below.

Super excited about getting a new package, but you can't stop thinking about ? You're looking for a simple tool to yell at you when your delivery guy is moving ? I think this might be a good solution for you.

![travis](https://travis-ci.org/P147x/delivery-cli.svg?branch=master) ![goreportcards](https://goreportcard.com/badge/github.com/P147x/delivery-cli)
## Installation
First things first. Get the necessary files before.
```
git clone https://github.com/P147x/delivery-cli
```
Make sure you have Golang installed
```
go version
```

Once you are ready, build the binary and launch it !
```
go build delivery
./delivery
```
## Configuration file
This program does not allow yet to manage delivery services neither with GUI nor CLI, but actually works with a configuration file in JSON format.
Create a folder `./config`in the root of the project and create a configuration file named `services.json` inside. Your file must indented like the following exemple :

```
{
    "services": [
        {
            "delivery": "yanwen", //Delivery service name
            "trackingID": "yourtrackingID" // Your tracking ID on string format
        }
    ]
}
```
## Supported delivery services
- Yanwen 

## Built with
- [Systray](https://github.com/getlantern/systray/) - Cross-plateform Systray
- [Notificator](https://github.com/0xAX/notificator) - Cross plateform desktop notifications with Go
- [Colly](https://github.com/gocolly/colly) - Scrapping framework

## Versioning
We use the [Sermver](https://semver.org/) for versionning this project.
