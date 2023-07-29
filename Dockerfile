FROM golang:1.16

ARG USERNAME=www-data
ARG USER_UID=1000
ARG USER_GID=$USER_UID
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u $USER_UID $USERNAME
RUN mkdir -p /var/www/.cache
RUN chown -R $USERNAME:$USERNAME /go
RUN chown -R $USERNAME:$USERNAME /var/www/.cache
USER $USERNAME

CMD ["tail", "-f", "/dev/null"]