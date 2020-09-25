package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/tangerine-network/go-tangerine/accounts/abi/bind"
	"github.com/tangerine-network/go-tangerine/common"
	"github.com/tangerine-network/go-tangerine/crypto"
	"github.com/tangerine-network/go-tangerine/ethclient"
)

const rpcURL = "https://mainnet-rpc.tangerine-network.io"

var governanceAddress = common.HexToAddress("0x246FcDE58581e2754f215A523C0718C4BFc8041F")
var minStake = new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18))

var client *ethclient.Client
var gov *Governance

func init() {
	var err error

	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		panic(err)
	}

	gov, err = NewGovernance(governanceAddress, client)
	if err != nil {
		panic(err)
	}
}

func networkStatusCmd(c *cli.Context) error {
	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	txOps := &bind.CallOpts{}
	roundHeight, err := gov.RoundHeight(
		txOps, new(big.Int).SetUint64(latestHeader.Round))
	if err != nil {
		return err
	}

	header, err := client.HeaderByNumber(context.Background(), roundHeight)
	if err != nil {
		return err
	}

	elapsed := time.Duration(time.Now().Unix())*time.Second -
		time.Duration(header.Time)*time.Millisecond

	fmt.Printf("Height:\t\t%s\n", latestHeader.Number.String())
	fmt.Printf("Round:\t\t%d\n", latestHeader.Round)
	fmt.Printf("Round Height:\t%s\n", roundHeight.String())
	fmt.Printf("Round Elapsed:\t%s\n", elapsed.String())
	return nil
}

func getNodeStatusCmd(c *cli.Context) error {
	txOps := &bind.CallOpts{}

	nodeLen, err := gov.NodesLength(txOps)
	if err != nil {
		return err
	}

	minStake, err := gov.MinStake(txOps)
	if err != nil {
		return err
	}

	for i := uint64(0); i < nodeLen.Uint64(); i++ {
		n, err := gov.Nodes(txOps, big.NewInt(int64(i)))
		if err != nil {
			return err
		}

		if n.Staked.Cmp(minStake) < 0 && n.Unstaked.Cmp(big.NewInt(0)) < 0 {
			continue
		}

		fmt.Println("*", n.Name)
		fmt.Println("  Address:", n.Owner.Hex())
		fmt.Println("  Email:", n.Email)
		fmt.Println("  Staked:", n.Staked)
		fmt.Println("  Unstaked:", n.Unstaked)
		fmt.Println("  UnstakedAt:", time.Unix(n.UnstakedAt.Int64(), 0))
		fmt.Println("  Fined:", n.Fined)
		fmt.Println("  Public Key:", hex.EncodeToString(n.PublicKey))
		fmt.Println("")
	}
	return nil
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
				Name:   "network-status",
				Action: networkStatusCmd,
			},
			{
				Name:   "node-status",
				Action: getNodeStatusCmd,
			},
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
