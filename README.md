# http-openai-tts

An HTTP microservice using OpenAI to generate TTS.

##### *More documentation added soon!*

## Caution
Requests with the max input length can cost you between 3 and 5 cents of OpenAI credits. If
config.json's "needKey" is turned off and/or your input length is not monitored bad actors could run up an expensive bill for you or use up all of your credits extremely fast.

*When testing it took me not even 3 minutes (OpenAI only allows 50 RPM) to send enough requests (~135) to spend $5 of credits. That's about $90-100 if unsupervised for 1 hour.*

## Models:
- tts-1
- tts-1-hd

## Voices:
- alloy
- echo
- fable
- onyx
- nova
- shimmer

## Response formats:
- mp3
- opus
- aac
- flac
- wav
- pcm

## Speed:
Select a value between 0.25 and 4.0.

## Input:
The maximum input is 4096 characters.

## Example requests:
- `GET /tts?model=tts-1-hd&voice=alloy&format=mp3&speed=1&text=Hello%20World`
- `GET /raw?model=tts-1&voice=echo&format=wav&speed=1.3&text=Foo%20Bar`

## Error response format for /tts & /raw:
```json
{
	"status": "Error description",
	"message": "Information about the response.",
	"error": "This is not always provided."
}
```

## Response for /tts:
```json
{
	"status": "OK",
	"message": "Information about the response.",
	"link": "Contains a link to the hosted file. (Only appears if status is 'OK')"
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
