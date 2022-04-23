package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/soupy-finance/noodle/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdatePrices_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdatePrices
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePrices{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdatePrices{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
