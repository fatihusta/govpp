// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package nat44_ed

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "git.fd.io/govpp.git/binapi/vpe"
)

// RPCService defines RPC service nat44_ed.
type RPCService interface {
	Nat44AddDelAddressRange(ctx context.Context, in *Nat44AddDelAddressRange) (*Nat44AddDelAddressRangeReply, error)
	Nat44AddDelIdentityMapping(ctx context.Context, in *Nat44AddDelIdentityMapping) (*Nat44AddDelIdentityMappingReply, error)
	Nat44AddDelInterfaceAddr(ctx context.Context, in *Nat44AddDelInterfaceAddr) (*Nat44AddDelInterfaceAddrReply, error)
	Nat44AddDelLbStaticMapping(ctx context.Context, in *Nat44AddDelLbStaticMapping) (*Nat44AddDelLbStaticMappingReply, error)
	Nat44AddDelStaticMapping(ctx context.Context, in *Nat44AddDelStaticMapping) (*Nat44AddDelStaticMappingReply, error)
	Nat44AddDelStaticMappingV2(ctx context.Context, in *Nat44AddDelStaticMappingV2) (*Nat44AddDelStaticMappingV2Reply, error)
	Nat44AddressDump(ctx context.Context, in *Nat44AddressDump) (RPCService_Nat44AddressDumpClient, error)
	Nat44DelSession(ctx context.Context, in *Nat44DelSession) (*Nat44DelSessionReply, error)
	Nat44DelUser(ctx context.Context, in *Nat44DelUser) (*Nat44DelUserReply, error)
	Nat44EdPluginEnableDisable(ctx context.Context, in *Nat44EdPluginEnableDisable) (*Nat44EdPluginEnableDisableReply, error)
	Nat44EdSetFqOptions(ctx context.Context, in *Nat44EdSetFqOptions) (*Nat44EdSetFqOptionsReply, error)
	Nat44EdShowFqOptions(ctx context.Context, in *Nat44EdShowFqOptions) (*Nat44EdShowFqOptionsReply, error)
	Nat44ForwardingEnableDisable(ctx context.Context, in *Nat44ForwardingEnableDisable) (*Nat44ForwardingEnableDisableReply, error)
	Nat44ForwardingIsEnabled(ctx context.Context, in *Nat44ForwardingIsEnabled) (*Nat44ForwardingIsEnabledReply, error)
	Nat44IdentityMappingDump(ctx context.Context, in *Nat44IdentityMappingDump) (RPCService_Nat44IdentityMappingDumpClient, error)
	Nat44InterfaceAddDelFeature(ctx context.Context, in *Nat44InterfaceAddDelFeature) (*Nat44InterfaceAddDelFeatureReply, error)
	Nat44InterfaceAddDelOutputFeature(ctx context.Context, in *Nat44InterfaceAddDelOutputFeature) (*Nat44InterfaceAddDelOutputFeatureReply, error)
	Nat44InterfaceAddrDump(ctx context.Context, in *Nat44InterfaceAddrDump) (RPCService_Nat44InterfaceAddrDumpClient, error)
	Nat44InterfaceDump(ctx context.Context, in *Nat44InterfaceDump) (RPCService_Nat44InterfaceDumpClient, error)
	Nat44InterfaceOutputFeatureDump(ctx context.Context, in *Nat44InterfaceOutputFeatureDump) (RPCService_Nat44InterfaceOutputFeatureDumpClient, error)
	Nat44LbStaticMappingAddDelLocal(ctx context.Context, in *Nat44LbStaticMappingAddDelLocal) (*Nat44LbStaticMappingAddDelLocalReply, error)
	Nat44LbStaticMappingDump(ctx context.Context, in *Nat44LbStaticMappingDump) (RPCService_Nat44LbStaticMappingDumpClient, error)
	Nat44PluginEnableDisable(ctx context.Context, in *Nat44PluginEnableDisable) (*Nat44PluginEnableDisableReply, error)
	Nat44SessionCleanup(ctx context.Context, in *Nat44SessionCleanup) (*Nat44SessionCleanupReply, error)
	Nat44SetSessionLimit(ctx context.Context, in *Nat44SetSessionLimit) (*Nat44SetSessionLimitReply, error)
	Nat44ShowRunningConfig(ctx context.Context, in *Nat44ShowRunningConfig) (*Nat44ShowRunningConfigReply, error)
	Nat44StaticMappingDump(ctx context.Context, in *Nat44StaticMappingDump) (RPCService_Nat44StaticMappingDumpClient, error)
	Nat44UserDump(ctx context.Context, in *Nat44UserDump) (RPCService_Nat44UserDumpClient, error)
	Nat44UserSessionDump(ctx context.Context, in *Nat44UserSessionDump) (RPCService_Nat44UserSessionDumpClient, error)
	NatControlPing(ctx context.Context, in *NatControlPing) (*NatControlPingReply, error)
	NatGetAddrAndPortAllocAlg(ctx context.Context, in *NatGetAddrAndPortAllocAlg) (*NatGetAddrAndPortAllocAlgReply, error)
	NatGetMssClamping(ctx context.Context, in *NatGetMssClamping) (*NatGetMssClampingReply, error)
	NatGetTimeouts(ctx context.Context, in *NatGetTimeouts) (*NatGetTimeoutsReply, error)
	NatHaFlush(ctx context.Context, in *NatHaFlush) (*NatHaFlushReply, error)
	NatHaGetFailover(ctx context.Context, in *NatHaGetFailover) (*NatHaGetFailoverReply, error)
	NatHaGetListener(ctx context.Context, in *NatHaGetListener) (*NatHaGetListenerReply, error)
	NatHaResync(ctx context.Context, in *NatHaResync) (*NatHaResyncReply, error)
	NatHaSetFailover(ctx context.Context, in *NatHaSetFailover) (*NatHaSetFailoverReply, error)
	NatHaSetListener(ctx context.Context, in *NatHaSetListener) (*NatHaSetListenerReply, error)
	NatIpfixEnableDisable(ctx context.Context, in *NatIpfixEnableDisable) (*NatIpfixEnableDisableReply, error)
	NatSetAddrAndPortAllocAlg(ctx context.Context, in *NatSetAddrAndPortAllocAlg) (*NatSetAddrAndPortAllocAlgReply, error)
	NatSetLogLevel(ctx context.Context, in *NatSetLogLevel) (*NatSetLogLevelReply, error)
	NatSetMssClamping(ctx context.Context, in *NatSetMssClamping) (*NatSetMssClampingReply, error)
	NatSetTimeouts(ctx context.Context, in *NatSetTimeouts) (*NatSetTimeoutsReply, error)
	NatSetWorkers(ctx context.Context, in *NatSetWorkers) (*NatSetWorkersReply, error)
	NatShowConfig(ctx context.Context, in *NatShowConfig) (*NatShowConfigReply, error)
	NatShowConfig2(ctx context.Context, in *NatShowConfig2) (*NatShowConfig2Reply, error)
	NatWorkerDump(ctx context.Context, in *NatWorkerDump) (RPCService_NatWorkerDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Nat44AddDelAddressRange(ctx context.Context, in *Nat44AddDelAddressRange) (*Nat44AddDelAddressRangeReply, error) {
	out := new(Nat44AddDelAddressRangeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44AddDelIdentityMapping(ctx context.Context, in *Nat44AddDelIdentityMapping) (*Nat44AddDelIdentityMappingReply, error) {
	out := new(Nat44AddDelIdentityMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44AddDelInterfaceAddr(ctx context.Context, in *Nat44AddDelInterfaceAddr) (*Nat44AddDelInterfaceAddrReply, error) {
	out := new(Nat44AddDelInterfaceAddrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44AddDelLbStaticMapping(ctx context.Context, in *Nat44AddDelLbStaticMapping) (*Nat44AddDelLbStaticMappingReply, error) {
	out := new(Nat44AddDelLbStaticMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44AddDelStaticMapping(ctx context.Context, in *Nat44AddDelStaticMapping) (*Nat44AddDelStaticMappingReply, error) {
	out := new(Nat44AddDelStaticMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44AddDelStaticMappingV2(ctx context.Context, in *Nat44AddDelStaticMappingV2) (*Nat44AddDelStaticMappingV2Reply, error) {
	out := new(Nat44AddDelStaticMappingV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44AddressDump(ctx context.Context, in *Nat44AddressDump) (RPCService_Nat44AddressDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44AddressDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44AddressDumpClient interface {
	Recv() (*Nat44AddressDetails, error)
	api.Stream
}

type serviceClient_Nat44AddressDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44AddressDumpClient) Recv() (*Nat44AddressDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44AddressDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44DelSession(ctx context.Context, in *Nat44DelSession) (*Nat44DelSessionReply, error) {
	out := new(Nat44DelSessionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44DelUser(ctx context.Context, in *Nat44DelUser) (*Nat44DelUserReply, error) {
	out := new(Nat44DelUserReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EdPluginEnableDisable(ctx context.Context, in *Nat44EdPluginEnableDisable) (*Nat44EdPluginEnableDisableReply, error) {
	out := new(Nat44EdPluginEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EdSetFqOptions(ctx context.Context, in *Nat44EdSetFqOptions) (*Nat44EdSetFqOptionsReply, error) {
	out := new(Nat44EdSetFqOptionsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44EdShowFqOptions(ctx context.Context, in *Nat44EdShowFqOptions) (*Nat44EdShowFqOptionsReply, error) {
	out := new(Nat44EdShowFqOptionsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44ForwardingEnableDisable(ctx context.Context, in *Nat44ForwardingEnableDisable) (*Nat44ForwardingEnableDisableReply, error) {
	out := new(Nat44ForwardingEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44ForwardingIsEnabled(ctx context.Context, in *Nat44ForwardingIsEnabled) (*Nat44ForwardingIsEnabledReply, error) {
	out := new(Nat44ForwardingIsEnabledReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Nat44IdentityMappingDump(ctx context.Context, in *Nat44IdentityMappingDump) (RPCService_Nat44IdentityMappingDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44IdentityMappingDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44IdentityMappingDumpClient interface {
	Recv() (*Nat44IdentityMappingDetails, error)
	api.Stream
}

type serviceClient_Nat44IdentityMappingDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44IdentityMappingDumpClient) Recv() (*Nat44IdentityMappingDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44IdentityMappingDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44InterfaceAddDelFeature(ctx context.Context, in *Nat44InterfaceAddDelFeature) (*Nat44InterfaceAddDelFeatureReply, error) {
	out := new(Nat44InterfaceAddDelFeatureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44InterfaceAddDelOutputFeature(ctx context.Context, in *Nat44InterfaceAddDelOutputFeature) (*Nat44InterfaceAddDelOutputFeatureReply, error) {
	out := new(Nat44InterfaceAddDelOutputFeatureReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44InterfaceAddrDump(ctx context.Context, in *Nat44InterfaceAddrDump) (RPCService_Nat44InterfaceAddrDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44InterfaceAddrDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44InterfaceAddrDumpClient interface {
	Recv() (*Nat44InterfaceAddrDetails, error)
	api.Stream
}

type serviceClient_Nat44InterfaceAddrDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44InterfaceAddrDumpClient) Recv() (*Nat44InterfaceAddrDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44InterfaceAddrDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44InterfaceDump(ctx context.Context, in *Nat44InterfaceDump) (RPCService_Nat44InterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44InterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44InterfaceDumpClient interface {
	Recv() (*Nat44InterfaceDetails, error)
	api.Stream
}

type serviceClient_Nat44InterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44InterfaceDumpClient) Recv() (*Nat44InterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44InterfaceDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44InterfaceOutputFeatureDump(ctx context.Context, in *Nat44InterfaceOutputFeatureDump) (RPCService_Nat44InterfaceOutputFeatureDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44InterfaceOutputFeatureDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44InterfaceOutputFeatureDumpClient interface {
	Recv() (*Nat44InterfaceOutputFeatureDetails, error)
	api.Stream
}

type serviceClient_Nat44InterfaceOutputFeatureDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44InterfaceOutputFeatureDumpClient) Recv() (*Nat44InterfaceOutputFeatureDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44InterfaceOutputFeatureDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44LbStaticMappingAddDelLocal(ctx context.Context, in *Nat44LbStaticMappingAddDelLocal) (*Nat44LbStaticMappingAddDelLocalReply, error) {
	out := new(Nat44LbStaticMappingAddDelLocalReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44LbStaticMappingDump(ctx context.Context, in *Nat44LbStaticMappingDump) (RPCService_Nat44LbStaticMappingDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44LbStaticMappingDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44LbStaticMappingDumpClient interface {
	Recv() (*Nat44LbStaticMappingDetails, error)
	api.Stream
}

type serviceClient_Nat44LbStaticMappingDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44LbStaticMappingDumpClient) Recv() (*Nat44LbStaticMappingDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44LbStaticMappingDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44PluginEnableDisable(ctx context.Context, in *Nat44PluginEnableDisable) (*Nat44PluginEnableDisableReply, error) {
	out := new(Nat44PluginEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44SessionCleanup(ctx context.Context, in *Nat44SessionCleanup) (*Nat44SessionCleanupReply, error) {
	out := new(Nat44SessionCleanupReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44SetSessionLimit(ctx context.Context, in *Nat44SetSessionLimit) (*Nat44SetSessionLimitReply, error) {
	out := new(Nat44SetSessionLimitReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44ShowRunningConfig(ctx context.Context, in *Nat44ShowRunningConfig) (*Nat44ShowRunningConfigReply, error) {
	out := new(Nat44ShowRunningConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Nat44StaticMappingDump(ctx context.Context, in *Nat44StaticMappingDump) (RPCService_Nat44StaticMappingDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44StaticMappingDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44StaticMappingDumpClient interface {
	Recv() (*Nat44StaticMappingDetails, error)
	api.Stream
}

type serviceClient_Nat44StaticMappingDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44StaticMappingDumpClient) Recv() (*Nat44StaticMappingDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44StaticMappingDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44UserDump(ctx context.Context, in *Nat44UserDump) (RPCService_Nat44UserDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44UserDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44UserDumpClient interface {
	Recv() (*Nat44UserDetails, error)
	api.Stream
}

type serviceClient_Nat44UserDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44UserDumpClient) Recv() (*Nat44UserDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44UserDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) Nat44UserSessionDump(ctx context.Context, in *Nat44UserSessionDump) (RPCService_Nat44UserSessionDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_Nat44UserSessionDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_Nat44UserSessionDumpClient interface {
	Recv() (*Nat44UserSessionDetails, error)
	api.Stream
}

type serviceClient_Nat44UserSessionDumpClient struct {
	api.Stream
}

func (c *serviceClient_Nat44UserSessionDumpClient) Recv() (*Nat44UserSessionDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *Nat44UserSessionDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) NatControlPing(ctx context.Context, in *NatControlPing) (*NatControlPingReply, error) {
	out := new(NatControlPingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatGetAddrAndPortAllocAlg(ctx context.Context, in *NatGetAddrAndPortAllocAlg) (*NatGetAddrAndPortAllocAlgReply, error) {
	out := new(NatGetAddrAndPortAllocAlgReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatGetMssClamping(ctx context.Context, in *NatGetMssClamping) (*NatGetMssClampingReply, error) {
	out := new(NatGetMssClampingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatGetTimeouts(ctx context.Context, in *NatGetTimeouts) (*NatGetTimeoutsReply, error) {
	out := new(NatGetTimeoutsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatHaFlush(ctx context.Context, in *NatHaFlush) (*NatHaFlushReply, error) {
	out := new(NatHaFlushReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatHaGetFailover(ctx context.Context, in *NatHaGetFailover) (*NatHaGetFailoverReply, error) {
	out := new(NatHaGetFailoverReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatHaGetListener(ctx context.Context, in *NatHaGetListener) (*NatHaGetListenerReply, error) {
	out := new(NatHaGetListenerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatHaResync(ctx context.Context, in *NatHaResync) (*NatHaResyncReply, error) {
	out := new(NatHaResyncReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatHaSetFailover(ctx context.Context, in *NatHaSetFailover) (*NatHaSetFailoverReply, error) {
	out := new(NatHaSetFailoverReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatHaSetListener(ctx context.Context, in *NatHaSetListener) (*NatHaSetListenerReply, error) {
	out := new(NatHaSetListenerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatIpfixEnableDisable(ctx context.Context, in *NatIpfixEnableDisable) (*NatIpfixEnableDisableReply, error) {
	out := new(NatIpfixEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatSetAddrAndPortAllocAlg(ctx context.Context, in *NatSetAddrAndPortAllocAlg) (*NatSetAddrAndPortAllocAlgReply, error) {
	out := new(NatSetAddrAndPortAllocAlgReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatSetLogLevel(ctx context.Context, in *NatSetLogLevel) (*NatSetLogLevelReply, error) {
	out := new(NatSetLogLevelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatSetMssClamping(ctx context.Context, in *NatSetMssClamping) (*NatSetMssClampingReply, error) {
	out := new(NatSetMssClampingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatSetTimeouts(ctx context.Context, in *NatSetTimeouts) (*NatSetTimeoutsReply, error) {
	out := new(NatSetTimeoutsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatSetWorkers(ctx context.Context, in *NatSetWorkers) (*NatSetWorkersReply, error) {
	out := new(NatSetWorkersReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatShowConfig(ctx context.Context, in *NatShowConfig) (*NatShowConfigReply, error) {
	out := new(NatShowConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatShowConfig2(ctx context.Context, in *NatShowConfig2) (*NatShowConfig2Reply, error) {
	out := new(NatShowConfig2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) NatWorkerDump(ctx context.Context, in *NatWorkerDump) (RPCService_NatWorkerDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_NatWorkerDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_NatWorkerDumpClient interface {
	Recv() (*NatWorkerDetails, error)
	api.Stream
}

type serviceClient_NatWorkerDumpClient struct {
	api.Stream
}

func (c *serviceClient_NatWorkerDumpClient) Recv() (*NatWorkerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *NatWorkerDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}