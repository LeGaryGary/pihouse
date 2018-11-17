# PiHouse

A voice/app/webapp controlled home automation system written in Golang to run on a variety of machines

### The applications included:
- pihouseserver
- pihouseclient/voice
- pihouseclient/enviroment

## Application: pihouseserver
This is a web server for the http API and react web app that will recieve data from the clients (nodes) and send instructions over websockets to the clients.

#### Requires:
- An MsSQL server - enviroment variable `SqlServerConnectionString`
- A functioning Go install

### Running on Linux:
```
export SQLSERVERCONNECTIONSTRING="sqlserver://YOUR_USERNAME_HERE:YOUR_PASSWORD_HERE@YOUR_IP_ADDRESS_HERE:YOUR_PORT_HERE?database=YOUR_PIHOUSE_DB_NAME_HERE"
go run github.com/Jordank321/pihouse/pihouseserver
```

### Running on Windows - powershell:
```
$env:SqlServerConnectionString="sqlserver://YOUR_USERNAME_HERE:YOUR_PASSWORD_HERE@YOUR_IP_ADDRESS_HERE:YOUR_PORT_HERE?database=YOUR_PIHOUSE_DB_NAME_HERE"
go run github.com/Jordank321/pihouse/pihouseserver
```

## pihouseclient/voice
This executable is used to issue commands to the pihouseserver by voice

The concept is to use [Porcupine](https://github.com/Picovoice/Porcupine) by [Picovoice](https://github.com/Picovoice/) through
the Golang bindings [here](https://github.com/charithe/porcupine-go/) by [charithe](https://github.com/charithe/)
to detect when the wake word (currently Franchesca) is said, then begin streaming the audio to the Google speech-to-text API.
The returned transcription is fed through [Wit.ai](https://wit.ai/) to produce intent tags (like `living_room_lights` and `ps4`)
and wit/on_off tags (on/off/toggle) to switch on or off lights, consoles, TVs and heating!

#### Requires:
- A wit.ai project token - enviroment variable `WIT_ACCESS_TOKEN`
- A microphone
- C libary and include files setup for porcupine (see [here](https://github.com/charithe/porcupine-go/) `Requires the pv_porcupine library to be available...` )
- A functioning Go install

### Running on Linux:

```
export WIT_ACCESS_TOKEN=YOUR_WIT_AI_TOKEN_HERE
(command which pipes 16bit little endian audio here examples coming soon) | go run github.com/Jordank321/pihouse/pihouseclient/voice
```

## pihouseclient/environment
This executable is used to update the pihouseserver on the temperature, humidity, bluetooth devices nearby and motions sensor input

... W.I.P
