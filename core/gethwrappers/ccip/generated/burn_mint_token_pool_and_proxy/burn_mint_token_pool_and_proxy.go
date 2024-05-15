// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_token_pool_and_proxy

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type PoolLockOrBurnInV1 struct {
	Receiver            []byte
	RemoteChainSelector uint64
	OriginalSender      common.Address
	Amount              *big.Int
}

type PoolLockOrBurnOutV1 struct {
	DestPoolAddress []byte
	DestPoolData    []byte
}

type PoolReleaseOrMintInV1 struct {
	OriginalSender      []byte
	RemoteChainSelector uint64
	Receiver            common.Address
	Amount              *big.Int
	SourcePoolAddress   []byte
	SourcePoolData      []byte
	OffchainTokenData   []byte
}

type PoolReleaseOrMintOutV1 struct {
	LocalToken        common.Address
	DestinationAmount *big.Int
}

type RateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type RateLimiterTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

type TokenPoolChainUpdate struct {
	RemoteChainSelector       uint64
	Allowed                   bool
	RemotePoolAddress         []byte
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
}

var BurnMintTokenPoolAndProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBurnMintERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"allowlist\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"AggregateValueMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"AggregateValueRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowListNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketOverfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotARampOnRouter\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CursedByRMN\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"DisabledNonZeroRateLimit\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"rateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"InvalidRatelimitRate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"}],\"name\":\"InvalidSourcePoolAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"NonExistentChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitMustBeDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenMaxCapacityExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minWaitInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRateLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"AllowListRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"name\":\"ChainConfigured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"ChainRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structRateLimiter.Config\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPoolPriorTo1_5\",\"name\":\"oldPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPoolPriorTo1_5\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"LegacyPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"previousPoolAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"RemotePoolSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"RouterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"TokensConsumed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"adds\",\"type\":\"address[]\"}],\"name\":\"applyAllowListUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundRateLimiterConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundRateLimiterConfig\",\"type\":\"tuple\"}],\"internalType\":\"structTokenPool.ChainUpdate[]\",\"name\":\"chains\",\"type\":\"tuple[]\"}],\"name\":\"applyChainUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentInboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getCurrentOutboundRateLimiterState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"tokens\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdated\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.TokenBucket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"getOnRamp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"onRampAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"getRemotePool\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRmnProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rmnProxy\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sourceChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"offRamp\",\"type\":\"address\"}],\"name\":\"isOffRamp\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"originalSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.LockOrBurnInV1\",\"name\":\"lockOrBurnIn\",\"type\":\"tuple\"}],\"name\":\"lockOrBurn\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"destPoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"destPoolData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.LockOrBurnOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"originalSender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sourcePoolData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainTokenData\",\"type\":\"bytes\"}],\"internalType\":\"structPool.ReleaseOrMintInV1\",\"name\":\"releaseOrMintIn\",\"type\":\"tuple\"}],\"name\":\"releaseOrMint\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"destinationAmount\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ReleaseOrMintOutV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"outboundConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"capacity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rate\",\"type\":\"uint128\"}],\"internalType\":\"structRateLimiter.Config\",\"name\":\"inboundConfig\",\"type\":\"tuple\"}],\"name\":\"setChainRateLimiterConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolPriorTo1_5\",\"name\":\"prevPool\",\"type\":\"address\"}],\"name\":\"setPreviousPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainSelector\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"remotePoolAddress\",\"type\":\"bytes\"}],\"name\":\"setRemotePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200466038038062004660833981016040819052620000349162000541565b83838383838383833380600081620000935760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000c657620000c6816200016b565b5050506001600160a01b0384161580620000e757506001600160a01b038116155b1562000106576040516342bcdf7f60e11b815260040160405180910390fd5b6001600160a01b0384811660805282811660a052600480546001600160a01b031916918316919091179055825115801560c052620001595760408051600081526020810190915262000159908462000216565b5050505050505050505050506200069f565b336001600160a01b03821603620001c55760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200008a565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60c05162000237576040516335f4a7b360e01b815260040160405180910390fd5b60005b8251811015620002c25760008382815181106200025b576200025b62000651565b602090810291909101015190506200027560028262000373565b15620002b8576040516001600160a01b03821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b506001016200023a565b5060005b81518110156200036e576000828281518110620002e757620002e762000651565b6020026020010151905060006001600160a01b0316816001600160a01b03160362000313575062000365565b6200032060028262000393565b1562000363576040516001600160a01b03821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101620002c6565b505050565b60006200038a836001600160a01b038416620003aa565b90505b92915050565b60006200038a836001600160a01b038416620004ae565b60008181526001830160205260408120548015620004a3576000620003d160018362000667565b8554909150600090620003e79060019062000667565b9050818114620004535760008660000182815481106200040b576200040b62000651565b906000526020600020015490508087600001848154811062000431576200043162000651565b6000918252602080832090910192909255918252600188019052604090208390555b855486908062000467576200046762000689565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506200038d565b60009150506200038d565b6000818152600183016020526040812054620004f7575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556200038d565b5060006200038d565b6001600160a01b03811681146200051657600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b80516200053c8162000500565b919050565b600080600080608085870312156200055857600080fd5b8451620005658162000500565b602086810151919550906001600160401b03808211156200058557600080fd5b818801915088601f8301126200059a57600080fd5b815181811115620005af57620005af62000519565b8060051b604051601f19603f83011681018181108582111715620005d757620005d762000519565b60405291825284820192508381018501918b831115620005f657600080fd5b938501935b828510156200061f576200060f856200052f565b84529385019392850192620005fb565b80985050505050505062000636604086016200052f565b915062000646606086016200052f565b905092959194509250565b634e487b7160e01b600052603260045260246000fd5b818103818111156200038d57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b60805160a05160c051613f4b620007156000396000818161045501528181611b9a0152611d5001526000818161042f015281816106ac0152610d130152600081816101fa015281816108260152818161098001528181610e6601528181611a5301528181611f400152611f870152613f4b6000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80639766b932116100ee578063c4bffe2b11610097578063dc0bd97111610071578063dc0bd9711461042d578063e0351e1314610453578063f2fde38b14610479578063f6e2145e1461048c57600080fd5b8063c4bffe2b146103f2578063c75eea9c14610407578063cf7401f31461041a57600080fd5b8063af58d59f116100c8578063af58d59f14610352578063b0f479a1146103c1578063c0d78655146103df57600080fd5b80639766b93214610317578063a7cd63b71461032a578063a8d87a3b1461033f57600080fd5b806378a010b21161015057806383826b2b1161012a57806383826b2b146102d35780638926f54f146102e65780638da5cb5b146102f957600080fd5b806378a010b21461029857806379ba5097146102ab5780637a5c972d146102b357600080fd5b806321df0da71161018157806321df0da7146101f85780634059f55b1461023f57806354c8a4f31461028357600080fd5b806301ffc9a7146101a85780630a2fd493146101d0578063181f5a77146101f0575b600080fd5b6101bb6101b6366004612f52565b61049f565b60405190151581526020015b60405180910390f35b6101e36101de366004612fb1565b610584565b6040516101c7919061303a565b6101e3610634565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c7565b61025261024d36600461304d565b610650565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260209283015192810192909252016101c7565b6102966102913660046130d4565b6109cb565b005b6102966102a6366004613140565b610a46565b610296610bba565b6102c66102c13660046131c3565b610cb7565b6040516101c791906131fe565b6101bb6102e136600461328b565b610f63565b6101bb6102f4366004612fb1565b611030565b60005473ffffffffffffffffffffffffffffffffffffffff1661021a565b6102966103253660046132c2565b611047565b6103326110d6565b6040516101c791906132df565b61021a61034d366004612fb1565b503090565b610365610360366004612fb1565b6110e7565b6040516101c7919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b60045473ffffffffffffffffffffffffffffffffffffffff1661021a565b6102966103ed3660046132c2565b6111bc565b6103fa611290565b6040516101c79190613339565b610365610415366004612fb1565b611348565b6102966104283660046134c1565b61141a565b7f000000000000000000000000000000000000000000000000000000000000000061021a565b7f00000000000000000000000000000000000000000000000000000000000000006101bb565b6102966104873660046132c2565b611432565b61029661049a366004613506565b611446565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167faff2afbf00000000000000000000000000000000000000000000000000000000148061053257507fffffffff0000000000000000000000000000000000000000000000000000000082167fb323973900000000000000000000000000000000000000000000000000000000145b8061057e57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b67ffffffffffffffff811660009081526007602052604090206004018054606091906105af90613548565b80601f01602080910402602001604051908101604052809291908181526020018280546105db90613548565b80156106285780601f106105fd57610100808354040283529160200191610628565b820191906000526020600020905b81548152906001019060200180831161060b57829003601f168201915b50505050509050919050565b604051806060016040528060238152602001613f1c6023913981565b60408051808201909152600080825260208201526106746040830160208401612fb1565b6040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015610708573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072c9190613595565b15610763576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61077b6107766040850160208601612fb1565b6118b5565b6107d561078e6040850160208601612fb1565b61079b60808601866135b2565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506119db92505050565b6107f26107e86040850160208601612fb1565b8460600135611a30565b60085473ffffffffffffffffffffffffffffffffffffffff166108e85773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166340c10f1961085b60608601604087016132c2565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260608601356024820152604401600060405180830381600087803b1580156108cb57600080fd5b505af11580156108df573d6000803e3d6000fd5b505050506108f9565b6108f96108f4846136b3565b611a77565b61090960608401604085016132c2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0856060013560405161096b91815260200190565b60405180910390a360405180604001604052807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168152602001846060013581525091505b50919050565b6109d3611b15565b610a4084848080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808802828101820190935287825290935087925086918291850190849080828437600092019190915250611b9892505050565b50505050565b610a4e611b15565b610a5783611030565b610a9e576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff841660048201526024015b60405180910390fd5b67ffffffffffffffff831660009081526007602052604081206004018054610ac590613548565b80601f0160208091040260200160405190810160405280929190818152602001828054610af190613548565b8015610b3e5780601f10610b1357610100808354040283529160200191610b3e565b820191906000526020600020905b815481529060010190602001808311610b2157829003601f168201915b5050505067ffffffffffffffff8616600090815260076020526040902091925050600401610b6d8385836137e6565b508367ffffffffffffffff167fdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf828585604051610bac93929190613900565b60405180910390a250505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610c3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610a95565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6040805180820190915260608082526020820152610cdb6040830160208401612fb1565b6040517f58babe3300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff821660048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906358babe3390602401602060405180830381865afa158015610d6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d939190613595565b15610dca576040517f53ad11d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610de2610ddd60608501604086016132c2565b611d4e565b610dfa610df56040850160208601612fb1565b611dd2565b610e17610e0d6040850160208601612fb1565b8460600135611f20565b60085473ffffffffffffffffffffffffffffffffffffffff16610edc576040517f42966c68000000000000000000000000000000000000000000000000000000008152606084013560048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906342966c6890602401600060405180830381600087803b158015610ebf57600080fd5b505af1158015610ed3573d6000803e3d6000fd5b50505050610eed565b610eed610ee884613964565b611f64565b6040516060840135815233907f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df79060200160405180910390a26040518060400160405280610f478560200160208101906101de9190612fb1565b8152604080516020808201909252600081529101529392505050565b600073ffffffffffffffffffffffffffffffffffffffff82163014806110295750600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff86169281019290925273ffffffffffffffffffffffffffffffffffffffff848116602484015216906383826b2b90604401602060405180830381865afa158015611005573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110299190613595565b9392505050565b600061057e600567ffffffffffffffff841661207e565b61104f611b15565b6008805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f81accd0a7023865eaa51b3399dd0eafc488bf3ba238402911e1659cfe860f22891015b60405180910390a15050565b60606110e26002612096565b905090565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845260028201546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600390910154808416606083015291909104909116608082015261057e906120a3565b6111c4611b15565b73ffffffffffffffffffffffffffffffffffffffff8116611211576040517f8579befe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f168491016110ca565b6060600061129e6005612096565b90506000815167ffffffffffffffff8111156112bc576112bc61337b565b6040519080825280602002602001820160405280156112e5578160200160208202803683370190505b50905060005b825181101561134157828181518110611306576113066139f4565b6020026020010151828281518110611320576113206139f4565b67ffffffffffffffff909216602092830291909101909101526001016112eb565b5092915050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915267ffffffffffffffff8216600090815260076020908152604091829020825160a08101845281546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000080830463ffffffff16958401959095527401000000000000000000000000000000000000000090910460ff16151594820194909452600190910154808416606083015291909104909116608082015261057e906120a3565b611422611b15565b61142d838383612155565b505050565b61143a611b15565b6114438161223f565b50565b61144e611b15565b60005b8181101561142d57600083838381811061146d5761146d6139f4565b905060200281019061147f9190613a23565b61148890613a61565b905061149d8160600151826020015115612334565b6114b08160800151826020015115612334565b80602001511561177b5780516114d29060059067ffffffffffffffff1661246d565b6115175780516040517f1d5ad3c500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a95565b6040805161012081018252606083810180516020908101516fffffffffffffffffffffffffffffffff90811660808087019182524263ffffffff90811660a0808a01829052875151151560c08b01528751870151861660e08b015296518a015185166101008a015292885288519586018952818a01805186015185168752868601939093528251511515868a01528251850151841686880152915188015183168582015283870194855288880151878901908152848a0151151587890152895167ffffffffffffffff1660009081526007865289902088518051825482890151838e01519289167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316177001000000000000000000000000000000009189168202177fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff908116740100000000000000000000000000000000000000009415158502178655848d015194880151948a16948a168202949094176001860155995180516002860180549b8301519f830151918b169b9093169a909a179d9097168a029c909c179091169615150295909517909855948101519401519381169316909102919091176003820155915190919060048201906116f79082613b0e565b5060609182015160059190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905581519082015160808301516040517f0f135cbb9afa12a8bf3bbd071c117bcca4ddeca6160ef7f33d012a81b9c0c4719361176e9390929091613c28565b60405180910390a16118ac565b80516117939060059067ffffffffffffffff16612479565b6117d85780516040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610a95565b805167ffffffffffffffff16600090815260076020526040812080547fffffffffffffffffffffff000000000000000000000000000000000000000000908116825560018201839055600282018054909116905560038101829055906118416004830182612f04565b5060050180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055805160405167ffffffffffffffff90911681527f5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d8599169060200160405180910390a15b50600101611451565b6118be81611030565b611900576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610a95565b600480546040517f83826b2b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925233602483015273ffffffffffffffffffffffffffffffffffffffff16906383826b2b90604401602060405180830381865afa15801561197f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119a39190613595565b611443576040517f728fe07b000000000000000000000000000000000000000000000000000000008152336004820152602401610a95565b6119e482610584565b80519060200120818051906020012014611a2c57806040517f24eb47e5000000000000000000000000000000000000000000000000000000008152600401610a95919061303a565b5050565b67ffffffffffffffff82166000908152600760205260409020611a2c90600201827f0000000000000000000000000000000000000000000000000000000000000000612485565b60085481516040808401516060850151602086015192517f8627fad600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90951694638627fad694611ae09490939291600401613cab565b600060405180830381600087803b158015611afa57600080fd5b505af1158015611b0e573d6000803e3d6000fd5b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314611b96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610a95565b565b7f0000000000000000000000000000000000000000000000000000000000000000611bef576040517f35f4a7b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015611c85576000838281518110611c0f57611c0f6139f4565b60200260200101519050611c2d81600261280890919063ffffffff16565b15611c7c5760405173ffffffffffffffffffffffffffffffffffffffff821681527f800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf75669060200160405180910390a15b50600101611bf2565b5060005b815181101561142d576000828281518110611ca657611ca66139f4565b60200260200101519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611cea5750611d46565b611cf560028261282a565b15611d445760405173ffffffffffffffffffffffffffffffffffffffff821681527f2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d89060200160405180910390a15b505b600101611c89565b7f00000000000000000000000000000000000000000000000000000000000000008015611d835750611d8160028261284c565b155b15611443576040517fd0d2597600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610a95565b611ddb81611030565b611e1d576040517fa9902c7e00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610a95565b600480546040517fa8d87a3b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84169281019290925273ffffffffffffffffffffffffffffffffffffffff169063a8d87a3b90602401602060405180830381865afa158015611e96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eba9190613d0c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611443576040517f728fe07b000000000000000000000000000000000000000000000000000000008152336004820152602401610a95565b67ffffffffffffffff82166000908152600760205260409020611a2c90827f0000000000000000000000000000000000000000000000000000000000000000612485565b6008546060820151611fb19173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081169291169061287b565b60085460408083015183516060850151602086015193517f9687544500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9095169463968754459461201994939291600401613d29565b6000604051808303816000875af1158015612038573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052611a2c9190810190613d89565b60008181526001830160205260408120541515611029565b6060600061102983612908565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261213182606001516fffffffffffffffffffffffffffffffff1683600001516fffffffffffffffffffffffffffffffff16846020015163ffffffff16426121159190613e26565b85608001516fffffffffffffffffffffffffffffffff16612963565b6fffffffffffffffffffffffffffffffff1682525063ffffffff4216602082015290565b61215e83611030565b6121a0576040517f1e670e4b00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff84166004820152602401610a95565b6121ab826000612334565b67ffffffffffffffff831660009081526007602052604090206121ce908361298d565b6121d9816000612334565b67ffffffffffffffff831660009081526007602052604090206121ff906002018261298d565b7f0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b83838360405161223293929190613c28565b60405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036122be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610a95565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b8151156123fb5781602001516fffffffffffffffffffffffffffffffff1682604001516fffffffffffffffffffffffffffffffff1610158061238a575060408201516fffffffffffffffffffffffffffffffff16155b156123c357816040517f70505e56000000000000000000000000000000000000000000000000000000008152600401610a959190613e39565b8015611a2c576040517f433fc33d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408201516fffffffffffffffffffffffffffffffff16151580612434575060208201516fffffffffffffffffffffffffffffffff1615155b15611a2c57816040517fd68af9cc000000000000000000000000000000000000000000000000000000008152600401610a959190613e39565b60006110298383612b2f565b60006110298383612b7e565b825474010000000000000000000000000000000000000000900460ff1615806124ac575081155b156124b657505050565b825460018401546fffffffffffffffffffffffffffffffff808316929116906000906124fc90700100000000000000000000000000000000900463ffffffff1642613e26565b905080156125bc578183111561253e576040517f9725942a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018601546125789083908590849070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16612963565b86547fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff167001000000000000000000000000000000004263ffffffff160217875592505b848210156126735773ffffffffffffffffffffffffffffffffffffffff841661261b576040517ff94ebcd10000000000000000000000000000000000000000000000000000000081526004810183905260248101869052604401610a95565b6040517f1a76572a000000000000000000000000000000000000000000000000000000008152600481018390526024810186905273ffffffffffffffffffffffffffffffffffffffff85166044820152606401610a95565b848310156127865760018681015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff169060009082906126b79082613e26565b6126c1878a613e26565b6126cb9190613e75565b6126d59190613e88565b905073ffffffffffffffffffffffffffffffffffffffff861661272e576040517f15279c080000000000000000000000000000000000000000000000000000000081526004810182905260248101869052604401610a95565b6040517fd0c8d23a000000000000000000000000000000000000000000000000000000008152600481018290526024810186905273ffffffffffffffffffffffffffffffffffffffff87166044820152606401610a95565b6127908584613e26565b86547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff82161787556040518681529093507f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a9060200160405180910390a1505050505050565b60006110298373ffffffffffffffffffffffffffffffffffffffff8416612b7e565b60006110298373ffffffffffffffffffffffffffffffffffffffff8416612b2f565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515611029565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261142d908490612c71565b60608160000180548060200260200160405190810160405280929190818152602001828054801561062857602002820191906000526020600020905b8154815260200190600101908083116129445750505050509050919050565b6000612982856129738486613ec3565b61297d9087613e75565b612d7d565b90505b949350505050565b81546000906129b690700100000000000000000000000000000000900463ffffffff1642613e26565b90508015612a5857600183015483546129fe916fffffffffffffffffffffffffffffffff80821692811691859170010000000000000000000000000000000090910416612963565b83546fffffffffffffffffffffffffffffffff919091167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116177001000000000000000000000000000000004263ffffffff16021783555b60208201518354612a7e916fffffffffffffffffffffffffffffffff9081169116612d7d565b83548351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffff000000000000000000000000000000009091166fffffffffffffffffffffffffffffffff92831617178455602083015160408085015183167001000000000000000000000000000000000291909216176001850155517f9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c1990612232908490613e39565b6000818152600183016020526040812054612b765750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561057e565b50600061057e565b60008181526001830160205260408120548015612c67576000612ba2600183613e26565b8554909150600090612bb690600190613e26565b9050818114612c1b576000866000018281548110612bd657612bd66139f4565b9060005260206000200154905080876000018481548110612bf957612bf96139f4565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612c2c57612c2c613eda565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061057e565b600091505061057e565b6000612cd3826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612d939092919063ffffffff16565b80519091501561142d5780806020019051810190612cf19190613595565b61142d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610a95565b6000818310612d8c5781611029565b5090919050565b60606129858484600085856000808673ffffffffffffffffffffffffffffffffffffffff168587604051612dc79190613f09565b60006040518083038185875af1925050503d8060008114612e04576040519150601f19603f3d011682016040523d82523d6000602084013e612e09565b606091505b5091509150612e1a87838387612e25565b979650505050505050565b60608315612ebb578251600003612eb45773ffffffffffffffffffffffffffffffffffffffff85163b612eb4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610a95565b5081612985565b6129858383815115612ed05781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a95919061303a565b508054612f1090613548565b6000825580601f10612f20575050565b601f01602090049060005260206000209081019061144391905b80821115612f4e5760008155600101612f3a565b5090565b600060208284031215612f6457600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461102957600080fd5b803567ffffffffffffffff81168114612fac57600080fd5b919050565b600060208284031215612fc357600080fd5b61102982612f94565b60005b83811015612fe7578181015183820152602001612fcf565b50506000910152565b60008151808452613008816020860160208601612fcc565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006110296020830184612ff0565b60006020828403121561305f57600080fd5b813567ffffffffffffffff81111561307657600080fd5b820160e0818503121561102957600080fd5b60008083601f84011261309a57600080fd5b50813567ffffffffffffffff8111156130b257600080fd5b6020830191508360208260051b85010111156130cd57600080fd5b9250929050565b600080600080604085870312156130ea57600080fd5b843567ffffffffffffffff8082111561310257600080fd5b61310e88838901613088565b9096509450602087013591508082111561312757600080fd5b5061313487828801613088565b95989497509550505050565b60008060006040848603121561315557600080fd5b61315e84612f94565b9250602084013567ffffffffffffffff8082111561317b57600080fd5b818601915086601f83011261318f57600080fd5b81358181111561319e57600080fd5b8760208285010111156131b057600080fd5b6020830194508093505050509250925092565b6000602082840312156131d557600080fd5b813567ffffffffffffffff8111156131ec57600080fd5b82016080818503121561102957600080fd5b60208152600082516040602084015261321a6060840182612ff0565b905060208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160408501526132558282612ff0565b95945050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461144357600080fd5b8035612fac8161325e565b6000806040838503121561329e57600080fd5b6132a783612f94565b915060208301356132b78161325e565b809150509250929050565b6000602082840312156132d457600080fd5b81356110298161325e565b6020808252825182820181905260009190848201906040850190845b8181101561332d57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016132fb565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561332d57835167ffffffffffffffff1683529284019291840191600101613355565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff811182821017156133cd576133cd61337b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561341a5761341a61337b565b604052919050565b801515811461144357600080fd5b80356fffffffffffffffffffffffffffffffff81168114612fac57600080fd5b60006060828403121561346257600080fd5b6040516060810181811067ffffffffffffffff821117156134855761348561337b565b604052905080823561349681613422565b81526134a460208401613430565b60208201526134b560408401613430565b60408201525092915050565b600080600060e084860312156134d657600080fd5b6134df84612f94565b92506134ee8560208601613450565b91506134fd8560808601613450565b90509250925092565b6000806020838503121561351957600080fd5b823567ffffffffffffffff81111561353057600080fd5b61353c85828601613088565b90969095509350505050565b600181811c9082168061355c57607f821691505b6020821081036109c5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000602082840312156135a757600080fd5b815161102981613422565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126135e757600080fd5b83018035915067ffffffffffffffff82111561360257600080fd5b6020019150368190038213156130cd57600080fd5b600067ffffffffffffffff8211156136315761363161337b565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261366e57600080fd5b813561368161367c82613617565b6133d3565b81815284602083860101111561369657600080fd5b816020850160208301376000918101602001919091529392505050565b600060e082360312156136c557600080fd5b6136cd6133aa565b823567ffffffffffffffff808211156136e557600080fd5b6136f13683870161365d565b83526136ff60208601612f94565b602084015261371060408601613280565b604084015260608501356060840152608085013591508082111561373357600080fd5b61373f3683870161365d565b608084015260a085013591508082111561375857600080fd5b6137643683870161365d565b60a084015260c085013591508082111561377d57600080fd5b5061378a3682860161365d565b60c08301525092915050565b601f82111561142d576000816000526020600020601f850160051c810160208610156137bf5750805b601f850160051c820191505b818110156137de578281556001016137cb565b505050505050565b67ffffffffffffffff8311156137fe576137fe61337b565b6138128361380c8354613548565b83613796565b6000601f841160018114613864576000851561382e5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355611b0e565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156138b35786850135825560209485019460019092019101613893565b50868210156138ee577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555050505050565b6040815260006139136040830186612ff0565b82810360208401528381528385602083013760006020858301015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116820101915050949350505050565b60006080823603121561397657600080fd5b6040516080810167ffffffffffffffff828210818311171561399a5761399a61337b565b8160405284359150808211156139af57600080fd5b506139bc3682860161365d565b8252506139cb60208401612f94565b602082015260408301356139de8161325e565b6040820152606092830135928101929092525090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee1833603018112613a5757600080fd5b9190910192915050565b60006101208236031215613a7457600080fd5b60405160a0810167ffffffffffffffff8282108183111715613a9857613a9861337b565b81604052613aa585612f94565b835260208501359150613ab782613422565b8160208401526040850135915080821115613ad157600080fd5b50613ade3682860161365d565b604083015250613af13660608501613450565b6060820152613b033660c08501613450565b608082015292915050565b815167ffffffffffffffff811115613b2857613b2861337b565b613b3c81613b368454613548565b84613796565b602080601f831160018114613b8f5760008415613b595750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556137de565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015613bdc57888601518255948401946001909101908401613bbd565b5085821015613c1857878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b67ffffffffffffffff8416815260e08101613c7460208301858051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b82511515608083015260208301516fffffffffffffffffffffffffffffffff90811660a084015260408401511660c0830152612985565b60a081526000613cbe60a0830187612ff0565b73ffffffffffffffffffffffffffffffffffffffff8616602084015284604084015267ffffffffffffffff841660608401528281036080840152600081526020810191505095945050505050565b600060208284031215613d1e57600080fd5b81516110298161325e565b73ffffffffffffffffffffffffffffffffffffffff8516815260a060208201526000613d5860a0830186612ff0565b60408301949094525067ffffffffffffffff9190911660608201528082036080909101526000815260200192915050565b600060208284031215613d9b57600080fd5b815167ffffffffffffffff811115613db257600080fd5b8201601f81018413613dc357600080fd5b8051613dd161367c82613617565b818152856020838501011115613de657600080fd5b613255826020830160208601612fcc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561057e5761057e613df7565b6060810161057e82848051151582526020808201516fffffffffffffffffffffffffffffffff9081169184019190915260409182015116910152565b8082018082111561057e5761057e613df7565b600082613ebe577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b808202811582820484141761057e5761057e613df7565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60008251613a57818460208701612fcc56fe4275726e4d696e74546f6b656e506f6f6c416e6450726f787920312e352e302d646576a164736f6c6343000818000a",
}

