FROM maven:3.9.9-eclipse-temurin-21 AS build
WORKDIR /workspace

COPY proto ./proto
COPY services/notification-service-java ./services/notification-service-java

WORKDIR /workspace/services/notification-service-java

RUN mvn -q -DskipTests package

FROM eclipse-temurin:21-jre
WORKDIR /app

COPY --from=build /workspace/services/notification-service-java/target/notification-service-java-0.1.0-SNAPSHOT.jar app.jar

EXPOSE 8094
EXPOSE 10084

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
