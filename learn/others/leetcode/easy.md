# 两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

https://leetcode-cn.com/problems/two-sum/

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