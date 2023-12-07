import run from "aocrunner";
import { splitLines } from "../utils/index.js";

const parseInput = (rawInput: string) => rawInput;

interface Card {
  id: number;
  winningNums: number[];
  haveNums: number[];
  wins: number;
  copies: number;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);

  const cards = parseLines(lines);

  let sum = 0;
  cards.forEach((card) => {
    const wins = getWins(card);
    // console.log(`Card ${card.id}: ${wins} wins`);
    if (wins > 0) sum += Math.pow(2, wins - 1);
  });

  return sum;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);

  const cards = parseLines(lines);

  let sum = 0;
  cards.forEach((card) => {
    card.wins = getWins(card);

    for (let i = 0; i < card.wins; i++) {
      cards[card.id + i].copies += card.copies;
    }

    // console.log(`Card ${card.id}: ${card.copies}`);

    sum += card.copies;
  });

  return sum;
};

function parseLines(lines: string[]): Card[] {
  const cards: Card[] = [];
  lines.forEach((line) => {
    const card: Card = {
      id: 0,
      winningNums: [],
      haveNums: [],
      wins: 0,
      copies: 1,
    };
    card.id = parseInt(line.match(/Card\s+\d+/)[0].split(/\s+/)[1], 10);

    const numbersString = line.split(": ")[1];
    const [winningNumsString, haveNumsString] = numbersString
      .split(" | ")
      .map((s) => s.trim());

    card.winningNums = winningNumsString
      .split(/\s+/)
      .map((numString) => parseInt(numString, 10));
    card.haveNums = haveNumsString
      .split(/\s+/)
      .map((numString) => parseInt(numString, 10));

    cards.push(card);
  });

  return cards;
}

function getWins(card: Card): number {
  let wins = 0;
  card.haveNums.forEach((haveNum) => {
    if (card.winningNums.includes(haveNum)) {
      wins++;
    }
  });
  return wins;
}

run({
  part1: {
    tests: [
      {
        input: `
        Card   1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
        Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
        Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
        Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
        `,
        expected: 13,
      },
      {
        input: `Card 110:  3 70 67  8 59 13 93 99 52 83 |  2 68 16 39  7 77 75 64 57 47 56 30 73 62 20 82  4 31 28 81  1 19  6 76 32`,
        expected: 0,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
    Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
    `,
        expected: 30,
      },
      // {{}
      //   input: ``,
      //   expected: "",
      // },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
