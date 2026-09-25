// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ValidatorSetABI is the input ABI used to generate the binding from.
const ValidatorSetABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialValidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enterMaintenance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitMaintenance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"faultTolerance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInMaintenance\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maintenanceSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalExecuted\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalId\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalIdAtVersion\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"versionNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"proposalVotes\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteAddValidator\",\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voteRemoveValidator\",\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MaintenanceEntered\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaintenanceExited\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAdded\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemoved\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorVote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operation\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"quorum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"

var ValidatorSetParsedABI, _ = abi.JSON(strings.NewReader(ValidatorSetABI))

// ValidatorSetBin is the compiled bytecode used for deploying new contracts.
var ValidatorSetBin = "0x608060405234801561001057600080fd5b50604051612c28380380612c288339818101604052810190610032919061048c565b6000815111610076576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006d90610532565b60405180910390fd5b60005b81518110156100b8576100ab82828151811061009857610097610552565b5b60200260200101516100bf60201b60201c565b8080600101915050610079565b505061067f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361012e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610125906105cd565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156101bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b29061065f565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600080549050600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610323826102da565b810181811067ffffffffffffffff82111715610342576103416102eb565b5b80604052505050565b60006103556102c1565b9050610361828261031a565b919050565b600067ffffffffffffffff821115610381576103806102eb565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103c282610397565b9050919050565b6103d2816103b7565b81146103dd57600080fd5b50565b6000815190506103ef816103c9565b92915050565b600061040861040384610366565b61034b565b9050808382526020820190506020840283018581111561042b5761042a610392565b5b835b81811015610454578061044088826103e0565b84526020840193505060208101905061042d565b5050509392505050565b600082601f830112610473576104726102d5565b5b81516104838482602086016103f5565b91505092915050565b6000602082840312156104a2576104a16102cb565b5b600082015167ffffffffffffffff8111156104c0576104bf6102d0565b5b6104cc8482850161045e565b91505092915050565b600082825260208201905092915050565b7f56616c696461746f725365743a20656d7074792076616c696461746f72730000600082015250565b600061051c601e836104d5565b9150610527826104e6565b602082019050919050565b6000602082019050818103600083015261054b8161050f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f56616c696461746f725365743a207a65726f2076616c696461746f7200000000600082015250565b60006105b7601c836104d5565b91506105c282610581565b602082019050919050565b600060208201905081810360008301526105e6816105aa565b9050919050565b7f56616c696461746f725365743a206475706c69636174652076616c696461746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b60006106496021836104d5565b9150610654826105ed565b604082019050919050565b600060208201905081810360008301526106788161063c565b9050919050565b61259a8061068e6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80639369d7de116100a2578063dfa45bdd11610071578063dfa45bdd14610282578063e6c166e3146102b2578063eac942a9146102e2578063f3513a3714610312578063facd743b146103305761010b565b80639369d7de1461020e5780639bf6a74a14610218578063a8b4b21014610234578063b7ab4db5146102645761010b565b806357456f22116100de57806357456f221461017457806358753136146101905780635e401b3e146101c057806375b76f1c146101f05761010b565b806304c4fec61461011057806322e283691461011a5780633068ef4d1461013857806354fd4d5014610156575b600080fd5b610118610360565b005b61012261052d565b60405161012f91906117ba565b60405180910390f35b610140610537565b60405161014d91906117ba565b60405180910390f35b61015e61055e565b60405161016b91906117ba565b60405180910390f35b61018e60048036038101906101899190611838565b610568565b005b6101aa60048036038101906101a591906118b6565b6106fe565b6040516101b79190611922565b60405180910390f35b6101da60048036038101906101d5919061193d565b610734565b6040516101e791906119ab565b60405180910390f35b6101f86107a9565b60405161020591906117ba565b60405180910390f35b6102166107ce565b005b610232600480360381019061022d9190611838565b6109e7565b005b61024e600480360381019061024991906119c6565b610b55565b60405161025b9190611922565b60405180910390f35b61026c610b69565b6040516102799190611ac4565b60405180910390f35b61029c60048036038101906102979190611838565b610ce8565b6040516102a991906119ab565b60405180910390f35b6102cc60048036038101906102c791906119c6565b610d3e565b6040516102d991906117ba565b60405180910390f35b6102fc60048036038101906102f791906119c6565b610d68565b60405161030991906119ab565b60405180910390f35b61031a610d9f565b6040516103279190611ac4565b60405180910390f35b61034a60048036038101906103459190611838565b610ea3565b60405161035791906119ab565b60405180910390f35b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e390611b69565b60405180910390fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610478576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046f90611bd5565b60405180910390fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600460008154809291906104e390611c24565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f0203ae0082a4b3f0b71e461ab5245258799bdb5d52d0f224a6aaed75deca688960405160405180910390a2565b6000600454905090565b600060016105436107a9565b600261054f9190611c4d565b6105599190611c8f565b905090565b6000600554905090565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166105f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105eb90611b69565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610663576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065a90611d0f565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e790611da1565b60405180910390fd5b6106fb600082610ef9565b50565b600083838360405160200161071593929190611ec8565b6040516020818303038152906040528051906020012090509392505050565b60006006600061074486866111fb565b815260200190815260200160002060020160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690509392505050565b6000600360016000805490506107bf9190611f05565b6107c99190611f68565b905090565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661085a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085190611b69565b60405180910390fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156108e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108de9061200b565b60405180910390fd5b6108ef6107a9565b60045410610932576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109299061209d565b60405180910390fd5b6001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506004600081548092919061099d906120bd565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f44aa86acf916265e97ff989ab2c2289b9f442a5a4769121a3c82d3daf7ea972b60405160405180910390a2565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6a90611b69565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610aff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af690612177565b60405180910390fd5b600160008054905011610b47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3e90612209565b60405180910390fd5b610b52600182610ef9565b50565b6000610b6183836111fb565b905092915050565b60606000600454600080549050610b809190611f05565b905060008167ffffffffffffffff811115610b9e57610b9d612229565b5b604051908082528060200260200182016040528015610bcc5781602001602082028036833780820191505090505b5090506000805b600080549050811015610cde576000808281548110610bf557610bf4612258565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610cd05780848481518110610c8757610c86612258565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508280610ccc906120bd565b9350505b508080600101915050610bd3565b5081935050505090565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600060066000610d4e85856111fb565b815260200190815260200160002060000154905092915050565b600060066000610d7885856111fb565b815260200190815260200160002060010160009054906101000a900460ff16905092915050565b60606000808054905067ffffffffffffffff811115610dc157610dc0612229565b5b604051908082528060200260200182016040528015610def5781602001602082028036833780820191505090505b50905060005b600080549050811015610e9b5760008181548110610e1657610e15612258565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828281518110610e5457610e53612258565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610df5565b508091505090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000610f0583836111fb565b905060006006600083815260200190815260200160002090508060010160009054906101000a900460ff1615610f70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f67906122d3565b60405180910390fd5b8060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610fff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff69061233f565b60405180910390fd5b60018160020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080600001600081548092919061106e906120bd565b9190505550600061107d610537565b90508373ffffffffffffffffffffffffffffffffffffffff168560018111156110a9576110a8611dc1565b5b847f88de5c1b4216be3907eda6a00c98004ce4a2076ccd207f384d40e92d288fdab6338660000154866040516110e19392919061236e565b60405180910390a4808260000154106111f45760018260010160006101000a81548160ff0219169083151502179055506000600181111561112557611124611dc1565b5b85600181111561113857611137611dc1565b5b0361118e5761114684611212565b8373ffffffffffffffffffffffffffffffffffffffff167fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec388498760405160405180910390a26111db565b61119784611414565b8373ffffffffffffffffffffffffffffffffffffffff167fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f160405160405180910390a25b600560008154809291906111ee906120bd565b91905055505b5050505050565b600061120a83836005546106fe565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611281576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611278906123f1565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561130e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130590612483565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600080549050600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166114a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149790612515565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611563576000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506004600081548092919061155d90611c24565b91905055505b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008190505b60016000805490506115bf9190611f05565b8110156116c3576000806001836115d69190611c8f565b815481106115e7576115e6612258565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050806000838154811061162957611628612258565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505080806001019150506115ad565b5060008054806116d6576116d5612535565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555050565b6000819050919050565b6117b4816117a1565b82525050565b60006020820190506117cf60008301846117ab565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611805826117da565b9050919050565b611815816117fa565b811461182057600080fd5b50565b6000813590506118328161180c565b92915050565b60006020828403121561184e5761184d6117d5565b5b600061185c84828501611823565b91505092915050565b6002811061187257600080fd5b50565b60008135905061188481611865565b92915050565b611893816117a1565b811461189e57600080fd5b50565b6000813590506118b08161188a565b92915050565b6000806000606084860312156118cf576118ce6117d5565b5b60006118dd86828701611875565b93505060206118ee86828701611823565b92505060406118ff868287016118a1565b9150509250925092565b6000819050919050565b61191c81611909565b82525050565b60006020820190506119376000830184611913565b92915050565b600080600060608486031215611956576119556117d5565b5b600061196486828701611875565b935050602061197586828701611823565b925050604061198686828701611823565b9150509250925092565b60008115159050919050565b6119a581611990565b82525050565b60006020820190506119c0600083018461199c565b92915050565b600080604083850312156119dd576119dc6117d5565b5b60006119eb85828601611875565b92505060206119fc85828601611823565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611a3b816117fa565b82525050565b6000611a4d8383611a32565b60208301905092915050565b6000602082019050919050565b6000611a7182611a06565b611a7b8185611a11565b9350611a8683611a22565b8060005b83811015611ab7578151611a9e8882611a41565b9750611aa983611a59565b925050600181019050611a8a565b5085935050505092915050565b60006020820190508181036000830152611ade8184611a66565b905092915050565b600082825260208201905092915050565b7f56616c696461746f725365743a2073656e646572206973206e6f742076616c6960008201527f6461746f72000000000000000000000000000000000000000000000000000000602082015250565b6000611b53602583611ae6565b9150611b5e82611af7565b604082019050919050565b60006020820190508181036000830152611b8281611b46565b9050919050565b7f56616c696461746f725365743a206e6f7420696e206d61696e74656e616e6365600082015250565b6000611bbf602083611ae6565b9150611bca82611b89565b602082019050919050565b60006020820190508181036000830152611bee81611bb2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611c2f826117a1565b915060008203611c4257611c41611bf5565b5b600182039050919050565b6000611c58826117a1565b9150611c63836117a1565b9250828202611c71816117a1565b91508282048414831517611c8857611c87611bf5565b5b5092915050565b6000611c9a826117a1565b9150611ca5836117a1565b9250828201905080821115611cbd57611cbc611bf5565b5b92915050565b7f56616c696461746f725365743a207a65726f2063616e64696461746500000000600082015250565b6000611cf9601c83611ae6565b9150611d0482611cc3565b602082019050919050565b60006020820190508181036000830152611d2881611cec565b9050919050565b7f56616c696461746f725365743a2063616e6469646174652069732076616c696460008201527f61746f7200000000000000000000000000000000000000000000000000000000602082015250565b6000611d8b602483611ae6565b9150611d9682611d2f565b604082019050919050565b60006020820190508181036000830152611dba81611d7e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028110611e0157611e00611dc1565b5b50565b6000819050611e1282611df0565b919050565b6000611e2282611e04565b9050919050565b60008160f81b9050919050565b6000611e4182611e29565b9050919050565b611e59611e5482611e17565b611e36565b82525050565b60008160601b9050919050565b6000611e7782611e5f565b9050919050565b6000611e8982611e6c565b9050919050565b611ea1611e9c826117fa565b611e7e565b82525050565b6000819050919050565b611ec2611ebd826117a1565b611ea7565b82525050565b6000611ed48286611e48565b600182019150611ee48285611e90565b601482019150611ef48284611eb1565b602082019150819050949350505050565b6000611f10826117a1565b9150611f1b836117a1565b9250828203905081811115611f3357611f32611bf5565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611f73826117a1565b9150611f7e836117a1565b925082611f8e57611f8d611f39565b5b828204905092915050565b7f56616c696461746f725365743a20616c726561647920696e206d61696e74656e60008201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b6000611ff5602483611ae6565b915061200082611f99565b604082019050919050565b6000602082019050818103600083015261202481611fe8565b9050919050565b7f56616c696461746f725365743a20746f6f206d616e792076616c696461746f7260008201527f7320696e206d61696e74656e616e636500000000000000000000000000000000602082015250565b6000612087603083611ae6565b91506120928261202b565b604082019050919050565b600060208201905081810360008301526120b68161207a565b9050919050565b60006120c8826117a1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120fa576120f9611bf5565b5b600182019050919050565b7f56616c696461746f725365743a2063616e646964617465206973206e6f74207660008201527f616c696461746f72000000000000000000000000000000000000000000000000602082015250565b6000612161602883611ae6565b915061216c82612105565b604082019050919050565b6000602082019050818103600083015261219081612154565b9050919050565b7f56616c696461746f725365743a2063616e6e6f742072656d6f7665206c61737460008201527f2076616c696461746f7200000000000000000000000000000000000000000000602082015250565b60006121f3602a83611ae6565b91506121fe82612197565b604082019050919050565b60006020820190508181036000830152612222816121e6565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f56616c696461746f725365743a2070726f706f73616c20657865637574656400600082015250565b60006122bd601f83611ae6565b91506122c882612287565b602082019050919050565b600060208201905081810360008301526122ec816122b0565b9050919050565b7f56616c696461746f725365743a20616c726561647920766f7465640000000000600082015250565b6000612329601b83611ae6565b9150612334826122f3565b602082019050919050565b600060208201905081810360008301526123588161231c565b9050919050565b612368816117fa565b82525050565b6000606082019050612383600083018661235f565b61239060208301856117ab565b61239d60408301846117ab565b949350505050565b7f56616c696461746f725365743a207a65726f2076616c696461746f7200000000600082015250565b60006123db601c83611ae6565b91506123e6826123a5565b602082019050919050565b6000602082019050818103600083015261240a816123ce565b9050919050565b7f56616c696461746f725365743a206475706c69636174652076616c696461746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b600061246d602183611ae6565b915061247882612411565b604082019050919050565b6000602082019050818103600083015261249c81612460565b9050919050565b7f56616c696461746f725365743a2076616c696461746f72206e6f7420666f756e60008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b60006124ff602183611ae6565b915061250a826124a3565b604082019050919050565b6000602082019050818103600083015261252e816124f2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122010ef9fcd63ea0d158ca96d87736aac0c3ac9911a27cfc09892a45b5805a9eed764736f6c634300081e0033"

