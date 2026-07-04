package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type TimeBucket string

const (
	Morning   TimeBucket = "Morning Bird (5AM - 11AM)"
	Afternoon TimeBucket = "Afternoon Hustler (12PM - 5PM)"
	Evening   TimeBucket = "Evening Coder (6PM - 10PM)"
	Night     TimeBucket = "Night Owl (11PM - 4AM)"
)

type GitHubEvent struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

type JSONOutput struct {
	TotalCommits   int            `json:"total_commits"`
	DominantBucket string         `json:"dominant_bucket"`
	Buckets        map[string]int `json:"buckets"`
}

func main() {
	targetDir := flag.String("dir", ".", "Target git directory to analyze")
	githubUser := flag.String("github", "", "GitHub username to fetch remote public events")
	jsonFlag := flag.Bool("json", false, "Output results as machine-readable JSON")
	flag.Parse()

	var hours []int
	var err error

	if *githubUser != "" {
		hours, err = fetchGitHubHours(*githubUser)
		if err != nil {
			fmt.Printf("Error fetching GitHub events: %v\n", err)
			os.Exit(1)
		}
	} else {
		hours, err = fetchLocalHours(*targetDir)
		if err != nil {
			fmt.Printf("Error analyzing local git log: %v\n", err)
			os.Exit(1)
		}
	}

	// We only care about the last 50 commits according to the spec
	if len(hours) > 50 {
		hours = hours[len(hours)-50:]
	}

	if len(hours) == 0 {
		if *jsonFlag {
			fmt.Println("{}")
		} else {
			fmt.Println("No commits found to analyze.")
		}
		os.Exit(0)
	}

	buckets := map[TimeBucket]int{
		Morning:   0,
		Afternoon: 0,
		Evening:   0,
		Night:     0,
	}

	for _, hour := range hours {
		if hour >= 5 && hour < 12 {
			buckets[Morning]++
		} else if hour >= 12 && hour < 18 {
			buckets[Afternoon]++
		} else if hour >= 18 && hour < 23 {
			buckets[Evening]++
		} else {
			buckets[Night]++ // 23 (11 PM) or 0-4 (12 AM - 4 AM)
		}
	}

	var dominantBucket TimeBucket
	var maxCount int

	keys := []TimeBucket{Morning, Afternoon, Evening, Night}
	for _, k := range keys {
		if buckets[k] > maxCount {
			maxCount = buckets[k]
			dominantBucket = k
		}
	}

	if *jsonFlag {
		printJSON(buckets, dominantBucket, len(hours))
	} else {
		printProfile(dominantBucket, maxCount, len(hours), *githubUser)
	}
}

func fetchGitHubHours(username string) ([]int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events/public", username)
	
	// Create request with custom User-Agent to avoid getting blocked
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "DevSelfie-CLI")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	var hours []int
	for _, event := range events {
		if event.Type == "PushEvent" {
			t, err := time.Parse(time.RFC3339, event.CreatedAt)
			if err == nil {
				// Convert UTC to local time
				t = t.Local()
				hours = append(hours, t.Hour())
			}
		}
	}
	
	// Events are newest first, let's reverse to keep chronological consistency 
	// (though order doesn't matter for buckets)
	return hours, nil
}

func fetchLocalHours(targetDir string) ([]int, error) {
	logPath := filepath.Join(targetDir, ".git", "logs", "HEAD")

	file, err := os.Open(logPath)
	if err != nil {
		return nil, fmt.Errorf("could not open %s (make sure you are in a valid git repository)", logPath)
	}
	defer file.Close()

	var hours []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "> ")
		if len(parts) < 2 {
			continue
		}

		metaParts := strings.Split(parts[1], " ")
		if len(metaParts) < 1 {
			continue
		}

		unixTime, err := strconv.ParseInt(metaParts[0], 10, 64)
		if err != nil {
			continue
		}

		commitTime := time.Unix(unixTime, 0)
		hours = append(hours, commitTime.Hour())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hours, nil
}

func printJSON(buckets map[TimeBucket]int, dominant TimeBucket, total int) {
	out := JSONOutput{
		TotalCommits:   total,
		DominantBucket: string(dominant),
		Buckets: map[string]int{
			string(Morning):   buckets[Morning],
			string(Afternoon): buckets[Afternoon],
			string(Evening):   buckets[Evening],
			string(Night):     buckets[Night],
		},
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
}

func printProfile(bucket TimeBucket, count int, total int, remoteUser string) {
	percentage := float64(count) / float64(total) * 100

	fmt.Println("======================================")
	fmt.Printf("   DevSelfie - Git Commit Profile   \n")
	fmt.Println("======================================")
	
	if remoteUser != "" {
		fmt.Printf("\nAnalyzed last %d public pushes for @%s.\n\n", total, remoteUser)
	} else {
		fmt.Printf("\nAnalyzed last %d local commits.\n\n", total)
	}

	switch bucket {
	case Morning:
		art := "\033[33m" + `        \   |   /
         \  |  /
          \ | /
     _.--"'""'--._
   .' _.-'""'-._ '.
  /.-'  .    .   '-.
 /   .   \  /   .   \
|  .  ~ ~~~~~~ ~  .  |
 \   .  ~~~~   .   /
  '.__________________.'
     ___/    \___
    /   \____/   \` + "\033[0m"
		fmt.Println(art)
		fmt.Println("\033[1;33m" + "\n   THE MORNING BIRD" + "\033[0m")
		fmt.Println("You code best when the sun comes up!")
	case Afternoon:
		art := "\033[38;5;208m" + `           \  |  /
         .   \ | /   .
       '.___ \|/ ___.'
       ______(_)______
   '--~-'   /   \   '-~--'
         .-'     '-.
        /   O   O   \
       |   .-----.   |
        \  '-----'  /
         '-.......-'
         ___|_|___` + "\033[0m"
		fmt.Println(art)
		fmt.Println("\033[1;38;5;208m" + "\n  THE AFTERNOON HUSTLER" + "\033[0m")
		fmt.Println("Powering through the midday slump!")
	case Evening:
		art := "\033[35m" + `      .-""""""-.
     /  ___   \
    |  /   \   |
    |  \ o o/  |    < merging PRs after sunset >
    |   \_/    |
    ,\  '-'  /,
   / '.|___|.' \
  /  /|     |\  \
 |__/ |__|__| \__|
      /_____\
     |_______|` + "\033[0m"
		fmt.Println(art)
		fmt.Println("\033[1;35m" + "\n   THE EVENING CODER" + "\033[0m")
		fmt.Println("Winding down with some solid commits.")
	case Night:
		art := "\033[36m" + `        .           .
     *  .  *   .     .    *
   .    .-""""-.   *    .
      / .====.  \    .
  .  /  |    |   \
    |   '----'   |     *
     \  ,--.  ,--.  /
   *  '-|  |-|  |-'   .
        |__| |__|
        /_/   \_\
   .    -~-~-~-~-      *` + "\033[0m"
		fmt.Println(art)
		fmt.Println("\033[1;36m" + "\n     THE NIGHT OWL" + "\033[0m")
		fmt.Println("The world sleeps, but your terminal glows.")
	}

	fmt.Printf("\n%.1f%% of recent commits were in the %s bucket.\n", percentage, string(bucket))
	fmt.Println("======================================")
}
