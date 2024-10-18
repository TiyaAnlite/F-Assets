package main

import (
	"fmt"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/xlab/tablewriter"
)

func printPos(pos []types.Position) {
	tb := tablewriter.CreateTable()
	tb.AddHeaders("ID", "Name")
	for _, row := range pos {
		tb.AddRow(row.ID, row.Name)
	}
	fmt.Println(tb.Render())
}

func listPosition() []types.Position {
	resp := readResp[[]types.Position](client.Get(cfg.Endpoint + "/position"))
	if resp == nil {
		return nil
	}
	printPos(*resp)
	return *resp
}

func position() {
	for {
		fmt.Print("[position]>>> ")
		inputScanner()
		switch buffer {
		case "h":
			fmt.Print("list: l\nadd: a\nquit: q")
		case "l":
			listPosition()
		case "a":
			fmt.Print("New position name: ")
			inputScanner()
			resp := readResp[types.Position](client.Post(
				cfg.Endpoint+"/position",
				JsonContentType,
				must2Reader(&types.PositionOptRequest{Name: buffer})))
			if resp == nil {
				break
			}
			printPos([]types.Position{*resp})
		case "q":
			return
		}
	}
}
