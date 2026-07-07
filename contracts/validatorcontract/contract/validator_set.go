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
const ValidatorSetABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialValidators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enterMaintenance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exitMaintenance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"faultTolerance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInMaintenance\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maintenanceSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalExecuted\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalId\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"proposalVotes\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint8\",\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteAddValidator\",\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voteRemoveValidator\",\"inputs\":[{\"name\":\"candidate\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MaintenanceEntered\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaintenanceExited\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAdded\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRemoved\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorVote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operation\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumValidatorSet.Operation\"},{\"name\":\"candidate\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"votes\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"quorum\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"

var ValidatorSetParsedABI, _ = abi.JSON(strings.NewReader(ValidatorSetABI))

// ValidatorSetBin is the compiled bytecode used for deploying new contracts.
var ValidatorSetBin = "0x608060405234801561001057600080fd5b50604051612ac3380380612ac38339818101604052810190610032919061048c565b6000815111610076576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006d90610532565b60405180910390fd5b60005b81518110156100b8576100ab82828151811061009857610097610552565b5b60200260200101516100bf60201b60201c565b8080600101915050610079565b505061067f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361012e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610125906105cd565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156101bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b29061065f565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600080549050600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610323826102da565b810181811067ffffffffffffffff82111715610342576103416102eb565b5b80604052505050565b60006103556102c1565b9050610361828261031a565b919050565b600067ffffffffffffffff821115610381576103806102eb565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103c282610397565b9050919050565b6103d2816103b7565b81146103dd57600080fd5b50565b6000815190506103ef816103c9565b92915050565b600061040861040384610366565b61034b565b9050808382526020820190506020840283018581111561042b5761042a610392565b5b835b81811015610454578061044088826103e0565b84526020840193505060208101905061042d565b5050509392505050565b600082601f830112610473576104726102d5565b5b81516104838482602086016103f5565b91505092915050565b6000602082840312156104a2576104a16102cb565b5b600082015167ffffffffffffffff8111156104c0576104bf6102d0565b5b6104cc8482850161045e565b91505092915050565b600082825260208201905092915050565b7f56616c696461746f725365743a20656d7074792076616c696461746f72730000600082015250565b600061051c601e836104d5565b9150610527826104e6565b602082019050919050565b6000602082019050818103600083015261054b8161050f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f56616c696461746f725365743a207a65726f2076616c696461746f7200000000600082015250565b60006105b7601c836104d5565b91506105c282610581565b602082019050919050565b600060208201905081810360008301526105e6816105aa565b9050919050565b7f56616c696461746f725365743a206475706c69636174652076616c696461746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b60006106496021836104d5565b9150610654826105ed565b604082019050919050565b600060208201905081810360008301526106788161063c565b9050919050565b6124358061068e6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80639bf6a74a11610097578063e6c166e311610066578063e6c166e31461024e578063eac942a91461027e578063f3513a37146102ae578063facd743b146102cc576100f5565b80639bf6a74a146101b4578063a8b4b210146101d0578063b7ab4db514610200578063dfa45bdd1461021e576100f5565b806357456f22116100d357806357456f22146101405780635e401b3e1461015c57806375b76f1c1461018c5780639369d7de146101aa576100f5565b806304c4fec6146100fa57806322e28369146101045780633068ef4d14610122575b600080fd5b6101026102fc565b005b61010c6104c9565b6040516101199190611706565b60405180910390f35b61012a6104d3565b6040516101379190611706565b60405180910390f35b61015a60048036038101906101559190611784565b6104fa565b005b610176600480360381019061017191906117d6565b610690565b6040516101839190611844565b60405180910390f35b610194610705565b6040516101a19190611706565b60405180910390f35b6101b261072a565b005b6101ce60048036038101906101c99190611784565b610943565b005b6101ea60048036038101906101e5919061185f565b610ab1565b6040516101f791906118b8565b60405180910390f35b610208610ae4565b6040516102159190611991565b60405180910390f35b61023860048036038101906102339190611784565b610c63565b6040516102459190611844565b60405180910390f35b6102686004803603810190610263919061185f565b610cb9565b6040516102759190611706565b60405180910390f35b6102986004803603810190610293919061185f565b610ce3565b6040516102a59190611844565b60405180910390f35b6102b6610d1a565b6040516102c39190611991565b60405180910390f35b6102e660048036038101906102e19190611784565b610e1e565b6040516102f39190611844565b60405180910390f35b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610388576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037f90611a36565b60405180910390fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b90611aa2565b60405180910390fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506004600081548092919061047f90611af1565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f0203ae0082a4b3f0b71e461ab5245258799bdb5d52d0f224a6aaed75deca688960405160405180910390a2565b6000600454905090565b600060016104df610705565b60026104eb9190611b1a565b6104f59190611b5c565b905090565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610586576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057d90611a36565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036105f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ec90611bdc565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610682576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067990611c6e565b60405180910390fd5b61068d600082610e74565b50565b6000600560006106a08686610ab1565b815260200190815260200160002060020160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1690509392505050565b60006003600160008054905061071b9190611c8e565b6107259190611cf1565b905090565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ad90611a36565b60405180910390fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610843576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083a90611d94565b60405180910390fd5b61084b610705565b6004541061088e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088590611e26565b60405180910390fd5b6001600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600460008154809291906108f990611e46565b91905055503373ffffffffffffffffffffffffffffffffffffffff167f44aa86acf916265e97ff989ab2c2289b9f442a5a4769121a3c82d3daf7ea972b60405160405180910390a2565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166109cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c690611a36565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a5b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5290611f00565b60405180910390fd5b600160008054905011610aa3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9a90611f92565b60405180910390fd5b610aae600182610e74565b50565b60008282604051602001610ac6929190612098565b60405160208183030381529060405280519060200120905092915050565b60606000600454600080549050610afb9190611c8e565b905060008167ffffffffffffffff811115610b1957610b186120c4565b5b604051908082528060200260200182016040528015610b475781602001602082028036833780820191505090505b5090506000805b600080549050811015610c59576000808281548110610b7057610b6f6120f3565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610c4b5780848481518110610c0257610c016120f3565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508280610c4790611e46565b9350505b508080600101915050610b4e565b5081935050505090565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600060056000610cc98585610ab1565b815260200190815260200160002060000154905092915050565b600060056000610cf38585610ab1565b815260200190815260200160002060010160009054906101000a900460ff16905092915050565b60606000808054905067ffffffffffffffff811115610d3c57610d3b6120c4565b5b604051908082528060200260200182016040528015610d6a5781602001602082028036833780820191505090505b50905060005b600080549050811015610e165760008181548110610d9157610d906120f3565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16828281518110610dcf57610dce6120f3565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610d70565b508091505090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000610e808383610ab1565b905060006005600083815260200190815260200160002090508060010160009054906101000a900460ff1615610eeb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee29061216e565b60405180910390fd5b8060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610f7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f71906121da565b60405180910390fd5b60018160020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550806000016000815480929190610fe990611e46565b91905055506000610ff86104d3565b90508373ffffffffffffffffffffffffffffffffffffffff1685600181111561102457611023611fb2565b5b847f88de5c1b4216be3907eda6a00c98004ce4a2076ccd207f384d40e92d288fdab63386600001548660405161105c93929190612209565b60405180910390a4808260000154106111575760018260010160006101000a81548160ff021916908315150217905550600060018111156110a05761109f611fb2565b5b8560018111156110b3576110b2611fb2565b5b03611109576110c18461115e565b8373ffffffffffffffffffffffffffffffffffffffff167fe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec388498760405160405180910390a2611156565b61111284611360565b8373ffffffffffffffffffffffffffffffffffffffff167fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f160405160405180910390a25b5b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036111cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c49061228c565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161561125a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112519061231e565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600080549050600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166113ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e3906123b0565b60405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156114af576000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600460008154809291906114a990611af1565b91905055505b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008190505b600160008054905061150b9190611c8e565b81101561160f576000806001836115229190611b5c565b81548110611533576115326120f3565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060008381548110611575576115746120f3565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505080806001019150506114f9565b506000805480611622576116216123d0565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555050565b6000819050919050565b611700816116ed565b82525050565b600060208201905061171b60008301846116f7565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061175182611726565b9050919050565b61176181611746565b811461176c57600080fd5b50565b60008135905061177e81611758565b92915050565b60006020828403121561179a57611799611721565b5b60006117a88482850161176f565b91505092915050565b600281106117be57600080fd5b50565b6000813590506117d0816117b1565b92915050565b6000806000606084860312156117ef576117ee611721565b5b60006117fd868287016117c1565b935050602061180e8682870161176f565b925050604061181f8682870161176f565b9150509250925092565b60008115159050919050565b61183e81611829565b82525050565b60006020820190506118596000830184611835565b92915050565b6000806040838503121561187657611875611721565b5b6000611884858286016117c1565b92505060206118958582860161176f565b9150509250929050565b6000819050919050565b6118b28161189f565b82525050565b60006020820190506118cd60008301846118a9565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61190881611746565b82525050565b600061191a83836118ff565b60208301905092915050565b6000602082019050919050565b600061193e826118d3565b61194881856118de565b9350611953836118ef565b8060005b8381101561198457815161196b888261190e565b975061197683611926565b925050600181019050611957565b5085935050505092915050565b600060208201905081810360008301526119ab8184611933565b905092915050565b600082825260208201905092915050565b7f56616c696461746f725365743a2073656e646572206973206e6f742076616c6960008201527f6461746f72000000000000000000000000000000000000000000000000000000602082015250565b6000611a206025836119b3565b9150611a2b826119c4565b604082019050919050565b60006020820190508181036000830152611a4f81611a13565b9050919050565b7f56616c696461746f725365743a206e6f7420696e206d61696e74656e616e6365600082015250565b6000611a8c6020836119b3565b9150611a9782611a56565b602082019050919050565b60006020820190508181036000830152611abb81611a7f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611afc826116ed565b915060008203611b0f57611b0e611ac2565b5b600182039050919050565b6000611b25826116ed565b9150611b30836116ed565b9250828202611b3e816116ed565b91508282048414831517611b5557611b54611ac2565b5b5092915050565b6000611b67826116ed565b9150611b72836116ed565b9250828201905080821115611b8a57611b89611ac2565b5b92915050565b7f56616c696461746f725365743a207a65726f2063616e64696461746500000000600082015250565b6000611bc6601c836119b3565b9150611bd182611b90565b602082019050919050565b60006020820190508181036000830152611bf581611bb9565b9050919050565b7f56616c696461746f725365743a2063616e6469646174652069732076616c696460008201527f61746f7200000000000000000000000000000000000000000000000000000000602082015250565b6000611c586024836119b3565b9150611c6382611bfc565b604082019050919050565b60006020820190508181036000830152611c8781611c4b565b9050919050565b6000611c99826116ed565b9150611ca4836116ed565b9250828203905081811115611cbc57611cbb611ac2565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611cfc826116ed565b9150611d07836116ed565b925082611d1757611d16611cc2565b5b828204905092915050565b7f56616c696461746f725365743a20616c726561647920696e206d61696e74656e60008201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b6000611d7e6024836119b3565b9150611d8982611d22565b604082019050919050565b60006020820190508181036000830152611dad81611d71565b9050919050565b7f56616c696461746f725365743a20746f6f206d616e792076616c696461746f7260008201527f7320696e206d61696e74656e616e636500000000000000000000000000000000602082015250565b6000611e106030836119b3565b9150611e1b82611db4565b604082019050919050565b60006020820190508181036000830152611e3f81611e03565b9050919050565b6000611e51826116ed565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e8357611e82611ac2565b5b600182019050919050565b7f56616c696461746f725365743a2063616e646964617465206973206e6f74207660008201527f616c696461746f72000000000000000000000000000000000000000000000000602082015250565b6000611eea6028836119b3565b9150611ef582611e8e565b604082019050919050565b60006020820190508181036000830152611f1981611edd565b9050919050565b7f56616c696461746f725365743a2063616e6e6f742072656d6f7665206c61737460008201527f2076616c696461746f7200000000000000000000000000000000000000000000602082015250565b6000611f7c602a836119b3565b9150611f8782611f20565b604082019050919050565b60006020820190508181036000830152611fab81611f6f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028110611ff257611ff1611fb2565b5b50565b600081905061200382611fe1565b919050565b600061201382611ff5565b9050919050565b60008160f81b9050919050565b60006120328261201a565b9050919050565b61204a61204582612008565b612027565b82525050565b60008160601b9050919050565b600061206882612050565b9050919050565b600061207a8261205d565b9050919050565b61209261208d82611746565b61206f565b82525050565b60006120a48285612039565b6001820191506120b48284612081565b6014820191508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f56616c696461746f725365743a2070726f706f73616c20657865637574656400600082015250565b6000612158601f836119b3565b915061216382612122565b602082019050919050565b600060208201905081810360008301526121878161214b565b9050919050565b7f56616c696461746f725365743a20616c726561647920766f7465640000000000600082015250565b60006121c4601b836119b3565b91506121cf8261218e565b602082019050919050565b600060208201905081810360008301526121f3816121b7565b9050919050565b61220381611746565b82525050565b600060608201905061221e60008301866121fa565b61222b60208301856116f7565b61223860408301846116f7565b949350505050565b7f56616c696461746f725365743a207a65726f2076616c696461746f7200000000600082015250565b6000612276601c836119b3565b915061228182612240565b602082019050919050565b600060208201905081810360008301526122a581612269565b9050919050565b7f56616c696461746f725365743a206475706c69636174652076616c696461746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b60006123086021836119b3565b9150612313826122ac565b604082019050919050565b60006020820190508181036000830152612337816122fb565b9050919050565b7f56616c696461746f725365743a2076616c696461746f72206e6f7420666f756e60008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b600061239a6021836119b3565b91506123a58261233e565b604082019050919050565b600060208201905081810360008301526123c98161238d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212205b825b33b91a6b548e067adf4c576ee2301f5f03372edcd2dccfc1e17b0231ff64736f6c634300081e0033"

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
// Solidity: function proposalId(uint8 operation, address candidate) pure returns(bytes32)
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
// Solidity: function proposalId(uint8 operation, address candidate) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetSession) ProposalId(operation uint8, candidate common.Address) ([32]byte, error) {
	return _ValidatorSet.Contract.ProposalId(&_ValidatorSet.CallOpts, operation, candidate)
}

// ProposalId is a free data retrieval call binding the contract method 0xa8b4b210.
//
// Solidity: function proposalId(uint8 operation, address candidate) pure returns(bytes32)
func (_ValidatorSet *ValidatorSetCallerSession) ProposalId(operation uint8, candidate common.Address) ([32]byte, error) {
	return _ValidatorSet.Contract.ProposalId(&_ValidatorSet.CallOpts, operation, candidate)
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
