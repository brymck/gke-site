FROM openjdk:8-slim as builder

WORKDIR /app

# Download Gradle. HTTPS may not work over a proxy, and --version is a no-op so we only download Gradle. This way, we
# should not re-download Gradle even if we change the build script.
COPY gradlew .
COPY gradle gradle
RUN sed -i 's/https/http/' gradle/wrapper/gradle-wrapper.properties
RUN ./gradlew --version

# Download dependencies
COPY build.gradle env.gradle ./
RUN ./gradlew --init-script env.gradle downloadRepos

# Install distribution
COPY . .
RUN ./gradlew installDist

FROM openjdk:8-alpine

RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=builder /app .

EXPOSE 9555
ENTRYPOINT ["/app/build/install/com/brymck/bin/CountService"]
