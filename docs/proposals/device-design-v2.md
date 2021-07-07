# A New Architecture of Edge IoT

## Motivation

In the Industry 4.0 era, IoT platforms need to support thousands of types of devices,  support AI analysis, and support automated control. To solve the above problems, many vendors or community teams have proposed their own solutions, for example EdgeX Foundry.  EdgeX Foundry is a highly flexible and scalable open source software framework that facilitates interoperability between devices and applications at the IoT Edge. Compared with KubeEdge, there are both similarities and differences.  The system capability comparison is as follows.

|                                       | KubeEdge                  | Edgex Foundry                                              |
| ------------------------------------- | ------------------------- | ---------------------------------------------------------- |
| Platform Requirements                 | CRI，K8s required         | 1G mem，3GB disk                                           |
| North API                             | MQTT                      | REST/MQTT                                                  |
| Data persistence*                     | -                         | redis/mongodb                                              |
| Control Command                       | -                         | supported                                                  |
| Data Export*                          | MQTT                      | REST/MQTT/云厂商                                           |
| Rule Engine*                          | -                         | kuiper supported                                           |
| Timing scheduling                     | -                         | Data Clean supported                                       |
| Automatic device discovery*           | -                         | supported                                                  |
| Logging                               | supported                 | deployment/orchestration tools required                    |
| Alerts & Notifications                | -                         | supported                                                  |
| Secret Store                          | supported                 | supported                                                  |
| Collaboration between cloud and edge* | supported                 | -                                                          |
| Suppported IoT protocols*             | Modbus, OPC-UA, Bluetooth | more than a dozen protocols such as Modbus, rest, snmp etc |

In a word, with the support of the IoT vendor such as Dell, hp, IBM etc and complete system capability,  EdgeX has more common application scenarios.In addtion,  compared with EdgeX, KubeEdge lacks some important capability, but cloud-native management and collaboration between cloud and edge are unique advantages to it.  To improve the use of KubeEdge in the IoT scenario,  this design doc proposes a new architecture of Edge IoT to KubeEdge IoT sig.



## Goals

KubeEdge manages the device logic model uniformly, and is responsible for the status update of the device model and the addition and deletion of devices. On the one hand, we hope the edge gateway platform which has been built before can quickly access KubeEdge so that it is compatible with the existing ecosystem. On the other hand, KubeEdge-IoT can develop its own end-side device ecology, providing a lightweight and simple access method, and at the same time manage devices in a more cloud-native management mode, and enable device-edge-cloud collaboration.

1. Supports the access of multiple IoT gateways
2. KubeEdge provides data persistence, data query and forwarding, and supports rule engine