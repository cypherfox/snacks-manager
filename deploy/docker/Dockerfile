# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.22-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy directory files i.e all files ending with .go
COPY ./ ./

# build tools and generate code
RUN \
  go mod download && \
  go generate ./... 

# download Go modules and dependencies
RUN \
  go mod tidy && \
  go build -o /snackmgr


##
## STEP 2 - DEPLOY
##

FROM scratch

WORKDIR /

COPY --from=build /snackmgr /snackmgr


EXPOSE 8080


ENTRYPOINT ["/godocker"]