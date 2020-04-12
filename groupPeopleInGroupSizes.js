/**
 * Problem Statement:
There are n people whose IDs go from 0 to n - 1 and each person belongs exactly to one group. Given the array groupSizes of length n telling the group size each person belongs to, return the groups there are and the people's IDs each group includes.

You can return any solution in any order and the same applies for IDs. Also, it is guaranteed that there exists at least one solution. 

 

Example 1:

Input: groupSizes = [3,3,3,3,3,1,3]
Output: [[5],[0,1,2],[3,4,6]]
Explanation: 
Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
Example 2:

Input: groupSizes = [2,1,3,3,3,2]
Output: [[1],[0,5],[2,3,4]]
 

Constraints:

groupSizes.length == n
1 <= n <= 500
1 <= groupSizes[i] <= n
-----------------------------------------

 * Solution:
 * @param {number[]} groupSizes
 * @return {number[][]}
 * Input is an [ ] of groupsize for each person & each person belongs to 1 group
 * we can loop thru the [ ] adding each input to a map of its group size eg: { groupsize: [idx] }
 * when the size is reached push the current group to the result [ ] & emplty the groupsize
 Eg: [3,3,3,3,3,1,3]
  var grp = {}, r = [];
  3 => grp{3:[0]}, r=[]
  3 => grp{3:[0,1]}, r=[]
  3 => grp{3:[0,1,2]}, r=[]  => grp{3:[]}, r=[[0,1,2]]
  3 => grp{3:[3]}, r=[[0,1,2]]
  3 => grp{3:[3,4]}, r=[[0,1,2]]
  1 => grp{3:[3,4], 1:[5]}, r=[[0,1,2]] => grp{3:[3,4,6], 1:[]}, r=[[0,1,2], [5]]
  3 => grp{3:[3,4,6], 1:[]}, r=[[0,1,2]. [5]] => 1 => grp{3:[], 1:[]}, r=[[0,1,2], [5], [3,4,6]]
  
  at the end of the for, we would have the map with keys with values as empty [ ] & result [ ] with [ ]s of groupSizes[idx]
  
  Complezity = 
  - loops = for(n) => O(n)
  
  => O(n)
 */

var groupThePeople = function(groupSizes) {
    //group = store each group in a map
    let grp = {};
    let r = [];
    
    for(let i=0; i<groupSizes.length; i++){
        //add current element ot the groupsize [ ]
        if(!grp[groupSizes[i]]){
            //if current groupSizes[i] is not present in the map, add it & set its value = [i]
            grp[groupSizes[i]] = [i]
            
        }else{
            //push i to the groupSizes[i] [ ]
            grp[groupSizes[i]].push(i);
            
        }
        
        // if groupSizes[i] matches the length, push it to the result & reset the []
        if(grp[groupSizes[i]].length == groupSizes[i]){
            r.push(grp[groupSizes[i]]);
            grp[groupSizes[i]] = [];
        }
    }
    
    return r;
};
