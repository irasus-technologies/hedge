#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ============================================================================

# Default value: placeholder to pull latest version for dev, will be overridden from Makefile
ARG ALPINE_BASE=alpine:latest
FROM frolvlad/alpine-glibc:alpine-3.21

# Switch to root user for installation
USER root

# Update and install basic required packages
RUN apk update && \
    apk upgrade && \
    apk add --no-cache libstdc++ hdf5-dev

# Set locale and environment variables
ENV LANG=C.UTF-8
ENV LC_ALL=en_US.utf8
ENV PYTHONPATH=/opt/hedge
ENV PATH=/opt/hedge/python/bin:$PATH

WORKDIR /opt/hedge

# Install glibc without overwriting gcompat files
#RUN apk add --no-cache libc6-compat libc-dev && \

# Download and install Miniconda
ARG SHA256_HASH=4da8dde69eca0d9bc31420349a204851bfa2a1c87aeb87fe0c05517797edaac4
RUN CONDA_VERSION="py311_24.3.0-0" && \
    wget https://repo.anaconda.com/miniconda/Miniconda3-${CONDA_VERSION}-Linux-x86_64.sh -O miniconda.sh && \
    echo "${SHA256_HASH}  miniconda.sh" | sha256sum -c - || (echo "Checksum failed!" >&2; exit 1) && \
        sh miniconda.sh -b -p /opt/hedge/python && \
        rm miniconda.sh

RUN /opt/hedge/python/bin/python -V && \
    chmod 755 /opt/hedge && \
    rm -rf /root/.wget-hsts glibc-2.38-r1-bmc.apk glibc-bin-2.38-r1-bmc.apk \
        glibc-i18n-2.38-r1-bmc.apk miniconda.sh /root/.cache/pip /opt/hedge/python/pkgs /opt/java && \
    find /opt/hedge/python/ -type f -name '*.a' -delete && \
    find /opt/hedge/python/ -type f -name '*.js.map' -delete && \
    rm -rf /opt/hedge/python/lib/python3.11/site-packages/Pillow-*-py*.egg-info && \
    rm -rf /opt/hedge/python/lib/python3.11/site-packages/future/backports/test/*.pem && \
    rm -rf /opt/hedge/python/bin/openssl \
            /opt/hedge/python/lib/python3.11/site-packages/tests/conda_env/support/requirements.txt && \
    pip uninstall cryptography -y && \
    conda clean -afy && \
    rm /opt/hedge/python/bin/xml2-config && \
    rm -rf /root/.cache /tmp/* /opt/hedge/python/share /opt/hedge/python/include /opt/hedge/python/doc

RUN pip install --upgrade pip && \
    pip install --upgrade idna tqdm certifi urllib3 watchdog flatdict

# Create a non-root user
# RUN addgroup -S mluser && adduser -S mluser -G mluser
# Create a user with specific UID and GID, need to ensure same UID and GID since we will use this to set
# right permission in hedge-ml-management and hedge-ml-sandbox
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

# Set the appropriate ownership for the directory
RUN chown -R edgex:edgex /opt/hedge

# Set user for the image
USER edgex

# Set environment variables for Python
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# Final command, to be overridden by the derived image
ENTRYPOINT ["python", "-V"]
