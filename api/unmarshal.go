package main

import (
	"fmt"

	types "github.com/vegaprotocol/api/go/generated/code.vegaprotocol.io/vega/proto"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

func unmarshalCommand(payload []byte, cmd Command) (proto.Message, error) {
	// first is that a signed or unsigned command?
	switch cmd {
	case SubmitOrderCommand:
		orderSubmission := &types.OrderSubmission{}
		err := proto.Unmarshal(payload, orderSubmission)
		if err != nil {
			return nil, err
		}
		return orderSubmission, nil
	case CancelOrderCommand:
		order := &types.OrderCancellation{}
		err := proto.Unmarshal(payload, order)
		if err != nil {
			return nil, err
		}
		return order, nil
	case AmendOrderCommand:
		amendment := &types.OrderAmendment{}
		err := proto.Unmarshal(payload, amendment)
		if err != nil {
			return nil, errors.Wrap(err, "error decoding order to proto")
		}
		return amendment, nil
	case WithdrawCommand:
		w := &types.WithdrawSubmission{}
		err := proto.Unmarshal(payload, w)
		if err != nil {
			return nil, errors.Wrap(err, "error decoding order to proto")
		}
		return w, nil
	case ProposeCommand:
		proposalSubmission := &types.Proposal{}
		err := proto.Unmarshal(payload, proposalSubmission)
		if err != nil {
			return nil, err
		}
		return proposalSubmission, nil
	case VoteCommand:
		voteSubmission := &types.Vote{}
		err := proto.Unmarshal(payload, voteSubmission)
		if err != nil {
			return nil, err
		}
		return voteSubmission, nil
	case RegisterNodeCommand:
		cmd := &types.NodeRegistration{}
		err := proto.Unmarshal(payload, cmd)
		if err != nil {
			return nil, err
		}
		return cmd, nil
	case NodeVoteCommand:
		vote := &types.NodeVote{}
		if err := proto.Unmarshal(payload, vote); err != nil {
			return nil, err
		}
		return vote, nil
	case NodeSignatureCommand:
		sig := &types.NodeSignature{}
		if err := proto.Unmarshal(payload, sig); err != nil {
			return nil, err
		}
		return sig, nil
	case ChainEventCommand:
		evt := &types.ChainEvent{}
		if err := proto.Unmarshal(payload, evt); err != nil {
			return nil, err
		}
		return evt, nil
	default:
		return nil, fmt.Errorf("unsupported command: %s", cmd)
	}
}
