FORMAT: 1A
HOST: http://api.coinding.com

# Coinding API
Coinding API is a *bitcoin service* designed to allow developers around the world to easily integrate the power of bitcoin into their games.

We allow multiple interactions with bitcoin through one API including:

* Send tips to game developers or fellow players
* Sending and requesting bitcoin by email or bitcoin address
* Access to raw bitcoin network data
* Fetch random bitcoin information

Some other features we are working on and will be available soon:
- Handle In-App purchases without fees
- Interact with streams and live e-sport matches

We want to help you make great games that work with bitcoin, so please send requests and suggestions to api@coinding.com.

# Group Developer
These services allow you to manage a game developer account.

## List Developers [/developer]
### Get all Developers [GET]
Returns a list with the name of registered developers.
+ Response 200 (application/json)

        [{name: Capybara Games},{name: Eric Chahi},{name: Ralph Baer},{name: Ron Gilbert},{name: Steve Feak},{name: Team Meat},{name: Yuji Horii}]

## Register Developer [/developer/register{?name,email,pass}]
### Create new Developer [POST]
Creates a new developer account

+ Parameters
    + name (required, string, `John Dev`) ... Name of the game developer.
    + email (required, string, `some@email.com`) ... Contact email.
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Account password. It is stored encripted but we recommend that you hash it, to avoid transfering it as plain text.

+ Response 200

## Get Balance [/developer/balance{?email,pass}]
### Retrieve Developer Balance [GET]
Authenticated resource which returns the current balance of the developer in Bitcoins

+ Parameters
    + email (required, string, `some@email.com`) ... Developer email
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request

+ Response 200 (application/json)

        {name: Eric Chahi, email: eric@chahi.com, address: 13jsKDJz2MmvxjuVqiS5KmTA8vCwDkPUj9, balance: 1.75}

## Withdraw funds [/developer/withdraw{?email,pass,amount,to}]
### Withdraw Developer Funds [GET]
Authenticated resource which lets you send money to an email or bitcoin address.

+ Parameters
    + email (required, string, `some@email.com`) ... Developer email
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request
    + amount (required, float, `0.4`) ... A string amount that will be converted to BTC, such as 1 or 1.234567. Also must be >= 0.01 or it will shown an error
    + to (required, string, `13jsKDJz2MmvxjuVqiS5KmTA8vCwDkPUj9`) ... An email address or a bitcoin address

+ Response 200

# Group Game
Basic game management

