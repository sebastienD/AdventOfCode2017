const fs = require('fs');

const data = fs
  .readFileSync('input.txt')
  .toString('utf-8')
  .split("\n")
  .filter(e => e !== '')
  .map(f => parseInt(f, 10));

let index = 0;
let counter = 0;

while (typeof(data[index]) !== 'undefined') {
  const oldIndex = index;
  const newIndex = index + data[index];
  /*/
  data[oldIndex]++;
  /*/
  if (data[index] >= 3) {
    data[oldIndex]--;
  } else {
    data[oldIndex]++;
  }
  //*/
  index = newIndex;
  counter++;
}

console.log(counter);