// DeployValidatorSet deploys a new Ethereum contract, binding an instance of ValidatorSet to it.
func DeployValidatorSet(auth *bind.TransactOpts, backend bind.ContractBackend, initialValidators []common.Address) (common.Address, *types.Transaction, *ValidatorSet, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorSetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorSetBin), backend, initialValidators)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// ValidatorSet is an auto generated Go binding around an Ethereum contract.
type ValidatorSet struct {
	ValidatorSetCaller     // Read-only binding to the contract
	ValidatorSetTransactor // Write-only binding to the contract
	ValidatorSetFilterer   // Log filterer for contract events
}

// ValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSetSession struct {
	Contract     *ValidatorSet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorSetCallerSession struct {
	Contract *ValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorSetTransactorSession struct {
	Contract     *ValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorSetRaw struct {
	Contract *ValidatorSet // Generic contract binding to access the raw methods on
}

// ValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorSetCallerRaw struct {
	Contract *ValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorSetTransactorRaw struct {
	Contract *ValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorSet creates a new instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSet(address common.Address, backend bind.ContractBackend) (*ValidatorSet, error) {
	contract, err := bindValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorSet{ValidatorSetCaller: ValidatorSetCaller{contract: contract}, ValidatorSetTransactor: ValidatorSetTransactor{contract: contract}, ValidatorSetFilterer: ValidatorSetFilterer{contract: contract}}, nil
}

// NewValidatorSetCaller creates a new read-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorSetCaller, error) {
	contract, err := bindValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetCaller{contract: contract}, nil
}

// NewValidatorSetTransactor creates a new write-only instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorSetTransactor, error) {
	contract, err := bindValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetTransactor{contract: contract}, nil
}

// NewValidatorSetFilterer creates a new log filterer instance of ValidatorSet, bound to a specific deployed contract.
func NewValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorSetFilterer, error) {
	contract, err := bindValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetFilterer{contract: contract}, nil
}

