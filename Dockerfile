FROM golang:1.16.5-stretch

WORKDIR /go/src/market-app
COPY . .

RUN CGO_ENABLED=0 go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...

RUN mkdir ~/.aws/
RUN touch ~/.aws/credentials

RUN printf "[default]\naws_access_key_id=FAKE_AWS_ACCESS_KEY_ID_PREPROD \n\
aws_secret_access_key=FAKE_AWS_ACCESS_SECRET_KEY_PREPROD" >> ~/.aws/credentials

CMD ["market-app"]