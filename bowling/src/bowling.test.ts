import { Game } from "./bowling";
import assert = require("assert");

var g: Game;

beforeEach(() => {
  g = new Game();
});

function rollMany(n: number, pins: number): void {
  for (let i = 0; i < n; i++) {
    g.roll(pins);
  }
}

function rollSpare(): void {
  rollMany(2, 5);
}

test("canRoll", () => {
  g.roll(0);
});

test("gutterGame", () => {
  rollMany(20, 0);
  assert.equal(g.score(), 0);
});

test("allOnes", () => {
  rollMany(20, 1);
  assert.equal(g.score(), 20);
});

test("oneSpare", () => {
  rollSpare();
  g.roll(7);
  rollMany(17, 0);
  assert.equal(g.score(), 24);
});

test("oneStrike", () => {
  g.roll(10);
  g.roll(2);
  g.roll(3);
  rollMany(16, 0);
  assert.equal(g.score(), 20);
});

test("perfectGame", () => {
  rollMany(12, 10);
  assert.equal(g.score(), 300);
});
