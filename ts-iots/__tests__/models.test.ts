import * as t from '../io-ts'
import { checkEncodeDecode } from '../util'

import {
  TMessage,
  Message,
  Choice,
  EnumFields,
  TEnumFields,
  Parent,
  TParent,
  Nested,
  NumericFields,
  TNumericFields,
  NonNumericFields,
  TNonNumericFields,
  ArrayFields,
  TArrayFields,
  MapFields,
  TMapFields,
  OptionalFields,
  TOptionalFields,
  OrderEvent,
  TOrderEvent
} from '../models';

describe('object', function() {
  let decoded: Message = {field: 123}
  let encoded = {'field': 123}
  checkEncodeDecode(TMessage, decoded, encoded)
  it('decode breaks', function() {
    expect(() => t.decode(TMessage, {})).toThrow();
  })
});

describe('enum fields', function() {
  let decoded: EnumFields = {enum_field: Choice.THIRD_CHOICE}
  let encoded = {'enum_field': 'THIRD_CHOICE'}
  checkEncodeDecode(TEnumFields, decoded, encoded)
  it('decode fail', function() {
    expect(() => t.decode(TEnumFields, {})).toThrowError('Decoding failed')
  })
});

describe('nested types', function() {
  let decoded: Parent = {
    'field': 'the string',
    'nested': {'field': "the nested string"}
  }
  let encoded = {
    'field': 'the string',
    'nested': {'field': "the nested string"}
  }
  checkEncodeDecode(TParent, decoded, encoded)
  it('decode fail', function() {
    expect(() => t.decode(TParent, {})).toThrowError('Decoding failed');
  })
});

describe('numeric fields', function() {
  let decoded: NumericFields = {
    int_field: 123,
    long_field: 123,
    float_field: 12.3,
    double_field: 12.3,
    decimal_field: 12.3,
  }
  let encoded = {
    int_field: 123,
    long_field: 123,
    float_field: 12.3,
    double_field: 12.3,
    decimal_field: 12.3,
  }
  checkEncodeDecode(TNumericFields, decoded, encoded)
  it('decode fail', function() {
    expect(() => t.decode(TNumericFields, {})).toThrowError('Decoding failed');

  })
});

describe('non numeric fields', function() {
  let decoded: NonNumericFields = {
    boolean_field: true,
    string_field: 'some string',
    uuid_field: '123e4567-e89b-12d3-a456-426655440000',
    date_field: '2021-01-01',
    datetime_field: new Date('2021-01-02T04:54:00.000Z'),
  }
  let encoded = {
    boolean_field: true,
    string_field: 'some string',
    uuid_field: '123e4567-e89b-12d3-a456-426655440000',
    date_field: '2021-01-01',
    datetime_field: '2021-01-02T04:54:00.000Z',
  }
  checkEncodeDecode(TNonNumericFields, decoded, encoded)
});

describe('array fields', function() {
  let decoded: ArrayFields = {
    int_array_field: [1, 2, 3],
    string_array_field: ['one', 'two', 'three'],
  }
  let encoded = {
    int_array_field: [1, 2, 3],
    string_array_field: ['one', 'two', 'three'],
  }
  checkEncodeDecode(TArrayFields, decoded, encoded)
});

describe('map fields', function() {
  let decoded: MapFields = {
    int_map_field: {'one': 1, 'two': 2, 'three': 3},
    string_map_field: {'one': 'first', 'two': 'second', 'three': 'third'},
  }
  let encoded = {
    int_map_field: {'one': 1, 'two': 2, 'three': 3},
    string_map_field: {'one': 'first', 'two': 'second', 'three': 'third'},
  }
  checkEncodeDecode(TMapFields, decoded, encoded)
});

describe('optional fields', function() {
  let decoded: OptionalFields = {
    int_option_field: 123,
    string_option_field: 'the string',
  }
  let encoded = {
    int_option_field: 123,
    string_option_field: 'the string',
  }
  checkEncodeDecode(TOptionalFields, decoded, encoded)
});

describe('optional fields null', function() {
  let decoded: OptionalFields = {
    int_option_field: null,
    string_option_field: null,
  }
  let encoded = {
    int_option_field: null,
    string_option_field: null,
  }
  checkEncodeDecode(TOptionalFields, decoded, encoded)
});

describe('optional fields missing', function() {
  let decoded: OptionalFields = {}
  let encoded = {}
  checkEncodeDecode(TOptionalFields, decoded, encoded)
});

describe('oneof types', function() {
  let decoded: OrderEvent = { changed: { id: '123e4567-e89b-12d3-a456-426655440000', quantity: 123 } }
  let encoded = { changed: {id: '123e4567-e89b-12d3-a456-426655440000', quantity: 123} }
  checkEncodeDecode(TOrderEvent, decoded, encoded)
  it('decode breaks', function() {
    expect(() => t.decode(TOrderEvent, { created: {} })).toThrowError('Decoding failed');
  })
});