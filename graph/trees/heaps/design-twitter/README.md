# Design Twitter

Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, and see the 10 most recent tweets in their news feed.


## Intuition

To implement this simplified version of Twitter, we need to manage users and their tweets. Each user has an ID, a list of followers, and a list of tweets. We can use a hash map to store users, where the key is the user ID and the value is a User struct. Each User struct contains the user's ID, a set of followers, and a list of tweets.

To retrieve the news feed for a user, we need to fetch the user's tweets and the tweets of the users they follow. We can use a min heap to efficiently retrieve the 10 most recent tweets from the combined list of tweets.

## Approach

1. Define structs for Tweet and User. Tweet struct contains the tweet ID and the time it was posted. User struct contains the user ID, a set of followers, and a list of tweets.
2. Implement methods for the Twitter class:
    - **Constructor**: Initialize a Twitter object with an empty map of users.
    - **PostTweet**: Create a new tweet and add it to the user's list of tweets. If the user doesn't exist, create a new user.
    - **GetNewsFeed**: Retrieve the 10 most recent tweets from the user's tweets and the tweets of the users they follow. Use a min heap to efficiently fetch the tweets.
    - **Follow**: Add the followee to the follower's set of followers.
    - **Unfollow**: Remove the followee from the follower's set of followers.
3. Initialize the Twitter object and call its methods as required.

## Time Complexity

- **PostTweet**: \(O(1)\) - Adding a tweet to a user's list of tweets is constant time.
- **GetNewsFeed**: \(O(n \log k)\) - Retrieving the news feed involves merging and sorting the user's tweets and the tweets of the users they follow, where \(n\) is the total number of tweets and \(k\) is the number of tweets to retrieve (10 in this case).
- **Follow** and **Unfollow**: \(O(1)\) - Adding or removing a follower is constant time.

## Space Complexity

- \(O(n + m)\), where \(n\) is the number of users and \(m\) is the total number of tweets across all users.