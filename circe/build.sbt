name := "echo-client"

scalaVersion := "2.13.2"
version := "0.0.1"

useJCenter := true

enablePlugins(SpecCirce)

specFile := file("./../spec.yaml")

libraryDependencies ++= specSttpDependencies ++ Seq(
  "org.scalatest" %% "scalatest" % "3.0.8" % Test
)

val junitxml = System.getProperty("junitxml", "target/junitxml")
(testOptions in Test) += Tests.Argument(TestFrameworks.ScalaTest, "-u", junitxml)