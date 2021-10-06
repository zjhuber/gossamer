package network

import (
	"fmt"
	"github.com/ChainSafe/gossamer/lib/common"
	"github.com/ChainSafe/gossamer/lib/common/optional"
	"github.com/ChainSafe/gossamer/lib/scale"
	scale2 "github.com/ChainSafe/gossamer/pkg/scale"
)

// Pair is a pair of arbitrary bytes.
type Pair struct {
	first  []byte
	second []byte
}

// LightRequest is all possible light client related requests.
type LightRequest struct {
	RmtCallRequest      *RemoteCallRequest
	RmtReadRequest      *RemoteReadRequest
	RmtHeaderRequest    *RemoteHeaderRequest
	RmtReadChildRequest *RemoteReadChildRequest
	RmtChangesRequest   *RemoteChangesRequestNew
}

// Request is a scale compatible struct containing all possible light client related requests.
type Request struct {
	RmtCallRequest      RemoteCallRequest
	RmtReadRequest      RemoteReadRequest
	RmtHeaderRequest    RemoteHeaderRequest
	RmtReadChildRequest RemoteReadChildRequest
	RmtChangesRequest   RemoteChangesRequestNew
}

// NewLightRequest returns a new LightRequest
func NewLightRequest() *LightRequest {
	rcr := NewRemoteChangesRequestNew()
	return &LightRequest{
		RmtCallRequest:      NewRemoteCallRequest(),
		RmtReadRequest:      NewRemoteReadRequest(),
		RmtHeaderRequest:    NewRemoteHeaderRequest(),
		RmtReadChildRequest: NewRemoteReadChildRequest(),
		RmtChangesRequest:   &rcr,
	}
}

func newRequest() *Request {
	return &Request{
		RmtCallRequest:      *NewRemoteCallRequest(),
		RmtReadRequest:      *NewRemoteReadRequest(),
		RmtHeaderRequest:    *NewRemoteHeaderRequest(),
		RmtReadChildRequest: *NewRemoteReadChildRequest(),
		RmtChangesRequest:   NewRemoteChangesRequestNew(),
	}
}

// SubProtocol returns the light sub-protocol
func (l *LightRequest) SubProtocol() string {
	return lightID
}

// Encode encodes a LightRequest message using SCALE and appends the type byte to the start
func (l *LightRequest) Encode() ([]byte, error) {
	req := Request{
		RmtCallRequest:      *l.RmtCallRequest,
		RmtReadRequest:      *l.RmtReadRequest,
		RmtHeaderRequest:    *l.RmtHeaderRequest,
		RmtReadChildRequest: *l.RmtReadChildRequest,
		RmtChangesRequest:   *l.RmtChangesRequest,
	}
	return scale2.Marshal(req)
}

// Decode the message into a LightRequest, it assumes the type byte has been removed
func (l *LightRequest) Decode(in []byte) error {
	msg := newRequest()
	err := scale2.Unmarshal(in, msg)
	if err != nil {
		return err
	}

	l.RmtCallRequest = &msg.RmtCallRequest
	l.RmtReadRequest = &msg.RmtReadRequest
	l.RmtHeaderRequest = &msg.RmtHeaderRequest
	l.RmtReadChildRequest = &msg.RmtReadChildRequest
	l.RmtChangesRequest = &msg.RmtChangesRequest
	return nil
}

// String formats a LightRequest as a string
func (l LightRequest) String() string {
	return fmt.Sprintf(
		"RemoteCallRequest=%s RemoteReadRequest=%s RemoteHeaderRequest=%s "+
			"RemoteReadChildRequest=%s RemoteChangesRequest=%s",
		l.RmtCallRequest, l.RmtReadRequest, l.RmtHeaderRequest, l.RmtReadChildRequest, l.RmtChangesRequest)
}

// LightResponse is all possible light client response messages.
type LightResponse struct {
	RmtCallResponse   *RemoteCallResponse
	RmtReadResponse   *RemoteReadResponse
	RmtHeaderResponse *RemoteHeaderResponse
	RmtChangeResponse *RemoteChangesResponse
}

