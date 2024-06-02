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
	message string
}

var users [maxUsers]user
var emails [maxEmails]email
var userCount, emailCount int = 0, 0

func main() {
	var choice int
	for {
		fmt.Println("1. Register User")
		fmt.Println("2. Persetujuan Admin")
		fmt.Println("3. Kirim Email")
		fmt.Println("4. Lihat Inbox")
		fmt.Println("5. Reply Email")
		fmt.Println("6. Delete Email")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			registerUser()
		case 2:
			persetujuanAdmin()
		case 3:
			kirimEmail()
		case 4:
			lihatInbox()
		case 5:
			replyEmail()
		case 6:
			deleteEmail()
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid, harap coba lagi.")
		}
	}
}

func registerUser() {
	var username, password, email string
	if userCount >= maxUsers {
		fmt.Println("User limit reached.")
		return
	}
	fmt.Print("Masukan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukan password: ")
	fmt.Scan(&password)
	fmt.Print("Masukan email: ")
	fmt.Scan(&email)
	users[userCount] = user{username, password, email, false}
	userCount++
	fmt.Println("User registered, menunggu persetujuan admin.")
}

func persetujuanAdmin() {
	var response string
	for i := 0; i < userCount; i++ {
		if !users[i].active {
			fmt.Printf("Terima user %s (yes/no)? ", users[i].username)
			fmt.Scan(&response)
			if strings.ToLower(response) == "yes" {
				users[i].active = true
				fmt.Println("User diterima.")
			} else {
				fmt.Println("User tidak diterima.")
			}
		}
	}
}

func kirimEmail() {
	var from, to, subject, message string
	fmt.Print("Masukan username kamu: ")
	fmt.Scan(&from)
	if !isUserActive(from) {
		fmt.Println("User tidak aktif atau tidak ketemu.")
		return
	}
	fmt.Print("Masukkan username penerima: ")
	fmt.Scan(&to)
	if !isUserActive(to) {
		fmt.Println("User tidak aktif atau tidak ketemu.")
		return
	}
	fmt.Print("Masukan subject: ")
	fmt.Scan(&subject)
	fmt.Print("Masukan Message: ")
	fmt.Scanln()
	message = readMultilineInput()

	if emailCount >= maxEmails {
		fmt.Println("Batas email tercapai.")
		return
	}

	emails[emailCount] = email{from, to, subject, message}
	emailCount++
	fmt.Println("Email terkirim.")
}
func lihatInbox() {
	var username string
	var inbox [maxEmails]email
	var inboxCount int
	fmt.Print("Masukan username kamu: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User tidak aktif atau tidak ketemu.")
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
		fmt.Printf("Dari: %s, Subject: %s\n", inbox[i].from, inbox[i].subject)
		fmt.Println("Pesan:", inbox[i].message)
	}
}

func replyEmail() {
	var username, to, subject, message string
	fmt.Print("Masukan username kamu: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User tidak aktif atau tidak ketemu.")
		return
	}

	fmt.Print("Masukkan username penerima: ")
	fmt.Scan(&to)
	if !isUserActive(to) {
		fmt.Println("User tidak aktif atau tidak ketemu.")
		return
	}
	fmt.Print("Masukan subject: ")
	fmt.Scan(&subject)
	fmt.Print("Tulis balasan Message: ")
	fmt.Scanln()
	message = readMultilineInput()

	emails[emailCount] = email{username, to, subject, message}
	emailCount++
	fmt.Println("Email terkirim.")
}

func deleteEmail() {
	var username, subject string
	fmt.Print("Masukan username kamu: ")
	fmt.Scan(&username)
	if !isUserActive(username) {
		fmt.Println("User tidak aktif atau tidak ketemu.")
		return
	}

	fmt.Print("Masukan subject dari email yang ingin di delete: ")
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
		fmt.Println("Email Tidak Ketemu.")
	}
}
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
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
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
func readMultilineInput() string {
	var lines []string
	var line string
	for line != "." {
		fmt.Scanf("%s", &line)
		if line != "." {
			lines = append(lines, line)
			line = ""
		}
	}

	return strings.Join(lines, " ")
}
