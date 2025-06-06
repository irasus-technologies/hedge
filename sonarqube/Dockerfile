#
# Copyright (c) 2023 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
####

# Default values: placeholders to pull latest version for dev, will be overridden from Makefile
ARG NODE_BASE=node:latest
ARG GO_BASE=golang:latest
ARG ALPINE_BASE=alpine:latest
ARG PYTHON_TEST_COVERAGE_BASE=hedge-python-test-coverage-base:latest

FROM ${GO_BASE} AS stage1_gocoverage

ARG ALPINE_PKG_BASE="make git"
ARG ALPINE_PKG_EXTRA=""

RUN apk add --update --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

WORKDIR /hedge

COPY ../go.mod .
COPY ../go.sum .

RUN go mod download

COPY .. .

# Delete the old report if it exists
RUN rm -f sonarqube/coverage.go.out

# Generate coverage report
RUN go test ./... -coverprofile=sonarqube/coverage.go.out -coverpkg=./...;


FROM ${PYTHON_TEST_COVERAGE_BASE} AS stage2_pycoverage
#FROM --platform=linux/amd64 python:3.11 AS stage2_pycoverage

## copy from prev step(s) to enforce sequential build
COPY --from=stage1_gocoverage /hedge/sonarqube/coverage.go.out /hedge/sonarqube/coverage.go.out
#
WORKDIR /hedge/edge-ml-service

# Copy the whole python-code directory
COPY ../edge-ml-service/python-code python-code

## Delete the old report if it exists
RUN rm -f ../sonarqube/coverage.py.out

WORKDIR /hedge/edge-ml-service/python-code/
#
### Upgrade pip with no cache
RUN pip install --no-cache-dir -U pip
RUN pip install --upgrade pip setuptools wheel

## Commented out the installation of requirements
## RUN python -m pip install -r tests/requirements.txt
RUN coverage run -m pytest tests && coverage xml -i -o ../../sonarqube/coverage.py.out


#FROM ${NODE_BASE} AS stage3_uicoverage
#
#RUN apk add --no-cache \
#      chromium \
#      nss \
#      freetype \
#      harfbuzz \
#      ca-certificates \
#      ttf-freefont \
#      nodejs \
#      yarn
#
#ENV CHROME_BIN=/usr/bin/chromium-browser
#ENV NODE_OPTIONS="--max_old_space_size=8192"
#
#WORKDIR /hedge/ui/edge-portal
#
#ADD ./ui/edge-portal/*.* ./
#COPY ../ui/edge-portal/projects ./projects
#
# RUN npm install -g @angular/cli@18.2.6
#
#RUN npm install
#
#RUN npm run test-headless;

## Run sonarqube scanner
FROM sonarsource/sonar-scanner-cli:5.0.1

WORKDIR /hedge

# Copy the UI files from the previous stages
#COPY --from=stage3_pycoverage /hedge/sonarqube/coverage.ui.out ./sonarqube/coverage.ui.out
COPY --from=stage2_pycoverage /hedge/sonarqube/coverage.go.out ./sonarqube/coverage.go.out
COPY --from=stage2_pycoverage /hedge/sonarqube/coverage.py.out ./sonarqube/coverage.py.out

COPY .. .
ADD ./sonarqube/run.sh .

# Run UI sonarqube Coverage
CMD ["./run.sh"]
