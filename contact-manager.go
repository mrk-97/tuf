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
	removed := false

	for i := 0; i < len(*contacts); {
		if (*contacts)[i].Name == name {
			// Remove the contact
			*contacts = append((*contacts)[:i], (*contacts)[i+1:]...)
			removed = true
			// Don't increment i since we need to check the new item at i
		} else {
			i++ // Only increment if we did not remove an item
		}
	}

	if removed {
		fmt.Printf("Removed all contacts with name: %s\n", name)
	} else {
		fmt.Printf("Contact not found: %s\n", name)
	}
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
	contacts := make([]Contact, 0)

	// Adding multiple contacts with the same name
	AddContact(&contacts, Contact{Name: "Manoj Reddy", Email: "manoj.reddy@beepkart.com", Phone: "7075198583"})
	AddContact(&contacts, Contact{Name: "Manoj Reddy", Email: "manoj2.reddy@beepkart.com", Phone: "7075198584"})

	// Print contacts before any removal
	PrintContacts(&contacts)

	// Remove all contacts with the name "Manoj Reddy"
	RemoveContact(&contacts, "Manoj Reddy")

	// Print remaining contacts after removal
	PrintContacts(&contacts)

	// Attempting to remove again to check behavior
	RemoveContact(&contacts, "Manoj Reddy")
	PrintContacts(&contacts)
}
