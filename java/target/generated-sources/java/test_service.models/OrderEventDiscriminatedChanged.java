package test_service.models;

import java.time.*;
import java.util.*;
import java.math.BigDecimal;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.annotation.JsonSubTypes.*;

public class OrderEventDiscriminatedChanged implements OrderEventDiscriminated {
	@JsonUnwrapped
	public OrderChanged data;

	public OrderEventDiscriminatedChanged() {
	}

	public OrderEventDiscriminatedChanged(OrderChanged data) {
		this.data = data;
	}

	public OrderChanged getData() {
		return data;
	}

	public void setData(OrderChanged data) {
		this.data = data;
	}

	@Override
	public boolean equals(Object o) {
		if (this == o) return true;
		if (!(o instanceof OrderEventDiscriminatedChanged)) return false;
		OrderEventDiscriminatedChanged changed = (OrderEventDiscriminatedChanged) o;
		return Objects.equals(getData(), changed.getData());
	}

	@Override
	public int hashCode() {
		return Objects.hash(getData());
	}

	@Override
	public String toString() {
		return String.format("OrderChanged{data=%s}", data);
	}
}
