
ARG test_py_major_minor_version='3.11'
FROM docker.io/python:3.11-slim as BUILD
COPY  . /tmp
WORKDIR /tmp

ARG PYTHONDONTWRITEBYTECODE="1" \
    PYTHONUNBUFFERED="1" \
    PYTHONFAULTHANDLER="1" \
    PYTHONHASHSEED="random" \
    PIP_DEFAULT_TIMEOUT="100"

# psycopg2 > libpq-dev, gcc
# pyodbc   > unixodbc-dev, g++
# git      > pip (installing remote)
RUN apt update \
    && apt install --no-install-recommends -y \
      apt-utils \
      git \
      libpq-dev \
      g++ \
      gcc \
      unixodbc-dev

RUN python3 -m pip install --upgrade pip
RUN ls -la && python3 -m pip install --no-cache-dir -r requirements.txt

RUN python3 -m pip uninstall pip wheel setuptools -y
RUN find / -type f -name "*.pyc" -exec rm -r {} \;

FROM gcr.io/distroless/python3-debian12:debug
LABEL maintainer "Oleg Butuzov <butuzov@made.ua>"

ENV PYTHONPATH=/usr/local/lib/python3.11/site-packages
COPY --from=BUILD ${PYTHONPATH} ${PYTHONPATH}

WORKDIR /app
COPY . .

CMD [ "-V" ]
ENTRYPOINT [ "python3" ]
