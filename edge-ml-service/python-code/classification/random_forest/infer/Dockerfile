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

# Use the base image which contains all the necessary installations
ARG PYTHON_BASE=hedge-ml-python-base:latest
FROM ${PYTHON_BASE}

USER root

# Set working directory for your application
WORKDIR /hedge

# Define the root folder for the application files
ARG ROOT_FOLDER=edge-ml-service/python-code/classification/random_forest

# Copy the application-specific files
COPY ${ROOT_FOLDER}/infer/requirements.txt requirements.txt
COPY ${ROOT_FOLDER}/env.yaml classification/random_forest/env.yaml

# Create directories as needed
RUN mkdir -p /hedge/common/ && \
    mkdir -p /hedge/infer/ && \
    mkdir -p /hedge/tmp/hedge

# Install Python packages
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.txt && \
    rm -rf /root/.cache /tmp/*

# Copy additional files
COPY ${ROOT_FOLDER}/../../common/ common/
COPY ${ROOT_FOLDER}/infer/ infer/

# Set the ownership and permissions for the application directory
RUN chown -R edgex:edgex /hedge && \
    chmod a+x infer/src/main/task.py

# Use the non-root user created in the base image
USER edgex

# Set the PYTHONPATH environment variable
ENV PYTHONPATH="/hedge"

# Run the application
ENTRYPOINT ["python", "./infer/src/main/task.py"]