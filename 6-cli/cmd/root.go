/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"os"

	"github.com/janainamai/learning-go/6-cli/database"
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	return db
}

func GetCategoryDb(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

var rootCmd = &cobra.Command{
	Use:   "6-cli",
	Short: "Learning Go Expert",
	Long:  `Starting to use cobra-cli`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
