build:
	GOOS=linux go build -o tts
	docker build -t gcr.io/ttsgrpc/tts:final .
	rm -f tts

push:
	docker push gcr.io/ttsgrpc/tts:final
	