package main

import (
	"fmt"
)

type Tweet struct {
	Id   int
	Time int
}

type User struct {
	Posts   []Tweet
	Follows map[int]bool
}

type Twitter struct {
	UserId map[int]*User
	count  int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{UserId: make(map[int]*User)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := Tweet{tweetId, this.count}
	this.count++
	if user, ok := this.UserId[userId]; ok {
		user.Posts = append(user.Posts, tweet)
	} else {
		newUser := User{Posts: []Tweet{tweet}, Follows: make(map[int]bool)}
		this.UserId[userId] = &newUser
	}
}

type TweetHeapNode struct {
	UserId int
	Post   Tweet
}

type TweetHeap struct {
	Heap []TweetHeapNode
}

func (this *TweetHeap) Add(userId int, tweet Tweet) {
	this.Heap = append(this.Heap, TweetHeapNode{userId, tweet})
	this.moveUp(len(this.Heap) - 1)
}

func (this *TweetHeap) GetTop() (int, Tweet) {
	res := this.Heap[0]
	this.Heap[0] = this.Heap[len(this.Heap)-1]
	this.Heap = this.Heap[:len(this.Heap)-1]
	this.moveDown(0)
	return res.UserId, res.Post
}

func (this *TweetHeap) moveUp(idx int) {
	parent := (idx - 1) / 2
	for parent >= 0 && this.Heap[parent].Post.Time < this.Heap[idx].Post.Time {
		this.Heap[parent], this.Heap[idx] = this.Heap[idx], this.Heap[parent]
		idx = parent
		parent = (idx - 1) / 2
	}
}

func (this *TweetHeap) moveDown(idx int) {
	left := idx*2 + 1
	right := idx*2 + 2
	for left < len(this.Heap) || right < len(this.Heap) {
		max := left
		if right < len(this.Heap) && this.Heap[max].Post.Time < this.Heap[right].Post.Time {
			max = right
		}
		if this.Heap[idx].Post.Time < this.Heap[max].Post.Time {
			this.Heap[idx], this.Heap[max] = this.Heap[max], this.Heap[idx]
			idx = max
			left = idx*2 + 1
			right = idx*2 + 2
		} else {
			break
		}
	}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	user, ok := this.UserId[userId]
	if !ok {
		return []int{}
	}
	nowAt := make(map[int]int)
	heap := TweetHeap{}
	if len(user.Posts) > 0 {
		nowAt[userId] = len(user.Posts) - 1
		heap.Add(userId, user.Posts[len(user.Posts)-1])
	}
	for k, _ := range user.Follows {
		followee := this.UserId[k]
		if len(followee.Posts) > 0 {
			nowAt[k] = len(followee.Posts) - 1
			heap.Add(k, followee.Posts[len(followee.Posts)-1])
		}
	}
	res := make([]int, 0)
	for i := 0; i < 10 && len(heap.Heap) > 0; i++ {
		recentId, recentTweet := heap.GetTop()
		res = append(res, recentTweet.Id)
		if nowAt[recentId] > 0 {
			nowAt[recentId]--
			followee := this.UserId[recentId]
			heap.Add(recentId, followee.Posts[nowAt[recentId]])
		}
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	follower, ok1 := this.UserId[followerId]
	followee, ok2 := this.UserId[followeeId]
	if !ok1 {
		follower = &User{make([]Tweet, 0), make(map[int]bool)}
		this.UserId[followerId] = follower
	}
	if !ok2 {
		followee = &User{make([]Tweet, 0), make(map[int]bool)}
		this.UserId[followeeId] = followee
	}

	follower.Follows[followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	follower, ok1 := this.UserId[followerId]
	_, ok2 := this.UserId[followeeId]
	if !ok1 || !ok2 {
		return
	}

	delete(follower.Follows, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {
	obj := Constructor()
	obj.PostTweet(1, 5)
	fmt.Println(obj.GetNewsFeed(1))

	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	fmt.Println(obj.GetNewsFeed(1))
	fmt.Println(obj.GetNewsFeed(2))

	obj.Unfollow(1, 2)
	fmt.Println(obj.GetNewsFeed(1))
	fmt.Println(obj.GetNewsFeed(2))
}
