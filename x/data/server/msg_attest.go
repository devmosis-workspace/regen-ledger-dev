package server

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

// Attest allows for digital signing of an arbitrary piece of data on the blockchain.
func (s serverImpl) Attest(ctx context.Context, request *data.MsgAttest) (*data.MsgAttestResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	timestamp := timestamppb.New(sdkCtx.BlockTime())

	iris := make([]string, 0) // only the IRIs for new attestations

	for _, ch := range request.ContentHashes {
		iri, id, _, err := s.anchorAndGetIRI(ctx, ch)
		if err != nil {
			return nil, err
		}

		addr, err := sdk.AccAddressFromBech32(request.Attestor)
		if err != nil {
			return nil, err
		}

		found, err := s.stateStore.DataAttestorTable().Has(ctx, id, addr)
		if err != nil {
			return nil, err
		} else if found {
			// an attestor attesting to the same piece of date is a no-op
			continue
		}

		err = s.stateStore.DataAttestorTable().Insert(ctx, &api.DataAttestor{
			Id:        id,
			Attestor:  addr,
			Timestamp: timestamp,
		})
		if err != nil {
			return nil, err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&data.EventAttest{
			Iri:      iri,
			Attestor: request.Attestor,
		})
		if err != nil {
			return nil, err
		}

		iris = append(iris, iri)

		sdkCtx.GasMeter().ConsumeGas(data.GasCostPerIteration, "data/Attest content hash iteration")
	}

	return &data.MsgAttestResponse{
		Iris:      iris,
		Timestamp: types.ProtobufToGogoTimestamp(timestamp),
	}, nil
}
