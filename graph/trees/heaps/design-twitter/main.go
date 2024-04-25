package main

import "container/heap"

type Tweet struct {
	tweetID int
	time    int
}

type User struct {
	userID  int
	follows map[int]bool
	tweets  []*Tweet
}

type Twitter struct {
	Time  int
	Users map[int]*User
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i].time > h[j].time }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor() Twitter {
	return Twitter{
		Users: make(map[int]*User),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Time++

	tweet := &Tweet{time: this.Time, tweetID: tweetId}
	if user, exists := this.Users[userId]; exists {
		user.tweets = append(user.tweets, &Tweet{
			time:    this.Time,
			tweetID: tweetId,
		})
	} else {
		this.Users[userId] = &User{
			userID:  userId,
			follows: make(map[int]bool),
			tweets:  []*Tweet{tweet},
		}
	}

}

func (this *Twitter) GetNewsFeed(userId int) []int {

	if user, exists := this.Users[userId]; exists {

		// Underlying implementation of heap is done using slice
		feedTweets := make(TweetHeap, len(user.tweets))
		copy(feedTweets, user.tweets)

		for followeeId := range user.follows {
			if followee, ok := this.Users[followeeId]; ok {
				feedTweets = append(feedTweets, followee.tweets...)
			}
		}

		var feeds []int
		heap.Init(&feedTweets)
		for i := 0; i < 10 && len(feedTweets) > 0; i++ {
			tweet := heap.Pop(&feedTweets).(*Tweet)
			feeds = append(feeds, tweet.tweetID)
		}

		return feeds
	}

	return nil
}

func (this *Twitter) Follow(followerId int, followeeId int) {

	if followerId != followeeId {
		// Check if follower first exists or not
		if user, exists := this.Users[followerId]; exists {
			user.follows[followeeId] = true
		} else {
			// Create the `follower` user in twitter and make him follow the `followeeId`
			this.Users[followerId] = &User{
				userID:  followeeId,
				follows: map[int]bool{followeeId: true},
			}
		}
	}

}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId != followeeId {
		if user, exists := this.Users[followerId]; exists {
			delete(user.follows, followeeId)
		}
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