## List Games [/game]
### Get all Games [GET]
Returns a list with the name of registered games.
+ Response 200 (application/json)

        [{name: Flappy bits, dev: Anonymous LLC., url: http://flappybits.com},{name: Dancing Disco Ball, dev: Paul, url: http://somegame.com},{name: Super Monkey Ball, dev: John Romero, url: http://supermonkeyball.com}]

## Register Game [/developer/register{?name,dev,url}]
### Create new Game [POST]
Adds a new game to the network

+ Parameters
    + name (required, string, `Flappy bits`) ... Name of the game developer.
    + dev (required, string, `Anonymous LLC.`) ... Developer name.
    + url (required, string, `http://flappybits.com`) ... URL where the game can be located

+ Response 200

# Group Player
These services allow you to manage a gamer account.

## Register Player [/player/register{?name,pass}]
### Create new Player [POST]
Creates a new player account

+ Parameters
    + name (required, string, `Gu Li`) ... Name of the player
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Account password. It is stored encripted but we recommend you to hash it, to avoid transfering it as plain text

+ Response 200

## Get Balance [/player/balance{?name,pass}]
### Retrieve Player Balance [GET]
Authenticated resource which returns the current balance of the player in Bitcoins

+ Parameters
    + name (required, string, `Monto Cavasini`) ... Player name
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request

+ Response 200 (application/json)

        {name: Monto Cavasini, address: 13jsKDJz2MmvxjuVqiS5KmTA8vCwDkPUj9, balance: 4.20}

## Withdraw funds [/player/withdraw{?name,pass,amount,to}]
### Withdraw Developer Funds [GET]
Authenticated resource which lets you send money to an email or bitcoin address.

+ Parameters
    + name (required, string, `Monto Cavasini`) ... Player name
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request
    + amount (required, float, `0.04`) ... A string amount that will be converted to BTC, such as 1 or 1.234567. Also must be >= 0.01 or it will shown an error
    + to (required, string, `13jsKDJz2MmvxjuVqiS5KmTA8vCwDkPUj9`) ... An email address or a bitcoin address

+ Response 200

# Group Tip
These services allow you to send bitcoins between the players and developers

## Tip [/tip{?name,pass,amount,to}]
### Player Tip Player [GET]
Authenticated resource which allows a player send money to another player

+ Parameters
    + name (required, string, `Lola Garcia`) ... Player name
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request
    + amount (required, float, `0.04`) ... A string amount that will be converted to BTC, such as 1 or 0.00001. It can be any amount
    + to (required, string, `Lucky John`) ... Name of the player that will receive the tip

+ Response 200

## Tip [/tip/dev{?name,pass,amount,to}]
### Player Tip Developer [GET]
Authenticated resource which allows a player send money to a game developer

+ Parameters
    + name (required, string, `Lola Garcia`) ... Player name
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request
    + amount (required, float, `0.04`) ... A string amount that will be converted to BTC, such as 1 or 0.00001. It can be any amount
    + to (required, string, `Lucky John`) ... Name of the developer that will receive the tip

+ Response 200

## Tip [/tip/player{?email,pass,amount,to}]
### Developer Tip Player [GET]
Authenticated resource which allows a developer send money to a player

+ Parameters
    + email (required, string, `game@dev.com`) ... Developer email
    + pass (required, string, `30565d8712bad1d13fb10133c5f0cd60`) ... Used to authorize the request
    + amount (required, float, `0.04`) ... A string amount that will be converted to BTC, such as 1 or 0.00001. It can be any amount
    + to (required, string, `Lucky John`) ... Name of the player that will receive the tip

+ Response 200

# Group Bitcoin
These services allow simple access to blockchain data

# Group Bitcoin/Address
The *Address* Object contains basic balance details for a Bitcoin address.

## Random Address [/bitcoin/address/random]
### Retrieve random Address [GET]
Returns basic balance details of a random Bitcoin addresses.
+ Response 200 (application/json)

        [{"address":"17x23dNjXJLzGMev6R63uyRhMWP1VHawKc","total":{"balance":5000000000,"received":5000000000,"sent":0},"confirmed":{"balance":5000000000,"received":5000000000,"sent":0}}]

## Get Address [/bitcoin/address/{hash}]
### Retrieve an Address [GET]
Returns basic balance details for a single Bitcoin addresses.

+ Parameters
    + hash (required, string, `1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND`) ... A Bitcoin address

+ Response 200 (application/json)

        [{"address":"17x23dNjXJLzGMev6R63uyRhMWP1VHawKc","total":{"balance":5000000000,"received":5000000000,"sent":0},"confirmed":{"balance":5000000000,"received":5000000000,"sent":0}}]


## Address Transactions [/bitcoin/address/{hash}/transactions]
### Retrieve Address Transactions [GET]
Returns a set of transactions for a Bitcoin address.

+ Parameters
    + hash (required, string, `1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND`) ... A Bitcoin address

+ Response 200 (application/json)

        [{"hash":"f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331","block_hash":"00000000bc919cfb64f62de736d55cf79e3d535b474ace256b4fbb56073f64db","block_height":30,"block_time":"2009-01-10T15:42:02Z","chain_received_at":null,"confirmations":339920,"lock_time":0,"inputs":[{"coinbase":"04ffff001d0121","value":5000000000,"sequence":4294967295}],"outputs":[{"transaction_hash":"f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331","output_index":0,"value":5000000000,"addresses":["17x23dNjXJLzGMev6R63uyRhMWP1VHawKc"],"script":"042cf59fafd089a348c5834283573608e89a305c60a034604c7d22dde50998f1b9bb74681986ca1884a6b1df8ce7f1b79a2277057de855a634626e7a5851c1e716 OP_CHECKSIG","script_hex":"41042cf59fafd089a348c5834283573608e89a305c60a034604c7d22dde50998f1b9bb74681986ca1884a6b1df8ce7f1b79a2277057de855a634626e7a5851c1e716ac","script_type":"pubkey","required_signatures":1,"spent":false}],"fees":0,"amount":5000000000}]

## Address Unspents [/bitcoin/address/{hash}/unspents]
### Retrieve Address Unspents [GET]
Returns a collection of unspent outputs for a Bitcoin address.

+ Parameters
    + hash (required, string, `17x23dNjXJLzGMev6R63uyRhMWP1VHawKc`) ... A Bitcoin address

+ Response 200 (application/json)

        [{"hash":"f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331","block_hash":"00000000bc919cfb64f62de736d55cf79e3d535b474ace256b4fbb56073f64db","block_height":30,"block_time":"2009-01-10T15:42:02Z","chain_received_at":null,"confirmations":339920,"lock_time":0,"inputs":[{"coinbase":"04ffff001d0121","value":5000000000,"sequence":4294967295}],"outputs":[{"transaction_hash":"f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331","output_index":0,"value":5000000000,"addresses":["17x23dNjXJLzGMev6R63uyRhMWP1VHawKc"],"script":"042cf59fafd089a348c5834283573608e89a305c60a034604c7d22dde50998f1b9bb74681986ca1884a6b1df8ce7f1b79a2277057de855a634626e7a5851c1e716 OP_CHECKSIG","script_hex":"41042cf59fafd089a348c5834283573608e89a305c60a034604c7d22dde50998f1b9bb74681986ca1884a6b1df8ce7f1b79a2277057de855a634626e7a5851c1e716ac","script_type":"pubkey","required_signatures":1,"spent":false}],"fees":0,"amount":5000000000}]

# Group Bitcoin/Transaction
The *Transaction* Object contains details about a Bitcoin transaction, including inputs and outputs.

## Last Transaction [/bitcoin/transaction]
### Retrieve last transactions [GET]
Returns an array with the latest transaction

+ Response 200 (application/json)

        [{"address":"17x23dNjXJLzGMev6R63uyRhMWP1VHawKc","total":{"balance":5000000000,"received":5000000000,"sent":0},"confirmed":{"balance":5000000000,"received":5000000000,"sent":0}},{"address":"17x23dNjXJLzGMev6R63uyRhMWP1VHawKc","total":{"balance":5000000000,"received":5000000000,"sent":0},"confirmed":{"balance":5000000000,"received":5000000000,"sent":0}}]
        
## Random Transaction [/bitcoin/transaction/random]
### Retrieve Random Transaction [GET]
Returns details about a Bitcoin transaction, including inputs and outputs.

+ Response 200 (application/json)

        {"hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","block_hash": "00000000000000001344545f30030f0e43477ca9d42d958a49bb8f4ebd3bf0ad","block_height": 340124,"block_time": "2015-01-23T02:35:11Z","chain_received_at": "2015-01-23T02:21:28.240Z","confirmations": 4,"lock_time": 0,"inputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_hash": "ee2aa75c909c5e99c0e762fbc3029109de890930b907c4b25e1f37155bcd6413","output_index": 0,"value": 103170000,"addresses": ["1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND"],"script_signature": "304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c74101 037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","script_signature_hex": "47304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c7410121037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","sequence": 4294967295}],"outputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 0,"value": 83800000,"addresses": ["1HU3fDRwn66pkP6ZhPaT34wkYyQrE8doDr"],"script": "OP_DUP OP_HASH160 b49f7688727d255952ea7d6ba2f37190702e869a OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914b49f7688727d255952ea7d6ba2f37190702e869a88ac","script_type": "pubkeyhash","required_signatures": 1,"spent": false},{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 1,"value": 19360000,"addresses": ["1MUQcwCUzEzfGefnj7wXUZXZkEHoJd9g1J"],"script": "OP_DUP OP_HASH160 e0917698271778c3ee832fe69b8aadb5c19ee275 OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914e0917698271778c3ee832fe69b8aadb5c19ee27588ac","script_type": "pubkeyhash","required_signatures": 1,"spent": true}],"fees": 10000,"amount": 103160000}

## Get Transaction [/bitcoin/transaction/{hash}]
### Retrieve a Transaction [GET]
Returns details about a Bitcoin transaction, including inputs and outputs.

+ Parameters
    + hash (required, string, `0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b`) ... A transaction hash

+ Response 200 (application/json)

        {"hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","block_hash": "00000000000000001344545f30030f0e43477ca9d42d958a49bb8f4ebd3bf0ad","block_height": 340124,"block_time": "2015-01-23T02:35:11Z","chain_received_at": "2015-01-23T02:21:28.240Z","confirmations": 4,"lock_time": 0,"inputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_hash": "ee2aa75c909c5e99c0e762fbc3029109de890930b907c4b25e1f37155bcd6413","output_index": 0,"value": 103170000,"addresses": ["1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND"],"script_signature": "304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c74101 037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","script_signature_hex": "47304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c7410121037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","sequence": 4294967295}],"outputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 0,"value": 83800000,"addresses": ["1HU3fDRwn66pkP6ZhPaT34wkYyQrE8doDr"],"script": "OP_DUP OP_HASH160 b49f7688727d255952ea7d6ba2f37190702e869a OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914b49f7688727d255952ea7d6ba2f37190702e869a88ac","script_type": "pubkeyhash","required_signatures": 1,"spent": false},{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 1,"value": 19360000,"addresses": ["1MUQcwCUzEzfGefnj7wXUZXZkEHoJd9g1J"],"script": "OP_DUP OP_HASH160 e0917698271778c3ee832fe69b8aadb5c19ee275 OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914e0917698271778c3ee832fe69b8aadb5c19ee27588ac","script_type": "pubkeyhash","required_signatures": 1,"spent": true}],"fees": 10000,"amount": 103160000}

