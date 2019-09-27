# Decentralized-and-Secured-Hostel-Allotment-System-for-NITTTR-Chandigarh
Successfully made Hostel Allotment System for NITTTR, Chandigarh (Ministry of HRD, Govt. of India), an Enterprise Blockchain based Project.
- Hyperledger Fabric was used to provide the blockchain back-end to the website and a distributed file system, Interplanetary File System (IPFS) was used to provide the utmost security to the data involved and Golang was used for implementation of Smart Contracts
- Server hardening was ensured by hosting the whole project on Apache Framework

# PROJECT - HOSTEL ALLOTMENT SYSTEM USING HYPERLEDGER FABRIC
In Hyperledger, transactions are submitted via an interface to the ordering service. This service collects transactions based on the consensus algorithm and configuration policy, which may define a time limit or specify the number of transactions allowed. Most of the time, for efficiency reasons, instead of outputting individual transactions, the ordering service will group multiple transactions into single block. In this case, the ordering service must impose and convey a deterministic ordering of the transactions within each block.

Advantages of using a Hyperledger based application:<br>
● Private Blockchain hence student records are not in public domain<br>
● Access levels can be customized as per requirement. Students won’t be able to allocate rooms to themselves<br>
● Beneficiary will only see the details of the diploma once it’s been issued<br>
● As Hyperledger is not coin (“token”) based blockchain, the environment is less complex to develop<br>
● Unlike Bitcoin or Ethereum blockchain Hyperledger does not require transferring a virtual currency to publish a transaction<br>
● You can query Hyperledger blockchain to extract details of students to whom hostel has been alloted<br>

How it Works -
The application has three main components-
1) Front End
2) Backend :-
a) Interplanetary File System (“IPFS”)
b) Hyperledger Fabric

One of the disadvantages of storing information on a blockchain based application is that you cannot store files, so we overcame this issue by integrating our system with IPFS. IPFS is a peer-to-peer method of storing and sharing media in a distributed file system. It is based on Bitcoin blockchain protocol and stores unalterable data, removes duplicates across the network. Student details such as Student ID, Student Name, Course Name, and Year of Graduation are added on the web application. The web application then generates an excel file containing all the details of student. The .xlsx file is uploaded to IPFS node and the hash-pointer (“IPFS link”) is saved on our blockchain. The JSON file (containing minimal details) is imported by Hyperledger to create an asset and an Invoke Event is displayed on the Blockchain Terminal. Furthermore, as the system is hosted on Hyperledger blockchain we can efficiently manage the access levels. The beneficiary i.e. the institution which saves the details of the student as a proof of credentials will have a clear view of the history of the candidate. The .xlsx file will contain immutable record of the allotment done by the hostel authority, date and time of allotment and other details.

