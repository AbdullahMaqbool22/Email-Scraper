package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// APIInfo represents information about an API including its URL and API key
type APIInfo struct {
	URL    string `json:"url"`
	APIKey string `json:"api_key"`
}

// DomainInfo represents the domain name
type DomainInfo struct {
	Domain string     `json:"domain"`
	APIs   []APIInfo `json:"apis"`
}

// Email represents an email address
type Email struct {
	Value string `json:"value"`
}

// EmailResponse represents the response from the API
type EmailResponse struct {
	Data struct {
		Emails []Email `json:"emails"`
	} `json:"data"`
}

func getEmailsFromDomain(domain DomainInfo) ([]string, error) {
	var allEmails []string
	for _, api := range domain.APIs {
		url := fmt.Sprintf("%s?domain=%s&api_key=%s", api.URL, domain.Domain, api.APIKey)
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var emailResponse EmailResponse
		err = json.Unmarshal(body, &emailResponse)
		if err != nil {
			return nil, err
		}

		for _, email := range emailResponse.Data.Emails {
			allEmails = append(allEmails, email.Value)
		}
	}
	return allEmails, nil
}

func saveEmailsToFile(emails []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, email := range emails {
		_, err := file.WriteString(email + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	domain := DomainInfo{
		Domain: "Your domain here",
		APIs: []APIInfo{
			{URL: "Your api here", APIKey: "Your api key here"},
			// Add more API URLs with their respective keys as needed
		},
	}
	emails, err := getEmailsFromDomain(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = saveEmailsToFile(emails, "domain_emails.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Email addresses saved to domain_emails.txt")
}
