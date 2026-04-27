package main

import (
	"fmt"
	"os"

	"github.com/kevincoj/Spiderweb_Simulations_Project/serpentapi"
)

func main() {
	apiKey := os.Getenv("SERPENT_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: SERPENT_API_KEY environment variable is not set.")
		fmt.Println("Get a free API key at https://apiserpent.com and set it:")
		fmt.Println("  export SERPENT_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	client := serpentapi.NewClient(apiKey, "google")

	// Search for different spider web types to compare structures.
	webTypes := []string{
		"orb weaver",
		"cobweb tangled",
		"funnel web",
		"sheet web",
		"triangle web",
	}

	for _, webType := range webTypes {
		fmt.Printf("\n=== Searching: %s spider web ===\n", webType)

		results, err := client.SearchSpiderWebs(webType)
		if err != nil {
			fmt.Printf("Error searching for %s: %v\n", webType, err)
			continue
		}

		if len(results.OrganicResults) == 0 {
			fmt.Printf("No results found for %s\n", webType)
			continue
		}

		// Display the top 3 results for each web type.
		limit := 3
		if len(results.OrganicResults) < limit {
			limit = len(results.OrganicResults)
		}
		for i, r := range results.OrganicResults[:limit] {
			fmt.Printf("  %d. %s\n     %s\n     %s\n", i+1, r.Title, r.Link, r.Snippet)
		}
	}

	// Direct search as specified in the task instructions.
	fmt.Println("\n=== Finding cheap third-party search engine APIs ===")
	resp, err := client.Search("cheapest third party search api")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for i, r := range resp.OrganicResults {
		fmt.Printf("  %d. %s\n     %s\n", i+1, r.Title, r.Link)
		if i >= 4 {
			break
		}
	}
}
