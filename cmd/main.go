package main

import (
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
	nodeFunc := node.NewKVApplicationNode

	// Create & start node
	rootCmd.AddCommand(cmd.NewRunNodeCmd(nodeFunc))

	cmd := cli.PrepareBaseCmd(rootCmd, "TM", os.ExpandEnv(filepath.Join("./", DefaultBCDir)))


	go func(){
		time.Sleep(5*time.Second)

		core.BroadcastTxCommit(&rpctypes.Context{}, []byte("test=test11111"))
	}()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
