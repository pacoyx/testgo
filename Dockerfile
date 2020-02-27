FROM golang:1.11.5

LABEL developer="carlos.bazan" 

ENV APP_NAME go-cron-test
ENV PORT 8080

COPY . /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ["mkdir", "pdfs"]

CMD ./${APP_NAME}

EXPOSE ${PORT}