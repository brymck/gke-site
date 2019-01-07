FROM python:alpine as base

FROM base as builder

COPY requirements.txt .
RUN pip install -r requirements.txt

FROM base as final

COPY --from=builder /usr/local/lib/python3.7/ /usr/local/lib/python3.7/

# Need libstdc++ for grpc
RUN apk add --no-cache libstdc++

WORKDIR /app
COPY . .

EXPOSE 8080
ENTRYPOINT ["python", "/app/app.py"]
