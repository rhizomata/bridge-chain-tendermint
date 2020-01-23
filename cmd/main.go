package main

import (
	"fmt"
	"github.com/rhizomata/bridge-chain-tendermint/node"
	rpctypes "github.com/tendermint/tendermint/rpc/lib/types"
	"os"
	"time"
	"path/filepath"

	"github.com/tendermint/tendermint/libs/cli"
	core "github.com/tendermint/tendermint/rpc/core"
	cmd "github.com/rhizomata/bridge-chain-tendermint/cmd/commands"
)

const (
	DefaultBCDir = "chainroot"
)

func main() {
	rootCmd := cmd.RootCmd
	rootCmd.AddCommand(
		cmd.GenValidatorCmd,
		cmd.InitFilesCmd,
		cmd.ProbeUpnpCmd,
		cmd.LiteCmd,
		cmd.ReplayCmd,
		cmd.ReplayConsoleCmd,
		cmd.ResetAllCmd,
		cmd.ResetPrivValidatorCmd,
		cmd.ShowValidatorCmd,
		cmd.TestnetFilesCmd,
		cmd.ShowNodeIDCmd,
		cmd.GenNodeKeyCmd,
		cmd.VersionCmd,
	)

	// NOTE:
	// Users wishing to:
	//	* Use an external signer for their validators
	//	* Supply an in-proc abci app
	//	* Supply a genesis doc file from another source
	//	* Provide their own DB implementation
	// can copy this file and use something other than the
	// DefaultNewNode function
	
	provider := &node.KVApplicationNodeProvider{}
	nodeFunc := provider.NewNode

	// Create & start node
	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeFunc))

	cmd := cli.PrepareBaseCmd(rootCmd, "TM", os.ExpandEnv(filepath.Join("./", DefaultBCDir)))
	
	
	go func(){
		time.Sleep(2*time.Second)
		for i:=0;i<100;i++{
			time.Sleep(20*time.Millisecond)
			stt , _ := core.Status(&rpctypes.Context{})
			core.BroadcastTxSync(&rpctypes.Context{}, []byte(fmt.Sprintf("test%d=%s%d",i,stt.NodeInfo.ID(), i)))
			//core.BroadcastTxCommit(&rpctypes.Context{}, []byte(fmt.Sprintf("test%d=%s%d",i,stt.NodeInfo.ID(), i)))
			
			//if i%5 ==0{
			//	core.BroadcastTxCommit(&rpctypes.Context{}, []byte(fmt.Sprintf("Commit%d=%s%d",i,stt.NodeInfo.ID(), i)))
			//}
			
			if i%5 ==0{
				time.Sleep(200*time.Millisecond)
			}
		}
	}()
	
	
	go func(){
		time.Sleep(2*time.Second)
		for i:=0;i<100;i++{
			time.Sleep(30*time.Millisecond)
			stt , _ := core.Status(&rpctypes.Context{})
			core.BroadcastTxSync(&rpctypes.Context{}, []byte(fmt.Sprintf("stest%d=%s%ds",i,stt.NodeInfo.ID(), i)))
			//core.BroadcastTxCommit(&rpctypes.Context{}, []byte(fmt.Sprintf("stest%d=%s%ds",i,stt.NodeInfo.ID(), i)))
			
			if i%7 ==0{
				time.Sleep(300*time.Millisecond)
			}
		}
		
	}()
	
	go func(){
		time.Sleep(5*time.Second)
		for i:=0;i<100;i++ {
			time.Sleep(30*time.Millisecond)
			iterator, _ := provider.App.DB.Iterator([]byte("kvPairKey:stest8"),[]byte("kvPairKey:stest999"))
			for iterator.Valid() {
				
				fmt.Println(" ^^ DB.Iterator: key=", string(iterator.Key()), ", value=", string(iterator.Value()))
				iterator.Next()
			}
			iterator.Close()
			if i%7 ==0{
				time.Sleep(300*time.Millisecond)
			}
		}
	}()
	
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
