import { Bowling } from './bowling';
import assert = require('assert');

test('nothing', () => {
    assert(new Bowling().ping() === "pong")
});
