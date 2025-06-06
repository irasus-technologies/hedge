ARG PYTHON_BASE=hedge-ml-python-base:latest
FROM ${PYTHON_BASE}

USER root

# Set working directory for your application
WORKDIR /hedge

ARG ROOT_FOLDER=edge-ml-service/python-code/timeseries/multivariate_deepVAR

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /hedge

RUN ls -altr 

RUN mkdir -p /hedge/${ROOT_FOLDER}

COPY ${ROOT_FOLDER}/train/requirements.txt requirements.txt
COPY ${ROOT_FOLDER}/env.yaml timeseries/multivariate_deepVAR/env.yaml
COPY ${ROOT_FOLDER}/train/src/main/tuning.py timeseries/multivariate_deepVAR/train/src/main/tuning.py

ENV LOCAL=False

# Install Python packages
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.txt && \
    rm -rf /root/.cache /tmp/*

# Copy additional files
RUN mkdir -p /hedge/common/
RUN mkdir -p /hedge/train/
RUN mkdir -p /hedge/tmp/hedge

COPY ${ROOT_FOLDER}/../../common/ common/
COPY ${ROOT_FOLDER}/train/ train/

# Set the ownership and permissions for the application directory
RUN chown -R edgex:edgex /hedge && \
    chmod a+x train/src/main/task.py

# Use the non-root user created in the base image
USER edgex

# Set the PYTHONPATH environment variable
ENV PYTHONPATH="/hedge"

ENTRYPOINT ["python", "./train/src/main/task.py"]