## Transaction Confidence [/bitcoin/transaction/{hash}/confidence]
### Estimated Transaction Confidence [GET]
Returns network propagation level and double spend information for a given Bitcoin transaction hash.

+ Parameters
    + hash (required, string, `0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b`) ... A transaction hash

+ Response 200 (application/json)

        {"hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","block_hash": "00000000000000001344545f30030f0e43477ca9d42d958a49bb8f4ebd3bf0ad","block_height": 340124,"block_time": "2015-01-23T02:35:11Z","chain_received_at": "2015-01-23T02:21:28.240Z","confirmations": 4,"lock_time": 0,"inputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_hash": "ee2aa75c909c5e99c0e762fbc3029109de890930b907c4b25e1f37155bcd6413","output_index": 0,"value": 103170000,"addresses": ["1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND"],"script_signature": "304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c74101 037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","script_signature_hex": "47304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c7410121037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","sequence": 4294967295}],"outputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 0,"value": 83800000,"addresses": ["1HU3fDRwn66pkP6ZhPaT34wkYyQrE8doDr"],"script": "OP_DUP OP_HASH160 b49f7688727d255952ea7d6ba2f37190702e869a OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914b49f7688727d255952ea7d6ba2f37190702e869a88ac","script_type": "pubkeyhash","required_signatures": 1,"spent": false},{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 1,"value": 19360000,"addresses": ["1MUQcwCUzEzfGefnj7wXUZXZkEHoJd9g1J"],"script": "OP_DUP OP_HASH160 e0917698271778c3ee832fe69b8aadb5c19ee275 OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914e0917698271778c3ee832fe69b8aadb5c19ee27588ac","script_type": "pubkeyhash","required_signatures": 1,"spent": true}],"fees": 10000,"amount": 103160000}

