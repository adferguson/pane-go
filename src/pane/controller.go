// Participatory Networking. Copyright (C) 2012-2013 Brown University
//
// Author: Andrew Ferguson <adf@cs.brown.edu>
// Author: Arjun Guha <arjun@cs.brown.edu>
//

package pane

import (
    "runtime/pprof"
    "goof/controller"
    "goof/of"
    "log"
)

func NewSwitch(sw *controller.Switch) {
    defer func() {
        pprof.StopCPUProfile()
        recover()
    }()

    // Learning switch
    routes := make(map[[of.EthAlen]uint8]uint16, 1000)

    sw.HandlePacketIn = func(msg *of.PacketIn) {
        routes[msg.EthFrame.SrcMAC] = msg.InPort
        outPort, found := routes[msg.EthFrame.DstMAC]
        if !found {
            err := sw.Send(&of.FlowMod{
                Xid: msg.Xid,
                Match: of.Match{
                    Wildcards: of.FwAll ^ of.FwDlSrc ^ of.FwDlDst,
                    DlSrc:     msg.EthFrame.SrcMAC,
                    DlDst:     msg.EthFrame.DstMAC},
                BufferId:    msg.BufferId,
                Flags:       of.FCAdd,
                HardTimeout: 5,
                Actions:     []of.Action{&of.ActionOutput{of.PortFlood, 0}}})
            if err != nil {
                log.Printf("Erroring sending: %v", err)
            }
            log.Printf("flooding %v", msg.EthFrame.EthernetHeader)
        } else {
            err := sw.Send(&of.FlowMod{
                Xid: msg.Xid,
                Match: of.Match{
                    Wildcards: of.FwAll ^ of.FwDlSrc ^ of.FwDlDst,
                    DlSrc:     msg.EthFrame.SrcMAC,
                    DlDst:     msg.EthFrame.DstMAC},
                BufferId:    msg.BufferId,
                Flags:       of.FCAdd,
                HardTimeout: 60,
                Actions:     []of.Action{&of.ActionOutput{outPort, 0}}})
            if err != nil {
                log.Printf("Erroring sending: %v", err)
            }

        }
    }
    sw.HandleSwitchFeatures = func(msg *of.SwitchFeatures) {
        log.Printf("Datapath %x online", msg.DatapathId)
    }

    sw.HandlePortStatus = func(msg *of.PortStatus) {
        // silently ignore
    }

    sw.Serve()
}
