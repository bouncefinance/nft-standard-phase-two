package solidity

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
	"math/big"
)

func ABIDecodeBigInt(data []byte) (*big.Int, error) {
	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, errors.Wrap(err, "abi.NewType error")
	}
	args := abi.Arguments{
		{
			Type: uint256Ty,
		},
	}
	unpacked, err := args.Unpack(data)
	if err != nil {
		return nil, errors.Wrap(err, "args.Unpack error")
	}
	var (
		s  *big.Int
		ok bool
	)
	if len(unpacked) != 0 {
		s, ok = unpacked[0].(*big.Int)
		if !ok {
			return nil, errors.New("arg is not *big.Int.")
		}
	}
	return s, nil
}

func ABIDecodeBool(data []byte) (bool, error) {
	boolTy, err := abi.NewType("bool", "", nil)
	if err != nil {
		return false, errors.Wrap(err, "abi.NewType error")
	}
	args := abi.Arguments{
		{
			Type: boolTy,
		},
	}
	unpacked, err := args.Unpack(data)
	if err != nil {
		return false, errors.Wrap(err, "args.Unpack error")
	}
	var (
		s  bool
		ok bool
	)
	if len(unpacked) != 0 {
		s, ok = unpacked[0].(bool)
		if !ok {
			return false, errors.New("arg is not bool.")
		}
	}
	return s, nil
}

func ABIDecodeString(data []byte) (string, error) {
	stringTy, err := abi.NewType("string", "", nil)
	if err != nil {
		return "", errors.Wrap(err, "abi.NewType error")
	}
	args := abi.Arguments{
		{
			Type: stringTy,
		},
	}

	unpacked, err := args.Unpack(data)
	if err != nil {
		return "", errors.Wrap(err, "args.Unpack error")
	}
	var (
		s  string
		ok bool
	)
	if len(unpacked) != 0 {
		s, ok = unpacked[0].(string)
		if !ok {
			return "", errors.New("arg is not string.")
		}
	}

	return s, nil
}