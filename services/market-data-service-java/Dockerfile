FROM maven:3.9.9-eclipse-temurin-21 AS build
WORKDIR /workspace

COPY services/market-data-service-java ./services/market-data-service-java

WORKDIR /workspace/services/market-data-service-java
RUN mvn -q -DskipTests package

FROM eclipse-temurin:21-jre
WORKDIR /app

COPY --from=build /workspace/services/market-data-service-java/target/market-data-service-java-0.1.0-SNAPSHOT.jar app.jar

EXPOSE 8095

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