// bindValidatorSet binds a generic wrapper to an already deployed contract.
func bindValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.ValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.ValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorSet *ValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorSet *ValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// FaultTolerance is a free data retrieval call binding the contract method 0x75b76f1c.
//
// Solidity: function faultTolerance() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) FaultTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "faultTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FaultTolerance is a free data retrieval call binding the contract method 0x75b76f1c.
//
// Solidity: function faultTolerance() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) FaultTolerance() (*big.Int, error) {
	return _ValidatorSet.Contract.FaultTolerance(&_ValidatorSet.CallOpts)
}

// FaultTolerance is a free data retrieval call binding the contract method 0x75b76f1c.
//
// Solidity: function faultTolerance() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) FaultTolerance() (*big.Int, error) {
	return _ValidatorSet.Contract.FaultTolerance(&_ValidatorSet.CallOpts)
}

// GetAllValidators is a free data retrieval call binding the contract method 0xf3513a37.
//
// Solidity: function getAllValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetCaller) GetAllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getAllValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllValidators is a free data retrieval call binding the contract method 0xf3513a37.
//
// Solidity: function getAllValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetSession) GetAllValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetAllValidators(&_ValidatorSet.CallOpts)
}

// GetAllValidators is a free data retrieval call binding the contract method 0xf3513a37.
//
// Solidity: function getAllValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetCallerSession) GetAllValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetAllValidators(&_ValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetSession) GetValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetValidators(&_ValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_ValidatorSet *ValidatorSetCallerSession) GetValidators() ([]common.Address, error) {
	return _ValidatorSet.Contract.GetValidators(&_ValidatorSet.CallOpts)
}