// NewLightResponse returns a new LightResponse
func NewLightResponse() *LightResponse {
	return &LightResponse{
		RmtCallResponse:   new(RemoteCallResponse),
		RmtReadResponse:   new(RemoteReadResponse),
		RmtHeaderResponse: new(RemoteHeaderResponse),
		RmtChangeResponse: new(RemoteChangesResponse),
	}
}

// SubProtocol returns the light sub-protocol
func (l *LightResponse) SubProtocol() string {
	return lightID
}

// Encode encodes a LightResponse message using SCALE and appends the type byte to the start
func (l *LightResponse) Encode() ([]byte, error) {
	return scale.Encode(l)
}

// Decode the message into a LightResponse, it assumes the type byte has been removed
func (l *LightResponse) Decode(in []byte) error {
	msg, err := scale.Decode(in, l)
	if err != nil {
		return err
	}

	l.RmtCallResponse = msg.(*LightResponse).RmtCallResponse
	l.RmtReadResponse = msg.(*LightResponse).RmtReadResponse
	l.RmtHeaderResponse = msg.(*LightResponse).RmtHeaderResponse
	l.RmtChangeResponse = msg.(*LightResponse).RmtChangeResponse
	return nil
}

// String formats a RemoteReadRequest as a string
func (l LightResponse) String() string {
	return fmt.Sprintf(
		"RemoteCallResponse=%s RemoteReadResponse=%s RemoteHeaderResponse=%s RemoteChangesResponse=%s",
		l.RmtCallResponse, l.RmtReadResponse, l.RmtHeaderResponse, l.RmtChangeResponse)
}

// RemoteCallRequest ...
type RemoteCallRequest struct {
	Block  []byte
	Method string
	Data   []byte
}

// NewLightRequest returns a new LightRequest
func NewRemoteCallRequest() *RemoteCallRequest{
	return &RemoteCallRequest{
		Block: []byte{},
		Method: "",
		Data: []byte{},
	}
}

// RemoteReadRequest ...
type RemoteReadRequest struct {
	Block []byte
	Keys  [][]byte
}

func NewRemoteReadRequest() *RemoteReadRequest{
	return &RemoteReadRequest{
		Block: []byte{},
	}
}

// RemoteReadChildRequest ...
type RemoteReadChildRequest struct {
	Block      []byte
	StorageKey []byte
	Keys       [][]byte
}

func NewRemoteReadChildRequest() *RemoteReadChildRequest{
	return &RemoteReadChildRequest{
		Block: []byte{},
		StorageKey: []byte{},
	}
}

// RemoteHeaderRequest ...
type RemoteHeaderRequest struct {
	Block []byte
}

func NewRemoteHeaderRequest() *RemoteHeaderRequest{
	return &RemoteHeaderRequest{
		Block: []byte{},
	}
}

// RemoteChangesRequest ...
type RemoteChangesRequest struct {
	FirstBlock *optional.Hash
	LastBlock  *optional.Hash
	Min        []byte
	Max        []byte
	StorageKey *optional.Bytes
	key        []byte
}

// RemoteChangesRequest ...
type RemoteChangesRequestNew struct {
	FirstBlock *common.Hash
	LastBlock  *common.Hash
	Min        []byte
	Max        []byte
	StorageKey *[]byte
	key        []byte
}

// NewRemoteChangesRequest returns a new RemoteChangesRequest
func NewRemoteChangesRequest() *RemoteChangesRequest {
	return &RemoteChangesRequest{
		FirstBlock: optional.NewHash(false, common.Hash{}),
		LastBlock:  optional.NewHash(false, common.Hash{}),
		Min:        []byte{},
		Max:        []byte{},
		StorageKey: optional.NewBytes(false, nil),
	}
}

// NewRemoteChangesRequest returns a new RemoteChangesRequest
func NewRemoteChangesRequestNew() RemoteChangesRequestNew {
	return RemoteChangesRequestNew{
		FirstBlock: nil,
		LastBlock:  nil,
		Min:        []byte{},
		Max:        []byte{},
		StorageKey: nil,
	}
}

// RemoteCallResponse ...
type RemoteCallResponse struct {
	Proof []byte
}

// RemoteReadResponse ...
type RemoteReadResponse struct {
	Proof []byte
}

// RemoteHeaderResponse ...
type RemoteHeaderResponse struct {
	Header []*optional.Header
	proof  []byte
}

