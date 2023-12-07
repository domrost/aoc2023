import run from "aocrunner";

const parseInput = (rawInput: string) => rawInput;
import { splitLines } from "../utils/index.js";

class Map {
  mapEntries: MapEntry[];

  constructor() {
    this.mapEntries = [];
  }

  getDest(source: number) {
    for (const entry of this.mapEntries) {
      if (
        entry.sourceRangeStart <= source &&
        entry.sourceRangeStart + entry.rangeLength - 1 >= source
      ) {
        return entry.destRangeStart + (source - entry.sourceRangeStart);
      }
    }
    return source;
  }
}

interface MapEntry {
  destRangeStart: number;
  sourceRangeStart: number;
  rangeLength: number;
}

interface SeedRange {
  rangeStart: number;
  rangeEnd: number;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);

  const seeds: number[] = extractSeeds(lines[0]);
  const maps = extractMaps(lines.slice(3));

  const locations: number[] = [];
  for (const seed of seeds) {
    locations.push(getLocation(seed, maps));
  }

  return Math.min(...locations);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const lines = splitLines(input);

  const seedRanges: SeedRange[] = extractSeedRanges(lines[0]);
  const maps = extractMaps(lines.slice(3));

  let minLocation = null;
  let seedAmount = 0;

  for (const seedRange of seedRanges) {
    seedAmount += seedRange.rangeEnd - seedRange.rangeStart;
  }

  let finishedSeeds = 0;
  for (const seedRange of seedRanges) {
    for (let seed = seedRange.rangeStart; seed <= seedRange.rangeEnd; seed++) {
      const location = getLocation(seed, maps);
      if (minLocation === null || location < minLocation) {
        minLocation = location;
        continue;
      }
      finishedSeeds++;
    }
    console.log(`Finished ${finishedSeeds / seedAmount} of seeds.`);
  }

  return minLocation;
};

function extractSeeds(line: string): number[] {
  return line
    .split(/\s+/)
    .slice(1)
    .map((e) => parseInt(e));
}

function extractSeedRanges(line: string): SeedRange[] {
  const seeds: SeedRange[] = [];
  const numbers = line
    .split(/\s+/)
    .slice(1)
    .map((e) => parseInt(e));
  for (let i = 0; i < numbers.length / 2; i++) {
    const range: SeedRange = {
      rangeStart: numbers[2 * i],
      rangeEnd: numbers[2 * i] + numbers[2 * i + 1] - 1,
    };
    seeds.push(range);
  }
  return seeds;
}

function extractMaps(lines: string[]) {
  const maps: Map[] = [];
  let currentMap = new Map();

  for (const line of lines) {
    if (line.match(/[a-z]/)) {
      //start of new map
      maps.push(currentMap);
      currentMap = new Map();
      continue;
    }

    if (line.match(/[0-9]/)) {
      const mapEntry: MapEntry = {
        destRangeStart: 0,
        sourceRangeStart: 0,
        rangeLength: 0,
      };
      [
        mapEntry.destRangeStart,
        mapEntry.sourceRangeStart,
        mapEntry.rangeLength,
      ] = line.split(/\s/).map((e) => parseInt(e));
      currentMap.mapEntries.push(mapEntry);
    }
  }
  maps.push(currentMap);
  return maps;
}

function getLocation(seed: number, maps: Map[]) {
  let currentNum = seed;

  for (const map of maps) {
    currentNum = map.getDest(currentNum);
  }

  return currentNum;
}

run({
  part1: {
    tests: [
      {
        input: `
        seeds: 79 14 55 13

        seed-to-soil map:
        50 98 2
        52 50 48

        soil-to-fertilizer map:
        0 15 37
        37 52 2
        39 0 15

        fertilizer-to-water map:
        49 53 8
        0 11 42
        42 0 7
        57 7 4

        water-to-light map:
        88 18 7
        18 25 70

        light-to-temperature map:
        45 77 23
        81 45 19
        68 64 13

        temperature-to-humidity map:
        0 69 1
        1 0 69

        humidity-to-location map:
        60 56 37
        56 93 4
      `,
        expected: 35,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      //      {
      {
        input: `
        seeds: 79 14 55 13

        seed-to-soil map:
        50 98 2
        52 50 48

        soil-to-fertilizer map:
        0 15 37
        37 52 2
        39 0 15

        fertilizer-to-water map:
        49 53 8
        0 11 42
        42 0 7
        57 7 4

        water-to-light map:
        88 18 7
        18 25 70

        light-to-temperature map:
        45 77 23
        81 45 19
        68 64 13

        temperature-to-humidity map:
        0 69 1
        1 0 69

        humidity-to-location map:
        60 56 37
        56 93 4
      `,
        expected: 46,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