// HasVoted is a free data retrieval call binding the contract method 0x5e401b3e.
//
// Solidity: function hasVoted(uint8 operation, address candidate, address voter) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) HasVoted(opts *bind.CallOpts, operation uint8, candidate common.Address, voter common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "hasVoted", operation, candidate, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x5e401b3e.
//
// Solidity: function hasVoted(uint8 operation, address candidate, address voter) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) HasVoted(operation uint8, candidate common.Address, voter common.Address) (bool, error) {
	return _ValidatorSet.Contract.HasVoted(&_ValidatorSet.CallOpts, operation, candidate, voter)
}

// HasVoted is a free data retrieval call binding the contract method 0x5e401b3e.
//
// Solidity: function hasVoted(uint8 operation, address candidate, address voter) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) HasVoted(operation uint8, candidate common.Address, voter common.Address) (bool, error) {
	return _ValidatorSet.Contract.HasVoted(&_ValidatorSet.CallOpts, operation, candidate, voter)
}

// IsInMaintenance is a free data retrieval call binding the contract method 0xdfa45bdd.
//
// Solidity: function isInMaintenance(address validator) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsInMaintenance(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "isInMaintenance", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInMaintenance is a free data retrieval call binding the contract method 0xdfa45bdd.
//
// Solidity: function isInMaintenance(address validator) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsInMaintenance(validator common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsInMaintenance(&_ValidatorSet.CallOpts, validator)
}

