import * as t from '../test-service/superstruct'

export const checkEncodeDecode = <T>(theType: t.Struct<T, unknown>, decoded: T, encoded: unknown) => {
  it('encode', function() {
    expect(t.encode(theType, decoded)).toStrictEqual(encoded);
  })
  it('decode', function() {
    expect(t.decode(theType, encoded)).toStrictEqual(decoded);
  })
}