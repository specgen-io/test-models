name := "echo-client"

scalaVersion := "2.13.2"
version := "0.0.1"

enablePlugins(SpecgenModels)

specgenSpecFile := file("./../spec.yaml")

libraryDependencies ++= specCirceDependencies ++ Seq(
  "org.scalatest" %% "scalatest" % "3.0.8" % Test
)

val junitxml = System.getProperty("junitxml", "target/junitxml")
(testOptions in Test) += Tests.Argument(TestFrameworks.ScalaTest, "-u", junitxml)
