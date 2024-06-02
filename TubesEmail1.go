package main

import (
	"fmt"
	"strings"
)

const (
	maxUsers  = 100
	maxEmails = 1000
)

type user struct {
	username string
	password string
	email    string
	active   bool
}

type email struct {
	from    string
	to      string
	subject string
	body    string
}

var users [maxUsers]user
var emails [maxEmails]email
var userCount, emailCount int = 0, 0

func main() {
	var choice int
	for {
		fmt.Println("1. Register User")
		fmt.Println("2. Admin Approval")
		fmt.Println("3. Send Email")
		fmt.Println("4. View Inbox")
		fmt.Println("5. Reply to Email")
		fmt.Println("6. Delete Email")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			registerUser()
		case 2:
			adminApproval()
		case 3:
			sendEmail() // masih harus dibenerin entah apa
		case 4:
			viewInbox()
		case 5:
			replyEmail()
		case 6:
			deleteEmail() // delete email gabisa milih email yg mana yg mau dihapus, cuma bisa milih subject nya apa
		case 7:
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func registerUser() {
	var username, password, email string
	if userCount >= maxUsers {
		fmt.Println("User limit reached.")
		return
	}
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	users[userCount] = user{username, password, email, false}
	userCount++
	fmt.Println("User registered, pending admin approval.")
}

func adminApproval() {
	var response string
	for i := 0; i < userCount; i++ {
		if !users[i].active {
			fmt.Printf("Approve user %s (yes/no)? ", users[i].username)
			fmt.Scan(&response)
			if strings.ToLower(response) == "yes" {
				users[i].active = true
				fmt.Println("User approved.")
			} else {
				fmt.Println("User not approved.")
			}
		}
	}
}

func sendEmail() {
	var from, to, subject, message string
	fmt.Print("Enter your username: ")
	fmt.Scan(&from)
	if !isUserActive(from) {
		fmt.Println("User not active or not found.")
		return
	}
	fmt.Print("Enter recipient username: ")
	fmt.Scan(&to)
	if !isUserActive(to) {
		fmt.Println("Recipient not active or not found.")
		return
	}
	fmt.Print("Enter subject: ")
	fmt.Scan(&subject)
	fmt.Print("Enter Message: ")
	fmt.Scanln()
	message = readMultilineInput()

	if emailCount >= maxEmails {
		fmt.Println("Email limit reached.")
		return
	}

	emails[emailCount] = email{from, to, subject, message}
	emailCount++
	fmt.Println("Email sent.")
}

func readMultilineInput() string {
	var lines []string
	var line string
	for line != "." {
		fmt.Scanf("%s\n", &line)
		if line != "." {
			lines = append(lines, line)
		}
	}

	return strings.Join(lines, " ")
}
func viewInbox() {
	var username string
	var inbox [maxEmails]email
	var inboxCount int
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User not active or not found.")
		return
	}

	fmt.Println("Inbox:")
	for i := 0; i < emailCount; i++ {
		if emails[i].to == username {
			inbox[inboxCount] = emails[i]
			inboxCount++
		}
	}

	selectionSortEmail(&inbox, inboxCount, true)

	for i := 0; i < inboxCount; i++ {
		i = i + 1
		fmt.Printf("No: %d\n", i)
		i = i - 1
		fmt.Printf("From: %s, Subject: %s\n", inbox[i].from, inbox[i].subject)
		fmt.Println("Body:", inbox[i].body)
	}
}

func replyEmail() {
	var username, to, subject, body string
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User not active or not found.")
		return
	}

	fmt.Print("Enter recipient username: ")
	fmt.Scan(&to)
	if !isUserActive(to) {
		fmt.Println("Recipient not active or not found.")
		return
	}
	fmt.Print("Enter subject: ")
	fmt.Scan(&subject)
	fmt.Print("Enter Reply Message: ")
	fmt.Scanln()
	body = readMultilineInput()

	emails[emailCount] = email{username, to, subject, body}
	emailCount++
	fmt.Println("Email sent.")
}

func deleteEmail() {
	var username, subject string
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User not active or not found.")
		return
	}

	fmt.Print("Enter subject of the email to delete: ")
	fmt.Scan(&subject)

	selectionSortEmail(&emails, emailCount, true)

	index := binarySearchEmail(emails[:emailCount], subject)

	if index != -1 {
		for i := index; i < emailCount-1; i++ {
			emails[i] = emails[i+1]
		}
		emailCount = emailCount - 1
		fmt.Println("Email deleted.")
	} else {
		fmt.Println("Email not found.")
	}
}

/*func printInbox() {
	var username string
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User not active or not found.")
		return
	}

	fmt.Println("Inbox:")
	inbox := []Email{}
	for i := 0; i < emailCount; i++ {
		if emails[i].to == username {
			inbox = append(inbox, emails[i])
		}
	}

	selectionSortEmails(inbox, true)

	for _, email := range inbox {
		fmt.Printf("From: %s, Subject: %s\n", email.from, email.subject)
	}
} */

func isUserActive(username string) bool {
	for i := 0; i < userCount; i++ {
		if users[i].username == username && users[i].active {
			return true
		}
	}
	return false
}

func selectionSortEmail(arr *[maxEmails]email, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j].subject < arr[minIdx].subject) ||
				(!ascending && arr[j].subject > arr[minIdx].subject) {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i] // kek gini biar gausah make temp
	}
}
func binarySearchEmail(arr []email, subject string) int {
	kiri, kanan := 0, len(arr)-1
	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		if arr[mid].subject == subject {
			return mid
		} else if arr[mid].subject < subject {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}
	return -1
}

// 1. reply ama message masih aneh huruf pertamanya ga ke output
// 2. delete yg pake binary butuh dipelajarin
