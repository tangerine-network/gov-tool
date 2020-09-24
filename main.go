package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const rpcURL = "https://mainnet-rpc.tangerine-network.io"

var governanceAddress = common.HexToAddress("0x246FcDE58581e2754f215A523C0718C4BFc8041F")
var minStake = new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18))

var conn *ethclient.Client
var gov *Governance

func init() {
	var err error

	conn, err = ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}

	gov, err = NewGovernance(governanceAddress, conn)
	if err != nil {
		panic(err)
	}
}

func registerCmd(c *cli.Context) error {
	if c.NArg() != 5 {
		return fmt.Errorf("Arguments: node.key Name Email Location Url")
	}

	args := c.Args()

	privKey, err := crypto.LoadECDSA(args.Get(0))
	if err != nil {
		return err
	}

	pk := crypto.FromECDSAPub(&privKey.PublicKey)

	auth := bind.NewKeyedTransactor(privKey)
	tx, err := gov.Register(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  minStake,
	}, pk, args.Get(1), args.Get(2), args.Get(3), args.Get(4))

	if err != nil {
		return err
	}

	fmt.Println("TxHash:", tx.Hash().String())
	return nil
}

func payFineCmd(c *cli.Context) error {
	if c.NArg() != 3 {
		return fmt.Errorf("Arguments: node.key NodeAddress Amount")
	}

	args := c.Args()

	privKey, err := crypto.LoadECDSA(args.Get(0))
	if err != nil {
		return err
	}

	nodeAddr := common.HexToAddress(args.Get(1))

	amount, err := strconv.ParseInt(args.Get(2), 10, 64)
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privKey)
	tx, err := gov.PayFine(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  new(big.Int).Mul(big.NewInt(amount), big.NewInt(1e18)),
	}, nodeAddr)

	if err != nil {
		return err
	}

	fmt.Println("TxHash:", tx.Hash().String())

	return nil
}

func unstakeCmd(c *cli.Context) error {
	if c.NArg() < 1 {
		return fmt.Errorf("Arguments: node.key [Amount]")
	}

	args := c.Args()

	privKey, err := crypto.LoadECDSA(args.Get(0))
	if err != nil {
		return err
	}

	var amount *big.Int

	if args.Len() == 1 {
		amount = minStake
	} else {
		amountInt, err := strconv.ParseInt(args.Get(1), 10, 64)
		if err != nil {
			return err
		}
		amount = new(big.Int).Mul(big.NewInt(amountInt), big.NewInt(1e18))
	}

	auth := bind.NewKeyedTransactor(privKey)
	tx, err := gov.Unstake(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, amount)

	if err != nil {
		return err
	}

	fmt.Println("TxHash:", tx.Hash().String())

	return nil
}

func withdrawCmd(c *cli.Context) error {
	if c.NArg() != 1 {
		return fmt.Errorf("Arguments: node.key")
	}

	args := c.Args()

	privKey, err := crypto.LoadECDSA(args.Get(0))
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privKey)
	tx, err := gov.Withdraw(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	})

	if err != nil {
		return err
	}

	fmt.Println("TxHash:", tx.Hash().String())
	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "register",
				Usage:  "node.key Name Email Location URL",
				Action: registerCmd,
			},
			{
				Name:   "payfine",
				Usage:  "node.key NodeAddress Amount",
				Action: payFineCmd,
			},
			{
				Name:   "unstake",
				Usage:  "node.key Amount",
				Action: unstakeCmd,
			},
			{
				Name:   "withdraw",
				Usage:  "node.key",
				Action: withdrawCmd,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
