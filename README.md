# http-openai-tts

An HTTP microservice using OpenAI to generate TTS.

##### *More documentation added soon!*

## Caution
Requests with the max input length can cost you between 3 and 5 cents of OpenAI credits. If
config.json's "needKey" is turned off and/or your input length is not monitored bad actors could run up an expensive bill for you or use up all of your credits extremely fast.

*When testing it took me not even 3 minutes (OpenAI only allows 50 RPM) to send enough requests (~135) to spend $5 of credits. That's about $90-100 if unsupervised for 1 hour.*

## Models (Optional, Defaults to tts-1):
- tts-1
- tts-1-hd

## Voices (Optional, Defaults to echo):
- alloy
- echo
- fable
- onyx
- nova
- shimmer

## Response formats (Optional, Defaults to mp3):
- mp3
- opus
- aac
- flac
- wav
- pcm

## Speed (Optional, Defaults to 1):
Select a value between 0.25 and 4.0.

## Text:
The maximum text is 4096 characters.

## Example requests:
- `GET /tts?model=tts-1-hd&voice=alloy&format=mp3&speed=1&text=Hello%20World`
- `GET /raw?model=tts-1&voice=echo&format=wav&speed=1.3&text=Foo%20Bar`

*If "allowHosting" is set to false /tts automatically returns the /raw output.*

## Error response format for /tts & /raw:
```json
{
	"status": "Error description",
	"message": "Information about the response.",
	"error": "Provides either an actual go error, a premade error or an OpenAI API reponse."
}
```

## Response for /tts:
```json
{
	"status": "OK",
	"message": "Information about the response.",
	"link": "Contains a link to the hosted file."
}
```
## Response for /raw:
Just returns the audio file content.

## Response of /ping:
```json
{
	"status": "OK",
	"message": "Pong!"
}
```
## Admin Controls (adding & removing keys)
- `POST /admin/add?name=Alvin&key=cm8f4a250816`
- `POST /admin/remove?name=Theodore`

## Response for /admin/add:
```json
{
	"status": "OK",
	"message": "Created the key Alvin successfully.",
	"key": "cm8f4a250816"
}
```

## Response for /admin/remove:
```json
{
	"status": "OK",
	"message": "Removed the key Theodore successfully."
}
```