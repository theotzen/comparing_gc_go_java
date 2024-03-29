# Use a base image with Java and Maven pre-installed
FROM maven:3.8.3-openjdk-17 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Maven project file
COPY pom.xml .

# Copy the source code
COPY src ./src

# Build the application
RUN mvn dependency:copy-dependencies package

# Use a smaller base image for the runtime
FROM openjdk:17-jdk-slim

# Set the working directory
WORKDIR /app

# Copy the compiled application from the build stage
COPY --from=build /app/target/java_gc-1.0-SNAPSHOT.jar /app/java_gc-1.0-SNAPSHOT.jar
COPY --from=build /app/target/dependency/* /app/lib/

# Set the entry point
ENTRYPOINT ["java", "-cp", "java_gc-1.0-SNAPSHOT.jar:lib/*", "com.javagcapp.LinkedListGenerator"]
