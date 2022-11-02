// ボウリングの 1 ゲーム分の投球記録が与えられたとき、そのゲームの合計スコアを生成するプログラム

export class Game {
  private frames: Frame[];
  private rolls: number[];
  private frameIndex: number;

  constructor() {
    this.frameIndex = 0;
    this.frames = new Array(10);
    for (let i = 0; i < this.frames.length; i++) {
      this.frames[i] = new Frame(i);
    }
    this.rolls = [];
  }

  public roll(pins: number) {
    this.rolls.push(pins);

    if (pins === 10) {
      this.frameIndex++;
    } else if (this.rolls.length >= 2) {
      this.frameIndex++;
    }
  }

  public getCurrentFrameIndex(): number {
    return this.frameIndex;
  }
}

class Frame {
  private frameIndex: number;
  private rolls: number[];

  constructor(n: number) {
    this.frameIndex = n;
    this.rolls = new Array(2);
  }
}
