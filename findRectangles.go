/*

Imagine we have an image. We'll represent this image as a simple 2D array where every pixel is a 1 or a 0.

The image you get is known to have potentially many distinct rectangles of 0s on a background of 1's. Write a function that takes in the image and returns the coordinates of all the 0 rectangles -- top-left and bottom-right; or top-left, width and height.

image1 = [
  [0, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [0, 1, 1, 0, 0, 0, 1],
  [1, 0, 1, 0, 0, 0, 1],
  [1, 0, 1, 1, 1, 1, 1],
  [1, 0, 1, 0, 0, 1, 1],
  [1, 1, 1, 0, 0, 1, 1],
  [1, 1, 1, 1, 1, 1, 0],
]

R : [[0,0], [2,3,2,4, 2,5,3,2,]] 
isNeighbour(R, [2,4])

Sample output variations (only one is necessary):

findRectangles(image1) =>
  // (using top-left-row-column and bottom-right):
  [
    [[0,0],[0,0]],
    [[2,0],[2,0]],
    [[2,3],[3,5]],
    [[3,1],[5,1]],
    [[5,3],[6,4]],
    [[7,6],[7,6]],
  ]
  // (using top-left-x-y and width/height):
  [
    [[0,0],[1,1]],
    [[0,2],[1,1]],
    [[3,2],[3,2]],
    [[1,3],[1,3]],
    [[3,5],[2,2]],
    [[6,7],[1,1]],
  ]

Other test cases:

image2 = [
  [0],
]

findRectangles(image2) =>
  // (using top-left-row-column and bottom-right):
  [
    [[0,0],[0,0]],
  ]

  // (using top-left-x-y and width/height):
  [
    [[0,0],[1,1]],
  ]

image3 = [
  [1],
]

findRectangles(image3) => []

image4 = [
  [1, 1, 1, 1, 1],
  [1, 0, 0, 0, 1],
  [1, 0, 0, 0, 1],
  [1, 0, 0, 0, 1],
  [1, 1, 1, 1, 1],
]

findRectangles(image4) =>
  // (using top-left-row-column and bottom-right or top-left-x-y and width/height):
  [
    [[1,1],[3,3]],
  ]

n: number of rows in the input image
m: number of columns in the input image

*/

package main
import "fmt"

func main() {


  image1 := [][]int {
      []int{0, 1, 1, 1, 1, 1, 1},
      []int{1, 1, 1, 1, 1, 1, 1},
      []int{0, 1, 1, 0, 0, 0, 1},
      []int{1, 0, 1, 0, 0, 0, 1},
      []int{1, 0, 1, 1, 1, 1, 1},
      []int{1, 0, 1, 0, 0, 1, 1},
      []int{1, 1, 1, 0, 0, 1, 1},
      []int{1, 1, 1, 1, 1, 1, 0},
  }

  image2 := [][]int {
      []int{0},
  }

  image3 := [][]int {
      []int{1},
  }

  image4 := [][]int {
    []int{1, 1, 1, 1, 1},
    []int{1, 0, 0, 0, 1},
    []int{1, 0, 0, 0, 1},
    []int{1, 0, 0, 0, 1},
    []int{1, 1, 1, 1, 1},
  }

  
//   fmt.Println("image 1 result = ", findRectangle(image1))
//   fmt.Println("image2 result = ", findRectangle(image2))
//   fmt.Println("image 3 result = ", findRectangle(image3))
//   fmt.Println("image 4 result = ", findRectangle(image4))
//   fmt.Println("image 5 result = ", findRectangle(image5))
  
}

// input m = matrix [][]int{image1}
// output [][]int{[2,4], [3,5]}
func findRectangle(m [][]int) [][]int{
  start := []int{}
  end := []int{}
  
  for i:=0; i < len(m); i++ {
    for j:=0; j< len(m[0]); j++ {
      
      if m[i][j] == 0 {
//         fmt.Println(start, end, i, j)
        if len(start) == 0 {
          start = append(start, i)
          start = append(start, j)
        }else{
          end = append(end, i)
          end = append(end, j)
        }
        
      }
      
    }
  }
  
  fmt.Println(start, end)
  result := [][]int{start}
  
  if len(end) > 0 {
    l := len(end)
//     fmt.Println("end length = ", l)
//     fmt.Println("adding end last 2 elements to result = ", result)
    result = append(result, []int{end[l-2], end[l-1]})
  }
  
  //if there is only 1 0
  if len(result) == 1 {
//     fmt.Println("adding start to result to account for one 0 case = ", result)
    result = append(result, start)
  }
  
  fmt.Println("result = ", result)
  return result
}


func findMultipleRectangle(m [][]int) [][]int{
  R := [][]int{}
  
  for i:=0; i < len(m); i++ {
    for j:=0; j< len(m[0]); j++ {
      
      if m[i][j] == 0 {
//         fmt.Println(start, end, i, j)
        add(R, []int{i,j})
        
      }
      
    }
    
  }
  
  for i, _ := range R {
    if len(R[i] > 2) {
      R[i] = 
    }
  }
  
  return R
}
  
func add(R [][]int, curr[]int) [][]int {
  for i:=0; i < len(R); i++ {
    for j:=0; j< len(R[0]); j++ {
      if R[i][j+1] == curr[0][1] || R[i+1][j] == curr[0][1]{
        //0 in next col
        R[i] = append(R[i], curr...)
      }else{
        R = append(R, curr)
      }

    }
  }
  
  return R
}