## Transaction Hex [/bitcoin/transaction/{hash}/hex]
### Raw Transaction Hex [GET]
Returns the raw transaction hex data for a given Bitcoin transaction hash.

+ Parameters
    + hash (required, string, `0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b`) ... A transaction hash

+ Response 200 (application/json)

        {"hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","block_hash": "00000000000000001344545f30030f0e43477ca9d42d958a49bb8f4ebd3bf0ad","block_height": 340124,"block_time": "2015-01-23T02:35:11Z","chain_received_at": "2015-01-23T02:21:28.240Z","confirmations": 4,"lock_time": 0,"inputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_hash": "ee2aa75c909c5e99c0e762fbc3029109de890930b907c4b25e1f37155bcd6413","output_index": 0,"value": 103170000,"addresses": ["1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND"],"script_signature": "304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c74101 037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","script_signature_hex": "47304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c7410121037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","sequence": 4294967295}],"outputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 0,"value": 83800000,"addresses": ["1HU3fDRwn66pkP6ZhPaT34wkYyQrE8doDr"],"script": "OP_DUP OP_HASH160 b49f7688727d255952ea7d6ba2f37190702e869a OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914b49f7688727d255952ea7d6ba2f37190702e869a88ac","script_type": "pubkeyhash","required_signatures": 1,"spent": false},{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 1,"value": 19360000,"addresses": ["1MUQcwCUzEzfGefnj7wXUZXZkEHoJd9g1J"],"script": "OP_DUP OP_HASH160 e0917698271778c3ee832fe69b8aadb5c19ee275 OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914e0917698271778c3ee832fe69b8aadb5c19ee27588ac","script_type": "pubkeyhash","required_signatures": 1,"spent": true}],"fees": 10000,"amount": 103160000}


# Group Bitcoin/Block
The *Block Object* contains details about a Bitcoin block, including all transaction hashes.

## Last Block [/bitcoin/block]
### Retrieve last block [GET]
Fetches the last mined block

+ Response 200 (application/json)

        [{"hash":"f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331","block_hash":"00000000bc919cfb64f62de736d55cf79e3d535b474ace256b4fbb56073f64db","block_height":30,"block_time":"2009-01-10T15:42:02Z","chain_received_at":null,"confirmations":339922,"lock_time":0,"inputs":[{"coinbase":"04ffff001d0121","value":5000000000,"sequence":4294967295}],"outputs":[{"transaction_hash":"f5e26c8b82401c585235c572ba8265f16f7d9304ed8e31c198eab571754f5331","output_index":0,"value":5000000000,"addresses":["17x23dNjXJLzGMev6R63uyRhMWP1VHawKc"],"script":"042cf59fafd089a348c5834283573608e89a305c60a034604c7d22dde50998f1b9bb74681986ca1884a6b1df8ce7f1b79a2277057de855a634626e7a5851c1e716 OP_CHECKSIG","script_hex":"41042cf59fafd089a348c5834283573608e89a305c60a034604c7d22dde50998f1b9bb74681986ca1884a6b1df8ce7f1b79a2277057de855a634626e7a5851c1e716ac","script_type":"pubkey","required_signatures":1,"spent":false}],"fees":0,"amount":5000000000}]

