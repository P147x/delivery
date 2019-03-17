# Delivery
A simple and cross-plateform delivery tracker, working with the delivery services listed below.

Super excited about getting a new package, but you can't stop thinking about ? You're looking for a simple tool to yell at you when your delivery guy is moving ? I think this might be a good solution for you.

![travis](https://travis-ci.org/P147x/delivery-cli.svg?branch=master) ![goreportcards](https://goreportcard.com/badge/github.com/P147x/delivery-cli)
## Installation
Did i told you this program is cross-plateform compatible ?
Get to the installation guide who matches with your system
### Linux / BSD
First things first. Assuming you already have Git, Golang and Make installed, you will need to get the necessary files.
```
git clone https://github.com/P147x/delivery-cli
cd delivery-cli
```
Once you have the files, you can use the `make` command to get the dependencies, testing (soon) and build the binary.
```
make
```
Actually, that's all. :)
### Windows
For this part, i'm assuming you are using ![Gitbash](https://gitforwindows.org/) and you have Golang installed on your computer.
In a first part, open a new terminal and use the following commands :
```
git clone https://github.com/P147x/delivery-cli
cd delivery-cli
```
Once you got the needed files, you will have to export your GOPATH and get manually the necessary dependancies (Makefile is not working on Windows.. sorry !)
```
export GOPATH=`pwd`
go get github.com/0xAX/notificator
go get github.com/getlantern/systray
go get github.com/gocolly/colly
```
Cross your fingers, you're almost done. It's compilation time !

```
go build delivery
```
Tadaa, You have got a beautiful .exe file !
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
