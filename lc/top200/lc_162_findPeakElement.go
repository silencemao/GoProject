package main

import "fmt"

/*
寻找峰值元素，峰值元素满足 nums[i-1] < nums[i] > nums[i+1]
默认
1、数组内nums[i] != nums[i+1]
2、nums[-1]=nums[n]=-∞

因为条件2只要数组有一个元素，就一定存在峰值元素
二分查找：
https://leetcode-cn.com/problems/find-peak-element/solution/xun-zhao-feng-zhi-by-leetcode/
*/
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1          // l不会取到r 所以mid < r-1 所以mid+1<r 不会超出阈值
		if nums[mid] > nums[mid+1] { // 峰值在[l, mid]内
			r = mid
		} else { // 峰值在[mid+1, r]内
			l = mid + 1
		}
	}
	return l
}

/*

class Solution {
    public int findPeakElement(int[] nums) {
        //本题从-inf到-inf, 相邻元素必不相同
        //不可能都是-inf,只要有一个元素就一定有上坡,就一定有峰顶
        // 使用二分法找峰顶
        if(nums.length==1) return 0;
        int l=0,r=nums.length-1;
        while(l<r){// [l,r] 型二分法,此处判别是开还是闭看l和r的取值
            // 虽然是闭区间上查,但用<号,所以最后退出循环时[l,l]或者[r,r]是不会查的
            int mid=(l+r)/2; // 向下取整,所以mid+1不会溢出
            if(nums[mid]>nums[mid+1]){// mid比右侧大, 峰顶在左侧或就在mid处
                r=mid;// [l,mid]
            }else if(nums[mid]<nums[mid+1]){
                l=mid+1;// mid比右侧小,峰顶在右侧[mid+1,r]
            }
        }// 退出循环时 l==r

        // 在l==r时,其实是没有判断当前是否就是答案, 但本题一定会有答案
        // ==>所以就没有去判断了
        return l;
    }
}
*/

func main() {
	nums := []int{1}
	fmt.Println(findPeakElement(nums))
}
