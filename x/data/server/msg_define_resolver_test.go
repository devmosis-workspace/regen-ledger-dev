//nolint:revive,stylecheck
package server

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/data/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/data/v2"
)

type defineResolverSuite struct {
	*baseSuite
	alice sdk.AccAddress
	bob   sdk.AccAddress
	err   error
}

func TestDefineResolver(t *testing.T) {
	gocuke.NewRunner(t, &defineResolverSuite{}).Path("./features/msg_define_resolver.feature").Run()
}

func (s *defineResolverSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
}

func (s *defineResolverSuite) AliceHasDefinedAResolverWithUrl(a string) {
	err := s.server.stateStore.ResolverTable().Insert(s.ctx, &api.Resolver{
		Url:     a,
		Manager: s.alice,
	})
	require.NoError(s.t, err)
}

func (s *defineResolverSuite) AliceAttemptsToDefineAResolverWithUrl(a string) {
	_, s.err = s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.alice.String(),
		ResolverUrl: a,
	})
}

func (s *defineResolverSuite) BobAttemptsToDefineAResolverWithUrl(a string) {
	_, s.err = s.server.DefineResolver(s.ctx, &data.MsgDefineResolver{
		Manager:     s.bob.String(),
		ResolverUrl: a,
	})
}

func (s *defineResolverSuite) ExpectTheResolverWithIdAndUrlAndManagerBob(a string, b string) {
	id, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	resolver, err := s.server.stateStore.ResolverTable().Get(s.ctx, id)
	require.NoError(s.t, err)
	require.Equal(s.t, b, resolver.Url)
	require.Equal(s.t, s.bob.Bytes(), resolver.Manager)
}

func (s *defineResolverSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *defineResolverSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event data.EventDefineResolver
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
