package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	//we need to specify the root command hear.
	var rootCmd = &cobra.Command{Use: "cli"}
	//Executing the root command

	//Hello command 
	var helloCmd = &cobra.Command{
		Use: "hello",
		Short: "Print simple greeting on screen",
		Long: ``,
		Run : func(cmd *cobra.Command, args [] string){
			fmt.Println("Hello from PAANDATD ")
		},
	}

	rootCmd.AddCommand(helloCmd)
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
	

}
