export class Game {
  private rolls: number[];
  private currentRoll: number;

  constructor() {
    this.rolls = [];
    this.currentRoll = 0;
  }

  public roll(pins: number): void {
    this.rolls[this.currentRoll++] = pins;
  }

  public score(): number {
    let score = 0;
    let frameIndex = 0;

    for (let frame = 0; frame < 10; frame++) {
      if (this.isStrike(frameIndex)) {
        score += 10 + this.strikeBonus(frameIndex);
        frameIndex++;
      } else if (this.isSpare(frameIndex)) {
        // spare
        score += 10 + this.spareBonus(frameIndex);
        frameIndex += 2;
      } else {
        score += this.twoBallsInFrame(frameIndex);
        frameIndex += 2;
      }
    }
    return score;
  }

  private twoBallsInFrame(frameIndex: number) {
    return this.rolls[frameIndex] + this.rolls[frameIndex + 1];
  }

  private spareBonus(frameIndex: number) {
    return this.rolls[frameIndex + 2];
  }

  private strikeBonus(frameIndex: number) {
    return this.rolls[frameIndex + 1] + this.rolls[frameIndex + 2];
  }

  private isStrike(frameIndex: number) {
    return this.rolls[frameIndex] === 10;
  }

  private isSpare(frameIndex: number): boolean {
    return this.rolls[frameIndex] + this.rolls[frameIndex + 1] === 10;
  }
}