var BurnMintTokenPoolAndProxyABI = BurnMintTokenPoolAndProxyMetaData.ABI

var BurnMintTokenPoolAndProxyBin = BurnMintTokenPoolAndProxyMetaData.Bin

func DeployBurnMintTokenPoolAndProxy(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, allowlist []common.Address, rmnProxy common.Address, router common.Address) (common.Address, *types.Transaction, *BurnMintTokenPoolAndProxy, error) {
	parsed, err := BurnMintTokenPoolAndProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintTokenPoolAndProxyBin), backend, token, allowlist, rmnProxy, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintTokenPoolAndProxy{address: address, abi: *parsed, BurnMintTokenPoolAndProxyCaller: BurnMintTokenPoolAndProxyCaller{contract: contract}, BurnMintTokenPoolAndProxyTransactor: BurnMintTokenPoolAndProxyTransactor{contract: contract}, BurnMintTokenPoolAndProxyFilterer: BurnMintTokenPoolAndProxyFilterer{contract: contract}}, nil
}

type BurnMintTokenPoolAndProxy struct {
	address common.Address
	abi     abi.ABI
	BurnMintTokenPoolAndProxyCaller
	BurnMintTokenPoolAndProxyTransactor
	BurnMintTokenPoolAndProxyFilterer
}

