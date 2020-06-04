import "fmt"

/*
leet code 362
https://leetcode.com/problems/design-hit-counter/submissions/
Design a hit counter which counts the number of hits received in the past 5 minutes.
Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing).
You may assume that the earliest timestamp starts at 1.
It is possible that several hits arrive roughly at the same time.

Example:

HitCounter counter = new HitCounter();

// hit at timestamp 1.
counter.hit(1);

// hit at timestamp 2.
counter.hit(2);

// hit at timestamp 3.
counter.hit(3);

// get hits at timestamp 4, should return 3.
counter.getHits(4);

// hit at timestamp 300.
counter.hit(300);

// get hits at timestamp 300, should return 4.
counter.getHits(300);

// get hits at timestamp 301, should return 3.
counter.getHits(301);
Follow up:
*/
const boundary int = 300

type Hit struct {
	Timestamp int
	Next      *Hit
}

type HitCounter struct {
	Count int
	Head  *Hit
	Tail  *Hit
}

/** Initialize your data structure here. */
func Constructor() HitCounter {
	hitCounter := HitCounter{
		Count: 0,
	}
	return hitCounter
}

/** Record a hit.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int) {
	fmt.Printf("Post T = %d \n", timestamp)
	hit := Hit{
		Timestamp: timestamp,
	}

	if this.Tail == nil {
		this.Tail = &hit
	} else {
		this.Tail.Next = &hit
		this.Tail = this.Tail.Next
	}

	if this.Head == nil {
		this.Head = this.Tail
	}

	this.Count++

	boundCheck := timestamp - this.Head.Timestamp
	fmt.Println("Post -------- boundCheck = ", boundCheck)
	if boundCheck <= boundary {
		fmt.Println("Post -------- in boundary so not moving head = ", this.Head.Timestamp)
		return
	}
	var curr *Hit
	curr = this.Head
	for boundCheck > boundary {
		curr = curr.Next
		this.Count--
		fmt.Printf("Post -------- moving curr = %d, reducting count = %d \n", curr.Timestamp, this.Count)
		if curr == nil {
			fmt.Println("Post -------- curr nil - count shuld be 0 ", this.Count)
			break
		}
		boundCheck = timestamp - curr.Timestamp
		fmt.Printf("Post -------- >>>>> curr.Timestamp = %d, boundCheck = %d \n", curr.Timestamp, boundCheck)

	}
	fmt.Println("Post -------- out boundary so moving head to curr, curr = ", curr.Timestamp)
	this.Head.Next = curr
	this.Head = this.Head.Next
	fmt.Println("Post -------- out boundary so moving head to = ", this.Head.Timestamp)
}

/** Return the number of hits in the past 5 minutes.
  @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	if this.Head == nil {
		return 0
	}
	boundCheck := timestamp - this.Head.Timestamp
	fmt.Printf("GET T = %d boundCheck = %d cnt = %d H = %d \n", timestamp, boundCheck, this.Count, this.Head.Timestamp)
	if boundCheck < boundary {
		fmt.Println("GET **** 1 retuning result - ", this.Count)
		return this.Count
	}

	//use this.Count value within boundary
	offBountCount := 0
	var curr *Hit
	curr = this.Head
	for curr.Next != nil {
		// fmt.Println("GET boundCheck failed, inc curr = ", curr)
		curr = curr.Next
		offBountCount++
		boundCheck = timestamp - curr.Timestamp
		fmt.Println("GET --2 boundCheck failed after inc curr = ", curr, curr.Timestamp, boundCheck)
		if boundCheck < boundary {
			fmt.Println("GET **** 2 returning result = ", this.Count-offBountCount)
			return this.Count - offBountCount
		}
	}
	fmt.Printf("GET -- *** returning 0 -- 2 boundCheck after for ends & > boundary = %d \n", boundCheck)

	return 0

}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */