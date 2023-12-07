import run from "aocrunner";
import { splitLines } from "../utils/index.js";

const parseInput = (rawInput: string) => rawInput;

interface Symbol {
  symbol: string;
  line: number;
  index: number;
}

interface Number {
  number: string;
  line: number;
  indexStart: number;
  indexEnd: number;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);

  let sum = 0;

  const symbols = getSymbols(lines);
  const numbers = getNumbers(lines);

  const adNums: Number[] = [];
  symbols.forEach((symbol) => {
    adNums.push(...getAdjecentNumbers(symbol, numbers));
  });

  adNums.forEach((adNum) => {
    sum += parseInt(adNum.number, 10);
  });
  return sum;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const lines = splitLines(input);

  let sum = 0;

  const symbols = getSymbols(lines);
  const numbers = getNumbers(lines);

  symbols.forEach((symbol) => {
    if (symbol.symbol === "*") {
      const adNums = getAdjecentNumbers(symbol, numbers);
      if (adNums.length == 2) {
        sum += parseInt(adNums[0].number) * parseInt(adNums[1].number);
      }
    }
  });

  return sum;
};

function getSymbols(lines: string[]) {
  const symbols: Symbol[] = [];
  lines.forEach((line, lineIndex) => {
    const symMatches = line.matchAll(/[^a-zA-Z0-9\. ]/g);
    for (let symMatch of symMatches) {
      const sym: Symbol = {
        index: symMatch.index,
        line: lineIndex,
        symbol: symMatch[0],
      };
      symbols.push(sym);
    }
  });
  return symbols;
}

function getNumbers(lines: string[]) {
  const numbers: Number[] = [];
  lines.forEach((line, lineIndex) => {
    const numMatches = line.matchAll(/[0-9]+/g);
    for (const numMatch of numMatches) {
      const num: Number = {
        number: numMatch[0],
        line: lineIndex,
        indexStart: numMatch.index,
        indexEnd: numMatch.index + numMatch[0].length - 1,
      };
      numbers.push(num);
    }
  });
  return numbers;
}

function getAdjecentNumbers(symbol: Symbol, numbers: Number[]): Number[] {
  const adNums: Number[] = [];
  numbers.forEach((number) => {
    if (
      number.line === symbol.line ||
      number.line === symbol.line - 1 ||
      number.line === symbol.line + 1
    ) {
      if (
        number.indexEnd === symbol.index - 1 ||
        number.indexEnd === symbol.index ||
        number.indexEnd === symbol.index + 1 ||
        number.indexStart === symbol.index - 1 ||
        number.indexStart === symbol.index ||
        number.indexStart === symbol.index + 1
      ) {
        adNums.push(number);
      }
    }
  });
  return adNums;
}

run({
  part1: {
    tests: [
      {
        input: `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
        `,
        expected: 4361,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..
        `,
        expected: 467835,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