// IsInMaintenance is a free data retrieval call binding the contract method 0xdfa45bdd.
//
// Solidity: function isInMaintenance(address validator) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsInMaintenance(validator common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsInMaintenance(&_ValidatorSet.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) IsValidator(validator common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsValidator(&_ValidatorSet.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _ValidatorSet.Contract.IsValidator(&_ValidatorSet.CallOpts, validator)
}

// MaintenanceSize is a free data retrieval call binding the contract method 0x22e28369.
//
// Solidity: function maintenanceSize() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) MaintenanceSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "maintenanceSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaintenanceSize is a free data retrieval call binding the contract method 0x22e28369.
//
// Solidity: function maintenanceSize() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) MaintenanceSize() (*big.Int, error) {
	return _ValidatorSet.Contract.MaintenanceSize(&_ValidatorSet.CallOpts)
}

// MaintenanceSize is a free data retrieval call binding the contract method 0x22e28369.
//
// Solidity: function maintenanceSize() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) MaintenanceSize() (*big.Int, error) {
	return _ValidatorSet.Contract.MaintenanceSize(&_ValidatorSet.CallOpts)
}

// ProposalExecuted is a free data retrieval call binding the contract method 0xeac942a9.
//
// Solidity: function proposalExecuted(uint8 operation, address candidate) view returns(bool)
func (_ValidatorSet *ValidatorSetCaller) ProposalExecuted(opts *bind.CallOpts, operation uint8, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "proposalExecuted", operation, candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalExecuted is a free data retrieval call binding the contract method 0xeac942a9.
//
// Solidity: function proposalExecuted(uint8 operation, address candidate) view returns(bool)
func (_ValidatorSet *ValidatorSetSession) ProposalExecuted(operation uint8, candidate common.Address) (bool, error) {
	return _ValidatorSet.Contract.ProposalExecuted(&_ValidatorSet.CallOpts, operation, candidate)
}

// ProposalExecuted is a free data retrieval call binding the contract method 0xeac942a9.
//
// Solidity: function proposalExecuted(uint8 operation, address candidate) view returns(bool)
func (_ValidatorSet *ValidatorSetCallerSession) ProposalExecuted(operation uint8, candidate common.Address) (bool, error) {
	return _ValidatorSet.Contract.ProposalExecuted(&_ValidatorSet.CallOpts, operation, candidate)
}

// ProposalId is a free data retrieval call binding the contract method 0xa8b4b210.
//
// Solidity: function proposalId(uint8 operation, address candidate) view returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) ProposalId(opts *bind.CallOpts, operation uint8, candidate common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "proposalId", operation, candidate)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProposalId is a free data retrieval call binding the contract method 0xa8b4b210.
//
// Solidity: function proposalId(uint8 operation, address candidate) view returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) ProposalId(operation uint8, candidate common.Address) ([32]byte, error) {
	return _ValidatorSet.Contract.ProposalId(&_ValidatorSet.CallOpts, operation, candidate)
}

// ProposalId is a free data retrieval call binding the contract method 0xa8b4b210.
//
// Solidity: function proposalId(uint8 operation, address candidate) view returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) ProposalId(operation uint8, candidate common.Address) ([32]byte, error) {
	return _ValidatorSet.Contract.ProposalId(&_ValidatorSet.CallOpts, operation, candidate)
}

// ProposalIdAtVersion is a free data retrieval call binding the contract method 0x58753136.
//
// Solidity: function proposalIdAtVersion(uint8 operation, address candidate, uint256 versionNumber) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCaller) ProposalIdAtVersion(opts *bind.CallOpts, operation uint8, candidate common.Address, versionNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "proposalIdAtVersion", operation, candidate, versionNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProposalIdAtVersion is a free data retrieval call binding the contract method 0x58753136.
//
// Solidity: function proposalIdAtVersion(uint8 operation, address candidate, uint256 versionNumber) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) ProposalIdAtVersion(operation uint8, candidate common.Address, versionNumber *big.Int) ([32]byte, error) {
	return _ValidatorSet.Contract.ProposalIdAtVersion(&_ValidatorSet.CallOpts, operation, candidate, versionNumber)
}

// ProposalIdAtVersion is a free data retrieval call binding the contract method 0x58753136.
//
// Solidity: function proposalIdAtVersion(uint8 operation, address candidate, uint256 versionNumber) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) ProposalIdAtVersion(operation uint8, candidate common.Address, versionNumber *big.Int) ([32]byte, error) {
	return _ValidatorSet.Contract.ProposalIdAtVersion(&_ValidatorSet.CallOpts, operation, candidate, versionNumber)
}

// ProposalVotes is a free data retrieval call binding the contract method 0xe6c166e3.
//
// Solidity: function proposalVotes(uint8 operation, address candidate) view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) ProposalVotes(opts *bind.CallOpts, operation uint8, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "proposalVotes", operation, candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0xe6c166e3.
//
// Solidity: function proposalVotes(uint8 operation, address candidate) view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) ProposalVotes(operation uint8, candidate common.Address) (*big.Int, error) {
	return _ValidatorSet.Contract.ProposalVotes(&_ValidatorSet.CallOpts, operation, candidate)
}

