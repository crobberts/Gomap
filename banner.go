package main

import "fmt"

//PrintBanner prints welcome banner
func PrintBanner() {
	banner := `
	#####                              
	#     #  ####  #    #   ##   #####  
	#       #    # ##  ##  #  #  #    # 
	#  #### #    # # ## # #    # #    # 
	#     # #    # #    # ###### #####  
	#     # #    # #    # #    # #      
	 #####   ####  #    # #    # #
	`

	fmt.Println(banner)
}
