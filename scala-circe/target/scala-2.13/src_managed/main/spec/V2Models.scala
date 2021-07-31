package testservice.models.v2

import enumeratum.values._
import java.time._
import java.time.format._
import java.util.UUID
import io.circe.Codec
import io.circe.generic.extras.{Configuration, JsonKey}
import io.circe.generic.extras.semiauto.{deriveConfiguredCodec, deriveUnwrappedCodec}

case class Message(
  @JsonKey("field") field: String
)

object Message {
  implicit val config = Configuration.default
  implicit val codec: Codec[Message] = deriveConfiguredCodec
}
