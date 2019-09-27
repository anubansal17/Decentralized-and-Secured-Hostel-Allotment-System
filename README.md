# Decentralized-and-Secured-Hostel-Allotment-System-for-NITTTR-Chandigarh (MHRD, Govt. of India)
## Highlights:
An Enterprise Blockchain based Project - 
- *Hyperledger Fabric* was used to provide the *blockchain* back-end to the website and a distributed file system, *Interplanetary File System (IPFS)* was used to provide the utmost security to the data involved and *Golang* was used for implementation of Smart Contracts
- Server hardening was ensured by hosting the whole project on *Apache Framework*

# PROJECT
In *Hyperledger*, transactions are submitted via an interface to the ordering service. This service collects transactions based on the consensus algorithm and configuration policy, which may define a time limit or specify the number of transactions allowed. Most of the time, for efficiency reasons, instead of outputting individual transactions, the ordering service will group multiple transactions into single block. In this case, the ordering service must impose and convey a deterministic ordering of the transactions within each block.

## Advantages of using a Hyperledger based application:<br>
● Private *Blockchain*, hence student records are not in public domain<br>
● Access levels can be customized as per requirement. Students won’t be able to allocate rooms to themselves<br>
● Beneficiary will only see the details of the diploma once it’s been issued<br>
● As *Hyperledger* is not coin (“token”) based *Blockchain*, the environment is less complex to develop<br>
● Unlike *Bitcoin or Ethereum *Blockchain* Hyperledger* does not require transferring a virtual currency to publish a transaction<br>
● You can query *Hyperledger Blockchain* to extract details of students to whom hostel has been alloted<br>

## How it Works -
The application has three main components-
1) Frontend
2) Backend :-<br>
a) *Interplanetary File System (“IPFS”)*<br>
b) *Hyperledger Fabric*<br>

One of the disadvantages of storing information on a *Blockchain* based application is that you cannot store files, so we overcame this issue by integrating our system with *IPFS*. *IPFS* is a peer-to-peer method of storing and sharing media in a distributed file system. It is based on Bitcoin *Blockchain* protocol and stores unalterable data, removes duplicates across the network.<br>
Student details such as Student ID, Student Name, Course Name, and Year of Graduation are added on the web application. The web application then generates an excel file containing all the details of student. The .xlsx file is uploaded to *IPFS* node and the hash-pointer (“*IPFS* link”) is saved on our *Blockchain*. The JSON file (containing minimal details) is imported by *Hyperledger* to create an asset and an Invoke Event is displayed on the *Blockchain* terminal.<br>
Furthermore, as the system is hosted on *Hyperledger Blockchain* we can efficiently manage the access levels. The beneficiary i.e. the institution which saves the details of the student as a proof of credentials will have a clear view of the history of the candidate. The .xlsx file will contain immutable record of the allotment done by the hostel authority, date and time of allotment and other details.<br>

![Image of Block Diagram](https://github.com/anubansal17/Decentralized-and-Secured-Hostel-Allotment-System/blob/master/images/BlockDiagram.PNG)

## Application Architecture
Now that we have an overview of how the application works let us take a closer look at implementation of each module.<br>
### InterPlanetary File System (IPFS):<br>
*IPFS* is a peer-to-peer distributed file system that seeks to connect all computing devices with the same system of files. In few ways *IPFS* is similar to World Wide Web, but *IPFS* could be seen as a single BitTorrent swarm exchanging objects within one Git repository. *IPFS* took advantage of *Bitcoin Blockchain* protocol and network infrastructure to store, unalterable data, remove duplicate files across the network and obtain address information for accessing storage nodes. *IPFS* has no single point of failure and nodes do not need to trust each other except for every node they are connected to. Distributed Content Delivery saves bandwidth and prevents DDos attacks, which HTTP struggles with. In-order to store information on *IPFS* network we first have to create a node, this can be achieved by downloading the *IPFS* infrastructure from https://ipfs.io/docs/install/. Once you have downloaded the files and have spun a node on your system we can run an *IPFS* daemon to connect to the Global Object Repository of *IPFS*. *IPFS* returns hash values of the documents uploaded from any node. This hash value is also the location pointer to each document. We have to add the stem of the HTTPS protocol followed by the hash of the file. This link is then stored in the asset class of *Hyperledger* and can be utilized to view the actual file of the participant.
### Hyperledger:<br>
*Hyperledger* is a Linux foundation project to produce an open *Blockchain* platform that is ready for business. It provides implementation of shared ledger, smart contracts, privacy and consensus mechanisms.<br>

## Process Flow :
In our model follows the following structure:<br>
### Participants:<br>
a) **Warden** -The warden of each hostel in the university network.<br>
b) **Head of Department (HOD)** - Head of the department in the university where student is currently enrolled.<br>
c) **Clerk** - Clerk of the each hostel issuing fee receipt to the students on successful payment of hostel fees.<br>
d) **Student** - Student of the university that wants hostel to be allocated to him/her.<br>
### Asset:
Person - Holds the person(Student, Warden, Clerk, HOD) details.
### Transactions:
a) **Invoke** — Executed when any person in the university registers himself/herself in the university network and when any change is made to the details.<br>
b) **Query** — Executed when any person requests to see the details to check the current status of the allocation process.<br>

