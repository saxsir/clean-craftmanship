import { Game } from "./bowling";
import assert = require("assert");

var g: Game;

beforeEach(() => {
  g = new Game();
});

test("canRoll", () => {
  g.roll(0);
});

test("2投するとフレームが進む", () => {
  g.roll(0);
  g.roll(0);
  assert.equal(g.getCurrentFrameIndex(), 1);
});

test("ストライクだと1投でもフレームが進む", () => {
  g.roll(10);
  assert.equal(g.getCurrentFrameIndex(), 1);
});

// describe("calcScore", function () {
//   test("2投球", () => {
//     g.roll(0);
//     g.score(1);
//   });
// });
