ARG PYTHON_BASE=hedge-ml-python-base:latest
FROM ${PYTHON_BASE}

USER root

# Set working directory for your application
WORKDIR /hedge

ARG ROOT_FOLDER=/tests

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /hedge

RUN ls -altr

RUN mkdir -p /hedge/${ROOT_FOLDER}

COPY ${ROOT_FOLDER}/requirements.txt requirements.txt

# Install Python packages
RUN pip install --no-cache-dir --upgrade pip && \
    pip install -r requirements.txt && \
    rm -rf /root/.cache /tmp/*

RUN pip install pytest-cov==6.0.0

# Copy additional files
RUN mkdir -p /hedge/anomaly/
RUN mkdir -p /hedge/classification/
RUN mkdir -p /hedge/clustering/
RUN mkdir -p /hedge/common/
RUN mkdir -p /hedge/regression/
RUN mkdir -p /hedge/timeseries/

COPY ${ROOT_FOLDER}/ tests/
COPY ${ROOT_FOLDER}/../anomaly/ anomaly/ 
COPY ${ROOT_FOLDER}/../classification/ classification/ 
COPY ${ROOT_FOLDER}/../clustering/ clustering/ 
COPY ${ROOT_FOLDER}/../common/ common/ 
COPY ${ROOT_FOLDER}/../regression/ regression/ 
COPY ${ROOT_FOLDER}/../timeseries/ timeseries/ 

COPY ${ROOT_FOLDER}/../.coveragerc .


# # Set the ownership and permissions for the application directory
# RUN chown -R mluser:mluser /hedge && \
#     chmod a+x /hedge/*

# # Use the non-root user created in the base image
# USER mluser

RUN ls -altr 

# Set the PYTHONPATH environment variable
ENV PYTHONPATH="/hedge"

ENTRYPOINT ["pytest", "--cov=.", "--cov-report=html:coverage_report", "--cov-report=term-missing" , "--cov-config=.coveragerc"]