## A quick walkthrough of the project:
1. First of all, participants will register themselves in the web app to become a part of the network. In the backend, invoke function will be called to make a new entry in the *Blockchain* Ledger (Refer to the code given in SmartContract.go). There would be different predefined formats for registration IDs of authority members and students.<br>
2. If a student registers himself/herself in the network, the request will be sent to the warden of the respective hostel to check the hostel availability for the student.<br>
3. After the verification done by the warden, the request will be sent to the head of the department (HOD) to verify the student enrollment in the department.<br>
4. When the HOD verifies the request, further request will be sent to the clerk of the respective hostel who will issue a fee receipt to the student after successful payment of the hostel fees.<br>
5. After the verification done by the clerk, the request will be sent to the warden of the respective hostel to allocate the available room to the student.<br>
6. Any change in the participant's form will lead to a separate transaction in *Blockchain* Ledger recorded along with the time stamp. In the backend, invoke function will be called if any change is made in the file (Refer to the code given in SmartContract.go). This leads to the immutability of the *Blockchain* Ledger.<br>
7. If the authority member wants to check the history of a student, he/she can query through the *Blockchain* using the student's registration number. In the backend, getHistory function will be called to fetch the full history of a student. (Refer to the code given in SmartContract.go)<br>
8. If the warden wants to check the list of all the residents of the respective hostel, he/she can query through the *Blockchain*. In the backend, getall function will be called to fetch the list of all the students residing in the respective hostel (Refer to the code given in SmartContract.go).<br>
9. To ensure security of the data involved in the model, *IPFS* is used. *IPFS* will store all the files containing details of the participants in a distributed manner. Stored data is safe even if the main server gets hacked. The file will come to the server's existence only if any query is made through *Blockchain* to check the participant's details. For instance, If there's a need to check the current status of a student's hostel allotment process. As soon as, an exit is made from the student details being shown, the file will get deleted from the server. In the backend, to implement this functionality a script has been written in *Golang*.<br>
10. The connectivity of *IPFS* with the *Blockchain* provides the utmost security to the whole project model.<br>

### System Requirements -
● This application has been made on Ubuntu 16.04 using *Hyperledger Fabric framework* and *GoLang*.<br>
● *Hyperledger Fabric* is a platform for distributed ledger solutions underpinned by a modular architecture delivering high degrees of
confidentiality, resiliency, flexibility and scalability. It is designed to support pluggable implementations of different components and accommodate the complexity and intricacies that exist across the economic ecosystem.<br>

#### Technology Stack Used -
● *Hyperledger Fabric*<br>
● *Golang*<br>
● *IPFS*<br>
● *PHP*<br>
● *Apache Server*<br>
● *Mongodb*<br>

### Conclusion:
During the recent few months *Blockchain* based enterprise solution has been a topic of prime interest. A hostel allotment system based on *Blockchain* is one of it's key use case. Using *Blockchain*, it is now possible for us to store our data in a secure way.
A similar infrastructure could be designed which could handle the transfer of medical documents. Prescriptions can be created by doctors; each prescription can be added as an asset in the *Hyperledger* application and file can be uploaded on IPFS. The patient can then show the prescription to any chemist and the chemist can verify the authenticity of the prescription with the history of the doctor who issued it. Such a system can further evolve in an established Healthcare Records Management System.
