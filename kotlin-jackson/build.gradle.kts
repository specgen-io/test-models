plugins {
    kotlin("jvm") version "1.5.31"
    id("io.specgen.gradle")
}

group = "io.specgen"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin:2.13.1")
    implementation("com.fasterxml.jackson.datatype:jackson-datatype-jsr310:2.13.1")
    testImplementation(kotlin("test"))
}

specgen {
    modelsKotlin {
        specFile.set(file("../spec.yaml"))
    }
}

tasks.test {
    useJUnitPlatform()
}