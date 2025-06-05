ARG NGINX_VERSION
FROM nginx:${NGINX_VERSION}

USER root

# Upgrade according to the security team recommendations
#
RUN apk upgrade --no-cache

# Create a group 'edgex' with GID 2001 and a user 'edgex' with UID 2002
RUN addgroup -S -g 2001 edgex && adduser -S -u 2002 -G edgex edgex

# Copy the scripts with 'edgex' user and group ownership, set permissions
COPY --chown=edgex:edgex external/nginx/njs /etc/nginx/njs
COPY --chown=edgex:edgex external/nginx/auth /etc/nginx/auth
COPY --chown=edgex:edgex external/nginx/nginx.conf /etc/nginx/nginx.conf

RUN chmod o+w /etc/nginx/auth/.htpasswd

RUN mkdir /etc/nginx/njs/logs && touch /etc/nginx/njs/logs/internal.log && \
    chmod o+w /etc/nginx/njs/logs/internal.log

# Switch to the non-root 'edgex' user
USER 2002

ENV USER_APP_MGMT_URL=http://hedge-user-app-mgmt:48111
