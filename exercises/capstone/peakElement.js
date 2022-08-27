function findPeakElement(nums)  {
  let left = 0;
  let right = nums.length - 1;
  
  if (nums.length === 1) return 0
  
  while (left <= right) {
    let mid = parseInt(left + ((right-left)/2), 10)
    
    if ((nums[mid] > nums[mid-1]) && (nums[mid] > nums[mid+1]))  {
      return mid
    } else if (nums[mid] < nums[mid+1])  {
      left = mid + 1
    } else {
      right = mid - 1
    }
    
    if (mid === left) return left
  }
}
