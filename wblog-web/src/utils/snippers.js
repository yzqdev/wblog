function getMost(str) {
  // 步骤1
  var result = {};
  for (let i in str) {
    if (str[i] in result) {
      // 步骤2
      result[str[i]]++;
    } else {
      // 步骤3
      var object = {};
      object[str[i]] = 1;
      result = Object.assign(result, object);
    }
  }
  return result;
}

var result = getMost("xyzzyxyz");
console.log(result); //{x: 2, y: 3, z: 3}
