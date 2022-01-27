package test_service.models;

public class Utils {
	public static String fixQuotes(String jsonStr) {
		return jsonStr.replaceAll("'", "\"");
	}
}