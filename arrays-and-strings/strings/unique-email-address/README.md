## OP

### Problem Statement
Given a list of email addresses, you need to count the number of unique email addresses. An email address is considered unique based on the following rules:
- In the local part of the email (before the '@' character), all characters after a '+' are ignored.
- All '.' characters in the local part are ignored.
- The domain part (after the '@' character) is not modified.

### Intuition
To determine the uniqueness of email addresses, we need to normalize each email address by processing its local part according to the rules. By storing the normalized emails in a set, we can easily count the unique ones.

### Approach
1. Initialize an empty set to store unique emails.
2. Iterate through each email in the list.
3. For each email, initialize an empty string for the modified email and a flag `skipUntil` set to `false`.
4. Iterate through the characters of the email:
   - If the character is '@', append the rest of the email (domain part) to `modifiedEmail` and break the loop.
   - If the character is '+', set `skipUntil` to `true`.
   - If `skipUntil` is `false` and the character is not '.', append the character to `modifiedEmail`.
5. Add the modified email to the set.
6. The number of unique emails is the size of the set.

### Time Complexity
- **O(n * m)**: where `n` is the number of emails and `m` is the average length of an email. We process each character of each email once.

### Space Complexity
- **O(n * m)**: We store each unique email in the set.
