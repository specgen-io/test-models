package test_service.models;

import org.junit.Test;

import static org.junit.Assert.assertThrows;

public class NullabilityTest {
    @Test
    public void oneOfItemNotNull() {
        assertThrows(IllegalArgumentException.class, () -> new OrderEventCanceled(null));
    }
}