// ProposalVotes is a free data retrieval call binding the contract method 0xe6c166e3.
//
// Solidity: function proposalVotes(uint8 operation, address candidate) view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) ProposalVotes(operation uint8, candidate common.Address) (*big.Int, error) {
	return _ValidatorSet.Contract.ProposalVotes(&_ValidatorSet.CallOpts, operation, candidate)
}

// QuorumSize is a free data retrieval call binding the contract method 0x3068ef4d.
//
// Solidity: function quorumSize() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) QuorumSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "quorumSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumSize is a free data retrieval call binding the contract method 0x3068ef4d.
//
// Solidity: function quorumSize() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) QuorumSize() (*big.Int, error) {
	return _ValidatorSet.Contract.QuorumSize(&_ValidatorSet.CallOpts)
}

// QuorumSize is a free data retrieval call binding the contract method 0x3068ef4d.
//
// Solidity: function quorumSize() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) QuorumSize() (*big.Int, error) {
	return _ValidatorSet.Contract.QuorumSize(&_ValidatorSet.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ValidatorSet *ValidatorSetCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorSet.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ValidatorSet *ValidatorSetSession) Version() (*big.Int, error) {
	return _ValidatorSet.Contract.Version(&_ValidatorSet.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_ValidatorSet *ValidatorSetCallerSession) Version() (*big.Int, error) {
	return _ValidatorSet.Contract.Version(&_ValidatorSet.CallOpts)
}

// EnterMaintenance is a paid mutator transaction binding the contract method 0x9369d7de.
//
// Solidity: function enterMaintenance() returns()
func (_ValidatorSet *ValidatorSetTransactor) EnterMaintenance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "enterMaintenance")
}

// EnterMaintenance is a paid mutator transaction binding the contract method 0x9369d7de.
//
// Solidity: function enterMaintenance() returns()
func (_ValidatorSet *ValidatorSetSession) EnterMaintenance() (*types.Transaction, error) {
	return _ValidatorSet.Contract.EnterMaintenance(&_ValidatorSet.TransactOpts)
}

// EnterMaintenance is a paid mutator transaction binding the contract method 0x9369d7de.
//
// Solidity: function enterMaintenance() returns()
func (_ValidatorSet *ValidatorSetTransactorSession) EnterMaintenance() (*types.Transaction, error) {
	return _ValidatorSet.Contract.EnterMaintenance(&_ValidatorSet.TransactOpts)
}

// ExitMaintenance is a paid mutator transaction binding the contract method 0x04c4fec6.
//
// Solidity: function exitMaintenance() returns()
func (_ValidatorSet *ValidatorSetTransactor) ExitMaintenance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "exitMaintenance")
}

// ExitMaintenance is a paid mutator transaction binding the contract method 0x04c4fec6.
//
// Solidity: function exitMaintenance() returns()
func (_ValidatorSet *ValidatorSetSession) ExitMaintenance() (*types.Transaction, error) {
	return _ValidatorSet.Contract.ExitMaintenance(&_ValidatorSet.TransactOpts)
}

// ExitMaintenance is a paid mutator transaction binding the contract method 0x04c4fec6.
//
// Solidity: function exitMaintenance() returns()
func (_ValidatorSet *ValidatorSetTransactorSession) ExitMaintenance() (*types.Transaction, error) {
	return _ValidatorSet.Contract.ExitMaintenance(&_ValidatorSet.TransactOpts)
}

// VoteAddValidator is a paid mutator transaction binding the contract method 0x57456f22.
//
// Solidity: function voteAddValidator(address candidate) returns()
func (_ValidatorSet *ValidatorSetTransactor) VoteAddValidator(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "voteAddValidator", candidate)
}

// VoteAddValidator is a paid mutator transaction binding the contract method 0x57456f22.
//
// Solidity: function voteAddValidator(address candidate) returns()
func (_ValidatorSet *ValidatorSetSession) VoteAddValidator(candidate common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.VoteAddValidator(&_ValidatorSet.TransactOpts, candidate)
}

// VoteAddValidator is a paid mutator transaction binding the contract method 0x57456f22.
//
// Solidity: function voteAddValidator(address candidate) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) VoteAddValidator(candidate common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.VoteAddValidator(&_ValidatorSet.TransactOpts, candidate)
}

// VoteRemoveValidator is a paid mutator transaction binding the contract method 0x9bf6a74a.
//
// Solidity: function voteRemoveValidator(address candidate) returns()
func (_ValidatorSet *ValidatorSetTransactor) VoteRemoveValidator(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _ValidatorSet.contract.Transact(opts, "voteRemoveValidator", candidate)
}

