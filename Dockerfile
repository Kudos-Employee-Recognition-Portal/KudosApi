# TESTING: DO NOT DEPLOY

FROM golang:latest

LABEL maintainer="Mat McDade <mathewmcdade@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN apt-get update && apt-get -y install texlive-latex-extra

COPY . .

# Test pdflatex binary: RUN pdflatex test.tex test.pdf

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]