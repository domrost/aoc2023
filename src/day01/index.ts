import run from "aocrunner";

const parseInput = (rawInput: string) => rawInput;

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  var lines = input.split(/\r?\n/);

  let sum = 0;

  for (const line of lines) {
    const numbers = line.match(/\d/g);
    const firstNumber = numbers[0];
    const lastNumber = numbers[numbers.length - 1];
    sum += parseInt(firstNumber + lastNumber, 10);
  }

  return sum;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  var lines = input.split(/\r?\n/);

  let sum: number = 0;

  for (const line of lines) {
    let rx = /(?=(\d|one|two|three|four|five|six|seven|eight|nine))/g;
    const numbers = Array.from(line.matchAll(rx), (x) => x[1]);

    const numbersClean = numbers.map((number) => {
      switch (number) {
        case "one":
          return "1";
        case "two":
          return "2";
        case "three":
          return "3";
        case "four":
          return "4";
        case "five":
          return "5";
        case "six":
          return "6";
        case "seven":
          return "7";
        case "eight":
          return "8";
        case "nine":
          return "9";
        default:
          return number;
      }
    });

    const firstNumber = numbersClean[0];
    const lastNumber = numbersClean[numbersClean.length - 1];
    sum += parseInt(firstNumber + lastNumber, 10);
  }

  return sum;
};

run({
  part1: {
    tests: [
      {
        input: `
          1abc2
          pqr3stu8vwx
          a1b2c3d49e5f
          treb7uchet
        `,
        expected: 142,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `
          two1nine
          eightwothree
          abcone2threexyz
          xtwone3four
          4nineeightseven2
          zoneight234
          7pqrstsixteen
          sevenine
        `,
        expected: 360,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
