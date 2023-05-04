/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/abibby/bob/bob-cli/util"
	"github.com/abibby/bob/migrate"
	"github.com/abibby/bob/schema"
	"github.com/spf13/cobra"
)

// makeMigrationCmd represents the makeMigration command
var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration",
	Short: "",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := fmt.Sprintf("%s-%s", time.Now().Format("20060102_150405"), strings.ReplaceAll(strings.Join(args, "_"), " ", "_"))

		root, err := util.PackageRoot()
		if err != nil {
			return err
		}

		packageName := "migrations"
		migrationsDir := path.Join(root, packageName)

		err = os.MkdirAll(migrationsDir, 0755)
		if err != nil {
			return err
		}
		src, err := migrate.SrcFile(name, packageName, schema.Table("", func(t *schema.Blueprint) {}), schema.Table("", func(t *schema.Blueprint) {}))
		if err != nil {
			return err
		}

		return os.WriteFile(path.Join(migrationsDir, name+".go"), []byte(src), 0644)
	},
}

func init() {
	rootCmd.AddCommand(makeMigrationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeMigrationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeMigrationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
