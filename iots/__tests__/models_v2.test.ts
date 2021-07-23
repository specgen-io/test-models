import * as t from '../io-ts'
import { checkEncodeDecode } from '../util'

import {
  Message, TMessage
} from '../models_v2';

describe('v2 object', function() {
  let decoded: Message = {field: 'the string'}
  let encoded = {'field': 'the string'}
  checkEncodeDecode(TMessage, decoded, encoded)
});