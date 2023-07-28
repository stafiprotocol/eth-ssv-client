package utils

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type NodeHash []byte

func (nodeHash *NodeHash) String() string {
	return hex.EncodeToString(*nodeHash)
}

func NodeHashFromHexString(hexStr string) (NodeHash, error) {
	bts, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

type Node struct {
	Hash   NodeHash
	Parent *Node
	left   *Node
	right  *Node
}

func (n *Node) String() string {
	return hex.EncodeToString(n.Hash[:])
}

type NodeHashList []NodeHash

func (nodeHashList NodeHashList) Len() int { return len(nodeHashList) }
func (nodeHashList NodeHashList) Less(i, j int) bool {
	return bytes.Compare(nodeHashList[i], nodeHashList[j]) < 0
}
func (nodeHashList NodeHashList) Swap(i, j int) {
	nodeHashList[i], nodeHashList[j] = nodeHashList[j], nodeHashList[i]
}

// leafNodes on layers[0], rootNode on layers[len(layers)-1]
type MerkleTree struct {
	layers        [][]*Node
	leafNodeIndex map[string]int
}

func NewMerkleTree(nodeHashList NodeHashList) *MerkleTree {
	m := MerkleTree{}
	sort.Sort(nodeHashList)
	m.layers = make([][]*Node, int64(math.Ceil(float64(nodeHashList.Len())/2)+1))

	m.buildLeafNodes(nodeHashList)

	realHeight := 0
	for i := 0; i < len(m.layers)-1; i++ {
		layer := make([]*Node, int64(math.Ceil(float64(len(m.layers[i]))/2)))
		for j := 0; j < len(m.layers[i]); j = j + 2 {
			if j+1 < len(m.layers[i]) {
				cHash := ConbinedHash(m.layers[i][j].Hash, m.layers[i][j+1].Hash)
				node := Node{
					Hash:   cHash,
					Parent: nil,
					left:   m.layers[i][j],
					right:  m.layers[i][j+1],
				}
				layer[j/2] = &node
				m.layers[i][j].Parent = &node
				m.layers[i][j+1].Parent = &node
			} else {
				layer[j/2] = m.layers[i][j]
			}
		}
		m.layers[i+1] = layer
		if len(layer) == 1 {
			realHeight = i + 1
			break
		}
	}
	m.layers = m.layers[0 : realHeight+1]
	return &m
}

func (m *MerkleTree) GetLayers() [][]*Node {
	return m.layers
}

func (m *MerkleTree) GetRootHash() (hash NodeHash, err error) {
	if (len(m.layers[len(m.layers)-1])) != 1 {
		err = errors.New("invalidate tree")
	}
	hash = m.layers[len(m.layers)-1][0].Hash
	return
}

func (m *MerkleTree) buildLeafNodes(nodeHashList NodeHashList) {
	m.leafNodeIndex = make(map[string]int)
	m.layers[0] = make([]*Node, nodeHashList.Len())
	for i, nodeHash := range nodeHashList {
		node := Node{
			nodeHash,
			nil,
			nil,
			nil,
		}
		m.layers[0][i] = &node
		m.leafNodeIndex[nodeHash.String()] = i
	}
}

func (m *MerkleTree) GetProof(leafNodeHash NodeHash) ([]NodeHash, error) {
	proof := make([]NodeHash, 0)
	if index, ok := m.leafNodeIndex[leafNodeHash.String()]; ok {

		for i := 0; i < len(m.layers)-1; i++ {
			node, err := m.getPairElement(index, i)
			if err != nil {
				index = index / 2
				continue
			}
			proof = append(proof, node.Hash)
			index = index / 2
		}

		return proof, nil

	} else {
		return nil, errors.New("leafnode not exist")
	}
}

func (m *MerkleTree) getPairElement(index, layer int) (*Node, error) {
	willUseIndex := 0
	if index%2 == 0 {
		willUseIndex = index + 1
	} else {
		willUseIndex = index - 1
	}
	if willUseIndex <= len(m.layers[layer])-1 {
		return m.layers[layer][willUseIndex], nil
	} else {
		return nil, fmt.Errorf("no pair index %d ,layer %d", index, layer)
	}
}

func VerifyProof(leafNode NodeHash, proof []NodeHash, root NodeHash) bool {
	result := leafNode
	for _, p := range proof {
		result = ConbinedHash(result, p)
	}
	return bytes.EqualFold(result, root)
}

func ConbinedHash(b0, b1 NodeHash) NodeHash {
	bts := make([]byte, 64)
	if bytes.Compare(b0, b1) <= 0 { //a<=b
		copy(bts[0:], b0)
		copy(bts[32:], b1)
	} else {
		copy(bts[0:], b1)
		copy(bts[32:], b0)
	}
	h := sha3.NewLegacyKeccak256()
	h.Write(bts)
	return h.Sum(nil)
}

func GetNodeHash(index *big.Int, account common.Address, rewardAmount, depositAmount *big.Int) NodeHash {
	hash := crypto.Keccak256Hash(common.LeftPadBytes(index.Bytes(), 32), account.Bytes(), common.LeftPadBytes(rewardAmount.Bytes(), 32), common.LeftPadBytes(depositAmount.Bytes(), 32))
	return hash[:]
}