// VoteRemoveValidator is a paid mutator transaction binding the contract method 0x9bf6a74a.
//
// Solidity: function voteRemoveValidator(address candidate) returns()
func (_ValidatorSet *ValidatorSetSession) VoteRemoveValidator(candidate common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.VoteRemoveValidator(&_ValidatorSet.TransactOpts, candidate)
}

// VoteRemoveValidator is a paid mutator transaction binding the contract method 0x9bf6a74a.
//
// Solidity: function voteRemoveValidator(address candidate) returns()
func (_ValidatorSet *ValidatorSetTransactorSession) VoteRemoveValidator(candidate common.Address) (*types.Transaction, error) {
	return _ValidatorSet.Contract.VoteRemoveValidator(&_ValidatorSet.TransactOpts, candidate)
}

// ValidatorSetMaintenanceEnteredIterator is returned from FilterMaintenanceEntered and is used to iterate over the raw logs and unpacked data for MaintenanceEntered events raised by the ValidatorSet contract.
type ValidatorSetMaintenanceEnteredIterator struct {
	Event *ValidatorSetMaintenanceEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetMaintenanceEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetMaintenanceEntered)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorSetMaintenanceEntered)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorSetMaintenanceEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetMaintenanceEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetMaintenanceEntered represents a MaintenanceEntered event raised by the ValidatorSet contract.
type ValidatorSetMaintenanceEntered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceEntered is a free log retrieval operation binding the contract event 0x44aa86acf916265e97ff989ab2c2289b9f442a5a4769121a3c82d3daf7ea972b.
//
// Solidity: event MaintenanceEntered(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) FilterMaintenanceEntered(opts *bind.FilterOpts, validator []common.Address) (*ValidatorSetMaintenanceEnteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "MaintenanceEntered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetMaintenanceEnteredIterator{contract: _ValidatorSet.contract, event: "MaintenanceEntered", logs: logs, sub: sub}, nil
}

var MaintenanceEnteredTopicHash = "0x44aa86acf916265e97ff989ab2c2289b9f442a5a4769121a3c82d3daf7ea972b"

// WatchMaintenanceEntered is a free log subscription operation binding the contract event 0x44aa86acf916265e97ff989ab2c2289b9f442a5a4769121a3c82d3daf7ea972b.
//
// Solidity: event MaintenanceEntered(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) WatchMaintenanceEntered(opts *bind.WatchOpts, sink chan<- *ValidatorSetMaintenanceEntered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "MaintenanceEntered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetMaintenanceEntered)
				if err := _ValidatorSet.contract.UnpackLog(event, "MaintenanceEntered", log); err != nil {
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

// ParseMaintenanceEntered is a log parse operation binding the contract event 0x44aa86acf916265e97ff989ab2c2289b9f442a5a4769121a3c82d3daf7ea972b.
//
// Solidity: event MaintenanceEntered(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) ParseMaintenanceEntered(log types.Log) (*ValidatorSetMaintenanceEntered, error) {
	event := new(ValidatorSetMaintenanceEntered)
	if err := _ValidatorSet.contract.UnpackLog(event, "MaintenanceEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorSetMaintenanceExitedIterator is returned from FilterMaintenanceExited and is used to iterate over the raw logs and unpacked data for MaintenanceExited events raised by the ValidatorSet contract.
type ValidatorSetMaintenanceExitedIterator struct {
	Event *ValidatorSetMaintenanceExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetMaintenanceExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetMaintenanceExited)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorSetMaintenanceExited)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorSetMaintenanceExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetMaintenanceExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetMaintenanceExited represents a MaintenanceExited event raised by the ValidatorSet contract.
type ValidatorSetMaintenanceExited struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceExited is a free log retrieval operation binding the contract event 0x0203ae0082a4b3f0b71e461ab5245258799bdb5d52d0f224a6aaed75deca6889.
//
// Solidity: event MaintenanceExited(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) FilterMaintenanceExited(opts *bind.FilterOpts, validator []common.Address) (*ValidatorSetMaintenanceExitedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "MaintenanceExited", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetMaintenanceExitedIterator{contract: _ValidatorSet.contract, event: "MaintenanceExited", logs: logs, sub: sub}, nil
}

var MaintenanceExitedTopicHash = "0x0203ae0082a4b3f0b71e461ab5245258799bdb5d52d0f224a6aaed75deca6889"

