#! /usr/bin/env node
const { readFileSync } = require('fs')
function main() {
  const data = readFileSync(__dirname + '/input.txt', { encoding: 'utf-8' })
  console.log('1-' + partOne(data))
  console.log('2-' + partTwo(data))
}

function partTwo(data) {
  return data.split('\n').filter(d => d !== '').reduce((p, c, i, a) => {
    const { elves, elfIndex } = p
    if (c === '\r') {
      return {
        elves,
        elfIndex: elfIndex + 1
      }
    }

    if (elves[elfIndex] === undefined) {
      elves[elfIndex] = parseInt(c)
    } else {
      elves[elfIndex] = elves[elfIndex] + parseInt(c)
    }
    return {
      elves,
      elfIndex
    }
  }, { elves: [], elfIndex: 0 }).elves.sort((a, b) => b - a).slice(0, 3).reduce((a, b) => a + b)
}

main()

function partOne(data) {
  const result = data.split('\n').filter(d => d !== '').reduce((p, c, i, a) => {
    const { currentElf, max } = p

    if (c === '\r') {
      // new elf time, calculate current max
      if (currentElf > max) {

        return {
          max: currentElf,
          currentElf: 0
        }
      } else {
        return {
          max: max,
          currentElf: 0
        }
      }
    } else {
      return {
        max: max,
        currentElf: currentElf + parseInt(c)
      }
    }
  }, { max: -Infinity, currentElf: 0 })
  return result.max
}
