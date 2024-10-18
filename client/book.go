package main

import (
	"fmt"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/xlab/tablewriter"
)

var (
	AssetStatusNameMap = map[types.AssetStatusType]string{
		types.AssetStatusOutbound: "出",
		types.AssetStatusInbound:  "在",
		types.AssetStatusAbandon:  "销",
	}
)

func printBook(book []types.Book) {
	tb := tablewriter.CreateTable()
	tb.AddHeaders("ID", "Code", "Name", "Position", "Status")
	for _, row := range book {
		tb.AddRow(row.Asset.ID, row.Asset.Code, row.Asset.Name, row.Asset.Position.Name, AssetStatusNameMap[row.Asset.Status])
	}
	fmt.Println(tb.Render())
}

func book() {
	for {
		fmt.Print("[book]>>> ")
		inputScanner()
		switch buffer {
		case "h":
			fmt.Print("list: l\nadd: a\nquit: q")
		case "l":
			resp := readResp[[]types.Book](client.Get(cfg.Endpoint + "/asset?type=BOOK"))
			if resp == nil {
				break
			}
			printBook(*resp)
		case "a":
			fmt.Println("adding book via fast mode")
			fmt.Print("New book name: ")
			inputScanner()
			name := buffer
			fmt.Print("New book code: ")
			inputScanner()
			code := buffer
			resp := readResp[types.Book](client.Post(
				cfg.Endpoint+"/asset?type=BOOK",
				JsonContentType,
				must2Reader(&types.BookOptRequest{
					ItemOptRequest: types.ItemOptRequest{
						Name: name,
						Code: code,
					},
					Author:         " ",
					Publisher:      " ",
					Specifications: " ",
					Language:       " ",
				})))
			if resp == nil {
				break
			}
			printBook([]types.Book{*resp})
		case "q":
			return
		}
	}
}
