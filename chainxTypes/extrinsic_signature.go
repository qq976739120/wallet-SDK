// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chainxTypes

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type ExtrinsicSignatureV3 struct {
	Signer    types.Address
	Signature types.Signature
	Era       types.ExtrinsicEra // extra via system::CheckEra
	Nonce     types.UCompact     // extra via system::CheckNonce (Compact<Index> where Index is u32))
	Tip       types.UCompact     // extra via balances::TakeFees (Compact<Balance> where Balance is u128))
}

type ExtrinsicSignatureV4 struct {
	Signer    types.Address
	Signature types.MultiSignature
	Era       types.ExtrinsicEra // extra via system::CheckEra
	Nonce     types.UCompact     // extra via system::CheckNonce (Compact<Index> where Index is u32))
	Tip       types.UCompact     // extra via balances::TakeFees (Compact<Balance> where Balance is u128))
}

type SignatureOptions struct {
	Era                types.ExtrinsicEra // extra via system::CheckEra
	Nonce              types.UCompact     // extra via system::CheckNonce (Compact<Index> where Index is u32)
	Tip                types.UCompact     // extra via balances::TakeFees (Compact<Balance> where Balance is u128)
	SpecVersion        types.U32          // additional via system::CheckSpecVersion
	GenesisHash        types.Hash         // additional via system::CheckGenesis
	BlockHash          types.Hash         // additional via system::CheckEra
	TransactionVersion types.U32          // additional via system::CheckTxVersion
}
