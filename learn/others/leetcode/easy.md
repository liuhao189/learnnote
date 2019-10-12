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