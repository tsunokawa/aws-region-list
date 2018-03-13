package main

import (
	"os"
	"github.com/olekukonko/tablewriter"
)

func main() {
  data := [][]string{
    []string{"us-east-1", "US East (N. Virginia)"},
    []string{"us-east-2", "US East (Ohio)"},
    []string{"us-west-1", "US West (N. California)"},
    []string{"us-west-2", "US West (Oregon)"},
    []string{"ca-central-1", "Canada (Central)"},
    []string{"eu-west-1", "EU (Ireland)"},
    []string{"eu-central-1", "EU (Frankfurt)"},
    []string{"eu-west-2", "EU (London)"},
    []string{"ap-northeast-1", "Asia Pacific (Tokyo)"},
    []string{"ap-northeast-2", "Asia Pacific (Seoul)"},
    []string{"ap-northeast-3", "Asia Pacific (Osaka-Local)"},
    []string{"ap-southeast-1", "Asia Pacific (Singapore)"},
    []string{"ap-southeast-2", "Asia Pacific (Sydney)"},
    []string{"ap-south-1", "Asia Pacific (Mumbai)"},
    []string{"sa-east-1", "South America (São Paulo)"},
  }

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"CODE", "NAME"})

  for _, v := range data {
      table.Append(v)
  }

  table.Render()
}
