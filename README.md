# Reinforcement Learning to Achieve High Scalability Using Multi-Leader Based BFT Consensus Protocols

## Introduction

During our research, we explored single-leader protocols and a variety of multi-leader protocols, ultimately selecting four multi-leader protocols—**RCC**, **DQBFT**, **Mir BFT**, and **Dispel**—to achieve consensus and scalability in a distributed system environment. Each protocol has unique advantages:

- **DQBFT**: Achieves consensus through a two-phase process: one for communication and another to achieve a global order. DQBFT introduces trusted subsystems, reducing the communication layers from 3 (as in PBFT and Mir BFT) to 2, eliminating the prepare phase. Replicas achieve consensus with only `2f+1` nodes (where `f` represents the confirmation required by the primary). Local order is maintained across replicas, with global order created by one of the replicas to achieve consensus. Key parameters for performance include network bandwidth, request size, response size, and the number of replicas, though high replica counts may increase latency.

- **Mir BFT**: Builds upon PBFT, introducing **multi-primary nodes** where every node participates in both prepare and commit phases, requiring `2f+1` confirmations to proceed. This distributed approach to responsibilities enhances fault tolerance and throughput, particularly under fluctuating network conditions and heavy system loads.

- **Dispel**: Optimizes communication between nodes, reducing overhead and latency by employing a pipeline architecture in the communication phase. This architecture allows faster decision-making through efficient message propagation and validation, which minimizes delays.

- **RCC**: Designed for concurrent requests and decisions, RCC allows multiple operations to proceed simultaneously without compromising system consistency. Each replica maintains a local order, which is used to calculate a global order, thereby supporting efficient concurrency crucial for improving distributed system responsiveness.

## Learning Algorithm Parameters

We identified several parameters that influence the performance of each protocol. These factors can be used to train our reinforcement learning (RL) model, allowing it to shift between protocols based on their strengths to achieve scalable consensus, high throughput, and support for an increasing number of replicas. Below, we outline the primary parameters considered for training; additional parameters may be added as research progresses.

### State Space

#### Workload-Based Parameters

- **Request Size**: `1KB` to `100KB` (Range: `[1KB, 10KB, 50KB, 100KB]`)
- **Reply Size**: Same as request size, depending on system response.
- **Execution Overhead**: Compared to single-leader protocols (e.g., `0.5x`, `1x`, `5x`, where `x` is the baseline overhead for single-leader protocols).

#### Fault Tolerance Parameters

- **Absence from Participation**: Number of instances absent from a consensus round or phase (Range: `1-5 instances`).
- **Message Delay**: The average or maximum delay for message delivery (Range: `[1ms, 10ms, 50ms, 100ms]`).
- **Rate of Message Loss**: The percentage of messages lost (Range: `[0%, 10%, 20%, 30%, 40%, 50%]`).

#### System Hardware Configuration

- **Network Bandwidth**: `[10Mbps, 50Mbps, 100Mbps]`
- **CPU Utilization**: `[10%, 30%, 50%, 70%]`
- **Memory Consumption**: `[500MB, 1GB, 2GB]`
- **Trusted Subsystems**: Binary indicator for trusted systems (`0` = not trusted, `1` = trusted).

#### Output Parameters (Input to Reward Function)

- **Degree of Concurrency**: Maximum number of concurrent instances processing in a given duration (Range: `[1, 2, 5, 10]`).
- **Time to Process Each Request/Query**: `[10ms, 50ms, 100ms, 200ms]`

### Action Space

The RL model will leverage the following multi-leader based BFT protocols to maximize scalability and throughput:

1. **RCC**
2. **Mir BFT**
3. **Dispel**
4. **DQBFT**

### Training Scenarios

Below is a sample scenario for training the RL model to adaptively select the optimal protocol based on current parameters:

#### Scenario 1 (Minimal)

- **Request Size**: `10KB`
- **Reply Size**: `5KB`
- **Execution Overhead**: `1x`
- **Absence from Participation**: `1 instance`
- **Message Delay**: `10ms`
- **Rate of Message Loss**: `10%`
- **Network Bandwidth**: `50Mbps`
- **CPU Utilization**: `30%`
- **Memory Consumption**: `1GB`
- **Trusted Subsystems**: `1`
- **Degree of Concurrency**: `2`
- **Request/Query Rate**: `100m/s`

## Future Work

Moving forward, we aim to expand the parameter set to incorporate additional factors, thereby enhancing the model's learning and adaptability. By training the RL model with these parameters, we aim to achieve scalable consensus across transactions by dynamically utilizing the best phase of each protocol as needed.
