package main

import "fmt"

type Contact struct {
	Name  string
	Email string
	Phone string
}

func AddContact(contacts *[]Contact, contact Contact) {
	*contacts = append(*contacts, contact)
}

func RemoveContact(contacts *[]Contact, name string) {
	for i := range *contacts {
		if (*contacts)[i].Name == name {
			*contacts = append((*contacts)[:i], (*contacts)[i+1:]...)
			fmt.Printf("Removed Contact : %v\n", name)
			return
		}
	}
	fmt.Printf("Contact not found: %s\n", name)
}

func PrintContacts(contacts *[]Contact) {
	if len(*contacts) == 0 {
		fmt.Printf("Phonebook is empty\n")
		return
	}
	fmt.Printf("Phonebook contacts:\n")
	for _, v := range *contacts {
		fmt.Printf("- Name: %s, Email: %s, Phone: %s\n", v.Name, v.Email, v.Phone)
	}
}

func main() {
	contact := make([]Contact, 0)

	AddContact(&contact, Contact{Name: "Manoj Reddy", Email: "manoj.reddy@beepkart.com", Phone: "7075198583"})
	RemoveContact(&contact, "Manoj Reddy")
	PrintContacts(&contact)

}
