package javascript

// http://www.java2s.com/Tutorials/Javascript/Javascript_Data_Structure/0240__Javascript_Array_Aggregate.htm

const DeepCopy = `
function copy(arr1, arr2) { 
    for (var i = 0; i < arr1.length; ++i) { 
        arr2[i] = arr1[i]; 
    } 
} 

var nums = []; 
for (var i = 0; i < 100; ++i) { 
   nums[i] = i+1; 
} 

var samenums = []; 
copy(nums, samenums); 
nums[0] = 400; 
console.log(samenums[0]); 
`
