# ---------- Stage 1: Build ----------
FROM gradle:8.4-jdk21 AS builder
WORKDIR /app

# Глобально отключаем toolchain-репозитории и авто-скачивание
RUN mkdir -p /home/gradle/.gradle && echo "\
org.gradle.java.installations.auto-detect=true\n\
org.gradle.java.installations.auto-download=false\n\
org.gradle.java.installations.fromEnv=JAVA_HOME\n" \
> /home/gradle/.gradle/gradle.properties

ENV JAVA_HOME=/opt/java/openjdk
ENV PATH=\"$JAVA_HOME/bin:$PATH\"

COPY build.gradle settings.gradle gradle/ ./
RUN gradle dependencies --no-daemon || true

COPY src ./src
COPY build/generated ./build/generated
RUN gradle bootJar --no-daemon -x generateProto

# ---------- Stage 2: Run ----------
FROM eclipse-temurin:21-jre-alpine
WORKDIR /app
COPY --from=builder /app/build/libs/*.jar app.jar
EXPOSE 8005
ENTRYPOINT [\"java\", \"-jar\", \"app.jar\"]
