package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)
	until, err := solana.SignatureFromBase58("5zJoMuVeM9KyNGZ5fCkzoerm6CRBcfGcS1GFJ4f9NP1NjniC3JjVoaFiKs8WUEnTmnuyAknULJAzHKaDof1VMpEY")
	if err != nil {
		panic(err)
	}
	out, err := client.GetSignaturesForAddressWithOpts(
		context.TODO(),
		solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"),
		&rpc.GetSignaturesForAddressOpts{
			Until: until,
		},
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}
