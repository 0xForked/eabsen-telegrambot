# Start from golang:1.12-alpine base image
# builder
FROM golang:1.16-alpine as builder

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk --update --no-cache add bash build-base
    

# Set the Current Working Directory inside the container
WORKDIR /telebot

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependancies to vendor folder.
RUN go mod vendor

# Build the Go app
RUN go build -o oksetda_telegram_svc .

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add tzdata && \
    mkdir /telebot

WORKDIR /telebot

COPY --from=builder /telebot/oksetda_telegram_svc /telebot
COPY --from=builder /telebot/.env /telebot

# Run the executable
CMD /telebot/oksetda_telegram_svc
