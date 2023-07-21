# Two-Phase Commit Service
This service implements the two-phase commit protocol for managing distributed transactions.

## Distributed Transactions
Distributed transactions are a method of coordinating stateful changes between multiple services. All changes must either succeed or fail together. This property is known as atomicity.

## Two-Phase Commit (2PC)
The two-phase commit protocol is a type of atomic commitment protocol (ACP). It involves two phases:

1. **Prepare phase**: The transaction manager (TM) asks all the participants whether they are ready to commit the transaction. They reply with a vote to either commit (Yes) or abort (No).

2. **Commit phase**: If all participants vote to commit, the TM sends a commit message to all participants. If any participant votes to abort, the TM sends an abort message to all participants.

## Trade-offs
While 2PC ensures atomicity and consistency, it does have a few trade-offs:

1. **Blocking**: If the TM crashes after some participants have responded "Yes" but before all have responded, those that have responded "Yes" will block until the TM recovers. This can lead to system-wide latency and can impact the overall performance of the service.

2. **Performance**: Waiting for all participants to respond can be slow, especially when there are a large number of participants or network latency is high. In such cases, the time taken to complete a transaction can increase substantially, potentially impacting the throughput of the service.
There are some optimizations that have been implemented to improve the performance of the 2PC protocol. These include techniques such as early abort, where the TM can decide to abort a transaction as soon as a single participant votes "No".

3. **Single Point of Failure**: The TM in 2PC acts as a single point of failure. If it crashes or becomes unresponsive, all ongoing transactions will be halted, and system operation will be disrupted.

## Handling Failures
The TM is designed to handle various failure scenarios. If any participant votes to abort, the TM will send an abort message to all participants. This has been tested using mock participants where the ability to commit or prepare a transaction can be toggled, and the system has been validated to respond correctly in all situations. Furthermore, if the TM crashes, it can recover and resume the transaction.

## Technical Debt and Future Work
This is a basic implementation of 2PC and does not include certain optimizations and fault-tolerant measures.

1. **Further Optimization**: Read-only participants can be excluded from the commit phase. This would improve performance by removing unnecessary participants. 

2. **Fault Tolerance**: Currently, the TM is a single point of failure. Implementing measures such as replication of the TM or using a distributed consensus algorithm could help improve the fault tolerance of the system.

3. **Three-Phase Commit Protocol (3PC)**: A more advanced protocol, 3PC adds a pre-commit phase to prevent blocking. This could be considered for future implementation to improve the robustness and performance of the service.

4. **Performance Improvements**: Future work could include measures to improve performance, such as more efficient message passing, parallel execution of transactions where possible, or use of non-blocking consensus algorithms.

5. **Security**: As a future enhancement, security measures could be implemented to ensure the authenticity and integrity of transactions. This could include the use of digital signatures, secure communication channels, and access control mechanisms.

## Running the Tests
You can run the unit tests with ```go test ./....``` Tests include checking the basic functionality of the two-phase commit protocol, as well as handling failure scenarios, such as a participant's inability to prepare or commit a transaction.