// RemoteChangesResponse ...
type RemoteChangesResponse struct {
	Max        []byte
	Proof      [][]byte
	Roots      [][]Pair
	RootsProof []byte
}

// String formats a RemoteCallRequest as a string
func (rc *RemoteCallRequest) String() string {
	return fmt.Sprintf("Block =%s method=%s Data=%s",
		string(rc.Block), rc.Method, string(rc.Data))
}

// String formats a RemoteChangesRequest as a string
func (rc *RemoteChangesRequest) String() string {
	return fmt.Sprintf("FirstBlock =%s LastBlock=%s Min=%s Max=%s Storagekey=%s key=%s",
		rc.FirstBlock.String(),
		rc.LastBlock.String(),
		string(rc.Min),
		string(rc.Max),
		rc.StorageKey.String(),
		string(rc.key),
	)
}

// String formats a RemoteChangesRequest as a string
func (rc *RemoteChangesRequestNew) String() string {
	first := common.Hash{}
	last := common.Hash{}
	storageKey := []byte{0}
	if rc.FirstBlock != nil {
		first = *rc.FirstBlock
	}
	if rc.LastBlock != nil {
		last = *rc.LastBlock
	}
	if rc.StorageKey != nil {
		storageKey = *rc.StorageKey
	}
	return fmt.Sprintf("FirstBlock =%s LastBlock=%s Min=%s Max=%s Storagekey=%s key=%s",
		first,
		last,
		string(rc.Min),
		string(rc.Max),
		storageKey,
		string(rc.key),
	)
}

// String formats a RemoteHeaderRequest as a string
func (rh *RemoteHeaderRequest) String() string {
	return fmt.Sprintf("Block =%s", string(rh.Block))
}

// String formats a RemoteReadRequest as a string
func (rr *RemoteReadRequest) String() string {
	return fmt.Sprintf("Block =%s", string(rr.Block))
}

// String formats a RemoteReadChildRequest as a string
func (rr *RemoteReadChildRequest) String() string {
	var strKeys []string
	for _, v := range rr.Keys {
		strKeys = append(strKeys, string(v))
	}
	return fmt.Sprintf("Block =%s StorageKey=%s Keys=%v",
		string(rr.Block),
		string(rr.StorageKey),
		strKeys,
	)
}

// String formats a RemoteCallResponse as a string
func (rc *RemoteCallResponse) String() string {
	return fmt.Sprintf("Proof =%s", string(rc.Proof))
}

// String formats a RemoteChangesResponse as a string
func (rc *RemoteChangesResponse) String() string {
	var strRoots []string
	var strProof []string
	for _, v := range rc.Proof {
		strProof = append(strProof, string(v))
	}
	for _, v := range rc.Roots {
		for _, p := range v {
			strRoots = append(strRoots, string(p.first), string(p.second))
		}
	}
	return fmt.Sprintf("Max =%s Proof =%s Roots=%v RootsProof=%s",
		string(rc.Max),
		strProof,
		strRoots,
		string(rc.RootsProof),
	)
}

// String formats a RemoteReadResponse as a string
func (rr *RemoteReadResponse) String() string {
	return fmt.Sprintf("Proof =%s", string(rr.Proof))
}

// String formats a RemoteHeaderResponse as a string
func (rh *RemoteHeaderResponse) String() string {
	return fmt.Sprintf("Header =%+v Proof =%s", rh.Header, string(rh.proof))
}

func remoteCallResp(req *RemoteCallRequest) (*RemoteCallResponse, error) {
	return &RemoteCallResponse{}, nil
}
func remoteChangeResp(req *RemoteChangesRequest) (*RemoteChangesResponse, error) {
	return &RemoteChangesResponse{}, nil
}
func remoteChangeRespNew(req *RemoteChangesRequestNew) (*RemoteChangesResponse, error) {
	return &RemoteChangesResponse{}, nil
}
func remoteHeaderResp(req *RemoteHeaderRequest) (*RemoteHeaderResponse, error) {
	return &RemoteHeaderResponse{}, nil
}
func remoteReadChildResp(req *RemoteReadChildRequest) (*RemoteReadResponse, error) {
	return &RemoteReadResponse{}, nil
}
func remoteReadResp(req *RemoteReadRequest) (*RemoteReadResponse, error) {
	return &RemoteReadResponse{}, nil
}
