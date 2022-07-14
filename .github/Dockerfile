
# Copyright 2022 Kirill Scherba <kirill@scherba.ru>.  All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
# Teonet teofortune-web server application docker file
#
# Docker build (included private repositories):
# --------------------------------------------
#
#   docker build -t teofortune-web -f ./.github/Dockerfile .
#
# Docker test run:
# ---------------
#
#   docker run --network host -it --rm teofortune-web teofortune-web
#
# Publish to github:
# -----------------
#
#   docker login ghcr.io -u USERNAME -p TOKEN_FOR_PACKAGES
#   #
#   docker tag teocfg ghcr.io/teonet-go/teofortune-web:0.0.0
#   docker push ghcr.io/teonet-go/teofortune-web:0.0.0
#   #
#   docker tag teofortune-web ghcr.io/teonet-go/teofortune-web:0.0.0
#   docker push ghcr.io/teonet-go/teofortune-web:0.0.0
#
# Run from gihub repo:
# -------------------
#
#   docker run --network=host -it --restart=always --name teofortune-web \
#     -v $HOME/.config/teonet/teofortune-web:/root/.config/teonet/teofortune-web \
#     ghcr.io/teonet-go/teofortune-web:0.0.1 teofortune-web -loglevel=debug
#

# Build stage
#
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .

# Add from arguments
ARG version
ENV version=$version

RUN go get
RUN go install ./

# Final stage
#
FROM alpine:latest
RUN apk --no-cache add ca-certificates fortune
COPY --from=builder /go/bin/* /usr/local/bin/
CMD "teofortune-web"
LABEL Name=teofortune-web Version=$(version)
