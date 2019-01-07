FROM openjdk:8-slim as builder

WORKDIR /app

COPY build.gradle gradlew ./
COPY gradle gradle
RUN sed -i 's/https/http/' gradle/wrapper/gradle-wrapper.properties
RUN ./gradlew downloadRepos

COPY . .
RUN ./gradlew installDist

FROM openjdk:8-alpine

RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=builder /app .

EXPOSE 9555
ENTRYPOINT ["/app/build/install/com/brymck/bin/CountService"]
