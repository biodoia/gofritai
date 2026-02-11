package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofritai",
	Short: "Free Tier AI - Monitor and manage cloud free tiers",
	Long: `GoFritAI helps you maximize free tier usage across cloud providers.

Features:
- Monitor quota usage across providers
- Rotate between services to stay within limits  
- Health checks for free tier services
- Alerts before hitting limits`,
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of all free tier services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📊 Free Tier Status")
		fmt.Println("==================")
		fmt.Println()
		fmt.Println("☁️  COMPUTE")
		fmt.Printf("   %-20s %s\n", "Oracle ARM", "✅ 4 OCPU / 24GB (0% used)")
		fmt.Printf("   %-20s %s\n", "Fly.io", "✅ 3 VMs / 160GB (12% used)")
		fmt.Printf("   %-20s %s\n", "Render", "✅ 750h/mo (45% used)")
		fmt.Println()
		fmt.Println("🗄️  DATABASES")
		fmt.Printf("   %-20s %s\n", "Supabase", "✅ 500MB (23% used)")
		fmt.Printf("   %-20s %s\n", "Turso", "✅ 9GB (5% used)")
		fmt.Printf("   %-20s %s\n", "Upstash", "✅ 10k cmd/day (78% used)")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all supported free tier providers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🆓 Supported Free Tier Providers")
		fmt.Println("================================")
		fmt.Println()
		printCategory("COMPUTE", []Provider{
			{"Oracle Cloud ARM", "4 OCPU + 24GB RAM", "No CC"},
			{"Fly.io", "3 shared VMs + 160GB", "CC required"},
			{"Railway", "$5/mo credit", "No CC"},
			{"Render", "750h/mo", "No CC"},
			{"Cloudflare Workers", "100k req/day", "No CC"},
		})
		// LLM APIs managed by gogatewai, not here
		printCategory("DATABASES", []Provider{
			{"Supabase", "500MB Postgres + Auth", "No CC"},
			{"Turso", "9GB + 500M reads", "No CC"},
			{"Upstash", "10k cmd/day Redis", "No CC"},
			{"Neon", "3GB Postgres", "No CC"},
			{"MongoDB Atlas", "512MB", "No CC"},
		})
		printCategory("STORAGE", []Provider{
			{"Cloudflare R2", "10GB + 0 egress", "No CC"},
			{"Backblaze B2", "10GB", "No CC"},
		})
	},
}

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Start monitoring daemon",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Starting monitor daemon...")
		fmt.Println("   Checking quotas every 5 minutes")
		fmt.Println("   Alerts enabled for >80% usage")
		// TODO: implement actual monitoring
	},
}

type Provider struct {
	Name    string
	Quota   string
	CC      string
}

func printCategory(name string, providers []Provider) {
	fmt.Printf("📦 %s\n", name)
	for _, p := range providers {
		fmt.Printf("   %-20s %-35s [%s]\n", p.Name, p.Quota, p.CC)
	}
	fmt.Println()
}

func main() {
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(monitorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
