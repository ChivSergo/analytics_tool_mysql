package main

import (
	"analytics_tool/db"
	"analytics_tool/handlers"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "analytics_tool",
		Usage: "A tool for analyzing employee timesheets",
		Commands: []*cli.Command{
			{
				Name:  "import",
				Usage: "Import data from CSV files",
				Subcommands: []*cli.Command{
					{
						Name:  "positions",
						Usage: "Import positions from CSV file",
						Action: func(c *cli.Context) error {
							handlers.ImportPositions(c.Args().First())
							return nil
						},
					},
					{
						Name:  "employees",
						Usage: "Import employees from CSV file",
						Action: func(c *cli.Context) error {
							handlers.ImportEmployees(c.Args().First())
							return nil
						},
					},
					{
						Name:  "timesheet",
						Usage: "Import timesheet from CSV file",
						Action: func(c *cli.Context) error {
							handlers.ImportTimesheet(c.Args().First())
							return nil
						},
					},
				},
			},
			{
				Name:  "list",
				Usage: "List data",
				Subcommands: []*cli.Command{
					{
						Name:  "employee",
						Usage: "List employees",
						Action: func(c *cli.Context) error {
							handlers.ListEmployees()
							return nil
						},
					},
				},
			},
			{
				Name:  "get",
				Usage: "Get data",
				Action: func(c *cli.Context) error {
					handlers.GetEmployeeTimesheet(c.Args().First())
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove data",
				Action: func(c *cli.Context) error {
					handlers.RemoveEmployee(c.Args().First())
					return nil
				},
			},
			{
				Name:  "report",
				Usage: "Generate reports",
				Subcommands: []*cli.Command{
					{
						Name:  "top5longTasks",
						Usage: "Top 5 longest tasks",
						Action: func(c *cli.Context) error {
							handlers.ReportTop5LongTasks()
							return nil
						},
					},
					{
						Name:  "top5costTasks",
						Usage: "Top 5 costliest tasks",
						Action: func(c *cli.Context) error {
							handlers.ReportTop5CostTasks()
							return nil
						},
					},
					{
						Name:  "top5employees",
						Usage: "Top 5 employees by total time worked",
						Action: func(c *cli.Context) error {
							handlers.ReportTop5Employees()
							return nil
						},
					},
				},
			},
		},
	}

	db.Init()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
