package main

import (
	"bufio"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"io"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"

	emintCrypto "github.com/cosmos/ethermint/crypto"
	tmcrypto "github.com/tendermint/tendermint/crypto"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagDryRun = "dry-run"
)

// keyCommands registers a sub-tree of commands to interact with
// local private key storage.
func keyCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Add or view local private keys",
		Long: `Keys allows you to manage your local keystore for tendermint.

    These keys may be in any format supported by go-crypto and can be
    used by light-clients, full nodes, or any other application that
    needs to sign with a private key.`,
	}
	addCmd := clientkeys.AddKeyCommand()
	addCmd.RunE = runAddCmd
	cmd.AddCommand(
		clientkeys.MnemonicKeyCommand(),
		addCmd,
		clientkeys.ExportKeyCommand(),
		clientkeys.ImportKeyCommand(),
		clientkeys.ListKeysCmd(),
		ShowKeysCmd(),
		flags.LineBreak,
		clientkeys.DeleteKeyCommand(),
		clientkeys.UpdateKeyCommand(),
		clientkeys.ParseKeyStringCommand(),
		clientkeys.MigrateCommand(),
		flags.LineBreak,
		exportEthKeyCommand(),
	)
	return cmd
}

func getKeybase(dryrun bool, buf io.Reader) (keys.Keybase, error) {
	fmt.Println("DRY RUN: ", dryrun)
	fmt.Println("SERVICE NAME: ", types.KeyringServiceName())
	fmt.Println("BACKEND: ", viper.GetString(flags.FlagKeyringBackend))
	fmt.Println("HOME: ", flags.FlagHome)
	if dryrun {
		return keys.NewInMemory(keys.WithKeygenFunc(ethermintKeygenFunc)), nil
	}

	return keys.NewKeyring(types.KeyringServiceName(),viper.GetString(flags.FlagKeyringBackend), flags.FlagHome, buf)
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	fmt.Println("***** run add command ******")
	kb, err := getKeybase(viper.GetBool(flagDryRun), inBuf)
	if err != nil {
		return err
	}
	fmt.Println("**************")
	return clientkeys.RunAddCmd(cmd, args, kb, inBuf)
}

//TODO check it
func ethermintKeygenFunc(bz []byte, algo keys.SigningAlgo) (tmcrypto.PrivKey, error) {
	return emintCrypto.PrivKeySecp256k1(bz[:]), nil
}
