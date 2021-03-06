package cmd

import (
	"fmt"
	"github.com/lusis/go-rundeck/src/rundeck.v12"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ListNodes(projectid string) {
	client := rundeck.NewClientFromEnv()
	data, err := client.ListNodes(projectid)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "Tags", "Hostname", "OsArch", "OsName", "OSVersion"})
	for _, d := range data.Nodes {
		table.Append([]string{d.Name, d.Description, d.Tags, d.Hostname, d.OsArch, d.OsName, d.OsVersion})
	}
	table.Render()
}
