package test_service.models;

import test_service.models.spec.Message;
import test_service.models.spec.Nested;
import test_service.models.spec.Parent;
import test_service.models.spec.EnumFields;
import test_service.models.spec.Choice;
import test_service.models.spec.NumericFields;
import test_service.models.spec.NonNumericFields;
import test_service.models.spec.ArrayFields;
import test_service.models.spec.MapFields;
import test_service.models.spec.OptionalFields;
import test_service.models.spec.RawJsonField;
import test_service.models.spec.OrderCreated;
import test_service.models.spec.OrderEvent;
import test_service.models.spec.OrderEventCanceled;
import test_service.models.spec.OrderCanceled;

import com.fasterxml.jackson.databind.*;
import org.junit.Test;

import java.io.*;
import java.math.BigDecimal;
import java.time.*;
import java.util.*;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.Assert.*;
import static test_service.models.Jsoner.*;

public class ModelsTest {

	public static String fixQuotes(String jsonStr) {
		return jsonStr.replaceAll("'", "\"");
	}

	public <T> void check(T data, String jsonStr, Class<T> tClass) throws IOException {
		ObjectMapper objectMapper = new ObjectMapper();

		String actualJson = serialize(objectMapper, data);
		jsonStr = fixQuotes(jsonStr);
		assertEquals(jsonStr, actualJson);

		T actualData = deserialize(objectMapper, jsonStr, tClass);
		assertThat(actualData).usingRecursiveComparison().isEqualTo(data);
	}

	@Test
	public void jsonMessageTest() throws IOException {
		Message data = new Message(123);

		String jsonStr = "{'field':123}";

		check(data, jsonStr, Message.class);
	}

	@Test
	public void jsonNestedTest() throws IOException {
		Nested data = new Nested("the string");

		String jsonStr = "{'field':'the string'}";

		check(data, jsonStr, Nested.class);
	}

	@Test
	public void jsonParentTest() throws IOException {
		Parent data = new Parent("the string", new Nested("the nested string"));

		String jsonStr = "{'field':'the string','nested':{'field':'the nested string'}}";

		check(data, jsonStr, Parent.class);
	}

	@Test
	public void jsonEnumTest() throws IOException {
		EnumFields data = new EnumFields(Choice.SECOND_CHOICE);

		String jsonStr = "{'enum_field':'SECOND_CHOICE'}";

		check(data, jsonStr, EnumFields.class);
	}

	@Test
	public void jsonNumericFieldsTest() throws IOException {
		NumericFields data = new NumericFields(123, 1234, 1.23f, 1.23, new BigDecimal("1.23"));

		String jsonStr = "{'int_field':123,'long_field':1234,'float_field':1.23,'double_field':1.23,'decimal_field':1.23}";

		check(data, jsonStr, NumericFields.class);
	}

	@Test
	public void jsonNonNumericFieldsTest() throws IOException {
		NonNumericFields data = new NonNumericFields(true, "the string", UUID.fromString("123e4567-e89b-12d3-a456-426655440000"), LocalDate.parse("2019-11-30"), LocalDateTime.parse("2019-11-30T17:45:55"));

		String jsonStr = "{'boolean_field':true,'string_field':'the string','uuid_field':'123e4567-e89b-12d3-a456-426655440000','date_field':'2019-11-30','datetime_field':'2019-11-30T17:45:55'}";

		check(data, jsonStr, NonNumericFields.class);
	}

	@Test
	public void jsonArrayFieldsTest() throws IOException {
		ArrayFields data = new ArrayFields(new int[]{1, 2, 3}, new String[]{"one", "two", "three"});

		String jsonStr = "{'int_array_field':[1,2,3],'string_array_field':['one','two','three']}";

		check(data, jsonStr, ArrayFields.class);
	}

	@Test
	public void jsonMapFieldsTest() throws IOException {
		MapFields data = new MapFields(new HashMap<String, Integer>() {{
			put("one", 1);
			put("two", 2);
		}}, new HashMap<String, String>() {{
			put("one", "first");
			put("two", "second");
		}});

		String jsonStr = "{'int_map_field':{'one':1,'two':2},'string_map_field':{'one':'first','two':'second'}}";

		check(data, jsonStr, MapFields.class);
	}

	@Test
	public void jsonOptionalFieldsTest() throws IOException {
		OptionalFields data = new OptionalFields(123, "the string");

		String jsonStr = "{'int_option_field':123,'string_option_field':'the string'}";

		check(data, jsonStr, OptionalFields.class);
	}

	@Test
	public void jsonRawJsonFieldTest() throws IOException {
		String jsonField = fixQuotes("{'the_array':[1,'some string'],'the_object':{'the_bool':true,'the_string':'some value'},'the_scalar':123}");
		RawJsonField data = new RawJsonField(jsonField);

		String jsonStr = "{'json_field':{'the_array':[1,'some string'],'the_object':{'the_bool':true,'the_string':'some value'},'the_scalar':123}}";

		check(data, jsonStr, RawJsonField.class);
	}

	@Test
	public void jsonOrderCreatedTest() throws IOException {
		OrderCreated data = new OrderCreated(UUID.fromString("58d5e212-165b-4ca0-909b-c86b9cee0111"), "SNI/01/136/0500", 3);

		String jsonStr = "{'id':'58d5e212-165b-4ca0-909b-c86b9cee0111','sku':'SNI/01/136/0500','quantity':3}";

		check(data, jsonStr, OrderCreated.class);
	}

	@Test
	public void jsonUnionTest() throws IOException {
		OrderEvent data = new OrderEventCanceled(new OrderCanceled(UUID.fromString("123e4567-e89b-12d3-a456-426655440000")));

		String jsonStr = "{'canceled':{'id':'123e4567-e89b-12d3-a456-426655440000'}}";

		check(data, jsonStr, OrderEvent.class);
	}
}
