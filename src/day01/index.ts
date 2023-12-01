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
    let re = /(?=(\d|one|two|three|four|five|six|seven|eight|nine))/g;
    const numbers = Array.from(line.matchAll(re), (x) => x[1]);

    const numberStrings = [
      "one",
      "two",
      "three",
      "four",
      "five",
      "six",
      "seven",
      "eight",
      "nine",
    ];
    const numbersClean = numbers.map((number) => {
      let num = number;
      for (const numberString of numberStrings) {
        num = num.replace(
          numberString,
          (numberStrings.indexOf(numberString) + 1).toString(),
        );
      }
      return num;
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
        input: `1abc2
        pqr3stu8vwx
      a1b2c3d49e5f
      treb7uchet`,
        expected: 142,
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
        input: `two1nine
        eightwothree
        abcone2threexyz
        xtwone3four
        4nineeightseven2
        zoneight234
        7pqrstsixteen
        sevenine`,
        expected: 360,
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
