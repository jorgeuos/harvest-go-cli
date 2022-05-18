/*
The MIT License (MIT)
Copyright © 2022 Jorge Powers <jorge.powers@gmail.com>
@license https://github.com/jorgeuos/harvest-go-cli/blob/main/LICENSE
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getProjectsCmd represents the getProjects command
var getProjectsCmd = &cobra.Command{
	Use:   "getProjects",
	Short: "List your projects.",
	Long: `List all Harvest Projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getProjects called")
		// We get our Harvestapp credentials from our config file
		account := viper.GetString("HARVEST_ACCOUNT_ID")
		token := viper.GetString("ACCESS_TOKEN")
	
		url := "https://api.harvestapp.com/v2/projects"
		method := "GET"
	
		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, nil)
	
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Harvest-Account-Id", account)
		req.Header.Add("Authorization", token)
	
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
	
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		projects := ProjectList{}
		jsonErr := json.Unmarshal([]byte(string(body)), &projects)
		if jsonErr != nil {
			panic(jsonErr)
		}

		// We only show id and name for now
		for i := 0; i < len(projects.Projects); i++ {
			fmt.Printf("Id: %+v, ", projects.Projects[i].ID)
			fmt.Printf("Name: %+v\n", projects.Projects[i].Name)
		}
		// fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(getProjectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getProjectsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getProjectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ProjectList contains an array Project off projects from Harvestapp
// and a summary of total projects and pagination options.
// id	integer	Unique ID for the project.
// name	string	Unique name for the project.
// code	string	The code associated with the project.
// is_active	boolean	Whether the project is active or archived.
// is_billable	boolean	Whether the project is billable or not.
// is_fixed_fee	boolean	Whether the project is a fixed-fee project or not.
// bill_by	string	The method by which the project is invoiced.
// budget	decimal	The budget in hours for the project when budgeting by time.
// budget_by	string	The method by which the project is budgeted.
// budget_is_monthly	boolean	Option to have the budget reset every month.
// notify_when_over_budget	boolean	Whether Project Managers should be notified when the project goes over budget.
// over_budget_notification_percentage	decimal	Percentage value used to trigger over budget email alerts.
// show_budget_to_all	boolean	Option to show project budget to all employees. Does not apply to Total Project Fee projects.
// created_at	datetime	Date and time the project was created.
// updated_at	datetime	Date and time the project was last updated.
// starts_on	date	Date the project was started.
// ends_on	date	Date the project will end.
// over_budget_notification_date	date	Date of last over budget notification. If none have been sent, this will be null.
// notes	string	Project notes.
// cost_budget	decimal	The monetary budget for the project when budgeting by money.
// cost_budget_include_expenses	boolean	Option for budget of Total Project Fees projects to include tracked expenses.
// hourly_rate	decimal	Rate for projects billed by Project Hourly Rate.
// fee	decimal	The amount you plan to invoice for the project. Only used by fixed-fee projects.
// client	object	An object containing the project’s client id, name, and currency.
type ProjectList struct {
	Projects []struct {
		ID                               int         `json:"id"`
		Name                             string      `json:"name"`
		Code                             string      `json:"code"`
		IsActive                         bool        `json:"is_active"`
		IsBillable                       bool        `json:"is_billable"`
		IsFixedFee                       bool        `json:"is_fixed_fee"`
		BillBy                           string      `json:"bill_by"`
		Budget                           interface{} `json:"budget"`
		BudgetBy                         string      `json:"budget_by"`
		BudgetIsMonthly                  bool        `json:"budget_is_monthly"`
		NotifyWhenOverBudget             bool        `json:"notify_when_over_budget"`
		OverBudgetNotificationPercentage float64     `json:"over_budget_notification_percentage"`
		ShowBudgetToAll                  bool        `json:"show_budget_to_all"`
		CreatedAt                        time.Time   `json:"created_at"`
		UpdatedAt                        time.Time   `json:"updated_at"`
		StartsOn                         interface{} `json:"starts_on"`
		EndsOn                           interface{} `json:"ends_on"`
		OverBudgetNotificationDate       interface{} `json:"over_budget_notification_date"`
		Notes                            string      `json:"notes"`
		CostBudget                       float64     `json:"cost_budget"`
		CostBudgetIncludeExpenses        bool        `json:"cost_budget_include_expenses"`
		HourlyRate                       float64     `json:"hourly_rate"`
		Fee                              interface{} `json:"fee"`
		Client                           struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Currency string `json:"currency"`
		} `json:"client"`
	} `json:"projects"`
	PerPage      int         `json:"per_page"`
	TotalPages   int         `json:"total_pages"`
	TotalEntries int         `json:"total_entries"`
	NextPage     int         `json:"next_page"`
	PreviousPage interface{} `json:"previous_page"`
	Page         int         `json:"page"`
	Links        struct {
		First    string      `json:"first"`
		Next     string      `json:"next"`
		Previous interface{} `json:"previous"`
		Last     string      `json:"last"`
	} `json:"links"`
}