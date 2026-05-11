// Package list contains methods to list variables in the variables.json
package list

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/iamkaran/tb-override/internal/config"
	"github.com/iamkaran/tb-override/internal/variables"
	"github.com/spf13/cobra"
)

func ListVariables(ctx context.Context, cfg *config.Config, cmd *cobra.Command) error {
	cssProperties, err := variables.LoadMap(cfg.TBOverride.Dirs.RootDirectory + "/" + cfg.TBOverride.Files.VariablesFilename)

	if err != nil {
		return err
	}
	if listCategories, _ := cmd.Flags().GetBool("list-categories"); listCategories {
		categories := cssProperties.FetchCategories()
		for _, c := range categories {
			_, _ = fmt.Println(c)
		}
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "NAME\tTYPE\tDEFAULT\tDESCRIPTION")

	if category, _ := cmd.Flags().GetString("by-category"); category != "" {
		properties := cssProperties.FetchItems(category)
		for _, c := range properties {
			printRow(w, c.Name, c.Type, c.Default, c.Description)
		}
	} else if listAll, _ := cmd.Flags().GetBool("list-all"); listAll {
		allVariables := cssProperties.FetchVariables()
		for category, vars := range allVariables {
			for _, v := range vars {
				_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
					category,
					v.Name,
					v.Default,
					v.Type,
					v.Description,
				)
			}
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}

func printRow(w *tabwriter.Writer, name, def, typ, desc string) {
	_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", name, typ, def, desc)
}
