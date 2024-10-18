package main

import (
	"fmt"
	"github.com/TiyaAnlite/F-Assests/types"
	"github.com/xlab/tablewriter"
	"strings"
)

func printBasicAsset(a []types.Asset) {
	tb := tablewriter.CreateTable()
	tb.AddHeaders("ID", "Code", "Name", "Type", "Position", "Status")
	for _, asset := range a {
		tb.AddRow(asset.ID, asset.Code, asset.Name, asset.Type, asset.Position.Name, AssetStatusNameMap[asset.Status])
	}
	fmt.Println(tb.Render())
}

func action() {
	for {
		fmt.Print("[action]>>>")
		inputScanner()
	actionSwitch:
		switch buffer {
		case "h":
			fmt.Println("Inbound: i\nOutbound: o\nAbandon: a\nCheck: c\nexit: q")
		case "i", "o", "a":
			t := types.AssetStatusType(strings.ToUpper(buffer))
			posId := ""
			posName := ""
			if t == types.AssetStatusInbound {
			positionSelect:
				for {
					fmt.Print("Enter position(q to exit and l to list position): ")
					inputScanner()
					switch buffer {
					case "q":
						break actionSwitch
					case "l":
						listPosition()
					default:
						// check
						for _, p := range listPosition() {
							if p.ID == buffer {
								posId = p.ID
								posName = fmt.Sprintf("(%s)", p.Name)
							}
						}
						if posId == "" {
							fmt.Println("Error: pos id not found")
							break
						}
						break positionSelect
					}
				}
			}
		actionLoop:
			for {
				fmt.Print(fmt.Sprintf("[%s]%sID or Code: ", AssetStatusNameMap[t], posName))
				inputScanner()
				switch buffer {
				case "q":
					break actionLoop
				default:
					if t != types.AssetStatusInbound {
						// fetch asset
						resp := readResp[types.Asset](client.Get(cfg.Endpoint + "/asset/" + buffer))
						if resp == nil {
							break
						}
						posId = resp.PositionID
						fmt.Println(fmt.Sprintf("Auto find asset position: [%s]%s", resp.PositionID, resp.Position.Name))
					}
					resp := readResp[types.Asset](client.Get(fmt.Sprintf("%s/action/%s/%s?position=%s", cfg.Endpoint, buffer, t, posId)))
					if resp == nil {
						break
					}
					printBasicAsset([]types.Asset{*resp})
				}
			}
		case "c":
		checkLoop:
			for {
				fmt.Print("[查]ID or Code: ")
				inputScanner()
				switch buffer {
				case "q":
					break checkLoop
				default:
					resp := readResp[types.Asset](client.Get(cfg.Endpoint + "/asset/" + buffer))
					if resp == nil {
						break
					}
					printBasicAsset([]types.Asset{*resp})
				}
			}
		case "q":
			return
		}
	}
}
