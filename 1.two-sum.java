import java.util.Map;

/*
 * @lc app=leetcode id=1 lang=java
 *
 * [1] Two Sum
 */
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> m = new HashMap<>();
        for (int i = 0; i < nums.length; i++)
        {
            int ans = target - nums[i];
            if (m.containsKey(ans)) return new int[] {i, m.get(ans)};
            m.put(nums[i], i);
        }
        return null;
    }
}

