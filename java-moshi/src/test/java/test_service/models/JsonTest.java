package test_service.models;

import com.squareup.moshi.*;
import org.junit.jupiter.api.Test;
import test_service.json.adapters.*;

import java.io.*;
import java.math.BigDecimal;
import java.time.*;
import java.util.*;

import static org.junit.jupiter.api.Assertions.*;
import static test_service.json.Json.setupMoshiAdapters;
import static utils.Utils.*;

public class JsonTest {
	public static Moshi createMoshiObject() {
		Moshi.Builder moshiBuilder = new Moshi.Builder();
		setupMoshiAdapters(moshiBuilder);

		moshiBuilder
			.add(new UnwrapFieldAdapterFactory(OrderEventDiscriminator.Created.class))
			.add(new UnwrapFieldAdapterFactory(OrderEventDiscriminator.Changed.class))
			.add(new UnwrapFieldAdapterFactory(OrderEventDiscriminator.Canceled.class))
			.add(UnionAdapterFactory.of(OrderEventDiscriminator.class).withDiscriminator("_type")
				.withSubtype(OrderEventDiscriminator.Created.class, "created")
				.withSubtype(OrderEventDiscriminator.Changed.class, "changed")
				.withSubtype(OrderEventDiscriminator.Canceled.class, "canceled"))
			.add(new UnwrapFieldAdapterFactory(OrderEventWrapper.Created.class))
			.add(new UnwrapFieldAdapterFactory(OrderEventWrapper.Changed.class))
			.add(new UnwrapFieldAdapterFactory(OrderEventWrapper.Canceled.class))
			.add(UnionAdapterFactory.of(OrderEventWrapper.class)
				.withSubtype(OrderEventWrapper.Created.class, "created")
				.withSubtype(OrderEventWrapper.Changed.class, "changed")
				.withSubtype(OrderEventWrapper.Canceled.class, "canceled"));

		return moshiBuilder.build();
	}

	public <T> void check(T data, String jsonStr, Class<T> tClass) throws IOException {
		Moshi moshi = createMoshiObject();
		JsonAdapter<T> jsonAdapter = moshi.adapter(tClass);

		String actualJson = jsonAdapter.toJson(data);
		jsonStr = fixQuotes(jsonStr);
		assertEquals(jsonStr, actualJson);

		T actualData = jsonAdapter.fromJson(jsonStr);
		assertEquals(actualData, data);
	}

	@Test
	public void objectModel() throws IOException {
		Message data = new Message(123);
		String jsonStr = "{'field':123}";
		check(data, jsonStr, Message.class);
	}

	@Test
	public void objectModelFieldCases() throws IOException {
		MessageCases data = new MessageCases("snake_case value", "camelCase value");
		String jsonStr = "{'camelCase':'camelCase value','snake_case':'snake_case value'}";
		check(data, jsonStr, MessageCases.class);
	}

	@Test
	public void nestedObject() throws IOException {
		Parent data = new Parent("the string", new Message(123));
		String jsonStr = "{'field':'the string','nested':{'field':123}}";
		check(data, jsonStr, Parent.class);
	}

	@Test
	public void enumModel() throws IOException {
		EnumFields data = new EnumFields(Choice.SECOND_CHOICE);
		String jsonStr = "{'enum_field':'Two'}";
		check(data, jsonStr, EnumFields.class);
	}

	@Test
	public void numericTypes() throws IOException {
		NumericFields data = new NumericFields(123, 1234, 1.23f, 1.23, new BigDecimal("1.23"));
		String jsonStr = "{'decimal_field':1.23,'double_field':1.23,'float_field':1.23,'int_field':123,'long_field':1234}";
		check(data, jsonStr, NumericFields.class);
	}

	@Test
	public void nonNumericTypes() throws IOException {
		NonNumericFields data = new NonNumericFields(true, "the string", UUID.fromString("123e4567-e89b-12d3-a456-426655440000"), LocalDate.parse("2019-11-30"), LocalDateTime.parse("2019-11-30T17:45:55"));
		String jsonStr = "{'boolean_field':true,'date_field':'2019-11-30','datetime_field':'2019-11-30T17:45:55','string_field':'the string','uuid_field':'123e4567-e89b-12d3-a456-426655440000'}";
		check(data, jsonStr, NonNumericFields.class);
	}

	@Test
	public void arrayType() throws IOException {
		ArrayFields data = new ArrayFields(new int[]{1, 2, 3}, new String[]{"one", "two", "three"});
		String jsonStr = "{'int_array_field':[1,2,3],'string_array_field':['one','two','three']}";
		check(data, jsonStr, ArrayFields.class);
	}

	@Test
	public void mapType() throws IOException {
		MapFields data = new MapFields(
			new HashMap<>() {{ put("one", 1); put("two", 2); }},
			new HashMap<>() {{ put("one", "first"); put("two", "second"); }}
		);
		String jsonStr = "{'int_map_field':{'one':1,'two':2},'string_map_field':{'one':'first','two':'second'}}";
		check(data, jsonStr, MapFields.class);
	}

	@Test
	public void optionalTypes() throws IOException {
		OptionalFields data = new OptionalFields(123, "the string");
		String jsonStr = "{'int_option_field':123,'string_option_field':'the string'}";
		check(data, jsonStr, OptionalFields.class);
	}

	@Test
	public void optionalTypesMissingFields() throws IOException {
		var data = new OptionalFields(null, null);
		String jsonStr = "{}";
		check(data, jsonStr, OptionalFields.class);
	}

	@Test
	public void jsonType() throws IOException {
		//String jsonField = fixQuotes("{'the_array':[true,'some string'],'the_object':{'the_bool':true,'the_string':'some value'},'the_scalar':'the value'}");
		var theObject = new HashMap<String, Object>();
		theObject.put("the_bool", true);
		theObject.put("the_string", "some value");
		var theArray = new ArrayList<Object>();
		theArray.add(true);
		theArray.add("some string");
		var map = new HashMap<String, Object>();
		map.put("the_array", theArray);
		map.put("the_object", theObject);
		map.put("the_scalar", "the value");
		RawJsonField data = new RawJsonField(map);
		String jsonStr = "{'json_field':{'the_array':[true,'some string'],'the_scalar':'the value','the_object':{'the_bool':true,'the_string':'some value'}}}";
		check(data, jsonStr, RawJsonField.class);
	}

	@Test
	public void oneOfItemType() throws IOException {
		OrderCreated data = new OrderCreated(UUID.fromString("58d5e212-165b-4ca0-909b-c86b9cee0111"), "SNI/01/136/0500", 3);
		String jsonStr = "{'id':'58d5e212-165b-4ca0-909b-c86b9cee0111','quantity':3,'sku':'SNI/01/136/0500'}";
		check(data, jsonStr, OrderCreated.class);
	}

	@Test
	public void oneOfWrapper() throws IOException {
		OrderEventWrapper data = new OrderEventWrapper.Canceled(new OrderCanceled(UUID.fromString("123e4567-e89b-12d3-a456-426655440000")));
		String jsonStr = "{'canceled':{'id':'123e4567-e89b-12d3-a456-426655440000'}}";
		check(data, jsonStr, OrderEventWrapper.class);
	}

	@Test
	public void oneOfDiscriminator() throws IOException {
		OrderEventDiscriminator data = new OrderEventDiscriminator.Canceled(new OrderCanceled(UUID.fromString("123e4567-e89b-12d3-a456-426655440000")));
		String jsonStr = "{'_type':'canceled','id':'123e4567-e89b-12d3-a456-426655440000'}";
		check(data, jsonStr, OrderEventDiscriminator.class);
	}
}