// WatchMaintenanceExited is a free log subscription operation binding the contract event 0x0203ae0082a4b3f0b71e461ab5245258799bdb5d52d0f224a6aaed75deca6889.
//
// Solidity: event MaintenanceExited(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) WatchMaintenanceExited(opts *bind.WatchOpts, sink chan<- *ValidatorSetMaintenanceExited, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "MaintenanceExited", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetMaintenanceExited)
				if err := _ValidatorSet.contract.UnpackLog(event, "MaintenanceExited", log); err != nil {
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

// ParseMaintenanceExited is a log parse operation binding the contract event 0x0203ae0082a4b3f0b71e461ab5245258799bdb5d52d0f224a6aaed75deca6889.
//
// Solidity: event MaintenanceExited(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) ParseMaintenanceExited(log types.Log) (*ValidatorSetMaintenanceExited, error) {
	event := new(ValidatorSetMaintenanceExited)
	if err := _ValidatorSet.contract.UnpackLog(event, "MaintenanceExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorSetValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the ValidatorSet contract.
type ValidatorSetValidatorAddedIterator struct {
	Event *ValidatorSetValidatorAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetValidatorAdded)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorSetValidatorAdded)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorSetValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetValidatorAdded represents a ValidatorAdded event raised by the ValidatorSet contract.
type ValidatorSetValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*ValidatorSetValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetValidatorAddedIterator{contract: _ValidatorSet.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

var ValidatorAddedTopicHash = "0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987"

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *ValidatorSetValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetValidatorAdded)
				if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) ParseValidatorAdded(log types.Log) (*ValidatorSetValidatorAdded, error) {
	event := new(ValidatorSetValidatorAdded)
	if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorSetValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the ValidatorSet contract.
type ValidatorSetValidatorRemovedIterator struct {
	Event *ValidatorSetValidatorRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetValidatorRemoved)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorSetValidatorRemoved)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorSetValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetValidatorRemoved represents a ValidatorRemoved event raised by the ValidatorSet contract.
type ValidatorSetValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*ValidatorSetValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetValidatorRemovedIterator{contract: _ValidatorSet.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

var ValidatorRemovedTopicHash = "0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1"

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorSetValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetValidatorRemoved)
				if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_ValidatorSet *ValidatorSetFilterer) ParseValidatorRemoved(log types.Log) (*ValidatorSetValidatorRemoved, error) {
	event := new(ValidatorSetValidatorRemoved)
	if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorSetValidatorVoteIterator is returned from FilterValidatorVote and is used to iterate over the raw logs and unpacked data for ValidatorVote events raised by the ValidatorSet contract.
type ValidatorSetValidatorVoteIterator struct {
	Event *ValidatorSetValidatorVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorSetValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSetValidatorVote)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorSetValidatorVote)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorSetValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSetValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSetValidatorVote represents a ValidatorVote event raised by the ValidatorSet contract.
type ValidatorSetValidatorVote struct {
	ProposalId [32]byte
	Operation  uint8
	Candidate  common.Address
	Voter      common.Address
	Votes      *big.Int
	Quorum     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorVote is a free log retrieval operation binding the contract event 0x88de5c1b4216be3907eda6a00c98004ce4a2076ccd207f384d40e92d288fdab6.
//
// Solidity: event ValidatorVote(bytes32 indexed proposalId, uint8 indexed operation, address indexed candidate, address voter, uint256 votes, uint256 quorum)
func (_ValidatorSet *ValidatorSetFilterer) FilterValidatorVote(opts *bind.FilterOpts, proposalId [][32]byte, operation []uint8, candidate []common.Address) (*ValidatorSetValidatorVoteIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}
	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}

	logs, sub, err := _ValidatorSet.contract.FilterLogs(opts, "ValidatorVote", proposalIdRule, operationRule, candidateRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorSetValidatorVoteIterator{contract: _ValidatorSet.contract, event: "ValidatorVote", logs: logs, sub: sub}, nil
}

var ValidatorVoteTopicHash = "0x88de5c1b4216be3907eda6a00c98004ce4a2076ccd207f384d40e92d288fdab6"

// WatchValidatorVote is a free log subscription operation binding the contract event 0x88de5c1b4216be3907eda6a00c98004ce4a2076ccd207f384d40e92d288fdab6.
//
// Solidity: event ValidatorVote(bytes32 indexed proposalId, uint8 indexed operation, address indexed candidate, address voter, uint256 votes, uint256 quorum)
func (_ValidatorSet *ValidatorSetFilterer) WatchValidatorVote(opts *bind.WatchOpts, sink chan<- *ValidatorSetValidatorVote, proposalId [][32]byte, operation []uint8, candidate []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}
	var candidateRule []interface{}
	for _, candidateItem := range candidate {
		candidateRule = append(candidateRule, candidateItem)
	}

	logs, sub, err := _ValidatorSet.contract.WatchLogs(opts, "ValidatorVote", proposalIdRule, operationRule, candidateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSetValidatorVote)
				if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorVote", log); err != nil {
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

// ParseValidatorVote is a log parse operation binding the contract event 0x88de5c1b4216be3907eda6a00c98004ce4a2076ccd207f384d40e92d288fdab6.
//
// Solidity: event ValidatorVote(bytes32 indexed proposalId, uint8 indexed operation, address indexed candidate, address voter, uint256 votes, uint256 quorum)
func (_ValidatorSet *ValidatorSetFilterer) ParseValidatorVote(log types.Log) (*ValidatorSetValidatorVote, error) {
	event := new(ValidatorSetValidatorVote)
	if err := _ValidatorSet.contract.UnpackLog(event, "ValidatorVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
