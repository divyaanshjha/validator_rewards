# Validator Rewards API 

##  A Brief Overview of Validator Rewards

**Validator rewards** are incentives provided to Ethereum validators for their role in maintaining the Ethereum 2.0 network. Validators participate in activities such as proposing and attesting blocks, synchronizing committees, and ensuring the overall security of the blockchain.

Validator rewards are made up of different components:

The following are the key components of validator rewards, along with a brief explanation:

- **Attestation Rewards**:

  - **Head Reward**: For attesting to the correct head of the chain.
  - **Source Reward**: For voting on the correct source checkpoint.
  - **Target Reward**: For voting on the correct target checkpoint.

- **Proposal Rewards**:

  - Earned by validators who propose new blocks on the Ethereum network.

- **Sync Committee Rewards**:

  - Rewards for participating in sync committee duties to help light clients stay updated.

- **Transaction Fee Rewards**:

  - Includes tips and priority fees from transactions included in blocks proposed by validators.

### **Penalty Components**

Penalties are applied when validators fail to perform their duties or behave maliciously:

- **Attestation Source/Target Penalty**: Applied when the validator fails to vote on the correct source or target checkpoint.
- **Finality Delay Penalty**: Applied when the network experiences delays in finalizing blocks.
- **Sync Committee Penalty**: Applied if the validator fails to perform sync committee duties.
- **Slashing Penalty**: Applied in cases of double signing or other malicious activities.

You can read about validator reward components in detail at https\://ethos.dev/beacon-chain#staking-rewards-and-penalties

## API Description

This API provides an endpoint to fetch detailed reward data for a specific Ethereum 2.0 validator. The API dynamically calculates rewards and penalties based on data fetched from the Ethereum 2.0 Beacon Chain. The endpoint accepts a validator public key or index as input and returns the reward details.

### **API Endpoint**
```bash
GET /validator-rewards/:validatorPublicKeyorIndex
```
Fetch the rewards and penalties for a specific validator.

#### **Parameters**

- `validatorPubKey` (string): The public key of the validator to fetch rewards for.

#### **Response**

- JSON object containing reward components and penalties.

**Example Request:**

```bash
GET http://localhost:8080/validator-rewards/0x80000d79dbfde36d1dcb492e74c4b55adbe7ffc4c4396dea0095f76f002092ee62e218dd0a31bccc04b00ed0ed8aefbe
```

**Example Response:**

```json
{
  "validatorAddress": "0xabc...",
  "totalRewards": "15.6 ETH",
  "attestationRewards": "10.4 ETH",
  "proposalRewards": "3.5 ETH",
  "syncCommitteeRewards": "1.7 ETH",
  "transactionFeeRewards": "0.8 ETH",
  "penalties": "-0.3 ETH",
}
```

## How to Run 

### **Prerequisites**

- Install **Golang** (v1.18 or later) on your device.
- Install a tool like **Postman** or **cURL** for API testing.

### **Steps to Run**

1. **Clone the Repository**

   ```bash
   git clone https://github.com/divyaanshjha/validator_rewards
   cd validator_rewards
   ```

2. **Install Dependencies**
   Ensure you have all required dependencies installed. The repository uses the **Gin** framework for the API. Run:

   ```bash
   go mod tidy
   ```

3. **Run the API**
   Start the server by running:

   ```bash
   go run main.go
   ```

   The API server will start on `http://localhost:8080` by default.

4. **Test the API**
   Use tools like Postman or your browser to send a GET request:

   ```bash
   GET http://localhost:8080/validator-rewards/{validatorPubKey}
   ```

###