## Random Block [/bitcoin/block/random]
### Retrieve a random Block [GET]
Returns details about a random Bitcoin block, including all transaction hashes.

+ Response 200 (application/json)

        {"hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","block_hash": "00000000000000001344545f30030f0e43477ca9d42d958a49bb8f4ebd3bf0ad","block_height": 340124,"block_time": "2015-01-23T02:35:11Z","chain_received_at": "2015-01-23T02:21:28.240Z","confirmations": 4,"lock_time": 0,"inputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_hash": "ee2aa75c909c5e99c0e762fbc3029109de890930b907c4b25e1f37155bcd6413","output_index": 0,"value": 103170000,"addresses": ["1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND"],"script_signature": "304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c74101 037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","script_signature_hex": "47304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c7410121037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","sequence": 4294967295}],"outputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 0,"value": 83800000,"addresses": ["1HU3fDRwn66pkP6ZhPaT34wkYyQrE8doDr"],"script": "OP_DUP OP_HASH160 b49f7688727d255952ea7d6ba2f37190702e869a OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914b49f7688727d255952ea7d6ba2f37190702e869a88ac","script_type": "pubkeyhash","required_signatures": 1,"spent": false},{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 1,"value": 19360000,"addresses": ["1MUQcwCUzEzfGefnj7wXUZXZkEHoJd9g1J"],"script": "OP_DUP OP_HASH160 e0917698271778c3ee832fe69b8aadb5c19ee275 OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914e0917698271778c3ee832fe69b8aadb5c19ee27588ac","script_type": "pubkeyhash","required_signatures": 1,"spent": true}],"fees": 10000,"amount": 103160000}

        
## Get Block [/bitcoin/block/{hash}]
### Retrieve a Block [GET]
Returns details about a Bitcoin block, including all transaction hashes.

+ Parameters
    + hash (required, string, `00000000000000009cc33fe219537756a68ee5433d593034b6dc200b34aa35fa`) ... A block hash

+ Response 200 (application/json)

        {"hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","block_hash": "00000000000000001344545f30030f0e43477ca9d42d958a49bb8f4ebd3bf0ad","block_height": 340124,"block_time": "2015-01-23T02:35:11Z","chain_received_at": "2015-01-23T02:21:28.240Z","confirmations": 4,"lock_time": 0,"inputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_hash": "ee2aa75c909c5e99c0e762fbc3029109de890930b907c4b25e1f37155bcd6413","output_index": 0,"value": 103170000,"addresses": ["1Msk11Bt3jvogNYmEPSvNyCf9dcp2eRKND"],"script_signature": "304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c74101 037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","script_signature_hex": "47304402203c0917e0801f67d3d5e8f3702ae4daed0c744b6b9ca6f17268bcf123739eaddc02202c3c664d46b4963bbab26370e1a831a0b2a904513e8cd12ae8bd4d2f4789c7410121037cf5415cff2ae015bc00a92f854fdb42f27f96c3169bf3d310432938daff2e0e","sequence": 4294967295}],"outputs": [{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 0,"value": 83800000,"addresses": ["1HU3fDRwn66pkP6ZhPaT34wkYyQrE8doDr"],"script": "OP_DUP OP_HASH160 b49f7688727d255952ea7d6ba2f37190702e869a OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914b49f7688727d255952ea7d6ba2f37190702e869a88ac","script_type": "pubkeyhash","required_signatures": 1,"spent": false},{"transaction_hash": "0463f5eb9def6b2568fc699f6faa2c4d6080ab10986a2bbecf2860e8c369bf2b","output_index": 1,"value": 19360000,"addresses": ["1MUQcwCUzEzfGefnj7wXUZXZkEHoJd9g1J"],"script": "OP_DUP OP_HASH160 e0917698271778c3ee832fe69b8aadb5c19ee275 OP_EQUALVERIFY OP_CHECKSIG","script_hex": "76a914e0917698271778c3ee832fe69b8aadb5c19ee27588ac","script_type": "pubkeyhash","required_signatures": 1,"spent": true}],"fees": 10000,"amount": 103160000}

        
        
        