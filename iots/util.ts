import * as t from './io-ts'

export const checkEncodeDecode = <A, O, I>(theType: t.Type<A, O, I>, decoded: A, encoded: I) => {
    it('encode', function() {
    expect(t.encode(theType, decoded)).toStrictEqual(encoded);
  })
  it('decode', function() {
    expect(t.decode(theType, encoded)).toStrictEqual(decoded);
  })
}