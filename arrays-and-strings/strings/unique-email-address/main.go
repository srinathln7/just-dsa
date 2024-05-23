package main

func numUniqueEmails(emails []string) int {
	emailSet := make(map[string]struct{})
	for _, email := range emails {
		modifiedEmail := ""
		skipUntil := false
		for i := 0; i < len(email); i++ {
			if email[i] == '@' {
				modifiedEmail += email[i:]
				break
			}

			if email[i] == '+' {
				skipUntil = true
			}

			if !skipUntil && email[i] != '.' {
				modifiedEmail += string(email[i])
			}
		}

		emailSet[modifiedEmail] = struct{}{}
	}

	return len(emailSet)
}
