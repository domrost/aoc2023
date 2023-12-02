import run from "aocrunner";
import { splitLines } from "../utils/index.js";

const parseInput = (rawInput: string) => rawInput;

interface Bag {
  red: number;
  green: number;
  blue: number;
}

const bag: Bag = { red: 12, green: 13, blue: 14 };

interface Game {
  id: number;
  draws: Draw[];
}

interface Draw {
  red: number;
  green: number;
  blue: number;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);

  let sum = 0;
  const games = parseGames(lines);

  for (const game of games) {
    let gameValid: boolean = true;
    for (const draw of game.draws) {
      if (
        draw.red > bag.red ||
        draw.green > bag.green ||
        draw.blue > bag.blue
      ) {
        gameValid = false;
        break;
      }
    }
    if (gameValid) sum += game.id;
  }

  return sum;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);
  const games = parseGames(lines);

  let sum = 0;
  for (const game of games) {
    const max: Bag = { red: 0, green: 0, blue: 0 };
    for (const draw of game.draws) {
      if (draw.red > max.red) max.red = draw.red;
      if (draw.green > max.green) max.green = draw.green;
      if (draw.blue > max.blue) max.blue = draw.blue;
    }
    const power = max.red * max.green * max.blue;
    sum += power;
  }

  return sum;
};

function parseGames(lines: string[]): Game[] {
  const games: Game[] = [];
  for (const line of lines) {
    const game: Game = { id: 0, draws: [] };
    game.id = parseInt(line.split(": ")[0].split(" ")[1], 10);

    const drawsStrings = line.split(": ")[1].split("; ");
    game.draws = parseDraws(drawsStrings);

    games.push(game);
  }
  return games;
}

function parseDraws(drawsStrings: string[]): Draw[] {
  const draws: Draw[] = [];
  for (const drawString of drawsStrings) {
    const draw: Draw = { blue: 0, green: 0, red: 0 };

    const cubeStrings = drawString.split(", ");
    for (const cubeString of cubeStrings) {
      draw[cubeString.split(" ")[1]] = parseInt(cubeString.split(" ")[0]);
    }
    draws.push(draw);
  }
  return draws;
}

run({
  part1: {
    tests: [
      {
        input: `
        Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
        `,
        expected: 8,
      },

      // {
      //   input: ``,
      //   expected: "",
      // },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
        `,
        expected: 2286,
      },
      // {
      //   input: ``,
      //   expected: "",
      // },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
