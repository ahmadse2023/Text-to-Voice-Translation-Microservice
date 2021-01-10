FROM alpine

RUN apk update && apk add translate-shell && apk add curl
COPY tts /tts
ENTRYPOINT [ "/tts" ]