package test_service.models;

import java.time.*;
import java.util.*;
import java.math.BigDecimal;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonSubTypes.*;

public class MapFields {
	@JsonProperty("int_map_field")
	private Map<String, Integer> intMapField;
	@JsonProperty("string_map_field")
	private Map<String, String> stringMapField;

	public MapFields() {
	}

	public MapFields(Map<String, Integer> intMapField, Map<String, String> stringMapField) {
		this.intMapField = intMapField;
		this.stringMapField = stringMapField;
	}

	public Map<String, Integer> getIntMapField() {
		return intMapField;
	}

	public void setIntMapField(Map<String, Integer> intMapField) {
		this.intMapField = intMapField;
	}

	public Map<String, String> getStringMapField() {
		return stringMapField;
	}

	public void setStringMapField(Map<String, String> stringMapField) {
		this.stringMapField = stringMapField;
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (!(o instanceof MapFields)) return false;
		MapFields mapFields = (MapFields) o;
		return Objects.equals(getIntMapField(), mapFields.getIntMapField()) && Objects.equals(getStringMapField(), mapFields.getStringMapField());
	}

	@Override
	public int hashCode() {
		return Objects.hash(getIntMapField(), getStringMapField());
	}

	@Override
	public String toString() {
		return String.format("MapFields{intMapField=%s, stringMapField=%s}", intMapField, stringMapField);
	}
}
