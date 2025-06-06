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

COPY ${ROOT_FOLDER}/infer/requirements.txt requirements.txt
COPY ${ROOT_FOLDER}/env.yaml timeseries/multivariate_deepVAR/env.yaml

# Install Python packages
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.txt && \
    rm -rf /root/.cache /tmp/*

# Copy additional files
RUN mkdir -p /hedge/common/
RUN mkdir -p /hedge/infer/
RUN mkdir -p /hedge/tmp/hedge

COPY ${ROOT_FOLDER}/../../common/ common/
COPY ${ROOT_FOLDER}/infer/ infer/

ENV PYTHONPATH="/hedge"

# Set the ownership and permissions for the application directory
RUN chown -R edgex:edgex /hedge && \
    chmod a+x infer/src/main/task.py

USER edgex

ENTRYPOINT ["python", "./infer/src/main/task.py"]
