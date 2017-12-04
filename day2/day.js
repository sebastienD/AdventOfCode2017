var fs = require('fs')
var os = require('os');

fs.readFile('input.txt', 'utf8', function (err,data) {
  if (err) {
    return console.log(err);
  }
  let lines = data.split(os.EOL);
  lines.pop();
  const sum = lines
    .map(value => value.split('\t').map(x => parseInt(x)))
    .map(nums => divide(nums))
    .reduce((a,c) => a+c);
    console.log(sum);
});

function divide(nums) {
    for (i = 0; i<nums.length; i++) {
        for (j = i+1; j<nums.length; j++) {
            if (Math.max(nums[i], nums[j]) % Math.min(nums[i], nums[j]) == 0) {
                return Math.max(nums[i], nums[j]) / Math.min(nums[i], nums[j])
            }
        }
    }
}

function first() {
    fs.readFile('input.txt', 'utf8', function (err,data) {
        if (err) {
          return console.log(err);
        }
        let lines = data.split(os.EOL);
        lines.pop();
        const sum = lines
          .map(value => value.split('\t').map(x => parseInt(x)))
          .map(nums => nums.reduce((acc, current) => current > acc ? current : acc) - nums.reduce((acc, current) => current > acc ? acc : current))
          .reduce((a,c) => a+c);
          console.log(sum);
      });
}
