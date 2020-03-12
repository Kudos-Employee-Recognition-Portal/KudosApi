# TESTING: DO NOT DEPLOY
# Image is very large, likely in large part due to the texlive libraries, but
#   it may be worth transitioning to a multistage build, eventually.

FROM golang:latest

LABEL maintainer="Mat McDade <mathewmcdade@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN apt-get update && apt-get -y install texlive-latex-extra

COPY . .

# Test pdflatex binary: RUN pdflatex award.tex test.pdf

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]