FROM alpine:3

RUN apk add --no-cache \
    bcc-tools \
    py3-pip \
    py3-bcc \
    py3-netifaces \
    python3


WORKDIR /app
COPY . .

ENV PYTHONUNBUFFERED=1

# CMD ["python3", "http-parse.py","-i","wlp0s20f3"]
