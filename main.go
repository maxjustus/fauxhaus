package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-faker/faker/v4"
)

func main() {
	fmt.Fprintln(os.Stderr, "Listening on stdin")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()
		t = strings.Split(t, "\t")[0]

		if t == "" {
			// ClickHouse sends empty input sometimes, just ignore
			continue
		}

		fakeData := generateFakeData(t)

		fmt.Println(fakeData)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func generateFakeData(method string) string {
	switch method {
	case "first_name":
		return faker.FirstName()
	case "last_name":
		return faker.LastName()
	case "name":
		return faker.Name()
	case "title_male":
		return faker.TitleMale()
	case "title_female":
		return faker.TitleFemale()
	case "phone_number":
		return faker.Phonenumber()
	case "e164_phone_number":
		return faker.E164PhoneNumber()
	case "address1":
		return faker.GetRealAddress().Address
	case "city":
		return faker.GetRealAddress().City
	case "state":
		return faker.GetRealAddress().State
	case "postal_code":
		return faker.GetRealAddress().PostalCode
	case "email":
		return faker.Email()
	case "ipv4":
		return faker.IPv4()
	case "ipv6":
		return faker.IPv6()
	case "mac_address":
		return faker.MacAddress()
	case "url":
		return faker.URL()
	default:
		return faker.Sentence()
	}
}
