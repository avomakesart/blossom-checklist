package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Section struct {
	Title       string
	Description string
	Items       []string
}

type BlossomChecklist struct {
	Intro    string
	Sections []Section
}

func main() {
	checklist := BlossomChecklist{
		Intro: "Make sure to follow these steps before you click on the merge button. We know that you are excited to merge your things, but let's be safe first.",
		Sections: []Section{
			{
				Title:       "Mobile",
				Description: "Make sure to verify your feature on the app and mobile browser",
				Items:       []string{"Check on Android", "Check on IOS", "Check on mobile browser"},
			},
			{
				Title:       "Unit test",
				Description: "Make sure to add proper unit tests to your feature.",
				Items: []string{
					"Add tests to your feature thinking about the customer actions (ex. clicks, hovers, focus, etc.)",
					"If there are more variants make sure to cover all of them.",
					"Make sure to verify a11y and (if necessary) add a11y tests.",
				},
			},
			{
				Title:       "Cross browser test",
				Description: "Sometimes browsers act different, and customers can have a different browser than you, so make sure to tests your features on the most common browsers.",
				Items:       []string{"Firefox", "Chrome", "Safari (if needed)"},
			},
			{
				Title:       "Different accounts",
				Description: "If you think is necessary to verify, make sure to login into a business account and test your feature.",
				Items:       []string{"Check with Business Account"},
			},
			{
				Title:       "Translations",
				Description: "Our texts needs to be translated so follow these steps.",
				Items: []string{
					"Add texts to rosetta (stg and pro)",
					"Request translations to french.",
				},
			},
			{
				Title:       "Mock data",
				Description: "We have generated mock data based on rosetta texts, but sometimes we need different edge cases, and the generated texts is not the best option, so make sure to follow these steps.",
				Items: []string{
					"Verify long texts (ex. product names, partner names)",
					"Verify special texts (ex. dynamic data)",
				},
			},
			{
				Title:       "Happy flows",
				Description: "Make sure to execute your whole flow, doing manual tests, like what happens if we have an error, what happen if we don't have data, or maybe we have a slow connection:",
				Items: []string{
					"Check what happens if error occured",
					"Check what happens if data isn't there (empy states?)",
					"Check what happen if we have slow connection.",
					"Check with long status/messages/names",
					"Double check by other frontend on feature env.",
					"Verify the story of your feature and make sure it is aligned, if something is not covered check with the reporter.",
					"Check if is necessary to add cypress tests (this can also be in a separate story).",
				},
			},
			{
				Title:       "After merging",
				Description: "You are done but sort of, your MR is merged but it is important to keep an eye on your feature, so follow these steps to make sure all is good:",
				Items: []string{
					"Check logs",
					"Check performance",
					"Check translations and make sure that your feature looks good on both languages, if not check with your favorite designer for the best option.",
					"Verify that tagging it is correct on stg and pro, and double check with your favorite business person. (optional)",
					"Checked with design before going to PRO",
				},
			},
		},
	}

	fmt.Println("Pre-Push Checklist:")
	fmt.Println(checklist.Intro)
	fmt.Println()

	allChecked := true
	uncheckedItems := []string{}

	reader := bufio.NewReader(os.Stdin)

	for _, section := range checklist.Sections {
		fmt.Printf("%s\n%s\n", section.Title, section.Description)

		for _, item := range section.Items {
			fmt.Printf("%s (y/N): ", item)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if strings.ToLower(answer) != "y" {
				allChecked = false
				uncheckedItems = append(uncheckedItems, fmt.Sprintf("%s: %s", section.Title, item))
			}
		}

		fmt.Println()
	}

	if allChecked {
		fmt.Println("Great! You have completed all checks. You can proceed with your push.")
		os.Exit(0)
	} else {
		fmt.Println("The following items are still unchecked:")
		for _, item := range uncheckedItems {
			fmt.Printf("- %s\n", item)
		}
		fmt.Println("Please complete all checks before pushing.")
		os.Exit(1)
	}
}
