# 两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

```js
// nums = [2, 7, 11, 15], target = 9
// result [0, 1]
var twoSum = function(nums, target) {
    if(!nums||!nums.length || typeof target!=='number') return [];
    let result=[];
    let numLen=nums.length
    for(let i=0;i<=numLen-2;++i){
        let firstNum=nums[i];
        for(let j=i+1;j<=numLen-1;++j){
            let secondNum=nums[j];
            if(secondNum+firstNum===target){
                result.push(i);
                result.push(j);
                break;
            }
        }
        if(result.length){
            break;
        }
        
    }
    return result;
};
```
# 整数反转

给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

输入: 123，输出: 321

输入: -123，输出: -321

输入: 120，输出: 21

```js
var reverse = function(x) {
    let belowZero=x<0;
    let xStr=x.toString();
    if(xStr.length===1) return x;
    if(xStr.length===2 && belowZero) return x;
    let startNum=belowZero?1:0;
    let numArr=xStr.split('');
    for(let startInx=startNum,endInx=xStr.length-1;startInx<endInx;startInx++,endInx--){
        let temp=numArr[startInx];
        numArr[startInx]=numArr[endInx];
        numArr[endInx]=temp;
    }
    let resultNum= Number.parseInt(numArr.join(''));
    if(resultNum>(Math.pow(2,31)-1)|| resultNum<(0-Math.pow(2,31))){
        return 0;
    }
    return resultNum;
};
```

## 回文子串

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

输入: 121
输出: true 

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

```js
var isPalindrome = function(x) {
    if(typeof x!=='number') return false;
    if(x<0) return false;
    let xStr=x.toString && x.toString()||'';
    let strLen=xStr.length;
    let result=true;
    for(let i=0,j=strLen-1;i<j;++i,--j){
        let leftChar=xStr.charAt(i);
        let rightChar=xStr.charAt(j);
        if(leftChar!==rightChar){
            result=false;
            break;
        }
    }
    return result;
};
```
# 罗马数字转整数

罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

```js
var romanToInt = function(s) {
    if(!s || !s.length) return;
    let specialNumMap={
        'IV':4,
        'IX':9,
        'XL':40,
        'XC':90,
        'CD':400,
        'CM':900
    }
    let numMap={
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000
    }
    let sum=0;
    for(let i=0;i<=s.length-1;++i){
        let currChar=s.charAt(i);
        if(i+1<s.length){
            let nextChar=s.charAt(i+1)   
            let specialStr=currChar+nextChar;
            if(specialNumMap[specialStr]){
                sum+=specialNumMap[specialStr];
                ++i;
            }else{
               sum+=    numMap[currChar];
            }
            continue;
        }
            sum+=numMap[currChar];
    }
    return sum;
};
```