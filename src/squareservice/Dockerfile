FROM python:3.7-alpine as base

FROM base as builder

# Install general Python development dependencies
RUN apk add --update --no-cache \
    gcc \
    g++ \
    linux-headers \
    make \
    musl-dev \
    python-dev

# Install Python dependencies
WORKDIR /src
COPY requirements.txt .
RUN pip install -r requirements.txt

# Install app
COPY . .
RUN pip install .

FROM base as final

# Enable unbuffered logging
ENV PYTHONUNBUFFERED=1

# We need libstdc++ for grpc
RUN apk add --no-cache libstdc++

COPY --from=builder /usr/local/bin/squareservice /usr/local/bin/squareservice
COPY --from=builder /usr/local/lib/python3.7/ /usr/local/lib/python3.7/

EXPOSE 5000
ENTRYPOINT ["/usr/local/bin/squareservice"]