type BurnMintTokenPoolAndProxyCaller struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolAndProxyTransactor struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolAndProxyFilterer struct {
	contract *bind.BoundContract
}

type BurnMintTokenPoolAndProxySession struct {
	Contract     *BurnMintTokenPoolAndProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolAndProxyCallerSession struct {
	Contract *BurnMintTokenPoolAndProxyCaller
	CallOpts bind.CallOpts
}

type BurnMintTokenPoolAndProxyTransactorSession struct {
	Contract     *BurnMintTokenPoolAndProxyTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintTokenPoolAndProxyRaw struct {
	Contract *BurnMintTokenPoolAndProxy
}

type BurnMintTokenPoolAndProxyCallerRaw struct {
	Contract *BurnMintTokenPoolAndProxyCaller
}

type BurnMintTokenPoolAndProxyTransactorRaw struct {
	Contract *BurnMintTokenPoolAndProxyTransactor
}

func NewBurnMintTokenPoolAndProxy(address common.Address, backend bind.ContractBackend) (*BurnMintTokenPoolAndProxy, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintTokenPoolAndProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintTokenPoolAndProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxy{address: address, abi: abi, BurnMintTokenPoolAndProxyCaller: BurnMintTokenPoolAndProxyCaller{contract: contract}, BurnMintTokenPoolAndProxyTransactor: BurnMintTokenPoolAndProxyTransactor{contract: contract}, BurnMintTokenPoolAndProxyFilterer: BurnMintTokenPoolAndProxyFilterer{contract: contract}}, nil
}

func NewBurnMintTokenPoolAndProxyCaller(address common.Address, caller bind.ContractCaller) (*BurnMintTokenPoolAndProxyCaller, error) {
	contract, err := bindBurnMintTokenPoolAndProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyCaller{contract: contract}, nil
}

func NewBurnMintTokenPoolAndProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintTokenPoolAndProxyTransactor, error) {
	contract, err := bindBurnMintTokenPoolAndProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyTransactor{contract: contract}, nil
}

func NewBurnMintTokenPoolAndProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintTokenPoolAndProxyFilterer, error) {
	contract, err := bindBurnMintTokenPoolAndProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyFilterer{contract: contract}, nil
}

func bindBurnMintTokenPoolAndProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintTokenPoolAndProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPoolAndProxy.Contract.BurnMintTokenPoolAndProxyCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.BurnMintTokenPoolAndProxyTransactor.contract.Transfer(opts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.BurnMintTokenPoolAndProxyTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintTokenPoolAndProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.contract.Transfer(opts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetAllowList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetAllowList(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetAllowList() ([]common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetAllowList(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetAllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getAllowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetAllowListEnabled(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetAllowListEnabled() (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetAllowListEnabled(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetCurrentInboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetCurrentInboundRateLimiterState(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", remoteChainSelector)

	if err != nil {
		return *new(RateLimiterTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(RateLimiterTokenBucket)).(*RateLimiterTokenBucket)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetCurrentOutboundRateLimiterState(remoteChainSelector uint64) (RateLimiterTokenBucket, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetCurrentOutboundRateLimiterState(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetOnRamp(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getOnRamp", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetOnRamp(arg0 uint64) (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetOnRamp(&_BurnMintTokenPoolAndProxy.CallOpts, arg0)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetOnRamp(arg0 uint64) (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetOnRamp(&_BurnMintTokenPoolAndProxy.CallOpts, arg0)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getRemotePool", remoteChainSelector)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetRemotePool(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetRemotePool(remoteChainSelector uint64) ([]byte, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetRemotePool(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetRmnProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getRmnProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetRmnProxy(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetRmnProxy() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetRmnProxy(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetRouter() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetRouter(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetRouter() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetRouter(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetSupportedChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getSupportedChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetSupportedChains(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetSupportedChains() ([]uint64, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetSupportedChains(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) GetToken() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetToken(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) GetToken() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.GetToken(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "isOffRamp", sourceChainSelector, offRamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.IsOffRamp(&_BurnMintTokenPoolAndProxy.CallOpts, sourceChainSelector, offRamp)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) IsOffRamp(sourceChainSelector uint64, offRamp common.Address) (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.IsOffRamp(&_BurnMintTokenPoolAndProxy.CallOpts, sourceChainSelector, offRamp)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "isSupportedChain", remoteChainSelector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.IsSupportedChain(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) IsSupportedChain(remoteChainSelector uint64) (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.IsSupportedChain(&_BurnMintTokenPoolAndProxy.CallOpts, remoteChainSelector)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) Owner() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.Owner(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) Owner() (common.Address, error) {
	return _BurnMintTokenPoolAndProxy.Contract.Owner(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SupportsInterface(&_BurnMintTokenPoolAndProxy.CallOpts, interfaceId)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SupportsInterface(&_BurnMintTokenPoolAndProxy.CallOpts, interfaceId)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintTokenPoolAndProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPoolAndProxy.Contract.TypeAndVersion(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintTokenPoolAndProxy.Contract.TypeAndVersion(&_BurnMintTokenPoolAndProxy.CallOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.AcceptOwnership(&_BurnMintTokenPoolAndProxy.TransactOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.AcceptOwnership(&_BurnMintTokenPoolAndProxy.TransactOpts)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "applyAllowListUpdates", removes, adds)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.ApplyAllowListUpdates(&_BurnMintTokenPoolAndProxy.TransactOpts, removes, adds)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) ApplyAllowListUpdates(removes []common.Address, adds []common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.ApplyAllowListUpdates(&_BurnMintTokenPoolAndProxy.TransactOpts, removes, adds)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "applyChainUpdates", chains)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.ApplyChainUpdates(&_BurnMintTokenPoolAndProxy.TransactOpts, chains)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) ApplyChainUpdates(chains []TokenPoolChainUpdate) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.ApplyChainUpdates(&_BurnMintTokenPoolAndProxy.TransactOpts, chains)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "lockOrBurn", lockOrBurnIn)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.LockOrBurn(&_BurnMintTokenPoolAndProxy.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) LockOrBurn(lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.LockOrBurn(&_BurnMintTokenPoolAndProxy.TransactOpts, lockOrBurnIn)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "releaseOrMint", releaseOrMintIn)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.ReleaseOrMint(&_BurnMintTokenPoolAndProxy.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) ReleaseOrMint(releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.ReleaseOrMint(&_BurnMintTokenPoolAndProxy.TransactOpts, releaseOrMintIn)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "setChainRateLimiterConfig", remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPoolAndProxy.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) SetChainRateLimiterConfig(remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetChainRateLimiterConfig(&_BurnMintTokenPoolAndProxy.TransactOpts, remoteChainSelector, outboundConfig, inboundConfig)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) SetPreviousPool(opts *bind.TransactOpts, prevPool common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "setPreviousPool", prevPool)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) SetPreviousPool(prevPool common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetPreviousPool(&_BurnMintTokenPoolAndProxy.TransactOpts, prevPool)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) SetPreviousPool(prevPool common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetPreviousPool(&_BurnMintTokenPoolAndProxy.TransactOpts, prevPool)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "setRemotePool", remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetRemotePool(&_BurnMintTokenPoolAndProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) SetRemotePool(remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetRemotePool(&_BurnMintTokenPoolAndProxy.TransactOpts, remoteChainSelector, remotePoolAddress)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "setRouter", newRouter)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetRouter(&_BurnMintTokenPoolAndProxy.TransactOpts, newRouter)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.SetRouter(&_BurnMintTokenPoolAndProxy.TransactOpts, newRouter)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.TransferOwnership(&_BurnMintTokenPoolAndProxy.TransactOpts, to)
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintTokenPoolAndProxy.Contract.TransferOwnership(&_BurnMintTokenPoolAndProxy.TransactOpts, to)
}

type BurnMintTokenPoolAndProxyAllowListAddIterator struct {
	Event *BurnMintTokenPoolAndProxyAllowListAdd

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyAllowListAddIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyAllowListAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyAllowListAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyAllowListAddIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyAllowListAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyAllowListAdd struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyAllowListAddIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyAllowListAddIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "AllowListAdd", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyAllowListAdd) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "AllowListAdd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyAllowListAdd)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolAndProxyAllowListAdd, error) {
	event := new(BurnMintTokenPoolAndProxyAllowListAdd)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "AllowListAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyAllowListRemoveIterator struct {
	Event *BurnMintTokenPoolAndProxyAllowListRemove

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyAllowListRemoveIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyAllowListRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyAllowListRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyAllowListRemoveIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyAllowListRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyAllowListRemove struct {
	Sender common.Address
	Raw    types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyAllowListRemoveIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyAllowListRemoveIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "AllowListRemove", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyAllowListRemove) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "AllowListRemove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyAllowListRemove)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolAndProxyAllowListRemove, error) {
	event := new(BurnMintTokenPoolAndProxyAllowListRemove)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "AllowListRemove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyBurnedIterator struct {
	Event *BurnMintTokenPoolAndProxyBurned

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyBurnedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyBurnedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyBurned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*BurnMintTokenPoolAndProxyBurnedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyBurnedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "Burned", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyBurned, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "Burned", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyBurned)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseBurned(log types.Log) (*BurnMintTokenPoolAndProxyBurned, error) {
	event := new(BurnMintTokenPoolAndProxyBurned)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyChainAddedIterator struct {
	Event *BurnMintTokenPoolAndProxyChainAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyChainAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyChainAddedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyChainAdded struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyChainAddedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyChainAddedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyChainAdded) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "ChainAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyChainAdded)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseChainAdded(log types.Log) (*BurnMintTokenPoolAndProxyChainAdded, error) {
	event := new(BurnMintTokenPoolAndProxyChainAdded)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ChainAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyChainConfiguredIterator struct {
	Event *BurnMintTokenPoolAndProxyChainConfigured

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyChainConfiguredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyChainConfigured)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyChainConfigured)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyChainConfiguredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyChainConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyChainConfigured struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	Raw                       types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyChainConfiguredIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyChainConfiguredIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "ChainConfigured", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyChainConfigured) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "ChainConfigured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyChainConfigured)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseChainConfigured(log types.Log) (*BurnMintTokenPoolAndProxyChainConfigured, error) {
	event := new(BurnMintTokenPoolAndProxyChainConfigured)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ChainConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyChainRemovedIterator struct {
	Event *BurnMintTokenPoolAndProxyChainRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyChainRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyChainRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyChainRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyChainRemovedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyChainRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyChainRemoved struct {
	RemoteChainSelector uint64
	Raw                 types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyChainRemovedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyChainRemovedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "ChainRemoved", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyChainRemoved) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "ChainRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyChainRemoved)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseChainRemoved(log types.Log) (*BurnMintTokenPoolAndProxyChainRemoved, error) {
	event := new(BurnMintTokenPoolAndProxyChainRemoved)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ChainRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyConfigChangedIterator struct {
	Event *BurnMintTokenPoolAndProxyConfigChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyConfigChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyConfigChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyConfigChanged struct {
	Config RateLimiterConfig
	Raw    types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyConfigChangedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyConfigChangedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyConfigChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyConfigChanged)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseConfigChanged(log types.Log) (*BurnMintTokenPoolAndProxyConfigChanged, error) {
	event := new(BurnMintTokenPoolAndProxyConfigChanged)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyLegacyPoolChangedIterator struct {
	Event *BurnMintTokenPoolAndProxyLegacyPoolChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyLegacyPoolChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyLegacyPoolChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyLegacyPoolChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyLegacyPoolChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyLegacyPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyLegacyPoolChanged struct {
	OldPool common.Address
	NewPool common.Address
	Raw     types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterLegacyPoolChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyLegacyPoolChangedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "LegacyPoolChanged")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyLegacyPoolChangedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "LegacyPoolChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchLegacyPoolChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyLegacyPoolChanged) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "LegacyPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyLegacyPoolChanged)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "LegacyPoolChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseLegacyPoolChanged(log types.Log) (*BurnMintTokenPoolAndProxyLegacyPoolChanged, error) {
	event := new(BurnMintTokenPoolAndProxyLegacyPoolChanged)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "LegacyPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyLockedIterator struct {
	Event *BurnMintTokenPoolAndProxyLocked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyLockedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyLockedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyLocked struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BurnMintTokenPoolAndProxyLockedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyLockedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "Locked", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyLocked, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "Locked", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyLocked)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Locked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseLocked(log types.Log) (*BurnMintTokenPoolAndProxyLocked, error) {
	event := new(BurnMintTokenPoolAndProxyLocked)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyMintedIterator struct {
	Event *BurnMintTokenPoolAndProxyMinted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyMintedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyMintedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyMinted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintTokenPoolAndProxyMintedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyMintedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "Minted", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "Minted", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyMinted)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Minted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseMinted(log types.Log) (*BurnMintTokenPoolAndProxyMinted, error) {
	event := new(BurnMintTokenPoolAndProxyMinted)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator struct {
	Event *BurnMintTokenPoolAndProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyOwnershipTransferRequested)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolAndProxyOwnershipTransferRequested, error) {
	event := new(BurnMintTokenPoolAndProxyOwnershipTransferRequested)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyOwnershipTransferredIterator struct {
	Event *BurnMintTokenPoolAndProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolAndProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyOwnershipTransferredIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyOwnershipTransferred)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolAndProxyOwnershipTransferred, error) {
	event := new(BurnMintTokenPoolAndProxyOwnershipTransferred)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyReleasedIterator struct {
	Event *BurnMintTokenPoolAndProxyReleased

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyReleasedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyReleasedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintTokenPoolAndProxyReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyReleasedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "Released", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyReleased)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Released", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseReleased(log types.Log) (*BurnMintTokenPoolAndProxyReleased, error) {
	event := new(BurnMintTokenPoolAndProxyReleased)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyRemotePoolSetIterator struct {
	Event *BurnMintTokenPoolAndProxyRemotePoolSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyRemotePoolSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyRemotePoolSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyRemotePoolSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyRemotePoolSetIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyRemotePoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyRemotePoolSet struct {
	RemoteChainSelector uint64
	PreviousPoolAddress []byte
	RemotePoolAddress   []byte
	Raw                 types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolAndProxyRemotePoolSetIterator, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyRemotePoolSetIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "RemotePoolSet", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error) {

	var remoteChainSelectorRule []interface{}
	for _, remoteChainSelectorItem := range remoteChainSelector {
		remoteChainSelectorRule = append(remoteChainSelectorRule, remoteChainSelectorItem)
	}

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "RemotePoolSet", remoteChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyRemotePoolSet)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseRemotePoolSet(log types.Log) (*BurnMintTokenPoolAndProxyRemotePoolSet, error) {
	event := new(BurnMintTokenPoolAndProxyRemotePoolSet)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "RemotePoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyRouterUpdatedIterator struct {
	Event *BurnMintTokenPoolAndProxyRouterUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyRouterUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyRouterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyRouterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyRouterUpdatedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyRouterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyRouterUpdated struct {
	OldRouter common.Address
	NewRouter common.Address
	Raw       types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyRouterUpdatedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyRouterUpdatedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "RouterUpdated", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyRouterUpdated) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "RouterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyRouterUpdated)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseRouterUpdated(log types.Log) (*BurnMintTokenPoolAndProxyRouterUpdated, error) {
	event := new(BurnMintTokenPoolAndProxyRouterUpdated)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "RouterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintTokenPoolAndProxyTokensConsumedIterator struct {
	Event *BurnMintTokenPoolAndProxyTokensConsumed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintTokenPoolAndProxyTokensConsumedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintTokenPoolAndProxyTokensConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(BurnMintTokenPoolAndProxyTokensConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *BurnMintTokenPoolAndProxyTokensConsumedIterator) Error() error {
	return it.fail
}

func (it *BurnMintTokenPoolAndProxyTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintTokenPoolAndProxyTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyTokensConsumedIterator, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &BurnMintTokenPoolAndProxyTokensConsumedIterator{contract: _BurnMintTokenPoolAndProxy.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _BurnMintTokenPoolAndProxy.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintTokenPoolAndProxyTokensConsumed)
				if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxyFilterer) ParseTokensConsumed(log types.Log) (*BurnMintTokenPoolAndProxyTokensConsumed, error) {
	event := new(BurnMintTokenPoolAndProxyTokensConsumed)
	if err := _BurnMintTokenPoolAndProxy.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintTokenPoolAndProxy.abi.Events["AllowListAdd"].ID:
		return _BurnMintTokenPoolAndProxy.ParseAllowListAdd(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["AllowListRemove"].ID:
		return _BurnMintTokenPoolAndProxy.ParseAllowListRemove(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["Burned"].ID:
		return _BurnMintTokenPoolAndProxy.ParseBurned(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["ChainAdded"].ID:
		return _BurnMintTokenPoolAndProxy.ParseChainAdded(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["ChainConfigured"].ID:
		return _BurnMintTokenPoolAndProxy.ParseChainConfigured(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["ChainRemoved"].ID:
		return _BurnMintTokenPoolAndProxy.ParseChainRemoved(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["ConfigChanged"].ID:
		return _BurnMintTokenPoolAndProxy.ParseConfigChanged(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["LegacyPoolChanged"].ID:
		return _BurnMintTokenPoolAndProxy.ParseLegacyPoolChanged(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["Locked"].ID:
		return _BurnMintTokenPoolAndProxy.ParseLocked(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["Minted"].ID:
		return _BurnMintTokenPoolAndProxy.ParseMinted(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnMintTokenPoolAndProxy.ParseOwnershipTransferRequested(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["OwnershipTransferred"].ID:
		return _BurnMintTokenPoolAndProxy.ParseOwnershipTransferred(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["Released"].ID:
		return _BurnMintTokenPoolAndProxy.ParseReleased(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["RemotePoolSet"].ID:
		return _BurnMintTokenPoolAndProxy.ParseRemotePoolSet(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["RouterUpdated"].ID:
		return _BurnMintTokenPoolAndProxy.ParseRouterUpdated(log)
	case _BurnMintTokenPoolAndProxy.abi.Events["TokensConsumed"].ID:
		return _BurnMintTokenPoolAndProxy.ParseTokensConsumed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintTokenPoolAndProxyAllowListAdd) Topic() common.Hash {
	return common.HexToHash("0x2640d4d76caf8bf478aabfa982fa4e1c4eb71a37f93cd15e80dbc657911546d8")
}

func (BurnMintTokenPoolAndProxyAllowListRemove) Topic() common.Hash {
	return common.HexToHash("0x800671136ab6cfee9fbe5ed1fb7ca417811aca3cf864800d127b927adedf7566")
}

func (BurnMintTokenPoolAndProxyBurned) Topic() common.Hash {
	return common.HexToHash("0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7")
}

func (BurnMintTokenPoolAndProxyChainAdded) Topic() common.Hash {
	return common.HexToHash("0x0f135cbb9afa12a8bf3bbd071c117bcca4ddeca6160ef7f33d012a81b9c0c471")
}

func (BurnMintTokenPoolAndProxyChainConfigured) Topic() common.Hash {
	return common.HexToHash("0x0350d63aa5f270e01729d00d627eeb8f3429772b1818c016c66a588a864f912b")
}

func (BurnMintTokenPoolAndProxyChainRemoved) Topic() common.Hash {
	return common.HexToHash("0x5204aec90a3c794d8e90fded8b46ae9c7c552803e7e832e0c1d358396d859916")
}

func (BurnMintTokenPoolAndProxyConfigChanged) Topic() common.Hash {
	return common.HexToHash("0x9ea3374b67bf275e6bb9c8ae68f9cae023e1c528b4b27e092f0bb209d3531c19")
}

func (BurnMintTokenPoolAndProxyLegacyPoolChanged) Topic() common.Hash {
	return common.HexToHash("0x81accd0a7023865eaa51b3399dd0eafc488bf3ba238402911e1659cfe860f228")
}

func (BurnMintTokenPoolAndProxyLocked) Topic() common.Hash {
	return common.HexToHash("0x9f1ec8c880f76798e7b793325d625e9b60e4082a553c98f42b6cda368dd60008")
}

func (BurnMintTokenPoolAndProxyMinted) Topic() common.Hash {
	return common.HexToHash("0x9d228d69b5fdb8d273a2336f8fb8612d039631024ea9bf09c424a9503aa078f0")
}

func (BurnMintTokenPoolAndProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintTokenPoolAndProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintTokenPoolAndProxyReleased) Topic() common.Hash {
	return common.HexToHash("0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52")
}

func (BurnMintTokenPoolAndProxyRemotePoolSet) Topic() common.Hash {
	return common.HexToHash("0xdb4d6220746a38cbc5335f7e108f7de80f482f4d23350253dfd0917df75a14bf")
}

func (BurnMintTokenPoolAndProxyRouterUpdated) Topic() common.Hash {
	return common.HexToHash("0x02dc5c233404867c793b749c6d644beb2277536d18a7e7974d3f238e4c6f1684")
}

func (BurnMintTokenPoolAndProxyTokensConsumed) Topic() common.Hash {
	return common.HexToHash("0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a")
}

func (_BurnMintTokenPoolAndProxy *BurnMintTokenPoolAndProxy) Address() common.Address {
	return _BurnMintTokenPoolAndProxy.address
}

type BurnMintTokenPoolAndProxyInterface interface {
	GetAllowList(opts *bind.CallOpts) ([]common.Address, error)

	GetAllowListEnabled(opts *bind.CallOpts) (bool, error)

	GetCurrentInboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, remoteChainSelector uint64) (RateLimiterTokenBucket, error)

	GetOnRamp(opts *bind.CallOpts, arg0 uint64) (common.Address, error)

	GetRemotePool(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)

	GetRmnProxy(opts *bind.CallOpts) (common.Address, error)

	GetRouter(opts *bind.CallOpts) (common.Address, error)

	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	IsOffRamp(opts *bind.CallOpts, sourceChainSelector uint64, offRamp common.Address) (bool, error)

	IsSupportedChain(opts *bind.CallOpts, remoteChainSelector uint64) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ApplyAllowListUpdates(opts *bind.TransactOpts, removes []common.Address, adds []common.Address) (*types.Transaction, error)

	ApplyChainUpdates(opts *bind.TransactOpts, chains []TokenPoolChainUpdate) (*types.Transaction, error)

	LockOrBurn(opts *bind.TransactOpts, lockOrBurnIn PoolLockOrBurnInV1) (*types.Transaction, error)

	ReleaseOrMint(opts *bind.TransactOpts, releaseOrMintIn PoolReleaseOrMintInV1) (*types.Transaction, error)

	SetChainRateLimiterConfig(opts *bind.TransactOpts, remoteChainSelector uint64, outboundConfig RateLimiterConfig, inboundConfig RateLimiterConfig) (*types.Transaction, error)

	SetPreviousPool(opts *bind.TransactOpts, prevPool common.Address) (*types.Transaction, error)

	SetRemotePool(opts *bind.TransactOpts, remoteChainSelector uint64, remotePoolAddress []byte) (*types.Transaction, error)

	SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterAllowListAdd(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyAllowListAddIterator, error)

	WatchAllowListAdd(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyAllowListAdd) (event.Subscription, error)

	ParseAllowListAdd(log types.Log) (*BurnMintTokenPoolAndProxyAllowListAdd, error)

	FilterAllowListRemove(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyAllowListRemoveIterator, error)

	WatchAllowListRemove(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyAllowListRemove) (event.Subscription, error)

	ParseAllowListRemove(log types.Log) (*BurnMintTokenPoolAndProxyAllowListRemove, error)

	FilterBurned(opts *bind.FilterOpts, sender []common.Address) (*BurnMintTokenPoolAndProxyBurnedIterator, error)

	WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyBurned, sender []common.Address) (event.Subscription, error)

	ParseBurned(log types.Log) (*BurnMintTokenPoolAndProxyBurned, error)

	FilterChainAdded(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyChainAddedIterator, error)

	WatchChainAdded(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyChainAdded) (event.Subscription, error)

	ParseChainAdded(log types.Log) (*BurnMintTokenPoolAndProxyChainAdded, error)

	FilterChainConfigured(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyChainConfiguredIterator, error)

	WatchChainConfigured(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyChainConfigured) (event.Subscription, error)

	ParseChainConfigured(log types.Log) (*BurnMintTokenPoolAndProxyChainConfigured, error)

	FilterChainRemoved(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyChainRemovedIterator, error)

	WatchChainRemoved(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyChainRemoved) (event.Subscription, error)

	ParseChainRemoved(log types.Log) (*BurnMintTokenPoolAndProxyChainRemoved, error)

	FilterConfigChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyConfigChangedIterator, error)

	WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyConfigChanged) (event.Subscription, error)

	ParseConfigChanged(log types.Log) (*BurnMintTokenPoolAndProxyConfigChanged, error)

	FilterLegacyPoolChanged(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyLegacyPoolChangedIterator, error)

	WatchLegacyPoolChanged(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyLegacyPoolChanged) (event.Subscription, error)

	ParseLegacyPoolChanged(log types.Log) (*BurnMintTokenPoolAndProxyLegacyPoolChanged, error)

	FilterLocked(opts *bind.FilterOpts, sender []common.Address) (*BurnMintTokenPoolAndProxyLockedIterator, error)

	WatchLocked(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyLocked, sender []common.Address) (event.Subscription, error)

	ParseLocked(log types.Log) (*BurnMintTokenPoolAndProxyLocked, error)

	FilterMinted(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintTokenPoolAndProxyMintedIterator, error)

	WatchMinted(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyMinted, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseMinted(log types.Log) (*BurnMintTokenPoolAndProxyMinted, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolAndProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintTokenPoolAndProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintTokenPoolAndProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintTokenPoolAndProxyOwnershipTransferred, error)

	FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*BurnMintTokenPoolAndProxyReleasedIterator, error)

	WatchReleased(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseReleased(log types.Log) (*BurnMintTokenPoolAndProxyReleased, error)

	FilterRemotePoolSet(opts *bind.FilterOpts, remoteChainSelector []uint64) (*BurnMintTokenPoolAndProxyRemotePoolSetIterator, error)

	WatchRemotePoolSet(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyRemotePoolSet, remoteChainSelector []uint64) (event.Subscription, error)

	ParseRemotePoolSet(log types.Log) (*BurnMintTokenPoolAndProxyRemotePoolSet, error)

	FilterRouterUpdated(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyRouterUpdatedIterator, error)

	WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyRouterUpdated) (event.Subscription, error)

	ParseRouterUpdated(log types.Log) (*BurnMintTokenPoolAndProxyRouterUpdated, error)

	FilterTokensConsumed(opts *bind.FilterOpts) (*BurnMintTokenPoolAndProxyTokensConsumedIterator, error)

	WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *BurnMintTokenPoolAndProxyTokensConsumed) (event.Subscription, error)

	ParseTokensConsumed(log types.Log) (*BurnMintTokenPoolAndProxyTokensConsumed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}